---
title: "Get Custom Workplace Access Data"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/workplace-v1/custom_workplace_access_data/search"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/workplace/v1/custom_workplace_access_data/search"
service: "workplace-v1"
resource: "custom_workplace_access_data"
section: "Workplace"
rate_limit: "20 per second"
scopes:
  - "workplace:workplace_using_data:read"
updated: "1688045199000"
---

# Get Custom Workplace Access Data

Get Custom Workplace Access Data

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/workplace/v1/custom_workplace_access_data/search |
| HTTP Method | POST |
| Rate Limit | 20 per second |
| Supported app types | custom,isv |
| Required scopes | `workplace:workplace_using_data:read` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `from_date` | `string` | Yes | **Example value**: 2023-03-01 |
| `to_date` | `string` | Yes | **Example value**: 2023-03-22 |
| `page_size` | `int` | Yes | **Example value**: 20 **Default value**: `20` **Data validation rules**: - Value range: `1` ～ `100` |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: ddowkdkl9w2d |
| `custom_workplace_id` | `string` | No | Custom workplace ID,which is not mandatory. When empty, all custom workplace data is responsed. How to obtain a custom workplace ID: You can go to Lark Admin>Workplace>Custom Workplace, click on the settings of the specified workplace to enter the settings page; Click the "Settings" button at the top three times with the mouse to display the ID, and then copy the ID. **Example value**: tpl_647184b585400013254c4ea6 | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `items` | `custom_workplace_access_data[]` | Custom Workplace Access Data |
|     `custom_workplace_id` | `string` | Custom Workplace ID |
|     `access_data` | `access_data` | Access Data |
|       `pv` | `int` | PV |
|       `uv` | `int` | UV |
|     `date` | `string` | Time, accurate to days, format yyyy MM dd |
|     `custom_workplace_name` | `i18n_name[]` | Custom Workplace Name |
|       `language` | `string` | Language code of ISO 639-1. For example, zh represents Chinese. |
|       `name` | `string` | Name |
|   `has_more` | `boolean` | Whether the response body has more parameters |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "items": [
            {
                "custom_workplace_id": "tpl_645b003aaa40001435b2ddw",
                "access_data": {
                    "pv": 100,
                    "uv": 30
                },
                "date": "2023-03-12",
                "custom_workplace_name": [
                    {
                        "language": "zh",
                        "name": "Name"
                    }
                ]
            }
        ],
        "has_more": true,
        "page_token": "ddowkdkl9w2d"
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1621001 | invalid  time format | Pass in the correct time format [yyyy MM dd], eg.2023-05-22 |
| 400 | 1621002 | from_date can't be greater than to_date | The end time of data retrieval should be less than the start time |
| 400 | 1621003 | internal service error | Backend service exception or network exception, can be requested again |
| 400 | 1621004 | invalid page_token | Please use the page_token returned by the previous request |
