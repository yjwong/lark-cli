---
title: "Resource introduction"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/instance-comment/overview"
service: "approval-v4"
resource: "instance-comment"
section: "Approval"
updated: "1675167414000"
---

# Resource introduction
Comments or comment responses made by employees in approval instances.

| Parameter | Type | Description |
| --- | --- | --- |
| `comment_id` | `string` | Filling in comment_id means to modify this comment (You can either fill in comment_id or parent_comment_id). **Example value**："7081516627711524883" |
| `parent_comment_id` | `string` | Filling in parent_comment_id means to reply to this comment. (You can either fill in comment_id or parent_comment_id.) **Example value**："7081516627711524883" |
| `content` | `string` | Comment content in json format **Example value**："{"text":"extra ","files":[{"url":"xxx","fileSize":155149,"title":"9a9fedc5cfb01a4a20c715098.png","type":"image","extra":""}]}" |
| `at_info_list` | `comment_at_info[]` | List of members who have been mentioned in the comments (list may be empty) |
|   `user_id` | `string` | ID of the member mentioned **Example value**："579fd9c4" |
|   `name` | `string` | Name of the member mentioned **Example value**："zhangsan" |
|   `offset` | `string` | Offset in the comment, starting at 0 **Example value**："1" |
| `disable_bot` | `boolean` | Only publish comments without sending bot (For custom approvals in Lark's Approval, fill in "False"; otherwise, fill in "True".) **Example value**：false |
| `extra` | `string` | Additional information **Example value**："{\"a\":\"a\"}" | ## Data example
```json
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
```

## User ID description
For how to distinguish between and use user_id, open_id, and union_id, refer to User-related IDs.
