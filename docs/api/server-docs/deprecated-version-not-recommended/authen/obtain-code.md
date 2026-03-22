---
title: "Obtain code"
url: "https://open.larksuite.com/document/ukTMukTMukTM/ukzN4UjL5cDO14SO3gTN"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/authen/v1/index?redirect_uri={REDIRECT_URI}&app_id={APPID}&state={STATE}"
section: "Deprecated Version (Not Recommended)"
updated: "1750924432000"
---

# Obtain Authorization Code

> This API is now deprecated. Its use is not recommended. Please use the latest version: Get Authorization Code.

This interface is called to obtain the user's login pre-authorization code, which is valid for 5 minutes and can only be used once.

> **Note:** **Note:** This interface is suitable for web browser. When users open a web application, it will jump to authorization interface. After user authorization, the browser will automatically jump to the redirect URL and carry the login pre-authorization code issued by Lark Open Platform. For more details about the redirect URL, please refer to Obtain user_access_token.

> **Note:** **Update**: The new authorization process supports the ability of developers to implement the incremental authorization by passing in the scope parameter when obtaining the authorization code. Scopes required by user_access_token can be granted only after the user agrees on the authorization page. You are advised to access the new version with more standardized and compliant procedures.

## Request
| HTTP URL | https://open.larksuite.com/open-apis/authen/v1/index?redirect_uri={REDIRECT_URI}&app_id={APPID}&state={STATE} |
| --- | --- |
| HTTP Method | GET |
| Interface frequency limit | 1000 req/min、50 req/s |
| Supported app types | custom,isv |
| Required scopes | None | ### Query parameters
 
**Parameter** | **Type** | **Required** | **Description**                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| ------------- | -------- | ------------ | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| redirect_uri  | string   | Yes          | The application Redirect URI  must be encoded in the format of "application/x-www form urlencoded". Before calling this interface, it is necessary to configure the redirect URI on the **security settings** page of the developer's backend, which supports configuring many URLs. Only URLs in the redirect URL list will pass open platform security verification |
| app_id        | string   | Yes          | The `app_id` can be found on the **Credentials & Basic Info** page of the Developer Console. For more information about the `app_id`, please refer to the General parameters                                                                                                                                                                                                                                                                                                                             |
| state         | string   | No           | An additional string used to maintain the request and callback status, which will be appended to the redirect uri. The application can determine the contextual relationship by this parameter

### Request example

```shell 
GET https://open.larksuite.com/open-apis/authen/v1/index?app_id=cli_9dff7f6ae02ad104&redirect_uri=https%3a%2f%2fopen.larksuite.com%2fdocument%2fuQjL04CN%2fucDOz4yN4MjL3gzM&state=RANDOMSTATE
``` 

### Response example

```
/ssl:ttdoc/uQjL04CN%2fucDOz4yN4MjL3gzM?code={AuthorizationCode}&state=RANDOMSTATE 
```

|**Name**|  **Description**|
| ---	|  --- 	|
| code | OAuth Code for User Access Token |
| state | state value in req |
