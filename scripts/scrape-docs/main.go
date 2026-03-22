package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/chromedp/chromedp"
)

// BrowserOpts holds chromedp allocator options for creating/restarting browsers.
type BrowserOpts []chromedp.ExecAllocatorOption

func main() {
	outputDir := flag.String("output", "../../docs/api", "Output directory for scraped docs")
	phase := flag.String("phase", "all", "Phase to run: discover, scrape, update, or all")
	delay := flag.Duration("delay", 2*time.Second, "Delay between page loads")
	headless := flag.Bool("headless", true, "Run browser in headless mode")
	sections := flag.String("sections", "", "Comma-separated section filter (e.g. 'Contacts,Calendar')")
	catalogFile := flag.String("catalog", "", "Path to existing catalog.json (skip discovery)")
	workers := flag.Int("workers", 8, "Number of concurrent HTTP workers (for http-scrape phase)")
	rps := flag.Float64("rps", 10, "Max requests per second (for http-scrape phase)")
	flag.Parse()

	absOutput, err := filepath.Abs(*outputDir)
	if err != nil {
		log.Fatalf("Failed to resolve output path: %v", err)
	}
	if err := os.MkdirAll(absOutput, 0o755); err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}

	opts := BrowserOpts(append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", *headless),
		chromedp.Flag("disable-gpu", true),
		chromedp.WindowSize(1920, 1080),
	))

	switch *phase {
	case "discover":
		catalog, err := runDiscover(opts, absOutput)
		if err != nil {
			log.Fatalf("Discovery failed: %v", err)
		}
		fmt.Printf("Catalog saved: %d pages discovered\n", catalog.TotalPages)

	case "scrape":
		catalog, err := loadCatalog(absOutput, *catalogFile)
		if err != nil {
			log.Fatalf("Failed to load catalog: %v", err)
		}
		if err := ScrapePages(opts, catalog, absOutput, *delay, *sections); err != nil {
			log.Fatalf("Scraping failed: %v", err)
		}
		if err := GenerateIndex(catalog, absOutput); err != nil {
			log.Printf("Warning: failed to generate index: %v", err)
		}
		fmt.Println("Scraping complete. Index generated.")

	case "http-scrape":
		catalog, err := loadCatalog(absOutput, *catalogFile)
		if err != nil {
			log.Fatalf("Failed to load catalog: %v", err)
		}
		if err := HTTPScrapePages(catalog, absOutput, *workers, *rps, *sections); err != nil {
			log.Fatalf("HTTP scraping failed: %v", err)
		}
		if err := GenerateIndex(catalog, absOutput); err != nil {
			log.Printf("Warning: failed to generate index: %v", err)
		}
		fmt.Println("HTTP scraping complete. Index generated.")

	case "http-all":
		catalog, err := runDiscover(opts, absOutput)
		if err != nil {
			log.Fatalf("Discovery failed: %v", err)
		}
		if err := HTTPScrapePages(catalog, absOutput, *workers, *rps, *sections); err != nil {
			log.Fatalf("HTTP scraping failed: %v", err)
		}
		if err := GenerateIndex(catalog, absOutput); err != nil {
			log.Printf("Warning: failed to generate index: %v", err)
		}
		fmt.Printf("Done. Scraped %d pages via HTTP.\n", catalog.TotalPages)

	case "all":
		catalog, err := runDiscover(opts, absOutput)
		if err != nil {
			log.Fatalf("Discovery failed: %v", err)
		}
		if err := ScrapePages(opts, catalog, absOutput, *delay, *sections); err != nil {
			log.Fatalf("Scraping failed: %v", err)
		}
		if err := GenerateIndex(catalog, absOutput); err != nil {
			log.Printf("Warning: failed to generate index: %v", err)
		}
		fmt.Printf("Done. Scraped %d pages.\n", catalog.TotalPages)

	case "update":
		if err := UpdatePages(opts, absOutput, *delay, *sections); err != nil {
			log.Fatalf("Update failed: %v", err)
		}
		catalog, err := loadCatalog(absOutput, "")
		if err == nil {
			if err := GenerateIndex(catalog, absOutput); err != nil {
				log.Printf("Warning: failed to generate index: %v", err)
			}
		}
		fmt.Println("Update complete.")

	case "validate":
		if err := ValidateOutput(absOutput); err != nil {
			log.Fatalf("Validation failed: %v", err)
		}

	default:
		log.Fatalf("Unknown phase: %s (use discover, scrape, http-scrape, http-all, update, validate, or all)", *phase)
	}
}

func runDiscover(opts BrowserOpts, outputDir string) (*Catalog, error) {
	ctx, cancel := newBrowser(opts)
	defer cancel()

	catalog, err := DiscoverCatalog(ctx)
	if err != nil {
		return nil, err
	}

	catalogPath := filepath.Join(outputDir, "catalog.json")
	data, err := json.MarshalIndent(catalog, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal catalog: %w", err)
	}
	if err := os.WriteFile(catalogPath, data, 0o644); err != nil {
		return nil, fmt.Errorf("failed to write catalog: %w", err)
	}
	log.Printf("Catalog saved to %s", catalogPath)
	return catalog, nil
}

func loadCatalog(outputDir, catalogFile string) (*Catalog, error) {
	path := catalogFile
	if path == "" {
		path = filepath.Join(outputDir, "catalog.json")
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("catalog not found at %s (run discovery first): %w", path, err)
	}
	var catalog Catalog
	if err := json.Unmarshal(data, &catalog); err != nil {
		return nil, fmt.Errorf("failed to parse catalog: %w", err)
	}
	log.Printf("Loaded catalog: %d pages", catalog.TotalPages)
	return &catalog, nil
}

// newBrowser creates a fresh chromedp browser context.
func newBrowser(opts BrowserOpts) (context.Context, context.CancelFunc) {
	allocCtx, allocCancel := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, ctxCancel := chromedp.NewContext(allocCtx)
	cancel := func() {
		ctxCancel()
		allocCancel()
	}
	return ctx, cancel
}
