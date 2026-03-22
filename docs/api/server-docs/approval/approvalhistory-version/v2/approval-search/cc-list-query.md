---
title: "Cc List Query"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUjMxYjL1ITM24SNyEjN"
method: "POST"
api_path: "https://www.larksuite.com/approval/openapi/v2/cc/search"
section: "Approval"
scopes:
  - "approval:approval.list:readonly"
updated: "1658309977000"
---

# Query CC list
> This API has been upgraded to enhance its security. We recommend that you migrate to the new version>>

This API is used to query matching results from the approval CC list based on query conditions.

##  Request
| HTTP URL | https://www.larksuite.com/approval/openapi/v2/cc/search |
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
|instance_external_id|string|No|Approval instance third-party  idNote: Take the union with approval_code
|group_external_id|string|No|Approval definition group third-party  idNote: Take the union with instance_code
|cc_title|string|No|Approval CC title (only for third-party approvals)
|read_status|string|No|Approval CC statusREAD:  readUNREAD: unreadNote: If this parameter is not set, all statuses will be queried.An error will be reported if there is a status that isn't in the collection.
|CcCreateTimeFrom|int|No|CC query start time (unix timestamp in ms)
|CcCreateTimeTo|int|No|CC query end time (unix timestamp in ms)
|locale|string|No|Region  (zh-CN, en-US, ja-JP)
|offset|int|No|Query offsetNote: It shall not exceed 10,000
|limit|int|No|Query limitNote: It shall not exceed 200The default value is 10 if this parameter is not set.

> Note:
> 1. user_id, approval_code, instance_code, instance_external_id, and group_external_id can't be empty at the same time.
> 2. Take the union of the query results of approval_code and group_external_id , take the union of the query results of instance_code and instance_external_id , and take the intersection of other query conditions.
> 3. The query time span can't exceed 30 days. You must set either both the start and end times or neither. 

### Request body example

```json
{
	"user_id": "lwiu098wj",
	"approval_code": "EB828003-9FFE-4B3F-AA50-2E199E2ED942",
	"insatnce_code": "EB828003-9FFE-4B3F-AA50-2E199E2ED943",
	"instance_external_id": "xxxxx",
	"group_external_id": "xxxx",
	"cc_title": "xxx",
	"CcCreateTimeFrom": 1547654251506,
	"CcCreateTimeTo": 1547654251506,
	"read_status": "READ",
	"locale": "zh-CN",
	"offset": 0,
	"limit": 50,
}
```

##  Response

###  Response body
|Parameter|Type|Description|
|-|-|-|-|
|code |int |Yes |Error code, a value other than 0 indicates failure. |
|msg | string |Yes| Return code description|
|data | map | Returned business information |
|  ∟data | map |Yes| Returned business information |
|  ∟count|int|Number of data entries returned by the query
|  ∟cc_list|list|Approval instance list
|  ∟approval|map|Approval definition
|    ∟code|string|Approval definition code
|    ∟name|string|Approval definition name
|    ∟is_external|bool|Whether it is a third-party approval
|    ∟external|map|Third-party approval information
|      ∟batch_cc_read|bool|Whether batch read is supported
|  ∟group|map|Approval definition group
|    ∟external_id|string|External ID of approval definition group
|    ∟name|string|Name of approval definition group
|  ∟instance|map|Approval instance information
|    ∟code|string|Approval instance code
|    ∟external_id|string|External ID of approval instance
|    ∟user_id|string|Initiator ID of approval instance
|    ∟start_time|int|Start time of approval instance
|    ∟end_time|int|End time of approval instance
|    ∟status|string|Approval instance status
|    ∟title|string|Approval instance name (only for third-party approvals)
|    ∟extra|string|Approval instance extension field
|    ∟serial_id|string|Approval serial number
|    ∟link|map|Approval instance link (only for third-party approvals)
|      ∟pc_link|string|Approval instance link on PC
|      ∟mobile_link|string|Approval instance link on mobile
|  ∟cc|map|Approval CC
|    ∟user_id|string|Approval CC initiator ID
|    ∟create_time|int|Approval CC start time
|    ∟read_status|string|Approval CC status
|    ∟title|string|Approval CC name
|    ∟extra|string|Approval CC extension field
|    ∟link|map|Approval CC link
|      ∟pc_link|string|Approval CC link on PC
|      ∟mobile_link|string|Approval CC link on mobile
###  Response body example

```json
{
	"code": 0,
	"msg": "success",
	"data": {
		"count": 10,
		"cc_list": [
			{
				"approval": {
					"code": "EB828003-9FFE-4B3F-AA50-2E199E2ED943",
					"name": "qingjia",
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
				},
				"cc": {
					"user_id": "lwiu098wj",
					"read_status": "READ",
					"create_time": 1547654251506,
					"title": "",
					"extra": "{xxxx}",
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
