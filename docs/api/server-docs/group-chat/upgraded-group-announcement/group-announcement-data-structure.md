---
title: "Group announcement data structure"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/group/upgraded-group-announcement/data-structure/group-announcement-data-structure"
section: "Group Chat"
updated: "1739935015000"
---

# Group announcement data structure

Group announcements are documents within a group carried by Docs. Each group has only one group announcement, and each announcement is uniquely identified by a `chat_id`. Structurally, each group announcement is formed as a Block tree with multiple blocks in a parent-child relationship.

## Group Announcement Data Structure

The basic metadata structure for an `group_announcement` is as follows. You can obtain this information by calling the Obtain the basic information of a group announcement API.
```JSON
{
    "revision_id": int,    // Current version number of the group announcement
    "create_time": string,    // Timestamp (in seconds) when the group announcement was created
    "update_time": string,    // Timestamp (in seconds) when the group announcement was last updated
    "owner_id": string,    // ID of the group announcement owner
    "owner_id_type": string,    // Type of the group announcement owner's ID
    "modifier_id": string,    // ID of the latest modifier of the group announcement
    "modifier_id_type": string,    // Type of the latest modifier's ID
    "announcement_type": string    // Type of group announcement, either 'docx' for upgraded versions or 'doc' for old versions
}
```

### Parameter Descriptions

The specific descriptions of the parameters in the above structure are shown in the table below.

| Parameter Name | Data Type | Example Value | Description |
| --- | --- | --- | --- |
| revision_id | int | 42 | Current version number of the group announcement |
| create_time | string | 1609296809 | Timestamp (in seconds) when the group announcement was created |
| update_time | string | 1609296809 | Timestamp (in seconds) when the group announcement was last updated |
| owner_id | string | ou_7d8a6e6df7621556ce0d21922b676706ccs | The group announcement owner ID. The ID value corresponds to the ID type of the owner_id_type. |
| owner_id_type | string | open_id | Group announcement owner ID type **Optional values are:** * `user_id`：dentifies a user's identity in a tenant. The User ID of the same user in tenant A and tenant B is different. In the same tenant, a user's User ID is consistent in all applications (including store applications). User ID is mainly used to connect user data between different applications. * `union_id`：Identifies a user's identity under a certain application developer. The Union ID of the same user in applications under the same developer is the same, and the Union ID in applications under different developers is different. Through the Union ID, application developers can associate the identities of the same user in multiple applications. * `open_id`：Identifies a user's identity in a certain application. The Open ID of the same user in different applications is different. |
| modifier_id | string | ou_7d8a6e6df7621556ce0d21922b676706ccs | Last group announcement modifier ID. The ID value corresponds to the ID type of modifier_id_type. |
| modifier_id_type | string | open_id | ID type of the last group announcement modifier **Optional values are:** * `user_id`：Identifies a user's identity in a tenant. The User ID of the same user in tenant A and tenant B is different. In the same tenant, a user's User ID is consistent in all applications (including store applications). User ID is mainly used to connect user data between different applications. * `union_id`：Identifies a user's identity under a certain application developer. The Union ID of the same user in applications under the same developer is the same, and the Union ID in applications under different developers is different. Through the Union ID, application developers can associate the identities of the same user in multiple applications. * `open_id`：Identifies a user's identity in a certain application. The Open ID of the same user in different applications is different. |
| announcement_type | string | docx | Group announcement type **Optional values are:** * `docx`：Upgraded group announcement * `doc`：Old version group announcement | ## Block Data Structure

The data structure of group announcement blocks is the same as the block data structure in documents. Refer to the Block Data Structure in Documents for detailed information.
