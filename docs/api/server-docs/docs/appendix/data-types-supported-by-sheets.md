---
title: "Data Types Supported By Sheets"
url: "https://open.larksuite.com/document/ukTMukTMukTM/ugjN1UjL4YTN14CO2UTN"
section: "Docs"
updated: "1647001439000"
---

# Data types supported by sheets

## String

```
"string"
```

## Number

```
123
```

## Link

### No text link

```
"http://www.dd.com"
```
### Link with text

```json
{
    "text": "Text",
    "link": "http://www.dd.com",
    "type": "url"
}
```

## Mailbox

```
"aaa@aa.com"
```

## Mention User

Only support users of same tenant;A single request support up to @50 people at the same time;

notify: whether to send lark messages, users without reading permission will not receive lark messages;

grantReadPermission: Whether to grant the user read permission (only support this field in a separate table);

textType: specify the incoming content of the text field, optional email, openId, unionId;

text: the information of the person who needs @, specified by textType

```json
{
    "type": "mention",
    "text": "aaa@aa.com",
    "textType": "email",
    "notify": true,
    "grantReadPermission": true
}
```

## Formula

The text field is the corresponding formula, does not support cross-table reference formulas (IMPORTRANGE)

```json
{
    "type": "formula",
    "text": "=A1"
}
```

## Mention Document

textType: fixed as fileToken

text: document token

objType: document type, optional sheet, doc, slide, bitable, mindnote

```json
{
    "type": "mention",
    "textType": "fileToken",
    "text": "shtxxxx",
    "objType": "sheet"
}
```

## Drop-down list

Values is an array, which can be filled with bool, string, and number types. String type data cannot contain ",". Before use, you need to use the set dropdown list interface to set the dropdown list.

```json
{
    "type": "multipleValue",
    "values": [
        1,
        "test"
    ]
}
```

## Segment style

### segmentStyle

bold: whether to bold

italic: whether it is italic

strikeThrough: whether it contains strikethrough

underline: whether it contains an underline

foreColor: font color, hexadecimal rgb color

fontSize: font size, the smallest font size is 9, the largest font size is 36

```json
{
    "bold": true,
    "italic": true,
    "strikeThrough": true,
    "underline": true,
    "foreColor": "#ff00ff",
    "fontSize": 30
}
```

### Write data with segment style

#### String

```json
{
    "text": "string",
    "type": "text",
    "segmentStyle": {
        "bold": true,
        "italic": true,
        "strikeThrough": true,
        "underline": true,
        "foreColor": "#ff00ff",
        "fontSize": 20
    }
}
```

#### Number

Does not support segment styles

#### Link

Supports writing different styles for different strings in the link. The foreColor field does not take effect and is fixed to blue; the underline field does not take effect, and it always contains an underline;

```json
{
    "text": "Text",
    "link": "http://www.dd.com",
    "type": "url",
    "texts": [
        {
            "text": "文",
            "segmentStyle": {
                "bold": true,
                "italic": true,
                "strikeThrough": true,
                "underline": true,
                "foreColor": "#ffffff",
                "fontSize": 20
            }
        },
        {
            "text": "本",
            "segmentStyle": {
                "bold": true,
                "italic": false,
                "strikeThrough": true,
                "underline": true,
                "foreColor": "#ffffff",
                "fontSize": 10
            }
        }
    ]
}
```

#### Mailbox

Supports writing different styles for different strings in the mailbox. The foreColor field does not take effect and is fixed to blue; the underline field does not take effect, and it always contains an underline;

```json
{
    "type": "url",
    "text": "aa@bytedance.com",
    "texts": [
        {
            "text": "aa",
            "segmentStyle": {
                "bold": true,
                "italic": true,
                "strikeThrough": true,
                "underline": true,
                "foreColor": "#ffffff",
                "fontSize": 20
            }
        },
        {
            "text": "@bytedance.com",
            "segmentStyle": {
                "bold": true,
                "italic": false,
                "strikeThrough": true,
                "underline": true,
                "foreColor": "#ffffff",
                "fontSize": 10
            }
        }
    ]
}
```

#### Mention users

Only support metion users of same tenant;A single request support mention up to 50 users at the same time

notify: whether to send lark messages, users without reading permission will not receive lark messages;

grantReadPermission: Whether to grant the user read permission (only support this field in a separate table);

textType: specify the incoming content of the text field, optional email, openId, unionId;

text: the information of the person who needs @, specified by textType;

Only supports setting partial styles for the whole; the foreColor field does not take effect and is fixed to blue;

```json
{
    "type": "mention",
    "text": "aaa@aa.com",
    "textType": "email",
    "notify": true,
    "grantReadPermission": true
    "segmentStyle": {
        "bold": true,
        "italic": true,
        "strikeThrough": true,
        "underline": true,
        "foreColor": "#ff00ff",
        "fontSize": 30
    }
}
```

#### Mention Docs

textType: fixed as fileToken

text: document token

objType: document type, optional sheet, doc, slide, bitable, mindnote

The foreColor field does not take effect and is fixed to blue;

```json
{
    "type": "mention",
    "textType": "fileToken",
    "text": "shtxxxx",
    "objType": "sheet",
    "segmentStyle": {
        "bold": true,
        "italic": true,
        "strikeThrough": true,
        "underline": true,
        "foreColor": "#ff00ff",
        "fontSize": 30
    }
}
```

#### Drop-down list

Does not support segment styles
