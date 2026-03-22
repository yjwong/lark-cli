---
title: "Obtain company assign information"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/tenant-v2/tenant-product_assign_info/query"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/tenant/v2/tenant/assign_info_list/query"
service: "tenant-v2"
resource: "tenant-product_assign_info"
section: "Company Information"
rate_limit: "100 per minute"
scopes:
  - "tenant:tenant.product_assign_info:read"
updated: "1692861219000"
---

# Obtain company seat information

Obtain the seat list to be allocated under the tenant, including seat name, subscription ID, quantity and validity period.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/tenant/v2/tenant/assign_info_list/query |
| HTTP Method | GET |
| Rate Limit | 100 per minute |
| Supported app types | custom,isv |
| Required scopes | `tenant:tenant.product_assign_info:read` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `assign_info_list` | `tenant_assign_info[]` | List of seats to be allocated by tenants |
|     `subscription_id` | `string` | subscription id |
|     `license_plan_key` | `string` | license key |
|     `product_name` | `string` | product name |
|     `i18n_name` | `product_i18n_name` | i18n name |
|       `zh_cn` | `string` | zh_cn name |
|       `ja_jp` | `string` | ja_jp name |
|       `en_us` | `string` | en_us name |
|     `total_seats` | `string` | total num of seats |
|     `assigned_seats` | `string` | assigned num of seats |
|     `start_time` | `string` | start time |
|     `end_time` | `string` | end time | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "assign_info_list": [
            {
                "subscription_id": "7079609167680782300",
                "license_plan_key": "suite_enterprise_e5",
                "product_name": "旗舰版 E5",
                "i18n_name": {
                    "zh_cn": "zh_cn_name",
                    "ja_jp": "ja_jp_name",
                    "en_us": "en_name"
                },
                "total_seats": "500",
                "assigned_seats": "20",
                "start_time": "1674981000",
                "end_time": "1674991000"
            }
        ]
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1180001 | param is invalid | please check the request parameters |
