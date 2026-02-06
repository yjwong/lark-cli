---
name: documents
description: Query and retrieve Lark document content - get documents as markdown or structured blocks, list folder contents. Use when user asks about a Lark doc, wants to read a document, list files in a folder, or mentions a document URL/ID.
---

# Lark Documents Skill

Query and retrieve Lark document content via the `lark` CLI.

## Running Commands

Ensure `lark` is in your PATH, or use the full path to the binary. Set the config directory if not using the default:

```bash
lark doc <command>
# Or with explicit config:
LARK_CONFIG_DIR=/Users/yingcong/Code/lark-cli/.lark lark doc <command>
```

## Commands Reference

### Search Documents

```bash
lark doc search <query> [--owner <user-id>] [--chat <chat-id>] [--type <doc-type>]
```

Searches for documents by keyword. Returns documents from Drive that match the query.

Options:
- `--owner`: Filter by owner user ID (can be repeated)
- `--chat`: Filter by chat ID where document is located (can be repeated)
- `--type`: Filter by document type (can be repeated). Valid types: `doc`, `sheet`, `slide`, `bitable`, `mindnote`, `file`

Output:
```json
{
  "query": "project plan",
  "results": [
    {
      "token": "doxcntan34DX4QoKJu7jJyabcef",
      "type": "docx",
      "title": "Q4 Project Plan",
      "owner_id": "ou_xxxx"
    }
  ],
  "total": 15,
  "count": 15
}
```

**Note:** Maximum 200 results can be returned per search.

### List Folder Items

```bash
lark doc list [folder-token]
```

Lists items in a Lark Drive folder. If no folder token is provided, lists items in the root cloud space.

Output:
```json
{
  "folder_token": "fldbcRho46N6MQ3mJkOAuPabcef",
  "items": [
    {
      "token": "doxcntan34DX4QoKJu7jJyabcef",
      "name": "My Document",
      "type": "docx",
      "parent_token": "fldbcRho46N6MQ3mJkOAuPabcef",
      "url": "https://larksuite.com/docx/doxcntan34DX4QoKJu7jJyabcef"
    }
  ],
  "count": 1
}
```

Item types: `doc`, `docx`, `sheet`, `bitable`, `mindnote`, `file`, `folder`, `shortcut`

For shortcuts, includes `shortcut_info` with `target_type` and `target_token`.

### Download File from Drive

```bash
lark doc download <file_token> -o <output_path>
```

Downloads a file from Lark Drive. The file token is obtained from `doc list` or `doc search` output. Only files with type "file" can be downloaded (not docs, sheets, etc - those are Lark native formats).

Options:
- `-o, --output`: Output file path (required)

Output:
```json
{
  "file_token": "FG3obxWuaoftXIx0CvxlQAabcef",
  "filename": "report.pdf",
  "content_type": "application/pdf",
  "size": 1048576
}
```

**Note:** You must have read access to the file. If you get a 403 error, the file may not be shared with you.

### Resolve Wiki Node to Document ID

```bash
lark doc wiki <node-token>
```

Resolves a wiki node token to get the underlying document ID. **Required for wiki URLs before fetching content.**

Output:
```json
{
  "node_token": "X8Tawq431ifOYSklP2tlamKsgNh",
  "obj_token": "QdqJdegH8oAKmuxqGBBlKjGMgAc",
  "obj_type": "docx",
  "title": "Document Title",
  "space_id": "7384661266473713695",
  "node_type": "origin",
  "has_child": false
}
```

Use the `obj_token` value with `doc get` or `doc blocks`.

### List Wiki Node Children

```bash
lark doc wiki-children <node-token>
```

Lists the immediate child nodes of a wiki node. Useful for browsing wiki hierarchies (e.g., listing all PRDs under a PRDs folder).

Output:
```json
{
  "parent_node_token": "RBCmwZEqhili9ZkKS5fl1Ov2gKc",
  "space_id": "7344964278161604639",
  "children": [
    {
      "node_token": "ABC123",
      "obj_token": "XYZ789",
      "obj_type": "docx",
      "title": "PRD: Feature X",
      "space_id": "7344964278161604639",
      "node_type": "origin",
      "has_child": false
    }
  ],
  "count": 5
}
```

The command automatically resolves the space_id from the node token. Use `obj_token` with `doc get` to read child documents.

### Search Wiki Nodes

```bash
lark doc wiki-search <query> [--space-id <space-id>] [--node-id <node-id>]
```

Searches for wiki nodes by keyword. Returns wiki nodes the user has permission to view.

**Note:** Results are limited to 50 items to avoid API rate limits.

Options:
- `--space-id`: Filter to a specific wiki space ID
- `--node-id`: Search within a node and its children (requires `--space-id`)

Output:
```json
{
  "query": "meeting notes",
  "results": [
    {
      "node_id": "BAgPwq6lIi5Nykk0E5fcJeabcef",
      "obj_token": "AcnMdexrlokOShxe40Fc0Oabcef",
      "obj_type": "docx",
      "title": "Weekly Meeting Notes",
      "url": "https://sample.larksuite.com/wiki/BAgPwq6lIi5Nykk0E5fcJeabcef",
      "space_id": "7307457194084925443"
    }
  ],
  "count": 1
}
```

Use the `obj_token` value with `doc get` to retrieve the document content.

### Get Document as Markdown

```bash
lark doc get <document-id>
```

Returns document content as markdown - compact and readable.

Output:
```json
{
  "document_id": "ABC123xyz",
  "title": "My Document",
  "content": "# Heading\n\nDocument content as markdown..."
}
```

### Get Document Block Structure

```bash
lark doc blocks <document-id>
```

Returns full block structure as JSON - useful for programmatic analysis.

Output:
```json
{
  "document_id": "ABC123xyz",
  "title": "My Document",
  "block_count": 42,
  "blocks": [...]
}
```

### Get Document Comments

```bash
lark doc comments <document-id>
```

Retrieves all comments from a document, including replies.

Output:
```json
{
  "file_token": "ABC123xyz",
  "comments": [
    {
      "comment_id": "6916106822734512356",
      "user_id": "ou_cc19b2bfb93f8a44db4b4d6eab",
      "create_time": "2026-01-02T09:00:00+08:00",
      "is_solved": false,
      "is_whole": true,
      "quote": "The quoted text",
      "replies": [
        {
          "reply_id": "6916106822734512357",
          "user_id": "ou_cc19b2bfb93f8a44db4b4d6eab",
          "create_time": "2026-01-02T09:05:00+08:00",
          "text": "This is the reply text"
        }
      ]
    }
  ],
  "count": 1
}
```

Fields:
- `is_whole`: `true` for document-level comments, `false` for inline/local comments
- `is_solved`: whether the comment thread has been resolved
- `quote`: the highlighted text from the document (for inline comments)

## Spreadsheet Commands

### List Sheets in a Spreadsheet

```bash
lark sheet list <spreadsheet_token>
```

Lists all sheets (tabs) within a Lark spreadsheet.

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
    }
  ],
  "count": 1
}
```

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
    ["Value1", "Value2", 123, true]
  ]
}
```

**Note:** Cell values preserve their types (string, number, boolean). Empty cells may be omitted from rows.

## Extracting IDs from URLs

| URL Type | Example | How to Get Content |
|----------|---------|----------------------------|
| Direct doc | `https://xxx.larksuite.com/docx/ABC123xyz` | Use `ABC123xyz` directly with `doc get` |
| Wiki | `https://xxx.larksuite.com/wiki/X8Tawq43...` | First run `doc wiki X8Tawq43...` to get `obj_token`, then use that with `doc get` |
| Spreadsheet | `https://xxx.larksuite.com/sheets/T4mHsr...` | Use `T4mHsr...` with `sheet list` or `sheet read` |

**Important**: Wiki URLs require a two-step process - resolve the node first, then fetch the document.

## Which Command to Use

| Use Case | Command | Why |
|----------|---------|-----|
| Search for documents | `doc search` | Find docs by keyword across Drive |
| Search wiki by keyword | `doc wiki-search` | Find wiki nodes by keyword |
| List folder contents | `doc list [folder-token]` | Browse Drive files and folders |
| Download a file | `doc download` | Save Drive files locally |
| Wiki URL | `doc wiki` then `doc get` | Must resolve wiki node first |
| List wiki sub-pages | `doc wiki-children` | Browse wiki hierarchy |
| Read/summarize content | `doc get` | Markdown is compact (~90KB) |
| Analyze structure | `doc blocks` | Full block hierarchy |
| Search for text | `doc get` | Grep-able markdown |
| Count elements | `doc blocks` | Block types enumerated |
| Read comments/feedback | `doc comments` | Get all comments and replies |
| List sheets in spreadsheet | `sheet list` | See all tabs and their sizes |
| Read spreadsheet data | `sheet read` | Get cell values as JSON |

**Default to `doc get`** - it's 2-3x smaller and sufficient for most tasks.

## Block Types Reference

When using `doc blocks`, key block types:

| Type | Description |
|------|-------------|
| 1 | Page (root block) |
| 2 | Text paragraph |
| 3-11 | Headings H1-H9 |
| 12 | Bullet list |
| 13 | Ordered list |
| 14 | Code block |
| 15 | Quote |
| 17 | Todo/checkbox |
| 22 | Divider |
| 27 | Image |
| 31 | Table |

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

## Required Permissions

This skill requires the `documents` scope group. If you see a `SCOPE_ERROR`, the user needs to add documents permissions:

```bash
lark auth login --add --scopes documents
```

To check current permissions:
```bash
lark auth status
```

If you get an `API_ERROR` after confirming scopes are granted, the user may need to ensure they have access to the specific document in Lark.

## Efficient Extraction with jq and grep

For large documents, use `jq` and `grep` to extract specific information without loading everything into context.

### Get Document Outline (Headings Only)

```bash
# Extract all headings from markdown
lark doc get <id> | jq -r '.content' | grep -E '^#{1,6} '
```

### Get Document Title Only

```bash
lark doc get <id> | jq -r '.title'
```

### Count Blocks by Type

```bash
# Count how many of each block type
lark doc blocks <id> | jq '.blocks | group_by(.block_type) | map({type: .[0].block_type, count: length})'
```

### Extract All Todo Items

```bash
# Get all todo/checkbox blocks (type 17)
lark doc blocks <id> | jq '[.blocks[] | select(.block_type == 17)]'
```

### Search for Specific Content

```bash
# Find lines mentioning a keyword
lark doc get <id> | jq -r '.content' | grep -i "keyword"

# With context
lark doc get <id> | jq -r '.content' | grep -i -C 3 "keyword"
```

### Extract Code Blocks

```bash
# Get all code blocks (type 14)
lark doc blocks <id> | jq '[.blocks[] | select(.block_type == 14)]'
```

### Get Document Stats

```bash
# Quick stats: title, block count
lark doc blocks <id> | jq '{title, block_count}'
```

## Best Practices

### When User Shares a Document URL
1. Check the URL type:
   - `/docx/` URL: Extract document ID directly
   - `/wiki/` URL: Run `doc wiki <node-token>` first to get `obj_token`
2. Start with outline: `doc get <id> | jq -r '.content' | grep -E '^#{1,6} '`
3. Fetch full content only if needed for detailed analysis

### For Large Documents
- **Get outline first** - headings give structure without full content
- **Search with grep** - find relevant sections before loading everything
- **Use jq filters** - extract only what's needed
- Documents can be 50-100KB+ as markdown; be selective

### Integration with Other Skills
- After reading a document, you can create calendar events based on meeting notes
- Look up mentioned colleagues with the contacts skill
