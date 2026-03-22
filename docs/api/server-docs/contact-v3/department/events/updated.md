---
title: "Department Updated"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/events/updated"
service: "contact-v3"
resource: "department"
section: "Contacts"
scopes:
  - "contact:contact:readonly_as_app"
  - "contact:contact:access_as_app"
field_scopes:
  - "contact:contact:readonly_as_app"
  - "contact:department.base:readonly"
  - "contact:department.organize:readonly"
updated: "1663221576000"
---

# Modify a department

This event is triggered in case of department updates. The old_object only displays the original values of updated fields. The scope to access contacts as an app only applies to older versions and is not recommended.{Usage Examples}(url=/api/tools/api_explore/api_explore_config?project=contact&version=v3&resource=department&event=updated)

## Event
| Facts |  |
| --- | --- |
| Event type | contact.department.updated_v3 |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `contact:contact:readonly_as_app` `contact:contact:access_as_app` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:contact:readonly_as_app` `contact:department.base:readonly` `contact:department.organize:readonly` |
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
|   `object` | `department_event` | Information after update |
|     `name` | `string` | Department name **Data validation rules**: - Minimum length: `1` characters **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.base:readonly` |
|     `parent_department_id` | `string` | Parent department ID **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.organize:readonly` |
|     `department_id` | `string` | Department's custom department ID **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.base:readonly` |
|     `open_department_id` | `string` | Department's open_id |
|     `leader_user_id` | `string` | Department manager's user ID **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.base:readonly` |
|     `chat_id` | `string` | Department group ID **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.organize:readonly` |
|     `order` | `int` | Department order **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.organize:readonly` |
|     `status` | `department_status` | Department status **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.base:readonly` |
|       `is_deleted` | `boolean` | Whether it is deleted |
|     `leaders` | `departmentLeader[]` | Head of department |
|       `leaderType` | `int` | department leaderType **Optional values are**:  - `1`: main leader - `2`: deputy leader  |
|       `leaderID` | `string` | leader id **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.organize:readonly` |
|   `old_object` | `department_event` | Information before update |
|     `name` | `string` | Department name **Data validation rules**: - Minimum length: `1` characters **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.base:readonly` |
|     `parent_department_id` | `string` | Parent department ID **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.organize:readonly` |
|     `department_id` | `string` | Department's custom department ID **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.base:readonly` |
|     `open_department_id` | `string` | Department's open_id |
|     `leader_user_id` | `string` | Department manager's user ID **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.base:readonly` |
|     `chat_id` | `string` | Department group ID **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.organize:readonly` |
|     `order` | `int` | Department order **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.organize:readonly` |
|     `status` | `department_status` | Department status **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.base:readonly` |
|       `is_deleted` | `boolean` | Whether it is deleted |
|     `leaders` | `departmentLeader[]` | Head of department |
|       `leaderType` | `int` | department leaderType **Optional values are**:  - `1`: main leader - `2`: deputy leader  |
|       `leaderID` | `string` | leader id **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:department.organize:readonly` | ### Event body example

{
    "schema": "2.0",
    "header": {
        "event_id": "5e3702a84e847582be8db7fb73283c02",
        "event_type": "contact.department.updated_v3",
        "create_time": "1608725989000",
        "token": "rvaYgkND1GOiu5MM0E1rncYC6PLtF7JV",
        "app_id": "cli_9f5343c580712544",
        "tenant_key": "2ca1d211f64f6438"
    },
    "event": {
        "object": {
            "name": "Test department",
            "parent_department_id": "jkfsd89782",
            "department_id": "yd7sa8yf2j",
            "open_department_id": "od_j10j52hjksd9g0isdfg43",
            "leader_user_id": "ou_3j1kh45jk18fgh23hf",
            "chat_id": "oc_uiy325uy23bnv48gdf",
            "order": 100,
            "status": {
                "is_deleted": false
            },
            "leaders": [
                {
                    "leaderType": 1,
                    "leaderID": "ou_7dab8a3d3cdcc9da365777c7ad535d62"
                }
            ]
        },
        "old_object": {
            "name": "Test department",
            "parent_department_id": "jkfsd89782",
            "department_id": "yd7sa8yf2j",
            "open_department_id": "od_j10j52hjksd9g0isdfg43",
            "leader_user_id": "ou_3j1kh45jk18fgh23hf",
            "chat_id": "oc_uiy325uy23bnv48gdf",
            "order": 100,
            "status": {
                "is_deleted": false
            },
            "leaders": [
                {
                    "leaderType": 1,
                    "leaderID": "ou_7dab8a3d3cdcc9da365777c7ad535d62"
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
	"github.com/larksuite/oapi-sdk-go/v3/service/contact/v3"
	larkws "github.com/larksuite/oapi-sdk-go/v3/ws"
)

// SDK 使用说明 SDK user guide：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/golang-sdk-guide/preparations
func main() {
	// 注册事件 Register event
	eventHandler := dispatcher.NewEventDispatcher("", "").
		OnP2DepartmentUpdatedV3(func(ctx context.Context, event *larkcontact.P2DepartmentUpdatedV3) error {
			fmt.Printf("[ OnP2DepartmentUpdatedV3 access ], data: %s\n", larkcore.Prettify(event))
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

def do_p2_contact_department_updated_v3(data: lark.contact.v3.P2ContactDepartmentUpdatedV3) -> None:
    print(f'[ do_p2_contact_department_updated_v3 access ], data: {lark.JSON.marshal(data, indent=4)}')

# 注册事件 Register event
event_handler = lark.EventDispatcherHandler.builder("", "") \
    .register_p2_contact_department_updated_v3(do_p2_contact_department_updated_v3) \
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
import com.lark.oapi.service.contact.ContactService;
import com.lark.oapi.service.contact.v3.model.P2DepartmentUpdatedV3;
import com.lark.oapi.event.EventDispatcher;
import com.lark.oapi.ws.Client;

// SDK 使用说明 SDK user guide：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/java-sdk-guide/preparations
public class Sample {
    // 注册事件 Register event
    private static final EventDispatcher EVENT_HANDLER = EventDispatcher.newBuilder("", "")
            .onP2DepartmentUpdatedV3(new ContactService.P2DepartmentUpdatedV3Handler() {
                @Override
                public void handle(P2DepartmentUpdatedV3 event) throws Exception {
                    System.out.printf("[ onP2DepartmentUpdatedV3 access ], data: %s\n", Jsons.DEFAULT.toJson(event.getEvent()));
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
        'contact.department.updated_v3': async (data) => {
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
	"github.com/larksuite/oapi-sdk-go/v3/service/contact/v3"
)

// SDK 使用说明 SDK user guide：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/golang-sdk-guide/preparations
func main() {
	// 注册事件 Register event
	eventHandler := dispatcher.NewEventDispatcher("", "").
		OnP2DepartmentUpdatedV3(func(ctx context.Context, event *larkcontact.P2DepartmentUpdatedV3) error {
			fmt.Printf("[ OnP2DepartmentUpdatedV3 access ], data: %s\n", larkcore.Prettify(event))
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

def do_p2_contact_department_updated_v3(data: lark.contact.v3.P2ContactDepartmentUpdatedV3) -> None:
    print(f'[ do_p2_contact_department_updated_v3 access ], data: {lark.JSON.marshal(data, indent=4)}')

# 注册事件 Register event
event_handler = lark.EventDispatcherHandler.builder("", "") \
    .register_p2_contact_department_updated_v3(do_p2_contact_department_updated_v3) \
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
import com.lark.oapi.service.contact.ContactService;
import com.lark.oapi.service.contact.v3.model.P2DepartmentUpdatedV3;
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
            .onP2DepartmentUpdatedV3(new ContactService.P2DepartmentUpdatedV3Handler() {
                @Override
                public void handle(P2DepartmentUpdatedV3 event) throws Exception {
                    System.out.printf("[ onP2DepartmentUpdatedV3 access ], data: %s\n", Jsons.DEFAULT.toJson(event.getEvent()));
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
    'contact.department.updated_v3': async (data) => {
        console.log(data);
        return 'success';
    },
});

const server = http.createServer();
// 创建路由处理器 Create route handler
server.on('request', lark.adaptDefault('/webhook/event', eventDispatcher));
server.listen(3000);
