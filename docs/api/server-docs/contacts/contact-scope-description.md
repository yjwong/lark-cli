---
title: "Contact Scope Description"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uETNz4SM1MjLxUzM/v3/guides/scope_authority"
section: "Contacts"
updated: "1646791338000"
---

# About the range of contacts data that an app can access
## 1. What is the range of contacts data that an app can access?
The **range of contacts data that an app can access** defines which department and user data the app can access via Contacts APIs. Data beyond the range is inaccessible. By default, the range of contacts data that an app can access is consistent with the app availability. Suppose Contacts contain departments A, B, and C. If the range of contacts data that an app can access contains departments B and C only, this app can access data of departments B and C only. An error message indicating no permission will be returned when this app accesses data of department A.

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/e8f078e4501bb85fe070e14f372c0d09_TxKf0f5Tmt.png?lazyload=true&width=2400&height=1320)
**Example case:**
John has developed a custom app for human resource management. For data confidentiality requirements, this app can be used by HR personnel only. However, this app needs to obtain contacts data of all employees so that it can manage human resource information of the entire company.
In this case, John can set the app availability to HR personnel, and ask the app administrator to make contacts data of all members accessible to this app. In this way, the app can be used only by specified personnel, while it can access all contacts data as needed.

## 2. How do users set the range of contacts data that an app can access?

### For tenant administrators
Tenant super administrators or tenant administrators installing this app can set the range on the app configuration details page in [Admin > Workplace > App Management](http://www.larksuite.com/admin/appCenter/manage).

The following video illustrates the setting method.

![2.gif](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/87e78168f7142c3eedebffec5ede894d_HvpKHLcuZZ.gif?lazyload=true&width=2364&height=1522)

### For app developers
App developers can set the range in Permissions & Scopes in [Developer Console](https://open.larksuite.com/app). The following video illustrates the setting method.

![tongxunlu.gif](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/08c0b3559cedf411cf9da6a7dde03476_5WQfS4XHBl.gif?lazyload=true&width=3092&height=1342)

> **Warning:** Note:
> - By default, the range of contacts data that an app can access is consistent with the app availability.
> - After the range of contacts data that an app can access is modified, an app version needs to be created. The modified range can take effect only after approval of the administrator.
> - Related contacts permissions must be configured to ensure the range configuration can take effect. Remember to apply for required contacts permissions in the Manage scopes section.
> 

## Concept explanation
**Note the differences between the range of contacts data that an app can access and the following concepts.**
- **App availability**App availability defines which users (departments, user groups, and users) can use this app, which is configured by the tenant super administrator or tenant administrator who installs this app in the Admin. To make an app access contacts data of users to which this app is accessible, enable "Keep consistent with the list of allowed app users" in Admin to synchronize the "Range of contacts data that an app can access" from "App availability". In this way, the range of contacts data that the app can access will be identical to the app availability and cannot be modified separately.

![Group 1321314833.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/cdf1f56f0f91f5fdaf60a6b9ed1dc3e5_vEsfNka0PJ.png?lazyload=true&width=1640&height=968?lazyload=true&width=1640&height=944)

- **Visible scope of organization**Visible scope of organization defines the organization scope visible to a user under Organization in Contacts from Lark, which is set by tenant administrators in [Admin > Security > User Permissions > Visible Scope of Organization](http://www.larksuite.com/admin/security/permission/visibility). Visible scope of organization restricts the visible scope of users, while the range of contacts data that an app can access restricts the visible scope of an app.

## 3. Classification of the range of contacts data that an app can access

### User-based contact scope

This scope is required when user delete, update, patch, get, as well as user group add and delete methods are used.
If an app can access contacts data of a department, it can access contacts data of all users under this department.

### Department-based contact scope

The latest version of Contacts categorizes the department-based contact scope into parent department-based scope and department-based scope.
The scope to access parent department contacts data is required when department create, delete, list, update, patch, and parent methods are used.
The scope to access department contacts data is required when department delete, update, patch, get, as well as user create, delete, update, patch, and list methods are used.
If an app can access contacts data of a department, it can access contacts data of all sub-departments under this department.

### All-staff contact scope
The scope to access all staff data is required in the following scenarios:
- Department create method: create a sub-department under a root department
- Department update and patch methods: update a parent department as the root department
- Department list method: obtain sub-departments under the root department
- Department delete method: delete a sub-department from the root department
- Department get method: obtain the root department information
- User create method: create a user under the root department
- User delete method: delete a user from the root department
- User update and patch methods: update a department as the root department
- User list method: obtain users under the root department
- User group create method
