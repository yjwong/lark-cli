---
title: "Get behavior audit log"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uQjM5YjL0ITO24CNykjN/audit_log/audit_data_get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/admin/v1/audit_infos"
section: "security_and_compliance"
rate_limit: "100 per minute"
scopes:
  - "admin:audit_info:readonly"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1723186668000"
---

# Obtain behavioral audit log

This interface is used to query member operation behavior logs. The behavior audit log records the time, location, operation objects and other information of members' operations. By querying member behavior logs, administrators can discover whether members have violated regulations to protect corporate data and information security.
- Performance Note: When querying, please shorten the query time range appropriately and control the query frequency appropriately (avoid repeated invalid queries and other situations)

## Request
| HTTP URL | https://open.larksuite.com/open-apis/admin/v1/audit_infos |
| --- | --- |
| HTTP Method | GET |
| Rate Limit | 100 per minute |
| Supported app types | custom |
| Required scopes Enable any scope from the list | `admin:audit_info:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | 必填 | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Query parameters
| Parameter | Type | 必填 | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | 否 | User ID categories **Example value**: user_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `user_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` |
| `latest` | `int` | No | Log time range: end time. Format: Second-level timestamp. Default: at this moment. The difference between the start and end dates cannot exceed 30 days **Example value**: 1668700799 |
| `oldest` | `int` | No | Log time range: starting time. Format: Second-level timestamp. Default value: 30 days ago at this moment. The difference between the start and end dates cannot exceed 30 days **Example value**: 1668528000 |
| `event_name` | `string` | No | Behavioral audit event name, such ascreate_doc. For names of optional events, see List of enumerated values **Example value**: space_create_doc |
| `operator_type` | `string` | No | Filter Operator: Operator type. Used in conjunction with operator_value, this item is required when filling in operator_value **Example value**: user **Optional values are**:  - `user`: user - `bot`: [Currently not open] Use bot_id to identify users  |
| `operator_value` | `string` | No | Operator value **Example value**: 55ed16fe |
| `event_module` | `int` | No | Behavioral audit event module. For the event modules of behavioral audit, see List of enumerated values **Example value**: 1 |
| `page_token` | `string` | No | Paging token. If this parameter has no value during the initial request, it indicates to traverse from the beginning. When returned has_more is true, new page_token will be returned. Recalling this API and passing in this page_token will obtain the data of next page **Example value**：LC39/f1%2B/Sz9Uv39Gf39/ew/cd5WY0gfGYFdixOW9cVk4bC79ituO/gx0qpPn1bYf92nz/kI0nNJOG3wCwDJKoNU%2BtyaXbpI8pV/9UNDMZT0BNeyanFH17Wv711Qh9anR3l2GjQfc2fUqXtxg1YPp63XyhYY4iRMv54ySRG7r%2BI89iS3zAoPzFuuU1MUJKsf |
| `page_size` | `int` | No | Paging size **Example value**: 20 **Default value**: `20` **Data validation rules**: - Value range: `1` ～ `200` |
| `user_type` | `int` | No | User type **Example value**：1 **Optional values are**：  - `0`: anyone on the internet - `1`: users within the organization - `2`: users outside the organization  |
| `object_type` | `int` | No | Action object type, see List of enumerated values. When you fill in object_value, this field is required. **Example value**：1 |
| `object_value` | `string` | No | Object value. When you fill in object_type, this field is required. **Example value**：55ed16fe | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `has_more` | `boolean` | Whether more data exists |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned |
|   `items` | `audit_info[]` | Behavioral audit log array |
|     `event_id` | `string` | Event ID, which is not unique and can be used for aggregation |
|     `unique_id` | `string` | Unique ID of event, which can be used for deduplication. This field is used to identify user behaviors |
|     `event_name` | `string` | Event name. For field details, see List of enumerated values |
|     `department_ids` | `string[]` | List of department IDs |
|     `event_module` | `int` | Event module. For field details, see List of enumerated values |
|     `operator_type` | `int` | Operator type **Optional values are**：  - `1`: inner user - `12`: bot - `1001`: external user  |
|     `operator_value` | `string` | Operator value. It is a desensitization value when operator_type is 1001 |
|     `objects` | `audit_object_entity[]` | List of action objects |
|       `object_type` | `string` | Action object type. For field details, see List of enumerated values |
|       `object_value` | `string` | Action object value |
|       `object_name` | `string` | Action object name, which is currently open for Docs, chats, and apps. eg: chat name, doc title. |
|       `object_owner` | `string` | Action object owner, which is currently open for Docs |
|       `object_detail` | `audit_object_detail` | Action object extension field. For field details, see List of enumerated values |
|     `recipients` | `audit_recipient_entity[]` | list of receiver objects |
|       `recipient_type` | `string` | Receiver object type, 1 represents user, 2 represents department, 4 represents chat |
|       `recipient_value` | `string` | Receiver object value |
|       `recipient_detail` | `audit_recipient_detail` | recipient detail |
|     `event_time` | `int` | event time |
|     `ip` | `string` | ip |
|     `operator_app` | `string` | Thrid-party isvID |
|     `audit_context` | `audit_context` | environmental information |
|       `terminal_type` | `int` | terminal type **Optional values are**：  - `0`: ios - `1`: android - `2`: pc - `3`: web  |
|       `ios_context` | `audit_ios_context` | Return iOS context information. For field details, see List of enumerated values |
|       `pc_context` | `audit_pc_context` | Returns PC context information. For field details, see List of enumerated values |
|       `web_context` | `audit_web_context` | Returns web context information. For field details, see List of enumerated values |
|       `android_context` | `audit_android_context` | Returns Android context information. For field details, see List of enumerated values |
|     `extend` | `audit_event_extend` | Event extension, refer to the information in common_drawers |
|     `operator_app_name` | `string` | app name called from open platform |
|     `common_drawers` | `api_audit_common_drawers` | Event extension, see List of enumerated values |
|     `audit_detail` | `audit_detail` | Device detail city: ip , city; device_model: Device version; mc: Mac address; os: operating system; |
|     `operator_tenant` | `string` | Operator's organization id | ### 响应体示例

{
    "code": 0,
  	"msg": "success",
    "data": {
        "has_more": false,
        "items": [
            {
  				"event_id": "7254062411181719572",
				"unique_id": "7254062413199179796",
                "event_module": 1,
                "event_name": "space_edit_doc",
		  		"operator_type": 1,
                "operator_value": "4a3b8541",
                "operator_tenant": "F686619755",
                "event_time": 1688968015,
                "audit_context": {
                    "terminal_type": 3,
                    "web_context": {
                        "IP": "fdbd:dc02:ff:1:1:174:246:126",
                        "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36"
                    }
                },
                "audit_detail": {
                    "city": "",
                    "device_model": "",
                    "mc": "",
                    "os": ""
                },
                "common_drawers": {
                    "common_draw_info_list": [
                        {
                            "info_key": "CCM_op_status",
                            "info_val": "success",
                            "key_i18n_key": "SuiteAdmin_DrawerKeyspaceeditdoc_Kccmopstatus",
                            "val_i18n_key": "SuiteAdmin_DrawerValspaceeditdoc_Kccmopstatusvsuccess",
                            "val_type": "1"
                        },
                        {
                            "info_key": "ccm_edit_part",
                            "info_val": "0",
                            "key_i18n_key": "SuiteAdmin_DrawerKeyspaceeditdoc_Kccmeditpart",
                            "val_i18n_key": "SuiteAdmin_DrawerValspaceeditdoc_Kccmeditpartv0",
                            "val_type": "1"
                        },
                        {
                            "info_key": "ccm_folder_id",
                            "info_val": "nodbcMQRCwBJ5aWEWxQsWsTA5zg",
                            "key_i18n_key": "SuiteAdmin_DrawerKeyspaceeditdoc_Kccmfolderid",
                            "val_i18n_key": "SuiteAdmin_DrawerValspaceeditdoc_Kccmfolderidvnodbcmqrcwbj5awewxqswsta5zg",
                            "val_type": "1"
                        },
                        {
                            "info_key": "ccm_folder_name",
                            "info_val": "",
                            "key_i18n_key": "SuiteAdmin_DrawerKeyspaceeditdoc_Kccmfoldername",
                            "val_i18n_key": "SuiteAdmin_DrawerValspaceeditdoc_Kccmfoldernamev",
                            "val_type": "1"
                        },
                        {
                            "info_key": "ccm",
                            "info_val": "[{\"key\":\"title\",\"keyName\":\"文件名称\",\"value\":\"\",\"valueType\":\"1\",\"valueExtInfo\":null,\"eventFieldInfoList\":null},{\"key\":\"owner\",\"keyName\":\"文件所有者\",\"value\":\"\",\"valueType\":\"3\",\"valueExtInfo\":\"{\\\"type\\\":3,\\\"id\\\":\\\"0\\\",\\\"name\\\":\\\"\\\",\\\"avatar\\\":\\\"\\\",\\\"employee_id\\\":\\\"\\\"}\",\"eventFieldInfoList\":null},{\"key\":\"create_time\",\"keyName\":\"创建时间\",\"value\":\"\",\"valueType\":\"8\",\"valueExtInfo\":null,\"eventFieldInfoList\":null},{\"key\":\"sec_label\",\"keyName\":\"密级标签\",\"value\":\"\",\"valueType\":\"1\",\"valueExtInfo\":null,\"eventFieldInfoList\":null},{\"key\":\"doc_belong_tenant\",\"keyName\":\"是否属于本组织\",\"value\":\"否\",\"valueType\":\"1\",\"valueExtInfo\":null,\"eventFieldInfoList\":null}]",
                            "key_i18n_key": "文件 ID",
                            "val_type": "Lwd1smp3nl01AndDEMzbsfqacBb"
                        },
                        {
                            "info_key": "ccm_type",
                            "info_val": "-",
                            "key_i18n_key": "文件类型"
                        }
                    ]
                },
                "department_ids": [
                    "0",
                    "od-ab89476fbcf901c3e5ccd78f788f85bf"
                ],
                
                "extend": {},
                "ip": "fdbd:dc02:ff:1:1:174:246:126",
                "objects": [
                    {
                        "object_detail": {},
                        "object_name": "",
                        "object_owner": "",
                        "object_type": "106",
                        "object_value": "Lwd1smp3nl01AndDEMzbsfqacBb"
                    }
                ],
                "operator_app": "",
                "operator_app_name": "",
                "recipients": [
  					{
                        "recipient_type": "",
                        "recipient_value": "",
                        "recipient_detail": {}
                    }
  				]
            }
        ],
        "page_token": "LC39/f1%2B/Sz9Uv39Gf39/ew/cd5WY0gfGYFdixOW9cW8qPuyh59mg7cXBJKifnqatoP8P2g8URSUi2/4fj5NPJh8VxaPH3yTCnbSZeXzNs4mpQYaUVbp7lTIy1YfIsZ1i0UPASt4qAPObiCewWErhlEG6Qg%2B/6zX7JFBCRsatVcctGTHeL9bb8ssmfdt0Yag"
    }
}

Note: The audit context information may be lost due to different Lark versions or user devices. We will try our best to guarantee complete access to context information. You can also report missing information to Lark at any time

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1050001 | TIME_CHECK_NOT_VALID | Check the request parameters latest and oldest |
| 500 | 1050002 | ErrCode_DATABASE_ERR | System error, please try again or contact relevant personnel |
| 400 | 1050004 | Error_Param_Error | Check the request parameters |
| 400 | 1050005 | Error_Page_Size_Invalid | Check the request parameter page_size |
| 400 | 1050006 | Error_Page_Token_Invalid | Check the request parameter page_token |
| 400 | 1050007 | Error_Event_Name_Not_Found | Check the request parameter event_name |
| 500 | 1050008 | Error_Open_Platform_RPC | System error, please try again or contact relevant personnel |
| 400 | 1050009 | Error_Lark_ID_Not_Found | Check the request parameter operator_value |
