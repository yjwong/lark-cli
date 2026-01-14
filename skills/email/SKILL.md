---
name: email
description: Read emails from Lark Mail - list messages, view email details, get attachment download URLs. Use when user asks about email, inbox, or messages.
---

# Email Management Skill

Read emails from Lark Mail via the `lark` CLI.

## Running Commands

Ensure `lark` is in your PATH, or use the full path to the binary. Set the config directory if not using the default:

```bash
lark email <command>
# Or with explicit config:
LARK_CONFIG_DIR=/path/to/.lark lark email <command>
```

## Commands Reference

### Check Authentication
```bash
lark auth status
```

### List Email IDs
```bash
# List emails from INBOX (default)
lark email list

# List only unread emails
lark email list --unread

# Specify folder
lark email list --folder INBOX

# Limit results per page (1-20)
lark email list --page-size 10

# Fetch all emails (pagination handled automatically)
lark email list --all

# Specify mailbox (default: "me" for current user)
lark email list --mailbox me
```

Returns a list of message IDs that can be used with `email show` to get full details.

### View Email Details
```bash
# Get full email content
lark email show --id <message_id>

# Example
lark email show --id ZWEyNGRmY2QtOTVlNy00NzlhLTg5MjItMjFjOTQ5ZjIzZjJl
```

Returns email metadata (subject, from, to, cc), decoded body text, and attachment info.

### Get Attachment Download URLs
```bash
# Get download URLs for all attachments in an email
lark email attachments --id <message_id>

# Example
lark email attachments --id ZWEyNGRmY2QtOTVlNy00NzlhLTg5MjItMjFjOTQ5ZjIzZjJl
```

**Important**: Download URLs expire after 2 hours and are limited to 2 uses each.

## Output Formats

All commands output JSON:

### email list Output
```json
{
  "message_ids": ["id1", "id2", "id3"],
  "count": 3,
  "has_more": true,
  "mailbox_id": "me"
}
```

### email show Output
```json
{
  "message_id": "ZWEy...",
  "subject": "Meeting Tomorrow",
  "from": {
    "email": "sender@example.com",
    "name": "John Doe"
  },
  "to": [
    {"email": "recipient@example.com", "name": "Jane Smith"}
  ],
  "cc": [],
  "date": "2024-01-15T10:30:00+08:00",
  "body": "Plain text content of the email...",
  "thread_id": "abc123",
  "attachments": [
    {
      "id": "att1",
      "filename": "document.pdf",
      "type": "standard",
      "is_inline": false
    }
  ]
}
```

### email attachments Output
```json
{
  "message_id": "ZWEy...",
  "mailbox_id": "me",
  "downloads": [
    {
      "attachment_id": "att1",
      "filename": "document.pdf",
      "download_url": "https://..."
    }
  ],
  "failed_ids": []
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
- `AUTH_ERROR` - Need to run `lark auth login`
- `NOT_FOUND` - Email or mailbox not found
- `VALIDATION_ERROR` - Missing required fields (e.g., --id)
- `API_ERROR` - Lark API issue

## Typical Workflow

1. **List emails** to get message IDs:
   ```bash
   lark email list --unread
   ```

2. **View email details** for a specific message:
   ```bash
   lark email show --id <message_id>
   ```

3. **Download attachments** if needed:
   ```bash
   lark email attachments --id <message_id>
   # Then use the download URLs to fetch files
   ```

## Required Scopes

The email commands require these Lark API scopes:
- `mail:user_mailbox.message:readonly` - Core email access
- `mail:user_mailbox.message.subject:read` - Email subjects
- `mail:user_mailbox.message.address:read` - From/To/CC fields
- `mail:user_mailbox.message.body:read` - Email body and attachments

If some fields are missing from the response, ensure the app has the required scopes configured.
