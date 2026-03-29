---
title: "Introduction"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/media/introduction"
service: "drive-v1"
resource: "media"
section: "Docs"
updated: "1750304297000"
---

#  Media overview

This document introduces the fundamental concepts related to media capabilities in cloud space, extra parameters, API list, and more.

## Basic concepts

Media refer to resource files used in documents, such as images, videos, or files in documents. Each media has a unique token as an identifier.

## Methods for obtaining media tokens

The Lark open platform supports uploading and downloading media in documents, spreadsheets, and Base. The methods for obtaining media tokens vary across different online documents. The specific methods are as follows:
| Document Type | Method to Obtain media Token |
| --- | --- |
| Document | Through the Get All Blocks in Document API to obtain the block inforamtion of the specified file block or image block. Within the block inforamtion, the `token` parameter is the media's token. |
| Spreadsheet | Through the Read Multiple Ranges API to obtain the `fileToken` parameter of the specified attachment, which is the media token. |
| Base | Through the Search records API to obtain the `file_token` parameter of the specified attachment, which is the media token. | ## Uploading media instructions

Uploading media refers to uploading various files from the local environment to online documents. The Lark open platform supports uploading media in documents, spreadsheets, and Base.

  :::html
  > The APIs related to uploading media do not support concurrent calls, and the call frequency limit is 5 QPS, 10,000 times per day. Otherwise, the error code 1061045 will be returned, which can be resolved by retrying later.
  :::
### Types of upload points and upload point tokens

To upload specified media to specified cloud documents, you need to first determine the type of upload point (`parent_type`) and the token of the upload point (`parent_node`), and then call the Upload media API or the Upload a media in blocks (Pre­uploading) API. The table below lists the types of upload points and the corresponding upload point tokens for different upload scenarios.

> **Note:** To upload images or files in the knowledge base document, you first need to call Get Knowledge Space Node Information to obtain the actual token of the current knowledge base document. Then, based on the following upload scenarios, determine the value of `parent_node`.

| Upload Scenario | Upload Point Type (parent_type) | Upload Point Token (parent_node) | Example Value |
| --- | --- | --- | --- |
| Upload Image in New Document | docx_image | The unique identifier `block_id` of the new document block, indicating that the image media is uploaded to the specified image block in the new document. For detailed steps on inserting images into a document, refer to Document FAQ - How to insert an image. | doxcnXgNGAtaAraIRVeCfmabcef |
| Upload File in New Document | docx_file | The unique identifier `block_id` of the new document block, indicating that the file media is uploaded to the specified file block in the new document. For detailed steps on inserting files into a document, refer to Document FAQ - How to insert a file. | doxcnXgNGAtaAraIRVeCfmabcef |
| Upload Image in Spreadsheet | sheet_image | The unique identifier `spreadsheet_token` of the spreadsheet, indicating that the image media is uploaded to the specified spreadsheet. Click here to learn how to get the cloud document token. After the image is uploaded, you can continue using the Write Image API to write the image to a specific location in the spreadsheet. | MRLOWBf6J47ZUjmwYRsN8uabcef |
| Upload File in Spreadsheet | sheet_file | The unique identifier `spreadsheet_token` of the spreadsheet, indicating that the file media is uploaded to the specified spreadsheet. Click here to learn how to get the cloud document token. | MRLOWBf6J47ZUjmwYRsN8uabcef |
| Upload Image in Base | bitable_image | The unique identifier `app_token` of the Base, indicating that the image media is uploaded to the specified Base. Click here to learn how to get the cloud document token. | Pc9OpwAV4nLdU7lTy71t6Kabcef |
| Upload File in Base | bitable_file | The unique identifier `app_token` of the Base, indicating that the file media is uploaded to the specified Base. Click here to learn how to get the cloud document token. | Pc9OpwAV4nLdU7lTy71t6Kabcef |
| Upload media to Cloud Space | ccm_import_open | No need to fill in, this scenario is used for importing files. For details, refer to File Import Overview. | / |
| Upload Image in Old Document | doc_image | The unique identifier `doc_token` of the old document, indicating that the image media is uploaded to the specified old document. Click here to learn how to get the cloud document token. | 2olt0Ts4Mds7j7iqzdwrqEabcef |
| Upload File in Old Document | doc_file | The unique identifier `doc_token` of the old document, indicating that the file media is uploaded to the specified old document. Click here to learn how to get the cloud document token. | 2olt0Ts4Mds7j7iqzdwrqEabcef | ### APIs Related to Uploading media

The Lark open platform provides several APIs for uploading media. Depending on the size of the media, you should choose an appropriate method:
- If the current network connection is stable and the media size does not exceed 20 MB, directly call the Upload media API;
- If the current network connection is unstable or the media size exceeds 20 MB, you need to sequentially call the following chunk upload media APIs to upload the media in fixed-length chunks. This reduces single bandwidth usage and improves the success rate of the upload. This upload method allows developers to display the overall upload progress based on the chunk upload progress and supports resuming the upload within a day.
  
  1. Chunk Upload media - Prepare for Upload
     
     Send an initialization request to obtain the upload transaction ID and chunk strategy, preparing for the chunking of the media. The platform fixes the chunk size at 4 MB.
     
  2. Chunk Upload media - Upload Part
     
     Upload the corresponding media chunk based on the upload transaction ID and chunk strategy returned by the prepare for upload API.
     
  3. Chunk Upload media - Complete Upload
    
     After uploading all chunks through the upload part API, you can call this API to trigger the completion of the upload.
     
## Explanation of the `extra` Parameter

For the following media-related scenarios, you need to fill in the `extra` parameter to adapt to special cases.
> - The extra parameter type is a string, and object types need to be converted to strings as input parameters.
> If accessed directly through a URL, the string needs to be URL encoded and placed in the URL query.
> - The extra parameters entered on the API Explorer page do not require encoding or escaping.

### Upload materials to cloud documents

In the following upload scenarios, you need to pass the token of the cloud document where the material is located through the extra parameter. The format of the extra parameter is `"{\"drive_route_token\":\"token of the cloud document where the material is located\"}"`. The example values below are not escaped; please pay attention to escaping in actual use.

| 
Upload scenario | Example value of extra parameter (not escaped) | How to obtain the cloud document token |
| --- | --- | --- |
| Upload an image in the document | {"drive_route_token":"doxcnXgNGAtaAraIRVeCfmabcef"} | Refer to Document Overview to obtain the token, which is the `document_id` of the document. |
| Upload a file in the document | {"drive_route_token":"doxcnXgNGAtaAraIRVeCfmabcef"} | Refer to Document Overview to obtain the token, which is the `document_id` of the document. |
| Upload an image in the spreadsheet | {"drive_route_token":"MRLOWBf6J47ZUjmwYRsN8uabcef"} | Refer to Spreadsheet Overview to obtain the token, which is the `spreadsheetToken` of the spreadsheet. |
| Upload a file in the spreadsheet | {"drive_route_token":"MRLOWBf6J47ZUjmwYRsN8uabcef"} | Refer to Spreadsheet Overview to obtain the token, which is the `spreadsheetToken` of the spreadsheet. |
| Upload an image in the Base | {"drive_route_token":"Pc9OpwAV4nLdU7lTy71t6Kabcef"} | Refer to Base Overview to obtain the Base token. |
| Upload a file in the Base | {"drive_route_token":"Pc9OpwAV4nLdU7lTy71t6Kabcef"} | Refer to Base Overview to obtain the Base token. | ### Uploading media to Cloud Space

Uploading media to cloud space is usually associated with file import scenarios. For details, refer to File Import Overview. In this scenario, the upload point type (`parent_type`) is `ccm_import_open`, and the upload point token (`parent_node`) does not need to be filled in. The `extra` parameter needs to be a JSON-serialized string. Examples of how to fill it in are shown below.

- Import a local file with a .txt extension as a new document:
    ```json
    "{ "obj_type": "docx", "file_extension": "txt" }"
    ```

- Import a local file with a .xlsx extension as a spreadsheet:
    ```json
    "{ "obj_type": "sheet", "file_extension": "xlsx" }"
    ```

- Import a local file with a .md extension as a new document:
    ```json
    "{ "obj_type": "docx", "file_extension": "md" }"
    ```

- Import a local file with a .markdown extension as a new document:
    ```json
    "{ "obj_type": "docx", "file_extension": "markdown" }"
    ```

### bitable_image/bitable_file
To download a media or obtain a temporary download URL for a bitable with advanced permissions, you must add an additional Extra parameter as a URL query parameter for authentication.
```Go
type Extra struct {
   BitablePerm struct {
      TableId string `json:"tableId"`
      Attachments map[string]map[string][]string `json:"attachments"`
   } `json:"bitablePerm"`
}
```
To construct the Extra parameter, perform the following steps:

Step 1. Specify the tableId parameter based on the table where the attachment is located, and construct the Extra object.

Step 2. Specify the attachments parameter based on the field id, record id and file token where the attachment is located. The first key of attachments is the field id, the second key is the record id, and the value is the file token array.

  :::html
  > `attachments` are a required parameter. Not filling them in may lead to excessive delays and interface call failures.
  :::
  
Step 3. Serialize the Extra object into a string.

Step 4. This string is URL-encoded so that it can be safely placed in a URL query.

Example:
```JSON
{"bitablePerm":{"tableId":"tblO6OeNZxfabcef","attachments":{"fld32zZi5I":{"rec0BuOHq":["boxbcsQNT0JsmrztOnX530abcef"]}}}}
// After escaping
https://open.larksuite.com/open-apis/drive/v1/media/boxbcsQNT0JsmrztOnX530abcef/download?extra=%7B%22bitablePerm%22%3A%7B%22tableId%22%3A%22tblO6OeNZxfabcef%22%2C%22attachments%22%3A%7B%22fld32zZi5I%22%3A%7B%22rec0BuOHq%22%3A%5B%22boxbcsQNT0JsmrztOnX530abcef%22%5D%7D%7D%7D%7D
```
## API list

Below is a list of methods related to media. "Store" represents app store applications; "Self-built" represents enterprise self-built applications. For more information on application types, refer to Introduction to Application Types. For an overview of calling server-side APIs, refer to Overview of Processes.

| `Upload mediaPOST /open-apis/drive/v1/media/upload_all` | `bitable:app` `docs:doc` `docs:document.media:upload` `drive:drive` `sheets:spreadsheet` | `tenant_access_token``user_access_token` | **✓** | **✓** |
| --- | --- | --- | --- | --- |
| `Download mediaGET /open-apis/drive/v1/media/:file_token/download` | `bitable:app` `sheets:spreadsheet:readonly` `docs:doc` `docs:doc:readonly` `docs:document.media:download` `drive:drive` `drive:drive:readonly` `sheets:spreadsheet` `sheets:spreadsheet:readonly` | `tenant_access_token``user_access_token` | **✓** | **✓** |
| `Get Temporary Download Link for mediaGET /open-apis/drive/v1/media/batch_get_tmp_download_url` | `bitable:app` `sheets:spreadsheet:readonly` `docs:doc` `docs:doc:readonly` `docs:document.media:download` `drive:drive` `drive:drive:readonly` `sheets:spreadsheet` `sheets:spreadsheet:readonly` | `tenant_access_token``user_access_token` | **✓** | **✓** |
| `Upload a media in blocks-Pre-uploadingPOST /open-apis/drive/v1/media/upload_prepare` | `bitable:app` `docs:doc` `docs:document.media:upload` `drive:drive` `sheets:spreadsheet` | `tenant_access_token``user_access_token` | **✓** | **✓** |
| `Upload a media in blocks-UploadingPOST /open-apis/drive/v1/media/upload_part` | `bitable:app` `docs:doc` `docs:document.media:upload` `drive:drive` `sheets:spreadsheet` | `tenant_access_token``user_access_token` | **✓** | **✓** |
| `Upload a media in blocks-CompletingPOST /open-apis/drive/v1/media/upload_finish` | `bitable:app` `docs:doc` `docs:document.media:upload` `drive:drive` `sheets:spreadsheet` | `tenant_access_token``user_access_token` | **✓** | **✓** | :::multi
