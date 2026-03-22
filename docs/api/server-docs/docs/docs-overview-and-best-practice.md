---
title: "Docs Overview and Best Practice"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uczNzUjL3czM14yN3MTN"
section: "Docs"
updated: "1637144898000"
---

# Docs API Overview
> **Note:** Through this doc, you can learn about cloud doc API capabilities.

## Docs API Introduction
Cloud doc is a series of authoring tools such as doc, table and space. By calling cloud doc API , you can complete a series of automated and batch operations. Cloud doc API consists of the following modules:

::: html
` File Management `: by files, folders, permissions, comments and other modules, you can complete the corresponding operation

- Files: Upload, download, copy files in cloud space
- Folders: Create and manage folders in the cloud space
- Permissions: Manage file permissions in cloud space
- Comments: Get, add, edit comments in the doc

::: html
` Docs `: used to create, obtain and edit online doc 
 Sheets : used to obtain and edit online forms 
 Bitable : used to obtain and edit Bitable 

## Permission Introduction
### App Scope
Before using the permissions of cloud doc, please first ensure that your application has applied for the corresponding permissions. Application permissions can be managed after the developer enters the application details page in the background. For details, please refer to: Get app and tenant access token

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/00a0658829b49d81c89f9b25475cd52e_1Cx269hN3t.png)

### Access Token
> **Note:** Maker sure the token has permission to access the document. Otherwise there will return an error.
::: html
Most interfaces of cloud doc support calling through ` user_access_token ` and ` tenant_access_token ` two identities. However, but still some interfaces that only support one identity. Please read the instructions of the corresponding interface doc carefully before calling. 

Take four scenarios to illustrate the use of access token: (take Batch update document as an example)
| User scenario & AccessToken | Realization method |
| --------- | --------- |
| Editing own document with user_access_token | 1) Get own user_access_token 2）Get the filetoken（see Introduction）3）Call Batch update document|
| Editing other people's document with user_access_token | 1) Get other people‘s  user_access_token 2）Get the filetoken（see Introduction）3）Call Batch update document|
| Editing own document with tenant_access_token | 1) Get the app's tenant_access_token 2）Get the filetoken（see Introduction）3）Call Batch update document|
| Editing other people's document with tenant_access_token | 1) Get the app's tenant_access_token 2) other people authorize the doc to the app  3）Get the filetoken（see Introduction）4）Call Batch update document| ## Request return value
The return value structure is as follows, if the error, please query the server error code description.
```json
{
    "code": 200,
    "msg": "sample err msg",
    "data": {}
}
```
