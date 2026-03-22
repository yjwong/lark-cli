---
title: "Delete Attendance Group"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/group/delete"
method: "DELETE"
api_path: "https://open.larksuite.com/open-apis/attendance/v1/groups/:group_id"
service: "attendance-v1"
resource: "group"
section: "Attendance"
scopes:
  - "attendance:rule"
updated: "1647419328000"
---

# Delete attendance group

Deletes a shift via shift ID.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/attendance/v1/groups/:group_id |
| HTTP Method | DELETE |
| Supported app types | custom |
| Required scopes | `attendance:rule` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `group_id` | `string` | Attendance group ID is obtained as follows: 1) Create or modify attendance groups 2) Query attendance group by name 3) Obtain attendance results **Example value**: "6919358128597097404" | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- | ### Response body example

```json
{
    "code": 0,
    "msg": "success",
    "data": {

    }
}
```

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1220001 | Parameter error | Please check if the parameters meet the requirements. |
| 400 | 1220002 | Tenant doesn't exist. | Please check if the tenant_access_token is correct. |
| 400 | 1220005 | No scope | Please go to [Attendance Admin](https://oa.larksuite.com/attendance/manage/member/list) to check permission scope on data |
| 500 | 1225000 | System error | See error information for details. |
| 500 | 1227000 | Management service system error | See error information for details. |
