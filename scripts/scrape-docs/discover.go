package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

const extractTreeJS = `(function() {
	var el = document.querySelector('.ud__tree');
	if (!el) return JSON.stringify({error: 'no tree element'});
	var fiberKey = Object.keys(el).find(function(k) { return k.startsWith('__reactFiber'); });
	if (!fiberKey) return JSON.stringify({error: 'no react fiber'});
	var fiber = el[fiberKey];
	for (var i = 0; i < 25; i++) {
		if (!fiber.return) break;
		fiber = fiber.return;
		var data = fiber.memoizedProps && fiber.memoizedProps.data;
		if (data && data.children) {
			function serialize(node) {
				var r = {id: node.id || '', name: node.name || '', path: node.path || '', targetPath: node.targetPath || '', type: node.type || ''};
				if (node.children && node.children.length) r.children = node.children.map(serialize);
				return r;
			}
			return JSON.stringify(serialize(data));
		}
	}
	return JSON.stringify({error: 'no tree data found (searched 25 levels)'});
})()`

// DiscoverCatalog navigates to a Lark doc page and extracts the full sidebar tree.
func DiscoverCatalog(ctx context.Context) (*Catalog, error) {
	const startURL = "https://open.larksuite.com/document/server-docs/getting-started/getting-started"

	log.Println("Navigating to documentation page...")
	var jsonResult string
	timeoutCtx, timeoutCancel := context.WithTimeout(ctx, 30*time.Second)
	defer timeoutCancel()
	err := chromedp.Run(timeoutCtx,
		chromedp.Navigate(startURL),
		chromedp.WaitVisible(".ud__tree", chromedp.ByQuery),
		chromedp.Sleep(3*time.Second),
		chromedp.Evaluate(extractTreeJS, &jsonResult),
	)
	if err != nil {
		return nil, fmt.Errorf("chromedp discovery failed: %w", err)
	}

	// Check for error response
	var errResp struct {
		Error string `json:"error"`
	}
	if err := json.Unmarshal([]byte(jsonResult), &errResp); err == nil && errResp.Error != "" {
		return nil, fmt.Errorf("tree extraction failed: %s", errResp.Error)
	}

	// Parse the tree
	var root CatalogNode
	if err := json.Unmarshal([]byte(jsonResult), &root); err != nil {
		return nil, fmt.Errorf("failed to parse tree JSON: %w", err)
	}

	// Flatten to page list
	var pages []*CatalogNode
	var totalNodes int
	flattenTree(root.Children, "", nil, &pages, &totalNodes)

	catalog := &Catalog{
		DiscoveredAt: time.Now(),
		TotalNodes:   totalNodes,
		TotalPages:   len(pages),
		Tree:         root.Children,
		Pages:        pages,
	}

	log.Printf("Discovered %d total nodes, %d scrapeable pages", totalNodes, len(pages))
	return catalog, nil
}

// flattenTree recursively walks the tree and collects DocumentType nodes.
func flattenTree(nodes []*CatalogNode, section string, breadcrumb []string, pages *[]*CatalogNode, total *int) {
	for _, node := range nodes {
		*total++
		sec := section
		if sec == "" {
			sec = node.Name
		}
		node.Section = sec
		node.ParentPath = append([]string{}, breadcrumb...)

		if node.Type == "DocumentType" {
			*pages = append(*pages, node)
		}

		if node.Children != nil {
			bc := append(breadcrumb, node.Name)
			flattenTree(node.Children, sec, bc, pages, total)
		}
	}
}
