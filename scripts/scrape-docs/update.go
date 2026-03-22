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

// JS to extract just the updateTime from documentData (lightweight check)
const extractUpdateTimeJS = `(function() {
	var el = document.querySelector('.ud__tree');
	if (!el) return JSON.stringify({error: 'no tree element'});
	var fiberKey = Object.keys(el).find(function(k) { return k.startsWith('__reactFiber'); });
	if (!fiberKey) return JSON.stringify({error: 'no react fiber'});
	var fiber = el[fiberKey];
	for (var i = 0; i < 25; i++) {
		if (!fiber.return) break;
		fiber = fiber.return;
		var dd = fiber.memoizedProps && fiber.memoizedProps.documentData;
		if (dd && typeof dd.updateTime !== 'undefined') {
			return JSON.stringify({updateTime: dd.updateTime});
		}
	}
	return JSON.stringify({error: 'no documentData found'});
})()`

// UpdatePages re-discovers the tree, diffs against stored progress,
// and only scrapes pages that are new or have a different updateTime.
func UpdatePages(opts BrowserOpts, outputDir string, delay time.Duration, sectionFilter string) error {
	// Phase 1: Re-discover the tree
	log.Println("Phase 1: Discovering current documentation tree...")
	ctx, cancel := newBrowser(opts)
	defer cancel()
	newCatalog, err := DiscoverCatalog(ctx)
	if err != nil {
		return fmt.Errorf("discovery failed: %w", err)
	}

	// Phase 2: Load old progress
	oldProgress := loadProgressFile(outputDir)

	// Build lookup of old pages by their doc path (targetPath or path)
	oldPages := make(map[string]*ScrapedPage)
	if oldProgress != nil {
		for i := range oldProgress.Pages {
			p := &oldProgress.Pages[i]
			key := p.TargetPath
			if key == "" {
				key = p.Path
			}
			oldPages[key] = p
		}
	}

	// Build section filter
	allowedSections := make(map[string]bool)
	if sectionFilter != "" {
		for _, s := range strings.Split(sectionFilter, ",") {
			allowedSections[strings.TrimSpace(s)] = true
		}
	}

	// Phase 3: Determine what needs updating
	var toCheck []*CatalogNode // existing pages that need timestamp check
	var toScrape []*CatalogNode // new pages that definitely need scraping
	var unchanged int
	var removed int

	newPageKeys := make(map[string]bool)
	for _, page := range newCatalog.Pages {
		key := page.TargetPath
		if key == "" {
			key = page.Path
		}
		newPageKeys[key] = true

		if len(allowedSections) > 0 && !allowedSections[page.Section] {
			continue
		}

		old, exists := oldPages[key]
		if !exists {
			toScrape = append(toScrape, page)
		} else if old.Status != "done" {
			toScrape = append(toScrape, page)
		} else {
			toCheck = append(toCheck, page)
		}
	}

	// Count removed pages
	if oldProgress != nil {
		for key := range oldPages {
			if !newPageKeys[key] {
				removed++
			}
		}
	}

	log.Printf("Phase 2: %d new pages, %d to check for updates, %d removed", len(toScrape), len(toCheck), removed)

	// Phase 4: Check timestamps for existing pages
	log.Println("Phase 3: Checking timestamps for existing pages...")
	for _, page := range toCheck {
		key := page.TargetPath
		if key == "" {
			key = page.Path
		}
		old := oldPages[key]

		remoteTime, err := checkUpdateTime(ctx, page, delay)
		if err != nil {
			log.Printf("  [CHECK ERROR] %s: %s (will re-scrape)", page.Name, err)
			toScrape = append(toScrape, page)
			continue
		}

		if remoteTime != old.UpdateTime {
			log.Printf("  [CHANGED] %s (stored: %s, remote: %s)", page.Name, old.UpdateTime, remoteTime)
			toScrape = append(toScrape, page)
		} else {
			unchanged++
		}
	}

	log.Printf("Phase 3 complete: %d unchanged, %d to scrape", unchanged, len(toScrape))

	// Always persist the refreshed catalog
	catalogPath := filepath.Join(outputDir, "catalog.json")
	data, err := json.MarshalIndent(newCatalog, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal catalog: %w", err)
	}
	if err := os.WriteFile(catalogPath, data, 0o644); err != nil {
		return fmt.Errorf("failed to write catalog: %w", err)
	}

	if len(toScrape) == 0 {
		log.Println("No pages need scraping.")
		mergeProgress(outputDir, oldProgress, newCatalog)
		return nil
	}

	// Phase 5: Scrape only the changed/new pages
	log.Printf("Phase 4: Scraping %d pages...", len(toScrape))

	// Build a mini-catalog for the scrape
	miniCatalog := &Catalog{
		DiscoveredAt: newCatalog.DiscoveredAt,
		TotalNodes:   newCatalog.TotalNodes,
		TotalPages:   len(toScrape),
		Tree:         newCatalog.Tree,
		Pages:        toScrape,
	}

	cancel() // close discovery browser before scraping
	if err := ScrapePages(opts, miniCatalog, outputDir, delay, ""); err != nil {
		return err
	}

	// Merge results back into the full progress
	mergeProgress(outputDir, oldProgress, newCatalog)

	return nil
}

// checkUpdateTime navigates to a page and extracts only the updateTime.
func checkUpdateTime(ctx context.Context, page *CatalogNode, delay time.Duration) (string, error) {
	docPath := page.TargetPath
	if docPath == "" {
		docPath = page.Path
	}
	url := "https://open.larksuite.com/document" + docPath

	var jsonResult string
	timeoutCtx, timeoutCancel := context.WithTimeout(ctx, 30*time.Second)
	defer timeoutCancel()
	err := chromedp.Run(timeoutCtx,
		chromedp.Navigate(url),
		chromedp.WaitReady("body", chromedp.ByQuery),
		chromedp.Sleep(3*time.Second),
		chromedp.Evaluate(extractUpdateTimeJS, &jsonResult),
	)
	if err != nil {
		return "", fmt.Errorf("navigation failed: %w", err)
	}

	var result struct {
		Error      string      `json:"error"`
		UpdateTime json.Number `json:"updateTime"`
	}
	if err := json.Unmarshal([]byte(jsonResult), &result); err != nil {
		return "", fmt.Errorf("failed to parse JSON: %w", err)
	}
	if result.Error != "" {
		return "", fmt.Errorf("extraction error: %s", result.Error)
	}

	time.Sleep(delay)
	return string(result.UpdateTime), nil
}

// loadProgressFile loads existing progress without requiring a catalog.
func loadProgressFile(outputDir string) *Progress {
	data, err := readFile(outputDir, "progress.json")
	if err != nil {
		return nil
	}
	var p Progress
	if json.Unmarshal(data, &p) != nil {
		return nil
	}
	return &p
}

// mergeProgress merges the mini-progress from the latest scrape (if any) into
// the full progress file, preserving data for pages that weren't re-scraped.
// oldProgress is the pre-scrape progress snapshot; it must not be nil.
func mergeProgress(outputDir string, oldProgress *Progress, fullCatalog *Catalog) {
	if oldProgress == nil {
		return // nothing to merge into
	}

	// Build lookup of old pages
	oldByKey := make(map[string]*ScrapedPage)
	for i := range oldProgress.Pages {
		p := &oldProgress.Pages[i]
		key := p.TargetPath
		if key == "" {
			key = p.Path
		}
		oldByKey[key] = p
	}

	// Read the mini-progress that ScrapePages just wrote (nil if nothing was scraped)
	miniProgress := loadProgressFile(outputDir)
	miniByKey := make(map[string]*ScrapedPage)
	if miniProgress != nil {
		for i := range miniProgress.Pages {
			p := &miniProgress.Pages[i]
			key := p.TargetPath
			if key == "" {
				key = p.Path
			}
			miniByKey[key] = p
		}
	}

	// Build merged progress from the full catalog
	pages := make([]ScrapedPage, 0, len(fullCatalog.Pages))
	done := 0
	errors := 0
	for _, catPage := range fullCatalog.Pages {
		key := catPage.TargetPath
		if key == "" {
			key = catPage.Path
		}

		// Prefer mini-progress (just scraped) over old progress
		if mp, ok := miniByKey[key]; ok {
			pages = append(pages, *mp)
			if mp.Status == "done" {
				done++
			} else if mp.Status == "error" {
				errors++
			}
		} else if op, ok := oldByKey[key]; ok {
			pages = append(pages, *op)
			if op.Status == "done" {
				done++
			} else if op.Status == "error" {
				errors++
			}
		} else {
			pages = append(pages, ScrapedPage{
				TargetPath: catPage.TargetPath,
				Path:       catPage.Path,
				Name:       catPage.Name,
				Section:    catPage.Section,
				Status:     "pending",
			})
		}
	}

	merged := &Progress{
		StartedAt:   oldProgress.StartedAt,
		LastUpdated: time.Now(),
		TotalPages:  len(fullCatalog.Pages),
		Done:        done,
		Errors:      errors,
		Pages:       pages,
	}
	saveProgress(outputDir, merged)
	log.Printf("Merged progress: %d/%d done, %d errors", done, len(fullCatalog.Pages), errors)
}

func readFile(dir, name string) ([]byte, error) {
	return os.ReadFile(filepath.Join(dir, name))
}
