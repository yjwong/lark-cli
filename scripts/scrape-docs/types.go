package main

import "time"

// CatalogNode represents a node in the documentation sidebar tree.
type CatalogNode struct {
	ID         string        `json:"id"`
	Name       string        `json:"name"`
	Path       string        `json:"path"`        // old-style URL path
	TargetPath string        `json:"target_path"`  // new-style URL path
	Type       string        `json:"type"`         // "DocumentType" or "DirectoryType"
	Section    string        `json:"section"`      // top-level section name
	ParentPath []string      `json:"parent_path"`  // breadcrumb trail
	Children   []*CatalogNode `json:"children,omitempty"`
}

// Catalog is the full discovered tree plus a flat list of scrapeable pages.
type Catalog struct {
	DiscoveredAt time.Time      `json:"discovered_at"`
	TotalNodes   int            `json:"total_nodes"`
	TotalPages   int            `json:"total_pages"`
	Tree         []*CatalogNode `json:"tree"`
	Pages        []*CatalogNode `json:"pages"`
}

// ScrapedPage tracks the scraping status of a single page.
type ScrapedPage struct {
	TargetPath string   `json:"target_path"`
	Path       string   `json:"path"`        // old-style path fallback
	Name       string   `json:"name"`
	Section    string   `json:"section"`
	ParentPath []string `json:"parent_path"` // breadcrumb trail
	Status     string   `json:"status"`      // "pending", "done", "error", "unchanged"
	Error      string   `json:"error,omitempty"`
	OutputFile string   `json:"output_file,omitempty"`
	ScrapedAt  string   `json:"scraped_at,omitempty"`
	UpdateTime string   `json:"update_time,omitempty"` // remote updateTime from last scrape
}

// Progress tracks which pages have been scraped for resumability.
type Progress struct {
	StartedAt   time.Time     `json:"started_at"`
	LastUpdated time.Time     `json:"last_updated"`
	TotalPages  int           `json:"total_pages"`
	Done        int           `json:"done"`
	Errors      int           `json:"errors"`
	Pages       []ScrapedPage `json:"pages"`
}
