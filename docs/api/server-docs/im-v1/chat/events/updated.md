---
title: "Group configuration changed"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat/events/updated"
service: "im-v1"
resource: "chat"
section: "Group Chat"
scopes:
  - "im:chat"
  - "im:chat:read"
  - "im:chat:readonly"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1717574948000"
---

# Group configuration changed

This event is triggered when the group configurations are modified, including:
- The owner is transferred.
- The basic information of the group is modified, such as the profile photo, name, description and international name of the group.
- The permission scopes of the group are modified, such as the permission scopes to add group members, edit the group, @mention all members, share the group, etc.{Usage Examples}(url=/api/tools/api_explore/api_explore_config?project=im&version=v1&resource=chat&event=updated)

> Notes:
> - The bot ability needs to be enabled.
> - You need to subscribe to Group configuration changed events under the IM category.
> - The event is pushed to bots who have subscribed to this event in the group.

## Event
| Facts |  |
| --- | --- |
| Event type | im.chat.updated_v1 |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `im:chat` `im:chat:read` `im:chat:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` |
| Push Mode | `Webhook` | ### Event body
| Parameter | Type | Description |
| --- | --- | --- |
| `schema` | `string` | Event schema |
| `header` | `event_header` | Event header |
|   `event_id` | `string` | Event ID |
|   `event_type` | `string` | Event type |
|   `create_time` | `string` | Event creation timestamp(in ms) |
|   `token` | `string` | Event token |
|   `app_id` | `string` | App ID |
|   `tenant_key` | `string` | Tenant key |
| `event` | `\-` | \- |
|   `chat_id` | `string` | Group ID. For details, refer to Group ID description. |
|   `operator_id` | `user_id` | User ID |
|     `union_id` | `string` | User's union_id |
|     `user_id` | `string` | user's user id **Required field scopes**: `contact:user.employee_id:readonly` |
|     `open_id` | `string` | User's open_id |
|   `external` | `boolean` | Whether it is an external group |
|   `operator_tenant_key` | `string` | The operator's tenant key |
|   `after_change` | `chat_change` | Updated group information |
|     `avatar` | `string` | Group profile photo |
|     `name` | `string` | Group name |
|     `description` | `string` | Group description |
|     `i18n_names` | `i18n_names` | Internationalized name of the group |
|       `zh_cn` | `string` | Chinese name |
|       `en_us` | `string` | English name |
|       `ja_jp` | `string` | Japanese name |
|     `add_member_permission` | `string` | Permission scopes to add group members (all_members/only_owner/unknown) |
|     `share_card_permission` | `string` | Group sharing scopes (allowed/not_allowed/unknown). |
|     `at_all_permission` | `string` | @mention all scopes (all_members/only_owner/unknown) |
|     `edit_permission` | `string` | Group edit scopes (all_members/only_owner/unknown) |
|     `membership_approval` | `string` | Group joining approval (no_approval_required/approval_required) |
|     `join_message_visibility` | `string` | Visibility of the messages on joining a group (only_owner/all_members/not_anyone) |
|     `leave_message_visibility` | `string` | Visibility of the messages on leaving a group (only_owner/all_members/not_anyone) |
|     `moderation_permission` | `string` | Speech scopes (all_members/only_owner) |
|     `owner_id` | `user_id` | User ID |
|       `union_id` | `string` | User's union_id |
|       `user_id` | `string` | user's user id **Required field scopes**: `contact:user.employee_id:readonly` |
|       `open_id` | `string` | User's open_id |
|     `restricted_mode_setting` | `restricted_mode_setting` | Restricted mode setting |
|       `status` | `boolean` | Whether the restricted mode is enabled |
|       `screenshot_has_permission_setting` | `string` | Whether to allow screenshots **Optional values are**:  - `all_members`: All members allow screenshots - `not_anyone`: Not anyone allow screenshots  |
|       `download_has_permission_setting` | `string` | Whether to allow downloading pictures, videos and files in messages **Optional values are**:  - `all_members`: All members allow downloading pictures, videos and files in messages - `not_anyone`: Not anyone allow downloading pictures, videos and files in messages  |
|       `message_has_permission_setting` | `string` | Whether to allow copying and forwarding of messages **Optional values are**:  - `all_members`: All members allow copying and forwarding of messages - `not_anyone`: Not anyone allow copying and forwarding of messages  |
|     `group_message_type` | `string` | Group message type (chat/thread) |
|   `before_change` | `chat_change` | Group information before update |
|     `avatar` | `string` | Group profile photo |
|     `name` | `string` | Group name |
|     `description` | `string` | Group description |
|     `i18n_names` | `i18n_names` | Internationalized name of the group |
|       `zh_cn` | `string` | Chinese name |
|       `en_us` | `string` | English name |
|       `ja_jp` | `string` | Japanese name |
|     `add_member_permission` | `string` | Permission scopes to add group members (all_members/only_owner/unknown) |
|     `share_card_permission` | `string` | Group sharing scopes (allowed/not_allowed/unknown). |
|     `at_all_permission` | `string` | @mention all scopes (all_members/only_owner/unknown) |
|     `edit_permission` | `string` | Group edit scopes (all_members/only_owner/unknown) |
|     `membership_approval` | `string` | Group joining approval (no_approval_required/approval_required) |
|     `join_message_visibility` | `string` | Visibility of the messages on joining a group (only_owner/all_members/not_anyone) |
|     `leave_message_visibility` | `string` | Visibility of the messages on leaving a group (only_owner/all_members/not_anyone) |
|     `moderation_permission` | `string` | Speech scopes (all_members/only_owner) |
|     `owner_id` | `user_id` | User ID |
|       `union_id` | `string` | User's union_id |
|       `user_id` | `string` | user's user id **Required field scopes**: `contact:user.employee_id:readonly` |
|       `open_id` | `string` | User's open_id |
|     `restricted_mode_setting` | `restricted_mode_setting` | Restricted mode setting |
|       `status` | `boolean` | Whether the restricted mode is enabled |
|       `screenshot_has_permission_setting` | `string` | Whether to allow screenshots **Optional values are**:  - `all_members`: All members allow screenshots - `not_anyone`: Not anyone allow screenshots  |
|       `download_has_permission_setting` | `string` | Whether to allow downloading pictures, videos and files in messages **Optional values are**:  - `all_members`: All members allow downloading pictures, videos and files in messages - `not_anyone`: Not anyone allow downloading pictures, videos and files in messages  |
|       `message_has_permission_setting` | `string` | Whether to allow copying and forwarding of messages **Optional values are**:  - `all_members`: All members allow copying and forwarding of messages - `not_anyone`: Not anyone allow copying and forwarding of messages  |
|     `group_message_type` | `string` | Group message type (chat/thread) |
|   `moderator_list` | `moderator_list` | Information about changes in the list of members who can speak in the group |
|     `added_member_list` | `list_event_moderator[]` | List of users who can speak after joining the group. The list must include the owner. |
|       `tenant_key` | `string` | Tenant key |
|       `user_id` | `user_id` | User ID |
|         `union_id` | `string` | User's union_id |
|         `user_id` | `string` | user's user id **Required field scopes**: `contact:user.employee_id:readonly` |
|         `open_id` | `string` | User's open_id |
|     `removed_member_list` | `list_event_moderator[]` | List of users who are removed from the list of users who can speak in the group |
|       `tenant_key` | `string` | Tenant key |
|       `user_id` | `user_id` | User ID |
|         `union_id` | `string` | User's union_id |
|         `user_id` | `string` | user's user id **Required field scopes**: `contact:user.employee_id:readonly` |
|         `open_id` | `string` | User's open_id | ### Event body example

{
    "schema": "2.0",
    "header": {
        "event_id": "5e3702a84e847582be8db7fb73283c02",
        "event_type": "im.chat.updated_v1",
        "create_time": "1608725989000",
        "token": "rvaYgkND1GOiu5MM0E1rncYC6PLtF7JV",
        "app_id": "cli_9f5343c580712544",
        "tenant_key": "2ca1d211f64f6438"
    },
    "event": {
        "chat_id": "oc_413871888e0d5492e25b173f0812efb7",
        "operator_id": {
            "union_id": "on_8ed6aa67826108097d9ee143816345",
            "user_id": "e33ggbyz",
            "open_id": "ou_84aad35d084aa403a838cf73ee18467"
        },
        "external": false,
        "operator_tenant_key": "86gwe65",
        "after_change": {
            "avatar": "default-avatar_0cda3662-875a-4354-94d2-83e7393c7123",
            "name": "Group name testing",
            "description": "Group description testing",
            "i18n_names": {
                "zh_cn": "群聊",
                "en_us": "group chat",
                "ja_jp": "グループチャット"
            },
            "add_member_permission": "all_members",
            "share_card_permission": "allowed",
            "at_all_permission": "only_owner",
            "edit_permission": "all_members",
            "membership_approval": "approval_required",
            "join_message_visibility": "all_members",
            "leave_message_visibility": "all_members",
            "moderation_permission": "all_members",
            "owner_id": {
                "union_id": "on_8ed6aa67826108097d9ee143816345",
                "user_id": "e33ggbyz",
                "open_id": "ou_84aad35d084aa403a838cf73ee18467"
            },
            "restricted_mode_setting": {
                "status": false,
                "screenshot_has_permission_setting": "all_members",
                "download_has_permission_setting": "all_members",
                "message_has_permission_setting": "all_members"
            },
            "group_message_type": "thread"
        },
        "before_change": {
            "avatar": "default-avatar_0cda3662-875a-4354-94d2-83e7393c7123",
            "name": "Group name testing",
            "description": "Group description testing",
            "i18n_names": {
                "zh_cn": "群聊",
                "en_us": "group chat",
                "ja_jp": "グループチャット"
            },
            "add_member_permission": "all_members",
            "share_card_permission": "allowed",
            "at_all_permission": "only_owner",
            "edit_permission": "all_members",
            "membership_approval": "approval_required",
            "join_message_visibility": "all_members",
            "leave_message_visibility": "all_members",
            "moderation_permission": "all_members",
            "owner_id": {
                "union_id": "on_8ed6aa67826108097d9ee143816345",
                "user_id": "e33ggbyz",
                "open_id": "ou_84aad35d084aa403a838cf73ee18467"
            },
            "restricted_mode_setting": {
                "status": false,
                "screenshot_has_permission_setting": "all_members",
                "download_has_permission_setting": "all_members",
                "message_has_permission_setting": "all_members"
            },
            "group_message_type": "thread"
        },
        "moderator_list": {
            "added_member_list": [
                {
                    "tenant_key": "86gwe65",
                    "user_id": {
                        "union_id": "on_8ed6aa67826108097d9ee143816345",
                        "user_id": "e33ggbyz",
                        "open_id": "ou_84aad35d084aa403a838cf73ee18467"
                    }
                }
            ],
            "removed_member_list": [
                {
                    "tenant_key": "86gwe65",
                    "user_id": {
                        "union_id": "on_8ed6aa67826108097d9ee143816345",
                        "user_id": "e33ggbyz",
                        "open_id": "ou_84aad35d084aa403a838cf73ee18467"
                    }
                }
            ]
        }
    }
}

### Sample code for event subscriptions

For event subscription process, refer to:Event Subscription overview，Guide for beginners:Tutorial

  Event subscription mode
  

  
	
    
package main

import (
	"context"
	"fmt"

	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	larkevent "github.com/larksuite/oapi-sdk-go/v3/event"
	"github.com/larksuite/oapi-sdk-go/v3/event/dispatcher"
	"github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
	larkws "github.com/larksuite/oapi-sdk-go/v3/ws"
)

// SDK 使用说明 SDK user guide：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/golang-sdk-guide/preparations
func main() {
	// 注册事件 Register event
	eventHandler := dispatcher.NewEventDispatcher("", "").
		OnP2ChatUpdatedV1(func(ctx context.Context, event *larkim.P2ChatUpdatedV1) error {
			fmt.Printf("[ OnP2ChatUpdatedV1 access ], data: %s\n", larkcore.Prettify(event))
			return nil
		})

	// 构建 client Build client
	cli := larkws.NewClient("YOUR_APP_ID", "YOUR_APP_SECRET",
		larkws.WithEventHandler(eventHandler),
		larkws.WithLogLevel(larkcore.LogLevelDebug),
	)

	// 建立长连接 Establish persistent connection
	err := cli.Start(context.Background())

	if err != nil {
		panic(err)
	}
}

    

    
# SDK 使用说明 SDK user guide：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/python--sdk/preparations-before-development
import lark_oapi as lark

def do_p2_im_chat_updated_v1(data: lark.im.v1.P2ImChatUpdatedV1) -> None:
    print(f'[ do_p2_im_chat_updated_v1 access ], data: {lark.JSON.marshal(data, indent=4)}')

# 注册事件 Register event
event_handler = lark.EventDispatcherHandler.builder("", "") \
    .register_p2_im_chat_updated_v1(do_p2_im_chat_updated_v1) \
    .build()

def main():
    # 构建 client Build client
    cli = lark.ws.Client("APP_ID", "APP_SECRET",
                        event_handler=event_handler, log_level=lark.LogLevel.DEBUG)
    # 建立长连接 Establish persistent connection
    cli.start()

if __name__ == "__main__":
    main()

    

    

package com.example.sample;

import com.lark.oapi.core.utils.Jsons;
import com.lark.oapi.service.im.ImService;
import com.lark.oapi.service.im.v1.model.P2ChatUpdatedV1;
import com.lark.oapi.event.EventDispatcher;
import com.lark.oapi.ws.Client;

// SDK 使用说明 SDK user guide：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/java-sdk-guide/preparations
public class Sample {
    // 注册事件 Register event
    private static final EventDispatcher EVENT_HANDLER = EventDispatcher.newBuilder("", "")
            .onP2ChatUpdatedV1(new ImService.P2ChatUpdatedV1Handler() {
                @Override
                public void handle(P2ChatUpdatedV1 event) throws Exception {
                    System.out.printf("[ onP2ChatUpdatedV1 access ], data: %s\n", Jsons.DEFAULT.toJson(event.getEvent()));
                }
            })
            .build();

    public static void main(String[] args) {
        // 构建 client Build client
        Client client = new Client.Builder("APP_ID", "APP_SECRET")
                .eventHandler(EVENT_HANDLER)
                .build();
        // 建立长连接 Establish persistent connection
        client.start();
    }
}
    

    
// SDK 使用说明 SDK user guide：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/nodejs-sdk/preparation-before-development
import * as Lark from '@larksuiteoapi/node-sdk';
const baseConfig = {
    appId: 'APP_ID',
    appSecret: 'APP_SECRET'
}
// 构建 client Build client
const wsClient = new Lark.WSClient(baseConfig);
// 建立长连接 Establish persistent connection
wsClient.start({
    // 注册事件 Register event
    eventDispatcher: new Lark.EventDispatcher({}).register({
        'im.chat.updated_v1': async (data) => {
            console.log(data);
        }
    })
});
    

  
  
	
    
package main

import (
	"context"
	"fmt"
	"net/http"

	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/core/httpserverext"
	larkevent "github.com/larksuite/oapi-sdk-go/v3/event"
	"github.com/larksuite/oapi-sdk-go/v3/event/dispatcher"
	"github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

// SDK 使用说明 SDK user guide：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/golang-sdk-guide/preparations
func main() {
	// 注册事件 Register event
	eventHandler := dispatcher.NewEventDispatcher("", "").
		OnP2ChatUpdatedV1(func(ctx context.Context, event *larkim.P2ChatUpdatedV1) error {
			fmt.Printf("[ OnP2ChatUpdatedV1 access ], data: %s\n", larkcore.Prettify(event))
			return nil
		})

	// 创建路由处理器 Create route handler
	http.HandleFunc("/webhook/event", httpserverext.NewEventHandlerFunc(handler, larkevent.WithLogLevel(larkcore.LogLevelDebug)))

	err := http.ListenAndServe(":7777", nil)

	if err != nil {
		panic(err)
	}
}

    

    
# SDK 使用说明 SDK user guide：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/python--sdk/preparations-before-development
from flask import Flask
from lark_oapi.adapter.flask import *
import lark_oapi as lark

app = Flask(__name__)

def do_p2_im_chat_updated_v1(data: lark.im.v1.P2ImChatUpdatedV1) -> None:
    print(f'[ do_p2_im_chat_updated_v1 access ], data: {lark.JSON.marshal(data, indent=4)}')

# 注册事件 Register event
event_handler = lark.EventDispatcherHandler.builder("", "") \
    .register_p2_im_chat_updated_v1(do_p2_im_chat_updated_v1) \
    .build()

# 创建路由处理器 Create route handler
@app.route("/webhook/event", methods=["POST"])
def event():
    resp = event_handler.do(parse_req())
    return parse_resp(resp)

if __name__ == "__main__":
    app.run(port=7777)

    

    

package com.lark.oapi.sample.event;

import com.lark.oapi.core.utils.Jsons;
import com.lark.oapi.service.im.ImService;
import com.lark.oapi.service.im.v1.model.P2ChatUpdatedV1;
import com.lark.oapi.sdk.servlet.ext.ServletAdapter;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

// SDK 使用说明 SDK user guide：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/java-sdk-guide/preparations
@RestController
public class EventController {

    // 注册事件 Register event
    private static final EventDispatcher EVENT_HANDLER = EventDispatcher.newBuilder("verificationToken", "encryptKey")
            .onP2ChatUpdatedV1(new ImService.P2ChatUpdatedV1Handler() {
                @Override
                public void handle(P2ChatUpdatedV1 event) throws Exception {
                    System.out.printf("[ onP2ChatUpdatedV1 access ], data: %s\n", Jsons.DEFAULT.toJson(event.getEvent()));
                }
            })
            .build();

    // 注入 ServletAdapter 实例 Inject ServletAdapter instance
    @Autowired
    private ServletAdapter servletAdapter;

    // 创建路由处理器 Create route handler
    @RequestMapping("/webhook/event")
    public void event(HttpServletRequest request, HttpServletResponse response)
            throws Throwable {
        // 回调扩展包提供的事件回调处理器 Callback handler provided by the extension package
        servletAdapter.handleEvent(request, response, EVENT_DISPATCHER);
    }
}
    

    
// SDK 使用说明 SDK user guide：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/nodejs-sdk/preparation-before-development
import http from 'http';
import * as lark from '@larksuiteoapi/node-sdk';

// 注册事件 Register event
const eventDispatcher = new lark.EventDispatcher({
    encryptKey: '',
    verificationToken: '',
}).register({
    'im.chat.updated_v1': async (data) => {
        console.log(data);
        return 'success';
    },
});

const server = http.createServer();
// 创建路由处理器 Create route handler
server.on('request', lark.adaptDefault('/webhook/event', eventDispatcher));
server.listen(3000);
