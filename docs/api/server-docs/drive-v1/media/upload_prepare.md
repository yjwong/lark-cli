---
title: "Multipart Upload Media (Prepare)"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/media/upload_prepare"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/drive/v1/medias/upload_prepare"
service: "drive-v1"
resource: "media"
section: "Docs"
scopes:
  - "bitable:app"
  - "docs:doc"
  - "drive:drive"
  - "sheets:spreadsheet"
  - "vc:material"
updated: "1647178822000"
---

# 

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/drive/v1/medias/upload_prepare |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `bitable:app` `docs:doc` `drive:drive` `sheets:spreadsheet` `vc:material` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `file_name` | `string` | Yes | **Example value**: "" **Data validation rules**: - Maximum length: `250` characters |
| `parent_type` | `string` | Yes | **Example value**: "" **Optional values are**:  - `doc_image`: - `sheet_image`: - `doc_file`: - `sheet_file`: - `vc_virtual_background`: - `bitable_image`: - `bitable_file`: - `moments`: - `ccm_import_open`:  |
| `parent_node` | `string` | Yes | **Example value**: "" |
| `size` | `int` | Yes | **Example value**: **Data validation rules**: - Minimum value: `0` |
| `extra` | `string` | No | **Example value**: "" | ### Request body example

{
    "file_name": "",
    "parent_type": "",
    "parent_node": "",
    "size": ,
    "extra": ""
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `upload_id` | `string` |  |
|   `block_size` | `int` |  |
|   `block_num` | `int` |  | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "upload_id": "",
        "block_size": ,
        "block_num": 
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 200 | 1061001 | unknown error. |  |
| 400 | 1061002 | params error. |  |
| 404 | 1061003 | not found. |  |
| 403 | 1061004 | forbidden. |  |
| 401 | 1061005 | auth failed. |  |
| 200 | 1061006 | internal time out, can retry. |  |
| 404 | 1061007 | file has been delete. |  |
| 400 | 1061008 | invalid file name. |  |
| 400 | 1061021 | upload id expire. |  |
| 400 | 1061041 | parent node has been deleted. |  |
| 400 | 1061042 | parent node out of limit. |  |
| 400 | 1061043 | file size beyond limit. | Check whether the length of the file is within the limit. For more information, see [File size limits in Drive](https://www.larksuite.com/hc/zh-CN/articles/360049067549). |
| 400 | 1061044 | parent node not exist. |  |
| 200 | 1061045 | can retry. |  |
| 400 | 1061109 | file name cqc not passed. |  |
| 400 | 1061113 | file cqc not passed. |  |
| 400 | 1061101 | file quota exceeded. |  |
| 202 | 1062004 | cover generating. |  |
| 202 | 1062005 | file type not support cover. |  |
| 202 | 1062006 | cover no exist. |  |
| 400 | 1062007 | upload user not match. |  |
| 400 | 1062008 | checksum param Invalid. |  |
| 400 | 1062009 | the actual size is inconsistent with the parameter declaration size. |  |
| 400 | 1062010 | block missing, please upload all blocks. |  |
| 400 | 1062011 | block num out of bounds. |  |
| 400 | 1061547 | attachment parent-child relation number exceed. |  |
| 400 | 1061061 | user quota exceeded. |  |
| 403 | 1061073 | no scope auth. |  |
| 400 | 1062012 | file copying. |  |
| 400 | 1062013 | file damaged. |  |
| 403 | 1062014 | dedupe no support. |  |
| 400 | 1062051 | client connect close. |  |
| 400 | 1062505 | parent node out of size. |  |
| 400 | 1062506 | parent node out of depth. |  |
| 400 | 1062507 | parent node out of sibling num. |  |
| 400 | 1061101 | file quota exceeded. |  |
| 400 | 1061061 | user quota exceeded. |  |
