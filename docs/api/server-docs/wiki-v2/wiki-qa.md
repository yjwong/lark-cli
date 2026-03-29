---
title: "API FAQ"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/wiki-qa"
service: "wiki-v2"
resource: "wiki-qa"
section: "Docs"
updated: "1660708283000"
---

# API FAQ

## How to edit a node document

Wiki nodes support a variety of document types, including documents, spreadsheets, multi-dimensional tables, thinking notes and files, etc. You can edit some types of nodes through the API. Currently, the types that support editing through the API are: documents, spreadsheets, multi-dimensional tables.

### Operation steps

By calling, [get node information](https://bytedance.feishu.cn/ssl:ttdoc/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space-node/get), interface, can be obtained from the return value to `objType` and `objToken`.

Depending on the type of node, `objToken` may be one of doc, sheet, and bitable. ObjToken is the corresponding `docToken, speadsheetToken, and bitable appToken`.

You can call the interfaces of each service through the corresponding tokens to add, delete and modify the content. For details, please refer to:

See Docs FAQ for more details.

## How to call API using tenant_access_token

::: html

 In the Wiki API, except for [Creating Workspace ](https://bytedance.feishu.cn/ssl:ttdoc/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/faq/create), most APIs support calling using ` tenant_access_token `. 

### Operation steps

::: html

 Before using  tenant_access_token , make sure that this application has been added as one of the workspace members, otherwise it will return No permission error.

The only way to add an app as a member of the workspace is:

-   Get the  user_access_token  of any workspace administrator who has permission to add workspace members.

-   Use this token to call the [Adding Workspace Member ](https://bytedance.feishu.cn/ssl:ttdoc/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space-member/create)interface to add the app as one of the members of the workspace. And ensure that the app, as a workspace member, has permission to get the node list and node content of the workspace.

-   After ensuring that the above steps are completed correctly, you can use  tenant_access_token  to call the relevant interface of the workspace.
