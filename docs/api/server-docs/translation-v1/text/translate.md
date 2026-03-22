---
title: "Text Translation"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/ai/translation-v1/text/translate"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/translation/v1/text/translate"
service: "translation-v1"
resource: "text"
section: "AI"
scopes:
  - "translation:text"
updated: "1704974643000"
---

# Translate with machine translation

The following languages are supported for translation: "Zh": Chinese ; 
 "Zh-Hant": Traditional Chinese ; 
 "En": English; 
 " Ja ": Japanese ; 
 " Ru ": Russian ; 
 " de ": German ; 
 " Fr ": French ; 
 "It": Italian ; " pl ": Polish ; 
 " Th ": Thai ; 
 "Hi": Hindi ; 
 "Id": Indonesian ; 
 " es ": Spanish ; 
 " Pt ": Portuguese ; 
 " Ko ": Korean ; 
 " vi ": Vietnamese

> Single-tenant traffic limit: 20 QPS. Apps under the same tenant share the 20 QPS. The free version does not support calling this interface.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/translation/v1/text/translate |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes | `translation:text` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `source_language` | `string` | Yes | Source language **Example value**: "zh" |
| `text` | `string` | Yes | Source text **Example value**: "尝试使用一下Lark吧" |
| `target_language` | `string` | Yes | Target language **Example value**: "en" |
| `glossary` | `term[]` | No | Request level glossary with at most 128 terms, valid only in this translation |
|   `from` | `string` | Yes | Associated text **Example value**: "Lark" |
|   `to` | `string` | Yes | Translation **Example value**: "Lark" | ### Request body example

{
    "source_language": "zh",
    "text": "尝试使用一下Lark吧",
    "target_language": "en",
    "glossary": [
        {
            "from": "Lark",
            "to": "Lark"
        }
    ]
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `text` | `string` | Translated text | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "text": "Try to use Lark"
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1140101 | invalid param | Parameter error. Check the input parameters according to the appropriate document. |
| 500 | 1140102 | network anomaly | Service or network error. Request again. |
