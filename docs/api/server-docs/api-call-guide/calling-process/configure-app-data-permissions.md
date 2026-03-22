---
title: "Configure app data permissions"
url: "https://open.larksuite.com/document/home/introduction-to-scope-and-authorization/configure-app-data-permissions"
resource: "home"
section: "API Call Guide"
updated: "1691064939000"
---

# Configure app data permissions

Application data permissions refer to the range of data that can be obtained when accessing business resources as an application. After the application has applied for some API permissions (for example, Contacts), it also needs to configure the corresponding data permissions and submit them for review. After the permission is approved, the API can be successfully called to obtain data. Otherwise, a permission error will be returned when calling the API.

## Introduction to Data Permissions

Both enterprise-built apps and store apps support configuring data permissions, You can use this chapter to learn about the details and configuration methods of data permissions for various applications.

### Data permissions for self-built applications

Self-built apps support configuring data permissions for **Contacts**.

| Data permission type | Permission Description | Management style |
| --- | --- | --- |
| Contacts Permission Scope | When calling the address book API as an application, the range of address book data that the application can obtain. For example, when calling the address book API to query user A's information as an application, the application needs to have the data permission of user A. By default, the permission scope of the address book is consistent with the availability scope of the application. | - Method 1: The app owner, administrator, or development role configures the address book permission range within the app. - Method 2: The enterprise administrator configures the address book permission scope of the specified application in the management background. | > **Note:** Organizational structure data such as departments and users are sensitive data for enterprises. Once the app authorizes the corresponding data permissions, it can add, delete, modify, and query these sensitive data. Therefore, after the app developer configures the data permissions, he needs to publish the app and wait for the app reviewer to pass the review before the data permissions can take effect.

### Data Permissions for Store Apps

The store app only supports the enterprise administrator to configure the data permissions of **Contacts**, and does not support the configuration of data permissions of other open capabilities. App developers can call Get Address Book Authorization Scope to view the data permissions of the app.

| Data permission type | Permission Description | Management style |
| --- | --- | --- |
| Contacts Permission Scope | When calling the address book API as an application, the range of address book data that the application can obtain. For example, when calling the address book API to query user A's information as an application, the application needs to have the data permission of user A. | The enterprise administrator configures the address book permission range of the specified application in the management background. | ## Enterprise self-built application

If the application you develop involves calling the Contacts API or Lark Personnel (Enterprise Edition) API as an application, you need to configure the corresponding data permissions.

> **Note:** - Data permissions need to be combined with API permissions to take effect, that is, before configuring data permissions, you need to ensure that you have applied for the corresponding API permissions for the application. For details, see Apply for API permissions.
> 
> - For more concept introduction about the scope of authority of the address book, see Detailed explanation of scope of authority.
> 
> - If you use the Testing Enterprise and Personnel capability of the app and switch the app to the test version for joint development and debugging, the app has the address book permission of all employees by default. 

### Procedure

1. Log in to [Developer Consle](https://open.larksuite.com/app). Find and enter the specified application details page.

2. In the left navigation bar, select **Development Configuration** > **Permissions & Scopes**.

1. In the **Range of contacts data to access** area, click **Configure**.

	![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/aa1a9d031ae096a6a6580a9793a7ccce_RxQmUCFPow.png?height=552&lazyload=true&maxWidth=600&width=1060)

2. Configure the permission range and click **OK**.
    
     By default, the data permission scope of the address book is **consistent with the available scope of the application**.

	![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/adbc12564a07cd5ef8331e857a6490e3_Cl1T3ZRaHk.png?height=570&lazyload=true&maxWidth=600&width=1810)

     - **Identical to the availability of app**: The permission scope of the address book is consistent with the available scope of the application. For a detailed introduction to application scope and authorization, see Configure Application Availability Scope.
    
     - **Some members**: After selecting this range, you also need to manually select a department or member to confirm the permission range of the address book.
    
     - **All Members**: The address book permission of all members in the company.

> **Note:** The enterprise administrator can find the specified application in the workbench of [Lark Admin](https://www.larksuite.com/admin) and modify the application's address book permission range.

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/d78c833ca7a751a61a9c397328dedc8d_KpgKgMbvxM.png?height=714&lazyload=true&maxWidth=600&width=2338)

## Store App

The store application only supports the enterprise administrator to find the specified application in the workbench of [Lark Admin](https://www.larksuite.com/admin) and modify the application's address book permission range.

> **Note:** App developers can call Get Address Book Authorization Scope to view the data permissions of the app.

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/7f25ab961a023c281077f205e32c3fdc_rCULw0ql44.png?height=628&lazyload=true&maxWidth=600&width=2334)
