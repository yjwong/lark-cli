---
title: "Get event's outbound IP"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uYDNxYjL2QTM24iN0EjN/event-v1/outbound_ip/list"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/event/v1/outbound_ip"
service: "event-v1"
resource: "outbound_ip"
section: "Events and callbacks"
scopes:
  - "event:ip_list"
updated: "1699845022000"
---

# Get event's outbound IP

Lark Open Platform sends events to the callback address configured for the app using specific IP addresses. Apps can use this interface to obtain all related IP addresses.

> IP addresses may change, so it is recommended to regularly fetch the latest IP addresses and automatically update them in the firewall rules. Additionally, when IP changes occur, the Open Platform will push card messages and publish update logs to notify developers.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/event/v1/outbound_ip |
| HTTP Method | GET |
| Supported app types | custom,isv |
| Required scopes | `event:ip_list` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `page_size` | `int` | No | Page size **Example value**: 10 **Data validation rules**: - Maximum value: `50` |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: "AQD9/Rn9eij9Pm39ED40/dk53s4Ebp882DYfFaPFbz00L4CMZJrqGdzNyc8BcZtDbwVUvRmQTvyMYicnGWrde9X56TgdBuS+JKiSIkdexPw=" | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `ip_list` | `string[]` | outbound ip |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned |
|   `has_more` | `boolean` | Whether the response body has more parameters | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "ip_list": [
            "1.1.1.1"
        ],
        "page_token": """",
        "has_more": false
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1810001 | param is invalid | check params |
