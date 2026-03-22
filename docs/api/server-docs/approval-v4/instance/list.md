---
title: "Obtain IDs of multiple approval instances"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/instance/list"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/approval/v4/instances"
service: "approval-v4"
resource: "instance"
section: "Approval"
scopes:
  - "approval:approval:readonly"
updated: "1675167402000"
---

# Obtain IDs of multiple approval instances

This API is used to obtain the instance_codes of multiple approval instances in batches based on a specified approval_code. This can be used to obtain all approval instances of a tenant's approval definition.The results are sorted by default in the order of when the approval was created.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/approval/v4/instances |
| HTTP Method | GET |
| Supported app types | custom,isv |
| Required scopes | `approval:approval:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `page_size` | `int` | No | Page size **Example value**: 100 **Default value**: `100` **Data validation rules**: - Maximum value: `100` |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: "nF1ZXJ5VGhlbkZldGNoCgAAAAAA6PZwFmUzSldvTC1yU" |
| `approval_code` | `string` | Yes | Unique identifier of approval definition **Example value**: "7C468A54-8745-2245-9675-08B7C63E7A85" |
| `start_time` | `string` | Yes | Approval instance creation time interval (in ms) **Example value**: "1567690398020" |
| `end_time` | `string` | Yes | Approval instance creation time interval (in ms) **Example value**: "1567690398020" | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `instance_code_list` | `string[]` | Approval instance code |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned |
|   `has_more` | `boolean` | Whether the response body has more parameters | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "instance_code_list": [
            "357C21A0-2069-4F6B-955F-1DFBE6710C51"
        ],
        "page_token": "nF1ZXJ5VGhlbkZldGNoCgAAAAAA6PZwFmUzSldvTC1yU",
        "has_more": false
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1390001 | param is invalid | Parameter error |
| 400 | 1390002 | approval code not found | Check that the review definition code is correct |
| 400 | 1395001 | There have been some errors. Please try again later | There was an error with the service, please try again later |
