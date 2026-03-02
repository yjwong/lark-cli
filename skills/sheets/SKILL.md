---
name: sheets
description: Read and write Lark Sheets (spreadsheets) - list sheets, read cell data, write values, create spreadsheets, add tabs, apply formatting, resize columns. Use when user asks about a spreadsheet, wants to read/write data in a Lark sheet, or mentions a spreadsheet URL/ID.
---

# Lark Sheets Skill

Read, write, and manage Lark Sheets (spreadsheets) via the `lark` CLI.

## Running Commands

Ensure `lark` is in your PATH, or use the full path to the binary. Set the config directory if not using the default:

```bash
lark sheet <command>
# Or with explicit config:
LARK_CONFIG_DIR=/Users/yingcong/Code/lark-cli/.lark lark sheet <command>
```

## Commands Reference

### List Sheets in a Spreadsheet

```bash
lark sheet list <spreadsheet_token>
```

Lists all sheets (tabs) within a Lark spreadsheet. Returns sheet IDs, titles, dimensions, and hidden status.

Output:
```json
{
  "spreadsheet_token": "T4mHsrFyzhXrj0tVzRslUGx8gkA",
  "sheets": [
    {
      "sheet_id": "abc123",
      "title": "Sheet1",
      "index": 0,
      "row_count": 100,
      "column_count": 10
    },
    {
      "sheet_id": "def456",
      "title": "Sheet2",
      "index": 1,
      "hidden": true,
      "row_count": 50,
      "column_count": 5
    }
  ],
  "count": 2
}
```

Fields:
- `sheet_id`: The unique ID of the sheet (use this with `--sheet` flag)
- `title`: Display name of the sheet
- `index`: Position of the sheet (0-indexed)
- `hidden`: Whether the sheet is hidden in the UI
- `row_count` / `column_count`: Dimensions of the sheet

### Read Sheet Data

```bash
lark sheet read <spreadsheet_token> [--sheet <sheet_id>] [--range A1:Z100]
```

Reads cell values from a Lark spreadsheet.

Options:
- `--sheet`: Sheet ID to read from (default: first sheet by index)
- `--range`: Cell range to read (e.g., `A1:Z100`). Default: all data up to 1000 rows

Output:
```json
{
  "spreadsheet_token": "T4mHsrFyzhXrj0tVzRslUGx8gkA",
  "sheet_id": "abc123",
  "range": "abc123!A1:D10",
  "row_count": 10,
  "column_count": 4,
  "values": [
    ["Header1", "Header2", "Header3", "Header4"],
    ["Value1", "Value2", 123, true],
    ["Row2Val1", null, 456, false]
  ]
}
```

**Note:** Cell values preserve their types (string, number, boolean). Empty cells may appear as `null` or be omitted from rows. Some cells with rich formatting may return structured objects instead of plain values.

### Write Cell Values

```bash
lark sheet write <spreadsheet_token> [--sheet <sheet_id>] [--range <A1:B2>] [--values '<JSON>']
```

Writes cell values to a Lark spreadsheet. Values are provided as a JSON array of arrays (rows of cells).

Options:
- `--sheet`: Sheet ID to write to (default: first sheet)
- `--range`: Target range in A1 notation (default: A1)
- `--values`: JSON array of arrays, e.g. `'[["a","b"],["c","d"]]'`. Can also be piped via stdin.

Examples:
```bash
# Write 2x2 data starting at A1
lark sheet write T4mHsrFyzhXrj0tVzRslUGx8gkA --values '[["hello","world"],["foo","bar"]]'

# Write to specific sheet and range
lark sheet write T4mHsrFyzhXrj0tVzRslUGx8gkA --sheet abc123 --range A1:B2 --values '[["a","b"],["c","d"]]'

# Pipe values from stdin
echo '[["a","b"]]' | lark sheet write T4mHsrFyzhXrj0tVzRslUGx8gkA
```

### Create a New Spreadsheet

```bash
lark sheet create --title "My Spreadsheet" [--folder <folder_token>]
```

Creates a new Lark spreadsheet. Optionally specify a folder to create it in.

Options:
- `--title`: Spreadsheet title (required)
- `--folder`: Folder token to create the spreadsheet in (default: root cloud space)

Example:
```bash
lark sheet create --title "My Spreadsheet"
lark sheet create --title "My Spreadsheet" --folder fldbcRho46N6MQ3mJkOAuPabcef
```

### Add a New Sheet Tab

```bash
lark sheet add-tab <spreadsheet_token> --title "Tab Name" [--index <position>]
```

Adds a new sheet tab to an existing spreadsheet.

Options:
- `--title`: Tab title (required)
- `--index`: Tab position, 0 = leftmost (default: 0)

Examples:
```bash
lark sheet add-tab T4mHsrFyzhXrj0tVzRslUGx8gkA --title "README"
lark sheet add-tab T4mHsrFyzhXrj0tVzRslUGx8gkA --title "README" --index 0
```

### Apply Cell Formatting (Style)

```bash
lark sheet style <spreadsheet_token> --range <A1:Z1> [--sheet <sheet_id>] [--bold]
```

Applies formatting to a range of cells. Currently supports bold formatting.

Options:
- `--range`: Cell range in A1 notation (required)
- `--sheet`: Sheet ID (default: first sheet)
- `--bold`: Apply bold formatting

Examples:
```bash
# Bold the header row on a specific sheet
lark sheet style T4mHsrFyzhXrj0tVzRslUGx8gkA --sheet abc123 --range A1:Q1 --bold

# Bold header row using default (first) sheet
lark sheet style T4mHsrFyzhXrj0tVzRslUGx8gkA --range A1:Z1 --bold
```

### Resize Column Widths

```bash
lark sheet resize <spreadsheet_token> [--sheet <sheet_id>] [--widths '<JSON>'] [--all <px> --cols <n>]
```

Sets pixel widths for columns in a spreadsheet.

Options:
- `--sheet`: Sheet ID (default: first sheet)
- `--widths`: JSON object mapping 0-based column index to pixel width
- `--all`: Set all columns to this pixel width (use with `--cols`)
- `--cols`: Number of columns to resize when using `--all`

Examples:
```bash
# Set specific column widths (0-based index)
lark sheet resize T4mHsrFyzhXrj0tVzRslUGx8gkA --sheet abc123 --widths '{"0":60,"1":280,"2":200}'

# Set all 17 columns to 200px wide
lark sheet resize T4mHsrFyzhXrj0tVzRslUGx8gkA --all 200 --cols 17
```

## Extracting IDs from URLs

The spreadsheet_token is from the spreadsheet URL:
- URL: `https://xxx.larksuite.com/sheets/T4mHsrFyzhXrj0tVzRslUGx8gkA`
- Token: `T4mHsrFyzhXrj0tVzRslUGx8gkA`

## Which Command to Use

| Use Case | Command | Notes |
|----------|---------|-------|
| Browse sheets/tabs | `sheet list` | See all sheets and dimensions |
| Read specific data | `sheet read --range` | Target specific cells |
| Read full sheet | `sheet read --sheet` | Up to 1000 rows |
| Read first sheet | `sheet read` | Auto-selects first by index |
| Write data to cells | `sheet write --values` | JSON array of arrays |
| Create new spreadsheet | `sheet create` | Needs title; optional folder |
| Add a new tab | `sheet add-tab` | Requires title flag |
| Bold headers | `sheet style --bold` | Requires --range |
| Adjust column widths | `sheet resize` | By index or uniformly |

## Workflow Examples

### Get Overview of a Spreadsheet

```bash
# List all sheets first
lark sheet list T4mHsrFyzhXrj0tVzRslUGx8gkA

# Then read specific sheet
lark sheet read T4mHsrFyzhXrj0tVzRslUGx8gkA --sheet abc123 --range A1:D20
```

### Read Header Row Only

```bash
lark sheet read T4mHsrFyzhXrj0tVzRslUGx8gkA --range A1:Z1
```

### Read First 50 Rows of Specific Sheet

```bash
lark sheet read T4mHsrFyzhXrj0tVzRslUGx8gkA --sheet def456 --range A1:Z50
```

### Create a Spreadsheet and Populate It

```bash
# 1. Create the spreadsheet
lark sheet create --title "Weekly Report"

# 2. Write headers (use the token from the create output)
lark sheet write <token> --values '[["Name","Status","Date"]]'

# 3. Bold the header row
lark sheet style <token> --range A1:C1 --bold

# 4. Resize columns
lark sheet resize <token> --widths '{"0":150,"1":100,"2":120}'

# 5. Add a second tab
lark sheet add-tab <token> --title "Details"
```

### Write Data from a Script

```bash
# Pipe JSON from a script
echo '[["Alice","Done","2026-03-01"],["Bob","In Progress","2026-03-02"]]' | \
  lark sheet write T4mHsrFyzhXrj0tVzRslUGx8gkA --sheet abc123 --range A2
```

## Efficient Extraction with jq

For large spreadsheets, use `jq` to extract specific data without loading everything into context.

### Get Column Headers

```bash
lark sheet read <token> --range A1:Z1 | jq '.values[0]'
```

### Get Row Count

```bash
lark sheet read <token> | jq '.row_count'
```

### Extract Specific Column (Column B)

```bash
lark sheet read <token> | jq '[.values[] | .[1]]'
```

### Find Rows Matching a Value

```bash
lark sheet read <token> | jq '[.values[] | select(.[0] == "SearchValue")]'
```

### Get First N Rows

```bash
lark sheet read <token> | jq '.values[:10]'
```

## Output Format

All commands output JSON. Format appropriately when presenting to user.

## Error Handling

Errors return JSON:
```json
{
  "error": true,
  "code": "ERROR_CODE",
  "message": "Description"
}
```

Common error codes:
- `AUTH_ERROR` - Need to run `lark auth login`
- `SCOPE_ERROR` - Missing documents permissions. Run `lark auth login --add --scopes documents`
- `API_ERROR` - Lark API issue (often permissions)
- `NO_SHEETS` - Spreadsheet has no sheets

## Required Permissions

This skill requires the `documents` scope group (uses `drive:drive:readonly` for reads, `sheets:spreadsheet` for writes). If you see a `SCOPE_ERROR`, the user needs to add documents permissions:

```bash
lark auth login --add --scopes documents
```

To check current permissions:
```bash
lark auth status
```

## Limitations

- Maximum 1000 rows read by default (use `--range` for specific cells)
- Column letters limited to A-Z (26 columns) when auto-detecting range
- Rich text cells may return structured objects instead of plain strings
- Some merged cells may have unexpected value placement
- `sheet style` currently only supports bold formatting
