---
title: "Batch Obtain Approval Instance IDs"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uQDOyUjL0gjM14CN4ITN"
method: "POST"
api_path: "https://www.larksuite.com/approval/openapi/v2/instance/list"
section: "Approval"
scopes:
  - "approval:approval:readonly"
updated: "1658309948000"
---

# Obtain approval instances in batches
> This API has been upgraded to enhance its security. We recommend that you migrate to the new version>>

This API is used to obtain the instance_codes of multiple approval instances in batches based on a specified approval_code. This can be used to obtain all approval instances of a tenant's approval definition.
Results are sorted by approval creation time by default.

##  Request
| HTTP URL | https://www.larksuite.com/approval/openapi/v2/instance/list |
| --- | --- |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes | `approval:approval:readonly` | ###  Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about how to obtain and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body

| Parameter         | Type           | Required        | Description        |
| --------- | --------------- | -------   | --------- |
|approval_code | string | Yes |  Unique identifier of approval definition |
|start_time |long | Yes | Approval instance creation time interval (in ms)  |
|end_time | long | Yes |Approval instance creation time interval (in ms) |
|offset | int | Yes | Query offset |
|limit | int | Yes | Query limit Note:It shall not exceed 100 | ### Request body example

```json
{
    "approval_code":"7C468A54-8745-2245-9675-08B7C63E7A85",
    "start_time":1567690398020,
    "end_time":1567690398020,
    "offset":0,
    "limit":100
}
````

##  Response

###  Response body

| Parameter         |Type         |Required  | Description        |
| --------- | ----------|----- | --------- |
|code |int |Yes |Error code. A value other than 0 indicates failure. |
|msg | string |Yes| Return code description|
|data | map |Yes| Returned business information |
|  ∟instance_code_list|list|Yes|Approval instance code| ###  Response body example

```json
{
    "code":0,
    "msg":"success",
    "data": {
    	"instance_code_list": [
            "357C21A0-2069-4F6B-955F-1DFBE6710C51",
            "A6F07B2A-0503-4197-8FD7-FEB40C79B47B",
            "3D9ED45B-0806-4B55-8356-BC0FF517408C",
            "33BF276D-0217-436D-A565-3081E5000E63",
            "F7765ABC-2609-49D3-B040-1BCBFD15E3B8",
        ]
    }
}
```
