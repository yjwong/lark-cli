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

	if resp.Code != 0 {
		return nil, fmt.Errorf("API error %d: %s", resp.Code, resp.Msg)
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

	if resp.Code != 0 {
		return "", fmt.Errorf("API error %d: %s", resp.Code, resp.Msg)
	}

	return resp.Data.Content, nil
}

// GetDocumentBlocks retrieves all blocks in a document with pagination
// documentID: the document ID (token from document URL)
func (c *Client) GetDocumentBlocks(documentID string) ([]DocumentBlock, error) {
	var allBlocks []DocumentBlock
	pageToken := ""

	for {
		path := fmt.Sprintf("/docx/v1/documents/%s/blocks?page_size=500",
			url.PathEscape(documentID))
		if pageToken != "" {
			path += "&page_token=" + url.QueryEscape(pageToken)
		}

		var resp DocumentBlocksResponse
		if err := c.Get(path, &resp); err != nil {
			return nil, err
		}

		if resp.Code != 0 {
			return nil, fmt.Errorf("API error %d: %s", resp.Code, resp.Msg)
		}

		allBlocks = append(allBlocks, resp.Data.Items...)

		if !resp.Data.HasMore || resp.Data.PageToken == "" {
			break
		}
		pageToken = resp.Data.PageToken
	}

	return allBlocks, nil
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

	if resp.Code != 0 {
		return nil, fmt.Errorf("API error %d: %s", resp.Code, resp.Msg)
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

	if resp.Code != 0 {
		return nil, 0, fmt.Errorf("API error %d: %s", resp.Code, resp.Msg)
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
	if resp.Code != 0 {
		return nil, false, "", fmt.Errorf("API error %d: %s", resp.Code, resp.Msg)
	}

	return resp.Data.Files, resp.Data.HasMore, resp.Data.NextPageToken, nil
}

// GetDocumentComments retrieves all comments for a document with pagination
// fileToken: the document token (same as document ID)
// fileType: document type (e.g., "docx", "doc", "sheet")
func (c *Client) GetDocumentComments(fileToken, fileType string) ([]DocumentComment, error) {
	var allComments []DocumentComment
	pageToken := ""

	for {
		path := fmt.Sprintf("/drive/v1/files/%s/comments?file_type=%s&page_size=100",
			url.PathEscape(fileToken), url.QueryEscape(fileType))
		if pageToken != "" {
			path += "&page_token=" + url.QueryEscape(pageToken)
		}

		var resp DocumentCommentsResponse
		if err := c.Get(path, &resp); err != nil {
			return nil, err
		}

		if resp.Code != 0 {
			return nil, fmt.Errorf("API error %d: %s", resp.Code, resp.Msg)
		}

		allComments = append(allComments, resp.Data.Items...)

		if !resp.Data.HasMore || resp.Data.PageToken == "" {
			break
		}
		pageToken = resp.Data.PageToken
	}

	return allComments, nil
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

	if resp.Code != 0 {
		return "", fmt.Errorf("API error %d: %s", resp.Code, resp.Msg)
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

	if uploadResp.Code != 0 {
		return "", fmt.Errorf("API error %d: %s", uploadResp.Code, uploadResp.Msg)
	}

	return uploadResp.Data.FileToken, nil
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

		if resp.Code != 0 {
			return nil, 0, fmt.Errorf("API error %d: %s", resp.Code, resp.Msg)
		}

		allResults = append(allResults, resp.Data.DocsEntities...)

		// Check if we should continue (has_more and offset+count < 200)
		if !resp.Data.HasMore || offset+pageSize >= 200 {
			return allResults, resp.Data.Total, nil
		}

		offset += pageSize
	}
}
