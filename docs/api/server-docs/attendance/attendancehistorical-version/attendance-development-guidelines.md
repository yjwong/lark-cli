---
title: "Attendance Development Guidelines"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/Attendance/Development_Guide"
section: "Attendance"
updated: "1646322856000"
---

# Clock-in/out developer's guide

## Attendance API capabilities
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
- 5.3 If the company has set a post review, the developer account of the app will be activated only after the approval of the company administrator.

**6. Go back to the Attendance Admin and click "API Integration" in the upper-right corner.**

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

**6. Go back to the Attendance Admin and click "API Integration" in the upper-right corner.**

## Terms
**employee_id**

Employee ID, same as a user_id. For details, see Terms.
