---
title: "Custom Field Overview"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/custom_field/custom-field-overview"
service: "task-v2"
resource: "custom_field"
section: "Tasks v2"
updated: "1699521745000"
---

# Custom Field Feature Overview
The task supports the expansion of custom fields in tasks, allowing users to add extra task information more clearly, manage tasks efficiently, and assist in collaborative progress. In addition to using system fields such as "task due" or "task assignees", users can define custom fields closely related to their usage scenario, such as "priority", "project release date", "price", etc.

Two related concepts about custom fields are need to be distinguished carefully: "**Custom Field Settings**" and "**Custom Field Values**."

Custom field settings are the metadata of custom fields, including name, type, and specific settings. Each custom field must have a name.

The currently supported custom field types are as follows:

| **Field Type** | **Description**                           |
| --------  | -------------------------------- |
| number     | Supports expressing a numeric value and displaying it in different formats in the App  |
| member      | Supports adding internal and external contacts; multiple members can be selected at the same time           |
| datetime    | Supports expressing a specific date                   |
| single_select | Used to add options to tasks; supports custom option name and color; only one option can be selected  |
| multi_select  | Used to add options to tasks; supports custom option name and color; multiple options can be selected at the same time |
| text | Used to add customized text to tasks | The settings for each type are different. For example, for the number type, you can set the number format, whether to display custom symbols, and the position of symbols; for the datetime type, you can set the display format of the date; for single/multi-select, you can define their options.

Custom field setting related APIs include:

* Create Custom Field
* Update Custom Field
* Get Custom Field
* List Custom Fields
* Add Custom Field to Resource
* Remove Custom Field from Resource

For the options in the single/multi-select field types, support:
* Create Custom Field Option
* Update Custom Field Option

Custom field values are similar to the values of built-in system fields. After defining a custom field in a tasklist, you can set the field value for the tasks in the tasklist. For example, you can use the Update Task API to set the custom field value, or you can set the custom field value when creating a task by invoking the Create Task API. Besides, users can view the custom field values of the task through the Get Task Details API.

## Data Model and Authorization

Custom fields and tasklists have a many-to-many relationship. A tasklist can define multiple custom fields. A custom field can be added to multiple tasklists, too.

Custom field permissions are defined by the tasklists they are added to. If a user has edit/read permissions for a custom field, it means that there is at least one user-tasklist-custom field relationship, and the user has edit/read permissions for that tasklist (for how tasklists are authorized, see Tasklist Feature Overview). If there are multiple "tasklists" between the user and the custom field, the user's permission for the custom field is taken from the one with the highest tasklist role. Therefore, when Create Custom Field, it must always be associated with a checklist in order to add custom fields for these tasklists; or you can use the Add Custom Field to Resource API to add a custom field to other tasklists. After adding, the field can be viewed in the "Field Settings" of the tasklist UI; meanwhile, all other collaborators of the tasklist will also get the corresponding permissions for the field. If you do not want a field to be displayed on a particular tasklist, you can use the Remove Custom Field from Resource API to remove it. After removal, the original tasklist collaborators will lose permission for the field unless they can access the custom field through other tasklists.

For example, there is a user U1 who is an editor of tasklist L1 and a viewer of tasklist L2. L1 has custom fields F1 and F2; L2 has custom fields F2 and F3. The user has edit permissions for F1 and F2; and read permissions for F3. If there is another tasklist L3 that U1 cannot access, and L3 has custom fields F3 and F4, the user has no permissions for F4. However, if there happens to be another user U2 who is an editor of both L1 and L3, U2 will gain edit permissions for F4. U2 can add F4 to L1, and then U1 will also have edit permissions for F4. As shown in the figure below:

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/429b6a3f72ec69a2db4812791f1cb86f_nJq9YEbngy.png?height=1446&lazyload=true&maxWidth=500&width=1798)

You can use the List Custom Fields API to query all custom fields that the caller has permission to access, or view custom fields in a specific tasklist by specifying the tasklist GUID.

If a custom field is removed from all checklists, it means that no users can access it anymore. From the user's perspective, a custom field that has been removed from all tasklists is equivalent to being deleted. However, direct deletion of custom fields is currently not supported.

Custom field values belong to both tasks and custom fields. To view a custom field value of a task, a user must have read permissions for the task (see Task Features Overview) and read permissions for the custom field; similarly, to modify a custom field value of a task, a user must have edit permissions for both the task and the custom field.

For example, a task T has three custom field values F1, F2, and F3. User U1 can edit T and has read permissions for F1 and F2. They can view the values of F1 and F2 through the Get task details API, but the value of F3 is not visible to them.

Custom field values cannot be deleted, but they can be "cleared", i.e., set to "none", which will be introduced in the following sections.

## Number

Number type custom fields can be used to display a number, such as

```json
{
  "guid": "d1fd19fd-4810-479d-a093-fb9b0442c6f6",
  "name": "price",
  "type": "number",
  "number_setting": {
    "format": "custom",
    "decimal_count": 2,
    "separator": "thousand",
    "custom_symbol": "€",
    "custom_symbol_position": "right"
  }
}
```
It means displaying the number with a custom symbol, with 2 decimal places. The integer part shows thousand separators, and the € symbol will be displayed on the right side of the number. Settings other than `decimal_count` only affect the display effect in the App UI and have no effect on the input and output of field values in openapi.

`decimal_count` affects the number of decimal places, supporting up to 6 decimal places. User-entered fields exceeding this setting will be automatically rounded.

In addition, when `format` is not "custom", `custom_symbol` can be left unset; but if `format` is "custom", `custom_symbol` must be set to a string with a maximum of 4 characters.

The vale of number supports regular numbers (considering the accuracy issue of json, it will be input and output in the form of a string), such as:
```
PATCH /task/v2/tasks/:task_guid
{
   "task": {
     "custom_fields": [
        {
           "guid": "b78a2bca-b580-4853-b3f8-2ebd477d91ca",
           "number_value": "12.34"
        }
      ]
   },
   "update_fields": ["custom_fields"]
}
```
It can set a task's number value of custom field (custom_field_guid=b78a2bca-b580-4853-b3f8-2ebd477d91ca) to 12.34.

Some valid number examples are:

* "1.23"
* "0.9248"
* "0"
* "-12.45"
* "+6"

The following inputs will be considered invalid:

* "1.2.3" (not a number)
* "1e9" (scientific notation not supported)
* " 99 " (contains spaces)
* "0x125" (hexadecimal not supported, symbols with non-numbers are considered invalid)
* "20%" (percentage input not supported, input actual value 0.2 instead)

Number value use empty string "" to represent empty value.

Numbers will be automatically stored in the simplest format after inputting value, for example:

* Input "1.200" will only retain "1.2"
* Input "054" will only retain "54", numbers starting with 0 will NOT be treated as octal, but as redundant 0s in the decimal place
* Input "+67.1" will only retain "67.1"

When entering a number that exceeds the value specified by decimal_count, it will be rounded:
* Input "20.124", decimal_count=2, will result in "20.12"
* Input "0.1558", decimal_count=2, will result in "0.16"

When entering a percentage with `format=percentage`, still input the original value of the number. For example:
* Input "0.2", decimal_count=0, the App UI will display "20%"
* Input "2.45", decimal_count=2, the App UI will display "245.00%"

Please note that when the input is a percentage value, decimal_count refers to the number of decimal places after displaying as a percentage. Therefore, the actual number of decimal places is always 2 more than for non-percentages. For example, input "0.24563", decimal_count=2;
* When `format` is not a percentage, it will result in "0.25"
* When `format` is a percentage, it will result in "0.2556" (displayed as 25.56% on the App UI)

Changing the `decimal_count` of a numeric custom field will not change the already set field values, but will affect the display on the App UI. For example, set `decimal_count` to 2, then write a value "1.23" for this field, and then change `decimal_count` to 1. The field value obtained through openapi is still "1.23", but the App UI will display "1.2", and filtering/sorting functions will be performed based on "1.2".

## Member

Member custom fields can be filled with one or more users. For example:

```json
{
  "guid": "b78a2bca-b580-4853-b3f8-2ebd477d91ca",
  "name": "reviewers",
  "type": "member",
  "member_setting": {
    "multi": true,
  }
}
```
This indicates a custom field for member types that supports multi-selection, allowing multiple users' information to be filled in.

The value of member types is represented by "member" data structure. It is similar to the member format in Adding Task Members, but does not support filling in `role`. That is only `id` and `type` can be passed in. The `type` only supports "user" and is the default value, so it can be omitted. The `id` must be filled in, and its format is defined by the `user_id_type` parameter (see "How to use user_id_type to modify the user ID format?" in Task Features Overview), and the default is "open_id".

```
PATCH /task/v2/tasks/:task_guid
{
   "task": {
     "custom_fields": [
        {
           "guid": "b78a2bca-b580-4853-b3f8-2ebd477d91ca",
           "member_value": [
              { "id": "ou_49dadcd6fd55da971d887087c4f3a37a" },
              { "id": "ou_2cefb2f014f8d0c6c2d2eb7bafb0e54f" }
           ]
        }
      ]
   },
   "update_fields": ["custom_fields"]
}
```

When `multi` is set to false, `member_value` can only have one element. If a field value is already set to multiple members, changing the field's `multi` from true to false will not affect the field value until its next modification.

When `multi` is set to true, the multiple elements in `member_value` cannot be duplicated.

An empty array is used to represent "empty" for member type field values.

## Datetime

Datetime custom field can be filled with a timestamp representing a date. The setting is:

```json
{
  "guid": "a0fddac1-0fd5-4903-8cae-5b93ccf0f797",
  "name": "release day",
  "type": "datetime",
  "datetime_setting": {
    "format": "yyyy/mm/dd",
  }
}
```

The settings for date custom fields only affect the display format in the App.

The value of date custom fields is represented in a timestamp format in milliseconds (since 1970-01-01 00:00:00 UTC), but the precision is only retained to the day.

```
PATCH /task/v2/tasks/:task_guid
{
   "task": {
     "custom_fields": [
        {
           "guid": "b78a2bca-b580-4853-b3f8-2ebd477d91ca",
           "datetime_value": "1666108800000"
        }
   },
   "update_fields": ["custom_fields"]
}
```

To set the value of a date field, you need to use the timestamp representation of the 00:00:00 UTC time of that date. For example, you can use the JavaScript moment library to generate it like this:

```
let m = require('moment');
m('2022-10-19T00:00:00Z').valueOf();
// Get 1666137600000, the timestamp of 2022-10-19
```

When setting the datetime field value, if the precision is greater than the day, it will be automatically truncated. For example, if you input "1666137600322", only "1666137600000" will be retained.

Datetime types fields only support positive timestamps (i.e., dates after 1970-01-01). An error will be prompted if a negative timestamp is given.

An empty string "" is used to represent empty for datetime field values.

## Single/Multi Select

The settings for single/multi select are similar, both providing a set of options. The difference is that the value of a single_select can only be set to at most one option, while value of multi_select can be set to multiple options.

Example of single_select:
```json
{
  "guid": "a0fddac1-0fd5-4903-8cae-5b93ccf0f797",
  "name": "priority",
  "type": "single_select",
  "single_select_setting": {
    "options": [
      {"guid": "734a0135-fcbb-4482-8909-f9589b03d9f1", "name": "high", "color_index": 23, "is_hidden": false},
      {"guid": "89a7d5b0-e29f-4363-a614-1fe47f2d3255", "name": "mid", "color_index": 24, "is_hidden": false},
      {"guid": "f911b99f-c6df-477a-a014-8b76ac40aa00", "name": "low", "color_index": 23, "is_hidden": false},
     ]
  }
}
```

Example of multi_select:
```json
{
  "guid": "7748b214-d299-4fc3-9d12-d95ef39979f7",
  "name": "publish regions",
  "type": "multi_select",
  "multi_select_setting": {
    "options": [
      {"guid": "345638bd-35b2-48a3-b3cd-ca17169a674f", "name": "america", "color_index": 10, "is_hidden": false},
      {"guid": "439615eb-2d11-41f6-b1c7-4418fd96dd62", "name": "asia", "color_index": 15, "is_hidden": false},
      {"guid": "f911b99f-c6df-477a-a014-8b76ac40aa00", "name": "europe", "color_index": 14, "is_hidden": false},
     ]
  }
}
```

The value of a single_select field is the one GUID of its options; the value of a multi_select field is an array of GUIDs of its options. The option GUIDs in the values of the same multi_select field cannot contain duplicated elements.
```
PATCH /task/v2/tasks/:task_guid
{
   "task": {
     "custom_fields": [
        {
           "guid": "a0fddac1-0fd5-4903-8cae-5b93ccf0f797",
           "single_select_value": "89a7d5b0-e29f-4363-a614-1fe47f2d3255"
        },
        {
           "guid": "7748b214-d299-4fc3-9d12-d95ef39979f7",
           "multi_select_value": [
             "345638bd-35b2-48a3-b3cd-ca17169a674f",
             "0f8165c6-b6f9-45e6-ad47-618189546007"
           ]
        }
   },
   "update_fields": ["custom_fields"]
}
```

A single_select field value uses an empty string "" to represent empty value; mutli_select field value use an empty array to represent empty value.

### About Options

Each option includes a name, a color index value, and a hidden flag (is_hidden). Options are not allowed to be deleted because they need to be associated with field values. If option could be deleted, the field value would not be able to find the corresponding option. However, options can be marked as hidden (is_hidden=true). Hidden options will not appear in the UI, and openapi is not allowed to set the value of a single/multi select field to a hidden option.

A single/multi select field can have up to 100 options (including both hidden and visible ones).

Options that are not hidden are called "visible options", and the settings for single/multi select fields must meet the following constraints: **All visible options for the same field must not have the same name**. Therefore, attempting to create a single/multile field with the same visible option name using openapi or restoring a hidden field to visible, which happens to have the same name as another visible option, will result in an error.

The color index (`color_index`) of the options is an enum value from 0 to 54. Each value represents a color. If not specified during creation, a random color will be selected. The following image shows the correspondence between the `color_index` value and the color.

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/cced3489eb55a2b84b11aa055d4790a9_PCApI3mDxw.png?height=596&lazyload=true&maxWidth=600&width=2108)

Options have orders. The order of visible options returned in openapi is the same as their display order in the App UI. The order of options can be adjusted by using `insert_before` or `insert_after` in Create Custom Field Option and Update Custom Field Option, or by directly setting the visible option order through the Update Custom Field API.

## Text

Text custom field allows users to add customized text to tasks. Task custom field has no setting besides the field name.

```json
{
  "guid": "a0fddac1-0fd5-4903-8cae-5b93ccf0f797",
  "name": "review comment",
  "type": "text",
  "text_setting": {}
}
```

Users can add text into the `text_value`. The format is exactly same as built-in `summary` field of task.

```
PATCH /task/v2/tasks/:task_guid
{
   "task": {
     "custom_fields": [
        {
           "guid": "b78a2bca-b580-4853-b3f8-2ebd477d91ca",
           "text_value": "this is a customized "
        }
      ]
   },
   "update_fields": ["custom_fields"]
}
```

A text field value uses an empty string "" to represent empty value.
