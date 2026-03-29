---
title: "User task status change event"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/Attendance/event/user-task-status-change-event"
section: "Attendance"
updated: "1646322873000"
---

# User task status change event
> To understand the usage scenarios and configuration process of event subscription, please refer to overview of event subscription.

**Event**

The notifications of the user's task status change will be sent when the user task is updated.

**Event body**
|Name|Type|Description|
|---|---|---|
|schema|string|Event mode|
|header|event_header|Event header|
|  ∟event_id|string|Event ID|
|  ∟event_type|string|Event type|
|  ∟create_time|string|Event creation timestamp (unit: millisecond)|
|  ∟token|string|Event Token|
|  ∟app_id|string|App ID|
|  ∟tenant_key|string|Tenant Key|
|event|-|Event body|
|  ∟employee_id|string|Employee ID|
|  ∟employee_no|string|Employee No|
|  ∟group_id|string|Attendance group ID|
|  ∟shift_id|string|Shift ID|
|  ∟date|int|Date|
|  ∟status_changes|list|Status change array|
|    ∟before_status|string|Attendance results before change, the value is: [NoNeedCheck (No need to record attendance), SystemCheck (System attendance), Normal (Normal), Early (Early out), Late (Late in), Lack (no records)]|
|    ∟current_status|string|Attendance results after change, the value is: [NoNeedCheck (No need to record attendance), SystemCheck (System attendance), Normal (Normal), Early (Early out), Late (Late in), Lack (no records)]|
|    ∟before_supplement|string|Results complement before change, the value is: [None (None), ManagerModification (administrator modification), CardReplacement (Correction passed), ShiftChange (Shift swap), Travel (Business trip), Leave (Leave), GoOut (Out of office), CardReplacementApplication (Correction requesting), FieldPunch (Offsite attendance)]|
|    ∟current_supplement|string|Results complement after change, the value is: [None (None), ManagerModification (administrator modification), CardReplacement (Correction passed), ShiftChange (Shift swap), Travel (Business trip), Leave (Leave), GoOut (Out of office), CardReplacementApplication (Correction requesting), FieldPunch (Offsite attendance)]|
|    ∟work_type|string|Change of work status, the value is: [on (on-duty), off (off-duty)]
|    ∟index|string|How many times of on-duty and off-duty during the task
**Event example**
```
{
  "schema": "2.0",
  "header": {
    "event_id": "2a0b662d75f87508a016a8da8b225a46",
    "token": "drIibiox5ZEZl5UvuRL3Uf3LNwD0fB6e",
    "create_time": "1615381917559",
    "event_type": "attendance.user_task.updated_v1",
    "tenant_key": "2fce678eb60d1651",
    "app_id": "cli_a0c27934a1f8100b"
  },
  "event": {
    "date": 20210310,
    "employee_id": "2b68933a",
    "employee_no": "",
    "group_id": "6935645428507557915",
    "shift_id": "6938006252085739521",
    "status_changes": [
      {
        "before_status": "Normal",
        "before_supplement": "None",
        "current_status": "Late",
        "current_supplement": "None",
        "index": 1,
        "work_type": "on"
      }
    ],
    "task_id": "6936733176245026817",
    "time_zone": "Asia/Shanghai"
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
	"github.com/larksuite/oapi-sdk-go/v3/service/attendance/v1"
	larkws "github.com/larksuite/oapi-sdk-go/v3/ws"
)

// SDK 使用说明 SDK user guide：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/golang-sdk-guide/preparations
func main() {
	// 注册事件 Register event
	eventHandler := dispatcher.NewEventDispatcher("", "").
		OnP2UserTaskUpdatedV1(func(ctx context.Context, event *larkattendance.P2UserTaskUpdatedV1) error {
			fmt.Printf("[ OnP2UserTaskUpdatedV1 access ], data: %s\n", larkcore.Prettify(event))
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

def do_p2_attendance_user_task_updated_v1(data: lark.attendance.v1.P2AttendanceUserTaskUpdatedV1) -> None:
    print(f'[ do_p2_attendance_user_task_updated_v1 access ], data: {lark.JSON.marshal(data, indent=4)}')

# 注册事件 Register event
event_handler = lark.EventDispatcherHandler.builder("", "") \
    .register_p2_attendance_user_task_updated_v1(do_p2_attendance_user_task_updated_v1) \
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
import com.lark.oapi.service.attendance.AttendanceService;
import com.lark.oapi.service.attendance.v1.model.P2UserTaskUpdatedV1;
import com.lark.oapi.event.EventDispatcher;
import com.lark.oapi.ws.Client;

// SDK 使用说明 SDK user guide：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/java-sdk-guide/preparations
public class Sample {
    // 注册事件 Register event
    private static final EventDispatcher EVENT_HANDLER = EventDispatcher.newBuilder("", "")
            .onP2UserTaskUpdatedV1(new AttendanceService.P2UserTaskUpdatedV1Handler() {
                @Override
                public void handle(P2UserTaskUpdatedV1 event) throws Exception {
                    System.out.printf("[ onP2UserTaskUpdatedV1 access ], data: %s\n", Jsons.DEFAULT.toJson(event.getEvent()));
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
        'attendance.user_task.updated_v1': async (data) => {
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
	"github.com/larksuite/oapi-sdk-go/v3/service/attendance/v1"
)

// SDK 使用说明 SDK user guide：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/golang-sdk-guide/preparations
func main() {
	// 注册事件 Register event
	eventHandler := dispatcher.NewEventDispatcher("", "").
		OnP2UserTaskUpdatedV1(func(ctx context.Context, event *larkattendance.P2UserTaskUpdatedV1) error {
			fmt.Printf("[ OnP2UserTaskUpdatedV1 access ], data: %s\n", larkcore.Prettify(event))
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

def do_p2_attendance_user_task_updated_v1(data: lark.attendance.v1.P2AttendanceUserTaskUpdatedV1) -> None:
    print(f'[ do_p2_attendance_user_task_updated_v1 access ], data: {lark.JSON.marshal(data, indent=4)}')

# 注册事件 Register event
event_handler = lark.EventDispatcherHandler.builder("", "") \
    .register_p2_attendance_user_task_updated_v1(do_p2_attendance_user_task_updated_v1) \
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
import com.lark.oapi.service.attendance.AttendanceService;
import com.lark.oapi.service.attendance.v1.model.P2UserTaskUpdatedV1;
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
            .onP2UserTaskUpdatedV1(new AttendanceService.P2UserTaskUpdatedV1Handler() {
                @Override
                public void handle(P2UserTaskUpdatedV1 event) throws Exception {
                    System.out.printf("[ onP2UserTaskUpdatedV1 access ], data: %s\n", Jsons.DEFAULT.toJson(event.getEvent()));
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
    'attendance.user_task.updated_v1': async (data) => {
        console.log(data);
        return 'success';
    },
});

const server = http.createServer();
// 创建路由处理器 Create route handler
server.on('request', lark.adaptDefault('/webhook/event', eventDispatcher));
server.listen(3000);
