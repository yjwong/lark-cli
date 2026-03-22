---
title: "Query Drop-down List in the Specific Range"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/datavalidation/query-datavalidation"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/dataValidation"
section: "Docs"
scopes:
  - "drive:drive"
  - "sheets:spreadsheet"
updated: "1755579401000"
---

# Query drop-down settings

Use this API to query drop-down settings information in a specified range based on spreadsheetToken and range. The range of a single query cannot exceed 5,000 rows and 100 columns.

## Request
| HTTP URL | https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/dataValidation |
| --- | --- |
| HTTP Method | GET |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `drive:drive` `sheets:spreadsheet` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `user_access_token`  or  `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about how to obtain and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
|Parameter|Type|Required|Description|
|--|-----|--|----|
|spreadsheetToken|string|Yes|Spreadsheet token. For more information about how to obtain the token, see  Sheets Developer's Guide.| ### Query parameters
|Parameter|Type|Required|Description|
|--|-----|--|----|
|range|string|Yes|Query range, includes the sheetId range and cell range. Four indexing methods are supported. For details, see Sheets Developer's Guide.|
|dataValidationType|string|Yes|The "list" of the drop-down| ###  cURL Request example
```
curl --location --request GET 'https://open.larksuite.com/open-apis/sheets/v2/spreadsheets/shtcngNygNfuqhxTBf588jwgWbJ/dataValidation?&range=BzY8T5!A2:A100' \
--header 'Authorization: Bearer t-5be16bd570d0437444c40d5e6b5584109e61b0b1' \
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
|  ∟revision|int|Yes|Version| 
|  ∟dataValidations|array|No|Drop-down array, blank when no value exists| 
|    ∟ranges|array<string>|Yes|The ranges associated with a dropdown list are aggregated by column. For example, in the subtable with ID 4d30c6, if cells A1, A2, A4, B1, and B2 belong to the same dropdown list, the corresponding ranges for that dropdown list are: ["4d30c6!A1:A2", "4d30c6!A4:A4", "4d30c6!B1:B2"].|
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
        "dataValidations": [
            {
                "conditionValues": [
                    "true",
                    "2",
                    "1",
                    "33.3333",
                    "ss"
                ],
                "dataValidationType": "list",
                "options": {
                    "colorValueMap": {
                        "1": "#b1e8fc",
                        "2": "#fed4a4",
                        "33.3333": "#f8e6ab",
                        "ss": "#a9efe6",
                        "true": "#bacefd"
                    },
                    "highlightValidData": true,
                    "multipleValues": true
                },
  				"ranges":[
        			"4d30c6!A1:A2",
        			"4d30c6!A4:A4",
        			"4d30c6!B1:B2"
        		]
            }
        ],
        "revision": 78,
        "sheetId": "4d30c6",
        "spreadsheetToken": "shtbckBcolBlRfkcMVZbolMdADe"
    },
    "msg": "Success"
}
```  
  
### Error code

For details, refer to Server-side error codes.
