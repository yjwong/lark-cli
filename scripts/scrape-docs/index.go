package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// GenerateIndex creates an index.md table of contents from the catalog.
func GenerateIndex(catalog *Catalog, outputDir string) error {
	var sb strings.Builder
	sb.WriteString("# Lark Server API Documentation\n\n")
	sb.WriteString(fmt.Sprintf("*%d pages across %d sections. Generated from [Lark Open Platform](https://open.larksuite.com/document/server-docs/getting-started/getting-started).*\n\n", catalog.TotalPages, len(catalog.Tree)))

	for _, section := range catalog.Tree {
		writeSection(&sb, section, 0)
	}

	outFile := filepath.Join(outputDir, "index.md")
	return os.WriteFile(outFile, []byte(sb.String()), 0o644)
}

func writeSection(sb *strings.Builder, node *CatalogNode, depth int) {
	if node.Type == "DocumentType" {
		sp := &ScrapedPage{
			TargetPath: node.TargetPath,
			Path:       node.Path,
			Name:       node.Name,
			Section:    node.Section,
			ParentPath: node.ParentPath,
		}
		relPath := deriveOutputPath(sp)
		indent := strings.Repeat("  ", depth)
		sb.WriteString(fmt.Sprintf("%s- [%s](%s)\n", indent, node.Name, relPath))
		return
	}

	// Directory node - write as heading or indented label
	if depth == 0 {
		sb.WriteString(fmt.Sprintf("## %s\n\n", node.Name))
	} else if depth == 1 {
		sb.WriteString(fmt.Sprintf("### %s\n\n", node.Name))
	} else {
		indent := strings.Repeat("  ", depth-2)
		sb.WriteString(fmt.Sprintf("%s- **%s**\n", indent, node.Name))
	}

	if node.Children != nil {
		for _, child := range node.Children {
			writeSection(sb, child, depth+1)
		}
		if depth <= 1 {
			sb.WriteString("\n")
		}
	}
}
