---
title: "Query Attendance Group by Name"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/group/search"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/attendance/v1/groups/search"
service: "attendance-v1"
resource: "group"
section: "Attendance"
scopes:
  - "attendance:rule"
  - "attendance:rule:readonly"
updated: "1646322852000"
---

# Query attendance group by name

Queries attendance group abstract by attendance group name. Available query methods include exact match and fuzzy match. The query results are sorted by attendance group modification time in descending order, and the maximum number of records is 10.

> There is a synchronization delay between the data on which this API is dependent and the master data of the attendance group (normal data synchronization is within 2 sec), so it is required to evaluate the potential risk of data delay when using this API.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/attendance/v1/groups/search |
| HTTP Method | POST |
| Supported app types | custom |
| Required scopes Enable any scope from the list | `attendance:rule` `attendance:rule:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `group_name` | `string` | Yes | Attendance group name **Example value**: "Attendance group 1" |
| `exactly_matched` | `boolean` | No | Whether it is an exact match; false: fuzzy match (default); true: exact match **Example value**: true | ### Request body example

```json
{
    "group_name": "Attendance group 1",
    "exactly_matched": true
}
```

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
| ∟ `group_list` | `group_meta[]` | Attendance group list |
| ∟ `group_id` | `string` | Attendance group ID |
| ∟ `group_name` | `string` | Attendance group name | ### Response body example

```json
{
    "code": 0,
    "msg": "success",
    "data": {
        "group_list": [
            {
                "group_id": "6919358128597097404",
                "group_name": "Attendance group 1"
            }
        ]
    }
}
```

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1220001 | Parameter error | Please check if the parameters meet the requirements. |
| 400 | 1220002 | Tenant doesn't exist. | Please check if the tenant_access_token is correct. |
