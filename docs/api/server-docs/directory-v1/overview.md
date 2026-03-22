---
title: "Organizational structure overview"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/directory-v1/overview"
service: "directory-v1"
resource: "overview"
section: "Organization"
updated: "1756439236000"
---

# Organizational structure overview

The OpenAPI of the management background - organizational structure serves enterprise administrators to manage the organizational structure data in Lark. Example:
- Create, update, query, delete department
- Create, update, query, leave, restore employees

You can view the specific open capabilities of the organization structure in API list

## Applicable Scenarios

There are various application scenarios for the organizational structure function, and a few typical application scenarios are listed below.

**Scenario 1: HR System Integration with Lark**

After completing the transfer to and from the HR system in the enterprise, it is necessary to synchronize the information and status changes of the Lark account. If you rely on HR to manually maintain, it is not only time-consuming and laborious, but also prone to errors and omissions. At this point, data synchronization can be achieved through the organization's API.

First of all, you need to create a Lark application, and secondly, determine the organization structure API to be called, and then turn on the application's permissions and get the application access credentials, so that you can operate on the departments and employees when the in-transfer transfer is completed, as well as view the completed synchronization of the organization structure.

**Scenario 2: Developing Enterprise Applications**

When an enterprise develops a self-built application, if it needs to obtain the organizational structure or employee data, it can use the organizational structure API. note that when using the organizational structure API to query data, in addition to obtaining the interface permissions, it is also necessary to obtain the permissions corresponding to specific fields.

**Scenario 3: Administrator Operations**

In addition to maintaining the organization architecture data in the Lark management background, administrators can also operate by calling the API. Note that when using the administrator's access credentials to call the API, it can only be called successfully if the management backend has the same operating privileges, and the scope of members and departments that can be operated will follow the administrator's management scope.

## Resource introduction

The organizational structure is resource-centric in its development approach. The currently available resources and relationships are as follows:

![EmployeeDepartmentRelation](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/de9c4afbdae82418db273c68cd274657_1dC3VfO6Tl.jpeg?height=1196&lazyload=true&width=1662)

## Permission requirements

The contact directory permission scope defines which department and user data the application can access by calling the organizational structure interface. Data that is not within the application's contact directory access scope cannot be accessed.

As shown in the diagram, the organizational structure has three departments: A, B, and C. If an application only has contact permissions for departments B and C, then that application can only retrieve data for departments B and C through the interface. Attempting to access data for department A would result in a "no permission" error. For a detailed introduction to the contact permission scope, please refer to Introduction to Contact Permission Scope。
![](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/8b71ac2df29d3ecf77856d90b19beaaa_JRarzEWeF6.png?height=1320&lazyload=true&maxWidth=600&width=2400)

## Integration Procedure

| procedure              | introduction                                                                                                                                                                                                                                                                                  |
| --------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Create an application          | For details about how to create a self-built application, see The development process of self-built applications  For details about how to create a store application, see Developing and listing store application                           |
| Invoke apis to manipulate organizational structures | Before calling the API, you need to obtain access credentials and enable the corresponding permissions, see How to use the server API You can easily try these apis in [API Explorer](https://open.larksuite.com/api-explorer?from=guide)seeAPI Explorer Guide |
| Listen to events and monitor organizational changes  | The organization structure does not currently provide events. If you want to monitor changes to the organization structure, you can use Contacat events Before listening to an event, you need to apply for the appropriate permissions, see details Event subscription overview                                                                              | ## Rate limits
Different levels of frequency control strategies have been set for the OpenAPIs of different organizational structures. You can find the interface frequency limit information in each interface document, and the Rate limits also explains how to deal with the situation when frequency control restrictions are encountered during the invocation of OpenAPIs.

**Rate Limits Overview**

The interfaces of the organizational structure adopt a dual-restriction approach at the application level and the enterprise level:
- When the invocation of an interface by a single application exceeds the preset rate limit threshold of the interface, the application-level restriction will be triggered.
- When multiple applications under the same enterprise invoke the same interface, the enterprise-level restriction is triggered cumulatively.
- When all enterprises across the entire platform simultaneously invoke the same interface, the platform-level restriction is triggered cumulatively.

**Interface Classification Description**

The general logic for interface rate limits is as follows. You can also find rate limit details in the documentation for each interface:

| API Type | Rate limits | Example |
| --- | --- | --- |
| Data Creation | 5 times/sec per enterprise per application | Create Employee, Create department |
| Data Update | 10 times/sec per enterprise per application | Update Employee, Delete employee, Reinstate departed employees, Update department, Delete department |
| Data Retrieval | 1000 times/min & 50 times/sec per enterprise per application | Obtain employee information in batches, Get employee list in batches,Search employee，Batch get department info, List department, Search department |
