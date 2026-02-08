## PR Review: feat: add Lark Tasks support (#12)

Overall this is a well-structured addition that follows many of the existing codebase patterns. The types, API client, scope registration, and command structure all align with how other features (calendar, messages, etc.) are implemented. A few issues to address:

### Issues

#### 1. `formatTaskTimestamp` does not apply user timezone (Medium)

**File:** `internal/cmd/task.go`, `formatTaskTimestamp` function

Every other command that formats timestamps loads the user's configured timezone via `config.GetTimezone()` and applies it with `time.LoadLocation()`. The task timestamp formatter converts Unix milliseconds to RFC3339 but uses the system default timezone rather than the user's configured one. This will produce timestamps in the wrong timezone for users who have configured a timezone in their Lark config.

Compare with `internal/api/events.go:317` or `internal/cmd/list.go:52-53` which all do:
```go
tz := config.GetTimezone()
loc, _ := time.LoadLocation(tz)
// ...
t := time.Unix(ts, 0).In(loc)
```

Suggestion: Load the timezone and apply `.In(loc)` before formatting, consistent with the rest of the codebase.

#### 2. `formatTaskTimestamp` — manual digit parsing instead of `strconv.ParseInt` (Low)

**File:** `internal/cmd/task.go`, lines ~170-178

The function manually iterates over characters to parse an integer:
```go
for i := 0; i < len(ts); i++ {
    if ts[i] >= '0' && ts[i] <= '9' {
        msec = msec*10 + int64(ts[i]-'0')
    } else {
        return ts
    }
}
```

The rest of the codebase uses `strconv.ParseInt` for this (see `internal/cmd/msg.go:199`). This manual parsing is harder to read and doesn't handle edge cases like integer overflow. Suggest replacing with:
```go
msec, err := strconv.ParseInt(ts, 10, 64)
if err != nil {
    return ts
}
```

#### 3. `created_at` and `completed_at` are not formatted as RFC3339 in output (Low)

**File:** `internal/cmd/task.go`, `taskToOutput` function

`CompletedAt` and `CreatedAt` are passed through as raw strings from the API response. The SKILL.md example output shows `"created_at": "1704067200000"` — raw Unix milliseconds in the JSON output. Other commands format timestamps for the user (e.g., calendar events output RFC3339). For consistency, these should also be run through `formatTaskTimestamp()`:

```go
CompletedAt: formatTaskTimestamp(t.CompletedAt),
CreatedAt:   formatTaskTimestamp(t.CreatedAt),
```

#### 4. `taskStatusToString` maps "done" → "completed" — potential confusion (Nit)

**File:** `internal/cmd/task.go`, `taskStatusToString`

The API returns `"done"` but the CLI outputs `"completed"`. This means consumers that receive a task from `task get` cannot use the status value to filter via `--completed` flag or match against the API. The SKILL.md documents only `todo` and `completed` as status values, but someone debugging against the Lark API might be confused. Consider either:
- Keeping the original API values (`"done"`) and documenting them, or
- Clearly noting this mapping in the SKILL.md

#### 5. `HasMore` in `OutputTaskList` may be inaccurate when `--limit` is used (Low)

**File:** `internal/cmd/task.go`, list command

When `--limit` is used and the limit is reached, the loop breaks early. The `hasMore` variable at that point reflects the API's pagination state, not whether there are more tasks beyond the user's limit. This could be misleading — it might say `has_more: false` even though the user only requested a subset. Consider always setting `HasMore: true` when the limit caused the early break, or documenting the behavior.

### Minor Notes

- The error message format `"API error %d: %s"` in `tasks.go` is consistent with the newer files (messages.go, contacts.go, sheets.go) but differs from the older pattern `"API error (code %d): %s"` used in calendar.go/events.go. Not a blocker — the codebase itself is inconsistent here.
- The `url.PathEscape(taskGUID)` in `GetTask` is good defensive practice.
- The SKILL.md is well-structured and follows a good format for Claude Code integration.
- Scope registration and `AllGroupNames()` update are correct and complete.

### Summary

The PR is in good shape architecturally. The main item to address is the timezone handling (#1), which is a functional correctness issue. The timestamp formatting issues (#2, #3) are quality-of-life improvements for consistency. The rest are minor.
