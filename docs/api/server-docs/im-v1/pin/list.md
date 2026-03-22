---
title: "Get pins in group"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/pin/list"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/im/v1/pins"
service: "im-v1"
resource: "pin"
section: "Messaging"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "im:message"
  - "im:message.pins:read"
  - "im:message:readonly"
updated: "1717574935000"
---

# Get pins in group

Get all pin data within the specified time range in the chat.

> Notes:
> - Bot ability must be enabled.
> - When you get pins  in a group, the bot must be in the group chat.
> - The obtained pins are sorted in descending order by the creation time of the pin.
> - The default query limit of the API is 50 QPS.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/im/v1/pins |
| HTTP Method | GET |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `im:message` `im:message.pins:read` `im:message:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `chat_id` | `string` | Yes | Group ID. For details, refer to Group ID description. **Example value**: oc_234jsi43d3ssi993d43545f |
| `start_time` | `string` | No | Start time of pin (millisecond timestamp). If not filled in, the earliest pin information in the group chat will be obtained by default **Example value**: 1658632251800 |
| `end_time` | `string` | No | End time of pin (millisecond timestamp). If not filled in, the default will be obtained from the latest pin information in the group chat **Note**: the `end_time` value should be greater than the `start_time` value **Example value**: 1658731646425 |
| `page_size` | `int` | No | **Example value**: 20 **Default value**: `20` **Data validation rules**: - Value range: `1` ～ `50` |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: GxmvlNRvP0NdQZpa7yIqf_Lv_QuBwTQ8tXkX7w-irAghVD_TvuYd1aoJ1LQph86O-XImC4X9j9FhUPhXQDvtrQ== | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `items` | `pin[]` | Information of pin |
|     `message_id` | `string` | Message ID of pin |
|     `chat_id` | `string` | Chat ID of the pined message |
|     `operator_id` | `string` | Operator ID of pin |
|     `operator_id_type` | `string` | Operator ID type of pin. If the pin operator is a user, it is open_id, and If the pin operator is a bot, it is app_id. |
|     `create_time` | `string` | Creation time of pin (millisecond timestamp) |
|   `has_more` | `boolean` | Whether the response body has more parameters |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "items": [
            {
                "message_id": "om_dc13264520392913993dd051dba21dcf",
                "chat_id": "oc_a0553eda9014c201e6969b478895c230",
                "operator_id": "ou_7d8a6e6df7621556ce0d21922b676706ccs",
                "operator_id_type": "open_id",
                "create_time": "1615380573211"
            }
        ],
        "has_more": true,
        "page_token": "GxmvlNRvP0NdQZpa7yIqf_Lv_QuBwTQ8tXkX7w-irAghVD_TvuYd1aoJ1LQph86O-XImC4X9j9FhUPhXQDvtrQ=="
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 230001 | Your request contains an invalid request parameter. | Parameter error. Please check the input parameter according to the error message returned by the interface and refer to the document. |
| 400 | 230002 | The bot can not be outside the group. | The bot isn't in the group chat. |
| 400 | 230006 | Bot ability is not activated. | Bot ability is not enabled. The bot ability can be added on the [Developer Console](https://open.larksuite.com/app) -> Features -> Add Features, and it will take effect after a new version is released. |
| 400 | 230013 | Bot has NO availability to this user. | The bot is not available for the user. You can edit the available scope of the application on the [Developer Console](https://open.larksuite.com/app)  -> App Versions -> Version Management & Release -> Create a version, and it will take effect after the new version is released. |
| 400 | 230027 | Lack of necessary permissions. | This operation is not currently supported in external groups. |
| 400 | 230048 | Get chat pins trigger app_id limit. | Rate-limiting is triggered. Please reduce the request speed and try again later. | > **Note:** For other server-side error codes, refer to Server-side error codes.
