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

// BatchCreateConditionFormats applies conditional formatting rules to sheet ranges
func (c *Client) BatchCreateConditionFormats(token string, formats []SheetConditionFormat) ([]ConditionFormatResponse, error) {
	path := fmt.Sprintf("/sheets/v2/spreadsheets/%s/condition_formats/batch_create", url.PathEscape(token))
	req := BatchCreateConditionFormatsRequest{SheetConditionFormats: formats}
	var resp BatchCreateConditionFormatsResponse
	if err := c.Post(path, req, &resp); err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("API error %d: %s", resp.Code, resp.Msg)
	}
	return resp.Data.Responses, nil
}
