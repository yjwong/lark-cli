---
title: "List system status"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/personal_settings-v1/system_status/list"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/personal_settings/v1/system_statuses"
service: "personal_settings-v1"
resource: "system_status"
section: "Personal Settings"
scopes:
  - "personal_settings:status:system_status_update"
updated: "1672726533000"
---

# List system statuses

List system statuses with a tenant-wide dimension.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/personal_settings/v1/system_statuses |
| HTTP Method | GET |
| Supported app types | custom |
| Required scopes | `personal_settings:status:system_status_update` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `page_size` | `int` | No | Page size **Example value**: 50 **Default value**: `50` **Data validation rules**: - Maximum value: `50` |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: "GxmvlNRvP0NdQZpa7yIqf_Lv_QuBwTQ8tXkX7w-irAghVD_TvuYd1aoJ1LQph86O-XImC4X9j9FhUPhXQDvtrQ==" | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `items` | `system_status[]` | Tenant system status list |
|     `system_status_id` | `string` | System status id |
|     `title` | `string` | System status title, the number of characters in the title should be in the range of 1 to 20. Titles of different system statuses cannot be repeated. **Notice:** - one Chinese character=two English character=two other language character |
|     `i18n_title` | `system_status_i18n_name` | System status i18n title, the number of characters in the title should be in the range of 1 to 20. Any title in i18n_title cannot be repeated between different system statuses. **Notice** - one Chinese character=two English character=two other language character. |
|       `zh_cn` | `string` | Chinese name |
|       `en_us` | `string` | English name |
|       `ja_jp` | `string` | Japanese name |
|     `icon_key` | `string` | Icon **Icon optional value** **Optional values are**:  - `GeneralDoNotDisturb`: GeneralDoNotDisturb - `GeneralInMeetingBusy`: GeneralInMeetingBusy - `Coffee`: Coffee - `GeneralBusinessTrip`: GeneralBusinessTrip - `GeneralWorkFromHome`: GeneralWorkFromHome - `StatusEnjoyLife`: StatusEnjoyLife - `GeneralTravellingCar`: GeneralTravellingCar - `StatusBus`: StatusBus - `StatusInFlight`: StatusInFlight - `Typing`: Typing - `EatingFood`: EatingFood - `SICK`: SICK - `GeneralSun`: GeneralSun - `GeneralMoonRest`: GeneralMoonRest - `StatusReading`: StatusReading - `Status_PrivateMessage`: Status_PrivateMessage - `StatusFlashOfInspiration`: StatusFlashOfInspiration - `GeneralVacation`: GeneralVacation  |
|     `color` | `string` | Color **Optional values are**:  - `BLUE`: BLUE - `GRAY`: GRAY - `INDIGO`: INDIGO - `WATHET`: WATHET - `GREEN`: GREEN - `TURQUOISE`: TURQUOISE - `YELLOW`: YELLOW - `LIME`: LIME - `RED`: RED - `ORANGE`: ORANGE - `PURPLE`: PURPLE - `VIOLET`: VIOLET - `CARMINE`: CARMINE  |
|     `priority` | `int` | The smaller the value, the higher the priority of the display. Different system status cannot have the same priority. |
|     `sync_setting` | `system_status_sync_setting` | Sync setting |
|       `is_open_by_default` | `boolean` | Is open by default |
|       `title` | `string` | System status sync title, the number of characters in the title should be in the range of 1 to 30. **Notice:** - one Chinese character=two English character=two other language character. |
|       `i18n_title` | `system_status_sync_i18n_name` | System status sync i18n title, the number of characters in the title should be in the range of 1 to 30. **Notice:** - one Chinese character=two English character=two other language character. |
|         `zh_cn` | `string` | Chinese name |
|         `en_us` | `string` | English name |
|         `ja_jp` | `string` | Japanese name |
|       `explain` | `string` | System status sync explain, the number of characters in the explain should be in the range of 1 to 60. **Notice:** - one Chinese character=two English character=two other language character. |
|       `i18n_explain` | `system_status_sync_i18n_explain` | System status sync i18n explain, the number of characters in the explain should be in the range of 1 to 60. **Notice:** - one Chinese character=two English character=two other language character. |
|         `zh_cn` | `string` | Chinese name |
|         `en_us` | `string` | English name |
|         `ja_jp` | `string` | Japanese name |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned |
|   `has_more` | `boolean` | Whether the response body has more parameters | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "items": [
            {
                "system_status_id": "7101214603622940633",
                "title": "出差",
                "i18n_title": {
                    "zh_cn": "出差",
                    "en_us": "On business trip",
                    "ja_jp": "出張中"
                },
                "icon_key": "GeneralBusinessTrip",
                "color": "BLUE",
                "priority": 1,
                "sync_setting": {
                    "is_open_by_default": true,
                    "title": "出差期间自动开启",
                    "i18n_title": {
                        "zh_cn": "出差期间自动开启",
                        "en_us": "Auto display Business Trip",
                        "ja_jp": "出張中に自動的にオンにする"
                    },
                    "explain": "出差审批通过后，将自动开启并优先展示该状态。",
                    "i18n_explain": {
                        "zh_cn": "出差审批通过后，该状态将自动开启并优先展示",
                        "en_us": "Auto-display after travel request is approved.",
                        "ja_jp": "申請が承認されると、このステータスが優先的に表示されます"
                    }
                }
            }
        ],
        "page_token": "GxmvlNRvP0NdQZpa7yIqf_Lv_QuBwTQ8tXkX7w-irAghVD_TvuYd1aoJ1LQph86O-XImC4X9j9FhUPhXQDvtrQ==",
        "has_more": true
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 2005001 | Your request contains an invalid request parameter. | The parameter is incorrect. Please check the input parameter according to the error message returned by the interface and refer to the documentation. |
| 400 | 2005007 | Tenant does not have permission to api. | Tenant does not have permission to api, Please apply to use. |
