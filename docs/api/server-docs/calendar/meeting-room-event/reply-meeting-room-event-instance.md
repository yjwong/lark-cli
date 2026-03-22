---
title: "Reply meeting room event instance"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uYzN4UjL2cDO14iN3gTN"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/meeting_room/instance/reply"
section: "Calendar"
scopes:
  - "calendar:room:readonly"
updated: "1736835311000"
---

# Reply to room event instance

Call this interface to reply to the conference room event instance. It supports replying to release without sign-in, early end release, set to accept by the administrator, and set to reject by the administrator.

> **Note:** **Note**: You can first call the Query Conference Room Busyness interface to obtain the uid and original_time information of a certain event in the specified conference room, and then call this interface to reply to the meeting room event.
>   
> **ID Note**: The format of the event ID (event_id) is `_`, so you can get the Uid and Original time of the event through event_id. For example, if the event ID is `c32537e6-e0a8-4506-b42f-47440655cdb4_0`, then the Uid is `c32537e6-e0a8-4506-b42f-47440655cdb4` and the Original time is `0`.

##  Request
| HTTP URL | https://open.larksuite.com/open-apis/meeting_room/instance/reply |
| --- | --- |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes | `calendar:room:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about how to obtain and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body

| **Parameter** | **Parameter type** | **Required** | **Description**                                                     |
| -------- | ------------ | -------- | ------------------------------------------------------------ |
| room_id | string | Yes | Room ID. |
| uid | string | Yes | The event Uid corresponding to the conference room. |
| original_time | int | Yes | The original time of the event instance. For non-recurring events and recurring events, pass 0 here; for exception events for recurring events, pass the corresponding original_time value (timestamp type) here. |
| status | string | Yes | Reply status. **Optional values are:** - NOT_CHECK_IN: Not signed in- ENDED_BEFORE_DUE: End early- ACCEPTED_BY_ADMIN: Set to accept by the administratorDECLINED_BY_ADMIN: Set to accept by the administrator Reject | ### Request body example

```json
{
        "room_id":"omm_c7e9ef91ec9f3c007f1dfb2503918981",
        "uid":"bff6b51f-b7c1-40c6-b8ef-aef966c9ffc7",
        "original_time":0,
        "status":"NOT_CHECK_IN"
} 
```

##  Response
### Response body

| **Parameters** |**Type**| **Description** |
| ----------- |---| -------------------- |
| code |int| Return code, non-0 indicates failure. **Note**: If 105003 is returned, it means original_time is illegal. This problem may be caused by the entire start time of the repeated event being modified. It is recommended that the application re-query the conference room event list, obtain the latest original_time, and then try to call this interface again. |
| msg |string| Description of the return code, `success` means success, others are error messages. | ### Response body example

```json
{
    "code":0,
    "msg":"success"
}
```

### Error code

You can go to General Error Code.
