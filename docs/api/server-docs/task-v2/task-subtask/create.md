---
title: "Create Subtask"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/task-subtask/create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/task/v2/tasks/:task_guid/subtasks"
service: "task-v2"
resource: "task-subtask"
section: "Tasks v2"
rate_limit: "10 per second"
scopes:
  - "task:task:write"
updated: "1699521610000"
---

# Create Subtask

Create a subtask for a task.

The api function is exactly the same as the Create Task except that the parent task GUID needs to be provided.

> Creating a subtask requires editing permissions to the parent task. Please refer to the "How are Tasks Authorized?" section in Task Features Overview.
> 
> If a new task is added to the tasklist, edit permissions to the tasklist are required. Please refer to the "How are tasklist authorized?" section in Tasklist Features Overview .

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/task/v2/tasks/:task_guid/subtasks |
| HTTP Method | POST |
| Rate Limit | 10 per second |
| Supported app types | custom,isv |
| Required scopes | `task:task:write` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `task_guid` | `string` | Parent task GUID **Example value**: "e297ddff-06ca-4166-b917-4ce57cd3a7a0" **Data validation rules**: - Maximum length: `100` characters | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Default value**: `open_id` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `summary` | `string` | Yes | Task summary **Example value**: "One review for full-year sales" **Data validation rules**: - Maximum length: `10000` characters |
| `description` | `string` | No | Task description **Example value**: "You need to read the review summary document in advance." |
| `due` | `due` | No | Task due **Example value**: 1675742789470 |
|   `timestamp` | `string` | No | The timestamp of the due time/date in milliseconds from 1970-01-01 00:00:00. If the expiration time is a date, you need to convert the date to timestamp and set is_all_day = true **Example value**: "1675454764000" |
|   `is_all_day` | `boolean` | No | Whether to due on a date. If set to true, only the date part of the timestamp will be parsed and stored. **Example value**: true |
| `origin` | `origin` | No | Task-associated third-party platform source information. See How to use Origin? |
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
|   `href` | `href` | No | Source of task association platform details page link |
|     `url` | `string` | No | The url corresponding to the task. **Example value**: "https://www.example.com" **Data validation rules**: - Length range: `0` ～ `1024` characters |
|     `title` | `string` | No | The title corresponding to the task **Example value**: "反馈一个问题，需要协助排查" **Data validation rules**: - Maximum length: `512` characters |
| `extra` | `string` | No | Any data that the caller can pass in attached to the task. It will be returned as is when the task details are obtained. **Example value**: "dGVzdA==" **Data validation rules**: - Maximum length: `65536` characters |
| `completed_at` | `string` | No | The completion time of the task timestamp (ms) **Example value**: "1675742789470" **Default value**: `0` **Data validation rules**: - Maximum length: `20` characters |
| `members` | `member[]` | No | Task member list **Example value**: [ou_2cefb2f014f8d0c6c2d2eb7bafb0e54f] **Data validation rules**: - Maximum length: `500` |
|   `id` | `string` | No | Indicates the id of member **Example value**: "ou_2cefb2f014f8d0c6c2d2eb7bafb0e54f" **Data validation rules**: - Maximum length: `100` characters |
|   `type` | `string` | No | Type of member **Example value**: "user" **Default value**: `user` |
|   `role` | `string` | No | Member role, supporting "assignee" or "follower" **Example value**: "editor" **Data validation rules**: - Maximum length: `20` characters |
| `repeat_rule` | `string` | No | Task repeat_rule. Please refer to the "How to use recurring task?" section in Feature Overview . **Example value**: "FREQ=WEEKLY;INTERVAL=1;BYDAY=MO,TU,WE,TH,FR" **Data validation rules**: - Maximum length: `1000` characters |
| `custom_complete` | `custom_complete` | No | Task custom complete configuration. Please refer to the "How to use custom completion?" section in Feature Overview . |
|   `pc` | `custom_complete_item` | No | pc client side custom configuration (including mac and windows) |
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
|   `ios` | `custom_complete_item` | No | Custom completion configuration of ios client |
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
|   `android` | `custom_complete_item` | No | Custom completion configuration of android client |
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
| `tasklists` | `task_in_tasklist_info[]` | No | Information about the tasklist where the task is added. If set, it means that the created task will be added directly to the specified tasklist. |
|   `tasklist_guid` | `string` | No | The GUID of the tasklist to which the task is to be added. **Example value**: "cc371766-6584-cf50-a222-c22cd9055004" **Data validation rules**: - Maximum length: `100` characters |
|   `section_guid` | `string` | No | The section GUID of the tasklist where the task is added. If the tasklist GUID is set but section GUID is unset, the default section of the tasklist is automatically used. **Example value**: "e6e37dcc-f75a-5936-f589-12fb4b5c80c2" |
| `client_token` | `string` | No | Idempotent token. If provided triggers the idempotent behavior.Please refer to the "Idempotent Invocation" section in Feature Overview . **Example value**: "daa2237f-8310-4707-a83b-52c8a81e0fb7" **Data validation rules**: - Length range: `10` ～ `100` characters |
| `start` | `start` | No | Task start time (ms) |
|   `timestamp` | `string` | No | The timestamp of the start time/date in milliseconds from 1970-01-01 00:00:00. If the start time is a date, you need to convert the date to timestamp and set is_all_day = true **Example value**: "1675454764000" |
|   `is_all_day` | `boolean` | No | Whether to start on a date. If set to true, only the date part of the timestamp will be parsed and stored. **Example value**: true |
| `reminders` | `reminder[]` | No | Task reminder |
|   `relative_fire_minute` | `int` | Yes | The number of minutes of reminder time relative to the deadline. For example, 30 means 30 minutes before the deadline; 0 means the deadline reminder. **Example value**: 30 | ### Request body example

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
            "title": "反馈一个问题，需要协助排查"
        }
    },
    "extra": "dGVzdA==",
    "completed_at": "1675742789470",
    "members": [
        {
            "id": "ou_2cefb2f014f8d0c6c2d2eb7bafb0e54f",
            "type": "user",
            "role": "editor"
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
    ]
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `subtask` | `task` | Created task |
|     `guid` | `string` | Task GUID |
|     `summary` | `string` | Task title |
|     `description` | `string` | task description |
|     `due` | `due` | Task dye |
|       `timestamp` | `string` | The timestamp of the expiration time/date in milliseconds from 1970-01-01 00:00:00. If the expiration time is a date, you need to convert the date to timestamp and set is_all_day = true |
|       `is_all_day` | `boolean` | Whether to expire on a date. If set to true, only the date part of the timestamp will be parsed and stored. |
|     `reminders` | `reminder[]` | List of reminder configurations for tasks. Currently there is a maximum of 1 per task. |
|       `id` | `string` | ID of reminder time setting |
|       `relative_fire_minute` | `int` | The number of minutes of reminder time relative to the deadline. For example, 30 means 30 minutes before the deadline; 0 means the deadline reminder. |
|     `creator` | `member` | Task creator |
|       `id` | `string` | Indicates the id of member |
|       `type` | `string` | Type of member |
|       `role` | `string` | Member role |
|     `members` | `member[]` | Task member list |
|       `id` | `string` | Indicates the id of member |
|       `type` | `string` | Type of member |
|       `role` | `string` | Member role |
|     `completed_at` | `string` | Task completion timestamp (ms) |
|     `attachments` | `attachment[]` | Task attachments |
|       `guid` | `string` | Attachment GUID |
|       `file_token` | `string` | Attachment token in drive system |
|       `name` | `string` | Attachment name |
|       `size` | `int` | Byte size of attachment |
|       `resource` | `resource` | Resources attributed to attachments |
|         `type` | `string` | resource type |
|         `id` | `string` | Resource ID |
|       `uploader` | `member` | Attachment uploader |
|         `id` | `string` | Indicates the id of member |
|         `type` | `string` | Type of member |
|         `role` | `string` | Member role |
|       `is_cover` | `boolean` | Is it a cover image |
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
|     `tasklists` | `task_in_tasklist_info[]` | The name of the tasklist to which the task belongs. The caller can only see the list of tasklist that have permission to access. |
|       `tasklist_guid` | `string` | The guid of the task list |
|       `section_guid` | `string` | Section guid for task list |
|     `repeat_rule` | `string` | If the task is a recurring task, return the configuration of the recurring task |
|     `parent_task_guid` | `string` | If the current task is a subtask of a task, return the guid of the parent task |
|     `mode` | `int` | The mode of the task. 1 - countersign tasks; 2 - or sign tasks |
|     `source` | `int` | Source of task creation **Optional values are**:  - `0`: Unknown source - `1`: Mission Center - `2`: Chat Task/Message Transfer Task - `6`: Tasks created with tenant_access_token authorization through the open platform - `7`: Tasks created with user_access_token authorization through the open platform - `8`: Document task  |
|     `custom_complete` | `custom_complete` | Custom completion configuration for tasks |
|       `pc` | `custom_complete_item` | Lark pc client side custom configuration (including mac and windows) |
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
|       `ios` | `custom_complete_item` | Custom completion configuration of Lark ios |
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
|       `android` | `custom_complete_item` | Custom completion configuration of Lark android |
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
|     `start` | `start` | The start time of the task. If you set both the start time and the deadline for the task, the start time must be 
{
    "code": 0,
    "msg": "success",
    "data": {
        "subtask": {
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
                "role": "editor"
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
            "parent_task_guid": "E297ddff-06ca-4166-b917-4ce57cd3a7a0",
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
            "subtask_count": 1
        }
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1470400 | Incorrect request parameters. For example, no task summary was left empty. | The reason for the error is shown in the returned ms. |
| 404 | 1470404 | Parent task data does not exist or has been deleted. | Confirm that the parent task data exists or has been deleted. |
| 500 | 1470500 | Server error. | Try retrying the call. If it continues to fail, please contact support for feedback. |
| 403 | 1470403 | Missing edit permissions for the parent task. | Check if the calling identity has edit permissions to the parent task. Please refer to the "How are Tasks Authorized?" section in Task Features Overview. |
| 500 | 1470422 | Invoke api concurrently using the same client_token. | Do not invoke api concurrently with same client_token. |
