---
title: "Advanced permission overview"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-role/advanced-permission-guide"
service: "bitable-v1"
resource: "app-role"
section: "Docs"
updated: "1753668118000"
---

# Overview of advanced permissions

**Owners** or collaborators with **management permissions** of a Base can set advanced permissions through the open platform interface, including setting custom roles and managing collaborators. This allows for specifying read or edit permissions for designated rows and columns for specific personnel in each Base.
## Precautions

- Before calling the custom roles and collaborator-related interfaces of advanced permissions, ensure that advanced permissions have been enabled for the Base. You can enable advanced permissions through the Update Base metadata interface.
- There is a delay in enabling advanced permissions. If you encounter an `OperationTypeError` error when calling the advanced permissions interface immediately after enabling advanced permissions, please try again later.
- The collaborators of advanced permissions are different from the collaborators of cloud document permissions. After adding advanced permissions collaborators, to ensure the normal setting of cloud document permissions, it is recommended to add cloud document permissions through the Add collaborator permissions interface.
- After enabling advanced permissions, you need to set custom roles first before adding collaborators.

## Usage restrictions

- Bases embedded in online documents and spreadsheets, and Bases in knowledge bases do not support enabling advanced permissions.
- The advanced permissions interface for Base does not currently support setting dashboard permissions.
## Resource description
The description of custom roles and collaborator resources in advanced permissions is as follows.
### Custom role

Add roles and set permissions in advanced permissions, and the role becomes a custom role. Each custom role has a unique identifier `role_id`. `role_id` needs to be obtained through the List custom roles interface.

### Collaborator

In advanced permission settings, a member in a custom role (role) is a collaborator (member). Each collaborator has a unique identifier `member_id`. `member_id` needs to be obtained through the List collaborators interface.
## Method list

The following is a list of methods for advanced permissions. "Store" represents store applications; "Self-built" represents enterprise self-built applications. For more information about applications, refer to Application types introduction. For the process of calling server-side APIs, refer to Process overview.

### Role

| `List roles `GET` /open-apis/bitable/v1/apps/:app_token/roles ` | `bitable:app` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| --- | --- | --- | --- | --- |
| `Create role` `POST` /open-apis/bitable/v1/apps/:app_token/roles/:role_id | `bitable:app` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Update role` `PUT` /open-apis/bitable/v1/apps/:app_token/roles/:role_id | `bitable:app` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Delete role` `DELETE` /open-apis/bitable/v1/apps/:app_token/roles/:role_id | `bitable:app` | `tenant_access_token` `user_access_token` | **✓** | **✓** | ### Member
| `List members `GET` /open-apis/bitable/v1/apps/:app_token/roles/:role_id/members ` | `bitable:app` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| --- | --- | --- | --- | --- |
| `Create member` `POST` /open-apis/bitable/v1/apps/:app_token/roles/:role_id/members | `bitable:app` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Delete member` `DELETE` /open-apis/bitable/v1/apps/:app_token/roles/:role_id/members/:member_id | `bitable:app` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Batch add collaborators` `POST` /open-apis/bitable/v1/apps/:app_token/roles/:role_id/members/batch_create | `bitable:app` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Batch delete collaborators` `DELETE` /open-apis/bitable/v1/apps/:app_token/roles/:role_id/members/batch_delete | `bitable:app` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
