---
title: "Delete User"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/delete"
method: "DELETE"
api_path: "https://open.larksuite.com/open-apis/contact/v3/users/:user_id"
service: "contact-v3"
resource: "user"
section: "Contacts"
scopes:
  - "contact:contact"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1731984590000"
---

# Delete a user

This API is used to delete a user from Contacts when an employee is terminated. FAQs.

> If the user belongs to departments A and B, the app's contact scope must include departments A and B so as to delete the user.
> 
> Users can specify the recipients of employee data (such as department groups, documents, schedules, etc.) when deleting employees. If not specified, the default processing logic will be triggered. The default processing logic for different data is different, see the description of each acceptor request parameter below for details

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/contact/v3/users/:user_id |
| HTTP Method | DELETE |
| Supported app types | custom |
| Required scopes | `contact:contact` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `user_id` | `string` | User ID, which must match the user_id_type in the query parameter. **Example value**: "ou_7dab8a3d3cdcc9da365777c7ad535d62" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: "open_id" **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `department_chat_acceptor_user_id` | `string` | No | Department group recipient. If the deleted user is the owner of a department group, the specified recipient will be assigned as the new owner. If no recipient is specified, the first member joining the group will be assigned as the new owner by default. **Example value**: "ou_7dab8a3d3cdcc9da365777c7ad535d62" |
| `external_chat_acceptor_user_id` | `string` | No | External group recipient. If the deleted user is the owner of an external group, the specified recipient will be assigned as the new owner. If no recipient is specified, the first member joining the group who is in the same organization as the deleted user will be assigned as the new owner by default. If the deleted user is the only one in the organization that is in the external group, the group will be deleted. **Example value**: "ou_7dab8a3d3cdcc9da365777c7ad535d62" |
| `docs_acceptor_user_id` | `string` | No | Document recipient. When a user is deleted, their documents will be transferred to the specified recipient. If no recipient is specified, the documents will be transferred to the user's direct leader by default. In case of no direct leader, the documents will be deleted directly. **Example value**: "ou_7dab8a3d3cdcc9da365777c7ad535d62" |
| `calendar_acceptor_user_id` | `string` | No | Event recipient. When a user is deleted, their events will be transferred to the specified recipient. If no recipient is specified, the events will be transferred by default to the user's direct leader. In case of no direct leader, the events will be deleted directly. **Example value**: "ou_7dab8a3d3cdcc9da365777c7ad535d62" |
| `application_acceptor_user_id` | `string` | No | App recipient. When a user is deleted, the apps they created will be transferred to the specified recipient. If no recipient is specified, the apps will be transferred to the user's direct leader by default. In case of no direct leader, the apps will not be transferred and will be unavailable **Example value**: "ou_7dab8a3d3cdcc9da365777c7ad535d62" |
| `minutes_acceptor_user_id` | `string` | No | Minute recipient's user ID. When a user is deleted, the Minute resources they own are transferred to the recipient. **Notice**: - The ID type needs to be consistent with the user_id_type type in the query parameter. For how to obtain user ID, please refer to How to obtain different user IDs. - If no recipient is specified, it will be transferred to the immediate superior of the deleted user by default. If the deleted user has no immediate superior, Minute will remain under the user name. - After the Minute is transferred, only the Minute owner is changed, and the shared Minute link is not affected. **Example value**: "ou_7dab8a3d3cdcc9da365777c7ad535d62" |
| `survey_acceptor_user_id` | `string` | No | The user ID of the survey recipient. When a user is deleted, the Survey resources they own are transferred to the recipient. **Notice**: - The ID type needs to be consistent with the user_id_type type in the query parameter. For how to obtain user ID, please refer to How to obtain different user IDs. - If no recipient is specified, it will be transferred to the immediate superior of the deleted user by default. If the deleted user has no immediate superior, the Survey resource will be deleted directly. **Example value**: "ou_7dab8a3d3cdcc9da365777c7ad535d62" |
| `email_acceptor` | `resource_acceptor` | No | User email resource processing method. This parameter is optional. If no value is passed, the email resources will be transferred to the deleted user's immediate superior by default. If the deleted user has no immediate superior, the email resources will be retained under the user's name. |
|   `processing_type` | `string` | Yes | Resource processing type. **Example value**: "1" **Optional values are**:  - `1`: Transfer Resource - `2`: Retain Resource - `3`: Delete Resource  |
|   `acceptor_user_id` | `string` | No | The user ID of the recipient of the mail resource. The ID type needs to be consistent with the user_id_type type in the query parameter. For how to obtain user ID, please refer to How to obtain different user IDs. **Note**: This parameter value needs to be set only when the value of `processing_type` is `1`. **Example value**: "ou_7dab8a3d3cdcc9da365777c7ad535d62" | ### Request body example

{
    "department_chat_acceptor_user_id": "ou_7dab8a3d3cdcc9da365777c7ad535d62",
    "external_chat_acceptor_user_id": "ou_7dab8a3d3cdcc9da365777c7ad535d62",
    "docs_acceptor_user_id": "ou_7dab8a3d3cdcc9da365777c7ad535d62",
    "calendar_acceptor_user_id": "ou_7dab8a3d3cdcc9da365777c7ad535d62",
    "application_acceptor_user_id": "ou_7dab8a3d3cdcc9da365777c7ad535d62",
    "minutes_acceptor_user_id": "ou_7dab8a3d3cdcc9da365777c7ad535d62",
    "survey_acceptor_user_id": "ou_7dab8a3d3cdcc9da365777c7ad535d62",
    "email_acceptor": {
        "processing_type": "1",
        "acceptor_user_id": "ou_7dab8a3d3cdcc9da365777c7ad535d62"
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
    "data": {

    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 40001 | param error | Parameter error. |
| 403 | 40004 | no dept authority error | The department must be within the range of contacts data that the app can access. Learn more |
| 403 | 41050 | no user authority error | The user must be within the range of contacts data that the app can access. Learn more. |
| 400 | 41052 | user resign acceptor is invalid error | Invalid user resignation approver. |
| 400 | 44037 | tenant manager cannot be deleted | Unable to delete tenant administrator. You need to go to the [Admin](https://www.larksuite.com/admin) > **Settings** > **Administrator Permissions** page, remove the user as an administrator and try again. |
| 400 | 44042 | User is in resurrect progress, retry later | The current user is in the resurrection process and cannot complete the operation. Please try again later. |
