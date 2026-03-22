---
title: "Search records"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-record/search"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/bitable/v1/apps/:app_token/tables/:table_id/records/search"
service: "bitable-v1"
resource: "app-table-record"
section: "Docs"
rate_limit: "20 per second"
scopes:
  - "base:record:retrieve"
  - "bitable:app"
  - "bitable:app:readonly"
field_scopes:
  - "contact:user.employee_id:readonly"
  - "contact:user.base:readonly"
  - "contact:contact:access_as_app"
  - "contact:contact:readonly"
  - "contact:contact:readonly_as_app"
updated: "1773323348000"
---

# Search records

This api is used to query existing records in the  table. A maximum of 500 rows of records can be queried at a time, and paging is supported.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/bitable/v1/apps/:app_token/tables/:table_id/records/search |
| HTTP Method | POST |
| Rate Limit | 20 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `base:record:retrieve` `bitable:app` `bitable:app:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` `contact:user.base:readonly` `contact:contact:access_as_app` `contact:contact:readonly` `contact:contact:readonly_as_app` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `app_token` | `string` | The unique identifier of the Bitable App. Different forms of Bitable have different ways of obtaining app_token: - If a Bitable URL begins with **larksuite.com/base**, the app_token of the Base is the string that follows **base**. - If Bitable's URL starts with **larksuite.com/base**, you need to call the Knowledge Base related Get Wiki Workspace Node Information interface to get Bitable's app_token. When the value of obj_type is bitable, the value of the obj_token field is the app_token of Bitable. **Example value**: "NQRxbRkBMa6OnZsjtERcxhNWnNh" **Data validation rules**: - Length range: `0` ～ `100` characters |
| `table_id` | `string` | Bitable data table unique identifier. How to get: - You can get the table_id through the Bitable URL. The highlighted part of the image below is the table_id of the current data table. ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/d4750d909d5a27493f6f5fe633d7cc95_CkBmQkbGch.png?height=765&lazyload=true&maxWidth=300&width=1731) - table_id can also be obtained through the list data table interface **Example value**: "tbl0xe5g8PP3U3cS" **Data validation rules**: - Length range: `0` ～ `50` characters | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: eVQrYzJBNDNONlk4VFZBZVlSdzlKdFJ4bVVHVExENDNKVHoxaVdiVnViQT0= |
| `page_size` | `int` | No | Page size. The maximum value is 500 **Example value**: 10 **Default value**: `20` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `view_id` | `string` | No | Base view unique device identifier view_id description **Example value**: "vewp7nmiS4" **Data validation rules**: - Length range: `0` ～ `50` characters |
| `field_names` | `string[]` | No | field_names **Example value**: ["fieldName1"] **Data validation rules**: - Length range: `0` ～ `200` |
| `sort` | `sort[]` | No | Data validation rules **Data validation rules**: - Length range: `0` ～ `100` |
|   `field_name` | `string` | No | field_name **Example value**: "fieldName1" **Data validation rules**: - Length range: `0` ～ `1000` characters |
|   `desc` | `boolean` | No | desc **Example value**: true **Default value**: `false` |
| `filter` | `filter_info` | No | Refer to the Record Filter Parameter Filling Guide to learn how to fill in the filter. |
|   `conjunction` | `string` | No | Conjunction. This parameter is required and please ignore the "No" in the Required column **Example value**: "and" **Optional values are**:  - `and`: meeting all of the conditions - `or`: meeting any of the conditions  **Data validation rules**: - Length range: `0` ～ `10` characters |
|   `conditions` | `condition[]` | No | conditions **Data validation rules**: - Length range: `0` ～ `50` |
|     `field_name` | `string` | Yes | field_name **Example value**: "fieldName1" **Data validation rules**: - Length range: `0` ～ `1000` characters |
|     `operator` | `string` | Yes | operator **Example value**: "is" **Optional values are**:  - `is`: is - `isNot`: is not. This value does not support date fields. To learn how to query date fields, refer to Date Field Entry Guide. - `contains`: contains. This value does not support date fields. - `doesNotContain`: does not contain. This value does not support date fields. - `isEmpty`: is empty - `isNotEmpty`: is not empty - `isGreater`: greater than. - `isGreaterEqual`: greater than or equal to. This value does not support date fields. - `isLess`: less than - `isLessEqual`: less than or equal to. This value does not support date fields. - `like`: LIKE operator. Not supported yet - `in`: IN operator. Not supported yet  |
|     `value` | `string[]` | No | value Record filter development guide **Example value**: ["text content"] **Data validation rules**: - Length range: `0` ～ `10` |
| `automatic_fields` | `boolean` | No | Whether to automatically calculate and return the four types of fields: creation time (`created_time`), modification time (`last_modified_time`), creator (`created_by`), and modifier (`last_modified_by`). The default is false, indicating they will not be returned. **Example value**: false | ### Request body example

{
  "view_id": "vewqhz51lk",
  "field_names": [
    "字段1",
    "字段2"
  ],
  "sort": [
    {
      "field_name": "多行文本",
      "desc": true
    }
  ],
  "filter": {
    "conjunction": "and",
    "children": [
      {
        "conjunction": "or",
        "conditions": [
          {
            "field_name": "职位",
            "operator": "is",
            "value": [
              "高级销售员"
            ]
          },
          {
            "field_name": "职位",
            "operator": "is",
            "value": [
              "初级销售员"
            ]
          }
        ]
      },
      {
        "conjunction": "or",
        "conditions": [
          {
            "field_name": "销售额",
            "operator": "is",
            "value": [
              "10000.0"
            ]
          },
          {
            "field_name": "销售额",
            "operator": "is",
            "value": [
              "20000.0"
            ]
          }
        ]
      }
    ]
  },
  "automatic_fields": false
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `items` | `app.table.record[]` | The list of records. The data type is array. For details, see Data structure. |
|     `fields` | `map<string, union>` | The record field |
|     `record_id` | `string` | record ID |
|     `created_by` | `person` | founder |
|       `id` | `string` | Person ID. Consistent with the type specified user_id_type query parameter. |
|       `name` | `string` | Chinese name |
|       `en_name` | `string` | English name |
|       `email` | `string` | email |
|       `avatar_url` | `string` | avatar link **Required field scopes (Satisfy any)**: `contact:user.base:readonly` `contact:contact:access_as_app` `contact:contact:readonly` `contact:contact:readonly_as_app` |
|     `created_time` | `int` | creation time |
|     `last_modified_by` | `person` | Modifier |
|       `id` | `string` | Person ID. Consistent with the type specified user_id_type query parameter. |
|       `name` | `string` | Chinese name |
|       `en_name` | `string` | English name |
|       `email` | `string` | email |
|       `avatar_url` | `string` | avatar link **Required field scopes (Satisfy any)**: `contact:user.base:readonly` `contact:contact:access_as_app` `contact:contact:readonly` `contact:contact:readonly_as_app` |
|     `last_modified_time` | `int` | Last update time |
|     `shared_url` | `string` | Record sharing link (the bulk fetch records interface will return this field) |
|     `record_url` | `string` | Record link (the retrieve record interface will return this field) |
|   `has_more` | `boolean` | Whether the response body has more parameters |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned |
|   `total` | `int` | total | ### Response body example

{
    "code":0,
    "data":{
        "has_more":false,
        "items":[
            {
                "created_by":{
                    "avatar_url":"https://internal-api-lark-file.larksuite.com/static-resource/v1/06d568cb-f464-4c2e-bd03-76512c545c5j~?image_size=72x72&cut_type=default-face&quality=&format=jpeg&sticker_format=.webp",
                    "email":"",
                    "en_name":"测试1",
                    "id":"ou_92945f86a98bba075174776959c90eda",
                    "name":"测试1"
                },
                "created_time":1691049973000,
                "fields":{
                    "人员":[
                        {
                            "avatar_url":"https://internal-api-lark-file.larksuite.com/static-resource/v1/b2-7619-4b8a-b27b-c72d90b06a2j~?image_size=72x72&cut_type=default-face&quality=&format=jpeg&sticker_format=.webp",
                            "email":"zhangsan.leben@bytedance.com",
                            "en_name":"ZhangSan",
                            "id":"ou_2910013f1e6456f16a0ce75ede950a0a",
                            "name":"张三"
                        },
                        {
                            "avatar_url":"https://internal-api-lark-file.larksuite.com/static-resource/v1/v2_q86-fcb6-4f18-85c7-87ca8881e50j~?image_size=72x72&cut_type=default-face&quality=&format=jpeg&sticker_format=.webp",
                            "email":"lisi.00@bytedance.com",
                            "en_name":"LiSi",
                            "id":"ou_e04138c9633dd0d2ea166d79f548ab5d",
                            "name":"李四"
                        }
                    ],
                    "修改人":[
                        {
                            "avatar_url":"https://internal-api-lark-file.larksuite.com/static-resource/v1/06d568cb-f464-4c2e-bd03-76512c545c5j~?image_size=72x72&cut_type=default-face&quality=&format=jpeg&sticker_format=.webp",
                            "email":"",
                            "en_name":"测试1",
                            "id":"ou_92945f86a98bba075174776959c90eda",
                            "name":"测试1"
                        }
                    ],
                    "创建人":[
                        {
                            "avatar_url":"https://internal-api-lark-file.larksuite.com/static-resource/v1/06d568cb-f464-4c2e-bd03-76512c545c5j~?image_size=72x72&cut_type=default-face&quality=&format=jpeg&sticker_format=.webp",
                            "email":"",
                            "en_name":"测试1",
                            "id":"ou_92945f86a98bba075174776959c90eda",
                            "name":"测试1"
                        }
                    ],
                    "创建时间":1691049973000,
                    "单向关联":{
                        "link_record_ids":[
                            "recnVYsuqV"
                        ]
                    },
                    "单选":"选项1",
                    "双向关联":{
                        "link_record_ids":[
                            "recqLvMaXT",
                            "recrdld32q"
                        ]
                    },
                    "地理位置":{
                        "address":"东长安街",
                        "adname":"东城区",
                        "cityname":"北京市",
                        "full_address":"天安门广场，北京市东城区东长安街",
                        "location":"116.397755,39.903179",
                        "name":"天安门广场",
                        "pname":"北京市"
                    },
                    "复选框":true,
                    "多行文本":[
                        {
                            "text":"多行文本内容1",
                            "type":"text"
                        },
                        {
                            "mentionNotify":false,
                            "mentionType":"User",
                            "name":"张三",
                            "text":"@张三",
                            "token":"ou_2910013f1e6456f16a0ce75ede950a0a",
                            "type":"mention"
                        }
                    ],
                    "多选":[
                        "选项1",
                        "选项2"
                    ],
                    "数字":2323.2323,
                    "日期":1690992000000,
                    "最后更新时间":1702455191000,
                    "条码":[
                        {
                            "text":"123",
                            "type":"text"
                        }
                    ],
                    "电话号码":"131xxxx6666",
                    "自动编号":"17",
                    "群组":[
                        {
                            "avatar_url":"https://internal-api-lark-file.larksuite.com/static-resource/v1/v2_c8d2cd50-ba29-476f-b7f1-5b5917cb18ej~?image_size=72x72&cut_type=&quality=&format=jpeg&sticker_format=.webp",
                            "id":"oc_cd07f55f14d6f4a4f1b51504e7e97f48",
                            "name":"武侠聊天组"
                        }
                    ],
                    "评分":3,
                    "货币":1,
                    "超链接":{
                        "link":"https://bitable.larksuite.com",
                        "text":"Lark多维表格官网"
                    },
                    "进度":0.66,
                    "附件":[
                        {
                            "file_token":"Vl3FbVkvnowlgpxpqsAbBrtFcrd",
                            "name":"Lark.jpeg",
                            "size":32975,
                            "tmp_url":"https://open.larksuite.com/open-apis/drive/v1/medias/batch_get_tmp_download_url?file_tokens=Vl3FbVk11owlgpxpqsAbBrtFcrd&extra=%7B%22bitablePerm%22%3A%7B%22tableId%22%3A%22tblBJyX6jZteblYv%22%2C%22rev%22%3A90%7D%7D",
                            "type":"image/jpeg",
                            "url":"https://open.larksuite.com/open-apis/drive/v1/medias/Vl3FbVk11owlgpxpqsAbBrtFcrd/download?extra=%7B%22bitablePerm%22%3A%7B%22tableId%22%3A%22tblBJyX6jZteblYv%22%2C%22rev%22%3A90%7D%7D"
                        }
                    ]
                },
                "last_modified_by":{
                    "avatar_url":"https://internal-api-lark-file.larksuite.com/static-resource/v1/06d568cb-f464-4c2e-bd03-76512c545c5j~?image_size=72x72&cut_type=default-face&quality=&format=jpeg&sticker_format=.webp",
                    "email":"",
                    "en_name":"测试1",
                    "id":"ou_92945f86a98bba075174776959c90eda",
                    "name":"测试1"
                },
                "last_modified_time":1702455191000,
                "record_id":"recyOaMB2F"
            }
        ],
        "total":1
    },
    "msg":"success"
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 200 | 1254000 | WrongRequestJson | Request error |
| 200 | 1254001 | WrongRequestBody | Request body error |
| 200 | 1254002 | Fail | Internal error, have any questions can be consulting service |
| 200 | 1254003 | WrongBaseToken | AppToken error |
| 200 | 1254004 | WrongTableId | Table id wrong |
| 200 | 1254005 | WrongViewId | View id wrong |
| 200 | 1254006 | WrongRecordId | Record id wrong |
| 200 | 1254007 | EmptyValue | Empty value |
| 200 | 1254008 | EmptyView | Empty view |
| 200 | 1254009 | WrongFieldId | Wrong fieldId |
| 200 | 1254010 | ReqConvError | Request error |
| 400 | 1254011 | Page size must greater than 0. | invalid page_size |
| 200 | 1254016 | InvalidSort | invalid sort |
| 200 | 1254018 | InvalidFilter | The filter parameter is incorrect. Please refer to Record filter development guide for information on how to fill in the filter parameter. |
| 200 | 1254024 | InvalidFieldNames | InvalidFieldNames |
| 200 | 1254030 | TooLargeResponse | TooLargeResponse |
| 400 | 1254036 | Base is copying, please try again later. | Base copy replicating, try again later |
| 200 | 1254040 | BaseTokenNotFound | AppToken not found |
| 200 | 1254041 | TableIdNotFound | Table not found |
| 200 | 1254042 | ViewIdNotFound | View not found |
| 200 | 1254043 | RecordIdNotFound | RecordIdNotFound |
| 200 | 1254044 | FieldIdNotFound | FieldIdNotFound |
| 200 | 1254045 | FieldNameNotFound | Field name does not exist |
| 200 | 1254060 | TextFieldConvFail | TextFieldConvFail |
| 200 | 1254061 | NumberFieldConvFail | NumberFieldConvFail |
| 200 | 1254062 | SingleSelectFieldConvFail | SingleSelectFieldConvFail |
| 200 | 1254063 | MultiSelectFieldConvFail | MultiSelectFieldConvFail |
| 200 | 1254064 | DatetimeFieldConvFail | DatetimeFieldConvFail |
| 200 | 1254065 | CheckboxFieldConvFail | CheckboxFieldConvFail |
| 200 | 1254066 | UserFieldConvFail | The value corresponding to the personnel field type is incorrect. The possible reasons are: - The ID type specified by the user_id_type parameter does not match the type of the provided ID. - An unrecognized type or structure was provided. Currently, only `id` is supported, and it must be passed as an array. - An `open_id` was passed across applications. If you are passing an ID across applications, it is recommended to use `user_id`. The `open_id` obtained from different applications cannot be used interchangeably. |
| 200 | 1254067 | LinkFieldConvFail | LinkFieldConvFail |
| 200 | 1254068 | URLFieldConvFail | URLFieldConvFail |
| 200 | 1254069 | AttachFieldConvFail | AttachFieldConvFail |
| 200 | 1254072 | Failed to convert phone field, please make sure it is correct. | invalid phone format |
| 200 | 1254100 | TableExceedLimit | TableExceedLimit, limited to 300 |
| 200 | 1254101 | ViewExceedLimit | ViewExceedLimit, limited to 200 |
| 200 | 1254102 | FileExceedLimit | FileExceedLimit |
| 200 | 1254103 | RecordExceedLimit | RecordExceedLimit, limited to 20,000 |
| 200 | 1254104 | RecordAddOnceExceedLimit | RecordAddOnceExceedLimit, limited to 500 |
| 200 | 1254107 | FilterLengthExceedLimit | FilterLengthExceedLimit, limited to 2,000 characters |
| 200 | 1254108 | SortLengthExceedLimit | SortLengthExceedLimit, limited to 1,000 characters |
| 200 | 1254109 | FormulaTableSizeExceedLimit | FormulaTableSizeExceedLimit |
| 200 | 1254130 | TooLargeCell | TooLargeCell |
| 200 | 1254290 | TooManyRequest | Request too fast, try again later |
| 200 | 1254291 | Write conflict | The same data table does not support concurrent calls to the write interface, please check whether there is a concurrent call to the write interface. The writing interface includes: adding, modifying, and deleting records; adding, modifying, and deleting fields; modifying forms; modifying views, etc. |
| 200 | 1254301 | OperationTypeError | Base does not have advanced permissions enabled or does not support enabling advanced permissions |
| 200 | 1254302 | Permission denied. | No access rights, usually caused by the table opening of advanced permissions, please add a group containing applications in the advanced permissions settings and give this group read and write permissions |
| 200 | 1254303 | AttachPermNotAllow | No attach permission |
| 200 | 1255001 | InternalError | Internal error, have any questions can be consulting service |
| 200 | 1255002 | RpcError | Internal error, have any questions can be consulting service |
| 200 | 1255003 | MarshalError | Serialization failed, have any questions can be consulting service |
| 200 | 1255004 | UmMarshalError | Deserialization failed, have any questions can be consulting service |
| 200 | 1255005 | ConvError | Internal error, have any questions can be consulting service |
| 504 | 1255040 | Request timed out, please try again later | Try again |
| 400 | 1254607 | Data not ready, please try again later | There are usually two situations when this error occurs: 1. The last submitted modification has not been processed; 2. The data is too large and the server calculation times out; This error code can be appropriately retried. |
