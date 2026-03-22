---
title: "Document Structure Introduction"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uAzM5YjLwMTO24CMzkjN"
section: "Deprecated Version (Not Recommended)"
updated: "1695369897000"
---

# Document data structure overview

This document introduces the data structure used by Lark Docs, including the different elements that compose Docs and the relations between the elements.

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/1ffbede126855e54a20779add375e638_FOq9UlUxNl.png?height=313&lazyload=true&width=633)

# Document

The level-1 Docs structure is composed of several parts:

```js 
document: {
    body: ...,
    title: ...
}
```

# Body

Body describes all the body text, which is composed of various blocks.

# Block
Each block is a group of structured elements. From the user's perspective, each block is a separate row of content. For example, a paragraph of text and an image in a document each have their own block.

```js 
{
    "paragraph": {object(Paragraph)},
    "horizontalLine": {object(HorizontalLine)},
    "embeddedPage": {object(EmbeddedPage)},
    "chatGroup": {object(ChatGroup)},
    "table": {object(Table)},
    "sheet": {object(Sheet)},
    "bitable": {object(Bitable)}
} 
```

# Location

We use a zoneId, startIndex, and endIndex to identify the location of each block in the body.
The different areas of a document are identified by zoneIds. The zoneId of the body text of the document is always "0". The zoneIds of table cells are irregular strings.
UTF-16 is used to encode the startIndex and endIndex and identify the start and end location offsets of a block. Each Block has a corresponding index to mark its absolute position.
Because Lark Docs supports collaborative editing, when you access a document and submit edit requests, the document may be modified by other users simultaneously, causing the offset to change. If modifications are made based on the reading status, text accuracy may be impacted. To avoid the failure of submission or incorrect results caused by the above problem, Lark Docs will automatically calculate the offset between the last request and the latest online content so developers don't have to perform additional work.

# Paragraph

Paragraph is a paragraph of the document body. It is composed of ParagraphElement (the paragraph element), Location, and ParagraphStyle (the paragraph style). The Paragraph structure can be described by the following data structure:

```js 
{
    "elements":  [{object (ParagraphElement)}],
    "loc": {object(Location)}
    "style": object (ParagraphStyle),
} 
```

> This is a sample text to describe a bold text, which is a common component style.
>
> - This is a bulleted list.
> - This bulleted list contains three steps.
> - Each bulleted list step is an independent paragraph.
>
> You can also use the API to insert such content.

The following figure graphically shows the format of the above text:

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/fe6981488134f643efcecc3b3c236cb1_N2b3fH02gj.png?height=233&lazyload=true&width=773)

This document introduces the data structure of documents. Next, you can read Document Data Structure Reference to learn more.
