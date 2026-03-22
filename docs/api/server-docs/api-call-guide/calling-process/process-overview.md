---
title: "Process overview"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uITNz4iM1MjLyUzM"
section: "API Call Guide"
updated: "1708226612000"
---

# Process overview

The APIs provided by Lark Open Platform follow the RESTful style, and the format of the request URL is as follows:

![](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/379b0797ca0cd9ed086efce6fd4d0b46_umwPiqSjkc.png?height=256&lazyload=true&maxWidth=750&width=2044)

The process of calling server APIs is illustrated in the following diagram:

![API概述en.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/767cf78bcf0d9ebcec219621b0f0a3b0_ryI2fGTexu.png?height=214&lazyload=true&maxWidth=800&width=2920)

- Create apps. In the [Developer Console](https://open.larksuite.com/app), create custom apps or store apps according to actual needs.

> **Note:** Only users with ISV qualifications can create store apps. For more information , see How to join Lark Open Platform as an ISV.

- Obtain an access token. Lark Open Platform provides various types of access tokens, which represent different resource access permissions. When calling APIs, the access token needs to be included in the HTTP header to access resources within the defined permissions.
	
- Apply for API scopes. If you want to call the API, you must get the scope to call the interface first. If it involves accessing sensitive fields, you also need to get the scope to access sensitive fields.

- (Optional) Configure data permissions. When an app requests permissions for certain APIs (e.g.,  Contacts, Lark HR Enterprise Edition), it needs to configure the corresponding data permissions and submit them for review. The permissions take effect after the review is approved, enabling successful data retrieval through API calls. Otherwise, an authorization error will be returned when invoking the API.

- (Optional) Set IP whitelist. To enhance app security, you can set an IP whitelist for your app. Only requests originating from whitelisted IP addresses will be responded to by the Lark Open Platform. Otherwise, the requests will be rejected.

- Call the API. After completing the above steps, you can call the open API. Users can refer to the API documentation to understand the specific functions of the API.
