---
title: "Obtain the Contacts permission scope"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/scope/list"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/contact/v3/scopes"
service: "contact-v3"
resource: "scope"
section: "Contacts"
scopes:
  - "contact:contact.base:readonly"
  - "contact:contact:access_as_app"
  - "contact:contact:readonly_as_app"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1701771780000"
---

# Obtain the range of contacts data that an app can access

This API is used to obtain the range of contacts data that an app is authorized to access, including the department list, user list, and user group list.
If the access range is all employees, all the first-level departments of the company will be returned. Otherwise, the departments that the administrator selected when setting the access range (excluding sub-departments of selected departments) will be returned.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/contact/v3/scopes |
| HTTP Method | GET |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `contact:contact.base:readonly` `contact:contact:access_as_app` `contact:contact:readonly_as_app` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: "open_id" **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` |
| `department_id_type` | `string` | No | Type of the returned department ID **Example value**: "department_id" **Optional values are**:  - `department_id`: Identify the department with the custom department_id - `open_department_id`: Identify the department with open_department_id  **Default value**: `open_department_id` |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: "AQD9/Rn9eij9Pm39ED40/dk53s4Ebp882DYfFaPFbz00L4CMZJrqGdzNyc8BcZtDbwVUvRmQTvyMYicnGWrde9X56TgdBuS+JKiSIkdexPw=" |
| `page_size` | `int` | No | Page size **Example value**: 50 **Default value**: `50` **Data validation rules**: - Maximum value: `100` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `department_ids` | `string[]` | List of departments that the app is authorized to access. If the access range is "Access to all", the list of all first-level departments of the company will be returned. |
|   `user_ids` | `string[]` | List of users that the app is authorized to access, which will be returned if the "Obtain user's user_id" scope is requested. If the access range is "Access to all", the list of users in all first-level departments of the company will be returned. |
|   `group_ids` | `string[]` | User groups that the app is authorized to access. If the access range is "Access to all", all the user groups of the company will be returned. |
|   `has_more` | `boolean` | Whether the response body has more parameters |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned | ### Response body example

{
    "code": 0,
    "data": {
        "department_ids": [
            "75217143312g9842"
        ],
        "group_ids": [
            "4bgga762"
        ],
        "has_more": true,
        "page_token": "AQD9/Rn9eij9Pm39ED40/e4EuQbBea/ybCE7/PQ1yzQXAVoq7bxTVC%2BI1oJqeoRY",
        "user_ids": [
            "9b2fabg6"
        ]
    },
    "msg": "success"
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 40011 | page size is invalid | The page size is between 1 and 50. |
| 400 | 40012 | page token is invalid error | Page token is invalid. |
