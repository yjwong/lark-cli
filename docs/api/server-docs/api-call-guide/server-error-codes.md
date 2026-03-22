---
title: "Server Error Codes"
url: "https://open.larksuite.com/document/ukTMukTMukTM/ugjM14COyUjL4ITN"
section: "API Call Guide"
updated: "1646296632000"
---

# Server Error Codes
> The following is a list of common error codes. Error codes not listed here can be found in the specific API interface documentation. You can also search for error code descriptions and troubleshooting suggestions globally through the search function in the upper right corner.

| Error Code | Description | Troubleshooting Suggestion |
| --- | --- | --- |
| 1002 | Failed | Deleted file cannot get document content |
| 1500 | internal error | Internal error, please try again later. If it still occurs then you can [consult customer service](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 1503 | internal error | Internal error, token update, but no authentication result is returned. Please try again later. If it still occurs then you can [consult customer service](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 1551 | internal error | Get tenant ID error, please try again later. If you have any questions,  you can [consult customer service](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 1557 | internal error | Get open user error, please try again later. If you have any questions,  you can [consult customer service](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 1642 | internal error | Internal error, failed to update the session, please try again later. If it still occurs then you can [consult customer service](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 1663 | internal error | Internal error, please try again later. If it still occurs then you can [consult customer service](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 1665 | internal error | Internal error, please [consult customer service](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 1668 | internal error | Internal error, please try again later. If it still occurs then you can [consult customer service](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 2200 | internal error | Internal error, please try again later. If it still occurs then you can [consult customer service](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 4000 | bad request | Please make sure request method, request query, request body is correct, e.g. rich text format |
| 4001 | Invalid token, please refresh | Check whether the token is filled in correctly and whether it is expired |
| 4002 | definition not found | Check that the approval code is correct |
| 4003 | instance not found | Check that the approval instance ID is correct |
| 4004 | user not found | Check userId is correct |
| 4005 | method not allowed | Check if the request method is POST |
| 4006 | service exception | Check service availability |
| 5000 | internal error | Reduce qps, wait a while then try again |
| 9499 | bad request OR invalid param | The request parameter is wrong, please check the correctness of the parameter |
| 10001 | Your request contains an invalid request parameter | Invalid request parameter |
| 10002 | Bot can NOT be out of the chat | Add the bot to the chat and re-send the request |
| 10003 | invalid parameter | If a request parameter is missing or incorrect, see the error message returned for the request for more information on the error |
| 10004 | user not found | The specified user was not found |
| 10005 | app id unauthorized | There is an error in the app information or tenant information in the current request |
| 10009 | get open_chat_id fail | The open chat ID is illegal |
| 10010 | send message forbidden | Sending inter-organization messages is forbidden |
| 10012 | get app access token fail | Please refer to API Access Token |
| 10013 | get tenant access token fail | Please refer to API Access Token |
| 10014 | app unauthorized | Check if the app is deactivated |
| 10015 | wrong app secret | Check the app secret for the current app in Lark Developer |
| 10016 | app resend fail | Check if the app is ISV |
| 10017 | Bot is NOT the owner of the resource | Bot is not the owner of the resource and cannot get the resource content |
| 10019 | urgent type not support | The type of urgent contact is unknown; at present, urgent app messages, text messages and phone calls are supported |
| 10020 | message id not exist | The message ID is invalid |
| 10023 | urgent scope unauthorized | The current request does not apply for relevant permissions. The opening of the permission needs to be submitted and released, and it will take effect after being reviewed by the administrator. For details, see How to open API scope. For applications in the development and testing phase, you can use the "Beta Application" function. You do not need to directly release the version to apply for permission. It can be completed in the test administrator. For details, see Tutorial. |
| 10029 | chat_id not exist | The open chat ID is invalid |
| 10030 | bot not in chat | Add the bot to the chat and re-send the request |
| 10032 | valid user is null | The user list in the request is empty, check your request parameters |
| 10034 | chat tenant id mismatch tenant access token | Requests cannot be sent between organizations |
| 10037 | open_message_id not exist | The open message ID is invalid |
| 10100 | invalid request | It is recommended to check whether the request parameters are legal |
| 10101 | internal error | Internal error, please try again later. If it still occurs then you can [consult customer service](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 10105 | internal error | Internal error, get open user error, please try again later. If you have any questions, you can [consult customer service](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 10200 | crop failed | Image cropping failed, please try again later. If it still occurs then you can [consult customer service](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 10204 | wrong code | The code is reused or expired |
| 10400 | bad request | Generally do not need to investigate |
| 10500 | internal server error | Recommended server troubleshooting |
| 11000 | get sso access token fail | See Authorization Tokens |
| 11001 | check security token fail | See Authorization Tokens |
| 11100 | check chat by open chat id fail | The open chat ID is invalid |
| 11101 | open_id not exist | The open user ID is invalid |
| 11102 | get open_id fail | The open user ID is invalid |
| 11103 | open department id not exist | The open department ID is invalid |
| 11104 | get open department id fail | The open department ID is invalid |
| 11105 | user_id not exist | The employee ID is invalid |
| 11106 | get employee_id fail | Failed to get the employee ID |
| 11200 | only admin can modify chat name | The user is not the group owner; only the group owner can update the group chat information |
| 11201 | bot is not chat owner | Add the bot to the chat and re-send the request |
| 11202 | only admin can add chatter | The bot is not the group owner and cannot add members |
| 11203 | send employee ids permission denied | Check that the “Get employee ID” permission has been enabled in Lark Developer |
| 11204 | send department list permission denied | Check that the “Get department ID” permission has been enabled in Lark Developer |
| 11205 | app do not have bot | Check the bot feature has been enabled in Lark Developer |
| 11206 | user not in chat | Add the user to the chat and re-send the request |
| 11207 | app is not usage | Check the app’s status in the backend |
| 11208 | only group admin can remove chatter | The bot is not the group owner and cannot remove members |
| 11209 | app not exist | App does not exist |
| 11210 | app usage info not exist | Contact the tenant administrator and check if the app is installed |
| 11211 | add chatters is empty | The list of members to add to the chat is empty |
| 11212 | invalid chat_id or user_id | The chat ID or user ID is incorrect |
| 11213 | chat not found | The specified chat does not exist |
| 11214 | image not given | The document list is empty |
| 11215 | chat_id is nil | The chat ID is invalid |
| 11216 | get chat id fail | The chat ID is invalid |
| 11217 | this group is set to only admin can add new member | The current user is not the group owner and cannot add members to the group |
| 11218 | the number of robots exceeds the limit | When the maximum number of bots allowed in a chat is reached, no more bots can be added to the chat |
| 11219 | cross tenant chat forbidden add bot | Bots cannot be added to chats between organizations |
| 11220 | only admin can disband group | The current user is not the group owner and cannot disband (i.e. close) the chat group |
| 11221 | this image does not belong to you | The bot does not have permission to get images that don’t belong to the bot |
| 11222 | bot owner not in chat | If the app owner is not in the group chat, the bot can’t be added to it |
| 11223 | do not have im write permission | The operator does not have permission to send messages. Enable the permission and re-send the request |
| 11224 | message_id invalid | The message ID is invalid |
| 11225 | the bot is invisible to the user | When the bot is invisible to the user, it can’t operate the user. You can edit the visibility of the application to users and publish it in the developer background-application release-version management and release. |
| 11226 | app_id is nil | The app ID is invalid |
| 11227 | image key not exist | The image key is invalid |
| 11228 | chat is not group | The current chat is not a group type chat |
| 11229 | no permission | Check that the app has enabled the permission to access the user’s information in Lark Developer |
| 11230 | internal error | The bot is forbidden from sending the current type of message |
| 11231 | chat not found | The requested chat does not exist |
| 11232 | create message service trigger rate limit | The request frequency has triggered the frequency control policy, see Frequency control |
| 11233 | create message chat trigger rate limit | The request frequency has triggered the frequency control policy, see Frequency control |
| 11234 | this message do not belong you | The bot does not have permission to read the status of a message that doesn’t belong to it |
| 11235 | this group is set to only admin can mention @All | The current chat forbids the use of @ mentions |
| 11236 | user is resigned | The user has left the organization |
| 11237 | group chat dissolved | The specified group chat has been disbanded and members cannot be added |
| 11238 | can not recall too old message | Messages that were sent too long ago can’t be recalled |
| 11239 | permission denied | If you try to get the information of another tenant, the operation will fail |
| 11240 | Robot not turned on | Check the bot feature has been enabled in Lark Developer |
| 11241 | no user_info scope | Check that the app has enabled the permission to access the user’s information in Lark Developer |
| 11242 | this chat has been banned | The chat has been banned; you cannot reply to messages in the chat |
| 11244 | file key not exist | File does not exist |
| 11245 | this file does not belong to you | If you try to get the information of a file that doesn’t belong to the user, the operation will fail |
| 11246 | create message content fail | See the error message returned for the request for more information on the error |
| 11247 | internal send message trigger rate limit | The request frequency has triggered the frequency control policy, see Frequency control |
| 11248 | message is sensitive | When the content of the message is sensitive, sending will fail |
| 11249 | bot not found | The specified bot is invalid |
| 11251 | chat name too short | The chat name must have at least 2 characters |
| 11252 | chat information review fail | If the message was not approved, check the message for sensitive information |
| 11310 | get card message fail | See the error code and message returned for the request and the card error code documentation for more information |
| 11311 | urgent message fail | 加急消息错误，仍然出现可以咨询客服 |
| 11312 | messages do not pass the audit | messages do not pass the audit，please check the content of message |
| 12001 | Your request contains an invalid request parameter. | Invalid request parameter |
| 12002 | Non-chat-owner can NOT edit chat information in the current situation. | The chat owner sets "Only the owner can edit chat information", non-chat owners cannot update the information |
| 12008 | Your request specifies a chat whose type is NOT supported currently | Chat type is not supported |
| 12011 | Operator can NOT be out of the chat. | The bot or authorized user is not in the chat, please add it to the chat |
| 12015 | Your request specifies a member_id_type which is NOT supported. | Only support open_id/user_id/union_id/app_id as request parameters |
| 12016 | Non-chat-owner can only edit certain parts | The chat owner sets "Only the owner can edit chat information", non-chat owners cannot update the information |
| 12018 | Updating announcement failed | Fail to update announcement, please check the log for specific reasons |
| 18017 | create urgent message fail | Failed to create urgent message, please make sure the request parameters are correct |
| 18033 | upload image fail | Failed to upload image, please make sure the request parameters are correct |
| 18034 | get chat id fail | Failed to get chat id, please make sure the request parameters are correct |
| 18035 | image_key require | Failed to get image,please make sure the request parameters are correct |
| 18053 | create message fail | Failed to create message, please make sure the request parameters are correct |
| 18054 | create message content fail | Failed to create the message content, please check carefully whether the message content involves sensitive information |
| 18113 | thread not found | Thread not found, please check if the current thread exists |
| 18117 | reach limit of chat size | reach limit of chat size |
| 18121 | create request is being processed | The same message create request is being processed, please try again later. |
| 19002 | incorrect input format or can't find msg_type | check input data format and msg_type |
| 19036 | The message exceeds the size limit of 30KB | check message size, make sure that the message size is below 30k |
| 20001 | invalid request | Please ensure that the request method, request information, request data format, etc. are correct |
| 20005 | invalid access_token | Please make sure the access_token is correct and valid. For details, please refer to API Access Token |
| 20007 | generate access_token fail | Please make sure that the code is not consumed repeatedly or expired; the code obtained by the applet tt.login() cannot be used for the server API to obtain user information, and vice versa. |
| 20009 | tenant app not installed | The app is not installed in the current tenant. |
| 20021 | user resigned | User has resigned, check user status. |
| 20022 | user frozen | User has been frozen, check user status. |
| 20023 | user not registered | User has not registered, check user status. |
| 20024 | code app_id not match | Please make sure that the appID of the generated code is consistent with the appID corresponding to the app access token in the request header |
| 30005 | interface is offline | Switch to the new version interface |
| 40001 | param error | param error |
| 40001 | market place app is not allowed to update user or department | Lark App Directory apps cannot modify Contacts, so attempts will fail |
| 40002 | process root dept error | Please check whether 0 is filled in the department ID of the request parameter |
| 40003 | no department authority | Check if permission to access the department’s Contacts has been granted. If you need to access the information, please contact the administrator to ask them to grant permission to the department’s Contacts |
| 40003 | internal error | Internal error. Please try again later. If it still occurs then you can [consult customer service](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D)  (Note:Use this meaning when the error code appears in the V3 version interface of the address book) |
| 40004 | no dept authority error | No department authority, check whether the department is within the address book authority |
| 40006 | internal error | Internal error. Please try again later. If it still occurs then you can [consult customer service](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 40007 | user_id %v is not exist | Check the source of the user ID and if the user belongs to the current tenant |
| 40008 | deptInfo is null error | Please [consult customer service](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D)   (Note:Use this meaning when the error code appears in the V3 version interface of the address book) |
| 40008 | open_id %v is not exist | Check the source of the user’s open_id |
| 40009 | page size is more more than 50 | The page_size parameter of the paging request is illegal, please check the size set by page_size |
| 40010 | chat id is invalid error | chat id invalid,please check |
| 40011 | page size is invalid error | please check page size value |
| 40012 | page token is invalid error | check page_token (Note:Use this meaning when the error code appears in the V3 version interface of the address book) |
| 40012 | no parent dept authority error | No parent department authority, check whether the department is within the address book authority |
| 40013 | param error | See the message for the cause of the error; if you have any questions, contact the administrator |
| 40014 | dept name can not be nul error | The department name cannot be empty, please check the department name setting in the request (Note:Use this meaning when the error code appears in the address book V1 or V2 version interface) |
| 40014 | no parent dept authority error | No parent department authority, check whether the department is within the address book authority (Note:Use this meaning when the error code appears in the V3 version interface of the address book) |
| 40016 | name can not be null in updateRequest error | Update the interface, need to bring the name field |
| 40018 | duplicate name error | The names of child departments of the same parent department do not conflict |
| 40021 | not a same request error | The two operations are not the same request, please check whether the request value has changed (Note:Use this meaning when the error code appears in the V3 version interface of the address book) |
| 40021 | forbid update department id | The department’s custom ID is the only department identifier. If the organization has other apps using Contacts, it can’t be modified |
| 40031 | task_id is not exist | Check the source of the task_id |
| 40032 | task_id invalid | Check the source of the ID and if it is the task_id generated when the current organization uploaded the task |
| 40040 | invalid page_token | The page_token parameter in the page request is invalid; check the source of the page_token |
| 40041 | request page_size invalid | The requested page size does not comply with the API requirements; check the API documentation and pass a valid value |
| 40042 | batch request object count error | The number of the batch request does not comply with the API requirements; check the API documentation and pass a valid value |
| 40051 | open_id %v is not valid | Check the source of the user’s open_id and if the open_id belongs to this app |
| 40052 | department id %v is not exist | Check the source of the department ID |
| 40053 | role_id is not exist | Check the source of the role ID |
| 40054 | user_id or open_id is needed | Either the user_id or the open_id must be passed |
| 40101 | mobile is already exist | Mobile numbers must be unique to an organization, so if the mobile number record has already been created, the user can’t use it again to create another new user |
| 40102 | email is already exist | Email addresses must be unique to an organization, so if the email address record has already been created, the user can’t use it again to create another new user |
| 40103 | mobile and email cp info not match | A user’s contact method cannot belong to two different Lark accounts, so adding failed. If you have any questions,  you can [consult customer service](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 40104 | mobile is not valid | Check the mobile number is in the correct format |
| 40105 | name can’t be null | Name fields, including the department and user name, cannot be empty |
| 40106 | email is not valid | Check the email address format is valid |
| 40107 | user count exceed certification tenant limit | Users cannot be created when the number of users exceeds the user limit for unverified organizations. If you have any questions,  you can [consult customer service](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 40108 | user count exceed bill limit | Users cannot be created when the number of users exceeds the user limit in that organization’s billing plan.  If you have any questions,  you can [consult customer service](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 40109 | neither email nor mobile is not provided | To create a user, either the mobile number or email address must be passed |
| 40110 | mobile is required | The mobile number is required to create a user |
| 40111 | user_id is already exist | The user_id is a unique identifier of a user within an organization and can’t be duplicated |
| 40112 | user_id %v is not vaild | The user_id format used to create the user is invalid |
| 40113 | must assign department for user | When creating a user, a department under which to create the user must be specified |
| 40114 | name contains sensitive words | If the name contains sensitive information, it can’t be created. If you have any questions,  you can [consult customer service](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 40115 | invalid join_time | The date the user joined the company is invalid |
| 40116 | invalid gender | Check the parameter descriptions in the API documentation for which parameters are allowed |
| 40117 | idp_type is required | If an IDP organization creates a user, the idp_type is required |
| 40118 | idp_type is invalid | Check the idp_type has been entered correctly |
| 40119 | support set or update user_id at most once | The user’s custom ID can only be changed once if it is not set when the user is created; if it was set when the user was created, it cannot be subsequently changed |
| 40120 | custom_attr id is not set | When setting custom user properties, a property ID must be set. The property ID can be obtained by querying the organization custom property API |
| 40121 | custom_attr %v value not set | When setting a custom property value, the value field must be passed |
| 40122 | custom_attr id %v is not exist | Check the source of the property ID. The property ID can be obtained by querying the organization custom property API |
| 40123 | href custom_attr %v text can’t be null | When setting an HREF-type custom property, the text field is required |
| 40124 | href custom_attr %v url can’t be null | When setting an HREF-type custom property, the URL field is required |
| 40125 | user position info invalid | When creating a user position, the position name, code and department are required. If there is a department leader, the leader’s ID must also be specified |
| 40126 | position department invalid, should be one of user’s department | The department ID of the user’s position must be the ID for a department to which the user belongs |
| 40127 | position code is not unique | The multiple position codes of the user must all be unique |
| 40128 | user has one main position at most | Only one main position can be set per user |
| 40129 | request contains users of different tenant | Check the source of the user ID and if it belongs to the same organization |
| 40130 | update user department conflict with position department, should update position department at the same time | The user must belong to the same department as the position. If you update the user’s department, the department to which the position belongs must also be updated, or the information cannot be updated |
| 40131 | update user position department conflict with user department, should update user department at the same time | The user’s position must belong to the same department as the user. If you update the department of the user’s position, it must belong to a department to which the user also belongs, or the information cannot be updated |
| 40132 | order department invalid | The department ID in the requested user sort must be one of the user’s department IDs |
| 40133 | idp check with account failed | Failed to check the IDP account |
| 40134 | create account failed | Failed to create the account |
| 40135 | idp user_id is required | The user_id is required to create an IDP account |
| 40136 | update user’s email and mobile empty is forbiden | When updating the user’s information, the user’s email address and mobile number can’t both be removed |
| 40137 | can’t update unactive user’s mobile empty | When updating the user’s information, an inactivated user’s mobile number can’t be removed |
| 40138 | %v is not valid | When a user leaves, the user resource recipient must be a current user of the current organization |
| 40139 | unable to join multiple departments, please upgrade Organizational Structure Visible info | Please upgrade the "Organizational Structure Visible Range" rule in the enterprise management background |
| 40140 | can not set leader to oneself | Please check the user's immediate superior parameter value |
| 40141 | can not set user position | The current team does not support setting user position information. If you have any questions,  you can [consult customer service](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 40142 | user department num exceed the limit | user's department cannot exceed the limit (50) |
| 40143 | unable to join multiple departments, please enable multiple departments feature | The current company does not support users to join multiple departments at the same time.  If you have any questions,  you can [consult customer service](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 40144 | department user count exceed | The number of direct members of at least one department in the created or updated department exceeds the limit. Please check the number of department members in the request |
| 40151 | custom department id %v is invaild | The custom department ID used to create or update the department is invalid |
| 40152 | custom department id %v is not unique | The department ID must be unique within the organization |
| 40153 | department name should be unique under same department | Department names must be unique under the same department |
| 40154 | department unit %v is not exist | Check the department unit ID |
| 40155 | duplicate department order | Department orders can’t be duplicated under the same department |
| 40156 | department id is required | No department ID was specified in the request |
| 40157 | forbidden operation to root department 0 | The root department is a fictitious department; it can’t be updated or deleted in a request |
| 40158 | support update custom department id at most once | The custom department ID can only be changed once |
| 40159 | department has active members, can’t delete | The department to be de |
| 40160 | department has sub department, can’t delete | When the department to be deleted has sub-departments, it is not allowed to delete the department directly, please delete the sub-department first and then delete the department |
| 40161 | can’t get root department info | The root department is a virtual department, and its details cannot be queried |
| 40162 | departments with more than 500 sub departments are not allowed to call recursively | Please implement the recursive query logic by yourself |
| 40163 | can not create department unit | The current team does not support the creation of department units. If you have any questions,  you can [consult customer service](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 40164 | department child count exceed | The parent department already have enough child department, please the count of child |
| 40170 | unit repeat, unit_type and unit_name is exist, can’t create | Under the same enterprise, the name must be unique under the same department unit type |
| 40171 | department unit name invalid | Illegal department name, please refer to the API document for description |
| 40172 | department unit type invalid | Illegal department type, please refer to the API document for description |
| 40173 | custom_id is already exist | In the same enterprise, the department unit id must be unique |
| 40174 | department unit is still using, can’t be deleted | The department unit is in use, for example, there are other departments associated with the department unit, it is not allowed to delete the department unit at this time |
| 40175 | department attach to more than one department unit | Currently only supports one department to associate at most one other department unit |
| 40180 | duplicate user_id in request | user_id is the unique identification of the user in the enterprise, and no repetition is allowed; please check the batch request parameters |
| 40181 | add user, leader must not user himself | Create a user, the user’s immediate superior cannot be the user himself |
| 40182 | MSyncUser must user_id must be assigned | To synchronize users in batches, user_id must be specified |
| 40183 | leader user_id circular reference | When synchronizing users in batches, and leader information forms a loop; for example, the same request is used to create users A and B, and the leader of A is designated as B, and the leader of B is designated as A. Please check the request parameters |
| 40184 | add department, parent_id must not equal to current department id | Create departments in batches, the parent department of the department cannot be the department itself |
| 40185 | add department, parent_id must be assigned | To create a department, the parent department must be specified |
| 40186 | duplicate department id in request | When creating departments in batches, the custom department_id of the departments are duplicated in the request, please check the request parameters |
| 40187 | batch request, parent department id incompatible | When creating departments in batches, department parameters conflict |
| 40501 | param required | check if the param is required |
| 40502 | invalid parameter | If a request parameter is missing or incorrect, see the error message returned for the request for more information on the error |
| 41001 | mobile already exist error | The mobile phone number must be unique within the enterprise, and users who have created the mobile phone number in the enterprise are not allowed to create new users |
| 41002 | email already exist error | The email must be unique in the enterprise |
| 41003 | user acount conflict error | The user's contact information belongs to two different Lark accounts. If you have any questions,  you can [consult customer service](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 41004 | mobile is invalid error | The phone number is illegal, please check if it is in the correct format |
| 41005 | email is invalid error | Email address that is not a valid email address, please check the validity of the email address |
| 41006 | no user name error | Username is missing or empty |
| 41007 | exceed uncertain tenant seat limit error | The number of users exceeds the limit of the number of non-certified companies, and users cannot be created.  If you have any questions,  you can [consult customer service](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 41008 | exceed bill seat limit error | The number of users exceeds the limit of the number of non-certified companies, and users cannot be created.  If you have any questions,  you can [consult customer service](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 41009 | no email or mobile error | Create a user, specify at least one of the mobile phone number and mailbox |
| 41010 | no mobile error | mobile phone number is required |
| 41011 | userID already exist error | user id is unique in an enterprise |
| 41012 | user id is invalid error | Characters between user_id1-64 bytes |
| 41013 | exceed userID update limit error | please retry later. |
| 41014 | user name sensitive error | Username contains sensitive information. If you have any questions,  you can [consult customer service](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 41015 | idp type invalid error | Invalid login type |
| 41016 | department has too may users error | There is a limit on the number of members directly under the department. If the limit is exceeded, it is not allowed to add |
| 41017 | department not exist | department not exist |
| 41018 | position info is invalid error | please check |
| 41019 | position department is invalid error | Invalid position department, check department information |
| 41020 | position code has already exist error | invalid position code |
| 41021 | position multiple main count error | A user can only have one master post |
| 41022 | user tenant not match error | Check whether to use other company's credentials to access the current company's resources |
| 41023 | update department conflict position department error | The user’s post department must be consistent with the user’s department. To update the user’s department, the corresponding post department must be updated at the same time, otherwise the update operation will be blocked. |
| 41024 | update position department conflict department error | The user’s post department must be consistent with the user’s department. To update the user’s department, the corresponding post department must be updated at the same time, otherwise the update operation will be blocked. |
| 41025 | order department invalid error | The department ID in the requested user ranking information must be one of the user's department ID |
| 41027 | create account failed error | Please try again later. If it still occurs then you can [consult customer service](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 41028 | user multi department need upgrade visibility error | The use of multiple departments requires an upgrade of the visibility version, which can be upgraded in the visible range of the corporate management background-organizational structure. If you have any questions,  you can [consult customer service](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 41029 | create or update user multi department error | The multi-department function may not be turned on, consult the company administrator |
| 41030 | set leader to oneself error | It is not allowed to set one's superior as oneself |
| 41031 | position feature not enable error | If you have any questions,  you can [consult customer service](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 41032 | user multi department feature not enable error | The switch to allow users to join multiple departments is not turned on. The current enterprise does not support users to join multiple departments at the same time. If you have any questions,  you can [consult customer service](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 41033 | user in too many departments error | Does not support users belonging to more than 50 departments at the same time, please check |
| 41034 | email prefix already exist error | please use another email prefix |
| 41035 | email prefix is invalid error | please use another email prefix |
| 41036 | avatar key is invalid error | The avatar is illegal, the key of the avatar is obtained through the upload interface |
| 41037 | avatar is sensitive error | Sensitive information |
| 41038 | gender is invalid error | please check |
| 41040 | user name is null error | please check |
| 41041 | department id is not assigned error | department id cannot be empty |
| 41042 | join time is invalid error | Join time should be a valid timestamp |
| 41043 | employee id is not valid error | The length of employee id should be between 1-64 characters and should not contain\t\n\r |
| 41044 | custom attr id is not set error | The custom field ID of the request parameter is empty |
| 41045 | Custom attribute id is not exist error | The custom attribute ID does not exist. Please confirm the source of the custom attribute ID. The custom attribute ID can be queried through the Get Enterprise Custom Attributes interface |
| 41046 | custom attr value is not set error | The value corresponding to the custom attribute is empty |
| 41047 | custom attribute href text is null error | HREF type custom attribute, link text is required |
| 41048 | custom attr href url is null error | HREF type custom attribute, url is a required field |
| 41050 | no user authority error | No user authority, check whether the user is within the address book authority |
| 41051 | user id info not provide error | please check |
| 41052 | acceptor id is invalid error | please check |
| 41053 | userID already exist error | user_id is the unique ID of the user in the enterprise and cannot be repeated |
| 41055 | employee type can not be null error | In the update interface, the employee type cannot be empty, please check |
| 41056 | no field authority error | There are unauthorised fields in the request, please check the scope of authority |
| 41057 | need send email but not set mail error | Need to send mail, but no mailbox is set |
| 41058 | need send sms but not set mobile error | Need to send SMS, but no mobile phone number is set |
| 41059 | invalid employee type error | The user’s employee type is wrong, please fill in the number between 1-5, 1 regular employee 2 intern 3 outsourcing 4 labor 5 consultant |
| 41060 | inactive employee type error | The current company does not have a corresponding employee type, and the company’s employee type consults the administrator |
| 42005 | user has exist in group error | The user already exists in the user group |
| 42006 | user has resigned error | User has left |
| 42008 | invalid tenant id error | Illegal tenant ID, please check if the tenant is available |
| 43004 | illegal unit error | Ensure that the unit information is legal and valid |
| 43005 | duplicate order error | please use another order |
| 43007 | dup dept unit custom id error | Duplicate unit_id in tenant or department_id, reused after admin |
| 43008 | custom dept id invalid error | use another department custome id |
| 43010 | big dept forbid recursion error | Too many sub-departments of the department, recursive query is not supported |
| 43011 | delete has member dept error | please retry later |
| 43013 | dept too many children error | The number of direct sub-departments of the department does not exceed 1000 |
| 45500 | internal error | Internal error, please try again later. If it still occurs then you can [consult customer service](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 47009 | duplicate name error | use another user group name |
| 50003 | invalid app_id | Incorrect parameter transmission, please make sure that app_id is correct and valid. |
| 50006 | invalid request | It is recommended to check whether the request parameters are legal |
| 55001 | server internal error | Reduce qps, wait a while then try again |
| 60001 | OpenApiErrorCodeParameterError | Check request parameters |
| 60002 | OpenApiErrorCodeDefinitionNotFound | Check that the approval code is correct |
| 60003 | OpenApiErrorCodeInstanceNotFound | Check that the approval instance ID is correct |
| 60004 | OpenApiErrorCodeUserNotFound | Check userId is correct |
| 60005 | OpenApiErrorCodeDepartmentNotFound | Check the user's department |
| 60006 | OpenApiErrorCodeValidateFormError | Check form permissions |
| 60007 | OpenApiErrorCodeSubscriptionExist | Check the subscribed list for duplicate subscriptions |
| 60008 | OpenApiErrorCodeSubscriptionNotExist | Check whether you have subscribed, you cannot cancel without subscribing |
| 60009 | OpenApiErrorCodeNoPermission | Check permissions |
| 60010 | OpenApiErrorCodeTaskNotFound | Verify that the taskId is correct |
| 60011 | OpenApiErrorCodeStartPremiumVersion | Paid approval Free users cannot initiate |
| 60012 | OpenApiErrorCodeUuidConflict | Uuid unique |
| 60013 | OpenApiErrorCodeUnsupportApporval | Free users do not support this type of approval |
| 60014 | get lock fail | concurrent update same instance error |
| 65001 | OpenApiErrorCodeInternalError | It is recommended to check the parameters and try again |
| 90201 | WrongRequestJson | Please make sure that the request body is in legal json format |
| 90202 | WrongRange | Check whether the range in the request is compliant |
| 90203 | Fail | Internal error, please try again later. If it still occurs then you can [consult customer service](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 90204 | WrongRequestBody | There is a problem with the request for entry, please check the entry carefully |
| 90205 | InvalidUsers | Check whether the user information is correct |
| 90206 | EmptySheetId | Must provide sheetid, correction request entry |
| 90207 | EmptySheetTitle | Check that the sheet name is correct |
| 90208 | SameSheetIdOrTitle | Adjust sheetId or title |
| 90209 | ExistSheetId | Retry request can be resolved |
| 90210 | ExistSheetTitle | The title of the sub-table is duplicated, please provide a non-duplicate sub-table title |
| 90211 | WrongSheetId | Check the correctness of the sheetId |
| 90212 | WrongRowOrCol | Check whether the row and column information is correct |
| 90213 | PermissionFail | Get file permissions |
| 90214 | SpreadSheetNotFound | Check spreadsheet token |
| 90215 | SheetIdNotFound | The subtable corresponding to the request entry parameter does not exist, and the request parameter is corrected |
| 90216 | EmptyValue | The request parameter is wrong or the request did not result in a data change, please check the parameter of the request |
| 90217 | TooManyRequest | The current is limited, please reduce the request running frequency |
| 90218 | LockedCell | Check if you have permission to edit the cell |
| 90219 | CellExcess | Check the number of cells |
| 90221 | TooLargeResponse | Reduce query volume |
| 90222 | TooLargeCell | Reduce cell content |
| 90223 | ColIdNotFound | Check whether the request body has set the ColId field |
| 90224 | RowIdNotFound | Check whether the request body sets the RowId field |
| 90225 | NotLinkSpreadSheet | ISV not associated |
| 90226 | ExcessLimit | Adjust the number of requests according to the returned message |
| 90227 | TooLargeRequest | Reduce request volume |
| 90228 | ImportProcessing | Waiting for import to complete |
| 90229 | WrongURLParam | URL parameter error |
| 90235 | Retry later | The server is busy, please try again later; if it still appears, you can [consult customer service for details](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 90242 | Spreadsheet in mix state | Loading document data, please wait for the document data to be loaded and try again |
| 90301 | FAILED | operation failed |
| 90302 | PARAM_ERROR | Parameter error |
| 90303 | FORBIDDEN | Permission denied |
| 90304 | META_DELETED | File deleted |
| 90305 | META_NOT_EXIST | file does not exist |
| 90306 | REVIEW_NOT_PASS | Comment content is not approved |
| 90399 | INTERNAL_ERROR | Internal error, please [consult customer service for details](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 91001 | OPEN_CODE_PARAM_ERROR | Check the parameters |
| 91002 | OPEN_CODE_FORBIDDEN | Operator should acquire permission before add permission |
| 91003 | OPEN_CODE_INVALID_OPERATION | Possible reason: document is not allowed to share externally; collaborator up to limit; etc |
| 91004 | OPEN_CODE_USER_NO_SHARE_PERM | Check the operator's share permission |
| 91005 | Resource is deleted | Resource is deleted |
| 91201 | FAILED | Process failed, retry later |
| 91202 | PARAMERR | Check the parameters |
| 91203 | NOTEXIST | Data not found, check the parameters |
| 91204 | FORBIDDEN | Check user's permission for doc and folder |
| 91205 | DELETED | Data has been deleted, check the parameters |
| 91206 | OUT_OF_LIMIT | exceed the limit |
| 91207 | DUPLICATE | Duplicate record |
| 91208 | REVIEW | Content review failed |
| 91401 | PARAMERR | Check parameter validity |
| 91402 | NOTEXIST | Check if the token is valid |
| 91403 | FORBIDDEN | Check for document read permissions |
| 91404 | LOGIN_REQUIRED | Need to log in |
| 93001 | param err | Check whether the information such as token and id is correct |
| 93002 | out of limit | Page list within 500 |
| 93003 | invalid user | Check if the user exists |
| 93004 | not found | Check that the nodeId and spaceId are correct |
| 93006 | internal err | Can directly find wiki development confirmation |
| 95001 | internal error | Internal error, please try again later. |
| 95003 | internal error | Internal error, please try again later. |
| 95005 | internal error | Internal error, please try again later. |
| 95006 | Failed | Check if the token is valid |
| 95007 | Failed | Deleted file cannot get document meta information |
| 95008 | internal error | Internal error, please try again later. |
| 95009 | Failed | Check for document read permissions. Add permissions |
| 95010 | internal error | Internal error, please try again later. |
| 95011 | internal error | Internal error, please try again later. |
| 95013 | Failed | Invalid folderToken or no permissions on directory |
| 95017 | Specific error message | Check if the revison is correct |
| 95018 | Specific error message | See specific error messages for details |
| 95020 | Specific error message | See specific error messages for details |
| 95023 | revision too old | The Revision is too old, Please use the latest version number |
| 95024 | Failed | Check parameter validity |
| 95025 | Failed | Check if the request is legal json |
| 95026 | Failed | Check batch_update interface request_type |
| 95050 | Specific error message | See specific error messages for details |
| 95051 | Specific error message | See specific error messages for details |
| 95201 | InternalError | Internal error. Please [consult customer service for details](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 95202 | InternalError | Internal error. Please [consult customer service for details](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 95203 | InternalError | Internal error, please try again later. If it still occurs then you can [consult customer service](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 95204 | InternalError | Internal error, please try again later. If it still occurs then you can [consult customer service](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 95205 | InternalError | Internal error, please try again later. If it still occurs then you can [consult customer service](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 95206 | InternalError | Internal error, please try again later. If it still occurs then you can [consult customer service](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 95207 | InternalError | Internal error, please try again later. If it still occurs then you can [consult customer service](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 95208 | InternalError | Internal error, please try again later. If it still occurs then you can [consult customer service](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 95209 | InternalError | Internal error, please try again later. If it still occurs then you can [consult customer service](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 95299 | DefaultCode | Other errors, please [consult customer service for details](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 96001 | OPEN_CODE_INTERNAL_ERROR | Retry and check the parameter. If it still occurs then you can [consult customer service](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 96002 | concurrency error,please retry | Document's permission cannot be operated concurrently, please retry |
| 96201 | LOCK | Don't call this api concurrently, please call it one by one |
| 96202 | RECOVER | Internal error, please [consult customer service for details](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 96401 | FAILED | See specific error messages for details |
| 96402 | TIMEOUT | exceed |
| 96403 | PROCESSING | Request is being processed |
| 100001 | page token invalid | Please check the token parameters |
| 100002 | Invalid field selection [illegal field] | Check the field names in fields |
| 100003 | time format must follow RFC3339 standard | Please check the time format |
| 100004 | building id invalid | Please check the Building ID parameter |
| 100005 | room id invalid | Please check the roomID parameter |
| 105001 | internal error | Internal error, please try again later. If it still occurs then you can [consult customer service](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 105002 | invalid request | Check request parameters |
| 121004 | data not exist | Please ensure that the request method, request information, request data format, etc. are correct |
| 154000 | bad request | Please make sure request method, request query, request body is correct. |
| 154001 | unauthorized | Please check authentication info, app and helpdesk should have the same tenant. |
| 190002 | invalid parameters in request | Please ensure that the parameter name and parameter value are correct, the parameters must be complete, etc. |
| 190003 | internal service error | Internal error, please try again later. If it still occurs then you can [consult customer service](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 190005 | app rate limited | Try again later, appropriately reduce the request qps |
| 191001 | invalid calendar_id | Check whether the calendar id is filled in correctly |
| 191002 | no calendar access_role | Make sure you have access permissions from the calendar, check the access role for the calendar |
| 193001 | event not found | Check whether the event id is correct and whether the event has been created |
| 193003 | event is deleted | Check if the event id is correct |
| 230003 | invalid_page_token | Invalid page token, please use to token which is get from the API |
| 230016 | the_message_was_sent_too_long | The message has been sent for more than 7days and cannot be checked. |
| 232002 | only_chat_owner_can_edit | The chat owner sets "Only the owner can edit chat information", non-chat owners cannot update the information |
| 232011 | operator_not_in_chat | The bot or authorized user is not in the chat, please add it to the chat |
| 232014 | operator_has_no_permission | Only when the user is the creator or bot is the creator and has the permission to update the information can remove members |
| 232016 | non_owner_can_only_edit_certain_parts | The bot or authorized user is not the owner of the chat and can only modify part of the information(avatar, name, description, i18n_names) |
| 232019 | The request has been rate limited. | Trigger chat current limit, please control the request speed |
| 1000001 | unexpected head in request | Request  header invalid. Please [consult customer service for details](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 1000002 | invalid parameters in request | Check whether the parameters meet API expectations |
| 1000003 | internal service error | Please [consult customer service for details](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 1000004 | method rate limited | Reduce request frequency |
| 1000005 | app rate limited | Reduce request frequency |
| 1000007 | app bot_id not found | Check if the app exists |
| 1001001 | invalid calendar_id | Check if the calendar id is correct |
| 1001002 | no calendar access_role | Check whether the corresponding calendar id is correct and whether the corresponding calendar has permission |
| 1001004 | invalid calendar type | Check if the calendar id is correct |
| 1001501 | user not found | Check if an incorrect open_id is passed |
| 1001502 | okr data not found | Check if an incorrect okr_id is passed |
| 1001901 | internal server error | Please [consult customer service for details](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 1001902 | no permission | Check if a valid token and okr_id are passed |
| 1001903 | invalid parameters | Check whether the OKR API parameter format is correct |
| 1001999 | system exception | Please [consult customer service for details](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 1002000 | acl not found | Check whether the acl_id is correct and whether it has been created correctly |
| 1003001 | event not found | Check if the event id is correct |
| 1003003 | event is deleted | Check if the event id is correct |
| 1004000 | attendee not found | Check whether the attendee id is correct |
| 1004004 | invalid attendee type | Check whether the attendee type is correct |
| 1050001 | TIME_CHECK_NOT_VALID | Fill in the accurate time range, the starting time should not be earlier than half a year ago, and the time selection range should not exceed one month |
| 1050002 | ErrCode_DATABASE_ERR | Please try again later |
| 1050004 | Error_Param_Error | Check whether the parameters are filled in incorrectly or missing |
| 1050005 | Error_Page_Size_Invalid | The page_size parameter of the paging request is illegal, please check the size set by page_size |
| 1050006 | Error_Page_Token_Invalid | Check whether the page token is filled in incorrectly |
| 1050007 | Error_Event_Name_Not_Found | Check whether the event name is filled in incorrectly |
| 1050008 | Error_Open_Platform_RPC | Please try again later |
| 1050009 | Error_Lark_ID_Not_Found | The provided user ID does not exist, please check the source of the user ID and check it is a user under the current tenant |
| 1050011 | Error_Event_Module_Invalid | Check whether the event module is filled in incorrectly |
| 1050012 | Error_Event_Module_Not_Match_Event_Name | Modify the event type or event module to ensure that the two are consistent |
| 1050019 | Error_Not_Support_Query | Check whether unsupported query fields are added |
| 1051002 | param validate error | The time range cannot exceed the limit. Please refer to the document for the specific limit, and the time format must be correct.The page_size and page_token should also be filled in correctly" |
| 1061041 | parent node has been deleted | Please confirm whether the node of parent_token is deleted |
| 1069301 | fail | Please try again later. If it still occurs then you can [consult customer service](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 1069302 | param error | Check parameter validity |
| 1069303 | forbidden | Check whether the document to be commented on has comment permissions (the default for common documents is to have reading permissions and comment permissions) |
| 1069304 | docs had been deleted | Check whether the document to be commented has been deleted |
| 1069305 | docs not exist | Check whether the document to be commented can be accessed normally |
| 1069306 | content review not pass | Check comments for illegal content |
| 1069307 | not exist | Check whether the document to be commented can be accessed normally, the comment may not exist / deleted |
| 1069308 | exceeded limit | The comment data exceeds the upper limit. Please [consult customer service for details](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 1069399 | internal error | Please try again later. If it still occurs then you can [consult customer service](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 1103003 | tenant no idp config | Check the SSO related configuration and turn on the tenant SSO switch after confirming that it is correct |
| 1103008 | idp user not bind idp cp | Call the binding interface to bind enterprise-level federated login credentials |
| 99991201 | resource not find | 404 resource not find. Check if the path is filled in correctly |
| 99991300 | invalid next parameter:%s | There is a problem with the next parameter in the callback URL, please check and carry the correct next parameter |
| 99991301 | request method doesn’t match | Check whether the request method (POST/GET/...) set by the interface is consistent with the one used in the request |
| 99991400 | request trigger frequency limit | Requests are too frequent, please reduce the frequency of requests |
| 99991401 | ip %s is denied by app setting | IP is restricted by the whitelist. Please check whether the application has enabled the IP whitelist. Once enabled, only the sources in the whitelist can request normal calls. You can get your server export IP in the following ways: 1. Please consult the operation and maintenance personnel of the IT department in the enterprise or the cloud service provider; 2. Log in to your server and execute with the command line: ```curl ifconfig.me``` or ```curl cip.cc``` or ```curl myip.ipip.net``` or ```curl ipinfo.io```. After obtaining the correct server export IP, you can refer to this document to modify the IP whitelist settings of your application |
| 99991501 | auth login invalid code | No login code, please log in again. If it still occurs then you can [consult customer service](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 99991503 | invalid next parameter: invalid host | There is a problem with the next parameter in the callback url, and the host of next and callback are different |
| 99991543 | login_user invalid | app_id or app_secret does not exist, please check the parameters |
| 99991611 | get session fail | Failed to get the session, please check whether the request carries the corresponding cookie (lobsession) |
| 99991612 | session invlaid | The gadget request authentication failed, and the requested session is invalid |
| 99991621 | get session fail | Failed to get the session, please check whether the request carries the corresponding cookie (lobsession) |
| 99991622 | session invlaid | Open platform SSO request authentication failed, the requested session is invalid |
| 99991631 | get session fail | Failed to get the session, please check whether the request carries the corresponding cookie (lobsession) |
| 99991632 | session invlaid | Lark SSO request session verification failed |
| 99991641 | auth open invalid session | The current login status is invalid, please log in again to use |
| 99991642 | auth open invalid session mina | The current login status has expired, please log in to use |
| 99991643 | auth open invalid session sso | Open id or refresh token input, please log in again. If it still occurs then you can [consult customer service](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 99991644 | auth open invalid config | Internal error, please [consult customer service for details](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 99991645 | auth open invalid user | The currently logged-in user is invalid, please log in again. If it still occurs then you can [consult customer service](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 99991661 | Need a token | The request needs to use token authentication; please check whether the correct authorization is filled in the request header, and be careful not to omit "Bearer" before the token value. For details, please refer to API Access Token |
| 99991662 | app not in use | Check if the app is disabled |
| 99991663 | tenant token invalid | Enter tenant_access_token to update the token. For details, please refer to API Access Token |
| 99991663 | tenant access token not support | The current request does not support tenant_access_token. Please check and try again. |
| 99991664 | invalid app token | Illegal app access token |
| 99991665 | invalid tenant code | Illegal tenant access token |
| 99991668 | token invalid | Enter user_access_token to update the token. For details, please refer to API Access Token |
| 99991668 | user access token not support | The current request does not support user_access_token, please check and try again |
| 99991669 | invalid user refresh token | Illegal user refresh token |
| 99991670 | invalid sso token | Illegal SSO token |
| 99991671 | Invalid token: must start with t-/u- | The token format used in the current request is incorrect. Please check and try again. For details, please refer to API Access Token |
| 99991672 | No permission | The current request does not apply for relevant permissions. The opening of the permission needs to be submitted and released, and it will take effect after being reviewed by the administrator. For details, see How to open API scope. For applications in the development and testing phase, you can use the "Beta Application" function. You do not need to directly release the version to apply for permission. It can be completed in the test administrator. For details, see Tutorial. |
| 99991673 | unauthorized app | The application status is not available, please check the current application installation status |
| 99991674 | user type not support | The user type of the current request is invalid, please check and try again |
| 99991675 | parse user error | The user_id format used in the current request is incorrect, please check and try again |
| 99991676 | token no permission | The token has no corresponding permissions, please check the permissions of the token |
| 99991677 | token expire | The token has expired, please token. For details, please refer to Update access_token |
| 99991681 | auth fail | Third-party authentication returns an error, please refer to the message for troubleshooting |
| 99991691 | get session fail | Authentication failed, request without cookie |
| 99991692 | auth fail | Authentication failed, please refer to message for troubleshooting |
| 99991693 | session not exist | Authentication failed, current session is invalid |
| 99992351 | these open ids not existed: %v | Some open_id does not exist, please check and try again |
| 99992352 | these lark ids not existed: %v | Some open_id does not exist, please check and try again |
| 99992353 | these lark ids not existed: %v | Some open_id does not exist, please check and try again |
| 99992354 | these ids not existed: %v | The id indicated in the message does not exist, please check and try again |
| 99992355 | these lark ids not existed: %v | The current chat_id does not exist, please check and try again |
| 99992356 | these open_chat ids not existed: %v | The current chat_id does not exist, please check and try again |
| 99992357 | these open_department ids not existed: %v | The current open_department_id does not exist, please check and try again |
| 99992358 | these ids not existed: %v | The current open_department_id does not exist, please check and try again |
| 99992359 | these user ids not existed: %v | The current user_id does not exist, please check and try again |
| 99992360 | these ids not existed: %v | The current user_id does not exist, please check and try again |
| 99992361 | open_id cross app | open_id does not belong to the current application, please check and try again |
| 99992363 | these ids not existed: %v | The current union_id does not exist, please check and try again |
| 99992364 | these ids not existed: %v | The current union_id does not exist, please check and try again |
| 99992364 | user id cross tenant | The user_id under other wells cannot be used, please check and try again |
| 99992370 | these ids not existed: %v | The current open_app_version_id does not exist, please check and try again |
| 99992375 | these open_room ids not existed: %v | Currently open_room_id does not exist, please check and try again |
| 99992376 | these ids not existed: %v | Currently open_room_id does not exist, please check and try again |
| 99992378 | these open_room ids not existed: %v | The current open_app_version_id does not exist, please check and try again |
| 99992379 | these ids not existed: %v | The current open_department_id does not exist, please check and try again |
| 99992380 | these ids not existed: %v | The current open_department_id does not exist, please check and try again |
| 99992381 | union_id cross tenant | union_id does not belong to the current tenant. When you use union_id, you need to pay attention to whether union_id belongs to the current tenant, please check and try again |
| 99992402 | Specific error message | If a request parameter is missing or incorrect, see the error message returned for the request for more information on the error |
| 99992500 | these ids not existed: %v | The current tenant_key does not exist, please check and try again |
| 99992501 | these ids not existed: %v | The current tenant_key does not exist, please check and try again |
| 99993201 | api don't support batch | Please confirm whether the batch API document contains the API |
| 99993202 | the id of subrequest can't be empty | Please refer to the example of the document to add the id of the subrequest |
| 99993203 | the id of subrequest is duplicated | Please confirm whether the id of the subrequest is repeated |
