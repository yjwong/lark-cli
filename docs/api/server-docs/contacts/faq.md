---
title: "FAQ"
url: "https://open.larksuite.com/document/ugTN1YjL4UTN24CO1UjN/uQzN1YjL0cTN24CN3UjN"
section: "Contacts"
updated: "1646791340000"
---

# FAQs

This document describes frequently asked questions when developers use APIs related to "Contacts".
> **Note:** You can quickly locate questions you are looking for via the categories listed on the right navigation bar.

## How do I efficiently use APIs related to Contacts?

A: The Contacts module contains several submodules such as Member, Department, and User Group, each of which provides related APIs for developers to intelligently use existing capabilities. For developers, permission scopes to use such APIs include the following **three types**:
1. Permission to use an API Indicates whether an app has the permission to call an API.
2. Permission scope to access data, for example, query or operate on data.
3. Permission to access sensitive data fields Due to the sensitivity of fields for some entities, for example, a user's mobile number, extra permission scope is required to obtain such fields. Before using APIs, developers can check whether they have related API permissions. Check the following:
1. Whether you have the permission scope to call an API. If no, the API cannot be accessed.
2. Whether the data you want to obtain or operate on is within the contacts permissions. If no, a message indicating no scope will be displayed.
3. Whether you have the permission scope to obtain the specified field of an entity. If no, the data field may fail to be obtained.

#### ***Example***
> Suppose that you expect to add users to a company via the [Create a user] API. Firstly, check whether you have the permission scope to call the "Update Contacts" API. Then, check whether you have the contacts permissions of the departments to add users, so that you can add users to these departments. Finally, check whether you have the permission scope to obtain sensitive user information.

## Range of contacts data that an app can access

**Q: What's the difference between using user_access_token and tenant_access_token to access Contacts APIs?**

A: They have different permission scopes.
1. **Tenant_access_token**: Whether data can be accessed depends on the app's contact scope. For example, if data of department A is to be obtained, the system will check whether department A is within the range of contacts data that an app can access. To configure the scope, go to Admin > Workplace > App Management.
2. **user_accss_token**: Whether data can be accessed depends on the user's visible scope of organization (i.e., the structure visible on Contacts > Organization in the app). To configure the visible scope, go to [Admin](https://www.larksuite.com/admin) > Security > User Permissions > Visible Scope of Organization.

**Q: How can I obtain the information about all employees in my company?**

A: Currently, no individual API is available to obtain information about all employees. If you need to do so, perform the following steps:
1. Check whether the contact scope that an app can access is all employees. If not, information about all employees cannot be obtained.
2. If the contact scope is all employees, call the Obtain department information list API (parent_department_id=0; fetch_child=true) to obtain all department IDs of the company.
3. With the department IDs obtained in step 2, call the Obtain user list API to obtain all employee information.

**Q: How can I obtain the Contacts of employees under the root department?**

A: Firstly, check whether the range of contacts data that an app can access contains all employees under the root department or in the entire company. Then, call the Obtain user list API (department_id=0) to obtain the information about employees under the root department.

## ID
**Q: Why do I fail to obtain information such as user_id, email, and mobile number even if permission scopes are set?**

A: Information such as user_id, email, and mobile number is sensitive, so extra permission scope is required to obtain such data. For details, see App scopes.

**Q: Does a user has the same open_id in different apps?**

A: No. open_id is a unique identifier for a user in an app. Therefore, a user has different open_id in different apps.

**Q: Does a user has the same open_id and union_id in different companies?**

A: No. A user has different open_id and union_id in different companies.

**Q: Can user_id (employee_id) be updated?**

A: Custom user_id (also called employee_id in some scenarios) cannot be updated. However, automatically generated user_id can be updated for **once**. To update this information, call the Update user information in whole API with the following parameters input:
```
{    
    "employee_id":"1fddc168",    
    "new_employee_id":"test2"
}
```
In principle, you are not advised to modify user_id, since it may have been used and saved by different apps. If user_id is changed while apps are unaware, these apps may work improperly.

**Q: How can I obtain user_id?**
A: You can obtain user_id via email or mobile number by calling the Obtain user ID via email or mobile number API. You can also log in to Lark [Admin](https://www.larksuite.com) and click "Member Details" to obtain the user ID.

## User/Member

**Q: In case where a department (or member) is deleted by mistake, can I recover it with its original ID?**

A: Currently, an original ID cannot be recovered. ID is unique in a company. You are advised to create a new ID.

**Q: Are there any detailed explanations of the status field?**

A: You can convert the value of the status field into a binary number, and each bit indicates a state. Bit 0 (the least significant bit): 1 indicates frozen and 0 indicates unfrozen; bit 1: 1 indicates terminated and 0 indicates in-service; bit 2: 1 indicates unactivated and 0 indicates activated.
For example, convert the value "2" into the binary number "010", which means "activated, terminated, unfrozen"; convert the value "0" into the binary number "000", which means "activated, in-service, unfrozen".

**Q: I have developed multiple apps, and I wonder whether a user is using different apps developed by me at the same time. Can I use any field to know that?**

A: Currently, union_id is unique for a user in different apps developed by the same developer.

**Q: Why does the Search users API fail to obtain any return value?**

A: You can do the following:
-  Check whether you can search for the specified user on Lark, and this user is not an external user or has been terminated.
-  Check whether the value of page_token is correctly configured. 

**Q: When I want to obtain user ID/information by calling the "Obtain user ID via email or mobile number" API or "Obtain user information" API, why can I obtain only mine?**

A: Ensure that the Range of contacts data that an app can access contains users whose information you want to obtain, not only yourself. Otherwise, a result of "email of mobile number does not exist" will be returned.

**Q: Why is a member added to other department groups in some cases?**

A: You can do the following:
- Check whether the member's department was modified.
- Check whether this member is the head of the department to which he/she is added.

**Q: If a userid is automatically generated for an employee, after this employee is deleted, will this userid be occupied by a newly added employee?**

A: A userid automatically generated ensures the uniqueness of a valid user (not terminated) within a company. When developers or administrators delete an existing user and then add a new user, userid of the new user may be the same as that of the deleted one. It is advised to use a unique ID as userid for a new user to avoid any confusion. (Note: This rule also applies to the custom ID of a department.)

## User APIs

**Q: Please explain the custom field in the request or response structure of the Create a user or [Obtain single user information]  API.**

A: custom_attrs is an extended user attribute, used for a company to flexibly extend its user description capabilities. Based on the value syntax, this attribute is categorized into text, web page, enumeration, image, and user types.
**Text type**: name attribute of a user, with a string value. Its value corresponds to the text field.
**Web page type**: Field reference link, which realizes redirection from a member's profile page. This type of field requires title text (text) and redirection URL (url). Due to the difference between PCs and mobile clients in redirection, pc_url is separately configured as the redirection URL on PCs. If it is left blank, PCs also use url as the redirection URL.
**Enumeration type:** This enables value selection from specific options, for example, selecting an employee type from employee, outsourcing, and consultant. The value corresponds to option_id, an option key configured by the administrators.
**Image type:** Similar to enumeration type, but the options are images displayed on a member's profile page. The value corresponds to option_id, and you can query the image ID from the admin console.
**User type:** Used to display reference to another user on a member's profile page, for example, displaying the HRBP field of "John" as "Richard" and supporting redirection to Richard's profile page by clicking. generic_user id indicates user_id to be referenced, and generic_user type is fixed to 1, indicating the user type.

**Notes:**
Ensure that your organization administrator has enabled "Allow to invoke by Open Platform's Contacts API" in [Admin > Organization > Member Profile > Custom Fields > Global Settings](http://www.larksuite.com/admin/contacts/employee-field-new/custom). Otherwise, this custom field will not take effect or its value will not be returned. If any custom field value is created or updated, ensure that the custom field key has been created by the company administrator and remains valid.

**Q: How do I use the enterprise_email field for the Create a user API?**

A: enterprise_email indicates the business email of a user. The domain name for a business email requires application on the admin console. If a company does not enable the business email service for the specified domain name, setting this field will fail. In this case, contact the company administrator to check whether the business email service has been enabled for the specified domain name.

**Q: Please explain the Obtain user list API.**

A: This API is used to obtain the list of direct members under a department. Before data acquisition, the range of contacts data that an app can access will be checked first. If the request contains a department ID (root department ID: 0), the system will check whether the app has the contacts permissions of this department. If yes, information about direct members under the department will be returned. If the request does not contain any department ID, the system will return all independent users within the permission scope (excluding departments of these users). When using this API, developers need to consider whether to contain a department ID in the request based on contacts permissions and the data to obtain.

## Department APIs

**Q: How can I obtain contact information of all employees under a parent department?**

A: Currently, no individual API is available to obtain contact information of all employees under a parent department. You need to firstly call the Obtain department information list API to obtain department_id of all departments under a parent department, and then call the Obtain user list  API to obtain the contact information of employees under each department.

**Q: Why does department_id obtained by the Obtain department information list API includes "od-" in some cases?**

A: Check whether the request specifies a different department_id_type. If no, the field returned is open_department_id by default.

**Q: Please explain the Obtain department information list API.**

A: This API is used to obtain the list of child departments directly under a department. Before data acquisition, the range of contacts data that the app can access will be checked first. If the request contains a parent department ID (root department ID: 0), the system will check whether the app has the contacts permissions of this department (if department ID is 0, the system will check whether the permission scope is all employees). If yes, information about child departments directly under the department will be returned (whether recursion is used depends on the value (true or false) of fetch_child). If the request does not contain any department ID, the system will return data based on the range of contacts data that an app can access. If the range is all employees, the system will return single root department ID 0 (by using which developers can continue requests). Otherwise, information about departments in the range will be returned. When using this API, developers need to consider whether to contain a parent department ID in the request based on contacts permissions and the data to obtain.

## Troubleshooting

**Q: Why is a success message returned even if partial information is returned by using an API to obtain information in batches?**

A: The return value for a batch operation API indicates the execution status of the batch operation, instead of whether the operation is successful. As long as the batch operation is started, the status will be marked as executed. For details, refer to Query the status of batch task execution.

**Q: What can I do if "40013 param error" is returned?**

A: You can do the following:
- Check whether the specific employee or department is terminated or deleted.
- Parameter error. Refer to the following descriptions. If you have any question, contact your administrator.
- Mobile number already exists. Check whether there is a duplicate mobile number. 
- Mobile number and email address conflict. Check whether the mobile number and email address belong to two different Lark accounts. If so, modify the mobile number or email address of either account, or delete either account.  
- Check whether a department name is repeatedly used.
- A child department cannot be updated as a parent department.
- Custom text field exceeds the length limit, 100 characters.

**Q: What can I do if the Create a user API returns "email and mobile account conflict"?**

A: When creating a user, setting his/her mobile number and email address at the same time will make them associated with the same account as the login credential. This error occurs when the mobile number and email address of a user have been used to register two different accounts, respectively. To fix this error, set the mobile number only when creating the user, and then update the email address. In this way, the email address only functions as an attribute instead of a login credential. Alternatively, instructing the user to delete either account registered before can also solve this problem.

**Q: Why does the Create a user API return "department id xxxxxxxx is not exist"?**

A: Generally, this error message is returned for any of the following reasons:
- Input department id does not exist.
- open_department_id and department_id are two different Ids. Check whether open_department_id is used. 
- The app has no scope to access contacts data of this department.
