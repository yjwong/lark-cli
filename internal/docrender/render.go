package docrender

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/yjwong/lark-cli/internal/api"
)

// blockNode wraps a DocumentBlock with resolved children for tree traversal
type blockNode struct {
	block    api.DocumentBlock
	children []*blockNode
}

// RenderBlocks converts a flat list of document blocks into markdown
func RenderBlocks(blocks []api.DocumentBlock) string {
	if len(blocks) == 0 {
		return ""
	}

	tree := buildBlockTree(blocks)
	if tree == nil {
		return ""
	}

	var sb strings.Builder
	renderChildren(&sb, tree, 0)
	return strings.TrimRight(sb.String(), "\n") + "\n"
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
func renderChildren(sb *strings.Builder, node *blockNode, depth int) {
	orderedCounter := 0
	for _, child := range node.children {
		if child.block.BlockType == 13 { // ordered list
			orderedCounter++
		} else {
			orderedCounter = 0
		}
		renderBlock(sb, child, depth, orderedCounter)
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
func renderBlock(sb *strings.Builder, node *blockNode, depth int, orderedNum int) {
	switch node.block.BlockType {
	case 1: // Page
		renderChildren(sb, node, depth)

	case 2: // Text
		text := renderTextBlock(node.block.Text)
		if text != "" {
			sb.WriteString(text)
		}
		sb.WriteString("\n")

	case 3, 4, 5, 6, 7, 8, 9, 10, 11: // Headings H1-H9
		level := node.block.BlockType - 2
		tb := getHeadingTextBlock(&node.block, level)
		text := renderTextBlock(tb)
		sb.WriteString(strings.Repeat("#", level))
		sb.WriteString(" ")
		sb.WriteString(text)
		sb.WriteString("\n\n")

	case 12: // Bullet
		text := renderTextBlock(node.block.Bullet)
		sb.WriteString(indentPrefix(depth))
		sb.WriteString("- ")
		sb.WriteString(text)
		sb.WriteString("\n")
		if len(node.children) > 0 {
			renderChildren(sb, node, depth+1)
		}

	case 13: // Ordered
		text := renderTextBlock(node.block.Ordered)
		sb.WriteString(indentPrefix(depth))
		sb.WriteString(fmt.Sprintf("%d. ", orderedNum))
		sb.WriteString(text)
		sb.WriteString("\n")
		if len(node.children) > 0 {
			renderChildren(sb, node, depth+1)
		}

	case 14: // Code
		if node.block.Code != nil {
			lang := ""
			if node.block.Code.Style != nil {
				lang = codeLanguageName(node.block.Code.Style.Language)
			}
			content := renderTextBlock(node.block.Code)
			sb.WriteString("```")
			sb.WriteString(lang)
			sb.WriteString("\n")
			sb.WriteString(content)
			sb.WriteString("\n```\n\n")
		}

	case 15: // Quote
		text := renderTextBlock(node.block.Quote)
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
			text := renderTextBlock(node.block.TodoBlock)
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
			renderChildren(&calloutSB, node, depth)
			for _, line := range strings.Split(strings.TrimRight(calloutSB.String(), "\n"), "\n") {
				sb.WriteString("> ")
				sb.WriteString(line)
				sb.WriteString("\n")
			}
			sb.WriteString("\n")
		}

	case 22: // Divider
		sb.WriteString("---\n\n")

	case 27: // Image
		if node.block.Image != nil {
			sb.WriteString(fmt.Sprintf("[image: %s]\n\n", node.block.Image.Token))
		}

	case 31: // Table
		renderTable(sb, node)

	case 32: // TableCell — handled by renderTable
		// Skip: rendered as part of parent table

	case 40: // AddOns
		renderAddOns(sb, &node.block)

	default:
		// Warn about unsupported block types to stderr
		fmt.Fprintf(os.Stderr, "warning: unsupported block type %d (block_id: %s), use --raw for full content\n",
			node.block.BlockType, node.block.BlockID)
	}
}

// renderTable renders a table block as a markdown table
func renderTable(sb *strings.Builder, node *blockNode) {
	if node.block.Table == nil || node.block.Table.Property == nil {
		return
	}

	prop := node.block.Table.Property
	rows := prop.RowSize
	cols := prop.ColumnSize

	if rows == 0 || cols == 0 || len(node.children) == 0 {
		return
	}

	// Children are table cells in row-major order
	cells := node.children

	// Check if table has a header row
	hasHeader := prop.HeaderRow

	for r := 0; r < rows; r++ {
		sb.WriteString("|")
		for c := 0; c < cols; c++ {
			idx := r*cols + c
			cellContent := ""
			if idx < len(cells) {
				cellContent = renderTableCell(cells[idx])
			}
			// Escape pipes in cell content
			cellContent = strings.ReplaceAll(cellContent, "|", "\\|")
			sb.WriteString(" ")
			sb.WriteString(cellContent)
			sb.WriteString(" |")
		}
		sb.WriteString("\n")

		// Add separator after first row (header or not — required by markdown)
		if r == 0 {
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
func renderTableCell(cell *blockNode) string {
	if len(cell.children) == 0 {
		return ""
	}

	var parts []string
	for _, child := range cell.children {
		text := ""
		switch child.block.BlockType {
		case 2: // Text
			text = renderTextBlock(child.block.Text)
		case 12: // Bullet in cell
			text = "- " + renderTextBlock(child.block.Bullet)
		case 13: // Ordered in cell
			text = renderTextBlock(child.block.Ordered)
		default:
			tb := getAnyTextBlock(&child.block)
			if tb != nil {
				text = renderTextBlock(tb)
			}
		}
		text = strings.TrimSpace(text)
		if text != "" {
			parts = append(parts, text)
		}
	}

	return strings.Join(parts, ", ")
}

// renderAddOns renders an AddOns block, detecting the type from component_type_id
func renderAddOns(sb *strings.Builder, block *api.DocumentBlock) {
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
func renderTextBlock(tb *api.TextBlock) string {
	if tb == nil {
		return ""
	}
	return renderTextElements(tb.Elements)
}

// renderTextElements converts text elements to markdown
func renderTextElements(elements []api.TextElement) string {
	var sb strings.Builder
	for _, elem := range elements {
		if elem.TextRun != nil {
			sb.WriteString(renderTextRun(elem.TextRun))
		} else if elem.MentionUser != nil {
			sb.WriteString("@")
			sb.WriteString(elem.MentionUser.UserID)
		} else if elem.MentionDoc != nil {
			title := elem.MentionDoc.Title
			if title == "" {
				title = elem.MentionDoc.Token
			}
			if elem.MentionDoc.URL != "" {
				sb.WriteString("[")
				sb.WriteString(title)
				sb.WriteString("](")
				sb.WriteString(elem.MentionDoc.URL)
				sb.WriteString(")")
			} else {
				sb.WriteString(title)
			}
		}
	}
	return sb.String()
}

// renderTextRun renders a single text run with markdown formatting
func renderTextRun(tr *api.TextRun) string {
	content := tr.Content
	if content == "" {
		return ""
	}

	style := tr.TextElementStyle
	if style == nil {
		return content
	}

	// Apply link
	if style.Link != nil && style.Link.URL != "" {
		content = "[" + content + "](" + style.Link.URL + ")"
	}

	// Apply inline code (don't combine with other styles)
	if style.InlineCode {
		return "`" + content + "`"
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

	return content
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

// codeLanguageName maps Lark code language IDs to markdown language names
func codeLanguageName(id int) string {
	switch id {
	case 1:
		return "plaintext"
	case 7:
		return "bash"
	case 9:
		return "cpp"
	case 10:
		return "c"
	case 12:
		return "css"
	case 22:
		return "go"
	case 24:
		return "html"
	case 28:
		return "json"
	case 29:
		return "java"
	case 30:
		return "javascript"
	case 32:
		return "kotlin"
	case 49:
		return "python"
	case 52:
		return "ruby"
	case 53:
		return "rust"
	case 56:
		return "sql"
	case 58:
		return "swift"
	case 63:
		return "typescript"
	case 67:
		return "yaml"
	default:
		return ""
	}
}
