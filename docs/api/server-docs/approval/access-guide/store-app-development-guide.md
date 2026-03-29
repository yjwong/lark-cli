---
title: "Store app development guide"
url: "https://open.larksuite.com/document/ukTMukTMukTM/ukzNyYjL5cjM24SO3IjN"
section: "Approval"
updated: "1675167385000"
---

# Store app development guide

## Target audience:

Vendors of third-party company apps in Lark

## Benefits:

When a third-party company app developed by an app vendor involves the approval workflow (taking attendance app as an example, when an employee forgets to clock in/out, they need to initiate a correction approval), the approval API can be called directly to implement the approval workflow without the need to build approval process-related functions in the app, thus reducing the development cost and enhancing approval experience.

## How to integrate:
1. The app vendor has created an app in the Lark Open Platform. For how to create an app, see the app creation documentation.

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/ffd9268bc3c9f21bc75c60b17e4232bc_oHAUGSI8gu.png?lazyload=true&width=1640&height=986)

2. Specify that the approval process is created through the API at a specific time (taking attendance app as an example, specify that the attendance app creates a correction approval process via the approval API when Attendance is enabled by the company). For how to create an approval process, refer to Create an approval definition.

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/a3dab00f1e02c089c614e30b3a2c4784_WWk5Hs7Liu.png?lazyload=true&width=1280&height=362)

3. When an employee initiates an approval in a third-party company app or on the approval initiation page, the third-party company app creates an approval instance via the API. For how to create an approval instance via API, refer to Create an approval instance.

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/55d35dabd74b3dd577746871981cbc29_r61BcqJ89N.png?lazyload=true&width=1538&height=1336)

4. The approver completes the approval in the Approval app, which then synchronizes the approval result to the third-party app (for example, the approver completes the correction approval in the Approval app, and the approval result is returned to the Attendance app, which changes the employee's no record status to normal). For how to sync the approval status, refer to Subscribe to approval events.

## The following APIs are used to integrate third-party company apps into Approval:

1. Create an approval definition
2. Create an approval instance
3. Subscribe to approval events
