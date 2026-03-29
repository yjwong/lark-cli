---
title: "Delete a Department"
url: "https://open.larksuite.com/document/ukTMukTMukTM/ugzNz4CO3MjL4czM"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/contact/v1/department/delete"
section: "Deprecated Version (Not Recommended)"
updated: "1646999099000"
---

# Delete a department
> In order to improve the security level, we have upgraded it. Please migrate to the new version as soon as possible. New version >>

This API is used to delete a department from Contacts. 

- The application should apply for `update contacts` and `access contacts by application identity` permissions to invoke this API. Contacts permission for this department is required. 
To invoke this API, the request header must specify the tenant_access_token. 
Marketplace apps do not have permission to invoke this API.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/contact/v1/department/delete |
| HTTP Method | POST |
| Required scopes Enable any scope from the list | update contacts access contacts by application identity | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value Format**: "Bearer `access_token`" **Example Value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body   
|Parameter|Type|Mandatory|Description|
|--|-----|--|----|
|id|string|Yes|The ID of the department to be deleted|
### Request body example
```json
{
    "id": "od-455efa262dc736b3e45a8b17fe945293"
}
```

## Response
### Response body 
|Parameter|Description|
|-|-|
|code|Error code, anything but 0 indicates failure|
|msg|Response code description|
### Response body example
```json
{
    "code": 0,
    "msg": "success"
}
```
### Error code

For details, please refer to: Service-side error codes
