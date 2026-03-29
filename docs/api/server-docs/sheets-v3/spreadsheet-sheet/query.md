---
title: "Get sheet"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet/query"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/query"
service: "sheets-v3"
resource: "spreadsheet-sheet"
section: "Docs"
rate_limit: "100 per minute"
scopes:
  - "drive:drive"
  - "drive:drive:readonly"
  - "sheets:spreadsheet"
  - "sheets:spreadsheet:readonly"
updated: "1710130588000"
---

# Get sheet

This interface is used to get all sheets and their properties under the spreadsheet.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/query |
| HTTP Method | GET |
| Rate Limit | 100 per minute |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `drive:drive` `drive:drive:readonly` `sheets:spreadsheet` `sheets:spreadsheet:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `spreadsheet_token` | `string` | Spreadsheet token **Example value**: "xxxxxxxxxxxxxxxxxxx" | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `sheets` | `sheet[]` | List of sheets |
|     `sheet_id` | `string` | Worksheet id |
|     `title` | `string` | Worksheet title |
|     `index` | `int` | Worksheet index position |
|     `hidden` | `boolean` | Is the worksheet hidden? |
|     `grid_properties` | `grid_properties` | Cell properties |
|       `frozen_row_count` | `int` | Number of rows frozen |
|       `frozen_column_count` | `int` | Number of columns frozen |
|       `row_count` | `int` | Number of rows in the worksheet |
|       `column_count` | `int` | Number of columns in the worksheet |
|     `resource_type` | `string` | Worksheet type |
|     `merges` | `merge_range[]` | Merge cells |
|       `start_row_index` | `int` | Start line |
|       `end_row_index` | `int` | End line |
|       `start_column_index` | `int` | Start column |
|       `end_column_index` | `int` | End column | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "sheets": [
            {
                "sheet_id": "sxj5ws",
                "title": "title",
                "index": 0,
                "hidden": false,
                "grid_properties": {
                    "frozen_row_count": 0,
                    "frozen_column_count": 0,
                    "row_count": 200,
                    "column_count": 20
                },
                "resource_type": "sheet",
                "merges": [
                    {
                        "start_row_index": 0,
                        "end_row_index": 0,
                        "start_column_index": 0,
                        "end_column_index": 0
                    }
                ]
            }
        ]
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1310251 | Invalid Parameters | Reference error message in response body |
| 400 | 1310213 | Permission Fail | No permission |
| 400 | 1310249 | Spreadsheet Deleted | Restore the form and try again |
| 400 | 1310214 | SpreadSheet Not Found | Check the spreadsheet token. |
| 500 | 1315201 | Server Error | Internal service error. For more information, [contact support](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952) |
| 500 | 1315203 | Server Error | Internal service error. For more information, [contact support](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952). |
| 500 | 1315210 | Server Error | Internal service errors, please consult customer service for details |
