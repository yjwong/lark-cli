package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yjwong/lark-cli/internal/api"
	"github.com/yjwong/lark-cli/internal/output"
)

// maxPaginationPages is a safety limit to prevent infinite pagination loops.
const maxPaginationPages = 200

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Chat/group commands",
	Long:  "Search and manage Lark chats and groups",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		validateScopeGroup("messages")
	},
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

		for page := 0; hasMore; page++ {
			if page >= maxPaginationPages {
				output.Fatal("PAGINATION_ERROR", fmt.Errorf("exceeded maximum page count (%d)", maxPaginationPages))
			}

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

			if more && nextToken == pageToken {
				output.Fatal("PAGINATION_ERROR", fmt.Errorf("API returned duplicate page token"))
			}
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

// --- chat list ---

var (
	chatListLimit int
)

var chatListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all chats the bot has joined",
	Long: `List all group chats that the bot is a member of.

Unlike 'chat search', this uses the /im/v1/chats API which returns
all chats the bot has joined, regardless of search visibility settings.

Examples:
  lark chat list
  lark chat list --limit 20`,
	Run: func(cmd *cobra.Command, args []string) {
		client := api.NewClient()

		opts := &api.ListChatsOptions{}

		var allChats []api.Chat
		var pageToken string
		hasMore := true
		remaining := chatListLimit

		for page := 0; hasMore; page++ {
			if page >= maxPaginationPages {
				output.Fatal("PAGINATION_ERROR", fmt.Errorf("exceeded maximum page count (%d)", maxPaginationPages))
			}

			pageSize := 50
			if remaining > 0 && remaining < pageSize {
				pageSize = remaining
			}
			opts.PageSize = pageSize
			opts.PageToken = pageToken

			chats, more, nextToken, err := client.ListChats(opts)
			if err != nil {
				output.Fatal("API_ERROR", err)
			}

			allChats = append(allChats, chats...)

			if more && nextToken == pageToken {
				output.Fatal("PAGINATION_ERROR", fmt.Errorf("API returned duplicate page token"))
			}
			hasMore = more
			pageToken = nextToken

			if chatListLimit > 0 {
				remaining = chatListLimit - len(allChats)
				if remaining <= 0 {
					break
				}
			}
		}

		if chatListLimit > 0 && len(allChats) > chatListLimit {
			allChats = allChats[:chatListLimit]
		}

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

		output.JSON(api.OutputChatList{
			Chats: outputChats,
			Count: len(outputChats),
		})
	},
}

// --- chat members ---

var (
	chatMembersLimit int
)

var chatMembersCmd = &cobra.Command{
	Use:   "members <chat_id>",
	Short: "List members of a chat",
	Long: `List all members of a group chat.

Examples:
  lark chat members oc_xxxxx
  lark chat members oc_xxxxx --limit 20`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client := api.NewClient()

		opts := &api.ListChatMembersOptions{
			ChatID:       args[0],
			MemberIDType: "open_id",
		}

		var allMembers []api.ChatMember
		var pageToken string
		hasMore := true
		remaining := chatMembersLimit

		for page := 0; hasMore; page++ {
			if page >= maxPaginationPages {
				output.Fatal("PAGINATION_ERROR", fmt.Errorf("exceeded maximum page count (%d)", maxPaginationPages))
			}

			pageSize := 50
			if remaining > 0 && remaining < pageSize {
				pageSize = remaining
			}
			opts.PageSize = pageSize
			opts.PageToken = pageToken

			members, more, nextToken, err := client.ListChatMembers(opts)
			if err != nil {
				output.Fatal("API_ERROR", err)
			}

			allMembers = append(allMembers, members...)

			if more && nextToken == pageToken {
				output.Fatal("PAGINATION_ERROR", fmt.Errorf("API returned duplicate page token"))
			}
			hasMore = more
			pageToken = nextToken

			if chatMembersLimit > 0 {
				remaining = chatMembersLimit - len(allMembers)
				if remaining <= 0 {
					break
				}
			}
		}

		if chatMembersLimit > 0 && len(allMembers) > chatMembersLimit {
			allMembers = allMembers[:chatMembersLimit]
		}

		outputMembers := make([]api.OutputChatMember, len(allMembers))
		for i, m := range allMembers {
			outputMembers[i] = api.OutputChatMember{
				OpenID: m.MemberID,
				Name:   m.Name,
			}
		}

		output.JSON(api.OutputChatMemberList{
			ChatID:  args[0],
			Members: outputMembers,
			Count:   len(outputMembers),
		})
	},
}

func init() {
	chatSearchCmd.Flags().IntVar(&chatSearchLimit, "limit", 0,
		"Maximum number of chats to retrieve (0 = no limit)")
	chatListCmd.Flags().IntVar(&chatListLimit, "limit", 0,
		"Maximum number of chats to retrieve (0 = no limit)")
	chatMembersCmd.Flags().IntVar(&chatMembersLimit, "limit", 0,
		"Maximum number of members to retrieve (0 = no limit)")

	chatCmd.AddCommand(chatSearchCmd)
	chatCmd.AddCommand(chatListCmd)
	chatCmd.AddCommand(chatMembersCmd)
}
