---
title: "Verify instances"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/external_instance/check"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/approval/v4/external_instances/check"
service: "approval-v4"
resource: "external_instance"
section: "Approval"
scopes:
  - "approval:approval:readonly"
updated: "1675167421000"
---

# Verify instances

Third-party approval instance validation is used to determine whether the server-side data is up-to-date. The user submits the last update time of the instance, and if the instance doesn't exist on the server-side, or the update time of the instance on the server-side isn't the latest, the corresponding instance ID is returned.

For example, users can compare the instances generated in the last 5 min using this API every 5 min.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/approval/v4/external_instances/check |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes | `approval:approval:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `instances` | `exteranl_instance_check[]` | Yes | Instance information |
|   `instance_id` | `string` | Yes | Approval instance ID **Example value**: "1234234234242423" |
|   `update_time` | `string` | Yes | Last update time of approval instance **Example value**: "1591603040000" |
|   `tasks` | `external_instance_task[]` | Yes | Task information |
|     `task_id` | `string` | Yes | Task id **Example value**: "112253" |
|     `update_time` | `string` | Yes | Last update time of task **Example value**: "1591603040000" | ### Request body example

{
    "instances": [
        {
            "instance_id": "1234234234242423",
            "update_time": "1591603040000",
            "tasks": [
                {
                    "task_id": "112253",
                    "update_time": "1591603040000"
                }
            ]
        }
    ]
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `diff_instances` | `exteranl_instance_check_response[]` | Instance information with inconsistent update times |
|     `instance_id` | `string` | Approval instance ID |
|     `update_time` | `string` | Last update time of approval instance |
|     `tasks` | `external_instance_task[]` | Task information |
|       `task_id` | `string` | Task id |
|       `update_time` | `string` | Last update time of task | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "diff_instances": [
            {
                "instance_id": "1234234234242423",
                "update_time": "1591603040000",
                "tasks": [
                    {
                        "task_id": "112253",
                        "update_time": "1591603040000"
                    }
                ]
            }
        ]
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1390001 | param is invalid | Parameter error |
| 400 | 1395001 | There have been some errors. Please try again later | There was an error with the service, please try again later |
