package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/yjwong/lark-cli/internal/api"
	"github.com/yjwong/lark-cli/internal/output"
)

var msgCmd = &cobra.Command{
	Use:   "msg",
	Short: "Message commands",
	Long:  "Retrieve and manage messages in Lark chats",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		validateScopeGroup("messages")
	},
}

// --- msg history ---

var (
	msgHistoryChatID    string
	msgHistoryType      string
	msgHistoryStartTime string
	msgHistoryEndTime   string
	msgHistorySort      string
	msgHistoryLimit     int
)

var msgHistoryCmd = &cobra.Command{
	Use:   "history",
	Short: "Get chat message history",
	Long: `Retrieve message history from a chat or thread.

Requires the bot to be in the group chat. For group chats, the app must have
the "Read all messages in associated group chat" permission scope.

Examples:
  lark msg history --chat-id oc_xxxxx
  lark msg history --chat-id oc_xxxxx --limit 50
  lark msg history --chat-id oc_xxxxx --start 1704067200 --end 1704153600
  lark msg history --chat-id oc_xxxxx --sort desc
  lark msg history --chat-id thread_xxxxx --type thread`,
	Run: func(cmd *cobra.Command, args []string) {
		if msgHistoryChatID == "" {
			output.Fatalf("VALIDATION_ERROR", "chat-id is required")
		}

		client := api.NewClient()

		// Build options
		opts := &api.ListMessagesOptions{}

		if msgHistoryStartTime != "" {
			opts.StartTime = parseTimeArg(msgHistoryStartTime)
		}
		if msgHistoryEndTime != "" {
			opts.EndTime = parseTimeArg(msgHistoryEndTime)
		}
		if msgHistorySort != "" {
			if msgHistorySort == "asc" {
				opts.SortType = "ByCreateTimeAsc"
			} else if msgHistorySort == "desc" {
				opts.SortType = "ByCreateTimeDesc"
			} else {
				output.Fatalf("VALIDATION_ERROR", "sort must be 'asc' or 'desc'")
			}
		}

		// Fetch messages with pagination
		var allMessages []api.Message
		var pageToken string
		hasMore := true
		remaining := msgHistoryLimit

		for hasMore {
			// Calculate page size
			pageSize := 50
			if remaining > 0 && remaining < pageSize {
				pageSize = remaining
			}
			opts.PageSize = pageSize
			opts.PageToken = pageToken

			messages, more, nextToken, err := client.ListMessages(msgHistoryType, msgHistoryChatID, opts)
			if err != nil {
				output.Fatal("API_ERROR", err)
			}

			allMessages = append(allMessages, messages...)
			hasMore = more
			pageToken = nextToken

			// Check limit
			if msgHistoryLimit > 0 {
				remaining = msgHistoryLimit - len(allMessages)
				if remaining <= 0 {
					break
				}
			}
		}

		// Trim to limit if needed
		if msgHistoryLimit > 0 && len(allMessages) > msgHistoryLimit {
			allMessages = allMessages[:msgHistoryLimit]
		}

		// Convert to output format
		outputMessages := make([]api.OutputMessage, len(allMessages))
		for i, m := range allMessages {
			outputMessages[i] = convertMessage(m)
		}

		result := api.OutputMessageList{
			Messages: outputMessages,
			Count:    len(outputMessages),
			ChatID:   msgHistoryChatID,
		}

		output.JSON(result)
	},
}

// parseTimeArg parses a time argument as either Unix timestamp or ISO 8601
func parseTimeArg(s string) string {
	// First try as Unix timestamp
	if _, err := strconv.ParseInt(s, 10, 64); err == nil {
		return s
	}

	// Try parsing as ISO 8601
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		// Try without timezone
		t, err = time.Parse("2006-01-02T15:04:05", s)
		if err != nil {
			// Try date only
			t, err = time.Parse("2006-01-02", s)
			if err != nil {
				output.Fatalf("PARSE_ERROR", "invalid time format: %s (use Unix timestamp or ISO 8601)", s)
			}
		}
	}

	return strconv.FormatInt(t.Unix(), 10)
}

// convertMessage converts an API message to CLI output format
func convertMessage(m api.Message) api.OutputMessage {
	out := api.OutputMessage{
		MessageID:  m.MessageID,
		MsgType:    m.MsgType,
		CreateTime: formatMessageTime(m.CreateTime),
		IsReply:    m.RootID != "" || m.ParentID != "",
		ThreadID:   m.ThreadID,
		Deleted:    m.Deleted,
	}

	if m.Body != nil {
		out.Content = m.Body.Content
	}

	if m.Sender != nil {
		out.Sender = &api.OutputMessageSender{
			ID:   m.Sender.ID,
			Type: m.Sender.SenderType,
		}
	}

	if len(m.Mentions) > 0 {
		out.Mentions = make([]api.OutputMessageMention, len(m.Mentions))
		for i, mention := range m.Mentions {
			out.Mentions[i] = api.OutputMessageMention{
				Key:  mention.Key,
				ID:   mention.ID,
				Name: mention.Name,
			}
		}
	}

	return out
}

// formatMessageTime converts Unix milliseconds to ISO 8601
func formatMessageTime(ms string) string {
	if ms == "" {
		return ""
	}

	msInt, err := strconv.ParseInt(ms, 10, 64)
	if err != nil {
		return ms
	}

	t := time.UnixMilli(msInt)
	return t.Format(time.RFC3339)
}

// --- msg resource ---

var (
	msgResourceMessageID string
	msgResourceFileKey   string
	msgResourceType      string
	msgResourceOutput    string
)

var msgResourceCmd = &cobra.Command{
	Use:   "resource",
	Short: "Download a resource file from a message",
	Long: `Download resource files (images, videos, audios, files) from messages.

The file_key can be found in the message content JSON returned by 'lark msg history'.
For image messages, use --type image. For file, audio, and video messages, use --type file.

Examples:
  lark msg resource --message-id om_xxx --file-key img_v2_xxx --type image --output ./image.png
  lark msg resource --message-id om_xxx --file-key file_v2_xxx --type file --output ./video.mp4`,
	Run: func(cmd *cobra.Command, args []string) {
		if msgResourceMessageID == "" {
			output.Fatalf("VALIDATION_ERROR", "message-id is required")
		}
		if msgResourceFileKey == "" {
			output.Fatalf("VALIDATION_ERROR", "file-key is required")
		}
		if msgResourceType == "" {
			output.Fatalf("VALIDATION_ERROR", "type is required (image or file)")
		}
		if msgResourceType != "image" && msgResourceType != "file" {
			output.Fatalf("VALIDATION_ERROR", "type must be 'image' or 'file'")
		}
		if msgResourceOutput == "" {
			output.Fatalf("VALIDATION_ERROR", "output is required")
		}

		client := api.NewClient()

		// Download the resource
		body, contentType, err := client.GetMessageResource(msgResourceMessageID, msgResourceFileKey, msgResourceType)
		if err != nil {
			output.Fatal("API_ERROR", err)
		}
		defer body.Close()

		// Create output file
		outFile, err := os.Create(msgResourceOutput)
		if err != nil {
			output.Fatalf("FILE_ERROR", "failed to create output file: %v", err)
		}
		defer outFile.Close()

		// Copy data to file
		bytesWritten, err := io.Copy(outFile, body)
		if err != nil {
			output.Fatalf("FILE_ERROR", "failed to write file: %v", err)
		}

		// Output result
		result := map[string]interface{}{
			"success":       true,
			"message_id":    msgResourceMessageID,
			"file_key":      msgResourceFileKey,
			"output_path":   msgResourceOutput,
			"content_type":  contentType,
			"bytes_written": bytesWritten,
		}
		output.JSON(result)
	},
}

// --- msg send ---

var (
	msgSendTo       string
	msgSendToType   string
	msgSendText     string
	msgSendMentions []string
	msgSendLink     string
	msgSendURL      string
)

var msgSendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send a message to a user or chat",
	Long: `Send a message to a user or chat as the bot.

Message types:
- Simple text: Use --text only
- Rich text with mentions/links: Use --text with --mention and/or --link/--url

Examples:
  # Send text to user
  lark msg send --to ou_xxx --text "Hello!"

  # Send to group chat with line breaks
  lark msg send --to oc_xxx --text "Line 1\nLine 2\nLine 3"

  # Mention users (auto-upgrades to post type)
  lark msg send --to oc_xxx --text "Team update" --mention ou_user1 --mention ou_user2

  # With link
  lark msg send --to ou_xxx --text "Check this" --link "Google" --url "https://google.com"

  # Combined
  lark msg send --to oc_xxx --text "Update" --mention ou_xxx --link "Details" --url "https://..."`,
	Run: func(cmd *cobra.Command, args []string) {
		if msgSendTo == "" {
			output.Fatalf("VALIDATION_ERROR", "--to is required")
		}
		if msgSendText == "" {
			output.Fatalf("VALIDATION_ERROR", "--text is required")
		}

		// Auto-detect receive_id_type if not specified
		receiveIDType := msgSendToType
		if receiveIDType == "" {
			receiveIDType = detectIDType(msgSendTo)
		}

		// Build message content
		hasMentions := len(msgSendMentions) > 0
		hasLink := msgSendLink != "" && msgSendURL != ""

		var msgType string
		var content string
		var err error

		if hasMentions || hasLink {
			// Use post (rich text) type
			msgType = "post"
			content, err = buildPostContent(msgSendText, msgSendMentions, msgSendLink, msgSendURL)
		} else {
			// Use simple text type
			msgType = "text"
			content, err = buildTextContent(msgSendText)
		}

		if err != nil {
			output.Fatal("VALIDATION_ERROR", err)
		}

		// Send message
		client := api.NewClient()
		resp, err := client.SendMessage(receiveIDType, msgSendTo, msgType, content)
		if err != nil {
			output.Fatal("API_ERROR", err)
		}

		// Format output
		result := api.OutputSendMessage{
			Success:    true,
			MessageID:  resp.Data.MessageID,
			ChatID:     resp.Data.ChatID,
			CreateTime: formatMessageTime(resp.Data.CreateTime),
		}

		output.JSON(result)
	},
}

// detectIDType infers the receive_id_type from the ID format
func detectIDType(id string) string {
	if strings.HasPrefix(id, "ou_") {
		return "open_id"
	}
	if strings.HasPrefix(id, "oc_") {
		return "chat_id"
	}
	if strings.Contains(id, "@") {
		return "email"
	}
	// Assume user_id if all numeric
	if _, err := strconv.Atoi(id); err == nil {
		return "user_id"
	}
	// Default to open_id
	return "open_id"
}

// unescapeString processes escape sequences like \n, \t, \r, etc.
// Users can send literal backslash-n by escaping as \\n
func unescapeString(s string) string {
	// Wrap string in quotes and use strconv.Unquote to process escape sequences
	// This handles \n -> newline, \t -> tab, \r -> carriage return, \\ -> backslash, etc.
	quoted := `"` + s + `"`
	unquoted, err := strconv.Unquote(quoted)
	if err != nil {
		// If unquoting fails (e.g., malformed escape), return original string
		return s
	}
	return unquoted
}

// buildTextContent creates JSON content for text message type
func buildTextContent(text string) (string, error) {
	// Unescape the text to handle \n, \t, etc.
	unescapedText := unescapeString(text)

	content := map[string]interface{}{
		"text": unescapedText,
	}
	jsonBytes, err := json.Marshal(content)
	if err != nil {
		return "", fmt.Errorf("failed to build text content: %w", err)
	}
	return string(jsonBytes), nil
}

// buildPostContent creates JSON content for post (rich text) message type
func buildPostContent(text string, mentions []string, linkText, linkURL string) (string, error) {
	// Build content array
	var elements []map[string]interface{}

	// Add main text
	if text != "" {
		// Unescape the text to handle \n, \t, etc.
		unescapedText := unescapeString(text)
		elements = append(elements, map[string]interface{}{
			"tag":  "text",
			"text": unescapedText,
		})
	}

	// Add mentions
	for _, userID := range mentions {
		// Add space before mention
		elements = append(elements, map[string]interface{}{
			"tag":  "text",
			"text": " ",
		})
		elements = append(elements, map[string]interface{}{
			"tag":     "at",
			"user_id": userID,
		})
	}

	// Add link
	if linkText != "" && linkURL != "" {
		// Add space before link
		elements = append(elements, map[string]interface{}{
			"tag":  "text",
			"text": " ",
		})
		elements = append(elements, map[string]interface{}{
			"tag":  "a",
			"text": linkText,
			"href": linkURL,
		})
	}

	content := map[string]interface{}{
		"zh_cn": map[string]interface{}{
			"title":   "",
			"content": [][]map[string]interface{}{elements},
		},
		"en_us": map[string]interface{}{
			"title":   "",
			"content": [][]map[string]interface{}{elements},
		},
	}

	jsonBytes, err := json.Marshal(content)
	if err != nil {
		return "", fmt.Errorf("failed to build post content: %w", err)
	}
	return string(jsonBytes), nil
}

func init() {
	// msg history flags
	msgHistoryCmd.Flags().StringVar(&msgHistoryChatID, "chat-id", "", "Chat ID or thread ID (required)")
	msgHistoryCmd.Flags().StringVar(&msgHistoryType, "type", "chat", "Container type: 'chat' or 'thread'")
	msgHistoryCmd.Flags().StringVar(&msgHistoryStartTime, "start", "", "Start time (Unix timestamp or ISO 8601)")
	msgHistoryCmd.Flags().StringVar(&msgHistoryEndTime, "end", "", "End time (Unix timestamp or ISO 8601)")
	msgHistoryCmd.Flags().StringVar(&msgHistorySort, "sort", "", "Sort order: 'asc' or 'desc' (default: asc)")
	msgHistoryCmd.Flags().IntVar(&msgHistoryLimit, "limit", 0, "Maximum number of messages to retrieve (0 = no limit)")

	// msg resource flags
	msgResourceCmd.Flags().StringVar(&msgResourceMessageID, "message-id", "", "Message ID containing the resource (required)")
	msgResourceCmd.Flags().StringVar(&msgResourceFileKey, "file-key", "", "Resource file key from message content (required)")
	msgResourceCmd.Flags().StringVar(&msgResourceType, "type", "", "Resource type: 'image' or 'file' (required)")
	msgResourceCmd.Flags().StringVar(&msgResourceOutput, "output", "", "Output file path (required)")

	// msg send flags
	msgSendCmd.Flags().StringVar(&msgSendTo, "to", "", "Recipient ID (user ID, open_id, email, or chat_id) (required)")
	msgSendCmd.Flags().StringVar(&msgSendToType, "to-type", "", "Recipient ID type: open_id, user_id, email, chat_id (auto-detected if not specified)")
	msgSendCmd.Flags().StringVar(&msgSendText, "text", "", "Message text (required)")
	msgSendCmd.Flags().StringArrayVar(&msgSendMentions, "mention", []string{}, "User ID to mention (can be repeated)")
	msgSendCmd.Flags().StringVar(&msgSendLink, "link", "", "Link text (requires --url)")
	msgSendCmd.Flags().StringVar(&msgSendURL, "url", "", "Link URL (requires --link)")

	// msg recall command
	var msgRecallCmd = &cobra.Command{
		Use:   "recall <message-id>",
		Short: "Recall a message sent by the bot",
		Long: `Recall a message sent by the bot.

The bot can only recall its own messages sent within the configurable time limit (default 24 hours).
For group chats, if the bot is a group owner/admin/creator, it can recall any message within 1 year.

Examples:
  lark msg recall om_dc13264520392913993dd051dba21dcf`,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			client := api.NewClient()
			if err := client.DeleteMessage(args[0]); err != nil {
				output.Fatal("DELETE_FAILED", err)
			}
			output.JSON(map[string]interface{}{
				"success":   true,
				"message":   "Message recalled",
				"messageId": args[0],
			})
		},
	}

	// Register subcommands
	msgCmd.AddCommand(msgHistoryCmd)
	msgCmd.AddCommand(msgResourceCmd)
	msgCmd.AddCommand(msgSendCmd)
	msgCmd.AddCommand(msgRecallCmd)
}
