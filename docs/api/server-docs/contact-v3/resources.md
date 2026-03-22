---
title: "Contacts overview"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/resources"
service: "contact-v3"
resource: "resources"
section: "Contacts"
updated: "1648389906000"
---

# Overview

## Services

On the Lark open platform, Contacts indicates the organizational structure of a company, including department and personnel information. With a series of secure, reliable Contacts APIs provided, you can do the following:

- View the organizational structure of a company

- Create users for a company

- Modify basic attributes of existing users in a company

- Maintain the relationship between users and departments

- Maintain the relationship between users and user groups

### How to integrate

|  | Steps | Introduction |
| --- | --- | --- |
| 1 | Create an app | - For creating a custom app, see Custom app development process. -   For creating a store app, see Develop and publish store apps. |
| 2 | Call APIs to operate Contacts. | Before calling APIs, you need to obtain the access token and enable corresponding permission scopes. For details, refer to How to call server-side APIs. You can also quickly debug these APIs using Postman: [![Run in Postman](https://run.pstmn.io/button.svg?lazyload=true&width=123&height=30)](https://god.gw.postman.com/run-collection/17195890-6fd42609-12b3-409b-8cbb-73656cf9d805?action=collection%2Ffork&collection-url=entityId%3D17195890-6fd42609-12b3-409b-8cbb-73656cf9d805%26entityType%3Dcollection%26workspaceId%3D55edcb9c-bbfe-45f5-a74a-efb882fe5384#?env%5Bfeishu-demo%5D=W3sia2V5IjoiYmFzZVVybCIsInZhbHVlIjoib3Blbi5mZWlzaHUuY24iLCJlbmFibGVkIjp0cnVlfSx7ImtleSI6InN0b3JlX2FwcF9pZCIsInZhbHVlIjoiY2xpX2EwYTUwODBjODRiODUwMTMiLCJlbmFibGVkIjp0cnVlfSx7ImtleSI6InN0b3JlX2FwcF9zZWNyZXQiLCJ2YWx1ZSI6InpFb0xLNHBjZE45VXBsNUpDOGNjcGZORlI3Q1FpbmNhIiwiZW5hYmxlZCI6dHJ1ZX0seyJrZXkiOiJhcHBfdGlja2V0IiwidmFsdWUiOiI3YTU1NjFlODdiMjkzYjFjOTEyZWM1NTQ2MDVjNDFlOWZhMjZkYzJmIiwiZW5hYmxlZCI6dHJ1ZX0seyJrZXkiOiJhcHBfYWNjZXNzX3Rva2VuIiwidmFsdWUiOiJhLWNlOTJjZTNhMmRjNmM2ZjQzYTVjNzM2YmRlMzAxM2FkYzdlZGM2MzQiLCJlbmFibGVkIjp0cnVlfSx7ImtleSI6InRlbmFudF9rZXkiLCJ2YWx1ZSI6IjczNjU4OGM5MjYwZjE3NWQiLCJlbmFibGVkIjp0cnVlfSx7ImtleSI6InRlbmFudF9hY2Nlc3NfdG9rZW4iLCJ2YWx1ZSI6InQtMmQ0OWY4ZjMyOTY2YTEzYmMzN2ZiMWJkZWFmZTBkNDdhMjAwZDZkZiIsImVuYWJsZWQiOnRydWV9LHsia2V5IjoiU1RBVEUiLCJ2YWx1ZSI6IjExIiwiZW5hYmxlZCI6dHJ1ZX0seyJrZXkiOiJSRURJUkVDVF9VUkkiLCJ2YWx1ZSI6Imh0dHBzJTNBJTJGJTJGd3d3LmJhaWR1LmNvbSUyRiIsImVuYWJsZWQiOnRydWV9LHsia2V5IjoidXNlcl9hY2Nlc3NfdG9rZW4iLCJ2YWx1ZSI6InUtMDJvYmhid01DSEl5c2ZhNzFWVGNUZCIsImVuYWJsZWQiOnRydWV9LHsia2V5IjoiYXBwX2lkIiwidmFsdWUiOiJjbGlfYTA3ZmM0ZDVhMmY5NTAwYyIsImVuYWJsZWQiOnRydWV9LHsia2V5IjoiYXBwX3NlY3JldCIsInZhbHVlIjoiVlpBcFd0ZXc2UUdHQm1SbmxJNTF2aEZtbUU0bkJScmwiLCJlbmFibGVkIjp0cnVlfV0=). For usage methods, see Postman template manual. |
| 3 | Listen events and obtain changes in Contacts. | You need to apply for relevant permission scope before listening on events. For details, see Event subscription overview. | > **Warning:** **For privacy security reasons, the following permission scope control requirements are configured for Contacts APIs:**
> - Departments and personnel that can be searched or modified by an app depend on the range of contacts data that an app can access, which is configured by the [company administrator](https://www.larksuite.com/hc/zh-CN/articles/360049067822). For details, see About the range of contacts data that an app can access.
> 
> - App store apps are unable to modify the contacts data of the company to which a user belongs.

## Resources

Contacts business domains are open around resources. The following figure illustrates the relationship among resources:

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/06db4625721248cf49ed59a2703f6f2e_f5G2YeBw6s.png?lazyload=true&width=1640&height=1284)
The following table describes different resources.
| User | A user of a company in Lark. |
| --- | --- |
| Department | A node in the organizational structure tree of a company |
| User group | Another way to group users besides department. User groups can be associated with users and are divided into [static groups](https://www.larksuite.com/hc/zh-CN/articles/360049067479) and [dynamic groups](https://www.larksuite.com/hc/zh-CN/articles/360049067874). |
| Workforce type | A special user attribute, used to mark the identity type of a user. Default values include employee, intern, outsourcing, contractor, and consultant. Enumerated values of this field can be customized for each company based on its actual conditions, and each user can be associated with only one enumerated value. |
| Custom user fields | User fields customized by a company, used to flexibly reflect user information. Custom user fields are created and updated by the [company administrator](https://www.larksuite.com/hc/zh-CN/articles/360049067822) in [Admin > Organization > Member Profile](http://www.larksuite.com/admin/contacts/employee-field-new/custom). You can obtain all custom fields of a company via the specific API. ![image (1).png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/d272ea8ae44d9c31b353e44186fbe2a9_ll3uSIwIGR.png?lazyload=true&width=1774&height=828?lazyload=true&width=1640&height=774) ![image (2).png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/fcf68d75fadfb5f50f267603976781a3_O79A5Bd69k.png?lazyload=true&width=1773&height=829?lazyload=true&width=1640&height=937) |
| Range of contacts data that an app can access | This resource refers to the scope of the organizational structure data that an app can access via the contacts API, which is defined by the [company administrator](https://www.larksuite.com/hc/zh-CN/articles/360049067822) in [Admin > Workplace > App Management](http://www.larksuite.com/admin/appCenter/manage). ![Group 1321314831.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/0adc604608902de9d887a0c571cb44a5_hpaRbffgyQ.png?lazyload=true&width=1640&height=1339) ![Group 1321314832.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/f1772a032c4ef8467d27bf889a3c2137_Hhk3bdgJ0m.png?lazyload=true&width=1640&height=937?lazyload=true&width=1277&height=436) | Below are details for fields, methods and events of each resource.

### Resource: User
View Resource fields and examples.

#### Methods
>  "Store" represents Store Apps; "Custom" represents Custom Apps.
| `Create a user `POST` /open-apis/contact/v3/users > Create a user in the Contacts, used for employment ` | Update Contacts | `tenant_access_token` | **X** | **✓** |
| --- | --- | --- | --- | --- |
| `Obtain single user information` `GET` /open-apis/contact/v3/users/:user_id > Obtain information about a single user in Contacts | Read contacts as an app | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Obtain the list of users directly under a department` `GET` /open-apis/contact/v3/users/find_by_department > Obtain the list of users directly under a department | Read contacts as an app Obtain department's organizational structure | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Modify user information in part` `PATCH` /open-apis/contact/v3/users/:user_id >Update the user field in the Contacts. Parameters not passed will not be updated | Update Contacts Update user's basic information | `tenant_access_token` `user_access_token` | **X** | **✓** |
| `Update user information in whole` `PUT` /open-apis/contact/v3/users/:user_id >Update the user field in the Contacts. | Update Contacts | `tenant_access_token` | **X** | **✓** |
| `Delete a user` `DELETE` /open-apis/contact/v3/users/:user_id >Delete a user's information in the Contacts, used for termination | Update Contacts | `tenant_access_token` | **X** | **✓** | #### Event list

| `Employee information modified` | When employee information is modified | contact.user.updated_v3 | **✓** | **✓** |
| --- | --- | --- | --- | --- |
| `Employment` | When an employee is employed | contact.user.created_v3 | **✓** | **✓** |
| `Termination` | When an employee is terminated | contact.user.deleted_v3 | **✓** | **✓** | ### Resource: Department

View Resource fields and examples.

#### Methods

| `Create a department` `POST` /open-apis/contact/v3/departments >Create a department in the Contacts | Update Contacts | `tenant_access_token` | **X** | **✓** |
| --- | --- | --- | --- | --- |
| `Obtain single department information` `GET` /open-apis/contact/v3/departments/:department_id >Obtain information about a single department in the Contacts | Read contacts as an app | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Obtain the list of sub-departments` `GET` open-apis/contact/v3/departments/:department_id/children >Obtain the list of sub-departments under the current department | Read contacts as an app Obtain department's organizational structure | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Obtain parent department information` `GET` /open-apis/contact/v3/departments/parent >Recursively obtain information about the parent department of a department | Read contacts as an app Obtain department's organizational structure | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Search for departments` `POST` /open-apis/contact/v3/departments/search >Search for visible department data by keyword | Read contacts as an app | `user_access_token` | **✓** | **✓** |
| `Modify department information in part` `PATCH` /open-apis/contact/v3/departments/:department_id >Modify any field of the department information in the Contacts | Update Contacts | `tenant_access_token` | **X** | **✓** |
| `Update department information in whole` `PUT` /open-apis/contact/v3/departments/:department_id >Used to update all information of the current department | Update Contacts | `tenant_access_token` | **X** | **✓** |
| `Delete a department` `DELETE` /open-apis/contact/v3/departments/:department_id >Delete a department in the Contacts | Update Contacts | `tenant_access_token` | **X** | **✓** | #### Event list

| `Department information modified` | When department information is modified | Read contacts as an app | contact.department.updated_v3 | **✓** | **✓** |
| --- | --- | --- | --- | --- | --- |
| `Department created` | When a department is created | Read contacts as an app | contact.department.created_v3 | **✓** | **✓** |
| `Department deleted` | When a department is deleted | Read contacts as an app | contact.department.deleted_v3 | **✓** | **✓** | ### Resource: User Group

View Resource fields and examples.

#### Methods

| `Create a user group` `POST` /open-apis/contact/v3/group >Create a user group | Update user groups information | `tenant_access_token` | **X** | **✓** |
| --- | --- | --- | --- | --- |
| `Update a user group` `PATCH` /open-apis/contact/v3/group/:group_id >Update information about a user group | Update user groups information | `tenant_access_token` | **X** | **✓** |
| `Delete a user group` `DELETE` /open-apis/contact/v3/group/:group_id >Delete a user group | Update user groups information | `tenant_access_token` | **X** | **✓** |
| `Query a user group` `GET` /open-apis/contact/v3/group/:group_id >Obtain information about a user group | Obtain user groups information | `tenant_access_token` | **✓** | **✓** |
| `Query the list of user groups` `GET` /open-apis/contact/v3/group/simplelist >Query the list of user groups in an organization | Obtain user groups information | `tenant_access_token` | **✓** | **✓** |
| `Add members to a user group` `POST` /open-apis/contact/v3/group/:group_id/member/add >Add members to a user group | Update user groups information | `tenant_access_token` | **X** | **✓** |
| `Remove members from a user group` `PUT` /open-apis/contact/v3/group/:group_id/member/remove >Remove members from a user group | Update user groups information | `tenant_access_token` | **X** | **✓** |
| `Obtain the list of members in a user group` `GET` /open-apis/contact/v3/group/:group_id/member/simplelist >Obtain the list of members in a user group | Obtain user groups information | `tenant_access_token` | **✓** | **✓** | ### Resource: Workforce Type

View Resource fields and examples.

#### Methods

| `Query the workforce type` `GET` /open-apis/contact/v3/employee_type_enums >Obtain the workforce type of an employee | Read contacts as an app | `tenant_access_token` | **✓** | **✓** |
| --- | --- | --- | --- | --- |
| `Update the workforce type` `PUT` /open-apis/contact/v3/employee_type_enums/:enum_id >Update the custom workforce type | Update Contacts | `tenant_access_token` | **X** | **✓** |
| `Delete a workforce type` `DELETE` /open-apis/contact/v3/employee_type_enums/:enum_id >This API is used to delete custom workforce types. | Update Contacts | `tenant_access_token` | **X** | **✓** |
| `Add a workforce type` `POST` /open-apis/contact/v3/employee_type_enums >Add a custom workforce type | Update Contacts | `tenant_access_token` | **X** | **✓** | #### Event list

| `Workforce type created` | When a workforce type is created | Read contacts as an app | contact.employee_type_enum.created_v3 | **✓** | **✓** |
| --- | --- | --- | --- | --- | --- |
| `Workforce type enabled` | When a workforce type is enabled | Read contacts as an app | contact.employee_type_enum.actived_v3 | **✓** | **✓** |
| `Workforce type disabled` | When a workforce type is disabled | Read contacts as an app | contact.employee_type_enum.deactivated_v3 | **✓** | **✓** |
| `Workforce type modified` | When a workforce type is modified | Read contacts as an app | contact.employee_type_enum.updated_v3 | **✓** | **✓** |
| `Workforce type deleted` | When a workforce type is deleted | Read contacts as an app | contact.employee_type_enum.deleted_v3 | **✓** | **✓** | ### Resource: Custom user fields

View Resource fields and examples.

#### Methods

| `Obtain custom user fields of a company` `GET` /open-apis/contact/v3/custom_attrs >Obtain custom user field configuration of a company | Read contacts as an app | `tenant_access_token` | **✓** | **✓** |
| --- | --- | --- | --- | --- | ### Resource: Contact Scope

View Resource fields and examples.

#### Event list

| `Contact scope changed` | When the contact scope is changed | Read contacts as an app | contact.scope.updated_v3 | **✓** | **✓** |
| --- | --- | --- | --- | --- | --- |
