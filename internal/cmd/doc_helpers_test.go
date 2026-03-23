package cmd

import (
	"testing"

	"github.com/yjwong/lark-cli/internal/api"
)

func TestParseSheetToken(t *testing.T) {
	tests := []struct {
		name       string
		token      string
		wantSpread string
		wantSheet  string
		wantOK     bool
	}{
		{"valid token", "ABC123_DEF456", "ABC123", "DEF456", true},
		{"token with multiple underscores", "ABC_DEF_GHI", "ABC", "DEF_GHI", true},
		{"no underscore", "ABCDEF", "", "", false},
		{"empty string", "", "", "", false},
		{"underscore only", "_", "", "", false},
		{"trailing underscore", "ABC_", "", "", false},
		{"leading underscore", "_DEF", "", "", false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			spread, sheet, ok := parseSheetToken(tc.token)
			if ok != tc.wantOK {
				t.Errorf("parseSheetToken(%q) ok = %v, want %v", tc.token, ok, tc.wantOK)
			}
			if spread != tc.wantSpread {
				t.Errorf("parseSheetToken(%q) spreadsheet = %q, want %q", tc.token, spread, tc.wantSpread)
			}
			if sheet != tc.wantSheet {
				t.Errorf("parseSheetToken(%q) sheet = %q, want %q", tc.token, sheet, tc.wantSheet)
			}
		})
	}
}

func TestMapBatchResults(t *testing.T) {
	t.Run("maps results by sheet ID", func(t *testing.T) {
		valueRanges := []api.ValueRange{
			{Range: "sheet1!A1:C3", Values: [][]any{{"a", "b"}}},
			{Range: "sheet2!A1:Z200", Values: [][]any{{"x"}, {"y"}}},
		}
		result := mapBatchResults(valueRanges)
		if len(result) != 2 {
			t.Fatalf("expected 2 entries, got %d", len(result))
		}
		if len(result["sheet1"]) != 1 {
			t.Errorf("sheet1: expected 1 row, got %d", len(result["sheet1"]))
		}
		if len(result["sheet2"]) != 2 {
			t.Errorf("sheet2: expected 2 rows, got %d", len(result["sheet2"]))
		}
	})

	t.Run("skips ranges without exclamation mark", func(t *testing.T) {
		valueRanges := []api.ValueRange{
			{Range: "malformed", Values: [][]any{{"a"}}},
			{Range: "sheet1!A1:C3", Values: [][]any{{"b"}}},
		}
		result := mapBatchResults(valueRanges)
		if len(result) != 1 {
			t.Fatalf("expected 1 entry, got %d", len(result))
		}
		if _, ok := result["sheet1"]; !ok {
			t.Error("expected sheet1 in results")
		}
	})

	t.Run("batch API omits ranges", func(t *testing.T) {
		// Simulates batch API returning fewer results than requested
		valueRanges := []api.ValueRange{
			{Range: "sheet1!A1:D10", Values: [][]any{{"only this one"}}},
			// sheet2 and sheet3 omitted by batch API
		}
		result := mapBatchResults(valueRanges)
		if len(result) != 1 {
			t.Fatalf("expected 1 entry, got %d", len(result))
		}
		if _, ok := result["sheet2"]; ok {
			t.Error("sheet2 should not be in results")
		}
	})

	t.Run("empty value ranges", func(t *testing.T) {
		result := mapBatchResults(nil)
		if len(result) != 0 {
			t.Fatalf("expected 0 entries, got %d", len(result))
		}
	})

	t.Run("range with nil values", func(t *testing.T) {
		valueRanges := []api.ValueRange{
			{Range: "sheet1!A1:C3", Values: nil},
		}
		result := mapBatchResults(valueRanges)
		if result["sheet1"] != nil {
			t.Error("expected nil values for sheet1")
		}
	})
}
