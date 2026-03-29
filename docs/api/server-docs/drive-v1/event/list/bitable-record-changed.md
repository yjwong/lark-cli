---
title: "Bitable Record Changed"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/event/list/bitable-record-changed"
service: "drive-v1"
resource: "event"
section: "Docs"
scopes:
  - "bitable:app"
  - "drive:drive"
updated: "1665221868000"
---

# Bitable record changed
> To understand the scenarios and configuration process of event subscription, see Event subscription overview.

This event is triggered when a bitable's record changed, **Changes in formula field values do not trigger events**

## Summary
| Supported app types | custom,isv |
| --- | --- |
| Required scopes Enable any scope from the list | `bitable:app` `drive:drive` | How to subscribe to the document, please click to view Subscribe Docs events
。

## Supported record change types

| Change type         | action           |
| --------- | --------------- |
|Add record | `record_added` | 
|Delete record | `record_deleted` | 
|Edit record | `record_edited` | The `field_value` field in the event structure is a JSON serialized string. Please see the original structure at Bitable structure.

## Event example
``` json
{
  "schema": "2.0",  // Version of event format. If there is no such field, the version is 1.0.
  "header": {
    "event_id": "85ec04b2f6cff27a1bf249b26fbabcef", // Unique identifier of the event
    "token": "L3bzVf0sz7qb9XoTavpsHe0uJv7abcef", // The verification token
    "create_time": "1652677457000", //  Event sending time
    "event_type": "drive.file.bitable_record_changed_v1", // Event type
    "tenant_key": "736588c9260abcef", // Company identifier 
    "app_id": "cli_a00c8400a7babcef" // App ID
  },
  "event": {
    "action_list": [ // Change action list
      {
        "action": "record_edited", // Action type：edit record
        "after_value": [
          {
            "field_id": "fld9Eabcef", // Changed field ID
            "field_value": "666" // Field value after edit
          },
          {
            "field_id": "fldqaabcef",// Changed field ID
            "field_value": "[{\"type\":\"text\",\"text\":\"after value\"}]" // Field value after edit
          }
        ],
        "before_value": [
          {
            "field_id": "fld9Eabcef", // Changed field ID
            "field_value": "123" // Field value before edit
          },
          {
            "field_id": "fldqaabcef", // Changed field ID
            "field_value": "[{\"type\":\"text\",\"text\":\"before value\"}]" // Field value before edit
          }
        ],
        "record_id": "rec9sabcef" // Edited record id
      },
      {
        "action": "record_added", // Action type：add record
        "after_value": [
          {
            "field_id": "fld9Eabcef", // Changed field id
            "field_value": "666" // Field value after added
          },
          {
            "field_id": "fldqaabcef",// Changed field id
            "field_value": "[{\"type\":\"text\",\"text\":\"field value\"}]" // Field value after added
          }
        ],
        "record_id": "rec9sabcef" // Added record id
      },
      {
        "action": "record_deleted", // Action type：delete record
        "before_value": [
          {
            "field_id": "fld9Eabcef", // Changed field id
            "field_value": "666" // Field value before deleted
          },
          {
            "field_id": "fldqaabcef",// Changed field id
            "field_value": "[{\"type\":\"text\",\"text\":\"field value\"}]" // Field value before deleted
          }
        ],
        "record_id": "rec9sabcef" // Deleted field id
      }
    ],
    "file_token": "bascnItn6oHUSEL8RDUdF6abcef", // Bitable token
    "file_type": "bitable", // File type: bitable
    "operator_id": { // operator
      "open_id": "ou_9bc587355789fc049904ae7c736abcef",
      "union_id": "on_8f71e0224c012a0365ad4e3c733abcef",
      "user_id": "638abcef"
    },
    "subscriber_id_list": [ // List of subscribers
      {
        "open_id": "ou_9bc587355789fc049904ae7c736abcef",
        "union_id": "on_8f71e0224c012a0365ad4e3c733abcef",
        "user_id": "638abcef"
      }
    ],
    "table_id": "tblOaqBWfGeabcef" // Table ID of changed records
  }
}
```
