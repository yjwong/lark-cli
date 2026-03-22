---
title: "recognize invoice"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/document_ai-v1/invoice/recognize"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/document_ai/v1/invoice/recognize"
service: "document_ai-v1"
resource: "invoice"
section: "AI"
rate_limit: "10 per second"
scopes:
  - "document_ai:invoice:recognize"
updated: "1710921012000"
---

# Identify invoices in documents

Invoice identification interface, supports one-time identification of five file types: JPG/JPEG/PNG/PDF

> The limited viewership of single tenant: 10QPS, there is no limited viewership of applications under the same tenant, and the 10QPS limited viewership of this tenant is shared

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/document_ai/v1/invoice/recognize |
| HTTP Method | POST |
| Rate Limit | 10 per second |
| Supported app types | custom,isv |
| Required scopes | `document_ai:invoice:recognize` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Example value**: "multipart/form-data; boundary=---7MA4YWxkTrZu0gW" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `file` | `file` | Yes | Identified invoice file (supports JPG/JPEG/PNG/BMP) **Example value**: file binary | ### Request body example

```HTTP
---7MA4YWxkTrZu0gW
Content-Disposition: form-data; name="file";
Content-Type: application/octet-stream

---7MA4YWxkTrZu0gW
```

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `invoice` | `invoice` | Invoice information |
|     `entities` | `invoice_entity[]` | Recognized entities |
|       `type` | `string` | Identified field types **Optional values are**:  - `customer_name`: Customer name - `invoice_id`: Invoice ID - `invoice_date`: Invoice date - `due_date`: Due date - `vendor_name`: Vendor name - `vendor_address`: Vendor address - `vendor_address_recipient`: Vendor address recipient - `billing_address`: Billing address - `billing_address_recipient`: Billing address recipient - `shipping_address`: Shipping address - `shipping_address_recipient`: Shipping Address Recipient - `sub_total`: Sub total - `total_tax`: Total tax - `invoice_total`: Invoice total - `sub_total_format`: Subtotal - Formatted result, e.g. CNY 100 - `sub_total_amount`: Subtotal - purely numerical result, e.g. 100 - `sub_total_currency_code`: Subtotal - Currency code, e.g. CNY - `sub_total_currency_symbol`: Subtotal - Currency identifier, e.g. ￥ - `total_tax_format`: Total amount - Formatted result, e.g. CNY 100 - `total_tax_amount`: Total amount - purely numerical result, e.g. 100 - `total_tax_currency_code`: Total amount - Currency code, e.g. CNY - `total_tax_currency_symbol`: Total amount - Currency identifier, e.g. ￥ - `invoice_total_format`: Total invoice - Formatted result, e.g. CNY 100 - `invoice_total_amount`: Total invoice - purely numerical result, e.g. 100 - `invoice_total_currency_code`: Total invoice - Currency code, e.g. CNY - `invoice_total_currency_symbol`: Total invoice - currency identifier, e.g., ￥ - `invoice_date_format`: Date of invoice - Formatted results yyyy-mm-dd - `due_date_format`: Invoice due Date - Formatted results yyyy-mm-dd  |
|       `value` | `string` | Identify the text information of the field |
|       `items` | `items_entity[]` | Identified cost breakdown |
|         `amount` | `currency_entity` | Amount |
|           `content` | `string` | Source Text |
|           `format` | `string` | Formatted result |
|           `currency_code` | `string` | Currency Code |
|           `currency_symbol` | `string` | Currency Symbol |
|           `amount` | `string` | Purely numerical result |
|         `date` | `common_entity` | Date |
|           `content` | `string` | Content |
|         `description` | `common_entity` | Description |
|           `content` | `string` | Source Text |
|         `product_code` | `common_entity` | Product Code |
|           `content` | `string` | Source Text |
|         `quantity` | `common_entity` | Purely numerical result |
|           `content` | `string` | Source Text |
|         `tax` | `currency_entity` | Tax |
|           `content` | `string` | Source Text |
|           `format` | `string` | Formatted Result |
|           `currency_code` | `string` | Currency Code |
|           `currency_symbol` | `string` | Currency Symbol |
|           `amount` | `string` | Purely numerical result |
|         `tax_rate` | `common_entity` | Tax Rate |
|           `content` | `string` | Source Text |
|         `unit` | `common_entity` | Unit |
|           `content` | `string` | Source Text |
|         `unit_price` | `currency_entity` | Unit Price |
|           `content` | `string` | Source Text |
|           `format` | `string` | Formatted Result |
|           `currency_code` | `string` | Currency Code |
|           `currency_symbol` | `string` | Currency Symbol |
|           `amount` | `string` | Purely numerical result | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "invoice": {
            "entities": [
                {
                    "type": "customer_name",
                    "value": "Tom",
                    "items": [
                        {
                            "amount": {
                                "content": "￥100",
                                "format": "￥100",
                                "currency_code": "CNY",
                                "currency_symbol": "￥",
                                "amount": "100"
                            },
                            "date": {
                                "content": "this is an example"
                            },
                            "description": {
                                "content": "this is an example"
                            },
                            "product_code": {
                                "content": "this is an example"
                            },
                            "quantity": {
                                "content": "this is an example"
                            },
                            "tax": {
                                "content": "￥100",
                                "format": "￥100",
                                "currency_code": "CNY",
                                "currency_symbol": "￥",
                                "amount": "100"
                            },
                            "tax_rate": {
                                "content": "this is an example"
                            },
                            "unit": {
                                "content": "this is an example"
                            },
                            "unit_price": {
                                "content": "￥100",
                                "format": "￥100",
                                "currency_code": "CNY",
                                "currency_symbol": "￥",
                                "amount": "100"
                            }
                        }
                    ]
                }
            ]
        }
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 2110001 | param is invalid | Input file error, refer to the documentation to check the input parameters |
| 400 | 2110002 | no valid entity | Invoice information is not detected, refer to the document to check whether the input file is valid |
| 500 | 2110010 | internal error, please try later | Backend service exception or network exception, can be rerequested |
| 400 | 2110003 | You have reached the Intelligent document parsing limit. To continue using this function, please contact sales to purchase more. | You have reached the Intelligent document parsing limit. To continue using this function, please contact sales to purchase more. |
