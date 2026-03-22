---
title: "Create card entity"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/cardkit-v1/card/create"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/cardkit/v1/cards"
service: "cardkit-v1"
resource: "card"
section: "Lark Card"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "cardkit:card:write"
updated: "1751018151000"
---

# Create card entity

Create a card entity based on the card JSON code.

## Use restrictions

This interface only supports schema 2.0 card structures.
- When calling this interface, it is not supported to set the card to exclusive card mode. That is, it is not supported to set the `update_multi` attribute in the card JSON data to `false`.
- The valid period of the card entity is 14 days. That is, after 14 days of creating the card entity, you will not be able to call the relevant interface to operate the card.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/cardkit/v1/cards |
| HTTP Method | POST |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes | `cardkit:card:write` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `type` | `string` | Yes | Type of card data. Take the fixed value `card_json`. **Example value**: "card_json" **Data validation rules**: - Length range: `1` ～ `50` characters |
| `data` | `string` | Yes | The contents of the card JSON data. Only Card JSON 2.0 structure is supported, i.e. you must declare the schema as 2.0. The following example values are not escaped, please be careful to convert them to JSON serialized strings when using them. **Example value**: "{\"schema\":\"2.0\",\"header\":{\"title\":{\"content\":\"Card title\",\"tag\":\"plain_text\"}},\"config\":{\"streaming_mode\":true,\"summary\":{\"content\":\"[Generating]\"}},\"body\":{\"elements\":[{\"tag\":\"markdown\",\"content\":\"Card content\",\"element_id\":\"markdown_1\"}]}}" **Data validation rules**: - Length range: `1` ～ `3000000` characters | ### Request body example

{
    "type": "card_json",
    "data": "{\"schema\":\"2.0\",\"header\":{\"title\":{\"content\":\"Card title\",\"tag\":\"plain_text\"}},\"config\":{\"streaming_mode\":true,\"summary\":{\"content\":\"[Generating]\"}},\"body\":{\"elements\":[{\"tag\":\"markdown\",\"content\":\"Card content\",\"element_id\":\"markdown_1\"}]}}"
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `card_id` | `string` | Card entity ID | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "card_id": "7355372766134157313"
    }
}

### Error Codes
| HTTP Status Code | Error Code | Description | Troubleshooting Suggestions |
| --- | --- | --- | --- |
| 400 | 10002 | Your request contains an invalid request parameter. | Parameter error. Please check the input parameters according to the error information returned by the interface and refer to the documentation. |
| 400 | 200860 | Card content exceeds limit. | The card size exceeds the limit. It is recommended to keep the card size within 30KB. |
| 400 | 300301 | Duplicate element_id in card component. | The ID (element_id) of the card internal component is duplicated. Please check and modify. |
| 400 | 300302 | update_multi property is false | In streaming update mode, the global card property update_multi must be set to true. |
| 400 | 300303 | Only schema 2.0 is supported | This interface only supports Schema v2.0 structure. For details, refer to Card JSON 2.0 Structure. |
| 400 | 200220 | Failed to generate card content | Failed to generate card content. Please check if the card JSON format is incorrect. |
| 400 | 300305 | The number of card components exceeds 200 | Exceeding card component limit. In the card JSON 2.0 structure, a single card supports a maximum of 200 elements (such as text elements with the tag `plain_text`) or components. Please ensure that the total number of components and elements does not exceed 200. |
| 400 | 300307 | The card DSL is empty | The card JSON data is empty. Please check and modify. |
