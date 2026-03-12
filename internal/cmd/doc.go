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
				CreatedTime:  formatUnixTimestamp(parseUnixStr(item.CreatedTime)),
				ModifiedTime: formatUnixTimestamp(parseUnixStr(item.ModifiedTime)),
				OwnerID:      item.OwnerID,
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

// parseUnixStr parses a unix timestamp string to int64
func parseUnixStr(s string) int64 {
	n, _ := strconv.ParseInt(s, 10, 64)
	return n
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

// --- doc create ---

var docCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new document",
	Long: `Create a new Lark document.

Creates an empty document with the specified title.
Optionally specify a folder to create the document in.

Examples:
  lark doc create --title "My Document"
  lark doc create --title "Project Plan" --folder fldbcRho46N6...`,
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		folderToken, _ := cmd.Flags().GetString("folder")

		if title == "" {
			output.Fatal("MISSING_ARG", fmt.Errorf("--title is required"))
		}

		client := api.NewClient()

		doc, err := client.CreateDocument(title, folderToken)
		if err != nil {
			output.Fatal("API_ERROR", err)
		}

		result := api.OutputDocumentCreate{
			Success:    true,
			DocumentID: doc.DocumentID,
			RevisionID: doc.RevisionID,
			Title:      doc.Title,
		}

		output.JSON(result)
	},
}

// --- doc append ---

// makeTextBlock creates a TextBlock with a single text run
func makeTextBlock(content string) *api.TextBlock {
	return &api.TextBlock{
		Elements: []api.TextElement{
			{
				TextRun: &api.TextRun{
					Content: content,
				},
			},
		},
	}
}

// makeLinkedTextBlock creates a TextBlock with a hyperlinked text run
func makeLinkedTextBlock(content, url string) *api.TextBlock {
	return &api.TextBlock{
		Elements: []api.TextElement{
			{
				TextRun: &api.TextRun{
					Content: content,
					TextElementStyle: &api.TextElementStyle{
						Link: &api.TextLink{URL: url},
					},
				},
			},
		},
	}
}

// blockTypeForHeadingLevel returns the block type number for a heading level
func blockTypeForHeadingLevel(level int) int {
	// Block types: heading1=3, heading2=4, ..., heading9=11
	return level + 2
}

// setHeadingField sets the appropriate heading field on a DocumentBlock based on level
func setHeadingField(block *api.DocumentBlock, level int, tb *api.TextBlock) {
	switch level {
	case 1:
		block.Heading1 = tb
	case 2:
		block.Heading2 = tb
	case 3:
		block.Heading3 = tb
	case 4:
		block.Heading4 = tb
	case 5:
		block.Heading5 = tb
	case 6:
		block.Heading6 = tb
	case 7:
		block.Heading7 = tb
	case 8:
		block.Heading8 = tb
	case 9:
		block.Heading9 = tb
	}
}

// BlockBuildOpts contains options for building document blocks
type BlockBuildOpts struct {
	TextContent    string
	HeadingContent string
	HeadingLevel   int
	CodeContent    string
	CodeLanguage   int
	BulletItems    []string
	OrderedItems   []string
	TodoContent    string
	AddDivider     bool
	QuoteContent   string
	LinkURL        string
}

// buildBlocks creates document blocks from the given options
func buildBlocks(opts BlockBuildOpts) []api.DocumentBlock {
	mkBlock := func(content string) *api.TextBlock {
		if opts.LinkURL != "" {
			return makeLinkedTextBlock(content, opts.LinkURL)
		}
		return makeTextBlock(content)
	}

	var blocks []api.DocumentBlock

	if opts.TextContent != "" {
		blocks = append(blocks, api.DocumentBlock{BlockType: 2, Text: mkBlock(opts.TextContent)})
	}

	if opts.HeadingContent != "" {
		level := opts.HeadingLevel
		if level < 1 || level > 9 {
			level = 1
		}
		block := api.DocumentBlock{BlockType: blockTypeForHeadingLevel(level)}
		setHeadingField(&block, level, mkBlock(opts.HeadingContent))
		blocks = append(blocks, block)
	}

	if opts.CodeContent != "" {
		tb := makeTextBlock(opts.CodeContent)
		tb.Style = &api.TextStyle{Language: opts.CodeLanguage}
		blocks = append(blocks, api.DocumentBlock{BlockType: 14, Code: tb})
	}

	for _, item := range opts.BulletItems {
		blocks = append(blocks, api.DocumentBlock{BlockType: 12, Bullet: mkBlock(item)})
	}

	for _, item := range opts.OrderedItems {
		blocks = append(blocks, api.DocumentBlock{BlockType: 13, Ordered: mkBlock(item)})
	}

	if opts.TodoContent != "" {
		blocks = append(blocks, api.DocumentBlock{BlockType: 17, TodoBlock: mkBlock(opts.TodoContent)})
	}

	if opts.AddDivider {
		blocks = append(blocks, api.DocumentBlock{BlockType: 22, Divider: &api.DividerBlock{}})
	}

	if opts.QuoteContent != "" {
		blocks = append(blocks, api.DocumentBlock{BlockType: 15, Quote: mkBlock(opts.QuoteContent)})
	}

	return blocks
}

// buildTableBlocks creates a table block with the given dimensions.
// The Lark API auto-generates cell blocks when a table is created.
// Cell content can be populated via separate append calls to each cell.
func buildTableBlocks(headers []string, rows []string) []api.DocumentBlock {
	if len(headers) == 0 {
		return nil
	}
	colSize := len(headers)
	rowSize := 1 + len(rows) // header row + data rows

	table := api.DocumentBlock{
		BlockType: 31,
		Table: &api.TableBlock{
			Property: &api.TableProperty{
				RowSize:    rowSize,
				ColumnSize: colSize,
				HeaderRow:  true,
			},
		},
	}

	return []api.DocumentBlock{table}
}

// readBlocksFromStdin reads a JSON array of DocumentBlocks from stdin
func readBlocksFromStdin() ([]api.DocumentBlock, error) {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return nil, fmt.Errorf("failed to read stdin: %w", err)
	}
	var blocks []api.DocumentBlock
	if err := json.Unmarshal(data, &blocks); err != nil {
		return nil, fmt.Errorf("invalid block JSON: %w", err)
	}
	return blocks, nil
}

// readBlockFromStdin reads a single DocumentBlock JSON from stdin
func readBlockFromStdin() (api.DocumentBlock, error) {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return api.DocumentBlock{}, fmt.Errorf("failed to read stdin: %w", err)
	}
	var block api.DocumentBlock
	if err := json.Unmarshal(data, &block); err != nil {
		return api.DocumentBlock{}, fmt.Errorf("invalid block JSON: %w", err)
	}
	return block, nil
}

func getStringFlag(cmd *cobra.Command, name string) string {
	v, _ := cmd.Flags().GetString(name)
	return v
}

func getIntFlag(cmd *cobra.Command, name string) int {
	v, _ := cmd.Flags().GetInt(name)
	return v
}

func getBoolFlag(cmd *cobra.Command, name string) bool {
	v, _ := cmd.Flags().GetBool(name)
	return v
}

var docAppendCmd = &cobra.Command{
	Use:   "append <document_id>",
	Short: "Append blocks to a document",
	Long: `Append content blocks to a Lark document.

Supports appending various block types: text, headings, code, bullets,
ordered lists, todos, dividers, and raw JSON blocks.

The document_id is the token from the document URL.

Examples:
  lark doc append ABC123xyz --text "Hello from CLI"
  lark doc append ABC123xyz --heading "Section Title" --level 2
  lark doc append ABC123xyz --code "fmt.Println(\"hello\")" --language 49
  lark doc append ABC123xyz --bullet "First item" --bullet "Second item"
  lark doc append ABC123xyz --ordered "Step 1" --ordered "Step 2"
  lark doc append ABC123xyz --todo "Buy groceries"
  lark doc append ABC123xyz --divider
  echo '[{"block_type":2,"text":{"elements":[{"text_run":{"content":"raw"}}]}}]' | lark doc append ABC123xyz --json`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		documentID := args[0]
		blockID, _ := cmd.Flags().GetString("block-id")
		useJSON, _ := cmd.Flags().GetBool("json")
		index, _ := cmd.Flags().GetInt("index")
		afterBlockID, _ := cmd.Flags().GetString("after")

		// Validate: --after and --index are mutually exclusive
		if afterBlockID != "" && cmd.Flags().Changed("index") {
			output.Fatal("VALIDATION_ERROR", fmt.Errorf("--after and --index are mutually exclusive"))
		}

		if blockID == "" {
			blockID = documentID
		}

		// If --after is set, resolve the target block's parent and index
		if afterBlockID != "" {
			client := api.NewClient()
			allBlocks, err := client.GetDocumentBlocks(documentID)
			if err != nil {
				output.Fatal("API_ERROR", err)
			}

			found := false
			for _, b := range allBlocks {
				if b.Children == nil {
					continue
				}
				for idx, childID := range b.Children {
					if childID == afterBlockID {
						blockID = b.BlockID
						index = idx + 1
						found = true
						break
					}
				}
				if found {
					break
				}
			}
			if !found {
				output.Fatal("NOT_FOUND", fmt.Errorf("block %s not found as a child of any block", afterBlockID))
			}
		}

		useMarkdown, _ := cmd.Flags().GetBool("markdown")

		var blocks []api.DocumentBlock
		if useJSON {
			var err error
			blocks, err = readBlocksFromStdin()
			if err != nil {
				output.Fatal("PARSE_ERROR", err)
			}
		} else if useMarkdown {
			data, err := io.ReadAll(os.Stdin)
			if err != nil {
				output.Fatal("READ_ERROR", err)
			}
			blocks = parseMarkdownToBlocks(data)
		} else {
			bulletItems, _ := cmd.Flags().GetStringArray("bullet")
			orderedItems, _ := cmd.Flags().GetStringArray("ordered")
			blocks = buildBlocks(BlockBuildOpts{
				TextContent:    getStringFlag(cmd, "text"),
				HeadingContent: getStringFlag(cmd, "heading"),
				HeadingLevel:   getIntFlag(cmd, "level"),
				CodeContent:    getStringFlag(cmd, "code"),
				CodeLanguage:   getIntFlag(cmd, "language"),
				BulletItems:    bulletItems,
				OrderedItems:   orderedItems,
				TodoContent:    getStringFlag(cmd, "todo"),
				AddDivider:     getBoolFlag(cmd, "divider"),
				QuoteContent:   getStringFlag(cmd, "quote"),
				LinkURL:        getStringFlag(cmd, "link"),
			})
		}

		// Handle table creation separately: create table, then populate cells
		tableHeaders, _ := cmd.Flags().GetStringArray("table-header")
		tableRows, _ := cmd.Flags().GetStringArray("table-row")

		if len(blocks) == 0 && len(tableHeaders) == 0 {
			output.Fatal("MISSING_ARG", fmt.Errorf("at least one content flag is required (--text, --heading, --code, --bullet, --ordered, --todo, --divider, --quote, --table-header, --json, or --markdown)"))
		}

		client := api.NewClient()

		// Create non-table blocks first
		var allCreatedBlocks []api.DocumentBlock
		var revisionID int
		if len(blocks) > 0 {
			var err error
			allCreatedBlocks, revisionID, err = client.CreateDocumentBlocks(documentID, blockID, blocks, index)
			if err != nil {
				output.Fatal("API_ERROR", err)
			}
		}

		// Create table and auto-populate cells
		if len(tableHeaders) > 0 {
			tableBlocks := buildTableBlocks(tableHeaders, tableRows)
			createdTableBlocks, tableRevID, err := client.CreateDocumentBlocks(documentID, blockID, tableBlocks, index)
			if err != nil {
				output.Fatal("API_ERROR", err)
			}
			revisionID = tableRevID

			// Populate cells with content
			for _, tb := range createdTableBlocks {
				if tb.BlockType == 31 && tb.Table != nil && len(tb.Table.Cells) > 0 {
					// Build flat list of cell contents: headers first, then rows
					var cellContents []string
					cellContents = append(cellContents, tableHeaders...)
					colSize := len(tableHeaders)
					for _, row := range tableRows {
						cells := strings.Split(row, "|")
						for i := 0; i < colSize; i++ {
							if i < len(cells) {
								cellContents = append(cellContents, strings.TrimSpace(cells[i]))
							} else {
								cellContents = append(cellContents, "")
							}
						}
					}

					// Populate each cell (with delay and retry to handle API rate limits)
					for i, cellID := range tb.Table.Cells {
						if i < len(cellContents) && cellContents[i] != "" {
							time.Sleep(200 * time.Millisecond)
							textBlock := []api.DocumentBlock{{BlockType: 2, Text: makeTextBlock(cellContents[i])}}
							_, _, cellErr := client.CreateDocumentBlocks(documentID, cellID, textBlock, -1)
							if cellErr != nil {
								// Retry once after a longer delay
								time.Sleep(500 * time.Millisecond)
								_, _, cellErr = client.CreateDocumentBlocks(documentID, cellID, textBlock, -1)
								if cellErr != nil {
									output.Fatal("API_ERROR", fmt.Errorf("failed to populate cell %d: %w", i, cellErr))
								}
							}
						}
					}
				}
			}
			allCreatedBlocks = append(allCreatedBlocks, createdTableBlocks...)
		}

		output.JSON(api.OutputDocumentAppend{
			Success:            true,
			DocumentRevisionID: revisionID,
			Blocks:             allCreatedBlocks,
		})
	},
}

// --- doc upload ---

var docUploadCmd = &cobra.Command{
	Use:   "upload <file_path>",
	Short: "Upload a file to Lark Drive",
	Long: `Upload a local file to Lark Drive (max 20MB).

Optionally specify a folder to upload into using --folder.

Examples:
  lark doc upload report.pdf
  lark doc upload data.csv --folder fldbcRho46N6...`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		folderToken, _ := cmd.Flags().GetString("folder")

		// Verify file exists
		info, err := os.Stat(filePath)
		if err != nil {
			output.Fatal("FILE_ERROR", err)
		}

		client := api.NewClient()

		fileToken, err := client.UploadDriveFile(filePath, folderToken, "explorer")
		if err != nil {
			output.Fatal("API_ERROR", err)
		}

		result := api.OutputDriveUpload{
			FileToken: fileToken,
			FileName:  info.Name(),
			Size:      info.Size(),
		}

		output.JSON(result)
	},
}

// --- doc info ---

var docInfoCmd = &cobra.Command{
	Use:   "info <token>",
	Short: "Get file or folder metadata",
	Long: `Get metadata for a Lark Drive file or folder.

Returns title, owner, creation time, last modified time, and URL.
The --type flag specifies the document type (default: file).

Supported types: doc, docx, sheet, bitable, folder, file, mindnote, slides, wiki

Examples:
  lark doc info fldbcRho46N6... --type folder
  lark doc info Mbxmsn4eRha6... --type sheet`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		token := args[0]
		docType, _ := cmd.Flags().GetString("type")

		client := api.NewClient()

		if !cmd.Flags().Changed("type") {
			// Auto-detect: try common types in order
			for _, t := range []string{"docx", "sheet", "bitable", "folder", "file", "doc", "mindnote", "slides", "wiki"} {
				meta, err := client.GetDriveMeta(token, t)
				if err == nil {
					result := api.OutputDriveInfo{
						Token:            meta.DocToken,
						Type:             meta.DocType,
						Title:            meta.Title,
						OwnerID:          meta.OwnerID,
						CreateTime:       formatUnixTimestamp(parseUnixStr(meta.CreateTime)),
						LatestModifyUser: meta.LatestModifyUser,
						LatestModifyTime: formatUnixTimestamp(parseUnixStr(meta.LatestModifyTime)),
						URL:              meta.URL,
					}
					output.JSON(result)
					return
				}
			}
			output.Fatal("API_ERROR", fmt.Errorf("no metadata found for token %s (tried all types)", token))
			return
		}

		if docType == "" {
			docType = "file"
		}

		meta, err := client.GetDriveMeta(token, docType)
		if err != nil {
			output.Fatal("API_ERROR", err)
		}

		result := api.OutputDriveInfo{
			Token:            meta.DocToken,
			Type:             meta.DocType,
			Title:            meta.Title,
			OwnerID:          meta.OwnerID,
			CreateTime:       formatUnixTimestamp(parseUnixStr(meta.CreateTime)),
			LatestModifyUser: meta.LatestModifyUser,
			LatestModifyTime: formatUnixTimestamp(parseUnixStr(meta.LatestModifyTime)),
			URL:              meta.URL,
		}
		output.JSON(result)
	},
}

// --- doc mkdir ---

var docMkdirCmd = &cobra.Command{
	Use:   "mkdir <name>",
	Short: "Create a new folder in Lark Drive",
	Long: `Create a new folder in Lark Drive.

By default creates in the root of your cloud space.
Use --folder to specify a parent folder.

Examples:
  lark doc mkdir "Project Files"
  lark doc mkdir "Reports" --folder fldbcRho46N6...`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		parentToken, _ := cmd.Flags().GetString("folder")

		client := api.NewClient()
		token, url, err := client.CreateFolder(name, parentToken)
		if err != nil {
			output.Fatal("API_ERROR", err)
		}

		output.JSON(api.OutputCreateFolder{
			Token: token,
			Name:  name,
			URL:   url,
		})
	},
}

// --- doc delete ---

var docDeleteCmd = &cobra.Command{
	Use:   "delete <document_id> <block_id> [block_id...]",
	Short: "Delete blocks from a document",
	Long: `Delete one or more blocks from a Lark document.

The document_id is the token from the document URL.
Block IDs can be found using 'doc blocks'.

Examples:
  lark doc delete ABC123xyz doxlgXYZ123
  lark doc delete ABC123xyz doxlgXYZ123 doxlgABC456`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		documentID := args[0]
		blockIDs := args[1:]

		client := api.NewClient()

		revisionID, err := client.DeleteDocumentBlocks(documentID, blockIDs)
		if err != nil {
			output.Fatal("API_ERROR", err)
		}

		result := api.OutputDocumentDelete{
			Success:            true,
			DocumentRevisionID: revisionID,
			DeletedBlockIDs:    blockIDs,
		}

		output.JSON(result)
	},
}

// --- doc update ---

var docUpdateCmd = &cobra.Command{
	Use:   "update <document_id> <block_id>",
	Short: "Update a block in a document",
	Long: `Update the content of a block in a Lark document.

Supports updating text, heading, code, bullet, ordered, and todo blocks.
The block type must match the existing block's type.

Block IDs can be found using 'doc blocks'.

Examples:
  lark doc update ABC123xyz doxlgXYZ123 --text "Updated content"
  lark doc update ABC123xyz doxlgXYZ123 --heading "New Title" --level 2
  lark doc update ABC123xyz doxlgXYZ123 --code "fmt.Println()" --language 22
  echo '{"block_type":2,"text":{"elements":[{"text_run":{"content":"raw"}}]}}' | lark doc update ABC123xyz doxlgXYZ123 --json`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		documentID := args[0]
		blockID := args[1]
		useJSON, _ := cmd.Flags().GetBool("json")

		var block api.DocumentBlock
		if useJSON {
			var err error
			block, err = readBlockFromStdin()
			if err != nil {
				output.Fatal("PARSE_ERROR", err)
			}
		} else {
			bulletContent := getStringFlag(cmd, "bullet")
			orderedContent := getStringFlag(cmd, "ordered")
			var bulletItems, orderedItems []string
			if bulletContent != "" {
				bulletItems = []string{bulletContent}
			}
			if orderedContent != "" {
				orderedItems = []string{orderedContent}
			}
			blocks := buildBlocks(BlockBuildOpts{
				TextContent:    getStringFlag(cmd, "text"),
				HeadingContent: getStringFlag(cmd, "heading"),
				HeadingLevel:   getIntFlag(cmd, "level"),
				CodeContent:    getStringFlag(cmd, "code"),
				CodeLanguage:   getIntFlag(cmd, "language"),
				BulletItems:    bulletItems,
				OrderedItems:   orderedItems,
				TodoContent:    getStringFlag(cmd, "todo"),
				LinkURL:        getStringFlag(cmd, "link"),
			})
			if len(blocks) == 0 {
				output.Fatal("MISSING_ARG", fmt.Errorf("at least one content flag is required (--text, --heading, --code, --bullet, --ordered, --todo, --quote, or --json)"))
			}
			block = blocks[0]
		}

		client := api.NewClient()
		revisionID, err := client.UpdateDocumentBlock(documentID, blockID, block)
		if err != nil {
			output.Fatal("API_ERROR", err)
		}
		output.JSON(api.OutputDocumentUpdate{
			Success:            true,
			DocumentRevisionID: revisionID,
		})
	},
}

// --- doc trash ---

var docTrashCmd = &cobra.Command{
	Use:   "trash <file_token>",
	Short: "Move a file to trash in Lark Drive",
	Long: `Move a file or document to trash in Lark Drive.

The file_token is the document or file token.
Use --type to specify the document type (default: docx).

Supported types: doc, docx, sheet, bitable, folder, file, mindnote, slides

Examples:
  lark doc trash ABC123xyz
  lark doc trash ABC123xyz --type sheet
  lark doc trash fldbcRho46N6... --type folder`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fileToken := args[0]
		docType, _ := cmd.Flags().GetString("type")

		client := api.NewClient()

		err := client.DeleteDriveFile(fileToken, docType)
		if err != nil {
			output.Fatal("API_ERROR", err)
		}

		result := api.OutputDriveTrash{
			Success:   true,
			FileToken: fileToken,
			Type:      docType,
		}

		output.JSON(result)
	},
}

// --- doc find ---

// extractBlockText extracts the text content from a block's elements
func extractBlockText(block api.DocumentBlock) string {
	var textBlock *api.TextBlock
	switch {
	case block.Text != nil:
		textBlock = block.Text
	case block.Heading1 != nil:
		textBlock = block.Heading1
	case block.Heading2 != nil:
		textBlock = block.Heading2
	case block.Heading3 != nil:
		textBlock = block.Heading3
	case block.Heading4 != nil:
		textBlock = block.Heading4
	case block.Heading5 != nil:
		textBlock = block.Heading5
	case block.Heading6 != nil:
		textBlock = block.Heading6
	case block.Heading7 != nil:
		textBlock = block.Heading7
	case block.Heading8 != nil:
		textBlock = block.Heading8
	case block.Heading9 != nil:
		textBlock = block.Heading9
	case block.Bullet != nil:
		textBlock = block.Bullet
	case block.Ordered != nil:
		textBlock = block.Ordered
	case block.TodoBlock != nil:
		textBlock = block.TodoBlock
	case block.Code != nil:
		textBlock = block.Code
	}

	if textBlock == nil {
		return ""
	}

	var parts []string
	for _, elem := range textBlock.Elements {
		if elem.TextRun != nil {
			parts = append(parts, elem.TextRun.Content)
		}
	}
	return strings.Join(parts, "")
}

var docFindCmd = &cobra.Command{
	Use:   "find <document_id> <query>",
	Short: "Search for blocks containing text",
	Long: `Search for blocks in a document that contain the given text.

Returns matching block IDs with their parent, index, type, and a content preview.
Useful for finding blocks before using doc replace or doc delete.

The search is case-insensitive.

Examples:
  lark doc find ABC123xyz "Section Title"
  lark doc find ABC123xyz "TODO" --type 17`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		documentID := args[0]
		query := args[1]
		filterType, _ := cmd.Flags().GetInt("type")

		client := api.NewClient()

		blocks, err := client.GetDocumentBlocks(documentID)
		if err != nil {
			output.Fatal("API_ERROR", err)
		}

		// Build parent -> children index map
		childIndex := make(map[string]map[string]int) // parentID -> blockID -> index
		for _, block := range blocks {
			if block.Children == nil {
				continue
			}
			m := make(map[string]int)
			for idx, childID := range block.Children {
				m[childID] = idx
			}
			childIndex[block.BlockID] = m
		}

		queryLower := strings.ToLower(query)
		var matches []api.OutputFindMatch

		for _, block := range blocks {
			if filterType > 0 && block.BlockType != filterType {
				continue
			}

			text := extractBlockText(block)
			if text == "" {
				continue
			}

			if !strings.Contains(strings.ToLower(text), queryLower) {
				continue
			}

			// Truncate preview
			preview := text
			if len(preview) > 120 {
				preview = preview[:120] + "..."
			}

			// Find index within parent
			idx := -1
			if m, ok := childIndex[block.ParentID]; ok {
				if i, ok := m[block.BlockID]; ok {
					idx = i
				}
			}

			matches = append(matches, api.OutputFindMatch{
				BlockID:   block.BlockID,
				ParentID:  block.ParentID,
				Index:     idx,
				BlockType: block.BlockType,
				Preview:   preview,
			})
		}

		result := api.OutputDocumentFind{
			DocumentID: documentID,
			Query:      query,
			Matches:    matches,
			Count:      len(matches),
		}

		output.JSON(result)
	},
}

// --- doc replace ---

var docReplaceCmd = &cobra.Command{
	Use:   "replace <document_id> <block_id>",
	Short: "Replace a block with new content",
	Long: `Replace a block in a document with new content.

Atomically deletes the specified block and inserts new content
at the same position. Supports the same content flags as append.

Block IDs can be found using 'doc find' or 'doc blocks'.

Examples:
  lark doc replace ABC123xyz doxlgXYZ123 --text "Updated content"
  lark doc replace ABC123xyz doxlgXYZ123 --heading "New Title" --level 2
  lark doc replace ABC123xyz doxlgXYZ123 --bullet "Item 1" --bullet "Item 2"
  echo '[{"block_type":2,"text":{"elements":[{"text_run":{"content":"raw"}}]}}]' | lark doc replace ABC123xyz doxlgXYZ123 --json`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		documentID := args[0]
		blockID := args[1]
		useJSON, _ := cmd.Flags().GetBool("json")

		useMarkdown, _ := cmd.Flags().GetBool("markdown")

		// Build the new blocks
		var newBlocks []api.DocumentBlock
		if useJSON {
			var err error
			newBlocks, err = readBlocksFromStdin()
			if err != nil {
				output.Fatal("PARSE_ERROR", err)
			}
		} else if useMarkdown {
			data, err := io.ReadAll(os.Stdin)
			if err != nil {
				output.Fatal("READ_ERROR", err)
			}
			newBlocks = parseMarkdownToBlocks(data)
		} else {
			bulletItems, _ := cmd.Flags().GetStringArray("bullet")
			orderedItems, _ := cmd.Flags().GetStringArray("ordered")
			newBlocks = buildBlocks(BlockBuildOpts{
				TextContent:    getStringFlag(cmd, "text"),
				HeadingContent: getStringFlag(cmd, "heading"),
				HeadingLevel:   getIntFlag(cmd, "level"),
				CodeContent:    getStringFlag(cmd, "code"),
				CodeLanguage:   getIntFlag(cmd, "language"),
				BulletItems:    bulletItems,
				OrderedItems:   orderedItems,
				TodoContent:    getStringFlag(cmd, "todo"),
				AddDivider:     getBoolFlag(cmd, "divider"),
				QuoteContent:   getStringFlag(cmd, "quote"),
				LinkURL:        getStringFlag(cmd, "link"),
			})
		}

		if len(newBlocks) == 0 {
			output.Fatal("MISSING_ARG", fmt.Errorf("at least one content flag is required (--text, --heading, --code, --bullet, --ordered, --todo, --divider, --quote, --json, or --markdown)"))
		}

		client := api.NewClient()

		createdBlocks, revisionID, err := client.ReplaceDocumentBlock(documentID, blockID, newBlocks)
		if err != nil {
			output.Fatal("API_ERROR", err)
		}

		output.JSON(api.OutputDocumentAppend{
			Success:            true,
			DocumentRevisionID: revisionID,
			Blocks:             createdBlocks,
		})
	},
}

// --- doc outline ---

var docOutlineCmd = &cobra.Command{
	Use:   "outline <document_id>",
	Short: "Show document heading outline",
	Long: `Show the heading outline of a Lark document.

Returns only heading blocks (H1-H9) with their block IDs, positions,
and text content. Useful for understanding document structure.

The document_id is the token from the document URL.

Examples:
  lark doc outline ABC123xyz`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		documentID := args[0]

		client := api.NewClient()

		doc, err := client.GetDocument(documentID)
		if err != nil {
			output.Fatal("API_ERROR", err)
		}

		blocks, err := client.GetDocumentBlocks(documentID)
		if err != nil {
			output.Fatal("API_ERROR", err)
		}

		// Build parent -> child index map
		childIndex := make(map[string]map[string]int)
		for _, block := range blocks {
			if block.Children == nil {
				continue
			}
			m := make(map[string]int)
			for idx, childID := range block.Children {
				m[childID] = idx
			}
			childIndex[block.BlockID] = m
		}

		var outline []api.OutlineEntry
		for _, block := range blocks {
			// Heading block types: 3=H1, 4=H2, ..., 11=H9
			if block.BlockType < 3 || block.BlockType > 11 {
				continue
			}

			level := block.BlockType - 2
			text := extractBlockText(block)

			idx := -1
			if m, ok := childIndex[block.ParentID]; ok {
				if i, ok := m[block.BlockID]; ok {
					idx = i
				}
			}

			outline = append(outline, api.OutlineEntry{
				BlockID: block.BlockID,
				Index:   idx,
				Level:   level,
				Text:    text,
			})
		}

		var title string
		if doc != nil {
			title = doc.Title
		}

		result := api.OutputDocumentOutline{
			DocumentID: documentID,
			Title:      title,
			Outline:    outline,
			Count:      len(outline),
		}

		output.JSON(result)
	},
}

// --- doc move ---

var docMoveCmd = &cobra.Command{
	Use:   "move <document_id> <block_id>",
	Short: "Move a block to a new position",
	Long: `Move a block to a new position in a Lark document.

Relocates a block by deleting it from its current position and
inserting it at the specified new position.

Use --index to specify an absolute position, or --after to insert
after a specific block.

Examples:
  lark doc move ABC123xyz doxlgXYZ123 --index 0
  lark doc move ABC123xyz doxlgXYZ123 --after doxlgABC456`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		documentID := args[0]
		blockID := args[1]
		index, _ := cmd.Flags().GetInt("index")
		afterBlockID, _ := cmd.Flags().GetString("after")

		hasIndex := cmd.Flags().Changed("index")
		hasAfter := afterBlockID != ""

		if hasIndex && hasAfter {
			output.Fatal("VALIDATION_ERROR", fmt.Errorf("--index and --after are mutually exclusive"))
		}
		if !hasIndex && !hasAfter {
			output.Fatal("MISSING_ARG", fmt.Errorf("either --index or --after is required"))
		}

		client := api.NewClient()

		// Get all blocks to find the target block
		blocks, err := client.GetDocumentBlocks(documentID)
		if err != nil {
			output.Fatal("API_ERROR", err)
		}

		// Find the block's current parent and index
		var srcParentID string
		var srcIndex int
		srcFound := false
		for _, b := range blocks {
			if b.Children == nil {
				continue
			}
			for idx, childID := range b.Children {
				if childID == blockID {
					srcParentID = b.BlockID
					srcIndex = idx
					srcFound = true
					break
				}
			}
			if srcFound {
				break
			}
		}
		if !srcFound {
			output.Fatal("NOT_FOUND", fmt.Errorf("block %s not found as a child of any block", blockID))
		}

		// Find the target block content to re-create it
		var targetBlock *api.DocumentBlock
		for i := range blocks {
			if blocks[i].BlockID == blockID {
				targetBlock = &blocks[i]
				break
			}
		}
		if targetBlock == nil {
			output.Fatal("NOT_FOUND", fmt.Errorf("block %s not found", blockID))
		}

		// Resolve --after to a parent + index
		destParentID := srcParentID
		destIndex := index
		if hasAfter {
			found := false
			for _, b := range blocks {
				if b.Children == nil {
					continue
				}
				for idx, childID := range b.Children {
					if childID == afterBlockID {
						destParentID = b.BlockID
						destIndex = idx + 1
						found = true
						break
					}
				}
				if found {
					break
				}
			}
			if !found {
				output.Fatal("NOT_FOUND", fmt.Errorf("--after block %s not found as a child of any block", afterBlockID))
			}
		}

		// Delete the block from its current position
		delPath := fmt.Sprintf("/docx/v1/documents/%s/blocks/%s/children/batch_delete?document_revision_id=-1",
			srcParentID, documentID)
		// Use the proper URL format
		delPath = fmt.Sprintf("/docx/v1/documents/%s/blocks/%s/children/batch_delete?document_revision_id=-1",
			documentID, srcParentID)
		delReq := api.DeleteBlocksRequest{
			StartIndex: srcIndex,
			EndIndex:   srcIndex + 1,
		}
		var delResp api.DeleteBlocksResponse
		if err := client.DeleteWithBody(delPath, delReq, &delResp); err != nil {
			output.Fatal("API_ERROR", fmt.Errorf("failed to delete block from current position: %w", err))
		}
		if err := delResp.Err(); err != nil {
			output.Fatal("API_ERROR", err)
		}

		// Adjust destination index if needed (if in same parent and source was before dest)
		if destParentID == srcParentID && srcIndex < destIndex {
			destIndex--
		}

		// Build a clean block for re-insertion (strip block_id, parent_id, children)
		newBlock := api.DocumentBlock{
			BlockType: targetBlock.BlockType,
			Page:      targetBlock.Page,
			Text:      targetBlock.Text,
			Heading1:  targetBlock.Heading1,
			Heading2:  targetBlock.Heading2,
			Heading3:  targetBlock.Heading3,
			Heading4:  targetBlock.Heading4,
			Heading5:  targetBlock.Heading5,
			Heading6:  targetBlock.Heading6,
			Heading7:  targetBlock.Heading7,
			Heading8:  targetBlock.Heading8,
			Heading9:  targetBlock.Heading9,
			Bullet:    targetBlock.Bullet,
			Ordered:   targetBlock.Ordered,
			Code:      targetBlock.Code,
			Quote:     targetBlock.Quote,
			TodoBlock: targetBlock.TodoBlock,
			Divider:   targetBlock.Divider,
			Image:     targetBlock.Image,
		}

		// Insert at new position
		createdBlocks, revisionID, err := client.CreateDocumentBlocks(documentID, destParentID, []api.DocumentBlock{newBlock}, destIndex)
		if err != nil {
			output.Fatal("API_ERROR", fmt.Errorf("deleted block but failed to insert at new position: %w", err))
		}

		output.JSON(api.OutputDocumentMove{
			Success:            true,
			DocumentRevisionID: revisionID,
			BlockID:            blockID,
			Blocks:             createdBlocks,
		})
	},
}

// codeLanguageHelp returns a string listing code language IDs for the help text
func codeLanguageHelp() string {
	return "Common language IDs: 1=PlainText, 7=Bash, 8=C#, 9=C++, 10=C, 12=CSS, 22=Go, 24=HTML, 28=JSON, 29=Java, 30=JavaScript, 32=Kotlin, 49=Python, 52=Ruby, 53=Rust, 56=SQL, 61=Swift, 63=TypeScript, 67=YAML"
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
	docCmd.AddCommand(docCreateCmd)
	docCmd.AddCommand(docAppendCmd)
	docCmd.AddCommand(docUploadCmd)
	docCmd.AddCommand(docInfoCmd)
	docCmd.AddCommand(docMkdirCmd)
	docCmd.AddCommand(docDeleteCmd)
	docCmd.AddCommand(docUpdateCmd)
	docCmd.AddCommand(docTrashCmd)
	docCmd.AddCommand(docFindCmd)
	docCmd.AddCommand(docReplaceCmd)
	docCmd.AddCommand(docOutlineCmd)
	docCmd.AddCommand(docMoveCmd)

	// Flags for doc update
	docUpdateCmd.Flags().String("text", "", "Update block with text content")
	docUpdateCmd.Flags().String("heading", "", "Update block with heading content")
	docUpdateCmd.Flags().Int("level", 1, "Heading level 1-9 (used with --heading)")
	docUpdateCmd.Flags().String("code", "", "Update block with code content")
	docUpdateCmd.Flags().Int("language", 0, "Code language ID (used with --code). "+codeLanguageHelp())
	docUpdateCmd.Flags().String("bullet", "", "Update block with bullet content")
	docUpdateCmd.Flags().String("ordered", "", "Update block with ordered list content")
	docUpdateCmd.Flags().String("todo", "", "Update block with todo content")
	docUpdateCmd.Flags().Bool("json", false, "Read raw block JSON from stdin")

	// Flags for doc info
	docInfoCmd.Flags().String("type", "file", "Document type: doc, docx, sheet, bitable, folder, file, mindnote, slides, wiki")

	// Flags for doc mkdir
	docMkdirCmd.Flags().String("folder", "", "Parent folder token (default: root)")

	// Flags for doc upload
	docUploadCmd.Flags().String("folder", "", "Folder token to upload into (default: root)")

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

	// Flags for doc create
	docCreateCmd.Flags().String("title", "", "Document title (required)")
	docCreateCmd.Flags().String("folder", "", "Folder token to create document in (default: root)")

	// Flags for doc append
	docAppendCmd.Flags().String("block-id", "", "Parent block ID (default: document root)")
	docAppendCmd.Flags().String("text", "", "Append a text block")
	docAppendCmd.Flags().String("heading", "", "Append a heading block")
	docAppendCmd.Flags().Int("level", 1, "Heading level 1-9 (used with --heading)")
	docAppendCmd.Flags().String("code", "", "Append a code block")
	docAppendCmd.Flags().Int("language", 0, "Code language ID (used with --code). "+codeLanguageHelp())
	docAppendCmd.Flags().StringArray("bullet", nil, "Append bullet list items (repeatable)")
	docAppendCmd.Flags().StringArray("ordered", nil, "Append ordered list items (repeatable)")
	docAppendCmd.Flags().String("todo", "", "Append a todo item")
	docAppendCmd.Flags().Bool("divider", false, "Append a divider")
	docAppendCmd.Flags().String("quote", "", "Append a quote block")
	docAppendCmd.Flags().Bool("json", false, "Read raw block JSON from stdin")
	docAppendCmd.Flags().Bool("markdown", false, "Read markdown from stdin and convert to blocks")
	docAppendCmd.Flags().Int("index", -1, "Insertion position (-1=end, 0=beginning)")
	docAppendCmd.Flags().String("link", "", "Hyperlink URL to apply to the text")
	docAppendCmd.Flags().String("after", "", "Insert after this block ID (mutually exclusive with --index)")
	docAppendCmd.Flags().StringArray("table-header", nil, "Table header cells (repeatable, one per column)")
	docAppendCmd.Flags().StringArray("table-row", nil, "Table row as pipe-separated cells: \"cell1|cell2|cell3\" (repeatable)")

	// Flags for doc update
	docUpdateCmd.Flags().String("link", "", "Hyperlink URL to apply to the text")

	// Flags for doc trash
	docTrashCmd.Flags().String("type", "docx", "Document type: doc, docx, sheet, bitable, folder, file, mindnote, slides")

	// Flags for doc find
	docFindCmd.Flags().Int("type", 0, "Filter by block type (e.g., 2=text, 12=bullet, 14=code)")

	// Flags for doc replace
	docReplaceCmd.Flags().String("text", "", "Replace with text block")
	docReplaceCmd.Flags().String("heading", "", "Replace with heading block")
	docReplaceCmd.Flags().Int("level", 1, "Heading level 1-9 (used with --heading)")
	docReplaceCmd.Flags().String("code", "", "Replace with code block")
	docReplaceCmd.Flags().Int("language", 0, "Code language ID (used with --code). "+codeLanguageHelp())
	docReplaceCmd.Flags().StringArray("bullet", nil, "Replace with bullet list items (repeatable)")
	docReplaceCmd.Flags().StringArray("ordered", nil, "Replace with ordered list items (repeatable)")
	docReplaceCmd.Flags().String("todo", "", "Replace with todo item")
	docReplaceCmd.Flags().Bool("divider", false, "Replace with divider")
	docReplaceCmd.Flags().String("quote", "", "Replace with quote block")
	docReplaceCmd.Flags().Bool("json", false, "Read raw block JSON from stdin")
	docReplaceCmd.Flags().Bool("markdown", false, "Read markdown from stdin and convert to blocks")
	docReplaceCmd.Flags().String("link", "", "Hyperlink URL to apply to the text")

	// Flags for doc move
	docMoveCmd.Flags().Int("index", -1, "Target position index")
	docMoveCmd.Flags().String("after", "", "Insert after this block ID (mutually exclusive with --index)")
}
