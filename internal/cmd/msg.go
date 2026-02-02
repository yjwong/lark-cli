package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/yjwong/lark-cli/internal/api"
	"github.com/yjwong/lark-cli/internal/config"
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
	msgSendTo     string
	msgSendToType string
	msgSendText   string
	msgSendImages []string
)

var msgSendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send a message to a user or chat",
	Long: `Send a message to a user or chat as the bot.

Message format:
- Markdown-lite (default): Use --text with **bold**, *italic*, [text](url), and @{ou_xxx} mentions
- Images: Use --image and place {{image}} in --text to position them

Examples:
	# Send text to user
	lark msg send --to ou_xxx --text "Hello!"

	# Send to group chat with line breaks
	lark msg send --to oc_xxx --text "Line 1\nLine 2\nLine 3"

	# Bold and italic
	lark msg send --to oc_xxx --text "**Status:** *Green*"

	# Mention users with @{open_id}
	lark msg send --to oc_xxx --text "Please review @{ou_user1}"

	# With link
	lark msg send --to ou_xxx --text "Check this [Google](https://google.com)"

	# Text + image
	lark msg send --to oc_xxx --text "Intro\n{{image}}\nMore details" --image ./diagram.png

	# Multiple images
	lark msg send --to oc_xxx --text "A\n{{image}}\nB\n{{image}}\nC" --image ./one.png --image ./two.png

	# Image only
	lark msg send --to oc_xxx --image ./screenshot.png`,
	Run: func(cmd *cobra.Command, args []string) {
		if msgSendTo == "" {
			output.Fatalf("VALIDATION_ERROR", "--to is required")
		}
		if msgSendText == "" && len(msgSendImages) == 0 {
			output.Fatalf("VALIDATION_ERROR", "--text or --image is required")
		}

		// Auto-detect receive_id_type if not specified
		receiveIDType := msgSendToType
		if receiveIDType == "" {
			receiveIDType = detectIDType(msgSendTo)
		}

		client := api.NewClient()
		imageKeys := make([]string, 0, len(msgSendImages))
		for _, imagePath := range msgSendImages {
			imageKey, err := client.UploadMessageImage(imagePath)
			if err != nil {
				if errors.Is(err, os.ErrNotExist) {
					output.Fatalf("FILE_ERROR", "image not found: %s", imagePath)
				}
				output.Fatal("API_ERROR", err)
			}
			imageKeys = append(imageKeys, imageKey)
		}

		// Build message content (markdown-lite post)
		msgType := "post"
		content, err := buildMarkdownPostContentWithImages(msgSendText, imageKeys)
		if err != nil {
			output.Fatal("VALIDATION_ERROR", err)
		}

		// Send message
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

// --- msg react ---

var (
	msgReactMessageID        string
	msgReactReactionID       string
	msgReactReactionType     string
	msgReactListMessageID    string
	msgReactListReactionID   string
	msgReactListLimit        int
	msgReactRemoveMessageID  string
	msgReactRemoveReactionID string
)

var msgReactCmd = &cobra.Command{
	Use:   "react",
	Short: "Add a reaction to a message",
	Long: `Add a reaction to a message as the bot.

Examples:
  lark msg react --message-id om_xxx --reaction smile
  lark msg react --message-id om_xxx --reaction "+1" --type emoji`,
	Run: func(cmd *cobra.Command, args []string) {
		if msgReactMessageID == "" {
			output.Fatalf("VALIDATION_ERROR", "message-id is required")
		}
		if msgReactReactionID == "" {
			output.Fatalf("VALIDATION_ERROR", "reaction is required")
		}
		if msgReactReactionType != "" && msgReactReactionType != "emoji" {
			output.Fatalf("VALIDATION_ERROR", "type must be 'emoji'")
		}

		client := api.NewClient()
		emojiType := strings.ToUpper(msgReactReactionID)
		reaction, err := client.AddMessageReaction(msgReactMessageID, emojiType)
		if err != nil {
			output.Fatal("API_ERROR", err)
		}

		result := api.OutputMessageReaction{
			Success:      true,
			MessageID:    msgReactMessageID,
			ReactionType: "emoji",
			ReactionID:   emojiType,
			EmojiType:    emojiType,
		}
		if reaction != nil {
			if reaction.ReactionID != "" {
				result.ReactionID = reaction.ReactionID
			}
			if reaction.ReactionType != nil && reaction.ReactionType.EmojiType != "" {
				result.EmojiType = reaction.ReactionType.EmojiType
			}
		}

		output.JSON(result)
	},
}

// --- msg react list ---

var msgReactListCmd = &cobra.Command{
	Use:   "list",
	Short: "List reactions for a message",
	Long: `List reactions for a message.

Examples:
  lark msg react list --message-id om_xxx
  lark msg react list --message-id om_xxx --reaction SMILE
  lark msg react list --message-id om_xxx --limit 50`,
	Run: func(cmd *cobra.Command, args []string) {
		if msgReactListMessageID == "" {
			output.Fatalf("VALIDATION_ERROR", "message-id is required")
		}

		client := api.NewClient()
		opts := &api.ListMessageReactionsOptions{}
		if msgReactListReactionID != "" {
			opts.ReactionType = strings.ToUpper(msgReactListReactionID)
		}

		var allReactions []api.MessageReaction
		var pageToken string
		hasMore := true
		remaining := msgReactListLimit

		for hasMore {
			pageSize := 20
			if remaining > 0 && remaining < pageSize {
				pageSize = remaining
			}
			opts.PageSize = pageSize
			opts.PageToken = pageToken

			reactions, more, nextToken, err := client.ListMessageReactions(msgReactListMessageID, opts)
			if err != nil {
				output.Fatal("API_ERROR", err)
			}

			allReactions = append(allReactions, reactions...)
			hasMore = more
			pageToken = nextToken

			if msgReactListLimit > 0 {
				remaining = msgReactListLimit - len(allReactions)
				if remaining <= 0 {
					break
				}
			}
		}

		if msgReactListLimit > 0 && len(allReactions) > msgReactListLimit {
			allReactions = allReactions[:msgReactListLimit]
		}

		outputReactions := make([]api.OutputMessageReactionItem, len(allReactions))
		for i, r := range allReactions {
			outputReactions[i] = convertMessageReaction(r)
		}

		result := api.OutputMessageReactionList{
			MessageID: msgReactListMessageID,
			Reactions: outputReactions,
			Count:     len(outputReactions),
		}

		output.JSON(result)
	},
}

// --- msg react remove ---

var msgReactRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a reaction from a message",
	Long: `Remove a reaction from a message.

Examples:
  lark msg react remove --message-id om_xxx --reaction-id ZCaCIjUBVVWSrm5L-3ZTwxxxx`,
	Run: func(cmd *cobra.Command, args []string) {
		if msgReactRemoveMessageID == "" {
			output.Fatalf("VALIDATION_ERROR", "message-id is required")
		}
		if msgReactRemoveReactionID == "" {
			output.Fatalf("VALIDATION_ERROR", "reaction-id is required")
		}

		client := api.NewClient()
		reaction, err := client.DeleteMessageReaction(msgReactRemoveMessageID, msgReactRemoveReactionID)
		if err != nil {
			output.Fatal("API_ERROR", err)
		}

		result := api.OutputMessageReaction{
			Success:      true,
			MessageID:    msgReactRemoveMessageID,
			ReactionType: "emoji",
			ReactionID:   msgReactRemoveReactionID,
			EmojiType:    "",
		}
		if reaction != nil {
			if reaction.ReactionType != nil && reaction.ReactionType.EmojiType != "" {
				result.EmojiType = reaction.ReactionType.EmojiType
			}
			if reaction.ReactionID != "" {
				result.ReactionID = reaction.ReactionID
			}
		}

		output.JSON(result)
	},
}

// --- msg react emojis ---

var msgReactEmojisCmd = &cobra.Command{
	Use:   "emojis",
	Short: "Show emoji catalog reference",
	Long: `Show the Lark emoji catalog reference for reaction emoji types.

Examples:
  lark msg react emojis`,
	Run: func(cmd *cobra.Command, args []string) {
		emojis := []string{
			"OK",
			"THUMBSUP",
			"THANKS",
			"MUSCLE",
			"FINGERHEART",
			"APPLAUSE",
			"FISTBUMP",
			"JIAYI",
			"DONE",
			"SMILE",
			"BLUSH",
			"LAUGH",
			"SMIRK",
			"LOL",
			"FACEPALM",
			"LOVE",
			"WINK",
			"PROUD",
			"WITTY",
			"SMART",
			"SCOWL",
			"THINKING",
			"SOB",
			"CRY",
			"ERROR",
			"NOSEPICK",
			"HAUGHTY",
			"SLAP",
			"SPITBLOOD",
			"TOASTED",
			"GLANCE",
			"DULL",
			"INNOCENTSMILE",
			"JOYFUL",
			"WOW",
			"TRICK",
			"YEAH",
			"ENOUGH",
			"TEARS",
			"EMBARRASSED",
			"KISS",
			"SMOOCH",
			"DROOL",
			"OBSESSED",
			"MONEY",
			"TEASE",
			"SHOWOFF",
			"COMFORT",
			"CLAP",
			"PRAISE",
			"STRIVE",
			"XBLUSH",
			"SILENT",
			"WAVE",
			"WHAT",
			"FROWN",
			"SHY",
			"DIZZY",
			"LOOKDOWN",
			"CHUCKLE",
			"WAIL",
			"CRAZY",
			"WHIMPER",
			"HUG",
			"BLUBBER",
			"WRONGED",
			"HUSKY",
			"SHHH",
			"SMUG",
			"ANGRY",
			"HAMMER",
			"SHOCKED",
			"TERROR",
			"PETRIFIED",
			"SKULL",
			"SWEAT",
			"SPEECHLESS",
			"SLEEP",
			"DROWSY",
			"YAWN",
			"SICK",
			"PUKE",
			"BETRAYED",
			"HEADSET",
			"EatingFood",
			"MeMeMe",
			"Sigh",
			"Typing",
			"Lemon",
			"Get",
			"LGTM",
			"OnIt",
			"OneSecond",
			"VRHeadset",
			"YouAreTheBest",
			"SALUTE",
			"SHAKE",
			"HIGHFIVE",
			"UPPERLEFT",
			"ThumbsDown",
			"SLIGHT",
			"TONGUE",
			"EYESCLOSED",
			"RoarForYou",
			"CALF",
			"BEAR",
			"BULL",
			"RAINBOWPUKE",
			"ROSE",
			"HEART",
			"PARTY",
			"LIPS",
			"BEER",
			"CAKE",
			"GIFT",
			"CUCUMBER",
			"Drumstick",
			"Pepper",
			"CANDIEDHAWS",
			"BubbleTea",
			"Coffee",
			"Yes",
			"No",
			"OKR",
			"CheckMark",
			"CrossMark",
			"MinusOne",
			"Hundred",
			"AWESOMEN",
			"Pin",
			"Alarm",
			"Loudspeaker",
			"Trophy",
			"Fire",
			"BOMB",
			"Music",
			"XmasTree",
			"Snowman",
			"XmasHat",
			"FIREWORKS",
			"2022",
			"REDPACKET",
			"FORTUNE",
			"LUCK",
			"FIRECRACKER",
			"StickyRiceBalls",
			"HEARTBROKEN",
			"POOP",
			"StatusFlashOfInspiration",
			"18X",
			"CLEAVER",
			"Soccer",
			"Basketball",
			"GeneralDoNotDisturb",
			"Status_PrivateMessage",
			"GeneralInMeetingBusy",
			"StatusReading",
			"StatusInFlight",
			"GeneralBusinessTrip",
			"GeneralWorkFromHome",
			"StatusEnjoyLife",
			"GeneralTravellingCar",
			"StatusBus",
			"GeneralSun",
			"GeneralMoonRest",
		}
		customEmojis := config.GetCustomEmojis()
		for emojiID := range customEmojis {
			emojis = append(emojis, emojiID)
		}
		output.JSON(map[string]interface{}{
			"source":        "im-v1/message-reaction/emojis-introduce",
			"url":           "https://open.larksuite.com/document/server-docs/im-v1/message-reaction/emojis-introduce",
			"examples":      []string{"SMILE", "LAUGH", "THUMBSUP", "CLAP", "OK", "HEART"},
			"count":         len(emojis),
			"emojis":        emojis,
			"custom_emojis": customEmojis,
		})
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

type postElement struct {
	tag    string
	text   string
	href   string
	userID string
	bold   bool
	italic bool
}

const imagePlaceholder = "{{image}}"

// buildMarkdownPostContent creates JSON content for markdown-lite post messages.
func buildMarkdownPostContent(text string) (string, error) {
	return buildMarkdownPostContentWithImages(text, nil)
}

// buildMarkdownPostContentWithImages creates JSON content with image placeholders.
func buildMarkdownPostContentWithImages(text string, imageKeys []string) (string, error) {
	unescapedText := unescapeString(text)
	var lines []string
	if unescapedText != "" {
		lines = strings.Split(unescapedText, "\n")
	}
	contentLines := make([][]map[string]interface{}, 0, len(lines)+len(imageKeys))
	usedImages := 0

	for _, line := range lines {
		if strings.Contains(line, imagePlaceholder) {
			segments := strings.Split(line, imagePlaceholder)
			for index, segment := range segments {
				if segment != "" {
					elements := parseMarkdownLine(segment)
					if len(elements) == 0 {
						elements = append(elements, postElement{tag: "text", text: ""})
					}
					contentLines = append(contentLines, buildPostElements(elements))
				}
				if index < len(segments)-1 {
					if usedImages >= len(imageKeys) {
						return "", fmt.Errorf("not enough images for %s placeholders", imagePlaceholder)
					}
					contentLines = append(contentLines, []map[string]interface{}{
						{
							"tag":       "img",
							"image_key": imageKeys[usedImages],
						},
					})
					usedImages++
				}
			}
			continue
		}

		elements := parseMarkdownLine(line)
		if len(elements) == 0 {
			elements = append(elements, postElement{tag: "text", text: ""})
		}
		contentLines = append(contentLines, buildPostElements(elements))
	}

	for usedImages < len(imageKeys) {
		contentLines = append(contentLines, []map[string]interface{}{
			{
				"tag":       "img",
				"image_key": imageKeys[usedImages],
			},
		})
		usedImages++
	}

	if len(contentLines) == 0 {
		return "", fmt.Errorf("message content cannot be empty")
	}

	content := map[string]interface{}{
		"zh_cn": map[string]interface{}{
			"title":   "",
			"content": contentLines,
		},
		"en_us": map[string]interface{}{
			"title":   "",
			"content": contentLines,
		},
	}

	jsonBytes, err := json.Marshal(content)
	if err != nil {
		return "", fmt.Errorf("failed to build post content: %w", err)
	}
	return string(jsonBytes), nil
}

func convertMessageReaction(r api.MessageReaction) api.OutputMessageReactionItem {
	item := api.OutputMessageReactionItem{
		ReactionID: r.ReactionID,
	}
	if r.ReactionType != nil {
		item.EmojiType = r.ReactionType.EmojiType
	}
	if r.Operator != nil {
		item.OperatorID = r.Operator.OperatorID
		item.OperatorType = r.Operator.OperatorType
	}
	if r.ActionTime != "" {
		item.ActionTime = formatMessageTime(r.ActionTime)
	}
	return item
}

func parseMarkdownLine(line string) []postElement {
	var elements []postElement
	var buffer strings.Builder

	flushText := func(text string) {
		if text == "" {
			return
		}
		elements = append(elements, postElement{tag: "text", text: text})
	}

	i := 0
	for i < len(line) {
		if strings.HasPrefix(line[i:], "@{") {
			closing := strings.Index(line[i+2:], "}")
			if closing >= 0 {
				flushText(buffer.String())
				buffer.Reset()
				userID := line[i+2 : i+2+closing]
				elements = append(elements, postElement{tag: "at", userID: userID})
				i += closing + 3
				continue
			}
		}

		if line[i] == '[' {
			closeBracket := strings.Index(line[i+1:], "]")
			if closeBracket >= 0 {
				closeBracket += i + 1
				if closeBracket+1 < len(line) && line[closeBracket+1] == '(' {
					closeParen := strings.Index(line[closeBracket+2:], ")")
					if closeParen >= 0 {
						closeParen += closeBracket + 2
						linkText := line[i+1 : closeBracket]
						linkURL := line[closeBracket+2 : closeParen]
						flushText(buffer.String())
						buffer.Reset()
						elements = append(elements, postElement{tag: "a", text: linkText, href: linkURL})
						i = closeParen + 1
						continue
					}
				}
			}
		}

		if strings.HasPrefix(line[i:], "**") {
			closing := strings.Index(line[i+2:], "**")
			if closing >= 0 {
				flushText(buffer.String())
				buffer.Reset()
				boldText := line[i+2 : i+2+closing]
				if boldText != "" {
					elements = append(elements, postElement{tag: "text", text: boldText, bold: true})
				}
				i += closing + 4
				continue
			}
		}

		if line[i] == '*' {
			closing := strings.Index(line[i+1:], "*")
			if closing >= 0 {
				flushText(buffer.String())
				buffer.Reset()
				italicText := line[i+1 : i+1+closing]
				if italicText != "" {
					elements = append(elements, postElement{tag: "text", text: italicText, italic: true})
				}
				i += closing + 2
				continue
			}
		}

		buffer.WriteByte(line[i])
		i++
	}

	flushText(buffer.String())
	return elements
}

func buildPostElements(elements []postElement) []map[string]interface{} {
	content := make([]map[string]interface{}, 0, len(elements))
	for _, elem := range elements {
		switch elem.tag {
		case "text":
			entry := map[string]interface{}{
				"tag":  "text",
				"text": elem.text,
			}
			style := buildPostTextStyle(elem.bold, elem.italic)
			if len(style) > 0 {
				entry["style"] = style
			}
			content = append(content, entry)
		case "a":
			entry := map[string]interface{}{
				"tag":  "a",
				"text": elem.text,
				"href": elem.href,
			}
			style := buildPostTextStyle(elem.bold, elem.italic)
			if len(style) > 0 {
				entry["style"] = style
			}
			content = append(content, entry)
		case "at":
			content = append(content, map[string]interface{}{
				"tag":     "at",
				"user_id": elem.userID,
			})
		}
	}
	return content
}

func buildPostTextStyle(isBold, isItalic bool) []string {
	var style []string
	if isBold {
		style = append(style, "bold")
	}
	if isItalic {
		style = append(style, "italic")
	}
	return style
}

// --- msg recall ---

var msgRecallCmd = &cobra.Command{
	Use:   "recall <message-id>",
	Short: "Recall a message",
	Long: `Recall a previously sent message.

Messages can be recalled within 24 hours of sending.
Group owners/admins can recall member messages within 1 year.

Examples:
  lark msg recall om_dc13264520392913993dd051dba21dcf`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		messageID := args[0]
		client := api.NewClient()

		if err := client.RecallMessage(messageID); err != nil {
			output.Fatal("API_ERROR", err)
		}

		output.JSON(map[string]interface{}{
			"success":    true,
			"message_id": messageID,
		})
	},
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
	msgSendCmd.Flags().StringVar(&msgSendText, "text", "", "Message text (markdown-lite). Use {{image}} to place images")
	msgSendCmd.Flags().StringSliceVar(&msgSendImages, "image", nil, "Image file path (repeatable)")

	// msg react flags
	msgReactCmd.Flags().StringVar(&msgReactMessageID, "message-id", "", "Message ID to react to (required)")
	msgReactCmd.Flags().StringVar(&msgReactReactionID, "reaction", "", "Reaction ID or emoji name (required)")
	msgReactCmd.Flags().StringVar(&msgReactReactionType, "type", "emoji", "Reaction type (default: emoji)")

	// msg react list flags
	msgReactListCmd.Flags().StringVar(&msgReactListMessageID, "message-id", "", "Message ID to list reactions for (required)")
	msgReactListCmd.Flags().StringVar(&msgReactListReactionID, "reaction", "", "Emoji type to filter (optional)")
	msgReactListCmd.Flags().IntVar(&msgReactListLimit, "limit", 0, "Maximum number of reactions to retrieve (0 = no limit)")

	// msg react remove flags
	msgReactRemoveCmd.Flags().StringVar(&msgReactRemoveMessageID, "message-id", "", "Message ID to remove reaction from (required)")
	msgReactRemoveCmd.Flags().StringVar(&msgReactRemoveReactionID, "reaction-id", "", "Reaction ID to remove (required)")

	// Register subcommands
	msgCmd.AddCommand(msgHistoryCmd)
	msgCmd.AddCommand(msgResourceCmd)
	msgCmd.AddCommand(msgSendCmd)
	msgCmd.AddCommand(msgReactCmd)
	msgCmd.AddCommand(msgRecallCmd)

	msgReactCmd.AddCommand(msgReactListCmd)
	msgReactCmd.AddCommand(msgReactRemoveCmd)
	msgReactCmd.AddCommand(msgReactEmojisCmd)
}
