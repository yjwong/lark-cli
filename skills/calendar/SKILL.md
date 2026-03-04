---
name: calendar
description: Manage Lark calendar - view schedule, create/update/delete events, check availability, find meeting slots, RSVP to invitations. Use when user asks about meetings, schedule, availability, or calendar.
---

# Calendar Management Skill

Manage Lark calendar via the `lark` CLI.

## Running Commands

Ensure `lark` is in your PATH, or use the full path to the binary. Set the config directory if not using the default:

```bash
lark cal <command>
# Or with explicit config:
LARK_CONFIG_DIR=/path/to/.lark lark cal <command>
```

## Before Running Commands

**Always check the current time first** before running calendar commands that involve dates or times:
```bash
date '+%Y-%m-%d %H:%M:%S %Z'
```

This ensures you use the correct current date when constructing queries like `--from`, `--to`, `--start`, etc.

## Commands Reference

### Check Authentication
```bash
lark auth status
```

### List Events
```bash
# Today's events with your RSVP status (recommended default)
lark cal list --rsvp

# This week with RSVP status
lark cal list --week --rsvp

# Date range (ISO 8601)
lark cal list --from 2024-01-15 --to 2024-01-20 --rsvp

# With conflict detection
lark cal list --week --detect-conflicts --buffer-minutes 15 --rsvp

# Events awaiting RSVP
lark cal list --pending --from 2024-01-15 --to 2024-01-31

# Include all attendee info (slower)
lark cal list --week --attendees
```

**Important**: Always include `--rsvp` when listing events to show the user's RSVP status. Highlight events where `rsvp_status` is `needs_action`.

### Show Event Details
```bash
lark cal show <event-id>
```

### Create Event
```bash
lark cal create \
  --summary "Meeting title" \
  --start 2024-01-20T14:00:00+08:00 \
  --duration 1h \
  --location "Conference Room A" \
  --description "Discussion topics" \
  --attendee user1@example.com \
  --attendee user2@example.com \
  --reminder 15 \
  --color "#9CA2A9" \
  --visibility private
```

Time format: ISO 8601 (e.g., `2024-01-20T14:00:00+08:00` or `2024-01-20`)
Duration formats: `30m`, `1h`, `1h30m`
Color format: Hex (e.g., `#9CA2A9`)
Visibility: `default`, `public`, or `private`
Attendee ability (guest permissions): `none`, `can_see_others`, `can_invite_others`, or `can_modify_event`

Additional flags:
- `--exclude-self` - Don't add yourself as an attendee
- `--attendee-ability` - Set guest permissions (see values above)

### Update Event
```bash
lark cal update <event-id> \
  --summary "New title" \
  --start 2024-01-20T15:00:00+08:00 \
  --end 2024-01-20T16:00:00+08:00 \
  --color "#34C724"

# Change visibility
lark cal update <event-id> --visibility public
```

Available flags: `--summary`, `--description`, `--start`, `--end`, `--location`, `--color`, `--visibility`, `--attendee-ability`, `--no-notify`

- `--no-notify` - Don't send notifications to attendees about the update

Visibility options: `default`, `public`, or `private`

Note: Use `--start` and `--end` together to reschedule. There is no `--duration` flag.

### Delete Event
```bash
lark cal delete <event-id>
```

### Search Events
```bash
lark cal search "keyword" --from 2024-01-08 --to 2024-01-22
```

### Look Up User
```bash
# Get user's open_id from email (needed for freebusy/common-freetime)
lark cal lookup-user --email user@example.com
```

### Check Availability
```bash
# Own availability
lark cal freebusy --from 2024-01-20T09:00:00+08:00 --to 2024-01-20T18:00:00+08:00

# Another user's availability
lark cal freebusy --from 2024-01-20T09:00:00+08:00 --to 2024-01-20T18:00:00+08:00 --user <open_id>

# Meeting room availability
lark cal freebusy --from 2024-01-20T09:00:00+08:00 --to 2024-01-20T18:00:00+08:00 --room <room_id>
```

### Find Common Free Time
```bash
# Find mutual availability with one or more users
lark cal common-freetime \
  --users ou_abc123 \
  --from 2024-01-20T09:00:00+08:00 \
  --to 2024-01-20T18:00:00+08:00 \
  --min-length 30

# Multiple users (comma-separated, max 10)
lark cal common-freetime \
  --users "ou_abc123,ou_def456" \
  --from 2024-01-20 \
  --to 2024-01-21 \
  --work-hours
```

Available flags: `--users` (required), `--from`, `--to`, `--min-length`, `--work-hours`, `--include-external`, `--limit`

### RSVP to Event
```bash
# Accept an invitation
lark cal rsvp <event-id> --accept

# Decline an invitation
lark cal rsvp <event-id> --decline

# Mark as tentative
lark cal rsvp <event-id> --tentative
```

### Manage Attendees
```bash
# Add yourself to an event
lark cal attendee add <event-id> --self

# Add someone by email (adds as third-party/external)
lark cal attendee add <event-id> --email user@example.com

# Add someone by Lark user ID (preferred - shows as proper Lark user)
lark cal attendee add <event-id> --user ou_xxxxxxxx

# Add as optional attendee
lark cal attendee add <event-id> --user ou_xxxxxxxx --optional

# Add without sending notification
lark cal attendee add <event-id> --self --no-notify

# List attendees (shows attendee IDs needed for removal)
lark cal attendee list <event-id>

# Remove yourself from an event
lark cal attendee remove <event-id> --self

# Remove by attendee ID (get ID from attendee list)
lark cal attendee remove <event-id> --id user_xxxxx
```

**Tip**: Just use email addresses with `--attendee` (create) or `--email` (attendee add). The CLI automatically resolves internal Lark users via the contacts API, falling back to third-party for external contacts.

## Output Formats

- **Default**: JSON (for programmatic processing)
- Note: All commands output JSON. Format the output appropriately when presenting to user.

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
- `SCOPE_ERROR` - Missing calendar permissions. Run `lark auth login --add --scopes calendar`
- `EVENT_NOT_FOUND` - Invalid event ID
- `VALIDATION_ERROR` - Missing required fields
- `API_ERROR` - Lark API issue

## Required Permissions

This skill requires the `calendar` scope group. If you see a `SCOPE_ERROR`, the user needs to add calendar permissions:

```bash
lark auth login --add --scopes calendar
```

To check current permissions:
```bash
lark auth status
```

## EA Best Practices

When managing the user's calendar, follow these guidelines:

### Before Scheduling
- **Always check for conflicts** using `--detect-conflicts` before creating events
- **Check attendee availability** with `freebusy` before inviting others
- For in-person meetings, consider travel/commute time

### Buffer Time
- Maintain **15 minutes buffer** between meetings by default
- Flag back-to-back meetings as a concern when showing schedule
- Suggest rescheduling if no buffer exists

### Schedule Health
- **Flag overloaded days**: More than 6 hours of meetings in a day is a concern
- Proactively suggest declining or rescheduling when schedule is packed
- Protect focus time - prefer scheduling meetings in afternoons when possible

### Finding Meeting Slots
- When asked to "find time", use `freebusy` to check availability
- Prefer morning slots for focus work, afternoon for meetings (unless user has different preference in profile)
- Account for timezone differences when scheduling with external attendees

### Rescheduling Meetings
- **Always check all attendees first** - Use `cal show <event-id>` to get the full attendee list before checking availability. Don't assume you know who's invited.
- **Watch for bogus busy blocks** - Freebusy data can contain erroneous all-day or multi-day blocks that aren't real OOO. If `common-freetime` returns no slots unexpectedly, check individual freebusy data to spot anomalies.

### Communication
- When showing the day's schedule, summarize key meetings and highlight any concerns
- After creating/updating events, confirm the details back to the user
- If authentication fails, remind user to run `lark auth login`
