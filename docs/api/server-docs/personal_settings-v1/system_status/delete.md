---
title: "Delete system status"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/personal_settings-v1/system_status/delete"
method: "DELETE"
api_path: "https://open.larksuite.com/open-apis/personal_settings/v1/system_statuses/:system_status_id"
service: "personal_settings-v1"
resource: "system_status"
section: "Personal Settings"
scopes:
  - "personal_settings:status:system_status_update"
updated: "1672726533000"
---

# Delete system status

Delete a system status of the full-tenant dimension.

> Notice:
> - The data to be operated is tenant dimension data, please operate with care.
> - After deleting the system state, it does not affect the display of the system state of the user who is using the state.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/personal_settings/v1/system_statuses/:system_status_id |
| HTTP Method | DELETE |
| Supported app types | custom |
| Required scopes | `personal_settings:status:system_status_update` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `system_status_id` | `string` | System stauts ID Get system status ID **Example value**: "7101214603622940633" | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {

    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 2005001 | Your request contains an invalid request parameter. | The parameter is incorrect. Please check the input parameter according to the error message returned by the interface and refer to the documentation. |
| 400 | 2005007 | Tenant does not have permission to api. | Tenant does not have permission to api, Please apply to use. |
