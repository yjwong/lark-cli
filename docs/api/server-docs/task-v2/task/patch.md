---
title: "Patch Task"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/task/patch"
method: "PATCH"
api_path: "https://open.larksuite.com/open-apis/task/v2/tasks/:task_guid"
service: "task-v2"
resource: "task"
section: "Tasks v2"
rate_limit: "10 per second"
scopes:
  - "task:task:write"
updated: "1699521563000"
---

# Patch Task

This api is used to update the summary, description, due, etc. of a task.

To update a task, fill in all the field names to be updated in the `update_fields` field, and fill in the new value of the field to update in the `task` field. If the field name to be changed is included in the `update_fields`, but no new value is set in `task`, the field will be cleaned to empty. For update api specification, please refer to the "About Resource Update" section in Feature Overview .

The updatable fields include:

* `summary`: task summary
* `description`: task description
* `start`: task start time
* `due`: task due time
* `completed_at`: complete or uncomplete task
* `extra`: user data attached to the task
* `custom_complete`: task custom complete configration
* `repeat_rule`: repeat rule of the task.
* `mode`:  completion mode of the task.
* `is_milestone`: whether the task is a milestone task or not.

This api can be used to complete tasks by updating `completed_at` of task to a timestamp and restore tasks to uncompleted by setting `completed_at` to 0 . However, at present, regardless of whether the task itself is a countersign task or a or-sign task, the oapi can only support overall completion" of the task, and does not support individual completion. Besides, cannot complete a task that has bee completed. But a completed task can be restored to uncompleted by setting `completed_at` to "0".

If you want to update the repeat_rule of a task, you must not also update the `completed_at` of the task or unset the due. See "How to use 
recuring rule of task?" section in Tasklist Features Overview .

Task member/reminder/tasklist data cannot be updated by this api.
* To add/remove task members, use Add Task Member
And Remove Task Member.
* To modify the task reminder, use the Add Task Reminder and Remove Task Reminder.
* To change the tasklist where the task belongs to, use Add Task to Tasklist and Remove Task from Tasklist.

> To update a task, the calling identity needs to have task's edit permission. Please refer to the "How are Tasks Authorized?" section in Task Features Overview.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/task/v2/tasks/:task_guid |
| HTTP Method | PATCH |
| Rate Limit | 10 per second |
| Supported app types | custom,isv |
| Required scopes | `task:task:write` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `task_guid` | `string` | GUID of task to update **Example value**: "e297ddff-06ca-4166-b917-4ce57cd3a7a0" **Data validation rules**: - Maximum length: `100` characters | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Default value**: `open_id` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `task` | `input_task` | No | To update the task data, just set the field that appears in the `update_fields`. If update_fields sets a field name to be changed, but no new value is set in the task, the field will be cleared. |
|   `summary` | `string` | No | Task summary. When updating, empty is not allowed. Supports up to 3000 utf8 characters. **Example value**: "One review for full-year sales" |
|   `description` | `string` | No | Task description. Supports up to 3000 utf8 characters. **Example value**: "You need to read the review summary document in advance." |
|   `due` | `due` | No | Task due time. Please refer to the "How to use start and due time?" section in Feature Overview . **Example value**: 1675742789470 |
|     `timestamp` | `string` | No | The timestamp of the due time/date in milliseconds from 1970-01-01 00:00:00. If the expiration time is a date, you need to convert the date to timestamp and set is_all_day = true **Example value**: "1675454764000" |
|     `is_all_day` | `boolean` | No | Whether to due on a date. If set to true, only the date part of the timestamp will be parsed and stored. **Example value**: true |
|   `extra` | `string` | No | Any data that the caller can pass in attached to the task. It will be returned as is when the task details are obtained. **Example value**: "dGVzdA==" **Data validation rules**: - Maximum length: `65536` characters |
|   `completed_at` | `string` | No | The completion time of the task timestamp (ms) **Example value**: "1675742789470" **Default value**: `0` **Data validation rules**: - Maximum length: `20` characters |
|   `repeat_rule` | `string` | No | Task repeat rule. If set, the task is "recurring task". Please refer to the "How to Use Recurring Tasks? " section in Task Feature Overview . **Example value**: "FREQ=WEEKLY;INTERVAL=1;BYDAY=MO,TU,WE,TH,FR" **Data validation rules**: - Maximum length: `1000` characters |
|   `custom_complete` | `custom_complete` | No | Task custom complete configuration. Please refer to the "How to use custom completion?" section in Task Feature Overview . |
|     `pc` | `custom_complete_item` | No | pc client side custom configuration (including mac and windows) |
|       `href` | `string` | No | Customize the completed redirect url **Example value**: "https://www.example.com" |
|       `tip` | `i18n_text` | No | The pop-up prompt for customization completion is |
|         `en_us` | `string` | No | English **Example value**: "Workbench" **Data validation rules**: - Maximum length: `1000` characters |
|         `zh_cn` | `string` | No | Chinese **Example value**: "Workbench" **Data validation rules**: - Maximum length: `1000` characters |
|         `zh_hk` | `string` | No | Chinese (Hong Kong) **Example value**: "工作臺" **Data validation rules**: - Maximum length: `1000` characters |
|         `zh_tw` | `string` | No | Chinese (Taiwan) **Example value**: "工作臺" **Data validation rules**: - Maximum length: `1000` characters |
|         `ja_jp` | `string` | No | Japanese **Example value**: "作業台" **Data validation rules**: - Maximum length: `1000` characters |
|         `fr_fr` | `string` | No | French **Example value**: "Table de travail" |
|         `it_it` | `string` | No | Italian **Example value**: "banco di lavoro" |
|         `de_de` | `string` | No | German **Example value**: "Werkbank" |
|         `ru_ru` | `string` | No | Russian **Example value**: "верстак" |
|         `th_th` | `string` | No | Thai **Example value**: "โต๊ะทำงาน" |
|         `es_es` | `string` | No | Spanish **Example value**: "banco de trabajo" |
|         `ko_kr` | `string` | No | Korean **Example value**: "작업대" |
|     `ios` | `custom_complete_item` | No | Custom completion configuration of ios |
|       `href` | `string` | No | Customize the completed redirect url **Example value**: "https://www.example.com" |
|       `tip` | `i18n_text` | No | The pop-up prompt for customization completion is |
|         `en_us` | `string` | No | English **Example value**: "Workbench" **Data validation rules**: - Maximum length: `1000` characters |
|         `zh_cn` | `string` | No | Chinese **Example value**: "Workbench" **Data validation rules**: - Maximum length: `1000` characters |
|         `zh_hk` | `string` | No | Chinese (Hong Kong) **Example value**: "工作臺" **Data validation rules**: - Maximum length: `1000` characters |
|         `zh_tw` | `string` | No | Chinese (Taiwan) **Example value**: "工作臺" **Data validation rules**: - Maximum length: `1000` characters |
|         `ja_jp` | `string` | No | Japanese **Example value**: "作業台" **Data validation rules**: - Maximum length: `1000` characters |
|         `fr_fr` | `string` | No | French **Example value**: "Table de travail" |
|         `it_it` | `string` | No | Italian **Example value**: "banco di lavoro" |
|         `de_de` | `string` | No | German **Example value**: "Werkbank" |
|         `ru_ru` | `string` | No | Russian **Example value**: "верстак" |
|         `th_th` | `string` | No | Thai **Example value**: "โต๊ะทำงาน" |
|         `es_es` | `string` | No | Spanish **Example value**: "banco de trabajo" |
|         `ko_kr` | `string` | No | Korean **Example value**: "작업대" |
|     `android` | `custom_complete_item` | No | Custom completion configuration of android |
|       `href` | `string` | No | Customize the redirect redirect url **Example value**: "https://www.example.com" |
|       `tip` | `i18n_text` | No | The pop-up prompt for customization completion is |
|         `en_us` | `string` | No | English **Example value**: "Workbench" **Data validation rules**: - Maximum length: `1000` characters |
|         `zh_cn` | `string` | No | Chinese **Example value**: "Workbench" **Data validation rules**: - Maximum length: `1000` characters |
|         `zh_hk` | `string` | No | Chinese (Hong Kong) **Example value**: "工作臺" **Data validation rules**: - Maximum length: `1000` characters |
|         `zh_tw` | `string` | No | Chinese (Taiwan) **Example value**: "工作臺" **Data validation rules**: - Maximum length: `1000` characters |
|         `ja_jp` | `string` | No | Japanese **Example value**: "作業台" **Data validation rules**: - Maximum length: `1000` characters |
|         `fr_fr` | `string` | No | French **Example value**: "Table de travail" |
|         `it_it` | `string` | No | Italian **Example value**: "banco di lavoro" |
|         `de_de` | `string` | No | German **Example value**: "Werkbank" |
|         `ru_ru` | `string` | No | Russian **Example value**: "верстак" |
|         `th_th` | `string` | No | Thai **Example value**: "โต๊ะทำงาน" |
|         `es_es` | `string` | No | Spanish **Example value**: "banco de trabajo" |
|         `ko_kr` | `string` | No | Korean **Example value**: "작업대" |
|   `start` | `start` | No | Task start time. Please refer to the "How to use start and due time?" section in Feature Overview . |
|     `timestamp` | `string` | No | The timestamp of the start time/date in milliseconds from 1970-01-01 00:00:00. If the start time is a date, you need to convert the date to timestamp and set is_all_day=true **Example value**: "1675454764000" |
|     `is_all_day` | `boolean` | No | Whether to start on a date. If set to true, only the date part of the timestamp will be parsed and stored. **Example value**: true |
|   `mode` | `int` | No | The completion mode of the task. 1: countersign task; 2: or-sign task **Example value**: 2 **Data validation rules**: - Value range: `1` ～ `2` |
|   `is_milestone` | `boolean` | No | Whether is a milestone task or not **Example value**: false |
|   `custom_fields` | `input_custom_field_value[]` | No | Custom field values. If updted, the value needs to be set to the corresponding field according to the field type. * When `type` is "number", the `number_value` field should be used to represent the value of number custom field; * When "type" is "member", the `member_value` field should be used to represent the value of member custom field; * When `type` is "datetime", `datetime_value` field should be used to represent the value of the datetime custom field. * When `type` is "single_select", the `single_select_value` field should be used to represent the value of single_select custom field. * When `type` is "multi_select", the `multi_select_value` field should be used to represent the value of multi_select custom field. * When `type` is "text", the `text_value` field should be used to represent the value of text custom field. |
|     `guid` | `string` | Yes | Custom field guid **Example value**: "73b21903-0041-4796-a11e-f8be919a7063" |
|     `number_value` | `string` | No | A custom field value of a number type, fill in the string representation of a valild number. Or unset the value with an empty string. **Example value**: "10.23" **Data validation rules**: - Maximum length: `20` characters |
|     `member_value` | `member[]` | No | Custom field values for member types. The one or multiple users can be added (following the member format, only user type is supported). Only one value can be passed in when the field setting is "multi=false". Value can be unset with an empty array. **Data validation rules**: - Maximum length: `50` |
|       `id` | `string` | No | Indicates the id of member **Example value**: "ou_2cefb2f014f8d0c6c2d2eb7bafb0e54f" **Data validation rules**: - Maximum length: `100` characters |
|       `type` | `string` | No | Type of member **Example value**: "user" **Default value**: `user` |
|     `datetime_value` | `string` | No | Date type Custom field value, you can enter a millisecond string representing the date. Set to empty string means set to empty. **Example value**: "1698192000000" |
|     `single_select_value` | `string` | No | single_select field value, you can fill in a option_guid of the field options, or unset it with an empty string. **Example value**: "73b21903-0041-4796-a11e-f8be919a7063" **Data validation rules**: - Maximum length: `100` characters |
|     `multi_select_value` | `string[]` | No | multi_select type field value, you can fill in one or more option_guid of this field, or unset it with an empty array. **Example value**: ["73b21903-0041-4796-a11e-f8be919a7063"] **Data validation rules**: - Maximum length: `50` |
|     `text_value` | `string` | No | Text custom field value, up to 3000 characters. Use empty string "" to unset the value. **Example value**: "文本类型字段值。可以输入一段文本。空字符串表示清空。" |
| `update_fields` | `string[]` | Yes | Set the fields that need to be modified  - `summary`: Task summary - `description`: Task description - `start`: Task start time - `due`: Task due time Task completion time - `extra`: Task attached custom data - `custom_complete`: Task custom completion config Task repeat rule  **Example value**: ["summary"] | ### Request body example

{
    "task": {
        "summary": "One review for full-year sales",
        "description": "You need to read the review summary document in advance.",
        "due": {
            "timestamp": "1675454764000",
            "is_all_day": true
        },
        "extra": "dGVzdA==",
        "completed_at": "1675742789470",
        "repeat_rule": "FREQ=WEEKLY;INTERVAL=1;BYDAY=MO,TU,WE,TH,FR",
        "custom_complete": {
            "pc": {
                "href": "https://www.example.com",
                "tip": {
                    "en_us": "Workbench",
                    "zh_cn": "Workbench",
                    "zh_hk": "工作臺",
                    "zh_tw": "工作臺",
                    "ja_jp": "作業台",
                    "fr_fr": "Table de travail",
                    "it_it": "banco di lavoro",
                    "de_de": "Werkbank",
                    "ru_ru": "верстак",
                    "th_th": "โต๊ะทำงาน",
                    "es_es": "banco de trabajo",
                    "ko_kr": "작업대"
                }
            },
            "ios": {
                "href": "https://www.example.com",
                "tip": {
                    "en_us": "Workbench",
                    "zh_cn": "Workbench",
                    "zh_hk": "工作臺",
                    "zh_tw": "工作臺",
                    "ja_jp": "作業台",
                    "fr_fr": "Table de travail",
                    "it_it": "banco di lavoro",
                    "de_de": "Werkbank",
                    "ru_ru": "верстак",
                    "th_th": "โต๊ะทำงาน",
                    "es_es": "banco de trabajo",
                    "ko_kr": "작업대"
                }
            },
            "android": {
                "href": "https://www.example.com",
                "tip": {
                    "en_us": "Workbench",
                    "zh_cn": "Workbench",
                    "zh_hk": "工作臺",
                    "zh_tw": "工作臺",
                    "ja_jp": "作業台",
                    "fr_fr": "Table de travail",
                    "it_it": "banco di lavoro",
                    "de_de": "Werkbank",
                    "ru_ru": "верстак",
                    "th_th": "โต๊ะทำงาน",
                    "es_es": "banco de trabajo",
                    "ko_kr": "작업대"
                }
            }
        },
        "start": {
            "timestamp": "1675454764000",
            "is_all_day": true
        },
        "mode": 2,
        "is_milestone": false,
        "custom_fields": [
            {
                "guid": "73b21903-0041-4796-a11e-f8be919a7063",
                "number_value": "10.23",
                "member_value": [
                    {
                        "id": "ou_2cefb2f014f8d0c6c2d2eb7bafb0e54f",
                        "type": "user"
                    }
                ],
                "datetime_value": "1698192000000",
                "single_select_value": "73b21903-0041-4796-a11e-f8be919a7063",
                "multi_select_value": [
                    "73b21903-0041-4796-a11e-f8be919a7063"
                ],
                "text_value": "文本类型字段值。可以输入一段文本。空字符串表示清空。"
            }
        ]
    },
    "update_fields": [
        "summary"
    ]
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `task` | `task` | Updated task |
|     `guid` | `string` | Task GUID |
|     `summary` | `string` | Task summary |
|     `description` | `string` | mission description |
|     `due` | `due` | Task due |
|       `timestamp` | `string` | The timestamp of the expiration time/date in milliseconds from 1970-01-01 00:00:00. If the expiration time is a date, you need to convert the date to timestamp and set is_all_day = true |
|       `is_all_day` | `boolean` | Whether to expire on a date. If set to true, only the date part of the timestamp will be parsed and stored. |
|     `reminders` | `reminder[]` | List of reminders for tasks. Currently there is a maximum of 1 per task. |
|       `id` | `string` | ID of reminder time setting |
|       `relative_fire_minute` | `int` | The number of minutes of reminder time relative to the deadline. For example, 30 means 30 minutes before the deadline; 0 means the deadline reminder. |
|     `creator` | `member` | Task creator |
|       `id` | `string` | Indicates the id of member |
|       `type` | `string` | Type of member |
|       `role` | `string` | task role |
|     `members` | `member[]` | Task member list |
|       `id` | `string` | Indicates the id of member |
|       `type` | `string` | Type of member |
|       `role` | `string` | task role |
|     `completed_at` | `string` | Task completion timestamp (ms) |
|     `attachments` | `attachment[]` | List of attachments for tasks |
|       `guid` | `string` | Attachment guid |
|       `file_token` | `string` | Attachment token in cloud document system |
|       `name` | `string` | Attachment name |
|       `size` | `int` | Byte size of attachment |
|       `resource` | `resource` | Resources attributed to attachments |
|         `type` | `string` | resource type |
|         `id` | `string` | Resource ID |
|       `uploader` | `member` | Attachment uploader |
|         `id` | `string` | Indicates the id of member |
|         `type` | `string` | Type of member |
|         `role` | `string` | task role |
|       `is_cover` | `boolean` | Is it a cover image? |
|       `uploaded_at` | `string` | Upload timestamp (ms) |
|     `origin` | `origin` | Third-party platform source information associated with the task. Created is set and cannot be changed. |
|       `platform_i18n_name` | `i18n_text` | The name of the task import source to display on the task center details page. A multilingual version is required. |
|         `en_us` | `string` | English |
|         `zh_cn` | `string` | Chinese |
|         `zh_hk` | `string` | Chinese (Hong Kong) |
|         `zh_tw` | `string` | Chinese (Taiwan) |
|         `ja_jp` | `string` | Japanese |
|         `fr_fr` | `string` | French |
|         `it_it` | `string` | Italian |
|         `de_de` | `string` | German |
|         `ru_ru` | `string` | Russian |
|         `th_th` | `string` | Thai |
|         `es_es` | `string` | Spanish |
|         `ko_kr` | `string` | Korean |
|       `href` | `href` | Source of task association Platform details page link |
|         `url` | `string` | The address corresponding to the link, if filled in, must start with https://or http:// |
|         `title` | `string` | The title corresponding to the link |
|     `extra` | `string` | Custom data that accompanies the task. |
|     `tasklists` | `task_in_tasklist_info[]` | The name of the manifest to which the task belongs. The caller can only see a list of manifests that have permission to access. |
|       `tasklist_guid` | `string` | The guid of the task list |
|       `section_guid` | `string` | Custom grouping guid for task list |
|     `repeat_rule` | `string` | repeat_rule of task |
|     `parent_task_guid` | `string` | If the current task is a subtask of a task, return the guid of the parent task |
|     `mode` | `int` | The mode of the task. 1: countersign tasks; 2: or-sign tasks |
|     `source` | `int` | Source of task creation **Optional values are**:  - `0`: Unknown source - `1`: Mission Center - `2`: Group Task/Message Transfer Task - `6`: Tasks created with tenant_access_token authorization through the open platform - `7`: Tasks created with user_access_token authorization through the open platform - `8`: Document task  |
|     `custom_complete` | `custom_complete` | Custom completion configuration for tasks |
|       `pc` | `custom_complete_item` | pc client side custom configuration (including mac and windows) |
|         `href` | `string` | Customize the completed redirect url |
|         `tip` | `i18n_text` | The pop-up prompt for customization completion is |
|           `en_us` | `string` | English |
|           `zh_cn` | `string` | Chinese |
|           `zh_hk` | `string` | Chinese (Hong Kong) |
|           `zh_tw` | `string` | Chinese (Taiwan) |
|           `ja_jp` | `string` | Japanese |
|           `fr_fr` | `string` | French |
|           `it_it` | `string` | Italian |
|           `de_de` | `string` | German |
|           `ru_ru` | `string` | Russian |
|           `th_th` | `string` | Thai |
|           `es_es` | `string` | Spanish |
|           `ko_kr` | `string` | Korean |
|       `ios` | `custom_complete_item` | Custom completion configuration of ios |
|         `href` | `string` | Customize the completed redirect url |
|         `tip` | `i18n_text` | The pop-up prompt for customization completion is |
|           `en_us` | `string` | English |
|           `zh_cn` | `string` | Chinese |
|           `zh_hk` | `string` | Chinese (Hong Kong) |
|           `zh_tw` | `string` | Chinese (Taiwan) |
|           `ja_jp` | `string` | Japanese |
|           `fr_fr` | `string` | French |
|           `it_it` | `string` | Italian |
|           `de_de` | `string` | German |
|           `ru_ru` | `string` | Russian |
|           `th_th` | `string` | Thai |
|           `es_es` | `string` | Spanish |
|           `ko_kr` | `string` | Korean |
|       `android` | `custom_complete_item` | Custom completion configuration of android |
|         `href` | `string` | Customize the completed redirect url |
|         `tip` | `i18n_text` | The pop-up prompt for customization completion is |
|           `en_us` | `string` | English |
|           `zh_cn` | `string` | Chinese |
|           `zh_hk` | `string` | Chinese (Hong Kong) |
|           `zh_tw` | `string` | Chinese (Taiwan) |
|           `ja_jp` | `string` | Japanese |
|           `fr_fr` | `string` | French |
|           `it_it` | `string` | Italian |
|           `de_de` | `string` | German |
|           `ru_ru` | `string` | Russian |
|           `th_th` | `string` | Thai |
|           `es_es` | `string` | Spanish |
|           `ko_kr` | `string` | Korean |
|     `task_id` | `string` | Code on the task interface |
|     `created_at` | `string` | Task creation timestamp (ms) |
|     `updated_at` | `string` | Last updated timestamp of the task (ms) |
|     `status` | `string` | The status of the task, supports "todo" and "done" states |
|     `url` | `string` | Task applink |
|     `start` | `start` | Start time of the task |
|       `timestamp` | `string` | The timestamp of the start time/date in milliseconds from 1970-01-01 00:00:00. If the start time is a date, you need to convert the date to timestamp and set is_all_day = true |
|       `is_all_day` | `boolean` | Whether to start on a date. If set to true, only the date part of the timestamp will be parsed and stored. |
|     `subtask_count` | `int` | The number of subtasks of this task. |
|     `is_milestone` | `boolean` | Whether is a milestone task or not |
|     `custom_fields` | `custom_field_value[]` | Custom field values for tasks |
|       `guid` | `string` | Field GUID |
|       `type` | `string` | Custom field types, support "member", "datetime", "number", "single_select", "multi_select" five types |
|       `number_value` | `string` | A custom field value of a numeric type, fill in the string representation of a legal number, and set the empty string representation to empty. |
|       `datetime_value` | `string` | Date type Custom field value. You can enter a millisecond string representing the date. Set to empty string means set to empty. |
|       `member_value` | `member[]` | Custom field value for person type, you can set the id of 1 or more users (follow member format, only user type is supported). Only one value can be entered when the field is set to "cannot select more". Set to empty Array means set to empty. |
|         `id` | `string` | Indicates the id of member |
|         `type` | `string` | Type of member |
|         `role` | `string` | Member role |
|       `single_select_value` | `string` | Radio Type field value, fill in the option_guid of a field option. Set to empty string means set to empty. |
|       `multi_select_value` | `string[]` | Multiple select type field value, you can fill in one or more option_guid of this field. Set to empty Array means set to empty. |
|       `name` | `string` | Custom field name |
|       `text_value` | `string` | text field value. |
|     `dependencies` | `task_dependency[]` | task dependence |
|       `type` | `string` | dependency type **Optional values are**:  - `prev`: prev-dependency - `next`: next-dependency  |
|       `task_guid` | `string` | GUIDs for dependent tasks | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "task": {
            "guid": "83912691-2E43-47fc-94a4-d512e03984fa",
            "summary": "Conduct a mid-year sales summary",
            "description": "Conduct a mid-year sales summary",
            "due": {
                "timestamp": "1675454764000",
                "is_all_day": true
            },
            "reminders": [
                {
                    "id": "10",
                    "relative_fire_minute": 30
                }
            ],
            "creator": {
                "id": "ou_2cefb2f014f8d0c6c2d2eb7bafb0e54f",
                "type": "user",
                "role": "Assignee"
            },
            "members": [
                {
                    "id": "ou_2cefb2f014f8d0c6c2d2eb7bafb0e54f",
                    "type": "user",
                    "role": "Assignee"
                }
            ],
            "completed_at": "1675742789470",
            "attachments": [
                {
                    "guid": "f860de3e-6881-4ddd-9321-070f36d1af0b",
                    "file_token": "boxcnTDqPaRA6JbYnzQsZ2doB2b",
                    "name": "foo.jpg",
                    "size": 62232,
                    "resource": {
                        "type": "task",
                        "id": "e6e37dcc-f75a-5936-f589-12fb4b5c80c2"
                    },
                    "uploader": {
                        "id": "ou_2cefb2f014f8d0c6c2d2eb7bafb0e54f",
                        "type": "user",
                        "role": "editor"
                    },
                    "is_cover": false,
                    "uploaded_at": "1675742789470"
                }
            ],
            "origin": {
                "platform_i18n_name": {
                    "en_us": "Workbench",
                    "zh_cn": "Workbench",
                    "zh_hk": "工作臺",
                    "zh_tw": "工作臺",
                    "ja_jp": "作業台",
                    "fr_fr": "Table de travail",
                    "it_it": "banco di lavoro",
                    "de_de": "Werkbank",
                    "ru_ru": "верстак",
                    "th_th": "โต๊ะทำงาน",
                    "es_es": "banco de trabajo",
                    "ko_kr": "작업대"
                },
                "href": {
                    "url": "https://www.example.com",
                    "title": "反馈一个问题，需要协助排查"
                }
            },
            "extra": "dGVzdA ==",
            "tasklists": [
                {
                    "tasklist_guid": "cc371766-6584-cf50-a222-c22cd9055004",
                    "section_guid": "e6e37dcc-f75a-5936-f589-12fb4b5c80c2"
                }
            ],
            "repeat_rule": "FREQ=WEEKLY;INTERVAL=1;BYDAY=MO,TU, WE,TH,FR",
            "parent_task_guid": "e297ddff-06ca-4166-b917-4ce57cd3a7a0",
            "mode": 2,
            "source": 6,
            "custom_complete": {
                "pc": {
                    "href": "https://www.example.com",
                    "tip": {
                        "en_us": "Workbench",
                        "zh_cn": "Workbench",
                        "zh_hk": "工作臺",
                        "zh_tw": "工作臺",
                        "ja_jp": "作業台",
                        "fr_fr": "Table de travail",
                        "it_it": "banco di lavoro",
                        "de_de": "Werkbank",
                        "ru_ru": "верстак",
                        "th_th": "โต๊ะทำงาน",
                        "es_es": "banco de trabajo",
                        "ko_kr": "작업대"
                    }
                },
                "ios": {
                    "href": "https://www.example.com",
                    "tip": {
                        "en_us": "Workbench",
                        "zh_cn": "Workbench",
                        "zh_hk": "工作臺",
                        "zh_tw": "工作臺",
                        "ja_jp": "作業台",
                        "fr_fr": "Table de travail",
                        "it_it": "banco di lavoro",
                        "de_de": "Werkbank",
                        "ru_ru": "верстак",
                        "th_th": "โต๊ะทำงาน",
                        "es_es": "banco de trabajo",
                        "ko_kr": "작업대"
                    }
                },
                "android": {
                    "href": "https://www.example.com",
                    "tip": {
                        "en_us": "Workbench",
                        "zh_cn": "Workbench",
                        "zh_hk": "工作臺",
                        "zh_tw": "工作臺",
                        "ja_jp": "作業台",
                        "fr_fr": "Table de travail",
                        "it_it": "banco di lavoro",
                        "de_de": "Werkbank",
                        "ru_ru": "верстак",
                        "th_th": "โต๊ะทำงาน",
                        "es_es": "banco de trabajo",
                        "ko_kr": "작업대"
                    }
                }
            },
            "task_id": "t6272302",
            "created_at": "1675742789470",
            "updated_at": "1675742789470",
            "status": "todo",
            "url": "https://applink.larksuite.com/client/todo/detail?guid=70577c8f-91ab-4c91-b359-a21a751054e8&suite_entity_num=t192012",
            "start": {
                "timestamp": "1675454764000",
                "is_all_day": true
            },
            "subtask_count": 1,
            "is_milestone": false,
            "custom_fields": [
                {
                    "guid": "a4f648d7-76ef-477f-bc8e-0601b5a60093",
                    "type": "number",
                    "number_value": "10.23",
                    "datetime_value": "1687708260000",
                    "member_value": [
                        {
                            "id": "ou_2cefb2f014f8d0c6c2d2eb7bafb0e54f",
                            "type": "user",
                            "role": "editor"
                        }
                    ],
                    "single_select_value": "4216f79b-3fda-4dc6-a0c4-a16022e47152",
                    "multi_select_value": [
                        "4216f79b-3fda-4dc6-a0c4-a16022e47152"
                    ],
                    "name": "priority",
                    "text_value": "Here are some texts you want to save."
                }
            ],
            "dependencies": [
                {
                    "type": "next",
                    "task_guid": "93b7bd05-35e6-4371-b3c9-6b7cbd7100c0"
                }
            ]
        }
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1470400 | Incorrect request parameters, such as providing an invalid task GUID, trying to set the task summary to empty, etc. | The reason for the error is shown in the returned msg. |
| 404 | 1470404 | Task data does not exist or has been deleted. | Confirm that the task data you want to access exists or has been deleted. |
| 500 | 1470500 | Server error. | Try retrying the call. If it continues to fail, please contact support for feedback. |
| 403 | 1470403 | Missing edit permission to the task. | Check if the calling identity has edit permissions on the task. Please refer to the "How are Tasks Authorized?" section in Task Features Overview. |
