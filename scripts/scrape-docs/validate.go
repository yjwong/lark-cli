package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type ValidationIssue struct {
	File  string
	Line  int
	Issue string
}

// ValidateOutput walks all .md files in outputDir and checks for known quality issues.
func ValidateOutput(outputDir string) error {
	var issues []ValidationIssue

	err := filepath.Walk(outputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || !strings.HasSuffix(path, ".md") || info.Name() == "index.md" {
			return nil
		}

		relPath, _ := filepath.Rel(outputDir, path)
		data, err := os.ReadFile(path)
		if err != nil {
			issues = append(issues, ValidationIssue{relPath, 0, "failed to read file"})
			return nil
		}

		content := string(data)
		lines := strings.Split(content, "\n")

		fileIssues := validateFile(relPath, content, lines)
		issues = append(issues, fileIssues...)

		return nil
	})
	if err != nil {
		return fmt.Errorf("walk failed: %w", err)
	}

	// Report
	if len(issues) == 0 {
		log.Println("Validation passed: no issues found")
		return nil
	}

	// Group by issue type
	byType := make(map[string]int)
	for _, issue := range issues {
		byType[issue.Issue]++
	}

	log.Printf("Validation found %d issues across %d types:", len(issues), len(byType))
	for issueType, count := range byType {
		log.Printf("  %s: %d occurrences", issueType, count)
	}

	// Print first 5 examples of each type
	shown := make(map[string]int)
	for _, issue := range issues {
		if shown[issue.Issue] < 5 {
			if issue.Line > 0 {
				log.Printf("  [%s] %s:%d", issue.Issue, issue.File, issue.Line)
			} else {
				log.Printf("  [%s] %s", issue.Issue, issue.File)
			}
			shown[issue.Issue]++
		}
	}

	return fmt.Errorf("validation failed: %d issues found", len(issues))
}

// reValidateHighlight matches ==text== highlight markers (== wrapping non-== content).
var reValidateHighlight = regexp.MustCompile(`==[^=]+==[^=]`)

// reValidateJunkFrontmatter matches values that look like hashes/base64: uppercase letters + no hyphens.
var reValidateJunkFrontmatter = regexp.MustCompile(`^[A-Z0-9a-z+/]{16,}$`)

func validateFile(relPath, content string, lines []string) []ValidationIssue {
	var issues []ValidationIssue

	add := func(line int, issueType string) {
		issues = append(issues, ValidationIssue{File: relPath, Line: line, Issue: issueType})
	}

	// 1. split-table-header: pipe cell immediately followed by whitespace+content on next line
	//    Detect "| \n" or a line that is only "| " (possibly with trailing spaces).
	for i, line := range lines {
		trimmed := strings.TrimRight(line, " \t")
		if strings.HasSuffix(trimmed, "|") && trimmed == "|" {
			// A line that is just "|" inside a table context
			if i+1 < len(lines) && strings.TrimSpace(lines[i+1]) != "" {
				add(i+1, "split-table-header")
			}
		}
	}
	// Also check raw content for | followed by newline + indented content
	reSplitTable := regexp.MustCompile(`\|\s*\n\s+\S`)
	if reSplitTable.MatchString(content) {
		// Find first occurrence line number
		for i, line := range lines {
			trimmed := strings.TrimRight(line, " \t\r")
			if trimmed == "|" && i+1 < len(lines) {
				next := lines[i+1]
				if len(next) > 0 && (next[0] == ' ' || next[0] == '\t') && strings.TrimSpace(next) != "" {
					add(i+1, "split-table-header")
				}
			}
		}
	}

	// 2. unfenced-json: a line starting with "{" after a blank line where the preceding non-blank
	//    line contains "example" (case-insensitive) and no ``` fence between them.
	inFence := false
	lastNonBlankLine := ""
	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "```") {
			inFence = !inFence
		}
		if !inFence {
			if trimmed == "" {
				// blank line: candidate boundary
			} else {
				if strings.HasPrefix(trimmed, "{") {
					// Check if previous non-blank line contained "example"
					if strings.Contains(strings.ToLower(lastNonBlankLine), "example") {
						add(i+1, "unfenced-json")
					}
				}
				lastNonBlankLine = trimmed
			}
		}
	}

	// 3. raw-html-fence: lines containing :::html
	for i, line := range lines {
		if strings.Contains(line, ":::html") {
			add(i+1, "raw-html-fence")
		}
	}

	// 4. raw-admonition: lines containing :::note or :::warning
	for i, line := range lines {
		lower := strings.ToLower(line)
		if strings.Contains(lower, ":::note") || strings.Contains(lower, ":::warning") {
			add(i+1, "raw-admonition")
		}
	}

	// 5. html-entities: &emsp; or &nbsp; remaining
	for i, line := range lines {
		if strings.Contains(line, "&emsp;") || strings.Contains(line, "&nbsp;") {
			add(i+1, "html-entities")
		}
	}

	// 6. ssl-link: /ssl:ttdoc remaining
	for i, line := range lines {
		if strings.Contains(line, "/ssl:ttdoc") {
			add(i+1, "ssl-link")
		}
	}

	// 7. escaped-brackets: \[\] remaining
	for i, line := range lines {
		if strings.Contains(line, `\[\]`) {
			add(i+1, "escaped-brackets")
		}
	}

	// 8. highlight-markers: ==text== pattern
	for i, line := range lines {
		if reValidateHighlight.MatchString(line) {
			add(i+1, "highlight-markers")
		}
	}

	// 9. empty-content: fewer than 20 lines total
	if len(lines) < 20 {
		add(0, "empty-content")
	}

	// 10. junk-frontmatter: frontmatter resource: or service: with hash/base64-like value
	if strings.HasPrefix(content, "---") {
		endFM := strings.Index(content[3:], "\n---")
		if endFM >= 0 {
			frontmatter := content[3 : 3+endFM]
			for i, line := range strings.Split(frontmatter, "\n") {
				line = strings.TrimSpace(line)
				for _, key := range []string{"resource:", "service:"} {
					if strings.HasPrefix(line, key) {
						val := strings.TrimSpace(strings.TrimPrefix(line, key))
						// Strip quotes if present
						val = strings.Trim(val, `"'`)
						if reValidateJunkFrontmatter.MatchString(val) && !strings.Contains(val, "-") {
							add(i+1, "junk-frontmatter")
						}
					}
				}
			}
		}
	}

	return issues
}
