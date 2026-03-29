---
title: "Exchange resource introduction"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/exchange_binding/introduction"
service: "calendar-v4"
resource: "exchange_binding"
section: "Calendar"
updated: "1736837326000"
---

# Exchange resource introduction

The calendar API provides the ability to bind, unbind, and query the status of Exchange accounts and Lark accounts. You can synchronize your Exchange calendar to Lark calendar for use. When an Exchange account is bound to a Lark account, the Lark account will asynchronously synchronize the Exchange calendar to the Lark calendar. During the synchronization, you can check the synchronization status of the calendar in real time.

## Field description

The properties included in Exchange binding and unbinding are described below.

| Name | Type | Description |
| --- | --- | --- |
| `exchange_binding_id` | `string` | The ID generated when creating an Exchange binding relationship is the unique identification ID of the admin account, exchange account, and user triplet. You can use this ID to query the binding relationship, calendar synchronization status, or unbind the relationship. **Getting method**: Get it from the returned result when calling the Bind Exchange account to Lark account interface. **Example value**: "ZW1haWxfYWRtaW5fZXhhbXBsZUBvdXRsb29rLmNvbSBlbWFpbF9hY2NvdW50X2V4YW1wbGVAb3V0bG9vay5jb20=" |
| `admin_account` | `string` | Exchange admin account. **Example value**: "email_admin_example@outlook.com" **Field permission requirements**: `contact:user.email:readonly` |
| `exchange_account` | `string` | Exchange account to be bound. **Example value**: "email_account_example@outlook.com" **Field permission requirements**: `contact:user.email:readonly` |
| `user_id` | `string` | User ID, which is the Lark account ID bound to the Exchange account. For information about user IDs, see User-related ID concepts. **Example value**: ou_xxxxxxxxxxxxxxxxxx |
| `status` | `string` | The synchronization status of the Exchange account. **Optional values**: - `doing`: Calendar is syncing - `cal_done`: calendar synchronization completed - `timespan_done`: recent time period synchronization completed - `done`: schedule synchronization completed - `err`: synchronization error **Example value**: "doing" | ## Data example

```json
{
     "admin_account": "email_admin_example@outlook.com",
     "exchange_account": "email_account_example@outlook.com",
     "user_id": "ou_xxxxxxxxxxxxxxxxxxxxx",
     "status": "doing",
     "exchange_binding_id": "ZW1haWxfYWRtaW5fZXhhbXBsZUBvdXRsb29rLmNvbSBlbWFpbF9hY2NvdW50X2V4YW1wbGVAb3V0bG9vay5jb20="
}
```
