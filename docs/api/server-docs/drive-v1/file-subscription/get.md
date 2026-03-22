---
title: "Get subscription status"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file-subscription/get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/drive/v1/files/:file_token/subscriptions/:subscription_id"
service: "drive-v1"
resource: "file-subscription"
section: "Docs"
scopes:
  - "drive:drive"
  - "drive:drive:readonly"
updated: "1647001287000"
---

# Get subscription status

Get the status of the subscription based on the subscription ID

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/drive/v1/files/:file_token/subscriptions/:subscription_id |
| HTTP Method | GET |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `drive:drive` `drive:drive:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `file_token` | `string` | file token **Example value**: "doxcnxxxxxxxxxxxxxxxxxxxxxx" |
| `subscription_id` | `string` | Subscription ID **Example value**: "1234567890987654321" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `subscription_id` | `string` | No | subsription id **Example value**: "1234567890987654321" |
| `subscription_type` | `string` | No | subscription type **Example value**: "comment_update" **Optional values are**: - `comment_update`: subscribe comment |
| `is_subcribe` | `boolean` | No | Whether to subscribe **Example value**: true |
| `file_type` | `string` | Yes | file type **Example value**: "doc" **Optional values are**: - `doc`: doc - `docx`: docx - `wiki`: wiki | ### Request body example

```json
{"file_type":"docx"}
```

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
| ∟ `subscription` | `file.subscription` | Document Subscription Information |
| ∟ `subscription_id` | `string` | subsription id |
| ∟ `subscription_type` | `string` | subscription type **Optional values are**: - `comment_update`: subscribe comment |
| ∟ `is_subcribe` | `boolean` | Whether to subscribe |
| ∟ `file_type` | `string` | file type **Optional values are**: - `doc`: doc - `docx`: docx - `wiki`: wiki | ### Response body example

```json
{
    "code": 0,
    "data": {
        "file_type": "docx",
        "is_subcribe": false,
        "subscription_id": "xxxxxxxx",
        "subscription_type": "comment_update"
    },
    "msg": "success"
}
```

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1064000 | Illegal parameter | Check parameter validity |
| 403 | 1064030 | Permission denied | Check if there is permission for the document to be read |
| 404 | 1064040 | Token not exist | Check whether the document can be accessed normally |
| 500 | 1065000 | Internal Server Error | Try again, if the stability fails, please contact the oncall staff of the relevant business party |
