---
title: "Cancel user migration"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/security_and_compliance-v1/user_migration/cancel"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/security_and_compliance/v1/user_migrations/cancel"
service: "security_and_compliance-v1"
resource: "user_migration"
section: "security_and_compliance"
scopes:
  - "security_and_compliance:user_migration"
  - "security_and_compliance:user_migration:multi-geo"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1752734936000"
---

# Cancel user migration

Cancel user migration. This operation can only be done for users whose migration has not been started .

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/security_and_compliance/v1/user_migrations/cancel |
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
| `user_ids` | `string[]` | Yes | List of User IDs. **Example value**: abcdefghi **Data validation rules**: - Length range: `1` ～ `100` | ### Request body example

{
    "user_ids": [
        "abcdefghi"
    ]
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
    "data": {

    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1781001 | Your request contains an invalid request parameter. | Parameter error. Check the input parameters according to the appropriate document. |
| 403 | 1781002 | Lack of necessary permissions. | Apply for Multi-Geo permission. |
| 400 | 1781003 | Multi-Geo is not enabled for this tenant. | Activate Multi-Geo service for your organization. |
| 500 | 1782001 | Internal service error. | Contact customer service for assistance. |
| 500 | 1782002 | User migration record not found. | Check if the user ID is correct. Contact customer service if in doubt. |
| 500 | 1783001 | User migration cannot be canceled. | The user has already been migrated to the destination geographic location. Contact customer service for help if the migration must be canceled. |
| 500 | 1782004 | The system or some other user is operating on the specified users. | Please try again later. |
