---
title: "Base app field changed"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file/events/bitable_field_changed"
service: "drive-v1"
resource: "file"
section: "Docs"
scopes:
  - "bitable:app"
  - "drive:drive"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1753668106000"
---

# Base app field change

This event is triggered when a subscribed Base app field changes. To understand the configuration process and usage scenarios for event subscriptions, refer to Event Overview.{Usage Examples}(url=/api/tools/api_explore/api_explore_config?project=drive&version=v1&resource=file&event=bitable_field_changed)

> File edited event is also triggered when a subscribed Base app field changes.

## Notes

If the app subscribes to events using `tenant_access_token`, you must request for both tenant token scope and user token scope to receive events.

## Prerequisite

Before adding this event, you need to ensure that you have called the Subscribe to Docs events API.

## Event
| Facts |  |
| --- | --- |
| Event type | drive.file.bitable_field_changed_v1 |
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
|   `file_type` | `string` | Document type |
|   `file_token` | `string` | Base token |
|   `table_id` | `string` | Base data table ID |
|   `operator_id` | `user_id` | User ID |
|     `union_id` | `string` | User's union_id |
|     `user_id` | `string` | user's user id **Required field scopes**: `contact:user.employee_id:readonly` |
|     `open_id` | `string` | User's open_id |
|   `action_list` | `bitable.table.field_action[]` | Field change type. The enumeration values are: - `field_added`: Field added - `field_edited`: Field edited - `field_deleted`: Field deleted |
|     `action` | `string` | Operation type |
|     `field_id` | `string` | Field ID |
|     `before_value` | `bitable.table.field_action.value` | Field value before operation |
|       `id` | `string` | Field ID |
|       `name` | `string` | Field name |
|       `type` | `int` | Field type |
|       `description` | `string` | Field description |
|       `property` | `bitable.table.field_action.value.property` | Field properties |
|         `formatter` | `string` | Display format of numbers and formula fields |
|         `date_formatter` | `string` | Display format for date, creation time, last updated time fields |
|         `auto_fill` | `boolean` | New records in the date field are automatically filled in Creation time |
|         `multiple` | `boolean` | Multiple members are allowed to be added in the personnel field, and multiple records are allowed in one-way association and two-way association |
|         `table_id` | `string` | The ID of the associated data table in the one-way association, two-way association field |
|         `table_name` | `string` | The name of the associated data table in a one-way association, two-way association field |
|         `back_field_name` | `string` | The name of the corresponding bidirectional association field in the associated data table in the bidirectional association field |
|         `input_type` | `string` | Geolocation input restrictions |
|         `back_field_id` | `string` | The id of the corresponding bidirectional association field in the associated data table in the bidirectional association field |
|         `auto_serial` | `bitable.table.field_action.value.property.auto_serial` | Automatic numbering type |
|           `type` | `string` | Automatic numbering type |
|           `options` | `bitable.table.field_action.value.property.auto_serial.options[]` | List of auto-numbering rules |
|             `type` | `string` | Optional rule item types for auto-numbering |
|             `value` | `string` | Value corresponding to type |
|         `options` | `bitable.table.field_action.value.property.option[]` | Option information for radio and multi-select fields |
|           `name` | `string` | Option name |
|           `id` | `string` | Option ID |
|           `color` | `int` | Option color |
|         `formula_expression` | `string` | Formula expressions for formula fields |
|     `after_value` | `bitable.table.field_action.value` | Field value after operation |
|       `id` | `string` | Field ID |
|       `name` | `string` | Field name |
|       `type` | `int` | Field type |
|       `description` | `string` | Field description |
|       `property` | `bitable.table.field_action.value.property` | Field properties |
|         `formatter` | `string` | Display format of numbers and formula fields |
|         `date_formatter` | `string` | Display format for date, creation time, last updated time fields |
|         `auto_fill` | `boolean` | New records in the date field are automatically filled in Creation time |
|         `multiple` | `boolean` | Multiple members are allowed to be added in the personnel field, and multiple records are allowed in one-way association and two-way association |
|         `table_id` | `string` | The ID of the associated data table in the one-way association, two-way association field |
|         `table_name` | `string` | The name of the associated data table in a one-way association, two-way association field |
|         `back_field_name` | `string` | The name of the corresponding bidirectional association field in the associated data table in the bidirectional association field |
|         `input_type` | `string` | Geolocation input restrictions |
|         `back_field_id` | `string` | The id of the corresponding bidirectional association field in the associated data table in the bidirectional association field |
|         `auto_serial` | `bitable.table.field_action.value.property.auto_serial` | Automatic numbering type |
|           `type` | `string` | Automatic numbering type |
|           `options` | `bitable.table.field_action.value.property.auto_serial.options[]` | List of auto-numbering rules |
|             `type` | `string` | Optional rule item types for auto-numbering |
|             `value` | `string` | Value corresponding to type |
|         `options` | `bitable.table.field_action.value.property.option[]` | Option information for radio and multi-select fields |
|           `name` | `string` | Option name |
|           `id` | `string` | Option ID |
|           `color` | `int` | Option color |
|         `formula_expression` | `string` | Formula expressions for formula fields |
|   `revision` | `int` | Base data table version number |
|   `subscriber_id_list` | `user_id[]` | Subscribe to user id list |
|     `union_id` | `string` | Union ID |
|     `user_id` | `string` | User ID |
|     `open_id` | `string` | Open ID |
|   `update_time` | `int` | Field change time | ### Event body example

{
    "schema": "2.0",
    "header": {
        "event_id": "5e3702a84e847582be8db7fb73283c02",
        "event_type": "drive.file.bitable_field_changed_v1",
        "create_time": "1608725989000",
        "token": "rvaYgkND1GOiu5MM0E1rncYC6PLtF7JV",
        "app_id": "cli_9f5343c580712544",
        "tenant_key": "2ca1d211f64f6438"
    },
    "event": {
        "file_type": "Base",
        "file_token": "bascntUPmnoH9kZbGJ8RWeabcef",
        "table_id": "tblWXe2d0I0abcef",
        "operator_id": {
            "union_id": "on_8ed6aa67826108097d9ee143816345",
            "user_id": "e33ggbyz",
            "open_id": "ou_84aad35d084aa403a838cf73ee18467"
        },
        "action_list": [
            {
                "action": "field_edited",
                "field_id": "fldmj5qNii",
                "before_value": {
                    "id": "fldmj5qNii",
                    "name": "field name",
                    "type": 20,
                    "description": "description",
                    "property": {
                        "formatter": "1,000",
                        "date_formatter": "yyyyMMdd",
                        "auto_fill": true,
                        "multiple": true,
                        "table_id": "tblIniLz0Ic8oXyN",
                        "table_name": "Table name",
                        "back_field_name": "Field name",
                        "input_type": "only_mobile",
                        "back_field_id": "fldmj5qNii",
                        "auto_serial": {
                            "type": "Custom",
                            "options": [
                                {
                                    "type": "created_time",
                                    "value": "yyyyMMdd"
                                }
                            ]
                        },
                        "options": [
                            {
                                "name": "option name",
                                "id": "optabcef",
                                "color": 3
                            }
                        ],
                        "formula_expression": "bitable::$table[tblIniLz0Ic8oXyN].$field[fldqatAwxx]*6+333"
                    }
                },
                "after_value": {
                    "id": "fldmj5qNii",
                    "name": "field name",
                    "type": 20,
                    "description": "description",
                    "property": {
                        "formatter": "1,000",
                        "date_formatter": "yyyyMMdd",
                        "auto_fill": true,
                        "multiple": true,
                        "table_id": "tblIniLz0Ic8oXyN",
                        "table_name": "Table name",
                        "back_field_name": "Field name",
                        "input_type": "only_mobile",
                        "back_field_id": "fldmj5qNii",
                        "auto_serial": {
                            "type": "Custom",
                            "options": [
                                {
                                    "type": "created_time",
                                    "value": "yyyyMMdd"
                                }
                            ]
                        },
                        "options": [
                            {
                                "name": "option name",
                                "id": "optabcef",
                                "color": 3
                            }
                        ],
                        "formula_expression": "bitable::$table[tblIniLz0Ic8oXyN].$field[fldqatAwxx]*6+333"
                    }
                }
            }
        ],
        "revision": 10,
        "subscriber_id_list": [
            {
                "union_id": "on_876b570a984d02ab1c0906a49e4a9d04",
                "user_id": "638474b8",
                "open_id": "ou_9bc587355789fc049904ae7c73619b89"
            }
        ],
        "update_time": 1663727688
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
		OnP2FileBitableFieldChangedV1(func(ctx context.Context, event *larkdrive.P2FileBitableFieldChangedV1) error {
			fmt.Printf("[ OnP2FileBitableFieldChangedV1 access ], data: %s\n", larkcore.Prettify(event))
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

def do_p2_drive_file_bitable_field_changed_v1(data: lark.drive.v1.P2DriveFileBitableFieldChangedV1) -> None:
    print(f'[ do_p2_drive_file_bitable_field_changed_v1 access ], data: {lark.JSON.marshal(data, indent=4)}')

# 注册事件 Register event
event_handler = lark.EventDispatcherHandler.builder("", "") \
    .register_p2_drive_file_bitable_field_changed_v1(do_p2_drive_file_bitable_field_changed_v1) \
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
import com.lark.oapi.service.drive.v1.model.P2FileBitableFieldChangedV1;
import com.lark.oapi.event.EventDispatcher;
import com.lark.oapi.ws.Client;

// SDK 使用说明 SDK user guide：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/java-sdk-guide/preparations
public class Sample {
    // 注册事件 Register event
    private static final EventDispatcher EVENT_HANDLER = EventDispatcher.newBuilder("", "")
            .onP2FileBitableFieldChangedV1(new DriveService.P2FileBitableFieldChangedV1Handler() {
                @Override
                public void handle(P2FileBitableFieldChangedV1 event) throws Exception {
                    System.out.printf("[ onP2FileBitableFieldChangedV1 access ], data: %s\n", Jsons.DEFAULT.toJson(event.getEvent()));
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
        'drive.file.bitable_field_changed_v1': async (data) => {
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
		OnP2FileBitableFieldChangedV1(func(ctx context.Context, event *larkdrive.P2FileBitableFieldChangedV1) error {
			fmt.Printf("[ OnP2FileBitableFieldChangedV1 access ], data: %s\n", larkcore.Prettify(event))
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

def do_p2_drive_file_bitable_field_changed_v1(data: lark.drive.v1.P2DriveFileBitableFieldChangedV1) -> None:
    print(f'[ do_p2_drive_file_bitable_field_changed_v1 access ], data: {lark.JSON.marshal(data, indent=4)}')

# 注册事件 Register event
event_handler = lark.EventDispatcherHandler.builder("", "") \
    .register_p2_drive_file_bitable_field_changed_v1(do_p2_drive_file_bitable_field_changed_v1) \
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
import com.lark.oapi.service.drive.v1.model.P2FileBitableFieldChangedV1;
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
            .onP2FileBitableFieldChangedV1(new DriveService.P2FileBitableFieldChangedV1Handler() {
                @Override
                public void handle(P2FileBitableFieldChangedV1 event) throws Exception {
                    System.out.printf("[ onP2FileBitableFieldChangedV1 access ], data: %s\n", Jsons.DEFAULT.toJson(event.getEvent()));
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
    'drive.file.bitable_field_changed_v1': async (data) => {
        console.log(data);
        return 'success';
    },
});

const server = http.createServer();
// 创建路由处理器 Create route handler
server.on('request', lark.adaptDefault('/webhook/event', eventDispatcher));
server.listen(3000);
