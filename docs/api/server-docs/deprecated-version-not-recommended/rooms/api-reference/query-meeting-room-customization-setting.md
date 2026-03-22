---
title: "Query Meeting Room Customization Setting"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uIjM5UjLyITO14iMykTN/query-meeting-room-customization-setting"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/meeting_room/room/customization"
section: "Deprecated Version (Not Recommended)"
scopes:
  - "calendar:room"
updated: "1686900917000"
---

# Obtain personalization settings of rooms

This API is used to obtain the personalization settings of rooms.

##  Request
| HTTP URL | https://open.larksuite.com/open-apis/meeting_room/room/customization |
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
| room_ids  | array      | Yes   | List of room IDs to be obtained | ### Request body example

```json
{
	"room_ids": [
  		"omm_2f3060afa404d3db7bb473xxxxxxxxxx"
  	]
}
```

## 	Response
### Response body

| Parameter         | Parameter type | Description                                                 |
| ------------ |----| ---------------------------------------------------- |
| error_room_ids         |-| Returns incorrect input parameter                                |
|    ∟building_id	      |string| Building ID |
|    ∟room_id	      |string| Room ID |
|    ∟error_msg	      |string| Error message |
| room_id_to_customization      |-| Room ID - mapping of personalization settings |
| ∟room_id	      |string| Room ID |
|    ∟contact_ids	      |array| List of IDs of members preparing personalization settings for rooms |
|    ∟customization_data	      |-| Personalization settings data |
|       ∟conditions	      |-| Display conditions of the questionnaire. When the conditions are satisfied, the questionnaire will appear for users to select | |          ∟custom_key	      |string| The questionnaire will only be displayed after the index_key of the custom_key is selected. | |          ∟option_keys	      |array| The questionnaire will only be displayed after all options are selected. |
|       ∟customization_type	      |int| Questionnaire type. 1 for single option, 2 for multiple options, and 3 for fill in the blanks. |
|       ∟index_key	      |string| Independent id | for each questionnaire
|       ∟input_content	      |string| When the type is fill in the blanks, it is required to be filled in. |
|       ∟is_required	      |bool| When the type is fill in the blanks, it is required to be filled in. |
|       ∟label	      |string| Questions for each questionnaire |
|       ∟options	      |-| Options for each questionnaire |
|          ∟is_others	      |bool| Another option or not |
|          ∟is_selected	      |bool| Check this option or not |
|          ∟option_image_url	      |string| Option image |
|          ∟option_key	      |string| Unique id  of each option|
|          ∟option_label	      |string| Options for each questionnaire |
|          ∟others_content	      |string| Input of other options |
|       ∟place_holder	      |string| Placeholders for fill in the blanks |
|    ∟preparation_time	      |int| Preparation time | ### Response body example

```json
{
    "error_room_ids": [
        {
            "building_id": "",
            "error_msg": "",
            "room_id": ""
        },
        {
            "building_id": "",
            "error_msg": "",
            "room_id": ""
        }
    ],
    "room_id_to_customization": {
        "room_id": {
            "contact_ids": [
                "",
                ""
            ],
            "customization_data": [
                {
                    "conditions": [
                        {
                            "custom_key": "",
                            "option_keys": [
                                "",
                                ""
                            ]
                        },
                        {
                            "custom_key": "",
                            "option_keys": [
                                "",
                                ""
                            ]
                        }
                    ],
                    "customization_type": 0,
                    "index_key": "",
                    "input_content": "",
                    "is_required": true,
                    "label": "",
                    "options": [
                        {
                            "is_others": true,
                            "is_selected": true,
                            "option_image_url": "",
                            "option_key": "",
                            "option_label": "",
                            "others_content": ""
                        },
                        {
                            "is_others": true,
                            "is_selected": true,
                            "option_image_url": "",
                            "option_key": "",
                            "option_label": "",
                            "others_content": ""
                        }
                    ],
                    "place_holder": ""
                }
            ],
            "preparation_time": 0
        }
    }
}
```

### Error code

For details, refer to Server-side error codes.
