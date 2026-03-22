---
title: "Update application information"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/application-v6/application/patch"
method: "PATCH"
api_path: "https://open.larksuite.com/open-apis/application/v6/applications/:app_id"
service: "application-v6"
resource: "application"
section: "App Information"
scopes:
  - "application:application"
updated: "1646274616000"
---

# Update app category information

Updates app category information (Category changes can affect app sorting in Workplace, so proceed with caution).

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/application/v6/applications/:app_id |
| HTTP Method | PATCH |
| Supported app types | custom |
| Required scopes | `application:application` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `app_id` | `string` | App ID **Example value**: "cli_9b445f5258795107" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `lang` | `string` | Yes | Specifies the language of returned results **Example value**: "zh_cn" **Optional values are**: - `zh_cn`: Chinese - `en_us`: English - `ja_jp`: Japanese | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `common_categories` | `string[]` | No | Internationalized description of categories **Example value**: Project Management **Data validation rules**: - Length range: `1` ～ `3` | ### Request body example

```json
{
    "common_categories": [
        "Analysis tools"
    ]
}
```

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- | ### Response body example

```json
{
    "code": 0,
    "msg": "success",
    "data": {

    }
}
```

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 210503 | invalid app_id | Check whether the app_id in the request path is valid. |
| 400 | 210504 | no such app in tenant | Check whether the queried app and the app used to call the API are in the same company. |
| 400 | 210505 | target app not a custom app | Check whether the queried app is a custom app. |
| 400 | 210506 | no such app | Check whether the app_id in the request path exists. |
| 400 | 211000 | size of common categories out of range, should be between 1 and 3 | Check whether the length of the specified categories list is in the range [1, 3]. |
| 400 | 211001 | common_categories[%d](%s) not exist (index starts from 0) | Verify whether the specified category value is correct based on the subscript in the prompt. The category language must match the specified lang parameter. |
