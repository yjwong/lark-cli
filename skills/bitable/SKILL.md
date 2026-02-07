---
name: bitable
description: Access Lark Bitable databases - list tables, view fields, and read records. Use when user asks about a Bitable, database, or wants to query structured data.
---

# Lark Bitable Skill

Access Lark Bitable (database) content via the `lark` CLI.

## Running Commands

Ensure `lark` is in your PATH, or use the full path to the binary:

```bash
lark bitable <command>
```

## Commands Reference

### List Tables

```bash
lark bitable tables <app_token>
```

Lists all tables in a Bitable app. The app_token is from the Bitable URL.

For example, if the URL is `https://xxx.larksuite.com/base/ABC123xyz`, the app_token is `ABC123xyz`.

Output:
```json
{
  "app_token": "ABC123xyz",
  "tables": [
    {
      "table_id": "tblXYZ789",
      "name": "Projects"
    },
    {
      "table_id": "tblABC456",
      "name": "Tasks"
    }
  ],
  "count": 2
}
```

### List Fields

```bash
lark bitable fields <app_token> <table_id>
```

Lists all fields (columns) in a Bitable table.

Output:
```json
{
  "app_token": "ABC123xyz",
  "table_id": "tblXYZ789",
  "fields": [
    {
      "field_id": "fldAAA111",
      "field_name": "Name",
      "type": "text",
      "is_primary": true
    },
    {
      "field_id": "fldBBB222",
      "field_name": "Status",
      "type": "select"
    },
    {
      "field_id": "fldCCC333",
      "field_name": "Due Date",
      "type": "date"
    }
  ],
  "count": 3
}
```

Field types: `text`, `number`, `select`, `multi_select`, `date`, `checkbox`, `person`, `phone`, `url`, `attachment`, `link`, `formula`, `duplex_link`, `location`, `group`, `created_time`, `created_user`, `modified_time`, `modified_user`, `auto_number`

### List Records

```bash
lark bitable records <app_token> <table_id> [--limit N] [--view <view_id>] [--filter <expression>]
```

Lists records (rows) in a Bitable table.

Options:
- `--limit`: Maximum number of records to retrieve (default: no limit)
- `--view`: View ID to filter records
- `--filter`: Filter expression (see Lark API docs for syntax)

Output:
```json
{
  "app_token": "ABC123xyz",
  "table_id": "tblXYZ789",
  "records": [
    {
      "record_id": "recAAA111",
      "fields": {
        "Name": "Project Alpha",
        "Status": "In Progress",
        "Due Date": 1704067200000
      }
    },
    {
      "record_id": "recBBB222",
      "fields": {
        "Name": "Project Beta",
        "Status": "Completed",
        "Due Date": 1703462400000
      }
    }
  ],
  "count": 2,
  "has_more": false
}
```

**Note:** Date fields return Unix timestamps in milliseconds.

## Extracting IDs from URLs

| URL Type | Example | How to Extract |
|----------|---------|----------------|
| Bitable app | `https://xxx.larksuite.com/base/ABC123xyz` | app_token is `ABC123xyz` |
| Specific table | `https://xxx.larksuite.com/base/ABC123xyz?table=tblXYZ789` | app_token is `ABC123xyz`, table_id is `tblXYZ789` |

## Workflow Example

```bash
# 1. List all tables in the Bitable
lark bitable tables ABC123xyz

# 2. See what fields a table has
lark bitable fields ABC123xyz tblXYZ789

# 3. Read records from the table
lark bitable records ABC123xyz tblXYZ789 --limit 100
```

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
- `SCOPE_ERROR` - Missing bitable permissions. Run `lark auth login --add --scopes bitable`
- `API_ERROR` - Lark API issue (often permissions on the specific Bitable)

## Required Permissions

This skill requires the `bitable` scope group with the `bitable:app:readonly` scope. If you see a `SCOPE_ERROR`, add the permissions:

```bash
lark auth login --add --scopes bitable
```

To check current permissions:
```bash
lark auth status
```

## Best Practices

### Reading Large Tables
- Use `--limit` to avoid fetching too many records at once
- For analysis, fetch fields first to understand the schema
- Filter by view if possible to reduce data transfer

### Working with Field Values
- Text fields return strings
- Number fields return floats
- Date fields return Unix timestamps in milliseconds
- Select fields return the selected option text
- Person fields return user IDs (use `contact get` to resolve names)
- Attachment fields return file metadata arrays
