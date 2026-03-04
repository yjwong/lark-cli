---
name: tasks
description: View Lark Tasks - list your tasks and view task details. Use when user asks about their todo list, tasks, or things to do.
---

# Lark Tasks Skill

View Lark Tasks via the `lark` CLI.

## Running Commands

Ensure `lark` is in your PATH, or use the full path to the binary:

```bash
lark task <command>
```

## Commands Reference

### List Tasks

```bash
lark task list [--limit N] [--completed]
```

Lists tasks assigned to you. By default, only shows incomplete tasks.

Options:
- `--limit`: Maximum number of tasks to retrieve (default: no limit)
- `--completed`: Include completed tasks

Output:
```json
{
  "tasks": [
    {
      "guid": "d300e75f-c56a-4be9-80d6-e47653b3e1a9",
      "summary": "Review Q4 budget proposal",
      "description": "Check the numbers and approve",
      "due_date": "2026-02-05T00:00:00+08:00",
      "is_all_day": true,
      "status": "todo",
      "creator_id": "ou_xxx",
      "creator_name": "John Doe",
      "created_at": "1704067200000"
    }
  ],
  "count": 1,
  "has_more": false
}
```

### Get Task Details

```bash
lark task get <task_guid>
```

Gets details of a specific task by its GUID.

Output:
```json
{
  "guid": "d300e75f-c56a-4be9-80d6-e47653b3e1a9",
  "summary": "Review Q4 budget proposal",
  "description": "Check the numbers and approve by EOD",
  "due_date": "2026-02-05T00:00:00+08:00",
  "is_all_day": true,
  "status": "todo",
  "creator_id": "ou_xxx",
  "creator_name": "John Doe",
  "created_at": "1704067200000"
}
```

## Task Status Values

| Status | Description |
|--------|-------------|
| `todo` | Task is incomplete |
| `completed` | Task has been completed |

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
- `SCOPE_ERROR` - Missing tasks permissions. Run `lark auth login --add --scopes tasks`
- `NOT_FOUND` - Task GUID not found
- `API_ERROR` - Lark API issue

## Required Permissions

This skill requires the `tasks` scope group. If you see a `SCOPE_ERROR`, add the permissions:

```bash
lark auth login --add --scopes tasks
```

To check current permissions:
```bash
lark auth status
```

## Best Practices

### Daily Task Review
```bash
# Show all incomplete tasks
lark task list

# Show including completed
lark task list --completed --limit 20
```

### Integration with Calendar
Tasks with due dates can be cross-referenced with calendar events to help plan the day.

### Working with Task Details
- Use `task get` to see the full description
- Creator information helps identify who assigned the task
- Check `due_date` and `is_all_day` for deadline information
