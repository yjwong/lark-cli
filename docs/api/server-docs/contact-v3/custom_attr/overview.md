---
title: "Overview"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/custom_attr/overview"
service: "contact-v3"
resource: "custom_attr"
section: "Contacts"
updated: "1646791352000"
---

#  Overview

##  Resource definition
User fields customized by a company, used to flexibly reflect user information. Custom user fields are created and updated by the [company administrator](https://www.larksuite.com/hc/en-US/articles/360048488029) in [Admin > Organization > Member Profile](https://www.larksuite.com/admin/contacts/employee-field-new/custom). You can obtain all custom fields of a company via the specific API.
  

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/d272ea8ae44d9c31b353e44186fbe2a9_hpUDdeNyqd.png?lazyload=true&width=1774&height=828?lazyload=true&width=1640&height=776)

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/fcf68d75fadfb5f50f267603976781a3_AVvcw3EwR9.png?lazyload=true&width=1773&height=829?lazyload=true&width=1640&height=933)

##  Field description
| Parameter | Type | Description |
| --- | --- | --- |
| `items` | `custom_attr[]` | Custom field definition |
| ∟ `id` | `string` | Custom field ID |
| ∟ `type` | `string` | Custom field type. Optional values are: - `TEXT`: Plain text, used to describe members in plain text, such as notes. - `HREF`: Static URL, used as a redirect link to member profile. - `ENUMERATION`: Enum, used to describe members in a structured way, such as ethnic group. - `GENERIC_USER`: User, used to describe the relationship between members, such as HRBP. - `PICTURE_ENUM`: Enum image, used to describe members with structured images, such as displaying badges of honor on the profile page. |
| ∟ `options` | `custom_attr_options` | Option definition. When type is `ENUMERATION` or `PICTURE_ENUM`, this field lists all options. |
| ∟ `default_option_id` | `string` | Default option ID |
| ∟ `option_type` | `string` | Option type **Optional values are**: - `TEXT`: Text - `PICTURE`: Image |
| ∟ `options` | `custom_attr_option[]` | Option list |
| ∟ `id` | `string` | Enum type option ID |
| ∟ `value` | `string` | Enum option value. When option_type is `TEXT`, it is a text value. When option_type is `PICTURE`, it is an image link. |
| ∟ `name` | `string` | Name. This is only valid when option_type is PICTURE. |
| ∟ `i18n_name` | `i18n_content[]` | Name of the custom field |
| ∟ `locale` | `string` | Language version |
| ∟ `value` | `string` | Field name |
| ∟ `page_token` | `string` | Paging token. A new page_token will be returned only when has_more is true. |
| ∟ `has_more` | `boolean` | Indicates whether there are more options | ### Data example
```json
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
]
```
