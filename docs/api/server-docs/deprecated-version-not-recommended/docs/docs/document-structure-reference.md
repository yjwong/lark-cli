---
title: "Document Structure Reference"
url: "https://open.larksuite.com/document/ukTMukTMukTM/ukDM2YjL5AjN24SOwYjN"
section: "Deprecated Version (Not Recommended)"
updated: "1695369901000"
---

# Document

Describes the rich text content of a document, which is composed of headings and text. [Obtain rich text content]
] 

```json
{
    "title": {object(Paragraph)}, 
    "body": {object(Body)}
}
```

**Field description**:  

| Field  | Type  | Description |
| ----- | ------ | -------- |
| title | object | Document heading |
| body  | object | Body | # Location

Describes the locations of the elements in a document.

```json
{
    "zoneId": string,
    "startIndex": int, 
    "endIndex": int
}
```

**Field description**:  

| Field            | Type        | Description                                                         |
| ---------- | ------ | ------------------------------------------------------------ |
| zoneId     | string | Indicates the zone to edit, "0" indicates bodyA cell ID, such as "xr1m4jw7egd9nefz1s0mdsetenl5fbe3lygxc1azupv81i5t2rjmosw5ta0esuwtn8ksya", indicates a cell |
| startIndex | int    | Element start location offset, starts from 0, in the range [0, EOF)             |
| endIndex | int    | Element end location offset, in the range [0, EOF)             | # Title

Document heading, only supports plain text and alignment.

**Field description**:  

|Field|Description|
| ------------------ | ---------------------------------------- |
| paragraph.style    | Only supports align properties                          |
| paragraph.elements | Only supports textRun. The input textStyle will be discarded. | # Body

The body of the document, composed of blocks.

```json
{
    "blocks": [{object(Block)}]
}
```

**Field description**:  

| Field            | Type        | Description                                                         |
| ------ | -------- | ------------------------------------------------------------ |
| blocks | []object | The document structure is ordered by row. Each row is a block. For details, see Document Data Structure Overview. | # Block

The basic elements of a document row.

```json
{
    "type": string,
    "paragraph": {object(Paragraph)},
    "horizontalLine": {object(HorizontalLine)},
    "embeddedPage": {object(EmbeddedPage)},
    "chatGroup": {object(ChatGroup)},
    "table": {object(Table)},
    "sheet": {object(Sheet)},
    "bitable": {object(Bitable)},
    "gallery": {object(Gallery)},
    "file": {object(File)},
    "diagram": {object(Diagram)},
    "jira": {object(Jira)},
    "poll": {object(Poll)},
    "code": {object(Code)},
    "docsApp": {object(DocsApp)},
    "callout": {object(Callout)},
    "undefinedBlock": {object{UndefinedBlock}}
}
```

**Field description**:  

| Field  | Type  | Description |
| -------------- | ------ | ------------------------------------------- |
| type           | string | One of the following types must be specified for each block |
| paragraph      | object | Paragraph of text                                    |
| gallery        | object | Image                                        |
| file           | object | File upload                                    |
| chatGroup      | object | Group card                                      |
| table          | object | Formatted sheet                                  |
| horizontalLine | object | Horizontal divider                                  |
| embeddedPage   | object | Embedded webpage, such as a Xigua Video                    |
| sheet          | object | Sheet                                    |
| bitable        | object | Bitable or kanban                                |
| diagram        | object | Diagram or UML diagram                                 |
| jira           | object | Jira filter or Jira issue                     |
| poll           | object | Poll                                        |
| code           | object | New code block                                      |
| docsApp           | object |               Team interaction app                      |
| callout           | object |               Callout                      |
| undefinedBlock | object | All unsupported blocks are labeled undefineBlock      | # Paragraph

A paragraph can be composed of multiple row elements.

```json
{
    "style": object (ParagraphStyle),
    "elements":  [{object (ParagraphElement)}],
    "location": {object(Location)}
}
```

**Field description**:  

| Field            | Type        | Description                                                         |
| -------- | -------- | ---------------------------- |
| style    | object   | Paragraph style                     |
| elements | []object | The paragraph elements of one paragraph          |
| location | Location | The element location, can't be written, only returned | # ParagraphStyle

The style of a paragraph, supports custom paragraph title, list, references, and row alignment.

```json
{
    "headingLevel": int,
    "collapse": bool,
    "list": {object (List)},
    "quote": bool,
    "align": string
}
```

**Field description**:  

| Field  | Type  | Description |
| ------------ | ------ | ---------------------------------------------------------- |
| headingLevel | int    | Heading level, level 1 to 9, value range: [1,9]                 |
|collapse| bool| Whether heading is collapsed, only valid when headingLevel is [1,9]|
| list         | object | Numbered list/Bulleted list/Task list/Old code block                          |
| quote        | bool   | Reference style                                                   |
| align        | string | Row alignmentAlign left (default): leftAlign right: rightAlign center: center | # List

List paragraph, including numbered lists, bulleted lists, task lists, and old code blocks.  
Note To insert a code block, use Block.Code. Paragraph.List.Code will soon be removed.

```json
{
    "type": string,
    "indentLevel": int,
    "number": int
}
```

**Field description**:  

| Field            | Type        | Description                                                         |
| ----------- | ------ | ------------------------------------------------------------ |
| type        | string | Numbered list: numberBulleted list: bulletTask list: checkBoxCompleted task list: checkedBoxOld code block: code |
| indentLevel | int    | List indent level, allows you to specify a row indentYou can set level 1 to 16 indents for lists other than code blocks, value range: [1,16]. |
| number      | Int    | This is used to specify the row numbers of a list, only valid for numbered lists and code blocksIf an indent is set for a numbered list, the row numbers may be displayed as letters or roman numerals.

# ParagraphElement

Row elements composing a paragraph. A single row can contain multiple row elements.

```json
{
    "type": string,
    "textRun": {object(TextRun)},
    "docsLink": {object(DocsLink)},
    "person": {object(Person)},
    "equation": {object(Equation)},
    "reminder": {object(Reminder)},
    "file": {object(File)},
    "jira": {object(Jira)},
    "undefinedElement": {object(UndefinedElement)}
}
```

**Field description**:  

| Field  | Type | Description |
| ---------------- | ------ | -------------------------------------------------------- |
| type           | string | One of the following types must be specified for each element |
| textRun          | object | Text                                                     |
| docsLink         | object | Document link, the title can be automatically recognized from the link                     |
| person           | object | @mention user                                                    |
| equation         | object | LaTeX equation                                               |
| reminder         | object | Task list deadline, only used with checkBox rows |
| file             | object | The file row display style, used to display the file name                |
| jira             | object | The row's jira issue                    |
| undefinedElement | object | Unsupported row element, can't be written                           | # TextRun

Text content of a paragraph.

```json
{
    "text": string,
    "style": {object(TextStyle)},
    "lineId": string,
    "location": {object(Location)}
}
```

**Field description**:  

| Field            | Type        | Description                                                         |
| -------- | -------- | ------------------------------------------------------------ |
| text     | string   | Specific text content                                               |
| style    | object   | Text content style, supports BIUS, colors, etc.                            |
| lineId   | string   | Can't be input, only returned. When the paragraph style is Heading, you can use the lineId to create redirection links. |
| location | Location | Can't be input, only returned. Describes the location of the element in the current area (specified by the zoneId). | # TextStyle

Text style.

```json
{
    "bold": bool,
    "italic": bool,
    "strikeThrough": bool,
    "underLine": bool,
    "codeInline": bool,
    "backColor": {object(RGBColor)},
    "textColor": {object(RGBColor)},
    "link": {object(Link)}
}
```

**Field description**:  

| Field  | Type | Description |
| ------------- | ------ | ------------ |
| bold          | bool   | Bold         |
| italic        | bool   | Italic         |
| strikeThrough | bool   | Strikethrough       |
| underLine     | bool   | Underline         |
| codeInline    | bool   | In-row code style |
| backColor     | object | Background color       |
| textColor     | object | Text color      |
| link          | object | Hyperlink       | # RGBColor

You can specify the text color. Colors will be checked to make sure they are among the colors supported for online text.

```json
{
    "red": int,
    "green": int,
    "blue": int,
    "alpha": float
}
```

**Field description**:  

| Field  | Type  | Description |
| ----- | ----- | ----------------- |
| red   | int   | Value range: [0,255] |
| green | int   | Value range: [0,255] |
| blue  | int   | Value range: [0,255] |
| alpha | float | Value range: [0,1]   | **Color verification**:

> Supported text colors: rgb(216, 57, 49), rgb(222, 120, 2), rgb(220, 155, 4), rgb(46, 161, 33), rgb(36, 91, 219), rgb(100, 37, 208), rgb(143, 149, 158)
>
> Supported background colors: rgb(251, 191, 188), rgb(254, 212, 164), rgb(255, 246, 122), rgb(183, 237, 177), rgb(186, 206, 253), rgb(205, 178, 250), rgb(222, 224, 227), rgb(247, 105, 100), rgb(255, 165, 61), rgb(255, 233, 40), rgb(98, 210, 86), rgb(78, 131, 253), rgb(147, 90, 246), rgb(187, 191, 196)

# Link

Hyperlink.

```json
{
    "url": string
}
```

**Field description**:  

| Field | Type   | Description                                                         |
| ---- | ------ | ------------------------------------------------------------ |
| url  | string | Hyperlink URL, query escape  encoding must be used for read and write, for example: https%3A%2F%2Fwww.baidu.com%2Fmaps%2Fembed%2Fv1%2Fplace%3Fkey%3DAIzaSyAfJZc8JxNRe909WC_QBILdlM55NqGnI30%26q%3DCentral%2BPark%26center%3D40.7828647%2C-73.9675438%26zoom%3D17z%26language%3Den | # DocsLink

Document link, must be a Docs link for any Docs type. Valid links have the format: https://sample.larksuite.com/docs/doccnByZP6puODElAYySJkPIfUb.

```json
{
    "url": string,
    "location": {object(Location)}
}
```

**Field description**:  

| Field            | Type        | Description                                                         |
| -------- | -------- | ---------------------------- |
| url      | string   | Docs link                   |
| location | Location | The element location, can't be written, only returned | # Person

```json
{
    "openId": string,
    "location": {object(Location)}
}
```

**Field description**:  

| Field            | Type        | Description                                                         |
| -------- | -------- | -------------------------------------------------------- |
| openId   | string   | The user's openId, for example: ou_3bbe8a09c20e89cce9bff989ed840674 |
| location | Location | The element location, can't be written, only returned | # Reminder

Date reminder, only valid for checkBox row styles.

```json
{
    "isWholeDay": bool,
    "timestamp": int,
    "shouldNotify": bool,
    "notifyType": int,
    "location": {object(Location)}
}
```

**Field description**:  

| Field            | Type        | Description                                                         |
| ------------ | -------- | ------------------------------------------------------------ |
| isWholeDay   | bool     | Whether it is a day or a specific hour                                             |
| timestamp    | int      | If set to a time, it must be on the hour or the half-hour. If IsWholeDay is not true, the time stamp formula is H:59:59, with H set based on the time zone. |
| shouldNotify | bool     | Whether or not to notify                                                     |
| notifyType   | int      | The notification triggering mechanism. When isWholeDay is false: 0: now, 1: 2 min before, 2: 15 min before, 3: 30 min before, 4: 1 hr before, 5: 2 hr before, 6: 1 day before, 7: 2 days before. When isWholeDay is true: 50: 09:00 that day, 51: 09:00 the day before, 52: 09:00 two days before, 53: 09:00 the Monday before |
| location | Location | The element location, can't be written, only returned | # Equation

LaTeX formula.

```json
{
    "equation": string,
    "location": {object(Location)}
}
```

**Field description**:  

| Field            | Type        | Description                                                         |
| -------- | -------- | ----------------------------------------------- |
| equation | string   | The LaTeX formula, must follow LaTeX syntax, such as: "E=mc^2" |
| location | Location | The element location, can't be written, only returned | # UndefinedElement

Unsupported row element

```json
{
    "location": {object(Location)}
}
```

**Field description**:  

| Field            | Type        | Description                                                         |
| -------- | -------- | ---------------------------- |
| location | Location | The element location, can't be written, only returned | # Gallery

Image.

> 

```json
{
    "galleryStyle": {object(GalleryStyle)},
    "imageList": [{object(ImageItem)}],
    "location": {object(Location)}
}
```

**Field description**:  

| Field            | Type        | Description                                                         |
| ------------ | ------------ | ------------------------------------------------------------ |
| galleryStyle | GalleryStyle | Image style, currently only supports alignment methods                             |
| imageList    | []ImageItem  | The image object. If multiple images are input at once, they are displayed in the same row. A single row can contain up to 6 images. |
| location | Location | The element location, can't be written, only returned | # GalleryStyle

Image style, currently only supports alignment methods.

```json
{
    "align": string
}
```

**Field description**:  

| Field            | Type        | Description                                                         |
| ----- | ------ | ------------------------------------------------------------ |
| align | string | Image alignment method, only valid when the row contains only one image. Center (default): centerAlign left: leftAlign right: right | # ImageItem

A single write action can write up to 50 images and attachments. Images can't exceed 20 MB. Local images must first be uploaded through Upload materials or Multipart upload materials. Supported image formats: jpg, jpeg, bmp, png, and gif.

**Field description**:  

| Field            | Type        | Description                                                         |
| --------- | ------ | ------------------------------------------------------- |
| fileToken | string | Image token, such as: boxcnOj88GDkmWGm2zsTyCBqoLb, can't be edited |
| width     | float  | Image width, in px                                          |
| height     | float  | Image height, in px                                          | > You can specify the image width and height. To ensure images are displayed properly, the width and height are checked to see that they maintain the original aspect ratio. Therefore, the final display dimensions may be slightly different than those manually set.

# File

A single write action can write up to 50 images and attachments. Local files must first be uploaded through Upload materials or Multipart upload materials.

```json
{
    "fileToken": string,
    "viewType": string,
    "fileName": string,
    "location": {object(Location)}
}
```

**Field description**:  

| Field            | Type        | Description                                                         |
| --------- | -------- | ------------------------------------------------------------ |
| fileToken | string | Attachment token, such as: boxcnOj88GDkmWGm2zsTyCBqoLb, can't be edited |
| viewType  | string   | File display style. Preview: previewCard: cardIn a row: inline |
| fileName  | string   | File name, can't be edited                                           |
| location | Location | The element location, can't be written, only returned | # HorizontalLine

Horizontal divider.

```json
{
    "location": {object(Location)}
}
```

**Field description**:  

| Field            | Type        | Description                                                         |
| -------- | -------- | ---------------------------- |
| location | Location | The element location, can't be written, only returned | # EmbeddedPage

Embedded webpage.

```json
{
    "type": string,
    "url" : string,
    "width": float,
    "height": float,
    "location": {object(Location)}
}
```

**Field description**:  

| Field            | Type        | Description                                                         |
| -------- | -------- | ------------------------------------------------------------ |
| type     | string   | Supports the following webpages: bilibili: "bilibili"Xigua Video: "xigua"Youku: "youku"Airtable: "airtable"Baidu Maps: "baidumap" |
| url  | string | Hyperlink URL of a third-party webpage, query escape encoding must be used for read and write, for example: https%3A%2F%2Fwww.baidu.com%2Fmaps%2Fembed%2Fv1%2Fplace%3Fkey%3DAIzaSyAfJZc8JxNRe909WC_QBILdlM55NqGnI30%26q%3DCentral%2BPark%26center%3D40.7828647%2C-73.9675438%26zoom%3D17z%26language%3Den |
| width    | float    | Window width, defaults to the default value if not specified. If specified, the height is calculated based on the default aspect ratio. |
| height    | float    | Window height, defaults to the default value if not specified |
| location | Location | The element location, can't be written, only returned | # ChatGroup

```json
{
    "openChatId": string,
    "location": {object(Location)}
}
```

**Field description**:  

| Field            | Type        | Description                                                         |
| ---------- | -------- | --------------------------------------------------------- |
| openChatId | string   | openId of group chat session, such as: oc_4149593da6ef5fe4de16b10cb4769c94When a document is copied, this is replaced with "none" if the user doesn't have permission to view the group card.During write actions, a scope error is returned if the value is "none" or the user is not a group member. |
| location | Location | The element location, can't be written, only returned | # Table

Normal Docs sheet. To learn about sheets and bitables, see the Sheet and Bitable sections.

Use one of the two supported creation methods:

1. Create a sheet with content by providing rowSize, columnSize, tableRows, tableStyle, and mergedCells information
2. Create an empty sheet by specifying rowSize and columnSize.

```json
{    
    "tableId": string,
    "rowSize": int,
    "columnSize": int,
    "tableRows": [{object(TableRow)}],
    "tableStyle": {object(TableStyle)},
    "mergedCells": [{object(MergedCell)}],
    "location": {object(Location)}
}
```

**Field description**:  

| Field            | Type        | Description                                                         |
| ----------- | -------- | ------------------------------------------------------- |
| tableId     | string   | Sheet ID, can't be edited                                        |
| rowSize     | int      | Sheet row quantity, maxpoint during empty sheet creation: 9                        |
| columnSize     | int      | Sheet column quantity, maxpoint during empty sheet creation: 9                        |
| tableRows   | []object | Sheet row content, must be consistent with rowSize and columnSize when creating a sheet with content |
| tableStyle  | object   | Sheet style                                                |
| mergedCellId | []object    | Merged cell information                                                |
| location | Location | The element location, can't be written, only returned | # TableRow

A row of a normal sheet.

```json
{
    "rowIndex": int,
    "tableCells":[{object(TableCell)}]
}
```

**Field description**:  

| Field            | Type        | Description                                                         |
| ---------- | -------- | --------------------------------------- |
| rowIndex   | int      | Row index, starts from 0, can't be edited. The first row is 0. |
| tableCells | []object | Sheet cell content                          | # TableCell

A cell of a normal sheet.

```json
{
    "columnIndex": int,
    "zoneId": string,
    "body": {object(Body)}
}
```

**Field description**:  

| Field            | Type        | Description                                                         |
| ----------- | ------ | --------------------------------------------------- |
| columnIndex   | int      | Column index, starts from 0, can't be edited. The first column is 0. |
| zoneId      | string | Cell ID                                            |
| body        | object | Cell content, supports blocks other than table, sheet, and bitable. | # TableStyle

Sheet style, currently only allows you to adjust the column width of the sheet.

```json
{
    "tableColumnProperties": [{object(TableColumnProperties)}]
}
```

**Field description**:  

| Field            | Type        | Description                                                         |
| --------------------- | -------- | ------ |
| tableColumnProperties | []object | Column properties | # TableColumnProperties

Normal sheet column width.

```json
{
    "width": int
}
```

**Field description**:  

| Field            | Type        | Description                                                         |
| ----- | ---- | ----------------------------------------------- |
| width | int  | Column width, in px. Minpoint: 50, maxpoint: 1,300, default value: 100. | # MergedCell

Merged cell.

```json
{
    "mergedCellId": string,
    "rowStartIndex": int,
    "rowEndIndex": int,
    "columnStartIndex": int,
    "columnEndIndex": int
}
```

**Field description**:  

| Field            | Type        | Description                                                         |
| ---------------- | ------ | ------------------------------- |
| mergedCellId | string    | ID of the merged cell, can't be edited.                                             |
| rowStartIndex | int    | Merged cell row start index, starts from 0                                    |
| rowEndIndex   | int    | Merged cells row end index                                               |
| columnStartIndex | int    | Merged cell column start index, starts from 0                                    |
| columnEndIndex   | int    | Merged cell column end index                                               | # Sheet

Sheet in a document. Use the Obtain rich text content of a document to obtain the sheet token in token_tableId format. Then, use this token to call the Sheets API. Be sure to separate the token and ID.
Read sheet content API. Use one of the supported creation methods:

1. Perform a deep copy based on the token. The token can only be used for a sheet in a document, not an independent sheet.
2. Create an empty sheet by specifying rowSize and columnSize.

```json
{
    "token": string,
    "rowSize": int,
    "columnSize": int,
    "location": {object(Location)}
}
```

**Field description**:  

| Field            | Type        | Description                                                         |
| ---------- | -------- | ------------------------------------------------------------ |
| token      | string   | Sheet token, such as: shtcnKjLimyn2fW0uGseasQYGgh_m7fMrN. The part before the underscore is the sheet token, and the part after the underscore is the tableId. Be sure the split up the token and ID. Input the token when calling the API to perform a deep copy. |
| rowSize     | int      | Sheet row quantity, maxpoint during empty sheet creation: 9                        |
| columnSize     | int      | Sheet column quantity, maxpoint during empty sheet creation: 9                        |
| location | Location | The element location, can't be written, only returned | # Bitable

Bitable in a document. Use the Obtain rich text content of a document to obtain the token. Read bitable content API. Use one of the two supported creation methods:

1. Perform a deep copy based on the token+viewType. The token can only be used for a bitable in a document, not an independent bitable.
2. Create an empty bitable by specifying a viewType but not a token.

```json
{
    "token": string,
    "viewType": string,
    "location": {object(Location)}
}
```

**Field description**:  

| Field            | Type        | Description                                                         |
| -------- | -------- | ------------------------------------------------------------ |
| token    | string   | Bitable token, such as: shtcnKjLimyn2fW0uGseasQYGgh_m7fMrN. When you don't use a space when calling the create API or insert API, a deep copy is performed. This can't be edited. |
| viewType | string   | Grid: grid Kanban: kanban                                |
| location | Location | The element location, can't be written, only returned | # Diagram

Diagram, including flowcharts and UML diagrams. Diagram write actions are not currently supported.

```json
{
    "token": string,
    "diagramType": string,
    "location": {object(Location)}
}
```

**Field description**:  

| Field            | Type        | Description                                                         |
| ----------- | -------- | ---------------------------------------------------------- |
| token       | string   | Diagram token, such as: diacnK1MYEHBopBbIdc6A5AOVCh, can't be edited |
| diagramType | string   | Diagram type.Flowchart: flowchartUML diagram: uml              |
| location | Location | The element location, can't be written, only returned | # Jira

Jira, including Jira filters and Jira issues. Jira write actions are not currently supported.

```json
{
    "token": string,
    "jiraType": string,
    "location": {object(Location)}
}
```

**Field description**:  

| Field            | Type        | Description                                                         |
| -------- | -------- | -------------------------------------------------------- |
| token    | string   | Jira token, jftcnsA0fY8V3CzYvtRPy9XsXxf, can't be edited |
| jiraType | string   | Jira type.Filter: filterIssue: issue             |
| location | Location | The element location, can't be written, only returned | # Poll

Poll. Poll write actions are not currently supported.

```json
{
    "token": string,
    "location": {object(Location)}
}
```

**Field description**:  

| Field            | Type        | Description                                                         |
| -------- | -------- | -------------------------------------------------------- |
| token    | string   | Poll token, such as: jsncnxuT7beirSpf33NfcKrSwAh, can't be edited |
| location | Location | The element location, can't be written, only returned | # Code
```json
{
    "language": string,
    "wrapContent": bool,
    "body": {object(Body)},
    "zoneId": string,
    "location": {object(Location)}
}
```
**Field description**:  

| Field            | Type        | Description                                                         |
| -------- | -------- | -------------------------------------------------------- |
| language    | string   | Coding language Supported languages: Plain Text, ABAP, Ada, Apache, Apex, Assembly language, Bash, C, C#, C++, COBOL, CSS, CoffeeScript, D, Dart, Delphi, Dockerfile, Django, Erlang, Fortran, FoxPro, Go, Groovy, HTML, HTMLBars, HTTP, Haskell, JSON, Java, JavaScript, Julia, Kotlin, LaTeX, Lisp, Logo, Lua, MATLAB, Makefile, Markdown, Nginx, Objective-C, OpenEdge ABL, PHP, Perl, PostScript, PowerShell, Prolog, ProtoBuf, Python, R, RPG, Ruby, Rust, SAS, SCSS, SQL, Scala, Scheme, Scratch, Shell, Swift, Thrift, TypeScript, VBScript, Visual Basic, XML, YAML |
| wrapContent | bool | Whether to automatically wrap content                             |
| body | Object | Code block content (only supported Paragraph block)                     |
| zoneId | string | Code block ID (read only), required to add, modify, and delete code blocks                   |
| location | Location | The element location, can't be written, only returned | # Callout
```json
{
	"calloutEmojiId": string,										
	"calloutBackgroundColor": {object(RGBColor)},	
	"calloutBorderColor" {object(RGBColor)},
	"CalloutTextColor" {object(RGBColor)}, 
	"body": {object(Body)},
	"zoneId": string,
	"location": {object(Location)}
}
```

| fields                 | type     | Remark |
| ---------------------- | -------- | ------ |
| calloutEmojiId         | string   | The following callout emojis are supported: [emoji snum table](https://bytedance.larksuite.com/sheets/shtcnkDBCFZJDyGliM7IOqPuSgd?sheet=7FyRq7)    |
| calloutBackgroundColor | RGBColor    | Callout background color (divided into dark colors and light colors, corresponding to dark and light border colors) |
| calloutBorderColor     | RGBColor    | Callout border color (divided into dark colors and light colors, corresponding to dark and light background colors) |
| calloutTextColor		 | RGBColor	| Callout text color | 
| body                   | Body   | Callout content (only supports Paragraph Block and HorizontalLine Block; ParagraphStyle supports all except Collapse) |
| zoneId                 | string | Callout ID (read only), required to add, modify, and delete code block content. |
| location               | Location | Callout location (read only) | | Font/Background/Border | enum                                                        | Color scheme                                                         | rgb                                                          | Alpha |
| -------------- | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ----- |
| Text color       | *CalloutRedText**CalloutOrangeText**CalloutYellowText**CalloutGreenText**CalloutBlueText**CalloutPurpleText**CalloutNeutralText* | None                                                           | rgb(216,57,49)rgb(216,57,49)rgb(220,155,4)rgb(46,161,33)rgb(36,91,219)rgb(100,37,208)rgb(143,149,158) | 1.0   |
| Background color         | *CalloutLightRedBackground**CalloutLightOrangeBackground**CalloutLightYellowBackground**CalloutLightGreenBackground**CalloutLightBlueBackground**CalloutLightPurpleBackground**CalloutLightNeutralBackground**CalloutDarkRedBackground**CalloutDarkOrangeBackground**CalloutDarkYellowBackground**CalloutDarkGreenBackground**CalloutDarkBlueBackground**CalloutDarkPurpleBackground**CalloutDarkNeutralBackground* | Light Dark                 | rgb(254,241,241)rgb(255,245,235)rgb(254,255,240)rgb(240,251,239)rgb(240,244,255)rgb(246,241,254)rgb(248,249,250)rgb(253,226,226)rgb(254,234,210)rgb(255,255,204)rgb(217,245,214)rgb(225,234,255)rgb(236,226,254)rgb(245,246,247) | 1.0   |
| Border color         | *CalloutRedBorder**CalloutOrangeBorder**CalloutYellowBorder**CalloutGreenBorder**CalloutBlueBorder**CalloutPurpleBorder**CalloutNeutralBorder* | LightDarkLightDarkLightDarkLightDarkLightDarkLightDarkLightDark | rgb(251,191,188)rgb(249,142,139)rgb(254,212,164)rgb(255,186,107)rgb(255,246,122)rgb(255,242,88)rgb(183,237,177)rgb(142,224,133)rgb(186,206,253)rgb(130,167,252)rgb(205,178,250)rgb(173,130,247)rgb(222,224,227)rgb(187,191,196) | 1.0   | # DocsApp

Team interaction app. App write actions are not currently supported.

```json
{
    "typeId": string,
    "instanceId": string,
    "location": {object(Location)}
}
```

**Field description**:  

| Field            | Type        | Description                                                         |
| -------- | -------- | -------------------------------------------------------- |
| typeId    | string   |  Team interaction app type, such as: Information collection "blk_5f992038c64240015d280958"|
|instanceId|string	|Unique ID of team interaction app|
| location | Location | The element location, can't be written, only returned | # UndefinedBlock

Unsupported block types

```json
{
    "location": {object(Location)}
}
```

**Field description**:  

| Field            | Type        | Description                                                         |
| -------- | -------- | ---------------------------- |
| location | Location | The element location, can't be written, only returned |
