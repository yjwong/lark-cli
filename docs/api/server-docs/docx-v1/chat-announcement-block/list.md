---
title: "Obtain all blocks of a group announcement"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/document-docx/docx-v1/chat-announcement-block/list"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/docx/v1/chats/:chat_id/announcement/blocks"
service: "docx-v1"
resource: "chat-announcement-block"
section: "Group Chat"
scopes:
  - "im:chat.announcement:read"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1739935021000"
---

# Obtain all blocks of a group announcement

Obtains the rich text content of all blocks in a specified group announcement and return it paginated.

> **Application frequency limit:** the maximum frequency of a single application call is 10 times per second, beyond which the interface will return HTTP status code 400 and error code 99991400. When a request is frequency-limited, the application needs to handle the frequency-limited status code and use an exponential fallback algorithm or some other frequency-control strategy to reduce the rate of calls to the API.

## Prerequisites

- The application needs to enable bot capability.
- The user or robot calling the current interface must be in the corresponding group.
- When obtaining internal group information, the user or robot calling the current interface must be in the same tenant as the corresponding group.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/docx/v1/chats/:chat_id/announcement/blocks |
| HTTP Method | GET |
| Supported app types | custom,isv |
| Required scopes | `im:chat.announcement:read` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `chat_id` | `string` | Group ID. How to get it: - Create a group, get the chat_id of the group from the returned result. - Call the Get the list of groups that the user or robot is in interface to query the chat_id of the group that the user or robot is in. - Call the Search the list of groups visible to the user or robot to search for the chat_id of the group that the user or robot is in and the group that is open to the user or robot. **Note**: Single chat (group type is `p2p`) does not support getting group announcements. **Example value**: "oc_5ad11d72b830411d72b836c20" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `page_size` | `int` | No | Page size **Example value**: 500 **Default value**: `500` **Data validation rules**: - Maximum value: `500` |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: "aw7DoMKBFMOGwqHCrcO8w6jCmMOvw6ILeADCvsKNw57Di8O5XGV3LG4_w5HCqhFxSnDCrCzCn0BgZcOYUg85EMOYcEAcwqYOw4ojw5QFwofCu8KoIMO3K8Ktw4IuNMOBBHNYw4bCgCV3U1zDu8K-J8KSR8Kgw7Y0fsKZdsKvW3d9w53DnkHDrcO5bDkYwrvDisOEPcOtVFJ-I03CnsOILMOoAmLDknd6dsKqG1bClAjDuS3CvcOTwo7Dg8OrwovDsRdqIcKxw5HDohTDtXN9w5rCkWo" |
| `revision_id` | `int` | No | The group announcement version to be queried. -1 indicates the latest version of the group announcement. Once the group announcement is created, the `revision_id ` is 1. If querying the latest version, reading permissions for the group announcement are required. If querying a historical version, editing permissions for the group announcement are necessary. You can use the Obtain the basic information of a group announcement api to get the latest revision ID **Example value**: -1 **Default value**: `-1` **Data validation rules**: - Minimum value: `-1` |
| `user_id_type` | `string` | No | User ID categories **Example value**: "open_id" **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ### Request example

```bash
curl --location --request GET 'https://open.larksuite.com/open-apis/docx/v1/chats/oc_5ad11d72b830411d72b836c20/announcement/blocks?page_size=100&revision_id=-1' \
--header 'Authorization: Bearer u-xxx'
#You need to replace 'u-xxx' in 'Authorization: Bearer u-xxx' with the real access token 
```

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `items` | `block[]` | Block information for the group announcement |
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
|       `url_type` | `string` | Link type **Optional values are**:  - `MessageLink`: Message link - `Undefined`: Undefined link type  |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned |
|   `has_more` | `boolean` | Whether the response body has more parameters | ### Response body example

{
    "code": 0,
    "data":
    {
        "has_more": false,
        "items":
        [
            {
                "block_id": "oc_5ad11d72b830411d72b836c20",
                "block_type": 1,
                "children":
                [
                    "EJ8pdBT6RoA2BtxKXfuc1TD0n8f",
                    "WoF4dolv0ohvegxctR9c1t0sn4b",
                    "UPD7dhkrLo8ZCfxoAYLcuEbDnzd"
                ],
                "page":
                {
                    "elements":
                    [
                        {
                            "text_run":
                            {
                                "content": "",
                                "text_element_style":
                                {
                                    "bold": false,
                                    "inline_code": false,
                                    "italic": false,
                                    "strikethrough": false,
                                    "underline": false
                                }
                            }
                        }
                    ],
                    "style":
                    {
                        "align": 1
                    }
                },
                "parent_id": ""
            },
            {
                "block_id": "EJ8pdBT6RoA2BtxKXfuc1TD0n8f",
                "block_type": 3,
                "heading1":
                {
                    "elements":
                    [
                        {
                            "text_run":
                            {
                                "content": "Group announcement",
                                "text_element_style":
                                {
                                    "bold": false,
                                    "inline_code": false,
                                    "italic": false,
                                    "strikethrough": false,
                                    "underline": false
                                }
                            }
                        }
                    ],
                    "style":
                    {
                        "align": 1,
                        "folded": false
                    }
                },
                "parent_id": "oc_5ad11d72b830411d72b836c20"
            },
            {
                "block_id": "WoF4dolv0ohvegxctR9c1t0sn4b",
                "block_type": 2,
                "parent_id": "oc_5ad11d72b830411d72b836c20",
                "text":
                {
                    "elements":
                    [
                        {
                            "text_run":
                            {
                                "content": "Use the group announcement feature to communicate ",
                                "text_element_style":
                                {
                                    "bold": false,
                                    "inline_code": false,
                                    "italic": false,
                                    "strikethrough": false,
                                    "underline": false
                                }
                            }
                        },
                        {
                            "text_run":
                            {
                                "content": "important information",
                                "text_element_style":
                                {
                                    "bold": true,
                                    "inline_code": false,
                                    "italic": false,
                                    "strikethrough": false,
                                    "underline": false
                                }
                            }
                        },
                        {
                            "text_run":
                            {
                                "content": " to all members of a group. New members who join later can also view the current group announcement.",
                                "text_element_style":
                                {
                                    "bold": false,
                                    "inline_code": false,
                                    "italic": false,
                                    "strikethrough": false,
                                    "underline": false
                                }
                            }
                        }
                    ],
                    "style":
                    {
                        "align": 1,
                        "folded": false
                    }
                }
            },
            {
                "block_id": "UPD7dhkrLo8ZCfxoAYLcuEbDnzd",
                "block_type": 2,
                "parent_id": "oc_5ad11d72b830411d72b836c20",
                "text":
                {
                    "elements":
                    [
                        {
                            "text_run":
                            {
                                "content": "The group announcement editor now ",
                                "text_element_style":
                                {
                                    "bold": false,
                                    "inline_code": false,
                                    "italic": false,
                                    "strikethrough": false,
                                    "underline": false
                                }
                            }
                        },
                        {
                            "text_run":
                            {
                                "content": "supports real-time saving and collaborative editing",
                                "text_element_style":
                                {
                                    "background_color": 4,
                                    "bold": false,
                                    "inline_code": false,
                                    "italic": false,
                                    "strikethrough": false,
                                    "underline": false
                                }
                            }
                        },
                        {
                            "text_run":
                            {
                                "content": ", allowing group members to ",
                                "text_element_style":
                                {
                                    "bold": false,
                                    "inline_code": false,
                                    "italic": false,
                                    "strikethrough": false,
                                    "underline": false
                                }
                            }
                        },
                        {
                            "text_run":
                            {
                                "content": "view the latest announcement content instantly.",
                                "text_element_style":
                                {
                                    "background_color": 4,
                                    "bold": false,
                                    "inline_code": false,
                                    "italic": false,
                                    "strikethrough": false,
                                    "underline": false
                                }
                            }
                        }
                    ],
                    "style":
                    {
                        "align": 1,
                        "folded": false
                    }
                }
            }
        ]
    },
    "msg": "success"
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1770001 | invalid param | Confirm whether the passed parameters are valid |
| 404 | 1770002 | not found | **For Document APIs**: The `document_id` of the document does not exist. Please verify that the document has not been deleted or that the `document_id` is correct. Refer to Document overview for information on how to get the `document_id` of a document. **For Group Announcement APIs**: The `chat_id` of the group announcement does not exist. Please verify that the group has not been disbanded or if the `chat_id` is correctly inputted. |
| 400 | 1772001 | Chat announcement can NOT be found in chat information. | The group announcement information is abnormal. Please retry the operation. If the issue continues, consult [Technical Support](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) for assistance. |
| 400 | 1772002 | Operator and chat can NOT be in different tenants. | The operator and the group being operated should be under the same tenant. |
| 400 | 1772003 | Operator can NOT be out of the chat. | The operator is not in the group. You need to add the application or user currently calling the API to the group to be operated and try again. |
| 400 | 1772004 | The operator or invited bots does NOT have the authority to manage external chats without the scope. | The app does not have permission to operate external groups. |
| 400 | 1772025 | Bot ability is not activated. | The app does not have bot capabilities enabled. You need to log in to the [Developer Console](https://open.larksuite.com/app), add the **Bot** capability in **Features** > **Add Features** on the app details page, and publish the app to make the configuration effective. For specific operations, see Bot Capabilities. |
| 400 | 1772006 | announcement type is not supported | Except for the "Obtain the basic information of a group announcement" api, other Upgraded group announcements apis cannot operate doc-type group announcements. If you want to operate doc-type group announcements, please refer to the "Old version of group announcements" api. |
| 400 | 1772034 | The app is unavailable or inactivated by the tenant. | The app is unavailable or inactivated by the tenant. |
| 403 | 1770032 | forbidden | **For Document APIs**: - Confirm whether the current access identity has permission to read or edit documents. Please refer to the following methods to resolve this: - If you are using a `tenant_access_token`, it means the current application does not have permission to read or edit documents. You need to add document permissions for the application through the cloud document webpage by navigating to the top right corner **"..."** -> **"... More"** -> **"Add applications"**. **Note**: Before adding a document application, you need to ensure that the target application has at least one cloud document API permission enabled. Otherwise, you will not be able to search for the target application in the document application window. ![](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/bb60f97ebb402475f2af1d3131d4914f_sLOzoqYRXX.png?height=1992&maxWidth=550&width=3278) - If you are using a `user_access_token`, it means the current user does not have permission to read or edit documents. You need to add document permissions for the current user through the **Share** entry in the top right corner of the cloud document webpage. ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/caceea2ac91c170555194d7a8dc2a317_GfTRc9xLAt.png?height=1992&maxWidth=550&width=3278) For more details on the specific steps or other methods to add permissions, refer to Cloud Document FAQ 3. - Confirm whether the current access identity has permission to read MentionDoc, i.e., @document. - Confirm whether the user in MentionUser, i.e., @user, is currently employed and is a contact of the current access identity. - Confirm whether the current access identity has permission to view and share group cards. - Confirm whether the current access identity has permission to access specific Wiki subdirectories. - Confirm whether the current access identity has permission to view document blocks such as OKR, ISV, Add-Ons, etc. **For Group Announcement APIs**: The current operator does not have permission to edit group announcemen. Solution: - Solution 1: Call the Specify group administrator interface to set the current operator as a group administrator, and then try again. - Solution 2: In **Lark Client > Group > Settings > Group Settings**, set **Who can edit group info** to **Everyone in this group**, and then try again. When working with APIs for creating and updating group announcement, ensure the following: - Confirm whether the current access identity has permission to read MentionDoc, i.e., @document. - Confirm whether the user in MentionUser, i.e., @user, is currently employed and is a contact of the current access identity. - Confirm whether the current access identity has permission to view and share group cards. - Confirm whether the current access identity has permission to access specific Wiki subdirectories. - Confirm whether the current access identity has permission to view document blocks such as OKR, ISV, Add-Ons, etc. |
| 500 | 1771001 | server internal error | Server internal error. Please retry the operation. If the issue continues, consult [Technical Support](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) for assistance. |
| 500 | 1771002 | gateway server internal error | Gateway service internal error. Please retry the operation. If the issue continues, consult [Technical Support](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) for assistance. |
| 500 | 1771003 | gateway marshal error | Gateway service parsing error. Please retry the operation. If the issue continues, consult [Technical Support](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) for assistance. |
| 500 | 1771004 | gateway unmarshal error | Gateway service unmarshalling error. Please retry the operation. If the issue continues, consult [Technical Support](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) for assistance. |
| 503 | 1771005 | system under maintenance | System service is under maintenance |
