package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"
)

// writeMiniProgress writes a Progress struct to progress.json in dir,
// simulating what ScrapePages would write after a partial scrape.
func writeMiniProgress(t *testing.T, dir string, p *Progress) {
	t.Helper()
	data, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		t.Fatalf("marshal mini progress: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dir, "progress.json"), data, 0o644); err != nil {
		t.Fatalf("write mini progress: %v", err)
	}
}

// makeCatalogNode builds a minimal *CatalogNode for testing.
func makeCatalogNode(targetPath, name string) *CatalogNode {
	return &CatalogNode{
		TargetPath: targetPath,
		Name:       name,
		Type:       "DocumentType",
	}
}

func TestMergeProgress(t *testing.T) {
	dir := t.TempDir()

	// Old progress: page-a done, page-b done, page-c pending
	startedAt := time.Now().Add(-1 * time.Hour)
	oldProgress := &Progress{
		StartedAt:  startedAt,
		TotalPages: 3,
		Done:       2,
		Pages: []ScrapedPage{
			{TargetPath: "/server-docs/svc/page-a", Name: "page-a", Status: "done", ScrapedAt: "2024-01-01T00:00:00Z"},
			{TargetPath: "/server-docs/svc/page-b", Name: "page-b", Status: "done", ScrapedAt: "2024-01-01T00:00:00Z"},
			{TargetPath: "/server-docs/svc/page-c", Name: "page-c", Status: "pending"},
		},
	}

	// Mini progress on disk: only page-c, now done (what ScrapePages wrote)
	miniProgress := &Progress{
		StartedAt:  time.Now(),
		TotalPages: 1,
		Done:       1,
		Pages: []ScrapedPage{
			{TargetPath: "/server-docs/svc/page-c", Name: "page-c", Status: "done", ScrapedAt: "2024-06-01T00:00:00Z"},
		},
	}
	writeMiniProgress(t, dir, miniProgress)

	// Full catalog with all 3 pages
	fullCatalog := &Catalog{
		TotalPages: 3,
		Pages: []*CatalogNode{
			makeCatalogNode("/server-docs/svc/page-a", "page-a"),
			makeCatalogNode("/server-docs/svc/page-b", "page-b"),
			makeCatalogNode("/server-docs/svc/page-c", "page-c"),
		},
	}

	mergeProgress(dir, oldProgress, fullCatalog)

	// Read merged result
	result := loadProgressFile(dir)
	if result == nil {
		t.Fatal("mergeProgress did not write a progress file")
	}

	if result.TotalPages != 3 {
		t.Errorf("TotalPages = %d, want 3", result.TotalPages)
	}
	if result.Done != 3 {
		t.Errorf("Done = %d, want 3", result.Done)
	}
	if result.Errors != 0 {
		t.Errorf("Errors = %d, want 0", result.Errors)
	}

	// Build lookup of merged pages by key
	byKey := make(map[string]ScrapedPage)
	for _, p := range result.Pages {
		byKey[p.TargetPath] = p
	}

	for _, key := range []string{"/server-docs/svc/page-a", "/server-docs/svc/page-b"} {
		p, ok := byKey[key]
		if !ok {
			t.Errorf("merged result missing %q", key)
			continue
		}
		if p.Status != "done" {
			t.Errorf("%q: status = %q, want %q", key, p.Status, "done")
		}
	}

	pageC, ok := byKey["/server-docs/svc/page-c"]
	if !ok {
		t.Fatal("merged result missing page-c")
	}
	if pageC.Status != "done" {
		t.Errorf("page-c: status = %q, want %q", pageC.Status, "done")
	}
	// Confirm it came from mini (has the mini ScrapedAt)
	if pageC.ScrapedAt != "2024-06-01T00:00:00Z" {
		t.Errorf("page-c: ScrapedAt = %q, want mini value %q", pageC.ScrapedAt, "2024-06-01T00:00:00Z")
	}
}

func TestMergeProgressNoScrape(t *testing.T) {
	dir := t.TempDir()

	// Old progress: all 3 pages done; nothing on disk (no mini progress written)
	startedAt := time.Now().Add(-1 * time.Hour)
	oldProgress := &Progress{
		StartedAt:  startedAt,
		TotalPages: 3,
		Done:       3,
		Pages: []ScrapedPage{
			{TargetPath: "/server-docs/svc/page-a", Name: "page-a", Status: "done", ScrapedAt: "2024-01-01T00:00:00Z"},
			{TargetPath: "/server-docs/svc/page-b", Name: "page-b", Status: "done", ScrapedAt: "2024-01-02T00:00:00Z"},
			{TargetPath: "/server-docs/svc/page-c", Name: "page-c", Status: "done", ScrapedAt: "2024-01-03T00:00:00Z"},
		},
	}

	// No mini progress.json on disk

	fullCatalog := &Catalog{
		TotalPages: 3,
		Pages: []*CatalogNode{
			makeCatalogNode("/server-docs/svc/page-a", "page-a"),
			makeCatalogNode("/server-docs/svc/page-b", "page-b"),
			makeCatalogNode("/server-docs/svc/page-c", "page-c"),
		},
	}

	mergeProgress(dir, oldProgress, fullCatalog)

	result := loadProgressFile(dir)
	if result == nil {
		t.Fatal("mergeProgress did not write a progress file")
	}

	if result.Done != 3 {
		t.Errorf("Done = %d, want 3", result.Done)
	}
	if result.Errors != 0 {
		t.Errorf("Errors = %d, want 0", result.Errors)
	}
	if len(result.Pages) != 3 {
		t.Errorf("len(Pages) = %d, want 3", len(result.Pages))
	}

	for _, p := range result.Pages {
		if p.Status != "done" {
			t.Errorf("%q: status = %q, want %q", p.TargetPath, p.Status, "done")
		}
	}
}

func TestMergeProgressNewPages(t *testing.T) {
	dir := t.TempDir()

	// Old progress: 2 pages done
	startedAt := time.Now().Add(-1 * time.Hour)
	oldProgress := &Progress{
		StartedAt:  startedAt,
		TotalPages: 2,
		Done:       2,
		Pages: []ScrapedPage{
			{TargetPath: "/server-docs/svc/page-a", Name: "page-a", Status: "done", ScrapedAt: "2024-01-01T00:00:00Z"},
			{TargetPath: "/server-docs/svc/page-b", Name: "page-b", Status: "done", ScrapedAt: "2024-01-02T00:00:00Z"},
		},
	}

	// Catalog has 3 pages — page-new is not in old progress
	fullCatalog := &Catalog{
		TotalPages: 3,
		Pages: []*CatalogNode{
			makeCatalogNode("/server-docs/svc/page-a", "page-a"),
			makeCatalogNode("/server-docs/svc/page-b", "page-b"),
			makeCatalogNode("/server-docs/svc/page-new", "page-new"),
		},
	}

	// No mini progress.json on disk (the new page wasn't scraped in this run)
	mergeProgress(dir, oldProgress, fullCatalog)

	result := loadProgressFile(dir)
	if result == nil {
		t.Fatal("mergeProgress did not write a progress file")
	}

	if result.TotalPages != 3 {
		t.Errorf("TotalPages = %d, want 3", result.TotalPages)
	}
	if result.Done != 2 {
		t.Errorf("Done = %d, want 2", result.Done)
	}

	byKey := make(map[string]ScrapedPage)
	for _, p := range result.Pages {
		byKey[p.TargetPath] = p
	}

	pageNew, ok := byKey["/server-docs/svc/page-new"]
	if !ok {
		t.Fatal("merged result missing page-new")
	}
	if pageNew.Status != "pending" {
		t.Errorf("page-new: status = %q, want %q", pageNew.Status, "pending")
	}

	for _, key := range []string{"/server-docs/svc/page-a", "/server-docs/svc/page-b"} {
		p, ok := byKey[key]
		if !ok {
			t.Errorf("merged result missing %q", key)
			continue
		}
		if p.Status != "done" {
			t.Errorf("%q: status = %q, want %q", key, p.Status, "done")
		}
	}
}
