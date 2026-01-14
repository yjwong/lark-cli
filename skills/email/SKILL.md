---
name: email
description: Read and search emails from Lark Mail via IMAP with local caching. Use when user asks about email, inbox, or messages.
---

# Email Management Skill

Read and search emails from Lark Mail via the `lark` CLI using IMAP with local caching.

## Setup

Before using mail commands, IMAP credentials must be configured:

```bash
lark mail setup
```

This will prompt for IMAP credentials. See: https://www.larksuite.com/hc/en-US/articles/378111206512-log-in-to-lark-mail-through-a-third-party-email-client

## Running Commands

Ensure `lark` is in your PATH, or use the full path to the binary. Set the config directory if not using the default:

```bash
lark mail <command>
# Or with explicit config:
LARK_CONFIG_DIR=/path/to/.lark lark mail <command>
```

## Commands Reference

### Check Status
```bash
lark mail status
```

Shows connection status and cache freshness. The `freshness` field indicates how stale the cached data is.

### List Mailboxes
```bash
lark mail list
```

### Sync Emails
Fetch new emails from the server into the local cache:

```bash
# Sync INBOX (default)
lark mail sync

# Sync specific mailbox
lark mail sync --mailbox Sent
```

**Important**: Run sync before searching if you need fresh data.

### Search Emails
Search the local cache (fast, no network calls):

```bash
# List recent emails
lark mail search

# Filter by sender
lark mail search --from alice@example.com

# Filter by subject
lark mail search --subject "Q4 Report"

# Filter by date range
lark mail search --since 2025-01-01 --before 2025-01-15

# Combined filters
lark mail search --from boss@example.com --since 2025-01-01 --limit 10

# Search different mailbox
lark mail search --mailbox Sent
```

### View Email Content
```bash
lark mail show --uid <uid>
```

The UID is obtained from search results.

### Download as .eml
```bash
lark mail fetch --uid <uid>
lark mail fetch --uid <uid> --output ./emails/
```

## Output Formats

All commands output JSON.

### mail status Output
```json
{
  "configured": true,
  "host": "imap.larksuite.com",
  "port": 993,
  "username": "user@example.com",
  "connection": "ok",
  "cache": {
    "last_sync": "2025-01-14T10:30:00+08:00",
    "freshness": "15 minutes ago",
    "uidvalidity": 12345,
    "last_uid": 4521
  }
}
```

### mail search Output
```json
{
  "mailbox": "INBOX",
  "last_sync": "2025-01-14T10:30:00+08:00",
  "freshness": "15 minutes ago",
  "total_cached": 1523,
  "results": [
    {
      "uid": 4521,
      "message_id": "<abc123@mail.example.com>",
      "date": "2025-01-14T09:15:00+08:00",
      "from_addr": "alice@example.com",
      "from_name": "Alice",
      "subject": "Q4 Report"
    }
  ],
  "count": 1
}
```

### mail show Output
```json
{
  "uid": 4521,
  "from": {
    "email": "alice@example.com",
    "name": "Alice"
  },
  "subject": "Q4 Report",
  "date": "2025-01-14T09:15:00+08:00",
  "message_id": "<abc123@mail.example.com>",
  "body": "Hi,\n\nPlease find the Q4 report attached...\n\nBest,\nAlice"
}
```

### mail sync Output
```json
{
  "mailbox": "INBOX",
  "new_messages": 5,
  "total_cached": 1523,
  "message": "synced 5 new messages"
}
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
- `CONNECTION_ERROR` - IMAP connection failed (check credentials)
- `SYNC_ERROR` - Failed to sync emails
- `SEARCH_ERROR` - Cache query failed
- `VALIDATION_ERROR` - Missing required fields (e.g., --uid)
- `IO_ERROR` - File system error

## Typical Workflow

1. **Check cache freshness**:
   ```bash
   lark mail status
   ```

2. **Sync if needed** (based on freshness):
   ```bash
   lark mail sync
   ```

3. **Search emails**:
   ```bash
   lark mail search --from boss@example.com --since 2025-01-01
   ```

4. **View email details** using UID from search:
   ```bash
   lark mail show --uid 4521
   ```

5. **Download if needed**:
   ```bash
   lark mail fetch --uid 4521
   ```

