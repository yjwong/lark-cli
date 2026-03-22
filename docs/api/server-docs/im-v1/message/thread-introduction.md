---
title: "Thread Introduction"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/thread-introduction"
service: "im-v1"
resource: "message"
section: "Messaging"
updated: "1717574892000"
---

# Thread Introduction

## What is "thread"?
> **Note:** Learn more: [Reply to messages in a thread](https://www.larksuite.com/hc/en-US/articles/622032823229-reply-to-messages-in-a-thread).

You can create a thread for a message in a group and then reply. All discussions under that thread will stay within the thread so other group members aren't disturbed by notifications that are not relevant to them. 

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/1eb04f14eb4e3e3bf7b69bf6fed77cc2_xPZLrEdsY0.png?height=718&lazyload=true&width=1088)

- **Focused conversations**: Your discussions will become more focused. You'll also be able to find all information related to a single discussion within the thread without having to search through group messages. 
- **Fewer disturbances**: Members who aren't interested in the topic won't receive any notifications, cutting down on irrelevant notifications. 

| Message Form | Example | Features |
| --------- | --------------- | ------- |
|Thread | ![img_v3_026u_1548640f-4459-4ef0-81ae-d13edff5fe5g.jpg](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/6d51657d39dd98e1ffe7843308d48a9a_243RGtQV8x.jpg?height=946&lazyload=true&width=1032)| All message discussions are aggregated in the replies of the thread. |
|Messages |![img_v3_026u_0be8678f-87e4-43ee-a531-68dab033962g.jpg](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/60f48f2c0de16a50505f863a56ba2179_DXrg5Tx4ga.jpg?height=706&lazyload=true&width=1344) | Tile display of each message. | ## What is a "topic mode" group and what are its application scenarios?

**"Topic Mode"** **Group**: If you want to create a "product feedback collection group", "sales experience sharing group", or "work order information synchronization group" to keep relevant personnel information synchronized, you can upgrade the ordinary group to a "thread mode group". All messages in the group will be sent in the form of threads, and all replies can only be replied in the form of threads.

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/8a2915ef7d5d782c55f2d8cb0df62873_lEvtc3pJSZ.png?height=1916&lazyload=true&width=2570)

**Application Scenarios**:
- In the work order management scenario, the information discussion of each work order can be precipitated through topic mode groups. Each topic corresponds to a work order information, and the work order information can be sent to the topic mode group through the interface. The contextual discussion of the work order revolves around the topic.
- Similarly, in scenarios such as project management, content management, and user feedback collection and discussion, managing each project, content, and feedback information through topics can make the discussion more focused.

## Thread ID description
Each thread has a unique `thread_id`. The format of `thread_id` is a string starting with `omt_`, for example: omt_d4be107c616a. Through `thread_id`, you can operate on the thread, such as getting all the messages in a thread, forwarding a thread, etc.

**How to get thread_id?**

`thread_id` can be obtained through the message resource. If the `thread_id` field exists in the message resource, it means that it is a message in a thread;
- **Method 1**: In the thread mode group, call Send message, Reply message, Forward a message and other APIs to obtain `thread_id` from the response body;

- **Method 2**: In the chat mode group, call the Reply Message API and fill in the parameter `"reply_in_thread": "true"` to obtain `thread_id` from the response body;

- **Method 3**: Listen to the Receive message event. If the message is a thread message,  the `thread_id` field of the message entity can be obtained from the event body;

- **Method 4**: Call Obtain the content of the specified message or Get chat history API, if the message is a thread message, the `thread_id` field can be obtained from the response body;

  

## What are the APIs capabilities related to "thread"?

  
1. **Set the group as a thread mode group**
     - Synchronously set the group to topic mode when creating a group: When calling the Create a group API, fill in the parameter `"group_message_type": "thread" `;
     - Set the existing group to thread mode: when calling the Update group information API, fill in the parameter `"group_message_type": "thread "`;
2. **Query group message type**
     - Call the Obtain group information API. If the `group_message_type` in the response body is `thread`, it is the thread mode group. If this field is not returned, , indicating that the group chat does not currently support the setting of group message type;
3. **Reply in thread (create a topic)**
     - When calling the Reply message API, fill in the parameter `"reply_in_thread": "true"`;
4. **Forwarding or merged forwarding**
     - Forward a thread to other group: call the Forward a thread API;
     - Forward a message to a thread: When calling the Forward a message API, fill in the parameters `"receive_id_type": "thread_id"`;
     - Merge and forward multiple messages into a thread: When calling Merge forward a message, fill in the parameters `"receive_id_type": "thread_id"`;
5. **Query topic information:**
    
     - Get all messages in the thread: Call the Get chat history API, fill in the parameter `"container_id_type": "thread"` and fill in the thread_id in `container_id` to get all the messages in the thread in pages;
     - Get all messages in the thread mode group:
         1. Call the Get chat history API, fill in the parameter `"container_id_type": "chat"` and fill in the chat_id in `container_id`, you can get all the root messages in the group in pages;
         2. According to each thread root message (the `thread_id` field exists in the message resource) obtained in a , call Get chat history API, fill in the parameter `"container_id_type": "thread"` and fill in the thread_id in `container_id`, you can get all the messages in the thread in pages;
6. **Event notification:**
     - Notify me when the group mode changes: Subscribe to the Group configuration changed event and you can monitor whether the group message type has changed.

## FAQ
1. **Topic Visibility**

	If the group has "New members can see chat history" turned off and the thread was created before the operator entered the group, the topic requires passive subscription by the operator to be visible, for example, other users in the group @ the operator in the thread.
    

2. **The difference between topic groups and thread mode groups**
	::: note
	It is recommended to use thread mode groups.
	:::

	A topic group is a group whose `chat_mode` is `topic`, while a thread mode group refers to an ordinary conversation group whose group message format `group_message_type` is `thread` and `chat_mode` is `group`. The former only supports sending topics after creation, while the latter supports switching between "chat" and "thread". The Create a group API only supports the creation of ordinary conversation groups, and the group message mode can be specified during setting.
