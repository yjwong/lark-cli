package docrender

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/yjwong/lark-cli/internal/api"
)

func TestRenderBlocks_EmptyInput(t *testing.T) {
	result := RenderBlocks(nil)
	if result != "" {
		t.Errorf("expected empty string, got %q", result)
	}
}

func TestRenderBlocks_TextBlock(t *testing.T) {
	blocks := []api.DocumentBlock{
		{BlockID: "page", BlockType: 1, Children: []string{"t1"}},
		{BlockID: "t1", ParentID: "page", BlockType: 2, Text: &api.TextBlock{
			Elements: []api.TextElement{
				{TextRun: &api.TextRun{Content: "Hello world"}},
			},
		}},
	}
	result := RenderBlocks(blocks)
	if !strings.Contains(result, "Hello world") {
		t.Errorf("expected 'Hello world' in output, got %q", result)
	}
}

func TestRenderBlocks_TextWithStyles(t *testing.T) {
	blocks := []api.DocumentBlock{
		{BlockID: "page", BlockType: 1, Children: []string{"t1"}},
		{BlockID: "t1", ParentID: "page", BlockType: 2, Text: &api.TextBlock{
			Elements: []api.TextElement{
				{TextRun: &api.TextRun{Content: "bold", TextElementStyle: &api.TextElementStyle{Bold: true}}},
				{TextRun: &api.TextRun{Content: " "}},
				{TextRun: &api.TextRun{Content: "italic", TextElementStyle: &api.TextElementStyle{Italic: true}}},
				{TextRun: &api.TextRun{Content: " "}},
				{TextRun: &api.TextRun{Content: "struck", TextElementStyle: &api.TextElementStyle{Strikethrough: true}}},
				{TextRun: &api.TextRun{Content: " "}},
				{TextRun: &api.TextRun{Content: "code", TextElementStyle: &api.TextElementStyle{InlineCode: true}}},
			},
		}},
	}
	result := RenderBlocks(blocks)
	if !strings.Contains(result, "**bold**") {
		t.Errorf("expected **bold**, got %q", result)
	}
	if !strings.Contains(result, "*italic*") {
		t.Errorf("expected *italic*, got %q", result)
	}
	if !strings.Contains(result, "~~struck~~") {
		t.Errorf("expected ~~struck~~, got %q", result)
	}
	if !strings.Contains(result, "`code`") {
		t.Errorf("expected `code`, got %q", result)
	}
}

func TestRenderBlocks_TextWithLink(t *testing.T) {
	blocks := []api.DocumentBlock{
		{BlockID: "page", BlockType: 1, Children: []string{"t1"}},
		{BlockID: "t1", ParentID: "page", BlockType: 2, Text: &api.TextBlock{
			Elements: []api.TextElement{
				{TextRun: &api.TextRun{
					Content:          "click here",
					TextElementStyle: &api.TextElementStyle{Link: &api.TextLink{URL: "https://example.com"}},
				}},
			},
		}},
	}
	result := RenderBlocks(blocks)
	if !strings.Contains(result, "[click here](https://example.com)") {
		t.Errorf("expected link markdown, got %q", result)
	}
}

func TestRenderBlocks_Headings(t *testing.T) {
	blocks := []api.DocumentBlock{
		{BlockID: "page", BlockType: 1, Children: []string{"h1", "h2", "h3"}},
		{BlockID: "h1", ParentID: "page", BlockType: 3, Heading1: &api.TextBlock{
			Elements: []api.TextElement{{TextRun: &api.TextRun{Content: "Title"}}},
		}},
		{BlockID: "h2", ParentID: "page", BlockType: 4, Heading2: &api.TextBlock{
			Elements: []api.TextElement{{TextRun: &api.TextRun{Content: "Section"}}},
		}},
		{BlockID: "h3", ParentID: "page", BlockType: 5, Heading3: &api.TextBlock{
			Elements: []api.TextElement{{TextRun: &api.TextRun{Content: "Subsection"}}},
		}},
	}
	result := RenderBlocks(blocks)
	if !strings.Contains(result, "# Title") {
		t.Errorf("expected '# Title', got %q", result)
	}
	if !strings.Contains(result, "## Section") {
		t.Errorf("expected '## Section', got %q", result)
	}
	if !strings.Contains(result, "### Subsection") {
		t.Errorf("expected '### Subsection', got %q", result)
	}
}

func TestRenderBlocks_BulletList(t *testing.T) {
	blocks := []api.DocumentBlock{
		{BlockID: "page", BlockType: 1, Children: []string{"b1", "b2"}},
		{BlockID: "b1", ParentID: "page", BlockType: 12, Bullet: &api.TextBlock{
			Elements: []api.TextElement{{TextRun: &api.TextRun{Content: "First"}}},
		}},
		{BlockID: "b2", ParentID: "page", BlockType: 12, Bullet: &api.TextBlock{
			Elements: []api.TextElement{{TextRun: &api.TextRun{Content: "Second"}}},
		}},
	}
	result := RenderBlocks(blocks)
	if !strings.Contains(result, "- First\n- Second") {
		t.Errorf("expected bullet list, got %q", result)
	}
}

func TestRenderBlocks_OrderedList(t *testing.T) {
	blocks := []api.DocumentBlock{
		{BlockID: "page", BlockType: 1, Children: []string{"o1", "o2", "o3"}},
		{BlockID: "o1", ParentID: "page", BlockType: 13, Ordered: &api.TextBlock{
			Elements: []api.TextElement{{TextRun: &api.TextRun{Content: "First"}}},
		}},
		{BlockID: "o2", ParentID: "page", BlockType: 13, Ordered: &api.TextBlock{
			Elements: []api.TextElement{{TextRun: &api.TextRun{Content: "Second"}}},
		}},
		{BlockID: "o3", ParentID: "page", BlockType: 13, Ordered: &api.TextBlock{
			Elements: []api.TextElement{{TextRun: &api.TextRun{Content: "Third"}}},
		}},
	}
	result := RenderBlocks(blocks)
	if !strings.Contains(result, "1. First\n2. Second\n3. Third") {
		t.Errorf("expected ordered list, got %q", result)
	}
}

func TestRenderBlocks_CodeBlock(t *testing.T) {
	blocks := []api.DocumentBlock{
		{BlockID: "page", BlockType: 1, Children: []string{"c1"}},
		{BlockID: "c1", ParentID: "page", BlockType: 14, Code: &api.TextBlock{
			Style:    &api.TextStyle{Language: 49},
			Elements: []api.TextElement{{TextRun: &api.TextRun{Content: "print('hello')"}}},
		}},
	}
	result := RenderBlocks(blocks)
	if !strings.Contains(result, "```python\nprint('hello')\n```") {
		t.Errorf("expected python code block, got %q", result)
	}
}

func TestRenderBlocks_CodeBlockWithTripleBackticks(t *testing.T) {
	blocks := []api.DocumentBlock{
		{BlockID: "page", BlockType: 1, Children: []string{"c1"}},
		{BlockID: "c1", ParentID: "page", BlockType: 14, Code: &api.TextBlock{
			Style:    &api.TextStyle{Language: 39}, // markdown
			Elements: []api.TextElement{{TextRun: &api.TextRun{Content: "```go\nfmt.Println()\n```"}}},
		}},
	}
	result := RenderBlocks(blocks)
	// Should use ```` (4 backticks) to wrap content containing ```
	if !strings.Contains(result, "````") {
		t.Errorf("expected 4+ backtick fence for content with triple backticks, got %q", result)
	}
	if !strings.Contains(result, "```go\nfmt.Println()\n```") {
		t.Errorf("expected content preserved verbatim, got %q", result)
	}
}

func TestRenderBlocks_HeadingLevelClamped(t *testing.T) {
	blocks := []api.DocumentBlock{
		{BlockID: "page", BlockType: 1, Children: []string{"h7"}},
		{BlockID: "h7", ParentID: "page", BlockType: 9, Heading7: &api.TextBlock{
			Elements: []api.TextElement{{TextRun: &api.TextRun{Content: "Deep heading"}}},
		}},
	}
	result := RenderBlocks(blocks)
	// H7 (block_type 9) should be clamped to ###### (6)
	if !strings.Contains(result, "###### Deep heading") {
		t.Errorf("expected H7 clamped to 6 #'s, got %q", result)
	}
	if strings.Contains(result, "####### ") {
		t.Errorf("should not have 7+ #'s, got %q", result)
	}
}

func TestCodeFence(t *testing.T) {
	if f := codeFence("normal code"); f != "```" {
		t.Errorf("expected ```, got %q", f)
	}
	if f := codeFence("has ``` backticks"); f != "````" {
		t.Errorf("expected ````, got %q", f)
	}
	if f := codeFence("has ```` four"); f != "`````" {
		t.Errorf("expected `````, got %q", f)
	}
}

func TestRenderBlocks_Todo(t *testing.T) {
	blocks := []api.DocumentBlock{
		{BlockID: "page", BlockType: 1, Children: []string{"td1", "td2"}},
		{BlockID: "td1", ParentID: "page", BlockType: 17, TodoBlock: &api.TextBlock{
			Style:    &api.TextStyle{Done: false},
			Elements: []api.TextElement{{TextRun: &api.TextRun{Content: "Pending"}}},
		}},
		{BlockID: "td2", ParentID: "page", BlockType: 17, TodoBlock: &api.TextBlock{
			Style:    &api.TextStyle{Done: true},
			Elements: []api.TextElement{{TextRun: &api.TextRun{Content: "Done"}}},
		}},
	}
	result := RenderBlocks(blocks)
	if !strings.Contains(result, "- [ ] Pending") {
		t.Errorf("expected unchecked todo, got %q", result)
	}
	if !strings.Contains(result, "- [x] Done") {
		t.Errorf("expected checked todo, got %q", result)
	}
}

func TestRenderBlocks_Divider(t *testing.T) {
	blocks := []api.DocumentBlock{
		{BlockID: "page", BlockType: 1, Children: []string{"d1"}},
		{BlockID: "d1", ParentID: "page", BlockType: 22, Divider: &api.DividerBlock{}},
	}
	result := RenderBlocks(blocks)
	if !strings.Contains(result, "---") {
		t.Errorf("expected divider, got %q", result)
	}
}

func TestRenderBlocks_Table(t *testing.T) {
	blocks := []api.DocumentBlock{
		{BlockID: "page", BlockType: 1, Children: []string{"tbl"}},
		{BlockID: "tbl", ParentID: "page", BlockType: 31,
			Children: []string{"c1", "c2", "c3", "c4"},
			Table: &api.TableBlock{
				Property: &api.TableProperty{RowSize: 2, ColumnSize: 2},
			},
		},
		// Header row
		{BlockID: "c1", ParentID: "tbl", BlockType: 32, Children: []string{"c1t"}},
		{BlockID: "c1t", ParentID: "c1", BlockType: 2, Text: &api.TextBlock{
			Elements: []api.TextElement{{TextRun: &api.TextRun{Content: "Name"}}},
		}},
		{BlockID: "c2", ParentID: "tbl", BlockType: 32, Children: []string{"c2t"}},
		{BlockID: "c2t", ParentID: "c2", BlockType: 2, Text: &api.TextBlock{
			Elements: []api.TextElement{{TextRun: &api.TextRun{Content: "Value"}}},
		}},
		// Data row
		{BlockID: "c3", ParentID: "tbl", BlockType: 32, Children: []string{"c3t"}},
		{BlockID: "c3t", ParentID: "c3", BlockType: 2, Text: &api.TextBlock{
			Elements: []api.TextElement{{TextRun: &api.TextRun{Content: "foo"}}},
		}},
		{BlockID: "c4", ParentID: "tbl", BlockType: 32, Children: []string{"c4t"}},
		{BlockID: "c4t", ParentID: "c4", BlockType: 2, Text: &api.TextBlock{
			Elements: []api.TextElement{{TextRun: &api.TextRun{Content: "bar"}}},
		}},
	}
	result := RenderBlocks(blocks)
	if !strings.Contains(result, "| Name | Value |") {
		t.Errorf("expected table header, got %q", result)
	}
	if !strings.Contains(result, "| --- | --- |") {
		t.Errorf("expected table separator, got %q", result)
	}
	if !strings.Contains(result, "| foo | bar |") {
		t.Errorf("expected table data row, got %q", result)
	}
}

func TestRenderBlocks_TableWithMultiLineCells(t *testing.T) {
	blocks := []api.DocumentBlock{
		{BlockID: "page", BlockType: 1, Children: []string{"tbl"}},
		{BlockID: "tbl", ParentID: "page", BlockType: 31,
			Children: []string{"c1", "c2"},
			Table: &api.TableBlock{
				Property: &api.TableProperty{RowSize: 1, ColumnSize: 2},
			},
		},
		{BlockID: "c1", ParentID: "tbl", BlockType: 32, Children: []string{"c1t1", "c1t2"}},
		{BlockID: "c1t1", ParentID: "c1", BlockType: 2, Text: &api.TextBlock{
			Elements: []api.TextElement{{TextRun: &api.TextRun{Content: "Line 1"}}},
		}},
		{BlockID: "c1t2", ParentID: "c1", BlockType: 2, Text: &api.TextBlock{
			Elements: []api.TextElement{{TextRun: &api.TextRun{Content: "Line 2"}}},
		}},
		{BlockID: "c2", ParentID: "tbl", BlockType: 32, Children: []string{"c2t"}},
		{BlockID: "c2t", ParentID: "c2", BlockType: 2, Text: &api.TextBlock{
			Elements: []api.TextElement{{TextRun: &api.TextRun{Content: "Single"}}},
		}},
	}
	result := RenderBlocks(blocks)
	if !strings.Contains(result, "Line 1, Line 2") {
		t.Errorf("expected multi-line cell joined with ', ', got %q", result)
	}
}

func TestRenderBlocks_AddOnsMermaid(t *testing.T) {
	blocks := []api.DocumentBlock{
		{BlockID: "page", BlockType: 1, Children: []string{"a1"}},
		{BlockID: "a1", ParentID: "page", BlockType: 40, AddOns: &api.AddOnsBlock{
			ComponentTypeID: "mermaid",
			Record:          `{"data":"graph TD\n    A-->B","theme":"default"}`,
		}},
	}
	result := RenderBlocks(blocks)
	if !strings.Contains(result, "```mermaid\ngraph TD\n    A-->B\n```") {
		t.Errorf("expected mermaid code block, got %q", result)
	}
}

func TestRenderBlocks_AddOnsNonMermaid(t *testing.T) {
	blocks := []api.DocumentBlock{
		{BlockID: "page", BlockType: 1, Children: []string{"a1"}},
		{BlockID: "a1", ParentID: "page", BlockType: 40, AddOns: &api.AddOnsBlock{
			ComponentTypeID: "survey",
			Record:          `{"data":"some survey data"}`,
		}},
	}
	result := RenderBlocks(blocks)
	if strings.Contains(result, "mermaid") {
		t.Errorf("non-mermaid add-on should not be labeled as mermaid, got %q", result)
	}
	if !strings.Contains(result, "```text") {
		t.Errorf("expected text code fence for non-mermaid add-on, got %q", result)
	}
}

func TestRenderBlocks_MentionUser(t *testing.T) {
	blocks := []api.DocumentBlock{
		{BlockID: "page", BlockType: 1, Children: []string{"t1"}},
		{BlockID: "t1", ParentID: "page", BlockType: 2, Text: &api.TextBlock{
			Elements: []api.TextElement{
				{TextRun: &api.TextRun{Content: "Hello "}},
				{MentionUser: &api.MentionUser{UserID: "ou_abc123"}},
			},
		}},
	}
	result := RenderBlocks(blocks)
	if !strings.Contains(result, "@ou_abc123") {
		t.Errorf("expected @mention, got %q", result)
	}
}

func TestRenderBlocks_UnknownBlocksSilentlySkipped(t *testing.T) {
	blocks := []api.DocumentBlock{
		{BlockID: "page", BlockType: 1, Children: []string{"t1", "unknown", "t2"}},
		{BlockID: "t1", ParentID: "page", BlockType: 2, Text: &api.TextBlock{
			Elements: []api.TextElement{{TextRun: &api.TextRun{Content: "Before"}}},
		}},
		{BlockID: "unknown", ParentID: "page", BlockType: 999},
		{BlockID: "t2", ParentID: "page", BlockType: 2, Text: &api.TextBlock{
			Elements: []api.TextElement{{TextRun: &api.TextRun{Content: "After"}}},
		}},
	}
	result := RenderBlocks(blocks)
	if strings.Contains(result, "<!--") {
		t.Errorf("expected no HTML comments for unknown blocks, got %q", result)
	}
	if !strings.Contains(result, "Before") || !strings.Contains(result, "After") {
		t.Errorf("expected surrounding text preserved, got %q", result)
	}
}

func TestRenderBlocks_MixedDocument(t *testing.T) {
	blocks := []api.DocumentBlock{
		{BlockID: "page", BlockType: 1, Children: []string{"h1", "t1", "b1", "b2", "div", "t2"}},
		{BlockID: "h1", ParentID: "page", BlockType: 3, Heading1: &api.TextBlock{
			Elements: []api.TextElement{{TextRun: &api.TextRun{Content: "Title"}}},
		}},
		{BlockID: "t1", ParentID: "page", BlockType: 2, Text: &api.TextBlock{
			Elements: []api.TextElement{{TextRun: &api.TextRun{Content: "Some text"}}},
		}},
		{BlockID: "b1", ParentID: "page", BlockType: 12, Bullet: &api.TextBlock{
			Elements: []api.TextElement{{TextRun: &api.TextRun{Content: "Item 1"}}},
		}},
		{BlockID: "b2", ParentID: "page", BlockType: 12, Bullet: &api.TextBlock{
			Elements: []api.TextElement{{TextRun: &api.TextRun{Content: "Item 2"}}},
		}},
		{BlockID: "div", ParentID: "page", BlockType: 22, Divider: &api.DividerBlock{}},
		{BlockID: "t2", ParentID: "page", BlockType: 2, Text: &api.TextBlock{
			Elements: []api.TextElement{{TextRun: &api.TextRun{Content: "After divider"}}},
		}},
	}
	result := RenderBlocks(blocks)
	expected := "# Title\n\nSome text\n- Item 1\n- Item 2\n---\n\nAfter divider\n"
	if result != expected {
		t.Errorf("expected:\n%q\ngot:\n%q", expected, result)
	}
}

func TestRenderBlocks_PipeInTableCell(t *testing.T) {
	blocks := []api.DocumentBlock{
		{BlockID: "page", BlockType: 1, Children: []string{"tbl"}},
		{BlockID: "tbl", ParentID: "page", BlockType: 31,
			Children: []string{"c1", "c2"},
			Table: &api.TableBlock{
				Property: &api.TableProperty{RowSize: 1, ColumnSize: 2},
			},
		},
		{BlockID: "c1", ParentID: "tbl", BlockType: 32, Children: []string{"c1t"}},
		{BlockID: "c1t", ParentID: "c1", BlockType: 2, Text: &api.TextBlock{
			Elements: []api.TextElement{{TextRun: &api.TextRun{Content: "A | B"}}},
		}},
		{BlockID: "c2", ParentID: "tbl", BlockType: 32, Children: []string{"c2t"}},
		{BlockID: "c2t", ParentID: "c2", BlockType: 2, Text: &api.TextBlock{
			Elements: []api.TextElement{{TextRun: &api.TextRun{Content: "C"}}},
		}},
	}
	result := RenderBlocks(blocks)
	if !strings.Contains(result, `A \| B`) {
		t.Errorf("expected escaped pipe in cell, got %q", result)
	}
}

func TestRenderBlocks_Quote(t *testing.T) {
	blocks := []api.DocumentBlock{
		{BlockID: "page", BlockType: 1, Children: []string{"q1"}},
		{BlockID: "q1", ParentID: "page", BlockType: 15, Quote: &api.TextBlock{
			Elements: []api.TextElement{{TextRun: &api.TextRun{Content: "A wise quote"}}},
		}},
	}
	result := RenderBlocks(blocks)
	if !strings.Contains(result, "> A wise quote") {
		t.Errorf("expected blockquote, got %q", result)
	}
}

// --- Inline Escaping Tests ---

func TestRenderBlocks_InlineEscaping(t *testing.T) {
	blocks := []api.DocumentBlock{
		{BlockID: "page", BlockType: 1, Children: []string{"t1"}},
		{BlockID: "t1", ParentID: "page", BlockType: 2, Text: &api.TextBlock{
			Elements: []api.TextElement{
				{TextRun: &api.TextRun{Content: "Use **kwargs and _name"}},
			},
		}},
	}
	result := RenderBlocks(blocks)
	if !strings.Contains(result, `\*\*kwargs`) {
		t.Errorf("expected escaped **, got %q", result)
	}
	if !strings.Contains(result, `\_name`) {
		t.Errorf("expected escaped _, got %q", result)
	}
}

func TestRenderBlocks_InlineEscapingWithStyles(t *testing.T) {
	blocks := []api.DocumentBlock{
		{BlockID: "page", BlockType: 1, Children: []string{"t1"}},
		{BlockID: "t1", ParentID: "page", BlockType: 2, Text: &api.TextBlock{
			Elements: []api.TextElement{
				{TextRun: &api.TextRun{
					Content:          "has *stars*",
					TextElementStyle: &api.TextElementStyle{Bold: true},
				}},
			},
		}},
	}
	result := RenderBlocks(blocks)
	// Content should be escaped then wrapped in bold markers
	if !strings.Contains(result, `**has \*stars\***`) {
		t.Errorf("expected escaped content wrapped in bold, got %q", result)
	}
}

func TestRenderBlocks_BoldWithTrailingSpace(t *testing.T) {
	blocks := []api.DocumentBlock{
		{BlockID: "page", BlockType: 1, Children: []string{"t1"}},
		{BlockID: "t1", ParentID: "page", BlockType: 2, Text: &api.TextBlock{
			Elements: []api.TextElement{
				{TextRun: &api.TextRun{
					Content:          "important text ",
					TextElementStyle: &api.TextElementStyle{Bold: true},
				}},
			},
		}},
	}
	result := RenderBlocks(blocks)
	// Space should be outside the bold markers, not inside
	if !strings.Contains(result, "**important text** ") {
		t.Errorf("expected trailing space outside bold markers, got %q", result)
	}
	if strings.Contains(result, "**important text **") {
		t.Errorf("trailing space should not be inside bold markers, got %q", result)
	}
}

func TestRenderBlocks_ItalicWithLeadingSpace(t *testing.T) {
	blocks := []api.DocumentBlock{
		{BlockID: "page", BlockType: 1, Children: []string{"t1"}},
		{BlockID: "t1", ParentID: "page", BlockType: 2, Text: &api.TextBlock{
			Elements: []api.TextElement{
				{TextRun: &api.TextRun{
					Content:          " emphasis",
					TextElementStyle: &api.TextElementStyle{Italic: true},
				}},
			},
		}},
	}
	result := RenderBlocks(blocks)
	if !strings.Contains(result, " *emphasis*") {
		t.Errorf("expected leading space outside italic markers, got %q", result)
	}
}

func TestRenderBlocks_NoEscapingInInlineCode(t *testing.T) {
	blocks := []api.DocumentBlock{
		{BlockID: "page", BlockType: 1, Children: []string{"t1"}},
		{BlockID: "t1", ParentID: "page", BlockType: 2, Text: &api.TextBlock{
			Elements: []api.TextElement{
				{TextRun: &api.TextRun{
					Content:          "**not bold**",
					TextElementStyle: &api.TextElementStyle{InlineCode: true},
				}},
			},
		}},
	}
	result := RenderBlocks(blocks)
	if !strings.Contains(result, "`**not bold**`") {
		t.Errorf("expected unescaped content in inline code, got %q", result)
	}
}

func TestRenderBlocks_NoEscapingInCodeBlock(t *testing.T) {
	blocks := []api.DocumentBlock{
		{BlockID: "page", BlockType: 1, Children: []string{"c1"}},
		{BlockID: "c1", ParentID: "page", BlockType: 14, Code: &api.TextBlock{
			Style:    &api.TextStyle{Language: 49},
			Elements: []api.TextElement{{TextRun: &api.TextRun{Content: "x = [1, 2]"}}},
		}},
	}
	result := RenderBlocks(blocks)
	// Code block content should NOT have brackets escaped
	if strings.Contains(result, `\[`) {
		t.Errorf("expected no escaping in code block, got %q", result)
	}
}

// --- Block-Level Escaping Tests ---

func TestRenderBlocks_BlockLevelEscaping(t *testing.T) {
	tests := []struct {
		name    string
		content string
		escaped string
	}{
		{"heading trigger", "# Not a heading", `\# Not a heading`},
		{"blockquote trigger", "> Not a quote", `\> Not a quote`},
		{"unordered list dash", "- Not a list", `\- Not a list`},
		{"unordered list plus", "+ Not a list", `\+ Not a list`},
		{"ordered list", "1. Not ordered", `1\. Not ordered`},
		{"horizontal rule", "---", `\---`},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			blocks := []api.DocumentBlock{
				{BlockID: "page", BlockType: 1, Children: []string{"t1"}},
				{BlockID: "t1", ParentID: "page", BlockType: 2, Text: &api.TextBlock{
					Elements: []api.TextElement{{TextRun: &api.TextRun{Content: tc.content}}},
				}},
			}
			result := RenderBlocks(blocks)
			if !strings.Contains(result, tc.escaped) {
				t.Errorf("expected %q, got %q", tc.escaped, result)
			}
		})
	}
}

func TestRenderBlocks_BlockLevelEscapingWithIndentation(t *testing.T) {
	blocks := []api.DocumentBlock{
		{BlockID: "page", BlockType: 1, Children: []string{"t1"}},
		{BlockID: "t1", ParentID: "page", BlockType: 2, Text: &api.TextBlock{
			Elements: []api.TextElement{{TextRun: &api.TextRun{Content: "   # Indented heading"}}},
		}},
	}
	result := RenderBlocks(blocks)
	if !strings.Contains(result, `   \# Indented heading`) {
		t.Errorf("expected escaped indented heading trigger, got %q", result)
	}
}

func TestRenderBlocks_NoBlockEscapingInHeadings(t *testing.T) {
	blocks := []api.DocumentBlock{
		{BlockID: "page", BlockType: 1, Children: []string{"h1"}},
		{BlockID: "h1", ParentID: "page", BlockType: 3, Heading1: &api.TextBlock{
			Elements: []api.TextElement{{TextRun: &api.TextRun{Content: "Real Heading"}}},
		}},
	}
	result := RenderBlocks(blocks)
	if !strings.Contains(result, "# Real Heading") {
		t.Errorf("heading should not be escaped, got %q", result)
	}
}

// --- New TextElement Tests ---

func TestRenderBlocks_Equation(t *testing.T) {
	blocks := []api.DocumentBlock{
		{BlockID: "page", BlockType: 1, Children: []string{"t1"}},
		{BlockID: "t1", ParentID: "page", BlockType: 2, Text: &api.TextBlock{
			Elements: []api.TextElement{
				{TextRun: &api.TextRun{Content: "The formula is "}},
				{Equation: &api.Equation{Content: "E = mc^2"}},
			},
		}},
	}
	result := RenderBlocks(blocks)
	if !strings.Contains(result, "$E = mc^2$") {
		t.Errorf("expected equation rendering, got %q", result)
	}
}

func TestRenderBlocks_Reminder(t *testing.T) {
	blocks := []api.DocumentBlock{
		{BlockID: "page", BlockType: 1, Children: []string{"t1"}},
		{BlockID: "t1", ParentID: "page", BlockType: 2, Text: &api.TextBlock{
			Elements: []api.TextElement{
				{Reminder: &api.InlineReminder{
					ExpireTime: api.FlexInt64(1710460800000), // 2024-03-15T00:00:00Z
					IsWholeDay: true,
				}},
			},
		}},
	}
	result := RenderBlocks(blocks)
	if !strings.Contains(result, "[reminder: 2024-03-15]") {
		t.Errorf("expected whole-day reminder, got %q", result)
	}
}

func TestRenderBlocks_ReminderWithTime(t *testing.T) {
	blocks := []api.DocumentBlock{
		{BlockID: "page", BlockType: 1, Children: []string{"t1"}},
		{BlockID: "t1", ParentID: "page", BlockType: 2, Text: &api.TextBlock{
			Elements: []api.TextElement{
				{Reminder: &api.InlineReminder{
					ExpireTime: api.FlexInt64(1710500400000), // 2024-03-15T11:00:00Z
					IsWholeDay: false,
				}},
			},
		}},
	}
	result := RenderBlocks(blocks)
	if !strings.Contains(result, "[reminder: 2024-03-15T11:00:00Z]") {
		t.Errorf("expected datetime reminder in UTC, got %q", result)
	}
}

func TestRenderBlocks_ReminderZeroExpireTime(t *testing.T) {
	// When ExpireTime is 0 (field omitted or unparseable), render plain [reminder]
	blocks := []api.DocumentBlock{
		{BlockID: "page", BlockType: 1, Children: []string{"t1"}},
		{BlockID: "t1", ParentID: "page", BlockType: 2, Text: &api.TextBlock{
			Elements: []api.TextElement{
				{Reminder: &api.InlineReminder{IsWholeDay: true}},
			},
		}},
	}
	result := RenderBlocks(blocks)
	if !strings.Contains(result, "[reminder]") {
		t.Errorf("expected plain [reminder] for zero expire time, got %q", result)
	}
}

func TestRenderBlocks_InlineFile(t *testing.T) {
	blocks := []api.DocumentBlock{
		{BlockID: "page", BlockType: 1, Children: []string{"t1"}},
		{BlockID: "t1", ParentID: "page", BlockType: 2, Text: &api.TextBlock{
			Elements: []api.TextElement{
				{InlineFile: &api.InlineFile{FileToken: "boxcnABC123"}},
			},
		}},
	}
	result := RenderBlocks(blocks)
	if !strings.Contains(result, "[file: boxcnABC123]") {
		t.Errorf("expected inline file placeholder, got %q", result)
	}
}

// --- New Block Type Tests ---

func TestRenderBlocks_Callout(t *testing.T) {
	blocks := []api.DocumentBlock{
		{BlockID: "page", BlockType: 1, Children: []string{"co1"}},
		{BlockID: "co1", ParentID: "page", BlockType: 19,
			Children: []string{"co1t"},
			Callout:  &api.CalloutBlock{EmojiID: "bulb"},
		},
		{BlockID: "co1t", ParentID: "co1", BlockType: 2, Text: &api.TextBlock{
			Elements: []api.TextElement{{TextRun: &api.TextRun{Content: "Important note"}}},
		}},
	}
	result := RenderBlocks(blocks)
	if !strings.Contains(result, "> Important note") {
		t.Errorf("expected callout with > prefix, got %q", result)
	}
}

func TestRenderBlocks_Image(t *testing.T) {
	blocks := []api.DocumentBlock{
		{BlockID: "page", BlockType: 1, Children: []string{"img1"}},
		{BlockID: "img1", ParentID: "page", BlockType: 27, Image: &api.ImageBlock{
			Token: "boxcnIMG456",
		}},
	}
	result := RenderBlocks(blocks)
	if !strings.Contains(result, "[image: boxcnIMG456]") {
		t.Errorf("expected image placeholder, got %q", result)
	}
}

func TestRenderBlocks_MentionDoc(t *testing.T) {
	blocks := []api.DocumentBlock{
		{BlockID: "page", BlockType: 1, Children: []string{"t1"}},
		{BlockID: "t1", ParentID: "page", BlockType: 2, Text: &api.TextBlock{
			Elements: []api.TextElement{
				{MentionDoc: &api.MentionDoc{
					Token: "doccnABC",
					Title: "Design Doc",
					URL:   "https://example.com/doc",
				}},
			},
		}},
	}
	result := RenderBlocks(blocks)
	if !strings.Contains(result, "[Design Doc](https://example.com/doc)") {
		t.Errorf("expected mention doc link, got %q", result)
	}
}

func TestRenderBlocks_MentionDocNoURL(t *testing.T) {
	blocks := []api.DocumentBlock{
		{BlockID: "page", BlockType: 1, Children: []string{"t1"}},
		{BlockID: "t1", ParentID: "page", BlockType: 2, Text: &api.TextBlock{
			Elements: []api.TextElement{
				{MentionDoc: &api.MentionDoc{Token: "doccnABC", Title: "Design Doc"}},
			},
		}},
	}
	result := RenderBlocks(blocks)
	if !strings.Contains(result, "Design Doc") {
		t.Errorf("expected mention doc title, got %q", result)
	}
	if strings.Contains(result, "[Design Doc]") {
		t.Errorf("should not render as link without URL, got %q", result)
	}
}

func TestRenderBlocks_Grid(t *testing.T) {
	blocks := []api.DocumentBlock{
		{BlockID: "page", BlockType: 1, Children: []string{"grid"}},
		{BlockID: "grid", ParentID: "page", BlockType: 24,
			Children: []string{"col1", "col2"},
			Grid:     &api.GridBlock{ColumnSize: 2},
		},
		{BlockID: "col1", ParentID: "grid", BlockType: 25,
			Children:   []string{"col1t"},
			GridColumn: &api.GridColumnBlock{WidthRatio: 50},
		},
		{BlockID: "col1t", ParentID: "col1", BlockType: 2, Text: &api.TextBlock{
			Elements: []api.TextElement{{TextRun: &api.TextRun{Content: "Column A"}}},
		}},
		{BlockID: "col2", ParentID: "grid", BlockType: 25,
			Children:   []string{"col2t"},
			GridColumn: &api.GridColumnBlock{WidthRatio: 50},
		},
		{BlockID: "col2t", ParentID: "col2", BlockType: 2, Text: &api.TextBlock{
			Elements: []api.TextElement{{TextRun: &api.TextRun{Content: "Column B"}}},
		}},
	}
	result := RenderBlocks(blocks)
	if !strings.Contains(result, "Column A") || !strings.Contains(result, "Column B") {
		t.Errorf("expected both column contents, got %q", result)
	}
}

func TestRenderBlocks_QuoteContainer(t *testing.T) {
	blocks := []api.DocumentBlock{
		{BlockID: "page", BlockType: 1, Children: []string{"qc1"}},
		{BlockID: "qc1", ParentID: "page", BlockType: 34,
			Children:       []string{"qc1t"},
			QuoteContainer: &api.QuoteContainerBlock{},
		},
		{BlockID: "qc1t", ParentID: "qc1", BlockType: 2, Text: &api.TextBlock{
			Elements: []api.TextElement{{TextRun: &api.TextRun{Content: "Quoted text"}}},
		}},
	}
	result := RenderBlocks(blocks)
	if !strings.Contains(result, "> Quoted text") {
		t.Errorf("expected quote container with > prefix, got %q", result)
	}
}

func TestRenderBlocks_FileBlock(t *testing.T) {
	blocks := []api.DocumentBlock{
		{BlockID: "page", BlockType: 1, Children: []string{"f1"}},
		{BlockID: "f1", ParentID: "page", BlockType: 23, File: &api.FileBlock{
			Token: "boxcnFILE",
			Name:  "report.pdf",
		}},
	}
	result := RenderBlocks(blocks)
	if !strings.Contains(result, "[file: report.pdf]") {
		t.Errorf("expected file placeholder with name, got %q", result)
	}
}

func TestRenderBlocks_IframeBlock(t *testing.T) {
	blocks := []api.DocumentBlock{
		{BlockID: "page", BlockType: 1, Children: []string{"ifr1"}},
		{BlockID: "ifr1", ParentID: "page", BlockType: 26, Iframe: &api.IframeBlock{
			Component: &api.IframeComponent{URL: "https://www.figma.com/embed"},
		}},
	}
	result := RenderBlocks(blocks)
	if !strings.Contains(result, "[embed: https://www.figma.com/embed]") {
		t.Errorf("expected iframe embed, got %q", result)
	}
}

func TestRenderBlocks_PlaceholderBlocks(t *testing.T) {
	tests := []struct {
		blockType int
		expected  string
	}{
		{18, "[bitable]"},
		{20, "[chat]"},
		{21, "[diagram]"},
		{28, "[isv]"},
		{29, "[mindnote]"},
		{30, "[sheet]"},
		{36, "[okr]"},
	}
	for _, tc := range tests {
		t.Run(tc.expected, func(t *testing.T) {
			// Capture stderr to verify no warnings
			oldStderr := os.Stderr
			r, w, _ := os.Pipe()
			os.Stderr = w

			blocks := []api.DocumentBlock{
				{BlockID: "page", BlockType: 1, Children: []string{"b1"}},
				{BlockID: "b1", ParentID: "page", BlockType: tc.blockType},
			}
			result := RenderBlocks(blocks)

			w.Close()
			var buf bytes.Buffer
			buf.ReadFrom(r)
			os.Stderr = oldStderr

			if !strings.Contains(result, tc.expected) {
				t.Errorf("expected %q in output, got %q", tc.expected, result)
			}
			if strings.Contains(buf.String(), "warning: unsupported") {
				t.Errorf("should not warn for known block type %d", tc.blockType)
			}
		})
	}
}

func TestRenderBlocks_IdentifierPlaceholders(t *testing.T) {
	blocks := []api.DocumentBlock{
		{BlockID: "page", BlockType: 1, Children: []string{"task1", "jira1", "wiki1"}},
		{BlockID: "task1", ParentID: "page", BlockType: 35, Task: &api.TaskBlock{TaskID: "t_12345"}},
		{BlockID: "jira1", ParentID: "page", BlockType: 41, JiraIssue: &api.JiraIssueBlock{Key: "PROJ-123"}},
		{BlockID: "wiki1", ParentID: "page", BlockType: 42, WikiCatalog: &api.WikiCatalogBlock{WikiToken: "wikcnABC"}},
	}
	result := RenderBlocks(blocks)
	if !strings.Contains(result, "[task: t_12345]") {
		t.Errorf("expected task with ID, got %q", result)
	}
	if !strings.Contains(result, "[jira: PROJ-123]") {
		t.Errorf("expected jira with key, got %q", result)
	}
	if !strings.Contains(result, "[wiki: wikcnABC]") {
		t.Errorf("expected wiki with token, got %q", result)
	}
}

// --- Language Map Tests ---

func TestCodeLanguageName_AllMapped(t *testing.T) {
	for id := 1; id <= 75; id++ {
		name := codeLanguageName(id)
		if name == "" {
			t.Errorf("language ID %d has no mapping", id)
		}
	}
}

func TestCodeLanguageName_SwiftScheme(t *testing.T) {
	if got := codeLanguageName(58); got != "scheme" {
		t.Errorf("language ID 58 should be 'scheme', got %q", got)
	}
	if got := codeLanguageName(61); got != "swift" {
		t.Errorf("language ID 61 should be 'swift', got %q", got)
	}
}

// --- Table.Cells Ordering Test ---

func TestRenderBlocks_TableWithCellsOrdering(t *testing.T) {
	blocks := []api.DocumentBlock{
		{BlockID: "page", BlockType: 1, Children: []string{"tbl"}},
		{BlockID: "tbl", ParentID: "page", BlockType: 31,
			// Children in wrong order
			Children: []string{"c2", "c1"},
			Table: &api.TableBlock{
				// Cells in correct order
				Cells:    []string{"c1", "c2"},
				Property: &api.TableProperty{RowSize: 1, ColumnSize: 2},
			},
		},
		{BlockID: "c1", ParentID: "tbl", BlockType: 32, Children: []string{"c1t"}},
		{BlockID: "c1t", ParentID: "c1", BlockType: 2, Text: &api.TextBlock{
			Elements: []api.TextElement{{TextRun: &api.TextRun{Content: "First"}}},
		}},
		{BlockID: "c2", ParentID: "tbl", BlockType: 32, Children: []string{"c2t"}},
		{BlockID: "c2t", ParentID: "c2", BlockType: 2, Text: &api.TextBlock{
			Elements: []api.TextElement{{TextRun: &api.TextRun{Content: "Second"}}},
		}},
	}
	result := RenderBlocks(blocks)
	// Table.Cells should dictate order: c1 first, c2 second
	if !strings.Contains(result, "| First | Second |") {
		t.Errorf("expected cells in Table.Cells order, got %q", result)
	}
}

func TestRenderBlocks_NestedTable(t *testing.T) {
	// Outer table has one cell containing a nested 1x2 table
	blocks := []api.DocumentBlock{
		{BlockID: "page", BlockType: 1, Children: []string{"outer"}},
		{BlockID: "outer", ParentID: "page", BlockType: 31,
			Children: []string{"oc1"},
			Table: &api.TableBlock{
				Cells:    []string{"oc1"},
				Property: &api.TableProperty{RowSize: 1, ColumnSize: 1},
			},
		},
		{BlockID: "oc1", ParentID: "outer", BlockType: 32, Children: []string{"inner"}},
		{BlockID: "inner", ParentID: "oc1", BlockType: 31,
			Children: []string{"ic1", "ic2"},
			Table: &api.TableBlock{
				Cells:    []string{"ic1", "ic2"},
				Property: &api.TableProperty{RowSize: 1, ColumnSize: 2},
			},
		},
		{BlockID: "ic1", ParentID: "inner", BlockType: 32, Children: []string{"ic1t"}},
		{BlockID: "ic1t", ParentID: "ic1", BlockType: 2, Text: &api.TextBlock{
			Elements: []api.TextElement{{TextRun: &api.TextRun{Content: "Alpha"}}},
		}},
		{BlockID: "ic2", ParentID: "inner", BlockType: 32, Children: []string{"ic2t"}},
		{BlockID: "ic2t", ParentID: "ic2", BlockType: 2, Text: &api.TextBlock{
			Elements: []api.TextElement{{TextRun: &api.TextRun{Content: "Beta"}}},
		}},
	}
	result := RenderBlocks(blocks)
	// Nested table cells should appear as flattened text
	if !strings.Contains(result, "Alpha") || !strings.Contains(result, "Beta") {
		t.Errorf("expected nested table content preserved, got %q", result)
	}
}

func TestRenderBlocks_EmptyCallout(t *testing.T) {
	blocks := []api.DocumentBlock{
		{BlockID: "page", BlockType: 1, Children: []string{"co1"}},
		{BlockID: "co1", ParentID: "page", BlockType: 19,
			Callout: &api.CalloutBlock{EmojiID: "star"},
		},
	}
	result := RenderBlocks(blocks)
	if !strings.Contains(result, "> [callout: star]") {
		t.Errorf("expected empty callout with emoji ID, got %q", result)
	}
}

func TestRenderBlocks_EmptyCalloutNoEmoji(t *testing.T) {
	blocks := []api.DocumentBlock{
		{BlockID: "page", BlockType: 1, Children: []string{"co1"}},
		{BlockID: "co1", ParentID: "page", BlockType: 19},
	}
	result := RenderBlocks(blocks)
	if !strings.Contains(result, "> [callout]") {
		t.Errorf("expected empty callout placeholder, got %q", result)
	}
}

func TestRenderBlocks_NestedQuoteContainer(t *testing.T) {
	blocks := []api.DocumentBlock{
		{BlockID: "page", BlockType: 1, Children: []string{"qc1"}},
		{BlockID: "qc1", ParentID: "page", BlockType: 34, Children: []string{"qc2"}},
		{BlockID: "qc2", ParentID: "qc1", BlockType: 34, Children: []string{"qt"}},
		{BlockID: "qt", ParentID: "qc2", BlockType: 2, Text: &api.TextBlock{
			Elements: []api.TextElement{{TextRun: &api.TextRun{Content: "nested quote"}}},
		}},
	}
	result := RenderBlocks(blocks)
	if !strings.Contains(result, "> > nested quote") {
		t.Errorf("expected nested blockquote > > prefix, got %q", result)
	}
}

func TestRenderBlocks_OrderedListAcrossEmptyText(t *testing.T) {
	// Ordered list items with an empty text block between them should continue numbering
	blocks := []api.DocumentBlock{
		{BlockID: "page", BlockType: 1, Children: []string{"o1", "empty", "o2"}},
		{BlockID: "o1", ParentID: "page", BlockType: 13, Ordered: &api.TextBlock{
			Elements: []api.TextElement{{TextRun: &api.TextRun{Content: "First"}}},
		}},
		{BlockID: "empty", ParentID: "page", BlockType: 2, Text: &api.TextBlock{}},
		{BlockID: "o2", ParentID: "page", BlockType: 13, Ordered: &api.TextBlock{
			Elements: []api.TextElement{{TextRun: &api.TextRun{Content: "Second"}}},
		}},
	}
	result := RenderBlocks(blocks)
	if !strings.Contains(result, "1. First") {
		t.Errorf("expected '1. First', got %q", result)
	}
	if !strings.Contains(result, "2. Second") {
		t.Errorf("expected '2. Second' (continued numbering), got %q", result)
	}
}

func TestRenderBlocks_TaskWithDetails(t *testing.T) {
	blocks := []api.DocumentBlock{
		{BlockID: "page", BlockType: 1, Children: []string{"t1", "t2"}},
		{BlockID: "t1", ParentID: "page", BlockType: 35, Task: &api.TaskBlock{TaskID: "guid-1"}},
		{BlockID: "t2", ParentID: "page", BlockType: 35, Task: &api.TaskBlock{TaskID: "guid-2"}},
	}
	opts := RenderOptions{
		TaskDetails: map[string]TaskInfo{
			"guid-1": {Summary: "Buy groceries", Completed: false},
			"guid-2": {Summary: "Send report", Completed: true},
		},
	}
	result := RenderBlocksWithOptions(blocks, opts)
	if !strings.Contains(result, "- [ ] Buy groceries") {
		t.Errorf("expected incomplete task checkbox, got %q", result)
	}
	if !strings.Contains(result, "- [x] ~~Send report~~") {
		t.Errorf("expected completed task checkbox with strikethrough, got %q", result)
	}
}

func TestRenderBlocks_TaskWithoutDetails(t *testing.T) {
	blocks := []api.DocumentBlock{
		{BlockID: "page", BlockType: 1, Children: []string{"t1"}},
		{BlockID: "t1", ParentID: "page", BlockType: 35, Task: &api.TaskBlock{TaskID: "guid-1"}},
	}
	result := RenderBlocks(blocks)
	if !strings.Contains(result, "[task: guid-1]") {
		t.Errorf("expected task UUID fallback, got %q", result)
	}
}
