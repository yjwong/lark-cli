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
		var parts []string
		for child := n.FirstChild(); child != nil; child = child.NextSibling() {
			t := nodeText(child, source)
			if t != "" {
				parts = append(parts, t)
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
