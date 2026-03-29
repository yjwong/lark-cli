---
title: "Event Callback Optimization Guide"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/event-subscription-guide/event-callback-optimization-guide"
section: "Events and callbacks"
updated: "1699845022000"
---

# Event callback optimization guide

## Why am I receiving notifications?

  
To help developers better use Lark's open capabilities, the open platform provides developers with the ability to monitor the health of event callbacks. In cases where the success rate of event callbacks is less than 100%, developers will receive card messages through the developer assistant to remind them to optimize their services.

You received a notification because your event callbacks failed to process successfully after four retries in the past week.

## Optimization strategies

You can refer to the following strategies to optimize your event callback service.

  
1. **Use servers located in mainland China**

    Lark's event callback service requires you to respond with an HTTP 200 status code within 3 seconds to determine the success of the request. Using servers located in mainland China can effectively reduce the time consumed on the network, provide margin for internal service processing, and thus reduce the possibility of your server timing out in response.

2. **Handle time-consuming operations asynchronously**

    For certain time-consuming operations, you may not be able to respond to the request within 3 seconds, causing the event callback to consider the request as a failure and make another request. You can convert certain time-consuming operations into asynchronous processing, respond with HTTP 200 first, and handle the request in subsequent asynchronous logic.

3. **Unsubscribe from unused event callbacks**

    Subscribing to too many events will result in your event callback service processing a larger number of requests. Unsubscribing from events that are no longer needed can help reduce the number of events processed by your service, reduce the service load, and improve the overall performance of the system.
