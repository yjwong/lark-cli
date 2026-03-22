---
title: "Language Recognition"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/ai/translation-v1/text/detect"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/translation/v1/text/detect"
service: "translation-v1"
resource: "text"
section: "AI"
scopes:
  - "translation:text"
updated: "1704974643000"
---

# Text language recognition

Machine translation (MT), supporting recognition of over 100 languages. The return value is ISO 639-1 compliant.

> Single-tenant traffic limit: 20 QPS. Apps under the same tenant share the 20 QPS. The free version does not support calling this interface.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/translation/v1/text/detect |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes | `translation:text` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `text` | `string` | Yes | Text whose language is to be recognized **Example value**: "Hello" | ### Request body example

{
    "text": "Hello"
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `language` | `string` | Recognized text language. The return value is ISO 639-1 compliant. | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "language": "en"
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1140101 | invalid param | Parameter error. Check the input parameters according to the appropriate document. |
| 500 | 1140102 | network anomaly | Service or network error. Request again. |
