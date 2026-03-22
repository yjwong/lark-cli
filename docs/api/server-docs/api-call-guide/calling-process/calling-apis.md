---
title: "Calling APIs"
url: "https://open.larksuite.com/document/ukTMukTMukTM/ukDNz4SO0MjL5QzM/get-"
section: "API Call Guide"
updated: "1699602841000"
---

# Calling APIs

To facilitate developers in calling the API conveniently, Lark Open Platform provides Java SDK, Golang SDK, and NodeJS SDK. For detailed information about the SDKs, see Server SDK. This article introduces the basic API calling methods.

## Prerequisites

- Before calling the server APIs, you need to complete the steps of creating an app, applying for permissions, obtaining an access token, and setting IP allowlist.
- When calling the server API, you need to include the access token in the request header `Authorization: Bearer `.
- When calling the server API, you need to use the HTTPS protocol and UTF-8 encoding.

## Calling examples

### Sending messages to employees within an organization

You can use the Send message interface to send messages to employees within an organization. Based on the interface documentation, prior to calling this interface, you need to obtain the `tenant_access_token`.

1. Refer to Obtaining an access token to obtain the `tenant_access_token`.

   The example request for obtaining the token is as follows. Please replace `app_id` and `app_secret` with the actual values.
   

   ```
   curl -X POST 'https://open.larksuite.com/open-apis/auth/v3/tenant_access_token/internal/'
   -H 'Content-Type: application/json; charset=utf-8'
   -d '{
     "app_id": "",
     "app_secret": ""
   }'
   ```

2. Based on the request parameter description in the documentation, call the Send message interface.

   - Method 1: Initiate API call in the API explorer

     ![](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/df0679d85b2b7a0bc5eea33692e2d5df_647GS3jnKI.png?height=329&lazyload=true&maxWidth=600&width=956)

   - Method 2: Send a curl request locally

     This API needs to be called using the POST method.

     ![](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/ea137cab4d70f089f9c7daa7e7f96d32_jQLCARdMWP.png?height=368&lazyload=true&maxWidth=600&width=666)

     The following is an example, please replace the parameter values with the actual ones.

     ```HTTP
     curl -X POST 'https://open.larksuite.com/open-apis/im/v1/messages?receive_id_type=user_id'
     -H 'Authorization: Bearer '
     -H 'Content-Type: application/json; charset=utf-8'
     -d '{
         "content": {
              "text": "Hello World"
         },
         "msg_type": "text",
         "receive_id": ""
     }'
     ```

     - `receive_id_type` is used as a query parameter.
     - `content`, `msg_type`, and `receive_id` are included in the request body.
     - The required `tenant_access_token` and `Content-Type` are included in the header.

### Querying user information

You can use the Get user information interface to query user information. Based on the interface documentation, prior to calling this interface, you need to obtain either the `tenant_access_token` or `user_access_token`. Choose the appropriate access token based on the scope of user information you need.

1. Refer to Obtaining an access token to obtain the `tenant_access_token` or `user_access_token`.

2. Based on the request parameter description in the documentation, call the Get user information interface.

   - Method 1: Initiate API call in the API explorer

     ![](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/10e69d12ee1cfa49104c95723f77d0e7_qVfCoqlAY8.png?height=158&lazyload=true&maxWidth=500&width=637)

   - Method 2: Send a curl request locally

     From the interface documentation, we can see that `user_id` is a path parameter for this interface. For example, to query user information for a `user_id` of `ou_c99c5f35d542efc7ee492afe11af19ef`, use the following example:

     ```
     curl -X GET 'https://open.larksuite.com/open-apis/contact/v3/users/ou_c99c5f35d542efc7ee492afe11af19ef?user_id_type=user_id' \
     -H 'Authorization: Bearer ' \
     -H 'Content-Type: application/json; charset=utf-8'
     ```

## Response result

The response structure of most APIs consists of three parts: `code`, `msg`, and `data`:
- `code`: Error code. If the response is successful, `code` is set to 0.
- `msg`: Error message. 
- `data`: Result of the API call. In some operation APIs, `data` may not exist.

> **Note:** - Do not rely on `msg` to determine if a request has failed.
> - After receiving a failed response, you can first refer to the Common error codes for troubleshooting. If the issue persists, you can provide the `x-tt-log` `id` value in the response header to the Lark Open Platform for assistance in problem diagnosis.

Example of a successful response:
```
{
  "code": 0,
  "msg": "success",
  "data": {
    // Specific data content in the response
  }
}
```
Example of a failed response:
```
{
  "code": 40004,
  "msg": "no dept authority error"
}
```
