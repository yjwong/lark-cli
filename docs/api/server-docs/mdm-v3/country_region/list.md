---
title: "Pagination Batch Query Country Region"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/mdm-v3/country_region/list"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/mdm/v3/country_regions"
service: "mdm-v3"
resource: "country_region"
section: "Feishu Master Data Management"
rate_limit: "10 per second"
updated: "1756439316000"
---

# Paging batch query country region

Paging batch query country region

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/mdm/v3/country_regions |
| HTTP Method | GET |
| Rate Limit | 10 per second |
| Supported app types | custom |
| Required scopes | mdm:country_region:read | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `languages` | `string[]` | Yes | The language type you want to return, the supported format is as follows: -Chinese: zh-CN -English: en-US -Japanese: ja-JP For multilingual text fields, if a specific language is passed in, the corresponding language text will be returned. **Example value**: en-US **Data validation rules**: - Length range: `0` ～ `100` |
| `fields` | `string[]` | Yes | Required query field set **Example value**: name **Data validation rules**: - Length range: `1` ～ `100` |
| `limit` | `int` | No | query page size **Example value**: 10 **Data validation rules**: - Value range: `1` ～ `1000` |
| `offset` | `int` | No | query start location **Example value**: 0 **Data validation rules**: - Value range: `0` ～ `100000` |
| `return_count` | `boolean` | No | Whether to return the total **Example value**: true |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: eVQrYzJBNDNONlk4VFZBZVlSdzlKdFJ4bVVHVExENDNKVHoxaVdiVnViQT0= | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `filter` | `filter` | No | Filter parameters |
|   `logic` | `string` | Yes | Logic Multiple expressions at the same level are determined by the logic parameter using "and/or" conditions. 0=and, 1=or **Example value**: "0" |
|   `expressions` | `expression[]` | No | filter condition **Data validation rules**: - Length range: `0` ～ `100` |
|     `field` | `string` | Yes | field name **Example value**: "name" |
|     `operator` | `string` | Yes | Operator 0=equal, 1=not equal, 2=greater than, 3=greater than or equal to, 4=less than, 5=less than or equal to, 6=any, 7=not any, 8=match, 9=prefix match, 10=suffix Match, 11=null, 12=not null **Example value**: "0" |
|     `value` | `value` | Yes | field value |
|       `string_value` | `string` | No | string value **Example value**: "Andorra" |
|       `bool_value` | `boolean` | No | Boolean **Example value**: true |
|       `int_value` | `string` | No | shaping value **Example value**: "111" |
|       `string_list_value` | `string[]` | No | String list value **Example value**: ["1"] |
|       `int_list_value` | `string[]` | No | integer list value **Example value**: ["1"] |
| `common` | `common` | No | This parameter can be ignored | ### Request body example

{
    "filter": {
        "logic": "0",
        "expressions": [
            {
                "field": "name",
                "operator": "0",
                "value": {
                    "string_value": "Andorra",
                    "bool_value": true,
                    "int_value": "111",
                    "string_list_value": [
                        "1"
                    ],
                    "int_list_value": [
                        "1"
                    ]
                }
            }
        ]
    },
    "common": {}
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `data` | `country_region[]` | Country directory listing |
|     `alpha_3_code` | `string` | Three-letter code |
|     `alpha_2_code` | `string` | Two-letter code |
|     `numeric_code` | `string` | Numeric code |
|     `name` | `i18n_string` | name |
|       `value` | `string` | Enter the value corresponding to the language ranked first in the languages ​​parameter. |
|       `multilingual_value` | `map<string, string>` | Enter the values ​​corresponding to all languages ​​in the languages ​​parameter. |
|       `return_language` | `string` | The language corresponding to the value actually returned by value, such as "zh-CN" |
|     `mdm_code` | `string` | Master data encoding (unique permanent code generated by the system in the format "MDCT + 8 digits") |
|     `full_name` | `i18n_string` | Full name of country/region |
|       `value` | `string` | Enter the value corresponding to the language ranked first in the languages ​​parameter. |
|       `multilingual_value` | `map<string, string>` | Enter the values ​​corresponding to all languages ​​in the languages ​​parameter. |
|       `return_language` | `string` | The language corresponding to the value actually returned by value, such as "zh-CN" |
|     `global_code` | `string` | International phone area code |
|     `status` | `string` | Is it effective？ |
|     `continents` | `enum` | Continent. Optional values ​​are as follows 1-Asia, 2-Europe, 3-Africa, 4-North America, 5-South America, 6-Oceania, 7-Antarctica) |
|       `value` | `string` | Enter the value corresponding to the language ranked first in the languages ​​parameter. |
|       `multilingual_name` | `map<string, string>` | Enter the values ​​corresponding to all languages ​​in the languages ​​parameter. |
|   `total` | `string` | total |
|   `next_page_token` | `string` | next paging parameter | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "data": [
            {
                "alpha_3_code": "-",
                "alpha_2_code": "-",
                "numeric_code": "-",
                "name": {
                    "value": "zh-name",
                    "multilingual_value": {
                        "EN-EN": "Andorra"
                    },
                    "return_language": "zh-CN"
                },
                "mdm_code": "-",
                "full_name": {
                    "value": "zh-name",
                    "multilingual_value": {
                        "EN-CN": "Principality of Andorra"
                    },
                    "return_language": "zh-CN"
                },
                "global_code": "-",
                "status": "-",
                "continents": {
                    "value": "-",
                    "multilingual_name": {
                        "EN-EN": "Andorra"
                    }
                }
            }
        ],
        "total": "0",
        "next_page_token": "token"
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1640201 | invalid params | The parameter is illegal, please check the parameter |
| 400 | 1640202 | param not found | The parameter is empty, please check whether the parameter is filled in completely. |
| 400 | 1640203 | object field not found | This field is not in the query range, please check |
| 400 | 1640204 | language is invalid | Illegal language, please fill in again |
| 400 | 1640205 | params is required | Required parameters are not filled in, please check |
| 400 | 1640206 | limit too large | The number of inquiries is too large, please limit it to 1000. |
| 400 | 1640207 | limit plus offset too large | The number of queries and the starting position are too large, please limit it to 10,000 |
| 400 | 1640208 | page token invalid | Invalid paging code, please check this parameter |
| 400 | 1640209 | record id not found | No eligible row-level records were found, please check the query parameters |
| 400 | 1640210 | sort by null is not allowed | The sorting parameters are filled in incorrectly, please check |
| 400 | 1640211 | biz codes is empty | The business code is empty, please check |
| 400 | 1640212 | mdm codes is empty | The main data code is empty, please check |
| 400 | 1640213 | biz id is invalid | The business code is illegal, please check it. |
| 500 | 1640101 | system error | System error, contact Lark technical support to check. |
| 500 | 1640102 | invalid generic method | The request method is wrong, contact Lark technical support to check. |
| 500 | 1640103 | call meta service error | The request metadata method is wrong, contact Lark technical support to troubleshoot. |
| 500 | 1640104 | meta data not found | Can't find the corresponding metadata, contact Lark technical support to check. |
| 500 | 1640105 | data source request invalid | The database request is illegal, contact Lark technical support to check. |
| 500 | 1640106 | db error | Database error, contact Lark technical support to troubleshoot. |
| 500 | 1640107 | func not supported yet | The service does not support this method, contact Lark technical support to check. |
| 500 | 1640108 | field type not supported yet | Field type is not supported, contact Lark technical support to check. |
| 500 | 1640109 | validator not found | No data verification rules were found, contact Lark technical support for investigation. |
| 500 | 1640110 | transformer not found | No data conversion rules found, contact Lark technical support to check. |
| 500 | 1640111 | executor not found | No actuator found, contact Lark technical support to check. |
| 500 | 1640112 | invalid data source type | Error database type, contact Lark technical support to troubleshoot |
| 500 | 1640113 | data source not found | No database found, contact Lark technical support for investigation. |
| 500 | 1640114 | data source invalid | The database is invalid, contact Lark technical support to check. |
| 500 | 1640115 | meta setting not found | No metadata configuration found, contact Lark technical support to check. |
| 500 | 1640116 | generate page token error | Page code generation error, contact Lark technical support to check |
| 500 | 1640117 | record duplicate error | Line-level records are repeated, contact Lark technical support to check. |
| 500 | 1640118 | generate idl error | IDL generated an error, contact Lark technical support to troubleshoot. |
| 500 | 1640119 | producer not found error | Message producer error, contact Lark technical support to troubleshoot. |
| 500 | 1640120 | custom func request error | Custom actuator request error, contact Lark technical support to troubleshoot |
| 500 | 1640121 | custom func response error | Custom actuator response error, contact Lark technical support to troubleshoot |
