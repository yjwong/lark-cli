package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/yjwong/lark-cli/internal/api"
	"github.com/yjwong/lark-cli/internal/output"
)

var sheetCmd = &cobra.Command{
	Use:   "sheet",
	Short: "Spreadsheet commands",
	Long:  "Read and query Lark Sheets (spreadsheets)",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		validateScopeGroup("documents")
	},
}

// --- sheet list ---

var sheetListCmd = &cobra.Command{
	Use:   "list <spreadsheet_token>",
	Short: "List all sheets in a spreadsheet",
	Long: `List all sheets (tabs) within a Lark spreadsheet.

The spreadsheet_token is from the spreadsheet URL.
For example, if the URL is https://xxx.larksuite.com/sheets/T4mHsrFyzhXrj0tVzRslUGx8gkA
then the spreadsheet_token is T4mHsrFyzhXrj0tVzRslUGx8gkA.

Examples:
  lark sheet list T4mHsrFyzhXrj0tVzRslUGx8gkA`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		token := args[0]

		client := api.NewClient()

		sheets, err := client.GetSpreadsheetSheets(token)
		if err != nil {
			output.Fatal("API_ERROR", err)
		}

		outputSheets := make([]api.OutputSheet, len(sheets))
		for i, s := range sheets {
			outputSheets[i] = api.OutputSheet{
				SheetID: s.SheetID,
				Title:   s.Title,
				Index:   s.Index,
				Hidden:  s.Hidden,
			}
			if s.GridProperties != nil {
				outputSheets[i].RowCount = s.GridProperties.RowCount
				outputSheets[i].ColumnCount = s.GridProperties.ColumnCount
			}
		}

		result := api.OutputSheetList{
			SpreadsheetToken: token,
			Sheets:           outputSheets,
			Count:            len(outputSheets),
		}

		output.JSON(result)
	},
}

// --- sheet read ---

var sheetReadCmd = &cobra.Command{
	Use:   "read <spreadsheet_token>",
	Short: "Read cell data from a sheet",
	Long: `Read cell values from a Lark spreadsheet.

By default, reads the first sheet (index 0) and up to 1000 rows.
Use --sheet to specify a sheet ID, and --range to specify a cell range.

The spreadsheet_token is from the spreadsheet URL.
For example, if the URL is https://xxx.larksuite.com/sheets/T4mHsrFyzhXrj0tVzRslUGx8gkA
then the spreadsheet_token is T4mHsrFyzhXrj0tVzRslUGx8gkA.

Range format: A1:Z100 (columns A-Z, rows 1-100)
If no range is specified, reads from A1 up to the sheet's actual dimensions (max 1000 rows).

Examples:
  lark sheet read T4mHsrFyzhXrj0tVzRslUGx8gkA
  lark sheet read T4mHsrFyzhXrj0tVzRslUGx8gkA --sheet abc123
  lark sheet read T4mHsrFyzhXrj0tVzRslUGx8gkA --range A1:D50
  lark sheet read T4mHsrFyzhXrj0tVzRslUGx8gkA --sheet abc123 --range A1:Z100`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		token := args[0]
		sheetID, _ := cmd.Flags().GetString("sheet")
		rangeSpec, _ := cmd.Flags().GetString("range")

		client := api.NewClient()

		// If no sheet ID specified, get the first sheet
		if sheetID == "" {
			sheets, err := client.GetSpreadsheetSheets(token)
			if err != nil {
				output.Fatal("API_ERROR", err)
			}
			if len(sheets) == 0 {
				output.Fatal("NO_SHEETS", fmt.Errorf("spreadsheet has no sheets"))
			}
			// Find the sheet with the lowest index (first sheet)
			firstSheet := sheets[0]
			for _, s := range sheets[1:] {
				if s.Index < firstSheet.Index {
					firstSheet = s
				}
			}
			sheetID = firstSheet.SheetID
		}

		// Build the range string
		var fullRange string
		if rangeSpec != "" {
			fullRange = sheetID + "!" + rangeSpec
		} else {
			// Default: read up to 1000 rows, determined by sheet size
			// Get sheet metadata to determine actual dimensions
			sheet, err := client.GetSheetMetadata(token, sheetID)
			if err != nil {
				// Fall back to a reasonable default if we can't get metadata
				fullRange = sheetID + "!A1:Z1000"
			} else if sheet.GridProperties != nil {
				rowCount := sheet.GridProperties.RowCount
				colCount := sheet.GridProperties.ColumnCount
				if rowCount > 1000 {
					rowCount = 1000
				}
				if colCount > 26 {
					colCount = 26 // Limit to column Z for simplicity
				}
				endCol := columnIndexToLetter(colCount)
				fullRange = fmt.Sprintf("%s!A1:%s%d", sheetID, endCol, rowCount)
			} else {
				fullRange = sheetID + "!A1:Z1000"
			}
		}

		// Get the data
		data, err := client.GetSheetData(token, fullRange)
		if err != nil {
			output.Fatal("API_ERROR", err)
		}

		// Extract values from response
		var values [][]any
		var actualRange string
		if data.ValueRange != nil {
			values = data.ValueRange.Values
			actualRange = data.ValueRange.Range
		}

		// Calculate dimensions
		rowCount := len(values)
		colCount := 0
		for _, row := range values {
			if len(row) > colCount {
				colCount = len(row)
			}
		}

		result := api.OutputSheetData{
			SpreadsheetToken: token,
			SheetID:          sheetID,
			Range:            actualRange,
			RowCount:         rowCount,
			ColumnCount:      colCount,
			Values:           values,
		}

		output.JSON(result)
	},
}

// columnIndexToLetter converts a 1-based column index to a letter (1=A, 26=Z)
func columnIndexToLetter(index int) string {
	if index <= 0 {
		return "A"
	}
	if index > 26 {
		index = 26
	}
	return string(rune('A' + index - 1))
}

// --- sheet write ---

var sheetWriteCmd = &cobra.Command{
	Use:   "write <spreadsheet_token>",
	Short: "Write cell data to a sheet",
	Long: `Write cell values to a Lark spreadsheet.

Values are provided as a JSON array of arrays via --values or stdin.

By default, writes to the first sheet starting at A1.
Use --sheet to specify a sheet ID, and --range to specify a target range.

The spreadsheet_token is from the spreadsheet URL.

Examples:
  lark sheet write T4mHsrFyzhXrj0tVzRslUGx8gkA --values '[["hello","world"],["foo","bar"]]'
  lark sheet write T4mHsrFyzhXrj0tVzRslUGx8gkA --sheet abc123 --range A1:B2 --values '[["a","b"],["c","d"]]'
  echo '[["a","b"]]' | lark sheet write T4mHsrFyzhXrj0tVzRslUGx8gkA`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		token := args[0]
		sheetID, _ := cmd.Flags().GetString("sheet")
		rangeSpec, _ := cmd.Flags().GetString("range")
		valuesJSON, _ := cmd.Flags().GetString("values")

		client := api.NewClient()

		// If no sheet ID specified, get the first sheet
		if sheetID == "" {
			sheets, err := client.GetSpreadsheetSheets(token)
			if err != nil {
				output.Fatal("API_ERROR", err)
			}
			if len(sheets) == 0 {
				output.Fatal("NO_SHEETS", fmt.Errorf("spreadsheet has no sheets"))
			}
			firstSheet := sheets[0]
			for _, s := range sheets[1:] {
				if s.Index < firstSheet.Index {
					firstSheet = s
				}
			}
			sheetID = firstSheet.SheetID
		}

		// Parse values from --values flag or stdin
		if valuesJSON == "" {
			// Try reading from stdin
			stat, _ := os.Stdin.Stat()
			if (stat.Mode() & os.ModeCharDevice) == 0 {
				data, err := io.ReadAll(os.Stdin)
				if err != nil {
					output.Fatal("INPUT_ERROR", fmt.Errorf("failed to read stdin: %w", err))
				}
				valuesJSON = string(data)
			}
		}

		if valuesJSON == "" {
			output.Fatal("MISSING_ARG", fmt.Errorf("--values is required or provide JSON via stdin"))
		}

		var values [][]any
		if err := json.Unmarshal([]byte(valuesJSON), &values); err != nil {
			output.Fatal("PARSE_ERROR", fmt.Errorf("invalid values JSON (must be array of arrays): %w", err))
		}

		// Build the range string
		if rangeSpec == "" {
			rangeSpec = "A1"
		}
		fullRange := sheetID + "!" + rangeSpec

		// Write the data
		data, err := client.SetSheetData(token, fullRange, values)
		if err != nil {
			output.Fatal("API_ERROR", err)
		}

		result := api.OutputSheetWrite{
			Success: true,
		}
		if data != nil {
			result.UpdatedRange = data.UpdatedRange
			result.UpdatedRows = data.UpdatedRows
			result.UpdatedColumns = data.UpdatedColumns
			result.UpdatedCells = data.UpdatedCells
			result.Revision = data.Revision
		}

		output.JSON(result)
	},
}

func init() {
	// Register subcommands
	sheetCmd.AddCommand(sheetListCmd)
	sheetCmd.AddCommand(sheetReadCmd)
	sheetCmd.AddCommand(sheetWriteCmd)

	// Flags for sheet read
	sheetReadCmd.Flags().String("sheet", "", "Sheet ID to read from (default: first sheet)")
	sheetReadCmd.Flags().String("range", "", "Cell range to read (e.g., A1:Z100)")

	// Flags for sheet write
	sheetWriteCmd.Flags().String("sheet", "", "Sheet ID to write to (default: first sheet)")
	sheetWriteCmd.Flags().String("range", "", "Target range in A1 notation (e.g., A1:C3, default: A1)")
	sheetWriteCmd.Flags().String("values", "", "JSON array of arrays (e.g., '[[\"a\",\"b\"],[\"c\",\"d\"]]')")
}
