---
title: "Upload progress record image"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/okr-v1/image/upload"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/okr/v1/images/upload"
service: "okr-v1"
resource: "image"
section: "OKR"
scopes:
  - "okr:okr"
updated: "1728975249000"
---

# Upload progress record image

Upload progress record image. After successfully calling this interface, you can continue to call Create OKR Progress Record or Update OKR Progress Record, and pass the returned `url` parameter and `file_token` parameter into the `imageList` parameter.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/okr/v1/images/upload |
| HTTP Method | POST |
| Supported app types | custom |
| Required scopes | `okr:okr` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Example value**: "multipart/form-data; boundary=---7MA4YWxkTrZu0gW" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `data` | `file` | Yes | Picture **Example value**: file binary |
| `target_id` | `string` | Yes | Target ID of the picture **Example value**: "6974586812998174252" |
| `target_type` | `int` | Yes | The target type used by the picture **Example value**: 1 **Optional values are**:  - `2`: Objective - `3`: Key result  | ### Request body example

```HTTP
---7MA4YWxkTrZu0gW
Content-Disposition: form-data; name="data";
Content-Type: application/octet-stream

file binary
---7MA4YWxkTrZu0gW
Content-Disposition: form-data; name="target_id";

6974586812998174252
---7MA4YWxkTrZu0gW
Content-Disposition: form-data; name="target_type";

1
---7MA4YWxkTrZu0gW
```

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `image_info` | \- |
|   `file_token` | `string` | Image token |
|   `url` | `string` | Image Download Link | ### Response body example

{{
"Code": 0,
"Data": {
"file_token": "boxbcc6DmPfgi4rNXIaGfptc9HX",
"Url": "https://internal-api-okr.larksuite-boe.cn/stream/api/downloadFile/?file_token=boxbcc6DmPfgi4rNXIaGfptc9HX&ticket=eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ0YXJnZXRfaWQiOiI3MDQxNTA5NTg4MzkwMzQ2NzcyIiwidGFyZ2V0X3R5cGUiOjEsImFjdGlvbiI6MiwiZmlsZV90b2tlbiI6ImJveGJjYzZEbVBmZ2k0ck5YSWFHZnB0YzlIWCIsInVzZXJfaWQiOiI2OTY5ODU1NTAxNzQ0ODM0MDkyIiwidGVuYW50X2lkIjoiNjg3NzUwMjY4NzYwOTQwNjk5MCIsImV4cCI6MTY0MDY4MzI4OX0.VqOLS7kDtCuhyU_WuWeXvxg1XIyJxskBfNGFQP8uGkCBhYh9scwcbWQJ4xubAZs3cmsrPMVm-aho3tz5d7NT5Q"
},
"Msg": "success"
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 500 | 1009999 | internal server error | Internal error |
| 500 | 1009998 | system exception | System anomaly |
| 400 | 1001001 | invalid parameters | Invalid parameter |
| 400 | 1001002 | no permission | No permission |
| 400 | 1001003 | user not found | User does not exist |
| 400 | 1001004 | okr data not found | Okr data does not exist |
