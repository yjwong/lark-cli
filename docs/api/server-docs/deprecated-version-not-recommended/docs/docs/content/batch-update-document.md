---
title: "Batch Update Document"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uYDM2YjL2AjN24iNwYjN"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/doc/v2/:docToken/batch_update"
section: "Deprecated Version (Not Recommended)"
updated: "1695369925000"
---

# Edit a document
> **Note:** Before using this API, carefully read Overview and Prepare to integrate the Docs API to understand the calling rules and limitations of the Docs API. This will ensure your document data is not lost or corrupted.
> For a definition of the document data structure, see: Document Data Structure Overview. 
> **Note:** This API does not support the Upgraded Docs. If you need to edit the content of the Upgraded Docs, please call the related API of the Upgraded Docs:
> - Create Block
> - Update Block
> - Batch Update Blocks
> - Delete Blocks
This API is used to batch edit and update document content, including updating headings, deleting a specified range, and inserting content.

## Request  
| HTTP URL | https://open.larksuite.com/open-apis/doc/v2/:docToken/batch_update |
| --- | --- |
| HTTP Method | POST |
| Required scopes Enable any scope from the list | View, comment, edit, and manage all files in My Space View, comment, edit, and manage Docs | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `user_access_token` or `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about how to access and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ::: note
For more information about how to call the AccessToken in the Docs API, see Get started with the Docs API.

### Request body    
| Parameter | Type | Required | Description | 
| ------ | ------------ | ---- | -----|
|docToken|string|Yes|File token. For more information about how to obtain the token, see How to get the token of docs resources.|
| Revision      | int          | Yes   | The specified version of the document. When first created, the document version is 0. It must be >=0 in the Obtain method, post body JSON field. |
| Requests      | list | Yes   | post body JSON, OperationRequest type serialized string array                         | #### OperationRequest

```
{
    "requestType": string,
    "updateTitleRequest": {object(UpdateTitleRequest)},
    "deleteContentRangeRequest": {object(DeleteContentRangeRequest)},
    "insertBlocksRequest": {object(InsertBlocksRequest)},
    "insertParagraphElementsRequest": {object(InsertParagraphElementsRequest)},
    "insertTableRowRequest": {object(InsertTableRowRequest)},
    "insertTableColumnRequest": {object(InsertTableColumnRequest)},
    "deleteTableRowsRequest": {object(DeleteTableRowsRequest)},
    "deleteTableColumnsRequest": {object(DeleteTableColumnsRequest)},
    "updateTableColumnPropertiesRequest": {object(UpdateTableColumnPropertiesRequest)},
    "mergeTableCellsRequest": {object(MergeTableCellsRequest)},
    "unmergeTableCellsRequest": {object(UnmergeTableCellsRequest)}
    "replaceAllTextRequest": {object{ReplaceAllTextRequest}}
    "updateParagraphStyleRequest": {object{UpdateParagraphStyleRequest}}
}
```

## Field description 
| Field | Type | Description |
| --- | --- | --- |
| requestType | string | **OperationRequest:** Pass in the request field (first letter must be capitalized) + 'Type', for example: 'UpdateTitleRequestType'. **Notes (an error will be reported if not passed in the proper format)** First letter must be capitalized Add 'Type' |
| updateTitleRequest | object | Update heading |
| deleteContentRangeRequest | object | Delete range |
| insertBlocksRequest | object | Insert block |
| insertParagraphElementsRequest | object | Insert element in row |
| insertTableRowRequest | object | Add one row to formatted sheet |
| insertTableColumnRequest | object | Add one column to formatted sheet |
| deleteTableRowsRequest | object | Delete multiple rows from formatted sheet |
| deleteTableColumnsRequest | object | Delete multiple columns from formatted sheet |
| updateTableColumnPropertiesRequest | object | Modify column width of formatted sheet |
| mergeTableCellsRequest | object | Merge cells in formatted sheet |
| unmergeTableCellsRequest | object | Split cells in formatted sheet |
| ReplaceAllTextRequest | object | Find and replace text content |
| UpdateParagraphStyleRequest | object | Update paragraph style | ### InsertLocation

```json
{
    "zoneId": string,
    "index": int,
    "startOfZone": bool,
    "endOfZone": bool
}
```

**Field description**:  

| Field   | Type   | Description                                                         |
| ------ | ------ | ------------------------------------------------------------ |
| zoneId | string | Indicates the zone to edit, "0" indicates body and an ID indicates a cell, for example: "xr1m4jw7egd9nefz1s0mdsetenl5fbe3lygxc1azupv81i5t2rjmosw5ta0esuwtn8ksya" |
| index  | int    | Character index, used by headings, starts from 0                         |
| startOfZone  | bool    | Start location of inserted zone, text header (not deemed as heading), or the start location of the cell. When true, the index input parameter is invalid. This is used for InsertBlocksRequestType convenient actions.                          | | endOfZone  | bool    | End location of inserted zone, text footer, or the end location of the cell. When true, the index input parameter is invalid. This is used for InsertBlocksRequestType convenient actions.                            | ### Range

Range deletion structure, the range to delete StartIndex, EndIndex)

```json
{
   "zoneId": string,
   "startIndex": int,
   "endIndex": int
}
```

**Field description**:  

| Field            | Type        | Description                                                         |
| ---------- | ------ | ------------------------------------------------------------ |
| zoneId | string | Indicates the zone to edit, "0" indicates body and an ID indicates a cell, for example: "xr1m4jw7egd9nefz1s0mdsetenl5fbe3lygxc1azupv81i5t2rjmosw5ta0esuwtn8ksya" |
| startIndex | int    | Deletion start location                                                 |
| endIndex   | int    | Deletion end location                                                 | ### UpdateTitleRequest

```json
{
    "payload": string
}
```

**Field description**:  

| Field            | Type        | Description                                                         |
| ------- | ------ | ------------------------------------------------------------ |
| payload | string | Title structure serialized string, see [Document Data Structure Overview | **Request example**
   :::note
This example changes the document heading to Updated Document Title.
  :::
  ```
{
    "Revision": 1,
    "Requests": [
        "{\"requestType\":\"UpdateTitleRequestType\",\"updateTitleRequest\":{\"payload\":\"{\\\"elements\\\":[{\\\"type\\\":\\\"textRun\\\",\\\"textRun\\\":{\\\"text\\\":\\\"Updated Document Title\\\",\\\"style\\\":{}}}],\\\"style\\\":{}}\"}}"
    ]
}
  ```

### DeleteContentRangeRequest

```json
{
    "deleteRange": {object(Range)}
}
```

**Field description**:  

| Field            | Type        | Description                                                         |
| ----------- | ------ | -------- |
| deleteRange | object | Range to delete | **Request example**
  :::note
  This example deletes the content of the selected range in the document.
  :::
  ```
  {
    "Revision": 1,
    "Requests": [
        "{\"requestType\":\"DeleteContentRangeRequestType\",\"deleteContentRangeRequest\":{\"deleteRange\":{\"zoneId\":\"0\",\"startIndex\":625,\"endIndex\":631}}}"
    ]
}    
  ```

### InsertBlocksRequest
It supports the return of lineIds in Paragraph TextRun.
  
```json
{
    "payload": string,
    "location": {object(InsertLocation)}
}
```

**Field description**:  

| Field            | Type        | Description                                                         |
| -------- | ------ | ------------------------------------------------------------ |
| payload | string | Body structure serialized string, see Document Data Structure Overview |
| location | object | Insertion location                                                     | **Request example**  
  :::note
This example inserts a line of text, a 3x3 normal sheet, and a 3x3 sheet in the document.
  :::
  ```
  {
    "Revision": 1,
    "Requests": [
        "{\"requestType\":\"InsertBlocksRequestType\",\"insertBlocksRequest\":{\"payload\":\"{\\\"blocks\\\":[{\\\"type\\\":\\\"paragraph\\\",\\\"paragraph\\\":{\\\"elements\\\":[{\\\"type\\\":\\\"textRun\\\",\\\"textRun\\\":{\\\"text\\\":\\\"Docs API Sample Content\\\",\\\"style\\\":{}}}],\\\"style\\\":{}}},{\\\"type\\\":\\\"table\\\",\\\"table\\\":{\\\"rowSize\\\":3, \\\"columnSize\\\":3}},{\\\"type\\\":\\\"sheet\\\",\\\"sheet\\\":{\\\"rowSize\\\":3, \\\"columnSize\\\":3}}]}\",\"location\":{\"zoneId\":\"0\",\"index\":0, \"endOfZone\": true}}}"
    ]
}
  ```

### InsertParagraphElementsRequest

```json
{
    "payload": string,
    "location": {object(InsertLocation)}
}
```

**Field description**:  

| Field            | Type        | Description                                                         |
| -------- | ------ | ------------------------------------------------------------ |
| payload  | string |The passed in string that complies with the Document Data Structure |
| location | object | Insertion location                                                     | **Request example**

> **Note:** This example inserts the code `Docs API Sample Content` in a single row in a file.
  ```
  {
    "Revision": 1,
    "Requests": [
        "{\"requestType\":\"InsertParagraphElementsRequestType\",\"insertParagraphElementsRequest\":{\"payload\":\"{\\\"elements\\\":[{\\\"type\\\":\\\"textRun\\\",\\\"textRun\\\":{\\\"text\\\":\\\"Docs API Sample Content\\\",\\\"style\\\":{\\\"codeInline\\\":true}}}],\\\"style\\\":{}}\",\"location\":{\"zoneId\":\"0\",\"index\":653}}}"
    ]
}
  ```
### InsertTableRowRequest

```json
{
    "tableId": string,
    "rowIndex": int
}
```

**Field description**:  

| Field            | Type        | Description                                                         |
| -------- | ------ | ------------------------------------------------------------ |
| tableId  | string | Table ID, see the sheet structure definition in Document Data Structure Overview |
| rowIndex | object | Insert row index, starts from 0. Enter 0 to insert before the first row.                 | **Request example**
  ::: note
  This example inserts a row before the first row in a sheet.
  :::
  ```
  {
        "Revision": 167,
        "Requests": [
        "{\"requestType\":\"InsertTableRowRequestType\",\"insertTableRowRequest\":{\"tableId\":\"rslqdc8170vgu2vsjj2544fwz54ybb3hz7-csc9cbethbmi129ukq31cko24r28wziaa6\", \"rowIndex\": 0}}"
        ]
}
  ```

  
### InsertTableColumnRequest

```json
{
    "tableId": string,
    "columnIndex": int
}
```

**Field description**:  

| Field            | Type        | Description                                                         |
| ----------- | ------ | ------------------------------------------------------------ |
| tableId | string | Table ID, see the sheet structure definition in Document Data Structure Overview |
| columnIndex | object | Insert column index, starts from 0. Enter 0 to insert before the first column.                 | **Request example**
  ::: note
  This example inserts a column to the left of the second column in a sheet.
  :::
  ```
  {
        "Revision": 178,
        "Requests": [
        "{\"requestType\":\"InsertTableColumnRequestType\",\"insertTableColumnRequest\":{\"tableId\":\"rslqdc8170vgu2vsjj2544fwz54ybb3hz7-csc9cbethbmi129ukq31cko24r28wziaa6\", \"columnIndex\": 1}}"
        ]
}
  ```

### DeleteTableRowsRequest

Delete multiple rows, range to delete RowStartIndex,RowEndIndex)

```json
{
    "tableId": string,
    "rowStartIndex": int,
    "rowEndIndex": int
}
```

**Field description**:  

| Field            | Type        | Description                                                         |
| ------------- | ------ | ------------------------------------------------------------ |
| tableId  | string | Table ID, see the sheet structure definition in [Document Data Structure Overview |
| rowStartIndex | int    | Delete row start index, starts from 0                                    |
| rowEndIndex   | int    | Delete row end index                                               | **Request example**
  ::: note
  This example deletes the first row of the selected sheet.
  :::
  ```
  {
        "Revision": 197,
        "Requests": [
        "{\"requestType\":\"DeleteTableRowsRequestType\",\"deleteTableRowsRequest\":{\"tableId\":\"rslqdc8170vgu2vsjj2544fwz54ybb3hz7-csc9cbethbmi129ukq31cko24r28wziaa6\", \"rowStartIndex\": 0, \"rowEndIndex\": 1}}"
        ]
}
  ```

### DeleteTableColumnsRequest

Delete multiple columns, range to delete ColumnStartIndex,ColumnEndIndex)

```json
{
    "tableId": string,
    "columnStartIndex": int,
    "columnEndIndex": int
}
```

**Field description**:  

| Field            | Type        | Description                                                         |
| ---------------- | ------ | ------------------------------------------------------------ |
| tableId  | string | Table ID, see the sheet structure definition in [Document Data Structure Overview |
| columnStartIndex | int    | Delete column start index, starts from 0                                    |
| columnEndIndex   | int    | Delete column end index                                               | **Request example**
  ::: note
  This example deletes the second and third columns of the selected sheet.
  :::
  ```
  {
        "Revision": 197,
        "Requests": [
        "{\"requestType\":\"DeleteTableColumnsRequestType\",\"deleteTableColumnsRequest\":{\"tableId\":\"rslqdc8170vgu2vsjj2544fwz54ybb3hz7-csc9cbethbmi129ukq31cko24r28wziaa6\", \"columnStartIndex\": 1, \"columnEndIndex\": 3}}"
        ]
}
  ```
  
  

### UpdateTableColumnPropertiesRequest

```json
{
  "tableId": string,
  "columnIndex": int,
  "columnWidth": int
}
```

**Field description**:  

| Field            | Type        | Description                                                         |
| ----------- | ------ | ------------------------------------------------------------ |
| tableId  | string | Table ID, see the sheet structure definition in Document Data Structure Overview |
| columnIndex | int    | The column index, starting from 0                                            |
| columnWidth | int    | The column width, in px                                              | **Request example**
  ::: note
  This example deletes adjusts the width of the second column in the selected sheet to 200.
  :::
  
  ```
  {
        "Revision": 212,
        "Requests": [
        "{\"requestType\":\"UpdateTableColumnPropertiesRequestType\",\"updateTableColumnPropertiesRequest\":{\"tableId\":\"rslqdc8170vgu2vsjj2544fwz54ybb3hz7-csc9cbethbmi129ukq31cko24r28wziaa6\", \"columnIndex\": 1, \"columnWidth\": 200}}"
        ]
}
  ```

### MergeTableCellsRequest

```json
{
  "tableId": string,
  "rowStartIndex": int,
  "rowEndIndex": int,
  "columnStartIndex": int,
  "columnEndIndex": int,
}
```

**Field description**:  

| Field            | Type        | Description                                                         |
| ---------------- | ------ | ------------------------------------------------------------ |
| tableId  | string | Table ID, see the sheet structure definition in Document Data Structure Overview |
| rowStartIndex | int    | Merge row start index, starts from 0                                    |
| rowEndIndex   | int    | Merge row end index                                               |
| columnStartIndex | int    | Merge column start index, starts from 0                                    |
| columnEndIndex   | int    | Merge column end index                                               | **Request example**
  ::: note
  This example merges the 4 cells in the upper-left corner of the selected sheet.
  :::
  ```
  {
        "Revision": 237,
        "Requests": [
        "{\"requestType\":\"MergeTableCellsRequestType\",\"mergeTableCellsRequest\":{\"tableId\":\"rslqdc8170vgu2vsjj2544fwz54ybb3hz7-csc9cbethbmi129ukq31cko24r28wziaa6\", \"rowStartIndex\": 0, \"rowEndIndex\": 2,\"columnStartIndex\": 0, \"columnEndIndex\": 2}}"
        ]
}
  ```
  
### UnmergeTableCellsRequest

```json
{
  "tableId": string,
  "mergedCellId": string
}
```

**Field description**:  

| Field            | Type        | Description                                                         |
| ------------ | ------ | ------------------------------------------------------------ |
| tableId  | string | Table ID, see the sheet structure definition in Document Data Structure Overview |
| mergedCellId | string | ID of the merged cell                                                | **Request example**
  ::: note
  This example splits the 4 cells previously merged in the top-left corner of the selected sheet.
  :::
  ```
  {
        "Revision": 253,
        "Requests": [
        "{\"requestType\":\"UnmergeTableCellsRequestType\",\"unmergeTableCellsRequest\":{\"tableId\":\"rslqdc8170vgu2vsjj2544fwz54ybb3hz7-csc9cbethbmi129ukq31cko24r28wziaa6\", \"mergedCellId\": \"iskfju5o\"}}"
        ]
}
  ```
  
### ReplaceAllTextRequest
|Field|Type|Required|Description|
|---|---|---|---|
|replaceText|string|Yes|Replacement text, can be blank. Unable to contain newlines or tabs, or exceed 1,000 characters.|
|containsText|SubstringMatchCriteria|Yes|Text to find| **SubstringMatchCriteria**  
|Field|Type|Required|Description|
|---|---|---|---|
|text|string |Yes|Content to find, can't be blank. Unable to contain newlines or tabs, or exceed 1,000 characters.|
|matchCase|bool|No|Whether to match text case, false (case insensitive) by default| **Request example**
  ::: note
  This example replaces the "API" field in the document with "Test".
  :::
  ```
  {
        "Revision": 269,
        "Requests": [
        "{\"requestType\":\"ReplaceAllTextRequestType\",\"replaceAllTextRequest\":{\"containsText\":{\"text\":\"API\",\"matchCase\":false},\"replaceText\": \"Test\"}}"
        ]
}
  ```

### UpdateParagraphStyleRequest
|Field|Type|Required|Description|
|---|---|---|---|
|payload|string |Yes|For details, see Document Data Structure ReferenceUpdateParagraphStyleRequest sub-command uses the ParagraphStyle structure body serialized string|
|range|Range|Yes|Range to modify|
|fields|FieldMask|Yes|Field to modify| **FieldMask**
|Field|Type|Required|Description|
|---|---|---|---|
|masks|[]string |Yes|Fields to modify, including "list", "headingLevel", "align", and "quote"| **Request example**
  ::: note
  This example changes TODO in the document to the completed status.
  :::
  ```
  {
        "Revision": 312,
        "Requests": ["{\"RequestType\":\"UpdateParagraphStyleRequestType\",\"updateParagraphStyleRequest\":{\"payload\":\"{\\\"list\\\":{\\\"type\\\":\\\"checkedBox\\\",\\\"indentLevel\\\":1}}\",\"range\":{\"zoneID\":\"0\",\"startIndex\":667,\"endIndex\":678},\"fields\":{\"masks\":[\"list\"]}}}"]
}
  ```
  
### Return body:  

```json
  {
    "code": 0,
    "msg": "Success",
    "data": {}
}
```
### Error code

For details, refer to Server-side error codes.

### batch_update action help
When performing batch_update actions, you must note the following:
1. The request must specify the correct revision version. Otherwise, the data may become confused. In addition, for a collaborative document, the revision version is sometimes used for verification. We recommend you always use the latest revision version when updating a document. Using an older version may cause the action to fail.
2. Some request elements have quantity restrictions:
	- The requests size of a single request can't exceed 10.
  	- The person count of a single insert action can't exceed 100.
	- The docsLink count of a single insert action can't exceed 100.
  	- The chatGroup count of a single insert action can't exceed 5.
	- The block count (including jira, poll, diagram, sheet, and bitable blocks) of a single insert action can't exceed 50.
	- A single row in an insert action can't have more than 6 images.
	- The image and attachment count of a single insert action can't exceed 50.
	- UpdateTitleRequestType can only appear once.
	- UpdateParagraphStyleRequest, ReplaceAllTextRequestType, DeleteTableRowsRequestType, and DeleteTableColumnsRequestType sub-commands can't coexist with other commands.
3. Index requirements of insert and delete actions
	- DeleteContentRangeRequestType sub-commands
		1. Two action types are supported: (1) Delete consecutive blocks (delete newlines, newline after the last block is optional); (2) Delete consecutive elements in a row (newlines are not deleted).
    	2. Blocks and elements must be completely deleted.
	    3. The actual deletion range is closed on the left and open on the right startIndex, endIndex).
    	4. The endIndex can use the last newline location in the document, but this newline is retained.
    	5. When deleting blocks, endIndex can include the last block newline (+1), but it doesn't have to. If included, the newline is deleted.
	- InsertBlocksRequestType sub-commands
    	1. You can only insert content in a new row, so the character at the start of the insert location is a newline. The actual insert location starts from Index. Index can be equal to the last newline (docLengh-1). It can't be greater than or equal to the docLengh.
	-  InsertParagraphElementsRequestType sub-commands
    	1. The insert location can't be in the middle of an element in a row. It must be before or after an element. This action supports empty rows or rows with only one row property. The actual insert location starts from Index. The Index can't be equal to the last newline character. Otherwise, there is only one newline character at the end of the document.
    	2. If a row starts with a placeholder (style), the startIndex of the InsertParagraphElementsRequestType sub-command can't be equal to this placeholder (a placeholder exists when a row starts with paragraphStyle).
	- Batch action ranges must be unique.
    	1. If there are two DeleteContentRangeRequestType actions (d1 and d2), there ranges must be unique. For example, if d2.StartIndex>=d1.StartIndex, then d2.StartIndex>=d1.EndIndex.
    	2. If there is one DeleteContentRangeRequestType action (d) and one Insert action of either type (i), their ranges must be unique. For example, if i.Index >= d.StartIndex, then i.Index >=d.StartIndex or i.Index == d.StartIndex.
	- UpdateParagraphStyleRequest sub-commands
		1. The range startIndex must be equal to the startIndex of a block, and the range endIndex must be equal to the endIndex of a block.
4. Scope verification
 	- If you don't have docsLink scope or the document is not named, the link will be downgraded to a hyperlink.
	- If you don't have person scope (when you are not in the same tenant or are not friends) or the person has been terminated, an error is returned.
	- If you try to insert a chatGroup that you are not in, an error is returned.
5. Formatted sheets and code blocks
	- A single batch_update action on one sheet can only perform one sheet structure modification (add or delete rows and columns, modify column properties, merge cells, or split cells).
	- You can only add one row or column, but can delete a range or rows and columns.
	- Deleting rows and columns clears the affected cells. Deleting all rows or columns will delete the entire sheet. Therefore, after setting the order of insert block, insert row elements, and delete range, you must make sure the ranges of the actions don't conflict.
	- When using DeleteContentRangeRequestType, InsertBlocksRequestType, and InsertParagraphElementsRequestType, you must use the cell zoneId in the zoneId field, for example: xr1m4jw7egd9nefz1s0mdsetenl5fbe3lygxc1azupv81i5t2rjmosw5ta0esuwtn8ksya.
	- You can't embed a table, sheet, or bitable in a sheet.
	- You can't perform a merge/unmerge action and edit action on the same cell.
	- You can't perform a delete row/column action and edit action on the same cell.
	- You can't perform a delete sheet action and edit cell action in a single request.
	- You can't perform a delete sheet action and edit sheet structure action in a single request.
	- You can't perform a delete code zone action and edit code zone action in a single request.

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
| 95009 | Failed | Check for document read permissions. [Add permissions |
| 95010 | internal error | Internal error, please try again later. |
| 95011 | internal error | Internal error, please try again later. |
| 95017 | Specific error message | Check if the revison is correct |
| 95018 | Specific error message | See specific error messages for details |
| 95020 | Specific error message | See specific error messages for details |
| 95023 | revision too old | The Revision is too old, Please use the latest version number |
| 95024 | Failed | Check parameter validity |
| 95025 | Failed | Check if the request is legal json |
| 95026 | Failed | Check batch_update interface request_type |
| 95050 | Specific error message | See specific error messages for details |
| 95053 | this API does not support the Upgraded Docs(docx) | This API does not support the Upgraded Docs(docx) | For details, refer to Server-side error codes.
