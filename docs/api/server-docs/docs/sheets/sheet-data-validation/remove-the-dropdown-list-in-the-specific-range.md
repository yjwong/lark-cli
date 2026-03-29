---
title: "Remove the Drop-down List in the Specific Range"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/datavalidation/delete-datavalidation"
method: "DELETE"
api_path: "https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/dataValidation"
section: "Docs"
scopes:
  - "drive:drive"
  - "sheets:spreadsheet"
updated: "1755579405000"
---

# Delete drop-down settings

Use this API to remove drop-down settings but retain option text for cells in the selected data range based on spreadsheetToken and range. A single delete range cannot include more than 5,000 cells, and each request cannot specify more than 100 ranges.

## Request
| HTTP URL | https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/dataValidation |
| --- | --- |
| HTTP Method | DELETE |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `drive:drive` `sheets:spreadsheet` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `user_access_token`  or  `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about how to obtain and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
|Parameter|Type|Required|Description|
|--|-----|--|----|
|spreadsheetToken|string|Yes|Spreadsheet token. For more information about how to obtain the token, see Sheets Developer's Guide.| ### Request body
|Parameter|Type|Required|Description|
|--|-----|--|----|
|dataValidationRanges|array||Range array, each range is processed separately and cannot exceed 5,000 cells. The failure of one range does not impact other ranges. The return results give the execution results of each range.
|  ∟range|string|Yes|Query range, includes the sheetId range and cell range. Four indexing methods are supported. For details, see Sheets Developer's Guide.

### Request body example

```json
{
    "dataValidationRanges":[
        {
            "range": "4d30c6!A10:C10"
        }
    ]
}
```
### cURL Request example
  ```
  curl --location --request DELETE 'https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/shtcngNygNfuqhxTBf588jwgWbJ/dataValidation' \
--header 'Authorization: Bearer t-ce3540c5f02ac074535f1f14d64fa90fa49621c0' \
--header 'Content-Type: application/json' \
--data-raw '{
    "dataValidationRanges":[
        {
            "range": "BzY8T5!A1:A1"
        }
    ]
}'
  ```
 ## Response

### Response body
  
|Parameter|Type|Required|Description|
|--|-----|--|----|
|code|int|Yes|Status code, 0 indicates success|
|msg|string|No|Status information|
|data||||
|  ∟rangeResults|array|||
|    ∟range|string|Yes|The execution range, corresponds to range in the request parameters|
|    ∟msg|string|No|Result information|
|    ∟success|bool|Yes|Execution results|
|    ∟updatedCells|int|Yes|No. of cells affected| ### Response body example  

```json
{
    "code": 0,
    "data": {
        "rangeResults": [
            {
                "msg": "",
                "range": "4d30c6!A10:C10",
                "success": true,
                "updatedCells": 3
            }
        ]
    },
    "msg": "Success"
}
```  
  
### Error code

For details, refer to Server-side error codes.
