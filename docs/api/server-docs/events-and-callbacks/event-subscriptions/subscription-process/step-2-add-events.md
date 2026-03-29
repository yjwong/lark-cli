---
title: "Step 2: Add events"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uYDNxYjL2QTM24iN0EjN/event-subscription-configure-/subscription-event-case"
section: "Events and callbacks"
updated: "1750063783000"
---

# Step 2: Add events

After adding the events to subscribe to, the open platform will push event information to the application when events occur, enabling the application to respond quickly.

## Event versions

The Lark Open Platform provides two versions of events: v2.0 and v1.0. The v2.0 version of the events is more comprehensive and reasonable. If you subscribe to the same event in different versions simultaneously, the application will receive duplicate events. Therefore, it is recommended to only subscribe to v2.0 events and unsubscribe from v1.0 events. For details on the event structures of different versions, refer to Event structure.

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/600871898bbf8cba49283c7d2f896bc4_RfUQsGCU7E.png?height=478&lazyload=true&maxWidth=600&width=1270)

## Subscription identity

To enhance the security of open capabilities, the Lark Open Platform has designed an access token mechanism. When calling APIs to get application resources, the identity of the caller needs to be authenticated using an access token, informing Lark who the current caller is and what tenant data they are accessing. For details on how to choose and obtain different types of access tokens, refer to Get access token. Similarly, when adding events, the access token mechanism is required, using different identities to subscribe to events.

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/2b372384af8cf1a42cf52cc91d005672_VaxcX6IQR3.png?height=360&lazyload=true&maxWidth=600&width=1270)

Introduction to Subscription Identities:

| Subscription identity type | Description | Scenario example |
| --- | --- | --- |
| Tenant Token-Based Subscription | Subscribe to events using tenant_access_token. This does not require user login and directly obtains event data using the application identity. The scope of accessible data is determined by the application's own data permissions. | Suppose there is an application named "My bot." When the application subscribes to cloud document events with a tenant_access_token subscription, it can only subscribe to document changes where "My bot" is the owner or administrator, and it cannot perceive changes in other documents. |
| User Token-Based Subscription | Subscribe to events using user_access_token. This requires user login and application authorization, and the scope of accessible event data is determined by the user’s own data permissions. | Suppose there is an application named "My bot." When the application subscribes to cloud document events with a user_access_token subscription, and the token represents a user named "Li Jian," it can only subscribe to document changes where "Li Jian" is the owner or administrator, and it cannot perceive changes in other documents. | Most events require subscribing using application identity. The scenarios that require subscribing using user identity are as follows:

- **Docs**: This service’s events support subscription through either application identity or user identity. You need to choose the appropriate identity according to the actual scenario.
- **Calendar**: This service’s events only support user identity subscriptions.
- **Email**: The events for this service only support subscription through user identity (event subscription for email is currently in beta testing; please be patient if you cannot use it).

## Procedure

1. Go to the [Developer Console](https://open.larksuite.com/app).
2. In the application list, click on the specified application to go to the application details page.
3. Go to **Events & Callbacks > Event Configuration** page, in the **Add Events** area, click **Add Event**.

    ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/549adfa840fc99b7ea74ca5da300620a_rx9jtdJHJ6.png?height=1038&lazyload=true&maxWidth=600&width=2770)

4. In the **Add Event** dialog box, choose **Tenant Token-Based Subscription** or **User Token-Based Subscription**.

5. After selecting the events you want to subscribe to, click **Confirm**.

    - Click **Add Scopes** to enable the required permissions for the events with one click.
    
		![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/af0c5687991c39c2d1f0480b1fb4f3fc_hG74MIYpPE.png?height=424&lazyload=true&maxWidth=600&width=1182)
        
    - If you choose **Maybe Later**, you can return to the **Event Configuration** tab later, go to the **Events added** area, click the required permissions for the event, and then click **Add**.

        ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/22c07a3f0e26c646d0693250b5ebd70d_7nrjolXlO9.png?height=1014&lazyload=true&maxWidth=600&width=1838)
    

    
6. Publish the application.
    
    :::warning
    Even if the relevant permissions have already been enabled (approved or exemption permissions obtained), you still need to publish the application for the configuration to take effect after filling in the event subscription method and adding events.
    :::
    
    1. In the left navigation bar, choose **App Versions > Version Management & Release**, then click **Create a version** at the top right of the page.

        ![](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/a01b27e311a91764bafe231c6db7535f_10pBhVzbZs.png?height=652&lazyload=true&maxWidth=600&width=2184)
    
    2. On the **Version Details** page, complete the application version number, default capabilities, update descriptions, and available range configurations as needed.
    
    3. Click **Save** at the bottom of the page, then click **Submit for release** in the pop-up dialog box.
        
        After successfully initiating the application, you need to wait for the enterprise administrator to complete the application review. Once approved, the application will be automatically released, and the event subscription configurations will take effect.
