package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/yjwong/lark-cli/internal/api"
	"github.com/yjwong/lark-cli/internal/browser"
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

var singleCellRangeRe = regexp.MustCompile(`^[A-Za-z]+[0-9]+$`)

func normalizeStyleRange(rangeSpec string) string {
	rangeSpec = strings.TrimSpace(rangeSpec)
	if singleCellRangeRe.MatchString(rangeSpec) {
		return rangeSpec + ":" + rangeSpec
	}
	return rangeSpec
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
		render, _ := cmd.Flags().GetString("render")

		// Map friendly flag values to Lark API values
		renderOptionMap := map[string]string{
			"":          "UnformattedValue",
			"value":     "UnformattedValue",
			"formula":   "Formula",
			"formatted": "FormattedValue",
		}
		renderOption, ok := renderOptionMap[render]
		if !ok {
			output.Fatal("INVALID_ARG", fmt.Errorf("invalid --render value %q: must be value, formula, or formatted", render))
		}

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
		data, err := client.GetSheetData(token, fullRange, renderOption)
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

// --- sheet validation ---

var sheetValidationCmd = &cobra.Command{
	Use:   "validation <spreadsheet_token>",
	Short: "Query dropdown validation settings from a sheet range",
	Long: `Query dropdown validation settings from a range in a Lark spreadsheet.

By default, queries dropdown ("list") validations from the first sheet. Use --sheet
to specify a sheet ID and --range to narrow the query range.

Examples:
  lark sheet validation T4mHsrFyzhXrj0tVzRslUGx8gkA
  lark sheet validation T4mHsrFyzhXrj0tVzRslUGx8gkA --sheet abc123 --range M:N
  lark sheet validation T4mHsrFyzhXrj0tVzRslUGx8gkA --range A1:B100 --type list`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		token := args[0]
		sheetID, _ := cmd.Flags().GetString("sheet")
		rangeSpec, _ := cmd.Flags().GetString("range")
		dataValidationType, _ := cmd.Flags().GetString("type")

		client := api.NewClient()

		sheetID = resolveSheetID(client, token, sheetID)

		fullRange := sheetID
		if rangeSpec != "" {
			fullRange = sheetID + "!" + rangeSpec
		}

		data, err := client.GetSheetDataValidation(token, fullRange, dataValidationType)
		if err != nil {
			output.Fatal("API_ERROR", err)
		}

		result := api.OutputSheetValidation{
			SpreadsheetToken:   token,
			SheetID:            sheetID,
			Range:              fullRange,
			DataValidationType: dataValidationType,
		}
		if result.DataValidationType == "" {
			result.DataValidationType = "list"
		}
		if data != nil {
			result.SpreadsheetToken = data.SpreadsheetToken
			result.SheetID = data.SheetID
			result.Revision = data.Revision
			result.DataValidations = data.DataValidations
			result.Count = len(data.DataValidations)
		}

		output.JSON(result)
	},
}

var sheetSetValidationCmd = &cobra.Command{
	Use:   "set-validation <spreadsheet_token>",
	Short: "Create a dropdown validation rule for a sheet range",
	Long: `Create a dropdown validation rule for a range in a Lark spreadsheet.

Provide dropdown items with --values as a JSON array or a comma-separated list.
Optionally provide one color per item with --colors.

Examples:
  lark sheet set-validation T4mHsrFyzhXrj0tVzRslUGx8gkA --sheet abc123 --range L2:L200 --values '["Migrate","Deprecate"]'
  lark sheet set-validation T4mHsrFyzhXrj0tVzRslUGx8gkA --range M2:M200 --values 'Y,N' --colors '#bacefd,#fed4a4'`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		token := args[0]
		sheetID, _ := cmd.Flags().GetString("sheet")
		rangeSpec, _ := cmd.Flags().GetString("range")
		valuesArg, _ := cmd.Flags().GetString("values")
		colorsArg, _ := cmd.Flags().GetString("colors")
		multiple, _ := cmd.Flags().GetBool("multiple")
		noHighlight, _ := cmd.Flags().GetBool("no-highlight")

		if rangeSpec == "" {
			output.Fatal("MISSING_ARG", fmt.Errorf("--range is required"))
		}
		if valuesArg == "" {
			output.Fatal("MISSING_ARG", fmt.Errorf("--values is required"))
		}

		values, err := parseStringListArg(valuesArg)
		if err != nil {
			output.Fatal("PARSE_ERROR", fmt.Errorf("invalid --values: %w", err))
		}
		if len(values) == 0 {
			output.Fatal("INVALID_ARG", fmt.Errorf("--values must contain at least one dropdown item"))
		}

		var colors []string
		if colorsArg != "" {
			colors, err = parseStringListArg(colorsArg)
			if err != nil {
				output.Fatal("PARSE_ERROR", fmt.Errorf("invalid --colors: %w", err))
			}
			if len(colors) != len(values) {
				output.Fatal("INVALID_ARG", fmt.Errorf("--colors must have the same number of items as --values"))
			}
		}

		client := api.NewClient()
		sheetID = resolveSheetID(client, token, sheetID)
		fullRange := sheetID + "!" + rangeSpec

		options := &api.SheetDataValidationOptions{
			MultipleValues:     multiple,
			HighlightValidData: !noHighlight,
		}
		if len(colors) > 0 {
			options.Colors = colors
		}

		if err := client.CreateSheetDataValidation(token, fullRange, values, options); err != nil {
			output.Fatal("API_ERROR", err)
		}

		data, err := client.GetSheetDataValidation(token, fullRange, "list")
		if err != nil {
			output.Fatal("API_ERROR", err)
		}

		result := api.OutputSheetValidationSet{
			Success:          true,
			SpreadsheetToken: token,
			SheetID:          sheetID,
			Range:            fullRange,
		}
		if data != nil {
			result.SpreadsheetToken = data.SpreadsheetToken
			result.SheetID = data.SheetID
			result.DataValidations = data.DataValidations
			for _, validation := range data.DataValidations {
				for _, validationRange := range validation.Ranges {
					if validationRange == fullRange {
						v := validation
						result.DataValidation = &v
						break
					}
				}
				if result.DataValidation != nil {
					break
				}
			}
		}

		output.JSON(result)
	},
}

var sheetClearValidationCmd = &cobra.Command{
	Use:   "clear-validation <spreadsheet_token>",
	Short: "Delete dropdown validation rules from a sheet range",
	Long: `Delete dropdown validation rules that overlap a range in a Lark spreadsheet.

Examples:
  lark sheet clear-validation T4mHsrFyzhXrj0tVzRslUGx8gkA --sheet abc123 --range L2:L200`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		token := args[0]
		sheetID, _ := cmd.Flags().GetString("sheet")
		rangeSpec, _ := cmd.Flags().GetString("range")

		if rangeSpec == "" {
			output.Fatal("MISSING_ARG", fmt.Errorf("--range is required"))
		}

		client := api.NewClient()
		sheetID = resolveSheetID(client, token, sheetID)
		fullRange := sheetID + "!" + rangeSpec

		data, err := client.GetSheetDataValidation(token, fullRange, "list")
		if err != nil {
			output.Fatal("API_ERROR", err)
		}

		if data == nil || len(data.DataValidations) == 0 {
			output.JSON(api.OutputSheetValidationClear{
				Success:          true,
				SpreadsheetToken: token,
				SheetID:          sheetID,
				Range:            fullRange,
			})
			return
		}

		deleteRanges := make([]api.SheetDeleteDataValidationRange, 0, len(data.DataValidations))
		for _, validation := range data.DataValidations {
			for _, validationRange := range validation.Ranges {
				deleteRanges = append(deleteRanges, api.SheetDeleteDataValidationRange{
					SheetID:          sheetID,
					DataValidationID: validation.DataValidationID,
					Range:            validationRange,
				})
			}
		}

		deleteData, err := client.DeleteSheetDataValidation(token, deleteRanges)
		if err != nil {
			output.Fatal("API_ERROR", err)
		}

		result := api.OutputSheetValidationClear{
			Success:          true,
			SpreadsheetToken: token,
			SheetID:          sheetID,
			Range:            fullRange,
		}
		if deleteData != nil {
			result.Results = deleteData.RangeResults
		}
		output.JSON(result)
	},
}

func parseStringListArg(raw string) ([]string, error) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return nil, nil
	}

	if strings.HasPrefix(raw, "[") {
		var values []string
		if err := json.Unmarshal([]byte(raw), &values); err != nil {
			return nil, err
		}
		return trimNonEmpty(values), nil
	}

	parts := strings.Split(raw, ",")
	return trimNonEmpty(parts), nil
}

func trimNonEmpty(values []string) []string {
	result := make([]string, 0, len(values))
	for _, value := range values {
		trimmed := strings.TrimSpace(value)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
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

		// Convert booleans to 1/0 (Lark API doesn't accept bool cell type)
		for i, row := range values {
			for j, cell := range row {
				if b, ok := cell.(bool); ok {
					if b {
						values[i][j] = 1
					} else {
						values[i][j] = 0
					}
				}
			}
		}

		autoType, _ := cmd.Flags().GetBool("auto-type")
		if autoType {
			values = autoTypeValues(values)
		}

		// Build the range string
		if rangeSpec == "" {
			// Auto-calculate range from values dimensions
			rows := len(values)
			cols := 0
			for _, row := range values {
				if len(row) > cols {
					cols = len(row)
				}
			}
			if cols == 0 {
				cols = 1
			}
			endCol := columnIndexToLetter(cols)
			rangeSpec = fmt.Sprintf("A1:%s%d", endCol, rows)
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
		wrap, _ := cmd.Flags().GetBool("wrap")
		noWrap, _ := cmd.Flags().GetBool("no-wrap")

		if rangeSpec == "" {
			output.Fatal("MISSING_ARG", fmt.Errorf("--range is required"))
		}
		rangeSpec = normalizeStyleRange(rangeSpec)
		if wrap && noWrap {
			output.Fatal("INVALID_ARG", fmt.Errorf("--wrap and --no-wrap are mutually exclusive"))
		}
		if !bold && format == "" && !wrap && !noWrap {
			output.Fatal("MISSING_ARG", fmt.Errorf("at least one style flag is required (e.g. --bold, --format, --wrap, --no-wrap)"))
		}

		client := api.NewClient()

		sheetID = resolveSheetID(client, token, sheetID)

		style := api.SheetStyle{}
		applyViaAPI := false
		if bold {
			style.Font = &api.SheetStyleFont{Bold: true}
			applyViaAPI = true
		}
		if format != "" {
			style.Formatter = format
			applyViaAPI = true
		}
		if noWrap {
			cleanValue := true
			style.Clean = &cleanValue
			applyViaAPI = true
		}

		var data *api.SheetStyleUpdateData
		if applyViaAPI {
			var err error
			data, err = client.SetSheetStyle(token, sheetID, rangeSpec, style)
			if err != nil {
				output.Fatal("API_ERROR", err)
			}
		}
		if wrap {
			if err := browser.WrapSheetRange(token, rangeSpec); err != nil {
				output.Fatal("BROWSER_ERROR", fmt.Errorf("failed to persist wrapped text via browser fallback: %w", err))
			}
		}

		result := api.OutputSheetStyle{Success: true}
		if wrap && data == nil {
			result.UpdatedRange = sheetID + "!" + rangeSpec
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

// --- sheet download ---

var sheetDownloadCmd = &cobra.Command{
	Use:   "download <file_token>",
	Short: "Download an attachment from a spreadsheet cell",
	Long: `Download a file attachment embedded in a Lark spreadsheet cell.

The file_token is obtained from 'sheet read' output. When a cell contains
an attachment, it appears as a structured object with type "attachment"
and a "fileToken" field.

You must specify the spreadsheet token with --spreadsheet so the API
can authenticate the download.

Examples:
  lark sheet download JwhmbXFVeoeIkixnmMKlguhzgLe --spreadsheet WgKQsHcLDhfUM9tqo11lXPYLghh -o contract.pdf
  lark sheet download JwhmbXFVeoeIkixnmMKlguhzgLe --spreadsheet WgKQsHcLDhfUM9tqo11lXPYLghh -o ~/Downloads/report.pdf`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fileToken := args[0]
		outputPath, _ := cmd.Flags().GetString("output")
		spreadsheetToken, _ := cmd.Flags().GetString("spreadsheet")

		if outputPath == "" {
			output.Fatal("MISSING_ARG", fmt.Errorf("--output/-o flag is required"))
		}
		if spreadsheetToken == "" {
			output.Fatal("MISSING_ARG", fmt.Errorf("--spreadsheet flag is required"))
		}

		client := api.NewClient()

		reader, contentType, err := client.DownloadMedia(fileToken, spreadsheetToken)
		if err != nil {
			output.Fatal("API_ERROR", err)
		}
		defer reader.Close()

		file, err := os.Create(outputPath)
		if err != nil {
			output.Fatal("FILE_ERROR", err)
		}
		defer file.Close()

		written, err := io.Copy(file, reader)
		if err != nil {
			output.Fatal("IO_ERROR", err)
		}

		result := struct {
			FileToken   string `json:"file_token"`
			Filename    string `json:"filename"`
			ContentType string `json:"content_type"`
			Size        int64  `json:"size"`
		}{
			FileToken:   fileToken,
			Filename:    outputPath,
			ContentType: contentType,
			Size:        written,
		}
		output.JSON(result)
	},
}

func init() {
	// Register subcommands
	sheetCmd.AddCommand(sheetListCmd)
	sheetCmd.AddCommand(sheetReadCmd)
	sheetCmd.AddCommand(sheetValidationCmd)
	sheetCmd.AddCommand(sheetSetValidationCmd)
	sheetCmd.AddCommand(sheetClearValidationCmd)
	sheetCmd.AddCommand(sheetWriteCmd)
	sheetCmd.AddCommand(sheetCreateCmd)
	sheetCmd.AddCommand(sheetStyleCmd)
	sheetCmd.AddCommand(sheetResizeCmd)
	sheetCmd.AddCommand(sheetAddTabCmd)
	sheetCmd.AddCommand(sheetDownloadCmd)

	// Flags for sheet read
	sheetReadCmd.Flags().String("sheet", "", "Sheet ID to read from (default: first sheet)")
	sheetReadCmd.Flags().String("range", "", "Cell range to read (e.g., A1:Z100)")
	sheetReadCmd.Flags().String("render", "value", "How to render cell values: value (computed, default), formula (raw =formula), formatted (display string)")

	// Flags for sheet validation
	sheetValidationCmd.Flags().String("sheet", "", "Sheet ID to query from (default: first sheet)")
	sheetValidationCmd.Flags().String("range", "", "Cell range to query (e.g., M:N or A1:B100)")
	sheetValidationCmd.Flags().String("type", "list", "Validation type to query (currently only: list)")

	// Flags for sheet set-validation
	sheetSetValidationCmd.Flags().String("sheet", "", "Sheet ID to write to (default: first sheet)")
	sheetSetValidationCmd.Flags().String("range", "", "Cell range to apply validation to (e.g., L2:L200)")
	sheetSetValidationCmd.Flags().String("values", "", `Dropdown items as JSON array or comma-separated string, e.g. '["A","B"]' or 'A,B'`)
	sheetSetValidationCmd.Flags().String("colors", "", `Optional colors as JSON array or comma-separated string, e.g. '["#bacefd","#fed4a4"]'`)
	sheetSetValidationCmd.Flags().Bool("multiple", false, "Allow selecting multiple dropdown values")
	sheetSetValidationCmd.Flags().Bool("no-highlight", false, "Disable colored dropdown highlighting")

	// Flags for sheet clear-validation
	sheetClearValidationCmd.Flags().String("sheet", "", "Sheet ID to clear from (default: first sheet)")
	sheetClearValidationCmd.Flags().String("range", "", "Cell range to clear validation from (e.g., L2:L200)")

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
	sheetStyleCmd.Flags().Bool("wrap", false, "Enable text wrapping for the range")
	sheetStyleCmd.Flags().Bool("no-wrap", false, "Disable text wrapping for the range")

	// Flags for sheet resize
	sheetResizeCmd.Flags().String("sheet", "", "Sheet ID (default: first sheet)")
	sheetResizeCmd.Flags().String("widths", "", `JSON object mapping 0-based column index to pixel width, e.g. '{"0":60,"1":280}'`)
	sheetResizeCmd.Flags().Int("all", 0, "Set all columns to this pixel width (use with --cols)")
	sheetResizeCmd.Flags().Int("cols", 0, "Number of columns to resize when using --all")

	// Flags for sheet add-tab
	sheetAddTabCmd.Flags().String("title", "", "Tab title (required)")
	sheetAddTabCmd.Flags().Int("index", 0, "Tab position (0 = leftmost, default: 0)")

	// Flags for sheet download
	sheetDownloadCmd.Flags().StringP("output", "o", "", "Output file path (required)")
	sheetDownloadCmd.Flags().String("spreadsheet", "", "Spreadsheet token the attachment belongs to (required)")
}
