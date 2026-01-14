package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/yjwong/lark-cli/internal/mail"
	"github.com/yjwong/lark-cli/internal/output"
)

var mailCmd = &cobra.Command{
	Use:   "mail",
	Short: "Email commands (IMAP)",
	Long:  "Read and search emails via IMAP with local caching",
}

// --- mail setup ---

var mailSetupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Configure IMAP credentials",
	Long: `Configure IMAP credentials for accessing Lark Mail.

See: https://www.larksuite.com/hc/en-US/articles/378111206512-log-in-to-lark-mail-through-a-third-party-email-client

To get your IMAP credentials:
1. Open Lark on desktop or web
2. Go to Mail > Settings (gear icon) > Mail settings
3. Select "Third-party email client"
4. Enable IMAP and generate a dedicated password

This command will prompt for:
- IMAP server host (e.g., imap.larksuite.com)
- Port (usually 993 for SSL)
- Username (your Lark email address)
- Password (dedicated password from step 4)`,
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)

		fmt.Println("Lark Mail IMAP Setup")
		fmt.Println("====================")
		fmt.Println()
		fmt.Println("See: https://www.larksuite.com/hc/en-US/articles/378111206512")
		fmt.Println()
		fmt.Println("To get your IMAP credentials:")
		fmt.Println("1. Open Lark on desktop or web")
		fmt.Println("2. Go to Mail > Settings (gear icon) > Mail settings")
		fmt.Println("3. Select \"Third-party email client\"")
		fmt.Println("4. Enable IMAP and generate a dedicated password")
		fmt.Println()

		creds := &mail.Credentials{
			UseSSL: true,
		}

		// Host
		fmt.Print("IMAP Host [imap.larksuite.com]: ")
		host, _ := reader.ReadString('\n')
		host = strings.TrimSpace(host)
		if host == "" {
			host = "imap.larksuite.com"
		}
		creds.Host = host

		// Port
		fmt.Print("IMAP Port [993]: ")
		portStr, _ := reader.ReadString('\n')
		portStr = strings.TrimSpace(portStr)
		if portStr == "" {
			creds.Port = 993
		} else {
			port, err := strconv.Atoi(portStr)
			if err != nil {
				output.Fatalf("VALIDATION_ERROR", "invalid port: %s", portStr)
			}
			creds.Port = port
		}

		// SSL
		fmt.Print("Use SSL? [Y/n]: ")
		sslStr, _ := reader.ReadString('\n')
		sslStr = strings.TrimSpace(strings.ToLower(sslStr))
		creds.UseSSL = sslStr != "n" && sslStr != "no"

		// Username
		fmt.Print("Username (email address): ")
		username, _ := reader.ReadString('\n')
		username = strings.TrimSpace(username)
		if username == "" {
			output.Fatalf("VALIDATION_ERROR", "username is required")
		}
		creds.Username = username

		// Password
		fmt.Print("Password (app-specific password): ")
		password, _ := reader.ReadString('\n')
		password = strings.TrimSpace(password)
		if password == "" {
			output.Fatalf("VALIDATION_ERROR", "password is required")
		}
		creds.Password = password

		// Test connection
		fmt.Println()
		fmt.Print("Testing connection... ")

		if err := mail.TestConnection(creds); err != nil {
			fmt.Println("FAILED")
			output.Fatal("CONNECTION_ERROR", err)
		}
		fmt.Println("OK")

		// Save credentials
		if err := mail.SaveCredentials(creds); err != nil {
			output.Fatal("SAVE_ERROR", err)
		}

		fmt.Println()
		fmt.Println("Credentials saved successfully!")
		fmt.Println("Run 'lark mail sync' to fetch your emails.")

		output.JSON(map[string]interface{}{
			"success": true,
			"message": "IMAP credentials configured successfully",
		})
	},
}

// --- mail status ---

var mailStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show mail configuration and cache status",
	Run: func(cmd *cobra.Command, args []string) {
		result := map[string]interface{}{
			"configured": mail.HasCredentials(),
		}

		if mail.HasCredentials() {
			creds, err := mail.LoadCredentials()
			if err == nil {
				result["host"] = creds.Host
				result["port"] = creds.Port
				result["username"] = creds.Username
				result["use_ssl"] = creds.UseSSL
			}

			// Test connection
			if err := mail.TestConnection(creds); err != nil {
				result["connection"] = "failed"
				result["connection_error"] = err.Error()
			} else {
				result["connection"] = "ok"
			}
		}

		// Cache status
		cache, err := mail.OpenCache()
		if err == nil {
			defer cache.Close()

			// Get INBOX state as default
			state, _ := cache.GetMailboxState("INBOX")
			if state != nil {
				result["cache"] = map[string]interface{}{
					"last_sync":   state.LastSync.Format(time.RFC3339),
					"freshness":   formatFreshness(state.LastSync),
					"uidvalidity": state.UIDValidity,
					"last_uid":    state.LastUID,
				}
			} else {
				result["cache"] = map[string]interface{}{
					"last_sync": nil,
					"freshness": "never synced",
				}
			}
		}

		output.JSON(result)
	},
}

// --- mail list ---

var mailListCmd = &cobra.Command{
	Use:   "list",
	Short: "List mailboxes/folders",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := mail.Connect()
		if err != nil {
			output.Fatal("CONNECTION_ERROR", err)
		}
		defer client.Close()

		mailboxes, err := client.ListMailboxes()
		if err != nil {
			output.Fatal("IMAP_ERROR", err)
		}

		output.JSON(map[string]interface{}{
			"mailboxes": mailboxes,
			"count":     len(mailboxes),
		})
	},
}

// --- mail sync ---

var (
	mailSyncMailbox string
)

var mailSyncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync emails from server to local cache",
	Long: `Fetch new emails from the IMAP server and store metadata in the local cache.

On first sync, fetches all email headers. On subsequent syncs, only fetches new messages.
The cache is used for fast local searching with 'lark mail search'.`,
	Run: func(cmd *cobra.Command, args []string) {
		result, err := mail.Sync(mailSyncMailbox)
		if err != nil {
			output.Fatal("SYNC_ERROR", err)
		}

		output.JSON(result)
	},
}

// --- mail search ---

var (
	mailSearchMailbox string
	mailSearchFrom    string
	mailSearchSubject string
	mailSearchSince   string
	mailSearchBefore  string
	mailSearchLimit   int
)

var mailSearchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search emails in local cache",
	Long: `Search cached email metadata locally (no network calls).

The search uses the local cache which is updated by 'lark mail sync'.
Results include cache freshness information so you know if data is stale.

Examples:
  lark mail search
  lark mail search --from alice@example.com
  lark mail search --subject "Q4 Report" --since 2025-01-01
  lark mail search --mailbox INBOX --limit 20`,
	Run: func(cmd *cobra.Command, args []string) {
		opts, err := mail.ParseSearchOptions(
			mailSearchFrom, mailSearchSubject,
			mailSearchSince, mailSearchBefore,
			mailSearchLimit,
		)
		if err != nil {
			output.Fatal("VALIDATION_ERROR", err)
		}

		result, err := mail.Search(mailSearchMailbox, opts)
		if err != nil {
			output.Fatal("SEARCH_ERROR", err)
		}

		output.JSON(result)
	},
}

// --- mail show ---

var (
	mailShowMailbox string
	mailShowUID     uint32
)

var mailShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show email content",
	Long: `Fetch and display the full content of an email by UID.

The UID can be obtained from search results.

Examples:
  lark mail show --uid 12345
  lark mail show --mailbox INBOX --uid 12345`,
	Run: func(cmd *cobra.Command, args []string) {
		if mailShowUID == 0 {
			output.Fatalf("VALIDATION_ERROR", "--uid is required")
		}

		client, err := mail.Connect()
		if err != nil {
			output.Fatal("CONNECTION_ERROR", err)
		}
		defer client.Close()

		_, err = client.SelectMailbox(mailShowMailbox)
		if err != nil {
			output.Fatal("IMAP_ERROR", err)
		}

		body, envelope, err := client.FetchMessage(mail.UID(mailShowUID))
		if err != nil {
			output.Fatal("IMAP_ERROR", err)
		}

		result := map[string]interface{}{
			"uid":  mailShowUID,
			"body": string(body),
		}

		if envelope != nil {
			result["from"] = map[string]string{
				"email": envelope.FromAddr,
				"name":  envelope.FromName,
			}
			result["subject"] = envelope.Subject
			result["date"] = time.Unix(envelope.Date, 0).Format(time.RFC3339)
			result["message_id"] = envelope.MessageID
		}

		output.JSON(result)
	},
}

// --- mail fetch ---

var (
	mailFetchMailbox string
	mailFetchUID     uint32
	mailFetchOutput  string
)

var mailFetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Download email as .eml file",
	Long: `Fetch an email by UID and save it as a .eml file.

Examples:
  lark mail fetch --uid 12345
  lark mail fetch --uid 12345 -o /path/to/output/`,
	Run: func(cmd *cobra.Command, args []string) {
		if mailFetchUID == 0 {
			output.Fatalf("VALIDATION_ERROR", "--uid is required")
		}

		client, err := mail.Connect()
		if err != nil {
			output.Fatal("CONNECTION_ERROR", err)
		}
		defer client.Close()

		_, err = client.SelectMailbox(mailFetchMailbox)
		if err != nil {
			output.Fatal("IMAP_ERROR", err)
		}

		body, envelope, err := client.FetchMessage(mail.UID(mailFetchUID))
		if err != nil {
			output.Fatal("IMAP_ERROR", err)
		}

		// Build filename
		date := time.Now().Format("2006-01-02")
		subject := "email"
		if envelope != nil {
			if envelope.Date != 0 {
				date = time.Unix(envelope.Date, 0).Format("2006-01-02")
			}
			if envelope.Subject != "" {
				subject = sanitizeFilename(envelope.Subject)
				if len(subject) > 60 {
					subject = subject[:60]
				}
			}
		}

		filename := fmt.Sprintf("%s %s.eml", date, subject)
		outpath := filepath.Join(mailFetchOutput, filename)

		if err := os.MkdirAll(mailFetchOutput, 0755); err != nil {
			output.Fatal("IO_ERROR", err)
		}

		if err := os.WriteFile(outpath, body, 0644); err != nil {
			output.Fatal("IO_ERROR", err)
		}

		output.JSON(map[string]interface{}{
			"success":  true,
			"uid":      mailFetchUID,
			"filename": filename,
			"path":     outpath,
			"size":     len(body),
		})
	},
}

func sanitizeFilename(s string) string {
	replacer := strings.NewReplacer(
		"/", "-",
		"\\", "-",
		":", "-",
		"*", "",
		"?", "",
		"\"", "",
		"<", "",
		">", "",
		"|", "",
		"\n", " ",
		"\r", "",
	)
	return strings.TrimSpace(replacer.Replace(s))
}

func formatFreshness(t time.Time) string {
	if t.IsZero() {
		return "never synced"
	}

	d := time.Since(t)

	switch {
	case d < time.Minute:
		return "just now"
	case d < time.Hour:
		mins := int(d.Minutes())
		if mins == 1 {
			return "1 minute ago"
		}
		return fmt.Sprintf("%d minutes ago", mins)
	case d < 24*time.Hour:
		hours := int(d.Hours())
		if hours == 1 {
			return "1 hour ago"
		}
		return fmt.Sprintf("%d hours ago", hours)
	default:
		days := int(d.Hours() / 24)
		if days == 1 {
			return "1 day ago"
		}
		return fmt.Sprintf("%d days ago", days)
	}
}

func init() {
	// mail sync flags
	mailSyncCmd.Flags().StringVarP(&mailSyncMailbox, "mailbox", "m", "INBOX", "Mailbox to sync")

	// mail search flags
	mailSearchCmd.Flags().StringVarP(&mailSearchMailbox, "mailbox", "m", "INBOX", "Mailbox to search")
	mailSearchCmd.Flags().StringVar(&mailSearchFrom, "from", "", "Filter by sender address")
	mailSearchCmd.Flags().StringVar(&mailSearchSubject, "subject", "", "Filter by subject")
	mailSearchCmd.Flags().StringVar(&mailSearchSince, "since", "", "Emails since date (YYYY-MM-DD)")
	mailSearchCmd.Flags().StringVar(&mailSearchBefore, "before", "", "Emails before date (YYYY-MM-DD)")
	mailSearchCmd.Flags().IntVar(&mailSearchLimit, "limit", 50, "Maximum results")

	// mail show flags
	mailShowCmd.Flags().StringVarP(&mailShowMailbox, "mailbox", "m", "INBOX", "Mailbox")
	mailShowCmd.Flags().Uint32Var(&mailShowUID, "uid", 0, "Email UID (required)")

	// mail fetch flags
	mailFetchCmd.Flags().StringVarP(&mailFetchMailbox, "mailbox", "m", "INBOX", "Mailbox")
	mailFetchCmd.Flags().Uint32Var(&mailFetchUID, "uid", 0, "Email UID (required)")
	mailFetchCmd.Flags().StringVarP(&mailFetchOutput, "output", "o", ".", "Output directory")

	// Register subcommands
	mailCmd.AddCommand(mailSetupCmd)
	mailCmd.AddCommand(mailStatusCmd)
	mailCmd.AddCommand(mailListCmd)
	mailCmd.AddCommand(mailSyncCmd)
	mailCmd.AddCommand(mailSearchCmd)
	mailCmd.AddCommand(mailShowCmd)
	mailCmd.AddCommand(mailFetchCmd)
}
