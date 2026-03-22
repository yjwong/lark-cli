---
title: "Upload image"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/image/create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/im/v1/images"
service: "im-v1"
resource: "image"
section: "Messaging"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "im:resource"
  - "im:resource:upload"
updated: "1717574914000"
---

# Upload image

This API can be used to upload an image. You can upload JPEG, PNG, WEBP, GIF, TIFF, BMP, and ICO images.

> Notes:
> - Bot ability must be enabled.
> - The image size cannot exceed 10M, and uploading images with a size of 0 is not supported.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/im/v1/images |
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
| `image_type` | `string` | Yes | Image type **Example value**: "message" **Optional values are**:  - `message`: Sends a message. - `avatar`: Sets the profile photo.  |
| `image` | `file` | Yes | Image content **Note**: The image that you need to upload cannot be larger than 10 MB. **Example value**: Binary file | **cURL example**

```
curl --location --request POST 'https://open.larksuite.com/open-apis/im/v1/images' \
--header 'Authorization: Bearer t-39eec8560faa3dded7971873eb649fd40e70e0f1' \
--header 'Content-Type: multipart/form-data' \
--form 'image_type="message"' \
--form 'image=@"/path/image.jpg"'
```

**HTTP example**
```
POST /open-apis/im/v1/images HTTP/1.1
Host: open.larksuite.com
Authorization: Bearer t-ddf4732fda4aa8a8b1639ee42e8477001d8d5f7c
Content-Length: 285
Content-Type: multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW

----WebKitFormBoundary7MA4YWxkTrZu0gW
Content-Disposition: form-data; name="image_type"

message
----WebKitFormBoundary7MA4YWxkTrZu0gW
Content-Disposition: form-data; name="image"; filename="image.jpg"
Content-Type: image/jpeg

(data)
----WebKitFormBoundary7MA4YWxkTrZu0gW
```

**Python example code**
```
import requests
from requests_toolbelt import MultipartEncoder
# enter pip install requests_toolbelt to install dependent libraries
def uploadImage():
    url = "https://open.larksuite.com/open-apis/im/v1/images"
    form = {'image_type': 'message',
            'image': (open('path/testimage.png', 'rb'))}  # 需要替换具体的path 
    multi_form = MultipartEncoder(form)
    headers = {
        'Authorization': 'Bearer t-xxx',  ## 获取tenant_access_token, 需要替换为实际的token
    }
    headers['Content-Type'] = multi_form.content_type
    response = requests.request("POST", url, headers=headers, data=multi_form)
    print(response.headers['X-Tt-Logid'])  # for debug or oncall
    print(response.content)  # Print Response

if __name__ == '__main__':
    uploadImage()
```

### Request body example

```HTTP
---7MA4YWxkTrZu0gW
Content-Disposition: form-data; name="image_type";

message
---7MA4YWxkTrZu0gW
Content-Disposition: form-data; name="image";
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
|   `image_key` | `string` | Image key | ### Response body example

{
    "code": 0,
    "data": {
        "image_key": "img_v2_xxx"
    },
    "msg": "success"
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
| 400 | 234011 | Can't regonnize the image format. | Unsupported image format, currently only supports uploading images in JPEG, PNG, WEBP, GIF, TIFF, BMP, ICO format. |
| 400 | 234039 | Image resolution exceeds limit. | The resolution of GIF pictures should not be greater than 2000x2000, and the resolution of other pictures should not be greater than 12000x12000. Please use the Upload File interface to upload high-resolution images in file form. |
| 400 | 234041 | Tenant master key has been deleted, please contact the tenant administrator. | Tenant master key has been deleted, please contact the tenant administrator. |
| 400 | 234042 | Hybrid deployment tenant storage error, such as full storage space, please contact tenant administrator. | Hybrid deployment tenant storage error, such as full storage space, please contact tenant administrator. |
