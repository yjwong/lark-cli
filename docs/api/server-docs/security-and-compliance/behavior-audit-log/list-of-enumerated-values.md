---
title: "List of enumerated values"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uQjM5YjL0ITO24CNykjN/audit_log/appendix"
section: "security_and_compliance"
updated: "1737707180000"
---

# List of enumerated values

## **Event module**

| **Optional value**         | **Description**           |
| --------- | --------------- |
|1 | Docs |
|2 | Messenger |
|3 | Contacts |
|4 | Workplace |
|5 | Audio or video meetings|
|6 | Account Activity |
|8| Device |
|9| Moments |
|13| OKR |
|16|Email |
|18|Approve |
|21|Calendar |
|22|URL |
## **Event name**

| **Event name**         | **Event description**           | **Event module**|
| --------- | --------------- | ------|
|space_create_doc | Creates Docs | 1|
|space_read_doc | Reads Docs | 1|
|space_delete_doc | Deletes Docs | 1|
|space_edit_doc | Edits Docs | 1|
|space_comment_doc | Comments Docs |1| 
|space_update_collaborator_doc | Updates Docs collaborator (including internal/external collaborator) | 1|
|space_update_share_setting_doc | Updates Docs sharing settings |1| 
|space_export_doc | Exports Docs | 1|
|space_download_file | Downloads files | 1|
|space_import_doc | Imports Docs | 1|
|space_share_to_3rdApp | Shares to third-party apps |1|
|space_front_export_csv | Exports spreadsheet to CSV| 1|
|space_front_export_image | Exports and downloads images from mobile |1| 
|space_open_with_3rdApp | Opens preview files with third-party apps | 1| 
|space_print_doc | Prints Docs | 1| 
|space_copy_content | Copies contents | 1| 
| set_doc_sec_label | Set secure label for Docs | 1| 
| space_add_comment | Added comment | 1| 
| space_add_reply | Added reply | 1| 
| space_create_copy | Made a copy | 1| 
| space_create_spacecopy | Create workspace copy | 1| 
| space_delete_doccurrent | Delete files/folders (subfiles/subfolders not affected) | 1| 
| space_delete_reply | Deleted reply | 1| 
| space_delete_retlabel | Cleared retention label | 1| 
| space_demoslides | Present slides | 1| 
| space_download_history | Downloaded earlier versions | 1| 
| space_edit_reply | Edited reply | 1| 
| space_edit_retlabel | Edited retention label | 1| 
| space_edit_space_member | Edited workspace members | 1| 
| space_edit_space_member_auth | Edited permissions for workspace members | 1| 
| space_edit_space_owner | Edited workspace administrators | 1| 
| space_finish_comment | Resolved comment | 1| 
| space_modfiy_security_setting | Edited security settings for the workspace | 1| 
| space_move | Moved to folder/Wiki | 1| 
| space_move_owner | Transferred ownership | 1| 
| space_recover_history | Restored earlier versions | 1| 
| space_reopen_comment | Restored comment | 1| 
| space_save_module | Saved to My Templates | 1| 
| space_set_retlabel | Set retention label | 1| 
| space_share | Shared file/folder | 1| 
| space_share_content | Shared file contents | 1| 
| space_update_collaborator_auth | Edited collaborator permissions | 1| 
| space_update_spaceshare_status | Enabled or disabled space link sharing | 1| 
| space_view_history | Viewed earlier versions | 1| 
| turn_down_doc_sec_label | Lowered sensitivity level | 1| 
| turn_up_doc_sec_label | Increased sensitivity level | 1| 
| space_authapprove | Request Docs permissions | 1| 
|im_download_image | Downloads images | 2| 
|im_download_video | Downloads videos | 2| 
|im_open_link | Opens links in chats | 2|
|im_start_external_chat | Starts private chat with a stranger | 2|
|im_load_file_to_local | Loads server-side files to local (triggered by the first click to download, which is equivalent to caching to local) | 2|
|im_forward_file | Forwards files | 2|
|im_create_chat | Creates group chat | 2|
|im_delete_chat | Disbinds group chat | 2|
|im_join_chat | Joins group chat | 2|
|im_quit_chat | Quits group chat | 2|
| im_add_chatadmin | Added group administrator | 2| 
| im_addtochat | Added group member | 2| 
| im_admin_no_restrict_ctrl | Don't apply restricted mode of group chat to | 2| 
| im_chat_editimage | Edited image | 2| 
| im_chat_previewfile | Previewed non-image or video file online | 2| 
| im_chat_uploadfile | Uploaded file | 2| 
| im_chat_withdraw | Recalled | 2| 
| im_delete_chatadmin | Deleted group administrator | 2| 
| im_deletefromchat | Deleted group member | 2| 
| im_download | Downloaded file | 2| 
| im_forward_chatadmin | Set group owner | 2| 
| im_ocr | Extract text from screenshots | 2| 
| im_savetospace | Saved to My Space | 2| 
| im_screencap | Record screen | 2| 
| im_send_link | Sent URL in chat | 2| 
| im_snaphot | Take screenshots | 2| 
|im_copy_image | Copies image | 2|
|im_read_doc | Reads Docs in chat window | 2|
|im_download_doc | Downloads Docs in chat window | 2|
|im_open_with_3rdApp | Opens preview files in chat window with a third-party app | 2|
|im_chat_pin_show_pin_list      | Expand pinned area                                              | 2|
| im_chat_pin_show_in_chat       | View pinned items                                               | 2|
| im_chat_pin_stick              | Prioritize a pinned item                                        | 2|
| im_chat_pin_update_permission  | Change manage permissions for pinned area                       | 2|
| im_chat_pin_update             | Edit the name of the pinned link                                | 2|
| im_chat_pin_sidebar_search     | Click "Search" in the toolbar                                   | 2|
| im_chat_pin_click_open_browser | Click "Open in Browser" on the dropdown menu of the pinned link | 2|
| im_chat_pin_sidebar_file       | Click "Files" in the toolbar                                    | 2|
| im_chat_pin_unstick            | Cancel prioritize a pinned item                                 | 2|
| im_chat_pin_hover_show         | Hover over pinned card to preview                               | 2|
| im_chat_pin_reorder            | Drag to reorder the pinned area                                 | 2|
| im_chat_pin_create             | Pin an item                                                     | 2|
| im_chat_pin_click_copy_url     | Click "Copy Link" on the dropdown menu of the pinned link       | 2|
| im_chat_pin_sidebar_add_link   | Click "Add Pinned Link" in the toolbar                          | 2|
| im_chat_pin_click_back_to_chat | Click "View in Chat"                                            | 2|
| im_export_chat_chatter         | Export group member data                                        | 2|
| im_chat_pin_sidebar_announce   | Click "Add/View Group Announcement" in the toolbar              | 2|
| im_chat_pin_click_open_url     | Open pinned link                                                | 2|
| im_chat_pin_delete             | Remove a pinned item | 2|
|contact_add_external_friend | Adds external contacts as friends | 3|
|contact_obtain_mobile | Obtains mobile number | 3|
|contact_view_department_struct | Views department structure | 3|
|contact_view_personal_info | Views profile page | 3|
|contact_personal_avatar_update | Update profile photo | 3|
|workplace_open_app | Opens app | 4|
|workplace_app_download_doc | Downloads Docs in Workplace app | 4|
|workplace_open_with_3rdApp | Opens preview files in Workplace app with a third-party app | 4|
|vc_create_video_meeting | Creates audio/video calls |  5|
|vc_finish_video_meeting | Ends audio/video calls | 5|
|vc_quit_video_meeting | Quits audio/video calls | 5|
|vc_join_video_meeting | Joins audio/video calls | 5|
| vc_changerange | Change join meeting permission | 5| 
| vc_invite_video_meeting | Invited other users into meeting | 5| 
| vc_share_video_meeting | Shared meeting | 5| 
| vc_sharescreen_video_meeting | Share screen/document | 5| 
| vc_createclip | Create clips | 5| 
| vc_deleteclip | Delete clips | 5| 
| vc_deletemeta | Move a Minutes file to Trash | 5| 
| vc_deletemetafromtrash | Delete a Minutes file permanently | 5| 
| vc_drivebatchtranscribe | Generate Minutes by importing from Lark Docs | 5| 
| vc_exportmeta | Export Minutes files | 5| 
| vc_revertmeta | Restore a Minutes file from Trash | 5| 
| vc_setaudiostatus | Generate Minutes by recording on mobile | 5| 
| vc_share_to_room_video_meeting | share screen | 5| 
| vc_sharebychat | Share Minutes to chat | 5| 
| vc_sharebylink | Edit Minutes permissions via link | 5| 
| vc_shareclipbychat | Share clips to chat | 5| 
| vc_sharecollaborator | Manage Minutes' collaborators | 5| 
| vc_uploadbymeeting | Generate Minutes by recording Lark Meetings | 5| 
| vc_uploadmeta | Generate Minutes by uploading local files | 5| 
| vc_viewclip | Access a clip | 5| 
| vc_viewmeta | Access a Minutes file | 5| 
| vc_start_transcript         | Start transcribing                   |5| 
| vc_stop_transcript          | Stop transcribing                    |5| 
| vc_request_transcript       | Request transcribing                 |5| 
| vc_access_transcript        | Accept transcribing request          |5| 
| vc_refuse_transcript        | Reject transcribing request          |5| 
| vc_access_local_record      | Accept local recording request       |5| 
| vc_refuse_local_record      | Reject local recording request       |5| 
| vc_refuse_record            | Reject recording request             |5| 
| vc_access_record            | Accept recording request             |5| 
| vc_stop_record              | Stop recording                       |5| 
| vc_request_record           | Request recording                    |5| 
| vc_start_local_record       | Start local recording                |5| 
| vc_stop_local_record        | Stop local recording                 |5| 
| vc_request_local_record     | Request local recording              |5| 
| vc_manage_stop_local_record | Stop local recording by participants |5| 
| vc_start_record             | Start recording                      |5| 
| vc_start_screen_control     | Turn on remote control               |5| 
| vc_finish_screen_control    | Turn off remote control|5| 
|account_login | Logs in | 6|
| account_passport_changeaccount | Switched account | 6| 
| account_passport_logout | Logged out | 6| 
| account_passport_renew_logoncre | Updated login credentials | 6| 
| account_passport_renew_password | Updated login password | 6| 
| user_active_exit_team | Left organization | 6| 
| account_passport_update_otp     |Update OTP settings| 6| 
| account_passport_update_2fa    |Update two-step verification method| 6| 
| account_passport_updatesparecre |Update backup verification method| 6| 
| account_passport_revoke_appauth |Manage app authorization| 6| 
| account_passport_renew_fidocre|Update passkey settings| 6| 
| device_screen_recording | Screen record | 8| 
| device_screenshot | Screenshot | 8| 
| moments_copy_comment | Copy comment | 9| 
| moments_copy_image | Copy link | 9| 
| moments_copy_link | Copy moment | 9| 
| moments_copy_post | Post comment | 9| 
| moments_create_comment | Create official account | 9| 
| moments_create_official | Post moment | 9| 
| moments_create_post | Emoji response | 9| 
| moments_create_reaction | Delete comments | 9| 
| moments_delete_comment | Delete posts | 9| 
| moments_delete_post | Cancel emoji response | 9| 
| moments_delete_reaction | Edit official account | 9| 
| moments_download_image | Forwarded moment | 9| 
| moments_edit_official | Login to official account | 9| 
| moments_enter_post_detail | View new moments | 9| 
| moments_forward_post | Copy image | 9| 
| moments_login_official | Download image | 9| 
| moments_preview_image | Preview image | 9| 
| moments_preview_video | View moments | 9| 
| moments_view_feed_post | Enter moment's detail page | 9| 
| moments_viewmoment | Preview video | 9| 
| email_batchexport | Email batchexport | 16| 
| email_copy | Copy email content | 16| 
| email_delete | Delete email permanently | 16| 
| email_downloadfile | Download attachment | 16| 
| email_downloadmail | Download email | 16| 
| email_editforward | Set auto forward | 16| 
| email_foawardasatt | Forward email as attachment | 16| 
| email_localopen | Open attachment with local application | 16| 
| email_login | Log in to public mailbox | 16| 
| email_loginthird | Log in to third party email client | 16| 
| email_move | Move email | 16| 
| email_print | Print email | 16| 
| email_read | Read email | 16| 
| email_reply | Reply/Forward email | 16| 
| email_send | Send email | 16| 
| email_sharefile | Share attachment to chat | 16| 
| email_sharemail | Share email to chat | 16| 
| email_urlclick | Click links attached in email | 16| 
| officeapprove_operate_clear | Delete request | 18| 
| officeapprove_add_submanager | Add sub-administrator | 18| 
| officeapprove_process_intervene | Intervene in approval process | 18| 
| officeapprove_hand_over | Hand over approval | 18| 
| officeapprove_operate_save | Save request | 18| 
| officeapprove_bitable_add_data | Add approval data | 18| 
| officeapprove_admin_export_data | Export data from Admin Console | 18| 
| officeapprove_create_approval_group | Create approval process group | 18| 
| officeapprove_enable_approval | Activate approval process | 18| 
| officeapprove_update_approval_group | Update approval process group | 18| 
| officeapprove_update_approval | Edit approval process | 18| 
| officeapprove_delete_approval | Delete approval process | 18| 
| officeapprove_delete_approval_group | Delete approval process group | 18| 
| officeapprove_disable_approval | Deactivate approval process | 18| 
| officeapprove_export_data | Export data from app | 18| 
| officeapprove_create_approval | Create approval process | 18| 
| officeapprove_export_efficiency_task_detail | Export task details | 18| 
| officeapprove_process_diagnosis | View process diagnosis | 18| 
| officeapprove_export_efficiency_process_detail | Export process details | 18| 
| officeapprove_efficiency | View personal efficiency | 18| 
| officeapprove_share_approval | Share access to the approval submission page | 18| 
| officeapprove_operate_submit_again | Resubmit request | 18| 
| officeapprove_operate_recall | Recall request | 18| 
| officeapprove_operate_upload_attachment | Upload attachment | 18| 
| officeapprove_operate_upload_image | Upload image | 18| 
| officeapprove_operate_share | Share approval request | 18| 
| officeapprove_operate_urgent | Remind approver | 18| 
| officeapprove_operate_discuss | Start group chat | 18| 
| officeapprove_operate_addsign | Add approver | 18| 
| officeapprove_operate_rollback | Return request | 18| 
| officeapprove_operate_removesign | Remove approver | 18| 
| officeapprove_operate_comment | Leave comment | 18| 
| officeapprove_operate_cc | Cc to another approver | 18| 
| officeapprove_operate_reject | Reject request | 18| 
| officeapprove_operate_forward | Transfer request | 18| 
| officeapprove_operate_approve | Approve request | 18| 
| officeapprove_modify_instance | Edit approval request | 18| 
| officeapprove_create_instance | Submit approval request | 18| 
| officeapprove_detail_instance | View approval request | 18| 
| cal_create_event               | Create event                              |21| 
| cal_add_attachment             | Add event attachment                      |21| 
| cal_remove_attachment          | Remove event attachment                   |21| 
| cal_delete_event               | Delete event                              |21| 
| cal_add_remove_participants    | Add/Remove event guest                    |21| 
| cal_add_remove_meeting_room    | Add/Remove meeting room                   |21| 
| cal_mod_event_permission       | Change event guest's permissions          |21| 
| cal_create_meeting_group       | Create meeting group                      |21| 
| cal_transfer_event             | Transfer event                            |21| 
| cal_share_event                | Share event to chat                       |21| 
| cal_join_event                 | Join event                                |21| 
| calendar_create_scheduler      | Create schedule                           |21| 
| calendar_edit_scheduler        | Edit schedule                             |21| 
| calendar_del_scheduler         | Delete schedule                           |21| 
| cal_delete_calendar            | Create calendar                           |21| 
| cal_create_calendar            | Delete calendar                           |21| 
| cal_change_def_cal_permissions | Change calendar default permissions       |21| 
| cal_add_remove_calendar_sharee | Share/Stop sharing calendar with user     |21| 
| cal_modify_sharee_permissions  | Change users' calendar access permissions |21| 
| cal_add_3party_account        | Add third-party calendar account          |21| 
| cal_subscribe_calendar         | Subscribe to calendar                     |21| 
| cal_unsubscribe_calendar       | Unsubscribe from calendar                 |21| 
| cal_share_calendar_to_session  | Share calendar to chat                    |21| 
| cal_remove_3party_account     | Remove third-party calendar account|21| ## **Web context information**

| **Context field**         | **Description**          |
| --------- | --------------- | 
|user_agent | UA information | 
|IP | Public IP | ## **Android context information**
| **Context field**         | **Description**          |
| --------- | --------------- | 
|udid | UDID | 
|did | Device ID | 
|app_ver | App version | 
|ver | SecSDK version | 
|region | Device language | 
|id_i | Android version number | 
|id_r | Android version | 
|hw_brand | Brand | 
|hw_manuf | Manufacturer | 
|wifip | wifi ip | 
|route_iip | Route IP | 
|route_iip | Route gateway IP | 
|env_su | Root or not | 
|env_tz | Time zone of mobile system | 
|env_tz | Language of mobile system | 
|IP | Public IP | ## **IOS context information**
| **Context field**        | **Description**          |
| --------- | --------------- | 
| udid | UDID | 
|did | Device ID | 
|app_ver | App version | 
|ver | SecSDK version |
| os | System type and version | 
| STZone | System timezone |
| ML | Current language | 
| sjd | Root or not |
| proxyip | Proxy IP | 
| wifip | wif ip |
|IP | Public IP | ## **PC context information**
| **Context field**         | **Description**          |
| --------- | --------------- | 
| udid | UDID | 
|did | Device ID | 
|app_ver | App version | 
|ver | SecSDK version | 
| os | Client type | 
| wifip | wifi ip | 
|IP |Public IP | ## **Details of event extension fields**

### Event extension fields
**Field**                                    | **Event name**                                                                                                                                                                                                                                                                                | **Description**                            | **Value**                                                                                                                                                                                                                                                    |
| ---------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ApplyPermissionRole                      | space_authapprove                                                                                                                                                                                                                                                                         | Apply Permission                       | CollaboratorRole_CanView; CollaboratorRole_CanEdit; CollaboratorRole_FullAccess;                                                                                                                                                                         |
| ccm_authuser_auth_after                  | space_update_collaborator_auth                                                                                                                                                                                                                                                            | Permission after                       | CollaboratorRole_FullAccess; CollaboratorRole_CanView; CollaboratorRole_CanEdit;                                                                                                                                                                         |
| ccm_download_type                        | space_download_file space_download_history                                                                                                                                                                                                                                                | Download type                          | image;video;file;                                                                                                                                                                                                                                         |
| ccm_edit_part                            | space_edit_doc                                                                                                                                                                                                                                                                            | Type of edited content                 | 1: title; 0: text;                                                                                                                                                                                                                                        |
| ccm_share_type                           | space_share                                                                                                                                                                                                                                                                               | Share via                              | share_wecha;share_lark;share_encrypted_link;copy_link;share_qrcode;                                                                                                                                                                                      |
| ccm_space_default_permission             | space_edit_space_member_auth                                                                                                                                                                                                                                                              | Member permission after                | WikiNodeRole_VIEW; WikiNodeRole_EDIT;                                                                                                                                                                                                                    |
| CCMPermissionSettingType                 | space_update_share_setting_doc                                                                                                                                                                                                                                                            | Settings modified                      | see **Cloud document permission setting parameter description**                                                                                                                                                                                             |
| CCMPermissionSettingType_space           | space_modfiy_security_setting                                                                                                                                                                                                                                                             | Settings modified                      | see **Cloud space permission setting parameter description**                                                                                                                                                                                                 |
| CCMPermissionSettingTypeAftervalue       | space_update_share_setting_doc                                                                                                                                                                                                                                                            | Value after                            | see **Cloud document permission setting parameter description**                                                                                                                                                                                              |
| CCMPermissionSettingTypeAftervalue_space | space_modfiy_security_setting                                                                                                                                                                                                                                                             | Value after                            | see **Cloud space permission setting parameter description**                                                                                                                                                                                                 |
| ChatType                                 | im_chat_uploadfile im_chat_withdraw im_copy_image im_forward_file im_open_link im_savetospace email_sharefile im_chat_editimage im_chat_previewfile                                                                                                                                       | Chat type                              | P2P: private chat; GROUP: group chat;                                                                                                                                                                                                                     |
| exportType                               | vc_exportmeta                                                                                                                                                                                                                                                                             | Export file type                       | LarkDOC; TXT; SRT;                                                                                                                                                                                                                                     |
| FileType                                 | im_chat_uploadfile im_download im_chat_previewfile                                                                                                                                                                                                                                        | File type                              | Such as MEDIA. Use the FileName suffix to determine the file type more accurately.                                                                                                                                                                       |
| im_chat_restrict_whitelist_level         | im_admin_no_restrict_ctrl                                                                                                                                                                                                                                                                 | Restricted membership                  | unkown; no member; only admin and owne; white list;                                                                                                                                                                                                      |
| im_message_type                          | im_forward_file                                                                                                                                                                                                                                                                           | Type of messages forwarded             | See **Message type**                                                                                                                                                                                                                                         |
| im_withdraw_messagetype                  | im_chat_withdraw                                                                                                                                                                                                                                                                          | Type of messages recalled              | See **Message type**                                                                                                                                                                                                                                         |
| LabelType                                | space_delete_retlabel space_edit_retlabel space_set_retlabel                                                                                                                                                                                                                              | Label type                             | reserve; delete; reserve forever;                                                                                                                                                                                                                        |
| MessageType                              | im_copy_image                                                                                                                                                                                                                                                                             | Message type                           | See **Message type**                                                                                                                                                                                                                                         |
| MetaType                                 | vc_createclip vc_deleteclip vc_deletemeta vc_deletemetafromtrash vc_download_minutes vc_drivebatchtranscribe vc_exportmeta vc_revertmeta vc_setaudiostatus vc_sharebychat vc_sharebylink vc_shareclipbychat vc_sharecollaborator vc_uploadbymeeting vc_uploadmeta vc_viewclip vc_viewmeta | Minutes type                           | MediaType_Video; MediaType_Audio; MediaType_Text;                                                                                                                                                                                                        |
| OperateType                              | space_update_spaceshare_status                                                                                                                                                                                                                                                            | Publish the workplace to the internet? | off; on;                                                                                                                                                                                                                                                  |
| OpResult                                 | im_copy_image vc_sharescreen_video_meeting im_chat_editimage im_chat_previewfile im_download_image im_ocr vc_share_to_room_video_meeting                                                                                                                                                  | Action completed?                      | 0: success; 1: failure;                                                                                                                                                                                                                                   |
| roomsharetype                            | vc_share_to_room_video_meeting                                                                                                                                                                                                                                                            | Screen-sharing mode                    | roomsharetypevwireless; roomsharetypevwired;                                                                                                                                                                                                             |
| shareAuth                                | vc_sharebylink                                                                                                                                                                                                                                                                            | Auth                                   | PermLinkShareEntity_TenantReadable; PermLinkShareEntity_TenantEditable; PermLinkShareEntity_AnyoneReadable; PermLinkShareEntity_AnyoneEditable; PermLinkShareEntity_Off;                                                                                 |
| shareAuth                                | vc_sharecollaborator                                                                                                                                                                                                                                                                      | Permission granted/revoked             | PermissionCode_VIEW; PermissionCode_EDIT; Delete;                                                                                                                                                                                                        |
| shareAuth                                | vc_sharebychat                                                                                                                                                                                                                                                                            | Auth                                   | AuthViewSuccess; AuthEditSuccess; AuthNotAllowExternal; AuthNoShareAccess; AuthAlreadyHasPerm; AuthFailed;                                                                                                                                               |
| sharetype                                | vc_sharescreen_video_meeting                                                                                                                                                                                                                                                              | Sharing type                           | sharetypevsharetypescreen; sharetypevsharetypefollow; sharetypevsharetypewhiteboard;                                                                                                                                                                     |
| space_share_to_3rdApp                   | space_share_to_3rdApp                                                                                                                                                                                                                                                                    | Share via                              | share_wecha;share_lark;share_encrypted_link;copy_link;share_qrcode;                                                                                                                                                                                      |
| SwitchTenantType                         | account_login account_passport_changeaccount                                                                                                                                                                                                                                              | Tenant switching mode                  | Initiative 0: The user actively triggers the tenant cut; Passive 1: Other scenarios. It may be that session failure causes passive tenant switching; Login process 2: Triggered during the login process, generally refers to the login of the peer tenant; |
| Type                                     | vc_createclip vc_deleteclip vc_deletemeta vc_deletemetafromtrash vc_download_minutes vc_drivebatchtranscribe vc_exportmeta vc_revertmeta vc_setaudiostatus vc_sharebychat vc_sharebylink vc_shareclipbychat vc_sharecollaborator vc_uploadbymeeting vc_uploadmeta vc_viewclip vc_viewmeta | Created by                             | ObjectType_Normal: record meeting; ObjectType_Upload：upload local files; ObjectType_Upload_CCM：import from lark docs; ObjectType_RealTime_Recording: record on mobile; ObjectType_Clip：create clip ObjectType_Live: live;                                |
| UpdateType                               | account_passport_renew_logoncre account_passport_renew_password                                                                                                                                                                                                                           | Update type                            | Add: 0; Modify: 1; Delete: 2;
### Message type 
**Field**                    | **Description**                                                                                  |
| ---------------------- | -------------------------------------------------------------------------------------- |
| POST                   | It may contain multiple types, such as text, images, etc. and is considered rich text. |
| FILE                   | Files                                                                                  |
| TEXT                   | Text                                                                                   |
| IMAGE                  | Images                                                                                 |
| SYSTEM                 | System messages                                                                        |
| AUDIO                  | Audio                                                                                  |
| SHARE_GROUP_CHAT       | Group business cards                                                                   |
| STICKER                | Custom emojis                                                                          |
| MERGE_FORWARD          | Messages forwarded after merging                                                       |
| CALENDAR               | Calendar                                                                               |
| CLOUD_FILE             | Cloud files                                                                            |
| CARD                   | Card messages and emails                                                               |
| MEDIA                  | Videos                                                                                 |
| SHARE_CALENDAR_EVENT   | Calendar invitations                                                                   |
| HONGBAO                | Red envelopes                                                                          |
| GENERAL_CALENDER       | Assignment transfers                                                                   |
| VIDEO_CHAT             | Video conference invitations                                                           |
| LOCATION               | Location                                                                               |
| CUSTOMIZE              | Customized                                                                             |
| COMMERCIALIZED_HONGBAO | Commercial red envelopes                                                               |
| SHARE_USER_CARD        | Personal business cards                                                                |
| TODO                   | Tasks                                                                                  |
| FOLDER                 | Folders                                                                                |
| DIAGNOSE               | Diagnostic messages                                                                    |
| VOTE                   | Vote messages

### Cloud document permission setting parameter description
**Field**                          | **Parameter**                          | **Description**                                                                                        |
| ---------------------------------- | -------------------------------------- | ------------------------------------------------------------------------------------------------------ |
| CCMPermissionSettingType           | PermissionLinkShareType                | Link Sharing Settings                                                                                  |
| CCMPermissionSettingType           | PermissionSinglePageSettingType        | Single-page Link Sharing Settings                                                                      |
| CCMPermissionSettingType           | PermissionExternalAccessType           | External Access Settings                                                                               |
| CCMPermissionSettingType           | CommentEntityType                      | Comment Settings                                                                                       |
| CCMPermissionSettingType           | SecurityEntityType                     | Security Settings                                                                                      |
| CCMPermissionSettingType           | ShareEntityType                        | Add Collaborator Settings（Organization)                                                                |
| CCMPermissionSettingType           | ManageCollaboratorEntityType           | Add Collaborator Settings（Permission)                                                                  |
| CCMPermissionSettingType           | ShareExternalEntityType                | Share to External or Partner Tenant                                                                    |
| CCMPermissionSettingType           | PermissionSinglePageExternalAccessType | Single Page External Access Settings                                                                   |
| CCMPermissionSettingType           | ApplyEmbedEntityType                   | Quickly access permission-restricted reference documents and apply for permission setting              |
| CCMPermissionSettingType           | ShowCollaboratorInfoEntityType         | View document collaborative avatar and like avatar setting                                             |
| CCMPermissionSettingType           | LeaderCopySwitchType                   | Supervisor Copy Setting                                                                                |
| CCMPermissionSettingType           | LeaderLinkShareEntityType              | Supervisor Access Setting by link                                                                      |
| CCMPermissionSettingTypeAftervalue | DisableLinkShare                       | Disable Link Sharing                                                                                   |
| CCMPermissionSettingTypeAftervalue | CanReadByLinkInTenant                  | Get access to links within the Tenant                                                                  |
| CCMPermissionSettingTypeAftervalue | CanEditByLinkInTenant                  | Get edited access to links within the Tenant                                                           |
| CCMPermissionSettingTypeAftervalue | CanReadByLinkInPartnerTenant           | Get access to links within Partner Tenant                                                              |
| CCMPermissionSettingTypeAftervalue | CanEditByLinkInPartnerTenant           | Get edited access to links within Partner Tenant                                                       |
| CCMPermissionSettingTypeAftervalue | CanReadByLinkInInternet                | Get access to links on the internet                                                                    |
| CCMPermissionSettingTypeAftervalue | CanEditByLinkInInternet                | Get edited access to links on the internet                                                             |
| CCMPermissionSettingTypeAftervalue | CanReadBySinglePageLinkInTenant        | Get access to links within the Tenant on a single page                                                 |
| CCMPermissionSettingTypeAftervalue | CanEditBySinglePageLinkInTenant        | Get edited access to links within the Tenant on a single page                                          |
| CCMPermissionSettingTypeAftervalue | CanReadBySinglePageLinkInPartnerTenant | Get access to links within Partner Tenant on a single page                                             |
| CCMPermissionSettingTypeAftervalue | CanEditBySinglePageLinkInPartnerTenant | Get edited access to links within Partner Tenant on a single page                                      |
| CCMPermissionSettingTypeAftervalue | CanReadBySinglePageLinkInInternet      | Get access to links on the internet on a single page                                                   |
| CCMPermissionSettingTypeAftervalue | CanEditBySinglePageLinkInInternet      | Get edited access to links on the internet on a single page                                            |
| CCMPermissionSettingTypeAftervalue | ExternalAccessOn                       | Turn external access on                                                                                |
| CCMPermissionSettingTypeAftervalue | ExternalAccessOff                      | Turn external access off                                                                               |
| CCMPermissionSettingTypeAftervalue | ExternalAccessPartnerTenant            | Share with Partner Tenant                                                                              |
| CCMPermissionSettingTypeAftervalue | SinglePageExternalAccessOn             | Turn on single-page external access                                                                    |
| CCMPermissionSettingTypeAftervalue | SinglePageExternalAccessPartnerTenant  | Share single pages with Partner Tenant                                                                 |
| CCMPermissionSettingTypeAftervalue | CommentEntityView                      | Comment with read permission                                                                           |
| CCMPermissionSettingTypeAftervalue | CommentEntityEdit                      | Comment with edit permission                                                                           |
| CCMPermissionSettingTypeAftervalue | SecurityEntityView                     | Copy, create duplicates, print, download, and export with read permission                              |
| CCMPermissionSettingTypeAftervalue | SecurityEntityEdit                     | Copy, create duplicates, print, download, and export with edit permission                              |
| CCMPermissionSettingTypeAftervalue | SecurityEntityFullAccess               | Copy, create duplicates, print, download, and export with full permission                              |
| CCMPermissionSettingTypeAftervalue | ShareEntityFullAccess                  | Manage and invite collaborators                                                                        |
| CCMPermissionSettingTypeAftervalue | ShareEntitySameTenant                  | Add collaborators in the same tenant                                                                   |
| CCMPermissionSettingTypeAftervalue | ShareEntityAnyone                      | Add collaborators by anyone                                                                            |
| CCMPermissionSettingTypeAftervalue | ManageCollaboratorEntityView           | Add collaborators with read permission                                                                 |
| CCMPermissionSettingTypeAftervalue | ManageCollaboratorEntityEdit           | Add collaborators with edit permission                                                                 |
| CCMPermissionSettingTypeAftervalue | ManageCollaboratorEntityFullAccess     | Add collaborators with full permission                                                                 |
| CCMPermissionSettingTypeAftervalue | DetailValTrue                          | On                                                                                                     |
| CCMPermissionSettingTypeAftervalue | DetailValFalse                         | off                                                                                                    |
| CCMPermissionSettingTypeAftervalue | ApplyEmbedEntityView                   | Quickly access permission-restricted reference documents and apply for permission with read permission |
| CCMPermissionSettingTypeAftervalue | ApplyEmbedEntityEdit                   | Quickly access permission-restricted reference documents and apply for permission with edit permission |
| CCMPermissionSettingTypeAftervalue | ApplyEmbedEntityFullAccess             | Quickly access permission-restricted reference documents and apply for permission with full permission |
| CCMPermissionSettingTypeAftervalue | ShowCollaboratorInfoEntityView         | View document collaborative avatar and like avatars with read permission                               |
| CCMPermissionSettingTypeAftervalue | ShowCollaboratorInfoEntityEdit         | View document collaborative avatars and like avatars with edit permission                              |
| CCMPermissionSettingTypeAftervalue | ShowCollaboratorInfoEntityFullAccess   | View document collaborative avatars and like avatars with full permission                              |
| CCMPermissionSettingTypeAftervalue | LeaderLinkShareEntityClose             | Disable leaders access                                                                                 |
| CCMPermissionSettingTypeAftervalue | LeaderLinkShareEntityCanView           | Allow leaders to read                                                                                  |
| CCMPermissionSettingTypeAftervalue | LeaderLinkShareEntityCanEdit           | Allow leaders to edit                                                                                  |
| CCMPermissionSettingTypeAftervalue | LeaderLinkShareEntityFullAccess        | Allow leaders to manage                                                                                |
| CCMPermissionSettingTypeAftervalue | LeaderCopyOff                          | Disable leaders replication                                                                            |
| CCMPermissionSettingTypeAftervalue | LeaderCopyOn                           | Allow leaders replication
### Cloud space permission setting parameter description
**CCMPermissionSettingType_space**                                                                                                                                               | **CCMPermissionSettingTypeAftervalue_space**                                                                              |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------- |
| WikiCreateFirstNode: Who can create top-level nodes in space?                                                                                                                | Admin: only Administrator;AdminAndMember: Administrator and member;                                            |
| WikiSecurityEntity: Does the read-only user have permission to copy, create duplicates, print, or export the content of the page within the space (Word, PDF, images, etc.)? | Allow; Deny;                                                                                                          |
| WikiCommentEntity: Does the read-only user have permission to comment on the page?                                                                                           | Allow; Deny;                                                                                                          |
| WikiExternalAccess: Can the page be shared outside of the Tenant?                                                                                                             | Allow; Deny;                                                                                                        |
| WikiExternalAccessEntity: Can the content of the space be shared externally?                                                                                               | ExternalAccessEntity_None; ExternalAccessEntity_Open; ExternalAccessEntity_Close; ExternalAccessEntity_PartnerTenant; |
|WikiSpaceScope: Knowledge Space Visibility                                                                                                                                   | WikiSpaceScope_PRIVATE; WikiSpaceScope_PUBLIC;

## **Details of action object extension fields**
|**Field**         | **Type** | **Description**          |
| --------- | ------ | --------- | 
| clone_source | string | Copies source | 
| third_party_appID | string | Third party app ID | 
| contain_file_num | int | Number of files or folders | 
| permission_setting_type | string | Link sharing settings | 
| permission_external_access_Type | boolean | Enable external access settings or not | 
| permission_share_type | string | Permission scope sharing settings | 
| file_service_source | string | Uploads file sources | 
| okr_download_content |string | Downloads OKR contents|
|container_type|string| Container type. Wiki or not|
|container_id|string| Container ID. Wiki ID| ## **Action object type**
**Type** | **Description**                                                                                     | **Searchability** |
| -------- | --------------------------------------------------------------------------------------------------- | ----------------- |
| 1        | User                                                                                                | Yes               |
| 2        | Department                                                                                          | Yes               |
| 3        | Tenant                                                                                              | Yes               |
| 4        | Group chat                                                                                          | Yes               |
| 5        | Docs (doc)                                                                                          | Yes               |
| 6        | Docs (sheet)                                                                                        | Yes               |
| 7        | Docs (file)                                                                                         | Yes               |
| 8        | Docs (folder)                                                                                       | Yes               |
| 9        | Docs (mindnote)                                                                                     | Yes               |
| 10       | Docs (bitable)                                                                                      | Yes               |
| 11       | Docs (slide)                                                                                        | Yes               |
| 12       | Bot                                                                                                 | No                |
| 13       | Audio/video call                                                                                    | No                |
| 14       | Regular files, with value type of app_id and app_file_id, usually the preview files in the chat box | No                |
| 15       | Link                                                                                                | Yes               |
| 16       | IM User Identification type                                                                         | No                |
| 17       | Image                                                                                               | Yes               |
| 18       | Video                                                                                               | Yes               |
| 19       | Regular files, with value type of fileKey, usually the regular files sent by users in the chat box  | Yes               |
| 20       | Gadget and H5 app                                                                                   | Yes               |
| 21       | Opened local files                                                                                  | No                |
| 29       | OKR                                                                                                 | No                |
| 31       | Docs (docx)                                                                                         | Yes               |
| 106      | Docs (slides)                                                                                       | Yes               |
| 107      | Knowledge space                                                                                     | Yes               |
| 121      | IM scene file                                                                                       | Yes               |
| 425      | Minutes                                                                                             | Yes               |
| 426      | Sync Block                                                                                          | Yes               |
| 210001   | Calendar                                                                                            | Yes               |
| 210002   | Calendar Event                                                                                      | Yes               |
| 210003   | Calendar Scheduler                                                                                  | Yes               |
| 210004   | External Calendar Account                                                                           | Yes |
