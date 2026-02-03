package cmd

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/yjwong/lark-cli/internal/api"
	"github.com/yjwong/lark-cli/internal/output"
)

var docCmd = &cobra.Command{
	Use:   "doc",
	Short: "Document commands",
	Long:  "Query and retrieve Lark document content",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		validateScopeGroup("documents")
	},
}

// --- doc get ---

var docGetCmd = &cobra.Command{
	Use:   "get <document_id>",
	Short: "Get document content as markdown",
	Long: `Retrieve the content of a Lark document as markdown.

The document_id is the token from the document URL.
For example, if the URL is https://xxx.larksuite.com/docx/ABC123xyz
then the document_id is ABC123xyz.

Examples:
  lark doc get ABC123xyz`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		documentID := args[0]

		client := api.NewClient()

		// Get document metadata for title
		doc, err := client.GetDocument(documentID)
		if err != nil {
			output.Fatal("API_ERROR", err)
		}

		// Get document content as markdown
		content, err := client.GetDocumentContent(documentID)
		if err != nil {
			output.Fatal("API_ERROR", err)
		}

		var title string
		if doc != nil {
			title = doc.Title
		}

		result := api.OutputDocumentContent{
			DocumentID: documentID,
			Title:      title,
			Content:    content,
		}

		output.JSON(result)
	},
}

// --- doc blocks ---

var docBlocksCmd = &cobra.Command{
	Use:   "blocks <document_id>",
	Short: "Get document block structure",
	Long: `Retrieve the full block structure of a Lark document.

Returns the document as a tree of blocks, useful for
programmatic manipulation of document content.

The document_id is the token from the document URL.

Examples:
  lark doc blocks ABC123xyz`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		documentID := args[0]

		client := api.NewClient()

		// Get document metadata for title
		doc, err := client.GetDocument(documentID)
		if err != nil {
			output.Fatal("API_ERROR", err)
		}

		// Get all blocks
		blocks, err := client.GetDocumentBlocks(documentID)
		if err != nil {
			output.Fatal("API_ERROR", err)
		}

		var title string
		if doc != nil {
			title = doc.Title
		}

		result := api.OutputDocumentBlocks{
			DocumentID: documentID,
			Title:      title,
			BlockCount: len(blocks),
			Blocks:     blocks,
		}

		output.JSON(result)
	},
}

// --- doc list ---

var docListCmd = &cobra.Command{
	Use:   "list [folder_token]",
	Short: "List items in a Lark Drive folder",
	Long: `List items in a Lark Drive folder. If no folder_token is provided,
lists items in the root of the user's cloud space.

Item types include: doc, docx, sheet, bitable, mindnote, file, folder, shortcut

Examples:
  lark doc list                    # List root folder
  lark doc list fldbcRho46N6...    # List specific folder`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var folderToken string
		if len(args) > 0 {
			folderToken = args[0]
		}

		client := api.NewClient()

		var allItems []api.FolderItem
		var pageToken string
		for {
			items, hasMore, nextToken, err := client.ListFolderItems(folderToken, 200, pageToken)
			if err != nil {
				output.Fatal("API_ERROR", err)
			}
			allItems = append(allItems, items...)
			if !hasMore {
				break
			}
			pageToken = nextToken
		}

		outputItems := make([]api.OutputFolderItem, len(allItems))
		for i, item := range allItems {
			outputItems[i] = api.OutputFolderItem{
				Token:        item.Token,
				Name:         item.Name,
				Type:         item.Type,
				ParentToken:  item.ParentToken,
				URL:          item.URL,
				ShortcutInfo: item.ShortcutInfo,
			}
		}

		result := api.OutputFolderItemsList{
			FolderToken: folderToken,
			Items:       outputItems,
			Count:       len(outputItems),
		}

		output.JSON(result)
	},
}

// --- doc wiki ---

var docWikiCmd = &cobra.Command{
	Use:   "wiki <node_token>",
	Short: "Resolve wiki node to document token",
	Long: `Resolve a wiki node token to get the underlying document information.

The node_token is from the wiki URL.
For example, if the URL is https://xxx.larksuite.com/wiki/ABC123xyz
then the node_token is ABC123xyz.

This returns the obj_token (document ID) which can be used with 'doc get'.

Examples:
  lark doc wiki X8Tawq431ifOYSklP2tlamKsgNh`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		nodeToken := args[0]

		client := api.NewClient()

		node, err := client.GetWikiNode(nodeToken)
		if err != nil {
			output.Fatal("API_ERROR", err)
		}

		result := api.OutputWikiNode{
			NodeToken: node.NodeToken,
			ObjToken:  node.ObjToken,
			ObjType:   node.ObjType,
			Title:     node.Title,
			SpaceID:   node.SpaceID,
			NodeType:  node.NodeType,
			HasChild:  node.HasChild,
		}

		output.JSON(result)
	},
}

// --- doc wiki-children ---

var docWikiChildrenCmd = &cobra.Command{
	Use:   "wiki-children <node_token>",
	Short: "List child nodes of a wiki node",
	Long: `List the immediate child nodes of a wiki node.

The node_token is from the wiki URL.
For example, if the URL is https://xxx.larksuite.com/wiki/ABC123xyz
then the node_token is ABC123xyz.

This first resolves the node to get the space_id, then fetches its children.

Examples:
  lark doc wiki-children RBCmwZEqhili9ZkKS5fl1Ov2gKc`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		nodeToken := args[0]

		client := api.NewClient()

		// First resolve the node to get space_id
		node, err := client.GetWikiNode(nodeToken)
		if err != nil {
			output.Fatal("API_ERROR", err)
		}

		// Then get children
		children, err := client.GetWikiNodeChildren(node.SpaceID, nodeToken)
		if err != nil {
			output.Fatal("API_ERROR", err)
		}

		outputChildren := make([]api.OutputWikiNode, len(children))
		for i, child := range children {
			outputChildren[i] = api.OutputWikiNode{
				NodeToken: child.NodeToken,
				ObjToken:  child.ObjToken,
				ObjType:   child.ObjType,
				Title:     child.Title,
				SpaceID:   child.SpaceID,
				NodeType:  child.NodeType,
				HasChild:  child.HasChild,
			}
		}

		result := api.OutputWikiChildren{
			ParentNodeToken: nodeToken,
			SpaceID:         node.SpaceID,
			Children:        outputChildren,
			Count:           len(outputChildren),
		}

		output.JSON(result)
	},
}

// --- doc comments ---

var docCommentsCmd = &cobra.Command{
	Use:   "comments <document_id>",
	Short: "Get document comments",
	Long: `Retrieve all comments from a Lark document.

Returns comments and their replies, including user IDs, timestamps,
quoted text, and comment status (solved/unsolved).

The document_id is the token from the document URL.
For example, if the URL is https://xxx.larksuite.com/docx/ABC123xyz
then the document_id is ABC123xyz.

Examples:
  lark doc comments ABC123xyz`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		documentID := args[0]

		client := api.NewClient()

		comments, err := client.GetDocumentComments(documentID, "docx")
		if err != nil {
			output.Fatal("API_ERROR", err)
		}

		result := convertCommentsToOutput(documentID, comments)
		output.JSON(result)
	},
}

// convertCommentsToOutput converts API comments to CLI output format
func convertCommentsToOutput(fileToken string, comments []api.DocumentComment) api.OutputDocumentComments {
	outputComments := make([]api.OutputDocumentComment, len(comments))

	for i, c := range comments {
		// Convert replies
		replies := make([]api.OutputCommentReply, len(c.ReplyList.Replies))
		for j, r := range c.ReplyList.Replies {
			// Extract text from reply elements
			var text string
			for _, elem := range r.Content.Elements {
				switch elem.Type {
				case "text_run":
					if elem.TextRun != nil {
						text += elem.TextRun.Text
					}
				case "docs_link":
					if elem.DocsLink != nil {
						text += elem.DocsLink.URL
					}
				case "person":
					if elem.Person != nil {
						text += "@" + elem.Person.UserID
					}
				}
			}

			replies[j] = api.OutputCommentReply{
				ReplyID:    r.ReplyID,
				UserID:     r.UserID,
				CreateTime: formatUnixTimestamp(r.CreateTime),
				Text:       text,
			}
		}

		outputComments[i] = api.OutputDocumentComment{
			CommentID:  c.CommentID,
			UserID:     c.UserID,
			CreateTime: formatUnixTimestamp(c.CreateTime),
			IsSolved:   c.IsSolved,
			IsWhole:    c.IsWhole,
			Quote:      c.Quote,
			Replies:    replies,
		}
	}

	return api.OutputDocumentComments{
		FileToken: fileToken,
		Comments:  outputComments,
		Count:     len(outputComments),
	}
}

// formatUnixTimestamp converts a unix timestamp to RFC3339 format
func formatUnixTimestamp(ts int64) string {
	if ts == 0 {
		return ""
	}
	return time.Unix(ts, 0).Format(time.RFC3339)
}

// --- doc wiki-search ---

// objTypeToString converts wiki obj_type int to human-readable string
func objTypeToString(objType int) string {
	switch objType {
	case 1:
		return "doc"
	case 2:
		return "sheet"
	case 3:
		return "bitable"
	case 4:
		return "mindnote"
	case 5:
		return "file"
	case 6:
		return "slide"
	case 7:
		return "wiki"
	case 8:
		return "docx"
	case 9:
		return "folder"
	case 10:
		return "catalog"
	default:
		return fmt.Sprintf("unknown(%d)", objType)
	}
}

var docWikiSearchCmd = &cobra.Command{
	Use:   "wiki-search <query>",
	Short: "Search wiki nodes by keyword",
	Long: `Search for wiki nodes by keyword. Returns wiki nodes the user has permission to view.

Optionally filter by wiki space or search within a specific node's children.

Examples:
  lark doc wiki-search "meeting notes"
  lark doc wiki-search "PRD" --space-id 7344964278161604639`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		query := args[0]
		spaceID, _ := cmd.Flags().GetString("space-id")
		nodeID, _ := cmd.Flags().GetString("node-id")

		// Validate: node-id requires space-id
		if nodeID != "" && spaceID == "" {
			output.Fatal("VALIDATION_ERROR", fmt.Errorf("--node-id requires --space-id"))
		}

		client := api.NewClient()

		results, err := client.SearchWikiNodes(query, spaceID, nodeID)
		if err != nil {
			output.Fatal("API_ERROR", err)
		}

		outputItems := make([]api.OutputWikiSearchItem, len(results))
		for i, item := range results {
			outputItems[i] = api.OutputWikiSearchItem{
				NodeID:   item.NodeID,
				ObjToken: item.ObjToken,
				ObjType:  objTypeToString(item.ObjType),
				Title:    item.Title,
				URL:      item.URL,
				SpaceID:  item.SpaceID,
			}
		}

		result := api.OutputWikiSearchResult{
			Query:   query,
			SpaceID: spaceID,
			Results: outputItems,
			Count:   len(outputItems),
		}

		output.JSON(result)
	},
}

// --- doc search ---

var docSearchCmd = &cobra.Command{
	Use:   "search <query>",
	Short: "Search documents by keyword",
	Long: `Search for documents by keyword. Optionally filter by owner, chat, or document type.

The search returns documents from your Drive that match the query.
Maximum 200 results can be returned.

Document types: doc, sheet, slide, bitable, mindnote, file

Examples:
  lark doc search "project plan"
  lark doc search "budget" --type sheet
  lark doc search "meeting notes" --type doc --type sheet`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		query := args[0]
		ownerIDs, _ := cmd.Flags().GetStringSlice("owner")
		chatIDs, _ := cmd.Flags().GetStringSlice("chat")
		docTypes, _ := cmd.Flags().GetStringSlice("type")

		client := api.NewClient()

		results, total, err := client.SearchDocuments(query, ownerIDs, chatIDs, docTypes)
		if err != nil {
			output.Fatal("API_ERROR", err)
		}

		outputItems := make([]api.OutputDocSearchItem, len(results))
		for i, item := range results {
			outputItems[i] = api.OutputDocSearchItem{
				Token:   item.DocsToken,
				Type:    item.DocsType,
				Title:   item.Title,
				OwnerID: item.OwnerID,
			}
		}

		result := api.OutputDocSearchResult{
			Query:   query,
			Results: outputItems,
			Total:   total,
			Count:   len(outputItems),
		}

		output.JSON(result)
	},
}

// --- doc image ---

var docImageCmd = &cobra.Command{
	Use:   "image <image_token>",
	Short: "Download a document image",
	Long: `Download an image from a Lark document.

The image_token is obtained from the 'doc blocks' command output
for image blocks (block_type 27).

The --doc flag is required to specify which document the image belongs to.
This is needed for authentication with the Lark API.

By default, outputs the binary image data to stdout.
Use -o/--output to save to a file instead.

Examples:
  lark doc image K1TQbpmDuokIq3xq1WVl9J7ygkc --doc ABC123xyz > image.png
  lark doc image K1TQbpmDuokIq3xq1WVl9J7ygkc --doc ABC123xyz -o image.png`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		imageToken := args[0]
		outputFile, _ := cmd.Flags().GetString("output")
		documentID, _ := cmd.Flags().GetString("doc")

		if documentID == "" {
			output.Fatal("MISSING_ARG", fmt.Errorf("--doc flag is required to specify the document ID"))
		}

		client := api.NewClient()

		// Download the image
		reader, contentType, err := client.DownloadMedia(imageToken, documentID)
		if err != nil {
			output.Fatal("API_ERROR", err)
		}
		defer reader.Close()

		// Determine output destination
		var writer io.Writer
		if outputFile != "" {
			file, err := os.Create(outputFile)
			if err != nil {
				output.Fatal("FILE_ERROR", err)
			}
			defer file.Close()
			writer = file
		} else {
			writer = os.Stdout
		}

		// Copy image data
		_, err = io.Copy(writer, reader)
		if err != nil {
			output.Fatal("IO_ERROR", err)
		}

		// If writing to file, print success message to stderr
		if outputFile != "" {
			os.Stderr.WriteString("Downloaded image (" + contentType + ") to " + outputFile + "\n")
		}
	},
}

// --- doc download ---

var docDownloadCmd = &cobra.Command{
	Use:   "download <file_token>",
	Short: "Download a file from Lark Drive",
	Long: `Download a file from Lark Drive.

The file_token is obtained from 'doc list' or 'doc search' output.
Only files with type "file" can be downloaded (not docs, sheets, etc).

You must specify an output filename with -o/--output.

Examples:
  lark doc download FG3obxWuaoftXIx0CvxlQAabcef -o report.pdf
  lark doc download FG3obxWuaoftXIx0CvxlQAabcef -o ~/Downloads/report.pdf`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fileToken := args[0]
		outputPath, _ := cmd.Flags().GetString("output")

		if outputPath == "" {
			output.Fatal("MISSING_ARG", fmt.Errorf("--output/-o flag is required"))
		}

		client := api.NewClient()

		// Download the file
		reader, contentType, err := client.DownloadDriveFile(fileToken)
		if err != nil {
			output.Fatal("API_ERROR", err)
		}
		defer reader.Close()

		// Create output file
		file, err := os.Create(outputPath)
		if err != nil {
			output.Fatal("FILE_ERROR", err)
		}
		defer file.Close()

		// Copy file data
		written, err := io.Copy(file, reader)
		if err != nil {
			output.Fatal("IO_ERROR", err)
		}

		// Output result as JSON
		result := struct {
			FileToken   string `json:"file_token"`
			Filename    string `json:"filename"`
			ContentType string `json:"content_type"`
			Size        int64  `json:"size"`
		}{
			FileToken:   fileToken,
			Filename:    outputPath,
			ContentType: contentType,
			Size:        written,
		}
		output.JSON(result)
	},
}

func init() {
	// Register subcommands
	docCmd.AddCommand(docGetCmd)
	docCmd.AddCommand(docBlocksCmd)
	docCmd.AddCommand(docListCmd)
	docCmd.AddCommand(docWikiCmd)
	docCmd.AddCommand(docWikiChildrenCmd)
	docCmd.AddCommand(docCommentsCmd)
	docCmd.AddCommand(docSearchCmd)
	docCmd.AddCommand(docImageCmd)
	docCmd.AddCommand(docWikiSearchCmd)
	docCmd.AddCommand(docDownloadCmd)

	// Flags for doc wiki-search
	docWikiSearchCmd.Flags().String("space-id", "", "Filter to specific wiki space ID")
	docWikiSearchCmd.Flags().String("node-id", "", "Search within a node and its children (requires --space-id)")

	// Flags for doc search
	docSearchCmd.Flags().StringSlice("owner", nil, "Filter by owner user ID (can be repeated)")
	docSearchCmd.Flags().StringSlice("chat", nil, "Filter by chat ID (can be repeated)")
	docSearchCmd.Flags().StringSlice("type", nil, "Filter by doc type: doc, sheet, slide, bitable, mindnote, file (can be repeated)")

	// Flags for doc image
	docImageCmd.Flags().StringP("output", "o", "", "Output file path (default: stdout)")
	docImageCmd.Flags().StringP("doc", "d", "", "Document ID (required for authentication)")

	// Flags for doc download
	docDownloadCmd.Flags().StringP("output", "o", "", "Output file path (default: original filename)")
}
