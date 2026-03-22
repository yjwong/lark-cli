---
title: "Instance List Query"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uQjMxYjL0ITM24CNyEjN"
method: "POST"
api_path: "https://www.larksuite.com/approval/openapi/v2/instance/search"
section: "Approval"
scopes:
  - "approval:approval.list:readonly"
updated: "1658309974000"
---

# Query an instance list
> This API has been upgraded to enhance its security. We recommend that you migrate to the new version>>

This API is used to query matching results from the approval instance list based on query conditions.

##  Request
| HTTP URL | https://www.larksuite.com/approval/openapi/v2/instance/search |
| --- | --- |
| HTTP Method | POST |
| Supported app types | custom |
| Required scopes | `approval:approval.list:readonly` | ###  Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about how to obtain and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body

|Parameter|Type|Required|Description|
|-|-|-|-|
|user_id|string|Yes|User ID
|approval_code|string|No|Approval definition code
|instance_code|string|No|Approval instance code
|instance_external_id|string|No|Approval instance third-party IDNote: Take the union with approval_code
|group_external_id|string|No|Approval definition group third-party IDNote: Take the union with instance_code
|instance_title|string|No|Approval instance title (only for third-party approvals)
|instance_status|string|No|Approval instance statusREJECT: RejectedPENDING: Under reviewRECALL: RecalledDELETED: DeletedAPPROVED: ApprovedNote: If this parameter is not set, all statuses will be queried.An error will be reported if there is a status that isn't in the collection.
|instance_start_time_from|int|No|Instance query start time (unix timestamp in ms)
|instance_start_time_to|int|No|Instance query end time   (unix timestamp in ms)
|locale|string|No|Region  (zh-CN, en-US, ja-JP)
|offset|int|No|Query offset Note: It shall not exceed 10,000
|limit|int|No|Query limit Note: It shall not exceed 200The default value is 10 if this parameter is not set.
> Note:
> 1. user_id, approval_code, instance_code, instance_external_id, and group_external_id can't be empty at the same time.
> 2. Take the union of the query results of approval_code and group_external_id, take the union of the query results of instance_code and instance_external_id, and take the intersection of other query conditions.
> 3. The query time span can't exceed 30 days. You must set either both the start and end times or neither. 

### Request body example

```json
{
	"user_id": "lwiu098wj",
	"approval_code": "EB828003-9FFE-4B3F-AA50-2E199E2ED942",
	"instance_code": "EB828003-9FFE-4B3F-AA50-2E199E2ED943",
	"instance_external_id": "xxxxx",
	"group_external_id": "xxxx",
	"instance_title": "xxx",
	"instance_status": "PENDING",
	"instance_start_time_from": 1547654251506,
	"instance_start_time_to": 1547654251506,
	"locale": "zh-CN",
	"offset": 0,
	"limit": 50
}
```

##  Response

###  Response body

|Parameter|Type|Description|
|-|-|-|-|
|code |int |Error code. A value other than 0 indicates failure. |
|msg | string | Return code description|
|data | map | Returned business information |
|  ∟count|int|Number of data entries returned by the query
|  ∟instance_list|list|Approval instance list
|    ∟approval|map|Approval definition
|      ∟code|string|Approval definition code
|      ∟name|string|Approval definition name
|      ∟is_external|bool|Whether it is a third-party approval
|      ∟external|map|Third-party approval information
|        ∟batch_cc_read |bool|Whether batch read is supported
|    ∟group|map|Approval definition group
|        ∟external_id|string|External ID of approval definition group
|        ∟name|string|Name of approval definition group
|    ∟instance|map|Approval instance information
|      ∟code|string|Approval instance code
|      ∟external_id|string|External ID of approval instance
|      ∟user_id|string|Initiator ID of approval instance
|      ∟start_time|int|Start time of approval instance
|      ∟end_time|int|End time of approval instance
|      ∟status|string|Approval instance status
|      ∟title|string|Approval instance name (only for third-party approvals)
|      ∟extra|string|Approval instance extension field
|      ∟serial_id|string|Approval serial number
|      ∟link|map|Approval instance link (only for third-party approvals)
|        ∟pc_link|string|Approval instance link on PC
|        ∟mobile_link|string|Approval instance link on mobile

###  Response body example

```json
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
						 "batch_cc_read": false,
					}
				},
				"group": {
					 "external_id": "0004",
				 	"name": "groupA",
				},
				"instance": {
					"code": "EB828003-9FFE-4B3F-AA50-2E199E2ED943",
					"external_id": "0004_3ED52DC1-AA6C",
					"user_id": "lwiu098wj",
					"start_time": 1547654251506,
					"end_time": 1547654251506,
					"status": "PENDING",
					"title": "",
					"extra": "{}",
					"serial_id": "201902020001",
			 		"link": {
						"pc_link": "",
						"mobile_link": "",
					}
				}
			}
		]
	}
}
```
