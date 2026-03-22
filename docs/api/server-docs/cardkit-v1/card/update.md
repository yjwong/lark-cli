---
title: "Update card entity"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/cardkit-v1/card/update"
method: "PUT"
api_path: "https://open.larksuite.com/open-apis/cardkit/v1/cards/:card_id"
service: "cardkit-v1"
resource: "card"
section: "Lark Card"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "cardkit:card:write"
updated: "1751018163000"
---

# Update card entity

Update card entity

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/cardkit/v1/cards/:card_id |
| HTTP Method | PUT |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes | `cardkit:card:write` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ## Use restrictions

This interface only supports schema 2.0 card structures.
- When calling this interface, it is not supported to set the card to exclusive card mode. That is, it is not supported to set the 'update_multi 'attribute in the card JSON data to'false'.
- The identity of the application calling the interface must be the same as the identity of the application creating the target card entity.

### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `card_id` | `string` | Card entity ID. Get it by Create Card Entity **Example value**: "7355372766134157313" **Data validation rules**: - Length range: `1` ～ `20` characters | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `card` | `card` | Yes | Updated card content |
|   `type` | `string` | Yes | Type of card data. Take the fixed value card_json. **Example value**: "card_json" **Optional values are**:  - `card_json`: Card JSON  **Data validation rules**: - Length range: `1` ～ `50` characters |
|   `data` | `string` | Yes | The content of the card JSON data. Only card structures in schema 2.0 are supported. The following example values are not escaped. Please be careful to convert them to JSON serialized strings when using them. **Example value**: "{\"schema\":\"2.0\",\"header\":{\"title\":{\"content\":\"title\",\"tag\":\"plain_text\"}},\"body\":{\"elements\":[{\"tag\":\"markdown\",\"content\":\"body\"}]}}" **Data validation rules**: - Length range: `1` ～ `1000000` characters |
| `uuid` | `string` | No | Idempotent ID, which can be passed in a unique UUID to ensure that the same batch of operations is performed only once. **Example value**: "a0d69e20-1dd1-458b-k525-dfeca4015204" **Data validation rules**: - Length range: `1` ～ `64` characters |
| `sequence` | `int` | Yes | The sequence number of the card operation when the card is in streaming update mode. It is used to ensure the chronological order of multiple updates. **Note**: Please ensure that the value of `sequence` is strictly increasing compared to the previous operation when manipulating the same card through the Card OpenAPI. **Data validation rules**: Positive integers within the int32 range (`1`~`2147483647`). **Example value**: 1712578784 | ### Request body example

{
    "card": {
        "type": "card_json",
        "data": "{\"schema\":\"2.0\",\"header\":{\"title\":{\"content\":\"title\",\"tag\":\"plain_text\"}},\"body\":{\"elements\":[{\"tag\":\"markdown\",\"content\":\"body\"}]}}"
    },
    "uuid": "a0d69e20-1dd1-458b-k525-dfeca4015204",
    "sequence": 1712578784
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {}
}

### Error Codes
| HTTP Status Code | Error Code | Description | Troubleshooting Suggestions |
| --- | --- | --- | --- |
| 400 | 10002 | Your request contains an invalid request parameter. | Parameter error. Please check the input parameters according to the error information returned by the interface and refer to the documentation. |
| 400 | 200740 | The card entity does not exist | The card entity does not exist. Please check if the entity ID is correct. |
| 400 | 200750 | The card entity has expired | The card entity has expired. The validity period of the card entity is 14 days. You cannot call the relevant interface to operate the card after creating the card entity for more than 14 days. |
| 400 | 200770 | UUID conflict | UUID conflict. Please pass in a unique UUID to ensure that the same batch of operations is performed only once. |
| 400 | 200810 | The card is in an ongoing interaction and cannot be updated | The card is in an ongoing interaction and cannot be updated. |
| 400 | 200860 | Card content exceeds limit. | The card size exceeds the limit. Please keep the card size within 30KB. |
| 400 | 300301 | Duplicate element_id in card component. | Duplicate element ID (element_id) in the card component. |
| 400 | 300302 | update_multi property is false | In streaming update mode, the global card property update_multi must be set to true. |
| 400 | 300303 | Only schema 2.0 is supported | This interface only supports Schema v2.0 structure. For details, refer to Card JSON 2.0 Structure. |
| 400 | 200220 | Failed to generate card content | Failed to generate card content. Please check if the card JSON format is incorrect. |
| 400 | 300305 | The number of card components exceeds 200 | Exceeding card component limit. In the card JSON 2.0 structure, a single card supports a maximum of 200 elements (such as text elements with the tag `plain_text`) or components. Please ensure that the total number of components and elements does not exceed 200. |
| 400 | 300307 | The card DSL is empty | The card JSON data is empty. Please check the data. |
| 400 | 300311 | The current application does not have permission to update/use this card | The current application does not have permission to update or use this card. Only the application that created the card entity is allowed to call the relevant OpenAPI to send and operate the card. |
| 400 | 300312 | Unable to update element tag | Unable to update the tag attribute of the component when updating the card entity. |
| 400 | 300313 | Failed to update element properties | Failed to update component properties. Please check the input parameters according to the error information returned by the interface. |
| 400 | 300314 | Failed to delete element | Failed to delete component. Please check the input parameters according to the error information returned by the interface. |
| 400 | 300315 | Failed to add element | Failed to add component. Please check the input parameters according to the error information returned by the interface. |
| 400 | 300317 | The sequence number for operating on the card did not increment consecutively | The operation sequence number (`sequence`) of the card is not incrementing in order. Please ensure that when using the Card OpenAPI to operate on the same card, the `sequence` value strictly increases with each operation. |
| 400 | 300121 | Failed to replace element | Failed to replace component. Please check the input parameters according to the error information returned by the interface. |
| 400 | 300122 | Failed to update card configuration | Failed to update card configuration. Please check the input parameters according to the error information returned by the interface. |
