package cmd

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/yjwong/lark-cli/internal/api"
	"github.com/yjwong/lark-cli/internal/output"
)

var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Task commands",
	Long:  "View and manage Lark Tasks",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		validateScopeGroup("tasks")
	},
}

// --- task list ---

var (
	taskListLimit     int
	taskListCompleted bool
)

var taskListCmd = &cobra.Command{
	Use:   "list",
	Short: "List tasks",
	Long: `List tasks assigned to you.

By default, only shows incomplete tasks. Use --completed to include completed tasks.

Examples:
  lark task list
  lark task list --limit 50
  lark task list --completed`,
	Run: func(cmd *cobra.Command, args []string) {
		client := api.NewClient()

		opts := &api.TaskListOptions{
			PageSize:  50,
			Completed: taskListCompleted,
		}

		var allTasks []api.Task
		var pageToken string
		hasMore := true
		remaining := taskListLimit

		for hasMore {
			if remaining > 0 && remaining < opts.PageSize {
				opts.PageSize = remaining
			}
			opts.PageToken = pageToken

			tasks, more, nextToken, err := client.ListTasks(opts)
			if err != nil {
				output.Fatal("API_ERROR", err)
			}

			allTasks = append(allTasks, tasks...)
			hasMore = more
			pageToken = nextToken

			if taskListLimit > 0 {
				remaining = taskListLimit - len(allTasks)
				if remaining <= 0 {
					break
				}
			}
		}

		// Trim to limit if needed
		if taskListLimit > 0 && len(allTasks) > taskListLimit {
			allTasks = allTasks[:taskListLimit]
		}

		outputTasks := make([]api.OutputTask, len(allTasks))
		for i, t := range allTasks {
			outputTasks[i] = taskToOutput(t)
		}

		result := api.OutputTaskList{
			Tasks:   outputTasks,
			Count:   len(outputTasks),
			HasMore: hasMore,
		}

		output.JSON(result)
	},
}

// --- task get ---

var taskGetCmd = &cobra.Command{
	Use:   "get <task_guid>",
	Short: "Get task details",
	Long: `Get details of a specific task by its GUID.

The task GUID is returned in the list command output.

Examples:
  lark task get d300e75f-c56a-4be9-80d6-e47653b3e1a9`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskGUID := args[0]

		client := api.NewClient()

		task, err := client.GetTask(taskGUID)
		if err != nil {
			output.Fatal("API_ERROR", err)
		}

		if task == nil {
			output.Fatalf("NOT_FOUND", "Task not found: %s", taskGUID)
		}

		result := taskToOutput(*task)
		output.JSON(result)
	},
}

// taskToOutput converts a Task to OutputTask
func taskToOutput(t api.Task) api.OutputTask {
	out := api.OutputTask{
		GUID:        t.GUID,
		Summary:     t.Summary,
		Description: t.Description,
		Status:      taskStatusToString(t.Status),
		CompletedAt: t.CompletedAt,
		CreatedAt:   t.CreatedAt,
	}

	if t.Due != nil {
		out.DueDate = formatTaskTimestamp(t.Due.Timestamp)
		out.IsAllDay = t.Due.IsAllDay
	}

	if t.Creator != nil {
		out.CreatorID = t.Creator.ID
		out.CreatorName = t.Creator.Name
	}

	return out
}

// taskStatusToString converts task status to human-readable string
func taskStatusToString(status string) string {
	switch status {
	case "todo":
		return "todo"
	case "done":
		return "completed"
	default:
		return status
	}
}

// formatTaskTimestamp formats a task timestamp (Unix ms) to RFC3339
func formatTaskTimestamp(ts string) string {
	if ts == "" {
		return ""
	}
	// Task API returns timestamps as strings (Unix ms)
	var msec int64
	if _, err := time.Parse(time.RFC3339, ts); err == nil {
		// Already in RFC3339 format
		return ts
	}
	// Try parsing as Unix milliseconds
	if n, err := time.Parse("2006-01-02", ts); err == nil {
		return n.Format("2006-01-02")
	}
	// Try parsing as integer milliseconds
	for i := 0; i < len(ts); i++ {
		if ts[i] >= '0' && ts[i] <= '9' {
			msec = msec*10 + int64(ts[i]-'0')
		} else {
			return ts // Return as-is if can't parse
		}
	}
	if msec > 0 {
		return time.UnixMilli(msec).Format(time.RFC3339)
	}
	return ts
}

func init() {
	// task list flags
	taskListCmd.Flags().IntVar(&taskListLimit, "limit", 0,
		"Maximum number of tasks to retrieve (0 = no limit)")
	taskListCmd.Flags().BoolVar(&taskListCompleted, "completed", false,
		"Include completed tasks")

	// Register subcommands
	taskCmd.AddCommand(taskListCmd)
	taskCmd.AddCommand(taskGetCmd)
}
