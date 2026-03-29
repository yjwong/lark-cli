---
title: "Create OKR progress record"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/okr-v1/progress_record/create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/okr/v1/progress_records"
service: "okr-v1"
resource: "progress_record"
section: "OKR"
scopes:
  - "okr:okr"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1697014721000"
---

# Create OKR progress record

Create OKR progress record

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/okr/v1/progress_records |
| HTTP Method | POST |
| Supported app types | custom |
| Required scopes | `okr:okr` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: "open_id" **Optional values are**:  - `open_id`: user's open id - `union_id`: user's union id - `user_id`: user's user id  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `source_title` | `string` | Yes | Sources of progress **Example value**: "2021.12.20~2021.12.26 周报" |
| `source_url` | `string` | Yes | Progress Source Link **Example value**: "www.baidu.com" **Data validation rules**: - Regular expression: `^https?://.*$` |
| `target_id` | `string` | Yes | Target id, corresponding to target_type **Example value**: "7041430377642082323" |
| `target_type` | `int` | Yes | Target type **Example value**: 1 **Optional values are**:  - `2`: Objective - `3`: Key Result  |
| `content` | `content_block` | Yes | Details of progress, in rich text format |
|   `blocks` | `content_block_element[]` | No | The document structure is arranged in rows, each row of content is a Block |
|     `type` | `string` | No | Document element type **Example value**: "paragraph" **Optional values are**:  - `paragraph`: Text paragraph - `gallery`: Picture  |
|     `paragraph` | `content_paragraph` | No | Text paragraph |
|       `style` | `content_paragraph_style` | No | Paragraph style |
|         `list` | `content_list` | No | Numbered list/unordered list/task list |
|           `type` | `string` | No | List type **Example value**: "number" **Optional values are**:  - `number`: Numbered list - `bullet`: Unordered list - `checkBox`: Task list - `checkedBox`: List of completed tasks - `indent`: Tab indentation  |
|           `indentLevel` | `int` | No | The indentation level of the list, support to specify the indentation of a line, except for the code block list support to set the indentation, support 1-16 indentation, the value range: [1,16] **Example value**: 1 |
|           `number` | `int` | No | The line number used to specify the list, which is only effective for numbered list and code blocks. If indentation is set for numbered list, the line number may appear as letters or Roman numerals **Example value**: 1 |
|       `elements` | `content_paragraph_element[]` | No | Paragraph elements form a paragraph |
|         `type` | `string` | No | Element type **Example value**: "textRun" **Optional values are**:  - `textRun`: Text element - `docsLink`: Document link element - `person`: @User type element  |
|         `textRun` | `content_text_run` | No | Text |
|           `text` | `string` | No | Specific text content **Example value**: "周报内容" |
|           `style` | `content_text_style` | No | Styling of text content, support for BIUS, colors, etc |
|             `bold` | `boolean` | No | Bold or not **Example value**: true |
|             `strikeThrough` | `boolean` | No | Whether to delete **Example value**: true |
|             `backColor` | `content_color` | No | Background color |
|               `red` | `int` | No | Red, value range [0,255] **Example value**: 216 |
|               `green` | `int` | No | Green, value range [0,255] **Example value**: 191 |
|               `blue` | `int` | No | Blue, value range [0,255] **Example value**: 188 |
|               `alpha` | `float` | No | Transparency, value range [0,1] **Example value**: 0.1 |
|             `textColor` | `content_color` | No | Font color |
|               `red` | `int` | No | Red, value range [0,255] **Example value**: 216 |
|               `green` | `int` | No | Green, value range [0,255] **Example value**: 191 |
|               `blue` | `int` | No | Blue, value range [0,255] **Example value**: 188 |
|               `alpha` | `float` | No | Transparency, value range [0,1] **Example value**: 0.1 |
|             `link` | `content_link` | No | Link address |
|               `url` | `string` | No | Link address **Example value**: "https://www.baidu.com/" |
|         `docsLink` | `content_docs_link` | No | Lark Cloud Document |
|           `url` | `string` | No | Lark cloud document link address **Example value**: "https://www.baidu.com/" |
|           `title` | `string` | No | Lark Cloud Document Title **Example value**: "百度" |
|         `person` | `content_person` | No | @User |
|           `openId` | `string` | No | Employee OpenID **Example value**: "ou_3bbe8a09c20e89cce9bff989ed840674" |
|     `gallery` | `content_gallery` | No | Picture |
|       `imageList` | `content_image_item[]` | No | Picture elements |
|         `fileToken` | `string` | No | Image token, obtained through the upload image interface **Example value**: "boxcnOj88GDkmWGm2zsTyCBqoLb" |
|         `src` | `string` | No | src **Example value**: "https://example/drive/home/" |
|         `width` | `float` | No | Picture width in px **Example value**: 458 |
|         `height` | `float` | No | Picture height, unit px **Example value**: 372 | ### Request body example

{ "source_title":"abc", "source_url":"https://www.baidu.com/", "target_id":"7043693679567028244", "target_type":2, "content":{ "blocks":[ { "type":"paragraph", "paragraph":{ "elements":[ { "type":"textRun", "textRun":{ "text":"abc", "style":{ "bold":true } } } ] } }, { "type":"paragraph", "paragraph":{ "elements":[ { "type":"textRun", "textRun":{ "text":"*", "style":{ "bold":true } } }, { "type":"textRun", "textRun":{ "text":"abc", "style":{ "strikeThrough":true } } } ] } }, { "type":"paragraph", "paragraph":{ "elements":[ { "type":"textRun", "textRun":{ "text":"abc", "style":{ "textColor":{ "red":216, "green":57, "blue":49, "alpha":1 } } } } ] } }, { "type":"paragraph", "paragraph":{ "elements":[ { "type":"textRun", "textRun":{ "text":"abc", "style":{ "backColor":{ "red":251, "green":191, "blue":188, "alpha":1 } } } } ] } }, { "type":"paragraph", "paragraph":{ "elements":[ { "type":"textRun", "textRun":{ "text":"abc", "style":{ "bold":true, "strikeThrough":true, "backColor":{ "red":251, "green":191, "blue":188, "alpha":1 }, "textColor":{ "red":216, "green":57, "blue":49, "alpha":1 } } } } ] } }, { "type":"paragraph", "paragraph":{ "elements":[ { "type":"textRun", "textRun":{ "text":"abc", "style":{ } } }, { "type":"textRun", "textRun":{ "text":"abc", "style":{ "bold":true } } } ], "style":{ "list":{ "type":"number", "indentLevel":1, "number":1 } } } }, { "type":"paragraph", "paragraph":{ "elements":[ { "type":"textRun", "textRun":{ "text":"abc", "style":{ } } }, { "type":"textRun", "textRun":{ "text":"abc", "style":{ "strikeThrough":true } } } ], "style":{ "list":{ "type":"number", "indentLevel":1, "number":2 } } } }, { "type":"paragraph", "paragraph":{ "elements":[ { "type":"textRun", "textRun":{ "text":"abc", "style":{ } } }, { "type":"textRun", "textRun":{ "text":"abc", "style":{ "backColor":{ "red":251, "green":191, "blue":188, "alpha":1 }, "textColor":{ "red":216, "green":57, "blue":49, "alpha":1 } } } } ], "style":{ "list":{ "type":"number", "indentLevel":1, "number":3 } } } }, { "type":"paragraph", "paragraph":{ "elements":[ { "type":"textRun", "textRun":{ "text":"abc", "style":{ } } }, { "type":"textRun", "textRun":{ "text":"abc", "style":{ "bold":true } } } ], "style":{ "list":{ "type":"bullet", "indentLevel":1 } } } }, { "type":"paragraph", "paragraph":{ "elements":[ { "type":"textRun", "textRun":{ "text":"abc", "style":{ } } }, { "type":"textRun", "textRun":{ "text":"abc", "style":{ "strikeThrough":true } } } ], "style":{ "list":{ "type":"bullet", "indentLevel":1 } } } }, { "type":"paragraph", "paragraph":{ "elements":[ { "type":"textRun", "textRun":{ "text":"abc", "style":{ } } }, { "type":"textRun", "textRun":{ "text":"abc", "style":{ "backColor":{ "red":251, "green":191, "blue":188, "alpha":1 }, "textColor":{ "red":216, "green":57, "blue":49, "alpha":1 } } } } ], "style":{ "list":{ "type":"bullet", "indentLevel":1 } } } }, { "type":"paragraph", "paragraph":{ "elements":[ { "type":"textRun", "textRun":{ "text":"https://example/docx/doxcnO2Wkq0YPZQYLuJKyyOvLrh#doxcnSOui82swqk6c0o436Ak3nc", "style":{ "link":{ "url":"https://example/docx/doxcnO2Wkq0YPZQYLuJKyyOvLrh#doxcnSOui82swqk6c0o436Ak3nc" } } } } ] } }, { "type":"gallery", "gallery":{ "imageList":[ { "src":"https://internal-api-okr.feishu-boe.cn/stream/api/downloadFile/?file_token=boxbcMTBQO9ofLjWkDuPxkxOA2c&ticket=eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ0YXJnZXRfaWQiOiI3MDQxNDMwMzc3NjQyMDgyMzIzIiwidGFyZ2V0X3R5cGUiOjMsImFjdGlvbiI6MiwiZmlsZV90b2tlbiI6ImJveGJjTVRCUU85b2ZMaldrRHVQeGt4T0EyYyIsInVzZXJfaWQiOiI2OTY5ODU1NTAxNzQ0ODM0MDkyIiwidGVuYW50X2lkIjoiNjg3NzUwMjY4NzYwOTQwNjk5MCIsImV4cCI6MTY0MDE1NTk2M30.yc4qV2pkGUVwSO53-N_XGgeMucjmDn9iso1Ez_8vpghFz8YdeSDf4NHQpxOHYHc8RURvwI0a5UTNKKJ9CWagTQ", "fileToken":"boxbcMTBQO9ofLjWkDuPxkxOA2c", "width":458, "height":372 } ] } }, { "type":"paragraph", "paragraph":{ "elements":[ { "type":"textRun", "textRun":{ "text":"abc", "style":{ } } } ], "style":{ "list":{ "type":"checkBox", "indentLevel":1 } } } }, { "type":"paragraph", "paragraph":{ "elements":[ { "type":"textRun", "textRun":{ "text":"abc", "style":{ } } } ], "style":{ "list":{ "type":"checkedBox", "indentLevel":1 } } } } ] } }

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `progress_record` | \- |
|   `progress_id` | `string` | OKR progress ID |
|   `modify_time` | `string` | Progress update time, milliseconds |
|   `content` | `content_block` | Progress, Corresponding Content Details |
|     `blocks` | `content_block_element[]` | The document structure is arranged in rows, each row of content is a Block |
|       `type` | `string` | Document element type **Optional values are**:  - `paragraph`: Text paragraph - `gallery`: Picture  |
|       `paragraph` | `content_paragraph` | Text paragraph |
|         `style` | `content_paragraph_style` | Paragraph style |
|           `list` | `content_list` | Numbered list/unordered list/task list |
|             `type` | `string` | List type **Optional values are**:  - `number`: Numbered list - `bullet`: Unordered list - `checkBox`: Task list - `checkedBox`: List of completed tasks - `indent`: Tab indentation  |
|             `indentLevel` | `int` | The indentation level of the list, support to specify the indentation of a line, except for the code block list support to set the indentation, support 1-16 indentation, the value range: [1,16] |
|             `number` | `int` | The line number used to specify the list, which is only effective for numbered list and code blocks. If indentation is set for numbered list, the line number may appear as letters or Roman numerals |
|         `elements` | `content_paragraph_element[]` | Paragraph elements form a paragraph |
|           `type` | `string` | Element type **Optional values are**:  - `textRun`: Text element - `docsLink`: Document link element - `person`: @User type element  |
|           `textRun` | `content_text_run` | Text |
|             `text` | `string` | Specific text content |
|             `style` | `content_text_style` | Styling of text content, support for BIUS, colors, etc |
|               `bold` | `boolean` | Bold or not |
|               `strikeThrough` | `boolean` | Whether to delete |
|               `backColor` | `content_color` | Background color |
|                 `red` | `int` | Red, value range [0,255] |
|                 `green` | `int` | Green, value range [0,255] |
|                 `blue` | `int` | Blue, value range [0,255] |
|                 `alpha` | `float` | Transparency, value range [0,1] |
|               `textColor` | `content_color` | Font color |
|                 `red` | `int` | Red, value range [0,255] |
|                 `green` | `int` | Green, value range [0,255] |
|                 `blue` | `int` | Blue, value range [0,255] |
|                 `alpha` | `float` | Transparency, value range [0,1] |
|               `link` | `content_link` | Link address |
|                 `url` | `string` | Link address |
|           `docsLink` | `content_docs_link` | Lark Cloud Document |
|             `url` | `string` | Lark cloud document link address |
|             `title` | `string` | Lark Cloud Document Title |
|           `person` | `content_person` | @User |
|             `openId` | `string` | Employee OpenID |
|       `gallery` | `content_gallery` | Picture |
|         `imageList` | `content_image_item[]` | Picture elements |
|           `fileToken` | `string` | Image token, obtained through the upload image interface |
|           `src` | `string` | src |
|           `width` | `float` | Picture width in px |
|           `height` | `float` | Picture height, unit px | ### Response body example

{ "code": 0, "data": { "content": { "blocks": [ { "paragraph": { "elements": [ { "textRun": { "style": { "bold": true }, "text": "abc" }, "type": "textRun" } ] }, "type": "paragraph" }, { "paragraph": { "elements": [ { "textRun": { "style": { "bold": true }, "text": "*" }, "type": "textRun" }, { "textRun": { "style": { "strikeThrough": true }, "text": "abc" }, "type": "textRun" } ] }, "type": "paragraph" }, { "paragraph": { "elements": [ { "textRun": { "style": { "textColor": { "alpha": 1, "blue": 49, "green": 57, "red": 216 } }, "text": "abc" }, "type": "textRun" } ] }, "type": "paragraph" }, { "paragraph": { "elements": [ { "textRun": { "style": { "backColor": { "alpha": 1, "blue": 188, "green": 191, "red": 251 } }, "text": "abc" }, "type": "textRun" } ] }, "type": "paragraph" }, { "paragraph": { "elements": [ { "textRun": { "style": { "backColor": { "alpha": 1, "blue": 188, "green": 191, "red": 251 }, "bold": true, "strikeThrough": true, "textColor": { "alpha": 1, "blue": 49, "green": 57, "red": 216 } }, "text": "abc" }, "type": "textRun" } ] }, "type": "paragraph" }, { "paragraph": { "elements": [ { "textRun": { "style": {}, "text": "abc" }, "type": "textRun" }, { "textRun": { "style": { "bold": true }, "text": "abc" }, "type": "textRun" } ], "style": { "list": { "indentLevel": 1, "number": 1, "type": "number" } } }, "type": "paragraph" }, { "paragraph": { "elements": [ { "textRun": { "style": {}, "text": "abc" }, "type": "textRun" }, { "textRun": { "style": { "strikeThrough": true }, "text": "abc" }, "type": "textRun" } ], "style": { "list": { "indentLevel": 1, "number": 2, "type": "number" } } }, "type": "paragraph" }, { "paragraph": { "elements": [ { "textRun": { "style": {}, "text": "abc" }, "type": "textRun" }, { "textRun": { "style": { "backColor": { "alpha": 1, "blue": 188, "green": 191, "red": 251 }, "textColor": { "alpha": 1, "blue": 49, "green": 57, "red": 216 } }, "text": "abc" }, "type": "textRun" } ], "style": { "list": { "indentLevel": 1, "number": 3, "type": "number" } } }, "type": "paragraph" }, { "paragraph": { "elements": [ { "textRun": { "style": {}, "text": "abc" }, "type": "textRun" }, { "textRun": { "style": { "bold": true }, "text": "abc" }, "type": "textRun" } ], "style": { "list": { "indentLevel": 1, "type": "bullet" } } }, "type": "paragraph" }, { "paragraph": { "elements": [ { "textRun": { "style": {}, "text": "abc" }, "type": "textRun" }, { "textRun": { "style": { "strikeThrough": true }, "text": "abc" }, "type": "textRun" } ], "style": { "list": { "indentLevel": 1, "type": "bullet" } } }, "type": "paragraph" }, { "paragraph": { "elements": [ { "textRun": { "style": {}, "text": "abc" }, "type": "textRun" }, { "textRun": { "style": { "backColor": { "alpha": 1, "blue": 188, "green": 191, "red": 251 }, "textColor": { "alpha": 1, "blue": 49, "green": 57, "red": 216 } }, "text": "abc" }, "type": "textRun" } ], "style": { "list": { "indentLevel": 1, "type": "bullet" } } }, "type": "paragraph" }, { "paragraph": { "elements": [ { "textRun": { "style": { "link": { "url": "https://example/docx/doxcnO2Wkq0YPZQYLuJKyyOvLrh#doxcnSOui82swqk6c0o436Ak3nc" } }, "text": "https://example/docx/doxcnO2Wkq0YPZQYLuJKyyOvLrh#doxcnSOui82swqk6c0o436Ak3nc" }, "type": "textRun" } ] }, "type": "paragraph" }, { "gallery": { "imageList": [ { "fileToken": "boxbcMTBQO9ofLjWkDuPxkxOA2c", "height": 372, "src": "https://internal-api-okr.feishu-boe.cn/stream/api/downloadFile/?file_token=boxbcMTBQO9ofLjWkDuPxkxOA2c&ticket=eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ0YXJnZXRfaWQiOiI3MDQxNDMwMzc3NjQyMDgyMzIzIiwidGFyZ2V0X3R5cGUiOjMsImFjdGlvbiI6MiwiZmlsZV90b2tlbiI6ImJveGJjTVRCUU85b2ZMaldrRHVQeGt4T0EyYyIsInVzZXJfaWQiOiI2OTY5ODU1NTAxNzQ0ODM0MDkyIiwidGVuYW50X2lkIjoiNjg3NzUwMjY4NzYwOTQwNjk5MCIsImV4cCI6MTY0MDE1NTk2M30.yc4qV2pkGUVwSO53-N_XGgeMucjmDn9iso1Ez_8vpghFz8YdeSDf4NHQpxOHYHc8RURvwI0a5UTNKKJ9CWagTQ", "width": 458 } ] }, "type": "gallery" }, { "paragraph": { "elements": [ { "textRun": { "style": {}, "text": "abc" }, "type": "textRun" } ], "style": { "list": { "indentLevel": 1, "type": "checkBox" } } }, "type": "paragraph" }, { "paragraph": { "elements": [ { "textRun": { "style": {}, "text": "abc" }, "type": "textRun" } ], "style": { "list": { "indentLevel": 1, "type": "checkedBox" } } }, "type": "paragraph" } ] }, "modify_time": "1640675837810", "progress_id": "7046317985098760212" }, "msg": "success" }

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 500 | 1009999 | internal server error | Internal error |
| 500 | 1009998 | system exception | System anomaly |
| 400 | 1001001 | invalid parameters | Invalid parameter |
| 400 | 1001002 | no permission | No permission |
| 400 | 1001003 | user not found | User does not exist |
| 400 | 1001004 | okr data not found | Okr data does not exist |
