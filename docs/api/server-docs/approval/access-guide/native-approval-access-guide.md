---
title: "Native approval access guide"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uIjN4UjLyYDO14iM2gTN"
section: "Approval"
updated: "1675167382000"
---

# Native approval access guide
## Introduction
	This tutorial will show you how to complete a simple approval API process to help you integrate the approval API.
	Tool: Postman
## Guideline
### Procedures
You can integrate the approval API using the Open Platform as follows:

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/821fa46355eed8d7c8ac9f432cce0a70_DkJHxlnjTu.png?lazyload=true&width=1712&height=990)

### Preparations
First, create an app on the [Lark Open Platform](https://open.larksuite.com/).
1. Go to the Lark Open Platform > Developer Console and click "Create Custom App".

![7c4b58ef72c68298270ac5850904d7f5_U0zs2GLJrD.png]((//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/b33e12f92d97a1558a59bdcd0e808c95_Z62VsTN5sN.png)

2. Enter the app information with the name of "Contact Test".

![a2e57acf59d6fcb627a99c4bf7b22d40_YGqMtBc30n.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/c6ba4e7dbcc2814511dae188177bb5dc_6Ju7EChOpk.png?lazyload=true&width=1679&height=793)

3. After creation, the app will be added to the company's custom app directory and can be viewed on the "App Details" page.

![9b5cf58551a3e7a1589c80ac1fbb2746_tJs6tA7zCw.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/853e773229931265af2d233b8399739d_AU62URVxoT.png?lazyload=true&width=1679&height=792)

For information about using the app and granting permissions, see How to develop custom apps in the Open Platform documentation.
       
### Official start
1. Create an approval definition

Before creation, read [How to Approve an App with Lark](https://www.larksuite.com/hc/en-US/articles/360034502513), and Approval app developer's guide to have an initial understanding of approval first. Then, contact the company administrator to create the required approval definition at the [Approval Console](https://www.larksuite.com/approval/admin/approvalList?devMode=on).

Here, we will create a new form that contains two widgets, number and image. Approvers are selected by the approval initiator.

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/4f0e729494e70b1cd0e03b61d75d28be_5rSr4QQH2r.png?lazyload=true&width=1679&height=872)

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/9731e9ff20c81ddd5c98af1f95892b24_w6oInKhWDV.png?lazyload=true&width=1679&height=867)

2. Obtain the Approval Code

Go to [Approval Admin](https://www.larksuite.com/approval/admin/approvalList?devMode=on) to find the approval definition you created. Then, click "Edit" and find the Approval Code in the address bar.

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/e19e91a1de2207c08afc96529933d6cc_015yjunER2.png?lazyload=true&width=1008&height=77)

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/1904b9e5c774a98798e78055a296acd4_vb6juEuj45.png?lazyload=true&width=1679&height=910)

Input the parameter **&devMode=on** at the end of the address bar to enter developer mode. In developer mode, you can specify custom IDs for form widgets and process nodes to facilitate subsequent development work. The editing interface is shown below.

If you set custom IDs, you can use them to specify unique widgets or nodes. Custom IDs must be unique within an approval definition.

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/40b451392e1400e9747a9c2eda797a6b_ew8PD4dei8.png?lazyload=true&width=1679&height=909)

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/176a9992bbb1e648672eee57941ecfb2_5y3fALsFd6.png?lazyload=true&width=1679&height=898)

3. Obtain Open Platform token

On the Open Platform, you can find the app just approved for release and its App ID and App Secret.

Obtain the token that corresponds to your app according to the Access Token documentation of Open Platform. The token is valid for 2 hours. After it expires, you need to obtain a new one. The tenant_access_token is used for all the actions below.

First, set the request header in Postman. This request header is used for all approval API requests. Enter your tenant_access_token in the position indicated by the arrow.

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/395e7a148cada163f947d6dbb6516f18_aTPkJgX7DN.png?lazyload=true&width=1043&height=862)

4. View approval definition

You can use the View approval definition API to query the approval definition you just created.

Use the Approval Code to try the following: Enter the request URL in the Postman address bar, set the request method to POST, and select raw in the Body. Enter all required parameters in JSON format. Enter the Approval Code obtained in Step 2 in the approval_code field.

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/853576f5611e58f49fe2e3d9dd7faae0_15gjueowbF.png?lazyload=true&width=1043&height=862)
5. Upload files

To create a new approval instance based on the approval definition just created through the API, you must first upload a file because the approval definition includes an image widget.

Now, use the Upload files API to upload a file: Enter the request URL in the Postman address bar, set the request method to POST, and select form-data in the Body. Fill in the relevant fields and select File in the content field. Select a local file in the value field.

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/853576f5611e58f49fe2e3d9dd7faae0_15gjueowbF.png?lazyload=true&width=1043&height=862)

After a successful upload, you will receive the following response. Record the code in the response, as it will be used in subsequent instance creation.

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/57fb40084da4fefc1c644511bedfb6e7_k9KW1r2BUi.png?lazyload=true&width=1043&height=862)

6. Create an approval instance (obtain the approval instance's instance_code)

Next, let's start creating an instance. First, read the Create an approval instance document to familiarize yourself with the request body.

This is the request body used to create an approval instance based on the approval definition you just created: Enter the request URL in the Postman address bar, set the request method to POST, and select raw in the Body. Enter all required parameters in JSON format. Enter the Approval Code obtained in Step 2 in the approval_code field. Enter the user_id and department_id according to the initiator's information. Enter the node_approver_user_id_list and form fields according to the instructions in the documentation. This example involves "Approver selected by initiator". See the Approval app developer's guide for more information.

Here, you can use the custom form and node IDs set in Step 2. For example, if you set the custom ID of the number widget to "number" and the initiator selection node ID to "node1", you can replace the default value with "number" in the "form" field and replace the default initiator selection node ID ("855XXX") with "node1" in the "node_approver_user_id_list" field of the following request.

You must note that the default IDs of widgets may change if you modify and re-release the approval definition, but custom IDs won't change unless the widget is deleted or re-added after deletion. Therefore, we recommend using custom IDs during development.

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/57fb40084da4fefc1c644511bedfb6e7_ah5saGcXSU.png?lazyload=true&width=1043&height=862)

After you create an instance, you will receive the following return value:

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/805155bcfc2239724d393002a46d0aad_k1SzgOPZzf.png?lazyload=true&width=1043&height=862)

Here, instance_code is the code of the instance you created.

Now, you can go to the [Approval Frontend](https://applink.larksuite.com/client/mini_program/open?appId=cli_9c7cc8a9a9edd105&path=pages%2Fapproval-list%2Findex%3FselectIndex%3D3) > "Initiated by me" list page to see the approval instance you created through the API (here, the user_id and department_id values are your own).

7. Obtain approval instance details (obtain approval task ID)

Now, you can specify the instance ID to view approval instance details using the API: Enter the request URL in the Postman address bar, set the request method to POST, and select raw in the Body. Enter all required parameters in JSON format. Enter the code returned in Step 6 in the instance_code field.

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/1eb00e4a2f6a6b7375b0f832fefd358c_QUi2agPqFq.png?lazyload=true&width=1043&height=862)
The ID in the returned task_list is the ID of the approval task and used to perform subsequent actions on this approval task.

8. Subscribe and unsubscribe events

For information about subscribing to events, see Approval Event Listening Developer's Guide. After subscribing, you will receive approval messages from the request URL entered in the app.

The subscription is based on the approval definition. After subscribing, you will receive push notifications for all events under the approval definition.

The following is the request to subscribe to the approval definition you just created: Enter the request URL in the Postman address bar, set the request method to POST, and select raw in the Body. Enter all required parameters in JSON format. Enter the Approval Code obtained in Step 2 in the approval_code field.

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/4a6dd830b9178a6ce26bd544dbe9fcd9_6VsV90KzhC.png?lazyload=true&width=1043&height=862)

After subscribing, you can unsubscribe at any time.

The following is the request to unsubscribe to the approval definition you just created: Enter the request URL in the Postman address bar, set the request method to POST, and select raw in the Body. Enter all required parameters in JSON format. Enter the Approval Code obtained in Step 2 in the approval_code field.

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/b70ce489bc7451879c7b07ed87913d4d_QPpkllTlua.png?lazyload=true&width=1043&height=862)

## Conclusion

This ends the tutorial guideline. We have only described the basic applications of the approval API. The API also provides other functions, such as Approval actions and Obtain approval instance IDs in batches. For further information, see the Open Platform approval API documentation.
