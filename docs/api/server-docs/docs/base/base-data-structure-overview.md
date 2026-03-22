---
title: "Base data structure overview"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/bitable/development-guide/bitable-structure"
section: "Docs"
updated: "1753668070000"
---

# Data structure overview

This document introduces the data structure of records, fields, and views in Base tables. The data tables in Base consist of records and fields, and can have multiple views.

## Records
Records consist of two structures: record and fields.

### Record structure

A record is an object structure type.
| Parameter     | Data type         | Description        |
| --------- | -------   | --------- |
|`record_id`| string |  Record ID |
|`fields`| map |  Record fields | ### Fields structure

The fields attribute is of map type, consisting of key-value pairs of field names and their specific contents. For detailed structure and parameter descriptions of fields, refer to Base record data structure.
```json
{
  "Task summary": [
    {
      "text": "The website update task is being handled by Huang Paopao and is in progress",
      "type": "text"
    }
  ]
}
```

| Parameter | Data type | Description | Example value |
| --- | --- | --- | --- |
| key | string | Field name in the Base table. | "Task summary" |
| value | union | The specific content of a field, which can be a number, string, boolean, list of strings, or list of objects. For more details, refer to the following section. | This example value is a list of objects. For more examples, refer to Base record data structure. ```json [ { "text": "The website update task is in progress", "type": "text" } ] ``` | ## Fields

Fields are the "columns" in the Base table and are of object structure type. The basic structure of a field is shown below. For detailed structure and parameter descriptions of fields, refer to the Field editing guide.

```json
{
    "field_id": "fldYWaldeW", // Field ID
    "field_name": "Text",   // Field name
    "type": 1,  // Field type, see below for details
    "description": "Field description", // More information about the field
    "is_primary": true,   // Whether this field is the primary index field
    "property": null,   // Field property, see below for details
    "ui_type": "Text",  // The display type of the field in the interface, such as progress field being a display form of a number
    "is_hidden": false  // Whether the field is hidden
}
```
Parameter descriptions are as follows:
| Name | Type | Description |
| --- | --- | --- |
| field_id | string | Field ID |
| field_name | string | Field name |
| type | int | Field type: differentiated by ui_type for the same type: - 1: Text (default), Barcode (declare "ui_type": "Barcode"), Email (declare "ui_type": "Email") - 2: Number (default), Progress (declare "ui_type": "Progress"), Currency (declare "ui_type": "Currency"), Rating (declare "ui_type": "Rating") - 3: Single select - 4: Multi select - 5: Date - 7: Checkbox - 11: User - 13: Phone number - 15: URL - 17: Attachment - 18: Single link - 19: Lookup - 20: Formula - 21: Duplex link - 22: Location - 23: Group - 1001: Created time - 1002: Last modified time - 1003: Created by - 1004: Modified by - 1005: Auto number |
| description | Field description | More information about the field. |
| is_primary | true/false | Whether this field is the primary index field. |
| property | object | Field property, varies by field type. For details, refer to the Field editing guide. |
| ui_type | string | Field UI type: - "Text": Text - "Email": Email - "Barcode": Barcode - "Number": Number - "Progress": Progress - "Currency": Currency - "Rating": Rating - "SingleSelect": Single select - "MultiSelect": Multi select - "DateTime": Date - "Checkbox": Checkbox - "User": User - "GroupChat": Group - "Phone": Phone number - "Url": URL - "Attachment": Attachment - "SingleLink": Single link - "Formula": Formula - "Lookup": Lookup - "DuplexLink": Duplex link - "Location": Location - "CreatedTime": Created time - "ModifiedTime": Last modified time - "CreatedUser": Created by - "ModifiedUser": Modified by - "AutoNumber": Auto number |
| is_hidden | true/false | Whether the field is hidden. | ## Views

A view is an object structure type.
| Parameter | Type | Description |
| --- | --- | --- |
| view_id | string | View ID. `view_id` is unique within a Base table but not necessarily globally unique. Retrieval methods: - You can obtain the `view_id` through the Base table URL. The highlighted part in the image below is the unique identifier of the current view. ![](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/140668632c97e0095832219001d17c54_c76RMwZAgW.png?height=748&lazyload=true&maxWidth=700&width=2998) - You can also obtain the `view_id` through the list views API. Currently, it is not possible to obtain the `view_id` of a Base table embedded in a document. |
| view_name | string | View name |
| view_type | string | View type, supporting the following types, with grid as the default type. - grid: Grid view - kanban: Kanban view - gallery: Gallery view - gantt: Gantt view - form: Form view | ## Custom data structure

### delete_record
| Parameter         | Data type           |  Description         | 
| --------- | --------------- | ----------- | 
|`deleted` | `boolean` | Whether the deletion was successful |
|`record_id` | `string` | ID of a single record |
