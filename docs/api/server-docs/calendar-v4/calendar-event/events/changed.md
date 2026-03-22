---
title: "Schedule is changed"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/events/changed"
service: "calendar-v4"
resource: "calendar-event"
section: "Calendar"
scopes:
  - "calendar:calendar"
  - "calendar:calendar.event:read"
  - "calendar:calendar:readonly"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1736822159000"
---

# Event changes

After the user subscribes to the event change event, this event will be triggered when the event under the subscribed calendar changes.{Usage Examples}(url=/api/tools/api_explore/api_explore_config?project=calendar&version=v4&resource=calendar.event&event=changed)

## Precautions

Be sure to read this note before using this event.

- First call the Subscribe to event change event interface to subscribe to the event, and then go to the application to configure the event subscription, so that when the event is triggered Event data received. To learn about event subscription, see Event Subscription Overview.

- The calendar ID (calendar_event_id), change type (change_type), and change status of the user's reply to the schedule (rsvp_infos) parameters included in this event are all in the grayscale testing phase. If you need to use them, please consult your business contact or [technical support](https://applink.larksuite.com/TLJpeNdW).
	
    If you can currently only receive user information (user_id_list) and calendar ID (calendar_id) that have undergone event changes, then you can extract the user information in the user_id_list parameter after receiving the event request, and then use these user identities (user_access_token) to call Get event list interface to get the event information in the calendar through the calendar ID.

## Event
| Facts |  |
| --- | --- |
| Event type | calendar.calendar.event.changed_v4 |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `calendar:calendar` `calendar:calendar.event:read` `calendar:calendar:readonly` |
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
|   `calendar_id` | `string` | The calendar ID where the event is located. |
|   `user_id_list` | `user_id[]` | List of users who need to push events. For an introduction to different user IDs, see User Identity Overview. |
|     `union_id` | `string` | User's union_id. |
|     `user_id` | `string` | User's user_id. |
|     `open_id` | `string` | User's open_id. |
|   `calendar_event_id` | `string` | The event ID that changed. **Notice**: This parameter is in the grayscale testing phase. If you need to use it, please consult your business contact or [technical support](https://applink.larksuite.com/TLJpeNdW). |
|   `change_type` | `string` | Event change type. **Possible values are:** - create: The event is created on the calendar. Creating a new event or being invited to the event as a participant all belong to the create type. - update: The event has changed. - delete: The event disappears from the calendar. Deleting the event or being removed from the event as a participant belongs to the delete type. - rsvp: User-type participants actively respond to the event (including accept, decline, tentative). **Event aggregation strategy**: When actually pushing events, change events of the same calendar (calendarID) and the same event (eventID) will be aggregated and pushed in a window of 3 seconds. Within 3 seconds: - When the event is changed by create + delete, the event is not pushed. - When the event undergoes create + update changes, push events of create change type. - When the event undergoes delete + update changes, events of delete change type are pushed. - When the event undergoes update + update changes, only the event of the last update change type will be pushed. - When there are multiple rsvp changes, only the last rsvp change type event will be pushed. **Notice**: This parameter is in the grayscale testing phase. If you need to use it, please consult your business contact or [technical support](https://applink.larksuite.com/TLJpeNdW). |
|   `rsvp_infos` | `open_event_rsvp_info[]` | RSVP change details, that is, the schedule participant's reply status. **Notice**: - This parameter only contains change details for user type participants. - This parameter is in the grayscale testing phase. If you need to use it, please consult your business contact or [technical support](https://applink.larksuite.com/TLJpeNdW). |
|     `from_user_id` | `user_id` | User type The user ID of the participant. |
|       `union_id` | `string` | User's union_id |
|       `user_id` | `string` | user's user id **Required field scopes**: `contact:user.employee_id:readonly` |
|       `open_id` | `string` | User's open_id |
|     `rsvp_status` | `string` | RSVP operational status. **Possible values are:** - accept - decline - tentative | ### Event body example

{
    "schema": "2.0",
    "header": {
        "event_id": "5e3702a84e847582be8db7fb73283c02",
        "event_type": "calendar.calendar.event.changed_v4",
        "create_time": "1608725989000",
        "token": "rvaYgkND1GOiu5MM0E1rncYC6PLtF7JV",
        "app_id": "cli_9f5343c580712544",
        "tenant_key": "2ca1d211f64f6438"
    },
    "event": {
        "calendar_id": "larksuite.com_xxxxxxxxxx@group.calendar.larksuite.com",
        "user_id_list": [
            {
                "union_id": "on_xxxxxx",
                "user_id": "exxxxxxz",
                "open_id": "ou_xxxxxx"
            }
        ],
        "calendar_event_id": "efa67a98-06a8-4df5-8559-746c8f4477ef_0",
        "change_type": "create",
        "rsvp_infos": [
            {
                "from_user_id": {
                    "union_id": "on_8ed6aa67826108097d9ee143816345",
                    "user_id": "e33ggbyz",
                    "open_id": "ou_84aad35d084aa403a838cf73ee18467"
                },
                "rsvp_status": "accept"
            }
        ]
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
	"github.com/larksuite/oapi-sdk-go/v3/service/calendar/v4"
	larkws "github.com/larksuite/oapi-sdk-go/v3/ws"
)

// SDK 使用说明 SDK user guide：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/golang-sdk-guide/preparations
func main() {
	// 注册事件 Register event
	eventHandler := dispatcher.NewEventDispatcher("", "").
		OnP2CalendarEventChangedV4(func(ctx context.Context, event *larkcalendar.P2CalendarEventChangedV4) error {
			fmt.Printf("[ OnP2CalendarEventChangedV4 access ], data: %s\n", larkcore.Prettify(event))
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

def do_p2_calendar_calendar_event_changed_v4(data: lark.calendar.v4.P2CalendarCalendarEventChangedV4) -> None:
    print(f'[ do_p2_calendar_calendar_event_changed_v4 access ], data: {lark.JSON.marshal(data, indent=4)}')

# 注册事件 Register event
event_handler = lark.EventDispatcherHandler.builder("", "") \
    .register_p2_calendar_calendar_event_changed_v4(do_p2_calendar_calendar_event_changed_v4) \
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
import com.lark.oapi.service.calendar.CalendarService;
import com.lark.oapi.service.calendar.v4.model.P2CalendarEventChangedV4;
import com.lark.oapi.event.EventDispatcher;
import com.lark.oapi.ws.Client;

// SDK 使用说明 SDK user guide：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/java-sdk-guide/preparations
public class Sample {
    // 注册事件 Register event
    private static final EventDispatcher EVENT_HANDLER = EventDispatcher.newBuilder("", "")
            .onP2CalendarEventChangedV4(new CalendarService.P2CalendarEventChangedV4Handler() {
                @Override
                public void handle(P2CalendarEventChangedV4 event) throws Exception {
                    System.out.printf("[ onP2CalendarEventChangedV4 access ], data: %s\n", Jsons.DEFAULT.toJson(event.getEvent()));
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
        'calendar.calendar.event.changed_v4': async (data) => {
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
	"github.com/larksuite/oapi-sdk-go/v3/service/calendar/v4"
)

// SDK 使用说明 SDK user guide：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/golang-sdk-guide/preparations
func main() {
	// 注册事件 Register event
	eventHandler := dispatcher.NewEventDispatcher("", "").
		OnP2CalendarEventChangedV4(func(ctx context.Context, event *larkcalendar.P2CalendarEventChangedV4) error {
			fmt.Printf("[ OnP2CalendarEventChangedV4 access ], data: %s\n", larkcore.Prettify(event))
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

def do_p2_calendar_calendar_event_changed_v4(data: lark.calendar.v4.P2CalendarCalendarEventChangedV4) -> None:
    print(f'[ do_p2_calendar_calendar_event_changed_v4 access ], data: {lark.JSON.marshal(data, indent=4)}')

# 注册事件 Register event
event_handler = lark.EventDispatcherHandler.builder("", "") \
    .register_p2_calendar_calendar_event_changed_v4(do_p2_calendar_calendar_event_changed_v4) \
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
import com.lark.oapi.service.calendar.CalendarService;
import com.lark.oapi.service.calendar.v4.model.P2CalendarEventChangedV4;
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
            .onP2CalendarEventChangedV4(new CalendarService.P2CalendarEventChangedV4Handler() {
                @Override
                public void handle(P2CalendarEventChangedV4 event) throws Exception {
                    System.out.printf("[ onP2CalendarEventChangedV4 access ], data: %s\n", Jsons.DEFAULT.toJson(event.getEvent()));
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
    'calendar.calendar.event.changed_v4': async (data) => {
        console.log(data);
        return 'success';
    },
});

const server = http.createServer();
// 创建路由处理器 Create route handler
server.on('request', lark.adaptDefault('/webhook/event', eventDispatcher));
server.listen(3000);
