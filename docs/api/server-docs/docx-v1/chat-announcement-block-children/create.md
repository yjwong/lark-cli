---
title: "Create blocks in group announcement"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/document-docx/docx-v1/chat-announcement-block-children/create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/docx/v1/chats/:chat_id/announcement/blocks/:block_id/children"
service: "docx-v1"
resource: "chat-announcement-block-children"
section: "Group Chat"
scopes:
  - "im:chat.announcement:write_only"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1739935025000"
---

# Create blocks in group announcement

Creates a batch of child blocks for a given block and inserts them at a specific location. The return value of this API is the rich text content of the newly created child blocks. To learn about the parent-child relationship between blocks, see Group announcement overview-Basic concept.

> **Application frequency limit:** The maximum frequency of a single application call is 5 times per second, beyond which the HTTP status code 400 and error code 99991400 will be returned.
> 
> **Group announcement frequency limit**: The maximum number of concurrent edits per second for a single group announcement is 3, beyond which the HTTP status code 429 will be returned. The edit operations include:
> 
> - Create blocks
> - Batch update blocks
> - Delete blocks
> 
> When a request is frequency-limited, the application needs to handle the frequency-limited status code and use an exponential fallback algorithm or some other frequency-control strategy to reduce the rate of calls to the API.

## Prerequisites

- The application needs to enable bot capability.

- The user or robot calling the current interface must be in the corresponding group and must have the permission to read the group announcement document.

## Usage restrictions

- If the group is configured with **Only group owner and group admin can edit group info**, only group owners, group administrators, or robots that create groups and have the **Update the information of groups created by app (im:chat:operate_as_owner)** permission can update group announcement information.

- If the group is not configured with **Only group owner and group admin can edit group info**, all group members can update group announcement information.

- When operating an internal group, the operator and the operated group must be in the same tenant.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/docx/v1/chats/:chat_id/announcement/blocks/:block_id/children |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes | `im:chat.announcement:write_only` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `chat_id` | `string` | Group ID. How to get it: - Create a group, get the chat_id of the group from the returned result. - Call the Get the list of groups that the user or robot is in interface to query the chat_id of the group that the user or robot is in. - Call the Search the list of groups visible to the user or robot to search for the chat_id of the group that the user or robot is in and the group that is open to the user or robot. **Note**: Single chat (group type is `p2p`) does not support getting group announcements. **Example value**: "oc_5ad11d72b830411d72b836c20" |
| `block_id` | `string` | The `block_id` of the parent block. To create child blocks for the group announcement, you can fill in its `chat_id` here. You can use the Obtain all blocks of a group announcement to get the `block_id`. To learn about the parent-child relationship between blocks, see Document overview-Basic concept **Example value**: "doxcnO6UW6wAw2qIcYf4hZpFIth" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `revision_id` | `int` | No | The group announcement version to be queried. -1 indicates the latest version of the group announcement. Once the group announcement is created, the `revision_id` is 1. Make sure you have editing permission for the group announcement. You can use the Obtain the basic information of a group announcement api to get the latest revision ID **Example value**: -1 **Default value**: `-1` **Data validation rules**: - Minimum value: `-1` |
| `client_token` | `string` | No | The unique identifier of the operation, corresponding to the `client_token` returned by the API. It is used for idempotent update operations. When this value is null, a new request will be initiated; when this value is not null, idempotent updates will be performed. **Example value**: "fe599b60-450f-46ff-b2ef-9f6675625b97" |
| `user_id_type` | `string` | No | User ID categories **Example value**: "open_id" **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `children` | `block[]` | No | The child block list. To learn about the parent-child relationship between blocks, see Document overview-Basic concept. In a single request, you can create a maximum of 5 Sheet blocks. **Data validation rules**: - Length range: `1` ～ `50` |
|   `block_type` | `int` | Yes | Block type **Example value**: 1 **Optional values are**:  - `1`: Page Block - `2`: Text Block - `3`: Level 1 Heading Block - `4`: Level 2 Heading Block - `5`: Level 3 Heading Block - `6`: Level 4 Heading Block - `7`: Level 5 Heading Block - `8`: Level 6 Heading Block - `9`: Level 7 Heading Block - `10`: Level 8 Heading Block - `11`: Level 9 Heading Block - `12`: Bullet List Block - `13`: Ordered List Block - `14`: Code Block - `15`: Quote Block - `17`: Todo Block - `18`: Bitable Block - `19`: Callout Block - `20`: Chat Card Block - `21`: UML Diagram Block - `22`: Divider Block - `23`: File Block - `24`: Grid Block - `25`: Grid Column Block - `26`: Iframe Block - `27`: Image Block - `28`: ISV Block - `29`: Mindnote Block - `30`: Sheet Block - `31`: Table Block - `32`: Table Cell Block - `33`: View Block - `34`: Quote Container Block - `35`: Task Block - `36`: OKR Block - `37`: OKR Objective Block - `38`: OKR Key Result Block - `39`: OKR Progress Block - `40`: Add-Ons Block - `41`: Jira Issue Block - `42`: Wiki Catalog Block - `43`: Board - `44`: Agenda Block - `45`: Agenda Item Block - `46`: Agenda Item Title Block - `47`: Agenda Item Content Block - `48`: Link Preview Block - `999`: Blocks not supported  |
|   `text` | `text` | No | Text Block |
|     `style` | `text_style` | No | Text style |
|       `align` | `int` | No | Alignment **Example value**: 1 **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  **Default value**: `1` |
|       `done` | `boolean` | No | Todo's status. Supports Todo Blocks **Example value**: true **Default value**: `false` |
|       `folded` | `boolean` | No | Folded state of text. Supports Heading1~9 block and the Text, Bullet, Ordered and Todo blocks with children **Example value**: true **Default value**: `false` |
|       `language` | `int` | No | Code block language. Only supports Code block **Example value**: 1 **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|       `wrap` | `boolean` | No | Whether the code block wraps automatically. Only supports the Code block **Example value**: true **Default value**: `false` |
|       `background_color` | `string` | No | The background color of the block **Example value**: "LightGrayBackground" **Optional values are**:  - `LightGrayBackground`: Light gray - `LightRedBackground`: Light red - `LightOrangeBackground`: Light orange - `LightYellowBackground`: Light yellow - `LightGreenBackground`: Light green - `LightBlueBackground`: Light blue - `LightPurpleBackground`: Light purple - `PaleGrayBackground`: Medium gray - `DarkGrayBackground`: Gray - `DarkRedBackground`: Medium red - `DarkOrangeBackground`: Medium orange - `DarkYellowBackground`: Medium yellow - `DarkGreenBackground`: Medium green - `DarkBlueBackground`: Medium blue - `DarkPurpleBackground`: Medium purple  |
|       `indentation_level` | `string` | No | First line indentation level **Example value**: "NoIndent" **Optional values are**:  - `NoIndent`: No indent - `OneLevelIndent`: One Level indent  |
|     `elements` | `text_element[]` | Yes | Text element |
|       `text_run` | `text_run` | No | Text. Supports Page, Text, Heading1~9, Bullet, Ordered, Code, Quote and Todo blocks |
|         `content` | `string` | Yes | Text content **Example value**: "text" |
|         `text_element_style` | `text_element_style` | No | Text local style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `mention_user` | `mention_user` | No | @User. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|         `user_id` | `string` | Yes | User OpenID **Example value**: "ou_3bbe8a09c20e89cce9bff989ed840674" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `mention_doc` | `mention_doc` | No | @Document. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|         `token` | `string` | Yes | Docs token **Example value**: "doxbc873Y7cXD153gXqb76abcef" |
|         `obj_type` | `int` | Yes | Docs type **Example value**: 22 **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|         `url` | `string` | Yes | Docs link (url_encode) **Example value**: "https%3A%2F%2Fexample.larksuite.com%2Fdocx%2Fdoxbc873Y7cXD153gXqb76G1Y9b" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `reminder` | `reminder` | No | Date Reminder. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|         `create_user_id` | `string` | Yes | Creator user ID **Example value**: "0E2633a3-aa1a-4171-af9e-0768ff863566" |
|         `is_whole_day` | `boolean` | No | Whether it's a date or a whole hour **Example value**: true **Default value**: `false` |
|         `expire_time` | `string` | Yes | Time the event occurred (millisecond timestamp) **Example value**: "1641967200000" |
|         `notify_time` | `string` | Yes | Time to trigger notification (millisecond timestamp) **Example value**: "1643166000000" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `file` | `inline_file` | No | Inline file. Only supports deleting or moving the element, and creating new inline files is not supported |
|         `file_token` | `string` | No | File token **Example value**: "boxcnOj88GDkmWGm2zsTyCBqoLb" |
|         `source_block_id` | `string` | No | The ID of the block in which the file is located in the current document **Example value**: "doxcnM46kSWSkgUMW04ldKabcef" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `inline_block` | `inline_block` | No | Inline block. Only supports deleting or moving the element, and creating new inline blocks is not supported |
|         `block_id` | `string` | Yes | Block ID **Example value**: "doxcnPFi0R56ctbvh2Mjkkabcef" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `equation` | `equation` | No | Equation |
|         `content` | `string` | Yes | Equation content that conforms to KaTeX syntax. For more information, see [KaTex syntax rules](https://katex.org/docs/supported.html]. **Example value**: "E=mc^2\n" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|   `heading1` | `text` | No | Level 1 Heading Block |
|     `style` | `text_style` | No | Text style |
|       `align` | `int` | No | Alignment **Example value**: 1 **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  **Default value**: `1` |
|       `done` | `boolean` | No | Todo's status. Supports Todo Blocks **Example value**: true **Default value**: `false` |
|       `folded` | `boolean` | No | Folded state of text. Supports Heading1~9 block and the Text, Bullet, Ordered and Todo blocks with children **Example value**: true **Default value**: `false` |
|       `language` | `int` | No | Code block language. Only supports Code block **Example value**: 1 **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|       `wrap` | `boolean` | No | Whether the code block wraps automatically. Only supports the Code block **Example value**: true **Default value**: `false` |
|       `background_color` | `string` | No | The background color of the block **Example value**: "LightGrayBackground" **Optional values are**:  - `LightGrayBackground`: Light gray - `LightRedBackground`: Light red - `LightOrangeBackground`: Light orange - `LightYellowBackground`: Light yellow - `LightGreenBackground`: Light green - `LightBlueBackground`: Light blue - `LightPurpleBackground`: Light purple - `PaleGrayBackground`: Medium gray - `DarkGrayBackground`: Gray - `DarkRedBackground`: Medium red - `DarkOrangeBackground`: Medium orange - `DarkYellowBackground`: Medium yellow - `DarkGreenBackground`: Medium green - `DarkBlueBackground`: Medium blue - `DarkPurpleBackground`: Medium purple  |
|       `indentation_level` | `string` | No | First line indentation level **Example value**: "NoIndent" **Optional values are**:  - `NoIndent`: No indent - `OneLevelIndent`: One Level indent  |
|     `elements` | `text_element[]` | Yes | Text element |
|       `text_run` | `text_run` | No | Text. Supports Page, Text, Heading1~9, Bullet, Ordered, Code, Quote and Todo blocks |
|         `content` | `string` | Yes | Text content **Example value**: "text" |
|         `text_element_style` | `text_element_style` | No | Text local style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `mention_user` | `mention_user` | No | @User. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|         `user_id` | `string` | Yes | User OpenID **Example value**: "ou_3bbe8a09c20e89cce9bff989ed840674" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `mention_doc` | `mention_doc` | No | @Document. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|         `token` | `string` | Yes | Docs token **Example value**: "doxbc873Y7cXD153gXqb76abcef" |
|         `obj_type` | `int` | Yes | Docs type **Example value**: 22 **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|         `url` | `string` | Yes | Docs link (url_encode) **Example value**: "https%3A%2F%2Fexample.larksuite.com%2Fdocx%2Fdoxbc873Y7cXD153gXqb76G1Y9b" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `reminder` | `reminder` | No | Date Reminder. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|         `create_user_id` | `string` | Yes | Creator user ID **Example value**: "0E2633a3-aa1a-4171-af9e-0768ff863566" |
|         `is_whole_day` | `boolean` | No | Whether it's a date or a whole hour **Example value**: true **Default value**: `false` |
|         `expire_time` | `string` | Yes | Time the event occurred (millisecond timestamp) **Example value**: "1641967200000" |
|         `notify_time` | `string` | Yes | Time to trigger notification (millisecond timestamp) **Example value**: "1643166000000" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `file` | `inline_file` | No | Inline file. Only supports deleting or moving the element, and creating new inline files is not supported |
|         `file_token` | `string` | No | File token **Example value**: "boxcnOj88GDkmWGm2zsTyCBqoLb" |
|         `source_block_id` | `string` | No | The ID of the block in which the file is located in the current document **Example value**: "doxcnM46kSWSkgUMW04ldKabcef" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `inline_block` | `inline_block` | No | Inline block. Only supports deleting or moving the element, and creating new inline blocks is not supported |
|         `block_id` | `string` | Yes | Block ID **Example value**: "doxcnPFi0R56ctbvh2Mjkkabcef" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `equation` | `equation` | No | Equation |
|         `content` | `string` | Yes | Equation content that conforms to KaTeX syntax. For more information, see [KaTex syntax rules](https://katex.org/docs/supported.html]. **Example value**: "E=mc^2\n" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|   `heading2` | `text` | No | Level 2 Heading Block |
|     `style` | `text_style` | No | Text style |
|       `align` | `int` | No | Alignment **Example value**: 1 **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  **Default value**: `1` |
|       `done` | `boolean` | No | Todo's status. Supports Todo Blocks **Example value**: true **Default value**: `false` |
|       `folded` | `boolean` | No | Folded state of text. Supports Heading1~9 block and the Text, Bullet, Ordered and Todo blocks with children **Example value**: true **Default value**: `false` |
|       `language` | `int` | No | Code block language. Only supports Code block **Example value**: 1 **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|       `wrap` | `boolean` | No | Whether the code block wraps automatically. Only supports the Code block **Example value**: true **Default value**: `false` |
|       `background_color` | `string` | No | The background color of the block **Example value**: "LightGrayBackground" **Optional values are**:  - `LightGrayBackground`: Light gray - `LightRedBackground`: Light red - `LightOrangeBackground`: Light orange - `LightYellowBackground`: Light yellow - `LightGreenBackground`: Light green - `LightBlueBackground`: Light blue - `LightPurpleBackground`: Light purple - `PaleGrayBackground`: Medium gray - `DarkGrayBackground`: Gray - `DarkRedBackground`: Medium red - `DarkOrangeBackground`: Medium orange - `DarkYellowBackground`: Medium yellow - `DarkGreenBackground`: Medium green - `DarkBlueBackground`: Medium blue - `DarkPurpleBackground`: Medium purple  |
|       `indentation_level` | `string` | No | First line indentation level **Example value**: "NoIndent" **Optional values are**:  - `NoIndent`: No indent - `OneLevelIndent`: One Level indent  |
|     `elements` | `text_element[]` | Yes | Text element |
|       `text_run` | `text_run` | No | Text. Supports Page, Text, Heading1~9, Bullet, Ordered, Code, Quote and Todo blocks |
|         `content` | `string` | Yes | Text content **Example value**: "text" |
|         `text_element_style` | `text_element_style` | No | Text local style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `mention_user` | `mention_user` | No | @User. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|         `user_id` | `string` | Yes | User OpenID **Example value**: "ou_3bbe8a09c20e89cce9bff989ed840674" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `mention_doc` | `mention_doc` | No | @Document. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|         `token` | `string` | Yes | Docs token **Example value**: "doxbc873Y7cXD153gXqb76abcef" |
|         `obj_type` | `int` | Yes | Docs type **Example value**: 22 **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|         `url` | `string` | Yes | Docs link (url_encode) **Example value**: "https%3A%2F%2Fexample.larksuite.com%2Fdocx%2Fdoxbc873Y7cXD153gXqb76G1Y9b" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `reminder` | `reminder` | No | Date Reminder. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|         `create_user_id` | `string` | Yes | Creator user ID **Example value**: "0E2633a3-aa1a-4171-af9e-0768ff863566" |
|         `is_whole_day` | `boolean` | No | Whether it's a date or a whole hour **Example value**: true **Default value**: `false` |
|         `expire_time` | `string` | Yes | Time the event occurred (millisecond timestamp) **Example value**: "1641967200000" |
|         `notify_time` | `string` | Yes | Time to trigger notification (millisecond timestamp) **Example value**: "1643166000000" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `file` | `inline_file` | No | Inline file. Only supports deleting or moving the element, and creating new inline files is not supported |
|         `file_token` | `string` | No | File token **Example value**: "boxcnOj88GDkmWGm2zsTyCBqoLb" |
|         `source_block_id` | `string` | No | The ID of the block in which the file is located in the current document **Example value**: "doxcnM46kSWSkgUMW04ldKabcef" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `inline_block` | `inline_block` | No | Inline block. Only supports deleting or moving the element, and creating new inline blocks is not supported |
|         `block_id` | `string` | Yes | Block ID **Example value**: "doxcnPFi0R56ctbvh2Mjkkabcef" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `equation` | `equation` | No | Equation |
|         `content` | `string` | Yes | Equation content that conforms to KaTeX syntax. For more information, see [KaTex syntax rules](https://katex.org/docs/supported.html]. **Example value**: "E=mc^2\n" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|   `heading3` | `text` | No | Level 3 Heading Block |
|     `style` | `text_style` | No | Text style |
|       `align` | `int` | No | Alignment **Example value**: 1 **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  **Default value**: `1` |
|       `done` | `boolean` | No | Todo's status. Supports Todo Blocks **Example value**: true **Default value**: `false` |
|       `folded` | `boolean` | No | Folded state of text. Supports Heading1~9 block and the Text, Bullet, Ordered and Todo blocks with children **Example value**: true **Default value**: `false` |
|       `language` | `int` | No | Code block language. Only supports Code block **Example value**: 1 **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|       `wrap` | `boolean` | No | Whether the code block wraps automatically. Only supports the Code block **Example value**: true **Default value**: `false` |
|       `background_color` | `string` | No | The background color of the block **Example value**: "LightGrayBackground" **Optional values are**:  - `LightGrayBackground`: Light gray - `LightRedBackground`: Light red - `LightOrangeBackground`: Light orange - `LightYellowBackground`: Light yellow - `LightGreenBackground`: Light green - `LightBlueBackground`: Light blue - `LightPurpleBackground`: Light purple - `PaleGrayBackground`: Medium gray - `DarkGrayBackground`: Gray - `DarkRedBackground`: Medium red - `DarkOrangeBackground`: Medium orange - `DarkYellowBackground`: Medium yellow - `DarkGreenBackground`: Medium green - `DarkBlueBackground`: Medium blue - `DarkPurpleBackground`: Medium purple  |
|       `indentation_level` | `string` | No | First line indentation level **Example value**: "NoIndent" **Optional values are**:  - `NoIndent`: No indent - `OneLevelIndent`: One Level indent  |
|     `elements` | `text_element[]` | Yes | Text element |
|       `text_run` | `text_run` | No | Text. Supports Page, Text, Heading1~9, Bullet, Ordered, Code, Quote and Todo blocks |
|         `content` | `string` | Yes | Text content **Example value**: "text" |
|         `text_element_style` | `text_element_style` | No | Text local style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `mention_user` | `mention_user` | No | @User. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|         `user_id` | `string` | Yes | User OpenID **Example value**: "ou_3bbe8a09c20e89cce9bff989ed840674" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `mention_doc` | `mention_doc` | No | @Document. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|         `token` | `string` | Yes | Docs token **Example value**: "doxbc873Y7cXD153gXqb76abcef" |
|         `obj_type` | `int` | Yes | Docs type **Example value**: 22 **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|         `url` | `string` | Yes | Docs link (url_encode) **Example value**: "https%3A%2F%2Fexample.larksuite.com%2Fdocx%2Fdoxbc873Y7cXD153gXqb76G1Y9b" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `reminder` | `reminder` | No | Date Reminder. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|         `create_user_id` | `string` | Yes | Creator user ID **Example value**: "0E2633a3-aa1a-4171-af9e-0768ff863566" |
|         `is_whole_day` | `boolean` | No | Whether it's a date or a whole hour **Example value**: true **Default value**: `false` |
|         `expire_time` | `string` | Yes | Time the event occurred (millisecond timestamp) **Example value**: "1641967200000" |
|         `notify_time` | `string` | Yes | Time to trigger notification (millisecond timestamp) **Example value**: "1643166000000" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `file` | `inline_file` | No | Inline file. Only supports deleting or moving the element, and creating new inline files is not supported |
|         `file_token` | `string` | No | File token **Example value**: "boxcnOj88GDkmWGm2zsTyCBqoLb" |
|         `source_block_id` | `string` | No | The ID of the block in which the file is located in the current document **Example value**: "doxcnM46kSWSkgUMW04ldKabcef" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `inline_block` | `inline_block` | No | Inline block. Only supports deleting or moving the element, and creating new inline blocks is not supported |
|         `block_id` | `string` | Yes | Block ID **Example value**: "doxcnPFi0R56ctbvh2Mjkkabcef" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `equation` | `equation` | No | Equation |
|         `content` | `string` | Yes | Equation content that conforms to KaTeX syntax. For more information, see [KaTex syntax rules](https://katex.org/docs/supported.html]. **Example value**: "E=mc^2\n" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|   `heading4` | `text` | No | Level 4 Heading Block |
|     `style` | `text_style` | No | Text style |
|       `align` | `int` | No | Alignment **Example value**: 1 **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  **Default value**: `1` |
|       `done` | `boolean` | No | Todo's status. Supports Todo Blocks **Example value**: true **Default value**: `false` |
|       `folded` | `boolean` | No | Folded state of text. Supports Heading1~9 block and the Text, Bullet, Ordered and Todo blocks with children **Example value**: true **Default value**: `false` |
|       `language` | `int` | No | Code block language. Only supports Code block **Example value**: 1 **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|       `wrap` | `boolean` | No | Whether the code block wraps automatically. Only supports the Code block **Example value**: true **Default value**: `false` |
|       `background_color` | `string` | No | The background color of the block **Example value**: "LightGrayBackground" **Optional values are**:  - `LightGrayBackground`: Light gray - `LightRedBackground`: Light red - `LightOrangeBackground`: Light orange - `LightYellowBackground`: Light yellow - `LightGreenBackground`: Light green - `LightBlueBackground`: Light blue - `LightPurpleBackground`: Light purple - `PaleGrayBackground`: Medium gray - `DarkGrayBackground`: Gray - `DarkRedBackground`: Medium red - `DarkOrangeBackground`: Medium orange - `DarkYellowBackground`: Medium yellow - `DarkGreenBackground`: Medium green - `DarkBlueBackground`: Medium blue - `DarkPurpleBackground`: Medium purple  |
|       `indentation_level` | `string` | No | First line indentation level **Example value**: "NoIndent" **Optional values are**:  - `NoIndent`: No indent - `OneLevelIndent`: One Level indent  |
|     `elements` | `text_element[]` | Yes | Text element |
|       `text_run` | `text_run` | No | Text. Supports Page, Text, Heading1~9, Bullet, Ordered, Code, Quote and Todo blocks |
|         `content` | `string` | Yes | Text content **Example value**: "text" |
|         `text_element_style` | `text_element_style` | No | Text local style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `mention_user` | `mention_user` | No | @User. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|         `user_id` | `string` | Yes | User OpenID **Example value**: "ou_3bbe8a09c20e89cce9bff989ed840674" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `mention_doc` | `mention_doc` | No | @Document. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|         `token` | `string` | Yes | Docs token **Example value**: "doxbc873Y7cXD153gXqb76abcef" |
|         `obj_type` | `int` | Yes | Docs type **Example value**: 22 **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|         `url` | `string` | Yes | Docs link (url_encode) **Example value**: "https%3A%2F%2Fexample.larksuite.com%2Fdocx%2Fdoxbc873Y7cXD153gXqb76G1Y9b" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `reminder` | `reminder` | No | Date Reminder. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|         `create_user_id` | `string` | Yes | Creator user ID **Example value**: "0E2633a3-aa1a-4171-af9e-0768ff863566" |
|         `is_whole_day` | `boolean` | No | Whether it's a date or a whole hour **Example value**: true **Default value**: `false` |
|         `expire_time` | `string` | Yes | Time the event occurred (millisecond timestamp) **Example value**: "1641967200000" |
|         `notify_time` | `string` | Yes | Time to trigger notification (millisecond timestamp) **Example value**: "1643166000000" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `file` | `inline_file` | No | Inline file. Only supports deleting or moving the element, and creating new inline files is not supported |
|         `file_token` | `string` | No | File token **Example value**: "boxcnOj88GDkmWGm2zsTyCBqoLb" |
|         `source_block_id` | `string` | No | The ID of the block in which the file is located in the current document **Example value**: "doxcnM46kSWSkgUMW04ldKabcef" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `inline_block` | `inline_block` | No | Inline block. Only supports deleting or moving the element, and creating new inline blocks is not supported |
|         `block_id` | `string` | Yes | Block ID **Example value**: "doxcnPFi0R56ctbvh2Mjkkabcef" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `equation` | `equation` | No | Equation |
|         `content` | `string` | Yes | Equation content that conforms to KaTeX syntax. For more information, see [KaTex syntax rules](https://katex.org/docs/supported.html]. **Example value**: "E=mc^2\n" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|   `heading5` | `text` | No | Level 5 Heading Block |
|     `style` | `text_style` | No | Text style |
|       `align` | `int` | No | Alignment **Example value**: 1 **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  **Default value**: `1` |
|       `done` | `boolean` | No | Todo's status. Supports Todo Blocks **Example value**: true **Default value**: `false` |
|       `folded` | `boolean` | No | Folded state of text. Supports Heading1~9 block and the Text, Bullet, Ordered and Todo blocks with children **Example value**: true **Default value**: `false` |
|       `language` | `int` | No | Code block language. Only supports Code block **Example value**: 1 **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|       `wrap` | `boolean` | No | Whether the code block wraps automatically. Only supports the Code block **Example value**: true **Default value**: `false` |
|       `background_color` | `string` | No | The background color of the block **Example value**: "LightGrayBackground" **Optional values are**:  - `LightGrayBackground`: Light gray - `LightRedBackground`: Light red - `LightOrangeBackground`: Light orange - `LightYellowBackground`: Light yellow - `LightGreenBackground`: Light green - `LightBlueBackground`: Light blue - `LightPurpleBackground`: Light purple - `PaleGrayBackground`: Medium gray - `DarkGrayBackground`: Gray - `DarkRedBackground`: Medium red - `DarkOrangeBackground`: Medium orange - `DarkYellowBackground`: Medium yellow - `DarkGreenBackground`: Medium green - `DarkBlueBackground`: Medium blue - `DarkPurpleBackground`: Medium purple  |
|       `indentation_level` | `string` | No | First line indentation level **Example value**: "NoIndent" **Optional values are**:  - `NoIndent`: No indent - `OneLevelIndent`: One Level indent  |
|     `elements` | `text_element[]` | Yes | Text element |
|       `text_run` | `text_run` | No | Text. Supports Page, Text, Heading1~9, Bullet, Ordered, Code, Quote and Todo blocks |
|         `content` | `string` | Yes | Text content **Example value**: "text" |
|         `text_element_style` | `text_element_style` | No | Text local style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `mention_user` | `mention_user` | No | @User. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|         `user_id` | `string` | Yes | User OpenID **Example value**: "ou_3bbe8a09c20e89cce9bff989ed840674" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `mention_doc` | `mention_doc` | No | @Document. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|         `token` | `string` | Yes | Docs token **Example value**: "doxbc873Y7cXD153gXqb76abcef" |
|         `obj_type` | `int` | Yes | Docs type **Example value**: 22 **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|         `url` | `string` | Yes | Docs link (url_encode) **Example value**: "https%3A%2F%2Fexample.larksuite.com%2Fdocx%2Fdoxbc873Y7cXD153gXqb76G1Y9b" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `reminder` | `reminder` | No | Date Reminder. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|         `create_user_id` | `string` | Yes | Creator user ID **Example value**: "0E2633a3-aa1a-4171-af9e-0768ff863566" |
|         `is_whole_day` | `boolean` | No | Whether it's a date or a whole hour **Example value**: true **Default value**: `false` |
|         `expire_time` | `string` | Yes | Time the event occurred (millisecond timestamp) **Example value**: "1641967200000" |
|         `notify_time` | `string` | Yes | Time to trigger notification (millisecond timestamp) **Example value**: "1643166000000" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `file` | `inline_file` | No | Inline file. Only supports deleting or moving the element, and creating new inline files is not supported |
|         `file_token` | `string` | No | File token **Example value**: "boxcnOj88GDkmWGm2zsTyCBqoLb" |
|         `source_block_id` | `string` | No | The ID of the block in which the file is located in the current document **Example value**: "doxcnM46kSWSkgUMW04ldKabcef" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `inline_block` | `inline_block` | No | Inline block. Only supports deleting or moving the element, and creating new inline blocks is not supported |
|         `block_id` | `string` | Yes | Block ID **Example value**: "doxcnPFi0R56ctbvh2Mjkkabcef" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `equation` | `equation` | No | Equation |
|         `content` | `string` | Yes | Equation content that conforms to KaTeX syntax. For more information, see [KaTex syntax rules](https://katex.org/docs/supported.html]. **Example value**: "E=mc^2\n" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|   `heading6` | `text` | No | Level 6 Heading Block |
|     `style` | `text_style` | No | Text style |
|       `align` | `int` | No | Alignment **Example value**: 1 **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  **Default value**: `1` |
|       `done` | `boolean` | No | Todo's status. Supports Todo Blocks **Example value**: true **Default value**: `false` |
|       `folded` | `boolean` | No | Folded state of text. Supports Heading1~9 block and the Text, Bullet, Ordered and Todo blocks with children **Example value**: true **Default value**: `false` |
|       `language` | `int` | No | Code block language. Only supports Code block **Example value**: 1 **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|       `wrap` | `boolean` | No | Whether the code block wraps automatically. Only supports the Code block **Example value**: true **Default value**: `false` |
|       `background_color` | `string` | No | The background color of the block **Example value**: "LightGrayBackground" **Optional values are**:  - `LightGrayBackground`: Light gray - `LightRedBackground`: Light red - `LightOrangeBackground`: Light orange - `LightYellowBackground`: Light yellow - `LightGreenBackground`: Light green - `LightBlueBackground`: Light blue - `LightPurpleBackground`: Light purple - `PaleGrayBackground`: Medium gray - `DarkGrayBackground`: Gray - `DarkRedBackground`: Medium red - `DarkOrangeBackground`: Medium orange - `DarkYellowBackground`: Medium yellow - `DarkGreenBackground`: Medium green - `DarkBlueBackground`: Medium blue - `DarkPurpleBackground`: Medium purple  |
|       `indentation_level` | `string` | No | First line indentation level **Example value**: "NoIndent" **Optional values are**:  - `NoIndent`: No indent - `OneLevelIndent`: One Level indent  |
|     `elements` | `text_element[]` | Yes | Text element |
|       `text_run` | `text_run` | No | Text. Supports Page, Text, Heading1~9, Bullet, Ordered, Code, Quote and Todo blocks |
|         `content` | `string` | Yes | Text content **Example value**: "text" |
|         `text_element_style` | `text_element_style` | No | Text local style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `mention_user` | `mention_user` | No | @User. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|         `user_id` | `string` | Yes | User OpenID **Example value**: "ou_3bbe8a09c20e89cce9bff989ed840674" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `mention_doc` | `mention_doc` | No | @Document. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|         `token` | `string` | Yes | Docs token **Example value**: "doxbc873Y7cXD153gXqb76abcef" |
|         `obj_type` | `int` | Yes | Docs type **Example value**: 22 **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|         `url` | `string` | Yes | Docs link (url_encode) **Example value**: "https%3A%2F%2Fexample.larksuite.com%2Fdocx%2Fdoxbc873Y7cXD153gXqb76G1Y9b" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `reminder` | `reminder` | No | Date Reminder. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|         `create_user_id` | `string` | Yes | Creator user ID **Example value**: "0E2633a3-aa1a-4171-af9e-0768ff863566" |
|         `is_whole_day` | `boolean` | No | Whether it's a date or a whole hour **Example value**: true **Default value**: `false` |
|         `expire_time` | `string` | Yes | Time the event occurred (millisecond timestamp) **Example value**: "1641967200000" |
|         `notify_time` | `string` | Yes | Time to trigger notification (millisecond timestamp) **Example value**: "1643166000000" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `file` | `inline_file` | No | Inline file. Only supports deleting or moving the element, and creating new inline files is not supported |
|         `file_token` | `string` | No | File token **Example value**: "boxcnOj88GDkmWGm2zsTyCBqoLb" |
|         `source_block_id` | `string` | No | The ID of the block in which the file is located in the current document **Example value**: "doxcnM46kSWSkgUMW04ldKabcef" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `inline_block` | `inline_block` | No | Inline block. Only supports deleting or moving the element, and creating new inline blocks is not supported |
|         `block_id` | `string` | Yes | Block ID **Example value**: "doxcnPFi0R56ctbvh2Mjkkabcef" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `equation` | `equation` | No | Equation |
|         `content` | `string` | Yes | Equation content that conforms to KaTeX syntax. For more information, see [KaTex syntax rules](https://katex.org/docs/supported.html]. **Example value**: "E=mc^2\n" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|   `heading7` | `text` | No | Level 7 Heading Block |
|     `style` | `text_style` | No | Text style |
|       `align` | `int` | No | Alignment **Example value**: 1 **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  **Default value**: `1` |
|       `done` | `boolean` | No | Todo's status. Supports Todo Blocks **Example value**: true **Default value**: `false` |
|       `folded` | `boolean` | No | Folded state of text. Supports Heading1~9 block and the Text, Bullet, Ordered and Todo blocks with children **Example value**: true **Default value**: `false` |
|       `language` | `int` | No | Code block language. Only supports Code block **Example value**: 1 **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|       `wrap` | `boolean` | No | Whether the code block wraps automatically. Only supports the Code block **Example value**: true **Default value**: `false` |
|       `background_color` | `string` | No | The background color of the block **Example value**: "LightGrayBackground" **Optional values are**:  - `LightGrayBackground`: Light gray - `LightRedBackground`: Light red - `LightOrangeBackground`: Light orange - `LightYellowBackground`: Light yellow - `LightGreenBackground`: Light green - `LightBlueBackground`: Light blue - `LightPurpleBackground`: Light purple - `PaleGrayBackground`: Medium gray - `DarkGrayBackground`: Gray - `DarkRedBackground`: Medium red - `DarkOrangeBackground`: Medium orange - `DarkYellowBackground`: Medium yellow - `DarkGreenBackground`: Medium green - `DarkBlueBackground`: Medium blue - `DarkPurpleBackground`: Medium purple  |
|       `indentation_level` | `string` | No | First line indentation level **Example value**: "NoIndent" **Optional values are**:  - `NoIndent`: No indent - `OneLevelIndent`: One Level indent  |
|     `elements` | `text_element[]` | Yes | Text element |
|       `text_run` | `text_run` | No | Text. Supports Page, Text, Heading1~9, Bullet, Ordered, Code, Quote and Todo blocks |
|         `content` | `string` | Yes | Text content **Example value**: "text" |
|         `text_element_style` | `text_element_style` | No | Text local style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `mention_user` | `mention_user` | No | @User. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|         `user_id` | `string` | Yes | User OpenID **Example value**: "ou_3bbe8a09c20e89cce9bff989ed840674" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `mention_doc` | `mention_doc` | No | @Document. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|         `token` | `string` | Yes | Docs token **Example value**: "doxbc873Y7cXD153gXqb76abcef" |
|         `obj_type` | `int` | Yes | Docs type **Example value**: 22 **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|         `url` | `string` | Yes | Docs link (url_encode) **Example value**: "https%3A%2F%2Fexample.larksuite.com%2Fdocx%2Fdoxbc873Y7cXD153gXqb76G1Y9b" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `reminder` | `reminder` | No | Date Reminder. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|         `create_user_id` | `string` | Yes | Creator user ID **Example value**: "0E2633a3-aa1a-4171-af9e-0768ff863566" |
|         `is_whole_day` | `boolean` | No | Whether it's a date or a whole hour **Example value**: true **Default value**: `false` |
|         `expire_time` | `string` | Yes | Time the event occurred (millisecond timestamp) **Example value**: "1641967200000" |
|         `notify_time` | `string` | Yes | Time to trigger notification (millisecond timestamp) **Example value**: "1643166000000" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `file` | `inline_file` | No | Inline file. Only supports deleting or moving the element, and creating new inline files is not supported |
|         `file_token` | `string` | No | File token **Example value**: "boxcnOj88GDkmWGm2zsTyCBqoLb" |
|         `source_block_id` | `string` | No | The ID of the block in which the file is located in the current document **Example value**: "doxcnM46kSWSkgUMW04ldKabcef" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `inline_block` | `inline_block` | No | Inline block. Only supports deleting or moving the element, and creating new inline blocks is not supported |
|         `block_id` | `string` | Yes | Block ID **Example value**: "doxcnPFi0R56ctbvh2Mjkkabcef" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `equation` | `equation` | No | Equation |
|         `content` | `string` | Yes | Equation content that conforms to KaTeX syntax. For more information, see [KaTex syntax rules](https://katex.org/docs/supported.html]. **Example value**: "E=mc^2\n" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|   `heading8` | `text` | No | Level 8 Heading Block |
|     `style` | `text_style` | No | Text style |
|       `align` | `int` | No | Alignment **Example value**: 1 **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  **Default value**: `1` |
|       `done` | `boolean` | No | Todo's status. Supports Todo Blocks **Example value**: true **Default value**: `false` |
|       `folded` | `boolean` | No | Folded state of text. Supports Heading1~9 block and the Text, Bullet, Ordered and Todo blocks with children **Example value**: true **Default value**: `false` |
|       `language` | `int` | No | Code block language. Only supports Code block **Example value**: 1 **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|       `wrap` | `boolean` | No | Whether the code block wraps automatically. Only supports the Code block **Example value**: true **Default value**: `false` |
|       `background_color` | `string` | No | The background color of the block **Example value**: "LightGrayBackground" **Optional values are**:  - `LightGrayBackground`: Light gray - `LightRedBackground`: Light red - `LightOrangeBackground`: Light orange - `LightYellowBackground`: Light yellow - `LightGreenBackground`: Light green - `LightBlueBackground`: Light blue - `LightPurpleBackground`: Light purple - `PaleGrayBackground`: Medium gray - `DarkGrayBackground`: Gray - `DarkRedBackground`: Medium red - `DarkOrangeBackground`: Medium orange - `DarkYellowBackground`: Medium yellow - `DarkGreenBackground`: Medium green - `DarkBlueBackground`: Medium blue - `DarkPurpleBackground`: Medium purple  |
|       `indentation_level` | `string` | No | First line indentation level **Example value**: "NoIndent" **Optional values are**:  - `NoIndent`: No indent - `OneLevelIndent`: One Level indent  |
|     `elements` | `text_element[]` | Yes | Text element |
|       `text_run` | `text_run` | No | Text. Supports Page, Text, Heading1~9, Bullet, Ordered, Code, Quote and Todo blocks |
|         `content` | `string` | Yes | Text content **Example value**: "text" |
|         `text_element_style` | `text_element_style` | No | Text local style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `mention_user` | `mention_user` | No | @User. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|         `user_id` | `string` | Yes | User OpenID **Example value**: "ou_3bbe8a09c20e89cce9bff989ed840674" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `mention_doc` | `mention_doc` | No | @Document. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|         `token` | `string` | Yes | Docs token **Example value**: "doxbc873Y7cXD153gXqb76abcef" |
|         `obj_type` | `int` | Yes | Docs type **Example value**: 22 **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|         `url` | `string` | Yes | Docs link (url_encode) **Example value**: "https%3A%2F%2Fexample.larksuite.com%2Fdocx%2Fdoxbc873Y7cXD153gXqb76G1Y9b" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `reminder` | `reminder` | No | Date Reminder. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|         `create_user_id` | `string` | Yes | Creator user ID **Example value**: "0E2633a3-aa1a-4171-af9e-0768ff863566" |
|         `is_whole_day` | `boolean` | No | Whether it's a date or a whole hour **Example value**: true **Default value**: `false` |
|         `expire_time` | `string` | Yes | Time the event occurred (millisecond timestamp) **Example value**: "1641967200000" |
|         `notify_time` | `string` | Yes | Time to trigger notification (millisecond timestamp) **Example value**: "1643166000000" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `file` | `inline_file` | No | Inline file. Only supports deleting or moving the element, and creating new inline files is not supported |
|         `file_token` | `string` | No | File token **Example value**: "boxcnOj88GDkmWGm2zsTyCBqoLb" |
|         `source_block_id` | `string` | No | The ID of the block in which the file is located in the current document **Example value**: "doxcnM46kSWSkgUMW04ldKabcef" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `inline_block` | `inline_block` | No | Inline block. Only supports deleting or moving the element, and creating new inline blocks is not supported |
|         `block_id` | `string` | Yes | Block ID **Example value**: "doxcnPFi0R56ctbvh2Mjkkabcef" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `equation` | `equation` | No | Equation |
|         `content` | `string` | Yes | Equation content that conforms to KaTeX syntax. For more information, see [KaTex syntax rules](https://katex.org/docs/supported.html]. **Example value**: "E=mc^2\n" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|   `heading9` | `text` | No | Level 9 Heading Block |
|     `style` | `text_style` | No | Text style |
|       `align` | `int` | No | Alignment **Example value**: 1 **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  **Default value**: `1` |
|       `done` | `boolean` | No | Todo's status. Supports Todo Blocks **Example value**: true **Default value**: `false` |
|       `folded` | `boolean` | No | Folded state of text. Supports Heading1~9 block and the Text, Bullet, Ordered and Todo blocks with children **Example value**: true **Default value**: `false` |
|       `language` | `int` | No | Code block language. Only supports Code block **Example value**: 1 **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|       `wrap` | `boolean` | No | Whether the code block wraps automatically. Only supports the Code block **Example value**: true **Default value**: `false` |
|       `background_color` | `string` | No | The background color of the block **Example value**: "LightGrayBackground" **Optional values are**:  - `LightGrayBackground`: Light gray - `LightRedBackground`: Light red - `LightOrangeBackground`: Light orange - `LightYellowBackground`: Light yellow - `LightGreenBackground`: Light green - `LightBlueBackground`: Light blue - `LightPurpleBackground`: Light purple - `PaleGrayBackground`: Medium gray - `DarkGrayBackground`: Gray - `DarkRedBackground`: Medium red - `DarkOrangeBackground`: Medium orange - `DarkYellowBackground`: Medium yellow - `DarkGreenBackground`: Medium green - `DarkBlueBackground`: Medium blue - `DarkPurpleBackground`: Medium purple  |
|       `indentation_level` | `string` | No | First line indentation level **Example value**: "NoIndent" **Optional values are**:  - `NoIndent`: No indent - `OneLevelIndent`: One Level indent  |
|     `elements` | `text_element[]` | Yes | Text element |
|       `text_run` | `text_run` | No | Text. Supports Page, Text, Heading1~9, Bullet, Ordered, Code, Quote and Todo blocks |
|         `content` | `string` | Yes | Text content **Example value**: "text" |
|         `text_element_style` | `text_element_style` | No | Text local style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `mention_user` | `mention_user` | No | @User. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|         `user_id` | `string` | Yes | User OpenID **Example value**: "ou_3bbe8a09c20e89cce9bff989ed840674" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `mention_doc` | `mention_doc` | No | @Document. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|         `token` | `string` | Yes | Docs token **Example value**: "doxbc873Y7cXD153gXqb76abcef" |
|         `obj_type` | `int` | Yes | Docs type **Example value**: 22 **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|         `url` | `string` | Yes | Docs link (url_encode) **Example value**: "https%3A%2F%2Fexample.larksuite.com%2Fdocx%2Fdoxbc873Y7cXD153gXqb76G1Y9b" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `reminder` | `reminder` | No | Date Reminder. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|         `create_user_id` | `string` | Yes | Creator user ID **Example value**: "0E2633a3-aa1a-4171-af9e-0768ff863566" |
|         `is_whole_day` | `boolean` | No | Whether it's a date or a whole hour **Example value**: true **Default value**: `false` |
|         `expire_time` | `string` | Yes | Time the event occurred (millisecond timestamp) **Example value**: "1641967200000" |
|         `notify_time` | `string` | Yes | Time to trigger notification (millisecond timestamp) **Example value**: "1643166000000" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `file` | `inline_file` | No | Inline file. Only supports deleting or moving the element, and creating new inline files is not supported |
|         `file_token` | `string` | No | File token **Example value**: "boxcnOj88GDkmWGm2zsTyCBqoLb" |
|         `source_block_id` | `string` | No | The ID of the block in which the file is located in the current document **Example value**: "doxcnM46kSWSkgUMW04ldKabcef" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `inline_block` | `inline_block` | No | Inline block. Only supports deleting or moving the element, and creating new inline blocks is not supported |
|         `block_id` | `string` | Yes | Block ID **Example value**: "doxcnPFi0R56ctbvh2Mjkkabcef" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `equation` | `equation` | No | Equation |
|         `content` | `string` | Yes | Equation content that conforms to KaTeX syntax. For more information, see [KaTex syntax rules](https://katex.org/docs/supported.html]. **Example value**: "E=mc^2\n" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|   `bullet` | `text` | No | Bullet List Block |
|     `style` | `text_style` | No | Text style |
|       `align` | `int` | No | Alignment **Example value**: 1 **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  **Default value**: `1` |
|       `done` | `boolean` | No | Todo's status. Supports Todo Blocks **Example value**: true **Default value**: `false` |
|       `folded` | `boolean` | No | Folded state of text. Supports Heading1~9 block and the Text, Bullet, Ordered and Todo blocks with children **Example value**: true **Default value**: `false` |
|       `language` | `int` | No | Code block language. Only supports Code block **Example value**: 1 **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|       `wrap` | `boolean` | No | Whether the code block wraps automatically. Only supports the Code block **Example value**: true **Default value**: `false` |
|       `background_color` | `string` | No | The background color of the block **Example value**: "LightGrayBackground" **Optional values are**:  - `LightGrayBackground`: Light gray - `LightRedBackground`: Light red - `LightOrangeBackground`: Light orange - `LightYellowBackground`: Light yellow - `LightGreenBackground`: Light green - `LightBlueBackground`: Light blue - `LightPurpleBackground`: Light purple - `PaleGrayBackground`: Medium gray - `DarkGrayBackground`: Gray - `DarkRedBackground`: Medium red - `DarkOrangeBackground`: Medium orange - `DarkYellowBackground`: Medium yellow - `DarkGreenBackground`: Medium green - `DarkBlueBackground`: Medium blue - `DarkPurpleBackground`: Medium purple  |
|       `indentation_level` | `string` | No | First line indentation level **Example value**: "NoIndent" **Optional values are**:  - `NoIndent`: No indent - `OneLevelIndent`: One Level indent  |
|     `elements` | `text_element[]` | Yes | Text element |
|       `text_run` | `text_run` | No | Text. Supports Page, Text, Heading1~9, Bullet, Ordered, Code, Quote and Todo blocks |
|         `content` | `string` | Yes | Text content **Example value**: "text" |
|         `text_element_style` | `text_element_style` | No | Text local style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `mention_user` | `mention_user` | No | @User. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|         `user_id` | `string` | Yes | User OpenID **Example value**: "ou_3bbe8a09c20e89cce9bff989ed840674" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `mention_doc` | `mention_doc` | No | @Document. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|         `token` | `string` | Yes | Docs token **Example value**: "doxbc873Y7cXD153gXqb76abcef" |
|         `obj_type` | `int` | Yes | Docs type **Example value**: 22 **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|         `url` | `string` | Yes | Docs link (url_encode) **Example value**: "https%3A%2F%2Fexample.larksuite.com%2Fdocx%2Fdoxbc873Y7cXD153gXqb76G1Y9b" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `reminder` | `reminder` | No | Date Reminder. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|         `create_user_id` | `string` | Yes | Creator user ID **Example value**: "0E2633a3-aa1a-4171-af9e-0768ff863566" |
|         `is_whole_day` | `boolean` | No | Whether it's a date or a whole hour **Example value**: true **Default value**: `false` |
|         `expire_time` | `string` | Yes | Time the event occurred (millisecond timestamp) **Example value**: "1641967200000" |
|         `notify_time` | `string` | Yes | Time to trigger notification (millisecond timestamp) **Example value**: "1643166000000" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `file` | `inline_file` | No | Inline file. Only supports deleting or moving the element, and creating new inline files is not supported |
|         `file_token` | `string` | No | File token **Example value**: "boxcnOj88GDkmWGm2zsTyCBqoLb" |
|         `source_block_id` | `string` | No | The ID of the block in which the file is located in the current document **Example value**: "doxcnM46kSWSkgUMW04ldKabcef" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `inline_block` | `inline_block` | No | Inline block. Only supports deleting or moving the element, and creating new inline blocks is not supported |
|         `block_id` | `string` | Yes | Block ID **Example value**: "doxcnPFi0R56ctbvh2Mjkkabcef" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `equation` | `equation` | No | Equation |
|         `content` | `string` | Yes | Equation content that conforms to KaTeX syntax. For more information, see [KaTex syntax rules](https://katex.org/docs/supported.html]. **Example value**: "E=mc^2\n" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|   `ordered` | `text` | No | Ordered List Block |
|     `style` | `text_style` | No | Text style |
|       `align` | `int` | No | Alignment **Example value**: 1 **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  **Default value**: `1` |
|       `done` | `boolean` | No | Todo's status. Supports Todo Blocks **Example value**: true **Default value**: `false` |
|       `folded` | `boolean` | No | Folded state of text. Supports Heading1~9 block and the Text, Bullet, Ordered and Todo blocks with children **Example value**: true **Default value**: `false` |
|       `language` | `int` | No | Code block language. Only supports Code block **Example value**: 1 **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|       `wrap` | `boolean` | No | Whether the code block wraps automatically. Only supports the Code block **Example value**: true **Default value**: `false` |
|       `background_color` | `string` | No | The background color of the block **Example value**: "LightGrayBackground" **Optional values are**:  - `LightGrayBackground`: Light gray - `LightRedBackground`: Light red - `LightOrangeBackground`: Light orange - `LightYellowBackground`: Light yellow - `LightGreenBackground`: Light green - `LightBlueBackground`: Light blue - `LightPurpleBackground`: Light purple - `PaleGrayBackground`: Medium gray - `DarkGrayBackground`: Gray - `DarkRedBackground`: Medium red - `DarkOrangeBackground`: Medium orange - `DarkYellowBackground`: Medium yellow - `DarkGreenBackground`: Medium green - `DarkBlueBackground`: Medium blue - `DarkPurpleBackground`: Medium purple  |
|       `indentation_level` | `string` | No | First line indentation level **Example value**: "NoIndent" **Optional values are**:  - `NoIndent`: No indent - `OneLevelIndent`: One Level indent  |
|     `elements` | `text_element[]` | Yes | Text element |
|       `text_run` | `text_run` | No | Text. Supports Page, Text, Heading1~9, Bullet, Ordered, Code, Quote and Todo blocks |
|         `content` | `string` | Yes | Text content **Example value**: "text" |
|         `text_element_style` | `text_element_style` | No | Text local style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `mention_user` | `mention_user` | No | @User. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|         `user_id` | `string` | Yes | User OpenID **Example value**: "ou_3bbe8a09c20e89cce9bff989ed840674" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `mention_doc` | `mention_doc` | No | @Document. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|         `token` | `string` | Yes | Docs token **Example value**: "doxbc873Y7cXD153gXqb76abcef" |
|         `obj_type` | `int` | Yes | Docs type **Example value**: 22 **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|         `url` | `string` | Yes | Docs link (url_encode) **Example value**: "https%3A%2F%2Fexample.larksuite.com%2Fdocx%2Fdoxbc873Y7cXD153gXqb76G1Y9b" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `reminder` | `reminder` | No | Date Reminder. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|         `create_user_id` | `string` | Yes | Creator user ID **Example value**: "0E2633a3-aa1a-4171-af9e-0768ff863566" |
|         `is_whole_day` | `boolean` | No | Whether it's a date or a whole hour **Example value**: true **Default value**: `false` |
|         `expire_time` | `string` | Yes | Time the event occurred (millisecond timestamp) **Example value**: "1641967200000" |
|         `notify_time` | `string` | Yes | Time to trigger notification (millisecond timestamp) **Example value**: "1643166000000" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `file` | `inline_file` | No | Inline file. Only supports deleting or moving the element, and creating new inline files is not supported |
|         `file_token` | `string` | No | File token **Example value**: "boxcnOj88GDkmWGm2zsTyCBqoLb" |
|         `source_block_id` | `string` | No | The ID of the block in which the file is located in the current document **Example value**: "doxcnM46kSWSkgUMW04ldKabcef" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `inline_block` | `inline_block` | No | Inline block. Only supports deleting or moving the element, and creating new inline blocks is not supported |
|         `block_id` | `string` | Yes | Block ID **Example value**: "doxcnPFi0R56ctbvh2Mjkkabcef" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `equation` | `equation` | No | Equation |
|         `content` | `string` | Yes | Equation content that conforms to KaTeX syntax. For more information, see [KaTex syntax rules](https://katex.org/docs/supported.html]. **Example value**: "E=mc^2\n" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|   `code` | `text` | No | Code Block |
|     `style` | `text_style` | No | Text style |
|       `align` | `int` | No | Alignment **Example value**: 1 **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  **Default value**: `1` |
|       `done` | `boolean` | No | Todo's status. Supports Todo Blocks **Example value**: true **Default value**: `false` |
|       `folded` | `boolean` | No | Folded state of text. Supports Heading1~9 block and the Text, Bullet, Ordered and Todo blocks with children **Example value**: true **Default value**: `false` |
|       `language` | `int` | No | Code block language. Only supports Code block **Example value**: 1 **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|       `wrap` | `boolean` | No | Whether the code block wraps automatically. Only supports the Code block **Example value**: true **Default value**: `false` |
|       `background_color` | `string` | No | The background color of the block **Example value**: "LightGrayBackground" **Optional values are**:  - `LightGrayBackground`: Light gray - `LightRedBackground`: Light red - `LightOrangeBackground`: Light orange - `LightYellowBackground`: Light yellow - `LightGreenBackground`: Light green - `LightBlueBackground`: Light blue - `LightPurpleBackground`: Light purple - `PaleGrayBackground`: Medium gray - `DarkGrayBackground`: Gray - `DarkRedBackground`: Medium red - `DarkOrangeBackground`: Medium orange - `DarkYellowBackground`: Medium yellow - `DarkGreenBackground`: Medium green - `DarkBlueBackground`: Medium blue - `DarkPurpleBackground`: Medium purple  |
|       `indentation_level` | `string` | No | First line indentation level **Example value**: "NoIndent" **Optional values are**:  - `NoIndent`: No indent - `OneLevelIndent`: One Level indent  |
|     `elements` | `text_element[]` | Yes | Text element |
|       `text_run` | `text_run` | No | Text. Supports Page, Text, Heading1~9, Bullet, Ordered, Code, Quote and Todo blocks |
|         `content` | `string` | Yes | Text content **Example value**: "text" |
|         `text_element_style` | `text_element_style` | No | Text local style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `mention_user` | `mention_user` | No | @User. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|         `user_id` | `string` | Yes | User OpenID **Example value**: "ou_3bbe8a09c20e89cce9bff989ed840674" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `mention_doc` | `mention_doc` | No | @Document. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|         `token` | `string` | Yes | Docs token **Example value**: "doxbc873Y7cXD153gXqb76abcef" |
|         `obj_type` | `int` | Yes | Docs type **Example value**: 22 **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|         `url` | `string` | Yes | Docs link (url_encode) **Example value**: "https%3A%2F%2Fexample.larksuite.com%2Fdocx%2Fdoxbc873Y7cXD153gXqb76G1Y9b" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `reminder` | `reminder` | No | Date Reminder. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|         `create_user_id` | `string` | Yes | Creator user ID **Example value**: "0E2633a3-aa1a-4171-af9e-0768ff863566" |
|         `is_whole_day` | `boolean` | No | Whether it's a date or a whole hour **Example value**: true **Default value**: `false` |
|         `expire_time` | `string` | Yes | Time the event occurred (millisecond timestamp) **Example value**: "1641967200000" |
|         `notify_time` | `string` | Yes | Time to trigger notification (millisecond timestamp) **Example value**: "1643166000000" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `file` | `inline_file` | No | Inline file. Only supports deleting or moving the element, and creating new inline files is not supported |
|         `file_token` | `string` | No | File token **Example value**: "boxcnOj88GDkmWGm2zsTyCBqoLb" |
|         `source_block_id` | `string` | No | The ID of the block in which the file is located in the current document **Example value**: "doxcnM46kSWSkgUMW04ldKabcef" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `inline_block` | `inline_block` | No | Inline block. Only supports deleting or moving the element, and creating new inline blocks is not supported |
|         `block_id` | `string` | Yes | Block ID **Example value**: "doxcnPFi0R56ctbvh2Mjkkabcef" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `equation` | `equation` | No | Equation |
|         `content` | `string` | Yes | Equation content that conforms to KaTeX syntax. For more information, see [KaTex syntax rules](https://katex.org/docs/supported.html]. **Example value**: "E=mc^2\n" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|   `quote` | `text` | No | Quote Block |
|     `style` | `text_style` | No | Text style |
|       `align` | `int` | No | Alignment **Example value**: 1 **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  **Default value**: `1` |
|       `done` | `boolean` | No | Todo's status. Supports Todo Blocks **Example value**: true **Default value**: `false` |
|       `folded` | `boolean` | No | Folded state of text. Supports Heading1~9 block and the Text, Bullet, Ordered and Todo blocks with children **Example value**: true **Default value**: `false` |
|       `language` | `int` | No | Code block language. Only supports Code block **Example value**: 1 **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|       `wrap` | `boolean` | No | Whether the code block wraps automatically. Only supports the Code block **Example value**: true **Default value**: `false` |
|       `background_color` | `string` | No | The background color of the block **Example value**: "LightGrayBackground" **Optional values are**:  - `LightGrayBackground`: Light gray - `LightRedBackground`: Light red - `LightOrangeBackground`: Light orange - `LightYellowBackground`: Light yellow - `LightGreenBackground`: Light green - `LightBlueBackground`: Light blue - `LightPurpleBackground`: Light purple - `PaleGrayBackground`: Medium gray - `DarkGrayBackground`: Gray - `DarkRedBackground`: Medium red - `DarkOrangeBackground`: Medium orange - `DarkYellowBackground`: Medium yellow - `DarkGreenBackground`: Medium green - `DarkBlueBackground`: Medium blue - `DarkPurpleBackground`: Medium purple  |
|       `indentation_level` | `string` | No | First line indentation level **Example value**: "NoIndent" **Optional values are**:  - `NoIndent`: No indent - `OneLevelIndent`: One Level indent  |
|     `elements` | `text_element[]` | Yes | Text element |
|       `text_run` | `text_run` | No | Text. Supports Page, Text, Heading1~9, Bullet, Ordered, Code, Quote and Todo blocks |
|         `content` | `string` | Yes | Text content **Example value**: "text" |
|         `text_element_style` | `text_element_style` | No | Text local style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `mention_user` | `mention_user` | No | @User. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|         `user_id` | `string` | Yes | User OpenID **Example value**: "ou_3bbe8a09c20e89cce9bff989ed840674" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `mention_doc` | `mention_doc` | No | @Document. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|         `token` | `string` | Yes | Docs token **Example value**: "doxbc873Y7cXD153gXqb76abcef" |
|         `obj_type` | `int` | Yes | Docs type **Example value**: 22 **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|         `url` | `string` | Yes | Docs link (url_encode) **Example value**: "https%3A%2F%2Fexample.larksuite.com%2Fdocx%2Fdoxbc873Y7cXD153gXqb76G1Y9b" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `reminder` | `reminder` | No | Date Reminder. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|         `create_user_id` | `string` | Yes | Creator user ID **Example value**: "0E2633a3-aa1a-4171-af9e-0768ff863566" |
|         `is_whole_day` | `boolean` | No | Whether it's a date or a whole hour **Example value**: true **Default value**: `false` |
|         `expire_time` | `string` | Yes | Time the event occurred (millisecond timestamp) **Example value**: "1641967200000" |
|         `notify_time` | `string` | Yes | Time to trigger notification (millisecond timestamp) **Example value**: "1643166000000" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `file` | `inline_file` | No | Inline file. Only supports deleting or moving the element, and creating new inline files is not supported |
|         `file_token` | `string` | No | File token **Example value**: "boxcnOj88GDkmWGm2zsTyCBqoLb" |
|         `source_block_id` | `string` | No | The ID of the block in which the file is located in the current document **Example value**: "doxcnM46kSWSkgUMW04ldKabcef" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `inline_block` | `inline_block` | No | Inline block. Only supports deleting or moving the element, and creating new inline blocks is not supported |
|         `block_id` | `string` | Yes | Block ID **Example value**: "doxcnPFi0R56ctbvh2Mjkkabcef" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `equation` | `equation` | No | Equation |
|         `content` | `string` | Yes | Equation content that conforms to KaTeX syntax. For more information, see [KaTex syntax rules](https://katex.org/docs/supported.html]. **Example value**: "E=mc^2\n" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|   `equation` | `text` | No | Equation Block |
|     `style` | `text_style` | No | Text style |
|       `align` | `int` | No | Alignment **Example value**: 1 **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  **Default value**: `1` |
|       `done` | `boolean` | No | Todo's status. Supports Todo Blocks **Example value**: true **Default value**: `false` |
|       `folded` | `boolean` | No | Folded state of text. Supports Heading1~9 block and the Text, Bullet, Ordered and Todo blocks with children **Example value**: true **Default value**: `false` |
|       `language` | `int` | No | Code block language. Only supports Code block **Example value**: 1 **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|       `wrap` | `boolean` | No | Whether the code block wraps automatically. Only supports the Code block **Example value**: true **Default value**: `false` |
|       `background_color` | `string` | No | The background color of the block **Example value**: "LightGrayBackground" **Optional values are**:  - `LightGrayBackground`: Light gray - `LightRedBackground`: Light red - `LightOrangeBackground`: Light orange - `LightYellowBackground`: Light yellow - `LightGreenBackground`: Light green - `LightBlueBackground`: Light blue - `LightPurpleBackground`: Light purple - `PaleGrayBackground`: Medium gray - `DarkGrayBackground`: Gray - `DarkRedBackground`: Medium red - `DarkOrangeBackground`: Medium orange - `DarkYellowBackground`: Medium yellow - `DarkGreenBackground`: Medium green - `DarkBlueBackground`: Medium blue - `DarkPurpleBackground`: Medium purple  |
|       `indentation_level` | `string` | No | First line indentation level **Example value**: "NoIndent" **Optional values are**:  - `NoIndent`: No indent - `OneLevelIndent`: One Level indent  |
|     `elements` | `text_element[]` | Yes | Text element |
|       `text_run` | `text_run` | No | Text. Supports Page, Text, Heading1~9, Bullet, Ordered, Code, Quote and Todo blocks |
|         `content` | `string` | Yes | Text content **Example value**: "text" |
|         `text_element_style` | `text_element_style` | No | Text local style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `mention_user` | `mention_user` | No | @User. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|         `user_id` | `string` | Yes | User OpenID **Example value**: "ou_3bbe8a09c20e89cce9bff989ed840674" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `mention_doc` | `mention_doc` | No | @Document. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|         `token` | `string` | Yes | Docs token **Example value**: "doxbc873Y7cXD153gXqb76abcef" |
|         `obj_type` | `int` | Yes | Docs type **Example value**: 22 **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|         `url` | `string` | Yes | Docs link (url_encode) **Example value**: "https%3A%2F%2Fexample.larksuite.com%2Fdocx%2Fdoxbc873Y7cXD153gXqb76G1Y9b" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `reminder` | `reminder` | No | Date Reminder. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|         `create_user_id` | `string` | Yes | Creator user ID **Example value**: "0E2633a3-aa1a-4171-af9e-0768ff863566" |
|         `is_whole_day` | `boolean` | No | Whether it's a date or a whole hour **Example value**: true **Default value**: `false` |
|         `expire_time` | `string` | Yes | Time the event occurred (millisecond timestamp) **Example value**: "1641967200000" |
|         `notify_time` | `string` | Yes | Time to trigger notification (millisecond timestamp) **Example value**: "1643166000000" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `file` | `inline_file` | No | Inline file. Only supports deleting or moving the element, and creating new inline files is not supported |
|         `file_token` | `string` | No | File token **Example value**: "boxcnOj88GDkmWGm2zsTyCBqoLb" |
|         `source_block_id` | `string` | No | The ID of the block in which the file is located in the current document **Example value**: "doxcnM46kSWSkgUMW04ldKabcef" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `inline_block` | `inline_block` | No | Inline block. Only supports deleting or moving the element, and creating new inline blocks is not supported |
|         `block_id` | `string` | Yes | Block ID **Example value**: "doxcnPFi0R56ctbvh2Mjkkabcef" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `equation` | `equation` | No | Equation |
|         `content` | `string` | Yes | Equation content that conforms to KaTeX syntax. For more information, see [KaTex syntax rules](https://katex.org/docs/supported.html]. **Example value**: "E=mc^2\n" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|   `todo` | `text` | No | Todo Block |
|     `style` | `text_style` | No | Text style |
|       `align` | `int` | No | Alignment **Example value**: 1 **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  **Default value**: `1` |
|       `done` | `boolean` | No | Todo's status. Supports Todo Blocks **Example value**: true **Default value**: `false` |
|       `folded` | `boolean` | No | Folded state of text. Supports Heading1~9 block and the Text, Bullet, Ordered and Todo blocks with children **Example value**: true **Default value**: `false` |
|       `language` | `int` | No | Code block language. Only supports Code block **Example value**: 1 **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|       `wrap` | `boolean` | No | Whether the code block wraps automatically. Only supports the Code block **Example value**: true **Default value**: `false` |
|       `background_color` | `string` | No | The background color of the block **Example value**: "LightGrayBackground" **Optional values are**:  - `LightGrayBackground`: Light gray - `LightRedBackground`: Light red - `LightOrangeBackground`: Light orange - `LightYellowBackground`: Light yellow - `LightGreenBackground`: Light green - `LightBlueBackground`: Light blue - `LightPurpleBackground`: Light purple - `PaleGrayBackground`: Medium gray - `DarkGrayBackground`: Gray - `DarkRedBackground`: Medium red - `DarkOrangeBackground`: Medium orange - `DarkYellowBackground`: Medium yellow - `DarkGreenBackground`: Medium green - `DarkBlueBackground`: Medium blue - `DarkPurpleBackground`: Medium purple  |
|       `indentation_level` | `string` | No | First line indentation level **Example value**: "NoIndent" **Optional values are**:  - `NoIndent`: No indent - `OneLevelIndent`: One Level indent  |
|     `elements` | `text_element[]` | Yes | Text element |
|       `text_run` | `text_run` | No | Text. Supports Page, Text, Heading1~9, Bullet, Ordered, Code, Quote and Todo blocks |
|         `content` | `string` | Yes | Text content **Example value**: "text" |
|         `text_element_style` | `text_element_style` | No | Text local style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `mention_user` | `mention_user` | No | @User. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|         `user_id` | `string` | Yes | User OpenID **Example value**: "ou_3bbe8a09c20e89cce9bff989ed840674" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `mention_doc` | `mention_doc` | No | @Document. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|         `token` | `string` | Yes | Docs token **Example value**: "doxbc873Y7cXD153gXqb76abcef" |
|         `obj_type` | `int` | Yes | Docs type **Example value**: 22 **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|         `url` | `string` | Yes | Docs link (url_encode) **Example value**: "https%3A%2F%2Fexample.larksuite.com%2Fdocx%2Fdoxbc873Y7cXD153gXqb76G1Y9b" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `reminder` | `reminder` | No | Date Reminder. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|         `create_user_id` | `string` | Yes | Creator user ID **Example value**: "0E2633a3-aa1a-4171-af9e-0768ff863566" |
|         `is_whole_day` | `boolean` | No | Whether it's a date or a whole hour **Example value**: true **Default value**: `false` |
|         `expire_time` | `string` | Yes | Time the event occurred (millisecond timestamp) **Example value**: "1641967200000" |
|         `notify_time` | `string` | Yes | Time to trigger notification (millisecond timestamp) **Example value**: "1643166000000" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `file` | `inline_file` | No | Inline file. Only supports deleting or moving the element, and creating new inline files is not supported |
|         `file_token` | `string` | No | File token **Example value**: "boxcnOj88GDkmWGm2zsTyCBqoLb" |
|         `source_block_id` | `string` | No | The ID of the block in which the file is located in the current document **Example value**: "doxcnM46kSWSkgUMW04ldKabcef" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `inline_block` | `inline_block` | No | Inline block. Only supports deleting or moving the element, and creating new inline blocks is not supported |
|         `block_id` | `string` | Yes | Block ID **Example value**: "doxcnPFi0R56ctbvh2Mjkkabcef" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|       `equation` | `equation` | No | Equation |
|         `content` | `string` | Yes | Equation content that conforms to KaTeX syntax. For more information, see [KaTex syntax rules](https://katex.org/docs/supported.html]. **Example value**: "E=mc^2\n" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. **Example value**: ["1660030311959965796"] |
|   `bitable` | `bitable` | No | Bitable Block |
|     `view_type` | `int` | No | View type **Example value**: 1 **Optional values are**:  - `1`: Data sheet - `2`: Kanban  |
|   `callout` | `callout` | No | Callout Block |
|     `background_color` | `int` | No | Callout block background color **Example value**: 1 **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium gray - `8`: Medium Red - `9`: Medium orange - `10`: Medium yellow - `11`: Medium green - `12`: Medium blue - `13`: Medium purple - `14`: Gray - `15`: Light Gray  |
|     `border_color` | `int` | No | Border color **Example value**: 1 **Optional values are**:  - `1`: Red - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Gray  |
|     `text_color` | `int` | No | Text color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|   `chat_card` | `chat_card` | No | Chat Card Block |
|     `chat_id` | `string` | Yes | Chat ID **Example value**: "7052227140476993555" |
|     `align` | `int` | No | Alignment **Example value**: 1 **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  **Default value**: `1` |
|   `divider` | `divider` | No | Divider Block |
|   `file` | `file` | No | File Block |
|     `view_type` | `int` | No | View type **Example value**: 1 **Optional values are**:  - `1`: Card view - `2`: Preview view  |
|   `grid` | `grid` | No | Grid Block |
|     `column_size` | `int` | Yes | Number of columns **Example value**: 2 **Data validation rules**: - Value range: `2` ～ `5` |
|   `iframe` | `iframe` | No | Iframe Block |
|     `component` | `iframe_component` | Yes | Component elements of iframe |
|       `iframe_type` | `int` | No | Iframe Type **Example value**: 1 **Optional values are**:  - `1`: Bilibili - `2`: Xigua Video - `3`: Youku - `4`: Airtable - `5`: Baidu Map - `6`: Autonavi Map - `7`: TikTok - `8`: Figma - `9`: Ink knife - `10`: Canva - `11`: CodePen - `12`: Lark Survey - `13`: Gold data - `14`: GoogleMap - `15`: Youtube - `99`: Other  |
|       `url` | `string` | Yes | Iframe target url (requires url_encode) **Example value**: "https%3A%2F%2Fwww.bilibili.com%2Fvideo%2FBV1Hi4y1w7V7" |
|   `image` | `image` | No | Image Block |
|     `align` | `int` | No | alignment **Example value**: 2 **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  **Default value**: `2` |
|   `isv` | `isv` | No | ISV Block |
|   `add_ons` | `add_ons` | No | Add-ons |
|     `component_id` | `string` | No | Document widget ID **Example value**: "7056882725002051603" |
|     `component_type_id` | `string` | Yes | Document widget types, such as Q & A interactive "blk_636a0a6657db8001c8df5488" **Example value**: "blk_636a0a6657db8001c8df5488" |
|     `record` | `string` | No | Document widget content data, JSON string **Example value**: ""{}"" |
|   `sheet` | `sheet` | No | Sheet Block |
|     `row_size` | `int` | No | Number of spreadsheet rows **Example value**: 2 **Data validation rules**: - Value range: `1` ～ `9` |
|     `column_size` | `int` | No | Number of spreadsheet columns **Example value**: 2 **Data validation rules**: - Value range: `1` ～ `9` |
|   `table` | `table` | No | Table Block |
|     `property` | `table_property` | Yes | Table Properties |
|       `row_size` | `int` | Yes | Number of rows **Example value**: 1 **Data validation rules**: - Minimum value: `1` |
|       `column_size` | `int` | Yes | Number of columns **Example value**: 1 **Data validation rules**: - Minimum value: `1` |
|       `column_width` | `int[]` | No | Column width in px **Example value**: 100 **Default value**: `100` |
|       `merge_info` | `table_merge_info[]` | No | Cell merge information. This property is read-only when creating a table and is generated by the backend. If you need to merge cells, you can do so by using the subrequest merge_table_cells of the Update Block api. |
|         `row_span` | `int` | No | Number of consecutive rows merged from the current line **Example value**: 2 **Data validation rules**: - Minimum value: `1` |
|         `col_span` | `int` | No | Number of consecutive columns that are merged from the current pristine **Example value**: 2 **Data validation rules**: - Minimum value: `1` |
|       `header_row` | `boolean` | No | Set the first row as the header row. **Example value**: false **Default value**: `false` |
|       `header_column` | `boolean` | No | Set the first column as the header column. **Example value**: false **Default value**: `false` |
|   `quote_container` | `quote_container` | No | Quote Container Block |
|   `task` | `task` | No | Task Block |
|     `folded` | `boolean` | No | Folded state **Example value**: false **Default value**: `false` |
|   `okr` | `okr` | No | OKR blocks. It can only be created when using Authorization with `user_access_token` |
|     `okr_id` | `string` | No | OKR ID. Get the OKR ID that needs to be inserted and see Get the user's OKR list **Example value**: "7076349900476448796" |
|     `objectives` | `objective_id_with_kr_id[]` | No | The objective ID and key result ID in the OKR Block, insert all objective and key result under okr when objectives is null |
|       `objective_id` | `string` | No | Objective ID in OKR **Example value**: "7109022409227026460" |
|       `kr_ids` | `string[]` | No | List of IDs of key results, insert all key results under the current objective when kr_ids is null or empty **Example value**: ["7109022573011894300","7109022546444517404"] |
|   `comment_ids` | `string[]` | No | Comment ids **Example value**: [1660030311959965796] |
|   `wiki_catalog` | `wiki_catalog` | No | Wiki Catalog Block |
|     `wiki_token` | `string` | No | wiki token **Example value**: "Ub47wVl7AikG9wkgnpSbFy4EcAc" |
|   `board` | `board` | No | Board Block |
|     `align` | `int` | No | Alignment **Example value**: 2 **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  **Default value**: `1` |
|     `width` | `int` | No | The width of the board, measured in pixels (px); when left blank, it will automatically adapt to the document's width. If the specified value exceeds the maximum width of the document, the page rendering will be constrained to the document's maximum width. **Example value**: 300 **Data validation rules**: - Minimum value: `90` |
|     `height` | `int` | No | The height of the board, measured in pixels (px); when left blank, it will be automatically calculated based on the content of the board. If the specified value exceeds twice the height of the screen, the page rendering will be constrained to twice the screen height. **Example value**: 300 **Data validation rules**: - Minimum value: `90` |
|   `link_preview` | `link_preview` | No | Link Preview Block |
|     `url` | `string` | Yes | link **Example value**: "https://applink.larksuite.com/client/message/link/open?token=Al7F******Q%3D" **Data validation rules**: - Minimum length: `1` characters |
|     `url_type` | `string` | Yes | Link type **Example value**: "MessageLink" **Optional values are**:  - `MessageLink`: Message link - `Undefined`: Undefined link type  |
| `index` | `int` | No | Specifies the placement of the newly created child blocks in the parent block. The starting value of the index is 0, indicating the first position in the child block list; The maximum value of the index is the number of child blocks of the parent block, representing the last position in the child block list. For example, if a parent block has 5 child blocks, their indexes are 0, 1, 2, 3, and 4, respectively. If you want to place a newly created block in the first position of the parent block, the index value should be 0; If you want to place the newly created block in the last position, the index value should be -1 **Example value**: 0 **Default value**: `-1` | ### Request body example

{
  "index": 0,
  "children": [
    {
      "block_type": 2,
      "text": {
        "elements": [
          {
            "text_run": {
              "content": "When you need to convey important information to all members of the group, you can use the group announcement function. Subsequent members of the group can also see the current group announcement.",
              "text_element_style": {
                "background_color": 14,
                "text_color": 5
              }
            }
          },
          {
            "text_run": {
              "content": "Group announcement editing supports real-time saving and multi-person collaboration, and group members can view the latest group announcement content in real time.",
              "text_element_style": {
                "background_color": 14,
                "bold": true,
                "text_color": 5
              }
            }
          }
        ],
        "style": {}
      }
    }
  ]
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `children` | `block[]` | Block information for the newly created child blocks |
|     `block_id` | `string` | The unique identification of a block |
|     `parent_id` | `string` | The ID of a block's parent |
|     `children` | `string[]` | ID list of a block's children |
|     `block_type` | `int` | Block type **Optional values are**:  - `1`: Page Block - `2`: Text Block - `3`: Level 1 Heading Block - `4`: Level 2 Heading Block - `5`: Level 3 Heading Block - `6`: Level 4 Heading Block - `7`: Level 5 Heading Block - `8`: Level 6 Heading Block - `9`: Level 7 Heading Block - `10`: Level 8 Heading Block - `11`: Level 9 Heading Block - `12`: Bullet List Block - `13`: Ordered List Block - `14`: Code Block - `15`: Quote Block - `17`: Todo Block - `18`: Bitable Block - `19`: Callout Block - `20`: Chat Card Block - `21`: UML Diagram Block - `22`: Divider Block - `23`: File Block - `24`: Grid Block - `25`: Grid Column Block - `26`: Iframe Block - `27`: Image Block - `28`: ISV Block - `29`: Mindnote Block - `30`: Sheet Block - `31`: Table Block - `32`: Table Cell Block - `33`: View Block - `34`: Quote Container Block - `35`: Task Block - `36`: OKR Block - `37`: OKR Objective Block - `38`: OKR Key Result Block - `39`: OKR Progress Block - `40`: Add-Ons Block - `41`: Jira Issue Block - `42`: Wiki Catalog Block - `43`: Board - `44`: Agenda Block - `45`: Agenda Item Block - `46`: Agenda Item Title Block - `47`: Agenda Item Content Block - `48`: Link Preview Block - `999`: Blocks not supported  |
|     `page` | `text` | Document Block |
|       `style` | `text_style` | Text style |
|         `align` | `int` | Alignment **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  |
|         `done` | `boolean` | Todo's status. Supports Todo Blocks |
|         `folded` | `boolean` | Folded state of text. Supports Heading1~9 block and the Text, Bullet, Ordered and Todo blocks with children |
|         `language` | `int` | Code block language. Only supports Code block **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|         `wrap` | `boolean` | Whether the code block wraps automatically. Only supports the Code block |
|         `background_color` | `string` | The background color of the block **Optional values are**:  - `LightGrayBackground`: Light gray - `LightRedBackground`: Light red - `LightOrangeBackground`: Light orange - `LightYellowBackground`: Light yellow - `LightGreenBackground`: Light green - `LightBlueBackground`: Light blue - `LightPurpleBackground`: Light purple - `PaleGrayBackground`: Medium gray - `DarkGrayBackground`: Gray - `DarkRedBackground`: Medium red - `DarkOrangeBackground`: Medium orange - `DarkYellowBackground`: Medium yellow - `DarkGreenBackground`: Medium green - `DarkBlueBackground`: Medium blue - `DarkPurpleBackground`: Medium purple  |
|         `indentation_level` | `string` | First line indentation level **Optional values are**:  - `NoIndent`: No indent - `OneLevelIndent`: One Level indent  |
|         `sequence` | `string` | Used to determine the number of an ordered list, either a number or 'auto' - Start a new list, numbering starts at 1, sequence='1' - Manually change the numbering, the numbering is set to a specific number, e.g. sequence='3' - Continue numbering, numbering is automatically consecutive, sequence='auto' - Some historical data and ordered block created through OpenAPI do not return this field |
|       `elements` | `text_element[]` | Text element |
|         `text_run` | `text_run` | Text. Supports Page, Text, Heading1~9, Bullet, Ordered, Code, Quote and Todo blocks |
|           `content` | `string` | Text content |
|           `text_element_style` | `text_element_style` | Text local style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `mention_user` | `mention_user` | @User. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|           `user_id` | `string` | User OpenID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `mention_doc` | `mention_doc` | @Document. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|           `token` | `string` | Docs token |
|           `obj_type` | `int` | Docs type **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|           `url` | `string` | Docs link (url_encode) |
|           `title` | `string` | The title of the referenced document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `reminder` | `reminder` | Date Reminder. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|           `create_user_id` | `string` | Creator user ID |
|           `is_notify` | `boolean` | Whether to notify |
|           `is_whole_day` | `boolean` | Whether it's a date or a whole hour |
|           `expire_time` | `string` | Time the event occurred (millisecond timestamp) |
|           `notify_time` | `string` | Time to trigger notification (millisecond timestamp) |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `file` | `inline_file` | Inline file. Only supports deleting or moving the element, and creating new inline files is not supported |
|           `file_token` | `string` | File token |
|           `source_block_id` | `string` | The ID of the block in which the file is located in the current document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `undefined` | `undefined_element` | Unsupported TextElements |
|         `inline_block` | `inline_block` | Inline block. Only supports deleting or moving the element, and creating new inline blocks is not supported |
|           `block_id` | `string` | Block ID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `equation` | `equation` | Equation |
|           `content` | `string` | Equation content that conforms to KaTeX syntax. For more information, see [KaTex syntax rules](https://katex.org/docs/supported.html]. |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|     `text` | `text` | Text Block |
|       `style` | `text_style` | Text style |
|         `align` | `int` | Alignment **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  |
|         `done` | `boolean` | Todo's status. Supports Todo Blocks |
|         `folded` | `boolean` | Folded state of text. Supports Heading1~9 block and the Text, Bullet, Ordered and Todo blocks with children |
|         `language` | `int` | Code block language. Only supports Code block **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|         `wrap` | `boolean` | Whether the code block wraps automatically. Only supports the Code block |
|         `background_color` | `string` | The background color of the block **Optional values are**:  - `LightGrayBackground`: Light gray - `LightRedBackground`: Light red - `LightOrangeBackground`: Light orange - `LightYellowBackground`: Light yellow - `LightGreenBackground`: Light green - `LightBlueBackground`: Light blue - `LightPurpleBackground`: Light purple - `PaleGrayBackground`: Medium gray - `DarkGrayBackground`: Gray - `DarkRedBackground`: Medium red - `DarkOrangeBackground`: Medium orange - `DarkYellowBackground`: Medium yellow - `DarkGreenBackground`: Medium green - `DarkBlueBackground`: Medium blue - `DarkPurpleBackground`: Medium purple  |
|         `indentation_level` | `string` | First line indentation level **Optional values are**:  - `NoIndent`: No indent - `OneLevelIndent`: One Level indent  |
|         `sequence` | `string` | Used to determine the number of an ordered list, either a number or 'auto' - Start a new list, numbering starts at 1, sequence='1' - Manually change the numbering, the numbering is set to a specific number, e.g. sequence='3' - Continue numbering, numbering is automatically consecutive, sequence='auto' - Some historical data and ordered block created through OpenAPI do not return this field |
|       `elements` | `text_element[]` | Text element |
|         `text_run` | `text_run` | Text. Supports Page, Text, Heading1~9, Bullet, Ordered, Code, Quote and Todo blocks |
|           `content` | `string` | Text content |
|           `text_element_style` | `text_element_style` | Text local style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `mention_user` | `mention_user` | @User. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|           `user_id` | `string` | User OpenID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `mention_doc` | `mention_doc` | @Document. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|           `token` | `string` | Docs token |
|           `obj_type` | `int` | Docs type **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|           `url` | `string` | Docs link (url_encode) |
|           `title` | `string` | The title of the referenced document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `reminder` | `reminder` | Date Reminder. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|           `create_user_id` | `string` | Creator user ID |
|           `is_notify` | `boolean` | Whether to notify |
|           `is_whole_day` | `boolean` | Whether it's a date or a whole hour |
|           `expire_time` | `string` | Time the event occurred (millisecond timestamp) |
|           `notify_time` | `string` | Time to trigger notification (millisecond timestamp) |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `file` | `inline_file` | Inline file. Only supports deleting or moving the element, and creating new inline files is not supported |
|           `file_token` | `string` | File token |
|           `source_block_id` | `string` | The ID of the block in which the file is located in the current document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `undefined` | `undefined_element` | Unsupported TextElements |
|         `inline_block` | `inline_block` | Inline block. Only supports deleting or moving the element, and creating new inline blocks is not supported |
|           `block_id` | `string` | Block ID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `equation` | `equation` | Equation |
|           `content` | `string` | Equation content that conforms to KaTeX syntax. For more information, see [KaTex syntax rules](https://katex.org/docs/supported.html]. |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|     `heading1` | `text` | Level 1 Heading Block |
|       `style` | `text_style` | Text style |
|         `align` | `int` | Alignment **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  |
|         `done` | `boolean` | Todo's status. Supports Todo Blocks |
|         `folded` | `boolean` | Folded state of text. Supports Heading1~9 block and the Text, Bullet, Ordered and Todo blocks with children |
|         `language` | `int` | Code block language. Only supports Code block **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|         `wrap` | `boolean` | Whether the code block wraps automatically. Only supports the Code block |
|         `background_color` | `string` | The background color of the block **Optional values are**:  - `LightGrayBackground`: Light gray - `LightRedBackground`: Light red - `LightOrangeBackground`: Light orange - `LightYellowBackground`: Light yellow - `LightGreenBackground`: Light green - `LightBlueBackground`: Light blue - `LightPurpleBackground`: Light purple - `PaleGrayBackground`: Medium gray - `DarkGrayBackground`: Gray - `DarkRedBackground`: Medium red - `DarkOrangeBackground`: Medium orange - `DarkYellowBackground`: Medium yellow - `DarkGreenBackground`: Medium green - `DarkBlueBackground`: Medium blue - `DarkPurpleBackground`: Medium purple  |
|         `indentation_level` | `string` | First line indentation level **Optional values are**:  - `NoIndent`: No indent - `OneLevelIndent`: One Level indent  |
|         `sequence` | `string` | Used to determine the number of an ordered list, either a number or 'auto' - Start a new list, numbering starts at 1, sequence='1' - Manually change the numbering, the numbering is set to a specific number, e.g. sequence='3' - Continue numbering, numbering is automatically consecutive, sequence='auto' - Some historical data and ordered block created through OpenAPI do not return this field |
|       `elements` | `text_element[]` | Text element |
|         `text_run` | `text_run` | Text. Supports Page, Text, Heading1~9, Bullet, Ordered, Code, Quote and Todo blocks |
|           `content` | `string` | Text content |
|           `text_element_style` | `text_element_style` | Text local style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `mention_user` | `mention_user` | @User. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|           `user_id` | `string` | User OpenID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `mention_doc` | `mention_doc` | @Document. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|           `token` | `string` | Docs token |
|           `obj_type` | `int` | Docs type **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|           `url` | `string` | Docs link (url_encode) |
|           `title` | `string` | The title of the referenced document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `reminder` | `reminder` | Date Reminder. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|           `create_user_id` | `string` | Creator user ID |
|           `is_notify` | `boolean` | Whether to notify |
|           `is_whole_day` | `boolean` | Whether it's a date or a whole hour |
|           `expire_time` | `string` | Time the event occurred (millisecond timestamp) |
|           `notify_time` | `string` | Time to trigger notification (millisecond timestamp) |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `file` | `inline_file` | Inline file. Only supports deleting or moving the element, and creating new inline files is not supported |
|           `file_token` | `string` | File token |
|           `source_block_id` | `string` | The ID of the block in which the file is located in the current document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `undefined` | `undefined_element` | Unsupported TextElements |
|         `inline_block` | `inline_block` | Inline block. Only supports deleting or moving the element, and creating new inline blocks is not supported |
|           `block_id` | `string` | Block ID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `equation` | `equation` | Equation |
|           `content` | `string` | Equation content that conforms to KaTeX syntax. For more information, see [KaTex syntax rules](https://katex.org/docs/supported.html]. |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|     `heading2` | `text` | Level 2 Heading Block |
|       `style` | `text_style` | Text style |
|         `align` | `int` | Alignment **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  |
|         `done` | `boolean` | Todo's status. Supports Todo Blocks |
|         `folded` | `boolean` | Folded state of text. Supports Heading1~9 block and the Text, Bullet, Ordered and Todo blocks with children |
|         `language` | `int` | Code block language. Only supports Code block **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|         `wrap` | `boolean` | Whether the code block wraps automatically. Only supports the Code block |
|         `background_color` | `string` | The background color of the block **Optional values are**:  - `LightGrayBackground`: Light gray - `LightRedBackground`: Light red - `LightOrangeBackground`: Light orange - `LightYellowBackground`: Light yellow - `LightGreenBackground`: Light green - `LightBlueBackground`: Light blue - `LightPurpleBackground`: Light purple - `PaleGrayBackground`: Medium gray - `DarkGrayBackground`: Gray - `DarkRedBackground`: Medium red - `DarkOrangeBackground`: Medium orange - `DarkYellowBackground`: Medium yellow - `DarkGreenBackground`: Medium green - `DarkBlueBackground`: Medium blue - `DarkPurpleBackground`: Medium purple  |
|         `indentation_level` | `string` | First line indentation level **Optional values are**:  - `NoIndent`: No indent - `OneLevelIndent`: One Level indent  |
|         `sequence` | `string` | Used to determine the number of an ordered list, either a number or 'auto' - Start a new list, numbering starts at 1, sequence='1' - Manually change the numbering, the numbering is set to a specific number, e.g. sequence='3' - Continue numbering, numbering is automatically consecutive, sequence='auto' - Some historical data and ordered block created through OpenAPI do not return this field |
|       `elements` | `text_element[]` | Text element |
|         `text_run` | `text_run` | Text. Supports Page, Text, Heading1~9, Bullet, Ordered, Code, Quote and Todo blocks |
|           `content` | `string` | Text content |
|           `text_element_style` | `text_element_style` | Text local style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `mention_user` | `mention_user` | @User. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|           `user_id` | `string` | User OpenID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `mention_doc` | `mention_doc` | @Document. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|           `token` | `string` | Docs token |
|           `obj_type` | `int` | Docs type **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|           `url` | `string` | Docs link (url_encode) |
|           `title` | `string` | The title of the referenced document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `reminder` | `reminder` | Date Reminder. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|           `create_user_id` | `string` | Creator user ID |
|           `is_notify` | `boolean` | Whether to notify |
|           `is_whole_day` | `boolean` | Whether it's a date or a whole hour |
|           `expire_time` | `string` | Time the event occurred (millisecond timestamp) |
|           `notify_time` | `string` | Time to trigger notification (millisecond timestamp) |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `file` | `inline_file` | Inline file. Only supports deleting or moving the element, and creating new inline files is not supported |
|           `file_token` | `string` | File token |
|           `source_block_id` | `string` | The ID of the block in which the file is located in the current document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `undefined` | `undefined_element` | Unsupported TextElements |
|         `inline_block` | `inline_block` | Inline block. Only supports deleting or moving the element, and creating new inline blocks is not supported |
|           `block_id` | `string` | Block ID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `equation` | `equation` | Equation |
|           `content` | `string` | Equation content that conforms to KaTeX syntax. For more information, see [KaTex syntax rules](https://katex.org/docs/supported.html]. |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|     `heading3` | `text` | Level 3 Heading Block |
|       `style` | `text_style` | Text style |
|         `align` | `int` | Alignment **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  |
|         `done` | `boolean` | Todo's status. Supports Todo Blocks |
|         `folded` | `boolean` | Folded state of text. Supports Heading1~9 block and the Text, Bullet, Ordered and Todo blocks with children |
|         `language` | `int` | Code block language. Only supports Code block **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|         `wrap` | `boolean` | Whether the code block wraps automatically. Only supports the Code block |
|         `background_color` | `string` | The background color of the block **Optional values are**:  - `LightGrayBackground`: Light gray - `LightRedBackground`: Light red - `LightOrangeBackground`: Light orange - `LightYellowBackground`: Light yellow - `LightGreenBackground`: Light green - `LightBlueBackground`: Light blue - `LightPurpleBackground`: Light purple - `PaleGrayBackground`: Medium gray - `DarkGrayBackground`: Gray - `DarkRedBackground`: Medium red - `DarkOrangeBackground`: Medium orange - `DarkYellowBackground`: Medium yellow - `DarkGreenBackground`: Medium green - `DarkBlueBackground`: Medium blue - `DarkPurpleBackground`: Medium purple  |
|         `indentation_level` | `string` | First line indentation level **Optional values are**:  - `NoIndent`: No indent - `OneLevelIndent`: One Level indent  |
|         `sequence` | `string` | Used to determine the number of an ordered list, either a number or 'auto' - Start a new list, numbering starts at 1, sequence='1' - Manually change the numbering, the numbering is set to a specific number, e.g. sequence='3' - Continue numbering, numbering is automatically consecutive, sequence='auto' - Some historical data and ordered block created through OpenAPI do not return this field |
|       `elements` | `text_element[]` | Text element |
|         `text_run` | `text_run` | Text. Supports Page, Text, Heading1~9, Bullet, Ordered, Code, Quote and Todo blocks |
|           `content` | `string` | Text content |
|           `text_element_style` | `text_element_style` | Text local style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `mention_user` | `mention_user` | @User. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|           `user_id` | `string` | User OpenID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `mention_doc` | `mention_doc` | @Document. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|           `token` | `string` | Docs token |
|           `obj_type` | `int` | Docs type **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|           `url` | `string` | Docs link (url_encode) |
|           `title` | `string` | The title of the referenced document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `reminder` | `reminder` | Date Reminder. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|           `create_user_id` | `string` | Creator user ID |
|           `is_notify` | `boolean` | Whether to notify |
|           `is_whole_day` | `boolean` | Whether it's a date or a whole hour |
|           `expire_time` | `string` | Time the event occurred (millisecond timestamp) |
|           `notify_time` | `string` | Time to trigger notification (millisecond timestamp) |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `file` | `inline_file` | Inline file. Only supports deleting or moving the element, and creating new inline files is not supported |
|           `file_token` | `string` | File token |
|           `source_block_id` | `string` | The ID of the block in which the file is located in the current document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `undefined` | `undefined_element` | Unsupported TextElements |
|         `inline_block` | `inline_block` | Inline block. Only supports deleting or moving the element, and creating new inline blocks is not supported |
|           `block_id` | `string` | Block ID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `equation` | `equation` | Equation |
|           `content` | `string` | Equation content that conforms to KaTeX syntax. For more information, see [KaTex syntax rules](https://katex.org/docs/supported.html]. |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|     `heading4` | `text` | Level 4 Heading Block |
|       `style` | `text_style` | Text style |
|         `align` | `int` | Alignment **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  |
|         `done` | `boolean` | Todo's status. Supports Todo Blocks |
|         `folded` | `boolean` | Folded state of text. Supports Heading1~9 block and the Text, Bullet, Ordered and Todo blocks with children |
|         `language` | `int` | Code block language. Only supports Code block **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|         `wrap` | `boolean` | Whether the code block wraps automatically. Only supports the Code block |
|         `background_color` | `string` | The background color of the block **Optional values are**:  - `LightGrayBackground`: Light gray - `LightRedBackground`: Light red - `LightOrangeBackground`: Light orange - `LightYellowBackground`: Light yellow - `LightGreenBackground`: Light green - `LightBlueBackground`: Light blue - `LightPurpleBackground`: Light purple - `PaleGrayBackground`: Medium gray - `DarkGrayBackground`: Gray - `DarkRedBackground`: Medium red - `DarkOrangeBackground`: Medium orange - `DarkYellowBackground`: Medium yellow - `DarkGreenBackground`: Medium green - `DarkBlueBackground`: Medium blue - `DarkPurpleBackground`: Medium purple  |
|         `indentation_level` | `string` | First line indentation level **Optional values are**:  - `NoIndent`: No indent - `OneLevelIndent`: One Level indent  |
|         `sequence` | `string` | Used to determine the number of an ordered list, either a number or 'auto' - Start a new list, numbering starts at 1, sequence='1' - Manually change the numbering, the numbering is set to a specific number, e.g. sequence='3' - Continue numbering, numbering is automatically consecutive, sequence='auto' - Some historical data and ordered block created through OpenAPI do not return this field |
|       `elements` | `text_element[]` | Text element |
|         `text_run` | `text_run` | Text. Supports Page, Text, Heading1~9, Bullet, Ordered, Code, Quote and Todo blocks |
|           `content` | `string` | Text content |
|           `text_element_style` | `text_element_style` | Text local style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `mention_user` | `mention_user` | @User. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|           `user_id` | `string` | User OpenID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `mention_doc` | `mention_doc` | @Document. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|           `token` | `string` | Docs token |
|           `obj_type` | `int` | Docs type **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|           `url` | `string` | Docs link (url_encode) |
|           `title` | `string` | The title of the referenced document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `reminder` | `reminder` | Date Reminder. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|           `create_user_id` | `string` | Creator user ID |
|           `is_notify` | `boolean` | Whether to notify |
|           `is_whole_day` | `boolean` | Whether it's a date or a whole hour |
|           `expire_time` | `string` | Time the event occurred (millisecond timestamp) |
|           `notify_time` | `string` | Time to trigger notification (millisecond timestamp) |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `file` | `inline_file` | Inline file. Only supports deleting or moving the element, and creating new inline files is not supported |
|           `file_token` | `string` | File token |
|           `source_block_id` | `string` | The ID of the block in which the file is located in the current document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `undefined` | `undefined_element` | Unsupported TextElements |
|         `inline_block` | `inline_block` | Inline block. Only supports deleting or moving the element, and creating new inline blocks is not supported |
|           `block_id` | `string` | Block ID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `equation` | `equation` | Equation |
|           `content` | `string` | Equation content that conforms to KaTeX syntax. For more information, see [KaTex syntax rules](https://katex.org/docs/supported.html]. |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|     `heading5` | `text` | Level 5 Heading Block |
|       `style` | `text_style` | Text style |
|         `align` | `int` | Alignment **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  |
|         `done` | `boolean` | Todo's status. Supports Todo Blocks |
|         `folded` | `boolean` | Folded state of text. Supports Heading1~9 block and the Text, Bullet, Ordered and Todo blocks with children |
|         `language` | `int` | Code block language. Only supports Code block **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|         `wrap` | `boolean` | Whether the code block wraps automatically. Only supports the Code block |
|         `background_color` | `string` | The background color of the block **Optional values are**:  - `LightGrayBackground`: Light gray - `LightRedBackground`: Light red - `LightOrangeBackground`: Light orange - `LightYellowBackground`: Light yellow - `LightGreenBackground`: Light green - `LightBlueBackground`: Light blue - `LightPurpleBackground`: Light purple - `PaleGrayBackground`: Medium gray - `DarkGrayBackground`: Gray - `DarkRedBackground`: Medium red - `DarkOrangeBackground`: Medium orange - `DarkYellowBackground`: Medium yellow - `DarkGreenBackground`: Medium green - `DarkBlueBackground`: Medium blue - `DarkPurpleBackground`: Medium purple  |
|         `indentation_level` | `string` | First line indentation level **Optional values are**:  - `NoIndent`: No indent - `OneLevelIndent`: One Level indent  |
|         `sequence` | `string` | Used to determine the number of an ordered list, either a number or 'auto' - Start a new list, numbering starts at 1, sequence='1' - Manually change the numbering, the numbering is set to a specific number, e.g. sequence='3' - Continue numbering, numbering is automatically consecutive, sequence='auto' - Some historical data and ordered block created through OpenAPI do not return this field |
|       `elements` | `text_element[]` | Text element |
|         `text_run` | `text_run` | Text. Supports Page, Text, Heading1~9, Bullet, Ordered, Code, Quote and Todo blocks |
|           `content` | `string` | Text content |
|           `text_element_style` | `text_element_style` | Text local style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `mention_user` | `mention_user` | @User. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|           `user_id` | `string` | User OpenID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `mention_doc` | `mention_doc` | @Document. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|           `token` | `string` | Docs token |
|           `obj_type` | `int` | Docs type **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|           `url` | `string` | Docs link (url_encode) |
|           `title` | `string` | The title of the referenced document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `reminder` | `reminder` | Date Reminder. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|           `create_user_id` | `string` | Creator user ID |
|           `is_notify` | `boolean` | Whether to notify |
|           `is_whole_day` | `boolean` | Whether it's a date or a whole hour |
|           `expire_time` | `string` | Time the event occurred (millisecond timestamp) |
|           `notify_time` | `string` | Time to trigger notification (millisecond timestamp) |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `file` | `inline_file` | Inline file. Only supports deleting or moving the element, and creating new inline files is not supported |
|           `file_token` | `string` | File token |
|           `source_block_id` | `string` | The ID of the block in which the file is located in the current document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `undefined` | `undefined_element` | Unsupported TextElements |
|         `inline_block` | `inline_block` | Inline block. Only supports deleting or moving the element, and creating new inline blocks is not supported |
|           `block_id` | `string` | Block ID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `equation` | `equation` | Equation |
|           `content` | `string` | Equation content that conforms to KaTeX syntax. For more information, see [KaTex syntax rules](https://katex.org/docs/supported.html]. |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|     `heading6` | `text` | Level 6 Heading Block |
|       `style` | `text_style` | Text style |
|         `align` | `int` | Alignment **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  |
|         `done` | `boolean` | Todo's status. Supports Todo Blocks |
|         `folded` | `boolean` | Folded state of text. Supports Heading1~9 block and the Text, Bullet, Ordered and Todo blocks with children |
|         `language` | `int` | Code block language. Only supports Code block **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|         `wrap` | `boolean` | Whether the code block wraps automatically. Only supports the Code block |
|         `background_color` | `string` | The background color of the block **Optional values are**:  - `LightGrayBackground`: Light gray - `LightRedBackground`: Light red - `LightOrangeBackground`: Light orange - `LightYellowBackground`: Light yellow - `LightGreenBackground`: Light green - `LightBlueBackground`: Light blue - `LightPurpleBackground`: Light purple - `PaleGrayBackground`: Medium gray - `DarkGrayBackground`: Gray - `DarkRedBackground`: Medium red - `DarkOrangeBackground`: Medium orange - `DarkYellowBackground`: Medium yellow - `DarkGreenBackground`: Medium green - `DarkBlueBackground`: Medium blue - `DarkPurpleBackground`: Medium purple  |
|         `indentation_level` | `string` | First line indentation level **Optional values are**:  - `NoIndent`: No indent - `OneLevelIndent`: One Level indent  |
|         `sequence` | `string` | Used to determine the number of an ordered list, either a number or 'auto' - Start a new list, numbering starts at 1, sequence='1' - Manually change the numbering, the numbering is set to a specific number, e.g. sequence='3' - Continue numbering, numbering is automatically consecutive, sequence='auto' - Some historical data and ordered block created through OpenAPI do not return this field |
|       `elements` | `text_element[]` | Text element |
|         `text_run` | `text_run` | Text. Supports Page, Text, Heading1~9, Bullet, Ordered, Code, Quote and Todo blocks |
|           `content` | `string` | Text content |
|           `text_element_style` | `text_element_style` | Text local style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `mention_user` | `mention_user` | @User. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|           `user_id` | `string` | User OpenID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `mention_doc` | `mention_doc` | @Document. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|           `token` | `string` | Docs token |
|           `obj_type` | `int` | Docs type **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|           `url` | `string` | Docs link (url_encode) |
|           `title` | `string` | The title of the referenced document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `reminder` | `reminder` | Date Reminder. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|           `create_user_id` | `string` | Creator user ID |
|           `is_notify` | `boolean` | Whether to notify |
|           `is_whole_day` | `boolean` | Whether it's a date or a whole hour |
|           `expire_time` | `string` | Time the event occurred (millisecond timestamp) |
|           `notify_time` | `string` | Time to trigger notification (millisecond timestamp) |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `file` | `inline_file` | Inline file. Only supports deleting or moving the element, and creating new inline files is not supported |
|           `file_token` | `string` | File token |
|           `source_block_id` | `string` | The ID of the block in which the file is located in the current document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `undefined` | `undefined_element` | Unsupported TextElements |
|         `inline_block` | `inline_block` | Inline block. Only supports deleting or moving the element, and creating new inline blocks is not supported |
|           `block_id` | `string` | Block ID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `equation` | `equation` | Equation |
|           `content` | `string` | Equation content that conforms to KaTeX syntax. For more information, see [KaTex syntax rules](https://katex.org/docs/supported.html]. |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|     `heading7` | `text` | Level 7 Heading Block |
|       `style` | `text_style` | Text style |
|         `align` | `int` | Alignment **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  |
|         `done` | `boolean` | Todo's status. Supports Todo Blocks |
|         `folded` | `boolean` | Folded state of text. Supports Heading1~9 block and the Text, Bullet, Ordered and Todo blocks with children |
|         `language` | `int` | Code block language. Only supports Code block **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|         `wrap` | `boolean` | Whether the code block wraps automatically. Only supports the Code block |
|         `background_color` | `string` | The background color of the block **Optional values are**:  - `LightGrayBackground`: Light gray - `LightRedBackground`: Light red - `LightOrangeBackground`: Light orange - `LightYellowBackground`: Light yellow - `LightGreenBackground`: Light green - `LightBlueBackground`: Light blue - `LightPurpleBackground`: Light purple - `PaleGrayBackground`: Medium gray - `DarkGrayBackground`: Gray - `DarkRedBackground`: Medium red - `DarkOrangeBackground`: Medium orange - `DarkYellowBackground`: Medium yellow - `DarkGreenBackground`: Medium green - `DarkBlueBackground`: Medium blue - `DarkPurpleBackground`: Medium purple  |
|         `indentation_level` | `string` | First line indentation level **Optional values are**:  - `NoIndent`: No indent - `OneLevelIndent`: One Level indent  |
|         `sequence` | `string` | Used to determine the number of an ordered list, either a number or 'auto' - Start a new list, numbering starts at 1, sequence='1' - Manually change the numbering, the numbering is set to a specific number, e.g. sequence='3' - Continue numbering, numbering is automatically consecutive, sequence='auto' - Some historical data and ordered block created through OpenAPI do not return this field |
|       `elements` | `text_element[]` | Text element |
|         `text_run` | `text_run` | Text. Supports Page, Text, Heading1~9, Bullet, Ordered, Code, Quote and Todo blocks |
|           `content` | `string` | Text content |
|           `text_element_style` | `text_element_style` | Text local style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `mention_user` | `mention_user` | @User. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|           `user_id` | `string` | User OpenID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `mention_doc` | `mention_doc` | @Document. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|           `token` | `string` | Docs token |
|           `obj_type` | `int` | Docs type **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|           `url` | `string` | Docs link (url_encode) |
|           `title` | `string` | The title of the referenced document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `reminder` | `reminder` | Date Reminder. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|           `create_user_id` | `string` | Creator user ID |
|           `is_notify` | `boolean` | Whether to notify |
|           `is_whole_day` | `boolean` | Whether it's a date or a whole hour |
|           `expire_time` | `string` | Time the event occurred (millisecond timestamp) |
|           `notify_time` | `string` | Time to trigger notification (millisecond timestamp) |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `file` | `inline_file` | Inline file. Only supports deleting or moving the element, and creating new inline files is not supported |
|           `file_token` | `string` | File token |
|           `source_block_id` | `string` | The ID of the block in which the file is located in the current document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `undefined` | `undefined_element` | Unsupported TextElements |
|         `inline_block` | `inline_block` | Inline block. Only supports deleting or moving the element, and creating new inline blocks is not supported |
|           `block_id` | `string` | Block ID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `equation` | `equation` | Equation |
|           `content` | `string` | Equation content that conforms to KaTeX syntax. For more information, see [KaTex syntax rules](https://katex.org/docs/supported.html]. |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|     `heading8` | `text` | Level 8 Heading Block |
|       `style` | `text_style` | Text style |
|         `align` | `int` | Alignment **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  |
|         `done` | `boolean` | Todo's status. Supports Todo Blocks |
|         `folded` | `boolean` | Folded state of text. Supports Heading1~9 block and the Text, Bullet, Ordered and Todo blocks with children |
|         `language` | `int` | Code block language. Only supports Code block **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|         `wrap` | `boolean` | Whether the code block wraps automatically. Only supports the Code block |
|         `background_color` | `string` | The background color of the block **Optional values are**:  - `LightGrayBackground`: Light gray - `LightRedBackground`: Light red - `LightOrangeBackground`: Light orange - `LightYellowBackground`: Light yellow - `LightGreenBackground`: Light green - `LightBlueBackground`: Light blue - `LightPurpleBackground`: Light purple - `PaleGrayBackground`: Medium gray - `DarkGrayBackground`: Gray - `DarkRedBackground`: Medium red - `DarkOrangeBackground`: Medium orange - `DarkYellowBackground`: Medium yellow - `DarkGreenBackground`: Medium green - `DarkBlueBackground`: Medium blue - `DarkPurpleBackground`: Medium purple  |
|         `indentation_level` | `string` | First line indentation level **Optional values are**:  - `NoIndent`: No indent - `OneLevelIndent`: One Level indent  |
|         `sequence` | `string` | Used to determine the number of an ordered list, either a number or 'auto' - Start a new list, numbering starts at 1, sequence='1' - Manually change the numbering, the numbering is set to a specific number, e.g. sequence='3' - Continue numbering, numbering is automatically consecutive, sequence='auto' - Some historical data and ordered block created through OpenAPI do not return this field |
|       `elements` | `text_element[]` | Text element |
|         `text_run` | `text_run` | Text. Supports Page, Text, Heading1~9, Bullet, Ordered, Code, Quote and Todo blocks |
|           `content` | `string` | Text content |
|           `text_element_style` | `text_element_style` | Text local style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `mention_user` | `mention_user` | @User. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|           `user_id` | `string` | User OpenID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `mention_doc` | `mention_doc` | @Document. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|           `token` | `string` | Docs token |
|           `obj_type` | `int` | Docs type **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|           `url` | `string` | Docs link (url_encode) |
|           `title` | `string` | The title of the referenced document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `reminder` | `reminder` | Date Reminder. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|           `create_user_id` | `string` | Creator user ID |
|           `is_notify` | `boolean` | Whether to notify |
|           `is_whole_day` | `boolean` | Whether it's a date or a whole hour |
|           `expire_time` | `string` | Time the event occurred (millisecond timestamp) |
|           `notify_time` | `string` | Time to trigger notification (millisecond timestamp) |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `file` | `inline_file` | Inline file. Only supports deleting or moving the element, and creating new inline files is not supported |
|           `file_token` | `string` | File token |
|           `source_block_id` | `string` | The ID of the block in which the file is located in the current document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `undefined` | `undefined_element` | Unsupported TextElements |
|         `inline_block` | `inline_block` | Inline block. Only supports deleting or moving the element, and creating new inline blocks is not supported |
|           `block_id` | `string` | Block ID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `equation` | `equation` | Equation |
|           `content` | `string` | Equation content that conforms to KaTeX syntax. For more information, see [KaTex syntax rules](https://katex.org/docs/supported.html]. |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|     `heading9` | `text` | Level 9 Heading Block |
|       `style` | `text_style` | Text style |
|         `align` | `int` | Alignment **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  |
|         `done` | `boolean` | Todo's status. Supports Todo Blocks |
|         `folded` | `boolean` | Folded state of text. Supports Heading1~9 block and the Text, Bullet, Ordered and Todo blocks with children |
|         `language` | `int` | Code block language. Only supports Code block **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|         `wrap` | `boolean` | Whether the code block wraps automatically. Only supports the Code block |
|         `background_color` | `string` | The background color of the block **Optional values are**:  - `LightGrayBackground`: Light gray - `LightRedBackground`: Light red - `LightOrangeBackground`: Light orange - `LightYellowBackground`: Light yellow - `LightGreenBackground`: Light green - `LightBlueBackground`: Light blue - `LightPurpleBackground`: Light purple - `PaleGrayBackground`: Medium gray - `DarkGrayBackground`: Gray - `DarkRedBackground`: Medium red - `DarkOrangeBackground`: Medium orange - `DarkYellowBackground`: Medium yellow - `DarkGreenBackground`: Medium green - `DarkBlueBackground`: Medium blue - `DarkPurpleBackground`: Medium purple  |
|         `indentation_level` | `string` | First line indentation level **Optional values are**:  - `NoIndent`: No indent - `OneLevelIndent`: One Level indent  |
|         `sequence` | `string` | Used to determine the number of an ordered list, either a number or 'auto' - Start a new list, numbering starts at 1, sequence='1' - Manually change the numbering, the numbering is set to a specific number, e.g. sequence='3' - Continue numbering, numbering is automatically consecutive, sequence='auto' - Some historical data and ordered block created through OpenAPI do not return this field |
|       `elements` | `text_element[]` | Text element |
|         `text_run` | `text_run` | Text. Supports Page, Text, Heading1~9, Bullet, Ordered, Code, Quote and Todo blocks |
|           `content` | `string` | Text content |
|           `text_element_style` | `text_element_style` | Text local style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `mention_user` | `mention_user` | @User. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|           `user_id` | `string` | User OpenID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `mention_doc` | `mention_doc` | @Document. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|           `token` | `string` | Docs token |
|           `obj_type` | `int` | Docs type **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|           `url` | `string` | Docs link (url_encode) |
|           `title` | `string` | The title of the referenced document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `reminder` | `reminder` | Date Reminder. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|           `create_user_id` | `string` | Creator user ID |
|           `is_notify` | `boolean` | Whether to notify |
|           `is_whole_day` | `boolean` | Whether it's a date or a whole hour |
|           `expire_time` | `string` | Time the event occurred (millisecond timestamp) |
|           `notify_time` | `string` | Time to trigger notification (millisecond timestamp) |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `file` | `inline_file` | Inline file. Only supports deleting or moving the element, and creating new inline files is not supported |
|           `file_token` | `string` | File token |
|           `source_block_id` | `string` | The ID of the block in which the file is located in the current document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `undefined` | `undefined_element` | Unsupported TextElements |
|         `inline_block` | `inline_block` | Inline block. Only supports deleting or moving the element, and creating new inline blocks is not supported |
|           `block_id` | `string` | Block ID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `equation` | `equation` | Equation |
|           `content` | `string` | Equation content that conforms to KaTeX syntax. For more information, see [KaTex syntax rules](https://katex.org/docs/supported.html]. |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|     `bullet` | `text` | Bullet List Block |
|       `style` | `text_style` | Text style |
|         `align` | `int` | Alignment **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  |
|         `done` | `boolean` | Todo's status. Supports Todo Blocks |
|         `folded` | `boolean` | Folded state of text. Supports Heading1~9 block and the Text, Bullet, Ordered and Todo blocks with children |
|         `language` | `int` | Code block language. Only supports Code block **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|         `wrap` | `boolean` | Whether the code block wraps automatically. Only supports the Code block |
|         `background_color` | `string` | The background color of the block **Optional values are**:  - `LightGrayBackground`: Light gray - `LightRedBackground`: Light red - `LightOrangeBackground`: Light orange - `LightYellowBackground`: Light yellow - `LightGreenBackground`: Light green - `LightBlueBackground`: Light blue - `LightPurpleBackground`: Light purple - `PaleGrayBackground`: Medium gray - `DarkGrayBackground`: Gray - `DarkRedBackground`: Medium red - `DarkOrangeBackground`: Medium orange - `DarkYellowBackground`: Medium yellow - `DarkGreenBackground`: Medium green - `DarkBlueBackground`: Medium blue - `DarkPurpleBackground`: Medium purple  |
|         `indentation_level` | `string` | First line indentation level **Optional values are**:  - `NoIndent`: No indent - `OneLevelIndent`: One Level indent  |
|         `sequence` | `string` | Used to determine the number of an ordered list, either a number or 'auto' - Start a new list, numbering starts at 1, sequence='1' - Manually change the numbering, the numbering is set to a specific number, e.g. sequence='3' - Continue numbering, numbering is automatically consecutive, sequence='auto' - Some historical data and ordered block created through OpenAPI do not return this field |
|       `elements` | `text_element[]` | Text element |
|         `text_run` | `text_run` | Text. Supports Page, Text, Heading1~9, Bullet, Ordered, Code, Quote and Todo blocks |
|           `content` | `string` | Text content |
|           `text_element_style` | `text_element_style` | Text local style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `mention_user` | `mention_user` | @User. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|           `user_id` | `string` | User OpenID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `mention_doc` | `mention_doc` | @Document. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|           `token` | `string` | Docs token |
|           `obj_type` | `int` | Docs type **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|           `url` | `string` | Docs link (url_encode) |
|           `title` | `string` | The title of the referenced document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `reminder` | `reminder` | Date Reminder. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|           `create_user_id` | `string` | Creator user ID |
|           `is_notify` | `boolean` | Whether to notify |
|           `is_whole_day` | `boolean` | Whether it's a date or a whole hour |
|           `expire_time` | `string` | Time the event occurred (millisecond timestamp) |
|           `notify_time` | `string` | Time to trigger notification (millisecond timestamp) |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `file` | `inline_file` | Inline file. Only supports deleting or moving the element, and creating new inline files is not supported |
|           `file_token` | `string` | File token |
|           `source_block_id` | `string` | The ID of the block in which the file is located in the current document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `undefined` | `undefined_element` | Unsupported TextElements |
|         `inline_block` | `inline_block` | Inline block. Only supports deleting or moving the element, and creating new inline blocks is not supported |
|           `block_id` | `string` | Block ID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `equation` | `equation` | Equation |
|           `content` | `string` | Equation content that conforms to KaTeX syntax. For more information, see [KaTex syntax rules](https://katex.org/docs/supported.html]. |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|     `ordered` | `text` | Ordered List Block |
|       `style` | `text_style` | Text style |
|         `align` | `int` | Alignment **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  |
|         `done` | `boolean` | Todo's status. Supports Todo Blocks |
|         `folded` | `boolean` | Folded state of text. Supports Heading1~9 block and the Text, Bullet, Ordered and Todo blocks with children |
|         `language` | `int` | Code block language. Only supports Code block **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|         `wrap` | `boolean` | Whether the code block wraps automatically. Only supports the Code block |
|         `background_color` | `string` | The background color of the block **Optional values are**:  - `LightGrayBackground`: Light gray - `LightRedBackground`: Light red - `LightOrangeBackground`: Light orange - `LightYellowBackground`: Light yellow - `LightGreenBackground`: Light green - `LightBlueBackground`: Light blue - `LightPurpleBackground`: Light purple - `PaleGrayBackground`: Medium gray - `DarkGrayBackground`: Gray - `DarkRedBackground`: Medium red - `DarkOrangeBackground`: Medium orange - `DarkYellowBackground`: Medium yellow - `DarkGreenBackground`: Medium green - `DarkBlueBackground`: Medium blue - `DarkPurpleBackground`: Medium purple  |
|         `indentation_level` | `string` | First line indentation level **Optional values are**:  - `NoIndent`: No indent - `OneLevelIndent`: One Level indent  |
|         `sequence` | `string` | Used to determine the number of an ordered list, either a number or 'auto' - Start a new list, numbering starts at 1, sequence='1' - Manually change the numbering, the numbering is set to a specific number, e.g. sequence='3' - Continue numbering, numbering is automatically consecutive, sequence='auto' - Some historical data and ordered block created through OpenAPI do not return this field |
|       `elements` | `text_element[]` | Text element |
|         `text_run` | `text_run` | Text. Supports Page, Text, Heading1~9, Bullet, Ordered, Code, Quote and Todo blocks |
|           `content` | `string` | Text content |
|           `text_element_style` | `text_element_style` | Text local style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `mention_user` | `mention_user` | @User. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|           `user_id` | `string` | User OpenID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `mention_doc` | `mention_doc` | @Document. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|           `token` | `string` | Docs token |
|           `obj_type` | `int` | Docs type **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|           `url` | `string` | Docs link (url_encode) |
|           `title` | `string` | The title of the referenced document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `reminder` | `reminder` | Date Reminder. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|           `create_user_id` | `string` | Creator user ID |
|           `is_notify` | `boolean` | Whether to notify |
|           `is_whole_day` | `boolean` | Whether it's a date or a whole hour |
|           `expire_time` | `string` | Time the event occurred (millisecond timestamp) |
|           `notify_time` | `string` | Time to trigger notification (millisecond timestamp) |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `file` | `inline_file` | Inline file. Only supports deleting or moving the element, and creating new inline files is not supported |
|           `file_token` | `string` | File token |
|           `source_block_id` | `string` | The ID of the block in which the file is located in the current document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `undefined` | `undefined_element` | Unsupported TextElements |
|         `inline_block` | `inline_block` | Inline block. Only supports deleting or moving the element, and creating new inline blocks is not supported |
|           `block_id` | `string` | Block ID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `equation` | `equation` | Equation |
|           `content` | `string` | Equation content that conforms to KaTeX syntax. For more information, see [KaTex syntax rules](https://katex.org/docs/supported.html]. |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|     `code` | `text` | Code Block |
|       `style` | `text_style` | Text style |
|         `align` | `int` | Alignment **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  |
|         `done` | `boolean` | Todo's status. Supports Todo Blocks |
|         `folded` | `boolean` | Folded state of text. Supports Heading1~9 block and the Text, Bullet, Ordered and Todo blocks with children |
|         `language` | `int` | Code block language. Only supports Code block **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|         `wrap` | `boolean` | Whether the code block wraps automatically. Only supports the Code block |
|         `background_color` | `string` | The background color of the block **Optional values are**:  - `LightGrayBackground`: Light gray - `LightRedBackground`: Light red - `LightOrangeBackground`: Light orange - `LightYellowBackground`: Light yellow - `LightGreenBackground`: Light green - `LightBlueBackground`: Light blue - `LightPurpleBackground`: Light purple - `PaleGrayBackground`: Medium gray - `DarkGrayBackground`: Gray - `DarkRedBackground`: Medium red - `DarkOrangeBackground`: Medium orange - `DarkYellowBackground`: Medium yellow - `DarkGreenBackground`: Medium green - `DarkBlueBackground`: Medium blue - `DarkPurpleBackground`: Medium purple  |
|         `indentation_level` | `string` | First line indentation level **Optional values are**:  - `NoIndent`: No indent - `OneLevelIndent`: One Level indent  |
|         `sequence` | `string` | Used to determine the number of an ordered list, either a number or 'auto' - Start a new list, numbering starts at 1, sequence='1' - Manually change the numbering, the numbering is set to a specific number, e.g. sequence='3' - Continue numbering, numbering is automatically consecutive, sequence='auto' - Some historical data and ordered block created through OpenAPI do not return this field |
|       `elements` | `text_element[]` | Text element |
|         `text_run` | `text_run` | Text. Supports Page, Text, Heading1~9, Bullet, Ordered, Code, Quote and Todo blocks |
|           `content` | `string` | Text content |
|           `text_element_style` | `text_element_style` | Text local style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `mention_user` | `mention_user` | @User. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|           `user_id` | `string` | User OpenID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `mention_doc` | `mention_doc` | @Document. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|           `token` | `string` | Docs token |
|           `obj_type` | `int` | Docs type **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|           `url` | `string` | Docs link (url_encode) |
|           `title` | `string` | The title of the referenced document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `reminder` | `reminder` | Date Reminder. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|           `create_user_id` | `string` | Creator user ID |
|           `is_notify` | `boolean` | Whether to notify |
|           `is_whole_day` | `boolean` | Whether it's a date or a whole hour |
|           `expire_time` | `string` | Time the event occurred (millisecond timestamp) |
|           `notify_time` | `string` | Time to trigger notification (millisecond timestamp) |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `file` | `inline_file` | Inline file. Only supports deleting or moving the element, and creating new inline files is not supported |
|           `file_token` | `string` | File token |
|           `source_block_id` | `string` | The ID of the block in which the file is located in the current document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `undefined` | `undefined_element` | Unsupported TextElements |
|         `inline_block` | `inline_block` | Inline block. Only supports deleting or moving the element, and creating new inline blocks is not supported |
|           `block_id` | `string` | Block ID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `equation` | `equation` | Equation |
|           `content` | `string` | Equation content that conforms to KaTeX syntax. For more information, see [KaTex syntax rules](https://katex.org/docs/supported.html]. |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|     `quote` | `text` | Quote Block |
|       `style` | `text_style` | Text style |
|         `align` | `int` | Alignment **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  |
|         `done` | `boolean` | Todo's status. Supports Todo Blocks |
|         `folded` | `boolean` | Folded state of text. Supports Heading1~9 block and the Text, Bullet, Ordered and Todo blocks with children |
|         `language` | `int` | Code block language. Only supports Code block **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|         `wrap` | `boolean` | Whether the code block wraps automatically. Only supports the Code block |
|         `background_color` | `string` | The background color of the block **Optional values are**:  - `LightGrayBackground`: Light gray - `LightRedBackground`: Light red - `LightOrangeBackground`: Light orange - `LightYellowBackground`: Light yellow - `LightGreenBackground`: Light green - `LightBlueBackground`: Light blue - `LightPurpleBackground`: Light purple - `PaleGrayBackground`: Medium gray - `DarkGrayBackground`: Gray - `DarkRedBackground`: Medium red - `DarkOrangeBackground`: Medium orange - `DarkYellowBackground`: Medium yellow - `DarkGreenBackground`: Medium green - `DarkBlueBackground`: Medium blue - `DarkPurpleBackground`: Medium purple  |
|         `indentation_level` | `string` | First line indentation level **Optional values are**:  - `NoIndent`: No indent - `OneLevelIndent`: One Level indent  |
|         `sequence` | `string` | Used to determine the number of an ordered list, either a number or 'auto' - Start a new list, numbering starts at 1, sequence='1' - Manually change the numbering, the numbering is set to a specific number, e.g. sequence='3' - Continue numbering, numbering is automatically consecutive, sequence='auto' - Some historical data and ordered block created through OpenAPI do not return this field |
|       `elements` | `text_element[]` | Text element |
|         `text_run` | `text_run` | Text. Supports Page, Text, Heading1~9, Bullet, Ordered, Code, Quote and Todo blocks |
|           `content` | `string` | Text content |
|           `text_element_style` | `text_element_style` | Text local style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `mention_user` | `mention_user` | @User. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|           `user_id` | `string` | User OpenID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `mention_doc` | `mention_doc` | @Document. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|           `token` | `string` | Docs token |
|           `obj_type` | `int` | Docs type **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|           `url` | `string` | Docs link (url_encode) |
|           `title` | `string` | The title of the referenced document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `reminder` | `reminder` | Date Reminder. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|           `create_user_id` | `string` | Creator user ID |
|           `is_notify` | `boolean` | Whether to notify |
|           `is_whole_day` | `boolean` | Whether it's a date or a whole hour |
|           `expire_time` | `string` | Time the event occurred (millisecond timestamp) |
|           `notify_time` | `string` | Time to trigger notification (millisecond timestamp) |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `file` | `inline_file` | Inline file. Only supports deleting or moving the element, and creating new inline files is not supported |
|           `file_token` | `string` | File token |
|           `source_block_id` | `string` | The ID of the block in which the file is located in the current document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `undefined` | `undefined_element` | Unsupported TextElements |
|         `inline_block` | `inline_block` | Inline block. Only supports deleting or moving the element, and creating new inline blocks is not supported |
|           `block_id` | `string` | Block ID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `equation` | `equation` | Equation |
|           `content` | `string` | Equation content that conforms to KaTeX syntax. For more information, see [KaTex syntax rules](https://katex.org/docs/supported.html]. |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|     `equation` | `text` | Equation Block |
|       `style` | `text_style` | Text style |
|         `align` | `int` | Alignment **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  |
|         `done` | `boolean` | Todo's status. Supports Todo Blocks |
|         `folded` | `boolean` | Folded state of text. Supports Heading1~9 block and the Text, Bullet, Ordered and Todo blocks with children |
|         `language` | `int` | Code block language. Only supports Code block **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|         `wrap` | `boolean` | Whether the code block wraps automatically. Only supports the Code block |
|         `background_color` | `string` | The background color of the block **Optional values are**:  - `LightGrayBackground`: Light gray - `LightRedBackground`: Light red - `LightOrangeBackground`: Light orange - `LightYellowBackground`: Light yellow - `LightGreenBackground`: Light green - `LightBlueBackground`: Light blue - `LightPurpleBackground`: Light purple - `PaleGrayBackground`: Medium gray - `DarkGrayBackground`: Gray - `DarkRedBackground`: Medium red - `DarkOrangeBackground`: Medium orange - `DarkYellowBackground`: Medium yellow - `DarkGreenBackground`: Medium green - `DarkBlueBackground`: Medium blue - `DarkPurpleBackground`: Medium purple  |
|         `indentation_level` | `string` | First line indentation level **Optional values are**:  - `NoIndent`: No indent - `OneLevelIndent`: One Level indent  |
|         `sequence` | `string` | Used to determine the number of an ordered list, either a number or 'auto' - Start a new list, numbering starts at 1, sequence='1' - Manually change the numbering, the numbering is set to a specific number, e.g. sequence='3' - Continue numbering, numbering is automatically consecutive, sequence='auto' - Some historical data and ordered block created through OpenAPI do not return this field |
|       `elements` | `text_element[]` | Text element |
|         `text_run` | `text_run` | Text. Supports Page, Text, Heading1~9, Bullet, Ordered, Code, Quote and Todo blocks |
|           `content` | `string` | Text content |
|           `text_element_style` | `text_element_style` | Text local style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `mention_user` | `mention_user` | @User. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|           `user_id` | `string` | User OpenID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `mention_doc` | `mention_doc` | @Document. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|           `token` | `string` | Docs token |
|           `obj_type` | `int` | Docs type **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|           `url` | `string` | Docs link (url_encode) |
|           `title` | `string` | The title of the referenced document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `reminder` | `reminder` | Date Reminder. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|           `create_user_id` | `string` | Creator user ID |
|           `is_notify` | `boolean` | Whether to notify |
|           `is_whole_day` | `boolean` | Whether it's a date or a whole hour |
|           `expire_time` | `string` | Time the event occurred (millisecond timestamp) |
|           `notify_time` | `string` | Time to trigger notification (millisecond timestamp) |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `file` | `inline_file` | Inline file. Only supports deleting or moving the element, and creating new inline files is not supported |
|           `file_token` | `string` | File token |
|           `source_block_id` | `string` | The ID of the block in which the file is located in the current document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `undefined` | `undefined_element` | Unsupported TextElements |
|         `inline_block` | `inline_block` | Inline block. Only supports deleting or moving the element, and creating new inline blocks is not supported |
|           `block_id` | `string` | Block ID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `equation` | `equation` | Equation |
|           `content` | `string` | Equation content that conforms to KaTeX syntax. For more information, see [KaTex syntax rules](https://katex.org/docs/supported.html]. |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|     `todo` | `text` | Todo Block |
|       `style` | `text_style` | Text style |
|         `align` | `int` | Alignment **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  |
|         `done` | `boolean` | Todo's status. Supports Todo Blocks |
|         `folded` | `boolean` | Folded state of text. Supports Heading1~9 block and the Text, Bullet, Ordered and Todo blocks with children |
|         `language` | `int` | Code block language. Only supports Code block **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|         `wrap` | `boolean` | Whether the code block wraps automatically. Only supports the Code block |
|         `background_color` | `string` | The background color of the block **Optional values are**:  - `LightGrayBackground`: Light gray - `LightRedBackground`: Light red - `LightOrangeBackground`: Light orange - `LightYellowBackground`: Light yellow - `LightGreenBackground`: Light green - `LightBlueBackground`: Light blue - `LightPurpleBackground`: Light purple - `PaleGrayBackground`: Medium gray - `DarkGrayBackground`: Gray - `DarkRedBackground`: Medium red - `DarkOrangeBackground`: Medium orange - `DarkYellowBackground`: Medium yellow - `DarkGreenBackground`: Medium green - `DarkBlueBackground`: Medium blue - `DarkPurpleBackground`: Medium purple  |
|         `indentation_level` | `string` | First line indentation level **Optional values are**:  - `NoIndent`: No indent - `OneLevelIndent`: One Level indent  |
|         `sequence` | `string` | Used to determine the number of an ordered list, either a number or 'auto' - Start a new list, numbering starts at 1, sequence='1' - Manually change the numbering, the numbering is set to a specific number, e.g. sequence='3' - Continue numbering, numbering is automatically consecutive, sequence='auto' - Some historical data and ordered block created through OpenAPI do not return this field |
|       `elements` | `text_element[]` | Text element |
|         `text_run` | `text_run` | Text. Supports Page, Text, Heading1~9, Bullet, Ordered, Code, Quote and Todo blocks |
|           `content` | `string` | Text content |
|           `text_element_style` | `text_element_style` | Text local style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `mention_user` | `mention_user` | @User. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|           `user_id` | `string` | User OpenID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `mention_doc` | `mention_doc` | @Document. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|           `token` | `string` | Docs token |
|           `obj_type` | `int` | Docs type **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|           `url` | `string` | Docs link (url_encode) |
|           `title` | `string` | The title of the referenced document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `reminder` | `reminder` | Date Reminder. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|           `create_user_id` | `string` | Creator user ID |
|           `is_notify` | `boolean` | Whether to notify |
|           `is_whole_day` | `boolean` | Whether it's a date or a whole hour |
|           `expire_time` | `string` | Time the event occurred (millisecond timestamp) |
|           `notify_time` | `string` | Time to trigger notification (millisecond timestamp) |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `file` | `inline_file` | Inline file. Only supports deleting or moving the element, and creating new inline files is not supported |
|           `file_token` | `string` | File token |
|           `source_block_id` | `string` | The ID of the block in which the file is located in the current document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `undefined` | `undefined_element` | Unsupported TextElements |
|         `inline_block` | `inline_block` | Inline block. Only supports deleting or moving the element, and creating new inline blocks is not supported |
|           `block_id` | `string` | Block ID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `equation` | `equation` | Equation |
|           `content` | `string` | Equation content that conforms to KaTeX syntax. For more information, see [KaTex syntax rules](https://katex.org/docs/supported.html]. |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|     `bitable` | `bitable` | Bitable Block |
|       `token` | `string` | Bitable Block |
|       `view_type` | `int` | View type **Optional values are**:  - `1`: Data sheet - `2`: Kanban  |
|     `callout` | `callout` | Callout Block |
|       `background_color` | `int` | Callout block background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium gray - `8`: Medium Red - `9`: Medium orange - `10`: Medium yellow - `11`: Medium green - `12`: Medium blue - `13`: Medium purple - `14`: Gray - `15`: Light Gray  |
|       `border_color` | `int` | Border color **Optional values are**:  - `1`: Red - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Gray  |
|       `text_color` | `int` | Text color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|       `emoji_id` | `string` | Callout Block icon |
|     `chat_card` | `chat_card` | Chat Card Block |
|       `chat_id` | `string` | Chat ID |
|       `align` | `int` | Alignment **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  |
|     `diagram` | `diagram` | UML Diagram Block |
|       `diagram_type` | `int` | Diagram type **Optional values are**:  - `1`: Flowchart - `2`: UML diagram  |
|     `divider` | `divider` | Divider Block |
|     `file` | `file` | File Block |
|       `token` | `string` | Attachment Token |
|       `name` | `string` | File name |
|       `view_type` | `int` | View type **Optional values are**:  - `1`: Card view - `2`: Preview view  |
|     `grid` | `grid` | Grid Block |
|       `column_size` | `int` | Number of columns |
|     `grid_column` | `grid_column` | Grid Column Block |
|       `width_ratio` | `int` | The proportion of the current column to the total column |
|     `iframe` | `iframe` | Iframe Block |
|       `component` | `iframe_component` | Component elements of iframe |
|         `iframe_type` | `int` | Iframe Type **Optional values are**:  - `1`: Bilibili - `2`: Xigua Video - `3`: Youku - `4`: Airtable - `5`: Baidu Map - `6`: Autonavi Map - `7`: TikTok - `8`: Figma - `9`: Ink knife - `10`: Canva - `11`: CodePen - `12`: Lark Survey - `13`: Gold data - `14`: GoogleMap - `15`: Youtube - `99`: Other  |
|         `url` | `string` | Iframe target url (requires url_encode) |
|     `image` | `image` | Image Block |
|       `width` | `int` | Width unit px |
|       `height` | `int` | Height |
|       `token` | `string` | Image token |
|       `align` | `int` | alignment **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  |
|     `isv` | `isv` | ISV Block |
|       `component_id` | `string` | Team Interaction App Unique ID |
|       `component_type_id` | `string` | Team interaction application types, such as information collection "blk_5f992038c64240015d280958" |
|     `add_ons` | `add_ons` | Add-ons |
|       `component_id` | `string` | Document widget ID |
|       `component_type_id` | `string` | Document widget types, such as Q & A interactive "blk_636a0a6657db8001c8df5488" |
|       `record` | `string` | Document widget content data, JSON string |
|     `mindnote` | `mindnote` | Mindnote Block |
|       `token` | `string` | Mindnote token |
|     `sheet` | `sheet` | Sheet Block |
|       `token` | `string` | Sheet Block token |
|       `row_size` | `int` | Number of spreadsheet rows |
|       `column_size` | `int` | Number of spreadsheet columns |
|     `table` | `table` | Table Block |
|       `cells` | `string[]` | Array of cells, the array element is the ID of the Table Cell Block |
|       `property` | `table_property` | Table Properties |
|         `row_size` | `int` | Number of rows |
|         `column_size` | `int` | Number of columns |
|         `column_width` | `int[]` | Column width in px |
|         `merge_info` | `table_merge_info[]` | Cell merge information. This property is read-only when creating a table and is generated by the backend. If you need to merge cells, you can do so by using the subrequest merge_table_cells of the Update Block api. |
|           `row_span` | `int` | Number of consecutive rows merged from the current line |
|           `col_span` | `int` | Number of consecutive columns that are merged from the current pristine |
|         `header_row` | `boolean` | Set the first row as the header row. |
|         `header_column` | `boolean` | Set the first column as the header column. |
|     `table_cell` | `table_cell` | Table Cell Block |
|     `view` | `view` | View Block |
|       `view_type` | `int` | View type **Optional values are**:  - `1`: Card view - `2`: Preview view - `3`: inline  |
|     `undefined` | `undefined` | Blocks not supported |
|     `quote_container` | `quote_container` | Quote Container Block |
|     `task` | `task` | Task Block |
|       `task_id` | `string` | Task ID. For query task details, see: Get Task Details |
|       `folded` | `boolean` | Folded state |
|     `okr` | `okr` | OKR blocks. It can only be created when using Authorization with `user_access_token` |
|       `okr_id` | `string` | OKR ID. Get the OKR ID that needs to be inserted and see Get the user's OKR list |
|       `objectives` | `objective_id_with_kr_id[]` | The objective ID and key result ID in the OKR Block, insert all objective and key result under okr when objectives is null |
|         `objective_id` | `string` | Objective ID in OKR |
|         `kr_ids` | `string[]` | List of IDs of key results, insert all key results under the current objective when kr_ids is null or empty |
|       `period_display_status` | `string` | Period status **Optional values are**:  - `default`: Default - `normal`: Normal - `invalid`: Invalid - `hidden`: Hidden  |
|       `period_name_zh` | `string` | Chinese Period Name |
|       `period_name_en` | `string` | English Period Name |
|       `user_id` | `string` | The owner id of OKR |
|       `visible_setting` | `okr_visible_setting` | Visibility settings |
|         `progress_fill_area_visible` | `boolean` | Is the progress edit area visible |
|         `progress_status_visible` | `boolean` | Is the progress status visible |
|         `score_visible` | `boolean` | Is the score visible |
|     `okr_objective` | `okr_objective` | OKR Objective Block |
|       `objective_id` | `string` | Objective ID |
|       `confidential` | `boolean` | Whether private permissions have been set in OKR system |
|       `position` | `int` | The position number of the objective, corresponding to the 1 and 2 of O1 and O2 in the Block |
|       `score` | `int` | Score information |
|       `visible` | `boolean` | Whether the objective is displayed in OKR Block |
|       `weight` | `float` | Objective weight |
|       `progress_rate` | `okr_progress_rate` | Progress information |
|         `mode` | `string` | Status mode **Optional values are**:  - `simple`: Simple mode - `advanced`: Advanced mode  |
|         `current` | `float` | Current progress, used in advanced mode |
|         `percent` | `float` | Current progress percentage, used in simple mode |
|         `progress_status` | `string` | Status of progress **Optional values are**:  - `unset`: Not set - `normal`: Normal - `risk`: Risky - `extended`: Postponed  |
|         `start` | `float` | Progress starting value, used in advanced mode |
|         `status_type` | `string` | State Calculation Type **Optional values are**:  - `default`: Presenting the Highest Risk Key Result - `custom`: Custom  |
|         `target` | `float` | Progress target value, used in advanced mode |
|       `content` | `text` | Objective textual content |
|         `style` | `text_style` | Text style |
|           `align` | `int` | Alignment **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  |
|           `done` | `boolean` | Todo's status. Supports Todo Blocks |
|           `folded` | `boolean` | Folded state of text. Supports Heading1~9 block and the Text, Bullet, Ordered and Todo blocks with children |
|           `language` | `int` | Code block language. Only supports Code block **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|           `wrap` | `boolean` | Whether the code block wraps automatically. Only supports the Code block |
|           `background_color` | `string` | The background color of the block **Optional values are**:  - `LightGrayBackground`: Light gray - `LightRedBackground`: Light red - `LightOrangeBackground`: Light orange - `LightYellowBackground`: Light yellow - `LightGreenBackground`: Light green - `LightBlueBackground`: Light blue - `LightPurpleBackground`: Light purple - `PaleGrayBackground`: Medium gray - `DarkGrayBackground`: Gray - `DarkRedBackground`: Medium red - `DarkOrangeBackground`: Medium orange - `DarkYellowBackground`: Medium yellow - `DarkGreenBackground`: Medium green - `DarkBlueBackground`: Medium blue - `DarkPurpleBackground`: Medium purple  |
|           `indentation_level` | `string` | First line indentation level **Optional values are**:  - `NoIndent`: No indent - `OneLevelIndent`: One Level indent  |
|           `sequence` | `string` | Used to determine the number of an ordered list, either a number or 'auto' - Start a new list, numbering starts at 1, sequence='1' - Manually change the numbering, the numbering is set to a specific number, e.g. sequence='3' - Continue numbering, numbering is automatically consecutive, sequence='auto' - Some historical data and ordered block created through OpenAPI do not return this field |
|         `elements` | `text_element[]` | Text element |
|           `text_run` | `text_run` | Text. Supports Page, Text, Heading1~9, Bullet, Ordered, Code, Quote and Todo blocks |
|             `content` | `string` | Text content |
|             `text_element_style` | `text_element_style` | Text local style |
|               `bold` | `boolean` | Bold |
|               `italic` | `boolean` | Italic |
|               `strikethrough` | `boolean` | Strikethrough |
|               `underline` | `boolean` | Underline |
|               `inline_code` | `boolean` | Inline code |
|               `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|               `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|               `link` | `link` | Link |
|                 `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|               `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|           `mention_user` | `mention_user` | @User. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|             `user_id` | `string` | User OpenID |
|             `text_element_style` | `text_element_style` | Text element style |
|               `bold` | `boolean` | Bold |
|               `italic` | `boolean` | Italic |
|               `strikethrough` | `boolean` | Strikethrough |
|               `underline` | `boolean` | Underline |
|               `inline_code` | `boolean` | Inline code |
|               `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|               `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|               `link` | `link` | Link |
|                 `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|               `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|           `mention_doc` | `mention_doc` | @Document. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|             `token` | `string` | Docs token |
|             `obj_type` | `int` | Docs type **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|             `url` | `string` | Docs link (url_encode) |
|             `title` | `string` | The title of the referenced document |
|             `text_element_style` | `text_element_style` | Text element style |
|               `bold` | `boolean` | Bold |
|               `italic` | `boolean` | Italic |
|               `strikethrough` | `boolean` | Strikethrough |
|               `underline` | `boolean` | Underline |
|               `inline_code` | `boolean` | Inline code |
|               `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|               `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|               `link` | `link` | Link |
|                 `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|               `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|           `reminder` | `reminder` | Date Reminder. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|             `create_user_id` | `string` | Creator user ID |
|             `is_notify` | `boolean` | Whether to notify |
|             `is_whole_day` | `boolean` | Whether it's a date or a whole hour |
|             `expire_time` | `string` | Time the event occurred (millisecond timestamp) |
|             `notify_time` | `string` | Time to trigger notification (millisecond timestamp) |
|             `text_element_style` | `text_element_style` | Text element style |
|               `bold` | `boolean` | Bold |
|               `italic` | `boolean` | Italic |
|               `strikethrough` | `boolean` | Strikethrough |
|               `underline` | `boolean` | Underline |
|               `inline_code` | `boolean` | Inline code |
|               `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|               `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|               `link` | `link` | Link |
|                 `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|               `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|           `file` | `inline_file` | Inline file. Only supports deleting or moving the element, and creating new inline files is not supported |
|             `file_token` | `string` | File token |
|             `source_block_id` | `string` | The ID of the block in which the file is located in the current document |
|             `text_element_style` | `text_element_style` | Text element style |
|               `bold` | `boolean` | Bold |
|               `italic` | `boolean` | Italic |
|               `strikethrough` | `boolean` | Strikethrough |
|               `underline` | `boolean` | Underline |
|               `inline_code` | `boolean` | Inline code |
|               `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|               `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|               `link` | `link` | Link |
|                 `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|               `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|           `undefined` | `undefined_element` | Unsupported TextElements |
|           `inline_block` | `inline_block` | Inline block. Only supports deleting or moving the element, and creating new inline blocks is not supported |
|             `block_id` | `string` | Block ID |
|             `text_element_style` | `text_element_style` | Text element style |
|               `bold` | `boolean` | Bold |
|               `italic` | `boolean` | Italic |
|               `strikethrough` | `boolean` | Strikethrough |
|               `underline` | `boolean` | Underline |
|               `inline_code` | `boolean` | Inline code |
|               `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|               `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|               `link` | `link` | Link |
|                 `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|               `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|           `equation` | `equation` | Equation |
|             `content` | `string` | Equation content that conforms to KaTeX syntax. For more information, see [KaTex syntax rules](https://katex.org/docs/supported.html]. |
|             `text_element_style` | `text_element_style` | Text element style |
|               `bold` | `boolean` | Bold |
|               `italic` | `boolean` | Italic |
|               `strikethrough` | `boolean` | Strikethrough |
|               `underline` | `boolean` | Underline |
|               `inline_code` | `boolean` | Inline code |
|               `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|               `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|               `link` | `link` | Link |
|                 `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|               `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|     `okr_key_result` | `okr_key_result` | OKR Key Result Block |
|       `kr_id` | `string` | ID of key result |
|       `confidential` | `boolean` | Whether private permissions have been set  in OKR system |
|       `position` | `int` | The position number of the key result corresponds to 1 and 2 of KR1 and KR2 in the Block. |
|       `score` | `int` | Score information |
|       `visible` | `boolean` | Is this key result visible in OKR Block |
|       `weight` | `float` | Weight of key result |
|       `progress_rate` | `okr_progress_rate` | Progress information |
|         `mode` | `string` | Status mode **Optional values are**:  - `simple`: Simple mode - `advanced`: Advanced mode  |
|         `current` | `float` | Current progress, used in advanced mode |
|         `percent` | `float` | Current progress percentage, used in simple mode |
|         `progress_status` | `string` | Status of progress **Optional values are**:  - `unset`: Not set - `normal`: Normal - `risk`: Risky - `extended`: Postponed  |
|         `start` | `float` | Progress starting value, used in advanced mode |
|         `status_type` | `string` | State Calculation Type **Optional values are**:  - `default`: Presenting the Highest Risk Key Result - `custom`: Custom  |
|         `target` | `float` | Progress target value, used in advanced mode |
|       `content` | `text` | Key result textual content |
|         `style` | `text_style` | Text style |
|           `align` | `int` | Alignment **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  |
|           `done` | `boolean` | Todo's status. Supports Todo Blocks |
|           `folded` | `boolean` | Folded state of text. Supports Heading1~9 block and the Text, Bullet, Ordered and Todo blocks with children |
|           `language` | `int` | Code block language. Only supports Code block **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|           `wrap` | `boolean` | Whether the code block wraps automatically. Only supports the Code block |
|           `background_color` | `string` | The background color of the block **Optional values are**:  - `LightGrayBackground`: Light gray - `LightRedBackground`: Light red - `LightOrangeBackground`: Light orange - `LightYellowBackground`: Light yellow - `LightGreenBackground`: Light green - `LightBlueBackground`: Light blue - `LightPurpleBackground`: Light purple - `PaleGrayBackground`: Medium gray - `DarkGrayBackground`: Gray - `DarkRedBackground`: Medium red - `DarkOrangeBackground`: Medium orange - `DarkYellowBackground`: Medium yellow - `DarkGreenBackground`: Medium green - `DarkBlueBackground`: Medium blue - `DarkPurpleBackground`: Medium purple  |
|           `indentation_level` | `string` | First line indentation level **Optional values are**:  - `NoIndent`: No indent - `OneLevelIndent`: One Level indent  |
|           `sequence` | `string` | Used to determine the number of an ordered list, either a number or 'auto' - Start a new list, numbering starts at 1, sequence='1' - Manually change the numbering, the numbering is set to a specific number, e.g. sequence='3' - Continue numbering, numbering is automatically consecutive, sequence='auto' - Some historical data and ordered block created through OpenAPI do not return this field |
|         `elements` | `text_element[]` | Text element |
|           `text_run` | `text_run` | Text. Supports Page, Text, Heading1~9, Bullet, Ordered, Code, Quote and Todo blocks |
|             `content` | `string` | Text content |
|             `text_element_style` | `text_element_style` | Text local style |
|               `bold` | `boolean` | Bold |
|               `italic` | `boolean` | Italic |
|               `strikethrough` | `boolean` | Strikethrough |
|               `underline` | `boolean` | Underline |
|               `inline_code` | `boolean` | Inline code |
|               `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|               `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|               `link` | `link` | Link |
|                 `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|               `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|           `mention_user` | `mention_user` | @User. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|             `user_id` | `string` | User OpenID |
|             `text_element_style` | `text_element_style` | Text element style |
|               `bold` | `boolean` | Bold |
|               `italic` | `boolean` | Italic |
|               `strikethrough` | `boolean` | Strikethrough |
|               `underline` | `boolean` | Underline |
|               `inline_code` | `boolean` | Inline code |
|               `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|               `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|               `link` | `link` | Link |
|                 `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|               `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|           `mention_doc` | `mention_doc` | @Document. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|             `token` | `string` | Docs token |
|             `obj_type` | `int` | Docs type **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|             `url` | `string` | Docs link (url_encode) |
|             `title` | `string` | The title of the referenced document |
|             `text_element_style` | `text_element_style` | Text element style |
|               `bold` | `boolean` | Bold |
|               `italic` | `boolean` | Italic |
|               `strikethrough` | `boolean` | Strikethrough |
|               `underline` | `boolean` | Underline |
|               `inline_code` | `boolean` | Inline code |
|               `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|               `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|               `link` | `link` | Link |
|                 `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|               `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|           `reminder` | `reminder` | Date Reminder. Supports Text, Heading1~9, Bullet, Ordered, Quote and Todo blocks |
|             `create_user_id` | `string` | Creator user ID |
|             `is_notify` | `boolean` | Whether to notify |
|             `is_whole_day` | `boolean` | Whether it's a date or a whole hour |
|             `expire_time` | `string` | Time the event occurred (millisecond timestamp) |
|             `notify_time` | `string` | Time to trigger notification (millisecond timestamp) |
|             `text_element_style` | `text_element_style` | Text element style |
|               `bold` | `boolean` | Bold |
|               `italic` | `boolean` | Italic |
|               `strikethrough` | `boolean` | Strikethrough |
|               `underline` | `boolean` | Underline |
|               `inline_code` | `boolean` | Inline code |
|               `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|               `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|               `link` | `link` | Link |
|                 `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|               `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|           `file` | `inline_file` | Inline file. Only supports deleting or moving the element, and creating new inline files is not supported |
|             `file_token` | `string` | File token |
|             `source_block_id` | `string` | The ID of the block in which the file is located in the current document |
|             `text_element_style` | `text_element_style` | Text element style |
|               `bold` | `boolean` | Bold |
|               `italic` | `boolean` | Italic |
|               `strikethrough` | `boolean` | Strikethrough |
|               `underline` | `boolean` | Underline |
|               `inline_code` | `boolean` | Inline code |
|               `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|               `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|               `link` | `link` | Link |
|                 `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|               `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|           `undefined` | `undefined_element` | Unsupported TextElements |
|           `inline_block` | `inline_block` | Inline block. Only supports deleting or moving the element, and creating new inline blocks is not supported |
|             `block_id` | `string` | Block ID |
|             `text_element_style` | `text_element_style` | Text element style |
|               `bold` | `boolean` | Bold |
|               `italic` | `boolean` | Italic |
|               `strikethrough` | `boolean` | Strikethrough |
|               `underline` | `boolean` | Underline |
|               `inline_code` | `boolean` | Inline code |
|               `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|               `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|               `link` | `link` | Link |
|                 `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|               `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|           `equation` | `equation` | Equation |
|             `content` | `string` | Equation content that conforms to KaTeX syntax. For more information, see [KaTex syntax rules](https://katex.org/docs/supported.html]. |
|             `text_element_style` | `text_element_style` | Text element style |
|               `bold` | `boolean` | Bold |
|               `italic` | `boolean` | Italic |
|               `strikethrough` | `boolean` | Strikethrough |
|               `underline` | `boolean` | Underline |
|               `inline_code` | `boolean` | Inline code |
|               `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|               `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|               `link` | `link` | Link |
|                 `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|               `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|     `okr_progress` | `okr_progress` | OKR Progress Block |
|     `comment_ids` | `string[]` | Comment ids |
|     `jira_issue` | `jira_issue` | Jira Issue Block |
|       `id` | `string` | Jira Issue ID |
|       `key` | `string` | Jira Key |
|     `wiki_catalog` | `wiki_catalog` | Wiki Catalog Block |
|       `wiki_token` | `string` | wiki token |
|     `board` | `board` | Board Block |
|       `token` | `string` | Board Token |
|       `align` | `int` | Alignment **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  |
|       `width` | `int` | The width of the board, measured in pixels (px); when left blank, it will automatically adapt to the document's width. If the specified value exceeds the maximum width of the document, the page rendering will be constrained to the document's maximum width. |
|       `height` | `int` | The height of the board, measured in pixels (px); when left blank, it will be automatically calculated based on the content of the board. If the specified value exceeds twice the height of the screen, the page rendering will be constrained to twice the screen height. |
|     `agenda` | `agenda` | Agenda Block |
|     `agenda_item` | `agenda_item` | Agenda Item Blocks |
|     `agenda_item_title` | `agenda_item_title` | Agenda Item Title Block |
|       `elements` | `agenda_title_element[]` | Text element |
|         `text_run` | `text_run` | Text |
|           `content` | `string` | Text content |
|           `text_element_style` | `text_element_style` | Text local style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `mention_user` | `mention_user` | @user |
|           `user_id` | `string` | User OpenID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `mention_doc` | `mention_doc` | @Document |
|           `token` | `string` | Docs token |
|           `obj_type` | `int` | Docs type **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|           `url` | `string` | Docs link (url_encode) |
|           `title` | `string` | The title of the referenced document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `reminder` | `reminder` | Date Reminder. |
|           `create_user_id` | `string` | Creator user ID |
|           `is_notify` | `boolean` | Whether to notify |
|           `is_whole_day` | `boolean` | Whether it's a date or a whole hour |
|           `expire_time` | `string` | Time the event occurred (millisecond timestamp) |
|           `notify_time` | `string` | Time to trigger notification (millisecond timestamp) |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `file` | `inline_file` | Inline file. Only supports deleting or moving the element, and creating new inline files is not supported |
|           `file_token` | `string` | File token |
|           `source_block_id` | `string` | The ID of the block in which the file is located in the current document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `undefined` | `undefined_element` | Unsupported TextElement |
|         `inline_block` | `inline_block` | Inline block. Only supports deleting or moving the element, and creating new inline blocks is not supported |
|           `block_id` | `string` | Block ID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|         `equation` | `equation` | Equation |
|           `content` | `string` | Equation content that conforms to KaTeX syntax. For more information, see [KaTex syntax rules](https://katex.org/docs/supported.html]. |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Medium grey - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Light gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see: Get Comment API. |
|       `align` | `int` | alignment **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  |
|     `agenda_item_content` | `agenda_item_content` | Agenda Item Content Block |
|     `link_preview` | `link_preview` | Link Preview Block |
|       `url` | `string` | link |
|       `url_type` | `string` | Link type **Optional values are**:  - `MessageLink`: Message link - `Undefined`: Undefined link type  |
|   `revision_id` | `int` | The version number of the group announcement after the current block children is successfully created |
|   `client_token` | `string` | The unique identifier of the operation, which is used in the update request to indicate that the update is idempotent | ### Response body example

{
  "code": 0,
  "data": {
    "children": [
      {
        "block_id": "doxcnXgNGAtaAraIRVeCfmbx4Eo",
        "block_type": 2,
        "parent_id": "oc_5ad11d72b830411d72b836c20",
        "text": {
          "elements": [
            {
              "text_run": {
                "content": "When you need to convey important information to all members of the group, you can use the group announcement function. Subsequent members of the group can also see the current group announcement.",
                "text_element_style": {
                  "background_color": 14,
                  "text_color": 5
                }
              }
            },
            {
              "text_run": {
                "content": "Group announcement editing supports real-time saving and multi-person collaboration, and group members can view the latest group announcement content in real time.",
                "text_element_style": {
                  "background_color": 14,
                  "bold": true,
                  "text_color": 5
                }
              }
            }
          ],
          "style": {}
        }
      }
    ],
    "client_token": "ea403093-3af1-4e9d-8f5d-53c5a4e4c36e",
    "document_revision_id": 2
  },
  "msg": ""
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1770001 | invalid param | Confirm whether the passed parameters are valid |
| 404 | 1770002 | not found | **For Document APIs**: The `document_id` of the document does not exist. Please verify that the document has not been deleted or that the `document_id` is correct. Refer to Document overview for information on how to get the `document_id` of a document. **For Group Announcement APIs**: The `chat_id` of the group announcement does not exist. Please verify that the group has not been disbanded or if the `chat_id` is correctly inputted. |
| 400 | 1770003 | resource deleted | Confirm whether the resource has been deleted |
| 400 | 1770004 | too many blocks in document | Confirm whether the number of document blocks exceeds the limit |
| 400 | 1770005 | too deep level in document | Confirm whether the level of document blocks exceeds the limit |
| 400 | 1770006 | schema mismatch | Confirm whether the document structure is valid |
| 400 | 1770007 | too many children in block | Confirm whether the number of children for a specified block exceeds the limit |
| 400 | 1770008 | too big file size | Confirm whether the uploaded file size exceeds the limit |
| 400 | 1770010 | too many table column | Confirm whether the number of table columns exceeds the limit |
| 400 | 1770011 | too many table cell | Confirm whether the number of table cells exceeds the limit |
| 400 | 1770012 | too many grid column | Confirm whether the number of grid columns exceeds the limit |
| 400 | 1770013 | relation mismatch | Images, files, and other resources are incorrectly associated. Please make sure that you have uploaded the related image or file material to the corresponding document block at the same time when you create the image or file block. Please refer to the documentation FAQ 3 and 4 for details |
| 400 | 1770014 | parent children relation mismatch | Confirm whether the parent-child relationships of blocks are correct |
| 400 | 1770015 | single edit with multi document | Confirm whether the block belongs to the specified document |
| 400 | 1770029 | block not support to create | Confirm whether the specified block supports creation |
| 400 | 1770019 | repeated blockID in document | Confirm whether there are duplicate Block IDs in the document |
| 400 | 1770020 | operation denied on copying document | Confirm whether the document is currently being copied |
| 400 | 1770021 | too old document | Confirm if the specified document version is too old. The difference between the specified `revision_id` and the latest `revision_id` of the document cannot exceed 1000 |
| 400 | 1770041 | open schema mismatch | Confirm whether the block parent-child relationship is legal. To learn about the parent-child relationship between blocks, see Document overview-Basic concept. |
| 400 | 1770024 | invalid operation | Confirm whether the operation is valid |
| 400 | 1770025 | operation and block not match | Confirm whether the corresponding operation of the specified block is valid |
| 400 | 1770026 | row operation over range | Confirm if the row operation index is out of bounds |
| 400 | 1770027 | column operation over range | Confirm if the column operation index is out of bounds |
| 400 | 1770028 | block not support create children | Confirm if adding children to the specified block is valid |
| 400 | 1770029 | block not support to create | Confirm whether the specified block supports creation |
| 400 | 1770030 | invalid parent children relation | Confirm whether the parent-child relationship for the specified operation is valid |
| 400 | 1770031 | block not support to delete children | Confirm if the specified block supports deleting children |
| 400 | 1770033 | raw content size exceed limited | Confirm whether plain text content size exceeds limit |
| 400 | 1770034 | operation count exceed limited | Too many cells are involved in the current request. Please split into multiple requests |
| 400 | 1770035 | Resource count exceeds limit | The number of resources in the current request exceeds the limit, please split it into multiple requests. The upper limit of various resources is: 200 ChatCard, 200 File, 200 MentionDoc, 200 MentionUser, 20 Image, 20 ISV, 5 Sheet and 5 Bitable. |
| 400 | 1770038 | resource not found | The inserted resource was not found or the resource did not have permission to insert, please check that the resource identifier is correct. |
| 400 | 1772001 | Chat announcement can NOT be found in chat information. | The group announcement information is abnormal. Please retry the operation. If the issue continues, consult [Technical Support](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) for assistance. |
| 400 | 1772002 | Operator and chat can NOT be in different tenants. | The operator and the group being operated should be under the same tenant. |
| 400 | 1772003 | Operator can NOT be out of the chat. | The operator is not in the group. You need to add the application or user currently calling the API to the group to be operated and try again. |
| 400 | 1772004 | The operator or invited bots does NOT have the authority to manage external chats without the scope. | The app does not have permission to operate external groups. |
| 400 | 1772025 | Bot ability is not activated. | The app does not have bot capabilities enabled. You need to log in to the [Developer Console](https://open.larksuite.com/app), add the **Bot** capability in **Features** > **Add Features** on the app details page, and publish the app to make the configuration effective. For specific operations, see Bot Capabilities. |
| 400 | 1772006 | announcement type is not supported | Except for the "Obtain the basic information of a group announcement" api, other Upgraded group announcements apis cannot operate doc-type group announcements. If you want to operate doc-type group announcements, please refer to the "Old version of group announcements" api. |
| 400 | 1772005 | No Permission: Only chat owner or admin can edit chat information in the current situation. | Only the group owner or administrators are allowed to edit the group information. |
| 400 | 1772034 | The app is unavailable or inactivated by the tenant. | The app is unavailable or inactivated by the tenant. |
| 403 | 1770032 | forbidden | **For Document APIs**: - Confirm whether the current access identity has permission to read or edit documents. Please refer to the following methods to resolve this: - If you are using a `tenant_access_token`, it means the current application does not have permission to read or edit documents. You need to add document permissions for the application through the cloud document webpage by navigating to the top right corner **"..."** -> **"... More"** -> **"Add applications"**. **Note**: Before adding a document application, you need to ensure that the target application has at least one cloud document API permission enabled. Otherwise, you will not be able to search for the target application in the document application window. ![](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/bb60f97ebb402475f2af1d3131d4914f_sLOzoqYRXX.png?height=1992&maxWidth=550&width=3278) - If you are using a `user_access_token`, it means the current user does not have permission to read or edit documents. You need to add document permissions for the current user through the **Share** entry in the top right corner of the cloud document webpage. ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/caceea2ac91c170555194d7a8dc2a317_GfTRc9xLAt.png?height=1992&maxWidth=550&width=3278) For more details on the specific steps or other methods to add permissions, refer to Cloud Document FAQ 3. - Confirm whether the current access identity has permission to read MentionDoc, i.e., @document. - Confirm whether the user in MentionUser, i.e., @user, is currently employed and is a contact of the current access identity. - Confirm whether the current access identity has permission to view and share group cards. - Confirm whether the current access identity has permission to access specific Wiki subdirectories. - Confirm whether the current access identity has permission to view document blocks such as OKR, ISV, Add-Ons, etc. **For Group Announcement APIs**: The current operator does not have permission to edit group announcemen. Solution: - Solution 1: Call the Specify group administrator interface to set the current operator as a group administrator, and then try again. - Solution 2: In **Lark Client > Group > Settings > Group Settings**, set **Who can edit group info** to **Everyone in this group**, and then try again. When working with APIs for creating and updating group announcement, ensure the following: - Confirm whether the current access identity has permission to read MentionDoc, i.e., @document. - Confirm whether the user in MentionUser, i.e., @user, is currently employed and is a contact of the current access identity. - Confirm whether the current access identity has permission to view and share group cards. - Confirm whether the current access identity has permission to access specific Wiki subdirectories. - Confirm whether the current access identity has permission to view document blocks such as OKR, ISV, Add-Ons, etc. |
| 500 | 1771001 | server internal error | Server internal error. Please retry the operation. If the issue continues, consult [Technical Support](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) for assistance. |
| 500 | 1771002 | gateway server internal error | Gateway service internal error. Please retry the operation. If the issue continues, consult [Technical Support](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) for assistance. |
| 500 | 1771003 | gateway marshal error | Gateway service parsing error. Please retry the operation. If the issue continues, consult [Technical Support](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) for assistance. |
| 500 | 1771004 | gateway unmarshal error | Gateway service unmarshalling error. Please retry the operation. If the issue continues, consult [Technical Support](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) for assistance. |
| 503 | 1771005 | system under maintenance | System service is under maintenance |
