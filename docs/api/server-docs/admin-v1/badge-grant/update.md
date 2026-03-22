---
title: "Modify the grant list"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/admin-v1/badge-grant/update"
method: "PUT"
api_path: "https://open.larksuite.com/open-apis/admin/v1/badges/:badge_id/grants/:grant_id"
service: "admin-v1"
resource: "badge-grant"
section: "Admin"
rate_limit: "100 per minute"
scopes:
  - "admin:badge.grant"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1686287168000"
---

# Modify the grant list

Through this API, information about a specific grant list can be modified.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/admin/v1/badges/:badge_id/grants/:grant_id |
| HTTP Method | PUT |
| Rate Limit | 100 per minute |
| Supported app types | custom |
| Required scopes | `admin:badge.grant` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `badge_id` | `string` | Badge ID **Example value**: "m_DjMzaK" **Data validation rules**: - Length range: `1` ～ `64` characters |
| `grant_id` | `string` | Grant List ID **Example value**: "g_uS4yux" **Data validation rules**: - Length range: `1` ～ `64` characters | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` |
| `department_id_type` | `string` | No | The type of department ID used in this call. **Example value**: open_department_id **Optional values are**:  - `department_id`: Identify departments with custom department_id. - `open_department_id`: Identify departments by open_department_id.  **Default value**: `open_department_id` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `name` | `string` | Yes | Grant list names, up to 100 characters. **Example value**: "Incentive badge grant list." |
| `grant_type` | `int` | Yes | Grant List Type **Example value**: 0 **Optional values are**:  - `0`: Manually select valid period - `1`: Matching system date of employment  **Default value**: `0` **Data validation rules**: - Value range: `0` ～ `1` |
| `time_zone` | `string` | Yes | The time zone corresponding to the effective time of the grant list, used to check whether the value of RuleDetail's timestamp is standardized, and the value range is TZ database name. **Example value**: "Asia/Shanghai" **Data validation rules**: - Minimum length: `1` characters |
| `rule_detail` | `rule_detail` | Yes | Rules Details |
|   `effective_time` | `string` | No | The timestamp that takes effect. 1. Manually set the valid period type badge, the configuration validity period needs to configure this field; 2. timestamp must be the zero point timestamp of the day in the time zone, such as 1649606400 when the time zone is Asia/Shanghai time zone. **Example value**: "1649606400" |
|   `expiration_time` | `string` | No | The timestamp that ends in effect. 1. Manually set the valid period type badge, the configuration validity period needs to configure this field; 2. Maximum value: no more than effective_time + 100 years; 3. Non-permanent valid: timestamp must be 23:59:59 timestamp on the day of the time zone, such as 1649692799 when the time zone is Asia/Shanghai time zone; 4. Permanent valid: pass the value as 0. **Example value**: "1649692799" |
|   `anniversary` | `int` | No | Onboarding anniversary. Depending on the date of employment, this field needs to be configured. **Example value**: 1 **Default value**: `1` **Data validation rules**: - Value range: `1` ～ `60` |
|   `effective_period` | `int` | No | Validity period. Depending on the date of employment, this field needs to be configured. **Example value**: 1 **Optional values are**:  - `1`: Valid period is one year - `2`: Permanently valid  **Default value**: `1` **Data validation rules**: - Value range: `1` ～ `2` |
| `is_grant_all` | `boolean` | Yes | Whether to grant to all employees. 1. When false, 1 to 500 client bases need to be associated. 2. When true, users, user groups, and departments cannot be associated. **Example value**: false **Default value**: `false` |
| `user_ids` | `string[]` | No | Granted user ID list, this field is not returned in the return result of the grant list list interface, but only returned in the detail API. **Example value**: ["u273y71"] |
| `department_ids` | `string[]` | No | List of granted department IDs, this field is not returned in the return result of the grant list list interface, but only returned in the detail API. **Example value**: ["h121921"] |
| `group_ids` | `string[]` | No | Granted user group ID list, this field is not returned in the return result of the grant list list interface, but only returned in the details API. **Example value**: ["g122817"] | ### Request body example

{
    "name": "Incentive badge grant list.",
    "grant_type": 0,
    "time_zone": "Asia/Shanghai",
    "rule_detail": {
        "effective_time": "1649606400",
        "expiration_time": "1649692799",
        "anniversary": 1,
        "effective_period": 1
    },
    "is_grant_all": false,
    "user_ids": [
        "u273y71"
    ],
    "department_ids": [
        "h121921"
    ],
    "group_ids": [
        "g122817"
    ]
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `grant` | `grant` | Grant List |
|     `id` | `string` | A unique identifier of a grant list within a tenant, randomly generated by the system. |
|     `badge_id` | `string` | Unique ID for the enterprise badge |
|     `name` | `string` | Grant list names, up to 100 characters. |
|     `grant_type` | `int` | Grant List Type **Optional values are**:  - `0`: Manually select valid period - `1`: Matching system date of employment  |
|     `time_zone` | `string` | The time zone corresponding to the effective time of the grant list, used to check whether the value of RuleDetail's timestamp is standardized, and the value range is TZ database name. |
|     `rule_detail` | `rule_detail` | Rules Details |
|       `effective_time` | `string` | The timestamp that takes effect. 1. Manually set the valid period type badge, the configuration validity period needs to configure this field; 2. timestamp must be the zero point timestamp of the day in the time zone, such as 1649606400 when the time zone is Asia/Shanghai time zone. |
|       `expiration_time` | `string` | The timestamp that ends in effect. 1. Manually set the valid period type badge, the configuration validity period needs to configure this field; 2. Maximum value: no more than effective_time + 100 years; 3. Non-permanent valid: timestamp must be 23:59:59 timestamp on the day of the time zone, such as 1649692799 when the time zone is Asia/Shanghai time zone; 4. Permanent valid: pass the value as 0. |
|       `anniversary` | `int` | Onboarding anniversary. Depending on the date of employment, this field needs to be configured. |
|       `effective_period` | `int` | Validity period. Depending on the date of employment, this field needs to be configured. **Optional values are**:  - `1`: Valid period is one year - `2`: Permanently valid  |
|     `is_grant_all` | `boolean` | Whether to grant to all employees. 1. When false, 1 to 500 client bases need to be associated. 2. When true, users, user groups, and departments cannot be associated. |
|     `user_ids` | `string[]` | Granted user ID list, this field is not returned in the return result of the grant list list interface, but only returned in the detail API. |
|     `department_ids` | `string[]` | List of granted department IDs, this field is not returned in the return result of the grant list list interface, but only returned in the detail API. |
|     `group_ids` | `string[]` | Granted user group ID list, this field is not returned in the return result of the grant list list interface, but only returned in the details API. | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "grant": {
            "id": "g_49Z7CQ",
            "badge_id": "m_qTR2HM",
            "name": "Incentive badge grant list.",
            "grant_type": 0,
            "time_zone": "Asia/Shanghai",
            "rule_detail": {
                "effective_time": "1649606400",
                "expiration_time": "1649692799",
                "anniversary": 1,
                "effective_period": 1
            },
            "is_grant_all": false,
            "user_ids": [
                "u273y71"
            ],
            "department_ids": [
                "h121921"
            ],
            "group_ids": [
                "g122817"
            ]
        }
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 500 | 1051000 | unknown server error | Service internal error, please try again later |
| 400 | 1051001 | request contain invalid param | Request contain invalid param |
| 403 | 1051002 | request to exceed authority | Request to exceed authority |
| 400 | 1051113 | cannot find this badge | No information found for this badge |
| 400 | 1051205 | cannot find this grant | No information found for this grant list |
| 400 | 1051200 | grant name are duplicated | Grant list name conflict |
| 400 | 1051201 | cannot exceed the max length limit of grant name | Grant list names longer than limit |
| 400 | 1051202 | already choose all staff, cannot related other user entity | Do not grant type list associated client base for all staff |
| 400 | 1051203 | modification of grant type is prohibited | Modification of grant list type is prohibited |
