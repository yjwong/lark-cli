---
title: "Function introduction"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/admin-v1/badge/quick-start"
service: "admin-v1"
resource: "badge"
section: "Admin"
updated: "1686303056000"
---

# Function introduction

The enterprise badge aims to expand the scenario of "corporate culture + employee incentives", and build an honor system such as badges and an employee care system. Combining employee interaction with corporate culture to open up corporate customization capabilities.

We provide a safe and reliable API to facilitate you to manage badge information and configure medal award lists. Using our API interface, you can connect the enterprise badge function of Lark's management background with the user's own customized solution.

For how to quickly use the medal function, you can refer to the documentation [How do administrators upload and manage enterprise medals?](https://larksuite.com/hc/en-US/articles/360048488418), get up to speed quickly.

# Scope of authority
## Applicable applications
custom,isv

## Dependent permissions
| API permissions for badge resources | `admin:badge` |
| --- | --- |
| API permissions for badge grant. | `admin:badge.grant` |
| Grant the list API the dependency permission to use the address book entity. | `contact:user.employee_id:readonly` | # Development tutorial
**Tutorial steps dismantling**                                                                                                                                                             |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| 1.  The tenant must be an Ultimate type tenant. |
| 2.  Enable the relevant permissions that the application depends on.
| 3.  Get app access credentials.|
| 4.  Upload detailed pictures of badges and ornaments.|
| 5.  Create and update badge information.|
| 6.  View badge information.|
| 7.  Create and update award lists for the badge.

# Resources Introduction
The badge business domain is opened with "resources" as the center. The relationship diagram of resources is as follows:

![3f496166-8938-4b31-96b9-6c113b5aa4fb.svg](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/887474f60b888eee4e3a5d316091dfef_vi4gPxrzO9.svg?lazyload=true&width=1153&height=392)

Resources are defined as follows:
**Resource** | **Resource definition**                                                                        |
| ------ | ------------------------------------------------------------------------------- |
| badge     | A badge in an enterprise of Lark. Badges can be awarded to members of the company. After the badges are awarded, they will be displayed on the personal business card page of the corresponding member and on the badge wall. Members can also wear the badge on their avatars.|
| grant   | Multiple award lists can be created under a medal, and award lists represent the award of the medal to a specific group of members within a specific time frame.
