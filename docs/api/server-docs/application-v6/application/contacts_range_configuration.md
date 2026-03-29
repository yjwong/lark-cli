---
title: "Get the Range of Contacts Data an App Can Access"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/application-v6/application/contacts_range_configuration"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/application/v6/applications/:app_id/contacts_range_configuration"
service: "application-v6"
resource: "application"
section: "App Information"
rate_limit: "100 per minute"
scopes:
  - "admin:app.info:readonly"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1682647793000"
---

# Get the Range of Contacts Data an App Can Access

Get the contacts permissions configuration that is actually effective online for a self-built application in the current enterprise

> The permission range of the contact refers to the range of department and user data that can be obtained by the application when calling the related interface of the contact. Apps cannot obtain contact data beyond the scope of permission.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/application/v6/applications/:app_id/contacts_range_configuration |
| HTTP Method | GET |
| Rate Limit | 100 per minute |
| Supported app types | custom |
| Required scopes | `admin:app.info:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `app_id` | `string` | App's app_id, get it from [Developer Console](https://open.larksuite.com/app) > Credentials & Basic Info **Example value**: "cli_9b445f5258795107" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `page_size` | `int` | No | Page size **Example value**: 20 **Default value**: `50` **Data validation rules**: - Value range: `1` ～ `100` |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: "new-e3c5a0627cdf0c2e057da7257b90376a" |
| `department_id_type` | `string` | No | Type of the returned department ID **Example value**: "department_id" **Optional values are**:  - `department_id`: Identify the department with the custom department_id - `open_department_id`: Identify the department with open_department_id  **Default value**: `open_department_id` |
| `user_id_type` | `string` | No | User ID categories **Example value**: "open_id" **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `contacts_range` | `application.app_contacts_range` | Obtain the configuration of the contact permission range that has taken effect online |
|     `contacts_scope_type` | `string` | App version contacts range scope type **Optional values are**:  - `equal_to_availability`: Equal to availability, specific personnel can be queried by Obtain app availability in a company - `some`: Some, see visible_list for details - `all`: All  |
|     `visible_list` | `app_visible_list` | Visible list |
|       `open_ids` | `string[]` | List of open_ids of members to whom the app is visible |
|       `department_ids` | `string[]` | List of IDs of departments to which the app is visible |
|       `group_ids` | `string[]` | List of IDs of groups to which the app is visible |
|   `has_more` | `boolean` | Whether the response body has more parameters |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "contacts_range": {
            "contacts_scope_type": "some",
            "visible_list": {
                "open_ids": [
                    "ou_4065981088f8ef67a504ba8bd6b24d85"
                ],
                "department_ids": [
                    "od-4b4a6907ad726ea07b27b0d2882b7c65"
                ],
                "group_ids": [
                    "b6d1g5dd6fd26186"
                ]
            }
        },
        "has_more": true,
        "page_token": "new-e3c5a0627cdf0c2e057da7257b90376a"
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 210500 | invalid app_id | Check whether the app_id in the request path is valid. |
| 400 | 210501 | no such app in tenant | Check whether the queried app and the app used to call the API are in the same company. |
| 400 | 210503 | target app not a custom app | Check whether the queried app is a custom app. |
| 400 | 210504 | no such app | Check whether the app_id in the request path exists. |
| 400 | 210505 | invalid page_token | page_token is not universal between apps. Check whether this page_token is obtained by the app calling the API. |
| 400 | 210506 | page_token does not exist or has expired | Check whether the page_token is valid. A page_token expires in 2h. Obtain a new one if the current token expires. |
