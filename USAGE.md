# lark Usage Guide

A CLI tool for interacting with Lark APIs (calendar, contacts, documents). Designed for use by Claude Code.

## Setup

### 1. Create a Lark App

1. Go to https://open.larksuite.com and create a custom app
2. Enable these permissions:
   - `calendar:calendar` (read/write calendar and events)
   - `contact:contact.base:readonly` (read contacts)
   - `contact:department.base:readonly` (read departments)
   - `docx:document:readonly` (read documents)
   - `docs:document.content:read` (read document content)
   - `wiki:wiki:readonly` (read wiki nodes)
   - `space:document:retrieve` (list Drive folder contents)
   - `im:message:readonly` (read messages in chats)
   - `im:message` or `im:message:send_as_bot` (send messages)
   - `offline_access` (for refresh tokens)
3. Add redirect URI: `http://localhost:9999/callback`
4. Enable "Refresh user_access_token" in Security Settings
5. Note your App ID and App Secret

### 2. Configure

Copy the example config:
```bash
cp config.example.yaml .lark/config.yaml
```

Edit `.lark/config.yaml`:
```yaml
app_id: "cli_xxxxxxxxxx"  # Your App ID
```

Set your app secret as environment variable:
```bash
export LARK_APP_SECRET="your_app_secret"
```

### 3. Authenticate

```bash
./lark auth login
```

This opens a browser for OAuth authorization.

## Commands

All commands output JSON by default.

### Authentication

```bash
# Login with all permissions (default)
./lark auth login

# Login with specific scope groups only
./lark auth login --scopes calendar           # Only calendar permissions
./lark auth login --scopes calendar,contacts  # Calendar and contacts

# Add permissions incrementally (without losing existing ones)
./lark auth login --add --scopes messages

# Check authentication status (shows granted scopes)
./lark auth status
# Output: {"authenticated": true, "expires_at": "...", "granted_groups": ["calendar", "contacts"], ...}

# List available scope groups
./lark auth scopes

# Logout (clear stored tokens)
./lark auth logout
```

#### Scope Groups

| Group | Commands | Description |
|-------|----------|-------------|
| `calendar` | `cal *` | Calendar events and scheduling |
| `contacts` | `contact *` | Company directory lookup |
| `documents` | `doc *` | Lark Docs and Drive access |
| `messages` | `msg *`, `chat *` | Chat and messaging |
| `mail` | `mail *` | Email via IMAP |
| `minutes` | `minutes *` | Meeting recordings |

By default, `lark auth login` requests all scopes. Use `--scopes` for minimal permissions.

### Calendar

#### List Events

```bash
# Today's events (default)
./lark cal list

# This week
./lark cal list --week

# Custom date range (ISO 8601)
./lark cal list --from 2026-01-02 --to 2026-01-05

# Events awaiting your RSVP
./lark cal list --pending --from 2026-01-02 --to 2026-01-31

# With conflict detection
./lark cal list --week --detect-conflicts --buffer-minutes 15
```

Output format:
```json
{
  "events": [
    {
      "id": "abc123",
      "summary": "Team standup",
      "start": "2026-01-02T09:00:00+08:00",
      "end": "2026-01-02T09:30:00+08:00",
      "location": "Meeting Room A",
      "meeting_url": "https://..."
    }
  ],
  "count": 1
}
```

#### Show Event

```bash
./lark cal show <event-id>
```

#### Create Event

```bash
# With explicit end time
./lark cal create \
  --summary "Team standup" \
  --start "2026-01-03T09:00:00+08:00" \
  --end "2026-01-03T09:30:00+08:00" \
  --location "Meeting Room A" \
  --description "Daily sync"

# With duration
./lark cal create --summary "1:1" --start "2026-01-03T14:00:00+08:00" --duration 30m

# With attendees
./lark cal create \
  --summary "Sync" \
  --start "2026-01-03T14:00:00+08:00" \
  --duration 1h \
  --attendee user1@example.com \
  --attendee user2@example.com
```

Flags:
- `--summary` (required): Event title
- `--start` (required): Start time (ISO 8601)
- `--end`: End time
- `--duration`: Duration (e.g., `30m`, `1h`, `1h30m`)
- `--location`: Event location
- `--description`: Event description
- `--attendee`: Attendee email (can be repeated)
- `--reminder`: Minutes before event to remind
- `--visibility`: `default`, `public`, or `private`
- `--no-notify`: Don't send notifications

#### Update Event

```bash
./lark cal update <event-id> --summary "New title"
./lark cal update <event-id> --start "2026-01-03T10:00:00+08:00"
./lark cal update <event-id> --location "New location"
./lark cal update <event-id> --visibility public
```

Flags:
- `--summary`: Event title
- `--description`: Event description
- `--start`: Start time (ISO 8601)
- `--end`: End time
- `--location`: Event location
- `--color`: Event color (hex format, e.g., `#9CA2A9`)
- `--visibility`: `default`, `public`, or `private`
- `--no-notify`: Don't send notifications

#### Delete Event

```bash
./lark cal delete <event-id>
```

#### Search Events

```bash
./lark cal search "standup"
./lark cal search "1:1" --from 2026-01-01 --to 2026-01-31
```

#### Query Availability (Free/Busy)

```bash
# Check your own availability
./lark cal freebusy --from 2026-01-03T09:00:00+08:00 --to 2026-01-03T18:00:00+08:00

# Check another user's availability
./lark cal freebusy --from 2026-01-03 --to 2026-01-03 --user ou_xxxxxxxxxx

# Check a meeting room's availability
./lark cal freebusy --from 2026-01-06 --to 2026-01-10 --room omm_xxxxxxxxxx
```

#### RSVP to Event

```bash
# Accept an invitation
./lark cal rsvp <event-id> --accept

# Decline an invitation
./lark cal rsvp <event-id> --decline

# Mark as tentative
./lark cal rsvp <event-id> --tentative
```

### Contacts

#### Get User by ID

```bash
# Look up by open_id (default)
./lark contact get ou_xxxx

# Look up by user_id
./lark contact get 12345 --id-type user_id
```

#### List Users in Department

```bash
# List users in root department
./lark contact list-dept

# List users in specific department
./lark contact list-dept od_xxxx
```

#### Search Departments

```bash
./lark contact search-dept "Engineering"
```

### Messages

#### Get Chat History

```bash
# Get messages from a group chat
./lark msg history --chat-id oc_xxxxx

# Get messages with limit
./lark msg history --chat-id oc_xxxxx --limit 50

# Get messages in a time range (Unix timestamp)
./lark msg history --chat-id oc_xxxxx --start 1704067200 --end 1704153600

# Get messages in a time range (ISO 8601)
./lark msg history --chat-id oc_xxxxx --start 2026-01-02 --end 2026-01-03

# Sort by newest first
./lark msg history --chat-id oc_xxxxx --sort desc

# Get thread messages
./lark msg history --chat-id thread_xxxxx --type thread
```

Flags:
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
      "create_time": "2026-01-02T09:00:00+08:00",
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

Message types: `text`, `post`, `image`, `file`, `audio`, `media`, `sticker`, `interactive`, `share_chat`, `share_user`

**Note:** The bot must be in the group chat. For group messages, the app must have the "Read all messages in associated group chat" permission scope.

#### Download Message Resource

Download resource files (images, videos, audios, files) from messages.

```bash
# Download an image from a message
./lark msg resource --message-id om_xxx --file-key img_v2_xxx --type image --output ./image.png

# Download a file/video/audio from a message
./lark msg resource --message-id om_xxx --file-key file_v2_xxx --type file --output ./video.mp4
```

Flags:
- `--message-id` (required): Message ID containing the resource
- `--file-key` (required): Resource key (from message content JSON)
- `--type` (required): `image` for images, `file` for files/audio/video
- `--output` (required): Output file path

Output:
```json
{
  "success": true,
  "message_id": "om_xxx",
  "file_key": "img_v2_xxx",
  "output_path": "./image.png",
  "content_type": "image/png",
  "bytes_written": 12345
}
```

**Note:** The file_key can be found in the `content` field of messages returned by `lark msg history`. The maximum downloadable file size is 100MB. Emoji resources cannot be downloaded.

#### Send Message

Send messages to users or group chats as the bot.

```bash
# Send simple text to user
./lark msg send --to ou_xxxx --text "Hello!"

# Send to group chat
./lark msg send --to oc_xxxx --text "Meeting starting soon"

# Text with line breaks (\n creates actual newlines)
./lark msg send --to ou_xxxx --text "Line 1\nLine 2\nLine 3"

# For literal backslash-n, use double backslash
./lark msg send --to ou_xxxx --text "Show literal \\n text"

# Mention users (automatically uses rich text format)
./lark msg send --to oc_xxxx --text "Please review" --mention ou_user1 --mention ou_user2

# With link
./lark msg send --to ou_xxxx --text "Check this out" --link "Our Docs" --url "https://docs.example.com"

# Combined features
./lark msg send --to oc_xxxx \
  --text "Project milestone reached!" \
  --mention ou_user1 --mention ou_user2 \
  --link "View Details" --url "https://project.example.com"
```

Flags:
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

**Note:** Messages are sent as the bot/app. The bot must be added to group chats before it can send messages to them.

#### Recall Message

Recall a message sent by the bot.

```bash
# Recall a message by ID
./lark msg recall om_dc13264520392913993dd051dba21dcf
```

Output:
```json
{
  "success": true,
  "message": "Message recalled",
  "messageId": "om_dc13264520392913993dd051dba21dcf"
}
```

**Limitations:**
- The bot can only recall its own messages
- Messages must have been sent within the configurable time limit (default: 24 hours)
- For group chats, if the bot is a group owner/admin/creator, it can recall any message within 1 year
- Cannot recall messages sent via batch send API
- The bot must still be in the chat where the message was sent

### Documents

#### List Folder Items

```bash
# List items in root cloud space
./lark doc list

# List items in specific folder
./lark doc list fldbcRho46N6MQ3mJkOAuPabcef
```

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
    },
    {
      "token": "shtcnOko1Ad0HU48HH8KHuabcef",
      "name": "My Sheet",
      "type": "sheet",
      "parent_token": "fldbcRho46N6MQ3mJkOAuPabcef",
      "url": "https://larksuite.com/sheets/shtcnOko1Ad0HU48HH8KHuabcef"
    }
  ],
  "count": 2
}
```

Item types: `doc`, `docx`, `sheet`, `bitable`, `mindnote`, `file`, `folder`, `shortcut`

For shortcuts, a `shortcut_info` field is included with `target_type` and `target_token`.

#### Resolve Wiki Node

```bash
./lark doc wiki <node-token>
```

Resolve a wiki node token to get the underlying document information. The node_token is from the wiki URL. For example:
- URL: `https://xxx.larksuite.com/wiki/X8Tawq431ifOYSklP2tlamKsgNh`
- Node token: `X8Tawq431ifOYSklP2tlamKsgNh`

Output:
```json
{
  "node_token": "X8Tawq431ifOYSklP2tlamKsgNh",
  "obj_token": "DoccnzAaODNqyKC8g9hOWgSpprd",
  "obj_type": "docx",
  "title": "Document Title",
  "space_id": "6946843325487912356",
  "node_type": "origin",
  "has_child": false
}
```

Use the `obj_token` value with `doc get` to retrieve the document content.

#### Get Document as Markdown

```bash
./lark doc get <document-id>
```

The document_id is the token from the document URL. For example:
- URL: `https://xxx.larksuite.com/docx/ABC123xyz`
- Document ID: `ABC123xyz`

Output:
```json
{
  "document_id": "ABC123xyz",
  "title": "My Document",
  "content": "# Heading\n\nDocument content as markdown..."
}
```

#### Get Document Block Structure

```bash
./lark doc blocks <document-id>
```

Returns the full block structure as JSON. Useful for programmatic manipulation.

Output:
```json
{
  "document_id": "ABC123xyz",
  "title": "My Document",
  "block_count": 42,
  "blocks": [...]
}
```

Block types:
- 1: Page (root)
- 2: Text
- 3-11: Headings (H1-H9)
- 12: Bullet list
- 13: Ordered list
- 14: Code block
- 15: Quote
- 17: Todo
- 22: Divider
- 27: Image
- 31: Table

#### Get Document Comments

```bash
./lark doc comments <document-id>
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
      "quote": "The quoted text from the document",
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
- `quote`: the text from the document that was highlighted when commenting (for inline comments)

#### Efficient Extraction with jq and grep

For large documents, use `jq` and `grep` to extract specific information:

```bash
# Get document outline (headings only)
./lark doc get <id> | jq -r '.content' | grep -E '^#{1,6} '

# Get title only
./lark doc get <id> | jq -r '.title'

# Search for keyword with context
./lark doc get <id> | jq -r '.content' | grep -i -C 3 "keyword"

# Quick stats
./lark doc blocks <id> | jq '{title, block_count}'

# Extract all todo items (block type 17)
./lark doc blocks <id> | jq '[.blocks[] | select(.block_type == 17)]'

# Count blocks by type
./lark doc blocks <id> | jq '.blocks | group_by(.block_type) | map({type: .[0].block_type, count: length})'
```

#### Size Comparison

| Format | Typical Size | Use Case |
|--------|--------------|----------|
| Markdown (`doc get`) | ~90 KB | Reading content, summarizing |
| Blocks (`doc blocks`) | ~216 KB | Document manipulation, structure analysis |

Prefer `doc get` for most use cases - it's 2-3x smaller.

### Mail (IMAP)

Email access via IMAP with local caching for fast search.

#### Setup

```bash
# Configure IMAP credentials (interactive)
./lark mail setup
```

This prompts for:
- IMAP host (default: imap.larksuite.com)
- Port (default: 993)
- Username (your Lark email address)
- Password (dedicated password from Lark Mail settings)

To get your IMAP credentials, see: https://www.larksuite.com/hc/en-US/articles/378111206512-log-in-to-lark-mail-through-a-third-party-email-client

1. Open Lark on desktop or web
2. Go to Mail > Settings (gear icon) > Mail settings
3. Select "Third-party email client"
4. Enable IMAP and generate a dedicated password

#### Check Status

```bash
./lark mail status
```

Output:
```json
{
  "configured": true,
  "host": "imap.larksuite.com",
  "port": 993,
  "username": "user@example.com",
  "use_ssl": true,
  "connection": "ok",
  "cache": {
    "last_sync": "2026-01-14T10:30:00+08:00",
    "freshness": "15 minutes ago",
    "uidvalidity": 12345,
    "last_uid": 4521
  }
}
```

#### List Mailboxes

```bash
./lark mail list
```

Output:
```json
{
  "mailboxes": ["INBOX", "Sent", "Drafts", "Trash"],
  "count": 4
}
```

#### Sync Emails

Fetch new emails from the server into the local cache.

```bash
# Sync INBOX (default)
./lark mail sync

# Sync with more parallel connections (faster for large mailboxes)
./lark mail sync --workers 20

# Sync specific mailbox
./lark mail sync --mailbox Sent
```

Flags:
- `--mailbox`, `-m`: Mailbox to sync (default: INBOX)
- `--workers`, `-w`: Number of parallel connections (default: 10)

The sync is resumable - if interrupted, running sync again will only fetch messages not already cached. Progress is displayed during sync.

Output:
```json
{
  "mailbox": "INBOX",
  "new_messages": 5,
  "total_cached": 1523,
  "message": "synced 5 new messages"
}
```

#### Search Emails

Search the local cache (no network calls, very fast).

```bash
# List recent emails
./lark mail search

# Filter by sender
./lark mail search --from alice@example.com

# Filter by subject
./lark mail search --subject "Q4 Report"

# Filter by date range
./lark mail search --since 2026-01-01 --before 2026-01-15

# Combined filters with limit
./lark mail search --from boss@example.com --since 2026-01-01 --limit 10

# Different mailbox
./lark mail search --mailbox Sent --from me@example.com
```

Output:
```json
{
  "mailbox": "INBOX",
  "last_sync": "2026-01-14T10:30:00+08:00",
  "freshness": "15 minutes ago",
  "total_cached": 1523,
  "results": [
    {
      "uid": 4521,
      "message_id": "<abc123@mail.example.com>",
      "date": "2026-01-14T09:15:00+08:00",
      "from_addr": "alice@example.com",
      "from_name": "Alice",
      "subject": "Q4 Report"
    }
  ],
  "count": 1
}
```

**Note:** The `freshness` field indicates how stale the cache is. If data is stale, run `lark mail sync` first.

#### Show Email Content

Fetch and display the full content of an email by UID.

```bash
./lark mail show --uid 4521
```

Output:
```json
{
  "uid": 4521,
  "from": {
    "email": "alice@example.com",
    "name": "Alice"
  },
  "subject": "Q4 Report",
  "date": "2026-01-14T09:15:00+08:00",
  "message_id": "<abc123@mail.example.com>",
  "body": "Hi,\n\nPlease find the Q4 report attached...\n\nBest,\nAlice"
}
```

#### Download Email as .eml

Save an email as a standard .eml file.

```bash
# Download to current directory
./lark mail fetch --uid 4521

# Download to specific directory
./lark mail fetch --uid 4521 --output ./emails/
```

Output:
```json
{
  "success": true,
  "uid": 4521,
  "filename": "2026-01-14 Q4 Report.eml",
  "path": "./emails/2026-01-14 Q4 Report.eml",
  "size": 15234
}
```

### Minutes

Access Lark Minutes meeting recordings.

#### Get Minutes Metadata

```bash
./lark minutes get <minute-token>
```

The minute token is from the Minutes URL:
- URL: `https://bytedance.larksuite.com/minutes/obcnq3b9jl72l83w4f14xxxx`
- Token: `obcnq3b9jl72l83w4f14xxxx`

Output:
```json
{
  "token": "obcnq3b9jl72l83w4f14xxxx",
  "title": "Team Sync Meeting",
  "owner_id": "ou_612b787ccd3259fb3c816b3f678dxxxx",
  "create_time": "2026-01-20T10:30:00+08:00",
  "duration_seconds": 3600,
  "duration_display": "1h 0m 0s",
  "url": "https://bytedance.larksuite.com/minutes/obcnq3b9jl72l83w4f14xxxx"
}
```

#### Export Transcript

```bash
# Plain text transcript
./lark minutes transcript <minute-token>

# SRT format (for subtitles)
./lark minutes transcript <minute-token> --format srt

# Include speaker names
./lark minutes transcript <minute-token> --speaker

# Include timestamps
./lark minutes transcript <minute-token> --timestamp

# Save to file
./lark minutes transcript <minute-token> --output transcript.txt
```

Flags:
- `--format`: Output format - `txt` (default) or `srt`
- `--speaker`: Include speaker names
- `--timestamp`: Include timestamps
- `--output`: Write to file instead of JSON output

Output:
```json
{
  "token": "obcnq3b9jl72l83w4f14xxxx",
  "format": "txt",
  "content": "Full transcript text here..."
}
```

#### Get Media Download URL

```bash
./lark minutes media <minute-token>
```

Returns a temporary download URL valid for 24 hours:
```json
{
  "token": "obcnq3b9jl72l83w4f14xxxx",
  "download_url": "https://internal-api-drive-stream.larksuite.cn/..."
}
```

## Input Formats

### Dates and Times

- ISO 8601: `2026-01-02`, `2026-01-02T09:00:00+08:00`

### Durations

- `30m`, `1h`, `1h30m`, `2h`

### Event IDs

- String from Lark API, e.g., `efa67a98-06a8-4df5-8559-746c8f4477ef_0`

## Error Format

```json
{
  "error": true,
  "code": "EVENT_NOT_FOUND",
  "message": "Event with ID abc123 not found"
}
```

Error codes:
- `AUTH_ERROR`: Authentication failed
- `CONFIG_ERROR`: Configuration issue
- `CALENDAR_ERROR`: Calendar operation failed
- `EVENT_NOT_FOUND`: Event not found
- `API_ERROR`: Lark API error
- `PARSE_ERROR`: Failed to parse input
- `VALIDATION_ERROR`: Invalid input

## Configuration

Config file: `.lark/config.yaml`

```yaml
app_id: "cli_xxxxxxxxxx"
defaults:
  timezone: "Asia/Singapore"
  reminder_minutes: 15
oauth:
  redirect_port: 9999
```

Environment variables:
- `LARK_APP_ID`: Override app_id
- `LARK_APP_SECRET`: App secret (required, never store in file)
