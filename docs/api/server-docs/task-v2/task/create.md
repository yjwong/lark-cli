---
title: "Create Task"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/task/create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/task/v2/tasks"
service: "task-v2"
resource: "task"
section: "Tasks v2"
rate_limit: "10 per second"
scopes:
  - "task:task:write"
updated: "1699521555000"
---

# Create Task

This API can create a task. When creating a task, it supports filling in the basic information of the task (such as title, description, person in charge, etc.). In addition, it can also set the start time of the task, deadline reminders and other conditions. In addition, you can pass in the tasklists field to add new tasks to multiple lists.

When creating a task, "assignees" and "followers" can be added by specifying `members` fields. For the format of member, please refer to the " How to represent members of tasks and tasklists?" section in Feature Overview .

When setting task start or due time, you need to follow the specification and constraints of task time. Please refer to the "How to use start and due time?" section in Feature Overview .

This api supports "idempotent invocation" by specifying `client_token` field. Please refer to the "Idempotent Invocation" section in Feature Overview .

To create a subtask of a task, you need to use the Create Subtask API.

Creating tasks allows setting custom field values at the same time. However, according to the authorization model of custom fields, only values of fields associated with tasklists can only be added to tasks. Please refer to Custom Field Feature Overview for details.

> If want to add new created task to several tasklists, the calling identity must have edit permission on all these tasklists. Please refer to the "How are tasklist authorized?" section in Tasklist Features Overview .

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/task/v2/tasks |
| HTTP Method | POST |
| Rate Limit | 10 per second |
| Supported app types | custom,isv |
| Required scopes | `task:task:write` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Default value**: `open_id` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `summary` | `string` | Yes | Task summary. Empty is not allowed; supports up to 3000 utf8 characters. **Example value**: "One review for full-year sales" |
| `description` | `string` | No | Task description. Supports up to 3000 utf8 characters. **Example value**: "You need to read the review summary document in advance." |
| `due` | `due` | No | Task due time **Example value**: 1675742789470 |
|   `timestamp` | `string` | No | The timestamp of the due time/date, in milliseconds from 1970-01-01 00:00:00 UTC. If the expiration time is a date, you need to convert the date to timestamp and set is_all_day = true **Example value**: "1675454764000" |
|   `is_all_day` | `boolean` | No | Whether to due on a date. If set to true, only the date part of the timestamp will be parsed and stored. **Example value**: true |
| `origin` | `origin` | No | The third-party platform source information associated with the task is used to display the source information in the App Task UI. Origin can only be set when task is created and not changable. Please refer to the "How to use origin?" section in Feature Overview . |
|   `platform_i18n_name` | `i18n_text` | No | The name of the task import source to display on the task center details page. A multilingual version is required. |
|     `en_us` | `string` | No | English **Example value**: "Workbench" **Data validation rules**: - Maximum length: `1000` characters |
|     `zh_cn` | `string` | No | Chinese **Example value**: "Workbench" **Data validation rules**: - Maximum length: `1000` characters |
|     `zh_hk` | `string` | No | Chinese (Hong Kong) **Example value**: "Work" **Data validation rules**: - Maximum length: `1000` characters |
|     `zh_tw` | `string` | No | Chinese (Taiwan) **Example value**: "Work" **Data validation rules**: - Maximum length: `1000` characters |
|     `ja_jp` | `string` | No | Japanese **Example value**: "Workbench" **Data validation rules**: - Maximum length: `1000` characters |
|     `fr_fr` | `string` | No | French **Example value**: "Table de travail" |
|     `it_it` | `string` | No | Italian **Example value**: "Banco di lavoro" |
|     `de_de` | `string` | No | German **Example value**: "Werkbank" |
|     `ru_ru` | `string` | No | Russian **Example value**: "верстак" |
|     `th_th` | `string` | No | Thai **Example value**: "โต๊ะทำงาน" |
|     `es_es` | `string` | No | Spanish **Example value**: "banco de trabajo" |
|     `ko_kr` | `string` | No | Korean **Example value**: "작업대" |
|   `href` | `href` | No | Source of task association Platform details page link |
|     `url` | `string` | No | The address corresponding to the source link, if filled in, must start with https://or http://. **Example value**: "https://www.example.com" **Data validation rules**: - Length range: `0` ～ `1024` characters |
|     `title` | `string` | No | The title corresponding to the source link **Example value**: "feedback, need assistance" **Data validation rules**: - Maximum length: `512` characters |
| `extra` | `string` | No | Any data that the caller can pass in attached to the task. It will be returned as it is when getting the task details. If it is binary data, it can be encoded with Base64. **Example value**: "dGVzdA==" **Data validation rules**: - Maximum length: `65536` characters |
| `completed_at` | `string` | No | The completion time timestamp (ms) of the task. Fill in or set to 0 to create an unfinished task; fill in a specific timestamp to create a completed task. **Example value**: "1675742789470" **Default value**: `0` **Data validation rules**: - Maximum length: `20` characters |
| `members` | `member[]` | No | A list of task members, including assignees and followers. If not filled in, the task has no members. Support up to 50 members per request after de-duplication. Please refer to the " How to represent members of tasks and tasklists?" section in Feature Overview . **Example value**: [ou_2cefb2f014f8d0c6c2d2eb7bafb0e54f] |
|   `id` | `string` | Yes | Member ID **Example value**: "ou_2cefb2f014f8d0c6c2d2eb7bafb0e54f" **Data validation rules**: - Maximum length: `100` characters |
|   `type` | `string` | No | The type of member, which can be user or app. **Example value**: "user" **Default value**: `user` |
|   `role` | `string` | Yes | Member roles, which can be "assignee" or "follower". **Example value**: "assignee" **Data validation rules**: - Maximum length: `20` characters |
| `repeat_rule` | `string` | No | Task repeat rule. If set, the task is "recurring task". Please refer to the "How to Use Recurring Tasks? " section in Task Feature Overview . **Example value**: "FREQ=WEEKLY;INTERVAL=1;BYDAY=MO,TU,WE,TH,FR" **Data validation rules**: - Maximum length: `1000` characters |
| `custom_complete` | `custom_complete` | No | Task custom complete configuration. Please refer to the "How to use custom completion?" section in Task Feature Overview . |
|   `pc` | `custom_complete_item` | No | pc client side custom configuration (including mac and windows) |
|     `href` | `string` | No | Customize the completed redirect url **Example value**: "https://www.example.com" |
|     `tip` | `i18n_text` | No | Custom completed pop-up prompt |
|       `en_us` | `string` | No | English **Example value**: "Workbench" **Data validation rules**: - Maximum length: `1000` characters |
|       `zh_cn` | `string` | No | Chinese **Example value**: "Workbench" **Data validation rules**: - Maximum length: `1000` characters |
|       `zh_hk` | `string` | No | Chinese (Hong Kong) **Example value**: "Work" **Data validation rules**: - Maximum length: `1000` characters |
|       `zh_tw` | `string` | No | Chinese (Taiwan) **Example value**: "Work" **Data validation rules**: - Maximum length: `1000` characters |
|       `ja_jp` | `string` | No | Japanese **Example value**: "Workbench" **Data validation rules**: - Maximum length: `1000` characters |
|       `fr_fr` | `string` | No | French **Example value**: "Table de travail" |
|       `it_it` | `string` | No | Italian **Example value**: "Banco di lavoro" |
|       `de_de` | `string` | No | German **Example value**: "Werkbank" |
|       `ru_ru` | `string` | No | Russian **Example value**: "верстак" |
|       `th_th` | `string` | No | Thai **Example value**: "โต๊ะทำงาน" |
|       `es_es` | `string` | No | Spanish **Example value**: "banco de trabajo" |
|       `ko_kr` | `string` | No | Korean **Example value**: "작업대" |
|   `ios` | `custom_complete_item` | No | Custom completion configuration of ios |
|     `href` | `string` | No | Customize the completed redirect url **Example value**: "https://www.example.com" |
|     `tip` | `i18n_text` | No | The pop-up prompt for customization completion is |
|       `en_us` | `string` | No | English **Example value**: "Workbench" **Data validation rules**: - Maximum length: `1000` characters |
|       `zh_cn` | `string` | No | Chinese **Example value**: "Workbench" **Data validation rules**: - Maximum length: `1000` characters |
|       `zh_hk` | `string` | No | Chinese (Hong Kong) **Example value**: "Work" **Data validation rules**: - Maximum length: `1000` characters |
|       `zh_tw` | `string` | No | Chinese (Taiwan) **Example value**: "Work" **Data validation rules**: - Maximum length: `1000` characters |
|       `ja_jp` | `string` | No | Japanese **Example value**: "Workbench" **Data validation rules**: - Maximum length: `1000` characters |
|       `fr_fr` | `string` | No | French **Example value**: "Table de travail" |
|       `it_it` | `string` | No | Italian **Example value**: "Banco di lavoro" |
|       `de_de` | `string` | No | German **Example value**: "Werkbank" |
|       `ru_ru` | `string` | No | Russian **Example value**: "верстак" |
|       `th_th` | `string` | No | Thai **Example value**: "โต๊ะทำงาน" |
|       `es_es` | `string` | No | Spanish **Example value**: "banco de trabajo" |
|       `ko_kr` | `string` | No | Korean **Example value**: "작업대" |
|   `android` | `custom_complete_item` | No | Custom completion configuration of android |
|     `href` | `string` | No | Customize the completed redirect url **Example value**: "https://www.example.com" |
|     `tip` | `i18n_text` | No | The pop-up prompt for customization completion is |
|       `en_us` | `string` | No | English **Example value**: "Workbench" **Data validation rules**: - Maximum length: `1000` characters |
|       `zh_cn` | `string` | No | Chinese **Example value**: "Workbench" **Data validation rules**: - Maximum length: `1000` characters |
|       `zh_hk` | `string` | No | Chinese (Hong Kong) **Example value**: "Work" **Data validation rules**: - Maximum length: `1000` characters |
|       `zh_tw` | `string` | No | Chinese (Taiwan) **Example value**: "Work" **Data validation rules**: - Maximum length: `1000` characters |
|       `ja_jp` | `string` | No | Japanese **Example value**: "Workbench" **Data validation rules**: - Maximum length: `1000` characters |
|       `fr_fr` | `string` | No | French **Example value**: "Table de travail" |
|       `it_it` | `string` | No | Italian **Example value**: "Banco di lavoro" |
|       `de_de` | `string` | No | German **Example value**: "Werkbank" |
|       `ru_ru` | `string` | No | Russian **Example value**: "верстак" |
|       `th_th` | `string` | No | Thai **Example value**: "โต๊ะทำงาน" |
|       `es_es` | `string` | No | Spanish **Example value**: "banco de trabajo" |
|       `ko_kr` | `string` | No | Korean **Example value**: "작업대" |
| `tasklists` | `task_in_tasklist_info[]` | No | Information on the list of tasks |
|   `tasklist_guid` | `string` | No | Specifies that a task is created in a manifest that requires editable permissions to the manifest. If not filled in, the task will not be created in the manifest. **Example value**: "cc371766-6584-cf50-a222-c22cd9055004" **Data validation rules**: - Maximum length: `100` characters |
|   `section_guid` | `string` | No | The GUID of the custom group in the manifest is used to specify that the task is added to a specific group while creating the task in a manifest. If the GUID of the manifest is filled in, but the GUID of the group is not filled in, it will be automatically added to the default group of the manifest. **Example value**: "e6e37dcc-f75a-5936-f589-12fb4b5c80c2" |
| `client_token` | `string` | No | Idempotent token. If provided triggers the dempotent behavior. Please refer to the "Idempotent Invocation" section in Feature Overview . **Example value**: "daa2237f-8310-4707-a83b-52c8a81e0fb7" **Data validation rules**: - Length range: `10` ～ `100` characters |
| `start` | `start` | No | Task start time (ms). Please refer to the "How to use start and due time?" section in Feature Overview . |
|   `timestamp` | `string` | No | The timestamp of the start time/date in milliseconds from 1970-01-01 00:00:00. If the start time is a date, you need to convert the date to timestamp and set is_all_day = true. If you set both the start time and the deadline for the task, the start time must be 
{
    "summary": "One review for full-year sales",
    "description": "You need to read the review summary document in advance.",
    "due": {
        "timestamp": "1675454764000",
        "is_all_day": true
    },
    "origin": {
        "platform_i18n_name": {
            "en_us": "Workbench",
            "zh_cn": "Workbench",
            "zh_hk": "Work",
            "zh_tw": "Work",
            "ja_jp": "Workbench",
            "fr_fr": "Table de travail",
            "it_it": "Banco di lavoro",
            "de_de": "Werkbank",
            "ru_ru": "верстак",
            "th_th": "โต๊ะทำงาน",
            "es_es": "banco de trabajo",
            "ko_kr": "작업대"
        },
        "href": {
            "url": "https://www.example.com",
            "title": "feedback, need assistance"
        }
    },
    "extra": "dGVzdA==",
    "completed_at": "1675742789470",
    "members": [
        {
            "id": "ou_2cefb2f014f8d0c6c2d2eb7bafb0e54f",
            "type": "user",
            "role": "assignee"
        }
    ],
    "repeat_rule": "FREQ=WEEKLY;INTERVAL=1;BYDAY=MO,TU,WE,TH,FR",
    "custom_complete": {
        "pc": {
            "href": "https://www.example.com",
            "tip": {
                "en_us": "Workbench",
                "zh_cn": "Workbench",
                "zh_hk": "Work",
                "zh_tw": "Work",
                "ja_jp": "Workbench",
                "fr_fr": "Table de travail",
                "it_it": "Banco di lavoro",
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
                "zh_hk": "Work",
                "zh_tw": "Work",
                "ja_jp": "Workbench",
                "fr_fr": "Table de travail",
                "it_it": "Banco di lavoro",
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
                "zh_hk": "Work",
                "zh_tw": "Work",
                "ja_jp": "Workbench",
                "fr_fr": "Table de travail",
                "it_it": "Banco di lavoro",
                "de_de": "Werkbank",
                "ru_ru": "верстак",
                "th_th": "โต๊ะทำงาน",
                "es_es": "banco de trabajo",
                "ko_kr": "작업대"
            }
        }
    },
    "tasklists": [
        {
            "tasklist_guid": "cc371766-6584-cf50-a222-c22cd9055004",
            "section_guid": "e6e37dcc-f75a-5936-f589-12fb4b5c80c2"
        }
    ],
    "client_token": "daa2237f-8310-4707-a83b-52c8a81e0fb7",
    "start": {
        "timestamp": "1675454764000",
        "is_all_day": true
    },
    "reminders": [
        {
            "relative_fire_minute": 30
        }
    ],
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
            "text_value": "Here are something texts you want to save."
        }
    ]
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `task` | `task` | Generated tasks |
|     `guid` | `string` | Task guid, the unique ID of the task |
|     `summary` | `string` | Task title |
|     `description` | `string` | Task Notes |
|     `due` | `due` | Task due time. Please refer to the "How to use start and due time?" section in Feature Overview . |
|       `timestamp` | `string` | The timestamp of the due time/date in milliseconds from 1970-01-01 00:00:00. If the expiration time is a date, you need to convert the date to timestamp and set is_all_day = true |
|       `is_all_day` | `boolean` | Whether to due on a date. If set to true, only the date part of the timestamp will be parsed and stored. |
|     `reminders` | `reminder[]` | List of reminder configurations for tasks. Currently there is a maximum of 1 per task. |
|       `id` | `string` | Reminder ID |
|       `relative_fire_minute` | `int` | The number of reminder minutes relative to the due, is a non-negative integer. For example, set to 30 for "30 minutes before the deadline". Set to 0 for "reminder on deadline". |
|     `creator` | `member` | Task creator |
|       `id` | `string` | Indicates the id of member |
|       `type` | `string` | Type of member |
|       `role` | `string` | Member role |
|     `members` | `member[]` | Task member list |
|       `id` | `string` | Indicates the id of member |
|       `type` | `string` | Type of member |
|       `role` | `string` | Member role |
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
|         `role` | `string` | Role, which can be an assginee or a follower |
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
|         `url` | `string` | The address corresponding to the link. |
|         `title` | `string` | The title corresponding to the link |
|     `extra` | `string` | Custom data that accompanies the task. |
|     `tasklists` | `task_in_tasklist_info[]` | The name of the manifest to which the task belongs. The caller can only see a list of manifests that have permission to access. |
|       `tasklist_guid` | `string` | The guid of the task list |
|       `section_guid` | `string` | Custom grouping guid for task list |
|     `repeat_rule` | `string` | Repetition rules for tasks. |
|     `parent_task_guid` | `string` | If the current task is a subtask of a task, return the guid of the parent task |
|     `mode` | `int` | The mode of the task. 1 - countersign tasks; 2 - or sign tasks |
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
|         `href` | `string` | Customize the completed jump url |
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
|     `is_milestone` | `boolean` | Is it a milestone task |
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
|       `text_value` | `string` | text custom field value. |
|     `dependencies` | `task_dependency[]` | Task dependence |
|       `type` | `string` | dependency type **Optional values are**:  - `prev`: prev-dependency - `next`: next-dependency  |
|       `task_guid` | `string` | GUIDs for dependent tasks | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "task": {
            "guid": "83912691-2e43-47fc-94a4-d512e03984fa",
            "summary": "进行销售年中总结",
            "description": "进行销售年中总结",
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
                    "role": "editor"
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
                        "role": "Assignee"
                    },
                    "is_cover": false,
                    "uploaded_at": "1675742789470"
                }
            ],
            "origin": {
                "platform_i18n_name": {
                    "en_us": "Workbench",
                    "zh_cn": "Workbench",
                    "zh_hk": "Work",
                    "zh_tw": "Work",
                    "ja_jp": "Workbench",
                    "fr_fr": "Table de travail",
                    "it_it": "Banco di lavoro",
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
            "extra": "dGVzdA==",
            "tasklists": [
                {
                    "tasklist_guid": "cc371766-6584-cf50-a222-c22cd9055004",
                    "section_guid": "e6e37dcc-f75a-5936-f589-12fb4b5c80c2"
                }
            ],
            "repeat_rule": "FREQ=WEEKLY;INTERVAL=1;BYDAY=MO,TU,WE,TH,FR",
            "parent_task_guid": "e297ddff-06ca-4166-b917-4ce57cd3a7a0",
            "mode": 2,
            "source": 6,
            "custom_complete": {
                "pc": {
                    "href": "https://www.example.com",
                    "tip": {
                        "en_us": "Workbench",
                        "zh_cn": "Workbench",
                        "zh_hk": "Work",
                        "zh_tw": "Work",
                        "ja_jp": "Workbench",
                        "fr_fr": "Table de travail",
                        "it_it": "Banco di lavoro",
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
                        "zh_hk": "Work",
                        "zh_tw": "Work",
                        "ja_jp": "Workbench",
                        "fr_fr": "Table de travail",
                        "it_it": "Banco di lavoro",
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
                        "zh_hk": "Work",
                        "zh_tw": "Work",
                        "ja_jp": "Workbench",
                        "fr_fr": "Table de travail",
                        "it_it": "Banco di lavoro",
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
                    "name": "优先级",
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
| 400 | 1470400 | The request parameters are wrong, such as not providing required fields, the title is too long, trying to set a reminder but not providing a deadline, etc. | Refer to the returned error message for the specific error reason. |
| 404 | 1470404 | The join list/group is set when the task is created, but the list/group does not exist. | Refer to the returned error message for the specific error reason. |
| 500 | 1470500 | Failed to create task. | Internal server error. If there is a persistent error after retrying the call, you can contact support. |
| 403 | 1470403 | Permission error. When the task was created, it was set to add a list, but there was no editable permission for the list. | Check whether the calling identity has edit permission to the list to be added by the task. See [How is the list authenticated?] (/ssl: ttdoc/uAjLw4CM/ukTMukTMukTM/task-v2/faq) |
| 500 | 1470422 | Concurrent api calls using the same client_token | Do not invoke api concurrently with the same client_token. |
