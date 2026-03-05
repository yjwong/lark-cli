package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"

	"github.com/yjwong/lark-cli/internal/auth"
)

// GetDocument retrieves document metadata
// documentID: the document ID (token from document URL)
func (c *Client) GetDocument(documentID string) (*Document, error) {
	path := fmt.Sprintf("/docx/v1/documents/%s", url.PathEscape(documentID))

	var resp DocumentResponse
	if err := c.Get(path, &resp); err != nil {
		return nil, err
	}

	if err := resp.Err(); err != nil {
		return nil, err
	}

	return resp.Data.Document, nil
}

// GetDocumentContent retrieves document content as markdown
// documentID: the document ID (token from document URL)
func (c *Client) GetDocumentContent(documentID string) (string, error) {
	path := fmt.Sprintf("/docs/v1/content?doc_token=%s&doc_type=docx&content_type=markdown",
		url.QueryEscape(documentID))

	var resp DocumentContentResponse
	if err := c.Get(path, &resp); err != nil {
		return "", err
	}

	if err := resp.Err(); err != nil {
		return "", err
	}

	return resp.Data.Content, nil
}

// GetDocumentBlocks retrieves all blocks in a document with pagination
// documentID: the document ID (token from document URL)
func (c *Client) GetDocumentBlocks(documentID string) ([]DocumentBlock, error) {
	return PaginateWith(func(pageToken string) ([]DocumentBlock, bool, string, error) {
		path := fmt.Sprintf("/docx/v1/documents/%s/blocks?page_size=500",
			url.PathEscape(documentID))
		if pageToken != "" {
			path += "&page_token=" + url.QueryEscape(pageToken)
		}

		var resp DocumentBlocksResponse
		if err := c.Get(path, &resp); err != nil {
			return nil, false, "", err
		}
		if err := resp.Err(); err != nil {
			return nil, false, "", err
		}
		return resp.Data.Items, resp.Data.HasMore, resp.Data.PageToken, nil
	})
}

// CreateDocument creates a new document
// title: document title
// folderToken: optional folder token (empty for root)
func (c *Client) CreateDocument(title, folderToken string) (*Document, error) {
	req := CreateDocumentRequest{
		Title:       title,
		FolderToken: folderToken,
	}

	var resp DocumentResponse
	if err := c.Post("/docx/v1/documents", req, &resp); err != nil {
		return nil, err
	}

	if err := resp.Err(); err != nil {
		return nil, err
	}

	return resp.Data.Document, nil
}

// CreateDocumentBlocks creates child blocks under a parent block in a document
// documentID: the document ID
// blockID: the parent block ID (use documentID for root page block)
// children: blocks to create
// index: insertion position (-1 for end)
func (c *Client) CreateDocumentBlocks(documentID, blockID string, children []DocumentBlock, index int) ([]DocumentBlock, int, error) {
	path := fmt.Sprintf("/docx/v1/documents/%s/blocks/%s/children?document_revision_id=-1",
		url.PathEscape(documentID), url.PathEscape(blockID))

	req := CreateBlockChildrenRequest{
		Children: children,
		Index:    index,
	}

	var resp CreateBlockChildrenResponse
	if err := c.Post(path, req, &resp); err != nil {
		return nil, 0, err
	}

	if err := resp.Err(); err != nil {
		return nil, 0, err
	}

	return resp.Data.Children, resp.Data.DocumentRevisionID, nil
}

// ListFolderItems lists items in a Lark Drive folder
// folderToken: folder token (empty for root cloud space)
// pageSize: number of items per page (max 200)
// pageToken: pagination token
func (c *Client) ListFolderItems(folderToken string, pageSize int, pageToken string) ([]FolderItem, bool, string, error) {
	params := url.Values{}
	if folderToken != "" {
		params.Set("folder_token", folderToken)
	}
	if pageSize > 0 {
		params.Set("page_size", strconv.Itoa(pageSize))
	}
	if pageToken != "" {
		params.Set("page_token", pageToken)
	}

	path := "/drive/v1/files"
	if len(params) > 0 {
		path += "?" + params.Encode()
	}

	var resp ListFolderItemsResponse
	if err := c.Get(path, &resp); err != nil {
		return nil, false, "", err
	}
	if err := resp.Err(); err != nil {
		return nil, false, "", err
	}

	return resp.Data.Files, resp.Data.HasMore, resp.Data.NextPageToken, nil
}

// GetDocumentComments retrieves all comments for a document with pagination
// fileToken: the document token (same as document ID)
// fileType: document type (e.g., "docx", "doc", "sheet")
func (c *Client) GetDocumentComments(fileToken, fileType string) ([]DocumentComment, error) {
	return PaginateWith(func(pageToken string) ([]DocumentComment, bool, string, error) {
		path := fmt.Sprintf("/drive/v1/files/%s/comments?file_type=%s&page_size=100",
			url.PathEscape(fileToken), url.QueryEscape(fileType))
		if pageToken != "" {
			path += "&page_token=" + url.QueryEscape(pageToken)
		}

		var resp DocumentCommentsResponse
		if err := c.Get(path, &resp); err != nil {
			return nil, false, "", err
		}
		if err := resp.Err(); err != nil {
			return nil, false, "", err
		}
		return resp.Data.Items, resp.Data.HasMore, resp.Data.PageToken, nil
	})
}

// GetMediaTempDownloadURL gets a temporary download URL for a media file
// fileToken: the media token (e.g., image token from block)
// documentID: optional document ID for authentication (required for document images)
// Returns the temporary download URL (valid for 24 hours)
func (c *Client) GetMediaTempDownloadURL(fileToken, documentID string) (string, error) {
	path := fmt.Sprintf("/drive/v1/medias/batch_get_tmp_download_url?file_tokens=%s",
		url.QueryEscape(fileToken))

	// Add extra parameter with document ID if provided
	if documentID != "" {
		extra := fmt.Sprintf(`{"drive_route_token":"%s"}`, documentID)
		path += "&extra=" + url.QueryEscape(extra)
	}

	var resp MediaTempDownloadURLResponse
	if err := c.Get(path, &resp); err != nil {
		return "", err
	}

	if err := resp.Err(); err != nil {
		return "", err
	}

	if len(resp.Data.TmpDownloadURLs) == 0 {
		return "", fmt.Errorf("no download URL returned for token %s", fileToken)
	}

	return resp.Data.TmpDownloadURLs[0].TmpDownloadURL, nil
}

// DownloadMedia downloads a media file (image, attachment) from a document
// fileToken: the media token (e.g., image token from block)
// documentID: optional document ID for authentication (required for document images)
// Returns the file content as a ReadCloser and the content type
func (c *Client) DownloadMedia(fileToken, documentID string) (io.ReadCloser, string, error) {
	// Try direct download API first with extra parameter
	path := fmt.Sprintf("/drive/v1/medias/%s/download", url.PathEscape(fileToken))
	if documentID != "" {
		extra := fmt.Sprintf(`{"drive_route_token":"%s"}`, documentID)
		path += "?extra=" + url.QueryEscape(extra)
	}

	return c.Download(path)
}

// DownloadDriveFile downloads a file from Lark Drive
// fileToken: the file token from doc list or search
// Returns the file content as a ReadCloser and the content type
func (c *Client) DownloadDriveFile(fileToken string) (io.ReadCloser, string, error) {
	path := fmt.Sprintf("/drive/v1/files/%s/download", url.PathEscape(fileToken))
	// Try user token first, if that fails it might be a permission issue
	return c.Download(path)
}

// DownloadDriveFileWithTenant downloads a file using tenant token
// This may be needed for files shared with the bot
func (c *Client) DownloadDriveFileWithTenant(fileToken string) (io.ReadCloser, string, error) {
	path := fmt.Sprintf("/drive/v1/files/%s/download", url.PathEscape(fileToken))
	return c.DownloadWithTenantToken(path)
}

// UploadDriveFile uploads a file to Lark Drive using the upload_all API
// filePath: local file path
// parentToken: folder token to upload into (empty for root)
// parentType: "explorer" for Drive folder (default)
// Returns the file token of the uploaded file
func (c *Client) UploadDriveFile(filePath, parentToken, parentType string) (string, error) {
	if err := auth.EnsureValidToken(); err != nil {
		return "", err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return "", fmt.Errorf("failed to stat file: %w", err)
	}

	if stat.Size() > 20*1024*1024 {
		return "", fmt.Errorf("file size %d exceeds 20MB limit", stat.Size())
	}

	if parentType == "" {
		parentType = "explorer"
	}

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	_ = writer.WriteField("file_name", filepath.Base(filePath))
	_ = writer.WriteField("parent_type", parentType)
	if parentToken != "" {
		_ = writer.WriteField("parent_node", parentToken)
	}
	_ = writer.WriteField("size", strconv.FormatInt(stat.Size(), 10))

	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return "", fmt.Errorf("failed to create form file: %w", err)
	}
	if _, err = io.Copy(part, file); err != nil {
		return "", fmt.Errorf("failed to copy file data: %w", err)
	}
	writer.Close()

	url := baseURL + "/drive/v1/files/upload_all"
	req, err := http.NewRequest("POST", url, &body)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	token := auth.GetTokenStore().GetAccessToken()
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("upload request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	var uploadResp UploadDriveFileResponse
	if err := json.Unmarshal(respBody, &uploadResp); err != nil {
		return "", fmt.Errorf("failed to parse response: %w", err)
	}

	if err := uploadResp.Err(); err != nil {
		return "", err
	}

	return uploadResp.Data.FileToken, nil
}

// DeleteDocumentBlocks deletes blocks from a document by their IDs
// The Lark API deletes children by index range, so we need to resolve
// block IDs to their parent and index positions first.
// documentID: the document ID
// blockIDs: IDs of blocks to delete
func (c *Client) DeleteDocumentBlocks(documentID string, blockIDs []string) (int, error) {
	// Build a set for quick lookup
	toDelete := make(map[string]bool)
	for _, id := range blockIDs {
		toDelete[id] = true
	}

	// Get all blocks to find parent-child relationships
	blocks, err := c.GetDocumentBlocks(documentID)
	if err != nil {
		return 0, fmt.Errorf("failed to get blocks: %w", err)
	}

	// Build a map of block ID -> block for quick lookup
	blockMap := make(map[string]*DocumentBlock)
	for i := range blocks {
		blockMap[blocks[i].BlockID] = &blocks[i]
	}

	// Filter out blocks whose parent is also being deleted.
	// When a parent block is deleted, its children are deleted too,
	// so we only need to delete top-level blocks in the delete set.
	topLevelDelete := make(map[string]bool)
	for id := range toDelete {
		block, ok := blockMap[id]
		if !ok {
			continue
		}
		// Walk up the parent chain to check if any ancestor is also being deleted
		isDescendant := false
		parentID := block.ParentID
		for parentID != "" {
			if toDelete[parentID] {
				isDescendant = true
				break
			}
			parent, ok := blockMap[parentID]
			if !ok {
				break
			}
			parentID = parent.ParentID
		}
		if !isDescendant {
			topLevelDelete[id] = true
		}
	}

	// Group blocks to delete by parent, tracking their child indices
	// We need to delete from highest index to lowest to avoid shifting
	type deleteOp struct {
		parentID string
		index    int
	}
	var ops []deleteOp

	for _, block := range blocks {
		if block.Children == nil {
			continue
		}
		for idx, childID := range block.Children {
			if topLevelDelete[childID] {
				ops = append(ops, deleteOp{parentID: block.BlockID, index: idx})
			}
		}
	}

	if len(ops) == 0 {
		return 0, fmt.Errorf("none of the specified block IDs were found as children")
	}

	// Sort ops by index descending (delete from end first to preserve indices)
	for i := 0; i < len(ops)-1; i++ {
		for j := i + 1; j < len(ops); j++ {
			if ops[j].index > ops[i].index || (ops[j].index == ops[i].index && ops[j].parentID > ops[i].parentID) {
				ops[i], ops[j] = ops[j], ops[i]
			}
		}
	}

	// Delete one at a time from highest index to lowest
	var lastRevisionID int
	for _, op := range ops {
		path := fmt.Sprintf("/docx/v1/documents/%s/blocks/%s/children/batch_delete?document_revision_id=-1",
			url.PathEscape(documentID), url.PathEscape(op.parentID))

		req := DeleteBlocksRequest{
			StartIndex: op.index,
			EndIndex:   op.index + 1,
		}

		var resp DeleteBlocksResponse
		if err := c.DeleteWithBody(path, req, &resp); err != nil {
			return lastRevisionID, fmt.Errorf("failed to delete block at index %d: %w", op.index, err)
		}

		if err := resp.Err(); err != nil {
			return lastRevisionID, err
		}

		lastRevisionID = resp.Data.DocumentRevisionID
	}

	return lastRevisionID, nil
}

// UpdateDocumentBlock updates a single block's content in a document
// documentID: the document ID
// blockID: the block ID to update
// block: the updated block content
func (c *Client) UpdateDocumentBlock(documentID, blockID string, block DocumentBlock) (int, error) {
	path := fmt.Sprintf("/docx/v1/documents/%s/blocks/%s?document_revision_id=-1",
		url.PathEscape(documentID), url.PathEscape(blockID))

	updateReq := ConvertToUpdateRequest(block)

	var resp UpdateBlockResponse
	if err := c.Patch(path, updateReq, &resp); err != nil {
		return 0, err
	}

	if err := resp.Err(); err != nil {
		return 0, err
	}

	return resp.Data.DocumentRevisionID, nil
}

// ReplaceDocumentBlock atomically replaces a block: finds its position,
// deletes it, and inserts new blocks at the same index.
func (c *Client) ReplaceDocumentBlock(documentID, blockID string, newBlocks []DocumentBlock) ([]DocumentBlock, int, error) {
	// Get all blocks to find the target block's parent and index
	blocks, err := c.GetDocumentBlocks(documentID)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get blocks: %w", err)
	}

	// Find the block's parent and its index within the parent's children
	var parentID string
	var childIdx int
	found := false
	for _, block := range blocks {
		if block.Children == nil {
			continue
		}
		for idx, childID := range block.Children {
			if childID == blockID {
				parentID = block.BlockID
				childIdx = idx
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	if !found {
		return nil, 0, fmt.Errorf("block %s not found as a child of any block", blockID)
	}

	// Delete the block
	path := fmt.Sprintf("/docx/v1/documents/%s/blocks/%s/children/batch_delete?document_revision_id=-1",
		url.PathEscape(documentID), url.PathEscape(parentID))
	req := DeleteBlocksRequest{
		StartIndex: childIdx,
		EndIndex:   childIdx + 1,
	}
	var delResp DeleteBlocksResponse
	if err := c.DeleteWithBody(path, req, &delResp); err != nil {
		return nil, 0, fmt.Errorf("failed to delete block: %w", err)
	}
	if err := delResp.Err(); err != nil {
		return nil, 0, err
	}

	// Insert new blocks at the same index
	createdBlocks, revisionID, err := c.CreateDocumentBlocks(documentID, parentID, newBlocks, childIdx)
	if err != nil {
		return nil, delResp.Data.DocumentRevisionID, fmt.Errorf("deleted block but failed to insert replacement: %w", err)
	}

	return createdBlocks, revisionID, nil
}

// DeleteDriveFile moves a file to trash in Lark Drive
// fileToken: the file token
// docType: document type (e.g., "docx", "doc", "sheet", "bitable", "folder", "file")
func (c *Client) DeleteDriveFile(fileToken, docType string) error {
	path := fmt.Sprintf("/drive/v1/files/%s?type=%s",
		url.PathEscape(fileToken), url.QueryEscape(docType))

	var resp BaseResponse
	if err := c.Delete(path, &resp); err != nil {
		return err
	}

	if err := resp.Err(); err != nil {
		return err
	}

	return nil
}

// SearchDocuments searches for documents using the Lark Docs API
// query: search keyword (required)
// ownerIDs: optional filter by owner user IDs
// chatIDs: optional filter by chat IDs
// docTypes: optional filter by doc types (doc, sheet, slide, bitable, mindnote, file)
// Returns all matching documents (up to 200) and total count
func (c *Client) SearchDocuments(query string, ownerIDs, chatIDs, docTypes []string) ([]DocSearchEntity, int, error) {
	var allResults []DocSearchEntity
	offset := 0
	const pageSize = 50

	for {
		req := DocSearchRequest{
			SearchKey: query,
			Count:     pageSize,
			Offset:    offset,
		}
		if len(ownerIDs) > 0 {
			req.OwnerIDs = ownerIDs
		}
		if len(chatIDs) > 0 {
			req.ChatIDs = chatIDs
		}
		if len(docTypes) > 0 {
			req.DocsTypes = docTypes
		}

		var resp DocSearchResponse
		if err := c.Post("/suite/docs-api/search/object", req, &resp); err != nil {
			return nil, 0, err
		}

		if err := resp.Err(); err != nil {
			return nil, 0, err
		}

		allResults = append(allResults, resp.Data.DocsEntities...)

		// Check if we should continue (has_more and offset+count < 200)
		if !resp.Data.HasMore || offset+pageSize >= 200 {
			return allResults, resp.Data.Total, nil
		}

		offset += pageSize
	}
}


func (c *Client) GetDriveMeta(docToken, docType string) (*DriveMetaItem, error) {
	req := DriveMetaRequest{
		RequestDocs: []DriveMetaRequestDoc{{DocToken: docToken, DocType: docType}},
		WithURL:     true,
	}
	var resp DriveMetaResponse
	if err := c.Post("/drive/v1/metas/batch_query", req, &resp); err != nil {
		return nil, err
	}
	if err := resp.Err(); err != nil {
		return nil, err
	}
	if len(resp.Data.Metas) == 0 {
		return nil, fmt.Errorf("no metadata returned for token %s", docToken)
	}
	return &resp.Data.Metas[0], nil
}

func (c *Client) CreateFolder(name, parentToken string) (string, string, error) {
	req := CreateFolderRequest{
		Name:        name,
		FolderToken: parentToken,
	}
	var resp CreateFolderResponse
	if err := c.Post("/drive/v1/files/create_folder", req, &resp); err != nil {
		return "", "", err
	}
	if err := resp.Err(); err != nil {
		return "", "", err
	}
	return resp.Data.Token, resp.Data.URL, nil
}
