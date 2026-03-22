package main

import (
	"strings"
	"testing"
)

// TestCleanContent tests the cleanContent function in isolation via ParseContent.
// Since cleanContent is unexported, we test it through its effects on ParsedDoc.Content.
func TestCleanContent(t *testing.T) {
	t.Run("split table headers are joined", func(t *testing.T) {
		// A table cell split across a newline and leading whitespace should be collapsed.
		input := "| \n      Parameter |"
		got := cleanContent(input)
		if strings.Contains(got, "\n") {
			t.Errorf("expected newline inside cell to be removed, got: %q", got)
		}
		if !strings.Contains(got, "| Parameter |") {
			t.Errorf("expected '| Parameter |', got: %q", got)
		}
	})

	t.Run("HTML entities are decoded", func(t *testing.T) {
		// Note: &emsp;, &nbsp;, and &amp; are decoded to their plain-text forms.
		// &lt; and &gt; are decoded to < and >, but those are subsequently stripped by the
		// HTML tag remover (reHTMLTags) if they form tag-like sequences. Test cases here
		// use sequences that cannot be mistaken for HTML tags.
		cases := []struct {
			input string
			want  string
		}{
			{"foo&emsp;bar", "foo  bar"},
			{"foo&nbsp;bar", "foo bar"},
			{"price &amp; cost", "price & cost"},
		}
		for _, tc := range cases {
			got := cleanContent(tc.input)
			if !strings.Contains(got, tc.want) {
				t.Errorf("input %q: want %q in output, got %q", tc.input, tc.want, got)
			}
		}
	})

	t.Run("angle bracket entities decoded after tag stripping", func(t *testing.T) {
		// &lt;/&gt; must be decoded AFTER HTML tag stripping to preserve
		// type signatures like map<string, string> that would otherwise
		// be stripped as HTML tags.
		input := "`map&lt;string, string&gt;`"
		got := cleanContent(input)
		if !strings.Contains(got, "map<string, string>") {
			t.Errorf("expected map<string, string> to be preserved, got: %q", got)
		}

		// Also test simple inequality
		input2 := "value &lt;= 10"
		got2 := cleanContent(input2)
		if !strings.Contains(got2, "<=") {
			t.Errorf("expected '<=' preserved for non-tag-like sequence, got: %q", got)
		}
	})

	t.Run("highlight markers are stripped", func(t *testing.T) {
		input := "This is ==highlighted== text."
		got := cleanContent(input)
		if strings.Contains(got, "==") {
			t.Errorf("expected == markers to be removed, got: %q", got)
		}
		if !strings.Contains(got, "highlighted") {
			t.Errorf("expected 'highlighted' text preserved, got: %q", got)
		}
	})

	t.Run("escaped brackets are unescaped", func(t *testing.T) {
		input := `The type is string\[\].`
		got := cleanContent(input)
		if strings.Contains(got, `\[\]`) {
			t.Errorf("expected \\[\\] to be unescaped, got: %q", got)
		}
		if !strings.Contains(got, "[]") {
			t.Errorf("expected '[]' in output, got: %q", got)
		}
	})

	t.Run("ssl:ttdoc markdown links keep display text", func(t *testing.T) {
		input := "See [the guide](/ssl:ttdoc/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/overview) for details."
		got := cleanContent(input)
		if strings.Contains(got, "/ssl:ttdoc/") {
			t.Errorf("expected /ssl:ttdoc/ URL to be removed, got: %q", got)
		}
		if !strings.Contains(got, "the guide") {
			t.Errorf("expected link text 'the guide' to be preserved, got: %q", got)
		}
	})

	t.Run("bare ssl:ttdoc URLs in parentheses are removed", func(t *testing.T) {
		input := "Some text(/ssl:ttdoc/uAjLw4CM/path) more text."
		got := cleanContent(input)
		if strings.Contains(got, "/ssl:ttdoc/") {
			t.Errorf("expected bare /ssl:ttdoc/ URL to be removed, got: %q", got)
		}
	})

	t.Run(":::note converted to blockquote", func(t *testing.T) {
		input := ":::note\nThis is a note.\n:::"
		got := cleanContent(input)
		if strings.Contains(got, ":::") {
			t.Errorf("expected ::: markers to be removed, got: %q", got)
		}
		if !strings.Contains(got, "> **Note:**") {
			t.Errorf("expected '> **Note:**' in output, got: %q", got)
		}
	})

	t.Run(":::warning converted to blockquote", func(t *testing.T) {
		input := ":::warning\nBe careful.\n:::"
		got := cleanContent(input)
		if strings.Contains(got, ":::") {
			t.Errorf("expected ::: markers to be removed, got: %q", got)
		}
		if !strings.Contains(got, "> **Warning:**") {
			t.Errorf("expected '> **Warning:**' in output, got: %q", got)
		}
	})

	t.Run("multi-line admonition quotes all lines", func(t *testing.T) {
		input := ":::note\nFirst line.\nSecond line.\nThird line.\n:::"
		got := cleanContent(input)
		if strings.Contains(got, ":::") {
			t.Errorf("expected ::: markers to be removed, got: %q", got)
		}
		if !strings.Contains(got, "> **Note:** First line.") {
			t.Errorf("expected first line with label, got: %q", got)
		}
		if !strings.Contains(got, "> Second line.") {
			t.Errorf("expected second line blockquoted, got: %q", got)
		}
		if !strings.Contains(got, "> Third line.") {
			t.Errorf("expected third line blockquoted, got: %q", got)
		}
	})

	t.Run(":::html fencing is stripped but content preserved", func(t *testing.T) {
		input := ":::html\nSome HTML content here.\n:::"
		got := cleanContent(input)
		if strings.Contains(got, ":::") {
			t.Errorf("expected ::: markers to be removed, got: %q", got)
		}
		if !strings.Contains(got, "Some HTML content here.") {
			t.Errorf("expected content preserved, got: %q", got)
		}
	})

	t.Run("{Try it}(...) links are removed", func(t *testing.T) {
		input := "Use the API {Try it}(https://open.feishu.cn/api-explorer) here."
		got := cleanContent(input)
		if strings.Contains(got, "{Try it}") {
			t.Errorf("expected {Try it} link to be removed, got: %q", got)
		}
		if !strings.Contains(got, "Use the API") {
			t.Errorf("expected surrounding text preserved, got: %q", got)
		}
	})

	t.Run("tooltips are removed", func(t *testing.T) {
		input := `Some text <md-tooltip content="This is a tooltip">hover me</md-tooltip> after.`
		got := cleanContent(input)
		if strings.Contains(got, "md-tooltip") {
			t.Errorf("expected md-tooltip tags to be removed, got: %q", got)
		}
	})

	t.Run("multiple blank lines collapsed to two newlines", func(t *testing.T) {
		input := "Line one.\n\n\n\n\nLine two."
		got := cleanContent(input)
		if strings.Contains(got, "\n\n\n") {
			t.Errorf("expected 3+ blank lines collapsed to 2, got: %q", got)
		}
		if !strings.Contains(got, "Line one.") || !strings.Contains(got, "Line two.") {
			t.Errorf("expected both lines preserved, got: %q", got)
		}
	})
}

func TestFenceJSONBlocks(t *testing.T) {
	t.Run("JSON after Request body example heading gets fenced", func(t *testing.T) {
		input := "### Request body example\n\n{\n  \"key\": \"value\"\n}"
		got := fenceJSONBlocks(input)
		if !strings.Contains(got, "```json") {
			t.Errorf("expected ```json fence added, got:\n%s", got)
		}
		if !strings.Contains(got, "\"key\": \"value\"") {
			t.Errorf("expected JSON content preserved, got:\n%s", got)
		}
		if !strings.HasSuffix(strings.TrimSpace(got), "```") {
			t.Errorf("expected closing ``` fence, got:\n%s", got)
		}
	})

	t.Run("JSON after Response body example heading gets fenced", func(t *testing.T) {
		input := "### Response body example\n\n{\n  \"code\": 0\n}"
		got := fenceJSONBlocks(input)
		if !strings.Contains(got, "```json") {
			t.Errorf("expected ```json fence added, got:\n%s", got)
		}
		if !strings.Contains(got, "\"code\": 0") {
			t.Errorf("expected JSON content preserved, got:\n%s", got)
		}
	})

	t.Run("already-fenced JSON is not double-fenced", func(t *testing.T) {
		input := "### Request body example\n\n```json\n{\n  \"key\": \"value\"\n}\n```"
		got := fenceJSONBlocks(input)
		count := strings.Count(got, "```json")
		if count != 1 {
			t.Errorf("expected exactly one ```json fence, got %d in:\n%s", count, got)
		}
	})

	t.Run("non-JSON content after heading is left alone", func(t *testing.T) {
		input := "### Request body example\n\nThis is just text, not JSON."
		got := fenceJSONBlocks(input)
		if strings.Contains(got, "```json") {
			t.Errorf("expected no fence added for non-JSON content, got:\n%s", got)
		}
	})

	t.Run("array JSON after example heading gets fenced", func(t *testing.T) {
		input := "### Response body example\n\n[\n  {\"id\": 1}\n]"
		got := fenceJSONBlocks(input)
		if !strings.Contains(got, "```json") {
			t.Errorf("expected ```json fence added for array, got:\n%s", got)
		}
	})
}

func TestConvertTables(t *testing.T) {
	t.Run("basic md-table with header and data rows", func(t *testing.T) {
		input := `<md-table>
<md-tr>
<md-th>Name</md-th>
<md-th>Type</md-th>
</md-tr>
<md-tr>
<md-td>user_id</md-td>
<md-td>string</md-td>
</md-tr>
</md-table>`
		got := convertTables(input)

		if strings.Contains(got, "<md-table>") || strings.Contains(got, "</md-table>") {
			t.Errorf("expected md-table tags removed, got:\n%s", got)
		}
		if !strings.Contains(got, "| Name | Type |") {
			t.Errorf("expected header row '| Name | Type |', got:\n%s", got)
		}
		if !strings.Contains(got, "| --- |") {
			t.Errorf("expected separator row, got:\n%s", got)
		}
		if !strings.Contains(got, "| user_id | string |") {
			t.Errorf("expected data row '| user_id | string |', got:\n%s", got)
		}
	})

	t.Run("empty table returns empty string", func(t *testing.T) {
		input := "<md-table></md-table>"
		got := convertTables(input)
		if strings.TrimSpace(got) != "" {
			t.Errorf("expected empty output for empty table, got: %q", got)
		}
	})

	t.Run("multiple tables are all converted", func(t *testing.T) {
		input := `<md-table>
<md-tr><md-th>A</md-th></md-tr>
<md-tr><md-td>1</md-td></md-tr>
</md-table>
<md-table>
<md-tr><md-th>B</md-th></md-tr>
<md-tr><md-td>2</md-td></md-tr>
</md-table>`
		got := convertTables(input)
		if strings.Contains(got, "<md-table>") {
			t.Errorf("expected all md-table tags removed, got:\n%s", got)
		}
		if !strings.Contains(got, "| A |") || !strings.Contains(got, "| B |") {
			t.Errorf("expected both tables converted, got:\n%s", got)
		}
	})
}

func TestConvertDataTables(t *testing.T) {
	t.Run("md-dt-table with nested levels produces indented markdown table", func(t *testing.T) {
		input := `<md-dt-table>
<md-dt-th>Field</md-dt-th>
<md-dt-th>Description</md-dt-th>
<md-dt-tr level="0">
<md-dt-td>parent_field</md-dt-td>
<md-dt-td>A top-level field</md-dt-td>
</md-dt-tr>
<md-dt-tr level="1">
<md-dt-td>child_field</md-dt-td>
<md-dt-td>A nested field</md-dt-td>
</md-dt-tr>
</md-dt-table>`
		got := convertDataTables(input)

		if strings.Contains(got, "<md-dt-table>") || strings.Contains(got, "</md-dt-table>") {
			t.Errorf("expected md-dt-table tags removed, got:\n%s", got)
		}
		if !strings.Contains(got, "| Field | Description |") {
			t.Errorf("expected header row, got:\n%s", got)
		}
		if !strings.Contains(got, "| parent_field |") {
			t.Errorf("expected top-level row, got:\n%s", got)
		}
		// Level 1 should be indented with 2 spaces
		if !strings.Contains(got, "|   child_field |") {
			t.Errorf("expected indented child row with 2 spaces, got:\n%s", got)
		}
	})

	t.Run("empty md-dt-table returns empty string", func(t *testing.T) {
		input := "<md-dt-table></md-dt-table>"
		got := convertDataTables(input)
		if strings.TrimSpace(got) != "" {
			t.Errorf("expected empty output for empty dt-table, got: %q", got)
		}
	})
}

func TestParseContent(t *testing.T) {
	// A realistic Lark doc snippet containing the key fields ParseContent extracts.
	raw := `# Get calendar events

<md-table>
<md-tr>
<md-th>HTTP Method</md-th>
<md-td>GET</md-td>
</md-tr>
<md-tr>
<md-th>HTTP URL</md-th>
<md-td>https://open.feishu.cn/open-apis/calendar/v4/events</md-td>
</md-tr>
<md-tr>
<md-th>Rate Limit</md-th>
<md-td>100 QPS per app</md-td>
</md-tr>
</md-table>

<md-app-support types="self_built,marketplace">Supported app types</md-app-support>

## Required scopes

<md-table>
<md-tr>
<md-th>Required scopes</md-th>
<md-td><md-perm name="calendar:calendar" desc="Read calendar">Read calendar</md-perm></md-td>
</md-tr>
</md-table>

## Description

Retrieves events from the specified calendar.

## Request body example

{
  "calendar_id": "feishu_cal_abc123",
  "page_size": 50
}

## Response body example

{
  "code": 0,
  "data": {
    "items": []
  }
}
`

	doc := ParseContent(raw)

	t.Run("method extracted", func(t *testing.T) {
		if doc.Method != "GET" {
			t.Errorf("expected Method='GET', got %q", doc.Method)
		}
	})

	t.Run("API path extracted", func(t *testing.T) {
		if !strings.Contains(doc.APIPath, "/calendar/v4/events") {
			t.Errorf("expected APIPath to contain '/calendar/v4/events', got %q", doc.APIPath)
		}
	})

	t.Run("rate limit extracted", func(t *testing.T) {
		if !strings.Contains(doc.RateLimit, "100") {
			t.Errorf("expected RateLimit to contain '100', got %q", doc.RateLimit)
		}
	})

	t.Run("app types extracted", func(t *testing.T) {
		if !strings.Contains(doc.AppTypes, "self_built") {
			t.Errorf("expected AppTypes to contain 'self_built', got %q", doc.AppTypes)
		}
	})

	t.Run("scopes extracted", func(t *testing.T) {
		if len(doc.Scopes) == 0 {
			t.Errorf("expected at least one scope, got none")
		}
		found := false
		for _, s := range doc.Scopes {
			if s == "calendar:calendar" {
				found = true
			}
		}
		if !found {
			t.Errorf("expected scope 'calendar:calendar', got %v", doc.Scopes)
		}
	})

	t.Run("content is cleaned markdown", func(t *testing.T) {
		if strings.Contains(doc.Content, "<md-table>") {
			t.Errorf("expected md-table tags removed from content")
		}
		if strings.Contains(doc.Content, "<md-perm") {
			t.Errorf("expected md-perm tags removed from content")
		}
		if !strings.Contains(doc.Content, "Get calendar events") {
			t.Errorf("expected heading preserved in content, got:\n%s", doc.Content)
		}
	})

	t.Run("request body example is fenced as JSON", func(t *testing.T) {
		if !strings.Contains(doc.Content, "```json") {
			t.Errorf("expected JSON fence in content, got:\n%s", doc.Content)
		}
		if !strings.Contains(doc.Content, "\"calendar_id\"") {
			t.Errorf("expected request body JSON content, got:\n%s", doc.Content)
		}
	})

	t.Run("content is trimmed", func(t *testing.T) {
		if doc.Content != strings.TrimSpace(doc.Content) {
			t.Errorf("expected content to be trimmed, got leading/trailing whitespace")
		}
	})
}

func TestExtractScopes(t *testing.T) {
	t.Run("extracts scope names from Required scopes section", func(t *testing.T) {
		raw := `<md-tr>
<md-th>Required scopes</md-th>
<md-td>
<md-perm name="calendar:calendar.readonly" desc="Read calendar">Read calendar</md-perm>
<md-perm name="calendar:calendar" desc="Manage calendar">Manage calendar</md-perm>
</md-td>
</md-tr>`
		scopes := extractScopes(raw, "Required scopes")
		if len(scopes) != 2 {
			t.Fatalf("expected 2 scopes, got %d: %v", len(scopes), scopes)
		}
		if scopes[0] != "calendar:calendar.readonly" {
			t.Errorf("expected first scope 'calendar:calendar.readonly', got %q", scopes[0])
		}
		if scopes[1] != "calendar:calendar" {
			t.Errorf("expected second scope 'calendar:calendar', got %q", scopes[1])
		}
	})

	t.Run("returns nil when section is absent", func(t *testing.T) {
		raw := "<md-tr><md-th>Other section</md-th><md-td>something</md-td></md-tr>"
		scopes := extractScopes(raw, "Required scopes")
		if scopes != nil {
			t.Errorf("expected nil scopes when section missing, got %v", scopes)
		}
	})

	t.Run("deduplicates repeated scope names", func(t *testing.T) {
		raw := `<md-tr>
<md-th>Required scopes</md-th>
<md-td>
<md-perm name="calendar:calendar" desc="x">text</md-perm>
<md-perm name="calendar:calendar" desc="y">text</md-perm>
</md-td>
</md-tr>`
		scopes := extractScopes(raw, "Required scopes")
		if len(scopes) != 1 {
			t.Errorf("expected 1 deduplicated scope, got %d: %v", len(scopes), scopes)
		}
	})

	t.Run("extracts from Required field scopes section independently", func(t *testing.T) {
		raw := `<md-tr>
<md-th>Required scopes</md-th>
<md-td><md-perm name="scope.a" desc="">a</md-perm></md-td>
</md-tr>
<md-tr>
<md-th>Required field scopes</md-th>
<md-td><md-perm name="scope.b" desc="">b</md-perm></md-td>
</md-tr>`
		scopes := extractScopes(raw, "Required scopes")
		fieldScopes := extractScopes(raw, "Required field scopes")

		if len(scopes) != 1 || scopes[0] != "scope.a" {
			t.Errorf("expected ['scope.a'], got %v", scopes)
		}
		if len(fieldScopes) != 1 || fieldScopes[0] != "scope.b" {
			t.Errorf("expected ['scope.b'], got %v", fieldScopes)
		}
	})
}
