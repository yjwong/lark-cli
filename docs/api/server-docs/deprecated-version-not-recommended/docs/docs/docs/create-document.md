---
title: "Create Document"
url: "https://open.larksuite.com/document/ukTMukTMukTM/ugDM2YjL4AjN24COwYjN"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/doc/v2/create"
section: "Deprecated Version (Not Recommended)"
updated: "1737354744000"
---

# Create a document

> This API has been deprecated and apps are no longer permitted to call it. Continuing to call this api will result in an error code 95054.
> 
> If you want to create a document, see Create a document.

This API is used to create and initialize a document.

## Prerequisites

Before using this API, carefully read Overview and Prepare to integrate the Docs API to understand the calling rules and limitations of the Docs API. This will ensure your document data is not lost or corrupted.

For a definition of the document data structure, see Document Data Structure Overview. 

## Request
| HTTP URL | https://open.larksuite.com/open-apis/doc/v2/create |
| --- | --- |
| HTTP Method | POST |
| Required scopes Enable any scope from the list | View, comment, edit, and manage all files in My Space View, comment, edit, and manage Docs | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `user_access_token` or `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about how to access and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ::: note
For more information about how to call the AccessToken in the Docs API, see Get started with the Docs API.

### Request body
|Parameter|Type|Required|Description|Source|
|--|-----|--|----|----|
|FolderToken|string|No|Folder token. For more information about how to obtain the token, see How to get the token of docs resources. Defaults to the root directory when null. The tenant_access_token app scope only allows actions on directories created by the app. |post body JSON field | 
|Content|string|No|Passed in string that complies with the Document Data Structure. If it is null, a blank Docs is created. |post body JSON field| ### Curl request demo
```
curl -X POST 'https://open.larksuite.com/open-apis/doc/v2/create' \
	-H 'Authorization: Bearer u-s12okJw4R1VCZLWhk9Zyzg' \
	-H 'Content-Type: application/json; charset=utf-8' \
	-d '{"FolderToken":"","Content":"{\"title\":{\"elements\":[{\"type\":\"textRun\",\"textRun\":{\"text\":\"Hello Docs\",\"style\":{}}}]},\"body\":{\"blocks\":[{\"type\":\"paragraph\",\"paragraph\":{\"elements\":[{\"type\":\"textRun\",\"textRun\":{\"text\":\"Congratulations, you have successfully created an online Docs through the Open API. This Docs describes most supports API actions for Docs and provides sample code to help you understand how the Docs Open API works.\",\"style\":{}}}]}},{\"type\":\"paragraph\",\"paragraph\":{\"elements\":[],\"style\":{}}},{\"type\":\"paragraph\",\"paragraph\":{\"elements\":[{\"type\":\"textRun\",\"textRun\":{\"text\":\"You can use the URL to get the document token. You must provide this token when performing any API action on this Docs. For example:\",\"style\":{}}}]}},{\"type\":\"code\",\"code\":{\"language\":\"Plain Text\",\"body\":{\"blocks\":[{\"type\":\"paragraph\",\"paragraph\":{\"elements\":[{\"type\":\"textRun\",\"textRun\":{\"text\":\"https://xxx.larksuite.com/docs/\",\"style\":{}}},{\"type\":\"textRun\",\"textRun\":{\"text\":\"doccnmPiiJWfj1uk8bV7N5SzXaa\",\"style\":{\"backColor\":{\"red\":255,\"green\":246,\"blue\":122,\"alpha\":0.8}}}}],\"style\":{}}}]}}},{\"type\":\"paragraph\",\"paragraph\":{\"elements\":[],\"style\":{}}},{\"type\":\"paragraph\",\"paragraph\":{\"elements\":[{\"type\":\"textRun\",\"textRun\":{\"text\":\"As you can see from the above example, each line in the Docs is one block.\",\"style\":{}}}],\"style\":{\"headingLevel\":1}}},{\"type\":\"paragraph\",\"paragraph\":{\"elements\":[{\"type\":\"textRun\",\"textRun\":{\"text\":\"When a document has multiple block types, you must specify the block type for the current line when inserting content Block. The block type for text is "Paragraph".\",\"style\":{}}}],\"style\":{}}},{\"type\":\"paragraph\",\"paragraph\":{\"elements\":[{\"type\":\"textRun\",\"textRun\":{\"text\":\"A single row of text can be composed of multiple elements. For example, this line is  composed of the formula \",\"style\":{}}},{\"type\":\"equation\",\"equation\":{\"equation\":\"E=mc^2\"}},{\"type\":\"textRun\",\"textRun\":{\"text\":\" and code block in line \",\"style\":{}}},{\"type\":\"textRun\",\"textRun\":{\"text\":\"\",\"style\":{\"codeInline\":true}}},{\"type\":\"textRun\",\"textRun\":{\"text\":\". If some text contains a\",\"style\":{}}},{\"type\":\"textRun\",\"textRun\":{\"text\":\"unique style\",\"style\":{\"backColor\":{\"red\":255,\"green\":246,\"blue\":122,\"alpha\":0.8}}}},{\"type\":\"textRun\",\"textRun\":{\"text\":\", it is viewed as a unique element.\",\"style\":{}}}],\"style\":{}}},{\"type\":\"paragraph\",\"paragraph\":{\"elements\":[],\"style\":{}}},{\"type\":\"paragraph\",\"paragraph\":{\"elements\":[{\"type\":\"textRun\",\"textRun\":{\"text\":\"You can use APIs to insert various types of content.\",\"style\":{}}}],\"style\":{\"headingLevel\":1}}},{\"type\":\"paragraph\",\"paragraph\":{\"elements\":[{\"type\":\"textRun\",\"textRun\":{\"text\":\"You can insert numbered lists\",\"style\":{}}}],\"style\":{\"list\":{\"type\":\"number\",\"indentLevel\":1,\"number\":1}}}},{\"type\":\"paragraph\",\"paragraph\":{\"elements\":[{\"type\":\"textRun\",\"textRun\":{\"text\":\"You can specify your own numbered list numbers\",\"style\":{}}}],\"style\":{\"list\":{\"type\":\"number\",\"indentLevel\":1,\"number\":2}}}},{\"type\":\"paragraph\",\"paragraph\":{\"elements\":[{\"type\":\"textRun\",\"textRun\":{\"text\":\"You can also insert bulleted lists\",\"style\":{}}}],\"style\":{\"list\":{\"type\":\"bullet\",\"indentLevel\":1}}}},{\"type\":\"paragraph\",\"paragraph\":{\"elements\":[{\"type\":\"textRun\",\"textRun\":{\"text\":\"Insert more bulleted lists\",\"style\":{}}}],\"style\":{\"list\":{\"type\":\"bullet\",\"indentLevel\":1}}}},{\"type\":\"paragraph\",\"paragraph\":{\"elements\":[{\"type\":\"textRun\",\"textRun\":{\"text\":\"You can insert task lists and date reminders \",\"style\":{}}},{\"type\":\"reminder\",\"reminder\":{\"isWholeDay\":false,\"timestamp\":1609497000,\"shouldNotify\":true}}],\"style\":{\"list\":{\"type\":\"checkBox\",\"indentLevel\":1},\"todoUUID\":\"126f1471-f013-4ed1-9820-60f24da80677\"}}},{\"type\":\"paragraph\",\"paragraph\":{\"elements\":[{\"type\":\"textRun\",\"textRun\":{\"text\":\"Or insert referenced content\",\"style\":{}}}],\"style\":{\"quote\":true}}},{\"type\":\"paragraph\",\"paragraph\":{\"elements\":[{\"type\":\"textRun\",\"textRun\":{\"text\":\"Change document alignment method, for example: \",\"style\":{}}},{\"type\":\"textRun\",\"textRun\":{\"text\":\"Align right\",\"style\":{}}}],\"style\":{\"align\":\"right\"}}},{\"type\":\"paragraph\",\"paragraph\":{\"elements\":[{\"type\":\"textRun\",\"textRun\":{\"text\":\"In addition, you can flexibly apply various font styles to make the Docs more expressive. For example: \",\"style\":{}}}],\"style\":{\"align\":\"left\"}}},{\"type\":\"paragraph\",\"paragraph\":{\"elements\":[{\"type\":\"textRun\",\"textRun\":{\"text\":\"Bold \",\"style\":{\"bold\":true}}},{\"type\":\"textRun\",\"textRun\":{\"text\":\"Italics \",\"style\":{\"italic\":true}}},{\"type\":\"textRun\",\"textRun\":{\"text\":\"Underline \",\"style\":{\"underLine\":true}}},{\"type\":\"textRun\",\"textRun\":{\"text\":\"Strikethrough\",\"style\":{\"strikeThrough\":true}}},{\"type\":\"textRun\",\"textRun\":{\"text\":\" \",\"style\":{}}},{\"type\":\"textRun\",\"textRun\":{\"text\":\"inline code\",\"style\":{\"codeInline\":true}}},{\"type\":\"textRun\",\"textRun\":{\"text\":\" \",\"style\":{}}},{\"type\":\"textRun\",\"textRun\":{\"text\":\"Callout\",\"style\":{\"backColor\":{\"red\":255,\"green\":246,\"blue\":122,\"alpha\":0.8}}}},{\"type\":\"textRun\",\"textRun\":{\"text\":\"Color\",\"style\":{\"backColor\":{\"red\":255,\"green\":246,\"blue\":122,\"alpha\":0.8}}}},{\"type\":\"textRun\",\"textRun\":{\"text\":\" \",\"style\":{}}},{\"type\":\"textRun\",\"textRun\":{\"text\":\"Hyperlink\",\"style\":{\"link\":{\"url\":\"https%3A%2F%2Fopen.larksuite.com%2Fdocument%2FukTMukTMukTM%2FuUDN04SN0QjL1QDN%2Fdocs-doc-overview\"}}}}]}},{\"type\":\"paragraph\",\"paragraph\":{\"elements\":[{\"type\":\"textRun\",\"textRun\":{\"text\":\"Or a strikethrough\",\"style\":{}}}],\"style\":{\"align\":\"center\"}}},{\"type\":\"horizontalLine\",\"horizontalLine\":{}},{\"type\":\"paragraph\",\"paragraph\":{\"elements\":[],\"style\":{}}},{\"type\":\"paragraph\",\"paragraph\":{\"elements\":[{\"type\":\"textRun\",\"textRun\":{\"text\":\"You can also use the APIs to insert various tables. For example: \",\"style\":{}}}],\"style\":{\"headingLevel\":1}}},{\"type\":\"paragraph\",\"paragraph\":{\"elements\":[{\"type\":\"textRun\",\"textRun\":{\"text\":\"Normal table\",\"style\":{}}}],\"style\":{\"headingLevel\":2}}},{\"type\":\"table\",\"table\":{\"rowSize\":3,\"columnSize\":3,\"tableRows\":[{\"rowIndex\":0,\"tableCells\":[{\"columnIndex\":0,\"body\":{\"blocks\":[{\"type\":\"paragraph\",\"paragraph\":{\"elements\":[{\"type\":\"textRun\",\"textRun\":{\"text\":\"You can use tables to set the layout\",\"style\":{}}}],\"style\":{\"align\":\"center\"}}},{\"type\":\"paragraph\",\"paragraph\":{\"elements\":[{\"type\":\"textRun\",\"textRun\":{\"text\":\"Each row in a table is a block\",\"style\":{}}}],\"style\":{\"align\":\"center\"}}},{\"type\":\"paragraph\",\"paragraph\":{\"elements\":[{\"type\":\"textRun\",\"textRun\":{\"text\":\"In addition, each cell is identified by a unique zoneId\",\"style\":{}}}],\"style\":{\"align\":\"center\"}}}]}},{\"columnIndex\":1,\"body\":{\"blocks\":null}},{\"columnIndex\":2,\"body\":{\"blocks\":null}}]},{\"rowIndex\":1,\"tableCells\":[{\"columnIndex\":0,\"body\":{\"blocks\":[{\"type\":\"paragraph\",\"paragraph\":{\"elements\":[{\"type\":\"textRun\",\"textRun\":{\"text\":\"In\",\"style\":{\"codeInline\":true}}}],\"style\":{\"align\":\"center\"}}}]}},{\"columnIndex\":1,\"body\":{\"blocks\":[{\"type\":\"paragraph\",\"paragraph\":{\"elements\":[{\"type\":\"textRun\",\"textRun\":{\"text\":\"\",\"style\":{\"codeInline\":true}}}],\"style\":{\"align\":\"center\"}}}]}},{\"columnIndex\":2,\"body\":{\"blocks\":[{\"type\":\"paragraph\",\"paragraph\":{\"elements\":[{\"type\":\"textRun\",\"textRun\":{\"text\":\"Lark\",\"style\":{\"codeInline\":true}}}],\"style\":{\"align\":\"center\"}}}]}}]},{\"rowIndex\":2,\"tableCells\":[{\"columnIndex\":0,\"body\":{\"blocks\":[{\"type\":\"paragraph\",\"paragraph\":{\"elements\":[{\"type\":\"textRun\",\"textRun\":{\"text\":\" \",\"style\":{\"codeInline\":true}}}],\"style\":{\"align\":\"center\"}}}]}},{\"columnIndex\":1,\"body\":{\"blocks\":[{\"type\":\"paragraph\",\"paragraph\":{\"elements\":[{\"type\":\"textRun\",\"textRun\":{\"text\":\"Efficient\",\"style\":{\"codeInline\":true}}}],\"style\":{\"align\":\"center\"}}}]}},{\"columnIndex\":2,\"body\":{\"blocks\":[{\"type\":\"paragraph\",\"paragraph\":{\"elements\":[{\"type\":\"textRun\",\"textRun\":{\"text\":\"\",\"style\":{\"codeInline\":true}}}],\"style\":{\"align\":\"center\"}}}]}}]}],\"tableStyle\":{\"tableColumnProperties\":[{\"width\":100},{\"width\":100},{\"width\":100}]},\"mergedCells\":[{\"mergedCellId\":\"p47dtbcp\",\"rowStartIndex\":0,\"columnStartIndex\":0,\"rowEndIndex\":1,\"columnEndIndex\":3}]}},{\"type\":\"paragraph\",\"paragraph\":{\"elements\":[{\"type\":\"textRun\",\"textRun\":{\"text\":\"Datasheet\",\"style\":{}}}],\"style\":{\"headingLevel\":2}}},{\"type\":\"bitable\",\"bitable\":{\"viewType\":\"grid\"}},{\"type\":\"paragraph\",\"paragraph\":{\"elements\":[{\"type\":\"textRun\",\"textRun\":{\"text\":\"Dashboard\",\"style\":{}}}],\"style\":{\"headingLevel\":2}}},{\"type\":\"bitable\",\"bitable\":{\"viewType\":\"kanban\"}},{\"type\":\"paragraph\",\"paragraph\":{\"elements\":[]}},{\"type\":\"paragraph\",\"paragraph\":{\"elements\":[{\"type\":\"textRun\",\"textRun\":{\"text\":\"Sheets\",\"style\":{}}}],\"style\":{\"headingLevel\":2}}},{\"type\":\"sheet\",\"sheet\":{\"rowSize\":3,\"columnSize\":4}},{\"type\":\"paragraph\",\"paragraph\":{\"elements\":[],\"style\":{}}},{\"type\":\"paragraph\",\"paragraph\":{\"elements\":[{\"type\":\"textRun\",\"textRun\":{\"text\":\"I’m a webpage\",\"style\":{}}}],\"style\":{\"headingLevel\":2}}},{\"type\":\"embeddedPage\",\"embeddedPage\":{\"type\":\"xigua\",\"url\":\"https%3A%2F%2Fwww.ixigua.com%2Fiframe%2F6763174234282787341%3Fautoplay%3D0\",\"width\":637,\"height\":358}},{\"type\":\"paragraph\",\"paragraph\":{\"elements\":[]}}]}}"}'
```

## Response

### Response body  
|Parameter|Description|
|--|--|
|objToken|The token of the new Docs|
|url|The access link of the new Docs| ### Response body example   
```json
{
	"code": 0,
	"data": {
		"objToken": "doccn2EwbGEdcmryBvKzk0loaCd",
		"url": "https://xxx.larksuite.com/docs/doccn2EwbGEdcmryBvKzk0loaCd"
	},
	"msg": "Success"
}
```
### Error code
| Error code | Description | Troubleshooting suggestions |
| --- | --- | --- |
| 91401 | PARAMERR | Check parameter validity |
| 91402 | NOTEXIST | Check if the token is valid |
| 91403 | FORBIDDEN | Check for document read permissions |
| 91404 | LOGIN_REQUIRED | Need to log in |
| 95001 | internal error | Internal error, please try again later. |
| 95003 | internal error | Internal error, please try again later. |
| 95005 | internal error | Internal error, please try again later. |
| 95006 | Failed | Check if the token is valid |
| 95007 | Failed | Deleted file cannot get document meta information |
| 95008 | FORBIDDEN | Check user's permission for doc and folder |
| 95009 | Failed | Check for document read permissions. Add permissions |
| 95010 | internal error | Internal error, please try again later. |
| 95011 | internal error | Internal error, please try again later. |
| 95013 | Failed | Invalid folderToken or no permissions on directory or the same directory cannot be created concurrently or the maximum number of files in the directory exceeds 1500 |
| 95017 | Specific error message | Check if the revison is correct |
| 95018 | Specific error message | See specific error messages for details |
| 95020 | Specific error message | See specific error messages for details |
| 95023 | revision too old | The Revision is too old, Please use the latest version number |
| 95024 | Failed | Check parameter validity |
| 95025 | Failed | Check if the request is legal json |
| 95029 | folder locked | Concurrent creation of documents in the same folder is not supported, please call this api serially in the same folder |
| 95030 | folder size exceeded limit | The number of documents in the folder exceeds the limit |
| 95050 | Specific error message | See specific error messages for details |
| 95051 | Specific error message | See specific error messages for details |
| 95054 | this api is deprecated | This API has been deprecated. If you want to create a document, see Create a document | For details, refer to Server-side error codes.
