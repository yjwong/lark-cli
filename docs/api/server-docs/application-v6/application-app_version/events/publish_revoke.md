---
title: "Withdraw Application Release Application"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/application-v6/application-app_version/events/publish_revoke"
service: "application-v6"
resource: "application-app_version"
section: "App Information"
scopes:
  - "application:application.app_version:readonly"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1646720045000"
---

# App release application cancellation

This event is pushed when an app release application is canceled.{Usage Examples}(url=/api/tools/api_explore/api_explore_config?project=application&version=v6&resource=application.app_version&event=publish_revoke)

## Event
| Facts |  |
| --- | --- |
| Event type | application.application.app_version.publish_revoke_v6 |
| Supported app types | custom |
| Required scopes | `application:application.app_version:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` |
| Push Mode | `Webhook` | ### Event body
| Parameter | Type | Description |
| --- | --- | --- |
| `schema` | `string` | Event schema |
| `header` | `event_header` | Event header |
| ∟ `event_id` | `string` | Event ID |
| ∟ `event_type` | `string` | Event type |
| ∟ `create_time` | `string` | Event creation timestamp(in ms) |
| ∟ `token` | `string` | Event token |
| ∟ `app_id` | `string` | App ID |
| ∟ `tenant_key` | `string` | Tenant key |
| `event` | `\-` | \- |
| ∟ `operator_id` | `user_id` | User ID |
| ∟ `union_id` | `string` | User's union_id |
| ∟ `user_id` | `string` | user's user id **Required field scopes**: `contact:user.employee_id:readonly` |
| ∟ `open_id` | `string` | User's open_id |
| ∟ `creator_id` | `user_id` | User ID |
| ∟ `union_id` | `string` | User's union_id |
| ∟ `user_id` | `string` | user's user id **Required field scopes**: `contact:user.employee_id:readonly` |
| ∟ `open_id` | `string` | User's open_id |
| ∟ `app_id` | `string` | ID of the canceled app |
| ∟ `version_id` | `string` | Version ID of the canceled app | ### Event body example

```json
{
    "schema": "2.0",
    "header": {
        "event_id": "5e3702a84e847582be8db7fb73283c02",
        "event_type": "application.application.app_version.publish_revoke_v6",
        "create_time": "1608725989000",
        "token": "rvaYgkND1GOiu5MM0E1rncYC6PLtF7JV",
        "app_id": "cli_9f5343c580712544",
        "tenant_key": "2ca1d211f64f6438"
    },
    "event": {
        "operator_id": {
            "union_id": "on_8ed6aa67826108097d9ee143816345",
            "user_id": "e33ggbyz",
            "open_id": "ou_84aad35d084aa403a838cf73ee18467"
        },
        "creator_id": {
            "union_id": "on_8ed6aa67826108097d9ee143816345",
            "user_id": "e33ggbyz",
            "open_id": "ou_84aad35d084aa403a838cf73ee18467"
        },
        "app_id": "cli_9b445f5258795107",
        "version_id": "oav_d317f090b7258ad0372aa53963cda70d"
    }
}
```

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
	"github.com/larksuite/oapi-sdk-go/v3/service/application/v6"
	larkws "github.com/larksuite/oapi-sdk-go/v3/ws"
)

// SDK 使用说明 SDK user guide：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/golang-sdk-guide/preparations
func main() {
	// 注册事件 Register event
	eventHandler := dispatcher.NewEventDispatcher("", "").
		OnP2ApplicationAppVersionPublishRevokeV6(func(ctx context.Context, event *larkapplication.P2ApplicationAppVersionPublishRevokeV6) error {
			fmt.Printf("[ OnP2ApplicationAppVersionPublishRevokeV6 access ], data: %s\n", larkcore.Prettify(event))
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

def do_p2_application_application_app_version_publish_revoke_v6(data: lark.application.v6.P2ApplicationApplicationAppVersionPublishRevokeV6) -> None:
    print(f'[ do_p2_application_application_app_version_publish_revoke_v6 access ], data: {lark.JSON.marshal(data, indent=4)}')

# 注册事件 Register event
event_handler = lark.EventDispatcherHandler.builder("", "") \
    .register_p2_application_application_app_version_publish_revoke_v6(do_p2_application_application_app_version_publish_revoke_v6) \
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
import com.lark.oapi.service.application.ApplicationService;
import com.lark.oapi.service.application.v6.model.P2ApplicationAppVersionPublishRevokeV6;
import com.lark.oapi.event.EventDispatcher;
import com.lark.oapi.ws.Client;

// SDK 使用说明 SDK user guide：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/java-sdk-guide/preparations
public class Sample {
    // 注册事件 Register event
    private static final EventDispatcher EVENT_HANDLER = EventDispatcher.newBuilder("", "")
            .onP2ApplicationAppVersionPublishRevokeV6(new ApplicationService.P2ApplicationAppVersionPublishRevokeV6Handler() {
                @Override
                public void handle(P2ApplicationAppVersionPublishRevokeV6 event) throws Exception {
                    System.out.printf("[ onP2ApplicationAppVersionPublishRevokeV6 access ], data: %s\n", Jsons.DEFAULT.toJson(event.getEvent()));
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
        'application.application.app_version.publish_revoke_v6': async (data) => {
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
	"github.com/larksuite/oapi-sdk-go/v3/service/application/v6"
)

// SDK 使用说明 SDK user guide：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/golang-sdk-guide/preparations
func main() {
	// 注册事件 Register event
	eventHandler := dispatcher.NewEventDispatcher("", "").
		OnP2ApplicationAppVersionPublishRevokeV6(func(ctx context.Context, event *larkapplication.P2ApplicationAppVersionPublishRevokeV6) error {
			fmt.Printf("[ OnP2ApplicationAppVersionPublishRevokeV6 access ], data: %s\n", larkcore.Prettify(event))
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

def do_p2_application_application_app_version_publish_revoke_v6(data: lark.application.v6.P2ApplicationApplicationAppVersionPublishRevokeV6) -> None:
    print(f'[ do_p2_application_application_app_version_publish_revoke_v6 access ], data: {lark.JSON.marshal(data, indent=4)}')

# 注册事件 Register event
event_handler = lark.EventDispatcherHandler.builder("", "") \
    .register_p2_application_application_app_version_publish_revoke_v6(do_p2_application_application_app_version_publish_revoke_v6) \
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
import com.lark.oapi.service.application.ApplicationService;
import com.lark.oapi.service.application.v6.model.P2ApplicationAppVersionPublishRevokeV6;
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
            .onP2ApplicationAppVersionPublishRevokeV6(new ApplicationService.P2ApplicationAppVersionPublishRevokeV6Handler() {
                @Override
                public void handle(P2ApplicationAppVersionPublishRevokeV6 event) throws Exception {
                    System.out.printf("[ onP2ApplicationAppVersionPublishRevokeV6 access ], data: %s\n", Jsons.DEFAULT.toJson(event.getEvent()));
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
    'application.application.app_version.publish_revoke_v6': async (data) => {
        console.log(data);
        return 'success';
    },
});

const server = http.createServer();
// 创建路由处理器 Create route handler
server.on('request', lark.adaptDefault('/webhook/event', eventDispatcher));
server.listen(3000);
