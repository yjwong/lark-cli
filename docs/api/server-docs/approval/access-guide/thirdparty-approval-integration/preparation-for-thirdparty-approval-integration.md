---
title: "Preparation for third-party approval integration"
url: "https://open.larksuite.com/document/ukTMukTMukTM/ukjNyYjL5YjM24SO2IjN/tripartite-approval-and-access-preparation"
section: "Approval"
updated: "1687677440000"
---

# Preparation for third-party approval integration

## Sequence diagram
The following flow chart shows how to integrate a third-party system into the Lark Approval Center.

![0819时序图.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/c5ed2bccff4370950206e995a3d8f2dd_qimdqAOfi2.svg?lazyload=true&width=771&height=713)
## Preparations
###  Integration method
The Approval Center enables these capabilities via the [Lark Open Platform](https://open.larksuite.com/) APIs, and the third-party systems push the approval data to the Open Platform via the APIs.
###  Open Platform app
To use the Approval Center API, you need to apply for a company custom app on the Open Platform. Refer to Develop an app. The [Access Approval] permission scope needs to be enabled for the Approval Center API, and the new permission scope will only take effect after a new app version is released.

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/6a20db97961b536e46d858954e936886_oztufgEkoP.png?lazyload=true&width=3818&height=1686)

For authentication methods of the Open Platform APIs, refer to the API access token overview. The processes are as follows::

1. Obtain app_id and app_secret
2. Obtain app_access_token by using app_id and app_secret
3. Access the Open Platform approval API via app_access_token（The app_access_token of the custom app is the same as the tenant_access_token）

###  User system
The Approval Center uses the open platform user system. The user_id in the open platform API of the Approval Center is the user_id in Lark Open Platform. Refer to Terms for relevant definitions. Refer to Contacts for the method of obtaining user_id and department_id.
