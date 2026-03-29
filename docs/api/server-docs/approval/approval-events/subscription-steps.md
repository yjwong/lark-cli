---
title: "Subscription steps"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uIDO24iM4YjLygjN/event/subscription-steps"
section: "Approval"
updated: "1676430227000"
---

# Subscription steps

### 1. Configure event subscription callback URLs in the background

- Register the event callback address in "Developer Background" - "Event Subscription". For specific steps, please refer to the open API documentation：Event Subscription Overview。
### 2. Enable application permissions
- Search for "approval" in "Developer Background" - "Rights Management", and open the searched permissions.

### 3. Subscribe to approval events
- In "Developer Background" - "Event Subscription", add events, for example, you can subscribe to "Approval Instance Status Change" or "Approval Task Status Change" events.

### 4. Subscribe to the approval definition that needs to be listened to

After the application background subscribes to the approval event, you still need to subscribe to the specified approval definition again through the approval open interface to receive the approval event.

* Subscribe to approval events
* Unsubscribe to Approval Events

Approval interface authority authentication dependent tenant_access_token, the acquisition method is as follows:

* Obtain app_access_token (for custom apps)
* Obtain tenant_access_token (for store apps)
