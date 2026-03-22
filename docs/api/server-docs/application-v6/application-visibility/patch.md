---
title: "Update app availability"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/application-v6/application-visibility/patch"
method: "PATCH"
api_path: "https://open.larksuite.com/open-apis/application/v6/applications/:app_id/visibility"
service: "application-v6"
resource: "application-visibility"
section: "App Information"
rate_limit: "20 per minute"
scopes:
  - "admin:app.visibility"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1695019505000"
---

# Update app availability

Update the visibility of the custom App or installed Store App in the enterprise, including white list and block list. The update takes effect immediately online.

> When adding a user or department or group through the api, it is necessary to judge in advance whether the corresponding user or department is already in the block list. If it is already in the block list, even if the user or department is added to the white list, the user or department cannot see or use the App, that is to say the block list has priority over the white list.

> Same members (user_id) can not be added to the block list again within 30 seconds, otherwise it will cause Api call failed

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/application/v6/applications/:app_id/visibility |
| HTTP Method | PATCH |
| Rate Limit | 20 per minute |
| Supported app types | custom |
| Required scopes | `admin:app.visibility` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `app_id` | `string` | App's app_id, get it from [Developer Console](https://open.larksuite.com/app) > Credentials & Basic Info **Example value**: "cli_9b445f5258795107" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `department_id_type` | `string` | No | Department ID categories **Example value**: open_department_id **Optional values are**:  - `open_department_id`: Identify the department with open_department_id - `department_id`: Identify the department with the custom department_id  **Default value**: `open_department_id` |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `add_visible_list` | `app_visibility_id_list` | No | indicate some list which will add to the white list, members within white list can see and use the App this parameter will not take effect If the parameter "is_visible_to_all" is not specified and the current range of App avaliablity is  all members OR If the parameter "is_visible_to_all" is set to be "true" |
|   `user_ids` | `string[]` | No | ID list of members, the category of ID is indicated by query parameter "user_id_type" **Example value**: ["ou_84aad35d084aa403a838cf73ee18467"] **Data validation rules**: - Maximum length: `100` |
|   `department_ids` | `string[]` | No | ID list of departments, the category of ID is indicated by query parameter "department_id_type" **Example value**: ["od-4e6ac4d14bcd5071a37a39de902c7141"] **Data validation rules**: - Maximum length: `100` |
|   `group_ids` | `string[]` | No | ID list of groups **Example value**: ["g193821"] **Data validation rules**: - Maximum length: `100` |
| `del_visible_list` | `app_visibility_id_list` | No | indicate some list which will be deleted from the white list this parameter will not take effect If the parameter "is_visible_to_all" is not specified and the current range of App avaliablity is  all members OR If the parameter "is_visible_to_all" is set to be "true" |
|   `user_ids` | `string[]` | No | ID list of members, the category of ID is indicated by query parameter "user_id_type" **Example value**: ["ou_84aad35d084aa403a838cf73ee18467"] **Data validation rules**: - Maximum length: `100` |
|   `department_ids` | `string[]` | No | ID list of departments, the category of ID is indicated by query parameter "department_id_type" **Example value**: ["od-4e6ac4d14bcd5071a37a39de902c7141"] **Data validation rules**: - Maximum length: `100` |
|   `group_ids` | `string[]` | No | ID list of groups **Example value**: ["g193821"] **Data validation rules**: - Maximum length: `100` |
| `add_invisible_list` | `app_visibility_id_list` | No | indicate some list which will add to the block list, members within block list can not see or use the App |
|   `user_ids` | `string[]` | No | ID list of members, the category of ID is indicated by query parameter "user_id_type" Same members (user_id) can not be added to the block list again within 30 seconds, otherwise it will cause Api call failed **Example value**: ["ou_84aad35d084aa403a838cf73ee18467"] **Data validation rules**: - Maximum length: `100` |
|   `department_ids` | `string[]` | No | ID list of departments, the category of ID is indicated by query parameter "department_id_type" **Example value**: ["od-4e6ac4d14bcd5071a37a39de902c7141"] **Data validation rules**: - Maximum length: `100` |
|   `group_ids` | `string[]` | No | ID list of groups **Example value**: ["g193821"] **Data validation rules**: - Maximum length: `100` |
| `del_invisible_list` | `app_visibility_id_list` | No | indicate some list which will be deleted from the block list |
|   `user_ids` | `string[]` | No | ID list of members, the category of ID is indicated by query parameter "user_id_type" **Example value**: ["ou_84aad35d084aa403a838cf73ee18467"] **Data validation rules**: - Maximum length: `100` |
|   `department_ids` | `string[]` | No | ID list of departments, the category of ID is indicated by query parameter "department_id_type" **Example value**: ["od-4e6ac4d14bcd5071a37a39de902c7141"] **Data validation rules**: - Maximum length: `100` |
|   `group_ids` | `string[]` | No | ID list of groups **Example value**: ["g193821"] **Data validation rules**: - Maximum length: `100` |
| `is_visible_to_all` | `boolean` | No | whether the App is available to all members in enterprise true: Yes; false: No; not specified: Maintain current status parameter "add_visible_list/del_visible_list" will not take effect If this parameter is not specified and the current range of App avaliablity is all members OR If this parameter is set to be "true" **Example value**: false | ### Request body example

{
    "add_visible_list": {
        "user_ids": [
            "ou_84aad35d084aa403a838cf73ee18467"
        ],
        "department_ids": [
            "od-4e6ac4d14bcd5071a37a39de902c7141"
        ],
        "group_ids": [
            "g193821"
        ]
    },
    "del_visible_list": {
        "user_ids": [
            "ou_84aad35d084aa403a838cf73ee18467"
        ],
        "department_ids": [
            "od-4e6ac4d14bcd5071a37a39de902c7141"
        ],
        "group_ids": [
            "g193821"
        ]
    },
    "add_invisible_list": {
        "user_ids": [
            "ou_84aad35d084aa403a838cf73ee18467"
        ],
        "department_ids": [
            "od-4e6ac4d14bcd5071a37a39de902c7141"
        ],
        "group_ids": [
            "g193821"
        ]
    },
    "del_invisible_list": {
        "user_ids": [
            "ou_84aad35d084aa403a838cf73ee18467"
        ],
        "department_ids": [
            "od-4e6ac4d14bcd5071a37a39de902c7141"
        ],
        "group_ids": [
            "g193821"
        ]
    },
    "is_visible_to_all": false
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
| 400 | 210001 | invalid request | Please Check your request |
| 200 | 210002 | invalid app_id or app not exists | Please check if App identified by app_id is not exists or not intalled by enterprise |
| 200 | 210003 | Please check if param is empty or if there is conflicts between add and del list | Please Check if request body is empty or if there is conflicts between add_visible_list and del_visible_list or between add_invisible_list and del_invisible_list |
| 200 | 210004 | server internal error | Unexpected Error occured in server process |
| 200 | 210005 | invalid group_ids | Please Check if group_id is correct |
| 200 | 210006 | can not modify visibility of special app | check of App is official and Some App which is specially configured |
