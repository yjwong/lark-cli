---
title: "Delete message cards that are only visible to certain people"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uITOyYjLykjM24iM5IjN"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/ephemeral/v1/delete"
section: "Messaging"
updated: "1717574930000"
---

# Delete message cards that are only visible to certain people
Deletes a temporary message card that is visible only to a specified user in a group chat.
Temporary card messages can be deleted explicitly via this API, and will leave no trace once deleted.
> **Permission description**: You must enable the bot ability and the bot must be in the chat group.

## Request
| HTTP URL | https://open.larksuite.com/open-apis/ephemeral/v1/delete |
| --- | --- |
| HTTP Method | POST | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about how to access and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body
| Parameter | Type | Required | Description | Example |                                                   
| - | - | - | - | - | - |
| message_id | string | Yes | ID of the temporary message |om_5ad573a6411d72b8305fda3a9c15c70e|
### Request body example

```json
{
   "message_id": "om_xxxxxxxxxxxx"
}
```
## Response
### Response body example

```json
{
    "code": 0,
    "msg": "ok"
}
```
### Error code
| Error code | Description | Troubleshooting suggestions |
| --- | --- | --- |
| 18051 | The temporary message has been deleted. | Deletion failed because the temporary message has been deleted by a user. | For more information about other common error codes, see Server-side error codes.
