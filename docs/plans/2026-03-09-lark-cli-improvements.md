# lark-cli Document Improvements Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Add table block support, markdown input, quote flag, and missing block type fields to lark-cli.

**Architecture:** Extend the existing types → buildBlocks → CLI flags pipeline. Add a markdown parser that converts markdown text into `[]DocumentBlock` using the same types. Table support requires new struct types (TableBlock, TableCellBlock) and a multi-block creation flow since tables are container blocks with cell children.

**Tech Stack:** Go, Cobra CLI, `github.com/yuin/goldmark` for markdown parsing.

---

### Task 1: Add Table and Quote types to `types.go`

**Files:**
- Modify: `internal/api/types.go:650-685`

**Step 1: Add TableBlock and TableCellBlock structs**

Add after `DividerBlock` struct (line 659), before `DocumentBlock`:

```go
// TableBlock represents a table block (block_type 31)
type TableBlock struct {
	Property *TableProperty `json:"property,omitempty"`
	Cells    []string       `json:"cells,omitempty"`
}

// TableProperty defines table dimensions and layout
type TableProperty struct {
	RowSize      int     `json:"row_size,omitempty"`
	ColumnSize   int     `json:"column_size,omitempty"`
	ColumnWidth  []int   `json:"column_width,omitempty"`
	HeaderRow    bool    `json:"header_row,omitempty"`
	HeaderColumn bool    `json:"header_column,omitempty"`
}

// TableCellBlock represents a table cell (block_type 32) - content is in children
type TableCellBlock struct{}
```

**Step 2: Add Table and TableCell fields to DocumentBlock**

Add after `Image *ImageBlock` (line 684):

```go
	Table     *TableBlock     `json:"table,omitempty"`
	TableCell *TableCellBlock `json:"table_cell,omitempty"`
	Callout   *TextBlock      `json:"callout,omitempty"`
```

**Step 3: Build and verify it compiles**

Run: `cd /Users/yingcong/Code/lark-cli && make build`
Expected: Clean build, no errors.

**Step 4: Commit**

```bash
git add internal/api/types.go
git commit -m "doc: add Table, TableCell, and Callout block types"
```

---

### Task 2: Add `--quote` flag to append/replace commands

**Files:**
- Modify: `internal/cmd/doc.go:736-797` (BlockBuildOpts + buildBlocks)
- Modify: `internal/cmd/doc.go:1728-1764` (flag registration in init())

**Step 1: Add QuoteContent to BlockBuildOpts**

At line 745, add after `AddDivider bool`:

```go
	QuoteContent string
```

**Step 2: Add quote block building in buildBlocks**

At line 793 (after the divider block, before `return blocks`), add:

```go
	if opts.QuoteContent != "" {
		blocks = append(blocks, api.DocumentBlock{BlockType: 15, Quote: mkBlock(opts.QuoteContent)})
	}
```

**Step 3: Wire up --quote flag in docAppendCmd.Run**

In docAppendCmd.Run (line 916), add to the BlockBuildOpts construction:

```go
				QuoteContent:   getStringFlag(cmd, "quote"),
```

**Step 4: Register --quote flag for doc append**

At line 1738, add after the `--todo` flag:

```go
	docAppendCmd.Flags().String("quote", "", "Append a quote block")
```

**Step 5: Register --quote flag for doc replace**

At line 1761, add after the `--todo` flag:

```go
	docReplaceCmd.Flags().String("quote", "", "Replace with quote block")
```

**Step 6: Wire up --quote in docReplaceCmd.Run**

Find the `buildBlocks(BlockBuildOpts{` call in docReplaceCmd and add `QuoteContent: getStringFlag(cmd, "quote"),`.

**Step 7: Update the error message listing content flags**

At line 931, update the error message to include `--quote`:

```go
output.Fatal("MISSING_ARG", fmt.Errorf("at least one content flag is required (--text, --heading, --code, --bullet, --ordered, --todo, --quote, --divider, or --json)"))
```

**Step 8: Build and verify**

Run: `cd /Users/yingcong/Code/lark-cli && make build`
Expected: Clean build.

**Step 9: Test manually**

Run: `lark doc append <test-doc-id> --quote "This is a quote"`
Expected: Quote block appended to document.

**Step 10: Commit**

```bash
git add internal/cmd/doc.go
git commit -m "doc: add --quote flag for append and replace commands"
```

---

### Task 3: Add `--table` flag for creating tables

**Files:**
- Modify: `internal/cmd/doc.go` (BlockBuildOpts, buildBlocks, flag registration)

**Step 1: Add table fields to BlockBuildOpts**

```go
	TableHeader []string // pipe-separated header cells
	TableRows   []string // pipe-separated row cells (repeatable)
```

**Step 2: Add buildTableBlocks helper function**

Add after `buildBlocks` function (after line 797):

```go
// buildTableBlocks creates a table block with header and rows.
// Each row is pipe-separated: "Col1|Col2|Col3"
// Returns the table container block and cell content blocks.
// The Lark API requires creating the table first (which auto-creates cell blocks),
// then updating cells with content via separate API calls.
// For simplicity, we create the table structure and return it.
func buildTableBlocks(headers []string, rows []string) []api.DocumentBlock {
	if len(headers) == 0 {
		return nil
	}
	colSize := len(headers)
	rowSize := 1 + len(rows) // header row + data rows

	// Build cell content: header cells first, then row cells
	var cellBlocks []api.DocumentBlock
	for _, h := range headers {
		cellBlocks = append(cellBlocks, api.DocumentBlock{
			BlockType: 32,
			TableCell: &api.TableCellBlock{},
			Children:  nil, // Will contain text block IDs after creation
		})
		_ = h // Content will be set after table creation
	}
	for _, row := range rows {
		cells := strings.Split(row, "|")
		for i := 0; i < colSize; i++ {
			cellBlocks = append(cellBlocks, api.DocumentBlock{
				BlockType: 32,
				TableCell: &api.TableCellBlock{},
			})
			_ = cells // Content set after creation
			_ = i
		}
	}

	table := api.DocumentBlock{
		BlockType: 31,
		Table: &api.TableBlock{
			Property: &api.TableProperty{
				RowSize:    rowSize,
				ColumnSize: colSize,
				HeaderRow:  true,
			},
		},
	}

	return []api.DocumentBlock{table}
}
```

**Important note:** The Lark API for tables is complex - creating a table block auto-generates cell blocks. Cell content must be added via separate API calls to each cell's children. The initial implementation creates the table structure; cell content population will be a follow-up enhancement.

**Step 3: Wire up in docAppendCmd.Run**

After the `buildBlocks` call (line 927), before the empty blocks check (line 930), add:

```go
		// Build table blocks if --table-header is provided
		tableHeaders, _ := cmd.Flags().GetStringArray("table-header")
		tableRows, _ := cmd.Flags().GetStringArray("table-row")
		if len(tableHeaders) > 0 {
			tableBlocks := buildTableBlocks(tableHeaders, tableRows)
			blocks = append(blocks, tableBlocks...)
		}
```

**Step 4: Register table flags for doc append**

After the `--quote` flag, add:

```go
	docAppendCmd.Flags().StringArray("table-header", nil, "Table header cells (repeatable, one per column)")
	docAppendCmd.Flags().StringArray("table-row", nil, "Table row as pipe-separated cells: \"cell1|cell2|cell3\" (repeatable)")
```

**Step 5: Build and test**

Run: `cd /Users/yingcong/Code/lark-cli && make build`
Test: `lark doc append <test-doc-id> --table-header "Name" --table-header "Status" --table-row "API|Done" --table-row "CLI|WIP"`

**Step 6: Commit**

```bash
git add internal/cmd/doc.go
git commit -m "doc: add --table-header and --table-row flags for table creation"
```

---

### Task 4: Add `--markdown` flag for markdown input

This is the highest-impact feature. It parses markdown text into document blocks.

**Files:**
- Modify: `go.mod` (add goldmark dependency)
- Create: `internal/cmd/markdown.go` (markdown → blocks converter)
- Modify: `internal/cmd/doc.go` (wire up --markdown flag)

**Step 1: Add goldmark dependency**

Run: `cd /Users/yingcong/Code/lark-cli && go get github.com/yuin/goldmark`

**Step 2: Create `internal/cmd/markdown.go`**

```go
package cmd

import (
	"bytes"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
	"github.com/yjwong/lark-cli/internal/api"
)

// parseMarkdownToBlocks converts markdown text into Lark document blocks.
func parseMarkdownToBlocks(source []byte) []api.DocumentBlock {
	md := goldmark.New()
	reader := text.NewReader(source)
	doc := md.Parser().Parse(reader)

	var blocks []api.DocumentBlock
	for node := doc.FirstChild(); node != nil; node = node.NextSibling() {
		blocks = append(blocks, convertNode(node, source)...)
	}
	return blocks
}

// convertNode converts an AST node to document blocks.
func convertNode(node ast.Node, source []byte) []api.DocumentBlock {
	switch n := node.(type) {
	case *ast.Heading:
		content := nodeText(n, source)
		level := n.Level
		if level < 1 {
			level = 1
		}
		if level > 9 {
			level = 9
		}
		block := api.DocumentBlock{BlockType: blockTypeForHeadingLevel(level)}
		setHeadingField(&block, level, makeTextBlock(content))
		return []api.DocumentBlock{block}

	case *ast.Paragraph:
		// Check if parent is a list item - handled by list case
		if _, ok := node.Parent().(*ast.ListItem); ok {
			return nil
		}
		content := nodeText(n, source)
		if content == "" {
			return nil
		}
		return []api.DocumentBlock{{BlockType: 2, Text: makeTextBlock(content)}}

	case *ast.FencedCodeBlock:
		var buf bytes.Buffer
		lines := n.Lines()
		for i := 0; i < lines.Len(); i++ {
			line := lines.At(i)
			buf.Write(line.Value(source))
		}
		content := strings.TrimRight(buf.String(), "\n")
		tb := makeTextBlock(content)
		lang := string(n.Language(source))
		langID := markdownLangToID(lang)
		if langID > 0 {
			tb.Style = &api.TextStyle{Language: langID}
		}
		return []api.DocumentBlock{{BlockType: 14, Code: tb}}

	case *ast.CodeBlock:
		var buf bytes.Buffer
		lines := n.Lines()
		for i := 0; i < lines.Len(); i++ {
			line := lines.At(i)
			buf.Write(line.Value(source))
		}
		content := strings.TrimRight(buf.String(), "\n")
		return []api.DocumentBlock{{BlockType: 14, Code: makeTextBlock(content)}}

	case *ast.List:
		var blocks []api.DocumentBlock
		blockType := 12 // bullet
		if n.IsOrdered() {
			blockType = 13 // ordered
		}
		for child := n.FirstChild(); child != nil; child = child.NextSibling() {
			if li, ok := child.(*ast.ListItem); ok {
				content := nodeText(li, source)
				tb := makeTextBlock(content)
				block := api.DocumentBlock{BlockType: blockType}
				if blockType == 12 {
					block.Bullet = tb
				} else {
					block.Ordered = tb
				}
				blocks = append(blocks, block)
			}
		}
		return blocks

	case *ast.Blockquote:
		// Extract text from all children
		var parts []string
		for child := n.FirstChild(); child != nil; child = child.NextSibling() {
			text := nodeText(child, source)
			if text != "" {
				parts = append(parts, text)
			}
		}
		content := strings.Join(parts, "\n")
		return []api.DocumentBlock{{BlockType: 15, Quote: makeTextBlock(content)}}

	case *ast.ThematicBreak:
		return []api.DocumentBlock{{BlockType: 22, Divider: &api.DividerBlock{}}}

	default:
		return nil
	}
}

// nodeText extracts text content from an AST node.
func nodeText(node ast.Node, source []byte) string {
	var buf bytes.Buffer
	for child := node.FirstChild(); child != nil; child = child.NextSibling() {
		switch c := child.(type) {
		case *ast.Text:
			buf.Write(c.Segment.Value(source))
			if c.HardLineBreak() || c.SoftLineBreak() {
				buf.WriteByte('\n')
			}
		case *ast.CodeSpan:
			// Extract inline code text
			for gc := c.FirstChild(); gc != nil; gc = gc.NextSibling() {
				if t, ok := gc.(*ast.Text); ok {
					buf.Write(t.Segment.Value(source))
				}
			}
		case *ast.Emphasis:
			buf.WriteString(nodeText(c, source))
		case *ast.Link:
			buf.WriteString(nodeText(c, source))
		case *ast.AutoLink:
			buf.Write(c.URL(source))
		default:
			if c.HasChildren() {
				buf.WriteString(nodeText(c, source))
			}
		}
	}
	return buf.String()
}

// markdownLangToID maps common markdown language names to Lark code language IDs.
func markdownLangToID(lang string) int {
	lang = strings.ToLower(strings.TrimSpace(lang))
	switch lang {
	case "bash", "sh", "shell":
		return 7
	case "c":
		return 9
	case "cpp", "c++":
		return 10
	case "csharp", "c#", "cs":
		return 12
	case "css":
		return 13
	case "go", "golang":
		return 22
	case "html":
		return 25
	case "java":
		return 29
	case "javascript", "js":
		return 30
	case "json":
		return 31
	case "kotlin":
		return 32
	case "markdown", "md":
		return 35
	case "php":
		return 42
	case "python", "py":
		return 49
	case "ruby", "rb":
		return 51
	case "rust", "rs":
		return 53
	case "scala":
		return 54
	case "sql":
		return 56
	case "swift":
		return 58
	case "typescript", "ts":
		return 63
	case "xml":
		return 66
	case "yaml", "yml":
		return 67
	default:
		return 1 // PlainText
	}
}
```

**Step 3: Wire up `--markdown` flag in docAppendCmd.Run**

In `docAppendCmd.Run` (around line 906), after the `useJSON` check but before `buildBlocks`, add a `--markdown` path:

```go
		useMarkdown, _ := cmd.Flags().GetBool("markdown")

		var blocks []api.DocumentBlock
		if useJSON {
			var err error
			blocks, err = readBlocksFromStdin()
			if err != nil {
				output.Fatal("PARSE_ERROR", err)
			}
		} else if useMarkdown {
			data, err := io.ReadAll(os.Stdin)
			if err != nil {
				output.Fatal("READ_ERROR", err)
			}
			blocks = parseMarkdownToBlocks(data)
		} else {
```

**Step 4: Register --markdown flag for doc append**

```go
	docAppendCmd.Flags().Bool("markdown", false, "Read markdown from stdin and convert to blocks")
```

**Step 5: Register --markdown flag for doc replace**

```go
	docReplaceCmd.Flags().Bool("markdown", false, "Read markdown from stdin and convert to blocks")
```

**Step 6: Wire up --markdown in docReplaceCmd.Run**

Add the same `useMarkdown` path to docReplaceCmd's block reading logic.

**Step 7: Build and verify**

Run: `cd /Users/yingcong/Code/lark-cli && make build`
Expected: Clean build.

**Step 8: Test manually**

```bash
echo '# Hello World

This is a paragraph.

- bullet 1
- bullet 2

1. step one
2. step two

> A quote

---

```python
print("hello")
```' | lark doc append <test-doc-id> --markdown
```

Expected: All block types created correctly in the document.

**Step 9: Commit**

```bash
git add go.mod go.sum internal/cmd/markdown.go internal/cmd/doc.go
git commit -m "doc: add --markdown flag to parse markdown input into blocks"
```

---

### Task 5: Update doc replace with quote and markdown support

**Files:**
- Modify: `internal/cmd/doc.go` (docReplaceCmd)

Find `docReplaceCmd` and ensure it has the same `QuoteContent` and `useMarkdown` support as `docAppendCmd`. The replace command uses `buildBlocks` too, so it just needs the opts wired up.

**Step 1: Find the docReplaceCmd buildBlocks call and add QuoteContent**

**Step 2: Add markdown stdin path to docReplaceCmd**

**Step 3: Build and test**

**Step 4: Commit**

```bash
git add internal/cmd/doc.go
git commit -m "doc: add --quote and --markdown support to replace command"
```

---

### Task 6: Update skill docs

**Files:**
- Modify: `skills/documents/SKILL.md`

**Step 1: Add --quote to the append command documentation**

In the "Append Content to a Document" section, add:
```
- `--quote "content"`: Append a quote/blockquote
```

**Step 2: Add --markdown to the append command documentation**

```
- `--markdown`: Read markdown from stdin and convert to blocks
```

Add examples:
```bash
# Append from markdown
echo '# Title\n\nParagraph text\n\n- item 1\n- item 2' | lark doc append ABC123xyz --markdown

# Pipe a file
cat content.md | lark doc append ABC123xyz --markdown
```

**Step 3: Add --table-header/--table-row documentation**

```bash
# Create a table
lark doc append ABC123xyz --table-header "Name" --table-header "Status" --table-row "API|Done" --table-row "CLI|WIP"
```

**Step 4: Commit**

```bash
git add skills/documents/SKILL.md
git commit -m "doc: update skill docs with quote, markdown, and table flags"
```

---

## Summary of Changes

| Task | Files | Description |
|------|-------|-------------|
| 1 | types.go | Add TableBlock, TableCellBlock, TableProperty structs; add Table/TableCell/Callout fields to DocumentBlock |
| 2 | doc.go | Add --quote flag for append and replace |
| 3 | doc.go | Add --table-header and --table-row flags for table creation |
| 4 | markdown.go, doc.go, go.mod | Add --markdown flag with goldmark parser |
| 5 | doc.go | Wire quote + markdown into replace command |
| 6 | SKILL.md | Update documentation |
