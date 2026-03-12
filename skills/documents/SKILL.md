---
name: documents
description: Read and write Lark documents - get content as markdown or blocks, create new documents, append content (text, headings, lists, code), download images, list folders. Use when user asks about a Lark doc, wants to read/create/edit a document, or mentions a document URL/ID.
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

## Self-Diagnosis Before Doing Anything

Before attempting any doc command, verify the tool is available and authenticated. Run this first:

```bash
lark auth status
```

Use this to distinguish the three failure modes:

| Symptom | Cause | Fix |
|---------|-------|-----|
| `command not found` or binary missing | Tool not installed in this environment | Tell the user: "The lark CLI tool is not available in this environment." Do NOT attempt workarounds like asking them to paste content manually. |
| `AUTH_ERROR` or `access_token` is empty | Not logged in | Run `lark auth login` — this opens the browser automatically for the user to complete OAuth |
| `SCOPE_ERROR` on a command | Logged in but missing scope | Run `lark auth login --add --scopes documents` — opens the browser to re-authorize with the new scope |
| `API_ERROR (code 99991663)` on a specific file | Logged in and scoped, but no access to this file | Ask the file owner to share it (see Access Denied section below) |

**Key distinction**: "tool not available" and "not authenticated" are completely different problems. Never tell the user the tool can't access Lark when the real issue is the binary is missing, and vice versa.

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
      "url": "https://larksuite.com/docx/doxcntan34DX4QoKJu7jJyabcef",
      "created_time": "2024-03-17T01:54:45+08:00",
      "modified_time": "2026-03-02T13:07:17+08:00",
      "owner_id": "ou_xxxx"
    }
  ],
  "count": 1
}
```

Item types: `doc`, `docx`, `sheet`, `bitable`, `mindnote`, `file`, `folder`, `shortcut`

For shortcuts, includes `shortcut_info` with `target_type` and `target_token`.

### Get File/Folder Metadata

```bash
lark doc info <token> [--type <doc_type>]
```

Gets detailed metadata for a Lark Drive file or folder.

Options:
- `--type`: Document type (default: `file`). Supported: `doc`, `docx`, `sheet`, `bitable`, `folder`, `file`, `mindnote`, `slides`, `wiki`

Output:
```json
{
  "token": "Mbxmsn4eRha6ZXtqY9HlfVhsgab",
  "type": "sheet",
  "title": "API Task Inventory",
  "owner_id": "ou_xxxx",
  "create_time": "2026-03-02T13:07:16+08:00",
  "latest_modify_user": "ou_yyyy",
  "latest_modify_time": "2026-03-02T13:28:30+08:00",
  "url": "https://glints.sg.larksuite.com/sheets/Mbxmsn4eRha6ZXtqY9HlfVhsgab"
}
```

### Create Folder

```bash
lark doc mkdir <name> [--folder <parent_folder_token>]
```

Creates a new folder in Lark Drive. By default creates in the root cloud space.

Options:
- `--folder`: Parent folder token (default: root)

Output:
```json
{
  "token": "DyW7fcz0GlfQQwdODEIlw1TOgZh",
  "name": "Project Files",
  "url": "https://glints.sg.larksuite.com/drive/folder/DyW7fcz0GlfQQwdODEIlw1TOgZh"
}
```

### Upload File to Drive

```bash
lark doc upload <file_path> [--folder <folder_token>]
```

Uploads a local file to Lark Drive (max 20MB). Optionally specify a folder token to upload into.

Options:
- `--folder`: Folder token to upload into (default: root cloud space)

Output:
```json
{
  "file_token": "HjItbLxidogQ7BxQfbDlENvLg4f",
  "file_name": "report.pdf",
  "size": 1048576
}
```

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

### Create a New Document

```bash
lark doc create --title "My Document" [--folder <folder-token>]
```

Creates a new empty Lark document. Optionally specify a folder to create it in.

Options:
- `--title`: Document title (required)
- `--folder`: Folder token to create the document in (default: root cloud space)

Output:
```json
{
  "success": true,
  "document_id": "BKlmdClegoCan5x5Rzbl73QQgEC",
  "revision_id": 1,
  "title": "My Document"
}
```

### Append Content to a Document

```bash
lark doc append <document-id> [flags]
```

Appends content blocks to an existing document. Supports multiple block types.

Content flags (at least one required):
- `--text "content"`: Append a text paragraph
- `--heading "content" --level <1-9>`: Append a heading (default level 1)
- `--code "content" --language <id>`: Append a code block
- `--bullet "item"`: Append bullet list items (repeatable)
- `--ordered "item"`: Append ordered list items (repeatable)
- `--todo "item"`: Append a todo/checkbox item
- `--divider`: Append a horizontal divider
- `--quote "content"`: Append a quote/blockquote
- `--table-header "cell"`: Table header cells (repeatable, one per column)
- `--table-row "cell1|cell2"`: Table row as pipe-separated cells (repeatable)
- `--markdown`: Read markdown from stdin and convert to blocks
- `--json`: Read raw block JSON from stdin

Other flags:
- `--block-id`: Parent block ID to append under (default: document root)
- `--index`: Insertion position (-1=end, 0=beginning)
- `--after`: Insert after this block ID (mutually exclusive with `--index`)

Examples:
```bash
# Add a heading and text
lark doc append ABC123xyz --heading "Section Title" --level 2
lark doc append ABC123xyz --text "Hello from CLI"

# Add a bullet list
lark doc append ABC123xyz --bullet "First item" --bullet "Second item"

# Add a code block (Python)
lark doc append ABC123xyz --code "print('hello')" --language 49

# Add a quote block
lark doc append ABC123xyz --quote "This is important"

# Create a table (cells are auto-populated with content)
lark doc append ABC123xyz --table-header "Name" --table-header "Status" --table-row "API|Done" --table-row "CLI|WIP"

# Append from markdown (pipe from stdin)
echo '# Section Title

Paragraph text here.

- bullet 1
- bullet 2

> A blockquote

```python
print("hello")
```' | lark doc append ABC123xyz --markdown

# Pipe a markdown file
cat content.md | lark doc append ABC123xyz --markdown

# Add raw blocks via JSON stdin
echo '[{"block_type":2,"text":{"elements":[{"text_run":{"content":"raw"}}]}}]' | lark doc append ABC123xyz --json
```

Common code language IDs: 1=PlainText, 7=Bash, 22=Go, 29=Java, 30=JavaScript, 49=Python, 53=Rust, 56=SQL, 63=TypeScript, 67=YAML

Output:
```json
{
  "success": true,
  "document_revision_id": 5,
  "blocks": [...]
}
```

### Download Image from Document

```bash
lark doc image <image_token> --doc <document-id> [-o <output_path>]
```

Downloads an image from a Lark document. The image token is obtained from `doc blocks` output for image blocks (block_type 27).

Options:
- `-d, --doc`: Document ID that contains the image (required for authentication)
- `-o, --output`: Output file path (default: stdout - binary data)

Examples:
```bash
# Save image to file
lark doc image K1TQbpmDuokIq3xq1WVl9J7ygkc --doc ABC123xyz -o image.png

# Pipe to stdout (e.g., for piping to another command)
lark doc image K1TQbpmDuokIq3xq1WVl9J7ygkc --doc ABC123xyz > image.png
```

**Workflow:** Use `doc blocks` to find image blocks (block_type 27), extract the image token from the block data, then download with this command.

### Delete Blocks from a Document

```bash
lark doc delete <document-id> <block-id> [block-id...]
```

Deletes one or more blocks from a document. Block IDs can be found using `doc blocks`.

Examples:
```bash
lark doc delete ABC123xyz doxlgXYZ123
lark doc delete ABC123xyz doxlgXYZ123 doxlgABC456
```

Output:
```json
{
  "success": true,
  "document_revision_id": 15,
  "deleted_block_ids": ["doxlgXYZ123"]
}
```

### Update a Block in a Document

```bash
lark doc update <document-id> <block-id> [flags]
```

Updates the content of an existing block. Supports the same content flags as append: `--text`, `--heading`, `--code`, `--bullet`, `--ordered`, `--todo`, `--json`, `--link`.

**Note:** The Lark PATCH block API has strict format requirements and may return "invalid param" errors. If update fails, use `doc replace` instead (see below).

Examples:
```bash
lark doc update ABC123xyz doxlgXYZ123 --text "Updated content"
lark doc update ABC123xyz doxlgXYZ123 --heading "New Title" --level 2
lark doc update ABC123xyz doxlgXYZ123 --text "Click here" --link "https://example.com"
```

Output:
```json
{
  "success": true,
  "document_revision_id": 16
}
```

### Find Blocks by Content

```bash
lark doc find <document-id> <query> [--type <block-type>]
```

Searches block content for a text match (case-insensitive) and returns matching block IDs with their parent, index, type, and a content preview. Useful before using `doc replace` or `doc delete`.

Options:
- `--type`: Filter by block type number (e.g., 2=text, 12=bullet, 14=code, 17=todo)

Examples:
```bash
lark doc find ABC123xyz "Section Title"
lark doc find ABC123xyz "TODO" --type 17
```

Output:
```json
{
  "document_id": "ABC123xyz",
  "query": "Section Title",
  "matches": [
    {
      "block_id": "doxlgXYZ123",
      "parent_id": "ABC123xyz",
      "index": 5,
      "block_type": 4,
      "preview": "Section Title"
    }
  ],
  "count": 1
}
```

### Replace a Block

```bash
lark doc replace <document-id> <block-id> [content flags]
```

Atomically replaces a block: deletes it and inserts new content at the same position. Supports the same content flags as `append` (including `--quote`, `--markdown`, `--table-header`/`--table-row`). This is the recommended way to modify block content (instead of `doc update` which has API limitations).

Use `doc find` to locate the block ID first.

Examples:
```bash
# Find a block, then replace it
lark doc find ABC123xyz "Old Title"
lark doc replace ABC123xyz doxlgXYZ123 --heading "New Title" --level 2
lark doc replace ABC123xyz doxlgXYZ123 --text "Updated paragraph content"
echo '[{"block_type":14,"code":{"style":{"language":49},"elements":[{"text_run":{"content":"print(hello)","text_element_style":{}}}]}}]' | lark doc replace ABC123xyz doxlgXYZ123 --json
```

Output: Same format as `doc append`.

### Trash a File or Document

```bash
lark doc trash <file-token> [--type <doc-type>]
```

Moves a file or document to trash in Lark Drive. The file is moved to trash (not permanently deleted).

Options:
- `--type`: Document type (default: `docx`). Supported: `doc`, `docx`, `sheet`, `bitable`, `folder`, `file`, `mindnote`, `slides`

Examples:
```bash
lark doc trash ABC123xyz
lark doc trash ABC123xyz --type sheet
lark doc trash fldbcRho46N6... --type folder
```

Output:
```json
{
  "success": true,
  "file_token": "ABC123xyz",
  "type": "docx"
}
```

### Move a Block

```bash
lark doc move <document-id> <block-id> [--index <position>] [--after <block-id>]
```

Moves a block to a new position by deleting and reinserting it.

Options:
- `--index`: Target position index (-1=end, 0=beginning)
- `--after`: Insert after this block ID (mutually exclusive with `--index`)

Examples:
```bash
lark doc move ABC123xyz doxlgXYZ123 --index 0
lark doc move ABC123xyz doxlgXYZ123 --after doxlgABC456
```

### Show Document Outline

```bash
lark doc outline <document-id>
```

Returns only heading blocks (H1-H9) with block IDs, positions, and text. Faster than `doc blocks` for understanding document structure.

### Hyperlinks with --link

The `--link` flag adds a hyperlink URL to text content in `doc append` and `doc update`. It works with `--text`, `--heading`, `--bullet`, `--ordered`, `--todo`, and `--quote`.

Examples:
```bash
# Text with hyperlink
lark doc append ABC123xyz --text "Official docs" --link "https://docs.example.com"

# Bullet list item with hyperlink
lark doc append ABC123xyz --bullet "Read the guide" --link "https://example.com/guide"

# Heading with hyperlink
lark doc append ABC123xyz --heading "Project Page" --level 2 --link "https://example.com"
```

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
| Create a new document | `doc create` | Creates empty doc with title |
| Append content to doc | `doc append` | Add text, headings, lists, code, etc. |
| Read/summarize content | `doc get` | Markdown is compact (~90KB) |
| Analyze structure | `doc blocks` | Full block hierarchy |
| Search for text | `doc get` | Grep-able markdown |
| Count elements | `doc blocks` | Block types enumerated |
| Read comments/feedback | `doc comments` | Get all comments and replies |
| Download image from doc | `doc image` | Requires image token from `doc blocks` |
| Find blocks by content | `doc find` | Search blocks for text, get block IDs and indices |
| Replace a block | `doc replace` | Delete + insert in one step (preferred over `doc update`) |
| Delete blocks from doc | `doc delete` | Remove specific blocks by ID |
| Update a block | `doc update` | Modify existing block content (may fail, use `doc replace` instead) |
| Trash a file/document | `doc trash` | Move to trash in Drive |
| Move a block | `doc move` | Reorder blocks within a document |
| Get heading outline | `doc outline` | Fast structural overview (headings only) |
| Append from markdown | `doc append --markdown` | Pipe markdown, auto-converts to blocks |
| Create a table | `doc append --table-header/--table-row` | Creates table with auto-populated cells |

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
| 19 | Callout |
| 22 | Divider |
| 27 | Image |
| 31 | Table |
| 32 | Table cell |

## Output Format

All commands output JSON. Format appropriately when presenting to user.

## Authentication

Lark CLI uses OAuth 2.0 (browser-based flow). No API key auth is supported.

### Initial Setup

```bash
# Login with documents scope
lark auth login --scopes documents

# Or login with all default scopes, then add documents later
lark auth login
lark auth login --add --scopes documents
```

### Check Permissions

```bash
lark auth status
```

### Configuration

Config is stored in `~/.lark/config.yaml` with `app_id`. The app secret is stored in the `LARK_APP_SECRET` environment variable (or prompted during login).

### Available Scope Groups

- `calendar` - Calendar events and scheduling
- `contacts` - User and department lookups
- `documents` - Drive files and document operations
- `messages` - Chat messages and reactions
- `minutes` - Meeting recordings and transcripts

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

### Lark API Permission Errors (inside API_ERROR)

When you see `API_ERROR`, the message includes a Lark error code. Common ones:

| Lark Code | Meaning | What to tell the user |
|-----------|---------|----------------------|
| 99991663 | No read/write permission on this file | The document owner needs to share it with you. Ask them to open the doc → Share → add your Lark account with at least Viewer access. |
| 99991664 | No write permission (read-only access) | You have view-only access. Ask the owner to grant Editor access. |
| 99991400 | Invalid token / file does not exist | The document token may be wrong, or the file was deleted. Double-check the URL. |
| 99991401 | App not authorized to access tenant | The Lark app needs to be added to the workspace. Contact your Lark admin. |
| 99991668 | Rate limited | Too many requests. Wait a moment and retry. |

### When the User Gets Access Denied

If `doc get`, `doc blocks`, or any doc command returns a permission error, tell the user **exactly**:

1. **Who needs to act**: the document owner (or anyone with Editor+ access)
2. **What they need to do**: open the doc in Lark → click Share (top right) → add the user's Lark email → set role to at least Viewer
3. **Then retry**: the same command will work once access is granted - no re-login needed

Do NOT suggest copy-pasting content manually as a workaround - the lark-cli tool authenticates as the user and will work once access is properly shared.

## Required Permissions

This skill requires the `documents` scope group. If you see a `SCOPE_ERROR`, the user needs to add documents permissions:

```bash
lark auth login --add --scopes documents
```

To check current permissions:
```bash
lark auth status
```

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
