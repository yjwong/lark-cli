---
title: "Get OKR progress records"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/okr-v1/progress_record/get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/okr/v1/progress_records/:progress_id"
service: "okr-v1"
resource: "progress_record"
section: "OKR"
scopes:
  - "okr:okr:readonly"
  - "okr:okr"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1669270474000"
---

# Get OKR progress records

Obtain OKR progress record details according to ID

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/okr/v1/progress_records/:progress_id |
| HTTP Method | GET |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `okr:okr:readonly` `okr:okr` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `progress_id` | `string` | OKR progress record ID to be queried **Example value**: "7041857032248410131" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: "open_id" **Optional values are**:  - `open_id`: user's open id - `union_id`: user's union id - `user_id`: user's user id  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ## Response

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

{{
"Code": 0,
"Data": {
"Content": {
"Blocks": [
{{
"Paragraph": {
"Elements": [
{{
"TextRun": {
"Style": {},
"Text": "
},
"Type": "textRun"
},
{{
"Person": {
"OpenId": "ou_b1ba99a5340289d7af30839fd95ce1ee",
"UserId": "7012194140645721644"
},
"Type": "person"
}
]
},
"Type": "paragraph"
}
]
},
"modify_time": "1640677213095",
"progress_id": "7046518160811425812"
},
"Msg": "success"
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 500 | 1009999 | internal server error | Internal error |
| 500 | 1009998 | system exception | System anomaly |
| 400 | 1001001 | invalid parameters | Invalid parameter |
| 400 | 1001002 | no permission | No permission |
| 400 | 1001003 | user not found | User does not exist |
| 400 | 1001004 | okr data not found | Okr data does not exist |
