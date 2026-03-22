---
title: "Apply for API permission"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN"
section: "API Call Guide"
updated: "1699608969000"
---

# Apply for API permissions

Through this article, you can learn what  API permissions are and how to apply for them.

## What are API permissions

During app development, you may need to call server APIs of the open platform services, or listen to subscribed events, etc. These operations may require access to tenant and user privacy information or involve manipulation of tenant and user app data. For security reasons, you need to apply for the corresponding permissions for your app. The permissions will take effect only after the app review is approved, and subsequently the app can invoke server APIs or listen to subscribed event list. In simple terms, the API permissions (Scope) of an app determine which server-side capabilities the app can use.

### Permission levels

There are different permission levels for different app types: store apps and custom apps. You can refer to the table below.
  
> **Note:** For a detailed list of all permissions supported by the open platform, you can refer to the API permission list.

#### Custom apps

| Permission level | Description | Review rules |
| --- | --- | --- |
| Automatic approval permissions | Tenant administrators can configure automatic approval permissions based on the actual data governance requirements of the tenant to reduce the review burden. For specific configuration methods, see [How do administrators set up custom app review rules](https://www.larksuite.com/hc/zh-CN/articles/360048488346). | No review is required, and it takes effect immediately after applying. |
| Review-required permissions | For permissions involving sensitive data, include them in the review-required permission list. | After applying, you need to create a version and submit it for review. The app administrator will review and approve it before it takes effect. | #### Store apps

| Permission level | Description | Review rules |
| --- | --- | --- |
| Non-senstive permissions | Data access is of low sensitivity. | For store apps, all permission operations need to go through the following review processes: - App launch review: Reviewed by the Lark open platform. - Tenant installation review: Reviewed by the tenant administrator when updating the version. |
| Advanced permissions | Data access is of high sensitivity. Apps with no significant reason will not pass the review when applying for advanced permissions. | For store apps, all permission operations need to go through the following review processes: - App launch review: Reviewed by the Lark open platform. - Tenant installation review: Reviewed by the tenant administrator when updating the version. | ### Permission application principles

The application for permissions should follow the principle of minimum privilege. Applying for permissions with a large scope may pose a threat to enterprise data security control. Avoid applying for a large number of interface permissions directly without sufficient reasons. Therefore, in the Lark open system, permissions are strictly controlled through a review process.

- **Custom apps**: The permissions applied for by the app developer need to be reviewed by the tenant administrator. The tenant administrator can configure review rules and modes as needed. For more information, see [Custom app review guide](https://www.larksuite.com/hc/zh-CN/articles/360048488346).
    

	![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/84d0d014053876d0f56b95ae7b7c991f_pBVeGdaGsH.png?height=1436&lazyload=true&maxWidth=600&width=2882)

- **Store apps**: If you need to apply for permissions, you need to pass two review processes when publishing apps  on the Lark Open Platform and [installing apps by tenants](https://www.larksuite.com/hc/zh-CN/articles/777448915154).
    
	When a store app is installed for the first time or updated, it will receive an authorization prompt. Tenant administrators can set management rules as needed, refer to [Review application acquisition and use application](https://www.larksuite.com/hc/zh-CN/articles/360023575474).

### Determine the required permissions

When developing an app, you can obtain the corresponding permission information based on the API or event development documents.

Taking the Get user information interface as an example, in the corresponding API document, you can obtain information about the permission requirements and field permission requirements of the interface.

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/b8be909e627b8f3a28ba699f9bbf6485_JFHhMCpSyp.png?height=1134&lazyload=true&maxWidth=600&width=1666)

Among them:

- **Required scopes**: Represents the permission requirements of the entire interface, and there is an “OR” relationship between such permissions. That is, the application can call the interface after applying for any one of the permissions.

- **Required field scopes**: Displays the permission requirements for accessing field data in the response body.
    
     For example, as shown in the figure below, if a self-built application wants to obtain the value of the `user_id` field in the response body, it must enable the **Obtain user ID** permission.
    
     ![](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/415e06b85285796254389710c8b93049_Su0ztXHuDZ.png?height=576&lazyload=true&maxWidth=600&width=1756)

## Apply for API permissions for custom apps

### Step 1: Apply for permissions

1. Log in to the [Developer Consle](https://open.larksuite.com/app), and enter the designated self-built application.

2. Enter **Development Configuration** > **Permissions & Scopes** function page in the left navigation bar.

3. On the **API Scopes** tab, select the permission required to call the API.

	![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/219b1acd98b2ab8b09037f2052ae8b15_G2UCplIfCZ.png?height=880&lazyload=true&maxWidth=600&width=2350)

4. After the selection is complete, click **Add in bulk** in the upper right corner of the **Manage scopes** area.

### Step 2: Test and Debug APIs with approval required permissions

In the app testing and debugging phase, you can debug APIs that require review without waiting for review. You can do this by using the following methods.

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/d54a58d55c987aee770ab477498e1f70_OyFysJ1ccG.png?height=684&lazyload=true&maxWidth=600&width=3200)

#### Method 1: Test APIs using user_access_token

When applying for API permissions in bulk, if a permission supports calling APIs using the app developer's user_access_token without review (as shown in the figure below), you can click **Confirm** to use the user_access_token to debug APIs without the need to publish the app.

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/91fcb1f2c86a05b2b491f9305b916adb_83hevF8roQ.png?height=900&lazyload=true&maxWidth=600&width=1174)

Take List roles interface as an example. Before calling this API, you need to apply for **View, comment, edit and manage Bitable** or **View, comment, and export Bitable** permissions on the app’s **API Scopes**  page.

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/ad0b30c933006f8aba7ca39d4ce81fcf_sFRuOHSFVk.png?height=356&lazyload=true&maxWidth=600&width=1138)

After applying, you don't need to publish the app and wait for the approval. Instead, you can directly use the user_access_token of the app developer to debug the API.

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/1fe8a205db811a8a2cfe32408949e18e_ufpwBbszKa.png?height=582&lazyload=true&maxWidth=600&width=1850)

#### Method 2: Test APIs using test company function

Some APIs do not support authorization using user_access_token, and a small number of sensitive permissions in APIs that support user_access_token authorization cannot be debugged without review. You can find prompts when applying for such permissions (as shown in the figure below).

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/e2094d7ea16818176e4d976cba37fc36_Uy76NsiPRv.png?height=1090&lazyload=true&maxWidth=600&width=1168)

In this case, you can configure a test company for the app, switch to the test version . For specific operations, see Test companies and users. After switching, go to the **Development Configuration** > **Permissions & Scopes** page, and apply for the specified permissions on the **API Scopes** tab. The permissions applied for are all automatically approved and take effect immediately after clicking **Confirm**.

> **Note:** If you apply for permissions related to **Contacts** or **Lark HR (Enterprise)** APIs and need to call the corresponding APIs using the app identity (tenant_access_token), you also need to configure the corresponding data permissions for the app. For details, see Configure data permissions.

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/2272e952c1f0083b8a817c41fb3aa36d_1azbWVroJC.png?height=852&lazyload=true&maxWidth=600&width=2298)

### Step 3: Release the app

1. After ensuring that the permissions applied for during the test phase meet business expectations, go to the **App Versions** > **Version Management & Release** page, and click **Create a version**.
    

> **Note:** In the scenario of debugging APIs using a test company, the permissions enabled will not take effect in the official version of the app. Therefore, after completing testing and debugging, you need to switch back to the official version, enable the same permissions again, and then create a version release request.
    
2. On the **Version Details** page, configure the following fields, and click **Save**.

	![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/947492f5af79473afbdd7bf9861ef9e4_omBaVyjAXX.png?height=1298&lazyload=true&maxWidth=600&width=2198)

     - **App version**: Customize the app version number, for example: 1.0.0.
    
     - **Update notes**: Customize the update details of the current version.
    
     - **Features** and **Scope changes**: Check and confirm whether the added capabilities and permissions meet expectations.
    
     - **Availability**: The availability scope of the app. Click Edit to adjust the availability scope. For more information , see Configure availability.
    
     - **Reason for request**: It is used to help reviewers learn more about the app version.

3. In the pop-up dialog box, click **Submit for release**.

	![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/58bcda0547f9755fed516b397825cad2_0oGr3Jaaqg.png?height=748&lazyload=true&maxWidth=600&width=2226)

4. Wait for the enterprise administrator to review the application.
    
     All API permissions will take effect only after the app has passed the review.

	![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/b4c80cd8b6293a24da3eb03f8044468f_pUWSaDxNk3.png?height=562&lazyload=true&maxWidth=600&width=2316)

## Apply for API permission for store apps

### Step 1: Apply for permissions

1. Go to the specified store app in the [Developer Console](https://open.larksuite.com/app).

3. Go to **Development Configuration** > **Permissions & Scopes** in the left navigation pane.

3. Apply for permissions in the **Manage scopes** area.

	![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/34b97a2b34049c3826b3410551927460_9XyouCMpkW.png?height=866&lazyload=true&maxWidth=600&width=2874)

4. After the selection is complete, click **Add in bulk** in the upper right corner of the **Manage scopes** area.

5. In the **Reminder** dialog box, after confirming that the requested permission list is correct, click **Confirm**.

### Step 2: Test and debug APIs with approval required permissions

During the test and debugging phase of the app, you can generate and test the version through the test enterprise and personnel function of the store app. The permissions activated based on the test version can take effect without waiting for approval, so as to perform API debugging.

1. In the left navigation bar, select **Development Configuration** > **Test Companies and Users**, and create a test enterprise.
    
> **Note:** For details, see 6. Test store application.

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/65312c3edf52d2a73d906074fb6c8f04_FcAkMpxngW.png?height=490&lazyload=true&maxWidth=600&width=2240)

2. In the left navigation bar, select **App Release** > **Version Management & Release**, and then click **Create a version**.

	![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/66d04da4231b9afef2f60a111f7eda6b_roKzrzloCh.png?height=386&lazyload=true&maxWidth=600&width=2228)

3. Configure the version number and version description in turn, and confirm the integrity of the permissions to be activated, and finally click **Save** at the bottom of the page.

	![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/4a1430438cf5dec7b390e2aa5aac9fe9_hlmSPRaHxN.png?height=1152&lazyload=true&maxWidth=600&width=2422)

4. On the version details page, click **Set as test version**, and complete the setting of the test version in the pop-up dialog box.

     After the configuration is complete, the version status will change to **Testing**. At this time, the app has been installed in the corresponding test company, and the APIs can be called in the test company without review.

	![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/7281550ef8ecc67ce136854c84b1648f_YVduMRP1AG.png?height=352&lazyload=true&maxWidth=600&width=2284)
     
### Step 3: Release the app

After ensuring that the permissions applied for during the test phase meet business expectations, go to the **App Release** > **Version Management & Release** page to publish the app.

The official release of the store app is divided into targeted listing, full listing, and non-public listing. You need to choose the appropriate release option based on your actual business situation. For details, see 7. Publish store apps.
