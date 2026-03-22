---
title: "Resource introduction"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message-card/overview"
service: "im-v1"
resource: "message-card"
section: "Messaging"
updated: "1717574918000"
---

# Resource introduction
## Resource definition

Lark message cards are messages containing rich text, images, and interactive components.
With message cards, you can:
- Use the rich text style and image-text layout modules to send exquisite cards, such as eye-catching notifications and article lists with images and texts. In this way, it's easier for users to catch important information.
- Submit information with just one click on the card with the interactive components. In this way, users can perform system actions such as OA approval, vote statistics, and alarm processing in the chat window.

No client-side development experience is required to realize the above functions. Lark defines structured components and styles for card messages. Users can build exquisite and interactive message cards through JSON description on the server. For more information, see Introduction to message cards>>.

## Field description

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `content` | `string` | Yes | Message content. The data is the card structure in JSON format and is converted into String. For the description of each field of the card structure, please refer to Card Structure Introduction | ### Data example
The following sample code can be copied to the [message card builder](https://open.larksuite.com/tool/cardbuilder?from=howtoguide) for editing.

```json 
{
  "config": {
    "wide_screen_mode": true
  },
  "header": {
    "title": {
      "tag": "plain_text",
      "content": "A leave request requires your approval"
    }
  },
  "elements": [
    {
      "tag": "div",
      "fields": [
        {
          "is_short": true,
          "text": {
            "tag": "lark_md",
            "content": "**Applicant**\nWang Xiaolei"
          }
        },
        {
          "is_short": true,
          "text": {
            "tag": "lark_md",
            "content": "**Leave type:**\nAnnual leave"
          }
        },
        {
          "is_short": false,
          "text": {
            "tag": "lark_md",
            "content": ""
          }
        },
        {
          "is_short": false,
          "text": {
            "tag": "lark_md",
            "content": "**Duration:**\nFrom 2020-4-8 to 2020-4-10 (3 days in total)"
          }
        },
        {
          "is_short": false,
          "text": {
            "tag": "lark_md",
            "content": ""
          }
        },
        {
          "is_short": true,
          "text": {
            "tag": "lark_md",
            "content": "**Notes**\nReturn to hometown for a family emergency"
          }
        }
      ]
    },
    {
      "tag": "hr"
    },
    {
      "tag": "action",
      "layout": "bisected",
      "actions": [
        {
          "tag": "button",
          "text": {
            "tag": "plain_text",
            "content": "Approve"
          },
          "type": "primary",
          "value": {
            "chosen": "approve"
          }
        },
        {
          "tag": "button",
          "text": {
            "tag": "plain_text",
            "content": "Reject"
          },
          "type": "primary",
          "value": {
            "chosen": "decline"
          }
        }
      ]
    }
  ]
}
``` 

### Example of sending a message card
To send the approved message card shown in the preceding figure to a group, you can call the Send a message API. The cURL sample code is as follows:
```bash
curl --location --request POST 'https://open.larksuite.com/open-apis/im/v1/messages?receive_id_type=chat_id' \
--header 'Authorization: Bearer t-XXX' \
--header 'Content-Type: application/json; charset=utf-8' \
--data-raw '{
    "receive_id": "oc_84983ff6516d731e5b5f68d4ea2e1da5",
    "content": "{\"config\":{\"wide_screen_mode\":true},\"header\":{\"title\":{\"tag\":\"plain_text\",\"content\":\"A leave request requires your approval\"}},\"elements\":[{\"tag\":\"div\",\"fields\":[{\"is_short\":true,\"text\":{\"tag\":\"lark_md\",\"content\":\"**Applicant**\\nWang Xiaolei\"}},{\"is_short\":true,\"text\":{\"tag\":\"lark_md\",\"content\":\"**Leave type:**\\nAnnual leave\"}},{\"is_short\":false,\"text\":{\"tag\":\"lark_md\",\"content\":\"\"}},{\"is_short\":false,\"text\":{\"tag\":\"lark_md\",\"content\":\"**Time:**\\nApril 8, 2020 to April 10, 2020 (3 days)\"}},{\"is_short\":false,\"text\":{\"tag\":\"lark_md\",\"content\":\"\"}},{\"is_short\":true,\"text\":{\"tag\":\"lark_md\",\"content\":\"**Notes**\\nReturn to hometown for a family emergency\"}}]},{\"tag\":\"hr\"},{\"tag\":\"action\",\"layout\":\"bisected\",\"actions\":[{\"tag\":\"button\",\"text\":{\"tag\":\"plain_text\",\"content\":\"Approve\"},\"type\":\"primary\",\"value\":{\"chosen\":\"approve\"}},{\"tag\":\"button\",\"text\":{\"tag\":\"plain_text\",\"content\":\"Reject\"},\"type\":\"primary\",\"value\":{\"chosen\":\"decline\"}}]}]}",
    "msg_type": "interactive"
}'
``` 

### Preview a card

![Group 647.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/3b8c1a91dab728d66de7e6af4a3f34c6_RKPeA0PgXf.png?lazyload=true&width=1640&height=882)
