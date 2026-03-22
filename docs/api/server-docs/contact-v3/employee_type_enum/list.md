---
title: "Fetch Workforce Type"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/employee_type_enum/list"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/contact/v3/employee_type_enums"
service: "contact-v3"
resource: "employee_type_enum"
section: "Contacts"
scopes:
  - "contact:contact:readonly_as_app"
  - "contact:contact:access_as_app"
updated: "1646791352000"
---

# Query the workforce type

This API is used to obtain an employee's workforce type.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/contact/v3/employee_type_enums |
| HTTP Method | GET |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `contact:contact:readonly_as_app` `contact:contact:access_as_app` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: "3" |
| `page_size` | `int` | No | Page size **Example value**: 10 **Data validation rules**: - Maximum value: `100` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `items` | `employee_type_enum[]` | Enum data |
|     `enum_id` | `string` | Enum ID |
|     `enum_value` | `string` | Enum value, which is automatically generated for a newly created workforce type. This field corresponds to the employee_type field in user information in the API for creating users. |
|     `content` | `string` | Enum content |
|     `enum_type` | `int` | Type **Optional values are**:  - `1`: Built-in - `2`: Custom  |
|     `enum_status` | `int` | Status **Optional values are**:  - `1`: Activated - `2`: Unactivated  |
|     `i18n_content` | `i18n_content[]` | Internationalization definition |
|       `locale` | `string` | Language version |
|       `value` | `string` | Field name |
|   `has_more` | `boolean` | Whether the response body has more parameters |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "items": [
            {
                "enum_id": "exGeIjow7zIqWMy+ONkFxA

contact.v3.type.i18n_content.prop.value.string.example=$$$Expert",
                "enum_value": "2",
                "content": "Expert",
                "enum_type": 2,
                "enum_status": 1,
                "i18n_content": [
                    {
                        "locale": "zh_cn",
                        "value": "Expert"
                    }
                ]
            }
        ],
        "has_more": true,
        "page_token": "3"
    }
}
