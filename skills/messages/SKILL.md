---
name: messages
description: Retrieve chat message history, send messages, and recall messages in Lark - get messages from group chats, private chats, threads, send messages to users or chats, and recall messages sent by the bot. Use when user asks about chat messages, conversation history, what was discussed in a group, wants to send a message, or wants to recall a message.
---

# Messages Skill

Retrieve chat message history, send messages, recall messages, and search for chats/groups via the `lark` CLI.

## 🤖 Agent Capabilities

**Message Sending Features:**
- Send text messages with line breaks and formatting
- Mention users in group chats
- Include clickable links
- Send to individual users or group chats
- Auto-detect recipient types (email, user ID, chat ID)
- Support for escape sequences (\n, \t, \", \\)

**Use Cases:**
- Send notifications and updates
- Share links and resources
- Mention team members in group discussions
- Communicate with users via email or direct message
- Format messages with proper line breaks and structure

**Agent-Friendly Features:**
- Explicit flag-based interface (no ambiguous parsing)
- Clear error messages and validation
- Comprehensive help text with examples
- Consistent JSON output for easy parsing

## 🚀 Quick Reference

**Send a simple message:**
```bash
lark msg send --to user@example.com --text "Hello!"
```

**Send with formatting:**
```bash
lark msg send --to oc_12345 --text "Update:\n• Task completed\n• Next: Review" --mention ou_user1
```

**Send with links:**
```bash
lark msg send --to ou_12345 --text "Check this out" --link "Dashboard" --url "https://example.com"
```

**Find chats:**
```bash
lark chat search "project team"
```

**Read messages:**
```bash
lark msg history --chat-id oc_12345 --limit 10
```

## Running Commands

Ensure `lark` is in your PATH, or use the full path to the binary. Set the config directory if not using the default:

```bash
lark msg <command>
lark chat <command>
# Or with explicit config:
LARK_CONFIG_DIR=/path/to/.lark lark msg <command>
```

## Commands Reference

### Search for Chats/Groups
```bash
# Search for chats by name
lark chat search "project"

# Search with Chinese characters
lark chat search "团队"

# List all visible chats (no query)
lark chat search

# Limit results
lark chat search "team" --limit 10
```

Available flags:
- `--limit`: Maximum number of chats to retrieve (0 = no limit)

Output:
```json
{
  "chats": [
    {
      "chat_id": "oc_d8e62a81cd188199ab994080f1e0804f",
      "name": "project-team",
      "description": "Project discussion group",
      "owner_id": "ou_46538f635a314d611cdd028c9c293d21",
      "external": false,
      "chat_status": "normal"
    }
  ],
  "count": 1,
  "query": "project"
}
```

The search supports:
- Group name matching (including internationalized names)
- Group member name matching
- Multiple languages
- Fuzzy search (pinyin, prefix, etc.)

### Send Messages

Send messages to users or group chats as the bot.

```bash
# Send simple text to user
lark msg send --to ou_xxxx --text "Hello!"

# Send to group chat
lark msg send --to oc_xxxx --text "Meeting starting soon"

# Text with line breaks (\n is interpreted as newline)
lark msg send --to ou_xxxx --text "Update:\nPhase 1 complete\nPhase 2 starting"

# Supported escape sequences: \n (newline), \t (tab), \\ (backslash), \" (quote)
lark msg send --to ou_xxxx --text "Tab:\tindented\nNew line"

# Mention users in group chat
lark msg send --to oc_xxxx --text "Please review" --mention ou_user1 --mention ou_user2

# With link
lark msg send --to ou_xxxx --text "Check this out" --link "Our Docs" --url "https://docs.example.com"

# Combined features
lark msg send --to oc_xxxx \
  --text "Project milestone reached!" \
  --mention ou_user1 --mention ou_user2 \
  --link "View Details" --url "https://project.example.com"

# Using explicit ID type
lark msg send --to user@example.com --to-type email --text "Hello"
```

Available flags:
- `--to` (required): Recipient identifier (user ID, open_id, email, or chat_id)
- `--to-type`: Explicitly specify ID type (`open_id`, `user_id`, `email`, `chat_id`) - auto-detected if omitted
- `--text` (required): Message text content
- `--mention`: User open_id to mention (repeatable for multiple mentions)
- `--link`: Link display text (must be used with `--url`)
- `--url`: Link URL (must be used with `--link`)

Output:
```json
{
  "success": true,
  "message_id": "om_dc13264520392913993dd051dba21dcf",
  "chat_id": "oc_xxxxx",
  "create_time": "2026-01-14T10:30:00+08:00"
}
```

## 🔍 Agent Use Cases

**When to use message sending:**
- Notify users about task completion or status updates
- Share links to documents, dashboards, or resources
- Mention team members in group discussions
- Send automated alerts or reminders
- Communicate with users via email or direct message
- Format messages with proper structure and line breaks

**Message formatting tips for agents:**
- Use `\n` for line breaks: `--text "Line 1\nLine 2"`
- Use `\t` for indentation: `--text "•\tItem 1\n•\tItem 2"`
- Use `\"` for quotes: `--text "\"Important:\" message"`
- Use `\\` for literal backslash: `--text "Path: C:\\\\folder\\\\file"`

**Recipient discovery:**
- Use `lark chat search "team"` to find group chats
- Use `lark contact get <user_id>` to verify user information
- Use email addresses directly (auto-detected)
- Use open_id format for reliable user targeting

**Recipient ID Types:**
- `open_id`: User open ID (starts with `ou_`)
- `user_id`: User ID (numeric)
- `email`: User email address
- `chat_id`: Group chat ID (starts with `oc_`)

The CLI auto-detects the ID type based on format, but you can override with `--to-type`.

**Message Types:**
- **Simple text**: Use only `--text` flag (supports `\n` for line breaks)
- **Rich text**: Automatically activated when using `--mention` or `--link`/`--url`

**Notes:**
- Messages are sent as the bot/app
- The bot must be added to group chats before it can send messages
- Requires `im:message` or `im:message:send_as_bot` permission scope

### Get Chat History
```bash
# Get messages from a group chat
lark msg history --chat-id oc_xxxxx

# Get messages with limit
lark msg history --chat-id oc_xxxxx --limit 50

# Get messages in a time range (Unix timestamp)
lark msg history --chat-id oc_xxxxx --start 1704067200 --end 1704153600

# Get messages in a time range (ISO 8601)
lark msg history --chat-id oc_xxxxx --start 2024-01-15 --end 2024-01-16

# Sort by newest first
lark msg history --chat-id oc_xxxxx --sort desc

# Get thread messages
lark msg history --chat-id thread_xxxxx --type thread
```

Available flags:
- `--chat-id` (required): Chat ID or thread ID
- `--type`: Container type - `chat` (default) or `thread`
- `--start`: Start time (Unix timestamp or ISO 8601)
- `--end`: End time (Unix timestamp or ISO 8601)
- `--sort`: Sort order - `asc` (default) or `desc`
- `--limit`: Maximum number of messages (0 = no limit)

Output:
```json
{
  "messages": [
    {
      "message_id": "om_dc13264520392913993dd051dba21dcf",
      "msg_type": "text",
      "content": "{\"text\":\"Hello world\"}",
      "sender": {
        "id": "ou_155184d1e73cbfb8973e5a9e698e74f2",
        "type": "user"
      },
      "create_time": "2024-01-15T09:00:00+08:00",
      "mentions": [],
      "is_reply": false,
      "thread_id": "",
      "deleted": false
    }
  ],
  "count": 1,
  "chat_id": "oc_xxxxx"
}
```

## Reading Thread Replies

When a message has a `thread_id` field, it means the message is part of a thread (or is the root of a thread with replies). To fetch all replies in that thread:

1. Get chat history and note the `thread_id` from any message that has one
2. Use that `thread_id` with `--type thread` to get all messages in the thread

Example workflow:
```bash
# Get recent messages from a chat
lark msg history --chat-id oc_xxxxx --limit 10 --sort desc

# If a message has thread_id: "omt_1a3b99f9d2cfd982", fetch the thread
lark msg history --chat-id omt_1a3b99f9d2cfd982 --type thread
```

Thread messages will have `is_reply: true` for replies (the root message has `is_reply: false`).

## Message Types

The `msg_type` field indicates the message format:
- `text` - Plain text message
- `post` - Rich text post
- `image` - Image
- `file` - File attachment
- `audio` - Audio message
- `media` - Video/media
- `sticker` - Sticker/emoji
- `interactive` - Interactive card
- `share_chat` - Shared chat
- `share_user` - Shared user contact

## Parsing Message Content

The `content` field is a JSON string. Parse it based on `msg_type`:

### Text Messages
```json
{"text": "Hello world"}
```

### Post Messages (Rich Text)
```json
{
  "title": "Post Title",
  "content": [[{"tag": "text", "text": "paragraph text"}]]
}
```

### Image Messages
```json
{"image_key": "img_xxxx"}
```

### File Messages
```json
{"file_key": "file_xxxx", "file_name": "document.pdf"}
```

### Audio Messages
```json
{"file_key": "file_xxxx", "duration": 5000}
```

### Media (Video) Messages
```json
{"file_key": "file_xxxx", "image_key": "img_xxxx", "file_name": "video.mp4", "duration": 10000}
```

## Downloading Resource Files

Download images, files, audio, and video from messages using `msg resource`:

```bash
# Download an image
lark msg resource --message-id om_xxx --file-key img_v3_xxx --type image --output ./image.png

# Download a file, audio, or video
lark msg resource --message-id om_xxx --file-key file_v2_xxx --type file --output ./document.pdf
```

Available flags:
- `--message-id` (required): Message ID containing the resource
- `--file-key` (required): Resource key from message content (`image_key` or `file_key`)
- `--type` (required): `image` for images, `file` for files/audio/video
- `--output` (required): Output file path

Output:
```json
{
  "success": true,
  "message_id": "om_xxx",
  "file_key": "img_v3_xxx",
  "output_path": "./image.png",
  "content_type": "image/png",
  "bytes_written": 62419
}
```

### Workflow: View Images from Chat

1. Get message history to find image messages:
```bash
lark msg history --chat-id oc_xxxxx --limit 20
```

2. Find messages with `msg_type: "image"` and parse the content to get `image_key`

3. Download the image:
```bash
lark msg resource --message-id om_xxx --file-key img_v3_xxx --type image --output /tmp/image.png
```

4. View the image using the Read tool on the downloaded file

### Limitations

- Maximum file size: 100MB
- Emoji resources cannot be downloaded
- Resources from card messages, merged messages, or forwarded messages cannot be downloaded (API error 234043)
- The `message_id` and `file_key` must match (the file must belong to that message)

## Recall Messages

Recall a message sent by the bot.

```bash
# Recall a message by ID
lark msg recall om_dc13264520392913993dd051dba21dcf
```

Output:
```json
{
  "success": true,
  "message": "Message recalled",
  "messageId": "om_dc13264520392913993dd051dba21dcf"
}
```

### Limitations and Requirements

- **Own messages only**: The bot can only recall its own messages
- **Time limit**: Messages must be within the configurable recall time limit (default: 24 hours)
- **Group admin override**: In group chats where the bot is owner/admin/creator, it can recall any message within 1 year
- **Batch messages**: Cannot recall messages sent via batch send API
- **Still in chat**: The bot must still be in the chat where the message was sent

## Integration with Contacts

Enrich sender information by looking up user details:

```bash
# Get message history
lark msg history --chat-id oc_xxxxx --limit 10

# Look up sender details
lark contact get ou_sender_id
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
- `SCOPE_ERROR` - Missing messages permissions. Run `lark auth login --add --scopes messages`
- `VALIDATION_ERROR` - Missing required fields (e.g., chat-id)
- `API_ERROR` - Lark API issue (e.g., bot not in group, missing permissions)

## Required Permissions

This skill requires the `messages` scope group. If you see a `SCOPE_ERROR`, the user needs to add messages permissions:

```bash
lark auth login --add --scopes messages
```

To check current permissions:
```bash
lark auth status
```

Additional requirements:

**For reading messages:**
- The bot must be in the group chat
- For group chat messages, the app needs the "Read all messages in associated group chat" permission scope
- Private chat messages only require `im:message:readonly` scope

**For sending messages:**
- Requires `im:message` or `im:message:send_as_bot` permission scope
- The bot must be added to group chats before it can send messages to them
- Can send to users directly without being in a private chat first

## Notes

- Chat IDs typically start with `oc_` (for chats) or `thread_` (for threads)
- Time filters don't work for thread container type
- Messages are sorted by creation time ascending by default
- The `deleted` field indicates if a message was recalled
