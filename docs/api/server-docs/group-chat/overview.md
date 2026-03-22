---
title: "Overview"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/group/overview"
section: "Group Chat"
updated: "1717574939000"
---

# Overview
## Resource: Group
View Resource Fields and Examples.

### Methods
| `Create a group `POST` /open-apis/im/v1/chats > Create a group and set the profile photo, name, and description for the group. ` | Create group | `tenant_access_token` | **✓** | **✓** |
| --- | --- | --- | --- | --- |
| `Obtain group information` `GET` /open-apis/im/v1/chats/:chat_id > Obtain basic information such as group name, description, profile photo, and owner ID. | View group information | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Update group information` `PUT` /open-apis/im/v1/chats/:chat_id > Update the group profile photo, name, description, and configuration, or assign new group owner. | Update group information | `tenant_access_token` | **✓** | **✓** |
| `Delete a group` `DELETE` /open-apis/im/v1/chats/:chat_id >The bot recalls messages it sent or the group owner recalls messages within a group. | Disband group | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Obtain the list of groups with the user or bot` `GET` /open-apis/im/v1/chats >Updates the content of message cards sent by an app. | View group information | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Search for groups visible to a user or bot` `GET` /open-apis/im/v1/chats/search >Delays the update of message cards after interaction with users. | View group information | `tenant_access_token` `user _access_token` | **✓** | **✓** | ### Event list

| `Group deletion` | This event is triggered when a group is deleted. | View group information | im.chat.disbanded_v1 | **✓** | **✓** |
| --- | --- | --- | --- | --- | --- |
| `Modify group configurations` | This event is triggered when the group configurations are modified, including assignment of new group owner and modification of basic group information and permission scopes. | View group information | im.chat.updated_v1 | **✓** | **✓** | ## Resource: Group - group member
View Resource Fields and Examples.

### Methods

| `Add a user or bot to a group chat `POST` open-apis/im/v1/chats/:chat_id/members > Adds a user or bot to a group chat. ` | Add and remove group member | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| --- | --- | --- | --- | --- |
| `Remove a user or bot from a group chat` `DELETE` /open-apis/im/v1/chats/:chat_id/members > Removes a user or bot from a group chat. | Add and remove group member | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `A user or bot joins a group chat` `PATCH` /open-apis/im/v1/chats/:chat_id/members/me_join > Replies to a specific message, with text, rich text, cards, group cards, name cards, images, videos, audios, files, and emojis supported. | Add and remove group member | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Obtain group member list` `GET` /open-apis/im/v1/chats/:chat_id/members >The bot recalls messages it sent or the group owner recalls messages within a group. | View group members | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Determine whether a user or bot is in a group` `GET` /open-apis/im/v1/chats/:chat_id/members/is_in_chat >Queries the read status of a message. | View group members | `tenant_access_token` `usert_access_token` | **✓** | **✓** | ### Event list

| `Bots added to a group` | This event is triggered when a bot is added to a group chat. | Subscribe to events of bot joining and leaving group | im.chat.member.bot.added_v1 | **✓** | **✓** |
| --- | --- | --- | --- | --- | --- |
| `Remove bot from a group chat` | This event is triggered when a bot is removed from a group chat. | Subscribe to events of bot joining and leaving group | im.chat.member.bot.deleted_v1 | **✓** | **✓** |
| `Users join a group` | This event is triggered when a new user joins a group chat. | View group members | im.chat.member.user.added_v1 | **✓** | **✓** |
| `Users leave a group` | The event is pushed when users leaves a group voluntarily or are removed from a group chat. | View group members | im.chat.member.user.deleted_v1 | **✓** | **✓** |
| `Revoke invite` | This event is triggered when the invitation of a user to a group is revoked. | View group members | im.chat.member.user.withdrawn_v1 | **✓** | **✓** | ## Resource: Group - group announcement
View Resource fields and examples.

### Methods
  
| `Obtain group announcement information `GET` /open-apis/im/v1/chats/:chat_id/announcement > Obtain group announcement information in a chat, with the same format as Docs. ` | View group announcement | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| --- | --- | --- | --- | --- |
| `Update group announcement information` `PATCH` /open-apis/im/v1/chats/:chat_id/announcement > Update group announcement information in a chat, with the same format as Docs. | Update group announcement | `tenant_access_token` `user_access_token` | **✓** | **✓** |
