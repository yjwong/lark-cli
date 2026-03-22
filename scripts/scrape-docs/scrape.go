package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

const extractContentJS = `(function() {
	var el = document.querySelector('.ud__tree');
	if (!el) return JSON.stringify({error: 'no tree element'});
	var fiberKey = Object.keys(el).find(function(k) { return k.startsWith('__reactFiber'); });
	if (!fiberKey) return JSON.stringify({error: 'no react fiber'});
	var fiber = el[fiberKey];
	for (var i = 0; i < 25; i++) {
		if (!fiber.return) break;
		fiber = fiber.return;
		var dd = fiber.memoizedProps && fiber.memoizedProps.documentData;
		if (dd && dd.content) {
			return JSON.stringify({
				name: dd.name || '',
				content: dd.content || '',
				fullPath: dd.fullPath || '',
				updateTime: dd.updateTime || ''
			});
		}
	}
	return JSON.stringify({error: 'no documentData found (searched 25 levels)'});
})()`

const fallbackContentJS = `(function() {
	var mdRender = document.querySelector('.md-render');
	if (!mdRender) return JSON.stringify({error: 'no content element'});
	return JSON.stringify({
		content: mdRender.innerText,
		name: document.title.split(' - ')[0].trim(),
		fullPath: window.location.pathname.replace('/document', '')
	});
})()`

const maxConsecutiveErrors = 3

// ScrapePages visits each page in the catalog, extracts content, and saves as markdown.
// It manages its own browser lifecycle and restarts on consecutive failures.
func ScrapePages(opts BrowserOpts, catalog *Catalog, outputDir string, delay time.Duration, sectionFilter string) error {
	progress := loadProgress(outputDir, catalog)

	// Build section filter set
	allowedSections := make(map[string]bool)
	if sectionFilter != "" {
		for _, s := range strings.Split(sectionFilter, ",") {
			allowedSections[strings.TrimSpace(s)] = true
		}
	}

	pending := 0
	for i := range progress.Pages {
		if progress.Pages[i].Status == "pending" || progress.Pages[i].Status == "error" {
			if len(allowedSections) > 0 && !allowedSections[progress.Pages[i].Section] {
				continue
			}
			pending++
		}
	}
	log.Printf("Scraping %d pages (total: %d, done: %d, errors: %d)", pending, progress.TotalPages, progress.Done, progress.Errors)

	// Start browser
	ctx, cancel := newBrowser(opts)
	defer func() { cancel() }()
	consecutiveErrors := 0

	for i := range progress.Pages {
		p := &progress.Pages[i]
		if p.Status == "done" {
			continue
		}
		if len(allowedSections) > 0 && !allowedSections[p.Section] {
			continue
		}

		remoteUpdateTime, err := scrapePage(ctx, p, outputDir, delay)
		if err != nil {
			p.Status = "error"
			p.Error = err.Error()
			consecutiveErrors++
			log.Printf("[ERROR] %s: %s", p.Name, err)

			// Restart browser after consecutive failures
			if consecutiveErrors >= maxConsecutiveErrors {
				log.Printf("Restarting browser after %d consecutive errors...", consecutiveErrors)
				cancel()
				time.Sleep(2 * time.Second)
				ctx, cancel = newBrowser(opts)
				consecutiveErrors = 0
			} else {
				time.Sleep(delay)
			}
		} else {
			p.Status = "done"
			p.ScrapedAt = time.Now().Format(time.RFC3339)
			p.UpdateTime = remoteUpdateTime
			consecutiveErrors = 0
		}

		// Recompute counters from actual statuses
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

		if p.Status == "done" {
			log.Printf("[%d/%d] %s -> %s", progress.Done, progress.TotalPages, p.Name, p.OutputFile)
		}
	}

	return nil
}

func scrapePage(ctx context.Context, page *ScrapedPage, outputDir string, delay time.Duration) (string, error) {
	docPath := page.TargetPath
	if docPath == "" {
		docPath = page.Path
	}
	if docPath == "" {
		return "", fmt.Errorf("no URL path available")
	}
	url := "https://open.larksuite.com/document" + docPath

	log.Printf("  Loading %s", url)

	var jsonResult string
	timeoutCtx, timeoutCancel := context.WithTimeout(ctx, 30*time.Second)
	defer timeoutCancel()
	err := chromedp.Run(timeoutCtx,
		chromedp.Navigate(url),
		chromedp.WaitReady("body", chromedp.ByQuery),
		chromedp.Sleep(5*time.Second),
		chromedp.Evaluate(extractContentJS, &jsonResult),
	)
	if err != nil {
		return "", fmt.Errorf("navigation failed: %w", err)
	}

	var result struct {
		Error      string      `json:"error"`
		Name       string      `json:"name"`
		Content    string      `json:"content"`
		FullPath   string      `json:"fullPath"`
		UpdateTime json.Number `json:"updateTime"`
	}
	if err := json.Unmarshal([]byte(jsonResult), &result); err != nil {
		return "", fmt.Errorf("failed to parse content JSON: %w", err)
	}

	// Fallback if fiber extraction failed
	if result.Error != "" || result.Content == "" {
		log.Printf("[FALLBACK] %s: fiber failed (%s), trying innerText", page.Name, result.Error)
		var fallbackResult string
		timeoutCtx2, timeoutCancel2 := context.WithTimeout(ctx, 15*time.Second)
		defer timeoutCancel2()
		err := chromedp.Run(timeoutCtx2,
			chromedp.Evaluate(fallbackContentJS, &fallbackResult),
		)
		if err != nil {
			return "", fmt.Errorf("fallback extraction failed: %w", err)
		}
		if err := json.Unmarshal([]byte(fallbackResult), &result); err != nil {
			return "", fmt.Errorf("failed to parse fallback JSON: %w", err)
		}
		if result.Error != "" {
			return "", fmt.Errorf("content extraction failed: %s", result.Error)
		}
	}

	// Determine output path
	relPath := deriveOutputPath(page)
	outFile := filepath.Join(outputDir, relPath)

	// Validate output path stays within outputDir
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
	parsed := ParseContent(result.Content)
	md := buildMarkdown(page.Name, docPath, service, resource, page.Section, string(result.UpdateTime), parsed)

	// Write file
	if err := os.MkdirAll(filepath.Dir(outFile), 0o755); err != nil {
		return "", fmt.Errorf("mkdir failed: %w", err)
	}
	if err := os.WriteFile(outFile, []byte(md), 0o644); err != nil {
		return "", fmt.Errorf("write failed: %w", err)
	}

	page.OutputFile = relPath
	time.Sleep(delay)
	return string(result.UpdateTime), nil
}

func deriveOutputPath(page *ScrapedPage) string {
	docPath := page.TargetPath
	if docPath == "" {
		docPath = page.Path
	}

	// New-style paths are already clean
	// e.g. /server-docs/contact-v3/user/create -> server-docs/contact-v3/user/create.md
	p := strings.TrimPrefix(docPath, "/")
	if strings.HasPrefix(p, "server-docs/") {
		return p + ".md"
	}

	// Old-style: try to find service-version pattern in the URL
	// e.g. uAjLw4CM/ukTMukTMukTM/minutes-v1/minute/get -> server-docs/minutes-v1/minute/get.md
	parts := strings.Split(p, "/")
	for i, part := range parts {
		if strings.Contains(part, "-v") && len(part) > 2 {
			return "server-docs/" + strings.Join(parts[i:], "/") + ".md"
		}
	}

	// Fallback: build path from section + parent breadcrumb + name
	// e.g. Section="API Call Guide", ParentPath=["API Call Guide", "Calling Process"], Name="Process overview"
	// -> server-docs/api-call-guide/calling-process/process-overview.md
	pathParts := make([]string, 0, len(page.ParentPath)+1)
	for _, seg := range page.ParentPath {
		pathParts = append(pathParts, slugify(seg))
	}
	pathParts = append(pathParts, slugify(page.Name))
	return "server-docs/" + strings.Join(pathParts, "/") + ".md"
}

// slugify converts a human-readable name to a URL-safe path segment.
func slugify(s string) string {
	s = strings.ToLower(s)
	s = strings.Map(func(r rune) rune {
		if r >= 'a' && r <= 'z' || r >= '0' && r <= '9' {
			return r
		}
		if r == ' ' || r == '_' || r == '/' {
			return '-'
		}
		return -1
	}, s)
	// Collapse multiple dashes
	for strings.Contains(s, "--") {
		s = strings.ReplaceAll(s, "--", "-")
	}
	s = strings.Trim(s, "-")
	if s == "" {
		s = "index"
	}
	return s
}

// looksLikeJunkPath checks if a path segment looks like a base64/hash junk value
// rather than a meaningful service or resource name.
func looksLikeJunkPath(s string) bool {
	if len(s) < 3 {
		return false
	}
	// Real service names contain hyphens (e.g., "contact-v3", "im-v1")
	// Real resource names are lowercase with underscores (e.g., "user", "message", "user_task")
	// Junk segments are typically mixed-case alphanumeric without meaningful separators
	upper := 0
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			upper++
		}
	}
	// If more than 20% uppercase and no hyphens/underscores, likely junk
	if upper > len(s)/5 && !strings.Contains(s, "-") && !strings.Contains(s, "_") {
		return true
	}
	return false
}

func parseServiceResource(docPath string) (service, resource string) {
	// /server-docs/contact-v3/user/create -> service=contact-v3, resource=user
	p := strings.TrimPrefix(docPath, "/server-docs/")
	if p == docPath {
		// Old-style: find service-version pattern
		parts := strings.Split(strings.TrimPrefix(docPath, "/"), "/")
		for i, part := range parts {
			if strings.Contains(part, "-v") && len(part) > 2 {
				p = strings.Join(parts[i:], "/")
				break
			}
		}
	}
	parts := strings.Split(p, "/")
	if len(parts) >= 1 {
		service = parts[0]
	}
	if len(parts) >= 2 {
		resource = parts[1]
	}
	if looksLikeJunkPath(service) {
		service = ""
		resource = ""
	}
	if looksLikeJunkPath(resource) {
		resource = ""
	}
	return
}

func buildMarkdown(title, docPath, service, resource, section, updateTime string, parsed *ParsedDoc) string {
	var sb strings.Builder
	sb.WriteString("---\n")
	sb.WriteString(fmt.Sprintf("title: %q\n", title))
	sb.WriteString(fmt.Sprintf("url: \"https://open.larksuite.com/document%s\"\n", docPath))
	if parsed.Method != "" {
		sb.WriteString(fmt.Sprintf("method: %q\n", parsed.Method))
	}
	if parsed.APIPath != "" {
		sb.WriteString(fmt.Sprintf("api_path: %q\n", parsed.APIPath))
	}
	if service != "" {
		sb.WriteString(fmt.Sprintf("service: %q\n", service))
	}
	if resource != "" {
		sb.WriteString(fmt.Sprintf("resource: %q\n", resource))
	}
	sb.WriteString(fmt.Sprintf("section: %q\n", section))
	if parsed.RateLimit != "" {
		sb.WriteString(fmt.Sprintf("rate_limit: %q\n", parsed.RateLimit))
	}
	if len(parsed.Scopes) > 0 {
		sb.WriteString("scopes:\n")
		for _, s := range parsed.Scopes {
			sb.WriteString(fmt.Sprintf("  - %q\n", s))
		}
	}
	if len(parsed.FieldScopes) > 0 {
		sb.WriteString("field_scopes:\n")
		for _, s := range parsed.FieldScopes {
			sb.WriteString(fmt.Sprintf("  - %q\n", s))
		}
	}
	if updateTime != "" {
		sb.WriteString(fmt.Sprintf("updated: %q\n", updateTime))
	}
	sb.WriteString("---\n\n")
	sb.WriteString(parsed.Content)
	sb.WriteString("\n")
	return sb.String()
}

// loadProgress loads existing progress or creates fresh progress from the catalog.
func loadProgress(outputDir string, catalog *Catalog) *Progress {
	progressFile := filepath.Join(outputDir, "progress.json")
	data, err := os.ReadFile(progressFile)
	if err == nil {
		var p Progress
		if json.Unmarshal(data, &p) == nil && p.TotalPages == catalog.TotalPages {
			log.Printf("Resuming from progress: %d/%d done", p.Done, p.TotalPages)
			return &p
		}
	}

	// Create fresh progress
	pages := make([]ScrapedPage, len(catalog.Pages))
	for i, p := range catalog.Pages {
		pages[i] = ScrapedPage{
			TargetPath: p.TargetPath,
			Path:       p.Path,
			Name:       p.Name,
			Section:    p.Section,
			ParentPath: p.ParentPath,
			Status:     "pending",
		}
	}
	return &Progress{
		StartedAt:  time.Now(),
		TotalPages: len(catalog.Pages),
		Pages:      pages,
	}
}

func saveProgress(outputDir string, progress *Progress) {
	progressFile := filepath.Join(outputDir, "progress.json")
	data, err := json.MarshalIndent(progress, "", "  ")
	if err != nil {
		log.Printf("Warning: failed to marshal progress: %v", err)
		return
	}
	if err := os.WriteFile(progressFile, data, 0o644); err != nil {
		log.Printf("Warning: failed to save progress: %v", err)
	}
}
