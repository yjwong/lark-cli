---
title: "Get User Group"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/group/get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/contact/v3/group/:group_id"
service: "contact-v3"
resource: "group"
section: "Contacts"
scopes:
  - "contact:group:readonly"
updated: "1661854548000"
---

# Query a user group

This API is used to query the basic information of a user group by user group ID. You can query a common user group or a dynamic user groupsMake sure that the app's contact scope includes the user group or is "All employees". Click to learn about the range of contacts data that an app can access.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/contact/v3/group/:group_id |
| HTTP Method | GET |
| Supported app types | custom,isv |
| Required scopes | `contact:group:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `group_id` | `string` | User group ID **Example value**: "g193821" | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `group` | `group` | User group details |
|     `id` | `string` | User group ID |
|     `name` | `string` | User group name |
|     `description` | `string` | User group description |
|     `member_user_count` | `int` | Number of users among user group members |
|     `member_department_count` | `int` | Number of departments among user group members.Dynamic  groups do not contain departments. | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "group": {
            "id": "g193821",
            "name": "IT outsourcing group",
            "description": "IT outsourcing group. Fine-grained permission control is required for this group.",
            "member_user_count": 2,
            "member_department_count": 0
        }
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 500 | 40003 | internal error | Internal error. Please provide the X-Request-Id to our agent. [Contact support](https://applink.larksuite.com/client/helpdesk/open?id |
| 400 | 42002 | invalid group_id | Invalid user group ID. |
