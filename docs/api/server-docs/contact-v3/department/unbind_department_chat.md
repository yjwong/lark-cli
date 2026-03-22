---
title: "change department group to common group"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/unbind_department_chat"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/contact/v3/departments/unbind_department_chat"
service: "contact-v3"
resource: "department"
section: "Contacts"
scopes:
  - "contact:contact"
updated: "1683879166000"
---

# Change department group to common group

Use this api to change department group to common group.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/contact/v3/departments/unbind_department_chat |
| HTTP Method | POST |
| Supported app types | custom |
| Required scopes | `contact:contact` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `department_id_type` | `string` | No | The type of department ID used in this call **Example value**: "open_department_id" **Optional values are**:  - `department_id`: Identify the department with a custom department_id - `open_department_id`: Identify the department with open_department_id  | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `department_id` | `string` | Yes | ID of the parent department **Example value**: "od-4e6ac4d14bcd5071a37a39de902c7141" | ### Request body example

{
    "department_id": "od-4e6ac4d14bcd5071a37a39de902c7141"
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
| 400 | 40002 | process root dept error | Root departments cannot deleted. |
| 400 | 40003 | internal error | Internal error. Please provide the X-Request-Id to our agent. [Contact support](https://applink.larksuite.com/client/helpdesk/open?id |
| 403 | 40004 | no dept authority error | The department must be within the range of contacts data that the app can access. Learn more |
| 403 | 40020 | department do not has a chat error | The department do no has a department group. |
