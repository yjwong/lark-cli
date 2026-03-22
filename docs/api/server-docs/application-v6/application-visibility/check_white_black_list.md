---
title: "Obtain whether user in app visility white or black list"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/application-v6/application-visibility/check_white_black_list"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/application/v6/applications/:app_id/visibility/check_white_black_list"
service: "application-v6"
resource: "application-visibility"
section: "App Information"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "application:application:self_manage"
  - "admin:app.info:readonly"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1682583665000"
---

# Get if user or department is in an app's allowlist or blocklist

This api is used to query whether the user, department, or user group is in the white or black list of the application

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/application/v6/applications/:app_id/visibility/check_white_black_list |
| HTTP Method | POST |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `application:application:self_manage` `admin:app.info:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `app_id` | `string` | The AppID of the application can be viewed on the Developer Console > Credentials & Basic Info. When you need to query other application information, you must apply for `admin:app.info:readonly` Permissions, when only querying the information of this application, you can fill the AppID of the application itself. Example value: "cli_a3a3d00b40b8d01b" **Example value**: "cli_a3a3d00b40b8d01b" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: "open_id" **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` |
| `department_id_type` | `string` | No | Department id categories **Example value**: "department_id" **Optional values are**:  - `department_id`: Identifies a department with a custom department_id - `open_department_id`: Identifies a department with open_department_id  **Default value**: `department_id` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_ids` | `string[]` | No | The id list of the user you want to query, entered according to user_id_type, up to 100. You can call the Obtain the list of users directly under a department interface to obtain. **Example value**: ["ou_d317f090b7258ad0372aa53963cda70d"] **Data validation rules**: - Maximum length: `100` |
| `department_ids` | `string[]` | No | The id list of the department you want to query, up to 100. You can call the Obtain the list of sub-departments interface to obtain. **Example value**: ["od-aa2c50a04769feefededb7a05b7525a8"] **Data validation rules**: - Maximum length: `100` |
| `group_ids` | `string[]` | No | the id list of the user group you want to query, up to 100. You can call the Query the list of user groups interface to obtain it. **Example value**: ["96815a9cd9beg8g4"] **Data validation rules**: - Maximum length: `100` | ### Request body example

{
    "user_ids": [
        "ou_d317f090b7258ad0372aa53963cda70d"
    ],
    "department_ids": [
        "od-aa2c50a04769feefededb7a05b7525a8"
    ],
    "group_ids": [
        "96815a9cd9beg8g4"
    ]
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `user_visibility_list` | `application.visibility.user_white_black_info[]` | List of user visibility results for the query, the app is visible if the user is on the whitelist or paid list and not on the blacklist |
|     `user_id` | `string` | The user ID to be queried, the ID type is consistent with the user_id_type parameter type |
|     `in_white_list` | `boolean` | Whether it is in the whitelist. **Optional values**: - **true**: in the whitelist - **false**: not in the whitelist |
|     `in_black_list` | `boolean` | Whether it is in the blacklist. **Optional values**: - **true**: in the blacklist - **false**: not in the blacklist |
|     `in_paid_list` | `boolean` | Whether it is in the paid list. **Optional values**: - **true**: in the paid list - **false**: not in the paid list |
|   `department_visibility_list` | `application.visibility.department_white_black_info[]` | List of department visibility results for the query, the app is visible if the department is on the whitelist and not on the blacklist |
|     `department_id` | `string` | The department ID to be queried |
|     `in_white_list` | `boolean` | Whether it is in the whitelist. **Optional values**: - **true**: in the whitelist - **false**: not in the whitelist |
|     `in_black_list` | `boolean` | Whether it is in the blacklist. **Optional values**: - **true**: in the blacklist - **false**: not in the blacklist |
|   `group_visibility_list` | `application.visibility.group_white_black_info[]` | List of user group visibility results for the query, the app is visible if the user group is on the whitelist and not on the blacklist |
|     `group_id` | `string` | The user group ID to be queried |
|     `in_white_list` | `boolean` | Whether it is in the whitelist. **Optional values**: - **true**: in the whitelist - **false**: not in the whitelist |
|     `in_black_list` | `boolean` | Whether it is in the blacklist. **Optional values**: - **true**: in the blacklist - **false**: not in the blacklist | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "user_visibility_list": [
            {
                "user_id": "ou_d317f090b7258ad0372aa53963cda70d",
                "in_white_list": false,
                "in_black_list": false,
                "in_paid_list": false
            }
        ],
        "department_visibility_list": [
            {
                "department_id": "od-aa2c50a04769feefededb7a05b7525a8",
                "in_white_list": false,
                "in_black_list": false
            }
        ],
        "group_visibility_list": [
            {
                "group_id": "96815a9cd9beg8g4",
                "in_white_list": false,
                "in_black_list": false
            }
        ]
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 210001 | param is invalid | Check parameters |
