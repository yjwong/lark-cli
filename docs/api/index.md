# Lark Server API Documentation

*1133 pages across 30 sections. Generated from [Lark Open Platform](https://open.larksuite.com/document/server-docs/getting-started/getting-started).*

## API Call Guide

### Calling Process

    - [Process overview](server-docs/api-call-guide/calling-process/process-overview.md)
    - [Obtain access tokens](server-docs/api-call-guide/calling-process/obtain-access-tokens.md)
    - [Apply for API permission](server-docs/api-call-guide/calling-process/apply-for-api-permission.md)
    - [Configure app data permissions](server-docs/api-call-guide/calling-process/configure-app-data-permissions.md)
    - [Set IP allowlist](server-docs/api-call-guide/calling-process/set-ip-allowlist.md)
    - [Calling APIs](server-docs/api-call-guide/calling-process/calling-apis.md)

  - [Adjusting API call limits for custom apps](server-docs/api-call-guide/adjusting-api-call-limits-for-custom-apps.md)
  - [Server API list](server-docs/api-call-guide/server-api-list.md)
  - [Scope list](server-docs/api-call-guide/scope-list.md)
  - [Server SDK](server-docs/api-call-guide/server-sdk.md)
  - [Rate limits](server-docs/api-call-guide/rate-limits.md)
  - [General parameters](server-docs/api-call-guide/general-parameters.md)
  - [Server Error Codes](server-docs/api-call-guide/server-error-codes.md)

## Events and callbacks

### Event subscriptions

    - [Event Overview](server-docs/events-and-callbacks/event-subscriptions/event-overview.md)
    - [Event List](server-docs/events-and-callbacks/event-subscriptions/event-list.md)
- **Subscription Process**
  - **Step 1: Choose a subscription mode**
        - [Receive events through websocket](server-docs/events-and-callbacks/event-subscriptions/subscription-process/step-1-choose-a-subscription-mode/receive-events-through-websocket.md)
        - [Send notifications to developer's server](server-docs/events-and-callbacks/event-subscriptions/subscription-process/step-1-choose-a-subscription-mode/send-notifications-to-developers-server.md)
      - [Step 2: Add events](server-docs/events-and-callbacks/event-subscriptions/subscription-process/step-2-add-events.md)
      - [Step 3: Receive events](server-docs/events-and-callbacks/event-subscriptions/subscription-process/step-3-receive-events.md)
    - [Get event's outbound IP](server-docs/event-v1/outbound_ip/list.md)
    - [Event Callback Optimization Guide](server-docs/events-and-callbacks/event-subscriptions/event-callback-optimization-guide.md)

### Callback Subscription

    - [Callback overview](server-docs/events-and-callbacks/callback-subscription/callback-overview.md)
- **Step 1: Choose a subscription mode**
      - [Receive callbacks through websocket](server-docs/events-and-callbacks/callback-subscription/step-1-choose-a-subscription-mode/receive-callbacks-through-websocket.md)
      - [Send callbacks to developer's server](server-docs/events-and-callbacks/callback-subscription/step-1-choose-a-subscription-mode/send-callbacks-to-developers-server.md)
    - [Step 2: Add callback](server-docs/events-and-callbacks/callback-subscription/step-2-add-callback.md)
    - [Step 3: Receive callbacks](server-docs/events-and-callbacks/callback-subscription/step-3-receive-callbacks.md)

  - [FAQ](server-docs/events-and-callbacks/faq.md)

## Authenticate and Authorize

### Login state management

    - [Resource introduction](server-docs/passport-v1/session/usum.md)
    - [Get User Information](server-docs/authen-v1/user_info/get.md)
    - [Obtain desensitized user login information in batches](server-docs/passport-v1/session/query.md)
    - [Log out](server-docs/passport-v1/session/logout.md)

### Get Access Tokens

    - [Get custom app tenant_access_token](server-docs/auth-v3/auth/tenant_access_token_internal.md)
    - [Get custom app app_access_token](server-docs/auth-v3/auth/app_access_token_internal.md)
    - [Get authorization code](server-docs/authenticate-and-authorize/get-access-tokens/get-authorization-code.md)
    - [Get user_access_token](server-docs/authenticate-and-authorize/get-access-tokens/get-user-access-token.md)
    - [Refresh user_access_token](server-docs/authenticate-and-authorize/get-access-tokens/refresh-user-access-token.md)
    - [Re-pushing app_ticket](server-docs/auth-v3/auth/app_ticket_resend.md)
    - [Store applications get app_access_token](server-docs/auth-v3/auth/app_access_token.md)
    - [Store applications get tenant_access_token](server-docs/auth-v3/auth/tenant_access_token.md)


## Contacts

  - [Contacts overview](server-docs/contact-v3/resources.md)
  - [FAQ](server-docs/contacts/faq.md)
### Contact Scope

    - [Overview](server-docs/contact-v3/scope/overview.md)
    - [Obtain the Contacts permission scope](server-docs/contact-v3/scope/list.md)
- **Events**
      - [Contact Scope Updated](server-docs/contact-v3/scope/events/updated.md)

### User

    - [Overview](server-docs/contact-v3/user/field-overview.md)
    - [Country/Region Code Reference Table](server-docs/contact-v3/user/country-code-description.md)
    - [Create a user](server-docs/contact-v3/user/create.md)
    - [Modify user information in part](server-docs/contact-v3/user/patch.md)
    - [Update UserID](server-docs/contact-v3/user/update_user_id.md)
    - [Obtain single user's information](server-docs/contact-v3/user/get.md)
    - [Obtain multiple users' information](server-docs/contact-v3/user/batch.md)
    - [Obtain the list of users directly under a department](server-docs/contact-v3/user/find_by_department.md)
    - [Obtain user ID via email or mobile number](server-docs/contact-v3/user/batch_get_id.md)
    - [Search Users](server-docs/contacts/user/search-users.md)
    - [Delete User](server-docs/contact-v3/user/delete.md)
    - [Restore deleted users](server-docs/contact-v3/user/resurrect.md)
- **Events**
      - [Employment](server-docs/contact-v3/user/events/created.md)
      - [Termination](server-docs/contact-v3/user/events/deleted.md)
      - [Employee change](server-docs/contact-v3/user/events/updated.md)
    - [Fully Update User Information](server-docs/contact-v3/user/update.md)

### User group

    - [Overview](server-docs/contact-v3/group/overview.md)
    - [Add User Group](server-docs/contact-v3/group/create.md)
    - [Update User Group](server-docs/contact-v3/group/patch.md)
    - [Get User Group](server-docs/contact-v3/group/get.md)
    - [List User Group](server-docs/contact-v3/group/simplelist.md)
    - [Delete User Group](server-docs/contact-v3/group/delete.md)
- **User group member**
      - [Overview](server-docs/contact-v3/group-member/overview.md)
      - [Add User Group Member](server-docs/contact-v3/group-member/add.md)
      - [Batch add user group members](server-docs/contact-v3/group-member/batch_add.md)
      - [List User Group Member](server-docs/contact-v3/group-member/simplelist.md)
      - [Remove User Group Member](server-docs/contact-v3/group-member/remove.md)
      - [Batch remove user group members](server-docs/contact-v3/group-member/batch_remove.md)

### Custom user fields

    - [Overview](server-docs/contact-v3/custom_attr/overview.md)
    - [Obtain Custom User Fields](server-docs/contact-v3/custom_attr/list.md)
- **Events**
      - [Member field change event](server-docs/contact-v3/custom_attr_event/events/updated.md)

### Workforce type

    - [Overview](server-docs/contact-v3/employee_type_enum/overview.md)
    - [Create Workforce Type](server-docs/contact-v3/employee_type_enum/create.md)
    - [Update Workforce Type](server-docs/contact-v3/employee_type_enum/update.md)
    - [Fetch Workforce Type](server-docs/contact-v3/employee_type_enum/list.md)
    - [Delete Workforce Type](server-docs/contact-v3/employee_type_enum/delete.md)
- **Events**
      - [Add New Workforce Types](server-docs/contact-v3/employee_type_enum/events/created.md)
      - [Enable Workforce Types](server-docs/contact-v3/employee_type_enum/events/actived.md)
      - [Deactivate Workforce Types](server-docs/contact-v3/employee_type_enum/events/deactivated.md)
      - [Delete Workforce Types](server-docs/contact-v3/employee_type_enum/events/deleted.md)
      - [Change Names of Workforce Types](server-docs/contact-v3/employee_type_enum/events/updated.md)

### Department

    - [Overview](server-docs/contact-v3/department/field-overview.md)
    - [Create a department](server-docs/contact-v3/department/create.md)
    - [Modify department information in part](server-docs/contact-v3/department/patch.md)
    - [Update department information in whole](server-docs/contact-v3/department/update.md)
    - [Update DepartmentID](server-docs/contact-v3/department/update_department_id.md)
    - [change department group to common group](server-docs/contact-v3/department/unbind_department_chat.md)
    - [Obtain single department information](server-docs/contact-v3/department/get.md)
    - [Obtain bulk department information](server-docs/contact-v3/department/batch.md)
    - [Obtain the list of sub-departments](server-docs/contact-v3/department/children.md)
    - [Obtain parent department information](server-docs/contact-v3/department/parent.md)
    - [Search for departments](server-docs/contact-v3/department/search.md)
    - [Delete Department](server-docs/contact-v3/department/delete.md)
- **Events**
      - [Department Created](server-docs/contact-v3/department/events/created.md)
      - [Department Deleted](server-docs/contact-v3/department/events/deleted.md)
      - [Department Updated](server-docs/contact-v3/department/events/updated.md)

### role

    - [Resource introduction](server-docs/contact-v3/functional_role/resource-introduction.md)
    - [Create role](server-docs/contact-v3/functional_role/create.md)
    - [Update role name](server-docs/contact-v3/functional_role/update.md)
    - [Delete role](server-docs/contact-v3/functional_role/delete.md)

### role member

    - [Resource introduction](server-docs/contact-v3/functional_role-member/resource-introduction.md)
    - [Batch create role members](server-docs/contact-v3/functional_role-member/batch_create.md)
    - [Batch update role members scopes](server-docs/contact-v3/functional_role-member/scopes.md)
    - [The administrative scope of a member under a query role](server-docs/contact-v3/functional_role-member/get.md)
    - [Query all member information under the role](server-docs/contact-v3/functional_role-member/list.md)
    - [Delete members from roles](server-docs/contact-v3/functional_role-member/batch_delete.md)

  - [Contact Scope Description](server-docs/contacts/contact-scope-description.md)

## Organization

  - [Organizational structure overview](server-docs/directory-v1/overview.md)
  - [Field enumeration](server-docs/directory-v1/field-enumeration.md)
  - [Filter usage](server-docs/directory-v1/filter-usage.md)
### Employee management

    - [Overview](server-docs/directory-v1/employee/overview.md)
    - [Resources Introduction](server-docs/directory-v1/employee/resources-introduction.md)
    - [Create employee](server-docs/directory-v1/employee/create.md)
    - [Update employee information](server-docs/directory-v1/employee/patch.md)
    - [Delete employee](server-docs/directory-v1/employee/delete.md)
    - [Reinstate departed employees](server-docs/directory-v1/employee/resurrect.md)
    - [Update un-resigned employees to be resigned](server-docs/directory-v1/employee/to_be_resigned.md)
    - [Update pre-resigned members to un-resigned employees](server-docs/directory-v1/employee/regular.md)
    - [Obtain employee information in batches](server-docs/directory-v1/employee/mget.md)
    - [Get employee list in batches](server-docs/directory-v1/employee/filter.md)
    - [Search employee](server-docs/directory-v1/employee/search.md)
- **Events**
      - [Employment](server-docs/directory-v1/employee/events/created.md)
      - [Employee modified](server-docs/directory-v1/employee/events/updated.md)
      - [Employee resigned](server-docs/directory-v1/employee/events/resigned.md)
      - [Employee resurrected](server-docs/directory-v1/employee/events/resurrect.md)
      - [Employee to be resigned](server-docs/directory-v1/employee/events/to_be_resigned.md)
      - [To be resigned to regular](server-docs/directory-v1/employee/events/regular.md)

### Department management

    - [Overview](server-docs/directory-v1/department/overview.md)
    - [Resource Introduction](server-docs/directory-v1/department/resource-introduction.md)
    - [Create department](server-docs/directory-v1/department/create.md)
    - [Update department](server-docs/directory-v1/department/patch.md)
    - [Delete department](server-docs/directory-v1/department/delete.md)
    - [Batch get department info](server-docs/directory-v1/department/mget.md)
    - [List department](server-docs/directory-v1/department/filter.md)
    - [Search department](server-docs/directory-v1/department/search.md)
- **Events**
      - [Department Created](server-docs/directory-v1/department/events/created.md)
      - [Department Modified](server-docs/directory-v1/department/events/updated.md)
      - [Department Deleted](server-docs/directory-v1/department/events/deleted.md)


## Messaging

  - [Overview](server-docs/im-v1/introduction.md)
  - [FAQs](server-docs/im-v1/guide/faq.md)
### Message management

    - [Resource introduction](server-docs/im-v1/message/intro.md)
    - [Thread Introduction](server-docs/im-v1/message/thread-introduction.md)
    - [Send message](server-docs/im-v1/message/create.md)
    - [Reply message](server-docs/im-v1/message/reply.md)
    - [Edit message](server-docs/im-v1/message/update.md)
    - [Recall message](server-docs/im-v1/message/delete.md)
    - [Forward a message](server-docs/im-v1/message/forward.md)
    - [Merge forward messages](server-docs/im-v1/message/merge_forward.md)
    - [Forward a thread](server-docs/im-v1/thread/forward.md)
    - [Query the read status of message](server-docs/im-v1/message/read_users.md)
    - [Get chat history](server-docs/im-v1/message/list.md)
    - [Obtain resource files in messages](server-docs/im-v1/message-resource/get.md)
    - [Obtain the content of the specified message](server-docs/im-v1/message/get.md)
- **Message content introduction**
      - [Sent message content](server-docs/im-v1/message/create_json.md)
      - [Received message content](server-docs/im-v1/message/events/message_content.md)
- **Events**
      - [Receive message](server-docs/im-v1/message/events/receive.md)
      - [Message read](server-docs/im-v1/message/events/message_read.md)
      - [Recall message](server-docs/im-v1/message/events/recalled.md)

### Batch message

    - [Send messages in batches](server-docs/messaging/batch-message/send-messages-in-batches.md)
    - [Recall messages in batches](server-docs/im-v1/batch_message/delete.md)
    - [Query the number of pushers and readers of batch messages](server-docs/im-v1/batch_message/read_user.md)
    - [Query the overall progress of a batch message](server-docs/im-v1/batch_message/get_progress.md)

### Images message

    - [Upload image](server-docs/im-v1/image/create.md)
    - [Download image](server-docs/im-v1/image/get.md)

### File message

    - [Upload file](server-docs/im-v1/file/create.md)
    - [Download file](server-docs/im-v1/file/get.md)

### Buzz message

    - [Send buzz messages in apps](server-docs/im-v1/message/urgent_app.md)

### Message reaction

    - [Resource introduction](server-docs/im-v1/message-reaction/overview-of-message-reaction-resources.md)
    - [Emoji copy](server-docs/im-v1/message-reaction/emojis-introduce.md)
    - [Add a reaction for a message](server-docs/im-v1/message-reaction/create.md)
    - [List message reactions](server-docs/im-v1/message-reaction/list.md)
    - [Delete a reaction for a message](server-docs/im-v1/message-reaction/delete.md)
- **Events**
      - [Add a reaction for a message](server-docs/im-v1/message-reaction/events/created.md)
      - [Delete a reaction for a message](server-docs/im-v1/message-reaction/events/deleted.md)

### Pin

    - [Pin a message](server-docs/im-v1/pin/create.md)
    - [Unpin a message](server-docs/im-v1/pin/delete.md)
    - [Get pins in group](server-docs/im-v1/pin/list.md)

### Message card

    - [Resource introduction](server-docs/im-v1/message-card/overview.md)
    - [Update the message card sent by the app](server-docs/im-v1/message/patch.md)
    - [Delay update message card](server-docs/messaging/message-card/delay-update-message-card.md)
    - [Send message cards that are only visible to certain people](server-docs/messaging/message-card/send-message-cards-that-are-only-visible-to-certain-people.md)
    - [Delete message cards that are only visible to certain people](server-docs/messaging/message-card/delete-message-cards-that-are-only-visible-to-certain-people.md)


## Group Chat

  - [Overview](server-docs/group-chat/overview.md)
### Group management

    - [Resource introduction](server-docs/im-v1/chat/chat-info/intro.md)
    - [Chat ID description](server-docs/im-v1/chat-id-description.md)
    - [Create a group](server-docs/im-v1/chat/create.md)
    - [Delete a group](server-docs/im-v1/chat/delete.md)
    - [Update group information](server-docs/im-v1/chat/update.md)
    - [Updates group speech scopes](server-docs/im-v1/chat-moderation/update.md)
    - [Obtain group information](server-docs/im-v1/chat/get.md)
    - [Update group pinning](server-docs/im-v1/chat-top_notice/put_top_notice.md)
    - [Revoke group pinning](server-docs/im-v1/chat-top_notice/delete_top_notice.md)
    - [Obtain groups where the user or bot is a member](server-docs/im-v1/chat/list.md)
    - [Search for groups visible to a user or bot](server-docs/im-v1/chat/search.md)
    - [Obtains the group member speech scopes](server-docs/im-v1/chat-moderation/get.md)
    - [Get group share link](server-docs/im-v1/chat/link.md)
- **Events**
      - [Group disbanded](server-docs/im-v1/chat/events/disbanded.md)
      - [Group configuration changed](server-docs/im-v1/chat/events/updated.md)

### Group member

    - [Resource introduction](server-docs/im-v1/chat/chat-member/intro.md)
    - [Specify group administrators](server-docs/im-v1/chat-managers/add_managers.md)
    - [Delete group administrators](server-docs/im-v1/chat-managers/delete_managers.md)
    - [Add users or bots to a group](server-docs/im-v1/chat-members/create.md)
    - [Users or bots join a group voluntarily](server-docs/im-v1/chat-members/me_join.md)
    - [Remove users or bots from a group](server-docs/im-v1/chat-members/delete.md)
    - [Obtain group member list](server-docs/im-v1/chat-members/get.md)
    - [Determine whether a user or bot is in a group](server-docs/im-v1/chat-members/is_in_chat.md)
- **Events**
      - [Users join a group](server-docs/im-v1/chat-member-user/events/added.md)
      - [Users leave a group](server-docs/im-v1/chat-member-user/events/deleted.md)
      - [Invitation revoked](server-docs/im-v1/chat-member-user/events/withdrawn.md)
      - [Bots added to a group](server-docs/im-v1/chat-member-bot/events/added.md)
      - [Bots removed from a group](server-docs/im-v1/chat-member-bot/events/deleted.md)
      - [User and bot chat first created](server-docs/group-chat/group-member/events/user-and-bot-chat-first-created.md)

### Upgraded Group announcement

    - [Group announcement overview](server-docs/group-chat/upgraded-group-announcement/group-announcement-overview.md)
    - [Group announcement version changes](server-docs/group-announcement-version-changes.md)
    - [Group Announcement FAQs](server-docs/group-chat/upgraded-group-announcement/group-announcement-faqs.md)
    - [Group announcement data structure](server-docs/group-chat/upgraded-group-announcement/group-announcement-data-structure.md)
- **Group announcement**
      - [Obtain the basic information of a group announcement](server-docs/docx-v1/chat-announcement/get.md)
      - [Obtain all blocks of a group announcement](server-docs/docx-v1/chat-announcement-block/list.md)
- **Block**
      - [Create blocks in group announcement](server-docs/docx-v1/chat-announcement-block-children/create.md)
      - [Batch update blocks in group announcement](server-docs/docx-v1/chat-announcement-block/batch_update.md)
      - [Obtain the block content in group announcement](server-docs/docx-v1/chat-announcement-block/get.md)
      - [Obtain all the child blocks](server-docs/docx-v1/chat-announcement-block-children/get.md)
      - [Delete blocks in group announcement](server-docs/docx-v1/chat-announcement-block-children/batch_delete.md)

### Group announcement

    - [Resource introduction](server-docs/im-v1/chat/chat-announcement/intro.md)
    - [Update group announcement info](server-docs/im-v1/chat-announcement/patch.md)
    - [Obtain group announcement information](server-docs/im-v1/chat-announcement/get.md)

### Chat tab

    - [Add chat tab](server-docs/im-v1/chat-tab/create.md)
    - [Delete chat tab](server-docs/im-v1/chat-tab/delete_tabs.md)
    - [Update chat tab](server-docs/im-v1/chat-tab/update_tabs.md)
    - [Sort chat tab](server-docs/im-v1/chat-tab/sort_tabs.md)
    - [Pull chat tabs](server-docs/im-v1/chat-tab/list_tabs.md)

### Chat menu

    - [Resource introduction](server-docs/im-v1/chat-menu_tree/overview.md)
    - [Add chat menu](server-docs/im-v1/chat-menu_tree/create.md)
    - [Delete chat menu](server-docs/im-v1/chat-menu_tree/delete.md)
    - [Patch chat menu item info](server-docs/im-v1/chat-menu_item/patch.md)
    - [Sort chat menus](server-docs/im-v1/chat-menu_tree/sort.md)
    - [Get chat menus](server-docs/im-v1/chat-menu_tree/get.md)


## Lark Card

  - [Feishu Card resource overview](server-docs/cardkit-v1/feishu-card-resource-overview.md)
### Card

    - [Create card entity](server-docs/cardkit-v1/card/create.md)
    - [Update card config](server-docs/cardkit-v1/card/settings.md)
    - [Batch update card entity](server-docs/cardkit-v1/card/batch_update.md)
    - [Update card entity](server-docs/cardkit-v1/card/update.md)

### Element

    - [Insert element](server-docs/cardkit-v1/card-element/create.md)
    - [Update element](server-docs/cardkit-v1/card-element/update.md)
    - [Patch Update Element](server-docs/cardkit-v1/card-element/patch.md)
    - [Stream update text](server-docs/cardkit-v1/card-element/content.md)
    - [Delete element](server-docs/cardkit-v1/card-element/delete.md)


## Feed

  - [Overview](server-docs/im-v2/overview.md)
### Apps

    - [Create app feed card](server-docs/im-v2/app_feed_card/create.md)
    - [Update app feed card](server-docs/im-v2/app_feed_card-batch/update.md)
    - [Delete app feed card](server-docs/im-v2/app_feed_card-batch/delete.md)

### Groups & Bots

    - [Update feed card button](server-docs/im-v2/chat_button/update.md)
    - [Instant reminder](server-docs/im-v2/feed_card/patch.md)


## Docs

  - [Docs Overview](server-docs/docs/docs-overview.md)
  - [Docs Overview and Best Practice](server-docs/docs/docs-overview-and-best-practice.md)
### Space

    - [Overview](server-docs/docs/space/overview.md)
    - [FAQ](server-docs/drive-v1/faq.md)
- **Folder**
      - [Get Root Folder Meta](server-docs/docs/space/folder/get-root-folder-meta.md)
      - [List items in folder](server-docs/drive-v1/file/list.md)
      - [Get Folder Meta](server-docs/docs/space/folder/get-folder-meta.md)
      - [Create Folder](server-docs/drive-v1/file/create_folder.md)
      - [Move/Delete Folder](server-docs/docs/space/folder/move-delete-folder.md)
- **File**
      - [Obtain Metadata](server-docs/drive-v1/meta/batch_query.md)
      - [Get File Statistics](server-docs/drive-v1/file-statistics/get.md)
      - [Copy File](server-docs/drive-v1/file/copy.md)
      - [Move File/Folder](server-docs/drive-v1/file/move.md)
      - [Delete File/Folder](server-docs/drive-v1/file/delete.md)
      - [Create online document](server-docs/docs/space/file/create-online-document.md)
  - **Async task status**
        - [Query Task Status](server-docs/drive-v1/file/task_check.md)
- **Media**
      - [Introduction](server-docs/drive-v1/media/introduction.md)
      - [Upload Media](server-docs/drive-v1/media/upload_all.md)
      - [Download Media](server-docs/drive-v1/media/download.md)
      - [Get Temporary Download URL of Media](server-docs/drive-v1/media/batch_get_tmp_download_url.md)
  - **Multipart Upload**
        - [Introduction](server-docs/drive-v1/media/multipart-upload-media/introduction.md)
        - [Multipart Upload Media (Prepare)](server-docs/drive-v1/media/upload_prepare.md)
        - [Multipart Upload Media (Upload)](server-docs/drive-v1/media/upload_part.md)
        - [Multipart Upload Media (Finish)](server-docs/drive-v1/media/upload_finish.md)
- **Event**
      - [Subscribe Docs events](server-docs/drive-v1/file/subscribe.md)
      - [Get Docs events subscription status](server-docs/drive-v1/file/get_subscribe.md)
      - [Unsubscribe Docs events](server-docs/drive-v1/file/delete_subscribe.md)
  - **List**
        - [File title updated](server-docs/docs/space/event/list/file-title-updated.md)
        - [File read](server-docs/docs/space/event/list/file-read.md)
        - [File edited](server-docs/docs/space/event/list/file-edited.md)
        - [File collaborator added](server-docs/docs/space/event/list/file-collaborator-added.md)
        - [File collaborator removed](server-docs/docs/space/event/list/file-collaborator-removed.md)
        - [File deleted to trash can](server-docs/docs/space/event/list/file-deleted-to-trash-can.md)
        - [File deleted completely](server-docs/docs/space/event/list/file-deleted-completely.md)
        - [Bitable Record Changed](server-docs/drive-v1/event/list/bitable-record-changed.md)
        - [File created in folder](server-docs/drive-v1/file/events/created_in_folder.md)
        - [File collaborator permission application](server-docs/drive-v1/file/events/permission_member_applied.md)
- **Search**
      - [Document Search](server-docs/docs/space/search/document-search.md)
- **Upload**
      - [Upload Files](server-docs/drive-v1/file/upload_all.md)
  - **Multipart Upload**
        - [Introduction](server-docs/drive-v1/file/multipart-upload-file-/introduction.md)
        - [Multipart Upload Files (Prepare)](server-docs/drive-v1/file/upload_prepare.md)
        - [Multipart Upload Files (Upload)](server-docs/drive-v1/file/upload_part.md)
        - [Multipart Upload Files (Finish)](server-docs/drive-v1/file/upload_finish.md)
- **Download**
      - [Download Files](server-docs/drive-v1/file/download.md)
- **import**
      - [Overview](server-docs/drive-v1/import_task/import-user-guide.md)
      - [Create Import Task](server-docs/drive-v1/import_task/create.md)
      - [Query import task result](server-docs/drive-v1/import_task/get.md)
- **export**
      - [Overview](server-docs/drive-v1/export_task/export-user-guide.md)
      - [Create an export task](server-docs/drive-v1/export_task/create.md)
      - [get export task result](server-docs/drive-v1/export_task/get.md)
      - [download export file](server-docs/drive-v1/export_task/download.md)
- **Document Version**
      - [Overview](server-docs/drive-v1/file-version/overview.md)
      - [Create document version](server-docs/drive-v1/file-version/create.md)
      - [Delete document version](server-docs/drive-v1/file-version/delete.md)
      - [Get document version](server-docs/drive-v1/file-version/get.md)
      - [List document version](server-docs/drive-v1/file-version/list.md)
- **Subscription**
      - [Create Subscription](server-docs/drive-v1/file-subscription/create.md)
      - [Get subscription status](server-docs/drive-v1/file-subscription/get.md)
      - [update subscription status](server-docs/drive-v1/file-subscription/patch.md)

### Wiki

    - [Overview](server-docs/wiki-v2/getting-start.md)
    - [API FAQ](server-docs/wiki-v2/wiki-qa.md)
- **space**
      - [Get a list of Wiki spaces](server-docs/wiki-v2/space/list.md)
      - [Access to Wiki space information](server-docs/wiki-v2/space/get.md)
      - [Create Wiki space](server-docs/wiki-v2/space/create.md)
- **member**
      - [Add Wiki space members](server-docs/wiki-v2/space-member/create.md)
      - [Delete Wiki space members](server-docs/wiki-v2/space-member/delete.md)
- **Setting**
      - [Update Wiki space settings](server-docs/wiki-v2/space-setting/update.md)
- **node**
      - [Create node](server-docs/wiki-v2/space-node/create.md)
      - [Get Wiki node information](server-docs/wiki-v2/space/get_node.md)
      - [Get the list of child nodes](server-docs/wiki-v2/space-node/list.md)
      - [Move node](server-docs/wiki-v2/space-node/move.md)
      - [Update title](server-docs/wiki-v2/space-node/update_title.md)
      - [Create a node copy](server-docs/wiki-v2/space-node/copy.md)
- **task**
      - [Add an existing document to Wiki space](server-docs/wiki-v2/space-node/move_docs_to_wiki.md)
      - [Retrieve the result of Wiki task](server-docs/wiki-v2/task/get.md)
    - [Search Wiki](server-docs/docs/wiki/search-wiki.md)

### Document

    - [Document overview](server-docs/docs/document/document-overview.md)
    - [Document FAQs](server-docs/docx-v1/faq.md)
- **Data Structure**
      - [Document](server-docs/docx-v1/data-structure/document.md)
      - [Block](server-docs/docx-v1/data-structure/block.md)
      - [Emoji](server-docs/docx-v1/emoji.md)
- **API reference**
  - **Document**
        - [Create Document](server-docs/docx-v1/document/create.md)
        - [Query Document](server-docs/docx-v1/document/get.md)
        - [Query Document Raw Content](server-docs/docx-v1/document/raw_content.md)
        - [Query All Blocks](server-docs/docx-v1/document-block/list.md)
    - [Document version changes](server-docs/docs/document/document-version-changes.md)
- **Block**
      - [Create Block](server-docs/docx-v1/document-block-children/create.md)
      - [Create nested blocks](server-docs/docx-v1/document-block-descendant/create.md)
      - [Update Block](server-docs/docx-v1/document-block/patch.md)
      - [Query Block](server-docs/docx-v1/document-block/get.md)
      - [Batch Update Blocks](server-docs/docx-v1/document-block/batch_update.md)
      - [Query Block Children](server-docs/docx-v1/document-block-children/get.md)
      - [Delete Blocks](server-docs/docx-v1/document-block-children/batch_delete.md)
      - [Convert To Blocks](server-docs/docx-v1/document/convert.md)

### Sheets

- **spreadsheet**
      - [Create Spreadsheet](server-docs/sheets-v3/spreadsheet/create.md)
      - [Obtain Spreadsheet Metadata](server-docs/docs/sheets/spreadsheet/obtain-spreadsheet-metadata.md)
      - [Update Spreadsheet Properties](server-docs/docs/sheets/spreadsheet/update-spreadsheet-properties.md)
- **sheet**
      - [Query sheet](server-docs/sheets-v3/spreadsheet-sheet/get.md)
      - [Get sheet](server-docs/sheets-v3/spreadsheet-sheet/query.md)
      - [Operate Sheets](server-docs/docs/sheets/sheet/operate-sheets.md)
      - [Update Sheet Properties](server-docs/docs/sheets/sheet/update-sheet-properties.md)
- **Row Column**
      - [Add Rows or Columns](server-docs/docs/sheets/row-column/add-rows-or-columns.md)
      - [Insert Rows or Columns](server-docs/docs/sheets/row-column/insert-rows-or-columns.md)
      - [Update Rows or Columns](server-docs/docs/sheets/row-column/update-rows-or-columns.md)
      - [Move Dimension](server-docs/sheets-v3/spreadsheet-sheet/move_dimension.md)
      - [ Delete Rows or Columns](server-docs/docs/sheets/row-column/delete-rows-or-columns.md)
- **Cell**
      - [Prepend Data](server-docs/docs/sheets/cell/prepend-data.md)
      - [Append Data](server-docs/docs/sheets/cell/append-data.md)
      - [Reading a Single Range](server-docs/docs/sheets/cell/reading-a-single-range.md)
      - [Write Data to a Single Range](server-docs/docs/sheets/cell/write-data-to-a-single-range.md)
      - [Reading Multiple Ranges](server-docs/docs/sheets/cell/reading-multiple-ranges.md)
      - [Write Data to Multiple Ranges](server-docs/docs/sheets/cell/write-data-to-multiple-ranges.md)
      - [Set Cell Style](server-docs/docs/sheets/cell/set-cell-style.md)
      - [Batch Set Cell Style](server-docs/docs/sheets/cell/batch-set-cell-style.md)
      - [Write Images](server-docs/docs/sheets/cell/write-images.md)
      - [Merge Cells](server-docs/docs/sheets/cell/merge-cells.md)
      - [Split Cells](server-docs/docs/sheets/cell/split-cells.md)
      - [Find cells](server-docs/sheets-v3/spreadsheet-sheet/find.md)
      - [Replace cells](server-docs/sheets-v3/spreadsheet-sheet/replace.md)
- **Filter**
      - [Filter User Guide](server-docs/sheets-v3/spreadsheet-sheet-filter/filter-user-guide.md)
      - [Obtain Filter](server-docs/sheets-v3/spreadsheet-sheet-filter/get.md)
      - [Create Filter](server-docs/sheets-v3/spreadsheet-sheet-filter/create.md)
      - [Update Filter](server-docs/sheets-v3/spreadsheet-sheet-filter/update.md)
      - [Delete Filter](server-docs/sheets-v3/spreadsheet-sheet-filter/delete.md)
- **Filter view**
      - [Obtain Filter View](server-docs/sheets-v3/spreadsheet-sheet-filter_view/get.md)
      - [Query Filter View](server-docs/sheets-v3/spreadsheet-sheet-filter_view/query.md)
      - [Create Filter View](server-docs/sheets-v3/spreadsheet-sheet-filter_view/create.md)
      - [Update Filter View](server-docs/sheets-v3/spreadsheet-sheet-filter_view/patch.md)
      - [Delete Filter View](server-docs/sheets-v3/spreadsheet-sheet-filter_view/delete.md)
  - **filter view**
        - [Filter View Condition User Guide](server-docs/sheets-v3/spreadsheet-sheet-filter_view-condition/filter-view-condition-user-guide.md)
        - [Obtain Filter Condition](server-docs/sheets-v3/spreadsheet-sheet-filter_view-condition/get.md)
        - [Query Filter Condition](server-docs/sheets-v3/spreadsheet-sheet-filter_view-condition/query.md)
        - [Create Filter Condition](server-docs/sheets-v3/spreadsheet-sheet-filter_view-condition/create.md)
        - [Update Filter Condition](server-docs/sheets-v3/spreadsheet-sheet-filter_view-condition/update.md)
        - [Delete Filter Condition](server-docs/sheets-v3/spreadsheet-sheet-filter_view-condition/delete.md)
- **Row Columu - Protect Range**
      - [Add Locked Cells](server-docs/docs/sheets/row-columu-protect-range/add-locked-cells.md)
      - [Retrieve Protection Scopes](server-docs/docs/sheets/row-columu-protect-range/retrieve-protection-scopes.md)
      - [Modify Protection Scopes](server-docs/docs/sheets/row-columu-protect-range/modify-protection-scopes.md)
      - [Delete Protection Scopes](server-docs/docs/sheets/row-columu-protect-range/delete-protection-scopes.md)
- **Sheet - Data Validation**
      - [Data Validation User Guide (Drop-down List Included)](server-docs/docs/sheets/sheet-data-validation/data-validation-user-guide-dropdown-list-included.md)
      - [Set Drop-down List](server-docs/docs/sheets/sheet-data-validation/set-dropdown-list.md)
      - [Update Drop-down List Setting](server-docs/docs/sheets/sheet-data-validation/update-dropdown-list-setting.md)
      - [Query Drop-down List in the Specific Range](server-docs/docs/sheets/sheet-data-validation/query-dropdown-list-in-the-specific-range.md)
      - [Remove the Drop-down List in the Specific Range](server-docs/docs/sheets/sheet-data-validation/remove-the-dropdown-list-in-the-specific-range.md)
- **Sheet - Conditional Formatting**
      - [Conditional Formatting User Guide](server-docs/docs/sheets/sheet-conditional-formatting/conditional-formatting-user-guide.md)
      - [Create Conditional Formatting Rules](server-docs/docs/sheets/sheet-conditional-formatting/create-conditional-formatting-rules.md)
      - [Obtain Conditional Formatting Rules](server-docs/docs/sheets/sheet-conditional-formatting/obtain-conditional-formatting-rules.md)
      - [Update Conditional Formatting Rules](server-docs/docs/sheets/sheet-conditional-formatting/update-conditional-formatting-rules.md)
      - [Remove Conditional Formatting Rules](server-docs/docs/sheets/sheet-conditional-formatting/remove-conditional-formatting-rules.md)
- **Floating image**
      - [Float image user guide](server-docs/sheets-v3/spreadsheet-sheet-float_image/float-image-user-guide.md)
      - [Obtain Floating Image](server-docs/sheets-v3/spreadsheet-sheet-float_image/get.md)
      - [Query Floating Image](server-docs/sheets-v3/spreadsheet-sheet-float_image/query.md)
      - [Create Floating Image](server-docs/sheets-v3/spreadsheet-sheet-float_image/create.md)
      - [Update Floating Image](server-docs/sheets-v3/spreadsheet-sheet-float_image/patch.md)
      - [Delete Floating Image](server-docs/sheets-v3/spreadsheet-sheet-float_image/delete.md)
- **Guide**
      - [Overview](server-docs/docs/sheets/guide/overview.md)
      - [sheets-faq](server-docs/docs/sheets/guide/sheetsfaq.md)

### Base

    - [Base overview](server-docs/docs/base/base-overview.md)
    - [Base FAQs](server-docs/bitable-v1/faq.md)
    - [Base data structure overview](server-docs/docs/base/base-data-structure-overview.md)
- **App**
      - [Create App](server-docs/bitable-v1/app/create.md)
      - [Copy App](server-docs/bitable-v1/app/copy.md)
      - [Get App Info](server-docs/bitable-v1/app/get.md)
      - [Update App Name](server-docs/bitable-v1/app/update.md)
- **Table**
      - [Create table](server-docs/bitable-v1/app-table/create.md)
      - [Batch create table](server-docs/bitable-v1/app-table/batch_create.md)
      - [Update data table](server-docs/bitable-v1/app-table/patch.md)
      - [List all tables](server-docs/bitable-v1/app-table/list.md)
      - [Delete Table](server-docs/bitable-v1/app-table/delete.md)
      - [Batch delete table](server-docs/bitable-v1/app-table/batch_delete.md)
- **View**
      - [Add View](server-docs/bitable-v1/app-table-view/create.md)
      - [Update View](server-docs/bitable-v1/app-table-view/patch.md)
      - [List Views](server-docs/bitable-v1/app-table-view/list.md)
      - [Get View](server-docs/bitable-v1/app-table-view/get.md)
      - [Delete View](server-docs/bitable-v1/app-table-view/delete.md)
- **Record**
      - [Base record data structure](server-docs/bitable-v1/app-table-record/bitable-record-data-structure-overview.md)
      - [Record filter guide](server-docs/bitable-v1/app-table-record/record-filter-guide.md)
      - [Add a sub-record in a Base table](server-docs/bitable-v1/app-table-record/add-a-sub-record-in-a-base-table.md)
      - [Create a record](server-docs/bitable-v1/app-table-record/create.md)
      - [Update a record](server-docs/bitable-v1/app-table-record/update.md)
      - [Search records](server-docs/bitable-v1/app-table-record/search.md)
      - [Delete a record](server-docs/bitable-v1/app-table-record/delete.md)
      - [Create records](server-docs/bitable-v1/app-table-record/batch_create.md)
      - [Update records](server-docs/bitable-v1/app-table-record/batch_update.md)
      - [Batch get records](server-docs/bitable-v1/app-table-record/batch_get.md)
      - [Delete records](server-docs/bitable-v1/app-table-record/batch_delete.md)
- **Field**
      - [Field edit development guide](server-docs/bitable-v1/app-table-field/guide.md)
      - [Attachment Field](server-docs/bitable-v1/attachment.md)
      - [Create field](server-docs/bitable-v1/app-table-field/create.md)
      - [Update field](server-docs/bitable-v1/app-table-field/update.md)
      - [List fields](server-docs/bitable-v1/app-table-field/list.md)
      - [Delete field](server-docs/bitable-v1/app-table-field/delete.md)
- **Field Group**
      - [Create field group](server-docs/bitable-v1/app-table-field_group/create.md)
- **Dashboards**
      - [Copy dashboards](server-docs/bitable-v1/app-dashboard/copy.md)
      - [List dashboards](server-docs/bitable-v1/app-dashboard/list.md)
- **Form**
      - [Upgrade form](server-docs/bitable-v1/app-table-form/upgrade.md)
      - [Patch form](server-docs/bitable-v1/app-table-form/patch.md)
      - [List form](server-docs/bitable-v1/app-table-form/get.md)
      - [Patch form fields](server-docs/bitable-v1/app-table-form-field/patch.md)
      - [List form fields](server-docs/bitable-v1/app-table-form-field/list.md)
- **Advanced Permission**
      - [Advanced permission overview](server-docs/bitable-v1/app-role/advanced-permission-guide.md)
  - **Setting**
        - [Delete role](server-docs/bitable-v1/app-role/delete.md)
        - [List roles](server-docs/bitable-v1/app-role/list.md)
        - [Create role](server-docs/bitable-v1/app-role/create.md)
        - [Update role](server-docs/bitable-v1/app-role/update.md)
  - **Member**
        - [Create member](server-docs/bitable-v1/app-role-member/create.md)
        - [Batch create members](server-docs/bitable-v1/app-role-member/batch_create.md)
        - [List members](server-docs/bitable-v1/app-role-member/list.md)
        - [Delete member](server-docs/bitable-v1/app-role-member/delete.md)
        - [Batch delete members](server-docs/bitable-v1/app-role-member/batch_delete.md)
- **Workflow**
      - [List automations](server-docs/bitable-v1/app-workflow/list.md)
      - [Update Workflow status](server-docs/bitable-v1/app-workflow/update.md)
- **Workflow**
      - [List Workflows](server-docs/bitable-v1/app-block_workflow/list.md)
- **Events**
      - [Base app record changed](server-docs/drive-v1/file/events/bitable_record_changed.md)
      - [Base app field changed](server-docs/drive-v1/file/events/bitable_field_changed.md)

### Board

    - [Board overview](server-docs/board-v1/overview.md)
    - [data structure](server-docs/board-v1/data-structure.md)
- **nodes**
      - [list nodes](server-docs/board-v1/whiteboard-node/list.md)
- **Thumbnail**
      - [whiteboard image](server-docs/board-v1/whiteboard/download_as_image.md)

### Permission

- **Member**
      - [Add permissions](server-docs/drive-v1/permission-member/create.md)
      - [UpdatePermissionMember](server-docs/drive-v1/permission-member/update.md)
      - [DeletePermissionMember](server-docs/drive-v1/permission-member/delete.md)
      - [Obtain a Collaborator List](server-docs/docs/permission/member/obtain-a-collaborator-list.md)
      - [Transfer Ownership](server-docs/docs/permission/member/transfer-ownership.md)
- **Setting v1**
      - [Get Document Sharing Settings V2](server-docs/docs/permission/setting-v1/get-document-sharing-settings-v2.md)
      - [Update document sharing settings](server-docs/drive-v1/permission-public/patch.md)

### Comment

    - [Get Document Comments in Pages](server-docs/drive-v1/file-comment/list.md)
    - [Batch Query Comments](server-docs/drive-v1/file-comment/batch_query.md)
    - [Solve or Restore a Comment](server-docs/drive-v1/file-comment/patch.md)
    - [Add a Global Comment](server-docs/drive-v1/file-comment/create.md)
    - [Get a Global Comment](server-docs/drive-v1/file-comment/get.md)
    - [Get Replies List](server-docs/drive-v1/file-comment-reply/list.md)
    - [Update Reply](server-docs/drive-v1/file-comment-reply/update.md)
    - [Delete Reply](server-docs/drive-v1/file-comment-reply/delete.md)

### Common

    - [Get docs content](server-docs/docs-v1/content/get.md)

### Appendix

    - [Error Codes](server-docs/docs/appendix/error-codes.md)
    - [Data Types Supported By Sheets](server-docs/docs/appendix/data-types-supported-by-sheets.md)
    - [Data Formats Supported by Sheets](server-docs/docs/appendix/data-formats-supported-by-sheets.md)

  - [Subscribe to Events](server-docs/docs/subscribe-to-events.md)

## Calendar

  - [Calendar overview](server-docs/calendar/calendar-overview.md)
  - [Calendar FAQ](server-docs/calendar-v4/calendar-faq.md)
### Calendar management

    - [Calendar resources introduction](server-docs/calendar-v4/calendar/introduction.md)
    - [Create a shared calendar](server-docs/calendar-v4/calendar/create.md)
    - [Delete shared calendar](server-docs/calendar-v4/calendar/delete.md)
    - [Query primary calendar information](server-docs/calendar-v4/calendar/primary.md)
    - [Query calendar information](server-docs/calendar-v4/calendar/get.md)
    - [Query availability of the primary calendar](server-docs/calendar-v4/freebusy/list.md)
    - [Query the calendar list](server-docs/calendar-v4/calendar/list.md)
    - [Update calendar information](server-docs/calendar-v4/calendar/patch.md)
    - [Search Calendar](server-docs/calendar-v4/calendar/search.md)
    - [Subscribe Calendar](server-docs/calendar-v4/calendar/subscribe.md)
    - [Unsubscribe Calendar](server-docs/calendar-v4/calendar/unsubscribe.md)
    - [Subscribe Calendar Changes ](server-docs/calendar-v4/calendar/subscription.md)
    - [Unsubscribe Calendar Changes](server-docs/calendar-v4/calendar/unsubscription.md)
- **Events**
      - [Calendar is changed](server-docs/calendar-v4/calendar/events/changed.md)

### Calendar access control

    - [Introduction to access control resources](server-docs/calendar-v4/calendar-acl/introduction.md)
    - [Create ACL](server-docs/calendar-v4/calendar-acl/create.md)
    - [Delete ACL](server-docs/calendar-v4/calendar-acl/delete.md)
    - [Obtain the ACL](server-docs/calendar-v4/calendar-acl/list.md)
    - [Subscribe ACL Event](server-docs/calendar-v4/calendar-acl/subscription.md)
    - [Unsubscribe ACL Event](server-docs/calendar-v4/calendar-acl/unsubscription.md)
- **Events**
      - [Create an ACL](server-docs/calendar-v4/calendar-acl/events/created.md)
      - [Remove an ACL](server-docs/calendar-v4/calendar-acl/events/deleted.md)

### Event management

    - [Event resource introduction](server-docs/calendar-v4/calendar-event/introduction.md)
    - [Create event](server-docs/calendar-v4/calendar-event/create.md)
    - [Delete Event](server-docs/calendar-v4/calendar-event/delete.md)
    - [Update Event](server-docs/calendar-v4/calendar-event/patch.md)
    - [Get Event](server-docs/calendar-v4/calendar-event/get.md)
    - [Get Event List](server-docs/calendar-v4/calendar-event/list.md)
    - [Search Event](server-docs/calendar-v4/calendar-event/search.md)
    - [Subscribe Event Changes](server-docs/calendar-v4/calendar-event/subscription.md)
    - [Unsubscribe Event Changes](server-docs/calendar-v4/calendar-event/unsubscription.md)
    - [Get Event instances](server-docs/calendar-v4/calendar-event/instances.md)
    - [Query event view](server-docs/calendar-v4/calendar-event/instance_view.md)
- **Events**
      - [Schedule is changed](server-docs/calendar-v4/calendar-event/events/changed.md)

### Meeting Chat

    - [Create meeting chat](server-docs/calendar-v4/calendar-event-meeting_chat/create.md)
    - [Unbind meeting chat](server-docs/calendar-v4/calendar-event-meeting_chat/delete.md)

### Meeting Minute

    - [Create meeting minute](server-docs/calendar-v4/calendar-event-meeting_minute/create.md)

### Timeoff

    - [Introduction to leave event resources](server-docs/calendar-v4/timeoff_event/introduction.md)
    - [Create Timeoff Event](server-docs/calendar-v4/timeoff_event/create.md)
    - [Delete Timeoff Event](server-docs/calendar-v4/timeoff_event/delete.md)

### Meeting room event

    - [Query meeting room event topics and meeting details](server-docs/calendar/meeting-room-event/query-meeting-room-event-topics-and-meeting-details.md)
    - [Query room availability](server-docs/calendar/meeting-room-event/query-room-availability.md)
    - [Reply meeting room event instance](server-docs/calendar/meeting-room-event/reply-meeting-room-event-instance.md)
- **Event**
      - [Meeting Room Status Changed](server-docs/meeting_room-v1/meeting_room/events/status_changed.md)
      - [Third-party meeting room event changes](server-docs/meeting_room-v1/event/third-room-event-changes.md)

### Event attendee (Including meeting room)

    - [Event invitee resource introduction](server-docs/calendar-v4/calendar-event-attendee/introduction.md)
    - [Create event invitees](server-docs/calendar-v4/calendar-event-attendee/create.md)
    - [Delete event invitees](server-docs/calendar-v4/calendar-event-attendee/batch_delete.md)
    - [Obtain event invitee list](server-docs/calendar-v4/calendar-event-attendee/list.md)
    - [Obtain the list of members of group invitees of an event](server-docs/calendar-v4/calendar-event-attendee-chat_member/list.md)

### Sync to local calendar

    - [Generate CalDAV configuration](server-docs/calendar-v4/setting/generate_caldav_conf.md)

### Synchronize Exchange calendar information

    - [Exchange resource introduction](server-docs/calendar-v4/exchange_binding/introduction.md)
    - [Bind Exchange account to Lark account](server-docs/calendar-v4/exchange_binding/create.md)
    - [Unbind the Exchange account](server-docs/calendar-v4/exchange_binding/delete.md)
    - [Query the binding status of the Exchange account](server-docs/calendar-v4/exchange_binding/get.md)


## Video Conferencing

  - [Video Conferencing overview](server-docs/vc-v1/video-conferencing-overview.md)
### Meeting reservation

    - [Resource introduction](server-docs/vc-v1/reserve/schedule-meeting-overview.md)
    - [Schedule a meeting](server-docs/vc-v1/reserve/apply.md)
    - [Delete a schedule](server-docs/vc-v1/reserve/delete.md)
    - [Update a schedule](server-docs/vc-v1/reserve/update.md)
    - [Obtain a schedule](server-docs/vc-v1/reserve/get.md)
    - [Obtain an active meeting](server-docs/vc-v1/reserve/get_active_meeting.md)

### Meeting management

    - [Meeting overview](server-docs/vc-v1/meeting/meeting-overview.md)
    - [Invite participants](server-docs/vc-v1/meeting/invite.md)
    - [Remove participants](server-docs/vc-v1/meeting/kickout.md)
    - [Set meeting host](server-docs/vc-v1/meeting/set_host.md)
    - [End a meeting](server-docs/vc-v1/meeting/end.md)
    - [Obtain meeting details](server-docs/vc-v1/meeting/get.md)
    - [List meetings of same meeting number](server-docs/vc-v1/meeting/list_by_no.md)
- **Events**
      - [Start of corporate meeting](server-docs/vc-v1/meeting/events/all_meeting_started.md)
      - [End of corporate meeting](server-docs/vc-v1/meeting/events/all_meeting_ended.md)
      - [Meeting started](server-docs/vc-v1/meeting/events/meeting_started.md)
      - [Meeting ended](server-docs/vc-v1/meeting/events/meeting_ended.md)
      - [Join meeting](server-docs/vc-v1/meeting/events/join_meeting.md)
      - [Leave meeting](server-docs/vc-v1/meeting/events/leave_meeting.md)
      - [Start recording](server-docs/vc-v1/meeting/events/recording_started.md)
      - [Stop recording](server-docs/vc-v1/meeting/events/recording_ended.md)
      - [Finished recording](server-docs/vc-v1/meeting/events/recording_ready.md)
      - [Share started](server-docs/vc-v1/meeting/events/share_started.md)
      - [Meeting record](server-docs/vc-v1/meeting/events/share_ended.md)

### Meeting record

    - [Recording overview](server-docs/vc-v1/meeting-recording/recording-overview.md)
    - [Start recording](server-docs/vc-v1/meeting-recording/start.md)
    - [Stop recording](server-docs/vc-v1/meeting-recording/stop.md)
    - [Obtain recording files](server-docs/vc-v1/meeting-recording/get.md)
    - [Authorize recording files](server-docs/vc-v1/meeting-recording/set_permission.md)

### Meeting report

    - [Resource introduction](server-docs/vc-v1/report/meeting-report-overview.md)
    - [Obtain meeting reports](server-docs/vc-v1/report/get_daily.md)
    - [Get top user list](server-docs/vc-v1/report/get_top_user.md)

### Export

    - [Resource introduction](server-docs/vc-v1/export/export-overview.md)
    - [Export meeting details](server-docs/vc-v1/export/meeting_list.md)
    - [Export participant details](server-docs/vc-v1/export/participant_list.md)
    - [Export participant meeting quality data](server-docs/vc-v1/export/participant_quality_list.md)
    - [Export meeting room reservation data](server-docs/vc-v1/export/resource_reservation_list.md)
    - [Query export task results](server-docs/vc-v1/export/get.md)
    - [Download export file](server-docs/vc-v1/export/download.md)

### Meeting room level

    - [Resource introduction](server-docs/vc-v1/room_level/room-level-overview.md)
    - [Create room level](server-docs/vc-v1/room_level/create.md)
    - [Delete room level](server-docs/vc-v1/room_level/del.md)
    - [Update room level](server-docs/vc-v1/room_level/patch.md)
    - [Query room level details](server-docs/vc-v1/room_level/get.md)
    - [Batch query room level details](server-docs/vc-v1/room_level/mget.md)
    - [Query room level list](server-docs/vc-v1/room_level/list.md)
    - [Search room level](server-docs/vc-v1/room_level/search.md)
- **Events**
      - [Create room level](server-docs/vc-v1/room_level/events/created.md)
      - [Delete room level](server-docs/vc-v1/room_level/events/deleted.md)
      - [Update room level](server-docs/vc-v1/room_level/events/updated.md)

### Meeting room management

    - [Resource introduction](server-docs/vc-v1/room/room-overview.md)
    - [Create meeting room](server-docs/vc-v1/room/create.md)
    - [Delete meeting room](server-docs/vc-v1/room/delete.md)
    - [Update meeting room](server-docs/vc-v1/room/patch.md)
    - [Query meeting room details](server-docs/vc-v1/room/get.md)
    - [Batch query meeting room details](server-docs/vc-v1/room/mget.md)
    - [Query meeting room list](server-docs/vc-v1/room/list.md)
    - [Search meeting room](server-docs/vc-v1/room/search.md)
- **Events**
      - [Create meeting room](server-docs/vc-v1/room/events/created.md)
      - [Update meeting room](server-docs/vc-v1/room/events/updated.md)
      - [Delete meeting room](server-docs/vc-v1/room/events/deleted.md)

### Meeting room configuration

    - [Resource introduction](server-docs/vc-v1/scope_config/room-configuration-overview.md)
    - [Query room configuration](server-docs/vc-v1/scope_config/get.md)
    - [Set room configuration](server-docs/vc-v1/scope_config/create.md)
    - [Get reserve limitation](server-docs/vc-v1/reserve_config/reserve_scope.md)
    - [Update reserve limitation](server-docs/vc-v1/reserve_config/patch.md)
    - [Get reserve form](server-docs/vc-v1/reserve_config-form/get.md)
    - [Update reserve form](server-docs/vc-v1/reserve_config-form/patch.md)
    - [Get reserve admin](server-docs/vc-v1/reserve_config-admin/get.md)
    - [Update reserve admin](server-docs/vc-v1/reserve_config-admin/patch.md)
- **Events**
      - [Update room reserve limitation](server-docs/vc-v1/reserve_config/events/updated.md)

### Meeting data

    - [Resource introduction](server-docs/vc-v1/meeting-room-data/resource-introduction.md)
    - [Get meeting details](server-docs/vc-v1/meeting_list/get.md)
    - [Get participant details](server-docs/vc-v1/participant_list/get.md)
    - [Get participant meeting quality data](server-docs/vc-v1/participant_quality_list/get.md)
    - [Get meeting room reservation data](server-docs/vc-v1/resource_reservation_list/get.md)

### Alert center

    - [Get alert center history](server-docs/vc-v1/alert/list.md)


## Minutes

### Minutes audio or video file

    - [Download minutes audio or video file](server-docs/minutes-v1/minute-media/get.md)

### Minutes transcript

    - [Export minutes transcript](server-docs/minutes-v1/minute-transcript/get.md)

### Minutes statistics

    - [Get minutes statistics](server-docs/minutes-v1/minute-statistics/get.md)

### Minutes Meta

    - [Get minutes meta](server-docs/minutes-v1/minute/get.md)


## Attendance

  - [Overview](server-docs/attendance-v1/overview.md)
  - [Access guide](server-docs/attendance-v1/attendance-development-guidelines.md)
### Attendance Shift

    - [Query Attendance Shift by Name](server-docs/attendance-v1/shift/query.md)
    - [Get Attendance Shift Details](server-docs/attendance-v1/shift/get.md)
    - [Delete Attendance Shift](server-docs/attendance-v1/shift/delete.md)
    - [Create Attendance Shift](server-docs/attendance-v1/shift/create.md)

### Attendance Group

    - [Create or modify an Attendance Group](server-docs/attendance-v1/group/create.md)
    - [Query Attendance Group by Name](server-docs/attendance-v1/group/search.md)
    - [Get Attendance Group Details](server-docs/attendance-v1/group/get.md)
    - [Delete Attendance Group](server-docs/attendance-v1/group/delete.md)

### Attendance Schedule

    - [Create or modify schedule](server-docs/attendance-v1/user_daily_shift/batch_create.md)
    - [Query Schedule Information](server-docs/attendance-v1/user_daily_shift/query.md)

### Attendance Statistics

    - [Attendance Statistics Development Guide](server-docs/attendance-v1/user_stats_data/attendance-statistic-reference.md)
    - [Update Statistics Settings](server-docs/attendance-v1/user_stats_view/update.md)
    - [Query Statistics Header](server-docs/attendance-v1/user_stats_field/query.md)
    - [Query Statistics Settings](server-docs/attendance-v1/user_stats_view/query.md)
    - [Query statistical data](server-docs/attendance-v1/user_stats_data/query.md)

### Attendance Approval

    - [Get User Attendance Data](server-docs/attendance/attendance-approval/get-user-attendance-data.md)
    - [Notify Approval Status Update](server-docs/attendance/attendance-approval/notify-approval-status-update.md)
    - [Add Approved Data in Lark Attendance](server-docs/attendance-v1/user_approval/add-approved-data-in-feishu-attendance2.md)

### Attendance Correction

    - [Get Allowed Correction Time](server-docs/attendance-v1/user_task_remedy/query_user_allowed_remedys.md)
    - [Get Correction Records](server-docs/attendance-v1/user_task_remedy/query.md)
    - [Notify of Correction Request Submission](server-docs/attendance-v1/user_task_remedy/create.md)

### Attendance Records

    - [Batch Query Attendance Records](server-docs/attendance-v1/user_flow/query.md)
    - [Get Attendance Records](server-docs/attendance-v1/user_flow/get.md)
    - [Get Attendance Result](server-docs/attendance-v1/user_task/query.md)
    - [Import Attendance Records](server-docs/attendance-v1/user_flow/batch_create.md)

### Event

    - [User task status change event](server-docs/attendance/event/user-task-status-change-event.md)

### Attendance（Historical Version）

    - [Attendance Development Guidelines](server-docs/attendance/attendancehistorical-version/attendance-development-guidelines.md)
- **API Reference**
  - **Rule**
        - [Download Files](server-docs/attendance/attendancehistorical-version/api-reference/rule/download-files.md)
        - [Batch Query the User Settings](server-docs/attendance/attendancehistorical-version/api-reference/rule/batch-query-the-user-settings.md)
        - [Modify User Settings](server-docs/attendance/attendancehistorical-version/api-reference/rule/modify-user-settings.md)
        - [Upload Files](server-docs/attendance/attendancehistorical-version/api-reference/rule/upload-files.md)
        - [Create or Modify Attendance Groups](server-docs/attendance/attendancehistorical-version/api-reference/rule/create-or-modify-attendance-groups.md)
        - [Delete Attendance Groups](server-docs/attendance/attendancehistorical-version/api-reference/rule/delete-attendance-groups.md)
        - [Get Attendance Group Details](server-docs/attendance/attendancehistorical-version/api-reference/rule/get-attendance-group-details.md)
        - [Create a Shift](server-docs/attendance/attendancehistorical-version/api-reference/rule/create-a-shift.md)
        - [Delete a Shift](server-docs/attendance/attendancehistorical-version/api-reference/rule/delete-a-shift.md)
        - [Get Shift Details](server-docs/attendance/attendancehistorical-version/api-reference/rule/get-shift-details.md)
        - [Get the Shift Information by the Shift Name](server-docs/attendance/attendancehistorical-version/api-reference/rule/get-the-shift-information-by-the-shift-name.md)
  - **Task**
        - [Query Statistics](server-docs/attendance/attendancehistorical-version/api-reference/task/query-statistics.md)
        - [Query the Statistics Table Header](server-docs/attendance/attendancehistorical-version/api-reference/task/query-the-statistics-table-header.md)
        - [Update Statistics Settings](server-docs/attendance/attendancehistorical-version/api-reference/task/update-statistics-settings.md)
        - [Query Statistics Configuration](server-docs/attendance/attendancehistorical-version/api-reference/task/query-statistics-configuration.md)
        - [Query the Schedule Information](server-docs/attendance/attendancehistorical-version/api-reference/task/query-the-schedule-information.md)
        - [Get Users’ Attendance Results](server-docs/attendance/attendancehistorical-version/api-reference/task/get-users-attendance-results.md)
        - [Batch Query the Attendance Flow Records](server-docs/attendance/attendancehistorical-version/api-reference/task/batch-query-the-attendance-flow-records.md)
        - [Import Attendance Flow Records](server-docs/attendance/attendancehistorical-version/api-reference/task/import-attendance-flow-records.md)
        - [Get Attendance Flow Records](server-docs/attendance/attendancehistorical-version/api-reference/task/get-attendance-flow-records.md)
        - [Get Users’ Correction Records](server-docs/attendance/attendancehistorical-version/api-reference/task/get-users-correction-records.md)
        - [Create or Modify a Schedule](server-docs/attendance/attendancehistorical-version/api-reference/task/create-or-modify-a-schedule.md)
        - [Get User Attendance Data](server-docs/attendance/attendancehistorical-version/api-reference/task/get-user-attendance-data.md)
        - [Add Approved Data in Lark Attendance](server-docs/attendance/attendancehistorical-version/api-reference/task/add-approved-data-in-lark-attendance.md)
        - [Get Users’ Allowed Correction Time](server-docs/attendance/attendancehistorical-version/api-reference/task/get-users-allowed-correction-time.md)
        - [Notify Correction Request Submission](server-docs/attendance/attendancehistorical-version/api-reference/task/notify-correction-request-submission.md)
        - [Notify Approval Status Update](server-docs/attendance/attendancehistorical-version/api-reference/task/notify-approval-status-update.md)


## Approval

  - [Overview](server-docs/approval-v4/approval-overview.md)
  - [FAQs](server-docs/approval-v4/approval-related-faqs.md)
### Access guide

    - [Approval access guide](server-docs/approval/access-guide/approval-access-guide.md)
    - [Native approval access guide](server-docs/approval/access-guide/native-approval-access-guide.md)
    - [Store app development guide](server-docs/approval/access-guide/store-app-development-guide.md)
- **Third-party approval integration**
      - [Three-party approval access guide](server-docs/approval/access-guide/thirdparty-approval-integration/threeparty-approval-access-guide.md)
      - [Preparation for third-party approval integration](server-docs/approval/access-guide/thirdparty-approval-integration/preparation-for-thirdparty-approval-integration.md)

### Approval definition

    - [Resource introduction](server-docs/approval-v4/approval/overview-of-approval-resources.md)
    - [Create an approval definition](server-docs/approval-v4/approval/create.md)
    - [View approval definitions](server-docs/approval-v4/approval/get.md)
    - [Approval definition form control parameters](server-docs/approval-v4/approval/approval-definition-form-control-parameters.md)
    - [Associate external options](server-docs/approval/approval-definition/associate-external-options.md)

### Approval instances

    - [Resource introduction](server-docs/approval-v4/instance/overview-approval-instance.md)
    - [Create an approval instance](server-docs/approval-v4/instance/create.md)
    - [Revoke approval instances](server-docs/approval-v4/instance/cancel.md)
    - [Notify someone of approval instances](server-docs/approval-v4/instance/cc.md)
    - [Preview approval instances](server-docs/approval/approval-instances/preview-approval-instances.md)
    - [Obtain details of an approval instance](server-docs/approval-v4/instance/get.md)
    - [Obtain IDs of multiple approval instances](server-docs/approval-v4/instance/list.md)
    - [Approval instance form control parameters](server-docs/approval-v4/instance/approval-instance-form-control-parameters.md)

### Approval tasks

    - [Resource introduction](server-docs/approval-v4/task/introduction.md)
    - [Agree to approval task](server-docs/approval-v4/task/approve.md)
    - [Refuse to approve tasks](server-docs/approval-v4/task/reject.md)
    - [Transfer the approval task](server-docs/approval-v4/task/transfer.md)
    - [Return approval task](server-docs/approval-v4/instance/specified_rollback.md)
    - [Add approvers](server-docs/approval/approval-tasks/add-approvers.md)
    - [Resubmit the task for approval](server-docs/approval-v4/task/resubmit.md)

### Native approval document

    - [Resource introduction](server-docs/approval-v4/file/overview.md)
    - [Upload files](server-docs/approval/native-approval-document/upload-files.md)

### Native approval comments

    - [Resource introduction](server-docs/approval-v4/instance-comment/overview.md)
    - [Create comments](server-docs/approval-v4/instance-comment/create.md)
    - [Delete comments](server-docs/approval-v4/instance-comment/delete.md)
    - [Clear comments](server-docs/approval-v4/instance-comment/remove.md)
    - [Receive comments](server-docs/approval-v4/instance-comment/list.md)

### Third-party approval definition

    - [Resource introduction](server-docs/approval-v4/external_approval/overview.md)
    - [Create a third-party approval definition](server-docs/approval-v4/external_approval/create.md)
    - [Quick approval callback](server-docs/approval/thirdparty-approval-definition/quick-approval-callback.md)

### Third-party approval instances

    - [Resource introduction](server-docs/approval-v4/external_instance/overview.md)
    - [Sync instances](server-docs/approval-v4/external_instance/create.md)
    - [Verify instances](server-docs/approval-v4/external_instance/check.md)

### Third-party approval tasks

    - [Resource introduction](server-docs/approval-v4/external_task/overview.md)
    - [Status of third-party approval tasks](server-docs/approval-v4/external_task/list.md)

### Approval Bot messages

    - [Send Bot messages](server-docs/approval/approval-bot-messages/send-bot-messages.md)
    - [Update Bot messages](server-docs/approval/approval-bot-messages/update-bot-messages.md)

### Approval search

    - [List of instances](server-docs/approval-v4/instance/query.md)
    - [List of cc information](server-docs/approval-v4/instance/search_cc.md)
    - [List of tasks](server-docs/approval-v4/task/search.md)
    - [Query the user's task list](server-docs/approval-v4/task/query.md)

### Approval Events

    - [Function introduction](server-docs/approval/approval-events/function-introduction.md)
    - [Subscription steps](server-docs/approval/approval-events/subscription-steps.md)
    - [FAQs](server-docs/approval-v4/approval-common-problem/approval-events-faq.md)
- **Event interface**
      - [Subscribe to approval events](server-docs/approval-v4/approval/subscribe.md)
      - [Unsubscribe from approval events](server-docs/approval-v4/approval/unsubscribe.md)
- **Common event**
      - [Approval cc status change](server-docs/approval/approval-events/common-event/approval-cc-status-change.md)
      - [Approval task status change](server-docs/approval/approval-events/common-event/approval-task-status-change.md)
      - [Approval instance status change](server-docs/approval/approval-events/common-event/approval-instance-status-change.md)
      - [Approval definition updates](server-docs/approval/approval-events/common-event/approval-definition-updates.md)
- **Special event**
      - [Approval for going out](server-docs/approval/approval-events/special-event/approval-for-going-out.md)
      - [Business trip](server-docs/approval/approval-events/special-event/business-trip.md)
      - [Correction](server-docs/approval/approval-events/special-event/correction.md)
      - [Shift swap](server-docs/approval/approval-events/special-event/shift-swap.md)
      - [Overtime](server-docs/approval/approval-events/special-event/overtime.md)
      - [Leave](server-docs/approval/approval-events/special-event/leave.md)
    - [General Approval events](server-docs/approval/approval-events/general-approval-events.md)

### Approval（history version）

- **v2**
  - **Lark native approval**
        - [Subscribe to an Approvals Event ](server-docs/approval/approvalhistory-version/v2/lark-native-approval/subscribe-to-an-approvals-event.md)
        - [Cancel a Subscription to an Approvals Event](server-docs/approval/approvalhistory-version/v2/lark-native-approval/cancel-a-subscription-to-an-approvals-event.md)
        - [Approval task return](server-docs/approval/approvalhistory-version/v2/lark-native-approval/approval-task-return.md)
        - [Obtain Single Approval Form](server-docs/approval/approvalhistory-version/v2/lark-native-approval/obtain-single-approval-form.md)
        - [Create Approval Instance](server-docs/approval/approvalhistory-version/v2/lark-native-approval/create-approval-instance.md)
        - [Obtain Single Approval Instance Details](server-docs/approval/approvalhistory-version/v2/lark-native-approval/obtain-single-approval-instance-details.md)
        - [Batch Obtain Approval Instance IDs](server-docs/approval/approvalhistory-version/v2/lark-native-approval/batch-obtain-approval-instance-ids.md)
        - [CC Instance](server-docs/approval/approvalhistory-version/v2/lark-native-approval/cc-instance.md)
        - [Approval Instance Cancel](server-docs/approval/approvalhistory-version/v2/lark-native-approval/approval-instance-cancel.md)
        - [Approval Task Approve](server-docs/approval/approvalhistory-version/v2/lark-native-approval/approval-task-approve.md)
        - [Approval Task Reject](server-docs/approval/approvalhistory-version/v2/lark-native-approval/approval-task-reject.md)
        - [Approval Task Transfer](server-docs/approval/approvalhistory-version/v2/lark-native-approval/approval-task-transfer.md)
  - **Third-party approval integration**
        - [External Approval Create](server-docs/approval/approvalhistory-version/v2/thirdparty-approval-integration/external-approval-create.md)
        - [External Approval Instance Create](server-docs/approval/approvalhistory-version/v2/thirdparty-approval-integration/external-approval-instance-create.md)
        - [External Approval Instance Check](server-docs/approval/approvalhistory-version/v2/thirdparty-approval-integration/external-approval-instance-check.md)
        - [Get Third-party Approval Status](server-docs/approval/approvalhistory-version/v2/thirdparty-approval-integration/get-thirdparty-approval-status.md)
  - **Lark Store app integration**
        - [Create an Approval Definition](server-docs/approval/approvalhistory-version/v2/lark-store-app-integration/create-an-approval-definition.md)
  - **Approval Search**
        - [Instance List Query](server-docs/approval/approvalhistory-version/v2/approval-search/instance-list-query.md)
        - [Cc List Query](server-docs/approval/approvalhistory-version/v2/approval-search/cc-list-query.md)
        - [Task List Query](server-docs/approval/approvalhistory-version/v2/approval-search/task-list-query.md)

  - [Approval Error Code](server-docs/approval/approval-error-code.md)

## Bot

  - [Obtain Bot Info](server-docs/bot/obtain-bot-info.md)
### Events

    - [Bot Menu Event](server-docs/application-v6/bot/events/menu.md)


## Tasks v2

  - [Overview](server-docs/task-v2/overview.md)
  - [FAQ](server-docs/task-v2/faq.md)
### Task

    - [Task Feature Overview](server-docs/task-v2/task/overview.md)
    - [Create Task](server-docs/task-v2/task/create.md)
    - [Get Task Details](server-docs/task-v2/task/get.md)
    - [Patch Task](server-docs/task-v2/task/patch.md)
    - [Delete Task](server-docs/task-v2/task/delete.md)
    - [Add members to task](server-docs/task-v2/task/add_members.md)
    - [Remove members from task](server-docs/task-v2/task/remove_members.md)
    - [List tasks](server-docs/task-v2/task/list.md)
    - [List tasklists of task](server-docs/task-v2/task/tasklists.md)
    - [Add task to tasklist](server-docs/task-v2/task/add_tasklist.md)
    - [Remove task from tasklist](server-docs/task-v2/task/remove_tasklist.md)
    - [Add reminders to task](server-docs/task-v2/task/add_reminders.md)
    - [Remove task reminders](server-docs/task-v2/task/remove_reminders.md)
    - [Add Dependency](server-docs/task-v2/task/add_dependencies.md)
    - [Remove Dependency](server-docs/task-v2/task/remove_dependencies.md)

### Subtask

    - [Create Subtask](server-docs/task-v2/task-subtask/create.md)
    - [List Subtask](server-docs/task-v2/task-subtask/list.md)

### Tasklist

    - [Tasklist Features Overview](server-docs/task-v2/tasklist/overview.md)
    - [Create Tasklist](server-docs/task-v2/tasklist/create.md)
    - [Get Tasklist Details](server-docs/task-v2/tasklist/get.md)
    - [Patch Tasklist](server-docs/task-v2/tasklist/patch.md)
    - [Delete Tasklist](server-docs/task-v2/tasklist/delete.md)
    - [Add Tasklist Members](server-docs/task-v2/tasklist/add_members.md)
    - [Remove Tasklist Members](server-docs/task-v2/tasklist/remove_members.md)
    - [Get Tasks of Tasklist](server-docs/task-v2/tasklist/tasks.md)
    - [List Tasklists](server-docs/task-v2/tasklist/list.md)

### Tasklist Activity Subscription

    - [Create Activity Subscription](server-docs/task-v2/tasklist-activity_subscription/create.md)
    - [Get Activity Subscription](server-docs/task-v2/tasklist-activity_subscription/get.md)
    - [List Activity Subscription](server-docs/task-v2/tasklist-activity_subscription/list.md)
    - [Patch Activity Subscription](server-docs/task-v2/tasklist-activity_subscription/patch.md)
    - [Delete Activity Subscription](server-docs/task-v2/tasklist-activity_subscription/delete.md)

### Comment

    - [overview](server-docs/task-v2/comment/overview.md)
    - [Create Comment](server-docs/task-v2/comment/create.md)
    - [Get Comment](server-docs/task-v2/comment/get.md)
    - [Patch Comment](server-docs/task-v2/comment/patch.md)
    - [Delete Comment](server-docs/task-v2/comment/delete.md)
    - [List Comments](server-docs/task-v2/comment/list.md)

### attachment

    - [Attachment Feature Overview](server-docs/task-v2/attachment/attachment-feature-overview.md)
    - [Upload Attachment](server-docs/task-v2/attachment/upload.md)
    - [List Attachment](server-docs/task-v2/attachment/list.md)
    - [Get Attachment](server-docs/task-v2/attachment/get.md)
    - [Delete Attachment](server-docs/task-v2/attachment/delete.md)

### Section

    - [Section Feature Overview](server-docs/task-v2/section/section-feature-overview.md)
    - [Create Section](server-docs/task-v2/section/create.md)
    - [Get Section](server-docs/task-v2/section/get.md)
    - [Patch Section](server-docs/task-v2/section/patch.md)
    - [Delete Section](server-docs/task-v2/section/delete.md)
    - [List Section](server-docs/task-v2/section/list.md)
    - [List Tasks of Section](server-docs/task-v2/section/tasks.md)

### Custom Field

    - [Custom Field Overview](server-docs/task-v2/custom_field/custom-field-overview.md)
    - [Create Custom Field](server-docs/task-v2/custom_field/create.md)
    - [Get Custom Field](server-docs/task-v2/custom_field/get.md)
    - [Patch Custom Field](server-docs/task-v2/custom_field/patch.md)
    - [List Custom Field](server-docs/task-v2/custom_field/list.md)
    - [Add Custom Field To Resource](server-docs/task-v2/custom_field/add.md)
    - [Remove Custom Field From Resource](server-docs/task-v2/custom_field/remove.md)

### Custom Field Option

    - [Create Custom Field Option](server-docs/task-v2/custom_field-option/create.md)
    - [Patch Custom Field Option](server-docs/task-v2/custom_field-option/patch.md)


## Email

### Email Contacts

    - [Create Email Contact](server-docs/mail-v1/user_mailbox-mail_contact/create.md)
    - [Delete Email Contact](server-docs/mail-v1/user_mailbox-mail_contact/delete.md)
    - [Modify Email Contact's Info](server-docs/mail-v1/user_mailbox-mail_contact/patch.md)
    - [List Email Contacts](server-docs/mail-v1/user_mailbox-mail_contact/list.md)

### User Mailbox Alias

    - [Release Address From Recycle Bin](server-docs/mail-v1/user_mailbox/delete.md)
    - [Create A Member's Email Alias](server-docs/mail-v1/user_mailbox-alias/create.md)
    - [Delete A Member's Email Alias](server-docs/mail-v1/user_mailbox-alias/delete.md)
    - [Obtain All Member's Email Aliases](server-docs/mail-v1/user_mailbox-alias/list.md)
    - [Overview](server-docs/mail-v1/user_mailbox-alias/overview.md)

### email address

    - [Email Address Query](server-docs/mail-v1/user/query.md)

### Mail Folder

    - [Create Email Folder](server-docs/mail-v1/user_mailbox-folder/create.md)
    - [Delete Email Folder](server-docs/mail-v1/user_mailbox-folder/delete.md)
    - [Update Email Folder](server-docs/mail-v1/user_mailbox-folder/patch.md)
    - [List Email Folders](server-docs/mail-v1/user_mailbox-folder/list.md)

### User Message

    - [list mail of card](server-docs/mail-v1/user_mailbox-message/get_by_card.md)
    - [List Emails](server-docs/mail-v1/user_mailbox-message/list.md)
    - [Get Email Details](server-docs/mail-v1/user_mailbox-message/get.md)
    - [Send Message](server-docs/mail-v1/user_mailbox-message/send.md)
- **Attachment**
      - [Get Attachment Download Links](server-docs/mail-v1/user_mailbox-message-attachment/download_url.md)

### Event

    - [Subscribe Event](server-docs/mail-v1/user_mailbox-event/subscribe.md)
    - [Get Subscription Status](server-docs/mail-v1/user_mailbox-event/subscription.md)
    - [Cancel Subscribe](server-docs/mail-v1/user_mailbox-event/unsubscribe.md)
- **Event**
      - [message received](server-docs/mail-v1/user_mailbox-event/events/message_received.md)

### Auto Filter

    - [Create Auto FIlter](server-docs/mail-v1/user_mailbox-rule/create.md)
    - [Delete Auto Filter](server-docs/mail-v1/user_mailbox-rule/delete.md)
    - [Update Auto FIlter](server-docs/mail-v1/user_mailbox-rule/update.md)
    - [List Auto FIlters](server-docs/mail-v1/user_mailbox-rule/list.md)
    - [Reorder Auto Filters](server-docs/mail-v1/user_mailbox-rule/reorder.md)

### Mail Group

- **Mail Group Management**
      - [Create Mail Group](server-docs/mail-v1/mailgroup/create.md)
      - [Delete Mail Group](server-docs/mail-v1/mailgroup/delete.md)
      - [Modify Some Information Of Mail Group](server-docs/mail-v1/mailgroup/patch.md)
      - [Modify All Information Of Mail Group](server-docs/mail-v1/mailgroup/update.md)
      - [Query The Specified Mail Group](server-docs/mail-v1/mailgroup/get.md)
      - [Obtain Mailing Lists In Batch](server-docs/mail-v1/mailgroup/list.md)
      - [Overview](server-docs/mail-v1/mailgroup/overview.md)
- **Mail Group Manager**
      - [Create Mailing List Managers In Batch](server-docs/mail-v1/mailgroup-manager/batch_create.md)
      - [Batch Delete Mail Group Managers](server-docs/mail-v1/mailgroup-manager/batch_delete.md)
      - [Obtain Mailing List Managers In Batch](server-docs/mail-v1/mailgroup-manager/list.md)
- **Mail Group Member**
      - [Create A Mailing List Member](server-docs/mail-v1/mailgroup-member/create.md)
      - [Delete A Mailing List Member](server-docs/mail-v1/mailgroup-member/delete.md)
      - [Obtain Mailing List Member Information](server-docs/mail-v1/mailgroup-member/get.md)
      - [Obtain Mailing List Members In Batch](server-docs/mail-v1/mailgroup-member/list.md)
      - [Batch Create Mail Group Members](server-docs/mail-v1/mailgroup-member/batch_create.md)
      - [Batch Delete Mail Group Members](server-docs/mail-v1/mailgroup-member/batch_delete.md)
      - [Overview](server-docs/mail-v1/mailgroup-member/overview.md)
- **Mail Group Alias**
      - [Create A Mailing List Alias](server-docs/mail-v1/mailgroup-alias/create.md)
      - [Delete A Mailing List Alias](server-docs/mail-v1/mailgroup-alias/delete.md)
      - [Obtain All Mailing List Aliases](server-docs/mail-v1/mailgroup-alias/list.md)
      - [Overview](server-docs/mail-v1/mailgroup-alias/overview.md)
- **Mail Group Permission Member**
      - [Create A Mail Group Permission Member](server-docs/mail-v1/mailgroup-permission_member/create.md)
      - [Delete A Mail Group Permission Member](server-docs/mail-v1/mailgroup-permission_member/delete.md)
      - [Get A Mail Group Permission Member](server-docs/mail-v1/mailgroup-permission_member/get.md)
      - [List Permission Members Of A Mail Group](server-docs/mail-v1/mailgroup-permission_member/list.md)
      - [Batch Create Mail Group Permission Members](server-docs/mail-v1/mailgroup-permission_member/batch_create.md)
      - [Batch Delete Mail Group Permission Members](server-docs/mail-v1/mailgroup-permission_member/batch_delete.md)
      - [Overview](server-docs/mail-v1/mailgroup-permission_member/overview.md)

### Public Mailbox

- **Public Mailbox Management**
      - [Create A Public Mailbox](server-docs/mail-v1/public_mailbox/create.md)
      - [Modify Some Information Of Public Mailbox](server-docs/mail-v1/public_mailbox/patch.md)
      - [Modify All Information Of Public Mailbox](server-docs/mail-v1/public_mailbox/update.md)
      - [Query The Specified Public Mailbox](server-docs/mail-v1/public_mailbox/get.md)
      - [Check All Public Mailboxes](server-docs/mail-v1/public_mailbox/list.md)
      - [Move Public Mailbox To The Recycle Bin](server-docs/mail-v1/public_mailbox/remove_to_recycle_bin.md)
      - [Permanently Delete Public Mailbox Address](server-docs/mail-v1/public_mailbox/delete.md)
      - [Overview](server-docs/mail-v1/public_mailbox/overview.md)
- **Public Mailbox Member**
      - [Create A Public Mailbox Member](server-docs/mail-v1/public_mailbox-member/create.md)
      - [Delete A Public Mailbox Member](server-docs/mail-v1/public_mailbox-member/delete.md)
      - [Clear Public Mailbox Members](server-docs/mail-v1/public_mailbox-member/clear.md)
      - [Get A Public Mailbox Member](server-docs/mail-v1/public_mailbox-member/get.md)
      - [List Public Mailbox Members](server-docs/mail-v1/public_mailbox-member/list.md)
      - [Batch Create Public Mailbox Members](server-docs/mail-v1/public_mailbox-member/batch_create.md)
      - [Batch Delete Public Mailbox Members](server-docs/mail-v1/public_mailbox-member/batch_delete.md)
      - [Overview](server-docs/mail-v1/public_mailbox-member/overview.md)
- **Public Mailbox Alias**
      - [Create A Public Mailbox Alias](server-docs/mail-v1/public_mailbox-alias/create.md)
      - [Delete A Public Mailbox Alias](server-docs/mail-v1/public_mailbox-alias/delete.md)
      - [Obtain All Public Mailbox Aliases](server-docs/mail-v1/public_mailbox-alias/list.md)
      - [Overview](server-docs/mail-v1/public_mailbox-alias/overview.md)


## App Information

### Application

    - [Get Application Information](server-docs/application-v6/application/get.md)
    - [Get App Version Information](server-docs/application-v6/application-app_version/get.md)
    - [Obtain App Version List](server-docs/application-v6/application-app_version/list.md)
    - [Get the Range of Contacts Data to Access in an App's Version Release Request](server-docs/application-v6/application-app_version/contacts_range_suggest.md)

### Permissions & Scopes

    - [Apply for scopes from the admin](server-docs/application-v6/scope/apply.md)
    - [Query tenant authorization status](server-docs/application-v6/scope/list.md)

### Admin

    - [Obtain the Apps Installed by an Organization](server-docs/app-information/admin/obtain-the-apps-installed-by-an-organization.md)
    - [Obtain the Apps Available to a User](server-docs/app-information/admin/obtain-the-apps-available-to-a-user.md)
    - [Get Application Release Request List](server-docs/application-v6/application/underauditlist.md)
    - [Update version information](server-docs/application-v6/application-app_version/patch.md)
    - [Update application information](server-docs/application-v6/application/patch.md)
    - [Get the Range of Contacts Data an App Can Access](server-docs/application-v6/application/contacts_range_configuration.md)
    - [Update the contacts permission scope an app can access](server-docs/application-v6/application-contacts_range/patch.md)
    - [Obtain the App Availability in an Organization](server-docs/app-information/admin/obtain-the-app-availability-in-an-organization.md)
    - [Obtain whether user in app visility white or black list](server-docs/application-v6/application-visibility/check_white_black_list.md)
    - [Update app availability](server-docs/application-v6/application-visibility/patch.md)
    - [Enable or disable application](server-docs/application-v6/application-management/update.md)
    - [Verify App Admin](server-docs/app-information/admin/verify-app-admin.md)
    - [Obtain an App Admin’s Management Permissions](server-docs/app-information/admin/obtain-an-app-admins-management-permissions.md)
    - [Search App Admin List](server-docs/app-information/admin/search-app-admin-list.md)

### Appstore Paid Info

    - [Query a User's App Access](server-docs/app-information/appstore-paid-info/query-a-users-app-access.md)
    - [Query an App Tenant’s Paid Orders](server-docs/app-information/appstore-paid-info/query-an-app-tenants-paid-orders.md)
    - [Query Order Information](server-docs/app-information/appstore-paid-info/query-order-information.md)

### Event

    - [App Created](server-docs/application-v6/application/events/created.md)
    - [App First Enabled](server-docs/application-v6/event/app-first-enabled.md)
    - [App Enabled or Disabled](server-docs/application-v6/event/app-enabled-or-disabled.md)
    - [Public App Purchase](server-docs/application-v6/event/public-app-purchase.md)
    - [app_ticket Events](server-docs/application-v6/event/app_ticket-events.md)
    - [App Uninstalled](server-docs/application-v6/event/app-uninstalled.md)
    - [App availability scope extended](server-docs/application-v6/event/app-availability-scope-extended.md)
    - [Apply To Publish an Application](server-docs/application-v6/application-app_version/events/publish_apply.md)
    - [Withdraw Application Release Application](server-docs/application-v6/application-app_version/events/publish_revoke.md)
    - [App has been reviewed](server-docs/application-v6/application-app_version/events/audit.md)


## Company Information

### Tenant Product Assign Info

    - [Obtain company assign information](server-docs/tenant-v2/tenant-product_assign_info/query.md)

### Tenant

    - [Get Tenant Information](server-docs/tenant-v2/tenant/query.md)


## Personal Settings

### System status

    - [Function introduction](server-docs/personal_settings-v1/system_status/overview.md)
    - [Create system status](server-docs/personal_settings-v1/system_status/create.md)
    - [Delete system status](server-docs/personal_settings-v1/system_status/delete.md)
    - [Patch system status](server-docs/personal_settings-v1/system_status/patch.md)
    - [List system status](server-docs/personal_settings-v1/system_status/list.md)
    - [Batch open system status](server-docs/personal_settings-v1/system_status/batch_open.md)
    - [Batch close system status](server-docs/personal_settings-v1/system_status/batch_close.md)


## Search

### docs search

    - [search docs](server-docs/search-v2/doc_wiki/search.md)

### Suite Search

    - [search message](server-docs/search-v2/message/create.md)


## AI

### Document AI

- **Recognize  Japanese Driving License**
      - [Identify Japanese driving license in documents](server-docs/document_ai-v1/jp_driving_license/recognize.md)
- **Recognize invoice**
      - [recognize invoice](server-docs/document_ai-v1/invoice/recognize.md)
- **Business card recognition**
      - [Recognize business cards in pictures](server-docs/document_ai-v1/business_card/recognize.md)

### Optical Character Recognition

- **Image Recognition**
      - [Basic Recognition](server-docs/optical_char_recognition-v1/image/basic_recognize.md)

### Speech to Text

- **Speech Recognition**
      - [Stream Recognition](server-docs/speech_to_text-v1/speech/stream_recognize.md)
      - [File Recognition](server-docs/speech_to_text-v1/speech/file_recognize.md)

### Machine Translation

- **Text**
      - [Text Translation](server-docs/translation-v1/text/translate.md)
      - [Language Recognition](server-docs/translation-v1/text/detect.md)


## Admin

### password

    - [reset user enterprise email password](server-docs/admin-v1/password/reset.md)

### Data report management

    - [Obtain active users and function usage data in the department dimension](server-docs/admin-v1/admin_dept_stat/list.md)
    - [Obtain active users and function usage data in the user dimension](server-docs/admin-v1/admin_user_stat/list.md)

### Enterprise badge

    - [Function introduction](server-docs/admin-v1/badge/quick-start.md)
- **Badge management**
      - [Create a badge](server-docs/admin-v1/badge/create.md)
      - [Modify the badge information](server-docs/admin-v1/badge/update.md)
      - [Upload a badge image](server-docs/admin-v1/badge_image/create.md)
      - [Get a list of badges](server-docs/admin-v1/badge/list.md)
      - [Get the badge detail](server-docs/admin-v1/badge/get.md)
- **Badge grant**
      - [Create a grant list](server-docs/admin-v1/badge-grant/create.md)
      - [Delete the grant list](server-docs/admin-v1/badge-grant/delete.md)
      - [Modify the grant list](server-docs/admin-v1/badge-grant/update.md)
      - [List of grants to receive badge](server-docs/admin-v1/badge-grant/list.md)
      - [Get the grant list detail](server-docs/admin-v1/badge-grant/get.md)

### Data Report of Department Dimension

    - [Overview](server-docs/admin-v1/admin_dept_stat/overview.md)

### Data Report of User Dimension

    - [Overview](server-docs/admin-v1/admin_user_stat/overview.md)


## Moments

  - [Post and comment content format conversion](server-docs/moments-v1/moments-explanation.md)
### Post

    - [Query post information](server-docs/moments-v1/post/get.md)
- **Events**
      - [Moment posted](server-docs/moments-v1/post/events/created.md)
      - [Moment deleted](server-docs/moments-v1/post/events/deleted.md)

### Comment

- **Events**
      - [Comment posted](server-docs/moments-v1/comment/events/created.md)
      - [Comment deleted](server-docs/moments-v1/comment/events/deleted.md)

### Reaction

- **Events**
      - [Emoji reaction added](server-docs/moments-v1/reaction/events/created.md)
      - [Emoji reaction canceled](server-docs/moments-v1/reaction/events/deleted.md)

### Post Statistics

- **Events**
      - [Post statistics changed](server-docs/moments-v1/post_statistics/events/updated.md)


## OKR

  - [OKR Development Guide](server-docs/okr-v1/okr-guide.md)
### OKR Period

    - [Get OKR Period List](server-docs/okr-v1/period/list.md)

### OKR

    - [Batch Get OKR](server-docs/okr-v1/okr/batch_get.md)

### User OKR

    - [Get User OKR List](server-docs/okr-v1/user-okr/list.md)

### progress record

    - [Delete OKR progress](server-docs/okr-v1/progress_record/delete.md)
    - [Update OKR progress](server-docs/okr-v1/progress_record/update.md)
    - [Get OKR progress records](server-docs/okr-v1/progress_record/get.md)
    - [Create OKR progress record](server-docs/okr-v1/progress_record/create.md)

### image

    - [Upload progress record image](server-docs/okr-v1/image/upload.md)


## security_and_compliance

### OpenAPI Audit Log

    - [Obtain OpenAPI audit log](server-docs/security_and_compliance-v1/openapi_log/list_data.md)

### app dlp execute log export

    - [export app dlp execute log](server-docs/security_and_compliance-v1/app_dlp_execute_log/list.md)

### Behavior audit log

    - [Function introduction](server-docs/security-and-compliance/behavior-audit-log/function-introduction.md)
    - [Get behavior audit log](server-docs/security-and-compliance/behavior-audit-log/get-behavior-audit-log.md)
    - [List of enumerated values](server-docs/security-and-compliance/behavior-audit-log/list-of-enumerated-values.md)

### Data Residency

    - [Overview](server-docs/security_and_compliance-v1/user_migration/overview.md)
    - [Get a list of geographic locations](server-docs/security_and_compliance-v1/multi_geo_entity-tenant/get.md)
    - [Migrate users](server-docs/security_and_compliance-v1/user_migration/create.md)
    - [Get user migration status](server-docs/security_and_compliance-v1/user_migration/get.md)
    - [Get user migration status in batches](server-docs/security_and_compliance-v1/user_migration/search.md)
    - [Cancel user migration](server-docs/security_and_compliance-v1/user_migration/cancel.md)


## Workplace

### Workplace Access Data

    - [Search Workplace Access Data](server-docs/workplace-v1/workplace_access_data/search.md)
    - [Get Custom Workplace Access Data](server-docs/workplace-v1/custom_workplace_access_data/search.md)
    - [Get Block  Access Data](server-docs/workplace-v1/workplace_block_access_data/search.md)


## Feishu Master Data Management

### Common Data

- **country region**
      - [Resource Definition](server-docs/mdm-v3/country_region/resource-definition.md)
      - [batch get major by id](server-docs/mdm-v3/batch_country_region/get.md)
      - [Pagination Batch Query Country Region](server-docs/mdm-v3/country_region/list.md)


## Deprecated Version (Not Recommended)

### authen

    - [Get user_access_token](server-docs/authen-v1/oidc-access_token/create.md)
    - [Refresh user_access_token](server-docs/authen-v1/oidc-refresh_access_token/create.md)
    - [Get user_access_token (Bytedance only)](server-docs/deprecated-version-not-recommended/authen/get-user-access-token-bytedance-only.md)
    - [Refresh user_access_token (Bytedance only)](server-docs/deprecated-version-not-recommended/authen/refresh-user-access-token-bytedance-only.md)
    - [Obtain code](server-docs/deprecated-version-not-recommended/authen/obtain-code.md)
    - [Obtain user_access_token (v1)](server-docs/authen-v1/access_token/create.md)
    - [code2session](server-docs/deprecated-version-not-recommended/authen/code2session.md)
    - [Refresh user_access_token (v1)](server-docs/authen-v1/refresh_access_token/create.md)
    - [Obtain user_access_token (v1, ByteDance only)](server-docs/historic-version/authen/obtain-user_access_token-bytedance.md)
    - [Refresh user_access_token (v1, ByteDance only)](server-docs/historic-version/authen/refresh-user_access_token-bytedance.md)

### App Information

- **Admin**
      - [Update the Availability of an App](server-docs/deprecated-version-not-recommended/app-information/admin/update-the-availability-of-an-app.md)

### Contact

- **User**
      - [Get User List](server-docs/contact-v3/user/list.md)
      - [Batch Obtain User Information](server-docs/deprecated-version-not-recommended/contact/user/batch-obtain-user-information.md)
      - [Obtain a Department User List](server-docs/deprecated-version-not-recommended/contact/user/obtain-a-department-user-list.md)
      - [Obtain Department User Details](server-docs/deprecated-version-not-recommended/contact/user/obtain-department-user-details.md)
      - [Add a User](server-docs/deprecated-version-not-recommended/contact/user/add-a-user.md)
      - [Delete a User](server-docs/deprecated-version-not-recommended/contact/user/delete-a-user.md)
      - [Update a User’s Information](server-docs/deprecated-version-not-recommended/contact/user/update-a-users-information.md)
      - [Obtain a Role List](server-docs/deprecated-version-not-recommended/contact/user/obtain-a-role-list.md)
      - [Obtain a List of Role Members](server-docs/deprecated-version-not-recommended/contact/user/obtain-a-list-of-role-members.md)
      - [Get User IDs Using Email Addresses or Mobile Numbers](server-docs/deprecated-version-not-recommended/contact/user/get-user-ids-using-email-addresses-or-mobile-numbers.md)
      - [Get User IDs Using Union IDs](server-docs/deprecated-version-not-recommended/contact/user/get-user-ids-using-union-ids.md)
- **Department**
      - [Obtain Department Details](server-docs/deprecated-version-not-recommended/contact/department/obtain-department-details.md)
      - [Get Department Information List](server-docs/contact-v3/department/list.md)
      - [Obtain a Sub-department ID List](server-docs/deprecated-version-not-recommended/contact/department/obtain-a-subdepartment-id-list.md)
      - [Obtain a Sub-department List](server-docs/deprecated-version-not-recommended/contact/department/obtain-a-subdepartment-list.md)
      - [Batch Obtain Department Details](server-docs/deprecated-version-not-recommended/contact/department/batch-obtain-department-details.md)
      - [Add a Department](server-docs/deprecated-version-not-recommended/contact/department/add-a-department.md)
      - [Delete a Department](server-docs/deprecated-version-not-recommended/contact/department/delete-a-department.md)
      - [Update Department Information](server-docs/deprecated-version-not-recommended/contact/department/update-department-information.md)
- **Import API**
      - [Batch Add Departments](server-docs/deprecated-version-not-recommended/contact/import-api/batch-add-departments.md)
      - [Batch Add Users](server-docs/deprecated-version-not-recommended/contact/import-api/batch-add-users.md)
      - [Query the Execution Status of a Batch Task](server-docs/deprecated-version-not-recommended/contact/import-api/query-the-execution-status-of-a-batch-task.md)
    - [Obtain Custom User Properties](server-docs/deprecated-version-not-recommended/contact/obtain-custom-user-properties.md)
    - [Obtain the Contacts Permission Scope](server-docs/deprecated-version-not-recommended/contact/obtain-the-contacts-permission-scope.md)
- **event**
      - [user status changed](server-docs/deprecated-version-not-recommended/contact/event/user-status-changed.md)
      - [employee change](server-docs/deprecated-version-not-recommended/contact/event/employee-change.md)
      - [department update](server-docs/deprecated-version-not-recommended/contact/event/department-update.md)
      - [scope change](server-docs/deprecated-version-not-recommended/contact/event/scope-change.md)
    - [Contacts Events](server-docs/deprecated-version-not-recommended/contact/contacts-events.md)

### Docs

- **Upgraded Docs**
      - [Guide](server-docs/deprecated-version-not-recommended/docs/upgraded-docs/guide.md)
      - [Data Structure](server-docs/deprecated-version-not-recommended/docs/upgraded-docs/data-structure.md)
- **Base**
      - [Record filter development guide](server-docs/bitable-v1/filter.md)
      - [Get records](server-docs/bitable-v1/app-table-record/get.md)
      - [List records](server-docs/bitable-v1/app-table-record/list.md)
- **Docs**
      - [Overview](server-docs/deprecated-version-not-recommended/docs/docs/overview.md)
      - [Getting Started with Docs API](server-docs/deprecated-version-not-recommended/docs/docs/getting-started-with-docs-api.md)
      - [Document Structure Introduction](server-docs/deprecated-version-not-recommended/docs/docs/document-structure-introduction.md)
      - [Document Structure Reference](server-docs/deprecated-version-not-recommended/docs/docs/document-structure-reference.md)
  - **Docs**
        - [Create Document](server-docs/deprecated-version-not-recommended/docs/docs/docs/create-document.md)
        - [Obtain Document Meta](server-docs/deprecated-version-not-recommended/docs/docs/docs/obtain-document-meta.md)
        - [Obtain sheet meta info in doc](server-docs/deprecated-version-not-recommended/docs/docs/docs/obtain-sheet-meta-info-in-doc.md)
  - **Content**
        - [Obtain Document Content](server-docs/deprecated-version-not-recommended/docs/docs/content/obtain-document-content.md)
        - [Get Document](server-docs/deprecated-version-not-recommended/docs/docs/content/get-document.md)
        - [Batch Update Document](server-docs/deprecated-version-not-recommended/docs/docs/content/batch-update-document.md)
- **sheets**
  - **Sheet Operation**
        - [Import Spreadsheet](server-docs/deprecated-version-not-recommended/docs/sheets/sheet-operation/import-spreadsheet.md)
        - [Query Import Results](server-docs/deprecated-version-not-recommended/docs/sheets/sheet-operation/query-import-results.md)
  - **Data Validation**
        - [Query Drop-down List in the Specific Range](server-docs/historic-version/docs/sheets/datavalidation/query-datavalidation.md)
        - [Update Drop-down List Setting](server-docs/historic-version/docs/sheets/datavalidation/update-datavalidation.md)
        - [Remove the Drop-down List in the Specific Range](server-docs/historic-version/docs/sheets/datavalidation/delete-datavalidation.md)
- **Drive**
  - **permission**
        - [Querying if a Collaborator Has a Specific Permission](server-docs/deprecated-version-not-recommended/docs/drive/permission/querying-if-a-collaborator-has-a-specific-permission.md)
  - **Folder**
        - [Create a New Folder](server-docs/deprecated-version-not-recommended/docs/drive/folder/create-a-new-folder.md)
        - [Get Folder Children](server-docs/deprecated-version-not-recommended/docs/drive/folder/get-folder-children.md)
  - **File**
        - [Copy a Doc or Sheet](server-docs/deprecated-version-not-recommended/docs/drive/file/copy-a-doc-or-sheet.md)
        - [Delete a Doc](server-docs/deprecated-version-not-recommended/docs/drive/file/delete-a-doc.md)
        - [Delete a Sheet](server-docs/deprecated-version-not-recommended/docs/drive/file/delete-a-sheet.md)
        - [Obtain Metadata](server-docs/deprecated-version-not-recommended/docs/drive/file/obtain-metadata.md)
  - **Comment**
        - [Add a Whole Document Comment](server-docs/deprecated-version-not-recommended/docs/drive/comment/add-a-whole-document-comment.md)

### Rooms

- **API Reference**
      - [Query Meeting Room Customization Setting](server-docs/deprecated-version-not-recommended/rooms/api-reference/query-meeting-room-customization-setting.md)
      - [Obtain Building List](server-docs/deprecated-version-not-recommended/rooms/api-reference/obtain-building-list.md)
      - [Query building details](server-docs/deprecated-version-not-recommended/rooms/api-reference/query-building-details.md)
      - [Obtain meeting room list](server-docs/deprecated-version-not-recommended/rooms/api-reference/obtain-meeting-room-list.md)
      - [Query Meeting Room Details](server-docs/deprecated-version-not-recommended/rooms/api-reference/query-meeting-room-details.md)
      - [Create Building](server-docs/deprecated-version-not-recommended/rooms/api-reference/create-building.md)
      - [Update building](server-docs/deprecated-version-not-recommended/rooms/api-reference/update-building.md)
      - [Delete Building](server-docs/deprecated-version-not-recommended/rooms/api-reference/delete-building.md)
      - [Obtain Building ID](server-docs/deprecated-version-not-recommended/rooms/api-reference/obtain-building-id.md)
      - [Create Meeting Room](server-docs/deprecated-version-not-recommended/rooms/api-reference/create-meeting-room.md)
      - [Update Meeting Room](server-docs/deprecated-version-not-recommended/rooms/api-reference/update-meeting-room.md)
      - [Delete Meeting Room](server-docs/deprecated-version-not-recommended/rooms/api-reference/delete-meeting-room.md)
      - [Obtain Meeting Room ID](server-docs/deprecated-version-not-recommended/rooms/api-reference/obtain-meeting-room-id.md)
      - [Obtain Country/Region List](server-docs/deprecated-version-not-recommended/rooms/api-reference/obtain-country-region-list.md)
      - [Obtain City List](server-docs/deprecated-version-not-recommended/rooms/api-reference/obtain-city-list.md)
- **Event**
      - [Meeting room created](server-docs/meeting_room-v1/meeting_room/events/created.md)
      - [Meeting Room Updated](server-docs/meeting_room-v1/meeting_room/events/updated.md)
      - [Meeting Room Deleted](server-docs/meeting_room-v1/meeting_room/events/deleted.md)
    - [Error Codes](server-docs/deprecated-version-not-recommended/rooms/error-codes.md)
- **Rooms configuration**
      - [Rooms configuration overview](server-docs/vc-v1/room_config/rooms-configuration-overview.md)
      - [set_checkboard_access_code](server-docs/vc-v1/room_config/set_checkboard_access_code.md)
      - [set_room_access_code](server-docs/vc-v1/room_config/set_room_access_code.md)
      - [Query meeting room configuration](server-docs/vc-v1/room_config/query.md)
      - [Set up meeting room configuration](server-docs/vc-v1/room_config/set.md)

### IM & Chat

- **Message**
  - **Send Message**
        - [Send Message Card](server-docs/deprecated-version-not-recommended/im-chat/message/send-message/send-message-card.md)
        - [Send Text Message](server-docs/deprecated-version-not-recommended/im-chat/message/send-message/send-text-message.md)
        - [Send Image Message](server-docs/deprecated-version-not-recommended/im-chat/message/send-message/send-image-message.md)
        - [Send Post Message](server-docs/deprecated-version-not-recommended/im-chat/message/send-message/send-post-message.md)
        - [Send Share Group Message](server-docs/deprecated-version-not-recommended/im-chat/message/send-message/send-share-group-message.md)
        - [Recall Message](server-docs/deprecated-version-not-recommended/im-chat/message/send-message/recall-message.md)
        - [Query the Read Status of a Message](server-docs/deprecated-version-not-recommended/im-chat/message/send-message/query-the-read-status-of-a-message.md)
  - **Images**
        - [Upload Image](server-docs/deprecated-version-not-recommended/im-chat/message/images/upload-image.md)
        - [Get Image](server-docs/deprecated-version-not-recommended/im-chat/message/images/get-image.md)
  - **Files**
        - [Get File](server-docs/deprecated-version-not-recommended/im-chat/message/files/get-file.md)
      - [Events](server-docs/deprecated-version-not-recommended/im-chat/message/events.md)
- **Group Chat**
      - [Create Group](server-docs/deprecated-version-not-recommended/im-chat/group-chat/create-group.md)
      - [Obtain Group Info](server-docs/deprecated-version-not-recommended/im-chat/group-chat/obtain-group-info.md)
      - [Update Group Info](server-docs/deprecated-version-not-recommended/im-chat/group-chat/update-group-info.md)
      - [Disband Chat](server-docs/deprecated-version-not-recommended/im-chat/group-chat/disband-chat.md)
      - [Invite Users to Chat](server-docs/deprecated-version-not-recommended/im-chat/group-chat/invite-users-to-chat.md)
      - [Delete Users From Chat](server-docs/deprecated-version-not-recommended/im-chat/group-chat/delete-users-from-chat.md)
      - [Add Bot to Group](server-docs/deprecated-version-not-recommended/im-chat/group-chat/add-bot-to-group.md)
      - [Exit Bot From Group](server-docs/deprecated-version-not-recommended/im-chat/group-chat/exit-bot-from-group.md)
      - [Obtain Member List](server-docs/deprecated-version-not-recommended/im-chat/group-chat/obtain-member-list.md)
      - [Obtain Group List to Which User Belongs](server-docs/deprecated-version-not-recommended/im-chat/group-chat/obtain-group-list-to-which-user-belongs.md)
      - [Obtain Group List to Which Bot Belongs](server-docs/deprecated-version-not-recommended/im-chat/group-chat/obtain-group-list-to-which-bot-belongs.md)
      - [Search All Groups to Which User Belongs](server-docs/deprecated-version-not-recommended/im-chat/group-chat/search-all-groups-to-which-user-belongs.md)
      - [Events](server-docs/deprecated-version-not-recommended/im-chat/group-chat/events.md)


