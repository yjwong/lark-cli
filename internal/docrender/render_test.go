package docrender

import (
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
