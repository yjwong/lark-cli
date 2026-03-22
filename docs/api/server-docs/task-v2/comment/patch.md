---
title: "Patch Comment"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/comment/patch"
method: "PATCH"
api_path: "https://open.larksuite.com/open-apis/task/v2/comments/:comment_id"
service: "task-v2"
resource: "comment"
section: "Tasks v2"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "task:comment:write"
updated: "1699521686000"
---

# Patch Comment

Update a comment.

When updating, fill in the field names of all comments to be modified in the 'update_fields 'field, and fill in the new value of the field to be modified in the'comment' field. For update api specification, please refer to the "About Resource Update" section in Feature Overview .

Currently only the "conent" field for updating comments is supported.

> Updating comments requires the read permission of the comment attribution task, and can only update comments created by yourself. See "How are Tasks Authorized?" section in Task Feature Overview.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/task/v2/comments/:comment_id |
| HTTP Method | PATCH |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes | `task:comment:write` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `comment_id` | `string` | Comment ID to update **Example value**: "7198104824246747156" **Data validation rules**: - Maximum length: `100` characters | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Default value**: `open_id` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `comment` | `input_comment` | Yes | Comment data to update. |
|   `content` | `string` | No | Comment content to update. When updating, empty content is not allowed. Supports up to 3000 utf8 characters. **Example value**: "Raise a glass to invite the bright moon, face the shadow into three people" **Data validation rules**: - Maximum length: `10000` characters |
| `update_fields` | `string[]` | Yes | Fields to update, support   Comment content   **Example value**: ["Content"] | ### Request body example

{
    "comment": {
        "content": "Raise a glass to invite the bright moon, face the shadow into three people"
    },
    "update_fields": [
        "Content"
    ]
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `comment` | `comment` | Updated comments |
|     `id` | `string` | Comment id |
|     `content` | `string` | Comment content |
|     `creator` | `member` | Comment creator |
|       `id` | `string` | Indicates the id of member |
|       `type` | `string` | Type of member |
|       `role` | `string` | Member role |
|     `reply_to_comment_id` | `string` | The id of the comment being replied to. Empty if not replying to the comment. |
|     `created_at` | `string` | Comment creation timestamp (ms) |
|     `updated_at` | `string` | Comments Updatetimestamp (ms) |
|     `resource_type` | `string` | The resource type associated with the task |
|     `resource_id` | `string` | Resource ID associated with the task | ### Response body example

{
    "code": 0,
    "msg": "Success",
    "data": {
        "comment": {
            "id": "7197020628442939411",
            "content": "This is a comment",
            "creator": {
                "id": "ou_2cefb2f014f8d0c6c2d2eb7bafb0e54f",
                "type": "user",
                "role": "Creator"
            },
            "reply_to_comment_id": "7166825117308174356",
            "created_at": "1675742789470",
            "updated_at": "1675742789470",
            "resource_type": "task",
            "resource_id": "ccb55625-95d2-2e80-655f-0e40bf67953f"
        }
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1470400 | Incorrect request parameters, such as update_fields with an unsupported field name. | The reason for the error is shown in the returned msg prompt. |
| 404 | 1470404 | The comment to be updated does not exist or has been deleted. | Confirm that the comment you want to update exists or has been deleted. |
| 500 | 1470500 | Server error. | Try retrying the call. If it continues to fail, please contact support for feedback. |
| 403 | 1470403 | Missing permission to update comments. | Updating comments requires the read permission of the comment attribution task, and can only update comments sent by yourself. See [How is the task authenticated?] (/ssl: ttdoc/uAjLw4CM/ukTMukTMukTM/task-v2/faq) |
