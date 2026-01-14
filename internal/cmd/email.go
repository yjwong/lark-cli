package cmd

import (
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"github.com/yjwong/lark-cli/internal/api"
	"github.com/yjwong/lark-cli/internal/output"
)

var emailCmd = &cobra.Command{
	Use:   "email",
	Short: "Email commands",
	Long:  "Read and manage emails in Lark Mail",
}

// --- email list ---

var (
	emailListMailbox  string
	emailListFolder   string
	emailListUnread   bool
	emailListPageSize int
	emailListAll      bool
)

var emailListCmd = &cobra.Command{
	Use:   "list",
	Short: "List email message IDs",
	Long: `List email message IDs from a mailbox folder.

By default, lists emails from the INBOX folder of the current user's mailbox.
Returns message IDs that can be used with 'email show' to get full details.

Examples:
  lark email list
  lark email list --unread
  lark email list --folder INBOX
  lark email list --mailbox me --page-size 10
  lark email list --all`,
	Run: func(cmd *cobra.Command, args []string) {
		client := api.NewClient()

		opts := &api.ListEmailsOptions{
			FolderID:   emailListFolder,
			OnlyUnread: emailListUnread,
			PageSize:   emailListPageSize,
		}

		if emailListAll {
			// Fetch all pages
			var allIDs []string
			var pageToken string
			hasMore := true

			for hasMore {
				opts.PageToken = pageToken
				ids, more, nextToken, err := client.ListEmails(emailListMailbox, opts)
				if err != nil {
					output.Fatal("API_ERROR", err)
				}
				allIDs = append(allIDs, ids...)
				hasMore = more
				pageToken = nextToken
			}

			result := api.OutputEmailIDList{
				MessageIDs: allIDs,
				Count:      len(allIDs),
				HasMore:    false,
				MailboxID:  emailListMailbox,
			}
			output.JSON(result)
		} else {
			// Single page
			ids, hasMore, _, err := client.ListEmails(emailListMailbox, opts)
			if err != nil {
				output.Fatal("API_ERROR", err)
			}

			result := api.OutputEmailIDList{
				MessageIDs: ids,
				Count:      len(ids),
				HasMore:    hasMore,
				MailboxID:  emailListMailbox,
			}
			output.JSON(result)
		}
	},
}

// --- email show ---

var (
	emailShowMailbox   string
	emailShowMessageID string
)

var emailShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Get email details",
	Long: `Retrieve the full details of an email message.

Returns email metadata (subject, from, to, cc), body content, and attachment info.
The body is decoded from base64 and returned as plain text.

Examples:
  lark email show --id ZWEyNGRmY2QtOTVlNy00...
  lark email show --mailbox me --id ZWEyNGRmY2QtOTVlNy00...`,
	Run: func(cmd *cobra.Command, args []string) {
		if emailShowMessageID == "" {
			output.Fatalf("VALIDATION_ERROR", "--id is required")
		}

		client := api.NewClient()

		email, err := client.GetEmail(emailShowMailbox, emailShowMessageID)
		if err != nil {
			output.Fatal("API_ERROR", err)
		}

		if email == nil {
			output.Fatalf("NOT_FOUND", "email not found")
		}

		result := convertEmailToOutput(email)
		output.JSON(result)
	},
}

// --- email attachments ---

var (
	emailAttachmentsMailbox   string
	emailAttachmentsMessageID string
)

var emailAttachmentsCmd = &cobra.Command{
	Use:   "attachments",
	Short: "Get attachment download URLs",
	Long: `Get download URLs for all attachments in an email.

Returns temporary download URLs for each attachment.
Note: URLs expire after 2 hours and are limited to 2 uses each.

Examples:
  lark email attachments --id ZWEyNGRmY2QtOTVlNy00...
  lark email attachments --mailbox me --id ZWEyNGRmY2QtOTVlNy00...`,
	Run: func(cmd *cobra.Command, args []string) {
		if emailAttachmentsMessageID == "" {
			output.Fatalf("VALIDATION_ERROR", "--id is required")
		}

		client := api.NewClient()

		downloadURLs, failedIDs, attachments, err := client.GetAllAttachmentDownloadURLs(
			emailAttachmentsMailbox, emailAttachmentsMessageID)
		if err != nil {
			output.Fatal("API_ERROR", err)
		}

		if len(attachments) == 0 {
			result := api.OutputAttachmentDownload{
				MessageID: emailAttachmentsMessageID,
				MailboxID: emailAttachmentsMailbox,
				Downloads: []api.OutputAttachmentDownloadURL{},
			}
			output.JSON(result)
			return
		}

		// Build a map of attachment ID to filename
		filenameMap := make(map[string]string)
		for _, att := range attachments {
			filenameMap[att.ID] = att.Filename
		}

		// Convert to output format
		downloads := make([]api.OutputAttachmentDownloadURL, 0, len(downloadURLs))
		for _, dl := range downloadURLs {
			downloads = append(downloads, api.OutputAttachmentDownloadURL{
				AttachmentID: dl.AttachmentID,
				Filename:     filenameMap[dl.AttachmentID],
				DownloadURL:  dl.DownloadURL,
			})
		}

		result := api.OutputAttachmentDownload{
			MessageID: emailAttachmentsMessageID,
			MailboxID: emailAttachmentsMailbox,
			Downloads: downloads,
			FailedIDs: failedIDs,
		}
		output.JSON(result)
	},
}

// convertEmailToOutput converts an EmailMessage to OutputEmail format
func convertEmailToOutput(email *api.EmailMessage) api.OutputEmail {
	result := api.OutputEmail{
		MessageID: email.MessageID,
		Subject:   email.Subject,
		ThreadID:  email.ThreadID,
	}

	// Convert From address
	if email.From != nil {
		result.From = &api.OutputEmailAddress{
			Email: email.From.MailAddress,
			Name:  email.From.Name,
		}
	}

	// Convert To addresses
	if len(email.To) > 0 {
		result.To = make([]api.OutputEmailAddress, len(email.To))
		for i, addr := range email.To {
			result.To[i] = api.OutputEmailAddress{
				Email: addr.MailAddress,
				Name:  addr.Name,
			}
		}
	}

	// Convert CC addresses
	if len(email.CC) > 0 {
		result.CC = make([]api.OutputEmailAddress, len(email.CC))
		for i, addr := range email.CC {
			result.CC[i] = api.OutputEmailAddress{
				Email: addr.MailAddress,
				Name:  addr.Name,
			}
		}
	}

	// Convert date from Unix ms timestamp to ISO 8601
	if email.InternalDate != "" {
		if ms, err := strconv.ParseInt(email.InternalDate, 10, 64); err == nil {
			t := time.Unix(0, ms*int64(time.Millisecond))
			result.Date = t.Format(time.RFC3339)
		}
	}

	// Decode body - prefer plain text
	if email.BodyPlainText != "" {
		decoded, err := api.DecodeEmailBody(email.BodyPlainText)
		if err == nil {
			result.Body = decoded
		}
	}

	// Convert attachments
	if len(email.Attachments) > 0 {
		result.Attachments = make([]api.OutputEmailAttachment, len(email.Attachments))
		for i, att := range email.Attachments {
			attType := "standard"
			if att.AttachmentType == 2 {
				attType = "oversized"
			}
			result.Attachments[i] = api.OutputEmailAttachment{
				ID:       att.ID,
				Filename: att.Filename,
				Type:     attType,
				IsInline: att.IsInline,
			}
		}
	}

	return result
}

func init() {
	// email list flags
	emailListCmd.Flags().StringVarP(&emailListMailbox, "mailbox", "m", "me", "Mailbox ID (email address or 'me')")
	emailListCmd.Flags().StringVarP(&emailListFolder, "folder", "f", "INBOX", "Folder ID (default: INBOX)")
	emailListCmd.Flags().BoolVar(&emailListUnread, "unread", false, "Only list unread emails")
	emailListCmd.Flags().IntVar(&emailListPageSize, "page-size", 20, "Number of results per page (1-20)")
	emailListCmd.Flags().BoolVar(&emailListAll, "all", false, "Fetch all pages")

	// email show flags
	emailShowCmd.Flags().StringVarP(&emailShowMailbox, "mailbox", "m", "me", "Mailbox ID (email address or 'me')")
	emailShowCmd.Flags().StringVar(&emailShowMessageID, "id", "", "Message ID (required)")

	// email attachments flags
	emailAttachmentsCmd.Flags().StringVarP(&emailAttachmentsMailbox, "mailbox", "m", "me", "Mailbox ID (email address or 'me')")
	emailAttachmentsCmd.Flags().StringVar(&emailAttachmentsMessageID, "id", "", "Message ID (required)")

	// Register subcommands
	emailCmd.AddCommand(emailListCmd)
	emailCmd.AddCommand(emailShowCmd)
	emailCmd.AddCommand(emailAttachmentsCmd)
}
