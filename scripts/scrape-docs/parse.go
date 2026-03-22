package main

import (
	"regexp"
	"strings"
)

// ParsedDoc holds structured data extracted from Lark's custom markdown.
type ParsedDoc struct {
	Description string
	Method      string
	APIPath     string
	RateLimit   string
	AppTypes    string
	Scopes      []string
	FieldScopes []string
	Content     string // cleaned markdown
}

var (
	// Regex patterns for extraction
	rePermName         = regexp.MustCompile(`<md-perm\s+name="([^"]+)"[^>]*>`)
	reHTMLTags         = regexp.MustCompile(`<[^>]+>`)
	reTryIt            = regexp.MustCompile(`\{Try it\}\([^)]*\)`)
	reSSLLink          = regexp.MustCompile(`\[([^\]]*)\]\(/ssl:ttdoc/[^)]*\)`)
	reBareSSLLink      = regexp.MustCompile(`\(/ssl:ttdoc/[^)]*\)`)
	reFencedHTML       = regexp.MustCompile(`(?m)^:::html\s*\n`)
	reFenceClose       = regexp.MustCompile(`(?m)^:::\s*$`)
	reAdmonition       = regexp.MustCompile(`(?m)^:::(note|warning|info|tip)\s*\n`)
	reMultiBlank       = regexp.MustCompile(`\n{3,}`)
	reSplitTableHeader = regexp.MustCompile(`(?m)\|\s*\n\s+`)
	reEnumItem         = regexp.MustCompile(`<md-enum-item\s+key="([^"]*)"[^>]*>([^<]*)</md-enum-item>`)
	reFieldName        = regexp.MustCompile(`<md-text\s+type="field-name"[^>]*>([^<]*)</md-text>`)
	reFieldType        = regexp.MustCompile(`<md-text\s+type="field-type"[^>]*>([^<]*)</md-text>`)
	reAppSupport       = regexp.MustCompile(`<md-app-support\s+types="([^"]*)"[^>]*>`)
	reTableRow         = regexp.MustCompile(`(?s)<md-t[dh][^>]*>(.*?)</md-t[dh]>`)
	reDtLevel          = regexp.MustCompile(`<md-dt-tr\s+level="(\d+)"`)
	reAlertBlock       = regexp.MustCompile(`(?s)<md-alert[^>]*>(.*?)</md-alert>`)
	reTooltip          = regexp.MustCompile(`(?s)<md-tooltip[^>]*>.*?</md-tooltip>`)
	reTagInline        = regexp.MustCompile(`<md-tag[^>]*>([^<]*)</md-tag>`)
	reHighlight        = regexp.MustCompile(`==([^=]+)==`)
	rePermFull         = regexp.MustCompile(`<md-perm\s+name="([^"]*)"[^>]*>[^<]*</md-perm>`)
	reMdTr             = regexp.MustCompile(`(?s)<md-tr[^>]*>(.*?)</md-tr>`)
	reMdCell           = regexp.MustCompile(`(?s)<md-t[dh][^>]*>(.*?)</md-t[dh]>`)
	reCollapseWS       = regexp.MustCompile(`\s*\n\s*`)
	reDtTh             = regexp.MustCompile(`(?s)<md-dt-th[^>]*>(.*?)</md-dt-th>`)
	reDtTr             = regexp.MustCompile(`(?s)<md-dt-tr[^>]*>(.*?)</md-dt-tr>`)
	reDtTd             = regexp.MustCompile(`(?s)<md-dt-td[^>]*>(.*?)</md-dt-td>`)
)

// ParseContent transforms raw Lark custom markdown into clean, token-efficient markdown.
// It also extracts structured metadata (scopes, method, path, etc.) for frontmatter.
func ParseContent(raw string) *ParsedDoc {
	doc := &ParsedDoc{}

	// Extract scopes before cleaning
	doc.Scopes = extractScopes(raw, "Required scopes")
	doc.FieldScopes = extractScopes(raw, "Required field scopes")

	// Extract API facts
	doc.Method = extractFact(raw, "HTTP Method")
	doc.APIPath = extractFact(raw, "HTTP URL")
	doc.RateLimit = extractFact(raw, "Rate Limit")
	doc.AppTypes = extractAppTypes(raw)

	// Clean and convert content
	doc.Content = cleanContent(raw)

	return doc
}

// extractScopes finds all <md-perm name="..."> within a specific section.
func extractScopes(raw, sectionLabel string) []string {
	// Find the section containing the label
	idx := strings.Index(raw, sectionLabel)
	if idx < 0 {
		return nil
	}

	// Find the end of this section (next <md-tr> or </md-table>)
	section := raw[idx:]
	endIdx := strings.Index(section, "</md-tr>")
	if endIdx < 0 {
		endIdx = len(section)
	}
	section = section[:endIdx]

	// Extract all perm names, excluding legacy/history scopes
	matches := rePermName.FindAllStringSubmatch(section, -1)
	var scopes []string
	seen := make(map[string]bool)
	for _, m := range matches {
		name := m[1]
		if !seen[name] {
			seen[name] = true
			scopes = append(scopes, name)
		}
	}
	return scopes
}

func extractFact(raw, label string) string {
	// Find pattern: <md-th>label</md-th>\n      <md-td>value</md-td>
	idx := strings.Index(raw, ">"+label+"</md-th>")
	if idx < 0 {
		return ""
	}
	after := raw[idx:]
	tdStart := strings.Index(after, "<md-td>")
	if tdStart < 0 {
		return ""
	}
	tdEnd := strings.Index(after[tdStart:], "</md-td>")
	if tdEnd < 0 {
		return ""
	}
	val := after[tdStart+7 : tdStart+tdEnd]
	// Clean HTML from value
	val = reHTMLTags.ReplaceAllString(val, "")
	// Clean markdown links, keep text
	val = reSSLLink.ReplaceAllString(val, "$1")
	return strings.TrimSpace(val)
}

func extractAppTypes(raw string) string {
	m := reAppSupport.FindStringSubmatch(raw)
	if m == nil {
		return ""
	}
	return m[1]
}

func cleanContent(raw string) string {
	s := raw

	// Convert spacing HTML entities early (before tag processing)
	s = strings.ReplaceAll(s, "&emsp;", "  ")
	s = strings.ReplaceAll(s, "&nbsp;", " ")

	// Strip ==highlight== markers
	s = reHighlight.ReplaceAllString(s, "$1")

	// Remove {Try it}(...) links
	s = reTryIt.ReplaceAllString(s, "")

	// Remove tooltips
	s = reTooltip.ReplaceAllString(s, "")

	// Convert <md-tag> to plain text
	s = reTagInline.ReplaceAllString(s, "`$1`")

	// Convert <md-enum-item key="x">desc</md-enum-item> to "- `x`: desc"
	s = reEnumItem.ReplaceAllStringFunc(s, func(m string) string {
		parts := reEnumItem.FindStringSubmatch(m)
		if len(parts) >= 3 {
			return "- `" + parts[1] + "`: " + strings.TrimSpace(parts[2])
		}
		return m
	})

	// Convert <md-text type="field-name">x</md-text> to `x`
	s = reFieldName.ReplaceAllString(s, "`$1`")
	s = reFieldType.ReplaceAllString(s, "`$1`")

	// Convert <md-perm name="x" desc="y" ...>text</md-perm> to "`x`"
	s = rePermFull.ReplaceAllString(s, "`$1`")

	// Convert <md-app-support types="x"> to text
	s = reAppSupport.ReplaceAllString(s, "$1")
	s = strings.ReplaceAll(s, "</md-app-support>", "")

	// Convert alert blocks to blockquotes
	s = reAlertBlock.ReplaceAllStringFunc(s, func(m string) string {
		parts := reAlertBlock.FindStringSubmatch(m)
		if len(parts) < 2 {
			return ""
		}
		inner := strings.TrimSpace(parts[1])
		if inner == "" {
			return ""
		}
		// Convert to blockquote
		lines := strings.Split(inner, "\n")
		var result []string
		for _, l := range lines {
			result = append(result, "> "+strings.TrimSpace(l))
		}
		return strings.Join(result, "\n")
	})

	// Convert md-tables to markdown tables
	s = convertTables(s)

	// Convert md-dt-tables (data tables with levels) to markdown tables
	s = convertDataTables(s)

	// Convert :::note/:::warning/:::html blocks to blockquotes or plain text
	s = convertAdmonitions(s)
	// Remove any remaining fenced HTML markers and closing :::
	s = reFencedHTML.ReplaceAllString(s, "")
	s = reFenceClose.ReplaceAllString(s, "")

	// Clean /ssl:ttdoc links — keep text, remove URL
	s = reSSLLink.ReplaceAllString(s, "$1")
	// Remove any remaining bare /ssl:ttdoc URLs in parentheses
	s = reBareSSLLink.ReplaceAllString(s, "")

	// Remove any remaining HTML tags
	s = reHTMLTags.ReplaceAllString(s, "")

	// Decode angle-bracket entities AFTER HTML tag stripping to preserve
	// type signatures like map<string, string> that would otherwise be
	// stripped as HTML tags
	s = strings.ReplaceAll(s, "&lt;", "<")
	s = strings.ReplaceAll(s, "&gt;", ">")
	s = strings.ReplaceAll(s, "&amp;", "&")

	// Clean up whitespace
	s = reMultiBlank.ReplaceAllString(s, "\n\n")

	// Fix split table headers (newline + whitespace inside cells)
	s = reSplitTableHeader.ReplaceAllString(s, "| ")

	// Auto-fence JSON blocks after example headings
	s = fenceJSONBlocks(s)

	// Unescape brackets in type names
	s = strings.ReplaceAll(s, `\[\]`, "[]")

	s = strings.TrimSpace(s)

	return s
}

// convertAdmonitions handles :::note, :::warning, :::info, :::tip, and :::html blocks.
// Admonitions (note/warning/etc.) become blockquotes with all lines prefixed.
// :::html blocks are stripped (content preserved as plain text).
func convertAdmonitions(s string) string {
	lines := strings.Split(s, "\n")
	var result []string
	i := 0
	for i < len(lines) {
		line := lines[i]
		m := reAdmonition.FindStringSubmatch(line + "\n")
		if m == nil {
			result = append(result, line)
			i++
			continue
		}

		// Found an admonition opener — collect lines until closing :::
		label := strings.ToUpper(m[1][:1]) + m[1][1:]
		i++
		var blockLines []string
		for i < len(lines) {
			if strings.TrimSpace(lines[i]) == ":::" {
				i++ // skip closing :::
				break
			}
			blockLines = append(blockLines, lines[i])
			i++
		}

		// Emit as blockquote
		if len(blockLines) == 0 {
			continue
		}
		result = append(result, "> **"+label+":** "+strings.TrimSpace(blockLines[0]))
		for _, bl := range blockLines[1:] {
			result = append(result, "> "+bl)
		}
	}
	return strings.Join(result, "\n")
}

// convertTables converts <md-table> structures to markdown tables.
func convertTables(s string) string {
	// Process each table block
	for {
		start := strings.Index(s, "<md-table>")
		if start < 0 {
			break
		}
		end := strings.Index(s[start:], "</md-table>")
		if end < 0 {
			break
		}
		end += start + len("</md-table>")

		tableHTML := s[start:end]
		mdTable := parseTable(tableHTML)
		s = s[:start] + mdTable + s[end:]
	}
	return s
}

func parseTable(html string) string {
	var rows [][]string

	// Split by rows
	rowParts := reMdTr.FindAllStringSubmatch(html, -1)
	for _, rp := range rowParts {
		cells := reMdCell.FindAllStringSubmatch(rp[1], -1)
		var row []string
		for _, c := range cells {
			cell := strings.TrimSpace(c[1])
			// Collapse multiline to single line
			cell = reCollapseWS.ReplaceAllString(cell, " ")
			row = append(row, cell)
		}
		if len(row) > 0 {
			rows = append(rows, row)
		}
	}

	if len(rows) == 0 {
		return ""
	}

	return renderMarkdownTable(rows)
}

func renderMarkdownTable(rows [][]string) string {
	if len(rows) == 0 {
		return ""
	}

	// Normalize column count
	maxCols := 0
	for _, row := range rows {
		if len(row) > maxCols {
			maxCols = len(row)
		}
	}
	for i := range rows {
		for len(rows[i]) < maxCols {
			rows[i] = append(rows[i], "")
		}
	}

	var sb strings.Builder

	// Header row
	sb.WriteString("| ")
	sb.WriteString(strings.Join(rows[0], " | "))
	sb.WriteString(" |\n")

	// Separator
	sb.WriteString("|")
	for range rows[0] {
		sb.WriteString(" --- |")
	}
	sb.WriteString("\n")

	// Data rows
	for _, row := range rows[1:] {
		sb.WriteString("| ")
		sb.WriteString(strings.Join(row, " | "))
		sb.WriteString(" |\n")
	}

	return sb.String()
}

// convertDataTables converts <md-dt-table> structures (with nesting levels) to markdown tables.
func convertDataTables(s string) string {
	for {
		start := strings.Index(s, "<md-dt-table>")
		if start < 0 {
			break
		}
		end := strings.Index(s[start:], "</md-dt-table>")
		if end < 0 {
			break
		}
		end += start + len("</md-dt-table>")

		tableHTML := s[start:end]
		mdTable := parseDtTable(tableHTML)
		s = s[:start] + mdTable + s[end:]
	}
	return s
}

func parseDtTable(html string) string {
	var rows [][]string

	// Extract header
	headerParts := reDtTh.FindAllStringSubmatch(html, -1)
	if len(headerParts) > 0 {
		var header []string
		for _, h := range headerParts {
			header = append(header, strings.TrimSpace(h[1]))
		}
		rows = append(rows, header)
	}

	// Extract data rows
	rowParts := reDtTr.FindAllStringSubmatch(html, -1)
	for _, rp := range rowParts {
		// Get nesting level for indentation
		levelMatch := reDtLevel.FindStringSubmatch(rp[0])
		level := 0
		if len(levelMatch) > 1 {
			for _, c := range levelMatch[1] {
				level = level*10 + int(c-'0')
			}
		}

		cells := reDtTd.FindAllStringSubmatch(rp[1], -1)
		var row []string
		for i, c := range cells {
			cell := strings.TrimSpace(c[1])
			cell = reCollapseWS.ReplaceAllString(cell, " ")
			// Indent first column (field name) based on level
			if i == 0 && level > 0 {
				cell = strings.Repeat("  ", level) + cell
			}
			row = append(row, cell)
		}
		if len(row) > 0 {
			rows = append(rows, row)
		}
	}

	if len(rows) == 0 {
		return ""
	}

	return renderMarkdownTable(rows)
}

// fenceJSONBlocks wraps unfenced JSON blocks that follow "example" headings in ```json fences.
func fenceJSONBlocks(s string) string {
	lines := strings.Split(s, "\n")
	var result []string
	i := 0
	for i < len(lines) {
		line := lines[i]
		result = append(result, line)

		// Check if this line is an "example" heading or contains "example"
		trimmed := strings.TrimSpace(strings.ToLower(line))
		isExampleHeading := (strings.Contains(trimmed, "example") &&
			(strings.HasPrefix(trimmed, "#") || strings.HasPrefix(trimmed, "**")))

		if !isExampleHeading {
			i++
			continue
		}

		// Skip blank lines after heading
		j := i + 1
		for j < len(lines) && strings.TrimSpace(lines[j]) == "" {
			result = append(result, lines[j])
			j++
		}

		// Check if next non-blank line starts with { or [
		if j >= len(lines) {
			i = j
			continue
		}
		nextTrimmed := strings.TrimSpace(lines[j])
		if nextTrimmed != "{" && nextTrimmed != "[" {
			i = j
			continue
		}

		// Check it's not already fenced (previous non-blank line would be ```)
		alreadyFenced := false
		for k := len(result) - 1; k >= 0; k-- {
			t := strings.TrimSpace(result[k])
			if t == "" {
				continue
			}
			if strings.HasPrefix(t, "```") {
				alreadyFenced = true
			}
			break
		}
		if alreadyFenced {
			i = j
			continue
		}

		// Find the closing brace/bracket at the same level
		opener := nextTrimmed[0]
		var closer byte = '}'
		if opener == '[' {
			closer = ']'
		}
		depth := 0
		end := j
		for end < len(lines) {
			lt := strings.TrimSpace(lines[end])
			for _, c := range lt {
				if byte(c) == opener {
					depth++
				} else if byte(c) == closer {
					depth--
				}
			}
			if depth <= 0 && lt != "" {
				break
			}
			end++
		}

		// Insert fence
		result = append(result, "```json")
		for k := j; k <= end && k < len(lines); k++ {
			result = append(result, lines[k])
		}
		result = append(result, "```")
		i = end + 1
		continue
	}
	return strings.Join(result, "\n")
}
