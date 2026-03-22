---
title: "Upload Files"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file/upload_all"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/drive/v1/files/upload_all"
service: "drive-v1"
resource: "file"
section: "Docs"
scopes:
  - "drive:drive"
  - "drive:file"
updated: "1665221891000"
---

# Upload file

Uploads a small file to user's Space.

> Please do not use this interface to upload files larger than 20MB. If you have this requirement, you can try to use the multipart upload interface.

> This API supports up to 5 queries per second (QPS).

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/drive/v1/files/upload_all |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `drive:drive` `drive:file` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Example value**: "multipart/form-data; boundary=---7MA4YWxkTrZu0gW" | ::: note
For more information about how to call the AccessToken in the Docs API, see Get started with the Docs API.

### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `file_name` | `string` | Yes | File name. **Example value**: "demo.pdf" **Data validation rules**: - Maximum length: `250` characters |
| `parent_type` | `string` | Yes | Type of the upload point. **Example value**: "explorer" **Optional values are**:  - `explorer`: My Space.  |
| `parent_node` | `string` | Yes | Token of the folder. For more information about how to obtain the token, see Overview. **Example value**: "fldbcO1UuPz8VwnpPx5a92abcef" |
| `size` | `int` | Yes | The file size in bytes. **Example value**: 1024 **Data validation rules**: - Maximum value: `20971520` |
| `checksum` | `string` | No | Adler-32 checksum of the file. This field is optional. **Example value**: "123423882374238912356" |
| `file` | `file` | Yes | Binary content of the file. **Example value**: file binary | ### cURL example
```
curl --location --request POST 'https://open.larksuite.com/open-apis/drive/v1/files/upload_all' \
--header 'Authorization: Bearer t-e13d5ec1954e82e458f3ce04491c54ea8c9abcef' \
--header 'Content-Type: multipart/form-data' \
--form 'file_name="demo.pdf"' \
--form 'parent_type="explorer"' \
--form 'parent_node="fldbcO1UuPz8VwnpPx5a92abcef"' \
--form 'size="1024"' \
--form 'file=@"/path/demo.pdf"'
```

### Python example
```python
import os
import requests
from requests_toolbelt import MultipartEncoder

def upload_file():
    file_path = "/path/demo.pdf"
    file_size = os.path.getsize(file_path)
    url = "https://open.larksuite.com/open-apis/drive/v1/files/upload_all"
    form = {'file_name': 'demo.pdf',
            'parent_type': 'explorer',
            'parent_node': 'fldbcO1UuPz8VwnpPx5a92abcef',
            'size': str(file_size),
            'file': (open(file_path, 'rb'))}  
    multi_form = MultipartEncoder(form)
    headers = {
        'Authorization': 'Bearer t-e13d5ec1954e82e458f3ce04491c54ea8c9abcef',  ## replace with real tenant_access_token
    }
    headers['Content-Type'] = multi_form.content_type
    response = requests.request("POST", url, headers=headers, data=multi_form)

if __name__ == '__main__':
    upload_file()
```

### Request body example

```HTTP
---7MA4YWxkTrZu0gW
Content-Disposition: form-data; name="file_name";

demo.pdf
---7MA4YWxkTrZu0gW
Content-Disposition: form-data; name="parent_type";

explorer
---7MA4YWxkTrZu0gW
Content-Disposition: form-data; name="parent_node";

fldbcO1UuPz8VwnpPx5a92abcef
---7MA4YWxkTrZu0gW
Content-Disposition: form-data; name="size";

1024
---7MA4YWxkTrZu0gW
Content-Disposition: form-data; name="checksum";

123423882374238912356
---7MA4YWxkTrZu0gW
Content-Disposition: form-data; name="file";
Content-Type: application/octet-stream

file binary
---7MA4YWxkTrZu0gW
```

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `file_token` | `string` | Token of the new file. | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "file_token": "boxcnrHpsg1QDqXAAAyachabcef"
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 200 | 1061001 | internal error. | Internal service error, such as timeout or failure in processing error codes. |
| 400 | 1061002 | params error. | Check whether the request parameters are correct. |
| 404 | 1061003 | not found. | Check whether the resource exists. |
| 403 | 1061004 | forbidden. | Check whether the current user or app has the permissions for the destination upload node. For example, check whether the user has the edit permission for the specified Docs for document upload. |
| 400 | 1061021 | upload id expire. | The upload transaction has expired. Please upload again. |
| 500 | 1061022 | file version conflict. | File version number conflicts. |
| 400 | 1061041 | parent node has been deleted. | Check whether the upload point has been deleted. |
| 400 | 1061042 | parent node out of limit. | The number of materials to upload to the current upload node has reached the limit. Please change the upload point. |
| 400 | 1061043 | file size beyond limit. | Check whether the length of the file is within the limit. For more information, see [File size limits in Drive](https://www.feishu.cn/hc/zh-CN/articles/360049067549). |
| 400 | 1061044 | parent node not exist. | Check whether the upload point exists. |
| 200 | 1061045 | can retry. | Internal error. Please try again later. |
| 400 | 1061109 | file name cqc not passed. | Make sure that the file to upload and the file name meet compliance. |
| 400 | 1061101 | file quota exceeded. | The tenant has reached the maximum capacity. Make sure that the tenant has sufficient capacity and then upload again. |
| 403 | 1061500 | mount node point kill | The mount point does not exist. |
| 400 | 1062007 | upload user not match. | Make sure that the current request is sent by the same user or app as the upload task. |
| 400 | 1062008 | checksum param Invalid. | Make sure that the checksum of the file or file block is correct. |
| 400 | 1062009 | the actual size is inconsistent with the parameter declaration size. | The size of the file to transfer is inconsistent with that specified in the parameter. |
| 400 | 1062010 | block missing, please upload all blocks. | Some file blocks are missing. Make sure that all file blocks are uploaded. |
| 400 | 1062011 | block num out of bounds. | The number of file blocks to upload has reached the limit. Make sure that the file blocks belong to the specified file. |
| 400 | 1061061 | user quota exceeded. | You have reached your maximum personal capacity. Make sure that you have sufficient capacity and then upload again. |
| 403 | 1061073 | no scope auth. | You have no access to the API. |
| 200 | 1064230 | locked_for_data_migration | Data migrating, temporarily unable to upload. |
| 400 | 1062505 | parent node out of size. | The single tree in My Space has reached the maximum size of 400,000. |
| 400 | 1062506 | parent node out of depth. | My Space supports up to 15 levels of directories. |
| 400 | 1062507 | parent node out of sibling num. | The number of nodes mounted to the directory in My Space has reached the limit of **1,500** nodes per level. |
