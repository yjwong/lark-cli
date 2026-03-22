---
title: "Set Drop-down List"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/datavalidation/set-dropdown"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/dataValidation"
section: "Docs"
scopes:
  - "drive:drive"
  - "sheets:spreadsheet"
updated: "1647001427000"
---

# Set drop-down

Use this API to set drop-down rules for cells based on spreadsheetToken, range, and drop-down properties. The range of a single operation cannot exceed 5,000 rows and 100 columns. When data already exists in the selected range, valid values are converted to options.

## Request
| HTTP URL | https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/dataValidation |
| --- | --- |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `drive:drive` `sheets:spreadsheet` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `user_access_token` or `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about how to obtain and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
|Parameter|Type|Required|Description|
|--|-----|--|----|
|spreadsheetToken|string|Yes|Spreadsheet token. For more information about how to obtain the token, see Sheets Developer's Guide.| ### Request body
|Parameter|Type|Required|Description|
|--|-----|--|----|
|range|string|Yes|Query range, includes the sheetId range and cell range. Four indexing methods are supported. For details, see  Sheets Developer's Guide.| 
|dataValidationType|string|Yes|The "list" of the drop-down|
|dataValidation|||Drop-down rule properties
|  ∟conditionValues|array|Yes|Drop-down options, must be strings that do not include commas (,). Specify up to 500 options of up to 100 characters.|
|  ∟options||No|Option properties| 
|    ∟multipleValues|bool|No|Single option: false (default), multiple options: true| 
|    ∟highlightValidData|bool|No|Whether to set color and capsule style, defaults to false when not specified|
|    ∟colors|array|No|When highlightValidData is true, color must be specified. Values have one-to-one correspondence with conditionValues. Use RGB16 hex format, for example: "#fffd00".| ### Request body example

```json
{
    "range":"yuNGtr!A2:A100",
    "dataValidationType":"list",
    "dataValidation":{
        "conditionValues":["2", "89", "3","2"],
        "options":{
            "multipleValues":true,
            "highlightValidData":true,
            "colors":["#1FB6C1", "#F006C2", "#FB16C3","#FFB6C1"]
        }
    }
}
```
###  cURL  Request example
  ```
curl --location --request POST 'https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/shtcngNygNfuqhxTBf588jwgWbJ/dataValidation' \
--header 'Authorization: Bearer t-5be16bd570d0437444c40d5e6b5584109e61b0b1' \
--header 'Content-Type: application/json' \
--data-raw '{
    "range":"BzY8T5!A2:A100",
    "dataValidationType":"list",
    "dataValidation":{
        "conditionValues":["2", "89", "3"],
        "options":{
            "multipleValues":true,
            "highlightValidData":true,
            "colors":["#1FB6C1", "#F006C2", "#FB16C3"]
        }
    }
}'
  ```
 ## Response

### Response body
  
|Parameter|Type|Required|Description|
|--|-----|--|----|
|code|int|Yes|Status code, 0 indicates success|
|msg|string|No|Status information| ### Response body example  

```json
{
    "code": 0,
    "msg": "Success"
}
```  
  
### Error code

For details, refer to Server-side error codes.
