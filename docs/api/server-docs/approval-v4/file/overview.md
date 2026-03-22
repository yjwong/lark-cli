---
title: "Resource introduction"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/file/overview"
service: "approval-v4"
resource: "file"
section: "Approval"
updated: "1675167408000"
---

# Resource introduction
When there is a picture or attachment control in the review form, the developer needs to upload the file to the review system through the review upload file interface before creating the review instance.

| Parameter | Type | Description |
| --- | --- | --- |
| `name` | `string` | File name (need to include file extension **Example value**："123.doc" |
| `type` | `string` | File type **Example value**："attachment" **Optional values are**：  - `attachment`: attachment - `image`: image  |
| `content` | `file` | File **Example value**：123.doc | ## Data example
```json
{
	"name":"123.doc",
	"type":"attachment",
	"content":123.doc
}
```
