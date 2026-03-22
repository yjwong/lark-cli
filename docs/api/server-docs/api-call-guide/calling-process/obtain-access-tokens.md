---
title: "Obtain access tokens"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uMTNz4yM1MjLzUzM"
section: "API Call Guide"
updated: "1724147829000"
---

# Obtain access tokens

To improve API security, Lark uses an access token mechanism. When calling APIs, the access token authenticates the caller's identity and permissions. This informs Lark who is accessing app data and resources.
 
The access token is the key to accessing the Lark Open Platform, binding the caller identity, as well as all data access and API call permissions obtained by the app together, allowing the app to perform read and write operations on resources. It is recommended that developers have a thorough understanding of Lark's access token mechanism before starting formal development.

## Introduction to access tokens

  
The Lark Open Platform provides different types of access tokens for authenticating the calling party's identity. In the actual development process, you can choose the appropriate access token based on your business needs.

| Access Token Type | Requires User Authorization | Description |
| --- | --- | --- |
| tenant_access_token | No | The credentials required to call the API as an app. The range of readable and writable data is determined by the app's own data permission range. The value of this type of credential is prefixed with `t-`, and the example value is: `t-24b5bf4e00b2af1234`. |
| user_access_token | is | The credentials required when calling the API as a user. The range of readable and writable data is determined by the user's  data permission range. The value of this type of credential is prefixed with `u-`, and the example value is: `u-Lr1RT7S8fS03mT1234`. |
| app_access_token | No | The short-term token for app identity, the open platform identifies the caller's app identity based on the app_access_token. The value of this type of credential is prefixed with `a-` or `t-`, and the example value is: `a-24b5cef00b1234`. | ## How to choose different types of access tokens

> **Note:** APIs may support one or multiple types of access tokens. Different access tokens represent different resource access permissions. Therefore, different access tokens may provide different data information.

     
1. Confirm the supported access tokens for the API.
  
    When calling the API, the access token needs to be filled in the HTTP request header. In the request header section of the Lark Open Platform API documentation, you can check the supported access token types for that API.    

	![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/450f38a91c6d18ad9eb1b9110b5dcd16_5rMcLPefvM.png?height=335&lazyload=true&maxWidth=600&width=932)

2.  Choose the appropriate access token.

> **Warning:** Since `app_access_token` is rarely used (generally used to get user_access_token and refresh user_access_token), Lark Open Platform is gradually unifying the two credentials `app_access_token` and `tenant_access_token`. Therefore, `tenant_access_token` and `user_access_token` are mainly introduced here.
    
- **tenant_access_token**
        
    If your business logic does not require operating on user data resources and only needs to operate on the app's own resources, such as creating a cloud document in the app's document directory space, it is recommended to use the `tenant_access_token` without requiring additional authorization.
    
- **user_access_token**
        
	If your business logic needs to operate the user's data resources, for example, you need to create a cloud document under the user's document directory space, it is recommended to use `user_access_token` without additional application for authorization.
        
    If the `tenant_access_token` is used in this case, you need to additionally add the corresponding authorization for the app at the resource level. For example, if you call the contacts API using the `tenant_access_token`, you need to configure the app's contacts permission range on the Lark Developer Console's permission management page. However, if you use the `user_access_token`, you do not need to configure the contacts permission range separately. It follows the contacts permission range of the user_access_token owner.

## Obtain access tokens

  
This section describes how to obtain the corresponding access tokens for custom apps and store apps.
 
> **Note:** Access credentials have a validity period (after calling the access credential acquisition interface, the validity period field will be returned). Developers need to set up business logic to regularly refresh credentials on their own servers to prevent expiration. If you re-obtain the voucher before the current voucher expires, you will get the same value as the original value, but the remaining validity time will change each time it is returned, as shown in the figure below.

![](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/dba69fbde5e00970fdb60fdcba34bcc5_Q09daiPiIg.png?height=650&lazyload=true&maxWidth=600&width=2624)

### Obtain tenant_access_token for custom apps

  
1. Log in to the [Developer Console](https://open.larksuite.com/app/) and select the specified app.
  
2. On the **Credentials & Basic Info** page, obtain the **App ID** and **App Secret** of the app.

	![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/79114f27c5bf3b6c62b3df21e15b0fbf_p5UvswMTUd.png?height=376&lazyload=true&maxWidth=600&width=2842)
  
3. Call the obtain tenant_access_token for custom apps API and get the `tenant_access_token` for the custom app using the app credentials.

### Obtain app_access_token for custom apps

1. Log in to the [Developer Consle](https://open.larksuite.com/app/), and select the specified app.

2. On the **Credentials & Basic Info** page, obtain the app credentials **App ID** and **App Secret**.
    

	![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/79114f27c5bf3b6c62b3df21e15b0fbf_K208kpPRwP.png?height=376&lazyload=true&maxWidth=600&width=2842)

3. Call the obtain app_access_token for custom apps interface, and obtain the `app_access_token` of the custom app through the using the app credentials.

### Obtain app_access_token for store apps

1.  Obtain the `app_ticket`.
  
	1. Configure the event subscription request URL for the store app.
  
        For detailed steps, see Configure request URL. After completing the configuration, Lark will automatically push the latest `app_ticket` to the specified URL every hour. The app should handle the received app_ticket event to obtain and save the `app_ticket` for subsequent use.
    
      2. (Optional) Call the API for resending `app_ticket` to trigger Lark to resend the `app_ticket`.
      
      There may be delays in the push of `app_ticket`, so if you do not receive the push, you can use this method to resend the `app_ticket`.

2. Obtain the app credentials: App ID and App Secret.
    
     1. Log in to the [Developer Consle](https://open.larksuite.com/app/), and select the specified store app.
    
     2. On the **Credentials & Basic Info** page, obtain the **App ID** and **App Secret**.
        

		![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/79114f27c5bf3b6c62b3df21e15b0fbf_HVzsIfjHgN.png?height=376&lazyload=true&maxWidth=600&width=2842)

3. Call the obtain app_access_token for store apps  API to obtain the `app_access_token` of the store app using the app ID, app Secret, and app_ticket.

### Obtain tenant_access_token for store apps

  
> **Note:** Please obtain and save the `app_access_token` for the store app before obtaining the `tenant_access_token`.

1.  Obtain the `tenant_key`.
  
    `tenant_key` is the unique identifier for the tenant on Lark, used to obtain the corresponding `tenant_access_token`. It can also be used as a unique identifier for the tenant within the app and can be obtained in the following ways:
  
    - When an enterprise activates the app, Open Platform pushes the `tenant_key` data to the app. For details, see App enabled initially.

    - When a user logs into the gadget, H5 app, or browser app, obtain it from the user's identity information.
  
2. Call the obtain tenant_access_token for store apps API and get the `tenant_access_token` for the store app using the `app_access_token` and `tenant_key`.

### Obtain user_access_token

This section uses the web app as an example to illustrate how to obtain the `user_access_token`. 
  
> **Note:** The method for obtaining `user_access_token` is applicable to both custom and store apps.

1.  Log in to the [Developer Console](https://open.larksuite.com/app/) and select the specified app.
  
2. On the **Security** **Settings** page, configure the redirect URL.
  
    You need to configure the public address of the server that actually initiates the API request as the redirect URL of the application. Multiple redirect URLs can be configured. Only URLs in the redirect URL list will pass the security verification of the open platform when requested.
    
  	![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/499f98e70548ecc579991790f56b476e_BUMAmWoUgc.png?height=532&lazyload=true&maxWidth=600&width=2850)
3. Call the Obtain OAuth code API to obtain the login pre-authorization code.
  
    When calling the API, you need to use the app credentials App ID. You can obtain the app credentials App ID on the **Credentials & Basic Info** page.

    ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/79114f27c5bf3b6c62b3df21e15b0fbf_RcNJUgMsQu.png?height=376&lazyload=true&maxWidth=600&width=2842)
  
4. Call the obtain user_access_token API to obtain the `user_access_token`.
	
    At the same time, the validity period of the Token can be obtained (expires_in response field).

5. (Optional) Call the refresh user_access_token interface to refresh `user_access_token`.
    
     The validity period of `user_access_token` is about two hours. After expiration, it needs to be refreshed before use.
     
## How to use access tokens

To use an access token (access_token), you need to fill it in the **Request Header** parameter of the HTTP request when calling the API. The value needs to be prefixed with `Bearer`.
  
> **Note:** For security reasons, do not use access tokens at the front end of the app. Please initiate API access requests from the app server.

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/334579edc6370e34bb91961cbef94c4a_AXdxY0Mgad.png?height=428&lazyload=true&maxWidth=600&width=1972)

Each OpenAPI document declares the access token type and acquisition method required by the interface.

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/c8014b868f84b9f3fa4f2396dc247ccf_THrsduNUnj.png?height=616&lazyload=true&maxWidth=600&width=1720)

For example, a sample cURL request configuration is as follows:

> **Note:** For details on how to call the API, see Call API.

```Bash
curl --location --request GET 'https://open.larksuite.com/open-apis/contact/v3/departments/od-64242a18099d3a31acd24d8fce8dxxxx' \
--header 'Authorization: Bearer t-5888bcb7c421a695ca83c449caba5cxxxx' \
--header 'Content-Type: application/json; charset=utf-8'
```
