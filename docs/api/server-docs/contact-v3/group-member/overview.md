---
title: "Overview"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/group-member/overview"
service: "contact-v3"
resource: "group-member"
section: "Contacts"
updated: "1646791351000"
---

#  Overview
##  Resource definition
User group members can be users or departments in the organizational structure (only users are supported currently). You can define which users/departments are included in a user group by using "user group member".

## Description of user group member fields

| Parameter | Type | Description |
| --- | --- | --- |
| `member_type` | `string` | User group member type **Optional values of member_type**: - user: User - department: Department (available soon) **Example value**: "user" |
| `member_id` | `string` | Member ID **member_id description**: - When member_type is user, the user ID is supported, which can be open_id, union_id, or user_id (View ID types). **Example value**: "ou_7dab8a3d3cdcc9da365777c7ad535d62" | ## Data example

```json 
{
   "member_id": "u287xj12",
   "member_type": "user"
}
 
```
