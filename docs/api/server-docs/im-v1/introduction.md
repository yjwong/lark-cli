---
title: "Overview"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/introduction"
service: "im-v1"
resource: "introduction"
section: "Messaging"
updated: "1717574880000"
---

# Overview
## Services
Lark provides a series of APIs for instant messaging (IM) scenarios. By creating a Lark bot and calling Messenger APIs, you can do the following:
- Send different types of messages via Lark, such as text, rich text, images, files, interactive message cards, videos, audios, and emojis.
- Manage Lark groups, such as creating a group and adding a user or a bot to a group.

For more application scenarios, refer to the customer case:
- Huazhu: [From Notification to Efficient Collaboration, Lark Uniform Alarm Solution Guarantees Services](https://open.larksuite.com/solutions/detail/alert)
- NIO: [Lark Facilitates Project Management, Making Collaboration Much Easier](https://open.larksuite.com/solutions/detail/project)

###  How to integrate

|  | Step | Description |
| --- | --- | --- |
| 1 | Create an app | - For creating a custom app, see Custom app development process. -   For creating a store app, see Develop and publish store apps. |
| 2 | Call APIs to operate messages or groups | Before calling APIs, you need to obtain the access token and enable corresponding permission scopes. For details, refer to How to call server-side APIs. You can debug these APIs in [API Explorer![API en.svg](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/77417553b88af927342054c2a7ca02b9_ZLfujqpD69.svg?height=24&lazyload=true&width=139)](https://open.larksuite.com/api-explorer?from=guide). For usage methods, see API Explorer Guide. |
| 3 | Listen on events to obtain changes in messages or groups | You need to apply for relevant permission scope before listening on events. For details, see Event subscription overview. | ### Development tutorial

| Quick start |  |
| --- | --- |
| Bot creates a group and sends alarms automatically ![14机器人自动拉群报警.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/32d3403236206fb1d8d837adfdace264_BLCbiE1mjG.png?height=400&lazyload=true&width=752) | Send mass messages to specified departments ![13将企业组织架构同步到Larkcn.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/8acc06dd2dddd013321aaff68d9be544_99kA3JoDQx.png?height=400&lazyload=true&width=752) |
| Bot sends a welcome message to new group members ![13将企业组织架构同步到Larkcn.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/a930de540b9fc01bab32937c2fe67f95_1EG3i8YIlA.png?height=400&lazyload=true&width=752) | Chat-based interactive bots ![13将企业组织架构同步到Larkcn.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/64260520e74e8564e017476ad35256d8_Azq1dfnMya.png?height=400&lazyload=true&width=752) |
| Send interactive message cards (approval cards) ![13将企业组织架构同步到Larkcn.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/f549cc39058fed469396671ba251a1c9_ll5E8mleX0.png?height=400&lazyload=true&width=752) |  | ## Resources

Messenger business domains are open around resources. The following figure illustrates the relationship among resources:

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/35518a347fe1a1a2f48c52d864d5bd79_bvem0tBiZL.png?height=1854&lazyload=true&width=1640)

User resources are within the Contacts service. Refer to Contacts documentation for details.

Messenger related resources are describes as follows:

| Message | Message sent via IM service, with text, rich text, cards, group cards, name cards, images, videos, and files supported. |
| --- | --- |
| Message - Images | A static resource carrying image information in Lark. |
| Message- Files | A static resource carrying file information in Lark. |
| Message - Message cards | A message type that contain rich text, images, and interactive components. |
| Message - Message reactions | Emotions response to messages. |
| Group message | It refers to the chat groups in Lark. |
| Group - group member | A collection of members in a group (including users and bots), which describes the relationship between a group and its members. |
| Group - group announcement | An announcement file in a group. | Below are details for fields, methods and events of each resource.

> "Store" represents Store Apps; "Custom" represents Custom Apps. For details, please refer to Application Type Description.

### Resource: Message
View Resource Fields and Examples.

#### Methods
| `Send messages `POST` /open-apis/im/v1/messages > Sends messages to a specific user or chat, with text, rich text, cards, group cards, name cards, images, videos, audios, files, and emojis supported. ` | Send messages as an app | `tenant_access_token` | **✓** | **✓** |
| --- | --- | --- | --- | --- |
| `Send messages in batches` `POST` /open-apis/message/v4/batch_send/ > Sends messages to multiple users or departments. | Send batch messages to multiple users Send batch messages to members from one or more departments | `tenant_access_token` | **✓** | **✓** |
| `Reply to messages` `POST` /open-apis/im/v1/messages/:message_id/reply > Replies to a specific message, with text, rich text, cards, group cards, name cards, images, videos, audios, files, and emojis supported. | Send messages as an app | `tenant_access_token` | **✓** | **✓** |
| `Recall messages` `DELETE` /open-apis/im/v1/messages/:message_id >The bot recalls messages it sent or the group owner recalls messages within a group. | Recall message Send messages as an app | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Query the read status of a message` `GET` /open-apis/im/v1/messages/:message_id >Queries the read status of a message. | Read messages in private and group chats | `tenant_access_token` | **✓** | **✓** |
| `Read chat history` `GET` /open-apis/im/v1/messages >Reads chat history for private and group chats. | Read messages in private and group chats | `tenant_access_token` | **✓** | **✓** |
| `Obtain resource files in messages` `GET` /open-apis/im/v1/messages/:message_id/resources/:file_key >Obtain resource files in messages, including audios, videos, images, and files. Emoji resources cannot be downloaded, and the resource files for download cannot exceed 100 MB. | Read messages in private and group chats | `tenant_access_token` | **✓** | **✓** |
| `Obtain content of a specific message` `GET` /open-apis/im/v1/messages/:message_id >Query message content with message_id. | Read messages in private and group chats | `tenant_access_token` | **✓** | **✓** | #### Event list

| `Receive messages` | This event is triggered when the bot receives a message sent by a user. | Obtain private messages sent to the bot Read all messages in associated group chat Obtain group messages mentioning the bot | im.message.receive_v1 | **✓** | **✓** |
| --- | --- | --- | --- | --- | --- |
| `Message read` | This event is triggered when a user reads a private chat message sent by the bot. | Read messages in private and group chats | im.message.message_read_v1 | **✓** | **✓** | ### Resource: Message - Images
View Resource Fields and Examples.

#### Methods
| `Upload images `POST` /open-apis/im/v1/images >Upload images, with JPEG, PNG, and WEBP formats supported. ` | Read and upload images or other files  | `tenant_access_token` | **✓** | **✓** |
| --- | --- | --- | --- | --- |
| `Download images` `GET` /open-apis/im/v1/images/:image_key >Download image resources. Only images of the message type uploaded by an app itself can be downloaded. | None | `tenant_access_token` | **✓** | **✓** | ### Resource: Message - Files
View Resource Fields and Examples.

#### Methods
| `Upload files `POST` /open-apis/im/v1/files >Upload files, with videos, audios, and common file types supported. ` | Read and upload images or other files  | `tenant_access_token` | **✓** | **✓** |
| --- | --- | --- | --- | --- |
| `Download files` `GET` /open-apis/im/v1/files/:file_key >Download files. Only files uploaded by an app itself can be downloaded. | None | `tenant_access_token` | **✓** | **✓** | ### Resource: Message - Message cards
View Resource Fields and Examples.

#### Methods
| `Update app messages `PATCH` /open-apis/im/v1/messages/:message_id >Update the content of message cards sent by an app. ` | Update message Send messages as an app | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| --- | --- | --- | --- | --- |
| `Delay message card update` `POST` /open-apis/interactive/v1/card/update > Delay the update of message cards after interaction with users. | None | `tenant_access_token` | **✓** | **✓** |
| `Send temporary card messages` `POST` /open-apis/ephemeral/v1/send > A bot send message cards visible to specific users in a group chat. | None | `tenant_access_token` | **✓** | **✓** |
| `Delete temporary messages` `POST` /open-apis/ephemeral/v1/delete > Delete temporary message cards of specific users in a group chat. Temporary card messages can be deleted explicitly via this API, and will leave no trace once deleted. | None | `tenant_access_token` | **✓** | **✓** | ### Resource: Message - Message reactions
View Resource Fields and Examples

#### Methods
| `Add a reaction for a message `POST` /open-apis/im/v1/messages/:message_id/reactions > Adds a reaction of the specified type for a specified message. ` | Send and delete message reaction | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| --- | --- | --- | --- | --- |
| `Obtain a reaction for a message` `GET` /open-apis/im/v1/messages/:message_id/reactions > Obtains a list of reactions of the specified type for a specified message. | View message reaction Read messages in private and group chats | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Delete a reaction for a message` `DELETE` /open-apis/im/v1/messages/:message_id/reactions/:reaction_id > Deletes a reaction for a specified message. | Send and delete message reaction | `tenant_access_token` `user_access_token` | **✓** | **✓** | #### Event list
| ` Add a reaction for a message` | This event is triggered when a reaction for a message is added. | View message reaction Read messages in private and group chats | im.message.reaction.created_v1 | **✓** | **✓** |
| --- | --- | --- | --- | --- | --- |
| `Delete a reaction for a message` | This event is triggered when a reaction for a message is deleted. | View message reaction Read messages in private and group chats | im.message.reaction.deleted_v1 | **✓** | **✓** |
