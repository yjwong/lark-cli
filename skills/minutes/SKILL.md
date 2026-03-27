---
name: minutes
description: Access Lark Minutes recordings - get metadata, export transcripts, download audio/video. Use when user asks about meeting recordings, transcripts, or minutes.
---

# Minutes Skill

Access Lark Minutes recordings via the `lark` CLI.

## Running Commands

Ensure `lark` is in your PATH, or use the full path to the binary. Set the config directory if not using the default:

```bash
lark minutes <command>
# Or with explicit config:
LARK_CONFIG_DIR=/path/to/.lark lark minutes <command>
```

## Getting the Minute Token

The minute token is a 24-character identifier found in the Minutes URL:
- URL format: `https://bytedance.larksuite.com/minutes/obcnq3b9jl72l83w4f14xxxx`
- Token: `obcnq3b9jl72l83w4f14xxxx` (the last part of the URL)

## Commands Reference

### Check Authentication
```bash
lark auth status
```

### Get Minutes Metadata
```bash
lark minutes get <minute_token>
```

Returns:
```json
{
  "token": "obcnq3b9jl72l83w4f14xxxx",
  "title": "Team Sync Meeting",
  "owner_id": "ou_612b787ccd3259fb3c816b3f678dxxxx",
  "create_time": "2024-01-20T10:30:00+08:00",
  "duration_seconds": 3600,
  "duration_display": "1h 0m 0s",
  "url": "https://bytedance.larksuite.com/minutes/obcnq3b9jl72l83w4f14xxxx"
}
```

### Export Transcript
```bash
# Plain text transcript
lark minutes transcript <minute_token>

# SRT format (for subtitles)
lark minutes transcript <minute_token> --format srt

# Include speaker names
lark minutes transcript <minute_token> --speaker

# Include timestamps
lark minutes transcript <minute_token> --timestamp

# Full transcript with all details
lark minutes transcript <minute_token> --format srt --speaker --timestamp

# Save to file
lark minutes transcript <minute_token> --output transcript.txt
lark minutes transcript <minute_token> --format srt --output transcript.srt
```

Flags:
- `--format txt|srt` - Output format (default: txt)
- `--speaker` - Include speaker names in transcript
- `--timestamp` - Include timestamps
- `--output <file>` - Write to file instead of JSON output

Returns (when not using --output):
```json
{
  "token": "obcnq3b9jl72l83w4f14xxxx",
  "format": "txt",
  "content": "Full transcript text here..."
}
```

Returns (when using --output):
```json
{
  "token": "obcnq3b9jl72l83w4f14xxxx",
  "format": "txt",
  "file": "transcript.txt"
}
```

### Get Media Download URL
```bash
lark minutes media <minute_token>
```

Returns a temporary download URL valid for 24 hours:
```json
{
  "token": "obcnq3b9jl72l83w4f14xxxx",
  "download_url": "https://internal-api-drive-stream.larksuite.cn/space/api/box/stream/download/authcode/?code=xxx"
}
```

## Output Format

All commands output JSON for programmatic processing.

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
- `SCOPE_ERROR` - Missing minutes permissions. Run `lark auth login --scopes minutes`
- `NOT_FOUND` - Invalid minute token or minutes not found
- `API_ERROR` - Lark API issue (check if minute is ready/transcribed)

## Required Permissions

This skill requires the `minutes` scope group. If you see a `SCOPE_ERROR`, the user needs to add minutes permissions:

```bash
lark auth login --scopes minutes
```

To check current permissions:
```bash
lark auth status
```

## Common Use Cases

### Review Meeting Recording
1. Get metadata to see title and duration
2. Export transcript for reading/searching
3. Get media URL if need to watch the video

### Extract Meeting Notes
1. Export transcript with speaker names: `--speaker`
2. Process the transcript to extract action items, decisions, etc.

### Archive Meeting Content
1. Export transcript to file: `--output meeting-notes.txt`
2. Get media URL and download video for archival

## Notes

- Minutes must be fully transcribed before the API can access them
- Media download URLs expire after 24 hours
- Rate limit: 5 requests per second per endpoint
- Requires appropriate Minutes scopes in the Lark app configuration
