---
title: "List tasks"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/task/list"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/task/v2/tasks"
service: "task-v2"
resource: "task"
section: "Tasks v2"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "task:task:read"
  - "task:task:write"
updated: "1699521578000"
---

# List Tasks

List all tasks of a specific type based on the calling identity. Paging is supported.

Currently, only tasks of "my_tasks" are supported. The returned task data is in the order in which the tasks are list by "custom" order in the "Owned" in Task Center.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/task/v2/tasks |
| HTTP Method | GET |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `task:task:read` `task:task:write` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `page_size` | `int` | No | Number of tasks per page **Example value**: 50 **Default value**: `50` **Data validation rules**: - Value range: `1` ～ `100` |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: aWQ9NzEwMjMzMjMxMDE= |
| `completed` | `boolean` | No | Whether or not to filter by task completion. Filling in true means only completed tasks are listed; filling in false means only uncompleted tasks are listed. Fill in means no filtering. **Example value**: true |
| `type` | `string` | No | The type of column task currently only supports "my_tasks", that is, "Owned" tasks. **Example value**: my_tasks **Default value**: `my_tasks` |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Default value**: `open_id` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `items` | `task[]` | List of tasks returned |
|     `guid` | `string` | Task guid, the unique ID of the task |
|     `summary` | `string` | Task title |
|     `description` | `string` | Task Notes |
|     `due` | `due` | Task deadline |
|       `timestamp` | `string` | The timestamp of the expiration time/date, in milliseconds from 1970-01-01 00:00:00 UTC. If the expiration time is a date, you need to convert the date to timestamp and set is_all_day = true |
|       `is_all_day` | `boolean` | Whether to expire on a date. If set to true, only the date part of the timestamp will be parsed and stored. |
|     `reminders` | `reminder[]` | Task reminder list. Currently only 1 reminder is supported per task. |
|       `id` | `string` | Reminder ID |
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
|         `role` | `string` | Member role |
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
|       `tasklist_guid` | `string` | Tasklist GUID |
|       `section_guid` | `string` | Section GUID |
|     `repeat_rule` | `string` | If the task is a recurring task, return the repeat_rule of the recurring task |
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
|     `url` | `string` | Task applink. |
|     `start` | `start` | Start time of the task |
|       `timestamp` | `string` | The timestamp of the start time/date in milliseconds from 1970-01-01 00:00:00. If the start time is a date, you need to convert the date to timestamp and set is_all_day=true |
|       `is_all_day` | `boolean` | Whether to start on a date. If set to true, only the date part of the timestamp will be parsed and stored. |
|     `subtask_count` | `int` | The number of subtasks of this task. |
|     `is_milestone` | `boolean` | Is it a milestone task |
|     `custom_fields` | `custom_field_value[]` | Custom field values for tasks |
|       `guid` | `string` | Field GUID |
|       `type` | `string` | Custom field types, support "member", "datetime", "number", "single_select", "multi_select" five types |
|       `number_value` | `string` | A custom field value of a number type, you can fill in the string representation of a legal number, or unset it by an empty string |
|       `datetime_value` | `string` | Datetime type custom field value. You can enter a timestamp in ms string representing the date, or unset it by an empty string. |
|       `member_value` | `member[]` | Custom field value for member type, you can set the one or multiple users (follow member format, only user type is supported). Only one value can be entered when the field setting is "multi=false". Or you can unset the value by an empty array. |
|         `id` | `string` | Indicates the id of member |
|         `type` | `string` | Type of member |
|         `role` | `string` | Member role |
|       `single_select_value` | `string` | single_select value, you can fill in the option_guid of the field options or unset it by an empty string. |
|       `multi_select_value` | `string[]` | multi_select type field value, you can fill in one or more option_guid of this field, or unset it by an empty array. |
|       `name` | `string` | Custom field name |
|     `dependencies` | `task_dependency[]` | Task dependence |
|       `type` | `string` | dependency type **Optional values are**:  - `prev`: prev-dependency - `next`: post-dependency  |
|       `task_guid` | `string` | GUIDs for dependent tasks |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned |
|   `has_more` | `boolean` | Whether the response body has more parameters | ### Response body example

{
    "code": 0,
    "msg": "Success",
    "data": {
        "items": [
            {
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
                    "role": "assignee"
                },
                "members": [
                    {
                        "id": "ou_2cefb2f014f8d0c6c2d2eb7bafb0e54f",
                        "type": "user",
                        "role": "assignee"
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
                            "role": "assignee"
                        },
                        "is_cover": false,
                        "uploaded_at": "1675742789470"
                    }
                ],
                "origin": {
                    "platform_i18n_name": {
                        "en_us": "workbench",
                        "zh_cn": "工作台",
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
                        "title": "Need assistance."
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
                            "en_us": "workbench",
                            "zh_cn": "工作台",
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
                            "en_us": "workbench",
                            "zh_cn": "工作台",
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
                            "en_us": "workbench",
                            "zh_cn": "工作台",
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
                        "name": "优先级"
                    }
                ],
                "dependencies": [
                    {
                        "type": "next",
                        "task_guid": "93b7bd05-35e6-4371-b3c9-6b7cbd7100c0"
                    }
                ]
            }
        ],
        "page_token": "aWQ9NzEwMjMzMjMxMDE=",
        "has_more": true
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1470400 | Incorrect request parameters, such as page_size fill in negative numbers. | The reason for the error is shown in the returned msg. |
| 500 | 1470500 | Server error. | Try retrying the call. If it continues to fail, please contact support for feedback. |
