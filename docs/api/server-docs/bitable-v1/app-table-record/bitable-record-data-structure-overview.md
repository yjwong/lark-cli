---
title: "Base record data structure"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-record/bitable-record-data-structure-overview"
service: "bitable-v1"
resource: "app-table-record"
section: "Docs"
updated: "1753668076000"
---

# Base record data structure
Each row of data in a multi-dimensional table is a record. The data in a record is represented by the `fields` parameter. This document describes the data structure of the multi-dimensional table record data `fields`.

## Example of fields

In the above image, the highlighted data structure of the record is shown below:

```json
{
  "fields": {
    "Task Summary": [
      {
        "text": "The website update task is handled by Huang Paopao and is in progress.",
        "type": "text"
      }
    ],
    "Task Executor": [
      {
        "avatar_url": "https://s1-imfile.com/static-resource/v1/v3_00g2_058610dc-f65c-40c5-afac-46e83919630g~?image_size=72x72&cut_type=default-face&quality=&format=jpeg&sticker_format=.webp",
        "email": "amandahuang@bytedance.com",
        "en_name": "Amanda Huang",
        "id": "ou_8240099442cf5da49f04f4bf8f8abcef",
        "name": "Huang Paopao"
      }
    ],
    "Task Description": [
      {
        "text": "Update the company website",
        "type": "text"
      }
    ],
    "Start Date": 1675440000000,
    "Is Delayed": {
      "type": 1,
      "value": [
        {
          "text": "✅ Normal",
          "type": "text"
        }
      ]
    }
  }
}
```

## Structure of fields

The `fields` field is a map consisting of key-value pairs formed by the field names of the table and their specific content.
```json
{
  "Task Summary": [
    {
      "text": "The website update task is handled by Huang Paopao and is in progress.",
      "type": "text"
    }
  ]
}
```

| Parameter | Data Type | Description | Example Value |
| --- | --- | --- | --- |
| key | string | The field name in the multi-dimensional table. | "Task Summary" |
| value | union | The specific content of a field, which can be a number, string, boolean, list of strings, or list of objects. For more details, refer to the sections below. | This example value is a list of objects; for more examples, see below. ```json [ { "text": "The website update task is in progress.", "type": "text" } ] ``` | ### Structure of key

The key in the `fields` field is always of string type, corresponding to the title of each column in the multi-dimensional table, such as "Task Summary".

### Structure of value

This section introduces the data types and examples of the value structure corresponding to different field types in `fields`.

| Field Type type Enumeration | Field Type | value Data Type | value Description | value Example Value | Limitations |
| --- | --- | --- | --- | --- | --- |
| 1 | Text, Email, or Barcode. Use ui_type to distinguish; for details, refer to the description in the Text, Email, or Barcode section below. | When writing, it is a string; when returning, it is a list of objects. | Please refer to the description in the Text, Email, or Barcode section below. | Please refer to the examples in the Text, Email, or Barcode section below. | None |
| 2 | Number, Progress, Currency, or Rating. Use ui_type to distinguish: - When ui_type is "Number", the field type is a number - When ui_type is "Progress", the field type is progress - When ui_type is "Currency", the field type is currency - When ui_type is "Rating", the field type is rating | number | Numeric type | 0.5 | None |
| 3 | Single Choice | string | Text of the option name | "In Progress" | The total number of options in a single-choice field cannot exceed 5,000 |
| 4 | Multiple Choice | array<string> | An array containing multiple option name strings | ```json [ "Approval Integration", "Office Management", "Identity Management" ] ``` | - The total number of options in a multiple-choice field cannot exceed 5,000 - The number of options in a single cell cannot exceed 1,000 |
| 5 | Date | number | Unix timestamp, in milliseconds | 1675526400000 | None |
| 7 | Checkbox | boolean | Optional values: - true: Checked style - false: Unchecked style | true | None |
| 11 | Person | list of object | For fields of type person, the elements in the value are defined as follows: - id: User ID of the person, supports open_id, union_id, and user_id - name: Person's name (Not supported to pass in through the writing interface) - avatar_url: Link to the person's avatar (Not supported to pass in through the writing interface) - en_name: Person's English name (Not supported to pass in through the writing interface) - email: Person's email (Not supported to pass in through the writing interface) | ```json [ { "avatar_url": "https://s1-imfile.com/static-resource/v1/v3_00g2_058610dc-f65c-40c5-afac-46e83919630g~?image_size=72x72&cut_type=default-face&quality=&format=jpeg&sticker_format=.webp", "email": "amandahuang@bytedance.com", "en_name": "Amanda Huang", "id": "ou_8240099442cf5da49f04f4bf8f8abcef", "name": "Huang Paopao" } ] ``` | The number of people in a single cell cannot exceed 1,000 |
| 13 | Phone Number | string | Phone number, string matching the regular expression (\+)?\d* | "17899870000" | The length of the phone number cannot exceed 64 |
| 15 | Hyperlink | object | For fields of type hyperlink, the elements in the value are defined as follows: - text: Text displayed in the hyperlink - link: URL | ```json { "link": "/ssl:ttdoc/home/index", "text": "Lark Open Platform" } ``` | None |
| 17 | Attachment | list of object | Please refer to the description in the Attachment section below. | Please refer to the examples in the Attachment section below. | The number of attachments in a single cell cannot exceed 100 |
| 18 | One-way Association | object | For fields of type one-way association, the elements in the value are defined as follows: - link_record_ids: Associated record IDs, of array type, can contain multiple record ID strings. | ```json { "link_record_ids": [ "reclzUoBLn", "rec7bYQoX1", "recFIE3n52" ] } ``` | The number of One-way associations in a single cell cannot exceed 500 |
| 19 | Lookup Reference | object | Lookup references are essentially formulas. Its value definition is the same as that of a formula. Please refer to the description in the Formula or Lookup Reference section below. | Please refer to the description in the Formula or Lookup Reference section below. | None |
| 20 | Formula | object | Please refer to the description in the Formula or Lookup Reference section below. | Please refer to the examples in the Formula or Lookup Reference section below. | None |
| 21 | Bidirectional Association | object | For fields of type bidirectional association, the elements in the value are defined as follows: - link_record_ids: Associated record IDs, of array type, can contain multiple record ID strings. | ```json { "link_record_ids": [ "reclzUoBLn", "rec7bYQoX1", "recFIE3n52" ] } ``` | The number of bidirectional associations in a single cell cannot exceed 500 |
| 22 | Geolocation | object | For fields of type geolocation, the elements in the value are defined as follows: - location: Latitude and longitude - pname: Province - cityname: City - adname: District - address: Detailed address - name: Place name - full_address: Complete address | ```json { "address": "10 Xueqing Road, Xueqing Jiachang Building", "adname": "Haidian District", "cityname": "Beijing", "full_address": "ByteDance, Beijing, Haidian District, 10 Xueqing Road, Xueqing Jiachang Building", "location": "116.352681,40.01437", "name": "ByteDance", "pname": "Beijing" } ``` | None |
| 23 | Group | list of object | For fields of type group, the elements in the value are defined as follows: - name: Group name - avatar_url: Link to the group avatar - id: Group ID | ```json [ { "avatar_url": "https://s1-imfile.com/static-resource/avatar/default-avatar_9fb72564-d52a-49b0-9de8-f79071a02286_96.webp", "id": "oc_d2a947abb78bbbbb12d4cad55fbabcef", "name": "Test Department" } ] ``` | The number of groups in a single cell cannot exceed 10 |
| 1001 | Creation Time | number | Unix timestamp, in milliseconds. | 1675526400000 | - |
| 1002 | Last Update Time | number | Unix timestamp, in milliseconds. | 1675526400000 | - |
| 1003 | Creator | object | For fields of type person, the elements in the value are defined as follows: - id: User ID of the person, supports open_id, union_id, and user_id - name: Person's name (Not supported to pass in through the writing interface) - avatar_url: Link to the person's avatar (Not supported to pass in through the writing interface) - en_name: Person's English name (Not supported to pass in through the writing interface) - email: Person's email (Not supported to pass in through the writing interface) | ```json [ { "avatar_url": "https://s1-imfile.com/static-resource/v1/v3_00g2_058610dc-f65c-40c5-afac-46e83919630g~?image_size=72x72&cut_type=default-face&quality=&format=jpeg&sticker_format=.webp", "email": "amandahuang@bytedance.com", "en_name": "Amanda Huang", "id": "ou_8240099442cf5da49f04f4bf8f8abcef", "name": "Huang Paopao" } ] ``` | None |
| 1004 | Modifier | object | For fields of type person, the elements in the value are defined as follows: - id: User ID of the person, supports open_id, union_id, and user_id - name: Person's name (Not supported to pass in through the writing interface) - avatar_url: Link to the person's avatar (Not supported to pass in through the writing interface) - en_name: Person's English name (Not supported to pass in through the writing interface) - email: Person's email (Not supported to pass in through the writing interface) | ```json [ { "avatar_url": "https://s1-imfile.com/static-resource/v1/v3_00g2_058610dc-f65c-40c5-afac-46e83919630g~?image_size=72x72&cut_type=default-face&quality=&format=jpeg&sticker_format=.webp", "email": "amandahuang@bytedance.com", "en_name": "Amanda Huang", "id": "ou_8240099442cf5da49f04f4bf8f8abcef", "name": "Huang Paopao" } ] ``` | None |
| 1005 | Auto Number | string | A string composed of automatic numbering rules. | `"1"` | None | #### Text, Email, or Barcode

When the field type `type` enumerates to 1, the field type is distinguished based on `ui_type`, which can be text, email, or barcode. For details, refer to the following sections.

- When `ui_type` is "Text", the field is of text type, and the elements in its `value` are defined as follows:

  :::html
  | Parameter | Data Type | Description |
| --- | --- | --- |
| type | string | The text display type, optional values are: - text: plain text type - mention: mention user or cloud document type - url: hyperlink type |
| text | string | The text content |
| token | string | This field is valid when the type field is mention. - When mentionType is User, token is the user ID - When mentionType is Docx, token is the document's document_id - When mentionType is Sheet, token is the spreadsheet's spreadsheet_token - When mentionType is Bitable, token is the multidimensional table's app_token |
| link | string | The link. This field is valid when the type field is url |
| mentionType | string | This field is valid when the type field is mention. The optional values are: - User: mention user - Docx: mention document - Sheet: mention spreadsheet - Bitable: mention multidimensional table |
| mentionNotify | boolean | This field is valid when the type field is mention and the mentionType field is User. The optional values are: - false: do not mention this user - true: mention this user |
| name | string | The name of the mentioned user. This field is valid when the type field is mention and the mentionType field is User | :::

- When `ui_type` is "Barcode", the field type is barcode type, and an example of its value is as follows:
    - type: fixed value "text"
    - text: barcode number
      ```json
      [
        {
          "text": "FS0001",
          "type": "text"
        }
      ]
      ```
- When `ui_type` is "Email", the field type is email type, and an example of its value is as follows:
    - link: the URL link to the email
    - type: fixed value "url"
    - text: user email
      ```json
      {
        "link": "mailto:zhangmin@bytedance.com",
        "text": "zhangmin@bytedance.com",
        "type": "url"
      }
      ```

#### Attachments

For fields of attachment type, the elements in the value are defined as follows:
| Parameter | Data Type | Description |
| --- | --- | --- |
| file_token | string | The token of the attachment. You can use the download material interface to download this attachment |
| name | string | The name of the attachment |
| type | string | The mime type of the attachment, e.g., image/png |
| size | int | The size of the attachment. Unit: bytes |
| url | string | The attachment URL link, requires access token authentication. You can use the download material interface to download this attachment |
| tmp_url | string | The URL link for generating a temporary download link for the attachment, requires access token authentication. You can use the interface to obtain the temporary download link for materials | An example of the value for attachment type is as follows:
```json
[
  {
    "file_token": "J7GdbgNWWoD1fwx7oWccxdgknIe",
    "name": "58cc930b89.png",
    "size": 108867,
    "tmp_url": "https://open.larksuite.com/open-apis/drive/v1/medias/batch_get_tmp_download_url?file_tokens=J7GdbgNWWoD1fwx7oWccxdgknIe&extra=%7B%22bitablePerm%22%3A%7B%22tableId%22%3A%22tblx0Ed2NnBULN6a%22%2C%22rev%22%3A5%7D%7D",
    "type": "image/png",
    "url": "https://open.larksuite.com/open-apis/drive/v1/medias/J7GdbgNWWoD1fwx7oWccxdgknIe/download?extra=%7B%22bitablePerm%22%3A%7B%22tableId%22%3A%22tblx0Ed2NnBULN6a%22%2C%22rev%22%3A5%7D%7D"
  }
]
```

#### Formula or Lookup Reference

For fields of formula or lookup reference type, the elements in the value are defined as follows:
| Parameter | Data Structure | Description |
| --- | --- | --- |
| type | number | Used to specify the data type of the value, optional values are as follows (the same field type is distinguished by ui_type): - 1: Text, Barcode - 2: Number, Progress, Currency, Rating - 3: Single Select - 4: Multi Select - 5: Date - 7: Checkbox - 11: User - 13: Phone Number - 15: Hyperlink - 17: Attachment - 18: Single Link - 19: Lookup Reference - 20: Formula - 21: Duplex Link - 22: Geographic Location - 23: Group - 1001: Creation Time - 1002: Last Modified Time - 1003: Creator - 1004: Modifier - 1005: Auto Number |
| ui_type | string | The UI type of the field, optional values are as follows: - Text: Text - Barcode: Barcode - Number: Number - Progress: Progress - Currency: Currency - Rating: Rating - SingleSelect: Single Select - MultiSelect: Multi Select - DateTime: Date - Checkbox: Checkbox - User: User - GroupChat: Group - Phone: Phone Number - Url: Hyperlink - Attachment: Attachment - SingleLink: Single Link - Formula: Formula - DuplexLink: Duplex Link - Location: Geographic Location - CreatedTime: Creation Time - ModifiedTime: Last Modified Time - CreatedUser: Creator - ModifiedUser: Modifier - AutoNumber: Auto Number |
| value | list | The type field determines the data structure of the value; refer to the value structure in this document. Note: When the corresponding base type data structure is not in list format, this field will be in the corresponding data's list format. | An example of the value for formula or lookup reference type is as follows:
```json
{
  "type": 1,
  "value": [
    {
      "text": "✅ Normal",
      "type": "text"
    }
  ]
}
```
