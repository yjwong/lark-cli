package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"github.com/yjwong/lark-cli/internal/api"
	"github.com/yjwong/lark-cli/internal/output"
)

// resolveSheetID returns the given sheetID if non-empty, otherwise fetches the
// first sheet (by index) from the spreadsheet.
func resolveSheetID(client *api.Client, token, sheetID string) string {
	if sheetID != "" {
		return sheetID
	}
	sheets, err := client.GetSpreadsheetSheets(token)
	if err != nil {
		output.Fatal("API_ERROR", err)
	}
	if len(sheets) == 0 {
		output.Fatal("NO_SHEETS", fmt.Errorf("spreadsheet has no sheets"))
	}
	first := sheets[0]
	for _, s := range sheets[1:] {
		if s.Index < first.Index {
			first = s
		}
	}
	return first.SheetID
}

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

		sheetID = resolveSheetID(client, token, sheetID)

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

// --- sheet create ---

var sheetCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new spreadsheet",
	Long: `Create a new Lark spreadsheet.

Examples:
  lark sheet create --title "My Spreadsheet"
  lark sheet create --title "My Spreadsheet" --folder fldbcRho46N6MQ3mJkOAuPabcef`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		folderToken, _ := cmd.Flags().GetString("folder")

		if title == "" {
			output.Fatal("MISSING_ARG", fmt.Errorf("--title is required"))
		}

		client := api.NewClient()

		spreadsheet, err := client.CreateSpreadsheet(title, folderToken)
		if err != nil {
			output.Fatal("API_ERROR", err)
		}

		result := api.OutputSpreadsheetCreate{
			Success:          true,
			SpreadsheetToken: spreadsheet.SpreadsheetToken,
			Title:            spreadsheet.Title,
			URL:              spreadsheet.URL,
			FolderToken:      spreadsheet.FolderToken,
		}

		output.JSON(result)
	},
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
		sheetID = resolveSheetID(client, token, sheetID)

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

		autoType, _ := cmd.Flags().GetBool("auto-type")
		if autoType {
			values = autoTypeValues(values)
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

// autoTypeValues converts string values to appropriate types:
// - "YYYY-MM-DD" date strings become Excel serial date numbers
// - Integer strings become int
// - Float strings become float64
func autoTypeValues(values [][]any) [][]any {
	dateRegex := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	for i, row := range values {
		for j, cell := range row {
			s, ok := cell.(string)
			if !ok {
				continue
			}
			if dateRegex.MatchString(s) {
				t, err := time.Parse("2006-01-02", s)
				if err == nil {
					// Excel serial date: days since 1899-12-30
					epoch := time.Date(1899, 12, 30, 0, 0, 0, 0, time.UTC)
					values[i][j] = int(t.Sub(epoch).Hours() / 24)
					continue
				}
			}
			if n, err := strconv.Atoi(s); err == nil {
				values[i][j] = n
				continue
			}
			if f, err := strconv.ParseFloat(s, 64); err == nil {
				values[i][j] = f
				continue
			}
		}
	}
	return values
}

// --- sheet style ---

var sheetStyleCmd = &cobra.Command{
	Use:   "style <spreadsheet_token>",
	Short: "Apply style (e.g. bold) to a cell range",
	Long: `Apply formatting to a range of cells in a Lark spreadsheet.

Currently supports bold formatting and number/date format. Specify the range with --range and the sheet with --sheet.

Examples:
  lark sheet style T4mHsrFyzhXrj0tVzRslUGx8gkA --sheet abc123 --range A1:Q1 --bold
  lark sheet style T4mHsrFyzhXrj0tVzRslUGx8gkA --range H2:I122 --format "yyyy-MM-dd"
  lark sheet style T4mHsrFyzhXrj0tVzRslUGx8gkA --range J2:K122 --format "#,##0"`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		token := args[0]
		sheetID, _ := cmd.Flags().GetString("sheet")
		rangeSpec, _ := cmd.Flags().GetString("range")
		bold, _ := cmd.Flags().GetBool("bold")
		format, _ := cmd.Flags().GetString("format")

		if rangeSpec == "" {
			output.Fatal("MISSING_ARG", fmt.Errorf("--range is required"))
		}
		if !bold && format == "" {
			output.Fatal("MISSING_ARG", fmt.Errorf("at least one style flag is required (e.g. --bold, --format)"))
		}

		client := api.NewClient()

		sheetID = resolveSheetID(client, token, sheetID)

		style := api.SheetStyle{}
		if bold {
			style.Font = &api.SheetStyleFont{Bold: true}
		}
		if format != "" {
			style.Formatter = format
		}

		if err := client.SetSheetStyle(token, sheetID, rangeSpec, style); err != nil {
			output.Fatal("API_ERROR", err)
		}

		output.JSON(api.OutputSheetStyle{Success: true})
	},
}

// --- sheet resize ---

var sheetResizeCmd = &cobra.Command{
	Use:   "resize <spreadsheet_token>",
	Short: "Resize column widths in a sheet",
	Long: `Set pixel widths for columns in a Lark spreadsheet.

Provide widths as a JSON object mapping 0-based column index to pixel width,
or use --all to set a uniform width for all columns.

Examples:
  lark sheet resize T4mHsrFyzhXrj0tVzRslUGx8gkA --sheet abc123 --widths '{"0":60,"1":280,"2":200}'
  lark sheet resize T4mHsrFyzhXrj0tVzRslUGx8gkA --all 200 --cols 17`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		token := args[0]
		sheetID, _ := cmd.Flags().GetString("sheet")
		widthsJSON, _ := cmd.Flags().GetString("widths")
		allWidth, _ := cmd.Flags().GetInt("all")
		numCols, _ := cmd.Flags().GetInt("cols")

		client := api.NewClient()

		sheetID = resolveSheetID(client, token, sheetID)

		widths := map[int]int{}
		if widthsJSON != "" {
			var raw map[string]int
			if err := json.Unmarshal([]byte(widthsJSON), &raw); err != nil {
				output.Fatal("PARSE_ERROR", fmt.Errorf("invalid widths JSON: %w", err))
			}
			for k, v := range raw {
				
				idx, err := strconv.Atoi(k)
				if err != nil {
					output.Fatal("PARSE_ERROR", fmt.Errorf("invalid column index %q in --widths: must be a number", k))
				}
				widths[idx] = v
			}
		} else if cmd.Flags().Changed("all") {
			if numCols <= 0 {
				output.Fatal("MISSING_ARG", fmt.Errorf("--cols is required when using --all"))
			}
			for i := 0; i < numCols; i++ {
				widths[i] = allWidth
			}
		} else {
			output.Fatal("MISSING_ARG", fmt.Errorf("either --widths or --all with --cols is required"))
		}

		if err := client.SetSheetColumnWidths(token, sheetID, widths); err != nil {
			output.Fatal("API_ERROR", err)
		}

		output.JSON(api.OutputSheetResize{Success: true, ColumnsSet: len(widths)})
	},
}

// --- sheet add-tab ---

var sheetAddTabCmd = &cobra.Command{
	Use:   "add-tab <spreadsheet_token>",
	Short: "Add a new sheet tab to a spreadsheet",
	Long: `Add a new sheet tab to an existing Lark spreadsheet.

Examples:
  lark sheet add-tab T4mHsrFyzhXrj0tVzRslUGx8gkA --title "README"
  lark sheet add-tab T4mHsrFyzhXrj0tVzRslUGx8gkA --title "README" --index 0`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		token := args[0]
		title, _ := cmd.Flags().GetString("title")
		index, _ := cmd.Flags().GetInt("index")

		if title == "" {
			output.Fatal("MISSING_ARG", fmt.Errorf("--title is required"))
		}

		client := api.NewClient()

		result, err := client.AddSheetTab(token, title, index)
		if err != nil {
			output.Fatal("API_ERROR", err)
		}

		output.JSON(result)
	},
}

func init() {
	// Register subcommands
	sheetCmd.AddCommand(sheetListCmd)
	sheetCmd.AddCommand(sheetReadCmd)
	sheetCmd.AddCommand(sheetWriteCmd)
	sheetCmd.AddCommand(sheetCreateCmd)
	sheetCmd.AddCommand(sheetStyleCmd)
	sheetCmd.AddCommand(sheetResizeCmd)
	sheetCmd.AddCommand(sheetAddTabCmd)

	// Flags for sheet read
	sheetReadCmd.Flags().String("sheet", "", "Sheet ID to read from (default: first sheet)")
	sheetReadCmd.Flags().String("range", "", "Cell range to read (e.g., A1:Z100)")

	// Flags for sheet write
	sheetWriteCmd.Flags().String("sheet", "", "Sheet ID to write to (default: first sheet)")
	sheetWriteCmd.Flags().String("range", "", "Target range in A1 notation (e.g., A1:C3, default: A1)")
	sheetWriteCmd.Flags().String("values", "", "JSON array of arrays (e.g., '[[\"a\",\"b\"],[\"c\",\"d\"]]')")
	sheetWriteCmd.Flags().Bool("auto-type", false, "Auto-detect and convert date strings (YYYY-MM-DD) to serial numbers and numeric strings to numbers")

	// Flags for sheet create
	sheetCreateCmd.Flags().String("title", "", "Spreadsheet title (required)")
	sheetCreateCmd.Flags().String("folder", "", "Folder token to create the spreadsheet in (default: root)")

	// Flags for sheet style
	sheetStyleCmd.Flags().String("sheet", "", "Sheet ID (default: first sheet)")
	sheetStyleCmd.Flags().String("range", "", "Cell range in A1 notation (e.g., A1:Q1) - required")
	sheetStyleCmd.Flags().Bool("bold", false, "Apply bold formatting")
	sheetStyleCmd.Flags().String("format", "", `Number/date format string (e.g., "yyyy-MM-dd", "#,##0", "0.00%")`)

	// Flags for sheet resize
	sheetResizeCmd.Flags().String("sheet", "", "Sheet ID (default: first sheet)")
	sheetResizeCmd.Flags().String("widths", "", `JSON object mapping 0-based column index to pixel width, e.g. '{"0":60,"1":280}'`)
	sheetResizeCmd.Flags().Int("all", 0, "Set all columns to this pixel width (use with --cols)")
	sheetResizeCmd.Flags().Int("cols", 0, "Number of columns to resize when using --all")

	// Flags for sheet add-tab
	sheetAddTabCmd.Flags().String("title", "", "Tab title (required)")
	sheetAddTabCmd.Flags().Int("index", 0, "Tab position (0 = leftmost, default: 0)")
}
