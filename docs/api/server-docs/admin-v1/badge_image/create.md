---
title: "Upload a badge image"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/admin-v1/badge_image/create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/admin/v1/badge_images"
service: "admin-v1"
resource: "badge_image"
section: "Admin"
rate_limit: "20 per minute"
scopes:
  - "admin:badge"
updated: "1686287167000"
---

# Upload a badge image

Through this API, you can upload the files of the badge detail image and the decoration, and obtain the corresponding file key.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/admin/v1/badge_images |
| HTTP Method | POST |
| Rate Limit | 20 per minute |
| Supported app types | custom |
| Required scopes | `admin:badge` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Example value**: "multipart/form-data; boundary=---7MA4YWxkTrZu0gW" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `image_file` | `file` | Yes | The file of the badge image, only supports PNG format, 320 x 320 pixels, and the size does not exceed 1024 KB. **Example value**: file binary |
| `image_type` | `int` | Yes | Type of image **Example value**: 1 **Optional values are**:  - `1`: Badge detail image - `2`: Badge decoration image  **Data validation rules**: - Value range: `1` ～ `2` | ### Request body example

```HTTP
---7MA4YWxkTrZu0gW
Content-Disposition: form-data; name="image_file";
Content-Type: application/octet-stream

file binary
---7MA4YWxkTrZu0gW
Content-Disposition: form-data; name="image_type";

1
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
    "msg": "success",
    "data": {
        "image_key": "f02a98aa-1413-4af6-93ab-431ba9e5f2cg"
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 500 | 1051000 | unknown server error | Service internal error, please try again later |
| 400 | 1051001 | request contain invalid param | Request contain invalid param |
| 403 | 1051002 | request to exceed authority | Request to exceed authority |
| 400 | 1051100 | the image size is too large | The uploaded image size does not meet the requirements |
| 400 | 1051101 | the image format is illegal | The uploaded image format does not meet the requirements |
| 400 | 1051109 | the image height is too small | Image size height is too small |
| 400 | 1051110 | the image height is too large | Image size too high |
| 400 | 1051111 | the aspect ratio of the image is too small | Image size aspect ratio is too small |
| 400 | 1051112 | the aspect ratio of the image is too large | Image size aspect ratio is too large |
| 403 | 1051003 | the tenant's current payment plan does not support the use of this function | Confirm the current tenant package, or upgrade the current package. |
