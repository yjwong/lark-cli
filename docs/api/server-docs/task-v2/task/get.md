---
title: "Get Task Details"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/task/get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/task/v2/tasks/:task_guid"
service: "task-v2"
resource: "task"
section: "Tasks v2"
rate_limit: "10 per second"
scopes:
  - "task:task:read"
  - "task:task:write"
updated: "1699521559000"
---

# Get Task Details

This api is used to obtain task details, including task summary, description, time, members and other information.

> Getting the task details requires the calling identity to have read permissions on the task. Please refer to the "How are Tasks Authorized?" section in Task Features Overview.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/task/v2/tasks/:task_guid |
| HTTP Method | GET |
| Rate Limit | 10 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `task:task:read` `task:task:write` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `task_guid` | `string` | GUID of task to get **Example value**: "e297ddff-06ca-4166-b917-4ce57cd3a7a0" **Data validation rules**: - Maximum length: `100` characters | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Default value**: `open_id` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `task` | `task` | Task details |
|     `guid` | `string` | Task guid |
|     `summary` | `string` | Task summary |
|     `description` | `string` | Task description |
|     `due` | `due` | Task due |
|       `timestamp` | `string` | The timestamp of the expiration time/date in milliseconds from 1970-01-01 00:00:00 UTC. If the expiration time is a date, you need to convert the date to timestamp and set is_all_day = true |
|       `is_all_day` | `boolean` | Whether to due on a date. If set to true, only the date part of the timestamp will be parsed and stored. |
|     `reminders` | `reminder[]` | Task reminder list. Currently only 1 reminder is supported per task. |
|       `id` | `string` | Task Reminder ID |
|       `relative_fire_minute` | `int` | The number of minutes relative to the deadline. For example, 30 means 30 minutes before the due; 0 means the due time. |
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
|         `role` | `string` | role |
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
|     `tasklists` | `task_in_tasklist_info[]` | The tasklists to which the task belongs. The caller can only see a list of tasklist that have permission to access. |
|       `tasklist_guid` | `string` | The guid of the task list |
|       `section_guid` | `string` | Section GUID of the tasklist |
|     `repeat_rule` | `string` | If the task is a recurring task, return its repeat rule. |
|     `parent_task_guid` | `string` | If the current task is a subtask of a task, return the parent task GUID |
|     `mode` | `int` | The mode of the task. 1: countersign tasks; 2: or-sign tasks |
|     `source` | `int` | Source of task creation **Optional values are**:  - `0`: Unknown source - `1`: Task Center - `2`: Group Task/Message Transfer Task - `6`: Tasks created with tenant_access_token authorization through the open platform - `7`: Tasks created with user_access_token authorization through the open platform - `8`: Document task  |
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
|     `url` | `string` | Share link for task |
|     `start` | `start` | Start time of the task |
|       `timestamp` | `string` | The timestamp of the start time/date in milliseconds from 1970-01-01 00:00:00. If the start time is a date, you need to convert the date to timestamp and set is_all_day = true |
|       `is_all_day` | `boolean` | Whether to start on a date. If set to true, only the date part of the timestamp will be parsed and stored. |
|     `subtask_count` | `int` | The number of subtasks of this task. |
|     `is_milestone` | `boolean` | Is it a milestone task |
|     `custom_fields` | `custom_field_value[]` | Custom field values for tasks |
|       `guid` | `string` | Field GUID |
|       `type` | `string` | Custom field types, support "member", "datetime", "number", "single_select", "multi_select". |
|       `number_value` | `string` | A custom field value of a number type, fill in the string representation of a legal number, and set the empty string to unset the value. |
|       `datetime_value` | `string` | Datetime custom field value. You can enter a timestamp (in ms) string representing the date. Set to empty string means to unset the value. |
|       `member_value` | `member[]` | Custom field value for member type, you can set the id of 1 or more users (follow member format, only user type is supported). Only one value can be entered when the field setting is "multi=false". Set to empty array means set to empty. |
|         `id` | `string` | Indicates the id of member |
|         `type` | `string` | Type of member |
|         `role` | `string` | member role |
|       `single_select_value` | `string` | single select field value, fill in the `option_guid` of a field option. An empty string means to unset the value. |
|       `multi_select_value` | `string[]` | multi select type field value, you can fill in one or more option_guid of this field. Set to empty Array means to unset the value. |
|       `name` | `string` | Custom field name |
|       `text_value` | `string` | text field value |
|     `dependencies` | `task_dependency[]` | task dependence |
|       `type` | `string` | dependency type **Optional values are**:  - `prev`: prev-dependency - `next`: next-dependency  |
|       `task_guid` | `string` | GUIDs for dependent tasks | ### Response body example

{
    "code": 0,
    "msg": "Success",
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
                    "name": "Foo.jpg",
                    "size": 62232,
                    "resource": {
                        "type": "Task",
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
                    "es_es": "Banco de trabajo",
                    "ko_kr": "You're a cowboy"
                },
                "href": {
                    "url": "https://www.example.com",
                    "title": "Feedback a problem and need assistance in troubleshooting."
                }
            },
            "extra": "dGVzdA==",
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
                        "zh_hk": "Work",
                        "zh_tw": "Work",
                        "ja_jp": "Workbench",
                        "fr_fr": "Table de travail",
                        "it_it": "Banco di lavoro",
                        "de_de": "Werkbank",
                        "ru_ru": "верстак",
                        "th_th": "โต๊ะทำงาน",
                        "es_es": "Banco de trabajo",
                        "ko_kr": "You're a cowboy"
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
                        "es_es": "Banco de trabajo",
                        "ko_kr": "You're a cowboy"
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
                        "es_es": "Banco de trabajo",
                        "ko_kr": "You're a cowboy"
                    }
                }
            },
            "task_id": "t6272302",
            "created_at": "1675742789470",
            "updated_at": "1675742789470",
            "status": "Todo",
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
| 400 | 1470400 | Incorrect request parameters, such as an invalid task GUID passed in. | The reason for the error is shown in the returned msg. |
| 500 | 1470500 | Failed to get task details, server error. | Try retrying the call. If it continues to fail, please contact support for feedback. |
| 404 | 1470404 | The task does not exist or has been deleted. | Confirm that the task data you want to access exists or has been deleted. |
| 403 | 1470403 | Lack read permission to the task. | Check that the calling identity has read permission to the task. Please refer to the "How are Tasks Authorized?" section in Task Features Overview. |
