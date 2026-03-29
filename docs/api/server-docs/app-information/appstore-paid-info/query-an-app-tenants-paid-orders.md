---
title: "Query an App Tenant’s Paid Orders"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uETNwUjLxUDM14SM1ATN"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/pay/v1/order/list"
section: "App Information"
updated: "1646720026000"
---

#  Query paid orders under an app tenant
This API is used to query the paid orders under an app tenant by pages. Each purchase corresponds to a unique order which records the information of the plan purchased. The business party must process the validity period and price plan upgrades on their own.

## Request
| HTTP URL | https://open.larksuite.com/open-apis/pay/v1/order/list |
| --- | --- |
| HTTP Method | GET | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about how to access and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request parameters

| Parameter         | Type           | Required        | Description         | 
| --------- | --------------- | -------   | ----------- | 
|`status` | `string` | No |Filter condition set to obtain the information of the user's purchased plans. "normal" means normal orders; "refunded" means refunded orders; when this field is empty or "all", it means all orders. Unpaid orders cannot be queried. |
|`page_size` | `int` | Yes | Number of orders displayed per page |
|`page_token` | `string` | No | Paging token, which can be obtained from the response to the last request. If this is not specified or is left empty, the results will be obtained from the beginning. |
|`tenant_key` | `string` | No | Unique identifier of the tenant that purchased the app. If this is empty, all orders for the app will be obtained. If this field has a value, orders for the app placed by the tenant will be obtained. | ## Response
### Response body
| Parameter| Description | 
|-|-|
|code| Return code. A value other than 0 indicates failure. |
|msg | Return code description |
|data | Returned business information |
|  ∟total | Total number of orders|
|  ∟has_more | Whether there is more data. true: Yes; false: No. |
|  ∟page_token | Token for next page data, which can be used to request the next page of data. If has_more is false, this field is empty. |
|  ∟order_list | Order information list |
|    ∟order_id | Order ID, which is unique. |
|    ∟price_plan_id | Price plan ID, which is unique. |
|    ∟price_plan_type | Price plan type. "trial" - Trial; "permanent" - One-time purchase; "per_year" - Yearly fixed payment; "per_month" - Monthly fixed payment; "per_seat_per_year" - Yearly payment by account quantity; "per_seat_per_month" - Monthly payment by account quantity; "permanent_count" - Payment by value-added package. |
|    ∟seats | Actual number of accounts purchased. This is only valid when price_plan_type is per_seat_per_year or per_seat_per_month.|
|    ∟buy_count | Number of plans purchased. It is always 1. |
|    ∟create_time | Order creation timestamp |
|    ∟pay_time | Order payment timestamp |
|    ∟status | Order status. "normal" - Normal; "refund" - Refunded. |
|    ∟buy_type | Purchase type. "buy" - Normal purchase; "upgrade" - Upgrade purchase (only available when price_plan_type is per_year, per_month, per_seat_per_year, or per_seat_per_month); "renew" - Renewal purchase. |
|    ∟src_order_id | Source order ID. When the current order is an upgrade purchase (buy_type is upgrade), this field records the ID of the source order. |
|    ∟dst_order_id | New order ID after upgrade. If an upgrade purchase was made for the current order before, this field records the new order ID generated for the upgrade purchase. The current order remains valid. |
|    ∟order_pay_price | Actual order price paid, in cents |
|    ∟tenant_key | Unique identifier of the tenant |
### Response example

```json
{
    "code":0,
    "msg":"success",
    "data": {        
        "total":100,
        "has_more":true,
        "page_token":"10",
        "order_list": [
            {
                "order_id":"6708978506916697671",
                "price_plan_id":"price_9daf66c96968c003",
                "price_plan_type":"per_seat_per_year",
                "seats":2,
                "buy_count":1,
                "create_time":"1565003215",
                "pay_time":"1565003216",
                "status":"normal",
                "buy_type":"buy",
                "src_order_id":"6718900620264210445",
                "dst_order_id":"6717597079260102659",
                "order_pay_price":100,
                "tenant_key":"2e5c3a3ae38f175f"
            }
        ]
    }
} 
```
