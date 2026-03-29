---
title: "Update the contacts permission scope an app can access"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/application-v6/application-contacts_range/patch"
method: "PATCH"
api_path: "https://open.larksuite.com/open-apis/application/v6/applications/:app_id/contacts_range"
service: "application-v6"
resource: "application-contacts_range"
section: "App Information"
rate_limit: "20 per minute"
scopes:
  - "application:application.contacts_range:write"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1695029165000"
---

# Update the contacts permission scope an app can access

Update the contacts permission scope of the custom app or store app in the enterprise, This API is only available to custom apps. The update takes effect immediately online.

> The contacts permission scope  refers to the range of department and user data that can be obtained by the application when calling the related APIs of the contacts. Apps cannot obtain contacts beyond the scope of permission.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/application/v6/applications/:app_id/contacts_range |
| HTTP Method | PATCH |
| Rate Limit | 20 per minute |
| Supported app types | custom |
| Required scopes | `application:application.contacts_range:write` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `app_id` | `string` | App's app_id, get it from [Developer Console](https://open.larksuite.com/app) > Credentials & Basic Info **Example value**: "cli_dsfjksdfee1" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` |
| `department_id_type` | `string` | No | Department ID categories **Example value**: open_department_id **Optional values are**:  - `open_department_id`: Identify the department with open_department_id - `department_id`: Identify the department with the custom department_id  **Default value**: `open_department_id` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `contacts_range_type` | `string` | Yes | indicates the way how to update **Example value**: "some" **Optional values are**:  - `equal_to_availability`: the permission range of contacts will keep consistent with the range of App availability - `some`: the permission range of contacts will be partial, parameter "add_visible_list/del_visible_list" will take effect and do incremental updates - `all`: All contacts data can be accessed  |
| `add_visible_list` | `app_contacts_range_id_list` | No | indicate some list which will add to visible list，The infomation of members within visible list can be accessed by app This parameter take effect only if "contacts_range_type" is set to be "some" |
|   `user_ids` | `string[]` | No | ID list of members, the category of ID is indicated by query parameter "user_id_type" **Example value**: ["ou_7dab8a3d3cdcc9da365777c7ad535d62"] **Data validation rules**: - Maximum length: `100` |
|   `department_ids` | `string[]` | No | ID list of departments, the category of ID is indicated by query parameter "department_id_type" **Example value**: ["od-4e6ac4d14bcd5071a37a39de902c7141"] **Data validation rules**: - Maximum length: `100` |
|   `group_ids` | `string[]` | No | ID list of groups **Example value**: ["g12334"] **Data validation rules**: - Maximum length: `100` |
| `del_visible_list` | `app_contacts_range_id_list` | No | indicate some list which will be deleted from visible list This parameter take effect only if "contacts_range_type" is set to be "some" |
|   `user_ids` | `string[]` | No | ID list of members, the category of ID is indicated by query parameter "user_id_type" **Example value**: ["ou_7dab8a3d3cdcc9da365777c7ad535d62"] **Data validation rules**: - Maximum length: `100` |
|   `department_ids` | `string[]` | No | ID list of departments, the category of ID is indicated by query parameter "department_id_type" **Example value**: ["od-4e6ac4d14bcd5071a37a39de902c7141"] **Data validation rules**: - Maximum length: `100` |
|   `group_ids` | `string[]` | No | ID list of groups **Example value**: ["g12334"] **Data validation rules**: - Maximum length: `100` | ### Request body example

{
    "contacts_range_type": "some",
    "add_visible_list": {
        "user_ids": [
            "ou_7dab8a3d3cdcc9da365777c7ad535d62"
        ],
        "department_ids": [
            "od-4e6ac4d14bcd5071a37a39de902c7141"
        ],
        "group_ids": [
            "g12334"
        ]
    },
    "del_visible_list": {
        "user_ids": [
            "ou_7dab8a3d3cdcc9da365777c7ad535d62"
        ],
        "department_ids": [
            "od-4e6ac4d14bcd5071a37a39de902c7141"
        ],
        "group_ids": [
            "g12334"
        ]
    }
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {}
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 210001 | param is invalid | Please check your request parameters |
| 200 | 210002 | invalid app_id or app not exists | Please check if App identified by app_id is not exists or not intalled by enterprise |
| 200 | 210003 | please check if param is empty or if there is conflicts between add and del list | Please Check if request body is empty or if there is conflicts between add_visible_list and del_visible_list |
| 200 | 210004 | server internal error | Unexpected Error occured in server process |
| 200 | 210005 | invalid group_ids | Please Check if group_id is correct |
| 200 | 210006 | can not modify cantact of special app or official app | check of App is official and Some App which is specially configured |
