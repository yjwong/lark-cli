---
title: "Delete Attendance Shift"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/shift/delete"
method: "DELETE"
api_path: "https://open.larksuite.com/open-apis/attendance/v1/shifts/:shift_id"
service: "attendance-v1"
resource: "shift"
section: "Attendance"
scopes:
  - "attendance:rule"
updated: "1647419328000"
---

# Delete shift

Deletes a shift via shift ID.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/attendance/v1/shifts/:shift_id |
| HTTP Method | DELETE |
| Supported app types | custom |
| Required scopes | `attendance:rule` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `shift_id` | `string` | Shift ID, which can be obtained as follows: 1) Search shift by name 2) Create a shift **Example value**: "6919358778597097404" | ## Response

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
| 500 | 1226000 | Shift service system error | See error message for details. |
| 500 | 1226001 | The shift has been used | Please change the shift name |
| 400 | 1226003 | The shift does not exist | Please check if the shift_id is correct |
