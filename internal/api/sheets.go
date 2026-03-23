package api

import (
	"fmt"
	"net/url"
)

// GetSpreadsheetSheets retrieves all sheets in a spreadsheet
// token: the spreadsheet token from the URL
func (c *Client) GetSpreadsheetSheets(token string) ([]Sheet, error) {
	path := fmt.Sprintf("/sheets/v3/spreadsheets/%s/sheets/query", url.PathEscape(token))

	var resp SpreadsheetSheetsResponse
	if err := c.Get(path, &resp); err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("API error %d: %s", resp.Code, resp.Msg)
	}

	return resp.Data.Sheets, nil
}

// GetSheetMetadata retrieves metadata for a single sheet
// token: the spreadsheet token
// sheetID: the sheet ID within the spreadsheet
func (c *Client) GetSheetMetadata(token, sheetID string) (*Sheet, error) {
	path := fmt.Sprintf("/sheets/v3/spreadsheets/%s/sheets/%s",
		url.PathEscape(token), url.PathEscape(sheetID))

	var resp SheetMetadataResponse
	if err := c.Get(path, &resp); err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("API error %d: %s", resp.Code, resp.Msg)
	}

	return resp.Data.Sheet, nil
}

// GetSheetData retrieves cell values from a sheet
// token: the spreadsheet token
// rangeStr: the range in format "sheetId!A1:Z100" or just "sheetId" for all data
func (c *Client) GetSheetData(token, rangeStr string) (*SheetValues, error) {
	path := fmt.Sprintf("/sheets/v2/spreadsheets/%s/values/%s",
		url.PathEscape(token), url.PathEscape(rangeStr))

	var resp SheetValuesResponse
	if err := c.Get(path, &resp); err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("API error %d: %s", resp.Code, resp.Msg)
	}

	return resp.Data, nil
}

// GetSheetDataAsText retrieves cell values as plain text strings.
// Uses valueRenderOption=ToString and dateTimeRenderOption=FormattedString
// so the API returns pre-rendered text instead of rich text arrays.
// Ideal for embedded sheet rendering in documents.
func (c *Client) GetSheetDataAsText(token, rangeStr string) (*SheetValues, error) {
	path := fmt.Sprintf("/sheets/v2/spreadsheets/%s/values/%s?valueRenderOption=ToString&dateTimeRenderOption=FormattedString",
		url.PathEscape(token), url.PathEscape(rangeStr))

	var resp SheetValuesResponse
	if err := c.Get(path, &resp); err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("API error %d: %s", resp.Code, resp.Msg)
	}

	return resp.Data, nil
}

// BatchGetSheetDataAsText retrieves cell values for multiple ranges in one call.
// Uses the values_batch_get endpoint with ToString rendering.
// ranges: list of ranges in format "sheetId!A1:Z200"
func (c *Client) BatchGetSheetDataAsText(token string, ranges []string) ([]ValueRange, error) {
	params := url.Values{}
	for _, r := range ranges {
		params.Add("ranges", r)
	}
	params.Set("valueRenderOption", "ToString")
	params.Set("dateTimeRenderOption", "FormattedString")

	path := fmt.Sprintf("/sheets/v2/spreadsheets/%s/values_batch_get?%s",
		url.PathEscape(token), params.Encode())

	var resp SheetBatchValuesResponse
	if err := c.Get(path, &resp); err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("API error %d: %s", resp.Code, resp.Msg)
	}

	return resp.Data.ValueRanges, nil
}

// SetSheetData writes cell values to a sheet
// token: the spreadsheet token
// sheetRange: the range in format "sheetId!A1:C3"
// values: 2D array of values to write
func (c *Client) SetSheetData(token string, sheetRange string, values [][]any) (*SetSheetValuesData, error) {
	path := fmt.Sprintf("/sheets/v2/spreadsheets/%s/values", url.PathEscape(token))

	req := SetSheetValuesRequest{
		ValueRange: ValueRange{
			Range:  sheetRange,
			Values: values,
		},
	}

	var resp SetSheetValuesResponse
	if err := c.Put(path, req, &resp); err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("API error %d: %s", resp.Code, resp.Msg)
	}

	return resp.Data, nil
}

// CreateSpreadsheet creates a new spreadsheet
// title: the spreadsheet title
// folderToken: optional parent folder token (empty = root)
func (c *Client) CreateSpreadsheet(title, folderToken string) (*SpreadsheetInfo, error) {
	req := CreateSpreadsheetRequest{
		Title:       title,
		FolderToken: folderToken,
	}

	var resp CreateSpreadsheetResponse
	if err := c.Post("/sheets/v3/spreadsheets", req, &resp); err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("API error %d: %s", resp.Code, resp.Msg)
	}

	return resp.Data.Spreadsheet, nil
}
