---
title: "Overview"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/directory-v1/employee/overview"
service: "directory-v1"
resource: "employee"
section: "Organization"
updated: "1756439237000"
---

# Overview of employee management

The following APIs are provided for the Employee Management resource:

## Methods

**Create employee**

    This API is used to create employees under the enterprise, which can be understood as employee onboarding. An employee refers to a member of Lark who is an "Employee", which is equivalent to a "User" in the Contacts OpenAPI.

**Patch employee**

    This API is used to update the information of an incumbent or retired employee and freeze or restore an employee. Parameters that are not passed are not updated. An employee is a member of Lark who is an "Employee", which is equivalent to a "User" in the Contacts OpenAPI

**Delete employee**

    This API is used to delete an employee

**Resurrect employee**

    This interface is used to restore a member who has left the company

**Update unresigned employees to pre-resigned**

    This interface is used to handle the departure of current employees and update them to the status of "pending departure". "pending departure" employees will not automatically leave, and need to use the "delete employee" API to terminate and transfer resources.

**Update pre-resigned employee to un-resigned**

    This interface is used to cancel the departure of the pending employee and update it to the "on-the-job" status. When canceling the departure, the departure information will be cleared.

**Get employee information in bulk**

    This API is used to query the details of employees based on their IDs in batches. An employee is a member of Lark who is an "Employee", which is equivalent to a "User" in the Contacts OpenAPI

**Query the list of employees in batches**

    This API is used to obtain the detailed list of eligible employees in batches based on specified conditions. An employee is a member of Lark who is an "Employee", which is equivalent to a "User" in the Contacts OpenAPI

**Search for employees**

    This API is used to search for employee information. An employee is a member of Lark who is an "Employee", which is equivalent to a "User" in the Contacts OpenAPI
