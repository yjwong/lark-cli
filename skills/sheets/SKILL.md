---
name: sheets
description: Read and write Lark Sheets (spreadsheets) - list sheets, read cell data, write values, create spreadsheets, add tabs, apply formatting, resize columns. Use when user asks about a spreadsheet, wants to read/write data in a Lark sheet, or mentions a spreadsheet URL/ID.
---

# Lark Sheets Skill

Read, write, and manage Lark Sheets (spreadsheets) via the `lark` CLI.

## Tool Location

```
tools/bin/lark
```

Run from the vault root directory with required env var:
```bash
LARK_CONFIG_DIR=tools/lark/.lark tools/bin/lark sheet <command>
```

## Commands Reference

### List Sheets in a Spreadsheet

```bash
tools/bin/lark sheet list <spreadsheet_token>
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
tools/bin/lark sheet read <spreadsheet_token> [--sheet <sheet_id>] [--range A1:Z100]
```

Reads cell values from a Lark spreadsheet.

Options:
- `--sheet`: Sheet ID to read from (default: first sheet by index)
- `--range`: Cell range to read (e.g., `A1:Z100`). Default: all data up to 1000 rows
- `--render`: How to render cell values (default: `value`):
  - `value` - computed values; numbers stay numeric, strings stay string
  - `formula` - raw formula strings (e.g. `=SUM(A1:A10)`)
  - `formatted` - display strings with number/date formatting applied (e.g. `"1,234.56"`)

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

**Note:** Cell values preserve their types (string, number, boolean) with the default `--render value` mode. Use `--render formatted` to get display strings matching what the user sees in the sheet (e.g. `"-12%"` instead of `-0.116`), or `--render formula` to inspect raw formula expressions. Empty cells may appear as `null` or be omitted from rows. Some cells with rich formatting may return structured objects instead of plain values.

Cells containing file attachments return structured objects like:
```json
{
  "fileToken": "JwhmbXFVeoeIkixnmMKlguhzgLe",
  "mimeType": "application/pdf",
  "size": 930948,
  "text": "Contract.pdf",
  "type": "attachment"
}
```

Use the `fileToken` with `sheet download` to download these attachments.

### Download Cell Attachments

```bash
tools/bin/lark sheet download <file_token> --spreadsheet <spreadsheet_token> -o <output_path>
```

Downloads a file attachment embedded in a spreadsheet cell. The `file_token` is the `fileToken` from `sheet read` output.

Options:
- `--spreadsheet`: Spreadsheet token the attachment belongs to (required)
- `-o` / `--output`: Output file path (required)

Output:
```json
{
  "file_token": "JwhmbXFVeoeIkixnmMKlguhzgLe",
  "filename": "/tmp/contract.pdf",
  "content_type": "application/pdf",
  "size": 930948
}
```

### Write Cell Values

```bash
lark sheet write <spreadsheet_token> [--sheet <sheet_id>] [--range <A1:B2>] [--values '<JSON>'] [--auto-type]
```

Writes cell values to a Lark spreadsheet. Values are provided as a JSON array of arrays (rows of cells).

Options:
- `--sheet`: Sheet ID to write to (default: first sheet)
- `--range`: Target range in A1 notation (default: A1)
- `--values`: JSON array of arrays, e.g. `'[["a","b"],["c","d"]]'`. Can also be piped via stdin.
- `--auto-type`: Auto-detect and convert string values to proper types:
  - Date strings (`YYYY-MM-DD`) become Excel serial date numbers (for use with `--format "yyyy-MM-dd"`)
  - Integer strings become numbers
  - Float strings become numbers

**Important:** Lark Sheets stores dates as serial numbers (days since 1899-12-30). To display dates correctly, write with `--auto-type` then apply a date format with `sheet style --format "yyyy-MM-dd"`.

Examples:
```bash
# Write 2x2 data starting at A1
lark sheet write T4mHsrFyzhXrj0tVzRslUGx8gkA --values '[["hello","world"],["foo","bar"]]'

# Write with auto-type (converts "2025-01-15" to serial 45672, "42" to number)
lark sheet write T4mHsrFyzhXrj0tVzRslUGx8gkA --values '[["2025-01-15","42","text"]]' --auto-type

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
lark sheet style <spreadsheet_token> --range <A1:Z1> [--sheet <sheet_id>] [--bold] [--format <format_string>]
```

Applies formatting to a range of cells. Supports bold and number/date formatting.

Options:
- `--range`: Cell range in A1 notation (required, must be `A1:B1` format, not just `A1`)
- `--sheet`: Sheet ID (default: first sheet)
- `--bold`: Apply bold formatting
- `--format`: Number/date format string

Common format strings:
- `"yyyy-MM-dd"` - Date (2025-01-15)
- `"yyyy/MM/dd"` - Date with slashes
- `"yyyy/MM/dd HH:mm:ss"` - Datetime
- `"#,##0"` - Integer with thousands separator
- `"#,##0.00"` - Number with 2 decimals
- `"0%"` - Percentage
- `"0.00%"` - Percentage with decimals
- `"$#,##0"` - USD currency
- `"@"` - Plain text

Examples:
```bash
# Bold the header row
lark sheet style T4mHsrFyzhXrj0tVzRslUGx8gkA --range A1:Q1 --bold

# Format date columns
lark sheet style T4mHsrFyzhXrj0tVzRslUGx8gkA --range H2:I122 --format "yyyy-MM-dd"

# Format number columns
lark sheet style T4mHsrFyzhXrj0tVzRslUGx8gkA --range J2:K122 --format "#,##0"

# Bold and format together
lark sheet style T4mHsrFyzhXrj0tVzRslUGx8gkA --range A1:A1 --bold --format "#,##0"
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

The sheet_id can be found in the URL query parameter:
- URL: `https://xxx.larksuite.com/sheets/T4mHsrFyzhXrj0tVzRslUGx8gkA?sheet=abc123`
- Sheet ID: `abc123`

## Which Command to Use

| Use Case | Command | Notes |
|----------|---------|-------|
| Browse sheets/tabs | `sheet list` | See all sheets and dimensions |
| Read specific data | `sheet read --range` | Target specific cells |
| Read full sheet | `sheet read --sheet` | Up to 1000 rows |
| Read first sheet | `sheet read` | Auto-selects first by index |
| Read as user sees it | `sheet read --render formatted` | Display strings (e.g. "-12%" not -0.116) |
| Inspect formulas | `sheet read --render formula` | Raw formula expressions |
| Write data to cells | `sheet write --values` | JSON array of arrays |
| Create new spreadsheet | `sheet create` | Needs title; optional folder |
| Add a new tab | `sheet add-tab` | Requires title flag |
| Bold headers | `sheet style --bold` | Requires --range |
| Format dates/numbers | `sheet style --format` | Use with auto-typed values |
| Auto-type on write | `sheet write --auto-type` | Converts date/number strings |
| Adjust column widths | `sheet resize` | By index or uniformly |
| Download attachment | `sheet download` | From cell `fileToken` |

## Workflow Examples

### Get Overview of a Spreadsheet

```bash
# List all sheets first
LARK_CONFIG_DIR=tools/lark/.lark tools/bin/lark sheet list T4mHsrFyzhXrj0tVzRslUGx8gkA

# Then read specific sheet
LARK_CONFIG_DIR=tools/lark/.lark tools/bin/lark sheet read T4mHsrFyzhXrj0tVzRslUGx8gkA --sheet abc123 --range A1:D20
```

### Read Header Row Only

```bash
LARK_CONFIG_DIR=tools/lark/.lark tools/bin/lark sheet read T4mHsrFyzhXrj0tVzRslUGx8gkA --range A1:Z1
```

### Read First 50 Rows of Specific Sheet

```bash
LARK_CONFIG_DIR=tools/lark/.lark tools/bin/lark sheet read T4mHsrFyzhXrj0tVzRslUGx8gkA --sheet def456 --range A1:Z50
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
LARK_CONFIG_DIR=tools/lark/.lark tools/bin/lark sheet read <token> --range A1:Z1 | jq '.values[0]'
```

### Get Row Count

```bash
LARK_CONFIG_DIR=tools/lark/.lark tools/bin/lark sheet read <token> | jq '.row_count'
```

### Extract Specific Column (Column B)

```bash
LARK_CONFIG_DIR=tools/lark/.lark tools/bin/lark sheet read <token> | jq '[.values[] | .[1]]'
```

### Find Rows Matching a Value

```bash
LARK_CONFIG_DIR=tools/lark/.lark tools/bin/lark sheet read <token> | jq '[.values[] | select(.[0] == "SearchValue")]'
```

### Get First N Rows

```bash
LARK_CONFIG_DIR=tools/lark/.lark tools/bin/lark sheet read <token> | jq '.values[:10]'
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

### CLI Error Codes

- `AUTH_ERROR` - Not logged in. Run `lark auth login`
- `SCOPE_ERROR` - Missing OAuth scope. Run `lark auth login --add --scopes documents`
- `API_ERROR` - Lark API rejected the request (see message for Lark error code)
- `NO_SHEETS` - Spreadsheet has no sheets

### Lark API Permission Errors (inside API_ERROR)

When you see `API_ERROR`, the message includes a Lark error code. Common ones:

| Lark Code | Meaning | What to tell the user |
|-----------|---------|----------------------|
| 99991663 | No read/write permission on this file | The sheet owner needs to share it with you. Ask them to go to the sheet → Share → add your Lark account with at least Viewer access. |
| 99991664 | No write permission (read-only access) | You have view-only access. Ask the owner to grant Editor access. |
| 99991400 | Invalid token / file does not exist | The spreadsheet token in the URL may be wrong, or the file was deleted. Double-check the URL. |
| 99991401 | App not authorized to access tenant | The Lark app needs to be added to the workspace. Contact your Lark admin. |
| 99991668 | Rate limited | Too many requests. Wait a moment and retry. |
| 1254043  | Sheet range invalid | The `--range` value is out of bounds for this sheet. Use `sheet list` to check the sheet's actual dimensions. |

### When the User Gets Access Denied

If `sheet read` or `sheet list` returns a permission error, tell the user **exactly**:

1. **Who needs to act**: the spreadsheet owner (or anyone with Editor+ access)
2. **What they need to do**: open the sheet in Lark → click Share (top right) → add the user's Lark email → set role to at least Viewer
3. **Then retry**: the same command will work once access is granted - no re-login needed

Do NOT suggest exporting to CSV, pasting values manually, or using a public link as workarounds - the lark-cli tool authenticates as the user and will work once access is properly shared.

## Required Permissions

This skill requires the `documents` scope group (uses `drive:drive:readonly` for reads, `sheets:spreadsheet` for writes). If you see a `SCOPE_ERROR`, the user needs to add documents permissions:

```bash
LARK_CONFIG_DIR=tools/lark/.lark tools/bin/lark auth login --add --scopes documents
```

To check current permissions:
```bash
LARK_CONFIG_DIR=tools/lark/.lark tools/bin/lark auth status
```

## Limitations

- Maximum 1000 rows read by default (use `--range` for specific cells)
- Column letters limited to A-Z (26 columns) when auto-detecting range
- Rich text cells may return structured objects instead of plain strings
- Some merged cells may have unexpected value placement
- `sheet style --format` requires the range to use `A1:B1` notation (not just `A1`)
