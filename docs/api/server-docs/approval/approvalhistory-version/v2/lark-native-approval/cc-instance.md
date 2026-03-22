---
title: "CC Instance"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uADOzYjLwgzM24CM4MjN"
method: "POST"
api_path: "https://www.larksuite.com/approval/openapi/v2/instance/cc"
section: "Approval"
scopes:
  - "approval:approval:readonly"
updated: "1658309950000"
---

# Copy an approval instance
> This API has been upgraded to enhance its security. We recommend that you migrate to the new version>>

This API is used to copy the current approval instance to other users.

##  Request
| HTTP URL | https://www.larksuite.com/approval/openapi/v2/instance/cc |
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
|approval_code | string | Yes |  Approval definition code |
|instance_code |string | Yes | Approval instance code  |
|user_id | string | No | user_id of CC initiator |
|open_id | string | No | open_id of CC initiator. If a user_id is provided, the user_id should be used. Open ID and user ID can't be both empty. |
|cc_user_ids | list | No | user_id list of CC recipients |
|cc_open_ids | list | No | open_id list of CC recipients. This parameter and cc_user_ids can't be both empty. |
|comment | string | No | CC comment | ### Request body example

```json
{
    "approval_code":"7C468A54-8745-2245-9675-08B7C63E7A85",
    "instance_code":"7C468A54-8745-2245-9675-08B7C63E7A85",
    "user_id":"f7cb567e",
    "open_id":"ou_123456",
    "cc_user_ids": ["f7cb567e"],
    "cc_open_ids": ["ou_123456"],
    "comment": "123"
}
````

##  Response

###  Response body

| Parameter         |Type         |Required  | Description        |
| --------- | ----------|----- | --------- |
|code |int |Yes |Error code, a value other than 0 indicates failure. |
|msg | string |Yes| Return code description| ###  Response body example

```json
{
    "code":0,
    "msg":"success",
}
```
