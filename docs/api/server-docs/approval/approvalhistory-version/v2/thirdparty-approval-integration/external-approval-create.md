---
title: "External Approval Create"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uIDNyYjLyQjM24iM0IjN"
method: "POST"
api_path: "https://www.larksuite.com/approval/openapi/v3/external/approval/create"
section: "Approval"
scopes:
  - "approval:approval:readonly"
updated: "1658309962000"
---

# Create third-party approval definitions
> This API has been upgraded to enhance its security. We recommend that you migrate to the new version>>

Approval definition is the description of an approval, including the approval title, icon, description, and other basic information. After an approval definition is created, users can see the approval in the initiation page of the approval app. By clicking "Initiate", a user will be redirected to the configured system address for the third-party approval initiation to initiate an approval.

In addition, the approval definition also configures the callback URL for approval operations. When the approver performs [Approve] or [Reject] action in the pending approval list, the Approval Center will call the callback URL to notify the third-party system.

> Please note that the Approval Center is not responsible for approval transfer, but only for display, action, and message notification. Therefore, there is no information about the approval process when the approval definition is created.

## Request
| HTTP URL | https://www.larksuite.com/approval/openapi/v3/external/approval/create |
| --- | --- |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes | `approval:approval:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about how to obtain and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body
| `approval_name` `Required` | `string` | Approval definition name, which creates the value returned by the approval definition to indicate which process the instance belongs to. This field affects the title of the instance in the list, which is taken from the name field of the corresponding definition. | @i18n@1 |
| --- | --- | --- | --- |
| `approval_code` `Required` | `string` | Approval definition code, which is user-defined and is the unique identifier for the definition. If the code does not exist, it will be created; otherwise, it will be updated. | permission_test |
| `group_code` `Required` | `string` | The approval group to which the approval definition belongs, which is user-defined.  If the group_code does not exist, a new approval group will be created.  If the group_code already exists, the approval group name will be updated with the group_name. | work_group |
| `group_name` | `string` | Group name. Its value is in the format of i18n key, and the text is in i18n_resource. If the group_code does not exist, the group_name is required. Otherwise, the group name will only be updated if it is entered. The group name of the approval definition in the approval initiation page is from this field. | @i18n@2 |
| `description` | `string` | Description. Its value is in the format of i18n key, and the text is in i18n_resource. The description of the approval definition in the approval initiation page is from this field. | @i18n@2 |
| `external `Required` ` | `map` | Third-party approval related |  |
| ∟ `biz_name` | `string` | This field is used in the list to indicate where the approval comes from, whose value is in the format of i18n key. Please note that the prefix "From" is not required, since the prefix will be added by the Approval Center. | @i18n@3 |
| ∟ ` biz_type` | `string` | This field is used to identify which business party the approval definition belongs to. It is user-defined , and the same value can be used for multiple approval definitions. The field is to attribute multiple approval definitions to the same business party. Therefore, when searching for approvals, if the business party has configured an external approval search with the same biz_type parameter, the parameters with biz_type will be filtered out from the search provided by the Approval Center, and the search will be performed from the configured URL. | permission |
| ∟ `create_link_mobile `Required` ` | `string` | Initiation link in the mobile app. If the link is set, the approval will be displayed in the approval initiation page on the mobile, and users will be redirected to the link to initiate an approval after clicking.  If this field is left empty, the approval will not be displayed in themobile app. | https://applink.larksuite.com/client/mini_program/open?appId=cli_9c90fc38e07a9101&path=pages/approval-form/index?id=9999 |
| ∟ `     create_link_pc ` | `string` | Initiation link on PC. If the link is set, the approval will be displayed in the approval initiation page on PC, and users will be redirected to the link to initiate an approval after clicking. If this field is left empty, the approval will not be displayed on PC. | https://applink.larksuite.com/client/mini_program/open?mode=appCenter&appId=cli_9c90fc38e07a9101&path=pc/pages/create-form/index?id=9999 |
| ∟ `support_pc` | `bool` | Whether the approval instance, approval tasks, and approval CC are displayed on PC. If the value is  true, the instance information under this definition will be displayed on PC; otherwise, it will not be displayed. | true |
| ∟ ` support_mobile` | `bool` | Whether the approval instance, approval tasks, and approval CC are displayed on the mobile. If the value is  true, the instance information under this definition will be displayed on the mobile; otherwise, it will not be displayed. support_pcandsupport_mobile cannot be both false; otherwise, the instance information will not be displayed. | true |
| ∟ ` support_batch_read` | Whether marking as read in batches is supported | true |  |
| ∟ ` enable_mark_readed` | Whether marking as read is supported | true |  |
| ∟ ` enable_quick_operate` | Whether quick action is supported | true |  |
| ∟ ` allow_batch_operate` | Whether batch action is supported | true |  |
| ∟ ` action_callback_url` | Third-party's action callback URL. This URL is called by the Approval Center to notify the third-party system when the task approver of the [Pending] list clicks "Approve" or "Reject". | http://www.larksuite.com/approval/openapi/instanceOperate |  |
| ∟ ` action_callback_token` | Callback  token, used for the business system to validate whether the request comes from Approval. Refer to [Open Platform Documentation] for details. | sdjkljkx9lsadf110 |  |
| ∟ ` action_callback_key` | Request parameter encryption key. If this parameter is configured, the request parameters will be encrypted and the business party must decrypt the request. For the encryption and decryption methods, refer to Associate external options . | gfdqedvsadfgfsd |  |
| `viewers` | `list` | List of users who can view. Multiple viewers can be configured at the same time, and only users within the visible scope can see the approval in the approval initiation page. |  |
| ∟ `viewer_type` `Required` | `string` | Viewer type **Optional values are:** - `TENANT`: Users in a tenant - `DEPARTMENT`: Specified departments - `USER`: Specified users - `NONE`: No one | USER |
| ∟ `viewer_id` `Required` | `string` | Viewer ID. If view_type is TENANT and NONE, viewer_id can be empty.If view_type is DEPARTMENT, viewer_id is open_department_id.If view_type is USER, viewer_id is user_id. | 19a294c2 |
| `i18n_resources` `Required` | `list` | Internationalized text |  |
| ∟ `locale` `Required` | `string` | Language **Optional values are:** - `zh-CN`: Chinese - `en-US`: English - `ja-JP`: Japanese | zh-CN |
| ∟ `is_default` `Required` | `bool` | Whether to use the default language. If the default language is used, all keys should be contained; if a non-default language is used but its keys do not exist, the default language will be used instead. | true |
| ∟ `texts` `Required` | `map` | Text key, whose value is in the format of i18n key and starts with **@i18n@**. This field is mainly used for internationalized text. Word order users can pass texts in multiple languages at the same time, and the Approval Center will use the corresponding text according to the user's language environment. If the text of the user's language environment isn't passed, the text of default language will be used. | { "@i18n@1": "Permission scope application",      "@i18n@2": "OA approval", "@i18n@3": "Permission" } | ### Request body example

```json
{
    "approval_name": "@i18n@1",
    "approval_code": "permission_test",
    "group_code": "work_group",
    "group_name": "@i18n@2",
    "external": {
        "create_link_pc": "https://applink.larksuite.com/client/mini_program/open?mode=appCenter&appId=cli_9c90fc38e07a9101&path=pc%2Fpages%2Fcreate-form%2Findex%3Fid%3D9999",
        "create_link_mobile": "https://applink.larksuite.com/client/mini_program/open?appId=cli_9c90fc38e07a9101&path=pages%2Fapproval-form%2Findex%3Fid%3D9999",
        "support_pc": true,
        "support_mobile": true,
        "support_batch_read": false,
        "action_callback_url":"http://larksuite.com/approval/openapi/operate",
        "action_callback_token":"sdjkljkx9lsadf110",
        "action_callback_key":"gfdqedvsadfgfsd",
        "enable_mark_readed": false,
        "enable_quick_operate": false,
        "allow_batch_operate": false,
        "intranet_detect": false,
        "key": "",
        "token":""
    },
    "i18n_resources":[
     {
        "locale":"zh-CN",
        "is_default":true,
         "texts":{
            "@i18n@1":"people",
             "@i18n@2":"hr",
             "@i18n@3":"HR"
         }
      }
    ],
    "viewers": [
        {
            "viewer_type": "TENANT"
        }
    ]
}
```

## Response

### Response body
| `code` | `int` | Error code. A value other than 0 indicates failure. |
| --- | --- | --- |
| `msg` | `string` | Return code description |
| `data` | `map` | Returned business information |
| ∟ `approval_code` | `string` | Approval definition code, which is used to initiate an instance. | ### Response body example

```json
{
    "code": 0,
    "msg": "success",
    "data": {
        "approval_code": "C30381C8-7A5F-4717-A9CF-C233BF0202D4"
    }
}
```
