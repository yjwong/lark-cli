---
title: "Upload file"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/file/create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/im/v1/files"
service: "im-v1"
resource: "file"
section: "Messaging"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "im:resource"
  - "im:resource:upload"
updated: "1717574914000"
---

# Upload file

Upload files, with videos, audios, and common file types supported.

> Notes:
> - Bot ability must be enabled.
> - The file size can not exceed 30M, and uploading empty files is not allowed.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/im/v1/files |
| HTTP Method | POST |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `im:resource` `im:resource:upload` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Example value**: "multipart/form-data; boundary=---7MA4YWxkTrZu0gW" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `file_type` | `string` | Yes | File type **Example value**: "mp4" **Optional values are**:  - `opus`: Uploads an OPUS audio file. To upload an audio file in other formats, convert the audio file into the OPUS format. For more information about the conversion method, see the following sample code: `ffmpeg -i  SourceFile.mp3 -acodec libopus -ac 1 -ar 16000 TargetFile.opus`. - `mp4`: Uploads an MP4 video file. - `pdf`: Uploads a PDF file. - `doc`: Uploads a Doc file. - `xls`: Uploads an XLS file. - `ppt`: Upload a PPT file - `stream`: Uploads a stream file. That is, files in the stream format are supported in addition to the preceding formats.  |
| `file_name` | `string` | Yes | File name with a suffix **Example value**: "test video.mp4" |
| `duration` | `int` | No | Duration of the video or audio file, in milliseconds. If this field is not specified, no specific duration is displayed. **Example value**: 3000 |
| `file` | `file` | Yes | File content **Example value**: Binary file | > **Note:** In the sample code, you need to replace the file path and authentication token by yourself.
**cURL example**
```
curl --location --request POST 'https://open.larksuite.com/open-apis/im/v1/files' \
--header 'Authorization: Bearer t-39eec8560faa3dded7971873eb649fd40e70e0f1' \
--header 'Content-Type: multipart/form-data; boundary=---7MA4YWxkTrZu0gW' \
--form 'file_type="mp4"' \
--form 'file_name="test video.mp4"' \
--form 'duration="3000"' \
--form 'file=@"/path/test video.mp4"'
```

**HTTP example**
```
POST /open-apis/im/v1/files HTTP/1.1
Host: open.larksuite.com
Authorization: Bearer t-ddf4732fda4aa8a8b1639ee42e8477001d8d5f7c
Content-Type: multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW
Content-Length: 471

----WebKitFormBoundary7MA4YWxkTrZu0gW
Content-Disposition: form-data; name="file_type"

mp4
----WebKitFormBoundary7MA4YWxkTrZu0gW
Content-Disposition: form-data; name="file"; filename="test video.mp4"
Content-Type: 

(data)
----WebKitFormBoundary7MA4YWxkTrZu0gW
Content-Disposition: form-data; name="duration"

3000
----WebKitFormBoundary7MA4YWxkTrZu0gW
Content-Disposition: form-data; name="file_name"

test video.mp4
----WebKitFormBoundary7MA4YWxkTrZu0gW
```

**Python Example**
```
import requests
from requests_toolbelt import MultipartEncoder
# 请注意使用时替换文件path和Authorization
def upload():
  url = "https://open.larksuite.com/open-apis/im/v1/files"
  form = {'file_type': 'stream',
          'file_name': 'text.txt',
          'file':  ('text.txt', open('path/text.txt', 'rb'), 'text/plain')} # 需要替换具体的path  具体的格式参考  https://www.w3school.com.cn/media/media_mimeref.asp
  multi_form = MultipartEncoder(form)
  headers = {
    'Authorization': 'Bearer xxx', ## 获取tenant_access_token, 需要替换为实际的token
  }
  headers['Content-Type'] = multi_form.content_type
  response = requests.request("POST", url, headers=headers, data=multi_form)
  print(response.headers['X-Tt-Logid']) # for debug or oncall
  print(response.content) # Print Response
# Press the green button in the gutter to run the script.

if __name__ == '__main__':
    upload()
```

### Request body example

```HTTP
---7MA4YWxkTrZu0gW
Content-Disposition: form-data; name="file_type";

mp4
---7MA4YWxkTrZu0gW
Content-Disposition: form-data; name="file_name";

test video.mp4
---7MA4YWxkTrZu0gW
Content-Disposition: form-data; name="duration";

3000
---7MA4YWxkTrZu0gW
Content-Disposition: form-data; name="file";
Content-Type: application/octet-stream

Binary file
---7MA4YWxkTrZu0gW
```

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `file_key` | `string` | File key | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "file_key": "file_456a92d6-c6ea-4de4-ac3f-7afcf44ac78g"
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 232096 | Meta writing has stopped, please try again later. | App meta writing has stopped, please try again later. |
| 400 | 234001 | Invalid request param. | Check that the request parameters are correct. |
| 401 | 234002 | Unauthorized. | Authentication failed, contact Oncall to solve. |
| 400 | 234006 | The file size exceed the max value. | The file size exceeds the limit. A file can be up to 30 MB and an image can be up to 10 MB. |
| 400 | 234007 | App does not enable bot feature. | The bot ability is not enabled for the app. |
| 400 | 234010 | File's size can't be 0. | Do not upload a 0-byte file. |
| 400 | 234041 | Tenant master key has been deleted, please contact the tenant administrator. | Tenant master key has been deleted, please contact the tenant administrator. |
| 400 | 234042 | Hybrid deployment tenant storage error, such as full storage space, please contact tenant administrator. | Hybrid deployment tenant storage error, such as full storage space, please contact tenant administrator. |
