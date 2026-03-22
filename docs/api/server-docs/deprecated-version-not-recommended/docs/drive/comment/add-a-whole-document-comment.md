---
title: "Add a Whole Document Comment"
url: "https://open.larksuite.com/document/ukTMukTMukTM/ucDN4UjL3QDO14yN0gTN"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/comment/add_whole"
section: "Deprecated Version (Not Recommended)"
updated: "1647001470000"
---

# Add a whole document comment

This API is used to adding a comment to a document using the file token. The comment will be shown in the comments panel at the bottom of the document.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/comment/add_whole |
| HTTP Method | POST | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `user_access_token` **Value Format**: "Bearer `access_token`" **Example Value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body

|Parameter|Type|Mandatory|Description|
|--|-----|--|----|----|
|token|string|Yes|The file token, see page 4 of the Integration Overview for how to obtain it  |
|type|string|Yes|Document type, values include:  1) "doc": Document|
|content|string|Yes|Comment content (note: The comment content is required, and some special characters must be converted to escape characters; see the escape character table below)  | **Escape character table**:
|   Original character   |   Escape character   |
|--|--|
||  \>  |
|&| \&  |
|'| \&#x27;  |
|"| \&quot;   | ### Request body example
```json
{
	"type": "doc",
	"token": "doccnBKgoMyY5OMbUG6FioTXuBe",
	"content": ""
}
```
## Response

### Response body

|   Parameter   |Description|
|--|--|
|comment_id|  Comment ID  |
|reply_id|  Reply ID  |
|create_timestamp|  Creation timestamp  |
|update_timestamp|  Update timestamp  | ### Response body example
```json
{
	"code": 0,
	"data": {
		"comment_id": "1575537199207062072",
		"reply_id": "1575537230887731780",
		"create_timestamp": 1575537230,
		"update_timestamp": 1575537230
	},
	"msg": "Success"
}
```

### Error code

For details, please refer to: Service-side error codes
