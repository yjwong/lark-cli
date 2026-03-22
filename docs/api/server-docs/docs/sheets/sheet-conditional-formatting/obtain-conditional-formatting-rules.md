---
title: "Obtain Conditional Formatting Rules"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/conditionformat/condition-format-get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/condition_formats"
section: "Docs"
scopes:
  - "drive:drive"
  - "drive:drive:readonly"
  - "sheets:spreadsheet"
  - "sheets:spreadsheet:readonly"
updated: "1647001347000"
---

# Get conditional formatting

Use this API to get detailed conditional formatting based on a sheetId. Each operation can query up to 10 sheetIds.

## Request
| HTTP URL | https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/condition_formats |
| --- | --- |
| HTTP Method | GET |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `drive:drive` `drive:drive:readonly` `sheets:spreadsheet` `sheets:spreadsheet:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `user_access_token`  or  `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about how to obtain and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| spreadsheetToken | string | Spreadsheet token. For more information about how to obtain the token, see Sheets Developer's Guide | ### Query parameters
| Parameter             | Type   | Required | Description                                | 
| ---------------- | ------ | ---- | ------------------------------------------------------------ | 
| sheet_ids       |array|Yes   | Sheet ID, can be obtained through the Obtain sheet metadata API, separate multiple IDs with commas, for example: xxxID1,xxxID2 | 
###  cURL Request example
```
curl --location --request GET 'https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/shtcngNygNfuqhxTBf588jwgWbJ/condition_formats?sheet_ids=Q7PlXT' \
--header 'Authorization: Bearer t-e346617a4acfc3a11d4ed24dca0d0c0fc8e0067e' \
```
## Response  
### Response body
| Parameter                            |Type| Description                                       |
| ----------------------------- |-----| ---------------------------------------- |
|sheet_condition_formats|array|Spreadsheet conditional formatting information|
|   ∟sheet_id                      |string|Sheet ID                                 |
|   ∟condition_format              || Detailed information for a conditional formatting                              |
|     ∟cf_id                  |string| Conditional formatting ID                                  |
|     ∟ranges                 |array| Conditional formatting application scope, supported values: sheetId (for a whole sheet); sheetId!1:2 (for a whole row); sheetId!A:B (for a whole column); sheetId!A1:B2 (for a normal range); sheetId!A1:C (apply to the last row). The application scope cannot exceed the total rows and columns of a spreadsheet, and the sheetId must be consistent with the sheetId in the parameters. |
|     ∟rule_type              |string| Conditional formatting rule type, supports seven rule types: ***containsBlanks (contains blanks), notContainsBlanks (does not contain blanks), duplicateValues (duplicate values), uniqueValues (unique values), cellIs (limited value range), containsText (contains text content), timePeriod (dates)***                                  |
|     ∟attrs                  ||Specific property information for the corresponding rule_type, see Conditional Formatting User Guide                         for details|
|     ∟style                  || Conditional formatting style, only supports the following styles                           |
|        ∟font            || Font style                                     |
|            ∟bold     |bool| Bold                                       |
|           ∟italic   |bool| Italics                                       |
|        ∟text_decoration |int| Text decoration,  0: default, 1: underline, 2: strikethrough , 3: underline and strikethrough        |
|        ∟fore_color      |string| Text color                                     |
|        ∟back_color      |string| Background color                                     |
### Response body example

```json
{
    "code": 0,
    "msg": "Success"，
    "data": {
        "sheet_condition_formats": [
            {
                "condition_format": {
                    "cf_id": "r9sYuhgAl6",
                    "ranges": [
                        "uEnW3A!C4:C4"
                    ],
                    "rule_type": "timePeriod",
                    "attrs": [
                        {
                            "operator": "is",
                            "time_period": "today"
                        }
                    ],
                    "style": {
                        "back_color": "#d9f5d6",
                        "font": {
                            "bold": true,
                            "italic": false
                        },
                        "fore_color": "#faf1d1",
                        "text_decoration": 3
                    }
                },
                "sheet_id": "uEnW3A"
            }
        ]
    }
}
```
### Error code

For details, refer to Server-side error codes.
