---
title: "FAQs"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/guide/faq"
service: "im-v1"
resource: "guide"
section: "Messaging"
updated: "1717574884000"
---

# FAQs

This article lists the problems that users may encounter when using the 「Messages and Groups」 APIs to help you solve problems quickly.
## Notes before development

**Q: How should I use Lark OpenAPI?** 

Lark OpenAPI relies on RESTful APIs to provide services, please refer to How to call server-side APIs.

**Q: What are the differences between a company custom app and a store app (ISV app)?**

-   A company custom app can be used only by users within the same tenant, which is recommended for a company to meet all its needs.
-   A store app (ISV app) can be used by multiple tenants registered in the App Store. You need to make an application to create a store app. For details, refer to [Apply to become an ISV](https://open.larksuite.com/isv/).
> **Note:** For different tenants with the same appID or clientID, different tenantKey are used to obtain tenantToken. botID is uniquely identified by using appID and tenantKey for different tenants.

For details about custom apps and store apps, refer to Custom apps and store apps.

**Q: How can I debug the codes before releasing them to the production environment?**

Lark OpenAPI relies on RESTful APIs to provide services. To make it easier for developers to debug and test APIs, Lark provides [API Explorer](https://open.larksuite.com/api-explorer) to debug the API. If the interface debugging is correct, the parameters are filled correctly.

# ID

**Q: How to obtain "App ID" and "App Secret" ?**

1.  Open the [Lark Open Platform homepage](https://open.larksuite.com/?lang=en-US), go to [Developer Console](https://open.larksuite.com/app?lang=en-US), and select your own bot.
1.  Click "Credentials & Basic Info" in the left navigation bar to view the app information.

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/86105f1c9fb1abd0565e321d0c0b09a5_YQ8JUTrQta.png?height=176&lazyload=true&width=709)

**Q: How to obtain a group ID (chat_id)?**

You can obtain chat_id by using the following open APIs:

1.  Create a group
1.  Search for groups visible to a user or bot

For details about the usage scenarios and steps of chat_id, refer to Group ID description.
> **Note:** The Search for groups visible to a user or bot API is applicable only to group chats.

# Bot

**Q: Why can't I find my bot?**

After applying for a bot, you need to do the following to find your bot:

1.  Enable the bot capability.

Open the [Lark Open Platform homepage](https://open.larksuite.com/?lang=en-US), go to [Developer Console](https://open.larksuite.com/app?lang=en-US), and select your own bot app. Select Features > Bot in the left navigation bar, and then click "Enable bot".

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/139e637f7da535d1e4c8896dd63b917f_Y7VKYEoFc1.png?height=699&lazyload=true&width=1535)

2.  Release the app, and then you can find it.

> **Note:** Be sure to select the correct bot and environment (BOE/online/overseas).

**Q: How can I configure the visibility of my bot (error message: the bot is invisible to the user)?**

Visibility of a bot can be configured when you release an app version on the console.

1. Open the [Lark Open Platform homepage](https://open.larksuite.com/?lang=en-US), go to [Developer Console](https://open.larksuite.com/app?lang=en-US), and select your own bot.

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/d67b0b9bddd609733a7784fc9145bb19_5qXRgeux8n.png?height=1686&lazyload=true&width=3704)

2. Select "Version Management & Release" from the left navigation bar, and then click "Create a version" in the upper right corner.

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/c7f8ab6b0d5d8869a7518ef96cc8a802_YkbWHCv5C3.png?height=842&lazyload=true&width=1907)

3.  Click "Edit" under "Availability" to edit the app visibility to a specific user.

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/ddcf32f12df8d76368ae84cd942e2937_aezfgUjggX.png?height=836&lazyload=true&width=1812)

4.  Click "Save" and "Submit for release". The visibility configuration takes effect immediately after the version is approved for release.

**Q: What is a Webhook custom rot? How can I use it to send messages?**

A **custom bot** in a group and the **bot app** developed by you are different in the following aspects:

- A custom bot can only be used to automatically send notifications in a group chat. It is unable to reply to messages @mentioning it sent by users, or obtain any user or tenant information.

- A custom bot can be used in an external group, while a bot app can be used only within an internal group.

-   For details about how to use a custom bot, refer to Custom bot guide.

-   For usage scenarios of a bot, refer to the development tutorial in Overview.

**Q: Why can't I see a chat box in my bot after I configure the bot in the Developer Console?**

Your app must be configured with an event callback URL before the dialog box will appear.

**Q: How can I make a bot @mention a user (@mention all or @mention a specific user)?**

You can use the `at` tag in plain text messages (text), rich text messages (post), or message cards (interactive) to @mention a user.

Sample requests are shown below:

**(1) @mention specific users or @mention all in plain text messages:**

Sample code for `at` tag

```json 
// @mention a specific user
Name //You can fill in open_id, union_id or user_id to @mention a specific user.
// @mention all
Everyone
``` 

Sample code for request `content`:

```json 
{
	"content": " {\"text\":\"Tom text content\"}"
}
``` 
**(2) @mention specific users or @mention all in rich text messages:**

Sample code for `at` tag

```json 
// @mention a specific user
{
	"tag": "at",
	"user_id": "ou_xxxxxxx",//You can fill in open_id, union_id or user_id to @mention a specific user.
	"user_name": "tom"
}

// at everyone
{
	"tag": "at",
	"user_id": "all",//Use "all" to @mention all.
	"user_name": "all"
}
``` 
Sample code for request `content`:

```json 
{
  "zh_cn": {
    "title": "Title",
    "content": [
      [
        {
          "tag": "text",
          "text": "The first line:"
        },
        {
          "tag": "at",
          "user_id": "ou_xxxxxx",//You can fill in open_id, union_id or user_id to @mention a specific user.
          "user_name": "tom"
        }
      ],
      [
        {
          "tag": "text",
          "text": "The second line:"
        },
        {
          "tag": "at",
          "user_id": "all",
          "user_name": "all"//Use "all" to @mention all.
        }
      ]
    ]
  }
}
``` 
**(3) @mention specific users or @mention all in message cards:**
You can use the @ tag of the Markdown module to @mention users. Sample code:

```json 
// @mention a specific user
	 //Use open_id to @mention a specific user
	 //Use user_id to @mention a specific user
	 //Use email to @mention a specific user
// at everyone
	
``` 

Sample code for request `content`:

```json 
{
  "config": {
    "wide_screen_mode": true
  },
  "header": {
    "title": {
      "tag": "plain_text",
      "content": "Card title"
    },
    "template": "blue"
  },
  "elements": [
    {
      "tag": "div",
      "text": {
        "content": "@mention all \n@mention a specific user",
        "tag": "lark_md"
      }
    }
  ]
}
``` 
**Note:**
-   For custom bots, please refer to Custom bot usage guide.
-   If the group owner enables the `"Only group owner and admin can @mention all"` permission, a bot can't @mention all if they are not the group owner or admin.

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/bc70cbdcecbc5cfbf0ee0a7969b7e1f1_e0AYJ6wr5x.png?height=920&lazyload=true&width=686)

**Q: Does the text parameter of the custom bot webhook request has a length requirement?**

We recommend that the length of a JSON string should not be more than 30 KB, the serialized Protobuf (pb) size should not be more than 100 KB, and the image size should be less than 10 MB.

**Q: Can I use the same custom bot in different groups?**

No. A custom bot can be used only in a single group.

# Message

**Q: Is there any size limit on a message to be sent?**

We recommend that the length of a JSON string should not be more than 30 KB, the serialized Protobuf (pb) size should not be more than 100 KB, and the image size should be less than 10 MB.

**Q: Can I change messages already pushed?**

- Message card (Interactive) type messages: refer to the Interaction module documentation to see how to change

- Text, rich text type messages: refer to Edit message documentation

Other types of messages do not support changes yet. 

**Q: Are there any restrictions on images uploaded via APIs?**

- Supported formats of images uploaded: JPEG, PNG, WEBP, GIF, TIFF, BMP, and ICO
- Supported formats of output images: PNG, JPG, HEIC, WEBP, and WEBP
- Resolution of images uploaded should be no more than 200 MB (width x height), and the image size should be no more than 150 MB.
- Maximum size of profile photos uploaded: 4096 x 4096

**Q: Can I send messages to external contacts via APIs?**

Not supported.

**Q: How can I @mention all when sending messages?**

Refer to Sent message content. Specify  to @mention all.

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/4e57141e280e919e1425a67577218415_4Izk01Y3Tp.png?height=699&lazyload=true&width=1498)

**Q: Is idempotence supported for sending messages?**

Not supported.

**Q: Can I obtain the resource files in messages sent by others via the API for obtaining resource files in messages?**

Yes. You can obtain the resource files in any message visible to you.

# Message card

**Q: Can I send message cards to multiple people in batches?**

Yes. You can refer to the API Send messages in batches.
Note:
- For the request parameters, the value of `msg_type` should be `interactive`, and `content` should be the `card:{}` object.
- Message cards sent in batches do not support update and callback interaction.

**Q: Can I @mention a specific user in a message card?**

Yes. For the specific syntax, refer to Use the markdown tag.

**Q: What is the timeout interval of the message card callback?**

The callback timeout interval of the message card is **3 sec**, which cannot be changed. If a response can't be returned within 3 seconds, {} can be returned first, and the token in the callback can be used to update the card asynchronously.

# Group

**Q: Can I search for external groups from the group list of a user?**

External groups can be searched via the [Search for groups visible to a user or bot]] API, and cannot be searched via the 「**Search for groups of a user**」 (older version).

The older versions are no longer maintained. We recommend that you use the new version.

# Event

**Q: Why did I receive an empty message event?**

A message event is empty when the message is recalled.

**Q: What is the difference between new and old versions of events?**

The new version is updated based on the old version and uses a new protocol.

These two versions are incompatible. The old version will still be available but not be maintained any more. We recommend that you use the new version.

# Troubleshooting

**Q: What are the reasons for error codes returned during API calls?**

1.  Error code descriptions and troubleshooting suggestions are provided in each API documentation.
1.  All API error codes are listed in Server-side error codes, which can also be found in this document.

**Q: I have requested a scope, but why is a no scope message displayed?**

1. After requesting the scope, you need to release an app version to make the scope configuration take effect.
2. To make scope configuration of an ISV app take effect, you need to not only release an app version but also authorize the scope in [Tenant Management Background](https://admin.larksuite.com/).

**Q: Rate-limiting policies have been triggered during API calls. How do the policies work?**

Rate-limiting policies are used to ensure the stability of background services.

For details, refer to Rate-limiting Policies - Server-side Docs.

**Q: What are the mapping relationship between an error code and a rate-limiting policy triggered?**

| Error code      | Rate-limiting policy                  | Suggestion                                                                                        |
| -------- | --------------------- | ----------------------------------------------------------------------------------------- |
| 230020   | The message sending API (V1) triggers throttling for sending messages in a group. | QPS for sending messages in a single group must be no more than 5. For details, refer to Rate-limiting policies. |
| 11232    | The message sending API (V4) triggers throttling of the IM system.   | Try again later.                                                                                    |
| 11233    | The message sending API (V4) triggers throttling for sending messages in a group.  | QPS for sending messages in a single group must be no more than 5. For details, refer to Rate-limiting policies. |
| 11247    | The API for sending messages in batches triggers daily quota limit.        | No more than 500,000 messages can be sent by a single app via this API. Please control the quantity of messages sent in batches.                                              |
| 99991400 | Triggers API rate-limiting.                | Too many requests. Frequency not higher than Rate-limiting policies                                                                            | **Q: When I call the API for sending messages, a message appeared indicating message content errors in the request. What should I do?**

Rules of converting JSON to strings for rich text messages are provided in Sent message content. Here are the key points:

1.  Conclude the entire JSON content with " ".
2.  Use "`\`" for escaping.
3.  If you want to implement the newline function in the message content, you should add escape on the basis of "\n". For details, refer to Sent message content .

Example:

```json
{
"receive_id":"14f76322",
"msg_type": "post",
"content":"{\"zh_cn\":{\"title\":\"Title\",\"content\":[[{\"tag\":\"text\",\"text\":\"First row:\"},{\"tag\":\"a\",\"href\":\"[http://www.larksuite.com](http://www.larksuite.com/)\",\"text\":\"Hyperlink\"},{\"tag\":\"at\",\"user_id\":\"14f76322\",\"user_name\":\"tom\"}]]}}"
}
 
``` 

4.  A string must not contain [unicode](https://zh.wikipedia.org/wiki/Unicode).
4.  It is recommended to use the [JSON compressing and escaping tool](https://www.sojson.com/yasuo.html).
4.  To send special characters, use Markdown syntax. For details, refer to: Lark Open Platform > Developer Documentation > Common Capability > Markdown module.

# Others

**Q: How do I distinguish overseas OpenAPI from domestic OpenAPI?**

They are distinguished by the domain name.

For overseas OpenAPI, the domain name is https://open.larksuite.com/.

Take the Send messages API as an example:

-   CN -url: https://open.larksuite.com/open-apis/im/v1/messages
-   US -url: https://open.larksuite.com/open-apis/im/v1/messages

**Q: How should I set contentType?**

You need to set contentType to "application/json; charset=utf-8". Other formats may not be supported.

**For more questions, refer to Developer Documentation - FAQs.**
