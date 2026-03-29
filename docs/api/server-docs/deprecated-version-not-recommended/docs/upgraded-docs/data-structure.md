---
title: "Data Structure"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/document-docx/docx-structure"
section: "Deprecated Version (Not Recommended)"
updated: "1685011043000"
---

# Data Structure
# 1 Resources

| **Resource Name** | **Description** |
| ------------ | ------------------------------------------------------------ |
| Block | Documentation minimal building block. As a structural element of content, Block has various forms, which can be a piece of text, a table or a picture, etc. |
| Document | A new version of the document. The hierarchical content that can form a tree-like relationship between multiple blocks. From the structural point of view, the new version of the document is a block tree, which is called Document in OpenAPI.| # 2 Data Model

## 2.1 Documents

Document represents a new version of the document.
````json
"document": {
    "document_id": string,
    "revision_id": int,
    "title": string
}
````

| Name | Data Type | Properties | Default Value | Description |
| --------- | --------------- | :-: | :-: | --------- |
|`document_id` | `string` | required |`-`| **Document Unique ID**, and also the ID of the root Block of the document. |
|`revision_id` | `int` | required |`-`| **Document Revision ID**, used to specify the document version to query or update. If the interface is called multiple times and its return version ID has not changed, it means the document has not changed. On the contrary, a change in the version ID usually means that the document has been updated, but it should be noted that the ID change does not necessarily mean that the content of the document has changed. For example, the document may be commented by others. |
|`title` | `string` | required |`-`| **document title**, `Document` only supports returning plain text. | ## 2.2 Block
Block is the smallest building block of a document, which has various BlockData (content entities), such as `Text`, `Image`, and `Table`. For a specific Block, it has the following properties:
````json
{
    "block_id": string,
    "block_type": enum(BlockType),
    "parent_id": string,
    "children": [](string),
    "comment_ids": [](string),
    // BlockData can only be one of the following, and must correspond to BlockType:
    "page": object(Text),
    "text": object(Text),
    "heading1": object(Text),
    "heading2": object(Text),
    "heading3": object(Text),
    "heading4": object(Text),
    "heading5": object(Text),
    "heading6": object(Text),
    "heading7": object(Text),
    "heading8": object(Text),
    "heading9": object(Text),
    "bullet": object(Text),
    "ordered": object(Text),
    "code": object(Text),
    "quote": object(Text),
    "todo": object(Text),
    "bitable": object(Bitable),
    "callout": object(Callout),
    "chat_card": object(ChatCard),
    "diagram": object(Diagram),
    "divider": object(Divider),
    "file": object(File),
    "grid": object(Grid),
    "grid_column": object(GridColumn),
    "iframe": object(Iframe),
    "image": object(Image),
    "isv": object(ISV),
    "mindnote": object(Mindnote),
    "sheet": object(Sheet),
    "table": object(Table),
    "table_cell": object(TableCell),
    "view": object(View),
    "undefined": object(Undefined),
    "quote_container": object(QuoteContainer),
    "task": object(Task),
    "okr": object(OKR),
    "okr_objective": object(OkrObjective),
    "okr_key_result": object(OkrKeyResult),
    "okr_progress": object(OkrProgress)
}
````

| Name | Data Type | Properties | Default Value | Description |
| --------- | --------------- | ------- | ----------- | --------- |
|`block_id` | `string` | required | `-` | Block globally unique identifier, which will be automatically generated when a Block is created |
|`block_type` | `enum(BlockType)` | required | `-` | Block type, enum type |
|`parent_id` | `string` | optional | `-` | Parent BlockID, except the root node `PageBlock`, other blocks have ParentBlock |
|`children` | `array(string)` | optional | `-` | List of child BlockIDs |
| `comment_ids` | `array(string)` | optional | `-` | List of comment IDs. For comment content, please see:「Get Comment」。 |
| `page` | `object(Text)` | optional | `-` | Page Block |
| `text` | `object(Text)` | optional | `-` | Text Block |
| `heading1` | `object(Text)` | optional | `-` | Level 1 Heading Block |
| `heading2` | `object(Text)` | optional | `-` | Level 2 Heading Block |
| `heading3` | `object(Text)` | optional | `-` | Level 3 Heading Block |
| `heading4` | `object(Text)` | optional | `-` | Level 4 Heading Block |
| `heading5` | `object(Text)` | optional | `-` | Level 5 Heading Block |
| `heading6` | `object(Text)` | optional | `-` | Level 6 Heading Block |
| `heading7` | `object(Text)` | optional | `-` | Level 7 Heading Block |
| `heading8` | `object(Text)` | optional | `-` | Level 8 Heading Block |
| `heading9` | `object(Text)` | optional | `-` | Level 9 Heading Block |
| `bullet` | `object(Text)` | optional | `-` | Bullet List Block |
| `ordered` | `object(Text)` | optional | `-` | Ordered List Block |
| `code` | `object(Text)` | optional | `-` | Code Block |
| `quote` | `object(Text)` | optional| `-` | Quote Block |
| `todo` | `object(Text)` | optional | `-` | Todo Block |
| `bitable` | `object(Bitable)` | optional | `-` | Bitable Block |
| `callout` | `object(Callout)` | optional | `-` | Callout Block |
| `chat_card` | `object(ChatCard)` |optional | `-` | Chat Card Block |
| `diagram` | `object(Diagram)` | optional | `-` | UML Diagram Block |
| `divider` | `object(Divider)` | optional | `-` | Divider Block |
| `file` | `object(File)` | optional | `-` | File Block |
| `grid` | `object(Grid)` | optional | `-` | Grid Block |
| `grid_column`| `object(GridColumn)` | optional | `-` | Grid Column Block |
| `iframe` | `object(Iframe)` | optional | `-` | Iframe Block |
| `image` | `object(Image)` | optional | `-` | Image Block |
| `isv` | `object(ISV)` | optional | `-` | ISV Block |
| `mindnote` | `object(Mindnote)` | optional | `-` | Mindnote Block |
| `sheet` | `object(Sheet)` | optional | `-` | Sheet Block |
| `table` | `object(Table)` | optional | `-` | Table Block |
| `table_cell`| `object(TableCell)` |optional | `-` | Table Cell Block |
| `view` | `object(View)` | optional | `-` | View Block |
| `undefined` | `object(Undefined)` |ooptional | `-` | Undefined Block|
| `quote_container` | `object(QuoteContainer)` |ooptional | `-` | Quote Container Block|
| `task`      | `object(Task)` | optional | `-` | Task Block|
| `okr` | `object(OKR)` | optional | `-` | OKR Block |
| `okr_objective` | `object(OkrObjective)` | optional| `-` | OKR Objective Block |
| `okr_key_result` | `object(OkrKeyResult` | optional| `-` | OKR Key Result Block|
| `okr_progress` |  `object(OkrProgress)`  | optional| `-` | OKR Progress Block| Up to now, DocX has supported 30+ kinds of Blocks, some of which have the capacity to accommodate, such as `GridColumn` can accommodate `Image`, `Text`, etc.; some Block content entities are `Text`, such as ` heading1~9`, `ordered`, `code`, etc. 

## 2.3 BlockData

Each type of Block has its corresponding content entity, which is collectively referred to as **BlockData**.

### Bitable
Creating an empty Bitable by specifying a `view_type` is currently supported.
````json
{
    "token": string,
    "view_type": enum(BitableViewType)
}
````

| Name | Data Type | Properties | Default Value | Description |
| --------- | --------------- | ------- | ----------- | --------- |
|`token` | `string` | required | `-` |Bitable Token, readonly.The format is BitableToken**_**TableID, where BitableToken is the unique identifier of a Bitable, and TableID is the unique identifier of a data table. Please pay attention to splitting when using. |
|`view_type` | `enum(BitableViewType)` | required | `-` | view type | ### Callout

Callout Block
````json
{
    "background_color": enum(BackgroundColor),
    "border_color": enum(BorderColor),
    "text_color": enum(FontColor),
    "emoji_id": string
}
````

| Name | Data Type | Properties | Default Value | Description |
| --------- | --------------- | ------- | ----------- | --------- |
|`background_color` | `enum(BackgroundColor)` | optional | `-` | background color |
|`border_color` | `enum(BorderColor)` | optional | `-` | border color |
|`text_color` | `enum(FontColor)` | optional | `-` | font color |
|`emoji_id` | `enum(string)` | optional | `-` | EmojiID | ### ChatCard

Chat Card Block
````json
{
    "chat_id": string,
    "align": enum(Align)
}
````

| Name | Data Type | Properties | Default Value | Description |
| --------- | --------------- | ------- | ----------- | --------- |
|`chat_id` | `string` | required | `-` | Group chat session OpenID, whose value starts with 'oc_', represents the group ID defined for the open capability interface. For write operations, no permission error is returned if the user is not in the group. |
|`align` | `enum(Align)` | optional | `-` | alignment | ### Diagram

UML Diagram Block.
````json
{
    "diagram_type": enum(DiagramType)
}
````

| Name | Data Type | Properties | Default Value | Description |
| --------- | --------------- | ------- | ----------- | --------- |
|`diagram_type` | `enum(DiagramType)` | optional | `-` | drawing type | ### Divider

Divider Block, empty structure
````
{}
````
| **Attribute** | **Type** | **Description** | **Example** | **Parameter Check** |
| ------ | ------ | ------ | ------ | -------- | ### File

File Block. The file block cannot exist independently, it must appear together with the view block, and the file view is realized by controlling the `view_type` of the view block, such as card view, preview view, etc. When creating a file block, the system automatically generates a default view block.

````json
{
    "token": string,
    "name": string
}
````

| Name | Data Type | Properties | Default Value | Description |
| --------- | --------------- | ------- | ----------- | --------- |
|`token` | `string` | optional | `-` | Attachment token, readonly |
|`name` | `string` | optional | `-` | Attachment filename, readonly | ### Grid

Grid Block
````json
{
    "column_size": int
}
````

| Name | Data Type | Properties | Default Value | Description |
| --------- | --------------- | ------- | ----------- | --------- |
|`column_size` | `int` | required | `-` | number of columns, value range [2,5) | ### Grid Column
Grid Column Block.
````json
{
    "width_ratio": int
}
````

| Name | Data Type | Properties | Default Value | Description |
| --------- | --------------- | ------- | ----------- | --------- |
|`width_ratio` | `int` | optional | `-` | The ratio of the current column to the entire column, the value range is [1,99] | ### Iframe
Iframe Block
````json
{
    "component": {
      "type": enum(IframeComponentType),
      "url": string
    }
}
````
| Name | Data Type | Properties | Default Value | Description |
| --------- | --------------- | ------- | ----------- | --------- |
|`component` | `object(IframeComponent)` | required | `-` | Iframe component |
|`component.type` | `enum(IframeComponentType)` | required | `-` | Iframe component type |
|`component.url` | `string` | required | `-` | Target webpage URL, URL Encode is required for both reading and writing | ### Image
Image Block.

````json
{
    "token": string,
    "width": int,
    "height": int
}
````
| Name | Data Type | Properties | Default Value | Description |
| --------- | --------------- | ------- | ----------- | --------- |
|`token` | `string` | optional | `-` | image Token, readonly |
|`width` | `int` | optional | `100` | image width in px, readonly |
|`height` | `int` | optional | `100` | image height in px, readonly | ### ISV
ISV Block
````json
{
    "component_id": string,
    "component_type_id": string
}
````
| Name | Data Type | Properties | Default Value | Description |
| --------- | --------------- | ------- | ----------- | --------- |
|`component_id` | `string` | optional | `-` | Team interaction application unique ID |
|`component_type_id` | `string` | optional | `-` | Team interaction application type, such as information gathering | ### Mindnote
Mindnote Block. Currently only supports query to return placeholder information, does not support creation and editing. 
````json
{
    "token": string
}
````
| **Attribute** | **Type** | **Description** | **Example** | **Parameter Check** |
| ------ | ------ | ---------- | ----------------------------- | -------- |
| token | string | Thought Notes Token | `bmnbcJ7X9yUiNbuG8DfkgqxOtlf` | `Only output` | ### Sheet

Currently only specifying `row_size` and `column_size` creates an empty Sheet.
````json
{
    "token": string,
    "row_size": int,
    "column_size": int,
}
````
| Name | Data Type | Properties | Default Value | Description |
| --------- | --------------- | ------- | ----------- | --------- |
|`token` | `string` | optional | `-` | Spreadsheet Token, readonly. The format is SpreadsheetToken**_**SheetID. Among them, SpreadsheetToken is the unique identifier of a spreadsheet, SheetID is the unique identifier of a worksheet, please pay attention to splitting when using.|
|`row_size` | `int` | optional | `-` | Number of spreadsheet columns. Used when creating an empty spreadsheet, maximum 9. |
|`column_size` | `int` | optional | `-` | Number of spreadsheet columns. Used when creating an empty spreadsheet, maximum 9. | ### Table
Table Block. Supports specifying `property` to create an empty table.
````json
{
    "cells": [string, string, ...],
    "property": {
        "row_size": int,
        "column_size": int,
        "column_width":[](int),
        "merge_info": [](TableMergeInfo)
    }
}
// TableMergeInfo cell merge information
{
    "row_span": int,
    "col_span": int
}
````
| Name | Data Type | Properties | Default Value | Description |
| --------- | --------------- | ------- | ----------- | --------- |
|`cells` | `array(string)` | optional | `-` | Cell array, the array element is the ID of the Table Cell Block. |
|`property` | `object(TableProperty)` | required | `-` | Table property. |
|`property.row_size` | `int` | required | `-` | The number of rows. |
|`property.column_size` | `int` | required | `-` | Number of columns. |
|`property.column_width` | `array(int)` | optional | `-` | Column width in px. |
|`property.merge_info` | `object(TableMergeInfo)` | optional | `-` | Cell merge information, readonly. |
|`property.merge_info.row_span` | `int` | optional | `-` | The number of consecutive rows to be merged from the current row index. |
|`property.merge_info.col_span` | `int` | optional | `-` | The number of consecutive columns to be merged from the current column index. | ### Table Cell
Table Cell Block.
````json
{}
````
| Name | Data Type | Properties | Default Value | Description |
| --------- | --------------- | ------- | ----------- | --------- | ### View
View Block, which hosts the card / preview view.
````json
{
    "view_type": enum(ViewType)
}
````
| Name | Data Type | Properties | Default Value | Description |
| --------- | --------------- | ------- | ----------- | --------- |
| `view_type` | `enum(ViewType)` | optional | `-` | view type | ### Undefined
Unsupported Block.

### QuoteContainer
Quote Container Block.
````json
{}
````
| Name | Data Type | Properties | Default Value | Description |
| --------- | --------------- | ------- | ----------- | --------- | ### Text
Text Block, which has various types
````json
{
    "style": object(TextStyle),
    "elements": [
      object(TextElement)
    ]
}
````
| Name | Data Type | Properties | Default Value | Description |
| --------- | --------------- | ------- | ----------- | --------- |
| `style` | `object(TextStyle)` | optional | `-` | text style |
| `elements` | `array(TextElement)` | required | `-` | text elements | ### TextStyle
````json
{
    "align": enum(Align),
    "done": boolean,
    "folded": boolean,
    "language": enum(CodeLanguage),
    "wrap": boolean
}
````
| Name | Data Type | Properties | Default Value | Description |
| --------- | --------------- | ------- | ----------- | --------- |
| `align` | `enum(Align)` | optional | `1` | alignment|
| `done` | `boolean` | optional | `false` | Todo's done status |
| `folded` | `boolean` | optional | `false` | folded state of the text |
| `language` | `enum(CodeLanguage)` | optional | `-` | code block language |
| `wrap` | `boolean` | optional | `false` | whether the code block will automatically wrap | ### TextElement
Text element, supports multiple types.
````json
{
    "text_run": object(TextRun),
    "mention_user": object(MentionUser),
    "mention_doc": object(MentionDoc),
    "reminder": object(Reminder),
    "file": object(InlineFile),
    "inline_block": object(InlineBlock)
    "equation": object(Equation),
    "undefined_element": object(UndefinedElement)
}
````
| Name | Data Type | Properties | Default Value | Description |
| --------- | --------------- | ------- | ----------- | --------- |
| `text_run` | `object(TextRun)` | optional | `-` |Text|
| `mention_user` | `object(MentionUser)` | optional | `-` |@User|
| `mention_doc` | `object(MentionDoc)` | optional | `-` |@Docs|
| `reminder` | `object(Reminder)` | optional | `-` |Date Reminder|
| `file` | `object(InlineFile)` | optional | `-` | Inline Attachments |
| `inline_block` | `object(InlineBlock)` |optional |`-`|  Inline Block |
| `equation` | `object(Equation)` | optional | `-` | Equation |
| `undefined_element` | `object(UndefinedElement)` | optional | `-` |Unsupported TextElement| ### TextElementStyle
text local style
````json
{
    "bold": boolean,
    "italic": boolean,
    "strikethrough": boolean,
    "underline": boolean,
    "inline_code": boolean,
    "text_color": enum(FontColor),
    "background_color": enum(BackgroundColor),
    "link": object(Link),
    "comment_ids": [](string)
}
````
| Name | Data Type | Properties | Default Value | Description |
| --------- | --------------- | ------- | ----------- | --------- |
| `bold` | `boolean` | optional | `false` |bold|
| `italic` | `boolean` | optional | `false` | italic|
| `strikethrough` | `boolean` | optional | `false` |strikethrough|
| `underline` | `boolean` | optional | `false` | underline |
| `inline_code` | `boolean` | optional | `false` |inline code|
| `text_color` | `enum(FontColor)` | optional | `-` |font color|
| `background_color` | `enum(BackgroundColor)` | optional | `-` |background color|
| `link` | `object(Link)` | optional | `-` | link|
| `comment_ids` | `array(string)` | optional | `-` | list of comment ID| ### TextElementData
text element
#### TextRun
text
````json
{
    "content": string,
    "text_element_style": object(TextElementStyle)
}
````
| Name | Data Type | Properties | Default Value | Description |
| --------- | --------------- | ------- | ----------- | --------- |
| `content` | `string` | required | `-` |text|
| `text_element_style` | `object(TextElementStyle)` | optional | `-` | text local style| #### MentionUser
@User
````json
{
    "user_id": string,
    "text_element_style": object(TextElementStyle)
}
````
| Name | Data Type | Properties | Default Value | Description |
| --------- | --------------- | ------- | ----------- | --------- |
| `user_id` | `string` | required | `-` |user OpenID|
| `text_element_style` | `object(TextElementStyle)` | optional | `-` | text local style| #### MentionDoc
@Docs
````json
{
    "token": string,
    "obj_type": enum(MentionObjType),
    "url": string,
    "text_element_style": object(TextElementStyle)
}
````
| Name | Data Type | Properties | Default Value | Description |
| --------- | --------------- | ------- | ----------- | --------- |
| `token` | `string` | required | `-` |Cloud Document Token|
| `obj_type` | `enum(MentionObjType)` | required | `-` |Cloud Document Type|
| `url` | `string` | required | `-` |Cloud document link (requires URL Encode)|
| `text_element_style` | `object(TextElementStyle)` | optional | `-` | text local style| #### Reminder
Reminder
````json
{
    "create_user_id": string,
    "is_notify": boolean,
    "is_whole_day": boolean,
    "expire_time": int,
    "notify_time": int,
    "text_element_style": object(TextElementStyle)
}
````
| Name | Data Type | Properties | Default Value | Description |
| --------- | --------------- | ------- | ----------- | --------- |
| `create_user_id` | `string` | required | `-` |Creator User ID|
| `is_notify` | `boolean` | optional | `false` |whether to notify|
| `is_whole_day` | `boolean` | optional | `false` | day or hour |
| `expire_time` | `int` | required | `-` | time when the event occurred (event stamp in milliseconds) |
| `notify_time` | `int` | required | `-` | time when the notification was triggered (timestamp in milliseconds) |
| `text_element_style` | `object(TextElementStyle)` | optional | `-` | text local style| #### InlineFile (output only)
Inline Attachments
````json
{
    "file_token": string,
    "source_block_id": string,
    "text_element_style": object(TextElementStyle)
}
````
| Name | Data Type | Properties | Default Value | Description |
| --------- | --------------- | ------- | ----------- | --------- |
| `file_token` | `string` | optional | `-` |Attachment Token|
| `source_block_id` | `string` | optional | `-` |BlockID of this attachment in the current document|
| `text_element_style` | `object(TextElementStyle)` | optional | `-` | text local style| #### InlineBlock (output only)
Inline Block
```json
{
	"block_id": string,
	"text_element_style": object(TextElementStyle)
}
```
| Name | Data Type | Properties | Default Value | Description |
| --------- | --------------- | ------- | ----------- | --------- |
| `block_id` | `string` | optional | `-` |the BlockID of the inlined Block|
| `text_element_style` | `object(TextElementStyle)` | optional | `-` | text local style| #### Equation
Equation
```json
{
    "content": string,
    "text_element_style": object(TextElementStyle)
}
```
| Name | Data Type | Properties | Default Value | Description |
| --------- | --------------- | -------   | ----------- | --------- |
| `content` | `string` | required | `-` |Equation content that conforms to KaTeX syntax. Grammar rules refer to: https://katex.org/docs/supported.html|
| `text_element_style` | `object(TextElementStyle)` | optional | `-` | text local style| #### UndefinedElement (output only)
Unsupported TextElement
````json
{}
````
| Name | Data Type | Properties | Default Value | Description |
| --------- | --------------- | ------- | ----------- | --------- | ### Task
Task block. Only support the return task ID, and do not support the creation and editing.
```json
{
    "task_id": string
}
```

| Name | Data Type | Properties | Default value | Description |
| --------- | --------------- | -------   | ----------- | --------- |
| `task_id` | `string` | required | `-` |Task ID| ### OKR
OKR Block, only supported to create when using `user_access_token`

**Crate**

```json
{
    "okr_id": string,
    "objectives": [](Objective)
}
// Objective
{
    "objective_id": string,
    "kr_ids": [](string)
}
```

| Name | Data Type | Properties | Default value | Description |
| --------- | --------------- | -------   | ----------- | --------- |
| `okr_id`  | `string`  | required | `-` | OKR ID
| `objectives` | `array(Objective)` | optional | `-` | Objective ID and Key Result ID in OKR Block.When this value is null, it insert into all Objective and Key Result of the OKR.|
| `objectives.objective_id`| `string` | required | `-` | OKR's ID of Objective.| |
| `objectives.kr_ids`| `array(string)` | optional | `-` | The ID list of Key Result.When this value is null, it insert into all key results under the current Objective.| **Query**

```json
{
    "okr_id": string,
    "period_display_status": enum(OkrPeriodDisplayStatus),
    "period_name_zh": string,
    "period_name_en": string,
    "user_id": string,
    "visible_setting": {
        "progress_fill_area_visible": boolean,
        "progress_status_visible": boolean,
        "score_visible": boolean
    }
}
```

| Name | Data Type | Properties | Default value | Description |
| --------- | --------------- | -------   | ----------- | --------- |
| `okr_id`  | `string`  | optional | `-` | OKR ID |
| `period_display_status` | `enum(OkrPeriodDisplayStatus)` | optional | `-` | The status of period |
| `period_name_zh` | `string` | optional | `-` | Chinese period name |
| `period_name_en` | `string` | optional | `-` | English period name |
| `user_id`        | `string` | optional | `-` | The owner ID of the OKR |
| `visible_setting` | `object(VisibleSetting)` | optional | `-` | Visible settings |
| `visible_setting.progress_fill_area_visible` | `boolean` | optional | `true` | Whether the progress edit area is visible |
| `visible_setting.progress_status_visible` | `boolean` | optional | `true` | Whether the progress is visible |
| `visible_setting.score_visible` | `boolean` | optional | `true` | Whether the score is visible | ### OkrObjective
OKR Objective Block

```json
{
    "objective_id": string,
    "confidential": boolean,
    "position": int,
    "score": int,
    "visible": int,
    "weight": float,
    "progress_rate": object(ProgressRate),
    "content": object(Text)
}
```

| Name | Data Type | Properties | Default value | Description |
| --------- | --------------- | -------   | ----------- | --------- |
| `objective_id`  | `string`  | optional | `-` | OKR Objective ID |
| `confidential` | `boolean` | optional | `-` | Whether there is a confidential setting on the OKR platform |
| `position` | `int` | optional | `-` | Objective position number, corresponding to 1 and 2 of O1 and O2 in Block |
| `score` | `int` | optional | `-` | Scoring information |
| `visible` | `boolean` | optional | `true` | Whether the objective is displayed in OKR Block |
| `weight` | `float` | optional | `-` | The weight of the objective |
| `progress_rate` | `object(ProgressRate)` | optional | `-` | Progress information |
| `content` | `object(Text)` | optional | `-` | Objective text content | ### OkrKeyResult
OKR Key Result Block

```json
{
    "kr_id": string,
    "confidential": boolean,
    "position": int,
    "score": int,
    "visible": int,
    "weight": float,
    "progress_rate": object(ProgressRate),
    "content": object(Text)
}
```

| Name | Data Type | Properties | Default value | Description |
| --------- | --------------- | -------   | ----------- | --------- |
| `kr_id`  | `string`  | optional | `-` | OKR Key Result ID |
| `confidential` | `boolean` | optional | `-` | Whether there is a confidential setting on the OKR platform |
| `position` | `int` | optional | `-` | Key result position number, corresponding to 1 and 2 of KR1 and KR2 in Block |
| `score` | `int` | optional | `-` | Scoring information |
| `visible` | `boolean` | optional | `true` | Whether the key result is displayed in OKR Block |
| `weight` | `float` | optional | `-` | The weight of the key result |
| `progress_rate` | `object(ProgressRate)` | optional | `-` | Progress information |
| `content` | `object(Text)` | optional | `-` | Key Result text content | ### ProgressRate
OKR progress information

```json
{
    "mode": enum(OkrProgressRateMode),
    "current": float,
    "percent": float,
    "progress_status": enum(OkrProgressStatus),
    "status_type": enum(OkrProgressStatusType),
    "start": float,
    "target": float
}
```

| Name | Data Type | Properties | Default value | Description |
| --------- | --------------- | -------   | ----------- | --------- |
| `mode`  | `enum(OkrProgressRateMode)`  | optional | `-` | Status mode, divided into simple and advanced|
| `status_type` | `enum(OkrProgressStatusType)` | optional | `-` | The calculation type of OKR progress status |
| `progress_status` | `enum(OkrProgressStatus)` | optional | `-` | progress status |
| `percent` | `float` | optional | `-` | Current progress percentage, used in simple mode |
| `start`   | `float` | optional | `-` | The start value of progress, used in advanced mode |
| `current` | `float` | optional | `-` | Current progress, used in advanced mode |
| `target`  | `float` | optional | `-` | Progress target value, used in advanced mode| ### Progress
OKR Progress Block

```json
{}
```

# 3 Enumeration

## BlockType
| **Enums** | |
| ---------- | ----------- |
|1|Page Block|
|2|Text Block|
|3|Title 1 Block|
|4|Title 2 Block|
|5|Title 3 Block|
|6|Title 4 Block|
|7|Title 5 Block|
|8|Title 6 Block|
|9|Title 7 Block|
|10|Title 8 Block|
|11|Title 9 Block|
|12|Unordered List Block|
|13|Ordered List Block|
|14|Code Block Block|
|15|Cite Block|
|16|Formula Block|
|17|To Do Block|
|18|Multidimensional Table Block|
|19|Highlight Block|
|20|Session Card Block|
|21|Flowchart & UML Block|
|22|Separation Line Block|
|23|File Block|
|24|Grid Block|
|25|Grid Column Block|
|26|Embedded Block Block|
|27|Picture Block|
|28|Open Platform Widget Block|
|29|Thinking Notes Block|
|30|Spreadsheet Block|
|31|Form Block|
|32|Table Cell Block|
|33|View Block|
|34|QuoteContainer Block|
|35|Task Block|
|36|OKR Block|
|37|OKR Objective Block|
|38|OKR Key Result Block|
|39|OKR Progress Block|
|999|Block is not supported| ## TextElementType
| **Enums** | |
| ------------ | ----- |
| text_run | Text |
| mention_user | @User |
| mention_doc | @Docs |
| file | Inline Attachments |
| reminder | Date Reminder |
| undefined | Unsupported Element |
| equation    | Equation | ## Align
Block layout, such as left, etc.
| **Enums** | |
| --------- | ------ |
|1|Left layout|
|2|Centered Typesetting|
|3|Right layout| ## CalloutBackgroundColor
Callout block background color (divided into dark and light colors)
| **Enums** | |
| --------- | ------ |
|1|Light Red|
|2|Light Orange|
|3|Light Yellow|
|4|Light Green|
|5|Light Blue|
|6|Light Purple|
|7|Light Gray|
|8|Dark Red|
|9|Dark Orange|
|10|Dark Yellow|
|11|Dark Green|
|12|Dark Blue|
|13|Dark Purple|
|14|Dark Grey| ## CalloutBorderColor
The border color of the highlight block (the color system is the same as the background color of the highlight block)
| **Enums** | |
| --------- | ------ |
|1|Red|
|2|Orange|
|3|Yellow|
|4|Green|
|5|Blue|
|6|Purple|
|7|Grey| ## FontBackgroundColor
Font background color (divided into dark and light colors)
| **Enums** | |
| --------- | ------ |
|1|Light Pink|
|2|Light Orange|
|3|Light Yellow|
|4|Light Green|
|5|Light Blue|
|6|Light Purple|
|7|Light Gray|
|8|Dark Pink|
|9|Dark Orange|
|10|Dark Yellow|
|11|Dark Green|
|12|Dark Blue|
|13|Dark Purple|
|14|Dark Grey| ## FontColor
font color
| **Enums** | |
| --------- | ------ |
|1|pink|
|2|Orange|
|3|Yellow|
|4|Green|
|5|Blue|
|6|Purple|
|7|Grey| ## ViewType
The view type of the View Block
| **Enums** | |
| --------- | ------ |
|1|Card view, a view on a single line, there can be some simple interactions on the Card|
|2|Preview view, preview the inserted Block content directly on the current page without opening a new page|
|3|Inline attempt| ## BitableViewType
View type for Bitable Block
| **Enums** | |
| --------- | ------ |
|1|Data Sheet|
|2|Kanban| ## DiagramType
Diagram type
| **Enums** | |
| --------- | ------ |
|1|Flowchart|
|2|UML Diagram| ## IframeComponentType
Types supported by Iframe Blocks
| **Enums** | |
| --------- | ------ |
|1|Bilibili|
|2|Watermelon Video|
|3|Youku|
|4|Airtable|
|5|Baidu Map|
|6|Amap Map|
|7|Undefined|
|8|Figma|
|9|Ink Knife|
|10|Canva|
|11|CodePen|
|12|Lark Questionnaire|
|13|Gold Data|
|14|Undefined|
|15|Undefined| ## MentionObjType
Mention Doc Type
| **Enums** | |
| --------- | ------ |
|1|Doc|
|3|Sheet|
|8|Bitable|
|11|MindNote|
|12|File|
|15|Slide|
|16|Wiki|
|22|Docx|
## CodeLanguage
code block language
| **Enums** | |
| --------- | ------ |
|1|PlainText
|2|ABAP
|3|Ada
|4|Apache|
|5|Apex|
|6|Assembly| Language|
|7|Bash|
|8|CSharp|
|9|C++|
|10|C|
|11|COBOL|
|12|CSS
|13|CoffeeScript|
|14|D|
|15|Dart|
|16|Delphi|
|17|Django|
|18|Dockerfile|
|19|Erlang|
|20|Fortran|
|21|FoxPro|
|22|Go|
|23|Groovy|
|24|HTML|
|25|HTMLBars|
|26|HTTP|
|27|Haskell|
|28|JSON|
|29|Java|
|30|JavaScript|
|31|Julia|
|32|Kotlin|
|33|LateX|
|34|Lisp|
|35|Logo|
|36|Lua|
|37|MATLAB|
|38|Makefile|
|39|Markdown|
|40|Nginx|
|41|Objective|-C|
|42|OpenEdgeABL|
|43|PHP|
|44|Perl
|45|PostScript|
|46|Power| Shell|
|47|Prolog|
|48|ProtoBuf|
|49|Python
|50|R|
|51|RPG|
|52|Ruby|
|53|Rust|
|54|SAS|
|55|SCSS|
|56|SQL|
|57|Scala|
|58|Scheme|
|59|Scratch|
|60|Shell|
|61|Swift|
|62|Thrift|
|63|TypeScript|
|64|VBScript|
|65|Visual| Basic|
|66|XML|
|67|YAML| ## OkrPeriodDisplayStatus
Period display status
| **Enums** |        |
| --------- | ------ |
|default|default|
|normal|normal|
|invalid|invalid|
|hidden|hidden| ## OkrProgressRateMode
OKR progress status mode
| **Enums** |        |
| --------- | ------ |
|simple  |simple mode|
|advanced|advanced mode| ## OkrProgressStatus
OKR progress status
| **Enums** |        |
| --------- | ------ |
|unset|unset|
|normal|normal|
|risk|risk|
|extended|Extended| ## OkrProgressStatusType
The calculation type of OKR progress status
| **Enums** |        |
| --------- | ------ |
|default|Display with the most risky Key Result status|
|custom|custom|
