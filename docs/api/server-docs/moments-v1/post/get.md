---
title: "Query post information"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/moments-v1/post/get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/moments/v1/posts/:post_id"
service: "moments-v1"
resource: "post"
section: "Moments"
rate_limit: "20 per second"
scopes:
  - "moments:moments:readonly"
  - "moments:moments:access_all"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1760520778000"
---

# Query post information

Query post entity data information by post id

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/moments/v1/posts/:post_id |
| HTTP Method | GET |
| Rate Limit | 20 per second |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `moments:moments:readonly` `moments:moments:access_all` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `post_id` | `string` | Post ID, which can be got from the data returned by the "Publish moment" interface or the "Moment posted" event. **Example value**: "6934510454161014804" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: open_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `post` | `post` | Post entity |
|     `user_id` | `string` | Posting user ID (the type is the type passed in the request, only the value under the real name) |
|     `content` | `string` | Post content |
|     `image_key_list` | `string[]` | Image key list (Download is not supported currently) |
|     `media_file_token` | `string` | Token for media files (Not supported currently) |
|     `comment_count` | `int` | Number of comments |
|     `reaction_set` | `reaction_set` | Post reactions and their number |
|       `reactions` | `reaction_list[]` | Reaction list |
|         `type` | `string` | Type of reaction |
|         `count` | `int` | Number of reactions of this type |
|       `total_count` | `int` | Number of all reactions |
|     `id` | `string` | Post ID |
|     `create_time` | `string` | Post creation time, format rfc3339, E.g. "2006-01-02T15:04:05Z07:00" |
|     `media_cover_image_key` | `string` | Video cover image (Not supported currently) |
|     `category_ids` | `string[]` | The categories the post belongs to |
|     `link` | `string` | Post link |
|     `user_type` | `int` | Post User type **Optional values are**:  - `1`: Real name - `2`: Nickname - `3`: Anonymous - `4`: Official Account  |
|     `dislike_count` | `int` | Dislike count | ### Response body example

{
    "code": 0,
    "msg": "Success",
    "data": {
        "post": {
            "user_id": "ou_xxxxx",
            "content": "[[{\"tag\":\"text\",\"text\":\"Hello world.....\"}, {\"tag\":\"a\",\"text\":\"View original content\",\"href\":\"https://www.example.com/\"}]]",
            "image_key_list": [
                "Download is not supported currently"
            ],
            "media_file_token": "This field is not supported currently",
            "comment_count": 1,
            "reaction_set": {
                "reactions": [
                    {
                        "type": "OK",
                        "count": 12
                    }
                ],
                "total_count": 12
            },
            "id": "6934510454161014804",
            "create_time": "2022-05-23T00:00:00+08:00",
            "media_cover_image_key": "This field is not supported currently",
            "category_ids": [
                "71123"
            ],
            "link": "https://applink.larksuite.com/client/moments/detail?postId=6934510454161014804",
            "user_type": 1,
            "dislike_count": 0
        }
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1420001 | Invalid args | Check whether the post_id format is correct |
| 403 | 1420002 | Permission denied | Check whether you have permission to query posts |
| 404 | 1420004 | Post not found | Check whether the post id is correct or whether it has been deleted |
| 429 | 1420003 | Triggers the frequency limit | Please request less frequently and try again later |
