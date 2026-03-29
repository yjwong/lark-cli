---
title: "Create Auto FIlter"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/user_mailbox-rule/create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/mail/v1/user_mailboxes/:user_mailbox_id/rules"
service: "mail-v1"
resource: "user_mailbox-rule"
section: "Email"
rate_limit: "1 per second"
scopes:
  - "mail:user_mailbox.rule:write"
updated: "1745840324000"
---

# Create Auto Filter

Create Auto Filter

> When using the tenant_access_token, it is necessary to apply for data permissions for accessing auto filter resources.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/mail/v1/user_mailboxes/:user_mailbox_id/rules |
| HTTP Method | POST |
| Rate Limit | 1 per second |
| Supported app types | custom |
| Required scopes | `mail:user_mailbox.rule:write` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `user_mailbox_id` | `string` | User email address, use me when using user_access_token **Example value**: "user@xxx.xx 或 me" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `condition` | `rule_condition` | Yes | Match condition |
|   `match_type` | `int` | Yes | Match type **Example value**: 1 **Optional values are**:  - `1`: Meet all conditions - `2`: Meet any condition  **Data validation rules**: - Value range: `1` ～ `2` |
|   `items` | `rule_condition_item[]` | Yes | Match rule list **Data validation rules**: - Length range: `1` ～ `32` |
|     `type` | `int` | Yes | Match condition left value **Example value**: 1 **Optional values are**:  - `1`: Sender - `2`: Recipient - `3`: CC recipient - `4`: To or Cc recipient - `6`: Subject - `7`: Body - `8`: Attachment name - `9`: Attachment type - `10`: Any recipient - `12`: All mails - `13`: Mail Is external - `14`: Mail is spam - `15`: Mail is not spam - `16`: Mail has attachment  **Data validation rules**: - Value range: `1` ～ `16` |
|     `operator` | `int` | Yes | Match condition operator **Example value**: 1 **Optional values are**:  - `1`: Contains - `2`: Does not contains - `3`: Starts with - `4`: Ends with - `5`: Is - `6`: Is not - `7`: Include me - `10`: Is empty  **Data validation rules**: - Value range: `1` ～ `10` |
|     `input` | `string` | No | Match condition right value **Example value**: "hello@world.com" |
| `action` | `rule_action` | Yes | Action after match hit |
|   `items` | `rule_action_item[]` | Yes | List of actions after matching rules **Data validation rules**: - Length range: `1` ～ `32` |
|     `type` | `int` | Yes | Operation type **Example value**: 1 **Optional values are**:  - `1`: Archive - `2`: Delete it - `3`: Mark as read - `4`: Mark as spam - `5`: Never mark as spam - `8`: Label as (not support now) - `9`: Flag - `10`: No notification - `11`: Move to user folder - `12`: Automatically forward (not support now) - `13`: Share to chat (not support now)  **Data validation rules**: - Value range: `1` ～ `13` |
|     `input` | `string` | No | When the type is "Move to folder," fill this field with the ID of the folder. **Example value**: "283412371233" |
| `ignore_the_rest_of_rules` | `boolean` | Yes | Is terminal filter **Example value**: false |
| `name` | `string` | Yes | Name **Example value**: "Mark Sam's email as spam" **Data validation rules**: - Length range: `1` ～ `255` characters |
| `is_enable` | `boolean` | Yes | Is enable **Example value**: false | ### Request body example

{
    "condition": {
        "match_type": 1,
        "items": [
            {
                "type": 1,
                "operator": 1,
                "input": "hello@world.com"
            }
        ]
    },
    "action": {
        "items": [
            {
                "type": 1,
                "input": "283412371233"
            }
        ]
    },
    "ignore_the_rest_of_rules": false,
    "name": "Mark Sam's email as spam",
    "is_enable": false
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `rule` | `rule` | Auto filter |
|     `id` | `string` | ID |
|     `condition` | `rule_condition` | Match condition |
|       `match_type` | `int` | Match type **Optional values are**:  - `1`: Meet all conditions - `2`: Meet any condition  |
|       `items` | `rule_condition_item[]` | Match rule list |
|         `type` | `int` | Match condition left value **Optional values are**:  - `1`: Sender - `2`: Recipient - `3`: CC recipient - `4`: To or CC recipient - `6`: Subject - `7`: Body - `8`: Attachment name - `9`: Attachment type - `10`: Any recipient - `12`: All mails - `13`: Mail is external - `14`: Mail is spam - `15`: Mail is not spam - `16`: Mail has attachment  |
|         `operator` | `int` | Match condition operator **Optional values are**:  - `1`: Contains - `2`: Does not contains - `3`: Starts with - `4`: Ends with - `5`: Is - `6`: Is not - `7`: Include me - `10`: Is empty  |
|         `input` | `string` | Match condition right value |
|     `action` | `rule_action` | Action after match hit |
|       `items` | `rule_action_item[]` | List of actions after matching rules |
|         `type` | `int` | Operation type **Optional values are**:  - `1`: Archive - `2`: Delete it - `3`: Mark as read - `4`: Mark as spam - `5`: Never mark as spam - `8`: Label as (not support now) - `9`: Flag - `10`: No notification - `11`: Move to user folder - `12`: Automatically forward (not support now) - `13`: Share to chat (not support now)  |
|         `input` | `string` | When the type is "Move to folder," fill this field with the ID of the folder. |
|     `ignore_the_rest_of_rules` | `boolean` | Is terminal filter |
|     `name` | `string` | Name |
|     `is_enable` | `boolean` | Is enable | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "rule": {
            "id": "123124123123",
            "condition": {
                "match_type": 1,
                "items": [
                    {
                        "type": 1,
                        "operator": 1,
                        "input": "hello@world.com"
                    }
                ]
            },
            "action": {
                "items": [
                    {
                        "type": 1,
                        "input": "283412371233"
                    }
                ]
            },
            "ignore_the_rest_of_rules": false,
            "name": "将李三的邮件标记为垃圾邮件",
            "is_enable": false
        }
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1230001 | Parameter Error | Please retry after modifying the parameters. |
| 500 | 1230003 | Internal Error | Please try again later. |
| 403 | 1230002 | No Permission | Call this API after becoming a member of the public email or applying for relevant data permissions. |
