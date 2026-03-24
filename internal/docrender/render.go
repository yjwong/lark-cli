package docrender

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/yjwong/lark-cli/internal/api"
)

// RenderOptions configures the block renderer
type RenderOptions struct {
	// UserNames maps user IDs (e.g., "ou_xxx") to display names.
	// If nil or a user ID is missing, falls back to @user_id.
	UserNames map[string]string
	// SheetData maps embedded sheet block tokens to their cell values.
	// If nil or a token is missing, falls back to [sheet: token].
	SheetData map[string][][]any
}

// blockNode wraps a DocumentBlock with resolved children for tree traversal
type blockNode struct {
	block    api.DocumentBlock
	children []*blockNode
}

// renderer holds state for a single render pass
type renderer struct {
	opts      RenderOptions
	inCodeBlock bool
}

// RenderBlocks converts a flat list of document blocks into markdown
func RenderBlocks(blocks []api.DocumentBlock) string {
	return RenderBlocksWithOptions(blocks, RenderOptions{})
}

// RenderBlocksWithOptions converts blocks to markdown with custom options
func RenderBlocksWithOptions(blocks []api.DocumentBlock, opts RenderOptions) string {
	if len(blocks) == 0 {
		return ""
	}

	tree := buildBlockTree(blocks)
	if tree == nil {
		return ""
	}

	r := &renderer{opts: opts}
	var sb strings.Builder
	r.renderChildren(&sb, tree, 0)
	return strings.TrimRight(sb.String(), "\n") + "\n"
}

// ExtractUserIDs collects all unique mention_user IDs from a block list
func ExtractUserIDs(blocks []api.DocumentBlock) []string {
	seen := make(map[string]bool)
	var ids []string
	for _, b := range blocks {
		for _, tb := range allTextBlocks(&b) {
			for _, elem := range tb.Elements {
				if elem.MentionUser != nil && !seen[elem.MentionUser.UserID] {
					seen[elem.MentionUser.UserID] = true
					ids = append(ids, elem.MentionUser.UserID)
				}
			}
		}
	}
	return ids
}

// ExtractImageTokens collects all unique image tokens from a block list
func ExtractImageTokens(blocks []api.DocumentBlock) []string {
	seen := make(map[string]bool)
	var tokens []string
	for _, b := range blocks {
		if b.BlockType == 27 && b.Image != nil && b.Image.Token != "" && !seen[b.Image.Token] {
			seen[b.Image.Token] = true
			tokens = append(tokens, b.Image.Token)
		}
	}
	return tokens
}

// ExtractSheetTokens collects all unique embedded sheet tokens from a block list
func ExtractSheetTokens(blocks []api.DocumentBlock) []string {
	seen := make(map[string]bool)
	var tokens []string
	for _, b := range blocks {
		if b.BlockType == 30 && b.Sheet != nil && b.Sheet.Token != "" && !seen[b.Sheet.Token] {
			seen[b.Sheet.Token] = true
			tokens = append(tokens, b.Sheet.Token)
		}
	}
	return tokens
}

// allTextBlocks returns all TextBlock pointers from a DocumentBlock
func allTextBlocks(b *api.DocumentBlock) []*api.TextBlock {
	candidates := []*api.TextBlock{
		b.Page, b.Text, b.Heading1, b.Heading2, b.Heading3,
		b.Heading4, b.Heading5, b.Heading6, b.Heading7, b.Heading8, b.Heading9,
		b.Bullet, b.Ordered, b.Code, b.Quote, b.TodoBlock,
	}
	var result []*api.TextBlock
	for _, tb := range candidates {
		if tb != nil {
			result = append(result, tb)
		}
	}
	return result
}

// buildBlockTree indexes blocks by ID and builds a tree rooted at the Page block (type 1)
func buildBlockTree(blocks []api.DocumentBlock) *blockNode {
	index := make(map[string]*blockNode, len(blocks))
	var root *blockNode

	// Create nodes
	for i := range blocks {
		node := &blockNode{block: blocks[i]}
		index[blocks[i].BlockID] = node
		if blocks[i].BlockType == 1 {
			root = node
		}
	}

	if root == nil {
		return nil
	}

	// Resolve children
	for _, node := range index {
		for _, childID := range node.block.Children {
			if child, ok := index[childID]; ok {
				node.children = append(node.children, child)
			}
		}
	}

	return root
}

// renderChildren renders all children of a node
func (r *renderer) renderChildren(sb *strings.Builder, node *blockNode, depth int) {
	orderedCounter := 0
	for _, child := range node.children {
		if child.block.BlockType == 13 { // ordered list
			orderedCounter++
		} else {
			orderedCounter = 0
		}
		r.renderBlock(sb, child, depth, orderedCounter)
	}
}

// indentPrefix returns the indentation string for nested list items
func indentPrefix(depth int) string {
	if depth <= 0 {
		return ""
	}
	return strings.Repeat("  ", depth)
}

// renderBlock renders a single block to markdown
func (r *renderer) renderBlock(sb *strings.Builder, node *blockNode, depth int, orderedNum int) {
	switch node.block.BlockType {
	case 1: // Page
		r.renderChildren(sb, node, depth)

	case 2: // Text
		text := r.renderTextBlock(node.block.Text)
		if text != "" {
			sb.WriteString(escapeBlockSyntax(text))
		}
		sb.WriteString("\n")

	case 3, 4, 5, 6, 7, 8, 9, 10, 11: // Headings H1-H9
		level := node.block.BlockType - 2
		tb := getHeadingTextBlock(&node.block, level)
		text := r.renderTextBlock(tb)
		// CommonMark only supports ATX headings up to level 6
		mdLevel := level
		if mdLevel > 6 {
			mdLevel = 6
		}
		sb.WriteString(strings.Repeat("#", mdLevel))
		sb.WriteString(" ")
		sb.WriteString(text)
		sb.WriteString("\n\n")

	case 12: // Bullet
		text := r.renderTextBlock(node.block.Bullet)
		sb.WriteString(indentPrefix(depth))
		sb.WriteString("- ")
		sb.WriteString(text)
		sb.WriteString("\n")
		if len(node.children) > 0 {
			r.renderChildren(sb, node, depth+1)
		}

	case 13: // Ordered
		text := r.renderTextBlock(node.block.Ordered)
		sb.WriteString(indentPrefix(depth))
		sb.WriteString(fmt.Sprintf("%d. ", orderedNum))
		sb.WriteString(text)
		sb.WriteString("\n")
		if len(node.children) > 0 {
			r.renderChildren(sb, node, depth+1)
		}

	case 14: // Code
		if node.block.Code != nil {
			lang := ""
			if node.block.Code.Style != nil {
				lang = codeLanguageName(node.block.Code.Style.Language)
			}
			r.inCodeBlock = true
			content := r.renderTextBlock(node.block.Code)
			r.inCodeBlock = false
			fence := codeFence(content)
			sb.WriteString(fence)
			sb.WriteString(lang)
			sb.WriteString("\n")
			sb.WriteString(content)
			sb.WriteString("\n")
			sb.WriteString(fence)
			sb.WriteString("\n\n")
		}

	case 15: // Quote
		text := r.renderTextBlock(node.block.Quote)
		for _, line := range strings.Split(text, "\n") {
			sb.WriteString("> ")
			sb.WriteString(line)
			sb.WriteString("\n")
		}
		sb.WriteString("\n")

	case 17: // Todo
		if node.block.TodoBlock != nil {
			done := false
			if node.block.TodoBlock.Style != nil {
				done = node.block.TodoBlock.Style.Done
			}
			text := r.renderTextBlock(node.block.TodoBlock)
			if done {
				sb.WriteString("- [x] ")
			} else {
				sb.WriteString("- [ ] ")
			}
			sb.WriteString(text)
			sb.WriteString("\n")
		}

	case 19: // Callout
		if len(node.children) > 0 {
			var calloutSB strings.Builder
			r.renderChildren(&calloutSB, node, depth)
			for _, line := range strings.Split(strings.TrimRight(calloutSB.String(), "\n"), "\n") {
				sb.WriteString("> ")
				sb.WriteString(line)
				sb.WriteString("\n")
			}
			sb.WriteString("\n")
		}

	case 18: // Bitable
		sb.WriteString("[bitable]\n\n")

	case 20: // ChatCard
		sb.WriteString("[chat]\n\n")

	case 21: // Diagram
		sb.WriteString("[diagram]\n\n")

	case 22: // Divider
		sb.WriteString("---\n\n")

	case 23: // File
		if node.block.File != nil {
			name := node.block.File.Name
			if name == "" {
				name = node.block.File.Token
			}
			sb.WriteString(fmt.Sprintf("[file: %s]\n\n", name))
		} else {
			sb.WriteString("[file]\n\n")
		}

	case 24: // Grid — render columns sequentially
		r.renderChildren(sb, node, depth)

	case 25: // GridColumn — render children
		r.renderChildren(sb, node, depth)

	case 26: // Iframe
		if node.block.Iframe != nil && node.block.Iframe.Component != nil && node.block.Iframe.Component.URL != "" {
			sb.WriteString(fmt.Sprintf("[embed: %s]\n\n", node.block.Iframe.Component.URL))
		} else {
			sb.WriteString("[embed]\n\n")
		}

	case 27: // Image
		if node.block.Image != nil {
			sb.WriteString(fmt.Sprintf("[image: %s]\n\n", node.block.Image.Token))
		}

	case 28: // ISV
		sb.WriteString("[isv]\n\n")

	case 29: // Mindnote
		sb.WriteString("[mindnote]\n\n")

	case 30: // Sheet
		if node.block.Sheet != nil && node.block.Sheet.Token != "" {
			values, resolved := r.opts.SheetData[node.block.Sheet.Token]
			if resolved && len(values) > 0 {
				r.renderSheetAsTable(sb, values)
			} else if resolved {
				// Fetched successfully but sheet is empty
				sb.WriteString("[empty sheet]\n\n")
			} else {
				sb.WriteString(fmt.Sprintf("[sheet: %s]\n\n", node.block.Sheet.Token))
			}
		} else {
			sb.WriteString("[sheet]\n\n")
		}

	case 31: // Table
		r.renderTable(sb, node)

	case 32: // TableCell — handled by renderTable
		// Skip: rendered as part of parent table

	case 33: // View — associated with File blocks, no user-facing content
		// Skip silently

	case 34: // QuoteContainer — render children with > prefix
		if len(node.children) > 0 {
			var qcSB strings.Builder
			r.renderChildren(&qcSB, node, depth)
			for _, line := range strings.Split(strings.TrimRight(qcSB.String(), "\n"), "\n") {
				sb.WriteString("> ")
				sb.WriteString(line)
				sb.WriteString("\n")
			}
			sb.WriteString("\n")
		}

	case 35: // Task
		if node.block.Task != nil && node.block.Task.TaskID != "" {
			sb.WriteString(fmt.Sprintf("[task: %s]\n\n", node.block.Task.TaskID))
		} else {
			sb.WriteString("[task]\n\n")
		}

	case 36, 37, 38, 39: // OKR, OKR Objective, OKR Key Result, OKR Progress
		sb.WriteString("[okr]\n\n")

	case 40: // AddOns
		r.renderAddOns(sb, &node.block)

	case 41: // JiraIssue
		if node.block.JiraIssue != nil {
			key := node.block.JiraIssue.Key
			if key == "" {
				key = node.block.JiraIssue.ID
			}
			if key != "" {
				sb.WriteString(fmt.Sprintf("[jira: %s]\n\n", key))
			} else {
				sb.WriteString("[jira]\n\n")
			}
		} else {
			sb.WriteString("[jira]\n\n")
		}

	case 42: // WikiCatalog
		if node.block.WikiCatalog != nil && node.block.WikiCatalog.WikiToken != "" {
			sb.WriteString(fmt.Sprintf("[wiki: %s]\n\n", node.block.WikiCatalog.WikiToken))
		} else {
			sb.WriteString("[wiki]\n\n")
		}

	default:
		// Warn about unsupported block types to stderr
		fmt.Fprintf(os.Stderr, "warning: unsupported block type %d (block_id: %s), use --raw for full content\n",
			node.block.BlockType, node.block.BlockID)
	}
}

// renderTable renders a table block as a markdown table
func (r *renderer) renderTable(sb *strings.Builder, node *blockNode) {
	if node.block.Table == nil || node.block.Table.Property == nil {
		return
	}

	prop := node.block.Table.Property
	rows := prop.RowSize
	cols := prop.ColumnSize

	if rows == 0 || cols == 0 || len(node.children) == 0 {
		return
	}

	// Prefer Table.Cells for cell ordering (canonical per API docs), fall back to children
	cells := node.children
	if len(node.block.Table.Cells) > 0 {
		// Build an index of children by block ID for lookup
		childIndex := make(map[string]*blockNode, len(node.children))
		for _, child := range node.children {
			childIndex[child.block.BlockID] = child
		}
		ordered := make([]*blockNode, 0, len(node.block.Table.Cells))
		for _, cellID := range node.block.Table.Cells {
			if cell, ok := childIndex[cellID]; ok {
				ordered = append(ordered, cell)
			}
		}
		if len(ordered) > 0 {
			cells = ordered
		}
	}

	// Check if table has a header row
	hasHeader := prop.HeaderRow

	for row := 0; row < rows; row++ {
		sb.WriteString("|")
		for c := 0; c < cols; c++ {
			idx := row*cols + c
			cellContent := ""
			if idx < len(cells) {
				cellContent = r.renderTableCell(cells[idx])
			}
			// Escape pipes in cell content
			cellContent = strings.ReplaceAll(cellContent, "|", "\\|")
			sb.WriteString(" ")
			sb.WriteString(cellContent)
			sb.WriteString(" |")
		}
		sb.WriteString("\n")

		// Add separator after first row (header or not — required by markdown)
		if row == 0 {
			sb.WriteString("|")
			for c := 0; c < cols; c++ {
				sb.WriteString(" --- |")
			}
			sb.WriteString("\n")

			// If no header row, the first data row was just rendered as header.
			// Markdown requires a header, so we accept this tradeoff.
			_ = hasHeader
		}
	}
	sb.WriteString("\n")
}

// renderTableCell renders the content of a table cell (its child blocks) as inline text
func (r *renderer) renderTableCell(cell *blockNode) string {
	if len(cell.children) == 0 {
		return ""
	}

	var parts []string
	for _, child := range cell.children {
		r.collectCellText(child, &parts)
	}

	result := strings.Join(parts, ", ")
	// Markdown table cells must be single-line; replace any newlines with spaces
	result = strings.ReplaceAll(result, "\n", " ")
	return result
}

// collectCellText recursively extracts inline text from a block node and all
// its descendants, appending results to parts. Handles nested lists, container
// blocks, and non-text leaf blocks inside table cells.
func (r *renderer) collectCellText(node *blockNode, parts *[]string) {
	text := ""
	switch node.block.BlockType {
	// Text blocks
	case 2: // Text
		text = r.renderTextBlock(node.block.Text)
	case 12: // Bullet
		text = "- " + r.renderTextBlock(node.block.Bullet)
	case 13: // Ordered
		text = r.renderTextBlock(node.block.Ordered)

	// Non-text leaf blocks — emit placeholders mirroring renderBlock
	case 23: // File
		if node.block.File != nil {
			name := node.block.File.Name
			if name == "" {
				name = node.block.File.Token
			}
			text = fmt.Sprintf("[file: %s]", name)
		} else {
			text = "[file]"
		}
	case 26: // Iframe
		if node.block.Iframe != nil && node.block.Iframe.Component != nil && node.block.Iframe.Component.URL != "" {
			text = fmt.Sprintf("[embed: %s]", node.block.Iframe.Component.URL)
		} else {
			text = "[embed]"
		}
	case 27: // Image
		if node.block.Image != nil {
			text = fmt.Sprintf("[image: %s]", node.block.Image.Token)
		} else {
			text = "[image]"
		}
	case 30: // Sheet
		if node.block.Sheet != nil && node.block.Sheet.Token != "" {
			text = fmt.Sprintf("[sheet: %s]", node.block.Sheet.Token)
		} else {
			text = "[sheet]"
		}
	case 35: // Task
		if node.block.Task != nil && node.block.Task.TaskID != "" {
			text = fmt.Sprintf("[task: %s]", node.block.Task.TaskID)
		} else {
			text = "[task]"
		}
	case 41: // JiraIssue
		if node.block.JiraIssue != nil {
			key := node.block.JiraIssue.Key
			if key == "" {
				key = node.block.JiraIssue.ID
			}
			if key != "" {
				text = fmt.Sprintf("[jira: %s]", key)
			} else {
				text = "[jira]"
			}
		} else {
			text = "[jira]"
		}
	case 42: // WikiCatalog
		if node.block.WikiCatalog != nil && node.block.WikiCatalog.WikiToken != "" {
			text = fmt.Sprintf("[wiki: %s]", node.block.WikiCatalog.WikiToken)
		} else {
			text = "[wiki]"
		}

	// Container blocks — no own text, just recurse
	case 19, 24, 25, 33, 34: // Callout, Grid, GridColumn, View, QuoteContainer
		// Fall through to recursion below

	default:
		tb := getAnyTextBlock(&node.block)
		if tb != nil {
			text = r.renderTextBlock(tb)
		}
	}

	text = strings.TrimSpace(text)
	if text != "" {
		*parts = append(*parts, text)
	}

	// Recurse into children
	for _, child := range node.children {
		r.collectCellText(child, parts)
	}
}

// MaxInlineSheetRows is the maximum rows to render for embedded sheets in documents.
// Smaller than the standalone sheet read cap (1000) to keep document output readable.
// Exported so that the fetch logic in cmd/doc.go can use the same cap.
const MaxInlineSheetRows = 200

// renderSheetAsTable renders spreadsheet cell values as a markdown table.
func (r *renderer) renderSheetAsTable(sb *strings.Builder, values [][]any) {
	if len(values) == 0 {
		return
	}

	// Determine max column count
	cols := 0
	for _, row := range values {
		if len(row) > cols {
			cols = len(row)
		}
	}
	if cols == 0 {
		return
	}

	// Cap rows for inline rendering
	truncated := 0
	rows := values
	if len(rows) > MaxInlineSheetRows {
		truncated = len(rows) - MaxInlineSheetRows
		rows = rows[:MaxInlineSheetRows]
	}

	for i, row := range rows {
		sb.WriteString("|")
		for c := 0; c < cols; c++ {
			cellContent := ""
			if c < len(row) && row[c] != nil {
				cellContent = stringifySheetCell(row[c])
			}
			// Markdown safety: escape inline markdown, pipes, and flatten newlines
			cellContent = escapeMarkdown(cellContent)
			cellContent = strings.ReplaceAll(cellContent, "|", "\\|")
			cellContent = strings.ReplaceAll(cellContent, "\n", " ")
			sb.WriteString(" ")
			sb.WriteString(cellContent)
			sb.WriteString(" |")
		}
		sb.WriteString("\n")

		// Separator after header row
		if i == 0 {
			sb.WriteString("|")
			for c := 0; c < cols; c++ {
				sb.WriteString(" --- |")
			}
			sb.WriteString("\n")
		}
	}

	if truncated > 0 {
		sb.WriteString(fmt.Sprintf("*[%d more rows truncated]*\n", truncated))
	}
	sb.WriteString("\n")
}

// stringifySheetCell converts a sheet cell value to a plain text string.
// With valueRenderOption=ToString, most cells are strings. This handles
// residual types as fallback.
func stringifySheetCell(v any) string {
	switch val := v.(type) {
	case string:
		return val
	case float64:
		return fmt.Sprintf("%g", val)
	case bool:
		if val {
			return "true"
		}
		return "false"
	case []any:
		// Rich text segments — extract "text" field from each segment
		var parts []string
		for _, seg := range val {
			if m, ok := seg.(map[string]any); ok {
				if text, ok := m["text"].(string); ok {
					parts = append(parts, text)
				}
			}
		}
		return strings.Join(parts, "")
	case map[string]any:
		// Embedded objects (e.g., images)
		if t, ok := val["type"].(string); ok && t == "embed-image" {
			return "[image]"
		}
		if text, ok := val["text"].(string); ok && text != "" {
			return text
		}
		return "[embedded object]"
	default:
		return fmt.Sprint(v)
	}
}

// renderAddOns renders an AddOns block, detecting the type from component_type_id
func (r *renderer) renderAddOns(sb *strings.Builder, block *api.DocumentBlock) {
	if block.AddOns == nil || block.AddOns.Record == "" {
		return
	}

	var record map[string]interface{}
	if err := json.Unmarshal([]byte(block.AddOns.Record), &record); err != nil {
		return
	}

	data, ok := record["data"].(string)
	if !ok || data == "" {
		return
	}

	// Determine fence language based on component type
	lang := "text"
	if isMermaidAddOn(block.AddOns.ComponentTypeID) {
		lang = "mermaid"
	}

	sb.WriteString("```")
	sb.WriteString(lang)
	sb.WriteString("\n")
	sb.WriteString(data)
	if !strings.HasSuffix(data, "\n") {
		sb.WriteString("\n")
	}
	sb.WriteString("```\n\n")
}

// isMermaidAddOn checks if the component type ID corresponds to a Mermaid add-on
func isMermaidAddOn(componentTypeID string) bool {
	// Known Mermaid component type IDs from Lark
	return strings.Contains(strings.ToLower(componentTypeID), "mermaid")
}

// renderTextBlock renders a TextBlock's elements to markdown
func (r *renderer) renderTextBlock(tb *api.TextBlock) string {
	if tb == nil {
		return ""
	}
	return r.renderTextElements(tb.Elements)
}

// renderTextElements converts text elements to markdown
func (r *renderer) renderTextElements(elements []api.TextElement) string {
	var sb strings.Builder
	for _, elem := range elements {
		if elem.TextRun != nil {
			sb.WriteString(r.renderTextRun(elem.TextRun))
		} else if elem.MentionUser != nil {
			sb.WriteString("@")
			if name, ok := r.opts.UserNames[elem.MentionUser.UserID]; ok && name != "" {
				sb.WriteString(name)
			} else {
				sb.WriteString(elem.MentionUser.UserID)
			}
		} else if elem.MentionDoc != nil {
			title := elem.MentionDoc.Title
			if title == "" {
				title = elem.MentionDoc.Token
			}
			if elem.MentionDoc.URL != "" {
				sb.WriteString("[")
				sb.WriteString(title)
				sb.WriteString("](")
				sb.WriteString(decodeURL(elem.MentionDoc.URL))
				sb.WriteString(")")
			} else {
				sb.WriteString(title)
			}
		} else if elem.Equation != nil {
			sb.WriteString("$")
			sb.WriteString(elem.Equation.Content)
			sb.WriteString("$")
		} else if elem.Reminder != nil {
			t := time.UnixMilli(elem.Reminder.ExpireTime).UTC()
			if elem.Reminder.IsWholeDay {
				sb.WriteString(fmt.Sprintf("[reminder: %s]", t.Format("2006-01-02")))
			} else {
				sb.WriteString(fmt.Sprintf("[reminder: %s]", t.Format(time.RFC3339)))
			}
		} else if elem.InlineFile != nil {
			sb.WriteString(fmt.Sprintf("[file: %s]", elem.InlineFile.FileToken))
		} else if elem.InlineBlock != nil {
			sb.WriteString(fmt.Sprintf("[block: %s]", elem.InlineBlock.BlockID))
		}
	}
	result := sb.String()

	// Collapse adjacent identical style markers from consecutive runs with the same style.
	// e.g., **text1****text2** → **text1text2**
	result = strings.ReplaceAll(result, "****", "")
	result = strings.ReplaceAll(result, "~~~~", "")

	return result
}

// renderTextRun renders a single text run with markdown formatting
func (r *renderer) renderTextRun(tr *api.TextRun) string {
	content := tr.Content
	if content == "" {
		return ""
	}

	style := tr.TextElementStyle

	// Skip escaping inside code blocks (content is inside fenced code blocks)
	if r.inCodeBlock {
		if style == nil {
			return content
		}
		return content
	}

	if style == nil {
		return escapeMarkdown(content)
	}

	// Apply inline code (don't combine with other styles, no escaping needed)
	if style.InlineCode {
		return "`" + content + "`"
	}

	// Escape markdown-significant characters in plain text
	content = escapeMarkdown(content)

	// Extract leading/trailing whitespace so style markers stay adjacent to text.
	// Markdown parsers require e.g. **text** not ** text ** to render bold.
	leading := ""
	trailing := ""
	trimmed := strings.TrimRight(content, " \t\n")
	if len(trimmed) < len(content) {
		trailing = content[len(trimmed):]
		content = trimmed
	}
	trimmed = strings.TrimLeft(content, " \t\n")
	if len(trimmed) < len(content) {
		leading = content[:len(content)-len(trimmed)]
		content = trimmed
	}

	// Apply link
	if style.Link != nil && style.Link.URL != "" {
		content = "[" + content + "](" + decodeURL(style.Link.URL) + ")"
	}

	// Apply styles from inside out
	if style.Strikethrough {
		content = "~~" + content + "~~"
	}
	if style.Italic {
		content = "*" + content + "*"
	}
	if style.Bold {
		content = "**" + content + "**"
	}

	return leading + content + trailing
}

// getHeadingTextBlock returns the TextBlock for a heading at the given level
func getHeadingTextBlock(block *api.DocumentBlock, level int) *api.TextBlock {
	switch level {
	case 1:
		return block.Heading1
	case 2:
		return block.Heading2
	case 3:
		return block.Heading3
	case 4:
		return block.Heading4
	case 5:
		return block.Heading5
	case 6:
		return block.Heading6
	case 7:
		return block.Heading7
	case 8:
		return block.Heading8
	case 9:
		return block.Heading9
	default:
		return nil
	}
}

// getAnyTextBlock returns whichever TextBlock field is set on a DocumentBlock
func getAnyTextBlock(block *api.DocumentBlock) *api.TextBlock {
	if block.Text != nil {
		return block.Text
	}
	if block.Bullet != nil {
		return block.Bullet
	}
	if block.Ordered != nil {
		return block.Ordered
	}
	if block.Code != nil {
		return block.Code
	}
	if block.Quote != nil {
		return block.Quote
	}
	if block.TodoBlock != nil {
		return block.TodoBlock
	}
	return getHeadingTextBlock(block, block.BlockType-2)
}

// mdInlineReplacer escapes markdown-significant inline characters
var mdInlineReplacer = strings.NewReplacer(
	`\`, `\\`,
	`*`, `\*`,
	`_`, `\_`,
	"`", "\\`",
	`[`, `\[`,
	`]`, `\]`,
	`~`, `\~`,
)

// urlPattern matches URLs in plain text for preservation during escaping
var urlPattern = regexp.MustCompile(`https?://[^\s<>\[\]` + "`" + `]+`)

// escapeMarkdown escapes inline markdown-significant characters in plain text,
// preserving URLs by wrapping them in angle brackets for reliable autolinking.
func escapeMarkdown(s string) string {
	matches := urlPattern.FindAllStringIndex(s, -1)
	if len(matches) == 0 {
		return mdInlineReplacer.Replace(s)
	}

	var sb strings.Builder
	lastEnd := 0
	for _, m := range matches {
		// Escape text before the URL
		if m[0] > lastEnd {
			sb.WriteString(mdInlineReplacer.Replace(s[lastEnd:m[0]]))
		}
		// Write URL as-is wrapped in angle brackets (markdown autolink)
		sb.WriteString("<")
		sb.WriteString(s[m[0]:m[1]])
		sb.WriteString(">")
		lastEnd = m[1]
	}
	// Escape text after the last URL
	if lastEnd < len(s) {
		sb.WriteString(mdInlineReplacer.Replace(s[lastEnd:]))
	}
	return sb.String()
}

// decodeURL decodes a URL-encoded string from the Lark API.
// The API returns URLs with percent-encoding (e.g., https%3A%2F%2F).
func decodeURL(s string) string {
	decoded, err := url.QueryUnescape(s)
	if err != nil {
		return s
	}
	return decoded
}

// codeFence returns a backtick fence string long enough to safely wrap content.
// If content contains ```, the fence uses more backticks than the longest run found.
func codeFence(content string) string {
	fence := "```"
	for {
		if !strings.Contains(content, fence) {
			return fence
		}
		fence += "`"
	}
}

// Block-level syntax regexes that match markdown triggers at line start (up to 3 leading spaces)
var (
	blockHeadingRe = regexp.MustCompile(`(?m)^( {0,3})(#{1,6}[ \t])`)
	blockQuoteRe   = regexp.MustCompile(`(?m)^( {0,3})(> )`)
	blockListRe    = regexp.MustCompile(`(?m)^( {0,3})([+\-] )`)
	blockOrderedRe = regexp.MustCompile(`(?m)^( {0,3}\d+)(\. )`)
	blockHRuleRe   = regexp.MustCompile(`(?m)^( {0,3})(---)`)
)

// escapeBlockSyntax escapes line-start patterns that would be misinterpreted
// as markdown block syntax. Processes line-by-line.
func escapeBlockSyntax(s string) string {
	s = blockHeadingRe.ReplaceAllString(s, `${1}\${2}`)
	s = blockQuoteRe.ReplaceAllString(s, `${1}\${2}`)
	s = blockListRe.ReplaceAllString(s, `${1}\${2}`)
	s = blockOrderedRe.ReplaceAllString(s, `${1}\${2}`)
	s = blockHRuleRe.ReplaceAllString(s, `${1}\${2}`)
	return s
}

// codeLanguages maps Lark code language IDs to markdown fence language names.
// Complete list from https://open.larksuite.com/document/server-docs/docs/docx-v1/data-structure/block
var codeLanguages = map[int]string{
	1: "plaintext", 2: "abap", 3: "ada", 4: "apache", 5: "apex",
	6: "asm", 7: "bash", 8: "csharp", 9: "cpp", 10: "c",
	11: "cobol", 12: "css", 13: "coffeescript", 14: "d", 15: "dart",
	16: "delphi", 17: "django", 18: "dockerfile", 19: "erlang", 20: "fortran",
	21: "foxpro", 22: "go", 23: "groovy", 24: "html", 25: "htmlbars",
	26: "http", 27: "haskell", 28: "json", 29: "java", 30: "javascript",
	31: "julia", 32: "kotlin", 33: "latex", 34: "lisp", 35: "logo",
	36: "lua", 37: "matlab", 38: "makefile", 39: "markdown", 40: "nginx",
	41: "objectivec", 42: "openedgeabl", 43: "php", 44: "perl", 45: "postscript",
	46: "powershell", 47: "prolog", 48: "protobuf", 49: "python", 50: "r",
	51: "rpg", 52: "ruby", 53: "rust", 54: "sas", 55: "scss",
	56: "sql", 57: "scala", 58: "scheme", 59: "scratch", 60: "shell",
	61: "swift", 62: "thrift", 63: "typescript", 64: "vbscript", 65: "vb",
	66: "xml", 67: "yaml", 68: "cmake", 69: "diff", 70: "gherkin",
	71: "graphql", 72: "glsl", 73: "properties", 74: "solidity", 75: "toml",
}

// codeLanguageName maps a Lark code language ID to a markdown language name
func codeLanguageName(id int) string {
	return codeLanguages[id]
}
