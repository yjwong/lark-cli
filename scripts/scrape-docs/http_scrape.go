package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

const apiBase = "https://open.larksuite.com/document_portal/v1/document/get_detail"

// httpAPIResponse is the top-level JSON returned by the Lark portal API.
type httpAPIResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Title      string      `json:"title"`
		Content    string      `json:"content"`
		UpdateTime json.Number `json:"updateTime"`
		FullPath   string      `json:"fullPath"`
		Name       string      `json:"name"`
	} `json:"data"`
}

// httpFetchPage fetches a single page via the Lark portal HTTP API and writes the markdown file.
// It returns the updateTime string from the API response.
func httpFetchPage(client *http.Client, page *ScrapedPage, outputDir string) (updateTime string, err error) {
	docPath := page.TargetPath
	if docPath == "" {
		docPath = page.Path
	}
	if docPath == "" {
		return "", fmt.Errorf("no URL path available")
	}

	// Build the API URL
	apiURL := apiBase + "?fullPath=" + url.QueryEscape(docPath)

	req, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		return "", fmt.Errorf("failed to build request: %w", err)
	}
	req.Header.Set("Cookie", "open_locale=en-US")
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; lark-cli-scraper/1.0)")

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected HTTP status %d for %s", resp.StatusCode, docPath)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	var apiResp httpAPIResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return "", fmt.Errorf("failed to parse API response: %w", err)
	}
	if apiResp.Code != 0 {
		return "", fmt.Errorf("API error code %d: %s", apiResp.Code, apiResp.Msg)
	}
	if apiResp.Data.Content == "" {
		return "", fmt.Errorf("empty content returned for %s", docPath)
	}

	// Determine output path
	relPath := deriveOutputPath(page)
	outFile := filepath.Join(outputDir, relPath)

	// Validate output path stays within outputDir (path traversal protection)
	absOutFile, err := filepath.Abs(outFile)
	if err != nil {
		return "", fmt.Errorf("failed to resolve output path: %w", err)
	}
	absOutput, err := filepath.Abs(outputDir)
	if err != nil {
		return "", fmt.Errorf("failed to resolve output dir: %w", err)
	}
	rel, err := filepath.Rel(absOutput, absOutFile)
	if err != nil || strings.HasPrefix(rel, "..") {
		return "", fmt.Errorf("path traversal detected: %s escapes output directory", relPath)
	}

	// Parse content and build markdown with frontmatter
	service, resource := parseServiceResource(docPath)
	parsed := ParseContent(apiResp.Data.Content)
	updateTimeStr := string(apiResp.Data.UpdateTime)
	md := buildMarkdown(page.Name, docPath, service, resource, page.Section, updateTimeStr, parsed)

	// Write file
	if err := os.MkdirAll(filepath.Dir(outFile), 0o755); err != nil {
		return "", fmt.Errorf("mkdir failed: %w", err)
	}
	if err := os.WriteFile(outFile, []byte(md), 0o644); err != nil {
		return "", fmt.Errorf("write failed: %w", err)
	}

	page.OutputFile = relPath
	return updateTimeStr, nil
}

// tickerLimiter is a simple token-bucket-style rate limiter backed by a time.Ticker.
type tickerLimiter struct {
	ticker *time.Ticker
	done   chan struct{}
}

func newTickerLimiter(rps float64) *tickerLimiter {
	if rps <= 0 {
		rps = 1
	}
	interval := time.Duration(float64(time.Second) / rps)
	return &tickerLimiter{
		ticker: time.NewTicker(interval),
		done:   make(chan struct{}),
	}
}

func (l *tickerLimiter) Wait() {
	<-l.ticker.C
}

func (l *tickerLimiter) Stop() {
	l.ticker.Stop()
	close(l.done)
}

// HTTPScrapePages fetches each page in the catalog using plain HTTP (no browser).
// It uses a worker pool of `workers` goroutines and rate-limits to `rps` req/s.
func HTTPScrapePages(catalog *Catalog, outputDir string, workers int, rps float64, sectionFilter string) error {
	progress := loadProgress(outputDir, catalog)

	// Build section filter set
	allowedSections := make(map[string]bool)
	if sectionFilter != "" {
		for _, s := range strings.Split(sectionFilter, ",") {
			allowedSections[strings.TrimSpace(s)] = true
		}
	}

	// Collect pending pages
	var pending []*ScrapedPage
	for i := range progress.Pages {
		p := &progress.Pages[i]
		if p.Status != "pending" && p.Status != "error" {
			continue
		}
		if len(allowedSections) > 0 && !allowedSections[p.Section] {
			continue
		}
		pending = append(pending, p)
	}

	log.Printf("HTTP scraping %d pages (total: %d, done: %d, errors: %d)",
		len(pending), progress.TotalPages, progress.Done, progress.Errors)

	if len(pending) == 0 {
		return nil
	}

	// Shared HTTP client (connection pooling)
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	limiter := newTickerLimiter(rps)
	defer limiter.Stop()

	// Channel feeds pages to workers
	pageCh := make(chan *ScrapedPage, len(pending))
	for _, p := range pending {
		pageCh <- p
	}
	close(pageCh)

	var mu sync.Mutex
	var wg sync.WaitGroup

	if workers < 1 {
		workers = 1
	}

	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for page := range pageCh {
				// Rate limit: one tick per request across all workers
				limiter.Wait()

				remoteUpdateTime, err := httpFetchPage(client, page, outputDir)
				if err != nil {
					// Retry once after a short pause
					time.Sleep(2 * time.Second)
					limiter.Wait()
					remoteUpdateTime, err = httpFetchPage(client, page, outputDir)
				}

				mu.Lock()
				if err != nil {
					page.Status = "error"
					page.Error = err.Error()
					log.Printf("[ERROR] %s: %s", page.Name, err)
				} else {
					page.Status = "done"
					page.Error = ""
					page.ScrapedAt = time.Now().Format(time.RFC3339)
					page.UpdateTime = remoteUpdateTime
					log.Printf("[DONE] %s -> %s", page.Name, page.OutputFile)
				}

				// Recompute counters
				progress.Done = 0
				progress.Errors = 0
				for j := range progress.Pages {
					switch progress.Pages[j].Status {
					case "done":
						progress.Done++
					case "error":
						progress.Errors++
					}
				}
				progress.LastUpdated = time.Now()
				saveProgress(outputDir, progress)
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	return nil
}

// HTTPCheckUpdateTimes checks the updateTime of multiple pages concurrently via HTTP
// and returns the subset that are new or have changed since the last scrape.
func HTTPCheckUpdateTimes(catalog *Catalog, oldPages map[string]*ScrapedPage, workers int, rps float64) ([]*CatalogNode, error) {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	limiter := newTickerLimiter(rps)
	defer limiter.Stop()

	type result struct {
		node       *CatalogNode
		updateTime string
		err        error
	}

	pageCh := make(chan *CatalogNode, len(catalog.Pages))
	for _, p := range catalog.Pages {
		pageCh <- p
	}
	close(pageCh)

	resultCh := make(chan result, len(catalog.Pages))

	if workers < 1 {
		workers = 1
	}

	var wg sync.WaitGroup
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for node := range pageCh {
				limiter.Wait()

				docPath := node.TargetPath
				if docPath == "" {
					docPath = node.Path
				}

				ut, err := httpFetchUpdateTime(client, docPath)
				if err != nil {
					// Retry once
					time.Sleep(2 * time.Second)
					limiter.Wait()
					ut, err = httpFetchUpdateTime(client, docPath)
				}
				resultCh <- result{node: node, updateTime: ut, err: err}
			}
		}()
	}

	// Close resultCh once all workers finish
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	var needsScrape []*CatalogNode
	var checkErrors int

	for r := range resultCh {
		if r.err != nil {
			log.Printf("  [CHECK ERROR] %s: %s (will re-scrape)", r.node.Name, r.err)
			needsScrape = append(needsScrape, r.node)
			checkErrors++
			continue
		}

		key := r.node.TargetPath
		if key == "" {
			key = r.node.Path
		}
		old, exists := oldPages[key]
		if !exists || old.Status != "done" || old.UpdateTime != r.updateTime {
			if exists && old.UpdateTime != r.updateTime {
				log.Printf("  [CHANGED] %s (stored: %s, remote: %s)", r.node.Name, old.UpdateTime, r.updateTime)
			}
			needsScrape = append(needsScrape, r.node)
		}
	}

	log.Printf("HTTP check complete: %d need re-scraping (%d check errors)", len(needsScrape), checkErrors)
	return needsScrape, nil
}

// httpFetchUpdateTime fetches only the updateTime for a given docPath.
func httpFetchUpdateTime(client *http.Client, docPath string) (string, error) {
	apiURL := apiBase + "?fullPath=" + url.QueryEscape(docPath)

	req, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		return "", fmt.Errorf("failed to build request: %w", err)
	}
	req.Header.Set("Cookie", "open_locale=en-US")
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; lark-cli-scraper/1.0)")

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected HTTP status %d for %s", resp.StatusCode, docPath)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	var apiResp httpAPIResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return "", fmt.Errorf("failed to parse API response: %w", err)
	}
	if apiResp.Code != 0 {
		return "", fmt.Errorf("API error code %d: %s", apiResp.Code, apiResp.Msg)
	}

	return string(apiResp.Data.UpdateTime), nil
}
