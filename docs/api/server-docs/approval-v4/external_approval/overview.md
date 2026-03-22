---
title: "Resource introduction"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/external_approval/overview"
service: "approval-v4"
resource: "external_approval"
section: "Approval"
updated: "1675167418000"
---

# Resource introduction
Approval definition is the description of an approval, including the approval title, icon, description, group, and other basic information.The three parties will create and synchronize three-party approval instances.

| Parameter | Type | Description |
| --- | --- | --- |
| `approval_name` | `string` | Approval definition name, which creates the value returned by the approval definition to indicate which process the instance belongs to. This field affects the title of the instance in the list, which is taken from the name field of the corresponding definition. **Example value**："@i18n@1" |
| `approval_code` | `string` | Approval definition code, which is user-defined and is the unique identifier for the definition. If the code does not exist, it will be created; otherwise, it will be updated. **Example value**："permission_test" |
| `group_code` | `string` | The approval group to which the approval definition belongs, which is user-defined.If the group_code does not exist, a new approval group will be created.If the group_code already exists, the approval group name will be updated with the group_name. **Example value**："work_group" |
| `group_name` | `string` | Group name. Its value is in the format of i18n key, and the text is in i18n_resource.If the group_code does not exist, the group_name is required. Otherwise, the group name will only be updated if it is entered.The group name of the approval definition in the approval initiation page is from this field. **Example value**："@i18n@2" |
| `description` | `string` | Description. Its value is in the format of i18n key, and the text is in i18n_resource.The description of the approval definition in the approval initiation page is from this field. **Example value**："@i18n@2" |
| `external` | `approval_create_external` | Third-party approval related |
|   `biz_name` | `string` | This field is used in the list to indicate where the approval comes from, whose value is in the format of i18n key. Please note that the prefix "From" is not required, since the prefix will be added by the Approval Center. **Example value**："@i18n@3" |
|   `biz_type` | `string` | This field is used to identify which business party the approval definition belongs to. It is user-defined , and the same value can be used for multiple approval definitions.The field is to attribute multiple approval definitions to the same business party. Therefore, when searching for approvals, if the business party has configured an external approval search with the same biz_type parameter, the parameters with biz_type will be filtered out from the search provided by the Approval Center, and the search will be performed from the configured URL. **Example value**："permission" |
|   `create_link_mobile` | `string` | Initiation link in the mobile app. If the link is set, the approval will be displayed in the approval initiation page on the mobile, and users will be redirected to the link to initiate an approval after clicking.If this field is left empty, the approval will not be displayed in themobile app. **Example value**："https://applink.larksuite.com/client/mini_program/open?appId=cli_9c90fc38e07a9101&path=pages/approval-form/index?id=9999" |
|   `create_link_pc` | `string` | Initiation link on PC. If the link is set, the approval will be displayed in the approval initiation page on PC, and users will be redirected to the link to initiate an approval after clicking.If this field is left empty, the approval will not be displayed on PC. **Example value**："https://applink.larksuite.com/client/mini_program/open?mode=appCenter&appId=cli_9c90fc38e07a9101&path=pc/pages/create-form/index?id=9999" |
|   `support_pc` | `boolean` | Whether the approval instance, approval tasks, and approval CC are displayed on PC. If the value is true, the instance information under this definition will be displayed on PC; otherwise, it will not be displayed. **Example value**：true |
|   `support_mobile` | `boolean` | Whether the approval instance, approval tasks, and approval CC are displayed on the mobile. If the value is true, the instance information under this definition will be displayed on the mobile; otherwise, it will not be displayed.support_pcandsupport_mobile cannot be both false; otherwise, the instance information will not be displayed. **Example value**：true |
|   `support_batch_read` | `boolean` | Whether marking as read in batches is supported **Example value**：true |
|   `enable_mark_readed` | `boolean` | Whether marking as read is supported（This field is invalid） **Example value**：true |
|   `enable_quick_operate` | `boolean` | Whether quick action is supported **Example value**：true |
|   `action_callback_url` | `string` | Third-party's action callback URL. This URL is called by the Approval Center to notify the third-party system when the task approver of the [Pending] list clicks "Approve" or "Reject".Refer to Quick Approval Callback **Example value**："http://www.larksuite.com/approval/openapi/instanceOperate" |
|   `action_callback_token` | `string` | Callback token, used for the business system to validate whether the request comes from Approval. Refer to Open Platform Documentation **Example value**："sdjkljkx9lsadf110" |
|   `action_callback_key` | `string` | Request parameter encryption key. If this parameter is configured, the request parameters will be encrypted and the business party must decrypt the request. For the encryption and decryption methods, refer to Associate external options **Example value**："gfdqedvsadfgfsd" |
|   `allow_batch_operate` | `boolean` | Whether batch action is supported **Example value**：false |
| `viewers` | `approval_create_viewers[]` | List of users who can view. Multiple viewers can be configured at the same time, and only users within the visible scope can see the approval in the approval initiation page. |
|   `viewer_type` | `string` | Viewer type **Example value**："USER" **Optional values are**：  - `TENANT`: Users in a tenant - `DEPARTMENT`: Specified departments - `USER`: Specified users - `NONE`: No one  |
|   `viewer_user_id` | `string` | When view_type is USER, fill in the user id according to user_id_type **Example value**："19a294c2" |
|   `viewer_department_id` | `string` | When view_type is DEPARTMENT, fill in the department id according to department_id_type **Example value**："od-ac9d697abfa990b715dcc33d58a62a9d" |
| `i18n_resources` | `i18n_resource[]` | Internationalized text |
|   `locale` | `string` | Language **Example value**："zh-CN" **Optional values are**：  - `zh-CN`: Chinese - `en-US`: English - `ja-JP`: Japanese  |
|   `texts` | `i18n_resource_text[]` | Key, value, and i18n key of copy starts with @i18n@.This field is mainly used for internationalized text. Word order users can pass texts in multiple languages at the same time, and the Approval Center will use the corresponding text according to the user's language environment. If the text of the user's language environment isn't passed, the text of default language will be used. |
|     `key` | `string` | Copywriting key **Example value**："@i18n@1" |
|     `value` | `string` | Copywriting **Example value**："people" |
|   `is_default` | `boolean` | Whether to use the default language. If the default language is used, all keys should be contained; if a non-default language is used but its keys do not exist, the default language will be used instead. **Example value**：true | ## Data example

```json
{
    "approval_name": "@i18n@1",
    "approval_code": "permission_test",
    "group_code": "work_group",
    "group_name": "@i18@2",
    "external": {
        "create_link_pc": "https://applink.larksuite.com/client/mini_program/open?mode=appCenter&appId=cli_9c90fc38e07a9101&path=pc%2Fpages%2Fcreate-form%2Findex%3Fid%3D9999",
        "create_link_mobile": "https://applink.larksuite.com/client/mini_program/open?appId=cli_9c90fc38e07a9101&path=pages%2Fapproval-form%2Findex%3Fid%3D9999",
        "support_pc": true,
        "support_mobile": true,
        "support_batch_read": false,
        "action_callback_url":"http://www.larksuite.com/approval/openapi/operate",
        "action_callback_token":"sdjkljkx9lsadf110",
        "action_callback_key":"gfdqedvsadfgfsd",
        "enable_mark_readed": false,
        "key": "",
        "token":""
    },
    "i18n_resources":[
     {
        "locale":"zh-CN",
        "is_default":true,
         "texts":[
             {
             "key": "@i18n@1",
             "value":"people"  
            },
            {
             "key": "@i18n@2",
             "value":"hr"  
            }
         ]
      }
    ],
    "viewers": [
        {
            "viewer_type": "TENANT"
        }
    ]
}

```
