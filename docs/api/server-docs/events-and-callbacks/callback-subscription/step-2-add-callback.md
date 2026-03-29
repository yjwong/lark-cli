---
title: "Step 2: Add callback"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/event-subscription-guide/callback-subscription/add-callback"
section: "Events and callbacks"
updated: "1744337104000"
---

# Step 2: Add callback

After configuring the callback request address, you can add the required callbacks according to actual needs. When the callback is triggered later, the Lark server will push the callback information to the callback request address so that you can quickly respond to the callback in the business server.

## Procedure

1. Log in to [Developer Console](https://open.larksuite.com/app).
2. Enter the specified application and select **Development Configuration** > **Events and Callbacks** in the left navigation bar of the application details page.
3. On the **Events and callbacks** page, click the **Callback Configuration** tab.

     ![](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/36491a2b67426433d52e44b6b7b443dd_Q0l2M5pBBj.png?height=606&lazyload=true&maxWidth=600&width=2240)

4. In the **Subscribed callbacks** area at the bottom of the page, click **Add callback**.

     ![](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/901451a96ed8b7d37160379319174aa3_YHm18FCQj0.png?height=357&lazyload=true&maxWidth=600&width=1218)

5. Select the required callback and click **Confirm**.
    
    Currently supports the selection of Pull link preview data, Message card callback communication (legacy), Card postback interaction callback.
    
     ![](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/f745bc9d6689cee343ad7403a71a61a6_i4aVyJCCHa.png?height=1170&lazyload=true&maxWidth=600&width=1646)

6. Publish the application.
    
    
    1. In the left navigation bar, choose **App Versions > Version Management & Release**, then click **Create a version** at the top right of the page.

    
    2. On the **Version Details** page, complete the application version number, default capabilities, update descriptions, and available range configurations as needed.
    
    3. Click **Save** at the bottom of the page, then click **Submit for release** in the pop-up dialog box.
        
        After successfully initiating the request, you need to wait for the enterprise administrator to complete the application review. Once the review is approved, the application will be automatically published, and the callback subscription configurations will take effect.
