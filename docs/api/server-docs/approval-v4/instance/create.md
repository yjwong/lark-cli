---
title: "Create an approval instance"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/instance/create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/approval/v4/instances"
service: "approval-v4"
resource: "instance"
section: "Approval"
rate_limit: "100 per minute"
scopes:
  - "approval:approval"
  - "approval:instance"
updated: "1760339459000"
---

# Create an approval instance

To call this interface and create an approval instance using a specified approval definition code, the caller needs to have a detailed understanding of the approval definition form. According to the defined form structure, the form values should be passed in through this interface.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/approval/v4/instances |
| HTTP Method | POST |
| Rate Limit | 100 per minute |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `approval:approval` `approval:instance` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `approval_code` | `string` | Yes | Approval definition Code. Acquisition methods: - Call the Create approval definition API, and get the approval_code from the response parameters. - Log in to the approval management backend, and retrieve it from the URL of the specified approval definition. For specific operations, refer to What is Approval Code. **Example value**: "7C468A54-8745-2245-9675-08B7C63E7A85" |
| `user_id` | `string` | No | The initiator's user_id or open_id must be passed in, with priority given to user_id if both are provided. For the way to get it, refer to How to Obtain User ID. **Example value**: "f7cb567e" |
| `open_id` | `string` | No | The initiator's open_id or user_id must be passed in, with priority given to user_id if both are provided. For the way to get it, refer to How to Obtain Open ID. **Example value**: "ou_123456" |
| `department_id` | `string` | No | The initiator's department ID. If the user belongs to only one department, this field can be left empty. If the user belongs to multiple departments and no value is provided, the first department in the list will be selected by default. For the way to get it, refer to Department ID. **Notes**: - Filling in the root department is not supported. - The department ID should be of department_id type. **Example value**: "9293493ccacbdb9a" |
| `form` | `string` | Yes | Values of filled approval form controls must be provided as a JSON array, which needs to be compressed and escaped into a string when passed. Parameter descriptions for each control value can be found in Approval Instance Form Control Parameters. **Example value**: "[{\"id\":\"111\", \"type\": \"input\", \"value\":\"test\"}]" |
| `node_approver_user_id_list` | `node_approver[]` | No | If the approval process definition includes a node where the initiator needs to select an approver, this parameter should be used to specify the approver for the corresponding node (by designating the approver using the user_id). **Notes**: If both node_approver_user_id_list and node_approver_open_id_list are provided, the approvers will be the union of both lists. |
|   `key` | `string` | No | The node_id or custom_node_id of the node can be obtained by calling the View Specified Approval Definition API and retrieving it from the node_list parameter in the interface's response. **Example value**: "46e6d96cfa756980907209209ec03b64" |
|   `value` | `string[]` | No | The approver list requires the user_id of users. Refer to How to Obtain User ID for the method to acquire user_ids. **Example value**: ["f7cb567e"] |
| `node_approver_open_id_list` | `node_approver[]` | No | If the approval process definition includes a node where the initiator needs to select an approver, this parameter should be used to specify the approver for the corresponding node (by designating the approver using the open_id). **Notes**: If both node_approver_user_id_list and node_approver_open_id_list are provided, the approvers will be the union of both lists. |
|   `key` | `string` | No | The node_id or custom_node_id of the node can be obtained by calling the View Specified Approval Definition API and retrieving it from the node_list parameter in the interface's response. **Example value**: "46e6d96cfa756980907209209ec03b64" |
|   `value` | `string[]` | No | The approver list requires the user open_id. Refer to How to Obtain User's Open ID for the method to acquire open_ids. **Example value**: ["f7cb567e"] |
| `node_cc_user_id_list` | `node_cc[]` | No | If the approval process definition includes a node where the initiator needs to select the CC recipient, this parameter should be used to specify the CC recipient for the corresponding node (by designating the recipient using the user_id). **Notes**: If both node_cc_user_id_list and node_cc_open_id_list are provided, the CC recipients will be the union of both lists. **Data validation rules**: - Maximum length: `20` |
|   `key` | `string` | No | The node_id for a node can be obtained by calling the View Specified Approval Definition API and retrieving it from the node_list parameter in the interface's response. **Example value**: "46e6d96cfa756980907209209ec03b64" |
|   `value` | `string[]` | No | The CC recipient list requires the user_id. Refer to How to Obtain User's User ID for the method to acquire user_ids. **Example value**: ["f7cb567e"] |
| `node_cc_open_id_list` | `node_cc[]` | No | If the approval process definition includes a node where the initiator needs to select the CC recipient, this parameter should be used to specify the CC recipient for the corresponding node (by designating the recipient using the open_id). **Notes**: If both node_cc_user_id_list and node_cc_open_id_list are provided, the CC recipients will be the union of both lists. **Data validation rules**: - Maximum length: `20` |
|   `key` | `string` | No | The node_id of a node can be obtained by calling the View Specified Approval Definition API and retrieving it from the node_list parameter in the response of the API. **Example value**: "46e6d96cfa756980907209209ec03b64" |
|   `value` | `string[]` | No | The CC recipient list requires the user's open_id. Refer to How to Obtain User's Open ID for the method to acquire open_ids. **Example value**: ["f7cb567e"] |
| `uuid` | `string` | No | Approval instance uuid, used for idempotent action. A single uuid can only be used to create 1 approval instance. In the case of a conflict, the error code 60012 will be returned. The format must be: XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX (case-insensitive). **Example value**: "7C468A54-8745-2245-9675-08B7C63E7A87" **Data validation rules**: - Length range: `1` ～ `64` characters |
| `allow_resubmit` | `boolean` | No | Whether the **Submit** button is configured is applicable when the approver of the task returns the approval document, allowing the approval submitter to click **Submit** within the same approval instance to submit the document again. **Example value**: true |
| `allow_submit_again` | `boolean` | No | Create an approval instance interface to support configuration whether to allow submit  again or not, applicable to periodic bill of lading scenarios, where a new instance is launched again based on the current form content **Example value**: true |
| `cancel_bot_notification` | `string` | No | Configure the bot to suppress specified notification results. Optional values: - 1: Cancel push. - 2: Cancel rejection push. - 4: Cancel the instance and cancel the push. Supports canceling a single bot notification, and also supports canceling multiple bot notifications at the same time. Bit operation, that is, if you need to cancel the two notifications 1 and 2, you need to pass in the sum value 3. **Example value**: "1" |
| `forbid_revoke` | `boolean` | No | Configure whether revocation can be prohibited **Example value**: false **Default value**: `false` |
| `i18n_resources` | `i18n_resource[]` | No | Internationalized text. Only the value of input and textarea are supported. |
|   `locale` | `string` | Yes | Language **Example value**: "zh-CN" **Optional values are**:  - `zh-CN`: Chinese - `en-US`: English - `ja-JP`: Japanese  |
|   `texts` | `i18n_resource_text[]` | Yes | Copywriting Key:Value. The Key needs to start with @i18n@ and the Value should be passed in according to the requirements of each parameter. **Explanation**: This field is mainly used for internationalization and allows setting copywriting for multiple languages simultaneously. The approval center will use the matched copywriting based on the actual language environment of the current user. If the copywriting for the user's current language environment is not set, the default language copywriting will be used. **Example value**: { "@i18n@1": "权限申请", "@i18n@2": "OA审批", "@i18n@3": "Permission" } |
|     `key` | `string` | Yes | The copywriting Key needs to match the Keys of each parameter. **Example value**: "@i18n@1" |
|     `value` | `string` | Yes | Copywriting **Example value**: "People" |
|   `is_default` | `boolean` | Yes | Whether to use the default language. If the default language is used, all keys should be contained; if a non-default language is used but its keys do not exist, the default language will be used instead. **Example value**: true |
| `title` | `string` | No | Display name of the approval instance. If this parameter is filled in, the approval name in the approval list will use this parameter. If this parameter is not filled in, the approval name will use the name defined by the approval. **Explanation**: The Key for the internationalized copywriting is passed here (i.e., the Key in the i18n_resources.texts parameter). It must start with @i18n@ and also needs to be assigned in the i18n_resources.texts parameter in the format of Key:Value. **Example value**: "@i18n@1" |
| `title_display_method` | `int` | No | Display mode for the title on the approval details page. **Example value**: 0 **Optional values are**:  - `0`: If both the approval definition and approval instance have a title, display both, separated by a vertical line. - `1`: If both the approval definition and approval instance have a title, only display the title of the approval instance.  **Default value**: `0` |
| `node_auto_approval_list` | `node_auto_approval[]` | No | Set up auto-approval nodes. **Data validation rules**: - Maximum length: `10` |
|   `node_id_type` | `string` | No | Node ID type **Example value**: "NON_CUSTOM" **Optional values are**:  - `CUSTOM`: Custom node ID - `NON_CUSTOM`: Non-custom node ID  |
|   `node_id` | `string` | No | The value of the node ID can be obtained by calling the Get Specified Approval Definition API. The node ID can be found in the node_list parameter of the API response. **Example value**: "manager_node_id" | ### Request body example

{
    "approval_code":"4202AD96-9EC1-4284-9C48-B923CDC4F30B",
    "user_id":"59a92c4a",
    "open_id":"ou_806a18fb5bdf525e38ba219733bdbd73",
    "form":"[{\"id\":\"111\",\"type\":\"input\",\"value\":\"11111\"},{\"id\":\"222\",\"required\":true,\"type\":\"dateInterval\",\"value\":{\"start\":\"2019-10-01T08:12:01+08:00\",\"end\":\"2019-10-02T08:12:01+08:00\",\"interval\": 2.0}},{\"id\":\"333\",\"type\":\"radioV2\",\"value\":\"1\"},{\"id\":\"444\",\"type\":\"number\", \"value\":\"4\"},{\"id\":\"555\",\"type\":\"textarea\",\"value\":\"fsafs\"}]",
    "node_approver_user_id_list":[
        {"key": "46e6d96cfa756980907209209ec03b64","value":["59a92c4a"]},
        {"key": "manager_node_id","value":["59a92c4a"]}
    ],
    "node_approver_open_id_list":[
        {"key": "46e6d96cfa756980907209209ec03b64","value":["ou_806a18fb5bdf525e38ba219733bdbd73"]},
        {"key": "manager_node_id","value":["ou_806a18fb5bdf525e38ba219733bdbd73"]}
    ],
    "node_cc_user_id_list":[
        {"key": "46e6d96cfa756980907209209ec03b64","value":["59a92c4a"]},
        {"key": "manager_node_id","value":["59a92c4a"]}
    ],
    "node_cc_open_id_list":[
        {"key": "46e6d96cfa756980907209209ec03b64","value":["ou_806a18fb5bdf525e38ba219733bdbd73"]},
        {"key": "manager_node_id","value":["ou_806a18fb5bdf525e38ba219733bdbd73"]}
    ]
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `instance_code` | `string` | Approval Instance Code | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "instance_code": "81D31358-93AF-92D6-7425-01A5D67C4E71"
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1390001 | param is invalid | Parameter error. Troubleshooting steps: - Check the parameters passed during the request according to the parameter descriptions in the API documentation. - If form parameters are passed, verify whether the form control data within the parameters is correct. If the error message includes a control ID (e.g., `Control= widget17261088448220001`), you can call the View specific approval definition or Get single approval instance details API to retrieve the response parameter form value, search for the problematic control ID, and check whether the configuration of the control is correct. |
| 400 | 1390015 | approval is not active | The approval definition has been deactivated. Please ensure that the current approval definition is activated before retrying. You can log in to the [Lark Approval Admin](https://www.larksuite.com/approval/admin/approvalList) to check whether the corresponding approval definition has been deactivated. ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/eac4092463fb565dd36ebe48cfadb95e_omtoYin2Gk.png) |
| 400 | 1390013 | unsupported approval for free process | Free flow not supported |
| 400 | 1395001 | There have been some errors. Please try again later | Troubleshooting steps for service errors: 1. Refer to the parameter descriptions in the API documentation and check whether the parameters passed in the request are correct. If form parameters are passed (form), check whether the form control data passed in is correct. 2. Reduce the request frequency and retry. If the error persists after retrying, please contact [Technical Support](https://applink.larksuite.com/TLJpeNdW). | For more error code information, see General error codes.
