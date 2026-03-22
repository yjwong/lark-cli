---
title: "Access guide"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/attendance-development-guidelines"
service: "attendance-v1"
resource: "attendance-development-guidelines"
section: "Attendance"
updated: "1693473324000"
---

# Access guide

## API capability
The attendance API provides a wide range of open capabilities. Developers can use this API to obtain information such as users' attendance results, attendance records, attendance group details, and shift details within their authorization scopes, and can also create, modify, and delete attendance groups and shifts.

## API integration development process
**1. Go to [Lark Open Platform > Developer Console](https://open.larksuite.com/app/)**

**2. On the Open Platform page, click "Create app"**.
- 2.1 Select "Custom app".
- 2.2 Fill in the app name. We recommend "Attendance API".
- 2.3 Fill in the app subheading. We recommend "Create developer account".
- 2.4 Click "Confirm" to create the app.

**3. On the Open Platform page, click the app you just created**.
- 3.1 View the App ID and App Secret.
- 3.2 Upload a new app icon.
- 3.2 Click "Security settings" and set an IP allowlist for calling the attendance API.

**4. On the Open Platform page, grant read and write permission scopes to the attendance app.**
- 4.1 Click "Permissions & Scopes" and grant read and write permission scopes to the attendance app.

**5. On the Open Platform page, publish the custom app.**
- 5.1 Click "Features" > "Bot" to enable the bot.
- 5.2 Click "Version Management & Release" and release a version.
- 5.3 If the company has set a post review, the developer account of the app will be activated only after the approval of the **company administrator**.

**6. Go back to the [Attendance Admin](https://www.larksuite.com/attendance/manage/statistics/day) and click "API Integration" in the upper-right corner.** If no such access is found, access preparations are complete and the following content can be ignored.

> If you haven't completed API integration, you will receive an error (code: 1220004) if you call the attendance API.

- 6.1 Example of successful API integration.

![2.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/8ca1a2a4026168f946ccabe31a6ee2ab_DRbUk5pliP.png?height=603&lazyload=true&width=839)

- 6.2 Example of incomplete API integration.

![iShot_2022-06-23_16.41.02.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/8349ffe4427590343f9e339f7a596d04_NaaYlUuSxR.png?height=1410&lazyload=true&width=1754)

## Event subscription development process
**1. Go to [Lark Open Platform > Developer Console](https://open.larksuite.com/app/)**

**2. On the Open Platform page, click "Create app" (If you have already created the app, you can skip this step).
- 2.1 Select "Custom app".
- 2.2 Fill in the app name. We recommend "Attendance API".
- 2.3 Fill in the app subheading. We recommend "Create developer account".
- 2.4 Click "Confirm" to create the app.

**3. On the Open Platform page, grant read and write permission scopes to the attendance app** (If you have already done so, you can skip this step).
- 3.1 Click "Permissions & Scopes" and grant read and write permission scopes to the attendance app.

**4. Subscribe to events on the Open Platform page.** 
- 4.1 Click "Event Subscriptions" and fill in the callback address in the "Request URL" field.
- 4.2 To add events, select "Attendance" and select the sub-events to subscribe to

**5. On the Open Platform page, publish the custom app.**
- 5.1 Click "Features" > "Bot" to enable a bot (If you have already done so, you can skip this step).
- 5.2 Click "Version Management & Release" and release a version.
- 5.3 If the company has set a post review, the developer account of the app will be activated only after the approval of the company administrator.

**6. Go back to the [Attendance Admin](https://www.larksuite.com/attendance/manage/statistics/day) and click "API Integration" in the upper-right corner. If no such access is found, access preparations are complete and the following content can be ignored**

## Terms
**employee_id**

Employee ID, same as a user_id. For details, see Terms.
