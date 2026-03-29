---
title: "Obtain Custom User Fields"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/custom_attr/list"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/contact/v3/custom_attrs"
service: "contact-v3"
resource: "custom_attr"
section: "Contacts"
scopes:
  - "contact:contact:readonly_as_app"
updated: "1646791352000"
---

# Obtain custom user fields of a company

Obtain customer user field configuration of a company

> Before calling this API, make sure that the [company administrator](https://www.larksuite.com/hc/en-US/articles/360048488029) has enabled "Allow to invoke by Open Platform's Contacts API" in the Custom Fields section in [Admin > Organization > Member Profile](http://www.larksuite.com/admin/contacts/employee-field-new/custom).
> 
> ![image.gif](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/2322aa4cce5abe4f08087e22d13c3cbe_NMgO1puCYK.gif)

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/contact/v3/custom_attrs |
| HTTP Method | GET |
| Supported app types | custom,isv |
| Required scopes | `contact:contact:readonly_as_app` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `page_size` | `int` | No | Page size **Example value**: 10 **Data validation rules**: - Maximum value: `100` |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: "AQD9/Rn9eij9Pm39ED40/RYU5lvOM4s6zgbeeNNaWd%2BVKwAsoreeRWk0J2noGvJy" | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `items` | `custom_attr[]` | Custom field definition |
|     `id` | `string` | Custom field ID |
|     `type` | `string` | Custom field type. Optional values are: - `TEXT`: Plain text, used to describe members in plain text, such as notes. - `HREF`: Static URL, used as a redirect link to member profile. - `ENUMERATION`: Enum, used to describe members in a structured way, such as ethnic group. - `GENERIC_USER`: User, used to describe the relationship between members, such as HRBP. - `PICTURE_ENUM`: Enum image, used to describe members with structured images, such as displaying badges of honor on the profile page. |
|     `options` | `custom_attr_options` | Option definition. When type is `ENUMERATION` or `PICTURE_ENUM`, this field lists all options. |
|       `default_option_id` | `string` | Default option ID |
|       `option_type` | `string` | Option type **Optional values are**:  - `TEXT`: Text - `PICTURE`: Image  |
|       `options` | `custom_attr_option[]` | Option list |
|         `id` | `string` | Enum type option ID |
|         `value` | `string` | Enum option value. When option_type is `TEXT`, it is a text value. When option_type is `PICTURE`, it is an image link. |
|         `name` | `string` | Name. This is only valid when option_type is PICTURE. |
|     `i18n_name` | `i18n_content[]` | Name of the custom field |
|       `locale` | `string` | Language version |
|       `value` | `string` | Field name |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned |
|   `has_more` | `boolean` | Whether the response body has more parameters | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "items": [
            {
                "id": "C-6965457429001748507",
                "type": "TEXT",
                "options": {
                    "default_option_id": "qasdefgr",
                    "option_type": "TEXT",
                    "options": [
                        {
                            "id": "qasdefgr",
                            "value": "Option",
                            "name": "Name"
                        }
                    ]
                },
                "i18n_name": [
                    {
                        "locale": "zh_cn",
                        "value": "Expert"
                    }
                ]
            }
        ],
        "page_token": "AQD9/Rn9eij9Pm39ED40/RYU5lvOM4s6zgbeeNNaWd%2BVKwAsoreeRWk0J2noGvJy",
        "has_more": true
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 40012 | page token is invalid error | Page token is invalid. |
| 400 | 40011 | page size is invalid | The page size is between 1 and 50. |
