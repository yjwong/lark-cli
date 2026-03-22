---
title: "Add reminders to task"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/task/add_reminders"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/task/v2/tasks/:task_guid/add_reminders"
service: "task-v2"
resource: "task"
section: "Tasks v2"
rate_limit: "100 per minute"
scopes:
  - "task:task:write"
updated: "1699521594000"
---

# Add Task Reminder

Add a reminder for a task. A reminder is a moment calculated based on the task due time. In order to set a reminder, the task must first have a due. You can set the due when Create Task, or by Update Task set a deadline.

At present, a task can only set 1 reminder. But the form of the interface can be expanded to support multiple reminders for a task in the future.

If the current task already has a reminder, to update the settings of the reminder, you need to call the Remove Task Reminder api to remove the original reminder. Then call this interface to add reminders.

> Requires edit permission for the task. Please refer to the "How are Tasks Authorized?" section in Task Features Overview.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/task/v2/tasks/:task_guid/add_reminders |
| HTTP Method | POST |
| Rate Limit | 100 per minute |
| Supported app types | custom,isv |
| Required scopes | `task:task:write` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `task_guid` | `string` | GUID of task to add reminder **Example value**: "d300a75f-c56a-4be9-80d1-e47653028ceb" **Data validation rules**: - Maximum length: `100` characters | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Default value**: `open_id` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `reminders` | `reminder[]` | Yes | Task reminder list. Currently only 1 reminder is supported per task. **Example value**: [ou_2cefb2f014f8d0c6c2d2eb7bafb0e54f] **Data validation rules**: - Length range: `1` ～ `1` |
|   `relative_fire_minute` | `int` | Yes | The number of minutes relative to the deadline. For example, 30 means 30 minutes before the due; 0 means the due time. Negative number are not supported. **Example value**: 30 | ### Request body example

{
    "reminders": [
        {
            "relative_fire_minute": 30
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
|   `task` | `task` | Update the completed task entity |
|     `guid` | `string` | Task GUID |
|     `summary` | `string` | Task summary |
|     `description` | `string` | Task description |
|     `due` | `due` | Task due |
|       `timestamp` | `string` | The timestamp of the due time/date in milliseconds from 1970-01-01 00:00:00. If the expiration time is a date, you need to convert the date to timestamp and set is_all_day = true |
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
|       `type` | `string` | Member role |
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
|         `type` | `string` | member type |
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
|     `tasklists` | `task_in_tasklist_info[]` | The name of the tasklist to which the task belongs. The caller can only see tasklists that have permission to read. |
|       `tasklist_guid` | `string` | The guid of the task list |
|       `section_guid` | `string` | Section GUID |
|     `repeat_rule` | `string` | Repeat task configuration. |
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
|     `subtask_count` | `int` | The number of subtasks of this task. | ### Response body example

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
                        "role": "editor"
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
            "subtask_count": 1
        }
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 404 | 1470404 | Task data does not exist or has been deleted. | Confirm that the task data you want to access exists or has been deleted. |
| 400 | 1470400 | Request parameter error, such as setting relative_fire_minute to a negative number. | The reason for the error is shown in the returned msg prompt. |
| 500 | 1470500 | Server error. | Try retrying the call. If it continues to fail, please contact support for feedback. |
| 400 | 1470613 | Task has a reminder. | You need to call the Remove Task Reminder to remove the original reminder. Then call this api to add reminders. |
| 403 | 1470403 | Missing task edit permission. | Check whether the calling identity has task edit permission. Please refer to the "How are Tasks Authorized?" section in Task Features Overview. |
