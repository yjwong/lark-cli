---
title: "Create Workforce Type"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/employee_type_enum/create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/contact/v3/employee_type_enums"
service: "contact-v3"
resource: "employee_type_enum"
section: "Contacts"
scopes:
  - "contact:contact"
updated: "1661854549000"
---

# Add a workforce type

This API is used to add a workforce type.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/contact/v3/employee_type_enums |
| HTTP Method | POST |
| Supported app types | custom |
| Required scopes | `contact:contact` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `content` | `string` | Yes | Enum content **Example value**: "Expert" **Data validation rules**: - Length range: `1` ～ `100` characters |
| `enum_type` | `int` | Yes | Type **Example value**: 2 **Optional values are**:  - `1`: Built-in - `2`: Custom  |
| `enum_status` | `int` | Yes | Status **Example value**: 1 **Optional values are**:  - `1`: Activated - `2`: Unactivated  |
| `i18n_content` | `i18n_content[]` | No | Internationalization definition |
|   `locale` | `string` | No | Language version **Example value**: "zh_cn" |
|   `value` | `string` | No | Field name **Example value**: "Expert" | ### Request body example

{
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

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `employee_type_enum` | `employee_type_enum` | New workforce type |
|     `enum_id` | `string` | Enum ID |
|     `enum_value` | `string` | Enum value, which is automatically generated for a newly created workforce type. This field corresponds to the employee_type field in user information in the API for creating users. |
|     `content` | `string` | Enum content |
|     `enum_type` | `int` | Type **Optional values are**:  - `1`: Built-in - `2`: Custom  |
|     `enum_status` | `int` | Status **Optional values are**:  - `1`: Activated - `2`: Unactivated  |
|     `i18n_content` | `i18n_content[]` | Internationalization definition |
|       `locale` | `string` | Language version |
|       `value` | `string` | Field name | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "employee_type_enum": {
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
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 42301 | param content duplicate | "content" must be unique. |
| 400 | 42302 | param i18n_content duplicate | "i18n_content" must be unique. |
