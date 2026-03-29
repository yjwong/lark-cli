---
title: "Get Block  Access Data"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/workplace-v1/workplace_block_access_data/search"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/workplace/v1/workplace_block_access_data/search"
service: "workplace-v1"
resource: "workplace_block_access_data"
section: "Workplace"
rate_limit: "20 per second"
scopes:
  - "workplace:workplace_using_data:read"
updated: "1688436598000"
---

# Get Custom Workplace Block Access Data

Get Custom Workplace Block Access Data

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/workplace/v1/workplace_block_access_data/search |
| HTTP Method | POST |
| Rate Limit | 20 per second |
| Supported app types | custom,isv |
| Required scopes | `workplace:workplace_using_data:read` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `from_date` | `string` | Yes | **Example value**: 2023-02-01 |
| `to_date` | `string` | Yes | **Example value**: 2023-03-02 |
| `page_size` | `int` | Yes | **Example value**: 20 **Default value**: `20` **Data validation rules**: - Value range: `1` ～ `200` |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: ddowkdkl9w2d |
| `block_id` | `string` | No | BlockID. You can go to Lark Admin>Workplace>Custom Workplace, select the specified workplace and enter the Workplace Builder. Click on a block to view the "BlockID" below the block name in the right panel of the page **Example value**: 283438293839422334 | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `items` | `block_access_data[]` | Access data of block. |
|     `date` | `string` | Time, accurate to days, format yyyy MM dd |
|     `block_id` | `string` | BlockID. You can go to Management Background>Workplace>Custom Workplace, select the specified workplace and enter the Workplace Editor. Click on a block to view the "BlockID" below the widget name in the right panel of the page |
|     `access_data` | `access_data` | Access data of block |
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
                "block_id": "283438293839422334",
                "access_data": {
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
| 400 | 1629001 | invalid  time format | Pass in the correct time format [yyyy MM dd], eg.2023-05-22 |
| 400 | 1629002 | from_date can't be greater than to_date | The end time of data retrieval should be less than the start time |
| 400 | 1629005 | invalid page_token | Please use the page_token returned by the previous request |
| 400 | 1629006 | invalid page_size | Page size must be a positive integer |
| 400 | 1629004 | internal service error | Backend service exception or network exception, can be requested again |
