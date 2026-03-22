---
title: "Remove User Group Member"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/group-member/remove"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/contact/v3/group/:group_id/member/remove"
service: "contact-v3"
resource: "group-member"
section: "Contacts"
scopes:
  - "contact:group"
updated: "1647437749000"
---

# Remove members from a user group

This API is used to remove members from a user group (members can only be users currently, and can also be departments in the future). If the app's contact scope is "All employees", you can remove any member from any user group. If not, you can only remove members in the contact scope from the user groups in the contact scope. Click to learn about the range of contacts data that an app can access.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/contact/v3/group/:group_id/member/remove |
| HTTP Method | POST |
| Supported app types | custom |
| Required scopes | `contact:group` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `group_id` | `string` | User group ID **Example value**: "g198123" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `member_type` | `string` | Yes | The type of user group members, the value is user **Example value**: "user" **Optional values are**: - `user`: user **Default value**: `user` |
| `member_id` | `string` | Yes | ID of the member to be removed from a user group **Example value**: "xj82871k" |
| `member_id_type` | `string` | Yes | When member_type is user, member_id_type represents user_id_type, and the enum values are open_id, union_id, and user_id. **Example value**: "open_id" **Optional values are**: - `open_id`: Indicates open_id of the user when member_type is user - `union_id`: Indicates union_id of the user when member_type is user. - `user_id`: Indicates user_id of the user when member_type is user **Default value**: `open_id` | ### Request body example

```json
{
    "member_type": "user",
    "member_id": "xj82871k",
    "member_id_type": "open_id"
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
| 404 | 40004 | no dept authority error | The department must be within the range of contacts data that the app can access. Learn more |
| 403 | 41050 | no user authority error | The user must be within the range of contacts data that the app can access. Learn more. |
