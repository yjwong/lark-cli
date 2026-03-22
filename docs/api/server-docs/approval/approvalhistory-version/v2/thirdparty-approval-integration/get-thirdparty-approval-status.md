---
title: "Get Third-party Approval Status"
url: "https://open.larksuite.com/document/ukTMukTMukTM/ukjNyYjL5YjM24SO2IjN/external_status"
method: "POST"
api_path: "https://www.larksuite.com/approval/openapi/v2/external/list"
section: "Approval"
scopes:
  - "approval:approval:readonly"
updated: "1658826314000"
---

# Obtain third-party approval status
> This API has been upgraded to enhance its security. We recommend that you migrate to the new version>>

This API is used to obtain the third-party approval status. After a user enters the query condition, the API returns the status of the approval instances that meet the condition.
The API supports a combination of multiple parameters,  including the following:
 1. Obtain the task status of a specified instance by using instance_ids
 2. Obtain the task status of a specified user by using user_ids
 3. Obtain all the tasks of a specified status by using status
 4. Obtain the next batch of data

##  Request
| HTTP URL | https://www.larksuite.com/approval/openapi/v2/external/list |
| --- | --- |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes | `approval:approval:readonly` | ###  Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about how to obtain and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body

|Parameter|Type|Required|Description|
|-|-|-|-|
|approval_codes|list|No|Approval definition codes, which are used to specify that only the data under these definitions will be obtained.|
|instance_ids|list|No|Approval instance IDs, which are used to specify that only the data under these instances will be obtained. Up to 20  IDs are supported.|
|user_ids|list|No|Approver's user_ids, which are used to specify that only the data of these users will be obtained.|
|status|string|No|Approval task status, which is used to specify that data under this status will be obtained. For status values, refer to  Enum of third-party approval task status.|
|scroll_id|string|No|Data is returned in batches when requests of all tasks are obtained by using status. Use scroll_id to obtain the next batch of data.|
Note: instance_ids is required for obtaining the task status of a specified instance by using instance_ids. approval_codes, user_ids, andstatusare required for obtaining the task status of a specified user by using user_ids. approval_codes and status are required for obtaining all tasks of a specified status by using status. scroll_id is required for obtaining the next batch of data.

### Request body example

1.   Obtain the task status of a specified instance by using instance_ids
```json
{
    "instance_ids": ["oa_159160304"]
}
```

2.   Obtain the task status of a specified user by using user_ids
```json
{
    "approval_codes": ["B7B65FFE-C2GC-452F-9F0F-9AA8352363D6"],
    "user_ids": ["112321"],
    "status": "PENDING"
}
```
3.  Obtain the next batch of data
```json
{
    "scroll_id": "nF1ZXJ5VGhlbkZldGNoCgAAAAAA6PZwFmUzSldvTC1yU"
}
```

##  Response

###  Response body
|Parameter|Type|Required|Description|
|-|-|-|-|
|code|int|Yes|Error code. A value other than 0 indicates failure.|
|msg|String|No|Return code description|
|data|object|Yes||
|  ∟data|list|Yes||
|    ∟instance_id|string|Yes|Approval instance ID|
|    ∟approval_code|String|Yes|Approval's  approval_code|
|    ∟approval_id|String|Yes|Approval ID|
|    ∟status|String|Yes|Current status of an approval instance. Refer to Enum of third-party approval task status.|
|    ∟update_time|int|Yes|Last update time of an approval instance (in milliseconds)|
|    ∟tasks|list|No|Approval tasks under an approval instance|
|      ∟id|String|Yes|Approval task ID|
|      ∟status|String|Yes|Approval task status. For status values, refer to  Enum of third-party approval task status.|
|      ∟update_time|int|Yes|Last update time of an approval task (in  ms)|
|  ∟scroll_id|string|No|Data is returned in batches when requests of all tasks are obtained by using status. Use scroll_id to obtain the next batch of data until scroll_id is empty.|
###  Response body example

```json
{
    "code": 0,
    "msg": "success",
    "data": {
        "data": [
            {
                "approval_code": "60D0F7A2-052D-49E9-B570-C29C836CDF8E",
                "approval_id": "fwwweffff33111133xxx",
                "instance_id": "29075",
                "status": "PENDING",
                "tasks": [
                    {
                        "id": "310",
                        "status": "PENDING",
                        "update_time": 1621863215000
                    }
                ],
                "update_time": 1621863215000
            }
        ],
        "scroll_id": "DnF1ZXJ5VGhlbkZldGNoCgAAAAAFSe_nFjlFcHhJU2dXVEJlbzRhUDd2MHFEX2cAAAAABUnKKBZ2dkJLRkFHQVRrdTJWTGF3M2JFeENnAAAAAAOGMY0WUS1XNmw3bFlUZ2VORjNVT2cwOWtxUQAAAAADhjGOFlEtVzZsN2xZVGdlTkYzVU9nMDlrcVEAAAAAA8VIKxZ4VEFGaHRHRVE5V0s0ek9lNE9nOWpRAAAAAAQ9zgEWTDRLcUJ4c2VUU21ZRV9FQlRLWmNCQQAAAAAEfNk8Fno4emowUExBUzJTaFhPTkprU2RBaXcAAAAABQpkOxZ4V1drX2M5UVEycW5XUmpvNXJweG13AAAAAAPFSCwWeFRBRmh0R0VROVdLNHpPZTRPZzlqUQAAAAAEPlQ8FmZNc0Rsdm1TU2t5VnFrWjFsYjRhdlE="
    }
}
```
