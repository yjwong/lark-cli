---
title: "Invitation revoked"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-member-user/events/withdrawn"
service: "im-v1"
resource: "chat-member-user"
section: "Group Chat"
scopes:
  - "im:chat"
  - "im:chat.members:read"
  - "im:chat:readonly"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1717574953000"
---

# Invitation revoked

This event is triggered when the addition of a user to a group is revoked.{Usage Examples}(url=/api/tools/api_explore/api_explore_config?project=im&version=v1&resource=chat.member.user&event=withdrawn)

> Notes:
> - The bot ability needs to be enabled.
> - The Invitation revoked event under the Messenger category needs to be subscribed to.
> - The event is pushed to bots who have subscribed to this event in the group.

## Event
| Facts |  |
| --- | --- |
| Event type | im.chat.member.user.withdrawn_v1 |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `im:chat` `im:chat.members:read` `im:chat:readonly` |
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
|   `operator_tenant_key` | `string` | Operator tenant key, which is the unique identifier of the tenant on Lark, which is used to exchange for the corresponding tenant_access_token, and can also be used as the unique identifier of the tenant in the application |
|   `users` | `chat.member.user[]` | List of users whose addition to a group is revoked |
|     `name` | `string` | User name |
|     `tenant_key` | `string` | Tenant Key, which is the unique identifier of the tenant on Lark, which is used to exchange for the corresponding tenant_access_token, and can also be used as the unique identifier of the tenant in the application |
|     `user_id` | `user_id` | User ID |
|       `union_id` | `string` | User's union_id |
|       `user_id` | `string` | user's user id **Required field scopes**: `contact:user.employee_id:readonly` |
|       `open_id` | `string` | User's open_id |
|   `name` | `string` | Group name |
|   `i18n_names` | `i18n_names` | Internationalized name of the group |
|     `zh_cn` | `string` | Chinese name |
|     `en_us` | `string` | English name |
|     `ja_jp` | `string` | Japanese name | ### Event body example

{
    "schema": "2.0",
    "header": {
        "event_id": "5e3702a84e847582be8db7fb73283c02",
        "event_type": "im.chat.member.user.withdrawn_v1",
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
        "operator_tenant_key": "879dsyud8",
        "users": [
            {
                "name": "user name",
                "tenant_key": "3774yuu3743",
                "user_id": {
                    "union_id": "on_8ed6aa67826108097d9ee143816345",
                    "user_id": "e33ggbyz",
                    "open_id": "ou_84aad35d084aa403a838cf73ee18467"
                }
            }
        ],
        "name": "Group name testing",
        "i18n_names": {
            "zh_cn": "群聊",
            "en_us": "group chat",
            "ja_jp": "グループチャット"
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
		OnP2ChatMemberUserWithdrawnV1(func(ctx context.Context, event *larkim.P2ChatMemberUserWithdrawnV1) error {
			fmt.Printf("[ OnP2ChatMemberUserWithdrawnV1 access ], data: %s\n", larkcore.Prettify(event))
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

def do_p2_im_chat_member_user_withdrawn_v1(data: lark.im.v1.P2ImChatMemberUserWithdrawnV1) -> None:
    print(f'[ do_p2_im_chat_member_user_withdrawn_v1 access ], data: {lark.JSON.marshal(data, indent=4)}')

# 注册事件 Register event
event_handler = lark.EventDispatcherHandler.builder("", "") \
    .register_p2_im_chat_member_user_withdrawn_v1(do_p2_im_chat_member_user_withdrawn_v1) \
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
import com.lark.oapi.service.im.v1.model.P2ChatMemberUserWithdrawnV1;
import com.lark.oapi.event.EventDispatcher;
import com.lark.oapi.ws.Client;

// SDK 使用说明 SDK user guide：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/java-sdk-guide/preparations
public class Sample {
    // 注册事件 Register event
    private static final EventDispatcher EVENT_HANDLER = EventDispatcher.newBuilder("", "")
            .onP2ChatMemberUserWithdrawnV1(new ImService.P2ChatMemberUserWithdrawnV1Handler() {
                @Override
                public void handle(P2ChatMemberUserWithdrawnV1 event) throws Exception {
                    System.out.printf("[ onP2ChatMemberUserWithdrawnV1 access ], data: %s\n", Jsons.DEFAULT.toJson(event.getEvent()));
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
        'im.chat.member.user.withdrawn_v1': async (data) => {
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
		OnP2ChatMemberUserWithdrawnV1(func(ctx context.Context, event *larkim.P2ChatMemberUserWithdrawnV1) error {
			fmt.Printf("[ OnP2ChatMemberUserWithdrawnV1 access ], data: %s\n", larkcore.Prettify(event))
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

def do_p2_im_chat_member_user_withdrawn_v1(data: lark.im.v1.P2ImChatMemberUserWithdrawnV1) -> None:
    print(f'[ do_p2_im_chat_member_user_withdrawn_v1 access ], data: {lark.JSON.marshal(data, indent=4)}')

# 注册事件 Register event
event_handler = lark.EventDispatcherHandler.builder("", "") \
    .register_p2_im_chat_member_user_withdrawn_v1(do_p2_im_chat_member_user_withdrawn_v1) \
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
import com.lark.oapi.service.im.v1.model.P2ChatMemberUserWithdrawnV1;
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
            .onP2ChatMemberUserWithdrawnV1(new ImService.P2ChatMemberUserWithdrawnV1Handler() {
                @Override
                public void handle(P2ChatMemberUserWithdrawnV1 event) throws Exception {
                    System.out.printf("[ onP2ChatMemberUserWithdrawnV1 access ], data: %s\n", Jsons.DEFAULT.toJson(event.getEvent()));
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
    'im.chat.member.user.withdrawn_v1': async (data) => {
        console.log(data);
        return 'success';
    },
});

const server = http.createServer();
// 创建路由处理器 Create route handler
server.on('request', lark.adaptDefault('/webhook/event', eventDispatcher));
server.listen(3000);
