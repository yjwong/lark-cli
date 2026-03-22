---
title: "Basic Recognition"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/ai/optical_char_recognition-v1/image/basic_recognize"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/optical_char_recognition/v1/image/basic_recognize"
service: "optical_char_recognition-v1"
resource: "image"
section: "AI"
scopes:
  - "optical_char_recognition:image"
updated: "1744706460000"
---

# Basic image recognition (OCR)

Basic picture recognition interface, recognize the text in the picture, and return the text list by area

> Single tenant current limit: 20QPS, there is no current limit for applications under the same tenant, and share the current tenant’s 20QPS current limit

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/optical_char_recognition/v1/image/basic_recognize |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes | `optical_char_recognition:image` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `image` | `string` | No | Picture data after base64 **Example value**: "Image binary data after base64" | ### Request body example

{
    "image": "Image binary data after base64"
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `text_list` | `string[]` | Recognize by area, return to text list | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "text_list": [
            "hello"
        ]
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1630101 | invalid param | The parameter is wrong, please refer to the document to check the input parameter |
| 500 | 1630102 | network anomaly | The back-end service is abnormal or the network is abnormal, you can request again |
