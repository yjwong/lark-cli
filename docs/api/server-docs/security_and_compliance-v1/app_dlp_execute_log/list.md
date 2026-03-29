---
title: "export app dlp execute log"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/security_and_compliance-v1/app_dlp_execute_log/list"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/security_and_compliance/v1/app_dlp_execute_logs"
service: "security_and_compliance-v1"
resource: "app_dlp_execute_log"
section: "security_and_compliance"
rate_limit: "100 per minute"
scopes:
  - "security_and_compliance:app_dlp_execute_log:read"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1695268535000"
---

# APP DLP Execute Log Export

Call this interface to export the execution log of the APP DLP, pageSize should not exceed 50

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/security_and_compliance/v1/app_dlp_execute_logs |
| HTTP Method | GET |
| Rate Limit | 100 per minute |
| Supported app types | custom,isv |
| Required scopes | `security_and_compliance:app_dlp_execute_log:read` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `page_size` | `int` | No | paging size **Example value**: 10 |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: LC39/f1% 2B/Sz9Uv39Gf39/eem/IpRCy1f4hD6Z7L1TOhnnM%2BFvbHra1Ql4URUft1kk2Oabl%2BrDSUvbDGKIrbjDQ== |
| `hit_policy_names` | `string[]` | No | Hit strategy list **Example value**: "Strategy A" |
| `start_time` | `string` | No | Start time **Example value**: 1688100723000 (ms level timestamp) |
| `end_time` | `string` | No | End time **Example value**: 1688100723000 (ms level timestamp) |
| `user_id` | `string` | No | User ID **Example value**: ou_ 7d8a6e6df7621556ce0d21922b676706ccs |
| `app_apply_module` | `int` | No | Application module **Example value**: 1 Apply dlp, currently only this value is supported |
| `apply_module` | `int` | No | Application range module **Example value**: 99 application The application range of dlp is 99, which means "all", and currently only this value is supported |
| `actions` | `int[]` | No | Execute action list **Example value**: 0: release 1: interception 2: timeout interception 3: abnormal interception 4: only record release 5: only record timeout interception 6: only record abnormal interception, when the action needs to be filtered, apply_module, app_apply_module need to ensure that there are values at the same time. |
| `locale` | `string` | No | locale **Example value**: zh-CN **Optional values are**:  - `en-US`: English - `zh-CN`: Chinese - `ja-JP`: Japanese  |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `items` | `app_dlp_execute_log[]` | Details of returned log data |
|     `event_name` | `string` | Event name |
|     `user_id` | `string` | User's open_id |
|     `execute_time` | `string` | execute time |
|     `action_name` | `string` | action name |
|     `hit_policies` | `string[]` | Hit strategy list |
|     `entity_id` | `string` | entity id |
|     `evidences` | `dlp_execute_evidence` | evidence |
|       `keyword_hits` | `string[]` | keyword |
|       `regular_hits` | `string[]` | regular list |
|       `sensitive_hits` | `string[]` | sensitive hits |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned |
|   `has_more` | `boolean` | Whether the response body has more parameters | ### Response body example

{
    "code": 0,
    "msg": "Success",
    "data": {
        "items": [
            {
                "event_name": "MyAIDLP",
                "user_id": "ou_ 04444ce34fe049fb495d150dddeac36a",
                "execute_time": "1689685605",
                "action_name": "reject",
                "hit_policies": [
                    "Policy 1"
                ],
                "entity_id": "7d8a6e6df7621556ce0d21922b676706ccs",
                "evidences": {
                    "keyword_hits": [
                        "keyword1"
                    ],
                    "regular_hits": [
                        "reqular A"
                    ],
                    "sensitive_hits": [
                        "ID Card"
                    ]
                }
            }
        ],
        "page_token": "LC39/f1% 2B/Sz9Uv39Gf39/eem/IpRCy1f4hD6Z7L1TOhnnM%2BFvbHra1Ql4URUft1kk2Oabl%2BrDSUvbDGKIrbjDQ==",
        "has_more": false
    }
}

Response successful

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 200 | 1781000 | param is invalid,please check your param | The request parameter is wrong, please check the parameter configuration. Millsecond-level timestamps are required for the start and end; the requested pageSize cannot exceed 50; action needs to be passed in together with module and applyModule; |
| 500 | 1781001 | internal server error | Internal interface processing failed, please check the parameter configuration. Millsecond-level timestamps are required for the start and end; the requested pageSize cannot exceed 50; action needs to be passed in together with module and applyModule;if it not work,please seek help from customer service |
