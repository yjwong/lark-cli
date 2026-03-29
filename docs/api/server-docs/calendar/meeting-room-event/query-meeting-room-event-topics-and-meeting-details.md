---
title: "Query meeting room event topics and meeting details"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uIjM5UjLyITO14iMykTN/"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/meeting_room/summary/batch_get"
section: "Calendar"
scopes:
  - "calendar:room"
updated: "1736835302000"
---

# Query meeting room event topics and meeting details

Call this interface to query the meeting room event theme and details using the event's Uid and Original time.

> **Note:** The format of the event ID (event_id) is `_`, so you can get the Uid and Original time of the event through event_id. For example, if the event ID is `c32537e6-e0a8-4506-b42f-47440655cdb4_0`, then the Uid is `c32537e6-e0a8-4506-b42f-47440655cdb4` and the Original time is `0`.

##  Request
| HTTP URL | https://open.larksuite.com/open-apis/meeting_room/summary/batch_get |
| --- | --- |
| HTTP Method | POST |
| Supported app types | custom |
| Required scopes | `calendar:room` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about how to obtain and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body

| **Parameter** | **Parameter type** | **Required** | **Description**                                                     |
| -------- | ------------ | -------- | ------------------------------------------------------------ |
| EventUids | - | Yes | List of event Uid and Original time to be queried. |
| ∟uid | string | Yes| The unique ID of the event. |
| ∟original_time | int | Yes | The original time of the event instance. For non-recurring events and recurring events, pass 0 here; for exception events for recurring events, you need to pass in the corresponding original_time value (timestamp type). | ### Request body example

```json
{
	"EventUids":[
		{
			"uid":"a04dbea1-86b9-4372-aa8d-64ebe801be2a",
			"original_time":0
		},
        {
			"uid":"d7a44c9b-7ae0-4a97-bf80-b4f050cedffa",
			"original_time":0
		}
	]
}
```

##  Response

### Response body

| **Parameter**    | **Parameter type** |**Description**                                             |
| ----------- | -------- |---------------------------------------------------- |
| code | int|Return code, non-0 indicates failure. |
| msg | string|Description of return code, `success` means success, others are error messages. |
| data | -|Return information. |
| ∟EventInfos | -|The queried event information. |
| ∟∟uid | string|The unique ID of the event. |
| ∟∟original_time | int|Original time of the event instance. For non-recurring events and recurring events, here is 0; for exception events for recurring events, here is the corresponding original_time value (timestamp type). |
| ∟∟summary | string|Event theme. |
| ∟∟vchat | vhcat|Video conference information. |
| ∟∟∟vc_type| string | Video conference type.  **Optional values are:** - `vc`: Lark video conferencing. When this type is selected, other fields in vchat are invalid.  - `third_party`: Third party link video conferencing. When this type is selected, only the icon_type, description, and meeting_url fields in vchat will take effect.  - `no_meeting`: No video conferencing. When this type is selected, other fields in vchat are invalid. - `lark_live`: Lark live broadcast. This value is a client-side, read-only parameter. - `unknown`: Unknown type. This value is used by the client for compatibility and is a read-only parameter. |
| ∟∟∟icon_type| string | Third-party video conferencing icon type.  **Optional values are:** - `vc`: Lark video conference icon- `live`: Live video conference icon - `default`: Default icon|
| ∟∟∟description|string|Third-party video conferencing copywriting. |
| ∟∟∟meeting_url|string|Video conference URL. |
| ∟ErrorEventUids | - |No event information found. |
| ∟∟uid | string |The unique ID of the event. |
| ∟∟original_time | int |Event instance original time. For non-recurring events and recurring events, here is 0; for exception events for recurring events, here is the corresponding original_time value (timestamp type). |
| ∟∟error_msg | string | Error message. | ### Response body example

```json
{
    "code": 0,
    "data": {
        "ErrorEventUids": [],
        "EventInfos": [
            {
                "original_time": 0,
                "summary": "test",
                "uid": "a04dbea1-86b9-4372-aa8d-64ebe801be2a",
                "vchat": {
                    "meeting_url": "https://vc.larksuite.com/j/935314044",
                    "vc_type": "vc"
                }
            },
            {
                "original_time": 0,
                "summary": "event",
                "uid": "d7a44c9b-7ae0-4a97-bf80-b4f050cedffa",
                "vchat": {
                    "meeting_url": "https://vc.larksuite.com/j/777110140",
                    "vc_type": "vc"
                }
            }
        ]
    },
    "msg": ""
}
```

### Error code

You can go to General Error Code.
