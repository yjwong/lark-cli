---
title: "Events"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uMTNxYjLzUTM24yM1EjN"
section: "Deprecated Version (Not Recommended)"
updated: "1622191953000"
---

# Bot and Chat Message Events

## Bot Added to Group
This event is triggered when a bot is invited to join a group.   

- Prerequisites: The app must have enabled the bot feature.

**Callback example:**
```json
{ 
    "ts": "1502199207.7171419", // the time the event is sent; this is usually close to the time the event is generated 
    "uuid": "bc447199585340d1f3728d26b1c0297a",  // unique event ID
    "token": "41a9425ea7df4536a7623e38fa321bae", // authentication token 
    "type": "event_callback", // event callbacks always use event_callback
    "event": { 
        "app_id": "cli_9c8609450f78d102", 
        "chat_i18n_names": { // localised group name field 
            "en_us": "English Title", 
            "zh_cn": "Chinese Title"
        }, 
        "chat_name": "Group Name",
        "chat_owner_employee_id": "ca51d83b",// the group owner’s employee_id (i.e. user ID). If the group owner is a bot, this field won’t exist. This field is only returned for custom apps 
        "chat_owner_name": "xxx", // group owner name
        "chat_owner_open_id": "ou_18eac85d35a26f989317ad4f02e8bbbb", // the group owner’s open_id 
        "open_chat_id": "oc_e520983d3e4f5ec7306069bffe672aa3",  // the chat ID
        “operator_employee_id": "ca51d83b", // the operator’s employee_id; only returned for custom apps. The operator is the person that invited the bot
        "operator_name": "yyy", // the operator’s name 
        "operator_open_id": "ou_18eac85d35a26f989317ad4f02e8bbbb",// the operator's open_id 
        "owner_is_bot": false, //if the group owner is a bot 
        "tenant_key": "736588c9260f175d",  // organization ID
        "type": "add_bot" // event type
    }
}

```

## Bot Removed from Group
This event is triggered when a bot is removed from a group chat.

- Prerequisites: The app must have enabled the bot feature.

**Callback example:**
```json
{ 
    "ts": "1502199207.7171419", // the time the event is sent; this is usually close to the time the event is generated 
    "uuid": "bc447199585340d1f3728d26b1c0297a",  // unique event ID
    "token": "41a9425ea7df4536a7623e38fa321bae", // authentication token 
    "type": "event_callback", // event callbacks always use event_callback
    "event": { 
        "app_id": "cli_9c8609450f78d102", 
        "chat_i18n_names": { // localized group name field 
            "en_us": "English Title", 
            "zh_cn": "Chinese Title"
        }, 
        "chat_name": "Group Name", 
        "chat_owner_employee_id": "ca51d83b",// the group owner’s employee_id (i.e. user ID). If the group owner is a bot, this field won’t exist. This field is only returned for custom apps 
        "chat_owner_name": "xxx", // group owner name
        "chat_owner_open_id": "ou_18eac85d35a26f989317ad4f02e8bbbb", // the group owner’s open_id 
        "open_chat_id": "oc_e520983d3e4f5ec7306069bffe672aa3",  // the chat ID
        "operator_employee_id": "ca51d83b", // the operator’s emplolyee_id; only returned for custom apps. The operator is the person who removed the bot from the group
        "operator_name": "yyy", // the operator’s name 
        "operator_open_id": "ou_18eac85d35a26f989317ad4f02e8bbbb",// the operator's open_id 
        "owner_is_bot": false, //if the group owner is a bot 
        "tenant_key": "736588c9260f175d",  // organization ID
        "type": "remove_bot" // bot removed: remove_bot
        
    }
}
```

## User and Bot Chat First Created
The first chat message is a key opportunity for the user to understand how the bot works. You can use this to send instructions or a configuration link to guide the user in getting started.

::: Note
Public apps must subscribe and respond to this event.

**Callback example:**
```json
{ 
    "ts": "1502199207.7171419", // the time the event is sent; this is usually close to the time the event is generated 
    "uuid": "bc447199585340d1f3728d26b1c0297a",  // unique event ID
    "token": "41a9425ea7df4536a7623e38fa321bae", // authentication token 
    "type": "event_callback", // event callbacks always use event_callback
    "event": { 
        "app_id": "cli_9c8609450f78d102",
        "chat_id": "oc_26b66a5eb603162b849f91bcd8815b20", // the ID of the chat between the bot and user
        "operator": { // the person who launched the chat. This could be the user or the bot
            "open_id":"ou_2d2c0399b53d06fd195bb393cd1e38f2" // the user’s unique ID for this app; each user has a separate open_id for each app
            "user_id": "gfa21d92"  // user ID; only returned for custom apps 
        },
        "tenant_key": "736588c9260f175c",  // organization ID 
        "type": "p2p_chat_create",  // event type
        "user": {  // the user in the chat
            "name": "Jason",
            "open_id":"ou_7dede290d6a27698b969a7fd70ca53da"  // the user’s unique ID for this app; each user has a separate open_id for each app
            "user_id": "gfa21d92"  // user ID; only returned for custom apps 
        }
    }
}
```

## Receive Message
This event is triggered when a user sends a message to a bot or mentions a bot in a group chat.

- Permissions required: Obtain private messages sent from a user to a bot and Obtain user messages @ a bot in a group chat. The messages can only be received when these permissions have been enabled.
- Other requirements: The app must have enabled the bot feature。

### Text Message 

**Callback example:**
```json
{ 
    "uuid": "41b5f371157e3d5341b38b20396e77e3", 
    "token": "2g7als3DgPW6Xp1xEpmcvgVhQG621bFY", // authentication token 
    "ts": "1550038209.428520",  // timestamp 
    "type": "event_callback",// event callbacks always use event_callback 
    "event": { 
        "type": "message", // event type 
        "app_id": "cli_xxx", 
        "tenant_key": "xxx",  // organization ID 
        "root_id": "", 
        "parent_id": "", 
        "open_chat_id": "oc_5ce6d572455d361153b7cb51da133945", 
        "chat_type": "private",// private or group 
        "msg_type": "text",    // message type 
        "open_id": "ou_18eac85d35a26f989317ad4f02e8bbbb", 
        "employee_id":"xxx",    // i.e. the user’s ID within the organization; only returned for custom apps
        "union_id": "xxx" // the user’s unique ID for this ISV; 
        "open_message_id": "om_36686ee62209da697d8775375d0c8e88", 
        "is_mention": false, 
        "text": "@bot message text@Jason",      // the message text, which may include an @ mention of a person or bot
        "text_without_at_bot":"message text @Jason" / message content with the bot mention filtered out 
   } 
} 
```
### Rich Text and Post Message 

**Callback example:**
```json
{ 
    "uuid": "bc447199585340d1f3728d26b1c0297a", 
    "token": "2g7als3DgPW6Xp1xEpmcvgVhQG621bFY", 
    "ts": "1550032496.527917", 
    "type": "event_callback", 
    "event": { 
        "type": "message", 
        "app_id": "cli_xxx", 
        "tenant_key": "xxx",  // organization ID 
        "root_id": "", 
        "parent_id": "", 
        "open_chat_id": "oc_5ce6d572455d361153b7cb51da133945",// the open_chat_id of the message 
        "chat_type": "private", 
        "msg_type": "post",// rich_text and post messages are here treated as post 
        "open_id": "ou_18eac85d35a26f989317ad4f02e8bbbb",// the open_id of the user that sent the message 
        "employee_id":"xxx",    // i.e. the user’s ID within the organization; only returned for custom apps
        "union_id": "xxx" // the user’s unique ID for this ISV; 
        "open_message_id": "om_a81fa00ee467b1084ffd80b197b58f4b",// message ID 
        "is_mention": false, 
        "text": "\u003cp\u003e测试1212\u003c/p\u003e\u003cfigure\u003e\u003cimg src=\"http://p4.pstatp.com/origin/daff000d4967d033705b\" origin-width=\"2456\" origin-height=\"2458\"/\u003e\u003c/figure\u003e",// message content 
        "text_without_at_bot":"消息内容",// message content with the bot mention filtered out 
        "title": "测试" ,// message title 
        "image_keys": [ // the keys to the images in the rich text
           "1867eac8-8006-40be-8549-b7beae0d3c4a",
           "434593af-5269-4db4-8b94-65c6dfd4f35e"
         ],
  } 
} 
```
### Image Message

**Callback example:** 
```json
{ 
    "uuid": "c58e17e9e84824a48e51a562cf969fb3", 
    "token": "2g7als3DgPW6Xp1xEpmcvgVhQG621bFY", 
    "ts": "1550038110.478493", 
    "type": "event_callback", 
    "event": { 
        "type": "message", 
        "app_id": "cli_xxx", 
        "tenant_key": "xxx",  // organization ID 
        "root_id": "", 
        "parent_id": "", 
        "open_chat_id": "oc_5ce6d572455d361153b7cb51da133945", 
        "chat_type": "private", 
        "msg_type": "image", // image message 
        "image_height" :"300",
        "image_width" :"300",
        "open_id": "ou_18eac85d35a26f989317ad4f02e8bbbb", 
        "employee_id":"xxx",    // i.e. the user’s ID within the organization; only returned for custom apps
        "union_id": "xxx" // the user’s unique ID for this ISV; 
        "open_message_id": "om_340057d660022bf141eb470859c6114c", 
        "is_mention": false, 
        "image_key": "cd1ce282-94d1-4154-a326-121b07e4721e", // image_key. For how to obtain the image, see /ssl:ttdoc/ukTMukTMukTM/uYzN5QjL2cTO04iN3kDN

  } 
} 
```
### File Message

**Callback example:** 
```json
{
    "uuid": "8bbb4b5ba77c316991c41fa9ef0ced8b",
    "token": "ZYcIUHYh8lmZen6jStLgvcXAXbqn2tle",
    "ts": "1576725930.795192",
    "type": "event_callback",
    "event": {
        "app_id": "cli_xxx",
        "tenant_key": "xxx",
        "chat_type": "private",
        "is_mention": false,
        "msg_type": "file",
        "open_chat_id": "oc_a1e061372f7745a543dsd5h3d6d2349a",
        "open_id": "ou_2b940d169b7a4a0c76633984b08ced73",
        "employee_id":"xxx",    // i.e. the user’s ID within the organization; only returned for custom apps
        "union_id": "xxx" // the user’s unique ID for this ISV; 
        "open_message_id": "om_d85c4sd7b209tbb98g693a958bc7185f",
        "parent_id": "",
        "root_id": "",
        "type": "message",
        "file_key": "file_b4do9r9b-3526-4bc4-a568-65fe3695b05g"
  }
}
```
### Merged and Forwarded Message Content 

**Callback example:**
```json
{ 
    "uuid": "d465df13905458e361ff39af85d96d09", 
    "token": "2g7als3DgPW6Xp1xEpmcvgVhQG621bFY", 
    "ts": "1550111453.764646", 
    "type": "event_callback", 
    "event": { 
        "type": "message", 
        "app_id": "cli_xxx", 
        "tenant_key": "xxx",  // organization ID    
        "root_id": "", 
        "parent_id": "", 
        "open_chat_id": "oc_5ce6d572455d361153b7cb51da133945", 
        "msg_type": "merge_forward", 
        "open_id": "ou_18eac85d35a26f989317ad4f02e8bbbb", 
        "employee_id":"xxx",    // i.e. the user’s ID within the organization; only returned for custom apps
        "union_id": "xxx" // the user’s unique ID for this ISV;
        "open_message_id": "om_b3961b120d67259e7495d8eb69488189", 
        "is_mention": false, 
        "chat_type": "private", 
        "user": "6610187460791558158", 
        "msg_list": [ 
        { 
            "root_id": "", 
            "parent_id": "", 
            "open_chat_id": "oc_b74c59c68d0f2d0ac65846272499d651", 
            "msg_type": "image", 
            "open_id": "", 
            "open_message_id": "be1000265b014075a869134b20c87633", 
            "is_mention": false, 
            "image_key": "99295878-5e85-41a3-bb00-0ad63b5b156d", 
            "image_url": "https://oapi-staging.zjurl.cn/open-apis/api/v2/file/f/99295878-5e85-41a3-bb00-0ad63b5b156d", 
            "create_time": 1550044148
       }, 
       { 
            "root_id": "", 
            "parent_id": "", 
            "open_chat_id": "oc_b74c59c68d0f2d0ac65846272499d651", 
            "msg_type": "text", 
            "open_id": "", 
            "open_message_id": "om_a96c620f2aa036e3c08abebaec7f09d1", 
            "is_mention": false, 
            "text": "文本1", 
            "create_time": 1550044749
       }, 
       { 
            "root_id": "", 
            "parent_id": "", 
            "open_chat_id": "oc_b74c59c68d0f2d0ac65846272499d651", 
            "msg_type": "post", 
            "open_id": "", 
            "open_message_id": "om_5d5b1732aa9b997dff23d63146427bb2", 
            "is_mention": false, 
            "text": "\u003cp\u003e富文本内容\u003c/p\u003e\u003cfigure\u003e\u003cimg src=\"http://p2.pstatp.com/origin/dad90010d9c8fc72f0b0\" origin-width=\"888\" origin-height=\"888\"/\u003e\u003c/figure\u003e", 
            "title": "富文本", 
            "create_time": 1550044772
       } 
       ] 
   } 
}
```

## Message Read 
This event is triggered when the user reads a message from a bot.

**Callback example:**
```json
{
    "ts": "1502199207.7171419", // the time the event is sent; this is usually close to the time the event is generated 
    "uuid": "bc447199585340d1f3728d26b1c0297a",  // unique event ID
    "token": "41a9425ea7df4536a7623e38fa321bae", // authentication token 
    "type": "event_callback", // event callbacks always use event_callback
	"event": {
		"app_id": "cli_9c8609450f78d102",
		"open_chat_id": "oc_e520983d3e4f5ec7306069bffe672aa3",
		"open_id": "ou_2d2c0399b53d06fd195bb393cd1e38f2",              
		"open_message_ids": ["om_dc13264520392913993dd051dba21dcf"],   // list of messages read
		"tenant_key": "xxx",
		"type": "message_read"
	}
}
```
