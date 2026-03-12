---
name: contacts
description: Look up employee information via Lark - find colleagues by ID, list department members, search users by name, search departments. Use when user asks about a person, colleague, job title, department, or org structure.
---

# Contacts Lookup Skill

Look up employee information via the `lark` CLI.

## Running Commands

Ensure `lark` is in your PATH, or use the full path to the binary. Set the config directory if not using the default:

```bash
lark contact <command>
# Or with explicit config:
LARK_CONFIG_DIR=/path/to/.lark lark contact <command>
```

## Commands Reference

### Get User by ID
```bash
# Look up by open_id (default)
lark contact get ou_xxxx

# Look up by user_id
lark contact get 12345 --id-type user_id
```

Output:
```json
{
  "user_id": "ou_xxx",
  "open_id": "ou_xxx",
  "name": "Jane Doe",
  "en_name": "Jane Doe",
  "email": "jane@example.com",
  "job_title": "Data Analyst",
  "department": "Business Intelligence"
}
```

### List Users in Department
```bash
# List users in root department
lark contact list-dept

# List users in specific department
lark contact list-dept od_xxxx

# Specify page size (max 50, default 50)
lark contact list-dept od_xxxx --page-size 25
```

Available flags: `--page-size` (number of results per page, max 50)

Output:
```json
{
  "contacts": [
    {
      "user_id": "ou_xxx",
      "name": "Alice",
      "job_title": "Engineer",
      "department": "Engineering"
    }
  ],
  "count": 1
}
```

### Search Users by Name
```bash
lark contact search "Jane"
lark contact search "John Smith"
```

Output:
```json
{
  "contacts": [
    {
      "user_id": "ou_xxx",
      "open_id": "ou_xxx",
      "name": "Jane Doe",
      "department": "Engineering"
    }
  ],
  "count": 1
}
```

### Search Departments
```bash
lark contact search-dept "Engineering"
```

Output:
```json
{
  "departments": [
    {
      "department_id": "od_xxx",
      "name": "Engineering",
      "member_count": 42
    }
  ],
  "count": 1
}
```

## Integration with Calendar

When showing calendar events with attendees, you can enrich attendee info:

1. Get attendee `open_id` from calendar event
2. Use `contact get <open_id>` to fetch job title and department
3. Present enriched attendee info to user

Example workflow:
```bash
# Get event with attendees
lark cal show <event_id>
# Returns attendees with open_id

# Look up each attendee
lark contact get ou_attendee_id
# Returns name, job_title, department
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
- `SCOPE_ERROR` - Missing contacts permissions. Run `lark auth login --add --scopes contacts`
- `NOT_FOUND` - User or department not found
- `API_ERROR` - Lark API issue

## Required Permissions

This skill requires the `contacts` scope group. If you see a `SCOPE_ERROR`, the user needs to add contacts permissions:

```bash
lark auth login --add --scopes contacts
```

To check current permissions:
```bash
lark auth status
```

## Notes

- The `search` and `search-dept` commands require user authentication (OAuth via `lark auth login`)
- The `get` and `list-dept` commands use tenant token (no login required)
- Department IDs typically start with `od_`
- User open_ids typically start with `ou_`
