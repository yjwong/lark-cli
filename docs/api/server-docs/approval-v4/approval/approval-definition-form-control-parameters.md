---
title: "Approval definition form control parameters"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/approval/approval-definition-form-control-parameters"
service: "approval-v4"
resource: "approval"
section: "Approval"
updated: "1760339457000"
---

# Approval definition form control parameters

In the API for Creating approval definition and Viewing specified approval definition, the form parameter containing the JSON data of approval definition form controls must be used. This article summarizes the parameter descriptions in the JSON data of various form controls for your reference.

## Controls unsupported by approval definition API

The approval definition API does not fully support all approval form controls. Unsupported controls are listed below. If you need to use unsupported controls, please proceed to [Lark Approval Admin](https://www.larksuite.com/approval/admin/approvalList?devMode=on) for operations.

**Control Name** | **Type**     |
| -------------- | ------------ |
| Formula        | formula      | 
| Data from Base   | mutableGroup |
| Serial no.  | serialNumber | 
| Shift swap | shiftGroupV2                |
| Leave(View only) | leaveGroupV2                |
| Overtime(View only) | workGroup                   | 
| Business trip | tripGroup                   | 
| Out-of-office(View only) | outGroup                    | ## General parameters

General parameters are those included in the JSON data of each control, and they are centrally explained in this chapter.

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| id | string | Yes | Control ID, which must be unique within the same approval definition. |
| name | string | Yes | Control name's i18n key, must start with `@i18n@` and correspond to the key in i18n_resources.texts parameter in the relevant interface. For example, when creating an approval definition, if the `name` value of the control is `@i18n@demo`, then the same value `@i18n@demo` must be passed to the key of i18n_resources.texts parameter, and the value corresponding to the key will be assigned to `name`. |
| type | string | Yes | Control type. Refer to the detailed parameter descriptions of each control type in the following sections. |
| required | boolean | Yes | Whether the control is required during approval instance creation. **Optional values**: - true - false |
| custom_id | string | No | Custom control ID. |
| printable | boolean | No | Whether the control can be printed. **Optional values**: - true - false **Default value**: false | ## Parameters of different controls

This chapter provides the `type` parameter values, JSON examples, and descriptions of non-general parameters for different controls.

> **Short answer**, **Paragraph**, **Single select**, **Contacts**, and **Department** controls support setting default values when creating approval definitions. Specific configuration instructions can be found in the [API description for setting default values for controls](https://feishu.feishu.cn/docx/GTcAdkmPZobyTNxHsIhcs1zhnCb).

### Short answer

Control type `type` is `input`, the JSON example is as follows:

```json
{
    "id":"widget123456",
    "name":"@i18n@demo_name",
    "type":"input",
    "required":true
}
```

### Paragraph

Control type `type` is `textarea`, the JSON example is as follows:

```json
{
    "id":"widget123456",
    "name":"@i18n@demo_name",
    "type":"textarea",
    "required":true
}
```

### Number

Control type `type` is `number`, the JSON example is as follows:

```json
{
  "id": "widget123456",
  "name": "@i18n@demo_name",
  "type": "number",
  "required": true
}
```

### Image/Video

Control type `type` is `image`, the JSON example is as follows:

```json
{
  "id": "widget123456",
  "name": "@i18n@demo_name",
  "type": "image",
  "required": true
}
```

### Attachment

Control type `type` is `attachmentV2`, the JSON example is as follows:

```json
{
  "id": "widget123456",
  "name": "@i18n@demo_name",
  "type": "attachmentV2",
  "required": true
}
```

### Amount

Control type `type` is `amount`, the JSON example is as follows:

```json
{
  "id": "widget123456",
  "name": "@i18n@demo_name",
  "type": "amount",
  "required": true,
  "value": "CNY",
  "option": {
    "currencyRange": [
      "CNY",
      "USD"
    ]
  }
}
```

Non-general parameter descriptions:

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| value | string | Yes | Control value, including: - CNY: Chinese Yuan - USD: US Dollar - EUR: Euro - JPY: Japanese Yen - CAD: Canadian Dollar - CHF: Swiss Franc - SGD: Singapore Dollar - AUD: Australian Dollar - KBW: South Korean Won - INR: Indian Rupee - TWD: New Taiwan Dollar - HKD: Hong Kong Dollar - MOP: Macanese Pataca - THB: Thai Baht - IDR: Indonesian Rupiah - PHP: Philippine Peso - MYR: Malaysian Ringgit |
| option | object | Yes | Option configuration. |
| └ currencyRange | string[] | Yes | Currency range. Optional values: - CNY: Chinese Yuan - USD: US Dollar - EUR: Euro - JPY: Japanese Yen - CAD: Canadian Dollar - CHF: Swiss Franc - SGD: Singapore Dollar - AUD: Australian Dollar - KBW: South Korean Won - INR: Indian Rupee - TWD: New Taiwan Dollar - HKD: Hong Kong Dollar - MOP: Macanese Pataca - THB: Thai Baht - IDR: Indonesian Rupiah - PHP: Philippine Peso - MYR: Malaysian Ringgit |
| └ isCapital | boolean | No | Whether to display numbers in uppercase, recommended to be true for RMB. |
| └ isThousandSeparator | boolean | No | Whether to display thousand separator. |
| └ keepDecimalPlaces | int | No | Set the number of decimal places to display. For example, setting to 2 indicates displaying 2 decimal places. |
| └ maxValue | int | No | Maximum value of the amount range. |
| └ minValue | int | No | Minimum value of the amount range. | ### Description

The widget type "text" is illustrated in the JSON example below:

```json
{
    "id":"widget123456",
    "name":"@i18n@demo_name",
    "type":"text",
    "required":true,
    "value":"@i18n@text"
}
```

Description of Non-Standard Parameters:

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| value | string | Yes | The internationalization key for the content, which must start with `@i18n@`, and should be assigned in the format of Key:Value in the i18n_resources.texts parameter of the corresponding interface. The length must not be less than 9 characters. | ### Single select, Multiple select

- The widget type for single-choice is "radioV2", as shown in the following JSON example:

	```json
    {
        "id":"widget123456",
        "name":"@i18n@demo_name",
        "type":"radioV2",
        "required":true,
        "value":[{"key":"1","text":"@i18n@choice1"},{"key":"2","text":"@i18n@choice2"}]
    }
	```

- The widget type for multi-choice is "checkboxV2", as shown in the following JSON example:

    ```json
    {
        "id":"widget123456",
        "name":"@i18n@demo_name",
        "type":"checkboxV2",
        "required":true,
        "value":[{"key":"1","text":"@i18n@choice1"},{"key":"2","text":"@i18n@choice2"}]
    }
    ```

Description of Non-Standard Parameters:

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| value | object[] | Yes | Configuration items for single-choice and multi-choice widgets. |
| └ key | string | Yes | Option key, which must be unique. |
| └ text | string | Yes | The internationalization key for the content of the option, which must start with `@i18n@`, and should be assigned in the format of Key:Value in the i18n_resources.texts parameter of the corresponding interface. | If the company uses multiple systems simultaneously (e.g., Lark Approval, HR system, Sales Management system), and needs to sync data from other systems to approval forms as options, it can configure external data sources for single-choice and multi-choice widgets to avoid maintaining the same data across multiple systems. For more details, refer to Associated External Option Description. After association, the parameters for single-choice and multi-choice widgets change, as shown in the following JSON example:

```json
{
    "id":"widget123456",
    "name":"@i18n@demo_name",
    "type":"radioV2",
    "required":true,
    "value":[],
    "externalData":{
        "externalUrl":"https://xxx.xxx.xxx/",
        "token":"t",
        "key":"k",
        "linkageConfigs":[
            {
                "linkageWidgetID":"widget1",
                "key":"linkageWidget1",
                "value":"example"
            }
        ],
        "externalDataLinkage":true
    }
}
```

Description of Non-Standard Parameters:

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| value | object[] | No | Used for fixed options. When associated with external options, the externalData parameter needs to be configured. |
| externalData | object | Yes | Information for associating external options. |
| └ externalUrl | string | Yes | The URL of the external data source interface. |
| └ token | string | Yes | The token for the external data source interface. |
| └ key | string | No | The key for the external data source interface. |
| └ linkageConfigs | object | No | Configuration for linkage parameters. |
| └ └ linkageWidgetID | string | No | The ID of the widget corresponding to the linkage parameter. If detail widgets are included, it is recommended to use the custom ID of the associated widget. |
| └ └ key | string | No | Parameter code. | ### Date

The widget type "date" is illustrated in the JSON example below:

```json
{
    "id":"widget123456",
    "name":"@i18n@demo_name",
    "type":"date",
    "required":true,
    "value": "YYYY-MM-DD"
}
```

Description of Non-Standard Parameters:

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| value | string | Yes | Date format. Optional values are: - YYYY-MM-DD: Year-Month-Day - YYYY-MM-DD a: Year-Month-Day AM/PM - YYYY-MM-DD hh:mm: Year-Month-Day Hour:Minute | ### Link approvals

The control type is `connect`. Below is an example in JSON format:

```json
{
    "id":"widget123456",
    "name":"@i18n@demo_name",
    "type":"connect",
    "required":true,
    "value":["code1","code2"]
}
```

Description of non-general parameters:

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| value | string[] | Yes | List of codes for the approval definitions to be associated. To obtain the approval definition codes: - Call the Create Approval Definition API, and get the `approval_code` from the response parameters. - Login to the approval management backend and get it from the URL of the specified approval definition. For detailed operations, see What is Approval Code. | ### Contacts

The control type is `contact`. Below is an example in JSON format:

```json
{
    "id":"widget123456",
    "name":"@i18n@demo_name",
    "type":"contact",
    "required":true,
    "value":{
        "ignore": true,
        "multi": false
      }
}
```

Description of non-general parameters:

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| value | object | No | Configuration options for the contact control. |
| └ ignore | boolean | No | Whether the user can select themself as a contact. Default is `false`, meaning the user can select themself. |
| └ multi | boolean | No | Whether the user can select multiple contacts. Default is `false`, meaning the user cannot select multiple contacts. | ### Address

The control type is `address`. Below is an example in JSON format:

```json
{
    "id":"widget123456",
    "name":"@i18n@demo_name",
    "type":"address",
    "required":true,
    "value":{
        "enableDetailAddress": false,
        "requiredDetailAddress": false,
        "preLocating": false
      }
}
```
Description of non-general parameters:

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| value | object | No | Configuration options for the address control. |
| └ enableDetailAddress | boolean | No | Whether to enable the detailed address configuration option. Default is `false`, meaning the detailed address option is not enabled. |
| └ requiredDetailAddress | boolean | No | Whether the detailed address is required. Default is `false`, meaning it is not required. |
| └ preLocating | boolean | No | Whether to enable automatic location. Default is `false`, meaning automatic location is not enabled. | ### Date range

The control type is `dateInterval`. Below is an example in JSON format:

```json
{
    "id":"widget123456",
    "name":"@i18n@demo_name",
    "type":"dateInterval",
    "required":true,
    "value":{
        "format": "YYYY-MM-DD",
        "intervalAllowModify": false,
      }
}
```

Description of non-general parameters:

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| value | object | Yes | Configuration options for the date interval control. |
| └ format | string | Yes | Date format. Options are: - YYYY-MM-DD: Year-Month-Day - YYYY-MM-DD a: Year-Month-Day AM/PM - YYYY-MM-DD hh:mm: Year-Month-Day Hours:Minutes |
| └ intervalAllowModify | boolean | No | Whether the user can modify the duration automatically calculated by the system when initiating approval. Default is `false`, meaning modification is not allowed. | ### Contact no.

The control type is `telephone`. Below is an example in JSON format:

```json
{
    "id":"widget123456",
    "name":"@i18n@demo_name",
    "type":"telephone",
    "required":true,
    "option":{
        "availableType": "FIXED_LINE_OR_MOBILE"
      }
}
```

Description of non-general parameters:

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| option | object | Yes | Configuration options for the telephone control. |
| └ availableType | string | Yes | Type of telephone available. Options are: - MOBILE: Mobile Phone - FIXED_LINE: Fixed Phone - FIXED_LINE_OR_MOBILE: Mobile or Fixed Phone | ### Details/Table

The control type is fieldList, and the JSON example is as follows:

```json
{
    "id": "widget123456",
    "name": "@i18n@demo_name",
    "type": "fieldList",
    "required": true,
    "value":
    [
        {
            "id": "",
            "name": "",
            "type": "input",
            "required": true
        }
    ],
    "option":
    {
        "inputType": "LIST",
        "printType": "LIST"
    }
}
```
Explanation of non-generic parameters:

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| value | object[] | Yes | Information about other controls added in the details/table control. **Note**: The details cannot include other details/table controls, payment account controls, serial number controls, and any control groups. |
| └ id | string | Yes | Control ID, should be unique within the same approval definition. |
| └ name | string | Yes | The internationalization key for the control name, which must start with @i18n@ and correspond to the key parameter in the i18n_resources.texts corresponding interface. For example, when creating an approval definition, if the control's name is @i18n@demo, the same key @i18n@demo should be passed in the i18n_resources.texts parameter, and the corresponding value of the key will be assigned to the name. |
| └ type | string | Yes | Control type. |
| └ required | boolean | Yes | Whether the current control is mandatory when creating approval instances. **Possible values**: - true - false |
| option | object | Yes | Configuration options for the details control. |
| └ inputType | string | Yes | Input format for the details control. Possible values are: - LIST: vertical input - FORM: horizontal input |
| └ printType | string | Yes | Print format for the details control. Possible values are: - LIST: vertical print - FORM: horizontal print | ## Widget Group
A widget group is a special widget that combines several sub-widgets (basic widgets, such as radio, text, etc.) and a large amount of built-in logic (such as automatic value assignment of a sub-widget). The definition format and usage are different from those of basic controls. For details, please refer to the widget group parameter description.

### leave widget group
![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/41227687c0169aa4a161f5e075dd47a3_9biFnPSiVc.png?lazyload=true&width=2310&height=1686)

#### create approval definition

Creating definitions containing leave widget group through OpenAPI is not currently supported

#### approval definition detail

widget type is leaveGroupV2，JSON sample：

```json
{
  "id": "widgetLeaveGroupV2",
  "name": "",
  "option": null,
  "printable": true,
  "required": true,
  "type": "leaveGroupV2",
  "value": [{"id":"widgetLeaveGroupType","name":"leave type","option":[],"printable":true,"required":true,"type":"radioV2","visible":true},{"id":"widgetLeaveGroupStartTime","name":"start time","options":{"dateCheckEnd":0,"dateCheckStart":0,"dateCheckType":0},"printable":true,"required":true,"type":"date","value":"YYYY-MM-DD hh:mm","visible":true},{"id":"widgetLeaveGroupEndTime","name":"end time","options":{"dateCheckEnd":0,"dateCheckStart":0,"dateCheckType":0},"printable":true,"required":true,"type":"date","value":"YYYY-MM-DD hh:mm","visible":true},{"id":"widgetLeaveGroupInterval","name":"duration","option":[],"printable":true,"required":true,"type":"radioV2","visible":true},{"id":"widgetLeaveGroupReason","name":"leave reason","printable":true,"required":true,"type":"textarea","visible":true},{"id":"widgetLeaveGroupUnit","name":"unit","option":[{"value":"DAY","text":"day"},{"value":"HOUR","text":"hour"}],"printable":true,"required":true,"type":"radioV2","visible":true},{"default_value_type":"","display_condition":null,"enable_default_value":false,"id":"widgetLeaveGroupFeedingArrivingLate","name":"arrival late","option":[{"value":"0","text":"0"},{"value":"15","text":"15"},{"value":"30","text":"30"},{"value":"45","text":"45"},{"value":"60","text":"60"},{"value":"75","text":"75"},{"value":"90","text":"90"},{"value":"105","text":"105"},{"value":"120","text":"120"}],"printable":true,"required":false,"type":"radioV2","visible":true},{"id":"widgetLeaveGroupFeedingOffLeaveEarly","name":"leave early","option":[{"value":"0","text":"0"},{"value":"15","text":"15"},{"value":"30","text":"30"},{"value":"45","text":"45"},{"value":"60","text":"60"},{"value":"75","text":"75"},{"value":"90","text":"90"},{"value":"105","text":"105"},{"value":"120","text":"120"}],"printable":true,"required":false,"type":"radioV2","visible":true},{"id":"widgetLeaveGroupFeedingRestDaily","name":"rest time per day","printable":true,"required":false,"type":"input","visible":true},{"id":"widgetLeaveCertification","name":"leave cerificition","printable":true,"required":false,"type":"image","visible":true}],
  "visible": true
}
```

Widget group parameter description：

| parameter | type | description |
| --- | --- | --- |
| id | string | Leave widget group ID，fixed as widgetLeaveGroupV2 |
| type | string | Leave widget type，fixed as leaveGroupV2 |
| value | object[] | Sub-widget list, composed of basic controls, refer to the sub-widget parameter description | Sub-widget parameter descriptio:

| id | widget type | JSON sample | description |
| --- | --- | --- | --- |
| widgetLeaveGroupType | radioV2 | ``` { "id": "widgetLeaveGroupType", "name": "leave type", "option": [], "printable": true, "required": true, "type": "radioV2", "visible": true } ``` | Holiday type, radio type control, different from ordinary radio control, this control has no preset options, the options come from the vacation management background configuration |
| widgetLeaveGroupType | radioV2 | ``` { "id": "widgetLeaveGroupStartTime", "name": "start time", "printable": true, "required": true, "type": "date", "value": "YYYY-MM-DD hh:mm", "visible": true } ``` | Start time, date type widget |
| widgetLeaveGroupType | radioV2 | ``` { "id": "widgetLeaveGroupEndTime", "name": "end time", "printable": true, "required": true, "type": "date", "value": "YYYY-MM-DD hh:mm", "visible": true } ``` | End time, date type widget |
| widgetLeaveGroupType | radioV2 | ``` { "id": "widgetLeaveGroupInterval", "name": "duration", "option": [], "printable": true, "required": true, "type": "radioV2", "visible": true } ``` | Leave duration, single-select type widget, this widget does not need to be filled in when filing, the duration and unit will be automatically calculated based on the leave type, start and end time |
| widgetLeaveGroupType | radioV2 | ``` { "id": "widgetLeaveGroupReason", "name": "leave reason", "printable": true, "required": true, "type": "textarea", "visible": true } ``` | Reason for leave, text type widget |
| widgetLeaveGroupType | radioV2 | ``` { "id": "widgetLeaveGroupFeedingArrivingLate", "name": "Late for work (minutes)", "option": [{"value":"0","text":"0"},{"value":"15","text":"15"},{"value":"30","text":"30"},{"value":"45","text":"45"},{"value":"60","text":"60"},{"value":"75","text":"75"},{"value":"90","text":"90"},{"value":"105","text":"105"},{"value":"120","text":"120"}], "printable": true, "required": false, "type": "radioV2", "visible": true } ``` | The duration of being late for work is a radio-select type control. The options are built-in time ranges and cannot be edited. It is only used when submitting a breastfeeding leave application. |
| widgetLeaveGroupType | radioV2 | ``` { "id": "widgetLeaveGroupFeedingOffLeaveEarly", "name": "Leave work early (minutes)", "option": [{"value":"0","text":"0"},{"value":"15","text":"15"},{"value":"30","text":"30"},{"value":"45","text":"45"},{"value":"60","text":"60"},{"value":"75","text":"75"},{"value":"90","text":"90"},{"value":"105","text":"105"},{"value":"120","text":"120"}], "printable": true, "required": false, "type": "radioV2", "visible": true } ``` | The length of time to leave work early is a single-select type control. The options are built-in time ranges and cannot be edited. It is only used when submitting a breastfeeding leave application. | ### Overtime work widget group

When associating overtime rules, there is no overtime type
![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/26484184760379c033ef3e7b658b685f_3FlyBwM1OY.png?lazyload=true&width=2252&height=1110)

When overtime rules are not associated, you can set the overtime type and the associated vacation type.
![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/84fc62da0e7639541c51fa398ca5a216_OFHhtjfeaq.png?lazyload=true&width=2282&height=1382)

#### Create approval definition
Currently, it is not supported to create definitions containing overtime work widget group through OpenAPI

#### Approval definition detail
widget type is leaveGroupV2，JSON sample: 
```
{
    "id": "widgetWorkGroup",
    "name": "overtime",
    "option": {
      "allowInsteadMultiUser": 0,
      "allowMultiTimeRange": 1,
      "isSetRule": 1
    },
    "printable": true,
    "required": true,
    "type": "workGroup",
    "value": [{"id":"widgetWorkGroupType","name":"overtime type","option":[],"printable":true,"required":true,"type":"radioV2","visible":true,"widget_default_value":""},{"children":[{"id":"widgetWorkGroupStartTime","name":"start time","options":{"dateCheckEnd":0,"dateCheckStart":0,"dateCheckType":0},"printable":true,"required":true,"type":"date","value":"YYYY-MM-DD hh:mm","visible":true,"widget_default_value":""},{"default_value_type":"","display_condition":null,"enable_default_value":false,"id":"widgetWorkGroupEndTime","name":"end time","options":{"dateCheckEnd":0,"dateCheckStart":0,"dateCheckType":0},"printable":true,"required":true,"type":"date","value":"YYYY-MM-DD hh:mm","visible":true,"widget_default_value":""}],"default_value_type":"","display_condition":null,"enable_default_value":false,"id":"widgetWorkGroupTimeRangeFieldList","name":"time range","option":{"input_type":"LIST","mobile_detail_type":"CARD","print_type":"LIST"},"printable":true,"required":false,"type":"fieldList","visible":true,"widget_default_value":""},{"id":"widgetWorkGroupInterval","name":"Overtime hours","printable":true,"required":true,"type":"number","visible":true,"widget_default_value":""},{"id":"widgetWorkGroupReason","name":"Reasons for overtime","printable":true,"required":true,"type":"textarea","visible":true,"widget_default_value":""},{"id":"widgetWorkGroupUnit","name":"overtime unit","option":[{"value":"HOUR","text":"hour"},{"value":"DAY","text":"day"},{"value":"MINUTE","text":"minute"}],"printable":true,"required":true,"type":"radioV2","visible":true,"widget_default_value":""},{"id":"widgetWorkGroupRule","name":"overtime rule","printable":true,"required":true,"type":"input","visible":true,"widget_default_value":""},{"children":[{"id":"widgetWorkGroupDetailType","name":"overtime type","option":[{"value":"LEAVE","text":"Adjustment of leave"},{"value":"PAY","text":"overtime pay"},{"value":"NONE","text":"none"}],"printable":true,"required":true,"type":"radioV2","visible":true,"widget_default_value":""},{"id":"widgetWorkGroupDetailStartTime","name":"start time","options":{"dateCheckEnd":0,"dateCheckStart":0,"dateCheckType":0},"printable":true,"required":true,"type":"date","value":"YYYY-MM-DD hh:mm","visible":true,"widget_default_value":""},{"id":"widgetWorkGroupDetailEndTime","name":"end time","options":{"dateCheckEnd":0,"dateCheckStart":0,"dateCheckType":0},"printable":true,"required":true,"type":"date","value":"YYYY-MM-DD hh:mm","visible":true,"widget_default_value":""},{"default_value_type":"","display_condition":null,"enable_default_value":false,"id":"widgetWorkGroupDetailInterval","name":"Overtime hours","option":[],"printable":true,"required":true,"type":"radioV2","visible":true,"widget_default_value":""},{"id":"widgetWorkGroupDetailCategory","name":"date type","option":[{"value":"0","text":"none"},{"value":"1","text":"workday"},{"value":"2","text":"rest day"},{"value":"3","text":"vacation"}],"printable":true,"required":true,"type":"radioV2","visible":true,"widget_default_value":""}],"id":"widgetWorkGroupDetail","name":"overtime detail","option":{"input_type":"LIST","mobile_detail_type":"CARD","print_type":"LIST"},"printable":true,"required":true,"type":"fieldList","visible":true,"widget_default_value":""},{"id":"widgetWorkGroupOvertimeWorkers","name":"overtime worker","printable":true,"required":false,"type":"contact","visible":true,"widget_default_value":""}],
    "visible": true,
    "widget_default_value": ""
}
```

Widget group parameter description：

| parameter | type | description |
| --- | --- | --- |
| id | string | Overtime widget group ID, id is fixed widgetWorkGroup |
| type | string | Overtime widget group type，type is fixed workGroup |
| value | object[] | Sub-widget list, composed of basic controls, refer to the sub-widget parameter description |
| option | object | Widget group properties - allowInsteadMultiUser: Allows submission on behalf of multiple people, only available when associated with overtime rules - allowMultiTimeRange: Allows you to submit multiple overtime periods. This is only available when overtime rules are associated - isSetRule: Whether the overtime rule is associated, 1 means associated, otherwise it means not associated | Sub-widget parameter description：

| id | widget type | JSON sample | description |
| --- | --- | --- | --- |
| widgetWorkGroupType | radioV2 | ``` { "id": "widgetWorkGroupType", "name": "Overtime type", "option": [ { "value": "-1", "text": "compensated days off" }, ], "printable": true, "required": true, "type": "radioV2", "visible": true, "widget_default_value": "" } ``` | Overtime type, a radio widget. If overtime rules are associated, there is no option in the definition and no need to fill it out when making a bill of lading. Otherwise, the option will be returned through the option of this widget. In this case, this widget is required |
| widgetWorkGroupType | contact | ``` { "id": "widgetWorkGroupOvertimeWorkers", "name": "overtime worker", "printable": true, "required": true, "type": "contact", "visible": true, } ``` | Overtime person, contact person widget. If the space group allows submission on behalf of multiple people, you need to carry this widget when you submit the bill. If you cannot submit on behalf of others, you do not need to carry this control |
| widgetWorkGroupType | fieldList | ``` { "id": "widgetWorkGroupTimeRangeFieldList", "name": "overtime time range", "option": {}, "printable": true, "required": false, "type": "fieldList", "visible": true, "children": [{"default_value_type":"","display_condition":null,"enable_default_value":false,"id":"widgetWorkGroupStartTime","name":"start time","options":{"dateCheckEnd":0,"dateCheckStart":0,"dateCheckType":0},"printable":true,"required":true,"type":"date","value":"YYYY-MM-DD hh:mm","visible":true,"widget_default_value":""},{"default_value_type":"","display_condition":null,"enable_default_value":false,"id":"widgetWorkGroupEndTime","name":"end time","options":{"dateCheckEnd":0,"dateCheckStart":0,"dateCheckType":0},"printable":true,"required":true,"type":"date","value":"YYYY-MM-DD hh:mm","visible":true,"widget_default_value":""}] } ``` | Overtime period, detail control, sub-widgets are start and end time controls of date type |
| └widgetWorkGroupStartTime | number | ``` { "id": "widgetWorkGroupStartTime", "name": "start time", "options": {}, "printable": true, "required": true, "type": "date", "value": "YYYY-MM-DD hh:mm", "visible": true, } ``` | Start time, date type, start time of overtime |
| └widgetWorkGroupEndTime | number | ``` { "id": "widgetWorkGroupEndTime", "name": "end time", "options": {}, "printable": true, "required": true, "type": "date", "value": "YYYY-MM-DD hh:mm", "visible": true, } ``` | Start time, date type, start time of overtime |
| widgetWorkGroupInterval | number | ``` { "id": "widgetWorkGroupInterval", "name": "Overtime hours", "printable": true, "required": true, "type": "number", "visible": true, } ``` | Overtime hours, a digital widget, is automatically calculated based on the overtime type (or automatically associated overtime rules), overtime start and end time. This control does not need to be included when making a bill of lading |
| widgetWorkGroupReason | textarea | ``` { "id": "widgetWorkGroupReason", "name": "Reasons for overtime", "printable": true, "required": true, "type": "textarea", "visible": true, } ``` | Overtime reason, text widget, in the settings, the visible and required attributes of the overtime reason correspond to visible and required respectively. If it is not visible, the bill of lading does not need to carry this control, otherwise the control is required | ### Out of office widget group
Set outing type
![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/0c934beb6c8e60586b3f50fa57b7baa1_W6umxgAMZJ.png?lazyload=true&width=2256&height=1234)
Out of office type not set
![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/a40f967014411e7fb08ece11fe723c6b_2sdo9fJKQ1.png?lazyload=true&width=2216&height=1166)

#### Create approval definition
Currently, it is not supported to create definitions containing out widget groups through OpenAPI.

#### Approval definition detial
```
{
    "id": "widgetOutGroup",
    "name": "out of office widget group",
    "option": {
      "defaultUnit": "DAY",
      "isSetType": 1,
      "unitMap": {
        "meijufjq-iv2c5qrlm1i-0": "DAY",
        "meijuivb-aqhae0ptrt-0": "HOUR",
        "meijuivb-mhca5ofoj8-0": "HALFDAY"
      }
    },
    "printable": true,
    "required": true,
    "type": "outGroup",
    "value": [{"id":"widgetOutGroupType","name":"out type","option":[{"value":"meijuivb-aqhae0ptrt-0","text":"hour"},{"value":"meijuivb-mhca5ofoj8-0","text":"half day"},{"value":"meijufjq-iv2c5qrlm1i-0","text":"day"}],"printable":true,"required":true,"type":"radioV2","visible":true},{"id":"widgetOutGroupUnit","name":"duration unit","option":[{"value":"HOUR","text":"hour"},{"value":"DAY","text":"day"}],"printable":true,"required":true,"type":"radioV2","visible":true},{"id":"widgetOutGroupStartTime","name":"start time","options":{"dateCheckEnd":0,"dateCheckStart":0,"dateCheckType":0},"printable":true,"required":true,"type":"date","value":"YYYY-MM-DD hh:mm","visible":true},{"id":"widgetOutGroupEndTime","name":"end time","options":{"dateCheckEnd":0,"dateCheckStart":0,"dateCheckType":0},"printable":true,"required":true,"type":"date","value":"YYYY-MM-DD hh:mm","visible":true},{"id":"widgetOutGroupInterval","name":"duration","option":[],"printable":true,"required":true,"type":"radioV2","visible":true},{"id":"widgetOutGroupReason","name":"out reason","printable":true,"required":true,"type":"textarea","visible":true},{"id":"widgetOutGroupImage","name":"Photo proof","printable":true,"required":false,"type":"image","visible":true}]
}
```

Widget group parameter description：

| parameter | type | description |
| --- | --- | --- |
| id | string | Out widget group ID, id is fixedwidgetOutGroup |
| type | string | Out widget group type，type is fixed outGroup |
| value | object[] | Sub-widget list, composed of basic controls, refer to the sub-widget parameter description |
| option | object | Widget group properties - isSetType: Is the outgoing type set - defaultUnit: Unit of outing duration, available when outing type is not set - unitMap: Mapping between outing type and duration unit | Subcontrol parameter description：

| id | widget type | JSON sample | description |
| --- | --- | --- | --- |
| widgetOutGroupType | radioV2 | ``` { "id": "widgetOutGroupType", "name": "out type", "option": [ { "value": "meijuivb-aqhae0ptrt-0", "text": "out all day" }, { "value": "meijuivb-mhca5ofoj8-0", "text": "out half day" }, { "value": "meijufjq-iv2c5qrlm1i-0", "text": "out one hour" } ], "printable": true, "required": true, "type": "radioV2", "visible": true } ``` | Out type, radio button. If the out type is set, the optional out type will be returned through the option of the widget. Otherwise, there is no optional value and the widget does not need to be included when submitting. |
| widgetOutGroupStartTime | date | ``` { "id": "widgetOutGroupStartTime", "name": "start time", "options": {}, "printable": true, "required": true, "type": "date", "value": "YYYY-MM-DD hh:mm" } ``` | out start time，date type widget |
| widgetOutGroupEndTime | date | ``` { "id": "widgetOutGroupEndTime", "name": "end time", "options": {}, "printable": true, "required": true, "type": "date", "value": "YYYY-MM-DD hh:mm" } ``` | out end time，date type widget |
| widgetOutGroupInterval | radioV2 | ``` { "id": "widgetOutGroupInterval", "name": "duration", "options": {}, "printable": true, "required": true, "type": "radioV2", "value": "YYYY-MM-DD hh:mm" } ``` | out duration, single-select type, automatically calculated based on outing start and end time |
| widgetOutGroupReason | textarea | ``` { "id": "widgetOutGroupReason", "name": "out reason", "printable": true, "required": true, "type": "textarea", "visible": true } ``` | The reason for going out, text widget, visibility and mandatory are controlled by the visible and required fields. If it is not visible, the widget does not need to be included in the submission. |
| widgetOutGroupImage | image | ``` { "id": "widgetOutGroupImage", "name": "photo prove", "printable": true, "required": false, "type": "image", "visible": true } ``` | When taking photos outside, the visibility and mandatory nature of the image control are controlled by the visible and required fields. If it is not visible, the widget does not need to be included in the submission. |
