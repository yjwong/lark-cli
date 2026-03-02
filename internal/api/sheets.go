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

// SetSheetColumnWidths resizes columns in a sheet. widths is a map of 0-based column index to pixel width.
func (c *Client) SetSheetColumnWidths(token, sheetID string, widths map[int]int) error {
	path := fmt.Sprintf("/sheets/v2/spreadsheets/%s/dimension_range", url.PathEscape(token))

	for colIndex, pixelSize := range widths {
		req := SheetDimensionRangeRequest{
			Dimension: SheetDimension{
				SheetID:        sheetID,
				MajorDimension: "COLUMNS",
				StartIndex:     colIndex + 1, // API is 1-based
				EndIndex:       colIndex + 2,
			},
			DimensionProperties: SheetDimensionProperties{
				FixedSize: pixelSize,
			},
		}
		var resp SheetDimensionRangeResponse
		if err := c.Put(path, req, &resp); err != nil {
			return fmt.Errorf("column %d: %w", colIndex, err)
		}
		if resp.Code != 0 {
			return fmt.Errorf("column %d: API error %d: %s", colIndex, resp.Code, resp.Msg)
		}
	}
	return nil
}

// SetSheetStyleBold applies bold formatting to a range in a sheet.
func (c *Client) SetSheetStyleBold(token, sheetID, rangeSpec string) error {
	path := fmt.Sprintf("/sheets/v2/spreadsheets/%s/styles_batch_update", url.PathEscape(token))

	fullRange := sheetID + "!" + rangeSpec
	req := SheetStyleBatchUpdateRequest{
		Data: []SheetStyleItem{
			{
				Ranges: []string{fullRange},
				Style:  SheetStyle{Font: &SheetStyleFont{Bold: true}},
			},
		},
	}

	var resp SheetStyleBatchUpdateResponse
	if err := c.Put(path, req, &resp); err != nil {
		return err
	}
	if resp.Code != 0 {
		return fmt.Errorf("API error %d: %s", resp.Code, resp.Msg)
	}
	return nil
}

// SetSheetStyle applies a style to a range of cells
func (c *Client) SetSheetStyle(token, sheetID, rangeSpec string, style SheetStyle) error {
	path := fmt.Sprintf("/sheets/v2/spreadsheets/%s/styles_batch_update", url.PathEscape(token))
	fullRange := sheetID + "!" + rangeSpec
	req := SheetStyleBatchUpdateRequest{
		Data: []SheetStyleItem{
			{
				Ranges: []string{fullRange},
				Style:  style,
			},
		},
	}
	var resp SheetStyleBatchUpdateResponse
	if err := c.Put(path, req, &resp); err != nil {
		return err
	}
	if resp.Code != 0 {
		return fmt.Errorf("API error %d: %s", resp.Code, resp.Msg)
	}
	return nil
}

// AddSheetTab adds a new sheet tab to a spreadsheet.
func (c *Client) AddSheetTab(token, title string, index int) (*OutputSheetAddTab, error) {
	path := fmt.Sprintf("/sheets/v2/spreadsheets/%s/sheets_batch_update", url.PathEscape(token))

	item := AddSheetRequestItem{}
	item.AddSheet.Properties = AddSheetProperties{Title: title, Index: index}
	req := AddSheetBatchRequest{
		Requests: []AddSheetRequestItem{item},
	}

	var resp AddSheetBatchResponse
	if err := c.Post(path, req, &resp); err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("API error %d: %s", resp.Code, resp.Msg)
	}

	result := &OutputSheetAddTab{Success: true, Title: title, Index: index}
	if len(resp.Data.Replies) > 0 {
		props := resp.Data.Replies[0].AddSheet.Properties
		result.SheetID = props.SheetID
		result.Title = props.Title
		result.Index = props.Index
	}
	return result, nil
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
