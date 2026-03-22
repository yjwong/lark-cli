---
title: "Delete Meeting Room"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUzMxYjL1MTM24SNzEjN"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/meeting_room/room/delete"
section: "Deprecated Version (Not Recommended)"
scopes:
  - "calendar:room"
updated: "1686900957000"
---

# Delete room

This API is used to delete rooms. 

##  Request
| HTTP URL | https://open.larksuite.com/open-apis/meeting_room/room/delete |
| --- | --- |
| HTTP Method | POST |
| Supported app types | custom |
| Required scopes | `calendar:room` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about how to obtain and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body

| Parameter       | Parameter type | Required | Description                                                         |
| ---------- | -------- | ---- | ------------------------------------------------------------ |
| room_id  | string      | Yes   | Room ID  to be deleted| ###  Request body example

```json
{
	"room_id":"omb_8d020b12fe49e82847c2af3c193d5754"
}
```

##  Response
### Response body

| Parameter         | Description                                                 |
| ------------ | ---------------------------------------------------- |
| code         | Return code. A value other than 0 indicates failure.                                |
| msg          | Description of the return code. "success" indicates success, and others are error messages. | ### Response body example

```json
{
    "code":0,
    "msg":"success",
}
```

### Error code

For details, refer to Server-side error codes.
