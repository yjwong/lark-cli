---
title: "Batch Get OKR"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/okr-v1/okr/batch_get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/okr/v1/okrs/batch_get"
service: "okr-v1"
resource: "okr"
section: "OKR"
scopes:
  - "okr:okr:readonly"
  - "okr:okr"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1670425881000"
---

# Get OKR in bulk

Get OKR in batches based on OKR id

> ` tenant_access_token ` requires additional application permissions  to access OKR information as an application 

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/okr/v1/okrs/batch_get |
| HTTP Method | GET |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `okr:okr:readonly` `okr:okr` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: "open_id" **Optional values are**:  - `open_id`: user's open id - `union_id`: user's union id - `user_id`: user's user id - `people_admin_id`: Identify users by people_admin_id  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` |
| `okr_ids` | `string[]` | Yes | OKR ID list, up to 10 **Example value**: 7043693679567028244 **Data validation rules**: - Maximum length: `10` |
| `lang` | `string` | No | Request the language version of OKR (such as @'s name), lang = en_us/zh_cn, request Query **Example value**: "zh_cn" **Default value**: `zh_cn` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `okr_list` | `okr_batch[]` | OKR List |
|     `id` | `string` | Id |
|     `permission` | `int` | Access to OKR **Optional values are**:  - `0`: No permission - `1`: Has permission  |
|     `period_id` | `string` | period_id |
|     `name` | `string` | Name |
|     `objective_list` | `okr_objective[]` | Objective list |
|       `id` | `string` | Objective ID |
|       `permission` | `int` | Permissions **Optional values are**:  - `0`: No permission - `1`: Has permission  |
|       `content` | `string` | Objective content |
|       `progress_report` | `string` | Objective progress record content |
|       `score` | `int` | Objective score (0 - 100) |
|       `weight` | `float` | Objective weight (0 - 100) |
|       `progress_rate` | `okr_objective_progress_rate` | Objective progress |
|         `percent` | `int` | Objective progress percent > = 0 |
|         `status` | `string` | Objective progress status - "-1" not yet - "0" normal - "1" risk - "2" postpone |
|       `kr_list` | `okr_objective_kr[]` | Objective KeyResult List |
|         `id` | `string` | Key Result ID |
|         `content` | `string` | KeyResult |
|         `score` | `int` | KeyResult (0 - 100) |
|         `weight` | `int` | KeyResult weight (0 - 100) (abandoned) |
|         `kr_weight` | `float` | KeyResult weight (0 - 100) |
|         `progress_rate` | `okr_objective_progress_rate` | KR Progress |
|           `percent` | `int` | Objective progress percent > = 0 |
|           `status` | `string` | Objective progress status - "-1" not yet - "0" normal - "1" risk - "2" postpone |
|         `progress_record_list` | `progress_record_simplify[]` | Progress list for this KeyResult |
|           `id` | `string` | OKR progress record ID |
|         `progress_rate_percent_last_updated_time` | `string` | Last progress percentage update time, milliseconds |
|         `progress_rate_status_last_updated_time` | `string` | Last status update in milliseconds |
|         `progress_record_last_updated_time` | `string` | The last time you added or edited progress in the sidebar, in milliseconds |
|         `progress_report_last_updated_time` | `string` | Time of last edit of progress/notes, milliseconds |
|         `score_last_updated_time` | `string` | Last score update time, milliseconds |
|         `deadline` | `string` | Deadline, milliseconds |
|         `mentioned_user_list` | `okr_objective_aligned_objective_owner[]` | List of persons referred to in the Objective |
|           `open_id` | `string` | User open_id |
|           `user_id` | `string` | User user_id |
|       `aligned_objective_list` | `okr_objective_aligned_objective[]` | List of Objectives aligned to the Objective |
|         `id` | `string` | Objective ID |
|         `okr_id` | `string` | OKR ID |
|         `owner` | `okr_objective_aligned_objective_owner` | Owner of the Objective |
|           `open_id` | `string` | User open_id |
|           `user_id` | `string` | User user_id |
|       `aligning_objective_list` | `okr_objective_aligned_objective[]` | List of Objectives to which the Objective is aligned |
|         `id` | `string` | Objective ID |
|         `okr_id` | `string` | OKR ID |
|         `owner` | `okr_objective_aligned_objective_owner` | Owner of the Objective |
|           `open_id` | `string` | User open_id |
|           `user_id` | `string` | User user_id |
|       `progress_record_list` | `progress_record_simplify[]` | Progress list for this Objective |
|         `id` | `string` | OKR progress record ID |
|       `progress_rate_percent_last_updated_time` | `string` | Last progress percentage update time, milliseconds |
|       `progress_rate_status_last_updated_time` | `string` | Last status update in milliseconds |
|       `progress_record_last_updated_time` | `string` | The last time you added or edited progress in the sidebar, in milliseconds |
|       `progress_report_last_updated_time` | `string` | Time of last edit of progress/notes, milliseconds |
|       `score_last_updated_time` | `string` | Last score update time, milliseconds |
|       `deadline` | `string` | Deadline, milliseconds |
|       `mentioned_user_list` | `okr_objective_aligned_objective_owner[]` | List of persons referred to in the Objective |
|         `open_id` | `string` | User open_id |
|         `user_id` | `string` | User user_id |
|     `confirm_status` | `int` | OKR Confirmation Status **Optional values are**:  - `0`: Initial state - `1`: Pending/uncommitted - `2`: Pending / pending confirmation - `3`: Rejected/suggested - `4`: Passed/confirmed  | ### Response body example

{
    "code": 0,
    "data": {
        "okr_list": [
            {
                "confirm_status": 4,
                "id": "7072252816005349396",
                "name": "2022 年 3 月",
                "objective_list": [
                    {
                        "aligned_objective_list": [],
                        "aligning_objective_list": [],
                        "content": "需求@刘三",
                        "deadline": "1648656000000",
                        "id": "7073360513731690515",
                        "kr_list": [
                            {
                                "content": "1111@张三9",
                                "deadline": "1648656000000",
                                "id": "7073360471990140948",
                                "kr_weight": 50,
                                "mentioned_user_list": [
                                    {
                                        "open_id": "ou_a79faffdb6aee3618f0da4d42b192466",
                                        "user_id": "ou_a79faffdb6aee3618f0da4d42b192466"
                                    }
                                ],
                                "progress_rate": {
                                    "percent": 60,
                                    "status": "1"
                                },
                                "progress_rate_percent_last_updated_time": "1646907176099",
                                "progress_rate_status_last_updated_time": "1646907176099",
                                "progress_record_last_updated_time": "1646907586253",
                                "progress_record_list": [
                                    {
                                        "id": "7073411057431199764"
                                    },
                                    {
                                        "id": "7073410950174392340"
                                    },
                                    {
                                        "id": "7073360480580010004"
                                    },
                                    {
                                        "id": "7073360513731805203"
                                    }
                                ],
                                "progress_report_last_updated_time": "0",
                                "score": 100,
                                "score_last_updated_time": "1646907586244",
                                "weight": 50
                            }
                        ],
                        "mentioned_user_list": [
                            {
                                "open_id": "ou_ab08720df94e64045cc8c2b7694ef2a0",
                                "user_id": "ou_ab08720df94e64045cc8c2b7694ef2a0"
                            }
                        ],
                        "permission": 1,
                        "progress_rate": {
                            "percent": 30,
                            "status": "0"
                        },
                        "progress_rate_percent_last_updated_time": "1646907261326",
                        "progress_rate_status_last_updated_time": "1646907261326",
                        "progress_record_last_updated_time": "1646907590448",
                        "progress_record_list": [
                            {
                                "id": "7073410950447120403"
                            },
                            {
                                "id": "7073410950174474260"
                            },
                            {
                                "id": "7073360502990094355"
                            },
                            {
                                "id": "7073360502990061587"
                            }
                        ],
                        "progress_report": "红豆泥",
                        "progress_report_last_updated_time": "1646907387911",
                        "score": 100,
                        "score_last_updated_time": "1646907590472",
                        "weight": 40
                    }
                ],
                "period_id": "7067724095781142548",
                "permission": 1
            }
        ]
    },
    "msg": "success"
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 500 | 1009999 | internal server error | Internal error |
| 500 | 1009998 | system exception | System anomaly |
| 400 | 1001001 | invalid parameters | Invalid parameter |
| 400 | 1001002 | no permission | No permission |
| 400 | 1001003 | user not found | User does not exist |
| 400 | 1001004 | okr data not found | Okr data does not exist |
