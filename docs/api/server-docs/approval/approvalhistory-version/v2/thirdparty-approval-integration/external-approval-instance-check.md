---
title: "External Approval Instance Check"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUDNyYjL1QjM24SN0IjN"
method: "POST"
api_path: "https://www.larksuite.com/approval/openapi/v3/external/instance/check"
section: "Approval"
scopes:
  - "approval:approval:readonly"
updated: "1658309967000"
---

# Third-party approval instance validation
> This API has been upgraded to enhance its security. We recommend that you migrate to the new version>>

Third-party approval instance validation is used to determine whether the server-side data is up-to-date. The user submits the last update time of the instance, and if the instance doesn't exist on the server-side, or the update time of the instance on the server-side isn't the latest, the corresponding instance ID is returned.

For example, users can compare the instances generated in the last 5 min using this API every 5 min.

##  Request
| HTTP URL | https://www.larksuite.com/approval/openapi/v3/external/instance/check |
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
|update_times|list|Yes|Instance information|
|∟instance_id|String|Yes|Approval instance ID|
|∟update_time|int|Yes|Last update time of approval instance|
|∟tasks|list|Yes|Task information|
|  ∟task_id|String|Yes|Task ID|
|  ∟update_time|int|Yes|Last update time of task| ### Request body example

```json
{
    "update_times": [
        {
            "instance_id": "1234234234242423",
            "update_time": 1591603040000,
            "tasks": [
                {
                    "task_id": "112253",
                    "update_time": 1591603040000
                },
                {
                    "task_id": "112255",
                    "update_time": 1591603040000
                }
            ]
        }
    ]
}
```

##  Response

###  Response body
|Parameter|Type|Required|Description|
|-|-|-|-|
|code|int|Yes|Error code. A value other than 0 indicates failure.|
|msg|String|Yes|Return code description.|
|data|map|Yes|Returned business information|
|  diff_times|list|No|Instance information|
|    ∟instance_id|String|Yes|Approval instance ID|
|    ∟update_time|int|No|Last update time of approval instance|
|    ∟tasks|list|No|Task information|
|      ∟task_id|String|Yes|Task ID|
|      ∟update_time|int|No|Last update time of task|
###  Response body example

```json
{
    "code":0,
    "msg": "success",
    "data":{
      "diff_times": [
              {
                  "instance_id": "1234234234242423",
                  "update_time": 1591603040000,
                  "tasks": [
                      {
                          "task_id": "112255",
                          "update_time": 0
                      }
                  ]
              }
          ]
      }
}
```
