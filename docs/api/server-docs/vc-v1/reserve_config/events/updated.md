---
title: "Update room reserve limitation"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/reserve_config/events/updated"
service: "vc-v1"
resource: "reserve_config"
section: "Video Conferencing"
scopes:
  - "vc:room"
  - "vc:room:readonly"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1689326170000"
---

# Update room reserve limitation

This event is triggered when room reserve limitation is updated.{Usage Examples}(url=/api/tools/api_explore/api_explore_config?project=vc&version=v1&resource=reserve_config&event=updated)

## Event
| Facts |  |
| --- | --- |
| Event type | vc.reserve_config.updated_v1 |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `vc:room` `vc:room:readonly` |
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
|   `scope_id` | `string` | room or room level id |
|   `scope_type` | `int` | room or room level: 1 for room level, 2 for room **Data validation rules**: - Value range: `1` ～ `2` |
|   `approve_config` | `approval_config_event` | reservation approval settings |
|     `approval_switch` | `int` | approval_switch: 0 means closed, 1 means open **Data validation rules**: - Value range: `0` ～ `1` |
|     `approval_condition` | `int` | approval_condition: 0 means that all reservations require approval, 1 means that approval is required for those that meet the conditions **Data validation rules**: - Value range: `0` ～ `1` |
|     `meeting_duration` | `float` | bookings exceeding meeting_duration need to be approved (unit: hour, value range [0.1-99]) |
|     `approvers` | `subscribe_user_event[]` | approvers list, when the approval_switch is 1, at least one approver needs to be set |
|       `user_id` | `user_id` | User ID |
|         `union_id` | `string` | User's union_id |
|         `user_id` | `string` | user's user id **Required field scopes**: `contact:user.employee_id:readonly` |
|         `open_id` | `string` | User's open_id |
|   `time_config` | `time_config` | reservation time settings |
|     `time_switch` | `int` | reservation time switch: 0 for off, 1 for on **Data validation rules**: - Value range: `0` ～ `1` |
|     `days_in_advance` | `int` | the earliest time when the conference room can be reserved in advance (unit: day, value range [1-730]) Note: When not filled in, the default update is 365 |
|     `opening_hour` | `string` | You can start booking from opening_hour on the opening day (unit: second, value range [0,86400]) Note:  1.  If not filled in, it will be updated to 28800 by default 2.  If the filled value is not a multiple of 60, it will be automatically updated to the nearest integer multiple of 60. |
|     `start_time` | `string` | The start time of the daily scheduled time range (unit: second, value range [0,86400]) Note:  1.  When not filled in, it will be updated to 0 by default, and the end_time filled in at this time must not be less than 30. 2.  When both start_time and end_time are filled in, end_time should be at least 30 times longer than start_time. 3.   If the filled value is not a multiple of 60, it will be automatically updated to the nearest integer multiple of 60. |
|     `end_time` | `string` | End time of daily pre-determinable time range (unit: second, value range [0,86400]) Note:  1.  When not filled in, the default update is 86400, and the start_time filled in at this time must not be greater than or equal to 86370. 2.  When both start_time and end_time are filled in, end_time must be at least 30 longer than start_time. 3.  If the filled value is not a multiple of 60, it will be automatically updated to the nearest integer multiple of 60. |
|     `max_duration` | `int` | The upper limit of the duration of a single conference room reservation (unit: hour, value range [1,99]) Note: The default update is 2 when not filled in |
|   `reserve_scope_config` | `reserve_scope_config_event` | reservation range settings |
|     `allow_all_users` | `int` | the range of members that can book conference: 0 for some members, 1 for all members **Data validation rules**: - Value range: `0` ～ `1` |
|     `allow_users` | `subscribe_user_event[]` | list of available members |
|       `user_id` | `user_id` | User ID |
|         `union_id` | `string` | User's union_id |
|         `user_id` | `string` | user's user id **Required field scopes**: `contact:user.employee_id:readonly` |
|         `open_id` | `string` | User's open_id |
|     `allow_depts` | `subscribe_department[]` | list of available departments |
|       `department_id` | `string` | Reserve admin department ID | ### Event body example

{
    "schema": "2.0",
    "header": {
        "event_id": "5e3702a84e847582be8db7fb73283c02",
        "event_type": "vc.reserve_config.updated_v1",
        "create_time": "1608725989000",
        "token": "rvaYgkND1GOiu5MM0E1rncYC6PLtF7JV",
        "app_id": "cli_9f5343c580712544",
        "tenant_key": "2ca1d211f64f6438"
    },
    "event": {
        "scope_id": "omm_3c5dd7e09bac0c1758fcf9511bd1a771",
        "scope_type": 2,
        "approve_config": {
            "approval_switch": 0,
            "approval_condition": 0,
            "meeting_duration": 1,
            "approvers": [
                {
                    "user_id": {
                        "union_id": "on_8ed6aa67826108097d9ee143816345",
                        "user_id": "e33ggbyz",
                        "open_id": "ou_84aad35d084aa403a838cf73ee18467"
                    }
                }
            ]
        },
        "time_config": {
            "time_switch": 1,
            "days_in_advance": 30,
            "opening_hour": "27900",
            "start_time": "0",
            "end_time": "86400",
            "max_duration": 24
        },
        "reserve_scope_config": {
            "allow_all_users": 1,
            "allow_users": [
                {
                    "user_id": {
                        "union_id": "on_8ed6aa67826108097d9ee143816345",
                        "user_id": "e33ggbyz",
                        "open_id": "ou_84aad35d084aa403a838cf73ee18467"
                    }
                }
            ],
            "allow_depts": [
                {
                    "department_id": "od-47d8b570b0a011e9679a755efcc5f61a"
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
	"github.com/larksuite/oapi-sdk-go/v3/service/vc/v1"
	larkws "github.com/larksuite/oapi-sdk-go/v3/ws"
)

// SDK 使用说明 SDK user guide：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/golang-sdk-guide/preparations
func main() {
	// 注册事件 Register event
	eventHandler := dispatcher.NewEventDispatcher("", "").
		OnP2ReserveConfigUpdatedV1(func(ctx context.Context, event *larkvc.P2ReserveConfigUpdatedV1) error {
			fmt.Printf("[ OnP2ReserveConfigUpdatedV1 access ], data: %s\n", larkcore.Prettify(event))
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

def do_p2_vc_reserve_config_updated_v1(data: lark.vc.v1.P2VcReserveConfigUpdatedV1) -> None:
    print(f'[ do_p2_vc_reserve_config_updated_v1 access ], data: {lark.JSON.marshal(data, indent=4)}')

# 注册事件 Register event
event_handler = lark.EventDispatcherHandler.builder("", "") \
    .register_p2_vc_reserve_config_updated_v1(do_p2_vc_reserve_config_updated_v1) \
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
import com.lark.oapi.service.vc.VcService;
import com.lark.oapi.service.vc.v1.model.P2ReserveConfigUpdatedV1;
import com.lark.oapi.event.EventDispatcher;
import com.lark.oapi.ws.Client;

// SDK 使用说明 SDK user guide：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/java-sdk-guide/preparations
public class Sample {
    // 注册事件 Register event
    private static final EventDispatcher EVENT_HANDLER = EventDispatcher.newBuilder("", "")
            .onP2ReserveConfigUpdatedV1(new VcService.P2ReserveConfigUpdatedV1Handler() {
                @Override
                public void handle(P2ReserveConfigUpdatedV1 event) throws Exception {
                    System.out.printf("[ onP2ReserveConfigUpdatedV1 access ], data: %s\n", Jsons.DEFAULT.toJson(event.getEvent()));
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
        'vc.reserve_config.updated_v1': async (data) => {
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
	"github.com/larksuite/oapi-sdk-go/v3/service/vc/v1"
)

// SDK 使用说明 SDK user guide：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/golang-sdk-guide/preparations
func main() {
	// 注册事件 Register event
	eventHandler := dispatcher.NewEventDispatcher("", "").
		OnP2ReserveConfigUpdatedV1(func(ctx context.Context, event *larkvc.P2ReserveConfigUpdatedV1) error {
			fmt.Printf("[ OnP2ReserveConfigUpdatedV1 access ], data: %s\n", larkcore.Prettify(event))
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

def do_p2_vc_reserve_config_updated_v1(data: lark.vc.v1.P2VcReserveConfigUpdatedV1) -> None:
    print(f'[ do_p2_vc_reserve_config_updated_v1 access ], data: {lark.JSON.marshal(data, indent=4)}')

# 注册事件 Register event
event_handler = lark.EventDispatcherHandler.builder("", "") \
    .register_p2_vc_reserve_config_updated_v1(do_p2_vc_reserve_config_updated_v1) \
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
import com.lark.oapi.service.vc.VcService;
import com.lark.oapi.service.vc.v1.model.P2ReserveConfigUpdatedV1;
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
            .onP2ReserveConfigUpdatedV1(new VcService.P2ReserveConfigUpdatedV1Handler() {
                @Override
                public void handle(P2ReserveConfigUpdatedV1 event) throws Exception {
                    System.out.printf("[ onP2ReserveConfigUpdatedV1 access ], data: %s\n", Jsons.DEFAULT.toJson(event.getEvent()));
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
    'vc.reserve_config.updated_v1': async (data) => {
        console.log(data);
        return 'success';
    },
});

const server = http.createServer();
// 创建路由处理器 Create route handler
server.on('request', lark.adaptDefault('/webhook/event', eventDispatcher));
server.listen(3000);
