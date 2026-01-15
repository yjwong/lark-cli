# Claude Code Instructions for lark

A Go CLI tool for interacting with Lark APIs, designed for use by Claude Code.

## Project Structure

```
cmd/lark/         # Main entry point
internal/
  api/            # Lark API client
  auth/           # OAuth authentication
  cmd/            # Cobra command implementations
  config/         # Configuration handling
  conflicts/      # Conflict detection
  mail/           # IMAP mail client with caching
  output/         # JSON/human-readable output formatting
  time/           # Date/time parsing
```

## Build & Test

```bash
make build        # Build binary to ./lark
make test         # Run tests
make deps         # Tidy and download dependencies
make install      # Install to $GOPATH/bin
```

## Code Conventions

- JSON output by default
- Error responses use structured format: `{"error": true, "code": "...", "message": "..."}`
- Date parsing supports ISO 8601 formats
- Config in `.lark/config.yaml`, secrets via `LARK_APP_SECRET` env var
- Use Cobra for commands, Viper for config

## Commands

See `USAGE.md` for full CLI documentation. Main commands:

### Auth
- `auth login|status|logout` - Authentication

### Calendar (`cal`)
- `cal list` - List events (supports `--week`, `--from`, `--to`)
- `cal show <id>` - Show event details
- `cal create` - Create event
- `cal update <id>` - Update event
- `cal delete <id>` - Delete event
- `cal search <query>` - Search events
- `cal freebusy` - Query availability
- `cal lookup-user` - Get user ID from email
- `cal common-freetime` - Find mutual availability
- `cal rsvp <id>` - Accept/decline invitation

### Contacts (`contact`)
- `contact get <user_id>` - Get user info by ID
- `contact list-dept [dept_id]` - List users in department
- `contact search-dept <query>` - Search departments by name

### Messages (`msg`)
- `msg history` - Get chat message history
- `msg resource` - Download message attachments
- `msg send` - Send messages to users or chats
- `msg recall <message-id>` - Recall a message sent by the bot

**Message Sending Features:**
- Send text messages with line breaks (`\n` creates newlines)
- Mention users with `--mention` flag
- Include links with `--link` and `--url` flags
- Send to emails, user IDs, or chat IDs
- Auto-detect recipient types
- Support escape sequences (\n, \t, \", \\)

### Documents (`doc`)
- `doc list [folder_token]` - List items in a Drive folder
- `doc get <document_id>` - Get document content as markdown
- `doc blocks <document_id>` - Get document block structure (JSON)
- `doc wiki <node_token>` - Resolve wiki node to document token

### Mail (`mail`)
- `mail setup` - Configure IMAP credentials (interactive)
- `mail status` - Show connection and cache status
- `mail list` - List mailboxes/folders
- `mail sync` - Fetch new emails into local cache (supports `--workers` for parallel sync)
- `mail search` - Search cached emails (fast, local)
- `mail show --uid <uid>` - Show email content
- `mail fetch --uid <uid>` - Download as .eml file

**Mail Architecture:**
- Uses IMAP (not REST API) for better search capability
- Local SQLite cache for O(1) search performance
- Credentials stored in `$LARK_CONFIG_DIR/mail.json`
- Cache stored in `$LARK_CONFIG_DIR/mail_cache.db`
- Search results include `freshness` field for cache staleness
- Sync supports parallel connections (`--workers`) for faster initial sync
- Sync is resumable - only fetches messages not already cached
