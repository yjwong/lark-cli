---
title: "Delete employee"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/directory-v1/employee/delete"
method: "DELETE"
api_path: "https://open.larksuite.com/open-apis/directory/v1/employees/:employee_id"
service: "directory-v1"
resource: "employee"
section: "Organization"
rate_limit: "10 per second"
scopes:
  - "directory:employee.resign:write"
  - "directory:employee:write"
field_scopes:
  - "directory:employee.base.external_id:read"
updated: "1756439253000"
---

# Resign employee

This interface is for resign employees.

> Attention:
> - This interface supports tenant_access_token and user_access_token.
> - When using tenant_access_token, employees can only leave within the scope of the address book authorization currently applied.
> - If an employee belongs to multiple departments, the app needs to have the authority of all departments to which the employee belongs in order to leave successfully.
> - When using user_access_token, the default is the administrator user, and the administrator management scope will be verified. When the user has multiple administrator identities and can leave employees, the administrator management scope takes the maximum set.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/directory/v1/employees/:employee_id |
| HTTP Method | DELETE |
| Rate Limit | 10 per second |
| Supported app types | custom |
| Required scopes Enable any scope from the list | `directory:employee.resign:write` `directory:employee:write` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `directory:employee.base.external_id:read` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `employee_id` | `string` | The ID of the employee to leave. Needs to be consistent with the employee_id_type type in the query parameter **Example value**: "eesdasjd" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `employee_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies the identity of a user in an application. The same user has different Open IDs in different applications. Learn more: How to get an Open ID - `union_id`: Identifies the identity of a user under an application developer. The Union ID of the same user in an application under the same developer is the same, and the Union ID in an application under different developers is different. With Union ID, application developers can associate the identity of the same user in multiple applications. Learn more: How to get Union ID? - `employee_id`: The unique identifier of an employee within the enterprise. Supports customization, the system automatically generates it if it is not customized. ID supports modification. How to get employee_id: - Enterprise administrators go to the Admin Console > Organizational Structure > Members and Departments page, click Member Details to query employee IDs - Through the interface of Get employee list in batches, query employee ID by mobile phone number or email.  **Default value**: `open_id` **When the value is `employee_id`, the following field scopes are required**: `directory:employee.base.external_id:read` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `options` | `delete_employee_options` | No | interface extension options |
|   `resigned_employee_resource_receiver` | `resigned_user_resouce_receiver` | No | Resource transfer methods for departing employees. |
|     `department_chat_acceptor_employee_id` | `string` | No | Department group receiver. The ID value corresponds to the employee_id_type in the query parameter. When the deleted user is the group owner of the department group, the group owner is transferred to the designated recipient, and if the recipient is not specified, it is transferred to the first person in the group by default. **Example value**: "ou_ 7dab8a3d3cdcc9da365777c7ad535d62" |
|     `external_chat_acceptor_employee_id` | `string` | No | External group receiver. The ID value corresponds to the employee_id_type in the query parameter. When the deleted user is an external group owner, the group owner is transferred to the designated recipient. If the recipient is not specified, it will be transferred to the first person in the group who joins the group in the same organization as the deleted user. If only the user is in the group, the external group will be disbanded. **Example value**: "eehsdna" |
|     `docs_acceptor_employee_id` | `string` | No | Document recipient. The ID value corresponds to the employee_id_type in the query parameter. When a user is deleted, the documents owned by the user are transferred to the recipient. If no recipient is specified, the document resources will be retained under the user name. **Example value**: "EeehsDNA" |
|     `calendar_acceptor_employee_id` | `string` | No | Schedule recipient. The ID value corresponds to the employee_id_type in the query parameter. When a user is deleted, the schedules owned by the user are transferred to the recipient. If no recipient is specified, the schedule resources will be retained under the user name. **Example value**: "eehsdna" |
|     `application_acceptor_employee_id` | `string` | No | Apply the recipient. The ID value corresponds to the employee_id_type in the query parameter. When a user is deleted, the applications created by the user are transferred to the recipient. If the recipient is not specified, the application will be retained under the user name, but the user cannot log in to the developer backend to manage the application. The administrator can manually transfer the application to other people in the management backend. **Example value**: "eehsdna" |
|     `helpdesk_acceptor_employee_id` | `string` | No | Help desk resource receiver. The ID value corresponds to the employee_id_type in the query parameter. When a user is deleted, the help desk resources they own are transferred to the recipient, and when the recipient is not specified, the help desk resources are retained under the user's name. **Example value**: "eehsdna" |
|     `approval_acceptor_employee_id` | `string` | No | Approval of resource recipients. ID values correspond to employee_id_type in query parameters. When a user is deleted, the approval resources it owns are transferred to the recipient, and the approval resources are retained under the user's name when the recipient is not specified. **Example value**: "eehsdna" |
|     `email_acceptor_employee_id` | `string` | No | User mail resource recipient. The ID value corresponds to the employee_id_type in the query parameter. When a user is deleted, the email resources owned by the user are transferred to the recipient. If the recipient is not specified, the email resources are retained under the user name. **Example value**: "eehsdna" |
|     `minutes_acceptor_employee_id` | `string` | No | Wonder receiver. The ID value corresponds to the employee_id_type in the query parameter. When a user is deleted, the minutes resources owned by him are transferred to the receiver. If the receiver is not specified, the minutes will be retained under the user name. **Example value**: "eehsdna" |
|     `survey_acceptor_employee_id` | `string` | No | Lark Survey recipient. The ID value corresponds to the employee_id_type in the query parameters. When a user is deleted, the Lark survey resources owned by him are transferred to the recipient. If no recipient is specified, the Lark survey resources will be deleted directly. **Example value**: "eehsdna" |
|     `anycross_acceptor_employee_id` | `string` | No | Anycross recipient **Example value**: "eehsdna" | ### Request body example

{
    "options": {
        "resigned_employee_resource_receiver": {
            "department_chat_acceptor_employee_id": "ou_ 7dab8a3d3cdcc9da365777c7ad535d62",
            "external_chat_acceptor_employee_id": "eehsdna",
            "docs_acceptor_employee_id": "EeehsDNA",
            "calendar_acceptor_employee_id": "eehsdna",
            "application_acceptor_employee_id": "eehsdna",
            "helpdesk_acceptor_employee_id": "eehsdna",
            "approval_acceptor_employee_id": "eehsdna",
            "email_acceptor_employee_id": "eehsdna",
            "minutes_acceptor_employee_id": "eehsdna",
            "survey_acceptor_employee_id": "eehsdna",
            "anycross_acceptor_employee_id": "eehsdna"
        }
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
    "data": {}
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 2221183 | Unable to delete tenant founder | Unable to delete tenant creator |
| 400 | 2221185 | Frequent deletions. User recovered less than an hour ago | Delete frequently. Users recovered less than an hour ago, please try again in an hour |
| 400 | 2224001 | No permission to operate | No operation permission. Please check the permission of the current application or whether the enterprise version is the business professional version or above. |
| 400 | 2224002 | No permission to operate records | You do not have permission to operate this record. Please check the permissions of the data management scope of the current application or whether the members operated by the current application can be restored. |
| 400 | 2224003 | No permission to operate dependent objects | No operation permission for the dependent object. Please check whether the department to be restored has the permission. |
