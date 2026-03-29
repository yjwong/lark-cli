---
title: "Base app record changed"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file/events/bitable_record_changed"
service: "drive-v1"
resource: "file"
section: "Docs"
scopes:
  - "bitable:app"
  - "drive:drive"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1753668105000"
---

# Base app record change

Base record change event. This event is triggered when a subscribed multi-dimensional table record changes. To understand the configuration process and usage scenarios for event subscriptions, refer to Event Overview. {Usage Examples}(url=/api/tools/api_explore/api_explore_config?project=drive&version=v1&resource=file&event=bitable_record_changed)

> - A change in the value of a formula field in a Base table does not trigger an event.
> - The value of the formula field is not included in the event body when the Base table record is changed.
> - File edited event is also triggered when a subscribed Base app record changes.

## Notes

If the app subscribes to events using `tenant_access_token`, you must request for both tenant token scope and user token scope to receive events.

## Prerequisite

Before adding this event, you need to ensure that you have called the Subscribe to Docs events API.

## Event
| Facts |  |
| --- | --- |
| Event type | drive.file.bitable_record_changed_v1 |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `bitable:app` `drive:drive` |
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
|   `file_type` | `string` | File type. Fixed as bitable. |
|   `file_token` | `string` | Base's token |
|   `table_id` | `string` | The table ID which changes |
|   `revision` | `int` | The version ID for the table |
|   `operator_id` | `user_id` | The operator's ID |
|     `union_id` | `string` | User's union_id |
|     `user_id` | `string` | user's user id **Required field scopes**: `contact:user.employee_id:readonly` |
|     `open_id` | `string` | User's open_id |
|   `action_list` | `bitable.table.record_action[]` | The action list of the record changes |
|     `record_id` | `string` | The ID of changed record |
|     `action` | `string` | The change type. Can be: - record_added：add a new record - record_deleted：delete a record - record_edited：edit a record |
|     `before_value` | `bitable.table.record_action.field[]` | The record value before change occurs |
|       `field_id` | `string` | The ID of field which changes |
|       `field_value` | `string` | The field value after change occurs. The field is a JSON serialised string, see data-structure for the structure before serialisation. |
|       `field_identity_value` | `bitable.table.record_action.field.identity` | Additional information for the Person field. Returned when there is a change to the Person, Creator, Modifier type fields |
|         `users` | `bitable.table.record_action.field.identity.user[]` | User list |
|           `user_id` | `user_id` | User's ID |
|             `union_id` | `string` | User's union_id |
|             `user_id` | `string` | user's user id **Required field scopes**: `contact:user.employee_id:readonly` |
|             `open_id` | `string` | User's open_id |
|           `name` | `string` | User's name |
|           `en_name` | `string` | User's English name |
|           `avatar_url` | `string` | The URL of user's avatar's image |
|     `after_value` | `bitable.table.record_action.field[]` | The field value after change occurs |
|       `field_id` | `string` | The ID of field which changes |
|       `field_value` | `string` | The field value after change occurs. The field is a JSON serialised string, see data-structure for the structure before serialisation. |
|       `field_identity_value` | `bitable.table.record_action.field.identity` | Additional information for the Person field. Returned when there is a change to the Person, Creator, Modifier type fields |
|         `users` | `bitable.table.record_action.field.identity.user[]` | User list |
|           `user_id` | `user_id` | User's ID |
|             `union_id` | `string` | User's union_id |
|             `user_id` | `string` | user's user id **Required field scopes**: `contact:user.employee_id:readonly` |
|             `open_id` | `string` | User's open_id |
|           `name` | `string` | User's name |
|           `en_name` | `string` | User's English name |
|           `avatar_url` | `string` | The URL of user's avatar's image |
|   `subscriber_id_list` | `user_id[]` | List of subscribed users |
|     `union_id` | `string` | Union ID |
|     `user_id` | `string` | User ID |
|     `open_id` | `string` | Open ID |
|   `update_time` | `int` | The change time. The format is timestamp and the units are in seconds. | ### Event body example

{
    "schema": "2.0",
    "header": {
        "event_id": "5e3702a84e847582be8db7fb73283c02",
        "event_type": "drive.file.bitable_record_changed_v1",
        "create_time": "1608725989000",
        "token": "rvaYgkND1GOiu5MM0E1rncYC6PLtF7JV",
        "app_id": "cli_9f5343c580712544",
        "tenant_key": "2ca1d211f64f6438"
    },
    "event": {
        "file_type": "bitable",
        "file_token": "bascnItn6oHUSEL8RDUdF6abcef",
        "table_id": "tblOaqBWfGeabcef",
        "revision": 41,
        "operator_id": {
            "union_id": "on_8ed6aa67826108097d9ee143816345",
            "user_id": "e33ggbyz",
            "open_id": "ou_84aad35d084aa403a838cf73ee18467"
        },
        "action_list": [
            {
                "record_id": "rec9sabcef",
                "action": "record_edited",
                "before_value": [
                    {
                        "field_id": "fld9Eabcef",
                        "field_value": "666",
                        "field_identity_value": {
                            "users": [
                                {
                                    "user_id": {
                                        "union_id": "on_8ed6aa67826108097d9ee143816345",
                                        "user_id": "e33ggbyz",
                                        "open_id": "ou_84aad35d084aa403a838cf73ee18467"
                                    },
                                    "name": "张敏",
                                    "en_name": "Zhangmin",
                                    "avatar_url": "https://internal-api-lark-file.larksuite.com/static-resource/v1/v2_q86-fcb6-4f18-85c7-87ca8881e50j~?image_size=72x72&cut_type=default-face&quality=&format=jpeg&sticker_format=.webp"
                                }
                            ]
                        }
                    }
                ],
                "after_value": [
                    {
                        "field_id": "fld9Eabcef",
                        "field_value": "666",
                        "field_identity_value": {
                            "users": [
                                {
                                    "user_id": {
                                        "union_id": "on_8ed6aa67826108097d9ee143816345",
                                        "user_id": "e33ggbyz",
                                        "open_id": "ou_84aad35d084aa403a838cf73ee18467"
                                    },
                                    "name": "张敏",
                                    "en_name": "Zhangmin",
                                    "avatar_url": "https://internal-api-lark-file.larksuite.com/static-resource/v1/v2_q86-fcb6-4f18-85c7-87ca8881e50j~?image_size=72x72&cut_type=default-face&quality=&format=jpeg&sticker_format=.webp"
                                }
                            ]
                        }
                    }
                ]
            }
        ],
        "subscriber_id_list": [
            {
                "union_id": "on_876b570a984d02ab1c0906a49e4a9d04",
                "user_id": "638474b8",
                "open_id": "ou_9bc587355789fc049904ae7c73619b89"
            }
        ],
        "update_time": 1717040601
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
	"github.com/larksuite/oapi-sdk-go/v3/service/drive/v1"
	larkws "github.com/larksuite/oapi-sdk-go/v3/ws"
)

// SDK 使用说明 SDK user guide：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/golang-sdk-guide/preparations
func main() {
	// 注册事件 Register event
	eventHandler := dispatcher.NewEventDispatcher("", "").
		OnP2FileBitableRecordChangedV1(func(ctx context.Context, event *larkdrive.P2FileBitableRecordChangedV1) error {
			fmt.Printf("[ OnP2FileBitableRecordChangedV1 access ], data: %s\n", larkcore.Prettify(event))
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

def do_p2_drive_file_bitable_record_changed_v1(data: lark.drive.v1.P2DriveFileBitableRecordChangedV1) -> None:
    print(f'[ do_p2_drive_file_bitable_record_changed_v1 access ], data: {lark.JSON.marshal(data, indent=4)}')

# 注册事件 Register event
event_handler = lark.EventDispatcherHandler.builder("", "") \
    .register_p2_drive_file_bitable_record_changed_v1(do_p2_drive_file_bitable_record_changed_v1) \
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
import com.lark.oapi.service.drive.DriveService;
import com.lark.oapi.service.drive.v1.model.P2FileBitableRecordChangedV1;
import com.lark.oapi.event.EventDispatcher;
import com.lark.oapi.ws.Client;

// SDK 使用说明 SDK user guide：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/java-sdk-guide/preparations
public class Sample {
    // 注册事件 Register event
    private static final EventDispatcher EVENT_HANDLER = EventDispatcher.newBuilder("", "")
            .onP2FileBitableRecordChangedV1(new DriveService.P2FileBitableRecordChangedV1Handler() {
                @Override
                public void handle(P2FileBitableRecordChangedV1 event) throws Exception {
                    System.out.printf("[ onP2FileBitableRecordChangedV1 access ], data: %s\n", Jsons.DEFAULT.toJson(event.getEvent()));
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
        'drive.file.bitable_record_changed_v1': async (data) => {
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
	"github.com/larksuite/oapi-sdk-go/v3/service/drive/v1"
)

// SDK 使用说明 SDK user guide：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/golang-sdk-guide/preparations
func main() {
	// 注册事件 Register event
	eventHandler := dispatcher.NewEventDispatcher("", "").
		OnP2FileBitableRecordChangedV1(func(ctx context.Context, event *larkdrive.P2FileBitableRecordChangedV1) error {
			fmt.Printf("[ OnP2FileBitableRecordChangedV1 access ], data: %s\n", larkcore.Prettify(event))
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

def do_p2_drive_file_bitable_record_changed_v1(data: lark.drive.v1.P2DriveFileBitableRecordChangedV1) -> None:
    print(f'[ do_p2_drive_file_bitable_record_changed_v1 access ], data: {lark.JSON.marshal(data, indent=4)}')

# 注册事件 Register event
event_handler = lark.EventDispatcherHandler.builder("", "") \
    .register_p2_drive_file_bitable_record_changed_v1(do_p2_drive_file_bitable_record_changed_v1) \
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
import com.lark.oapi.service.drive.DriveService;
import com.lark.oapi.service.drive.v1.model.P2FileBitableRecordChangedV1;
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
            .onP2FileBitableRecordChangedV1(new DriveService.P2FileBitableRecordChangedV1Handler() {
                @Override
                public void handle(P2FileBitableRecordChangedV1 event) throws Exception {
                    System.out.printf("[ onP2FileBitableRecordChangedV1 access ], data: %s\n", Jsons.DEFAULT.toJson(event.getEvent()));
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
    'drive.file.bitable_record_changed_v1': async (data) => {
        console.log(data);
        return 'success';
    },
});

const server = http.createServer();
// 创建路由处理器 Create route handler
server.on('request', lark.adaptDefault('/webhook/event', eventDispatcher));
server.listen(3000);
