---
title: "Patch system status"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/personal_settings-v1/system_status/patch"
method: "PATCH"
api_path: "https://open.larksuite.com/open-apis/personal_settings/v1/system_statuses/:system_status_id"
service: "personal_settings-v1"
resource: "system_status"
section: "Personal Settings"
scopes:
  - "personal_settings:status:system_status_update"
updated: "1672726533000"
---

# Patch system status

Patch a system state with a tenant-wide dimension.

> Notice：
> - The data to be operated is tenant dimension data, please operate with care.
> - After modifying the system state, it does not affect the user who is using it. After the user's system state availability time expires, the updated content will not be synchronized until it is enabled again.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/personal_settings/v1/system_statuses/:system_status_id |
| HTTP Method | PATCH |
| Supported app types | custom |
| Required scopes | `personal_settings:status:system_status_update` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `system_status_id` | `string` | System stauts ID Get system status ID **Example value**: "7101214603622940633" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `system_status` | `system_status` | Yes | System status |
|   `title` | `string` | Yes | System status title, the number of characters in the title should be in the range of 1 to 20. Titles of different system statuses cannot be repeated. **Notice:** - one Chinese character=two English character=two other language character **Example value**: "出差" |
|   `i18n_title` | `system_status_i18n_name` | No | System status i18n title, the number of characters in the title should be in the range of 1 to 20. Any title in i18n_title cannot be repeated between different system statuses. **Notice** - one Chinese character=two English character=two other language character. |
|     `zh_cn` | `string` | No | Chinese name **Example value**: "出差" |
|     `en_us` | `string` | No | English name **Example value**: "On business trip" |
|     `ja_jp` | `string` | No | Japanese name **Example value**: "出張中" |
|   `icon_key` | `string` | Yes | Icon **Icon optional value** **Example value**: "GeneralBusinessTrip" **Optional values are**:  - `GeneralDoNotDisturb`: GeneralDoNotDisturb - `GeneralInMeetingBusy`: GeneralInMeetingBusy - `Coffee`: Coffee - `GeneralBusinessTrip`: GeneralBusinessTrip - `GeneralWorkFromHome`: GeneralWorkFromHome - `StatusEnjoyLife`: StatusEnjoyLife - `GeneralTravellingCar`: GeneralTravellingCar - `StatusBus`: StatusBus - `StatusInFlight`: StatusInFlight - `Typing`: Typing - `EatingFood`: EatingFood - `SICK`: SICK - `GeneralSun`: GeneralSun - `GeneralMoonRest`: GeneralMoonRest - `StatusReading`: StatusReading - `Status_PrivateMessage`: Status_PrivateMessage - `StatusFlashOfInspiration`: StatusFlashOfInspiration - `GeneralVacation`: GeneralVacation  |
|   `color` | `string` | No | Color **Example value**: "BLUE" **Optional values are**:  - `BLUE`: BLUE - `GRAY`: GRAY - `INDIGO`: INDIGO - `WATHET`: WATHET - `GREEN`: GREEN - `TURQUOISE`: TURQUOISE - `YELLOW`: YELLOW - `LIME`: LIME - `RED`: RED - `ORANGE`: ORANGE - `PURPLE`: PURPLE - `VIOLET`: VIOLET - `CARMINE`: CARMINE  **Default value**: `BLUE` |
|   `priority` | `int` | No | The smaller the value, the higher the priority of the display. Different system status cannot have the same priority. **Example value**: 1 **Default value**: `0` **Data validation rules**: - Value range: `0` ～ `9` |
|   `sync_setting` | `system_status_sync_setting` | No | Sync setting |
|     `is_open_by_default` | `boolean` | No | Is open by default **Example value**: true **Default value**: `true` |
|     `title` | `string` | No | System status sync title, the number of characters in the title should be in the range of 1 to 30. **Notice:** - one Chinese character=two English character=two other language character. **Example value**: "出差期间自动开启" **Default value**: `自动开启` |
|     `i18n_title` | `system_status_sync_i18n_name` | No | System status sync i18n title, the number of characters in the title should be in the range of 1 to 30. **Notice:** - one Chinese character=two English character=two other language character. |
|       `zh_cn` | `string` | No | Chinese name **Example value**: "出差期间自动开启" |
|       `en_us` | `string` | No | English name **Example value**: "Auto display Business Trip" |
|       `ja_jp` | `string` | No | Japanese name **Example value**: "出張中に自動的にオンにする" |
|     `explain` | `string` | No | System status sync explain, the number of characters in the explain should be in the range of 1 to 60. **Notice:** - one Chinese character=two English character=two other language character. **Example value**: "出差审批通过后，将自动开启并优先展示该状态。" **Default value**: `从相关系统进行信息同步，同步后将自动开启并优先展示该状态。` |
|     `i18n_explain` | `system_status_sync_i18n_explain` | No | System status sync i18n explain, the number of characters in the explain should be in the range of 1 to 60. **Notice:** - one Chinese character=two English character=two other language character. |
|       `zh_cn` | `string` | No | Chinese name **Example value**: "出差审批通过后，该状态将自动开启并优先展示" |
|       `en_us` | `string` | No | English name **Example value**: "Auto-display after travel request is approved." |
|       `ja_jp` | `string` | No | Japanese name **Example value**: "申請が承認されると、このステータスが優先的に表示されます" |
| `update_fields` | `string[]` | Yes | Fields that need to be updated **Example value**: ['TITLE'] **Optional values are**:  - `TITLE`: System status's name - `I18N_TITLE`: System status's i18n name - `ICON`: Icon - `COLOR`: Color - `PRIORITY`: Priority - `SYNC_SETTING`: Sync setting  **Data validation rules**: - Minimum length: `1` | ### Request body example

{
    "system_status": {
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
    },
    "update_fields": [
        "Icon"
    ]
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `system_status` | `system_status` | System status |
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
|         `ja_jp` | `string` | Japanese name | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "system_status": {
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
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 2005001 | Your request contains an invalid request parameter. | The parameter is incorrect. Please check the input parameter according to the error message returned by the interface and refer to the documentation. |
| 400 | 2005002 | The same name or i18n name has already be created within your tenant. | The system state corresponding to this name or internationalized name has already been created, please use another name. |
| 400 | 2005003 | Name or i18n name contains sensitive words. | The name or internationalized name contains sensitive words, please use another name. |
| 400 | 2005005 | The priority of tenant system status has already be created within your tenant. | Use another priority, or adjust the priority of other system status. |
| 400 | 2005007 | Tenant does not have permission to api. | Tenant does not have permission to api, Please apply to use. |
