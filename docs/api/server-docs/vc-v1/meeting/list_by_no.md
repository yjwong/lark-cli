---
title: "List meetings of same meeting number"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting/list_by_no"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/vc/v1/meetings/list_by_no"
service: "vc-v1"
resource: "meeting"
section: "Video Conferencing"
rate_limit: "100 per minute"
scopes:
  - "vc:meeting:readonly"
updated: "1692787315000"
---

# List meetings of same meeting number

Obtains the meeting brief list associated with the meeting number for a specified time period (within 90 days).

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/vc/v1/meetings/list_by_no |
| HTTP Method | GET |
| Rate Limit | 100 per minute |
| Supported app types | custom,isv |
| Required scopes | `vc:meeting:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `meeting_no` | `string` | Yes | 9-digit meeting number **Example value**: 123456789 |
| `start_time` | `string` | Yes | Query start time (unix time in sec) **Example value**: 1608888867 |
| `end_time` | `string` | Yes | Query end time (unix time in sec) **Example value**: 1608888867 |
| `page_token` | `string` | No | Page identifier. It is not filled in the first request, indicating traversal from the beginning; when there will be more groups, the new page_token will be returned at the same time, and the next traversal can use the page_token to get more groups **Example value**: 5 |
| `page_size` | `int` | No | **Example value**: 10 **Default value**: `20` **Data validation rules**: - Maximum value: `50` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `has_more` | `boolean` | Whether the response body has more parameters |
|   `page_token` | `string` | Page identifier, when has_more is true, a new page_token will also be returned. Otherwise, page_token will not be returned |
|   `meeting_briefs` | `meeting[]` | Meeting brief list |
|     `id` | `string` | Meeting ID (Unique identifier of video conference, generated after the conference is started) |
|     `topic` | `string` | Meeting topic |
|     `url` | `string` | Meeting link (Lark users can join the meeting by clicking the meeting link) |
|     `meeting_no` | `string` | Meeting number | ### Response body example

{
    "code": 0,
    "data": {
        "has_more": true,
        "meeting_briefs": [
            {
                "id": "7011030664708603907",
                "meeting_no": "146453285",
                "topic": "Test Title"
            },
            {
                "id": "7011031045899354115",
                "meeting_no": "146453285",
                "topic": "Test Title"
            }
        ],
        "page_token": "2"
    },
    "msg": ""
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 500 | 121001 | internal error | Internal server error. If retry still doesn't work, contact the administrator. |
| 400 | 121003 | param error | Parameter error. Check the value range of parameters (please check yourself according to the field description above). |
