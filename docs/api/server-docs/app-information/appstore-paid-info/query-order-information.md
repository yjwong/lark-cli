---
title: "Query Order Information"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uITNwUjLyUDM14iM1ATN"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/pay/v1/order/get"
section: "App Information"
updated: "1646720028000"
---

#  Query order information
This API is used to query details of an order.

## Request
| HTTP URL | https://open.larksuite.com/open-apis/pay/v1/order/get |
| --- | --- |
| HTTP Method | GET | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about how to access and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request parameters

| Parameter         | Type           | Required        | Description         |
| --------- | --------------- | -------   | ----------- |
|`order_id` | `string` | Yes |Order ID | ## Response

### Response body
| Parameter| Description | 
| --------- | --------------- | 
|code | Return code. A value other than 0 indicates failure. |
|msg | Return code description |
|data | Returned business information |
|  ∟order | Order information |
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
|    ∟tenant_key | Unique identifier of the tenant | ### Response example
```json
{
    "code":0,
    "msg":"success",
    "data": {        
        "order": {
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
                "tenant_key":"2e5c3a3ae38f175f",
            }
    }
} 
```
