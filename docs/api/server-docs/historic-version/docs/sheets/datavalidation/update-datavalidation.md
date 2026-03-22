---
title: "Update Drop-down List Setting"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/historic-version/docs/sheets/datavalidation/update-datavalidation"
method: "PUT"
api_path: "https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/dataValidation/:sheetId"
service: "historic-version"
resource: "docs"
section: "Deprecated Version (Not Recommended)"
scopes:
  - "drive:drive"
  - "sheets:spreadsheet"
updated: "1755579370000"
---

#  ~~Update drop-down settings~~

**Note: This API will become unavailable starting September 20, 2025. Please update your app to use the latest API Update drop-down settings.**

Use this API to update drop-down properties based on spreadsheetToken, sheetId, and dataValidationId.

## Request
| HTTP URL | https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/dataValidation/:sheetId |
| --- | --- |
| HTTP Method | PUT |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `drive:drive` `sheets:spreadsheet` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `user_access_token` or `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about how to obtain and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
|Parameter|Type|Required|Description|
|--|-----|--|----|
|spreadsheetToken|string|Yes|Spreadsheet token. For more information about how to obtain the token, see Sheets Developer's Guide.| 
|sheetId|string|Yes|The unique identifier of the sheet

### Request body
|Parameter|Type|Required|Description|
|--|-----|--|----|
|ranges|array<string>|Yes|Update to Range Handling for Dropdown Lists. When updating a range, note that if the specified range does not currently have a dropdown list configured, the update will effectively create a new dropdown list for that range.|
|dataValidationType|string|Yes|The "list" of the drop-down|
|dataValidation|||Drop-down rule properties
|  ∟conditionValues|array|Yes|Drop-down options, must be strings that do not include commas (,). Specify up to 500 options of up to 100 characters.|
|  ∟options||No|Option properties| 
|    ∟multipleValues|bool|No|Single option: false (default), multiple options: true| 
|    ∟highlightValidData|bool|No|Whether to set color and capsule style, defaults to false when not specified|
|    ∟colors|array|No|When highlightValidData is true, color must be specified. Values have one-to-one correspondence with conditionValues. Use RGB16 hex format, for example: "#fffd00".| ### Request body example

```json
{
    "ranges":["BzY8T5!A1:A2", "BzY8T5!B1:B1"],
    "dataValidationType":"list",
    "dataValidation":{
        "conditionValues":["1", "2", "4","2"],
        "options":{
            "multipleValues":false,
            "highlightValidData":true,
            "colors":["#1FB6C1", "#1006C2", "#FB16C3","#FFB6C1"]
        }
    }
}
```
###  cURL Request example
  ```
  curl --location --request PUT 'https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/shtcngNygNfuqhxTBf588jwgWbJ/dataValidation/BzY8T5' \
--header 'Authorization: Bearer t-5be16bd570d0437444c40d5e6b5584109e61b0b1' \
--header 'Content-Type: application/json' \
--data-raw '{
    "ranges":["BzY8T5!A1:A2", "BzY8T5!B1:B1"],
    "dataValidationType":"list",
    "dataValidation":{
        "conditionValues":["1", "2", "4"],
        "options":{
            "multipleValues":false,
            "highlightValidData":true,
            "colors":["#1FB6C1", "#1006C2", "#FB16C3"]
        }
    }
}'
  ```
 ## Response

### Response body
  
|Parameter|Type|Required|Description|
|--|-----|--|----|
|code|int|Yes|Status code, 0 indicates success|
|msg|string|No|Status information|
|data||||
|   ∟spreadsheetToken|string|Yes|The token of the spreadsheet| 
|   ∟sheetId|string|Yes|ID of the sheet| 
|  ∟dataValidation|||| 
|    ∟dataValidationType|string|Yes|The "list" of the drop-down|
|    ∟conditionValues|array|Yes|Drop-down options|
|    ∟options||No|Option properties| 
|      ∟multipleValues|bool|No|Single option: false (default), multiple options: true| 
|      ∟highlightValidData|bool|No|Whether to set color and capsule style|
|      ∟colorValueMap|map|No|When highlightValidData is true, the colorValueMap key has a one-to-one correspondence with conditionValues. Value is the corresponding color parameter.| ### Response body example  

```json
{
    "code": 0,
    "data": {
        "dataValidation": {
            "conditionValues": [
                "1",
                "2",
                "4"
            ],
            "dataValidationType": "list",
            "options": {
                "colorValueMap": {
                    "1": "#1FB6C1",
                    "2": "#1006C2",
                    "4": "#FB16C3"
                },
                "highlightValidData": true,
                "multipleValues": false
            }
        },
        "sheetId": "yuNGtr",
        "spreadsheetToken": "shtbckBcolBlRfkcMVZbolMdADe"
    },
    "msg": "Success"
}
```  
  
### Error code

For details, refer to Server-side error codes.
