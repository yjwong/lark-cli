---
title: "Received message content"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/im-v1/message/events/message_content"
service: "im-v1"
resource: "message"
section: "Messaging"
updated: "1717574909000"
---

# Received message content

> **Note:** This document describes the "content" field of different types of messages obtained from the "Read chat history", "Obtain content of a specific message", and other API.

## Example

Content is of the String type and indicates message content in JSON format.

By using a text-type message as an example, the example is as follows:

```json 
{
    "content": "{\"text\":\"Text message\"}"
}
```

## JSON structures for different message types

### Text

```json 
{
    "text": "@_user_1 Text message"
}

```
> **Note:** - In the text message, the format of the hyperlink is `[hyperlink text](hyperlink address)`, such as: `[Lark Open Platform](https://open.larksuite.com)`. In particular, for hyperlinks of the mailbox type, the format is `[mailbox text](mailto:mail address)`.
> - The @ in the message will be replaced with content in the form of "@_user_x", which represents the serial number of the @ user or bot. For example, the value of the third @-received member is "@_user_3". The detailed information of the @ user or bot can be obtained in the message mentions field based on the serial number. For details, see Resource Introduction.
> - Text styles such as bold, underline, etc. will be ignored, and only the text content will be displayed.

### Rich text post
> **Note:** The content when the message is obtained (including the message content in the response body of the sent message) is **not completely consistent with the content when it is sent**. The `md` tag only supports sending. This tag will not be included when getting the message content, and will be converted into other tags based on the content in `md`; in addition, quote, ordered and unordered lists will degenerate into text tags for output.

```json 
{
    "title": "Title",
    "content":[
        [
            {
                "tag":"text",
                "text": "The first row:",
                "style": ["bold", "underline"]
            },
            {
                "tag":"a",
                "href":"http://www.larksuite.com",
                "text": "Hyperlink",
                "style": ["bold", "italic"]
            },
            {
                "tag":"at",
                "user_id":"@_user_1",
                "user_name":"",
                "style": []
              
            }
        ],
        [
            {
                "tag":"img",
                "image_key":"img_47354fbc-a159-40ed-86ab-2ad0f1acb42g"
            }
        ],
        [
            {
                "tag":"text",
                "text": "The second row:",
                "style": ["bold", "underline"]
            },
            {
                "tag":"text",
                "text": "Text test",
                "style": []
            }
        ],
        [
            {
                "tag":"img",
                "image_key":"img_47354fbc-a159-40ed-86ab-2ad0f1acb42g"
            }
        ],
        [
            {
                "tag":"media",
                "file_key": "file_v2_0dcdd7d9-fib0-4432-a519-41d25aca542j",
                "image_key": "img_7ea74629-9191-4176-998c-2e603c9c5e8g"
            }
        ],
        [
            {
                "tag": "emotion",
                "emoji_type": "SMILE"
            }
        ],
        [
            {
                "tag": "hr"
            }
        ],
        [
            {
                "tag": "code_block",
                "language": "GO",
                "text": "func main() int64 {\n    return 0\n}"
            }
        ]
    ]
}

```

 **Tags and parameters supported by rich text**:

 **text** 

| Field | Type | Description |
| --- | --- | --- |
| text | string | Text content |
| un\_escape | bool | Whether to decode with unescape |
| style | []string | Bold, underline, strikethrough and italic style of the text content, the optional values are `bold`, `underline`, `lineThrough` and `italic` respectively, if there is no style, it is an empty list | **a** 

| Field | Type | Description |
| --- | --- | --- |
| text | string | Text content displayed by the link |
| href | string | Link address |
| style | []string | Bold, underline, strikethrough and italic style of the text content, the optional values are `bold`, `underline`, `lineThrough` and `italic` respectively, if there is no style, it is an empty list | **at** 

| Field | Type | Description |
| --- | --- | --- |
| user\_id | string | The serial number of the @ user or bot. For example, the value of the third @-received member is "@_user_3"; the detailed information of the @ member can be obtained according to the serial number in the message mentions field. For details, see Resource Introduction.|
| user\_name | string | User name |
| style | []string | Bold, underline, strikethrough and italic style of the text content, the optional values are `bold`, `underline`, `lineThrough` and `italic` respectively, if there is no style, it is an empty list | **img** 

| Field | Type | Description |
| --- | --- | --- |
| image\_key | string | Unique identifier of the image | **media** 
 
| Field | Type | Description |
| --- | --- | --- | --- | 
| file\_key | string | Unique identifier of the media file |
| image\_key | string | Unique identifier of the media cover image | **emotion** 

| Field | Type | Description |
| --- | --- | --- | --- | 
| emoji_type | string | Emoji type. For some optional values, please refer to Emoji copy | **code_block**

| Field | Type | Description |
| --- | ---| --- |
| language | string | Code block language, supports PYTHON, C, CPP, GO, JAVA, KOTLIN, SWIFT, PHP, RUBY, RUST, JAVASCRIPT, TYPESCRIPT, BASH, SHELL, SQL, JSON, XML, YAML, HTML, THRIFT, etc. |
| text | string | Code block content | #### hr

Rich text supports the hr tag, indicating a horizontal rule, or divider. It does not accept any other parameters.

### Image

```json 
{
    "image_key": "img_4adb3cc3-902b-4187-b0f1-842f67fd017g"
}

```

### File

```json 
{
    "file_key": "75235e0c-4f92-430a-a99b-8446610223cg",
    "file_name": "test.txt" // File name
}

```

### Folder

```json 
{
    "file_key": "75235e0c-4f92-430a-a99b-8446610223cg",
    "file_name": "folder" //Folder name
}

```

### Audio

```json 
{
    "file_key": "75235e0c-4f92-430a-a99b-8446610223cg", // File key
    "duration": 2000                                    // Duration in ms
}

```

### Media

```json 
{
    "file_key": "75235e0c-4f92-430a-a99b-8446610223cg",  // File key
    "image_key": "img_xxxxxx",                           // Key of the video cover image
    "file_name":"Test video.mp4",                            // File name
    "duration": 2000                                     // Video duration in ms
}

```

### Sticker

```json 
{
    "file_key": "75235e0c-4f92-430a-a99b-8446610223cg"
}

```

### Interactive
> **Note:** **Note**: The card structure is inconsistent with the card JSON obtained in the card building tool, and returning the original card JSON is not currently supported.

For the description of each field of the card structure, refer to Understand the card structure.
```json 
{
    "title": "Title",
    "elements": [
        [
            {
                "tag": "button",
                "text": "Primary button",
                "type": "primary"
            },
            {
                "tag": "button",
                "text": "Secondary button",
                "type": "default"
            },
            {
                "tag": "button",
                "text": "SOS button",
                "type": "danger"
            }
        ],
        [
            {
                "tag": "a",
                "href": "https://www.larksuite.com",
                "text": "Lark"
            },
            {
                "tag": "text",
                "text": "Integrates functions such as instant communication, calendar, audio and video conferencing, Docs, Drive, and Workplace to meet your organizational and individual needs."
            },
            {
                "tag": "at",
                "user_id": "@_user_1",
                "user_name": ""
            },
            {
                "tag": "text",
                "text": "More efficient and enjoyable"
            }
        ],
        [
            {
                "tag": "hr"
            }
        ],
        [
            {
                "tag": "text",
                "text": "Image title"
            },
            {
                "tag": "img",
                "image_key": "img_acd8a194-3e63-49ca-bcf6-224624457a3g"
            }
        ],
        [
            {
                "tag": "note",
                "elements": [
                    {
                        "tag": "img",
                        "image_key": "img_acd8a194-3e63-49ca-bcf6-224624457a3g"
                    },
                    {
                        "tag": "text",
                        "text": "Notes"
                    }
                ]
            }
        ],
        [
            {
                "tag": "text",
                "text": "Deeply integrates frequently used office tools so that company members can communicate and collaborate efficiently in one place."
            },
            {
                "tag": "img",
                "image_key": "img_acd8a194-3e63-49ca-bcf6-224624457a3g"
            }
        ],
        [
            {
                "tag": "text",
                "text": "You can also easily communicate, interact, and collaborate on your mobile devices, and your mobile phone and computer will be synchronized anytime and anywhere."
            },
            {
                "tag": "select_static",
                "options": [
                    "Option 1",
                    "Option 2",
                    "Option 3",
                    "Option 4"
                ],
                "placeholder": "Default prompt message"
            }
        ],
        [
            {
                "tag": "text",
                "text": "With ISV product access and independent development of companies, you can better interface with your existing system and meet the needs of different organizations."
            },
            {
                "tag": "overflow",
                "options": [
                    "Open Lark App Directory",
                    "Open Lark developer documentation",
                    "Open the Lark official website"
                ]
            }
        ],
        [
            {
                "tag": "text",
                "text": "International authoritative security certification and information security management systems provide companies with full lifecycle security."
            },
            {
                "tag": "date_picker",
                "placeholder": "Please select a date",
                "initial_date": "2021-1-1"
            }
        ]
    ]
}

```

### Red packet

```json 
{
    "text": "[Red packet]"
}

```
### Calendar

#### Share_calendar_event
Used for event sharing cards.
```json 
{
    "summary": "Event Sharing Card Test",
    "start_time": "1608265395000", // Timestamp in ms
    "end_time": "1608267015000"    // Timestamp in ms
}

```

#### Calendar
Used for event invitation cards.
```json 
{
    "summary": "Event Invitation Card Test",
    "start_time": "1608265395000", // Timestamp in ms
    "end_time": "1608267015000"    // Timestamp in ms
}

```

#### General_calendar
Used for event transfer cards/event postscript/switching the event where the calendar is located.
```json 
{
    "summary": "Event Transfer Card Test",
    "start_time": "1608265395000", // Timestamp in ms
    "end_time": "1608267015000"    // Timestamp in ms
}

```

### Group card share_chat

```json 
{
    "chat_id": "oc_0dd200d32fdaxxxxxxxx32f76"
}

```

### Name card share_user

```json 
{
    "user_id": "ou_0dd200d32xxxxx6d2c2ef1ddb32f76" // open_id of the user
}

```

### System message

```json 
{
        "template": "{from_user} invited {to_chatters} to this chat.",
        "from_user": ["botName"],
        "to_chatters": ["Sam", "John", "Mary"]
}

```

### Location

```json 
{
    "name": "xx province, xx city",
    "longitude": "xxx.xxx",
    "latitude": "xxx.xxx"
}

```

### Video_chat

```json 

{
    "topic": "Video call message",
    "start_time": "6745784522794598413", // Timestamp in ms
}

```

### Todo

```json 
{
    "task_id": "acd096a5-a157-4b9d-80e2-5b317456f005",
    "summary": {"title":"","content":[[{"tag":"text","text":"drink water"}]]}, 
    "due_time": "1623124318000"
}
```

**Parameters supported by todo**:

| Field | Type | Description |
| --- | --- | --- |
| task_id | string | Task ID, use this ID to operate the task, please see Task Feature Overview for details |
| summary | post | Task title in [post]() format |
| due_time | string | millisecond timestamp for task deadline | ### Vote

```json 
{
    "topic": "vote test",
    "options": ["option 1","option 2","option 3"]
}
```
**Parameters supported by vote**:

| Field | Type | Description |
| --- | --- | --- |
| topic | string | voting topic |
| options | list | voting options | ### Merge_forward
```json
{
    "content": "Merged and Forwarded Message"
}
```

The message content has no practical meaning. You can call Obtain the content of the specified message to obtain the sub-messages in the merged and forwarded message.
