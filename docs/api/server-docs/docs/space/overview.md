---
title: "Overview"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/files/guide/introduction"
section: "Docs"
updated: "1665221844000"
---

# Space Overview

Space is an enterprise cloud disk for convenient management of knowledge resources. All documents in Lark are centrally stored in the cloud, supporting high-speed uploading, downloading and previewing of multi-format files, and viewing, editing, sharing and collaboration through computers and mobile phones anytime and anywhere, creating a knowledge think tank at your fingertips.

| **Folder** | A container for managing files and other folders. |
| --- | --- |
| **File** | A collective term for various types of files, which generally refers to all files in the Space. | ## Resources: Folder

Folders are containers for managing files and other folders. Each folder has a unique token as identification.

> **Note:** 
> Due to the existence of new and old specifications for online resources, the naming of folder token in some interfaces may be folder_token, token, folderToken. Please read the interface documentation carefully when calling to avoid errors caused by naming issues.
> 

## Field description

| Name | Type | Description |
| --- | --- | --- |
| `folder_token` | `string` | The unique identifier of a folder.  Due to the existence of new and old specifications for online resources, the naming of folder token in some interfaces may be folder_token, token, folderToken. Please read the interface documentation carefully when calling to avoid errors caused by naming issues.  **Sample value**: "fldcnK0sP9zb1TejQsaN0S54cHc" **Field permission requirements (choose one)**: `drive:drive` `drive:drive:readonly ` |
| `name` | `string` | The name of the folder. | ### Method list

> "Store" represents store apps; "Custom" represents Custom apps

| `Create folder `POST` /open-apis/drive/v1/files/create_folder ` | View, comment, edit and manage all files in the Space | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| --- | --- | --- | --- | --- |
| `Get the space root directory` `GET` /open-apis/drive/explorer/v2/root_folder/meta > Get the root directory of the Space | View, comment, edit and manage all files in the Space | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Get folder meta information` `GET` /open-apis/drive/explorer/v2/folder/:folderToken/meta > Get the meta information of a folder | `drive:drive` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `List items in folder` `GET` /open-apis/drive/v1/files | `drive:drive` | `tenant_access_token` `user_access_token` | **✓** | **✓** | ## Resources: File

File is a collective term for various types of files, which generally refers to all files in the Space. Each file has a unique token as identification.

## Field description

| Name | Type | Description |
| --- | --- | --- |
| `file_token` | `string` | The unique identifier of a file.  **How to Obtain**: 1. Open a file, from the address bar of the browser, you can get the file_token: ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/a337f74a9f139608f79eddb85ade92e8_59TM59ESuy.png?lazyload=true&width=1919&height=101) 2. Get the file_token from the return value of the related interface calls of folders and files.  **Sample value**: "boxcnK7G8kasZRac70Wo50y6NGh" **Field permission requirements (choose one)**: `drive:drive` `drive:drive:readonly ` |
| `file_name` | `string` | The name of the file. |
| `parent_node` | `string` | The token of the folder where the file is located. | ### Method list

> "Store" represents store apps; "Custom" represents Custom apps

| `Upload file` `POST` /open-apis/drive/v1/files/upload_all > Used to upload files within 20M | `drive:drive` `drive:file` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| --- | --- | --- | --- | --- |
| Shard upload file `POST` /open-apis/drive/v1/files/upload_prepare `POST` /open-apis/drive/v1/files/upload_part `POST` /open-apis/drive/v1/files/upload_finish >When uploading large files (>20M), it is recommended to use multipart upload | `drive:drive` `drive:file` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Download file` `GET` /open-apis/drive/v1/files/:file_token/download | `drive:drive` `drive:file` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Copy file` `POST` /open-apis/drive/v1/files/:file_token/copy | `drive:drive` `drive:file` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Move file` `POST` /open-apis/drive/v1/files/:file_token/move | `drive:drive` `drive:file` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Delete file` `DELETE` /open-apis/drive/v1/files/:file_token | `drive:drive` `drive:file` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Get File Metadata` `POST` /open-apis/drive/v1/metas/batch_query >Support batch acquisition of file metadata | `drive:drive` `drive:file` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
