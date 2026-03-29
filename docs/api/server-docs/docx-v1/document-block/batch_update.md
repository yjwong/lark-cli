---
title: "Batch Update Blocks"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/document-docx/docx-v1/document-block/batch_update"
method: "PATCH"
api_path: "https://open.larksuite.com/open-apis/docx/v1/documents/:document_id/blocks/batch_update"
service: "docx-v1"
resource: "document-block"
section: "Docs"
scopes:
  - "docx:document"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1703648376000"
---

# Batch Update Blocks

Batch update the specified block and return its updated rich text content.

> **Application frequency limit:** the maximum frequency of a single application call is 3 times per second, beyond which the interface will return HTTP status code 400 and error code 99991400.
> 
> **Document frequency limit**: The maximum number of concurrent edits per second for a single document is 3. If the frequency limit is exceeded, the interface will return HTTP status code 429. The edit operations include:
> 
> - Create Blocks
> - Delete Blocks
> - Update Block
> - Batch Update Blocks
> 
> When a request is frequency-limited, the application needs to handle the frequency-limited status code and use an exponential fallback algorithm or some other frequency-control strategy to reduce the rate of calls to the API.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/docx/v1/documents/:document_id/blocks/batch_update |
| HTTP Method | PATCH |
| Supported app types | custom,isv |
| Required scopes | `docx:document` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `document_id` | `string` | Unique identification of the document. Corresponding to the token of the Upgraded Docs, Click to learn how to get document_id. **Example value**: "doxcnePuYufKa49ISjhD8Ih0ikh" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `document_revision_id` | `int` | No | The document version of the operation, -1 indicates the latest version of the document. **Example value**: -1 **Default value**: `-1` **Data validation rules**: - Minimum value: `-1` |
| `client_token` | `string` | No | The unique identifier of the operation, corresponding to the client_token of the interface return value, is used for idempotent update operations. This value is null to indicate that a new request will be initiated, and this value is not null to indicate idempotent update operations. **Example value**: "0e2633a3-aa1a-4171-af9e-0768ff863566" |
| `user_id_type` | `string` | No | User ID categories **Example value**: "open_id" **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `requests` | `update_block_request[]` | Yes | Batch update block requests **Data validation rules**: - Maximum length: `200` |
|   `update_text_elements` | `update_text_elements_request` | No | Update text element request |
|     `elements` | `text_element[]` | Yes | Updated text element list, a single update reminder limit of 30, mention_doc limit of 50, mention_user limit of 100 **Data validation rules**: - Minimum length: `1` |
|       `text_run` | `text_run` | No | Text |
|         `content` | `string` | Yes | Text content **Example value**: "text" |
|         `text_element_style` | `text_element_style` | No | Text local style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. **Example value**: ["1660030311959965796"] |
|       `mention_user` | `mention_user` | No | @User |
|         `user_id` | `string` | Yes | User OpenID **Example value**: "ou_3bbe8a09c20e89cce9bff989ed840674" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. **Example value**: ["1660030311959965796"] |
|       `mention_doc` | `mention_doc` | No | @Document |
|         `token` | `string` | Yes | Cloud document token **Example value**: "doxbc873Y7cXD153gXqb76G1Y9b" |
|         `obj_type` | `int` | Yes | Cloud Document Type **Example value**: 22 **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|         `url` | `string` | Yes | Docs link (url_encode) **Example value**: "https%3A%2F%2Fbytedance.larksuite.com%2Fdocx%2Fdoxbc873Y7cXD153gXqb76G1Y9b" |
|         `title` | `string` | No | The title of the referenced document **Example value**: "undefined" **Data validation rules**: - Length range: `0` ～ `800` characters |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. **Example value**: ["1660030311959965796"] |
|       `reminder` | `reminder` | No | Date Reminder |
|         `create_user_id` | `string` | Yes | Creator user ID **Example value**: "0E2633a3-aa1a-4171-af9e-0768ff863566" |
|         `is_notify` | `boolean` | No | Whether to notify **Example value**: true **Default value**: `false` |
|         `is_whole_day` | `boolean` | No | Is it the date or the hour? **Example value**: true **Default value**: `false` |
|         `expire_time` | `string` | Yes | Time the event occurred (millisecond timestamp) **Example value**: "1641967200000" |
|         `notify_time` | `string` | Yes | Time to trigger notification (millisecond timestamp) **Example value**: "1643166000000" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. **Example value**: ["1660030311959965796"] |
|       `file` | `inline_file` | No | inline attachment |
|         `file_token` | `string` | No | Attachment token **Example value**: "boxcnOj88GDkmWGm2zsTyCBqoLb" |
|         `source_block_id` | `string` | No | The id of the block in which the attachment is located in the current document **Example value**: "doxcnM46kSWSkgUMW04ldKsJDsc" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. **Example value**: ["1660030311959965796"] |
|       `undefined` | `undefined_element` | No | Unsupported TextElements |
|       `inline_block` | `inline_block` | No | Inline block |
|         `block_id` | `string` | Yes | Block ID **Example value**: "doxcnPFi0R56ctbvh2MjkkROFWf" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. **Example value**: ["1660030311959965796"] |
|       `equation` | `equation` | No | Equation |
|         `content` | `string` | Yes | Equation content that conforms to KaTeX syntax, please refer to the syntax rules: https://katex.org/docs/supported.html **Example value**: "E=mc^2\n" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. **Example value**: ["1660030311959965796"] |
|   `update_text_style` | `update_text_style_request` | No | Update text style request |
|     `style` | `text_style` | Yes | Text style |
|       `align` | `int` | No | Alignment **Example value**: 1 **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  **Default value**: `1` |
|       `done` | `boolean` | No | Todo completion status **Example value**: true **Default value**: `false` |
|       `folded` | `boolean` | No | Folded state of text **Example value**: true **Default value**: `false` |
|       `language` | `int` | No | Code block language **Example value**: 1 **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|       `wrap` | `boolean` | No | Whether the code block wraps automatically **Example value**: true **Default value**: `false` |
|     `fields` | `int[]` | Yes | Fields that should be updated must specify at least one field. For example, to adjust Block alignment, set fields to [1]. **Example value**: Modified text style properties **Optional values are**:  - `1`: Modify Block Alignment - `2`: Modify todo block completion status - `3`: Modify the folded state of the block - `4`: Modify the language type of the code block - `5`: Modify the collapsed state of code blocks  |
|   `update_table_property` | `update_table_property_request` | No | Update table property request |
|     `column_width` | `int` | No | Table column width **Example value**: 100 **Data validation rules**: - Minimum value: `50` |
|     `column_index` | `int` | No | Index of table columns whose column width needs to be modified **Example value**: 0 **Data validation rules**: - Minimum value: `0` |
|     `header_row` | `boolean` | No | Set first line header line **Example value**: false |
|     `header_column` | `boolean` | No | Set first column header column **Example value**: false |
|   `insert_table_row` | `insert_table_row_request` | No | Table insert new row request |
|     `row_index` | `int` | Yes | The index of the inserted row in the table. (-1 means to insert a row at the end of the table) **Example value**: -1 **Data validation rules**: - Minimum value: `-1` |
|   `insert_table_column` | `insert_table_column_request` | No | Table insert new column request |
|     `column_index` | `int` | Yes | The index of the inserted column in the table. (-1 means to insert a column at the end of the table) **Example value**: -1 **Data validation rules**: - Minimum value: `-1` |
|   `delete_table_rows` | `delete_table_rows_request` | No | Table batch delete row request |
|     `row_start_index` | `int` | Yes | Row start index (interval closed left and right open) **Example value**: 0 **Data validation rules**: - Minimum value: `0` |
|     `row_end_index` | `int` | Yes | Row end index (interval closed left and right open) **Example value**: 1 **Data validation rules**: - Minimum value: `1` |
|   `delete_table_columns` | `delete_table_columns_request` | No | Table batch delete column request |
|     `column_start_index` | `int` | Yes | Column start index (interval closed left and right open) **Example value**: 0 **Data validation rules**: - Minimum value: `0` |
|     `column_end_index` | `int` | Yes | Column end index (interval closed left and right open) **Example value**: 1 **Data validation rules**: - Minimum value: `1` |
|   `merge_table_cells` | `merge_table_cells_request` | No | Table merge cell request |
|     `row_start_index` | `int` | Yes | Row start index (interval closed left and right open) **Example value**: 0 **Data validation rules**: - Minimum value: `0` |
|     `row_end_index` | `int` | Yes | Row end index (interval closed left and right open) **Example value**: 1 **Data validation rules**: - Minimum value: `1` |
|     `column_start_index` | `int` | Yes | Column start index (interval closed left and right open) **Example value**: 0 **Data validation rules**: - Minimum value: `0` |
|     `column_end_index` | `int` | Yes | Column end index (interval closed left and right open) **Example value**: 1 **Data validation rules**: - Minimum value: `1` |
|   `unmerge_table_cells` | `unmerge_table_cells_request` | No | Table Cancel Cell Merge Status Request |
|     `row_index` | `int` | Yes | Table row index **Example value**: 0 **Data validation rules**: - Minimum value: `0` |
|     `column_index` | `int` | Yes | Table column index **Example value**: 0 **Data validation rules**: - Minimum value: `0` |
|   `insert_grid_column` | `insert_grid_column_request` | No | Column insert new column request |
|     `column_index` | `int` | Yes | Insert the column index, starting from 1, such as 1 means inserting after the first column, note that 0 is not allowed (-1 means inserting after the last column) **Example value**: 1 **Data validation rules**: - Minimum value: `-1` |
|   `delete_grid_column` | `delete_grid_column_request` | No | Column deletion request |
|     `column_index` | `int` | Yes | Delete the column index, starting from 0, if 0 means deleting the first column (-1 means deleting the last column) **Example value**: 0 **Data validation rules**: - Minimum value: `-1` |
|   `update_grid_column_width_ratio` | `update_grid_column_width_ratio_request` | No | Update column width ratio request |
|     `width_ratios` | `int[]` | Yes | When updating the column width ratio, you need to pass in all column width ratios **Example value**: 50 **Data validation rules**: - Length range: `1` ～ `99` |
|   `replace_image` | `replace_image_request` | No | Replace image request |
|     `token` | `string` | Yes | Picture token **Example value**: "BoxbckbfvfcqEg22hAzN8Dh9gJd" |
|     `width` | `int` | No | Image width(px) **Example value**: 100 |
|     `height` | `int` | No | Image height(px) **Example value**: 100 |
|     `align` | `int` | No | alignment **Example value**: 2 **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  |
|   `replace_file` | `replace_file_request` | No | Request for replacement attachment |
|     `token` | `string` | Yes | Attachment token **Example value**: "BoxbckbfvfcqEg22hAzN8Dh9gJd" |
|   `block_id` | `string` | No | Block unique identification **Example value**: "doxcnSS4ouQkQEouGSUkTg9NJPe" |
|   `update_text` | `update_text_request` | No | Update text elements and style requests |
|     `elements` | `text_element[]` | Yes | Updated list of text elements, up to 30 reminders, up to 50 mention_doc, and up to 100 mention_user in a single update **Data validation rules**: - Minimum length: `1` |
|       `text_run` | `text_run` | No | Text |
|         `content` | `string` | Yes | Text content **Example value**: "text" |
|         `text_element_style` | `text_element_style` | No | Text local style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. **Example value**: ["1660030311959965796"] |
|       `mention_user` | `mention_user` | No | @User |
|         `user_id` | `string` | Yes | User OpenID **Example value**: "ou_3bbe8a09c20e89cce9bff989ed840674" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. **Example value**: ["1660030311959965796"] |
|       `mention_doc` | `mention_doc` | No | @Document |
|         `token` | `string` | Yes | Cloud document token **Example value**: "doxbc873Y7cXD153gXqb76G1Y9b" |
|         `obj_type` | `int` | Yes | Cloud Document Type **Example value**: 22 **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|         `url` | `string` | Yes | Docs link (url_encode) **Example value**: "https%3A%2F%2Fbytedance.larksuite.com%2Fdocx%2Fdoxbc873Y7cXD153gXqb76G1Y9b" |
|         `title` | `string` | No | The title of the referenced document **Example value**: "undefined" **Data validation rules**: - Length range: `0` ～ `800` characters |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. **Example value**: ["1660030311959965796"] |
|       `reminder` | `reminder` | No | Date Reminder |
|         `create_user_id` | `string` | Yes | Creator user ID **Example value**: "0E2633a3-aa1a-4171-af9e-0768ff863566" |
|         `is_notify` | `boolean` | No | Whether to notify **Example value**: true **Default value**: `false` |
|         `is_whole_day` | `boolean` | No | Is it the date or the hour? **Example value**: true **Default value**: `false` |
|         `expire_time` | `string` | Yes | Time the event occurred (millisecond timestamp) **Example value**: "1641967200000" |
|         `notify_time` | `string` | Yes | Time to trigger notification (millisecond timestamp) **Example value**: "1643166000000" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. **Example value**: ["1660030311959965796"] |
|       `file` | `inline_file` | No | inline attachment |
|         `file_token` | `string` | No | Attachment token **Example value**: "boxcnOj88GDkmWGm2zsTyCBqoLb" |
|         `source_block_id` | `string` | No | The id of the block in which the attachment is located in the current document **Example value**: "doxcnM46kSWSkgUMW04ldKsJDsc" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. **Example value**: ["1660030311959965796"] |
|       `undefined` | `undefined_element` | No | Unsupported TextElements |
|       `inline_block` | `inline_block` | No | Inline block |
|         `block_id` | `string` | Yes | Block ID **Example value**: "doxcnPFi0R56ctbvh2MjkkROFWf" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. **Example value**: ["1660030311959965796"] |
|       `equation` | `equation` | No | Equation |
|         `content` | `string` | Yes | Equation content that conforms to KaTeX syntax, please refer to the syntax rules: https://katex.org/docs/supported.html **Example value**: "E=mc^2\n" |
|         `text_element_style` | `text_element_style` | No | Text element style |
|           `bold` | `boolean` | No | Bold **Example value**: true **Default value**: `false` |
|           `italic` | `boolean` | No | Italic **Example value**: true **Default value**: `false` |
|           `strikethrough` | `boolean` | No | Strikethrough **Example value**: true **Default value**: `false` |
|           `underline` | `boolean` | No | Underline **Example value**: true **Default value**: `false` |
|           `inline_code` | `boolean` | No | Inline code **Example value**: true **Default value**: `false` |
|           `background_color` | `int` | No | Background color **Example value**: 1 **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|           `text_color` | `int` | No | Font color **Example value**: 1 **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|           `link` | `link` | No | Link |
|             `url` | `string` | Yes | The url pointed to by hyperlinke (requires url_encode) **Example value**: "https%3A%2F%2Fopen.larksuite.com%2F" |
|           `comment_ids` | `string[]` | No | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. **Example value**: ["1660030311959965796"] |
|     `style` | `text_style` | Yes | Updated text style |
|       `align` | `int` | No | Alignment **Example value**: 1 **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  **Default value**: `1` |
|       `done` | `boolean` | No | Todo completion status **Example value**: true **Default value**: `false` |
|       `folded` | `boolean` | No | Folded state of text **Example value**: true **Default value**: `false` |
|       `language` | `int` | No | Code block language **Example value**: 1 **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|       `wrap` | `boolean` | No | Whether the code block wraps automatically **Example value**: true **Default value**: `false` |
|     `fields` | `int[]` | Yes | Fields in the text style that should be updated must specify at least one field. For example, to adjust Block alignment, set fields to [1]. **Example value**: [1] **Optional values are**:  - `1`: Modify the alignment of blocks - `2`: Modify the completion status of the todo block - `3`: Modify the collapsed state of the block - `4`: Modify the language type of the code block - `5`: Modify the collapsed state of the code block  | ### Request body example

{
    "requests": [
        {
            "block_id": "doxcnk0i44OMOaouw8AdXuXrp6b",
            "merge_table_cells": {
                "column_end_index": 2,
                "column_start_index": 0,
                "row_end_index": 1,
                "row_start_index": 0
            }
        },
        {
            "block_id": "doxcn0K8iGSMW4Mqgs9qlyTP50d",
            "update_text_elements": {
                "elements": [
                    {
                        "text_run": {
                            "content": "Hello",
                            "text_element_style": {
                                "background_color": 2,
                                "bold": true,
                                "italic": true,
                                "strikethrough": true,
                                "text_color": 2,
                                "underline": true
                            }
                        }
                    },
                    {
                        "text_run": {
                            "content": "World ",
                            "text_element_style": {
                                "italic": true
                            }
                        }
                    },
                    {
                        "mention_doc": {
                            "obj_type": 22,
                            "token": "doxcnAJ9VRRJqVMYZ1MyKnayXWe",
                            "url": "https%3A%2F%2Fexample.larksuite.com%2Fdocx%2FdoxcnAJ9VRRJqVMYZ1MyKnayXWe"
                        }
                    }
                ]
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
|   `blocks` | `block[]` | Batch update response blocks |
|     `block_id` | `string` | Block unique identification |
|     `parent_id` | `string` | Block parent id |
|     `children` | `string[]` | List of blocked child IDs |
|     `block_type` | `int` | Block Type **Optional values are**:  - `1`: Page Block - `2`: Text Block - `3`: Level 1 Heading Block - `4`: Level 2 Heading Block - `5`: Level 3 Heading Block - `6`: Level 4 Heading Block - `7`: Level 5 Heading Block - `8`: Level 6 Heading Block - `9`: Level 7 Heading Block - `10`: Level 8 Heading Block - `11`: Level 9 Heading Block - `12`: Bullet List Block - `13`: Ordered List Block - `14`: Code Block - `15`: Quote Block - `17`: Todo Block - `18`: Bitable Block - `19`: Callout Block - `20`: Chat Card Block - `21`: UML Diagram Block - `22`: Divider Block - `23`: File Block - `24`: Grid Block - `25`: Grid Column Block - `26`: Iframe Block - `27`: Image Block - `28`: ISV Block - `29`: Mindnote Block - `30`: Sheet Block - `31`: Table Block - `32`: Table Cell Block - `33`: View Block - `34`: Quote Container Block - `35`: Task Block - `36`: OKR Block - `37`: OKR Objective Block - `38`: OKR Key Result Block - `39`: OKR Progress Block - `40`: Add-Ons Block - `41`: Jira Issue Block - `42`: Wiki Catalog Block - `999`: Blocks not supported  |
|     `page` | `text` | Document Block |
|       `style` | `text_style` | Text style |
|         `align` | `int` | Alignment **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  |
|         `done` | `boolean` | Todo completion status |
|         `folded` | `boolean` | Folded state of text |
|         `language` | `int` | Code block language **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|         `wrap` | `boolean` | Whether the code block wraps automatically |
|       `elements` | `text_element[]` | Text element |
|         `text_run` | `text_run` | Text |
|           `content` | `string` | Text content |
|           `text_element_style` | `text_element_style` | Text local style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `mention_user` | `mention_user` | @User |
|           `user_id` | `string` | User OpenID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `mention_doc` | `mention_doc` | @Document |
|           `token` | `string` | Cloud document token |
|           `obj_type` | `int` | Cloud Document Type **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|           `url` | `string` | Docs link (url_encode) |
|           `title` | `string` | The title of the referenced document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `reminder` | `reminder` | Date Reminder |
|           `create_user_id` | `string` | Creator user ID |
|           `is_notify` | `boolean` | Whether to notify |
|           `is_whole_day` | `boolean` | Is it the date or the hour? |
|           `expire_time` | `string` | Time the event occurred (millisecond timestamp) |
|           `notify_time` | `string` | Time to trigger notification (millisecond timestamp) |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `file` | `inline_file` | inline attachment |
|           `file_token` | `string` | Attachment token |
|           `source_block_id` | `string` | The id of the block in which the attachment is located in the current document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `undefined` | `undefined_element` | Unsupported TextElements |
|         `inline_block` | `inline_block` | Inline block |
|           `block_id` | `string` | Block ID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `equation` | `equation` | Equation |
|           `content` | `string` | Equation content that conforms to KaTeX syntax, please refer to the syntax rules: https://katex.org/docs/supported.html |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|     `text` | `text` | Text Block |
|       `style` | `text_style` | Text style |
|         `align` | `int` | Alignment **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  |
|         `done` | `boolean` | Todo completion status |
|         `folded` | `boolean` | Folded state of text |
|         `language` | `int` | Code block language **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|         `wrap` | `boolean` | Whether the code block wraps automatically |
|       `elements` | `text_element[]` | Text element |
|         `text_run` | `text_run` | Text |
|           `content` | `string` | Text content |
|           `text_element_style` | `text_element_style` | Text local style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `mention_user` | `mention_user` | @User |
|           `user_id` | `string` | User OpenID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `mention_doc` | `mention_doc` | @Document |
|           `token` | `string` | Cloud document token |
|           `obj_type` | `int` | Cloud Document Type **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|           `url` | `string` | Docs link (url_encode) |
|           `title` | `string` | The title of the referenced document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `reminder` | `reminder` | Date Reminder |
|           `create_user_id` | `string` | Creator user ID |
|           `is_notify` | `boolean` | Whether to notify |
|           `is_whole_day` | `boolean` | Is it the date or the hour? |
|           `expire_time` | `string` | Time the event occurred (millisecond timestamp) |
|           `notify_time` | `string` | Time to trigger notification (millisecond timestamp) |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `file` | `inline_file` | inline attachment |
|           `file_token` | `string` | Attachment token |
|           `source_block_id` | `string` | The id of the block in which the attachment is located in the current document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `undefined` | `undefined_element` | Unsupported TextElements |
|         `inline_block` | `inline_block` | Inline block |
|           `block_id` | `string` | Block ID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `equation` | `equation` | Equation |
|           `content` | `string` | Equation content that conforms to KaTeX syntax, please refer to the syntax rules: https://katex.org/docs/supported.html |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|     `heading1` | `text` | Level 1 Heading Block |
|       `style` | `text_style` | Text style |
|         `align` | `int` | Alignment **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  |
|         `done` | `boolean` | Todo completion status |
|         `folded` | `boolean` | Folded state of text |
|         `language` | `int` | Code block language **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|         `wrap` | `boolean` | Whether the code block wraps automatically |
|       `elements` | `text_element[]` | Text element |
|         `text_run` | `text_run` | Text |
|           `content` | `string` | Text content |
|           `text_element_style` | `text_element_style` | Text local style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `mention_user` | `mention_user` | @User |
|           `user_id` | `string` | User OpenID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `mention_doc` | `mention_doc` | @Document |
|           `token` | `string` | Cloud document token |
|           `obj_type` | `int` | Cloud Document Type **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|           `url` | `string` | Docs link (url_encode) |
|           `title` | `string` | The title of the referenced document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `reminder` | `reminder` | Date Reminder |
|           `create_user_id` | `string` | Creator user ID |
|           `is_notify` | `boolean` | Whether to notify |
|           `is_whole_day` | `boolean` | Is it the date or the hour? |
|           `expire_time` | `string` | Time the event occurred (millisecond timestamp) |
|           `notify_time` | `string` | Time to trigger notification (millisecond timestamp) |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `file` | `inline_file` | inline attachment |
|           `file_token` | `string` | Attachment token |
|           `source_block_id` | `string` | The id of the block in which the attachment is located in the current document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `undefined` | `undefined_element` | Unsupported TextElements |
|         `inline_block` | `inline_block` | Inline block |
|           `block_id` | `string` | Block ID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `equation` | `equation` | Equation |
|           `content` | `string` | Equation content that conforms to KaTeX syntax, please refer to the syntax rules: https://katex.org/docs/supported.html |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|     `heading2` | `text` | Level 2 Heading Block |
|       `style` | `text_style` | Text style |
|         `align` | `int` | Alignment **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  |
|         `done` | `boolean` | Todo completion status |
|         `folded` | `boolean` | Folded state of text |
|         `language` | `int` | Code block language **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|         `wrap` | `boolean` | Whether the code block wraps automatically |
|       `elements` | `text_element[]` | Text element |
|         `text_run` | `text_run` | Text |
|           `content` | `string` | Text content |
|           `text_element_style` | `text_element_style` | Text local style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `mention_user` | `mention_user` | @User |
|           `user_id` | `string` | User OpenID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `mention_doc` | `mention_doc` | @Document |
|           `token` | `string` | Cloud document token |
|           `obj_type` | `int` | Cloud Document Type **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|           `url` | `string` | Docs link (url_encode) |
|           `title` | `string` | The title of the referenced document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `reminder` | `reminder` | Date Reminder |
|           `create_user_id` | `string` | Creator user ID |
|           `is_notify` | `boolean` | Whether to notify |
|           `is_whole_day` | `boolean` | Is it the date or the hour? |
|           `expire_time` | `string` | Time the event occurred (millisecond timestamp) |
|           `notify_time` | `string` | Time to trigger notification (millisecond timestamp) |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `file` | `inline_file` | inline attachment |
|           `file_token` | `string` | Attachment token |
|           `source_block_id` | `string` | The id of the block in which the attachment is located in the current document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `undefined` | `undefined_element` | Unsupported TextElements |
|         `inline_block` | `inline_block` | Inline block |
|           `block_id` | `string` | Block ID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `equation` | `equation` | Equation |
|           `content` | `string` | Equation content that conforms to KaTeX syntax, please refer to the syntax rules: https://katex.org/docs/supported.html |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|     `heading3` | `text` | Level 3 Heading Block |
|       `style` | `text_style` | Text style |
|         `align` | `int` | Alignment **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  |
|         `done` | `boolean` | Todo completion status |
|         `folded` | `boolean` | Folded state of text |
|         `language` | `int` | Code block language **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|         `wrap` | `boolean` | Whether the code block wraps automatically |
|       `elements` | `text_element[]` | Text element |
|         `text_run` | `text_run` | Text |
|           `content` | `string` | Text content |
|           `text_element_style` | `text_element_style` | Text local style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `mention_user` | `mention_user` | @User |
|           `user_id` | `string` | User OpenID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `mention_doc` | `mention_doc` | @Document |
|           `token` | `string` | Cloud document token |
|           `obj_type` | `int` | Cloud Document Type **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|           `url` | `string` | Docs link (url_encode) |
|           `title` | `string` | The title of the referenced document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `reminder` | `reminder` | Date Reminder |
|           `create_user_id` | `string` | Creator user ID |
|           `is_notify` | `boolean` | Whether to notify |
|           `is_whole_day` | `boolean` | Is it the date or the hour? |
|           `expire_time` | `string` | Time the event occurred (millisecond timestamp) |
|           `notify_time` | `string` | Time to trigger notification (millisecond timestamp) |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `file` | `inline_file` | inline attachment |
|           `file_token` | `string` | Attachment token |
|           `source_block_id` | `string` | The id of the block in which the attachment is located in the current document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `undefined` | `undefined_element` | Unsupported TextElements |
|         `inline_block` | `inline_block` | Inline block |
|           `block_id` | `string` | Block ID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `equation` | `equation` | Equation |
|           `content` | `string` | Equation content that conforms to KaTeX syntax, please refer to the syntax rules: https://katex.org/docs/supported.html |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|     `heading4` | `text` | Level 4 Heading Block |
|       `style` | `text_style` | Text style |
|         `align` | `int` | Alignment **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  |
|         `done` | `boolean` | Todo completion status |
|         `folded` | `boolean` | Folded state of text |
|         `language` | `int` | Code block language **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|         `wrap` | `boolean` | Whether the code block wraps automatically |
|       `elements` | `text_element[]` | Text element |
|         `text_run` | `text_run` | Text |
|           `content` | `string` | Text content |
|           `text_element_style` | `text_element_style` | Text local style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `mention_user` | `mention_user` | @User |
|           `user_id` | `string` | User OpenID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `mention_doc` | `mention_doc` | @Document |
|           `token` | `string` | Cloud document token |
|           `obj_type` | `int` | Cloud Document Type **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|           `url` | `string` | Docs link (url_encode) |
|           `title` | `string` | The title of the referenced document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `reminder` | `reminder` | Date Reminder |
|           `create_user_id` | `string` | Creator user ID |
|           `is_notify` | `boolean` | Whether to notify |
|           `is_whole_day` | `boolean` | Is it the date or the hour? |
|           `expire_time` | `string` | Time the event occurred (millisecond timestamp) |
|           `notify_time` | `string` | Time to trigger notification (millisecond timestamp) |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `file` | `inline_file` | inline attachment |
|           `file_token` | `string` | Attachment token |
|           `source_block_id` | `string` | The id of the block in which the attachment is located in the current document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `undefined` | `undefined_element` | Unsupported TextElements |
|         `inline_block` | `inline_block` | Inline block |
|           `block_id` | `string` | Block ID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `equation` | `equation` | Equation |
|           `content` | `string` | Equation content that conforms to KaTeX syntax, please refer to the syntax rules: https://katex.org/docs/supported.html |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|     `heading5` | `text` | Level 5 Heading Block |
|       `style` | `text_style` | Text style |
|         `align` | `int` | Alignment **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  |
|         `done` | `boolean` | Todo completion status |
|         `folded` | `boolean` | Folded state of text |
|         `language` | `int` | Code block language **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|         `wrap` | `boolean` | Whether the code block wraps automatically |
|       `elements` | `text_element[]` | Text element |
|         `text_run` | `text_run` | Text |
|           `content` | `string` | Text content |
|           `text_element_style` | `text_element_style` | Text local style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `mention_user` | `mention_user` | @User |
|           `user_id` | `string` | User OpenID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `mention_doc` | `mention_doc` | @Document |
|           `token` | `string` | Cloud document token |
|           `obj_type` | `int` | Cloud Document Type **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|           `url` | `string` | Docs link (url_encode) |
|           `title` | `string` | The title of the referenced document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `reminder` | `reminder` | Date Reminder |
|           `create_user_id` | `string` | Creator user ID |
|           `is_notify` | `boolean` | Whether to notify |
|           `is_whole_day` | `boolean` | Is it the date or the hour? |
|           `expire_time` | `string` | Time the event occurred (millisecond timestamp) |
|           `notify_time` | `string` | Time to trigger notification (millisecond timestamp) |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `file` | `inline_file` | inline attachment |
|           `file_token` | `string` | Attachment token |
|           `source_block_id` | `string` | The id of the block in which the attachment is located in the current document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `undefined` | `undefined_element` | Unsupported TextElements |
|         `inline_block` | `inline_block` | Inline block |
|           `block_id` | `string` | Block ID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `equation` | `equation` | Equation |
|           `content` | `string` | Equation content that conforms to KaTeX syntax, please refer to the syntax rules: https://katex.org/docs/supported.html |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|     `heading6` | `text` | Level 6 Heading Block |
|       `style` | `text_style` | Text style |
|         `align` | `int` | Alignment **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  |
|         `done` | `boolean` | Todo completion status |
|         `folded` | `boolean` | Folded state of text |
|         `language` | `int` | Code block language **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|         `wrap` | `boolean` | Whether the code block wraps automatically |
|       `elements` | `text_element[]` | Text element |
|         `text_run` | `text_run` | Text |
|           `content` | `string` | Text content |
|           `text_element_style` | `text_element_style` | Text local style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `mention_user` | `mention_user` | @User |
|           `user_id` | `string` | User OpenID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `mention_doc` | `mention_doc` | @Document |
|           `token` | `string` | Cloud document token |
|           `obj_type` | `int` | Cloud Document Type **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|           `url` | `string` | Docs link (url_encode) |
|           `title` | `string` | The title of the referenced document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `reminder` | `reminder` | Date Reminder |
|           `create_user_id` | `string` | Creator user ID |
|           `is_notify` | `boolean` | Whether to notify |
|           `is_whole_day` | `boolean` | Is it the date or the hour? |
|           `expire_time` | `string` | Time the event occurred (millisecond timestamp) |
|           `notify_time` | `string` | Time to trigger notification (millisecond timestamp) |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `file` | `inline_file` | inline attachment |
|           `file_token` | `string` | Attachment token |
|           `source_block_id` | `string` | The id of the block in which the attachment is located in the current document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `undefined` | `undefined_element` | Unsupported TextElements |
|         `inline_block` | `inline_block` | Inline block |
|           `block_id` | `string` | Block ID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `equation` | `equation` | Equation |
|           `content` | `string` | Equation content that conforms to KaTeX syntax, please refer to the syntax rules: https://katex.org/docs/supported.html |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|     `heading7` | `text` | Level 7 Heading Block |
|       `style` | `text_style` | Text style |
|         `align` | `int` | Alignment **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  |
|         `done` | `boolean` | Todo completion status |
|         `folded` | `boolean` | Folded state of text |
|         `language` | `int` | Code block language **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|         `wrap` | `boolean` | Whether the code block wraps automatically |
|       `elements` | `text_element[]` | Text element |
|         `text_run` | `text_run` | Text |
|           `content` | `string` | Text content |
|           `text_element_style` | `text_element_style` | Text local style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `mention_user` | `mention_user` | @User |
|           `user_id` | `string` | User OpenID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `mention_doc` | `mention_doc` | @Document |
|           `token` | `string` | Cloud document token |
|           `obj_type` | `int` | Cloud Document Type **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|           `url` | `string` | Docs link (url_encode) |
|           `title` | `string` | The title of the referenced document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `reminder` | `reminder` | Date Reminder |
|           `create_user_id` | `string` | Creator user ID |
|           `is_notify` | `boolean` | Whether to notify |
|           `is_whole_day` | `boolean` | Is it the date or the hour? |
|           `expire_time` | `string` | Time the event occurred (millisecond timestamp) |
|           `notify_time` | `string` | Time to trigger notification (millisecond timestamp) |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `file` | `inline_file` | inline attachment |
|           `file_token` | `string` | Attachment token |
|           `source_block_id` | `string` | The id of the block in which the attachment is located in the current document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `undefined` | `undefined_element` | Unsupported TextElements |
|         `inline_block` | `inline_block` | Inline block |
|           `block_id` | `string` | Block ID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `equation` | `equation` | Equation |
|           `content` | `string` | Equation content that conforms to KaTeX syntax, please refer to the syntax rules: https://katex.org/docs/supported.html |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|     `heading8` | `text` | Level 8 Heading Block |
|       `style` | `text_style` | Text style |
|         `align` | `int` | Alignment **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  |
|         `done` | `boolean` | Todo completion status |
|         `folded` | `boolean` | Folded state of text |
|         `language` | `int` | Code block language **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|         `wrap` | `boolean` | Whether the code block wraps automatically |
|       `elements` | `text_element[]` | Text element |
|         `text_run` | `text_run` | Text |
|           `content` | `string` | Text content |
|           `text_element_style` | `text_element_style` | Text local style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `mention_user` | `mention_user` | @User |
|           `user_id` | `string` | User OpenID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `mention_doc` | `mention_doc` | @Document |
|           `token` | `string` | Cloud document token |
|           `obj_type` | `int` | Cloud Document Type **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|           `url` | `string` | Docs link (url_encode) |
|           `title` | `string` | The title of the referenced document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `reminder` | `reminder` | Date Reminder |
|           `create_user_id` | `string` | Creator user ID |
|           `is_notify` | `boolean` | Whether to notify |
|           `is_whole_day` | `boolean` | Is it the date or the hour? |
|           `expire_time` | `string` | Time the event occurred (millisecond timestamp) |
|           `notify_time` | `string` | Time to trigger notification (millisecond timestamp) |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `file` | `inline_file` | inline attachment |
|           `file_token` | `string` | Attachment token |
|           `source_block_id` | `string` | The id of the block in which the attachment is located in the current document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `undefined` | `undefined_element` | Unsupported TextElements |
|         `inline_block` | `inline_block` | Inline block |
|           `block_id` | `string` | Block ID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `equation` | `equation` | Equation |
|           `content` | `string` | Equation content that conforms to KaTeX syntax, please refer to the syntax rules: https://katex.org/docs/supported.html |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|     `heading9` | `text` | Level 9 Heading Block |
|       `style` | `text_style` | Text style |
|         `align` | `int` | Alignment **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  |
|         `done` | `boolean` | Todo completion status |
|         `folded` | `boolean` | Folded state of text |
|         `language` | `int` | Code block language **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|         `wrap` | `boolean` | Whether the code block wraps automatically |
|       `elements` | `text_element[]` | Text element |
|         `text_run` | `text_run` | Text |
|           `content` | `string` | Text content |
|           `text_element_style` | `text_element_style` | Text local style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `mention_user` | `mention_user` | @User |
|           `user_id` | `string` | User OpenID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `mention_doc` | `mention_doc` | @Document |
|           `token` | `string` | Cloud document token |
|           `obj_type` | `int` | Cloud Document Type **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|           `url` | `string` | Docs link (url_encode) |
|           `title` | `string` | The title of the referenced document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `reminder` | `reminder` | Date Reminder |
|           `create_user_id` | `string` | Creator user ID |
|           `is_notify` | `boolean` | Whether to notify |
|           `is_whole_day` | `boolean` | Is it the date or the hour? |
|           `expire_time` | `string` | Time the event occurred (millisecond timestamp) |
|           `notify_time` | `string` | Time to trigger notification (millisecond timestamp) |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `file` | `inline_file` | inline attachment |
|           `file_token` | `string` | Attachment token |
|           `source_block_id` | `string` | The id of the block in which the attachment is located in the current document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `undefined` | `undefined_element` | Unsupported TextElements |
|         `inline_block` | `inline_block` | Inline block |
|           `block_id` | `string` | Block ID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `equation` | `equation` | Equation |
|           `content` | `string` | Equation content that conforms to KaTeX syntax, please refer to the syntax rules: https://katex.org/docs/supported.html |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|     `bullet` | `text` | Bullet List Block |
|       `style` | `text_style` | Text style |
|         `align` | `int` | Alignment **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  |
|         `done` | `boolean` | Todo completion status |
|         `folded` | `boolean` | Folded state of text |
|         `language` | `int` | Code block language **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|         `wrap` | `boolean` | Whether the code block wraps automatically |
|       `elements` | `text_element[]` | Text element |
|         `text_run` | `text_run` | Text |
|           `content` | `string` | Text content |
|           `text_element_style` | `text_element_style` | Text local style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `mention_user` | `mention_user` | @User |
|           `user_id` | `string` | User OpenID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `mention_doc` | `mention_doc` | @Document |
|           `token` | `string` | Cloud document token |
|           `obj_type` | `int` | Cloud Document Type **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|           `url` | `string` | Docs link (url_encode) |
|           `title` | `string` | The title of the referenced document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `reminder` | `reminder` | Date Reminder |
|           `create_user_id` | `string` | Creator user ID |
|           `is_notify` | `boolean` | Whether to notify |
|           `is_whole_day` | `boolean` | Is it the date or the hour? |
|           `expire_time` | `string` | Time the event occurred (millisecond timestamp) |
|           `notify_time` | `string` | Time to trigger notification (millisecond timestamp) |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `file` | `inline_file` | inline attachment |
|           `file_token` | `string` | Attachment token |
|           `source_block_id` | `string` | The id of the block in which the attachment is located in the current document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `undefined` | `undefined_element` | Unsupported TextElements |
|         `inline_block` | `inline_block` | Inline block |
|           `block_id` | `string` | Block ID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `equation` | `equation` | Equation |
|           `content` | `string` | Equation content that conforms to KaTeX syntax, please refer to the syntax rules: https://katex.org/docs/supported.html |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|     `ordered` | `text` | Ordered List Block |
|       `style` | `text_style` | Text style |
|         `align` | `int` | Alignment **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  |
|         `done` | `boolean` | Todo completion status |
|         `folded` | `boolean` | Folded state of text |
|         `language` | `int` | Code block language **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|         `wrap` | `boolean` | Whether the code block wraps automatically |
|       `elements` | `text_element[]` | Text element |
|         `text_run` | `text_run` | Text |
|           `content` | `string` | Text content |
|           `text_element_style` | `text_element_style` | Text local style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `mention_user` | `mention_user` | @User |
|           `user_id` | `string` | User OpenID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `mention_doc` | `mention_doc` | @Document |
|           `token` | `string` | Cloud document token |
|           `obj_type` | `int` | Cloud Document Type **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|           `url` | `string` | Docs link (url_encode) |
|           `title` | `string` | The title of the referenced document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `reminder` | `reminder` | Date Reminder |
|           `create_user_id` | `string` | Creator user ID |
|           `is_notify` | `boolean` | Whether to notify |
|           `is_whole_day` | `boolean` | Is it the date or the hour? |
|           `expire_time` | `string` | Time the event occurred (millisecond timestamp) |
|           `notify_time` | `string` | Time to trigger notification (millisecond timestamp) |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `file` | `inline_file` | inline attachment |
|           `file_token` | `string` | Attachment token |
|           `source_block_id` | `string` | The id of the block in which the attachment is located in the current document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `undefined` | `undefined_element` | Unsupported TextElements |
|         `inline_block` | `inline_block` | Inline block |
|           `block_id` | `string` | Block ID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `equation` | `equation` | Equation |
|           `content` | `string` | Equation content that conforms to KaTeX syntax, please refer to the syntax rules: https://katex.org/docs/supported.html |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|     `code` | `text` | Code Block |
|       `style` | `text_style` | Text style |
|         `align` | `int` | Alignment **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  |
|         `done` | `boolean` | Todo completion status |
|         `folded` | `boolean` | Folded state of text |
|         `language` | `int` | Code block language **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|         `wrap` | `boolean` | Whether the code block wraps automatically |
|       `elements` | `text_element[]` | Text element |
|         `text_run` | `text_run` | Text |
|           `content` | `string` | Text content |
|           `text_element_style` | `text_element_style` | Text local style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `mention_user` | `mention_user` | @User |
|           `user_id` | `string` | User OpenID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `mention_doc` | `mention_doc` | @Document |
|           `token` | `string` | Cloud document token |
|           `obj_type` | `int` | Cloud Document Type **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|           `url` | `string` | Docs link (url_encode) |
|           `title` | `string` | The title of the referenced document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `reminder` | `reminder` | Date Reminder |
|           `create_user_id` | `string` | Creator user ID |
|           `is_notify` | `boolean` | Whether to notify |
|           `is_whole_day` | `boolean` | Is it the date or the hour? |
|           `expire_time` | `string` | Time the event occurred (millisecond timestamp) |
|           `notify_time` | `string` | Time to trigger notification (millisecond timestamp) |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `file` | `inline_file` | inline attachment |
|           `file_token` | `string` | Attachment token |
|           `source_block_id` | `string` | The id of the block in which the attachment is located in the current document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `undefined` | `undefined_element` | Unsupported TextElements |
|         `inline_block` | `inline_block` | Inline block |
|           `block_id` | `string` | Block ID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `equation` | `equation` | Equation |
|           `content` | `string` | Equation content that conforms to KaTeX syntax, please refer to the syntax rules: https://katex.org/docs/supported.html |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|     `quote` | `text` | Quote Block |
|       `style` | `text_style` | Text style |
|         `align` | `int` | Alignment **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  |
|         `done` | `boolean` | Todo completion status |
|         `folded` | `boolean` | Folded state of text |
|         `language` | `int` | Code block language **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|         `wrap` | `boolean` | Whether the code block wraps automatically |
|       `elements` | `text_element[]` | Text element |
|         `text_run` | `text_run` | Text |
|           `content` | `string` | Text content |
|           `text_element_style` | `text_element_style` | Text local style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `mention_user` | `mention_user` | @User |
|           `user_id` | `string` | User OpenID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `mention_doc` | `mention_doc` | @Document |
|           `token` | `string` | Cloud document token |
|           `obj_type` | `int` | Cloud Document Type **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|           `url` | `string` | Docs link (url_encode) |
|           `title` | `string` | The title of the referenced document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `reminder` | `reminder` | Date Reminder |
|           `create_user_id` | `string` | Creator user ID |
|           `is_notify` | `boolean` | Whether to notify |
|           `is_whole_day` | `boolean` | Is it the date or the hour? |
|           `expire_time` | `string` | Time the event occurred (millisecond timestamp) |
|           `notify_time` | `string` | Time to trigger notification (millisecond timestamp) |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `file` | `inline_file` | inline attachment |
|           `file_token` | `string` | Attachment token |
|           `source_block_id` | `string` | The id of the block in which the attachment is located in the current document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `undefined` | `undefined_element` | Unsupported TextElements |
|         `inline_block` | `inline_block` | Inline block |
|           `block_id` | `string` | Block ID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `equation` | `equation` | Equation |
|           `content` | `string` | Equation content that conforms to KaTeX syntax, please refer to the syntax rules: https://katex.org/docs/supported.html |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|     `equation` | `text` | Equation Block |
|       `style` | `text_style` | Text style |
|         `align` | `int` | Alignment **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  |
|         `done` | `boolean` | Todo completion status |
|         `folded` | `boolean` | Folded state of text |
|         `language` | `int` | Code block language **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|         `wrap` | `boolean` | Whether the code block wraps automatically |
|       `elements` | `text_element[]` | Text element |
|         `text_run` | `text_run` | Text |
|           `content` | `string` | Text content |
|           `text_element_style` | `text_element_style` | Text local style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `mention_user` | `mention_user` | @User |
|           `user_id` | `string` | User OpenID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `mention_doc` | `mention_doc` | @Document |
|           `token` | `string` | Cloud document token |
|           `obj_type` | `int` | Cloud Document Type **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|           `url` | `string` | Docs link (url_encode) |
|           `title` | `string` | The title of the referenced document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `reminder` | `reminder` | Date Reminder |
|           `create_user_id` | `string` | Creator user ID |
|           `is_notify` | `boolean` | Whether to notify |
|           `is_whole_day` | `boolean` | Is it the date or the hour? |
|           `expire_time` | `string` | Time the event occurred (millisecond timestamp) |
|           `notify_time` | `string` | Time to trigger notification (millisecond timestamp) |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `file` | `inline_file` | inline attachment |
|           `file_token` | `string` | Attachment token |
|           `source_block_id` | `string` | The id of the block in which the attachment is located in the current document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `undefined` | `undefined_element` | Unsupported TextElements |
|         `inline_block` | `inline_block` | Inline block |
|           `block_id` | `string` | Block ID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `equation` | `equation` | Equation |
|           `content` | `string` | Equation content that conforms to KaTeX syntax, please refer to the syntax rules: https://katex.org/docs/supported.html |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|     `todo` | `text` | Todo Block |
|       `style` | `text_style` | Text style |
|         `align` | `int` | Alignment **Optional values are**:  - `1`: Left typesetting - `2`: Centered typesetting - `3`: Right typesetting  |
|         `done` | `boolean` | Todo completion status |
|         `folded` | `boolean` | Folded state of text |
|         `language` | `int` | Code block language **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|         `wrap` | `boolean` | Whether the code block wraps automatically |
|       `elements` | `text_element[]` | Text element |
|         `text_run` | `text_run` | Text |
|           `content` | `string` | Text content |
|           `text_element_style` | `text_element_style` | Text local style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `mention_user` | `mention_user` | @User |
|           `user_id` | `string` | User OpenID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `mention_doc` | `mention_doc` | @Document |
|           `token` | `string` | Cloud document token |
|           `obj_type` | `int` | Cloud Document Type **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|           `url` | `string` | Docs link (url_encode) |
|           `title` | `string` | The title of the referenced document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `reminder` | `reminder` | Date Reminder |
|           `create_user_id` | `string` | Creator user ID |
|           `is_notify` | `boolean` | Whether to notify |
|           `is_whole_day` | `boolean` | Is it the date or the hour? |
|           `expire_time` | `string` | Time the event occurred (millisecond timestamp) |
|           `notify_time` | `string` | Time to trigger notification (millisecond timestamp) |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `file` | `inline_file` | inline attachment |
|           `file_token` | `string` | Attachment token |
|           `source_block_id` | `string` | The id of the block in which the attachment is located in the current document |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `undefined` | `undefined_element` | Unsupported TextElements |
|         `inline_block` | `inline_block` | Inline block |
|           `block_id` | `string` | Block ID |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|         `equation` | `equation` | Equation |
|           `content` | `string` | Equation content that conforms to KaTeX syntax, please refer to the syntax rules: https://katex.org/docs/supported.html |
|           `text_element_style` | `text_element_style` | Text element style |
|             `bold` | `boolean` | Bold |
|             `italic` | `boolean` | Italic |
|             `strikethrough` | `boolean` | Strikethrough |
|             `underline` | `boolean` | Underline |
|             `inline_code` | `boolean` | Inline code |
|             `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|             `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|             `link` | `link` | Link |
|               `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|             `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|     `bitable` | `bitable` | Bitable Block |
|       `token` | `string` | Bitable Block |
|       `view_type` | `int` | View type **Optional values are**:  - `1`: Data sheet - `2`: Kanban  |
|     `callout` | `callout` | Callout Block |
|       `background_color` | `int` | Callout block background color **Optional values are**:  - `1`: Light red - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light neutral - `8`: Dark red - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark neutral - `15`: Dark slight gray  |
|       `border_color` | `int` | Border color **Optional values are**:  - `1`: Red - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Neutral  |
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
|         `iframe_type` | `int` | Iframe Type **Optional values are**:  - `1`: Bilibili - `2`: Xigua Video - `3`: Youku - `4`: Airtable - `5`: Baidu Map - `6`: Autonavi Map - `7`: TikTok - `8`: Figma - `9`: Ink knife - `10`: Canva - `11`: CodePen - `12`: Lark Survey - `13`: Gold data - `14`: Google Map - `15`: Youtube - `99`: Other  |
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
|           `done` | `boolean` | Todo completion status |
|           `folded` | `boolean` | Folded state of text |
|           `language` | `int` | Code block language **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|           `wrap` | `boolean` | Whether the code block wraps automatically |
|         `elements` | `text_element[]` | Text element |
|           `text_run` | `text_run` | Text |
|             `content` | `string` | Text content |
|             `text_element_style` | `text_element_style` | Text local style |
|               `bold` | `boolean` | Bold |
|               `italic` | `boolean` | Italic |
|               `strikethrough` | `boolean` | Strikethrough |
|               `underline` | `boolean` | Underline |
|               `inline_code` | `boolean` | Inline code |
|               `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|               `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|               `link` | `link` | Link |
|                 `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|               `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|           `mention_user` | `mention_user` | @User |
|             `user_id` | `string` | User OpenID |
|             `text_element_style` | `text_element_style` | Text element style |
|               `bold` | `boolean` | Bold |
|               `italic` | `boolean` | Italic |
|               `strikethrough` | `boolean` | Strikethrough |
|               `underline` | `boolean` | Underline |
|               `inline_code` | `boolean` | Inline code |
|               `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|               `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|               `link` | `link` | Link |
|                 `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|               `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|           `mention_doc` | `mention_doc` | @Document |
|             `token` | `string` | Cloud document token |
|             `obj_type` | `int` | Cloud Document Type **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|             `url` | `string` | Docs link (url_encode) |
|             `title` | `string` | The title of the referenced document |
|             `text_element_style` | `text_element_style` | Text element style |
|               `bold` | `boolean` | Bold |
|               `italic` | `boolean` | Italic |
|               `strikethrough` | `boolean` | Strikethrough |
|               `underline` | `boolean` | Underline |
|               `inline_code` | `boolean` | Inline code |
|               `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|               `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|               `link` | `link` | Link |
|                 `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|               `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|           `reminder` | `reminder` | Date Reminder |
|             `create_user_id` | `string` | Creator user ID |
|             `is_notify` | `boolean` | Whether to notify |
|             `is_whole_day` | `boolean` | Is it the date or the hour? |
|             `expire_time` | `string` | Time the event occurred (millisecond timestamp) |
|             `notify_time` | `string` | Time to trigger notification (millisecond timestamp) |
|             `text_element_style` | `text_element_style` | Text element style |
|               `bold` | `boolean` | Bold |
|               `italic` | `boolean` | Italic |
|               `strikethrough` | `boolean` | Strikethrough |
|               `underline` | `boolean` | Underline |
|               `inline_code` | `boolean` | Inline code |
|               `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|               `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|               `link` | `link` | Link |
|                 `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|               `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|           `file` | `inline_file` | inline attachment |
|             `file_token` | `string` | Attachment token |
|             `source_block_id` | `string` | The id of the block in which the attachment is located in the current document |
|             `text_element_style` | `text_element_style` | Text element style |
|               `bold` | `boolean` | Bold |
|               `italic` | `boolean` | Italic |
|               `strikethrough` | `boolean` | Strikethrough |
|               `underline` | `boolean` | Underline |
|               `inline_code` | `boolean` | Inline code |
|               `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|               `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|               `link` | `link` | Link |
|                 `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|               `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|           `undefined` | `undefined_element` | Unsupported TextElements |
|           `inline_block` | `inline_block` | Inline block |
|             `block_id` | `string` | Block ID |
|             `text_element_style` | `text_element_style` | Text element style |
|               `bold` | `boolean` | Bold |
|               `italic` | `boolean` | Italic |
|               `strikethrough` | `boolean` | Strikethrough |
|               `underline` | `boolean` | Underline |
|               `inline_code` | `boolean` | Inline code |
|               `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|               `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|               `link` | `link` | Link |
|                 `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|               `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|           `equation` | `equation` | Equation |
|             `content` | `string` | Equation content that conforms to KaTeX syntax, please refer to the syntax rules: https://katex.org/docs/supported.html |
|             `text_element_style` | `text_element_style` | Text element style |
|               `bold` | `boolean` | Bold |
|               `italic` | `boolean` | Italic |
|               `strikethrough` | `boolean` | Strikethrough |
|               `underline` | `boolean` | Underline |
|               `inline_code` | `boolean` | Inline code |
|               `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|               `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|               `link` | `link` | Link |
|                 `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|               `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
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
|           `done` | `boolean` | Todo completion status |
|           `folded` | `boolean` | Folded state of text |
|           `language` | `int` | Code block language **Optional values are**:  - `1`: PlainText - `2`: ABAP - `3`: Ada - `4`: Apache - `5`: Apex - `6`: Assembly Language - `7`: Bash - `8`: CSharp - `9`: C++ - `10`: C - `11`: COBOL - `12`: CSS - `13`: CoffeeScript - `14`: D - `15`: Dart - `16`: Delphi - `17`: Django - `18`: Dockerfile - `19`: Erlang - `20`: Fortran - `21`: FoxPro - `22`: Go - `23`: Groovy - `24`: HTML - `25`: HTMLBars - `26`: HTTP - `27`: Haskell - `28`: JSON - `29`: Java - `30`: JavaScript - `31`: Julia. - `32`: Kotlin - `33`: LateX - `34`: Lisp - `35`: Logo - `36`: Lua - `37`: MATLAB - `38`: Makefile - `39`: Markdown - `40`: Nginx - `41`: Objective-C - `42`: OpenEdgeABL - `43`: PHP - `44`: Perl - `45`: PostScript - `46`: Power Shell - `47`: Prolog - `48`: ProtoBuf - `49`: Python - `50`: R - `51`: RPG - `52`: Ruby - `53`: Rust - `54`: SAS - `55`: SCSS - `56`: SQL - `57`: Scala - `58`: Scheme - `59`: Scratch - `60`: Shell - `61`: Swift - `62`: Thrift - `63`: TypeScript - `64`: VBScript - `65`: Visual Basic - `66`: XML - `67`: YAML - `68`: CMake - `69`: Diff - `70`: Gherkin - `71`: GraphQL - `72`: OpenGL Shading Language - `73`: Properties - `74`: Solidity - `75`: TOML  |
|           `wrap` | `boolean` | Whether the code block wraps automatically |
|         `elements` | `text_element[]` | Text element |
|           `text_run` | `text_run` | Text |
|             `content` | `string` | Text content |
|             `text_element_style` | `text_element_style` | Text local style |
|               `bold` | `boolean` | Bold |
|               `italic` | `boolean` | Italic |
|               `strikethrough` | `boolean` | Strikethrough |
|               `underline` | `boolean` | Underline |
|               `inline_code` | `boolean` | Inline code |
|               `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|               `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|               `link` | `link` | Link |
|                 `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|               `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|           `mention_user` | `mention_user` | @User |
|             `user_id` | `string` | User OpenID |
|             `text_element_style` | `text_element_style` | Text element style |
|               `bold` | `boolean` | Bold |
|               `italic` | `boolean` | Italic |
|               `strikethrough` | `boolean` | Strikethrough |
|               `underline` | `boolean` | Underline |
|               `inline_code` | `boolean` | Inline code |
|               `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|               `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|               `link` | `link` | Link |
|                 `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|               `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|           `mention_doc` | `mention_doc` | @Document |
|             `token` | `string` | Cloud document token |
|             `obj_type` | `int` | Cloud Document Type **Optional values are**:  - `1`: Doc - `3`: Sheet - `8`: Bitable - `11`: MindNote - `12`: File - `15`: Slide - `16`: Wiki - `22`: Docx  |
|             `url` | `string` | Docs link (url_encode) |
|             `title` | `string` | The title of the referenced document |
|             `text_element_style` | `text_element_style` | Text element style |
|               `bold` | `boolean` | Bold |
|               `italic` | `boolean` | Italic |
|               `strikethrough` | `boolean` | Strikethrough |
|               `underline` | `boolean` | Underline |
|               `inline_code` | `boolean` | Inline code |
|               `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|               `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|               `link` | `link` | Link |
|                 `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|               `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|           `reminder` | `reminder` | Date Reminder |
|             `create_user_id` | `string` | Creator user ID |
|             `is_notify` | `boolean` | Whether to notify |
|             `is_whole_day` | `boolean` | Is it the date or the hour? |
|             `expire_time` | `string` | Time the event occurred (millisecond timestamp) |
|             `notify_time` | `string` | Time to trigger notification (millisecond timestamp) |
|             `text_element_style` | `text_element_style` | Text element style |
|               `bold` | `boolean` | Bold |
|               `italic` | `boolean` | Italic |
|               `strikethrough` | `boolean` | Strikethrough |
|               `underline` | `boolean` | Underline |
|               `inline_code` | `boolean` | Inline code |
|               `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|               `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|               `link` | `link` | Link |
|                 `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|               `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|           `file` | `inline_file` | inline attachment |
|             `file_token` | `string` | Attachment token |
|             `source_block_id` | `string` | The id of the block in which the attachment is located in the current document |
|             `text_element_style` | `text_element_style` | Text element style |
|               `bold` | `boolean` | Bold |
|               `italic` | `boolean` | Italic |
|               `strikethrough` | `boolean` | Strikethrough |
|               `underline` | `boolean` | Underline |
|               `inline_code` | `boolean` | Inline code |
|               `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|               `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|               `link` | `link` | Link |
|                 `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|               `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|           `undefined` | `undefined_element` | Unsupported TextElements |
|           `inline_block` | `inline_block` | Inline block |
|             `block_id` | `string` | Block ID |
|             `text_element_style` | `text_element_style` | Text element style |
|               `bold` | `boolean` | Bold |
|               `italic` | `boolean` | Italic |
|               `strikethrough` | `boolean` | Strikethrough |
|               `underline` | `boolean` | Underline |
|               `inline_code` | `boolean` | Inline code |
|               `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|               `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|               `link` | `link` | Link |
|                 `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|               `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|           `equation` | `equation` | Equation |
|             `content` | `string` | Equation content that conforms to KaTeX syntax, please refer to the syntax rules: https://katex.org/docs/supported.html |
|             `text_element_style` | `text_element_style` | Text element style |
|               `bold` | `boolean` | Bold |
|               `italic` | `boolean` | Italic |
|               `strikethrough` | `boolean` | Strikethrough |
|               `underline` | `boolean` | Underline |
|               `inline_code` | `boolean` | Inline code |
|               `background_color` | `int` | Background color **Optional values are**:  - `1`: Light pink - `2`: Light orange - `3`: Light yellow - `4`: Light green - `5`: Light blue - `6`: Light purple - `7`: Light grey - `8`: Dark pink - `9`: Dark orange - `10`: Dark yellow - `11`: Dark green - `12`: Dark blue - `13`: Dark purple - `14`: Dark gray - `15`: Dark slight gray  |
|               `text_color` | `int` | Font color **Optional values are**:  - `1`: Pink - `2`: Orange - `3`: Yellow - `4`: Green - `5`: Blue - `6`: Purple - `7`: Grey  |
|               `link` | `link` | Link |
|                 `url` | `string` | The url pointed to by hyperlinke (requires url_encode) |
|               `comment_ids` | `string[]` | Comments ID list. When creating a Block, incoming comments ID are not supported. When updating the Element of a text Block, it is allowed to move the existing comment ID of the corresponding version to any Element in the same Block, but does not support incoming new comments ID. For comment content, please see:「Get Comment」. |
|     `okr_progress` | `okr_progress` | OKR Progress Block |
|     `comment_ids` | `string[]` | Comment ids |
|     `jira_issue` | `jira_issue` | Jira Issue Block |
|       `id` | `string` | Jira Issue ID |
|       `key` | `string` | Jira Key |
|     `wiki_catalog` | `wiki_catalog` | Wiki Catalog Block |
|       `wiki_token` | `string` | wiki token |
|   `document_revision_id` | `int` | The version number of the document after the current update is successful |
|   `client_token` | `string` | The unique identifier of the operation, which is used in the update request to indicate that the update is idempotent | ### Response body example

{
    "code": 0,
    "data": {
        "blocks": [
            {
                "block_id": "doxcn0K8iGSMW4Mqgs9qlyTP50d",
                "block_type": 2,
                "parent_id": "doxcnAJ9VRRJqVMYZ1MyKnayXWe",
                "text": {
                    "elements": [
                        {
                            "text_run": {
                                "content": "Hello",
                                "text_element_style": {
                                    "background_color": 2,
                                    "bold": true,
                                    "inline_code": false,
                                    "italic": true,
                                    "strikethrough": true,
                                    "text_color": 2,
                                    "underline": true
                                }
                            }
                        },
                        {
                            "text_run": {
                                "content": "World ",
                                "text_element_style": {
                                    "bold": false,
                                    "inline_code": false,
                                    "italic": true,
                                    "strikethrough": false,
                                    "underline": false
                                }
                            }
                        },
                        {
                            "mention_doc": {
                                "obj_type": 22,
                                "title": "Demo",
                                "token": "doxcnAJ9VRRJqVMYZ1MyKnayXWe",
                                "url": "https%3A%2F%2Fexample.larksuite.com%2Fdocx%2FdoxcnAJ9VRRJqVMYZ1MyKnayXWe"
                            }
                        }
                    ],
                    "style": {
                        "done": false,
                        "folded": false,
                        "wrap": false
                    }
                }
            },
            {
                "block_id": "doxcnk0i44OMOaouw8AdXuXrp6b",
                "block_type": 31,
                "children": [
                    "doxcnO2UeYco4mu80sr6oRCiRpf",
                    "doxcnaAGMYMk6kcGeYXNfc1Rluc",
                    "doxcnCKuqMQOM0gAOYfysUgZD2d",
                    "doxcnMKg8Uk8wiAMIW8omQ06uoc"
                ],
                "parent_id": "doxcnAJ9VRRJqVMYZ1MyKnayXWe",
                "table": {
                    "cells": [
                        "doxcnO2UeYco4mu80sr6oRCiRpf",
                        "doxcnaAGMYMk6kcGeYXNfc1Rluc",
                        "doxcnCKuqMQOM0gAOYfysUgZD2d",
                        "doxcnMKg8Uk8wiAMIW8omQ06uoc"
                    ],
                    "property": {
                        "column_size": 2,
                        "column_width": [
                            100,
                            100
                        ],
                        "merge_info": [
                            {
                                "col_span": 2,
                                "row_span": 1
                            },
                            {
                                "col_span": 1,
                                "row_span": 1
                            },
                            {
                                "col_span": 1,
                                "row_span": 1
                            },
                            {
                                "col_span": 1,
                                "row_span": 1
                            }
                        ],
                        "row_size": 2
                    }
                }
            }
        ],
        "client_token": "e472907a-9ddc-4453-af28-22a6530b76b9",
        "document_revision_id": 387
    },
    "msg": ""
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1770001 | invalid param | Confirm whether the passed parameters are valid |
| 404 | 1770002 | not found | Confirm that the document has been deleted |
| 400 | 1770003 | resource deleted | Confirm that the resource has been deleted |
| 400 | 1770004 | too many blocks in document | Verify that the maximum number of Document Blocks is exceeded |
| 400 | 1770005 | too deep level in document | Confirm that the document Block level exceeds the upper limit |
| 400 | 1770006 | schema mismatch | Verify that the document structure is valid |
| 400 | 1770007 | too many children in block | Confirm that the maximum number of Children in the specified Block is exceeded |
| 400 | 1770008 | too big file size | Verify that the uploaded file size exceeds the upper limit |
| 400 | 1770010 | too many table column | Confirm whether the number of table columns exceeds the upper limit |
| 400 | 1770011 | too many table cell | Confirm that the number of table cells exceeds the upper limit |
| 400 | 1770012 | too many grid column | Confirm if the maximum number of Grid columns is exceeded |
| 400 | 1770013 | relation mismatch | Confirm whether the relationship between pictures, files and other resources is correct. Before inserting pictures and files, please upload the materials |
| 400 | 1770014 | parent children relation mismatch | Confirm that Block's father-son relationship is correct |
| 400 | 1770015 | single edit with multi document | Verify that the document to which the Block belongs is the same as the specified Document |
| 400 | 1770019 | repeated blockID in document | Verify Duplicate BlockIDs in Documents |
| 400 | 1770020 | operation denied on copying document | Confirm that the Document is creating a copy |
| 400 | 1770021 | too old document | Verify that the specified Document version is old |
| 400 | 1770022 | invalid page token | Confirm whether the page_token in the query parameter is valid |
| 400 | 1770024 | invalid operation | Confirm whether the operation is valid |
| 400 | 1770025 | operation and block not match | Confirm whether the corresponding operation of the specified Block application is valid |
| 400 | 1770026 | row operation over range | Confirm that row operation subscripts are out of bounds |
| 400 | 1770027 | column operation over range | Confirm that column operation subscripts are out of bounds |
| 400 | 1770028 | block not support create children | Confirm that adding Children to the specified Block is valid |
| 400 | 1770029 | block not support to create | Confirm that the specified Block supports creation |
| 400 | 1770030 | invalid parent children relation | Confirm whether the parent-child relationship of the specified operation is valid |
| 400 | 1770031 | block not support to delete children | Confirm that the specified block supports deleting Children |
| 400 | 1770033 | raw content size exceed limited | Raw content size exceed limited |
| 400 | 1770034 | operation count exceed limited | Too many cells are involved in the current request. Please split the request into multiple requests |
| 400 | 1770035 | Resource count exceeds limit | The number of resources in the current request exceeds the limit, please split it into multiple requests. The upper limit of various resources is: 200 ChatCard, 200 File, 200 MentionDoc, 200 MentionUser, 20 Image, 20 ISV, 5 Sheet and 5 Bitable. |
| 400 | 1770038 | resource not found | The inserted resource was not found or the resource did not have permission to insert, please check that the resource identifier is correct. |
| 403 | 1770032 | forbidden | 1. Confirm whether the operator has permissions for the document.  Click to learn about document permissions 2. Confirm whether the operator has view permission for the mentioned document. 3. Confirm whether the mentioned user is employed and is a contact with the current operator. 4. Confirm whether the operator has permission to view and share chat cards. 5. Confirm whether the operator has permission to access the specified Wiki subdirectory. 6. Confirm whether the operator has view permissions for OKR, ISV, and Add-Ons. |
| 500 | 1771001 | server internal error | Server internal error |
| 500 | 1771006 | mount folder failed | Failed to mount document to cloud space directory |
| 500 | 1771002 | gateway server internal error | Gateway service internal error |
| 500 | 1771003 | gateway marshal error | Gateway service analysis error |
| 500 | 1771004 | gateway unmarshal error | Gateway service back analysis error |
| 503 | 1771005 | system under maintenance | System service is under maintenance |
