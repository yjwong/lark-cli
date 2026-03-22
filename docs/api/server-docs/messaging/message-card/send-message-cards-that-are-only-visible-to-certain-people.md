---
title: "Send message cards that are only visible to certain people"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uETOyYjLxkjM24SM5IjN"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/ephemeral/v1/send"
section: "Messaging"
updated: "1717574926000"
---

# Send message cards that are only visible to certain people
Sends a message card that is visible only to a specified user in a group chat.The "Only visible to you" sign will be displayed on the card.

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/e55d303c1003a7fcbbf919e3bd086fa5_uQVcOD14bJ.png?lazyload=true&width=1366&height=448)
## Scenarios
A temporary message card is mostly used in the intermediate state of interaction between users and bots in group chats. For example, in a group chat, a user needs to use a to-do bot to create a reminder, and the bot sends an interactive message card in which the user can set the reminder date and reminder content. This card is visible to all members of the group before this card is set as a temporary card. That is, the interaction between the user and the bot is visible to the group. After this card is set as a temporary card, only the user can view the interaction process, and other members in the group can only view the reminder card that has been set by the user.
A temporary message card can reduce the disturbance of the message to irrelevant users in the group chat and reduce the noise caused by group messages.

> You must enable the bot, and the bot must be in the chat group.

> - A temporary message card is visible only to the user who triggers the temporary message card.
> - A temporary message card cannot be forwarded.
> - A temporary message card can only be used in the group chat.
> - A temporary message card is visible only when the user has logged in to Lark.
> - The Presentation capabilities and interaction capabilities of a temporary message card are the same as those of a message card.

## Request
| HTTP URL | https://open.larksuite.com/open-apis/ephemeral/v1/send |
| --- | --- |
| HTTP Method | POST | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about how to access and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body
| Parameter| Type   | Required | Description | Example |                                                   
| - | - | - | - | - | - |
| chat_id | string | Yes | ID of the group where you need to send the temporary message. The ID can be obtained by receiving pushed events. |oc_5ad573a6411d72b8305fda3a9c15c70e|
open_id user_id  email  | string | Yes | Specifies the user to whom you need to send the temporary message card. After you specify the user, the temporary message card is invisible to other members of the group. You only need to specify one of the open_id, email, and user_id parameters of the user. Open ID is recommended here. For details, refer to How to get Open ID? The server reads the parameters in the following sequence: **open_id** > **user_id** > **email** |ou_5ad573a6411d72b8305fda3a9c15c70e |
| msg_type                                          | string | Yes   | Message type, always "interactive" |interactive|
| card | object | Yes | Message card description. For more information, see Basic structure |
### Request body example

```json
{
   "chat_id": "oc_xxxxxxxxxxxx",
   "open_id": "ou_xxxxxxxxxxxx",
   "msg_type": "interactive",
   "card": {
        // card content
    }
}
```

## Response

### Response body
|Parameter|Type|Description|
|-|-|-|
|code|int|Return code. A value other than 0 indicates failure.|
|msg|string|Return code description|
data | - | -
  ∟message_id |string|Message ID

### Response body example

```json
{
    "code": 0,
    "data": {
        "message_id": "om_fd2057276f16243756ff8eb6fcd7b705"
    },
    "msg": "ok"
}
```

### Error code

For details, refer to Server-side error codes.

### Request body demo
```json
{
	 "open_id":"ou_515fbe9d04838174e2035f801453d07f",
	 "chat_id":"oc_286e0a771a76d99f1ec72737ef92ee8d",
     "msg_type": "interactive",
     "card": {
        "elements": [
            {
                "tag": "div",
                "text": {
                    "tag": "plain_text",
                    "content": "Overflow and datePicker features"
                },
                "fields": [
                    {
                        "is_short": true,
                        "text": {
                            "tag": "lark_md",
                            "content": "**Component:**\noverflow & datePicker"
                        }
                    },
                    {
                        "is_short": false,
                        "text": {
                            "tag": "lark_md",
                            "content": "**Feature points:**\n"
                        }
                    }
                ]
            },
            {
                "tag": "hr"
            },
            {
                "tag": "action",
                "actions": [
                    {
                        "tag": "overflow",
                        "confirm": {
                            "title": {
                                "tag": "plain_text",
                                "content": "Please check"
                            },
                            "text": {
                                "tag": "plain_text",
                                "content": "Redirection method"
                            }
                        },
                        "options": [
                            {
                                "text": {
                                    "tag": "plain_text",
                                    "content": "Multi-end redirection"
                                },
                                "multi_url": {
                                    "url": "https://www.baidu.com",
                                    "android_url": "https://developer.android.com/",
                                    "ios_url": "https://developer.apple.com/",
                                    "pc_url": "https://www.windows.com"
                                },
                                "value": "james"
                            },
                            {
                                "text": {
                                    "tag": "plain_text",
                                    "content": "Specify redirection"
                                },
                                "url": "https://www.baidu.com"
                            },
                            {
                                "text": {
                                    "tag": "plain_text",
                                    "content": "Text content"
                                },
                                "value": "joy"
                            },
                            {
                                "text": {
                                    "tag": "plain_text",
                                    "i18n": {
                                        "zh_cn": "中文文本-james",
                                        "en_us": "English text-james",
                                        "ja_jp": "日本語文案-james"
                                    }
                                },
                                "value": "james"
                            }
                        ]
                    },
                    {
                        "tag": "date_picker",
                        "placeholder": {
                            "tag": "plain_text",
                            "content": "Select a date (with the Confirm button)"
                        },
                        "confirm": {
                            "title": {
                                "tag": "plain_text",
                                "content": "Please check"
                            },
                            "text": {
                                "tag": "plain_text",
                                "content": "Redirection method"
                            }
                        }
                    },
                    {
                        "tag": "date_picker",
                        "placeholder": {
                            "tag": "plain_text",
                            "content": "Select a date"
                        },
                        "initial_date": "2002-2-22"
                    },
                    {
                        "tag": "picker_time",
                        "placeholder": {
                            "tag": "plain_text",
                            "content": "Select a date (with the Confirm button)"
                        },
                        "confirm": {
                            "title": {
                                "tag": "plain_text",
                                "content": "Please check"
                            },
                            "text": {
                                "tag": "plain_text",
                                "content": "Redirection method"
                            }
                        }
                    },
                    {
                        "tag": "picker_time",
                        "placeholder": {
                            "tag": "plain_text",
                            "content": "Select a date"
                        },
                        "initial_time": "19:09"
                    },
                    {
                        "tag": "picker_datetime",
                        "placeholder": {
                            "tag": "plain_text",
                            "content": "Select a date (with the Confirm button)"
                        },
                        "confirm": {
                            "title": {
                                "tag": "plain_text",
                                "content": "Please check"
                            },
                            "text": {
                                "tag": "plain_text",
                                "content": "Redirection method"
                            }
                        }
                    },
                    {
                        "tag": "picker_datetime",
                        "placeholder": {
                            "tag": "plain_text",
                            "content": "Select a date"
                        },
                        "initial_datetime": "2002-2-22 18:00"
                    }
                ]
            }
        ]
    }
 }
```
