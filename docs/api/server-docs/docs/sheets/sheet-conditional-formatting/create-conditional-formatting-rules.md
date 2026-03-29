---
title: "Create Conditional Formatting Rules"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/conditionformat/condition-format-set"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/condition_formats/batch_create"
section: "Docs"
scopes:
  - "drive:drive"
  - "sheets:spreadsheet"
updated: "1647001344000"
---

# Create conditional formatting

Use this API to create new conditional formatting. You can add up to 10 conditional formats at once. Each conditional formatting setting operation returns success or failed. Parameter verification information is returned for failed operations.

## Request
| HTTP URL | https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/condition_formats/batch_create |
| --- | --- |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `drive:drive` `sheets:spreadsheet` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `user_access_token`  or  `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about how to obtain and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| spreadsheetToken | string | spreadsheet , see [Sheets Developer's Guide]. | ### Request body
| Parameter                            |Type|Required| Description                                       |
| ----------------------------- |-----|----| ---------------------------------------- |
|sheet_condition_formats|array|Yes|Spreadsheet conditional formatting information|
|  ∟sheet_id                      |string|Yes|id                                  of  sheet|
|  ∟condition_format              ||Yes| Detailed information for a conditional formatting                              |
|    ∟ranges                 |array| | Conditional formatting application scope, supported values: sheetId (for a whole sheet); sheetId!1:2 (for a whole row); sheetId!A:B (for a whole column); sheetId!A1:B2 (for a normal range); sheetId!A1:C (apply to the last row). The application scope cannot exceed the total rows and columns of a spreadsheet, and the sheetId must be consistent with the sheetId in the parameters.| ∟rule_type              |string| Yes| Conditional formatting rule type, supports seven rule types: ***containsBlanks (contains blanks), notContainsBlanks (does not contain blanks), duplicateValues (duplicate values), uniqueValues (unique values), cellIs (limited value range), containsText (contains text content), timePeriod (dates)***                                 |
|    ∟attrs                  || Yes|Specific property information for the corresponding rule_type, see Conditional Formatting User Guide for details|
|      ∟operator|string| No|Operational approach|
|      ∟time_period|string| No|Time range|
|      ∟formula|array| No|Format|
|      ∟text|string| No|Text|
|    ∟style                  ||No| Conditional formatting style, only supports the following styles. Style parameters are all optional, but you cannot set an empty style .|
|       ∟font            || No| Font style                                     |
|         ∟bold     |bool| No| Bold                                       |
|         ∟italic   |bool| No| Italics                                       |
|       ∟text_decoration |int| No| Text decoration , 0: default, 1: underline, 2: strikethrough , 3: underline and strikethrough        |
|       ∟fore_color      |string| No| Text color                                     |
|       ∟back_color      |string| No| Background color                                     | ### Request body example

```json
{
    "sheet_condition_formats": [
        {
            "sheet_id": "40a7b0",
            "condition_format": {
                "ranges": [
                    "40a7b0!C3:C3"
                ],
                "rule_type": "timePeriod",
                "attrs": [
                    {
                        "operator": "is",
                        "time_period": "today", //Required when rule_type is timePeriod.
                        "formula": [], //Required when rule_type is cellIs.
                        "text": "" //Required when rule_type is containsText.
                    }
                ],
                "style": {
                    "font": {
                        "bold": true,
                        "italic": true
                    },
                    "fore_color": "#faf1d1",
                    "back_color": "#d9f5d6",
                    "text_decoration": 3
                }
            }
        }
    ]
}
```
###  cURL Request example
```
curl --location --request POST 'https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/shtcngNygNfuqhxTBf588jwgWbJ/condition_formats/batch_create' \
--header 'Authorization: Bearer t-e346617a4acfc3a11d4ed24dca0d0c0fc8e0067e' \
--header 'Content-Type: application/json' \
--data-raw '{
    "sheet_condition_formats": [
        {
            "sheet_id": "Q7PlXT",
            "condition_format": {
                "ranges": [
                    "Q7PlXT!C3:D9"
                ],
                "rule_type": "uniqueValues",
                "style": {
                    "font": {
                        "bold": true,
                        "italic": true
                    },
                    "fore_color": "#faf1d1",
                    "back_color": "#d9f5d6",
                    "text_decoration": 3
                }
            }
        }
    ]
}'
```
## Response
### Response body
| Parameter       | Type|Description                           |
| -------- |-----| ---------------------------- |
|responses|array|Response|
|   ∟sheet_id |string | Sheet ID                     |
|   ∟cf_id    |string| Successfully set conditional formatting id                  |
|   ∟res_code |int| Conditional formatting setting status code, 0 indicates success and a non-0 value indicates failure       |
|   ∟res_msg  |string| Status information returned for conditional formatting settings, empty indicates success and any content indicates the cause of failure | ### Response body example

```json
{
    "code": 0,
    "msg": "Success",
    "data": {
        "responses": [
            {
                "cf_id": "1gJelvenW9",
                "res_code": 0,
                "res_msg": "success",
                "sheet_id": "40a7b0"
            }
        ]
    }
}
```

### Error code

For details, refer to Server-side error codes.
