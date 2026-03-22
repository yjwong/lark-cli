---
title: "Send messages in batches"
url: "https://open.larksuite.com/document/ukTMukTMukTM/ucDO1EjL3gTNx4yN4UTM"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/message/v4/batch_send/"
section: "Messaging"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "im:message"
  - "im:message:send_as_bot"
  - "im:message:send_multi_users"
  - "im:message:send_multi_depts"
updated: "1717574914000"
---

# Send messages in batches

Send a message to multiple users or departments.

> **Note:** **Notes:**
>  - The bot ability needs to be enabled for the app.
>  - You must have the **Read and send messages in private and group chats** permission scope or have the **Send messages as an app** permission.
>  - You have at least one permission to send messages in batches:
>   	- To send batch messages to users, you must have the **Send batch messages to multiple users** permission scope.
>   	- To send batch messages to a department, you must have the **Send batch messages to members from one or more departments** permission scope.
>   - To send batch messages to the users or departments, the bot must be available for the users or departments.
> - Messages sent by calling this API **cannot be updated or replied**.
> - This API can send batch messages to users only, but cannot send batch messages to groups.
> - This API is an asynchronous API, which incurs a delay. Messages that are to be sent by each app are processed in sequence. Please arrange the batch sending scope and sequence as required. For the scenario of sending a message to a single user, please use Send Message.
> - Each app can send up to 500,000 messages per day by calling this API.

## Request
| HTTP URL | https://open.larksuite.com/open-apis/message/v4/batch_send/ |
| --- | --- |
| HTTP Method | POST |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any required permissions.Apply for corresponding optional permissions according to the sending target. | `im:message` `im:message:send_as_bot` `im:message:send_multi_users` `im:message:send_multi_depts` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about how to access and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body
Parameter | Type | Required | Description |  
--|--|--|--|  
msg_type | string | Yes | Type of the message. Multiple message types are supported. For more information, please see the "**Sample of message type and content**" part. |
content | object | No | Content of the message. Multiple types of message content are supported, except card messages. For more information, please see the "**Sample of message type and content**" part. |
card | object | No | Content of the card message. **Note**: either the card parameter or the content parameter must be specified. |
department_ids | list | No | Custom department IDs and open_department_ids are supported. And the list length is no more than 200.  **Note**: the message is also sent to the members of all the sub-departments under a department.**Example value:**["3dceba33a33226","d502aaa9514059", "od-5b91c9affb665451a16b90b4be367efa"]|
open_ids | list | No | List of open IDs of users. The list can contain up to 200 IDs.  For details, refer to How to get Open ID?**Example value:** ["ou_18eac85d35a26f989317ad4f02e8bbbb","ou_461cf042d9eedaa60d445f26dc747d5e"]|
user_ids | list | No | List of user IDs of the users. The list can contain up to 200 IDs. For details, refer to How to get User ID? **Example value:** ["7cdcc7c2","ca51d83b"]|
union_ids | list | No | List of union IDs of the users. The list can contain up to 200 IDs. For details, refer to How to get Union ID?**Example value:**["on_cad4860e7af114fb4ff6c5d496d1dd76","on_gdcq860e7af114fb4ff6c5d496dabcet"]| > **Note:** - Content, card field must choose one of them to fill in.
> - At least one of the department_ids, open_ids, user_ids, and union_ids fields must be specified.

#### Sample of message type and content
|     msg_type    |          Sample content        | Description |
| --------- | --------------- | -------   | ----------- | --------- |
|text  |{"content": { "text": "Text message that you need to send" } }|  |
|image  |{"content": { "image_key": "xxx-xxx-xxx-xxx-xxx" } }| The image key is obtained through the Upload Image API. |
|post  | {"content": { "post":{ POST_CONTENT }}} | For the post content format, refer to Sent message content. |
|share_chat  | {"content":{"share_chat_id": "oc_xxx"}} | Shares a group card. |
|interactive  | {"card":{...}} | Card messages. For the construction of a card message, refer to Message Card Builder. Note that the card field, instead of the content field, is used in card messages.| ### Request body example
> **Note:** The content of batch messages is formatted using JSON and therefore doesn't need escaping.

#### Text message
```json
{
    "department_ids": [
        "3dceba33a33226",
        "d502aaa9514059"
    ],
    "open_ids": [
        "ou_18eac85d35a26f989317ad4f02e8bbbb",
        "ou_461cf042d9eedaa60d445f26dc747d5e"
    ],
    "user_ids": [
        "7cdcc7c2",
        "ca51d83b"
    ],
    "union_ids": [
         "on_cad4860e7af114fb4ff6c5d496d1dd76",
         "on_gdcq860e7af114fb4ff6c5d496dabcet"
    ],
    "msg_type": "text",
    "content": {
        "text": "test content"
    }
}
```

#### Rich text message

```json 
{
	"open_ids": [
		"ou_a18fe85d22e7633852d8104226e99eac",
		"ou_9204a37300b3700d61effaa439f34295"
	],
	"department_ids": [
		"od-5b91c9affb665451a16b90b4be367efa"
	],
	"msg_type": "post",
	"content": {
		"post":{
			"zh_cn": {
				"title": "Title",
				"content": [
					[{
						"tag": "text",
						"text": "The first row"
					},
					{
						"tag": "a",
						"href": "http://www.larksuite.com",
						"text": "Lark"
					}
					]
				]
			},
			"en_us": {
				"title": "I am a title",
				"content": [
					[{
						"tag": "text",
					"text": "first line"
					},
					{
						"tag": "a",
						"href": "http://www.larksuite.com",
						"text": "Lark"
					}
					]
				]
			}
		}
	}
} 
``` 
#### Image

```json 
{
    "open_ids": [
        "ou_a18fe85d22e7633852d8104226e99eac",
        "ou_9204a37300b3700d61effaa439f34295"
    ],
    "department_ids": [
        "od-5b91c9affb665451a16b90b4be367efa"
    ],
    "msg_type": "image",
    "content": {
        "image_key": "img_v2_8e5a80ec-4d94-4bb2-84d5-e25d6c28dacj"
    }
} 
``` 

#### Group sharing message

```json 
{
    "open_ids": [
        "ou_a18fe85d22e7633852d8104226e99eac",
        "ou_9204a37300b3700d61effaa439f34295"
    ],
    "department_ids": [
        "od-5b91c9affb665451a16b90b4be367efa"
    ],
    "msg_type": "share_chat",
    "content": {
        "share_chat_id": "oc_84983ff6516d731e5b5f68d4ea2e1da5"
    }
} 
``` 
#### Card message
> **Note:** Note that the `card` field, instead of the `content` field, is used in the request body of card messages.

- Using the card JSON
```json 
{
	"open_ids": [
		"ou_a18fe85d22e7633852d8104226e99eac",
		"ou_9204a37300b3700d61effaa439f34295"
	],
	"department_ids": [
		"od-5b91c9affb665451a16b90b4be367efa"
	],
	"msg_type": "interactive",
	"card": {
		"config": {
			"wide_screen_mode": true
		},
		"elements": [
		{
			"tag": "div",
			"fields": [
			{
				"is_short": true,
				"text": {
					"tag": "lark_md",
					"content": "** Ticket source: **\nEvent and repair reports"
				}
			},
			{
				"is_short": true,
				"text": {
					"tag": "lark_md",
					"content": "**Ticket type: **\nAgent ticket"
				}
			}
			]
		}
		],

		"header": {
			"template": "turquoise",
			"title": {
				"content": "Show off your favorite books and win reading gifts"
				"tag": "plain_text"
			}
		}
	}
} 
``` 

- Using the card template
```json 
{
	"open_ids": [
		"ou_a18fe85d22e7633852d8104226e99eac",
		"ou_9204a37300b3700d61effaa439f34295"
	],
	"department_ids": [
		"od-5b91c9affb665451a16b90b4be367efa"
	],
	"msg_type": "interactive",
	"card": {
        "type": "template", 
        "data": { 
            "template_id": "ctp_xxxxxx",
            "template_variable": {
              "var_xxx": "xxxxxx"
            }
        }
	}
} 
``` 

## Response
### Response body
Parameter |Type| Description| 
-- |--| -- 
|code|int|Error code. A value other than 0 indicates failure.
|msg|string|Error description
data | - | -
  ∟message_id |string|Batch message task ID, used to uniquely identify a task for sending messages in batches
  ∟invalid_department_ids |list| List of invalid department IDs
  ∟invalid_open_ids |list| List of invalid open IDs
  ∟invalid_user_ids |list| List of invalid user IDs. 
  ∟invalid_union_ids |list| List of invalid union IDs

### Response body example
```json
{
    "code": 0,
    "msg": "ok",
    "data":{
      "invalid_department_ids": [
          "d502aaa9514059"
      ],
      "invalid_open_ids": [
          "ou_456e168d61cec276083b357f7bd3f1491",
          "ou_f8cbdb26fb2e4eda075e003381a102a41"
      ],
      "invalid_user_ids": ["7cdcc7c22"],
      "invalid_union_ids": [
         "on_cad4860e7af114fb4ff6c5d496d1dd76",
         "on_gdcq860e7af114fb4ff6c5d496dabcet"
      ],
      "message_id": "bm-d4be107c616aed9c1da8ed8068570a9f"
      }
}
```

### Error code
HTTP status code| Error code| Description| Troubleshooting suggestions
-- |--| -- | --
200| 10002| Create or send message fail.| Failed to create or send a message. Please refer to the error message returned by the request to check the message format.
200| 10003| Invalid parameter.| The request parameter is missing or incorrect. For more error information, please refer to the error message returned by the request.
200| 11101| Open_id not exist.| Open user ID invalid.
200| 11203| Send employee ids permission denied.|Sending message by user_id failed. Please make sure that in [Developer Console](https://open.larksuite.com/app) > **Application Details Page** > **Permission & Scopes** > **API Scopes**, you have activated **Send batch messages to multiple users** permission.
200| 11204| Send department list permission denied.| Sending message via department_ids failed. Please make sure that in [Developer Console](https://open.larksuite.com/app) > **Application Details Page** > **Permission & Scopes** > **API Scopes**, you have activated **Send batch messages to members from one or more departments** Permission.
200| 11205| App do not have bot.| Bot ability is not enabled. The bot ability can be added on the [Developer Console](https://open.larksuite.com/app) -> Features -> Add Features, and it will take effect after a new version is released.
200| 11207| App is not usage.| The app is not available under the current tenant, please check the status of the app in the [Developer Console](https://open.larksuite.com/app). 
200| 11209| App not exist.| App not exist.
200| 11210| App usage info not exist.|The application is not installed under the current tenant and cannot be used. Contact the tenant administrator to confirm the installation status of the application.
200| 11223| Do not have im write permission.| The current operator does not have permission to send messages, please activate the permission and request again.
200| 11227| Image key not exist.| Image Key invalid.
200| 11231| Chat not found.| The requested group chat does not exist.
200| 11235| This group is set to only admin can mention @All.| @All is prohibited from the current group chat.
200| 11246| Create message content fail.| Failed to create the message content. For more error information, please refer to the error message returned by the request.
200| 11247| Internal send message trigger rate limit.| The total number of messages sent by the app through this API has exceeded the daily limit of 500,000.
200| 11312| Bot messages do not pass the audit.| The message content is invalid, please check the message content.

For details, refer to Server-side error codes.
