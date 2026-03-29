---
title: "Group announcement overview"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/group/upgraded-group-announcement/group-announcement-overview"
section: "Group Chat"
updated: "1739935002000"
---

# Group announcement overview

This document introduces the basic concepts, limitations, access process, and API list of the group announcement capabilities.

> This document introduces the capabilities of upgraded group announcement. To understand the differences between upgraded group announcement and old group announcement features, see Group announcement.

## Basic concepts

There are two basic concepts in the Group announcement APIs: group announcement and block.

### Group announcement

Group announcements are announcement documents in a group. They are hosted by Upgrade Docs. Each group has only one group announcement, and each announcement is uniquely identified by a `chat_id`. To obtain the `chat_id` of a group, follow these methods:
- Create a group, get the chat_id of the group from the returned result.
- Call the Get the list of groups that the user or robot is in interface to query the chat_id of the group that the user or robot is in.
- Call the Search the list of groups visible to the user or robot to search for the chat_id of the group that the user or robot is in and the group that is open to the user or robot.

The basic metadata structure of a group announcement is as follows:

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

### Block

Within a group announcement, there are multiple types of paragraphs, each defined as a block. A block is the smallest building unit in a group announcement, serving as a structured component with a clear meaning. Blocks can take various forms, such as a text paragraph, a sheet, an image, or a Base. Each block has a unique `block_id` for identification.

Every group announcement has a root block called Page block. The `block_id` of the Page block is identical to the `chat_id` of the group it belongs to. In the data structure, the Page block forms a parent-child relationship with other blocks, where the Page block is the parent (referred to as Parent), and other blocks are children (referred to as Children). Parent-child relationships can also exist among other blocks.

**Block category**

From a functional perspective, blocks can be categorized into the following types. To know about the specific blocks, refer to BlockType enums.

| Functional category | Typical blocks | Example |
| --- | --- | --- |
| Text type | Page, Text, Heading, Bullet list, Ordered list, Code, Todo, etc. | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/b5c839e19f8709fc8b196ec5c41ceeb6_YQvFfOPuiZ.png?height=204&lazyload=true&width=709) |
| Data type | Bitable, Sheet, Mindnote, etc. |  |
| Visual type | Divider |  |
| Media type | Image, File, Iframe, etc. |  |
| Collaboration type | ChatCard |  |
| Container type | TableCell, GridColumn, Callout, View, QuoteContainer, etc. |  |
| Vertical type | Diagram (Flowchart & UML Diagram) |  |
| Auxiliary type | Table, Grid, etc. |  |
| Third-party type | ISV (Independent Software Vendor) block |  |
| Undefined type | / |  | **Parent-and-child relationships of block**
    
Parent-child relationships can exist among blocks. Before you use the Create blocks API, refer to the following rules to specify the parent block and child block.

**Rules for parent blocks**
    
Only blocks with the capacity to contain can serve as parent blocks to which child blocks can be added. These blocks include:

- **Text-type functional blocks:** Page, Text, Heading, Bullet, Ordered, Task, and Todo blocks.
- **Container-type functional blocks:** TableCell, GridColumn, Callout, and QuoteContainer blocks.
  
    
**Rules for child blocks**

The following blocks cannot be added as child blocks to a parent block:

  | Block | Description | | ------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | | Page | A group announcement has only one page block, which is automatically generated during group announcement creation. | | GridColumn | Can only be added through the **InsertGridColumnRequest** of Bath Update Blocks API.| | Mindnote | Not supported. | | TableCell | Can only be added through the **InsertTableRowRequest** or **InsertTableColumnRequest** operations of Bath Update Blocks interface. | | View | When adding a File block, a default view block is automatically generated.  | | Undefined | Invalid operation. | | Diagram | Not supported. | **Limitations for parent-child relations**

Other limitations for Table Cell, Grid Column, and Callout blocks are described below:
| Block | Restriction |
| --- | --- |
| Table Cell | It is not allowed to add the following blocks as child blocks to a Table Cell block:• Table• Sheet• Bitable• OKR |
| Grid Column | It is not allowed to add the following blocks as child blocks to a Grid Column block:• Grid• Bitable• OKR |
| Callout | Only the following blocks are allowed as child blocks to a Callout block:• Text• HeadingN• OrderedBullet• Task• Todo• Quote• QuoteContainer | ## Limitations

You can use group announcement API to operate on different types of blocks, including creating, reading, and editing block content. The support for different blocks varies in the group announcement APIs. For details, refer to the table below.

In the table, "/" represents that the corresponding operation is not required or has been covered in other capabilities, as follows:

- **Capability not required.** For example, the Divider block does not contain content, so there is no need to provide the ability to read or edit its content.
- **Capability alternatively supported.** For instance, for the TableCell block, you will directly create or delete the TableCell block while editing the content of the Table block. Therefore, the platform does not need to provide another creation capability for the TableCell block.
- **Capability covered in other open capabilities.** For example, for the Bitable block, you can create an empty Bitable block and then use the Base capabilities based on the returned token to use the corresponding read and write APIs. To understand Base capabilities, refer to Base method list.

  | Block | Create block | Read content | Edit content | |---|---|---|---| | Callout | Supported | Supported | Supported | | Table | Supported | Supported | Supported | | Text | Supported | Supported | Supported | | Divider | Supported | / | / | | Grid | Supported | Supported | Supported | | Iframe | Supported | Supported | Not supported | | ChatCard | Supported | / | / | | Image | Supported | / | / | | File | Supported | / | / | | TableCell | / | Supported | Supported | | GridColumn | / | Supported | Supported | | View | / | Supported | Not supported | | ISV | Supported | / | / | | Bitable | Supported | / | / | | Sheet | Supported | / | / | | Mindnote | Not supported | / | / | | Diagram | Not supported | Not supported | Not supported | | QuoteContainer | Supported | Supported | Supported | | Task | Not supported | / | / | | OKR | Supported | Supported | / | | OkrObjective | / | Supported | / | | OkrKeyResult | / | Supported | / | | OkrProgress | / | Supported | / | | AddOns | Supported | Supported | / | | JiraIssue | / | Supported | / | | WikiCatalog | Supported | / | / | | Board | Supported | Supported | / | | Undefined | / | / | / | ## Integration process
The process for integrating with the Group announcement API is illustrated in the diagram below. For details, refer to the "**Access process**" section in Group overview
![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/7c27953b2c7bfc8f8766b52721eee857_SDYnPkPfRd.png?height=214&lazyload=true&maxWidth=700&width=2920)

## API list

The following is a list of APIs for Group announcement and Block.
    

### Group announcement
| `GET` Obtain the basic information of a group announcement /open-apis/docx/v1/chats/:chat_id/announcement | `im:chat.announcement:read` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| --- | --- | --- | --- | --- |
| `GET` Obtain all blocks of a group announcement [/open-apis/docx/v1/chats/:chat_id/announcement/blocks](/ssl:ttdocopen-apis/docx/v1/chats/:chat_id/announcement/blocks) | `im:chat.announcement:read` | `tenant_access_token` `user_access_token` | **✓** | **✓** | ### Block

| `POST` Create blocks in group announcement /open-apis/docx/v1/chats/:chat_id/announcement/blocks/:block_id/children | `im:chat.announcement:write_only` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| --- | --- | --- | --- | --- |
| `PATCH` Batch update blocks in group announcement /open-apis/docx/v1/chats/:chat_id/announcement/blocks/batch_update | `im:chat.announcement:write_only` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `GET` Obtain the block content in group announcementopen-apis/docx/v1/chats/:chat_id/announcement/blocks/:block_id | `im:chat.announcement:read` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `GET` Obtain all the child blocks /open-apis/docx/v1/chats/:chat_id/announcement/blocks/:block_id/children | `im:chat.announcement:read` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `DELETE` Delete blocks in group announcement/open-apis/docx/v1/chats/:chat_id/announcement/blocks/:block_id/children/batch_delete | `im:chat.announcement:write_only` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
