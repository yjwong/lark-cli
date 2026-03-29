---
title: "Approval instance form control parameters"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/instance/approval-instance-form-control-parameters"
service: "approval-v4"
resource: "instance"
section: "Approval"
updated: "1760339462000"
---

# Approval instance form control parameters

When calling the Create approval instance API, you need to use form control parameters. This document helps you understand the parameters of various form controls within an approval instance.

## Preparation

The form control parameters of an approval instance are configured based on the approval definition form. For example, if the form design of the approval definition includes **Short answer** and **Date range** controls, then the form control parameters of the approval instance must be assigned for **Short answer** and **Date range** controls. Therefore, before operating the form control parameters of the approval instance, please first understand the Approval definition form control parameters.

## Unsupported controls in approval instance API

The Create approval instance API does not fully support all approval form controls. The unsupported controls are listed in the table below. If you must use controls that the API does not support, you will need to operate through the [Lark Approval Admin](https://www.larksuite.com/approval/admin/approvalList?devMode=on).

**Control/Control Group** | **Type**                    |
| ----------------------- | --------------------------- |
| Address                 | address                     |
| Description             | text                        |
| Data from Base | mutableGroup                |
| Payee account         | account                     |
| Serial no.          | serialNumber                |
| Business trip Control Group      | tripGroup                   |
| Employment Control Group | apaascorehrOnboardingGroup  |
| Conversion Control Group | apaascorehrRegularateGroup |
| Clock-in/out Correction Control Group | remedyGroupV2       |
| Transfer Control Group | apaascorehrJobAdjustGroup |
| Resignation Control Group | apaascorehrOffboardingGroup | ## General parameters

The form controls of an approval instance all include the parameters as shown in the table below.

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| id | string | Yes | ID of the control, which must match the ID of the control in the approval definition. |
| type | string | Yes | Type of the control. Refer to the **Parameters of Different Controls** section below for the values of different control types. |
| value | Varies by control type | Yes | Value of the control. The data type of `value` varies by control type. For example, for a single-line text control, `value` is a string; for a contact control, `value` is an array. Refer to the **Parameters of Different Controls** section below for details. | ## Parameters of different controls

This section provides the `type` parameter values, JSON examples, and non-general parameter descriptions for different controls.

### Short answer

Control type is `input`. JSON data example:

```json
{                                     
    "id": "widget1",
    "type": "input",
    "value": "data" // string type
}
```

### Paragraph

Control type is `textarea`. JSON data example:

```json
{                                     
    "id": "widget1",
    "type": "textarea",
    "value": "data" // string type
}
```

### Date

Control type is `date`. JSON data example:

```json
{                                     
    "id": "widget1",
    "type": "date",
    "value": "2019-10-01T08:12:01+08:00" // must be a string in RFC3339 format
}
```

### Date range

Control type is `dateInterval`. JSON data example:

```json
{
    "id": "widget1",
    "type": "dateInterval",
    "value": {
         "start":"2019-10-01T08:12:01+08:00",
         "end":"2019-10-02T08:12:01+08:00",
         "interval": 1.0
     }
}
```

The `value` parameter is an object type and includes the following parameters:

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| start | string | Yes | Start time, must be in RFC3339 format. |
| end | string | Yes | End time, must be in RFC3339 format. |
| interval | float | Yes | Duration (days). | ### Single select

Control type is `radio/radioV2`. JSON data example:

```json
{                                     
    "id": "widget1",
    "type": "radioV2",
    "value": "k2b8mkx0-h71x5gl1234-1" // string type
}
```

Here, `value` represents the option value. The range of values should refer to the `value` parameter of the **Single Choice** control options in the corresponding approval definition. You can call the View Approval Definition API to get the option values for single choice controls from the response's `form` parameter. If external options are linked, `value` should pass the external option’s options.id.

### Multiple select

Control type is `checkbox/checkboxV2`. JSON data example:

```json
{
    "id":"widget1",
    "type":"checkboxV2",
    "value": ["k2b8mkx0-h71x5gl4321-1"] // array of strings
}
```
Here, `value` represents the option values. The range of values should refer to the `value` parameter of the **Multiple Choice** control options in the corresponding approval definition. You can call the View Approval Definition API to get the option values for multiple choice controls from the response's `form` parameter. If external options are linked, `value` should pass the external option’s options.id.

### Number

Control type is `number`. JSON data example:

```json
{
    "id": "widget1",
    "type": "number",
    "value": 1234.5678 // float type
}
```

### Amount

Control type is `amount`. JSON data example:

```json
{
    "id": "widget1",
    "type": "amount",
    "value": 1234.5678, // float type
    "currency":"USD"
}
```

Here, `currency` represents the type of currency. The range of values should refer to the `value` parameter of the **Amount** control in the corresponding approval definition. You can call the View Approval Definition API to get the available currency types for the amount control from the response's `form` parameter.

### Formula

For the widget type `formula`, here is an example of JSON data:

```json
{
    "id": "widget1",
    "type": "formula",
    "value": 1234.5678 // This value is computed based on the formula defined within the approval configuration. If there is a mismatch, an error will be returned.
}
```

### Contacts

For the widget type `contact`, here is an example of JSON data:

```json
{
    "id": "widget1",
    "type": "contact",
    "value": ["f8ca557e"], // An array of strings.
    "open_ids": ["ou_12345"] // An array of strings.
}
```

In this instance, `value` contains the user's user_id; `open_ids` contains the user's open_id. For details on different user IDs, refer to the User Identity Overview.

### Link approvals

For the widget type `connect`, here is an example of JSON data:

```json
{
    "id": "widget1",
    "type": "connect",
    "value": ["19EAC829-F1CB-527F-BE2A-1330422E60C0"] // An array of strings.
}
```

`value` contains the code of the associated approval instance. You can use the Get Single Approval Instance Details API to retrieve instance details based on the approval instance code.

### Document

For the widget type `document`, here is an example of JSON data:

```json
{
    "id": "widget1",
    "type": "document",
    "value": {
           "token":"TLLKdcpDro9ijQxA33ycNMabcef",
           "type":"docx",
    }
}
```

The `value` parameter is an object and includes the following fields:

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| token | string | yes | The document's document_id. For more details, see Document. |
| type | string | yes | Type of the document, supports `docx`. | ### Attachment

For the widget type `attachmentV2`, here is an example of JSON data:

```json
{
    "id":"widget1",
    "type":"attachmentV2",
    "value": ["D93653C3-2609-4EE0-8041-61DC1D84F0B5"] // An array of strings.
}
```

`value` contains the file code returned by the Upload File API.

### Image/Video

For the widget type `image/imageV2`, here is an example of JSON data:

```json
{
    "id":"widget1",
    "type":"image",
    "value": ["D93653C3-2609-4EE0-8041-61DC1D84F0B5"] // An array of strings.
}
```

`value` contains the file code returned by the Upload File API.

### Details/Table

For the widget type `fieldList`, here is an example of JSON data:

```json
{
    "id": "widget1",
    "type": "fieldList",
    "value": [
         [   
            {
                "id": "widget1",
                "type": "checkbox",
                "value": ["jxpsebqp-0"]
            }
         ]
     ]
}
```

`value` is a two-dimensional array, and the JSON values of the controls are set based on the controls contained within the **Detailed/Table** widget of the approval definition.

### Department

For the widget type `department`, here is an example of JSON data:

```json
{
    "id": "widget1",
    "type": "department",
    "value": [ 
        {
            "open_id": "od-xxx"
        }
    ]
}
```

`value` is an array of objects and includes the `open_id` to set the department's open_department_id. For details regarding the open_department_id, refer to the Department Resource Overview.

### Contact no.

For the widget type `telephone`, here is an example of JSON data:

```json
{
    "id": "widget1",
    "type": "telephone",
    "value": {
        "countryCode":"+86",
        "nationalNumber":"13122222222"
    }
}
```

The `value` parameter is an object and includes the following fields:

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| countryCode | string | yes | Country code. |
| nationalNumber | string | yes | Telephone number. | ### Shift swap

For the widget type `shiftGroup`, here is an example of JSON data:

```json
{
    "id": "widget1",
    "type": "shiftGroup",
    "value": {
         "shiftTime": "2019-10-01T08:12:01+08:00",
         "returnTime": "2019-10-02T08:12:01+08:00",
         "reason": "ask for leave"
     }
}
```

The `value` parameter is an object and includes the following fields:

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| shiftTime | string | yes | Shift time, must be in RFC3339 format. |
| returnTime | string | yes | Return date, must be in RFC3339 format. |
| reason | string | yes | Reason for the shift. | ### leave widget group

For the widget type `shiftGroup`, here is an example of JSON data:
```
{
    "id": "widgetLeaveGroupV2",
    "type": "leaveGroupV2",
    "value": [
      {
        "id": "widgetLeaveGroupType",
        "type": "radioV2",
        "value": "7488925543484620819"
      },
      {
        "id": "widgetLeaveGroupStartTime",
        "type": "date",
        "value": "2025-08-25T11:30:00+08:00"
      },
      {
        "id": "widgetLeaveGroupEndTime",
        "type": "date",
        "value": "2025-08-26T11:35:00+08:00"
      },
      {
        "id": "widgetLeaveGroupReason",
        "type": "textarea",
        "value": "123123"
      },
      {
        "id": "widgetLeaveCertification",
        "type": "image",
        "value": [
          "B69F8E26-0EAA-4A92-9B80-DA613CD36136"
        ]
      },
      {
        "id":"widgetLeaveCertification",
        "type":"image",
        "value": ["D93653C3-2609-4EE0-8041-61DC1D84F0B5"]
      },
      {                                     
        "id": "widgetLeaveGroupFeedingArrivingLate",
        "type": "radioV2",
        "value": "30"
      },
      {                                     
        "id": "widgetLeaveGroupFeedingOffLeaveEarly",
        "type": "radioV2",
        "value": "30"
      }        
    ]
}
```

The `leaveWidgetGroup` contains parameter descriptions:
| id | type | JSON sample | description |
| --- | --- | --- | --- |
| id | string | 是 | Group ID, fixed to widgetLeaveGroupV2 |
| type | string | 是 | Group type, fixed toleaveGroupV2 |
| value | object[] | yes | The value of the group, which is a list of multiple sub-widget values | Description of the sub-control value contained in value:
| id | type | JSON sample | description |
| --- | --- | --- | --- |
| widgetLeaveGroupType | radioV2 | ``` { "id": "widgetLeaveGroupType", "type": "radioV2", "value": "7488925543484620819" } ``` | Holiday type. For the specific format, please refer to the radio button. The options are obtained from the leave interface. This button must be included when making a bill of lading |
| widgetLeaveGroupStartTime | date | ``` { "id": "widgetLeaveGroupStartTime", "type": "date", "value": "2019-10-01T08:12:01+08:00", // 需满足 RFC3339 格式的 string 类型 } ``` | The start time of the leave. For the specific format, please refer to the date control. It will be automatically rounded according to the type of leave. If it is less than 12 o'clock for a half-day leave, it will be considered as the morning. For an hourly leave, it will be rounded up to half an hour. This control must be included when making a bill of lading |
| widgetLeaveGroupEndTime | date | ``` { "id": "widgetLeaveGroupEndTime", "type": "date", "value": "2019-10-01T08:12:01+08:00", } ``` | The leave end time, for the specific format, please refer to the date control. It will be automatically rounded according to the leave type. If it is less than 12 o'clock for a half-day leave, it will be considered as the morning. If it is an hourly leave, it will be rounded back to half an hour. |
| widgetLeaveGroupReason | textarea | ``` { "id": "widgetLeaveGroupReason", "type": "textarea", "value": "123123" } ``` | For the specific format of the leave request, please refer to the multi-line text control. No entry is required for breastfeeding leave. For other cases, it is determined by whether the control is visible and required in the control group configuration. |
| widgetLeaveCertification | image | ``` { "id":"widgetLeaveCertification", "type":"image", "value": ["D93653C3-2609-4EE0-8041-61DC1D84F0B5"] } ``` | Leave certificate, for the specific format, please refer to the image control. If the selected leave type requires additional proof in the leave management background configuration, this value must be passed. An error will be reported if it is missing ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/8ab97ae490ac7c3f2008a3b554478c45_3qS7W30Qxv.png?height=546&lazyload=true&width=1654) |
| widgetLeaveGroupFeedingArrivingLate | radioV2 | ``` { "id": "widgetLeaveGroupFeedingArrivingLate", "type": "radioV2", "value": "30" } ``` | The number of minutes late for work. For the specific format, refer to the radio button. This field is only required for breastfeeding leave. The value range is 0-120 minutes, with a granularity of 15 minutes. The options are obtained from the option of this control in the approval definition. |
| widgetLeaveGroupFeedingOffLeaveEarly | radioV2 | ``` { "id": "widgetLeaveGroupFeedingOffLeaveEarly", "type": "radioV2", "value": "30" } ``` | The number of minutes to leave work early. For the specific format, please refer to the radio button. This field is only required for breastfeeding leave. The value range is 0-120 minutes, and the granularity is 15 minutes. The option is the string corresponding to the minute. | **Special parameter verification error information**
message                                            | description                                                                                                        |
| -------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------ |
| leave type id parse error                          | Leave type is not int64                                                                                            |
| group value is invalid                             | The value of the current control group is invalid. Please check whether it is empty or the check type is an array. |
| start time format is not RFC3339                   | The start time and date format is not in RFC3339 format                                                            |
| end time format is not RFC3339                     | The end time and date format is not in RFC3339 format                                                              |
| start time is after end time                       | The start time is later than the end time                                                                          |
| user not in gray                                   | The applicant is not in the leave grayscale                                                                        |
| leave type not found                               | The leave type does not exist                                                                                      |
| reason is required                                 | The leave reason is not entered                                                                                    |
| leave quote should be bigger than 0                | The leave duration must be greater than 0                                                                          |
| leave is conflict                                  | The selected timeframe already has a leave request. Please select a different timeframe                            |
| balance is not enough                              | Insufficient leave balance for the current leave type                                                              |
| certification is required                          | Leave request certificate must be uploaded                                                                         |
| arriving late is required                          | Maternity leave requires the duration of late arrival                                                              |
| arriving late value is not in the optional items   | Late arrival time is not within the selectable range                                                               |
| leaving early is required                          | Breastfeeding leave requires filling in the length of time to leave work in advance                                |
| leaving early value is not in the optional items   | The early leaving time is not within the optional range                                                            |
| feeding rest daily is 0                            | The daily rest period for breastfeeding leave is 0. Please select a new one                                        |
| the operation is prohibited by the workforce rules | The current account has been closed on the leave and cannot be submitted

### 加班控件组
**For the widget type `overtimeWidgetGroup`, here is an example of JSON data:**
```
{
  "id": "widgetWorkGroup",
  "type": "workGroup",
  "value":[
    {
      "id":"widgetWorkGroupOvertimeWorkers",
      "type":"contact",
      "value": ["f8ca557e"], 
      "open_ids": ["ou_12345"] 
    },
    {
      "id": "widgetWorkGroupType",
      "type": "radioV2",
      "value": "7259635026038505475" 
    },
    {
      "id":"widgetWorkGroupTimeRangeFieldList",
      "type":"fieldList",
      "value":[
        [
          {
            "id":"widgetWorkGroupStartTime",
            "type":"date",
            "value":"2019-10-01T08:12:01+08:00"
          },
          {
            "id":"widgetWorkGroupEndTime",
            "type":"date",
            "value":"2019-10-01T08:12:01+08:00"
          }
        ]
      ]
    },
    {
      "id": "widgetWorkGroupReason",
      "type": "textarea",
      "value": "111"
    }
  ]
}
```

**The parameter description of overtime widget group：**
| parameter | type | required | description |
| --- | --- | --- | --- |
| id | string | yes | Group ID, fixed as widgetWorkGroup |
| type | string | yes | Group type, fixed as workGroup |
| value | object[] | yes | The value of the group, which is a list of multiple sub-control values | Description of the sub-widget value contained in value:
| id | type | JSON sample | description |
| --- | --- | --- | --- |
| widgetWorkGroupOvertimeWorkers | contact | ``` { "id":"widgetWorkGroupOvertimeWorkers", "type":"contact", "value": ["f8ca557e"], "open_ids": ["ou_12345"] } ``` | List of overtime personnel. For the specific format, please refer to the contact control. If the definition is configured to "allow submission on behalf of multiple people", this field is required. If the submitter submits to himself, the submitter's ID must be filled in |
| widgetWorkGroupType | radioV2 | ``` { "id": "widgetWorkGroupType", "type": "radioV2", "value": "7259635026038505475" // 对应的类型选项ID } ``` | Overtime type. For the specific format, please refer to the radio button. If "Associated Overtime Rules" is turned off in the definition, you need to fill in this field |
| widgetWorkGroupTimeRangeFieldList | fieldList | ``` { "id":"widgetWorkGroupTimeRangeFieldList", "type":"fieldList", "value":[ [ { "id":"widgetWorkGroupStartTime", "type":"date", "value":"2019-10-01T08:12:01+08:00" }, { "id":"widgetWorkGroupEndTime", "type":"date", "value":"2019-10-01T08:12:01+08:00" } ] ] } ``` | Overtime period. For the specific format, please refer to the detail control. If "Allow submission of multiple overtime periods" is turned on in the definition, you can submit multiple, up to 30. Otherwise, only the first one will be used. The single overtime period cannot exceed two days |
| widgetWorkGroupReason | textarea | ``` { "id": "widgetWorkGroupReason", "type": "textarea", "value": "111" } ``` | Overtime reason. If the definition configures "Overtime reason" as mandatory, this field must be filled in | **Special parameter verification error information**
message                                                                            | description                                                                                                |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| the time range list has more than 30 items                                         | The number of overtime hours exceeds 30                                                                    |
| group value is invalid                                                             | The current control group value is invalid. Please verify that it is empty or that the type is an array    |
| overtime type is required                                                          | If no overtime rules are associated, the overtime type is required                                         |
| work time range is required                                                        | At least one overtime period is required                                                                   |
| start time is after end time                                                       | The start time must be later than the end time                                                             |
| start time or end time of range is required                                        | Both the start and end times of the overtime period are required                                           |
| overtime duration is over 2 days                                                   | Overtime cannot exceed two days at a time                                                                  |
| overtime date time zone not support                                                | The date and time zone for the overtime period cannot be recognized                                        |
| {date} can not apply overtime                                                      | Overtime cannot be requested for the selected time                                                         |
| {date} already apply overtime                                                      | Overtime has already been recorded for the selected time                                                   |
| {date} no need approval                                                            | Overtime does not need to be requested for the selected date                                               |
| apply reason is required                                                           | The reason for overtime is set as required in the definition and cannot be left blank                      |
| {users} user follow different overtime rules, cannot be submitted in the same form | The selected overtime users are not in the same attendance group and cannot submit overtime simultaneously |
| invalid overtime work application                                                  | There is no valid overtime application. Please reselect the overtime date                                  |
| the overtime duration cannot be 0                                                  | Overtime duration cannot be 0                                                                              |
| the number of apply workers cannot exceed 50                                       | The number of overtime users requested at a time cannot exceed 50. An overtime user must be specified      |
| apply worker is required                                                           | If you configure the overtime user to submit on behalf of multiple users, you must specify the user        |
| resigned worker can not apply                                                      | Resigned employees cannot apply for overtime                                                               |
| overtime duration is over limit                                                    | Overtime hours exceed the limit

### out of office widget group

**For the widget type `outWidgetGroup`, here is an example of JSON data:**
```
{
    "id": "widgetOutGroup",
    "type": "outGroup",
    "value":[
        {
            "id": "widgetOutGroupType",
            "type": "radioV2",
            "value":  "me15yqrf-gmjgbml2vhp-0"      
        },
        {
            "id": "widgetOutGroupStartTime",
            "type": "date",
            "value":"2019-10-01T08:12:01+08:00"
        },
        {
            "id": "widgetOutGroupEndTime",
            "type": "date",
            "value":"2019-10-01T08:12:01+08:00"
        },
        {
            "id": "widgetOutGroupReason",
            "type": "textarea",
            "value":"123213"
        },
        {
            "id":"widgetOutGroupImage",
            "type":"image",
            "value": ["D93653C3-2609-4EE0-8041-61DC1D84F0B5"]
        }                    
    ]   
}

```

**The parameter description of out of office widget group：**
| parameter | type | required | description |
| --- | --- | --- | --- |
| id | string | yes | group ID, fixed as widgetOutGroup |
| type | string | yes | group type, fixed asoutGroup |
| value | object[] | yes | The value of the space group, which is a list of multiple sub-control values | Description of the sub-widget values contained in value:
| id | type | JSON sample | description |
| --- | --- | --- | --- |
| widgetOutGroupType | radioV2 | ``` { "id": "widgetOutGroupType", "type": "radioV2", "value":  "me15yqrf-gmjgbml2vhp-0" } ``` | Out of office type. For the specific format, please refer to the radio button. If "Out of office type" is configured, it is required. The unit associated with the selected out of office type will be selected for out of office duration. If "Out of office type" is not configured, this field does not need to be filled in. The unit configured for "Out of office duration" will be selected when calculating out of office duration. |
| widgetOutGroupStartTime | date | ``` { "id": "widgetOutGroupStartTime", "type": "date", "value":"2019-10-01T08:12:01+08:00" } ``` | The start time of going out. For the specific format, please refer to the date control. If the unit of going out is half a day, it is considered to be in the morning if it is less than 12 o'clock, otherwise it is considered to be in the afternoon. If the unit is hour, it will be rounded up to half an hour. |
| widgetOutGroupEndTime | date | ``` { "id": "widgetOutGroupEndTime", "type": "date", "value":"2019-10-01T08:12:01+08:00" } ``` | End time of absence. For the specific format, please refer to the date control. If the absence time unit is half a day, it is considered to be in the morning if it is less than 12 o'clock, otherwise it is considered to be in the afternoon. If the unit is hour, it will be rounded up to half an hour |
| widgetOutGroupReason | textarea | ``` { "id": "widgetOutGroupReason", "type": "textarea", "value":"123213" } ``` | Reason for going out. For the specific format, please refer to the multi-line text control. If the definition of "reason for going out" is required, then this control must be filled in. If the definition configuration does not require filling in, then this control does not need to be filled in |
| widgetOutGroupImage | image | ``` { "id":"widgetOutGroupImage", "type":"image", "value": ["D93653C3-2609-4EE0-8041-61DC1D84F0B5"] } ``` | The certificate of out. The specific format of the proof of going out can refer to the picture control. If the definition of "taking photos when going out" is required, then this control must be filled in. If the definition configuration does not require filling in, then this control does not need to be filled in. | **Special parameter verification error information**
message                                               | description                                                                                                    |
| ----------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| group value is invalid                                | The current control group value is invalid. Please verify that it is empty or that the type is an array        |
| start time format is not RFC3339                      | The start time and date format is not in RFC3339 format                                                        |
| end time format is not RFC3339                        | The end time and date format is not in RFC3339 format                                                          |
| start time and end time must be in the same time zone | The start and end times must be in the same time zone                                                          |
| out type is required                                  | If "Outgoing Type" is set in the definition, the Outgoing Type is required                                     |
| out start time is required                            | Outing start time is required                                                                                  |
| out end time is required                              | Out of office end time is required                                                                             |
| out duration must be greater than 0                   | The outing interval cannot be 0. Please check the start and end times and reselect                             |
| out reason is empty                                   | This field is required if "Outing Reason" is selected and set as required in the definition                    |
| photo is required                                     | This field is required if "Outing Reason" is selected and set as required in the definition                    |
| out time is conflict                                  | If there is a conflict in the outing time, please confirm whether you have applied for outing during that time
