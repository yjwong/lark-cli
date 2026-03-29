---
title: "Create comments"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/instance-comment/create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/approval/v4/instances/:instance_id/comments"
service: "approval-v4"
resource: "instance-comment"
section: "Approval"
scopes:
  - "approval:approval"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1675167415000"
---

# Create comments

Create or modify comments, or comment on replies under an approval instance (excluding approved, rejected, transferred, and other additional reasons or opinions).

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/approval/v4/instances/:instance_id/comments |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes | `approval:approval` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `instance_id` | `string` | Approval instance code (or tenant custom approval instance ID) **Example value**: "6A123516-FB88-470D-A428-9AF58B71B3C0" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: "open_id" **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` |
| `user_id` | `string` | Yes | User ID **Example value**: "e5286g26" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `content` | `string` | No | Comment content in json format **Example value**: "{\"text\":\"agree\",\"files\":[{\"url\":\"xxx\",\"fileSize\":155149,\"title\":\"9a9fedc5cfb01a4a20c715098.png\",\"type\":\"image\",\"extra\":\"\"}]}" |
| `at_info_list` | `comment_at_info[]` | No | List of members who have been mentioned in the comments (list may be empty) |
|   `user_id` | `string` | Yes | ID of the member mentioned **Example value**: "579fd9c4" |
|   `name` | `string` | Yes | Name of the member mentioned **Example value**: "zhangsan" |
|   `offset` | `string` | Yes | Offset in the comment, starting at 0 **Example value**: "1" |
| `parent_comment_id` | `string` | No | Filling in parent_comment_id means to reply to this comment. (You can either fill in comment_id or parent_comment_id.) **Example value**: "7081516627711524883" |
| `comment_id` | `string` | No | Filling in comment_id means to modify this comment (You can either fill in comment_id or parent_comment_id). **Example value**: "7081516627711524883" |
| `disable_bot` | `boolean` | No | Only publish comments without sending bot (For custom approvals in Lark's Approval, fill in "False"; otherwise, fill in "True".) **Example value**: false |
| `extra` | `string` | No | Additional information **Example value**: "{\"a\":\"a\"}" | ### Request body example

{
    "content": "{\"text\":\"agree\",\"files\":[{\"url\":\"xxx\",\"fileSize\":155149,\"title\":\"9a9fedc5cfb01a4a20c715098.png\",\"type\":\"image\",\"extra\":\"\"}]}",
    "at_info_list": [
        {
            "user_id": "579fd9c4",
            "name": "zhangsan",
            "offset": "1"
        }
    ],
    "parent_comment_id": "7081516627711524883",
    "comment_id": "7081516627711524883",
    "disable_bot": false,
    "extra": "{\"a\":\"a\"}"
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `comment_id` | `string` | Comment ID adding | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "comment_id": "7081516627711606803"
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1390001 | param is invalid | Parameter error |
