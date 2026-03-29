---
title: "Get OKR Period List"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/okr-v1/period/list"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/okr/v1/periods"
service: "okr-v1"
resource: "period"
section: "OKR"
scopes:
  - "okr:okr:readonly"
  - "okr:okr"
updated: "1669270473000"
---

# Get the OKR cycle list

Get the OKR cycle list

> ` tenant_access_token ` requires additional application permissions  to access OKR information as an application 

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/okr/v1/periods |
| HTTP Method | GET |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `okr:okr:readonly` `okr:okr` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `page_token` | `string` | No | Pagination flag page_token **Example value**: "xaasdasdax" |
| `page_size` | `int` | No | Page size, default 10 **Example value**: 10 **Default value**: `10` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `page_token` | `string` | Pagination flag |
|   `has_more` | `boolean` | Is there more |
|   `items` | `period[]` | Data item |
|     `id` | `string` | ID |
|     `zh_name` | `string` | Chinese name |
|     `en_name` | `string` | English name |
|     `status` | `int` | Enabled status **Optional values are**:  - `0`: Default - `1`: Normal - `2`: Invalid - `3`: Hidden  | ### Response body example

{
    "code": 0,
    "data": {
        "has_more": false,
        "items": [
            {
                "en_name": "Jan - Mar 2022",
                "id": "7071200999834255380",
                "status": 0,
                "zh_name": "2022 年 1 月 - 3 月"
            }
        ]
    },
    "msg": "success"
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 500 | 1009999 | internal server error | Internal error |
| 500 | 1009998 | system exception | System anomaly |
| 400 | 1001001 | invalid parameters | Invalid parameter |
| 400 | 1001002 | no permission | No permission |
| 400 | 1001003 | user not found | User does not exist |
| 400 | 1001004 | okr data not found | Okr data does not exist |
