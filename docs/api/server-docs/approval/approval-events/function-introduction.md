---
title: "Function introduction"
url: "https://open.larksuite.com/document/ukTMukTMukTM/ugDNyUjL4QjM14CO0ITN"
section: "Approval"
updated: "1676430224000"
---

# Function introduction

## Why monitor approval events?

1. After the developer creates an approval instance through the approval open interface, there is a lot of information that may need to be paid attention to in the life cycle of an approval order. The developer can monitor the status change information of the approval order through callbacks.
2. It avoids the constant polling of the approval details interface to obtain the latest status, which not only increases the development complexity, but also causes unnecessary interface query consumption.
3. Approval events include shared approval events and special approval events. Shared approval events include：**Approval instance status changes** / **Approval task status changes**、**Approval cc status changes**. Special events are generally used by applications that have had special interaction with approval. Under normal circumstances, we developers do not need to pay attention.

## The push path of the event
1. The developer registers the request callback address in "Developer Background" - "Development Configuration" - "Event Subscription".
2. Approval When the status of the approval document (a document is an approval instance) changes, such as the creation or flow of the document, the creation of a new task, and the copy, a change message will be sent to the open API.
3. The open API will store the message first, and then push the message in the format of the response to the callback address registered by the developer.
4. When the user server interface receives the message, it can store the message first, and then perform business logic processing to avoid response timeout.
5. The message may be repeated, developers need to do idempotent according to the event_id or uuid in the callback request, different event types idempotent keys are different, please determine according to the structure of the specific event, or you can do idempotent according to your own business.

## Push cycle and frequency
1. Details：Receive and process events
2. After the application receives the HTTP POST request, it needs to respond to the request with HTTP 200 status code within 1 second. Otherwise, Lark Open Platform considers this push as a failure and will re-push the event at intervals of 15s, 5m, 1h, 6h, with a maximum of 4 retries.
3. From the above description, it can be seen that the maximum time window for event retransmission is about 7.5 hours . Please check and handle the repeated events within 7.5 hours. Event uniqueness can be determined using the following: For version 1.0 events, event uniqueness is determined by the uuid field in the event structure.For version 2.0 events, the event uniqueness is determined by the event_id field in the event structure.
