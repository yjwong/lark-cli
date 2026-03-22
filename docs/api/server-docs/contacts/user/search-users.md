---
title: "Search Users"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uMTM4UjLzEDO14yMxgTN"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/search/v1/user?query=zhangsan&page_size=20"
section: "Contacts"
updated: "1646999081000"
---

# Search Users

Uses a user's identity to search for other users. Users external to the
organization or those that have left the organization cannot be found.

> Permission to search users must have been granted to invoke this API.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/search/v1/user?query=zhangsan&page_size=20 |
| HTTP Method | GET |
| Required scopes | search users  |
| Required field scopes Enable the scopes for fields you want to return | Obtain user ID | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `user_access_token` **Value Format**: "Bearer `access_token`" **Example Value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters

|      Parameters     |      Type     |      Required/optional     |      Example     |      Description                                                                                                                                                                                                                                                                                      |
|---------------------|---------------|----------------------------|------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
|     query           |     string    |     Required               |     zhangsan     |     The string with which to execute the   search, normally the user name                                                                                                                                                                                                                             |
|     page_size       |     int       |     Optional               |     20           |     The page size. The minimum is 1, the   maximum is 200, and the default is 20.                                                                                                                                                                                                                     |
|     page_token      |     string    |     Optional               |     20           |     Page identifier. This is not required   to get the first page of results. To get the next page of results, pass the   page ID returned in the previous response.     Note that the value of the field here has no particular meaning. Please use   the page ID obtained through the last call.    | ## Response
### Response body

|      Parameters          |      Description                                                                                                                |
|--------------------------|---------------------------------------------------------------------------------------------------------------------------------|
|     code                 |     Response code, anything but 0 indicates   failure                                                                           |
|     msg                  |     Text description of the response code                                                                                       |
|     data                 |     -                                                                                                                           |
|      ∟has_more           |     If there are more users. When true, it   means there is another page                                                        |
|      ∟page_token         |     The page ID, returned when there is   another page. The next call can pass this ID to obtain the next page of users         |
|      ∟users              |     The list of users found                                                                                                     |
|       ∟avatar            |     The user’s avatar                                                                                                           |
|        ∟avatar_72        |     The URL of the user's avatar image, 72   px * 72 px                                                                         |
|        ∟avatar_240       |     The URL of the user's avatar image, 240   px * 240 px                                                                       |
|        ∟avatar_640       |     The URL of the user's avatar image, 640   px * 640 px                                                                       |
|        ∟avatar_origin    |     The URL of the user's avatar image,   original size                                                                         |
|       ∟department_ids    |     The user’s department ID                                                                                                    |
|       ∟name              |     User name                                                                                                                   |
|       ∟open_id           |     User’s open_id                                                                                                              |
|       ∟user_id           |     The user's user_id. This field is only   returned for custom apps that have been granted permission to Obtain the UserID    |
### Response body example
```
{

"code": 0,

"msg": "ok",

"data": {

"has_more": true,

"page_token": "20",

"users": [

{

"avatar": {

"avatar_72":
"https:*//sf6-ttcdn-tos.pstatp.com/img/lark.avatar/d1ca00148ad2c2cf62ed~72x72.png",*

"avatar_240":
"https:*//sf6-ttcdn-tos.pstatp.com/img/lark.avatar/d1ca00148ad2c2cf62ed~240x240.png",*

"avatar_640":
"https:*//sf6-ttcdn-tos.pstatp.com/img/lark.avatar/d1ca00148ad2c2cf62ed~640x640.png",*

"avatar_origin":
"https:*//lf3-ttcdn-tos.pstatp.com/img/lark.avatar/d1ca00148ad2c2cf62ed~noop.png"*

},

"department_ids": [

"od-c02cc3b682a71cdb3a4f14fc4cdb76ac"

],

"name": "zhangsan",

"open_id": "ou_7d8a6e6df7621556ce0d21922b676706",

"user_id": "02a13a9d"

}

]

}

}
```
### Error code

For details, please refer to: Service-side error codes
