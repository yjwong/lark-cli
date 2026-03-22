---
title: "Batch get records"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-record/batch_get"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/bitable/v1/apps/:app_token/tables/:table_id/records/batch_get"
service: "bitable-v1"
resource: "app-table-record"
section: "Docs"
rate_limit: "20 per second"
scopes:
  - "base:record:retrieve"
  - "bitable:app"
  - "bitable:app:readonly"
field_scopes:
  - "contact:user.base:readonly"
  - "contact:contact:readonly_as_app"
updated: "1773323355000"
---

# Batch get records

Batch get records by record ids

> Only supports querying up to 100 records

## Precautions

If advanced permissions are enabled for Base, you need to ensure that the calling identity has manage permissions for Base. Otherwise, it may result in a successful call but return empty data. For detailed steps, refer to How to Grant Document Permissions for Apps or Users.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/bitable/v1/apps/:app_token/tables/:table_id/records/batch_get |
| HTTP Method | POST |
| Rate Limit | 20 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `base:record:retrieve` `bitable:app` `bitable:app:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.base:readonly` `contact:contact:readonly_as_app` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `app_token` | `string` | Form token **Example value**: "NQRxbRkBMa6OnZsjtERcxhNWnNh" **Data validation rules**: - Length range: `0` ～ `100` characters |
| `table_id` | `string` | Form ID **Example value**: "tbl0xe5g8PP3U3cS" **Data validation rules**: - Length range: `0` ～ `50` characters | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `record_ids` | `string[]` | Yes | record id list. See Query records **Example value**: ["recxq2GJwE"] **Data validation rules**: - Length range: `1` ～ `100` |
| `user_id_type` | `string` | No | **Example value**: "open_id" **Optional values are**:  - `user_id`: - `union_id`: - `open_id`:  |
| `with_shared_url` | `boolean` | No | Controls whether to return the recorded share link, true means return the share link **Example value**: True |
| `automatic_fields` | `boolean` | No | Controls whether to return automatically calculated fields, true means return **Example value**: True | ### Request body example

{
  "record_ids": [
    "recyOaMB2F",
    "rec111111",
    "recyOaMB2F"
  ],
  "user_id_type": "open_id",
  "with_shared_url": true,
  "automatic_fields": true
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `records` | `app.table.record[]` | record list |
|     `fields` | `map<string, union>` | record field |
|     `record_id` | `string` | Record Id |
|     `created_by` | `person` | founder |
|       `id` | `string` | Personnel IDs |
|       `name` | `string` | Chinese name |
|       `en_name` | `string` | English name |
|       `email` | `string` | email |
|       `avatar_url` | `string` | avatar link **Required field scopes (Satisfy any)**: `contact:user.base:readonly` `contact:contact:readonly_as_app` |
|     `created_time` | `int` | creation time |
|     `last_modified_by` | `person` | Modifier |
|       `id` | `string` | Personnel IDs |
|       `name` | `string` | Chinese name |
|       `en_name` | `string` | English name |
|       `email` | `string` | email |
|       `avatar_url` | `string` | avatar link **Required field scopes (Satisfy any)**: `contact:user.base:readonly` `contact:contact:readonly_as_app` |
|     `last_modified_time` | `int` | Last update time |
|   `forbidden_record_ids` | `string[]` | List of prohibited records (for documents with advanced permissions enabled) |
|   `absent_record_ids` | `string[]` | List of non-existent records | ### Response body example

{
  "code": 0,
  "msg": "success",
  "data": {
    "forbidden_record_ids": [
      "recyOaMB2F"
    ],
    "absent_record_ids": [
      "rec111111"
    ],
    "records": [
      {
        "created_by": {
          "avatar_url": "https://internal-api-lark-file.larksuite.com/static-resource/v1/06d568cb-f464-4c2e-bd03-76512c545c5j~?image_size=72x72&cut_type=default-face&quality=&format=jpeg&sticker_format=.webp",
          "email": "",
          "en_name": "测试1",
          "id": "ou_92945f86a98bba075174776959c90eda",
          "name": "测试1"
        },
        "created_time": 1691049973000,
        "fields": {
          "人员": [
            {
              "avatar_url": "https://internal-api-lark-file.larksuite.com/static-resource/v1/b2-7619-4b8a-b27b-c72d90b06a2j~?image_size=72x72&cut_type=default-face&quality=&format=jpeg&sticker_format=.webp",
              "email": "zhangsan.leben@bytedance.com",
              "en_name": "ZhangSan",
              "id": "ou_2910013f1e6456f16a0ce75ede950a0a",
              "name": "张三"
            },
            {
              "avatar_url": "https://internal-api-lark-file.larksuite.com/static-resource/v1/v2_q86-fcb6-4f18-85c7-87ca8881e50j~?image_size=72x72&cut_type=default-face&quality=&format=jpeg&sticker_format=.webp",
              "email": "lisi.00@bytedance.com",
              "en_name": "LiSi",
              "id": "ou_e04138c9633dd0d2ea166d79f548ab5d",
              "name": "李四"
            }
          ],
          "修改人": [
            {
              "avatar_url": "https://internal-api-lark-file.larksuite.com/static-resource/v1/06d568cb-f464-4c2e-bd03-76512c545c5j~?image_size=72x72&cut_type=default-face&quality=&format=jpeg&sticker_format=.webp",
              "email": "",
              "en_name": "测试1",
              "id": "ou_92945f86a98bba075174776959c90eda",
              "name": "测试1"
            }
          ],
          "创建人": [
            {
              "avatar_url": "https://internal-api-lark-file.larksuite.com/static-resource/v1/06d568cb-f464-4c2e-bd03-76512c545c5j~?image_size=72x72&cut_type=default-face&quality=&format=jpeg&sticker_format=.webp",
              "email": "",
              "en_name": "测试1",
              "id": "ou_92945f86a98bba075174776959c90eda",
              "name": "测试1"
            }
          ],
          "创建时间": 1691049973000,
          "单向关联": {
            "link_record_ids": [
              "recnVYsuqV"
            ]
          },
          "单选": "选项1",
          "双向关联": {
            "link_record_ids": [
              "recqLvMaXT",
              "recrdld32q"
            ]
          },
          "地理位置": {
            "address": "东长安街",
            "adname": "东城区",
            "cityname": "北京市",
            "full_address": "天安门广场，北京市东城区东长安街",
            "location": "116.397755,39.903179",
            "name": "天安门广场",
            "pname": "北京市"
          },
          "复选框": true,
          "多行文本": [
            {
              "text": "多行文本内容1",
              "type": "text"
            },
            {
              "mentionNotify": false,
              "mentionType": "User",
              "name": "张三",
              "text": "@张三",
              "token": "ou_2910013f1e6456f16a0ce75ede950a0a",
              "type": "mention"
            }
          ],
          "多选": [
            "选项1",
            "选项2"
          ],
          "数字": 2323.2323,
          "日期": 1690992000000,
          "最后更新时间": 1702455191000,
          "条码": [
            {
              "text": "123",
              "type": "text"
            }
          ],
          "电话号码": "131xxxx6666",
          "自动编号": "17",
          "群组": [
            {
              "avatar_url": "https://internal-api-lark-file.com/static-resource/v1/v2_c8d2cd50-ba29-476f-b7f1-5b5917cb18ej~?image_size=72x72&cut_type=&quality=&format=jpeg&sticker_format=.webp",
              "id": "oc_cd07f55f14d6f4a4f1b51504e7e97f48",
              "name": "武侠聊天组"
            }
          ],
          "评分": 3,
          "货币": 1,
          "超链接": {
            "link": "https://bitable.larksuite.com",
            "text": "Lark多维表格官网"
          },
          "进度": 0.66,
          "附件": [
            {
              "file_token": "Vl3FbVkvnowlgpxpqsAbBrtFcrd",
              "name": "Lark.jpeg",
              "size": 32975,
              "tmp_url": "https://open.larksuite.com/open-apis/drive/v1/medias/batch_get_tmp_download_url?file_tokens=Vl3FbVk11owlgpxpqsAbBrtFcrd&extra=%7B%22bitablePerm%22%3A%7B%22tableId%22%3A%22tblBJyX6jZteblYv%22%2C%22rev%22%3A90%7D%7D",
              "type": "image/jpeg",
              "url": "https://open.larksuite.com/open-apis/drive/v1/medias/Vl3FbVk11owlgpxpqsAbBrtFcrd/download?extra=%7B%22bitablePerm%22%3A%7B%22tableId%22%3A%22tblBJyX6jZteblYv%22%2C%22rev%22%3A90%7D%7D"
            }
          ]
        },
        "last_modified_by": {
          "avatar_url": "https://internal-api-lark-file.larksuite.com/static-resource/v1/06d568cb-f464-4c2e-bd03-76512c545c5j~?image_size=72x72&cut_type=default-face&quality=&format=jpeg&sticker_format=.webp",
          "email": "",
          "en_name": "测试1",
          "id": "ou_92945f86a98bba075174776959c90eda",
          "name": "测试1"
        },
        "last_modified_time": 1702455191000,
        "record_id": "recyOaMB2F",
        "shared_url": "https://example.larksuite.com/record/KBcNrNtpWePAlscCvdmb6ZcSc5b"
      }
    ]
  }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1254000 | WrongRequestJson | Request error |
| 400 | 1254001 | WrongRequestBody | Request body error |
| 500 | 1254002 | Fail | Internal error, have any questions can be consulting service |
| 400 | 1254003 | WrongBaseToken | AppToken error |
| 400 | 1254004 | WrongTableId | Table id wrong |
| 400 | 1254005 | WrongViewId | View id wrong |
| 400 | 1254006 | WrongRecordId | Record id wrong |
| 400 | 1254007 | EmptyValue | Empty value |
| 400 | 1254008 | EmptyView | Empty view |
| 400 | 1254009 | WrongFieldId | Wrong fieldId |
| 400 | 1254010 | ReqConvError | Request error |
| 400 | 1254011 | Page size must greater than 0. | invalid page_size |
| 400 | 1254016 | InvalidSort | invalid sort |
| 400 | 1254018 | InvalidFilter | invalid filter |
| 400 | 1254024 | InvalidFieldNames | InvalidFieldNames |
| 400 | 1254030 | InvalidPageToken | Invalid PageToken |
| 400 | 1254036 | Bitable is copying, please try again later. | Base copy replicating, try again later |
| 404 | 1254040 | BaseTokenNotFound | AppToken not found |
| 404 | 1254041 | TableIdNotFound | Table not found |
| 404 | 1254042 | ViewIdNotFound | View not found |
| 404 | 1254043 | RecordIdNotFound | RecordIdNotFound |
| 404 | 1254044 | FieldIdNotFound | FieldIdNotFound |
| 404 | 1254045 | FieldNameNotFound | Field name does not exist |
| 500 | 1254060 | TextFieldConvFail | TextFieldConvFail |
| 500 | 1254061 | NumberFieldConvFail | NumberFieldConvFail |
| 500 | 1254062 | SingleSelectFieldConvFail | SingleSelectFieldConvFail |
| 500 | 1254063 | MultiSelectFieldConvFail | MultiSelectFieldConvFail |
| 500 | 1254064 | DatetimeFieldConvFail | DatetimeFieldConvFail |
| 500 | 1254065 | CheckboxFieldConvFail | CheckboxFieldConvFail |
| 500 | 1254066 | UserFieldConvFail | The value corresponding to the personnel field type is incorrect. The possible reasons are: - The ID type specified by the user_id_type parameter does not match the type of the provided ID. - An unrecognized type or structure was provided. Currently, only `id` is supported, and it must be passed as an array. - An `open_id` was passed across applications. If you are passing an ID across applications, it is recommended to use `user_id`. The `open_id` obtained from different applications cannot be used interchangeably. |
| 500 | 1254067 | LinkFieldConvFail | LinkFieldConvFail |
| 500 | 1254068 | URLFieldConvFail | URLFieldConvFail |
| 500 | 1254069 | AttachFieldConvFail | AttachFieldConvFail |
| 400 | 1254072 | Failed to convert phone field, please make sure it is correct. | invalid phone format |
| 400 | 1254100 | TableExceedLimit | TableExceedLimit, limited to 300 |
| 400 | 1254101 | ViewExceedLimit | ViewExceedLimit, limited to 200 |
| 400 | 1254102 | FileExceedLimit | FileExceedLimit |
| 400 | 1254103 | RecordExceedLimit | RecordExceedLimit, limited to 20,000 |
| 400 | 1254104 | RecordAddOnceExceedLimit | RecordAddOnceExceedLimit, limited to 500 |
| 400 | 1254107 | FilterLengthExceedLimit | FilterLengthExceedLimit, limited to 2,000 characters |
| 400 | 1254108 | SortLengthExceedLimit | SortLengthExceedLimit, limited to 1,000 characters |
| 400 | 1254109 | FormulaTableSizeExceedLimit | FormulaTableSizeExceedLimit |
| 400 | 1254130 | TooLargeCell | TooLargeCell |
| 400 | 1254290 | TooManyRequest | Request too fast, try again later |
| 400 | 1254291 | LockNotObtainedError | The same data table does not support concurrent calls to the write interface, please check whether there is a concurrent call to the write interface. The writing interface includes: adding, modifying, and deleting records; adding, modifying, and deleting fields; modifying forms; modifying views, etc. |
| 400 | 1254301 | OperationTypeError | Base does not have advanced permissions enabled or does not support enabling advanced permissions |
| 400 | 1254302 | RolePermNotAllow | No access rights, usually caused by the table opening of advanced permissions, please add a group containing applications in the advanced permissions settings and give this group read and write permissions |
| 400 | 1254303 | AttachPermNotAllow | No attach permission |
| 500 | 1255001 | InternalError | Internal error, have any questions can be consulting service |
| 500 | 1255002 | Something went wrong. | Internal error, have any questions can be consulting service |
| 500 | 1255003 | MarshalError | Serialization failed, have any questions can be consulting service |
| 500 | 1255004 | UmMarshalError | Deserialization failed, have any questions can be consulting service |
| 500 | 1255005 | ConvError | Internal error, have any questions can be consulting service |
| 504 | 1255040 | Request timed out, please try again later | Try again |
| 500 | 1254607 | Data not ready, please try again later | There are usually two situations when this error occurs: 1. The last submitted modification has not been processed; 2. The data is too large and the server calculation times out;This error code can be appropriately retried. |
