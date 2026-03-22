---
title: "List User Group Member"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/group-member/simplelist"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/contact/v3/group/:group_id/member/simplelist"
service: "contact-v3"
resource: "group-member"
section: "Contacts"
scopes:
  - "contact:group:readonly"
updated: "1661854549000"
---

# Query the list of members in a user group

This API is used to query the member list of a user group (member type can be user or department)，supports normal user groups and dynamic user groups. If the app's contact scope is "All employees", you can query the member list of any user group in the company. If not, you can only query the member list of the user groups in the app's contact scope. Click to learn about the range of contacts data that an app can access.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/contact/v3/group/:group_id/member/simplelist |
| HTTP Method | GET |
| Supported app types | custom,isv |
| Required scopes | `contact:group:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `group_id` | `string` | User group ID **Example value**: "g128187" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `page_size` | `int` | No | Page size **Example value**: 50 **Default value**: `50` **Data validation rules**: - Maximum value: `100` |
| `page_token` | `string` | No | Paging token. No value at the first request indicates traversal from the beginning. When the paging query involves more results, a new page_token will be returned, which can be used to obtain more results in the next traversal. **Example value**: "AQD9/Rn9eij9Pm39ED40/dk53s4Ebp882DYfFaPFbz00L4CMZJrqGdzNyc8BcZtDbwVUvRmQTvyMYicnGWrde9X56TgdBuS+JKiSIkdexPw=" |
| `member_id_type` | `string` | No | ID type of the members to be obtained. When member_type is user, member_id_type represents user_id_type, and the enum values are open_id, union_id, and user_id. **Example value**: "open_id" **Optional values are**:  - `open_id`: Indicates open_id of the user when member_type is user - `union_id`: Indicates union_id of the user when member_type is user. - `user_id`: Indicates user_id of the user when member_type is user - `department_id`: Indicates department_id of the department when member_type is department  **Default value**: `open_id` |
| `member_type` | `string` | No | Type of user group members to be obtained. **Example value**: "user" **Optional values are**:  - `user`: user - `department`: department  **Default value**: `user` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `memberlist` | `memberlist[]` | Member list |
|     `member_id` | `string` | member ID |
|     `member_type` | `string` | User group member type. The value is user or department. |
|     `member_id_type` | `string` | When member type is user, member_id_type means user_id_type, optional values are open_id, union_id, user_id. Valid only in request parameters. |
|   `page_token` | `string` | The page_token used in the next call |
|   `has_more` | `boolean` | Whether to obtain more by pages | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "memberlist": [
            {
                "member_id": "u287xj12",
                "member_type": "user",
                "member_id_type": "user_id"
            }
        ],
        "page_token": "TDRRV9/Rn9eij9Pm39ED40/dk53s4Ebp882DYfFaPFbz00L4CMZJrqGdzNyc8BcZtDbwVUvRmQTvyMYicnGWrde9X56TgdBuS+JKiJDGexPw=",
        "has_more": true
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 500 | 40003 | internal error | Internal error. Please provide the X-Request-Id to our agent. [Contact support](https://applink.larksuite.com/client/helpdesk/open?id |
| 400 | 40011 | page size is invalid | The page size is between 1 and 50. |
| 400 | 40012 | page token is invalid error | Page token is invalid. |
| 400 | 42002 | invalid group_id | Invalid user group ID. |
| 400 | 41074 | invalid member_type | Invalid member type. |
| 400 | 41071 | en_name length exceed 64 character | English name exceeds 64 characters. |
| 400 | 41072 | nickname length exceed 64 character | Alias exceeds 64 characters. |
| 403 | 42009 | no group authority | No user group scope. The app's contact scope must include this user group or be "All employees". Click to learn more |
