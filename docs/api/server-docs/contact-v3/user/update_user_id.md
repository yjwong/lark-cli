---
title: "Update UserID"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/update_user_id"
method: "PATCH"
api_path: "https://open.larksuite.com/open-apis/contact/v3/users/:user_id/update_user_id"
service: "contact-v3"
resource: "user"
section: "Contacts"
rate_limit: "20 per second"
scopes:
  - "contact:contact:update_user_id"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1732603268000"
---

# Update User ID

Call this interface to update the user's user_id.

## Precautions

The updated user user_id needs to ensure that it is not occupied in the current tenant.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/contact/v3/users/:user_id/update_user_id |
| HTTP Method | PATCH |
| Rate Limit | 20 per second |
| Supported app types | custom |
| Required scopes | `contact:contact:update_user_id` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `user_id` | `string` | The user ID and ID type are consistent with the value of the query parameter user_id_type. **Example value**: "ou-938e3e4fdc5e1993bee01250076f0cc2" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `new_user_id` | `string` | Yes | Customize new user user_id. Length cannot exceed 64 characters. **Example value**: "3e3cf96b" | ### Request body example

{
    "new_user_id": "3e3cf96b"
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
| 400 | 40001 | param error | Parameter error. Please refer to the parameter description in the interface document to check whether the input parameters are filled in correctly. |
| 400 | 41009 | no email or mobile error | No email or mobile phone number is found, please add it and try again. |
| 400 | 41011 | userID already exist error | The user's user_id already exists, you need to change the ID and try again. |
| 400 | 41013 | exceed userID update limit error | User ID update limit exceeded. |
| 400 | 41050 | no user authority error | No user permissions. The user currently operating must be within the address book permissions of the app. For the introduction and setting method of address book permission scope, please refer to Permission Scope Resource Introduction. |
| 400 | 41054 | new user id is required  error | Error in updated user_id. Please change the user_id value and try again. | For more error code information, see General Error Codes.
