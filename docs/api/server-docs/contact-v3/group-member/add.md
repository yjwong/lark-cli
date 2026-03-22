---
title: "Add User Group Member"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/group-member/add"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/contact/v3/group/:group_id/member/add"
service: "contact-v3"
resource: "group-member"
section: "Contacts"
scopes:
  - "contact:group"
updated: "1647437749000"
---

# Add members to a user group

This API is used to add members to a user group (members can only be users currently, and can also be departments in the future). If the app's contact scope is "All employees", you can add any member to any user group. If not, you can only add members in the contact scope to the user groups in the contact scope. Click to learn about the range of contacts data that an app can access.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/contact/v3/group/:group_id/member/add |
| HTTP Method | POST |
| Supported app types | custom |
| Required scopes | `contact:group` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `group_id` | `string` | User group ID **Example value**: "g281721" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `member_type` | `string` | Yes | User group member type. The value is user. **Example value**: "user" **Optional values are**: - `user`: user **Default value**: `user` |
| `member_id_type` | `string` | Yes | When member_type = user, member_id_type represents user_id_type, and the enumeration values are open_id, union_id, user_id **Example value**: "open_id" **Optional values are**: - `open_id`: Indicates open_id of the user when member_type is user - `union_id`: Indicates union_id of the user when member_type is user. - `user_id`: Indicates user_id of the user when member_type is user |
| `member_id` | `string` | Yes | ID of the member to be added **Example value**: "ou_7dab8a3d3cdcc9da365777c7ad535d62" | ### Request body example

```json
{
    "member_type": "user",
    "member_id_type": "open_id",
    "member_id": "ou_7dab8a3d3cdcc9da365777c7ad535d62"
}
```

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions | ### Response body example

```json
{
    "code": 0,
    "data": {},
    "msg": "success"
}
```

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 500 | 40003 | internal error | Internal error. Please provide the X-Request-Id to our agent. [Contact support](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 400 | 42002 | invalid group_id | Invalid user group ID. |
| 400 | 41073 | invalid member_id | Invalid member ID. |
| 400 | 41074 | invalid member_type, must user | Invalid member type. The member type should be user. |
| 400 | 41071 | en_name length exceed 64 character | English name exceeds 64 characters. |
| 400 | 41072 | nickname length exceed 64 character | Alias exceeds 64 characters. |
| 403 | 42009 | no user group authority error | No user group scope. The app's contact scope must include this user group or be "All employees". Click to learn more |
| 403 | 40004 | no dept authority error | The department must be within the range of contacts data that the app can access. Learn more |
| 403 | 41050 | no user authority error | The user must be within the range of contacts data that the app can access. Learn more. |
| 400 | 42005 | member exist in group error | The member already exists in the user group. |
| 400 | 42006 | user has resigned error | The user has been terminated. |
| 400 | 42012 | group member user reached the upper limit | The number of users in the user group has reached the limit. |
| 400 | 42011 | group member department reached the upper limit | The number of departments in the user group has reached the limit. |
