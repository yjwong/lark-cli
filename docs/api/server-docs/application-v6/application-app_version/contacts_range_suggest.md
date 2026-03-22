---
title: "Get the Range of Contacts Data to Access in an App's Version Release Request"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/application-v6/application-app_version/contacts_range_suggest"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/application/v6/applications/:app_id/app_versions/:version_id/contacts_range_suggest"
service: "application-v6"
resource: "application-app_version"
section: "App Information"
rate_limit: "100 per minute"
scopes:
  - "application:application:self_manage"
  - "application:application.app_version:readonly"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1682647793000"
---

# Get the Range of Contacts Data to Access in an App's Version Release Request

This interface is used to obtain the permission scope of the contact of a certain version of the enterprise's self-built application according to the application's App ID and version ID.

> Since the permission range of the contact needs to be submitted for release of a new application version, and will only take effect after being approved by the enterprise administrator, this permission range may be different from the actual effective permission range. If you need to get the actual online contacts range, you can get them through Get the Range of Contacts Data an App Can Access

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/application/v6/applications/:app_id/app_versions/:version_id/contacts_range_suggest |
| HTTP Method | GET |
| Rate Limit | 100 per minute |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `application:application:self_manage` `application:application.app_version:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `app_id` | `string` | App's app_id, get it from[Developer Console](https://open.larksuite.com/app) > Credentials & Basic Info. To query the version information of other apps, you must request the`application:application.app_version:readonly` scope. To only query the version information of this app, enter "me" or the app_id. **Example value**: "cli_9f3ca975326b501b" |
| `version_id` | `string` | ID that uniquely identifies the app version, get it fromObtain app version list **Example value**: "oav_d317f090b7258ad0372aa53963cda70d" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `department_id_type` | `string` | No | Type of the returned department ID **Example value**: "department_id" **Optional values are**:  - `department_id`: Identify the department with the custom department_id - `open_department_id`: Identify the department with open_department_id  **Default value**: `open_department_id` |
| `user_id_type` | `string` | No | User ID categories **Example value**: "open_id" **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `contacts_range` | `application.app_contacts_range` | App version contacts range suggest information. If the developer modifies the permission scope of the contact when submitting this version, the application permission scope of the contact will be returned. It does not represent the scope of contact permissions that will take effect in the final application. Empty if not modified. [If the permission range of the contact is consistent with the available range of the application, and the last configuration is also the same, it is considered that there is no change. ] |
|     `contacts_scope_type` | `string` | App version contacts range scope type **Optional values are**:  - `equal_to_availability`: Equal to availability, specific personnel can be queried by Obtain app availability in a company - `some`: Some, see visible_list for details - `all`: All  |
|     `visible_list` | `app_visible_list` | Visible list of contact scopes |
|       `open_ids` | `string[]` | List of open_ids of members to whom the app is visible |
|       `department_ids` | `string[]` | List of IDs of departments to which the app is visible |
|       `group_ids` | `string[]` | List of IDs of groups to which the app is visible | ### Response body example

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
        }
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 210503 | invalid app_id | Check whether the app_id in the request path is valid. |
| 400 | 210504 | no such app in tenant | Check whether the queried app and the app used to call the API are in the same company. |
| 400 | 210505 | target app not a custom app | Check whether the queried app is a custom app. |
| 400 | 210506 | no such app | Check whether the app_id in the request path exists. |
| 400 | 210508 | insufficient permission level | Check the scope you have requested and the queried app_id. This error code is returned when the queried app_id does not belong to this app and you have not requested the `application:application.app_version:readonly` scope. |
| 400 | 211003 | no such version of desired app | Check whether the version_id matches the app specified by the app_id. |
