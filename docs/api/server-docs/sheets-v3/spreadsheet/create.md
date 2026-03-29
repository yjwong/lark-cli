---
title: "Create Spreadsheet"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet/create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/sheets/v3/spreadsheets"
service: "sheets-v3"
resource: "spreadsheet"
section: "Docs"
scopes:
  - "drive:drive"
  - "sheets:spreadsheet"
updated: "1698998306000"
---

# Create spreadsheet

Use this API to create an online spreadsheet under the specified directory.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/sheets/v3/spreadsheets |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `drive:drive` `sheets:spreadsheet` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `title` | `string` | No | Spreadsheet title **Example value**: "title" |
| `folder_token` | `string` | No | Folter token, see Overview to learn how to obtain it. **Example value**: "fldcnMsNb*****hIW9IjG1LVswg" | ### Request body example

{
    "title": "title",
    "folder_token": "fldcnMsNb*****hIW9IjG1LVswg"
}

```
curl --location --request POST 'https://open.larksuite.com/open-apis/sheets/v3/spreadsheets' \
--header 'Authorization: Bearer u-3iqkd6KWzRLzNdXfeuCMEb' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title":"title",
    "folder_token":"fldcnMsNb*****hIW9IjG1LVswg"
}'
```

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `spreadsheet` | `spreadsheet` | Spreadsheet |
|     `title` | `string` | Spreadsheet title |
|     `folder_token` | `string` | Folter token, see Overview to learn how to obtain it. |
|     `url` | `string` | Document URL |
|     `spreadsheet_token` | `string` | Spreadsheet token | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "spreadsheet": {
            "title": "title",
            "folder_token": "fldcnMsNb*****hIW9IjG1LVswg",
            "url": "https://example.larksuite.com/sheets/shtcnmBA*****yGehy8",
            "spreadsheet_token": "shtcnmBA*****yGehy8"
        }
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 500 | 1315203 | Server Error | Internal service error. For more information, [contact support](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952). |
| 400 | 1310229 | Wrong URL Param | Check the url parameter. |
| 400 | 1310204 | Wrong Request Body | Check the request body parameter. |
| 400 | 1310213 | Permission Fail | No permission |
| 400 | 1310235 | Retry Later | Please try again later. |
| 500 | 1315201 | Server Error | Internal service error. For more information, [contact support](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952) |
| 400 | 1310226 | Excess Limit | Exceeds the limit |
