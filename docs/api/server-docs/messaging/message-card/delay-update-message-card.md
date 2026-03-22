---
title: "Delay update message card"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uMDO1YjLzgTN24yM4UjN"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/interactive/v1/card/update"
section: "Messaging"
updated: "1717574922000"
---

# Delay update message card
After the user completes the interactive operation on the message card, update the card content to the specified user and give timely feedback.

### Use case
1. After a user clicks a card, the business side needs to take a long time to process the call and cannot return the requested card content within 3 seconds.
2. Specify to update only part of the card content seen by members (same `message_id`) who received this card. Delayed update uses the `token` in the interactive postback event to specify the message of the target update, without paying extra attention to the `message_id` of the original message.

> - Users must interact with the card to trigger the API. Unconditional updates are not supported.
> - The token used for update delay is valid for 30 minutes. After timeout, the card cannot be updated by using the token.
> - The update delay API can be called only later than synchronous return. Otherwise unpredictable behavior occurs.
> When the server processes the call, you can immediately return an empty string and then call the update delay API within 30 minutes to update the card.
> - Only cards interacted with a user can be updated by the user.
> - Card content cannot exceed 100KB after conversion.
> - The same token can only be used 3 times.

## Request
| HTTP URL | https://open.larksuite.com/open-apis/interactive/v1/card/update |
| --- | --- |
| HTTP Method | POST | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about how to access and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Request body
| Parameter | Type | Required | Description | Example |                                                   
| - | - | - | - | - | - |
| token | string | Yes | Token used to update the card. The value is different from the tenant_access_token. To obtain the value, you can call the Interact with a card to obtain the return content API. | c-xxxxxxx
| card | object | Yes | Message card description. For more information, see Card structure. |
|  ∟open_ids | array | No | Specifies the user to whom the card is updated. For a shared card, the card is updated for all users by default and this field can be left blank. Open ID is recommended here. For details, refer to How to get Open ID?| ["ou_5ad573a6411d72b8305fda3a9c15c70e"] | ### Request body example

```json
{
    "token":"c-515fbxxxx", // Card update token returned in the POST request
     "card": {
        "open_ids":["ou_515fbe9d04838174e2035f8xxxx53d07f"], // Specifies the list of users to whom the message needs to be updated. This field is optional.
        //New card content in JSON after the card is updated. For more information about the specific structure, choose Common Capability > Message Card.
        "elements": [
            {
                "tag": "div",
                "text": {
                    "tag": "plain_text",
                    "content": "Test the overflow and datePicker features"
                },
                "fields": [
                    {
                        "is_short": true,
                        "text": {
                            "tag": "lark_md",
                            "content": "**Agreed**"
                        }
                    }
                ]
            }
        ]
    }
 }
```

## Response

### Response body

|Parameter|Type|Description|
|-|-|-|
|code|int|Return code. A value other than 0 indicates failure.|
|msg|string|Return code description| ### Response body example

```json
{
    "code": 0,
    "msg": "ok"
}
```

### Error code
| Error code | Description | Troubleshooting suggestions |
| --- | --- | --- |
| 11311 | The card format doesn't meet the requirements. | For more information, see the specific content of the error message. For more information about the card structure and format, see Card structure. |
| 10002 | The card parameter is not specified. | Specify the card field. For more information about the format, see Card structure. |
| 100000 | The card content cannot exceed 100 KB after conversion. | Reduce the size of the card. |
| 100030 | The body parameter passed in doesn't meet the JSON specification. | Check the input parameter. |
| 200000 | The card message has been recalled. | You cannot update the card message because the card message has been recalled. |
| 200310 | Update cards sent by other apps. | You cannot update cards sent by other apps. |
| 200320 | Check whether the open_ids parameter of the non-shared card is correct. | Check whether the open_ids parameter is correct. |
| 300020 | The token format of the card that you need to update is incorrect. | Check whether the token is in the format of c-xxxx. You can obtain the value by calling the Interact with a card to obtain the return content API. |
| 300030 | The token of the card that you need to update has expired. | A token is valid for 30 minutes. Please check whether the token is within the period of validity. |
| 300040 | The token of the card that you need to update has been used for the maximum number of times. | A token can be used only for 3 times. Please check whether the token has been used for the maximum number of times. |
| 300090 | The open_ids field must be specified for the non-shared card. | Check whether the open_ids parameter is specified and correct. | For more information about other common error codes, see Server-side error codes.
