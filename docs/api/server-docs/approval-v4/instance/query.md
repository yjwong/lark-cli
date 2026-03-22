---
title: "List of instances"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/instance/query"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/approval/v4/instances/query"
service: "approval-v4"
resource: "instance"
section: "Approval"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "approval:approval.list:readonly"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1747795866000"
---

# List of instances

The interface queries the list of qualified review instances in the review system through different conditions.

## Usage restrictions

- This API will filter out revoked approval instances, so there might be a situation where instances exist but no data is returned.

    Specific manifestation of this situation: the number of data entries per page in the returned result might be less than the page_size value. For example, if page_size is set to 10, but the current page shows only 6 entries, it indicates that 4 entries are revoked approval instances.

- The query results of this API might be delayed and cannot guarantee real-time results. For real-time queries, it is recommended to use Batch Get Approval Instance ID (this API can only query approval instance IDs under a specific approval definition).

- When querying:

    - user_id, approval_code, instance_code, instance_external_id, group_external_id cannot all be empty.

    - Query results for approval_code and group_external_id are a union set; query results for instance_code and instance_external_id are a union set; other query conditions are intersected.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/approval/v4/instances/query |
| HTTP Method | POST |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom |
| Required scopes | `approval:approval.list:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `page_size` | `int` | No | Page size: If the current page contains revoked approval instances, the number of data entries per page in the query result might be less than the page_size value. For example, if page_size is set to 10, but the current page shows only 6 entries, it indicates that 4 entries are revoked approval instances. **Example value**: 10 **Default value**: `10` **Data validation rules**: - Value range: `5` ～ `200` |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: nF1ZXJ5VGhlbkZldGNoCgAAAAAA6PZwFmUzSldvTC1yU |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id` | `string` | No | Fill in the user id according to the user_id_type **Example value**: "lwiu098wj" |
| `approval_code` | `string` | No | Approval Instance Code. Obtaining method: - Call the Create Approval Instance API and obtain the instance_code from the response parameters. - Call the Batch Get Approval Instance ID API to get the required approval instance code. **Note**: - user_id, approval_code, instance_code, instance_external_id, group_external_id cannot all be empty simultaneously. - Query results for instance_code and instance_external_id are a union set. **Example value**: "EB828003-9FFE-4B3F-AA50-2E199E2ED942" |
| `instance_code` | `string` | No | Approval Instance Code. Obtaining method: - Call the Create Approval Instance API and obtain the instance_code from the response parameters. - Call the Batch Get Approval Instance ID API to get the required approval instance code. **Note**: - user_id, approval_code, instance_code, instance_external_id, and group_external_id cannot all be empty simultaneously. - Query results for instance_code and instance_external_id are a union set. **Example value**: "EB828003-9FFE-4B3F-AA50-2E199E2ED943" |
| `instance_external_id` | `string` | No | Third-party ID of the approval instance. **Note**: - user_id, approval_code, instance_code, instance_external_id, and group_external_id cannot all be empty simultaneously. - Query results for instance_code and instance_external_id are a union set. **Example value**: "EB828003-9FFE-4B3F-AA50-2E199E2ED976" |
| `group_external_id` | `string` | No | Third-party ID of the approval definition group. **Note**: - user_id, approval_code, instance_code, instance_external_id, and group_external_id cannot all be empty simultaneously. - Query results for approval_code and group_external_id are a union set. **Example value**: "1234567" |
| `instance_title` | `string` | No | Approval instance title. **Note**: The approval instance title only exists for third-party approvals. **Example value**: "test" |
| `instance_status` | `string` | No | Approval instance status. **Default value**: ALL **Example value**: "PENDING" **Optional values are**:  - `PENDING`: Under review - `RECALL`: Withdraw - `REJECT`: Reject - `DELETED`: Deleted - `APPROVED`: Pass - `ALL`: All states  |
| `instance_start_time_from` | `string` | No | Instance query start time, Unix millisecond timestamp. Together with the `instance_start_time_to` parameter forms the time period query condition and will only return approval instances within this time period. **Note**: The query time span must not exceed 30 days, and the start and end times must either be both set or both unset. **Example value**: "1547654251506" |
| `instance_start_time_to` | `string` | No | Instance query end time, Unix millisecond timestamp. Together with the `instance_start_time_from` parameter forms the time period query condition and will only return approval instances within this time period. **Note**: The query time span must not exceed 30 days, and the start and end times must either be both set or both unset. **Example value**: "1547654251506" |
| `locale` | `string` | No | Language. **Example value**: "zh-CN" **Optional values are**:  - `zh-CN`: Chinese - `en-US`: English - `ja-JP`: Japanese  | ### Request body example

{
    "user_id": "lwiu098wj",
    "approval_code": "EB828003-9FFE-4B3F-AA50-2E199E2ED942",
    "instance_code": "EB828003-9FFE-4B3F-AA50-2E199E2ED943",
    "instance_external_id": "EB828003-9FFE-4B3F-AA50-2E199E2ED976",
    "group_external_id": "1234567",
    "instance_title": "test",
    "instance_status": "PENDING",
    "instance_start_time_from": "1547654251506",
    "instance_start_time_to": "1547654251506",
    "locale": "zh-CN"
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `count` | `int` | The number of items returned by the query |
|   `instance_list` | `instance_search_item[]` | List of review examples |
|     `approval` | `instance_search_approval` | Review Definition |
|       `code` | `string` | Review the code |
|       `name` | `string` | Review Definition Name |
|       `is_external` | `boolean` | Whether it is a third-party review |
|       `external` | `instance_search_approval_external` | Third party review information |
|         `batch_cc_read` | `boolean` | Whether to support batch reading |
|       `approval_id` | `string` | Approval Definition Id |
|       `icon` | `string` | Approval definition icon information |
|     `group` | `instance_search_group` | Review definition grouping |
|       `external_id` | `string` | Review definition group external id |
|       `name` | `string` | Review definition group name |
|     `instance` | `instance_search_node` | Review example information |
|       `code` | `string` | Review example code |
|       `external_id` | `string` | Review instance external id |
|       `user_id` | `string` | Review instance originator id |
|       `start_time` | `string` | Review instance start time |
|       `end_time` | `string` | Review instance end time |
|       `status` | `string` | Review instance status **Optional values are**:  - `rejected`: Reject - `pending`: Under review - `canceled`: Withdraw - `deleted`: Deleted - `approved`: Pass  |
|       `title` | `string` | Review instance name (only third-party reviews have it) |
|       `extra` | `string` | Review instance extension field, string type json |
|       `serial_id` | `string` | Review serial number |
|       `link` | `instance_search_link` | Review example link (only third-party reviews have) |
|         `pc_link` | `string` | Review example pc link |
|         `mobile_link` | `string` | Review example mobile end link |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned |
|   `has_more` | `boolean` | Whether the response body has more parameters | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "count": 10,
        "instance_list": [
            {
                "approval": {
                    "code": "EB828003-9FFE-4B3F-AA50-2E199E2ED943",
                    "name": "approval",
                    "is_external": true,
                    "external": {
                        "batch_cc_read": false
                    },
                    "approval_id": "7090754740375519252",
                    "icon": "https://lf3-ea.bytetos.com/obj/goofy/ee/approval/approval-admin/image/iconLib/v3/person.png"
                },
                "group": {
                    "external_id": "0004",
                    "name": "groupA"
                },
                "instance": {
                    "code": "EB828003-9FFE-4B3F-AA50-2E199E2ED943",
                    "external_id": "0004_3ED52DC1-AA6C",
                    "user_id": "lwiu098wj",
                    "start_time": "1547654251506",
                    "end_time": "1547654251506",
                    "status": "pending",
                    "title": "test",
                    "extra": "{}",
                    "serial_id": "201902020001",
                    "link": {
                        "pc_link": "https://www.baidu.com/",
                        "mobile_link": "https://www.baidu.com/"
                    }
                }
            }
        ],
        "page_token": "nF1ZXJ5VGhlbkZldGNoCgAAAAAA6PZwFmUzSldvTC1yU",
        "has_more": false
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1390001 | param is invalid | Parameter error. Troubleshooting steps: - Check the parameters passed during the request according to the parameter descriptions in the API documentation. - If form parameters are passed, verify whether the form control data within the parameters is correct. If the error message includes a control ID (e.g., `Control= widget17261088448220001`), you can call the View specific approval definition or Get single approval instance details API to retrieve the response parameter form value, search for the problematic control ID, and check whether the configuration of the control is correct. |
| 400 | 1390002 | approval code not found | Cannot find approval definition code. Check whether the passed approval definition code is correct. Methods to obtain the approval definition code: - After calling the Create approval definition API, get it from the response parameter approval_code. - Log into the approval management backend and obtain it from the URL of the specified approval definition. For detailed operations, refer to What is approval code. |
| 400 | 1390003 | instance code not found | Cannot find approval instance code. Check whether the passed approval instance code is correct. Methods to obtain the approval instance code: - After calling the Create approval instance API, get it from the response parameter instance_code. - Call the Batch Get approval instance ID API to obtain the required approval instance codes. - Call the Query instance list API, set the filter conditions to query the specified approval instance codes. |
| 400 | 1395001 | There have been some errors. Please try again later | Service error encountered. Troubleshooting steps: 1. Refer to the parameter descriptions in the API documentation and check whether the parameters passed in the request are correct. If form parameters are passed (form), ensure the form control data is correct. 2. Lower the request frequency and retry. If the error persists after retrying, please contact technical support. | For more error code information, see General error codes.
