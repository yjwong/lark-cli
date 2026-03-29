---
title: "Delete Workforce Type"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/employee_type_enum/delete"
method: "DELETE"
api_path: "https://open.larksuite.com/open-apis/contact/v3/employee_type_enums/:enum_id"
service: "contact-v3"
resource: "employee_type_enum"
section: "Contacts"
scopes:
  - "contact:contact"
updated: "1663221577000"
---

# Delete workforce types

This API is used to delete custom workforce types.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/contact/v3/employee_type_enums/:enum_id |
| HTTP Method | DELETE |
| Supported app types | custom |
| Required scopes | `contact:contact` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `enum_id` | `string` | Enum ID **Example value**: "exGeIjow7zIqWMy+ONkFxA==" | ## Response

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
