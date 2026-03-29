---
title: "Callback overview"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/event-subscription-guide/callback-subscription/callback-overview"
section: "Events and callbacks"
updated: "1744337095000"
---

# Callback overview

Callbacks are suitable for business scenarios that require synchronous response to user behavior. That is, when a user triggers certain operations on Lark, the front-end loads and waits for the server to return response data. When the server returns the response result, the front-end loading is completed and the returned response result is displayed to the user.

In Lark business, typical usage scenarios of the callback function are as follows:

- **Card interaction scenario**: The user clicks on the interactive component on the card (such as the agree/reject button on the approval card). The developer's server will receive the click callback of the button and needs to respond immediately to the updated card content. , giving users operational feedback (such as changing the approval status to approved).

- **Link preview scenario**: The user views an application link in the chat, and the link supports returning the application configured preview data. At this time, the application's server will receive The callback to Pull link preview data needs to respond immediately and return the link preview content, so that the end user can see the link preview effect.

## The difference between callbacks and events

The callback is similar to event but different:

- **Similarity:**
     - Lark server actively pushes data to the developer server.
     - Callbacks and events have similar data structures and can reuse the same encryption and decryption strategy. Developers can use the same strategy when parsing the content returned by Lark.
- **Differences:**
     - After subscribing to the callback, the developer server needs to **immediately return** the response content to feedback the user operation, while the event does not require return.
     - The callback is a synchronous operation and does not provide a supplementary push mechanism. If the callback fails to respond after a timeout, the callback will be considered a failure. The front end will display error reports and other response strategies provided by the platform.
     - Events are asynchronous operations. Developers only need to simply respond to whether the Lark server receives the event. If the developer does not respond, the platform will push the event.

## Subscription process

| Step | Description |
| --- | --- |
| 1. Set the Callback Subscription Request URL | You need to provide the public network address of the server that will receive callback messages. When an application subscribes to a callback event, the open platform will send HTTP POST requests to the given public network address of the server, and the request will contain the callback data. For detailed configuration instructions, refer to Sending Callbacks to Developer's Server. |
| 2. Add Required Callbacks | Once the callback subscription setup is complete, you can add the necessary callback subscriptions to your application and publish the application to apply the changes. For detailed instructions, refer to Adding Callbacks. |
| 3. Receive Callbacks | You need to perform security verification based on the application's encryption strategy. If it is an encrypted callback, you need to decrypt the callback first and then parse the callback details. For detailed instructions, refer to Receiving and Handling Callbacks. | ## Callback structure

The data structure returned by the callback is similar to the event. Taking "Pull link preview data" (`url.preview.get`) as an example, the structure of the callback is as follows:

```json
{
   "schema": "2.0", //Indicates the callback version. 2.0 indicates that the callback structure is consistent in form with the 2.0 version of the event.
   "header": { //General parameters for callbacks
     "token": "vi57noNQoGbhxxxxxWmmWdlsSn3FTzk1", //Corresponding to Verification Token
     "create_time": "170134xxxxx18480", //The timestamp sent by the callback is approximate to the time when the callback is triggered
     "event_type": "url.preview.get", //Callback type
     "tenant_key": "736xxxxx260f175d", //The tenant ID of the application to which the callback belongs
     "app_id": "cli_a40xxxxxe57e100c" //The application ID of the application to which the callback belongs
   },
   "event": { //Record detailed context information returned by different callback types
     "operator": {
       "tenant_key": "736588cxxxx175d",
       "user_id": "c3xxxxd1",
       "open_id": "ou_xxxxx54182ea7b8319f4d39823b79d2"
     },
     "host": "im_message", //The host scene where the link is located. The enumeration includes 1.im_message chat message 2.im_top_notice group top
     "context": { //Specific context parameters in this scenario
       "url": "https://www.example.com/smartcard/test/111", //Original link matching URL rules
       "preview_token": "e28r7df2-xxxx-477d-a8d0-2e1eb99796c2", //Certificate used to identify link preview, used when returning link preview data
       "open_message_id": "om_191d914xxxxx81c97a609c663452dfdf", //Message ID that triggers link preview
       "open_chat_id": "oc_20443194b65f9c8cf2935818dae39999" //Group ID that triggers link preview
     }
   }
}
```

## Callback list

The currently supported callback list is as follows:

| Function module | Callback name | Description |
| --- | --- | --- |
| Card | Card postback interaction | This callback is triggered when the user clicks on the component configured with postback interaction on the card. User interaction can be fed back by returning toast, updated card content, etc. |
| Link preview | Pull link preview data | This callback is triggered when the user views a link in the chat that matches the URL rule registered by the app. You can preview content by returning text links, cards and other links to extend the link preview effect for naked links. |
| Card | Card postback interaction (old) | When a user clicks on a component with postback interaction added to the card, the server callback address registered by the developer will receive this callback. Developers can declare responses to user interactions by popping up toasts, updating cards, keeping the original content unchanged, etc. > This callback uses the old version of the protocol and is compatible with historical robots callback configuration. |
