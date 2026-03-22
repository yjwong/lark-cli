---
title: "Obtain Mailing Lists In Batch"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup/list"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/mail/v1/mailgroups"
service: "mail-v1"
resource: "mailgroup"
section: "Email"
rate_limit: "Special Rate Limit"
scopes:
  - "mail:mailgroup"
  - "mail:mailgroup:readonly"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1745840340000"
---

# Obtain mailing lists in batch

Obtains mailing lists by pages.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/mail/v1/mailgroups |
| HTTP Method | GET |
| Rate Limit | Special Rate Limit |
| Supported app types | custom |
| Required scopes Enable any scope from the list | `mail:mailgroup` `mail:mailgroup:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `manager_user_id` | `string` | No | User ID of the mailing list administrator, which is used to obtain the mailing lists for which the user has the management permission. **Example value**: ou_xxxxxx |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: xxx |
| `page_size` | `int` | No | **Example value**: 10 **Default value**: `20` **Data validation rules**: - Maximum value: `200` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `has_more` | `boolean` | Whether the response body has more parameters |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned |
|   `items` | `mailgroup[]` | Mailing lists |
|     `mailgroup_id` | `string` | Mailing list ID |
|     `email` | `string` | Mailing list address |
|     `name` | `string` | Mailing list name |
|     `description` | `string` | Mailing list description |
|     `direct_members_count` | `string` | Number of mailing list members |
|     `include_external_member` | `boolean` | Whether it includes external members |
|     `include_all_company_member` | `boolean` | Whether it is an all-staff mailing list |
|     `who_can_send_mail` | `string` | Who can send emails to this mailing list **Optional values are**:  - `ANYONE`: Anyone - `ALL_INTERNAL_USERS`: Organization members only - `ALL_GROUP_MEMBERS`: Mailing list members only - `CUSTOM_MEMBERS`: Specified members  | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "has_more": true,
        "page_token": "xxx",
        "items": [
            {
                "mailgroup_id": "xxxxxxxxxxxxxxx",
                "email": "test_mail_group@xxx.xx",
                "name": "test mail group",
                "description": "mail group for testing",
                "direct_members_count": "10",
                "include_external_member": true,
                "include_all_company_member": false,
                "who_can_send_mail": "ALL_INTERNAL_USERS"
            }
        ]
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 503 | 1235003 | Service unavailable | Please try again later. |
