---
title: "Resource introduction"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/room_level/room-level-overview"
service: "vc-v1"
resource: "room_level"
section: "Video Conferencing"
updated: "1689326161000"
---

# Resource introduction

Users can query meeting room levels, create meeting room levels, update meeting room levels, and other operations.

## Method

> The corresponding original regular muiti-layer meeting room API is as follows:
> * Obtain Building List
> * Query building details
> * Create Building
> * Update building
> * Delete Building
  If the tenant has switched to the flexible level of the meeting room, compatible replacement will be used when calling the fixed level interface, which may lead to a downgraded experience. It is recommended to use the flexible layer interface

**Query Room Level List**

    This API is used to query meeting room level list at a specified meeting room level.

**Query Room Level Details**

    This API is used to query a specified meeting room level details by room level ID.

**Batch Query Room Level Details**

    This API is used to query multi specified meeting room level details by room level IDs.

**Create Room Level**

    This API is used to create meeting room level.

**Update Room Level**

    This API is used to update a specified meeting room level details.

**Delete Room Level**

    This API is used to delete a specified meeting room level.

**Search Room Level**

    This API is used to search meeting room level. Support custom room level ID to search.

## Events

**Create Room Level**

      This event is triggered when a room level is created.

**Update Room Level**

      This event is triggered when a room level is updated.

**Delete Room Level**

      This event is triggered when a room level is deleted.

## Field description

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `room_level_id` | `string` | Yes | Room level ID **Example value**："omb_57c9cc7d9a81e27e54c8fabfd02759e7" |
| `name` | `string` | Yes | Room level name **Example value**："TestLevel" |
| `parent_id` | `string` | Yes | Parent room level ID **Example value**："omb_4ad1a2c7a2fbc5fc9570f38456931293" |
| `path` | `string[]` | Yes | Room level path **Example value**：["omb_4ad1a2c7a2fbc5fc9570f38456931293"] |
| `has_child` | `boolean` | Yes | Whether the room level has child levels **Example value**：true |
| `custom_group_id` | `string` | No | Custom room level ID **Example value**："10000" | ### Request body example

{
  "room_level_id": "omb_57c9cc7d9a81e27e54c8fabfd02759e7",
  "name": "TestLevel",
  "parent_id": "omb_4ad1a2c7a2fbc5fc9570f38456931293",
  "path": [
	 "omb_4ad1a2c7a2fbc5fc9570f38456931293"
  ],
  "has_child": false,
  "custom_group_id": "10000"
}
