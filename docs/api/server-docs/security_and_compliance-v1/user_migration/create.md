---
title: "Migrate users"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/security_and_compliance-v1/user_migration/create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/security_and_compliance/v1/user_migrations"
service: "security_and_compliance-v1"
resource: "user_migration"
section: "security_and_compliance"
scopes:
  - "security_and_compliance:user_migration"
  - "security_and_compliance:user_migration:multi-geo"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1752734935000"
---

# Migrate users

Migrate users' data to  the specified geographic location.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/security_and_compliance/v1/user_migrations |
| HTTP Method | POST |
| Supported app types | custom |
| Required scopes Enable any scope from the list | `security_and_compliance:user_migration` `security_and_compliance:user_migration:multi-geo` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | Yes | User ID categories **Example value**: "open_id" **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_ids` | `string[]` | Yes | List of user IDs of the users to be migrated. **Example value**: abcdefghi **Data validation rules**: - Length range: `1` ～ `100` |
| `dest_geo` | `string` | Yes | Destination geographic location. **Example value**: "sg" | ### Request body example

{
    "user_ids": [
        "abcdefghi"
    ],
    "dest_geo": "sg"
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `user_migrations` | `user_migration[]` | User migration info. |
|     `user_id` | `string` | User ID |
|     `dest_geo` | `string` | Destination geographic location. |
|     `task_id` | `string` | Migration task ID. |
|     `status` | `string` | Migration status. **Optional values are**:  - `0`: Migration is in progress. - `1`: Migration has been completed. - `2`: Migration has been canceled.  |
|     `progress` | `int` | Migration progress. | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "user_migrations": [
            {
                "user_id": "abcdefghi",
                "dest_geo": "sg",
                "task_id": "jklmnopq",
                "status": "0",
                "progress": 30
            }
        ]
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1781001 | Your request contains an invalid request parameter. | Parameter error. Check the input parameters according to the appropriate document. |
| 403 | 1781002 | Lack of necessary permissions. | Apply for Multi-Geo permission. |
| 400 | 1781003 | Multi-Geo is not enabled for this tenant. | Activate Multi-Geo service for your organization. |
| 400 | 1781004 | Destination geographic location is not available for this tenant. | Use open API to get a list of geographic locations available to your organization. |
| 400 | 1781005 | Users have already been migrated to the specified geographic location. | Check the user's current geographic location or migrate the user to a different location. |
| 400 | 1781006 | Users are currently being migrated. | Use relevant open API to check the users' migration status. |
| 500 | 1782001 | Internal service error. | Contact customer service for assistance. |
| 500 | 1782004 | The system or some other user is operating on the specified users. | Please try again later. |
| 400 | 1782005 | Request failed. The number of users to be migrated exceeds the seat limit of the data residency service. | Please contact Customer Service to purchase additional seats as needed, then try again. |
