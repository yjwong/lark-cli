---
title: "Search Workplace Access Data"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/workplace-v1/workplace_access_data/search"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/workplace/v1/workplace_access_data/search"
service: "workplace-v1"
resource: "workplace_access_data"
section: "Workplace"
rate_limit: "20 per second"
scopes:
  - "workplace:workplace_using_data:read"
updated: "1688131605000"
---

# Get Workplace Access Data

Get Workplace Access Data, including default workplace and custom workplace

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/workplace/v1/workplace_access_data/search |
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
| `page_size` | `int` | Yes | **Example value**: 20 **Default value**: `20` **Data validation rules**: - Value range: `1` ～ `200` |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: ddowkdkl9w2d | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `items` | `workplace_access_data[]` | workplace access data |
|     `date` | `string` | Time, accurate to days, format yyyy MM dd |
|     `all_workplace` | `access_data` | Access data for all workplaces.Includes default workplace and custom workplace. Due to historical reasons, in some cases, the sum of these two data does not equal the access data of all workplaces. Please contact our technical support if you have any questions. |
|       `pv` | `int` | pv |
|       `uv` | `int` | uv |
|     `default_workplace` | `access_data` | Access data for default workplace |
|       `pv` | `int` | pv |
|       `uv` | `int` | uv |
|   `has_more` | `boolean` | Whether the response body has more parameters |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "items": [
            {
                "date": "2023-03-12",
                "all_workplace": {
                    "pv": 100,
                    "uv": 30
                },
                "default_workplace": {
                    "pv": 100,
                    "uv": 30
                }
            }
        ],
        "has_more": true,
        "page_token": "ddowkdkl9w2d"
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1628001 | invalid  time format | Pass in the correct time format [yyyy MM dd], eg.2023-05-22 |
| 400 | 1628002 | from_date can't be greater than to_date | The end time of data retrieval should be less than the start time |
| 400 | 1628003 | internal service error | Backend service exception or network exception, can be requested again |
| 400 | 1628004 | invalid page_token | Please use the page returned by the previous request API response's Token |
