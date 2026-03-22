---
title: "Delete Building"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uMzMxYjLzMTM24yMzEjN"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/meeting_room/building/delete"
section: "Deprecated Version (Not Recommended)"
scopes:
  - "calendar:room"
updated: "1686900942000"
---

# Delete building

This API is used to delete buildings (office buildings). 

##  Request
| HTTP URL | https://open.larksuite.com/open-apis/meeting_room/building/delete |
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
| building_id  | string      | Yes   | Building ID  to be deleted| ### Request body example

```json
{
	"building_id":"omb_8ec170b937536a5d87c23b418b83f9bb"
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
