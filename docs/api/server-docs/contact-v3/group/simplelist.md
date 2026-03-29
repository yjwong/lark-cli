---
title: "List User Group"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/group/simplelist"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/contact/v3/group/simplelist"
service: "contact-v3"
resource: "group"
section: "Contacts"
scopes:
  - "contact:group:readonly"
updated: "1663221577000"
---

# Query the list of user groups

This API is used to query the user group list of a company. You can query common user groups or dynamic user groups respectively.If the app's contact scope is "All employees", you can obtain all user groups in the company. If not, you can only obtain the user groups in the app's contact scope. Click to learn about the range of contacts data that an app can access.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/contact/v3/group/simplelist |
| HTTP Method | GET |
| Supported app types | custom,isv |
| Required scopes | `contact:group:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `page_size` | `int` | No | Page size **Example value**: 50 **Default value**: `50` **Data validation rules**: - Maximum value: `100` |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: "AQD9/Rn9eij9Pm39ED40/dk53s4Ebp882DYfFaPFbz00L4CMZJrqGdzNyc8BcZtDbwVUvRmQTvyMYicnGWrde9X56TgdBuS+JKiSIkdexPw=" |
| `type` | `int` | No | User group type **Example value**: 1 **Optional values are**:  - `1`: Common user group  **Default value**: `1` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `grouplist` | `group[]` | List of user groups |
|     `id` | `string` | User group ID |
|     `name` | `string` | User group name |
|     `description` | `string` | User group description |
|     `member_user_count` | `int` | Number of users among user group members |
|     `member_department_count` | `int` | Number of departments among user group members.Dynamic  groups do not contain departments. |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned |
|   `has_more` | `boolean` | Whether the response body has more parameters | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "grouplist": [
            {
                "id": "g193821",
                "name": "IT outsourcing group",
                "description": "IT outsourcing group. Fine-grained permission control is required for this group.",
                "member_user_count": 2,
                "member_department_count": 0
            }
        ],
        "page_token": "AQD9/Rn9556539ED40/dk53s4Ebp882DYfFaPFbz00L4CMZJrqGdzNyc8BcZtDbwVUvRmQTvyMYicnGWrde9X56TgdBuS+JDTJJDDPw=",
        "has_more": true
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 500 | 40003 | internal error | Internal error. Please provide the X-Request-Id to our agent. [Contact support](https://applink.larksuite.com/client/helpdesk/open?id |
| 400 | 40011 | page size is invalid | The page size is between 1 and 50. |
| 400 | 40012 | page token is invalid error | Page token is invalid. |
