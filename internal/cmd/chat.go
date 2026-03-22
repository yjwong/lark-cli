package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yjwong/lark-cli/internal/api"
	"github.com/yjwong/lark-cli/internal/output"
)

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Chat/group commands",
	Long:  "Search and manage Lark chats and groups",
	// No PersistentPreRun: chat commands use tenant token (can't pre-validate user scopes).
}

// --- chat search ---

var (
	chatSearchLimit int
)

var chatSearchCmd = &cobra.Command{
	Use:   "search [query]",
	Short: "Search for chats/groups",
	Long: `Search for chats and groups visible to the user or bot.

The search supports:
- Group name matching (including internationalized names)
- Group member name matching
- Multiple languages
- Fuzzy search (pinyin, prefix, etc.)

If no query is provided, returns all visible chats.

Examples:
  lark chat search "project"
  lark chat search "团队"
  lark chat search --limit 50`,
	Run: func(cmd *cobra.Command, args []string) {
		client := api.NewClient()

		opts := &api.SearchChatsOptions{}
		if len(args) > 0 {
			opts.Query = args[0]
		}

		// Fetch chats with pagination
		var allChats []api.Chat
		var pageToken string
		hasMore := true
		remaining := chatSearchLimit

		for hasMore {
			pageSize := 50
			if remaining > 0 && remaining < pageSize {
				pageSize = remaining
			}
			opts.PageSize = pageSize
			opts.PageToken = pageToken

			chats, more, nextToken, err := client.SearchChats(opts)
			if err != nil {
				output.Fatal("API_ERROR", err)
			}

			allChats = append(allChats, chats...)
			hasMore = more
			pageToken = nextToken

			if chatSearchLimit > 0 {
				remaining = chatSearchLimit - len(allChats)
				if remaining <= 0 {
					break
				}
			}
		}

		// Trim to limit if needed
		if chatSearchLimit > 0 && len(allChats) > chatSearchLimit {
			allChats = allChats[:chatSearchLimit]
		}

		// Convert to output format
		outputChats := make([]api.OutputChat, len(allChats))
		for i, c := range allChats {
			outputChats[i] = api.OutputChat{
				ChatID:      c.ChatID,
				Name:        c.Name,
				Description: c.Description,
				OwnerID:     c.OwnerID,
				External:    c.External,
				ChatStatus:  c.ChatStatus,
			}
		}

		result := api.OutputChatList{
			Chats: outputChats,
			Count: len(outputChats),
		}
		if len(args) > 0 {
			result.Query = args[0]
		}

		output.JSON(result)
	},
}

func init() {
	chatSearchCmd.Flags().IntVar(&chatSearchLimit, "limit", 0,
		"Maximum number of chats to retrieve (0 = no limit)")

	chatCmd.AddCommand(chatSearchCmd)
}
