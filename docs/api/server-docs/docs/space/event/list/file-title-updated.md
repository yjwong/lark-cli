---
title: "File title updated"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/event/file-title-update"
section: "Docs"
scopes:
  - "drive:drive"
  - "docs:doc"
  - "sheets:spreadsheet"
  - "docs:event.document_edited:read"
field_scopes:
  - "contact:user.employee_id:readonly"
updated: "1768985068000"
---

# File title updated

The cloud document title change event. This event is triggered when the subscribed cloud document title changes.
> To understand the configuration process and usage scenarios for event subscriptions, refer to Event Overview.

## Prerequisite

Before adding this event, you need to ensure that you have called the Subscribe to Docs events API.
## Event
| Facts |  |
| --- | --- |
| Event type | drive.file.title_updated_v1 |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `drive:drive` `docs:doc` `sheets:spreadsheet` `docs:event.document_edited:read` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee_id:readonly` |
| Push Mode | `Webhook` | ## Event Body

| 
Name | Type | Description |
| --- | --- | --- |
| `schema` | `string` | Event format version. If this field is not present, it defaults to version 1.0. |
| `header` | `event_header` | Event header |
|   `event_id` | `string` | Event ID |
|   `event_type` | `string` | Event type |
|   `token` | `string` | Event token |
|   `create_time` | string | Event creation timestamp (in milliseconds) |
|   `tenant_key` | string | Tenant key, i.e., enterprise identifier |
|   `app_id` | string | Application ID |
|   `resource_id` | string | Subscribed cloud document token |
|   `user_list` | string | Subscribed user list |
|     union_id | string | User's Union ID |
| event | \- | \- |
|   file_type | string | Cloud document type, supporting the following enumerations: - `doc`: Old version document, not recommended for use - `docx`: New version document - `sheet`: Spreadsheet - `bitable`: Multi-dimensional table - `slides`:Slides - `file`:File |
|   file_token | string | Cloud document token |
|   operator_id | - | Operator ID |
|     union_id | string | User's Union ID |
|     user_id | string | User's User ID **Field Permission Requirements**: Get user user ID |
|     open_id | string | User's Open ID | ## Event Body Example

```json
{
  "schema": "2.0", // Version of the event format. If this field is absent, the version is 1.0
  "header": {
    "event_id": "88ba7a69073dcc5b0dea70be77xxxxxx", // Unique identifier of the event
    "token": "xxxxxx", // Event token
    "create_time": "1612246959000", // Event creation timestamp (in milliseconds)
    "event_type": "drive.file.title_updated_v1", // Event type
    "tenant_key": "xxxxxx", // Tenant key, i.e., enterprise identifier
    "app_id": "cli_xxxxxx", // Application ID
    "resource_id": "doccnfYZzTlvXqZIGTdAHKabcef", // Resource ID, i.e., subscribed cloud document token
    "user_list": [
      { // Subscribed user list
        "union_id": "on_xxxxxx"
      }
    ]
  },
  "event": {
    "file_token": "doccnfYZzTlvXqZIGTdAHKabcef", // File token
    "file_type": "docx", // File type, supporting doc, sheet, docx, and bitable
    "operator_id_list": { 
      // Operator ID list
        "open_id": "ou_xxxxxx",
        "union_id": "on_xxxxxx",
        "user_id": "xxxxxx"
      }
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
	"github.com/larksuite/oapi-sdk-go/v3/service/drive/v1"
	larkws "github.com/larksuite/oapi-sdk-go/v3/ws"
)

// SDK 使用说明 SDK user guide：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/golang-sdk-guide/preparations
func main() {
	// 注册事件 Register event
	eventHandler := dispatcher.NewEventDispatcher("", "").
		OnP2FileTitleUpdatedV1(func(ctx context.Context, event *larkdrive.P2FileTitleUpdatedV1) error {
			fmt.Printf("[ OnP2FileTitleUpdatedV1 access ], data: %s\n", larkcore.Prettify(event))
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

def do_p2_drive_file_title_updated_v1(data: lark.drive.v1.P2DriveFileTitleUpdatedV1) -> None:
    print(f'[ do_p2_drive_file_title_updated_v1 access ], data: {lark.JSON.marshal(data, indent=4)}')

# 注册事件 Register event
event_handler = lark.EventDispatcherHandler.builder("", "") \
    .register_p2_drive_file_title_updated_v1(do_p2_drive_file_title_updated_v1) \
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
import com.lark.oapi.service.drive.v1.model.P2FileTitleUpdatedV1;
import com.lark.oapi.event.EventDispatcher;
import com.lark.oapi.ws.Client;

// SDK 使用说明 SDK user guide：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/java-sdk-guide/preparations
public class Sample {
    // 注册事件 Register event
    private static final EventDispatcher EVENT_HANDLER = EventDispatcher.newBuilder("", "")
            .onP2FileTitleUpdatedV1(new DriveService.P2FileTitleUpdatedV1Handler() {
                @Override
                public void handle(P2FileTitleUpdatedV1 event) throws Exception {
                    System.out.printf("[ onP2FileTitleUpdatedV1 access ], data: %s\n", Jsons.DEFAULT.toJson(event.getEvent()));
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
        'drive.file.title_updated_v1': async (data) => {
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
		OnP2FileTitleUpdatedV1(func(ctx context.Context, event *larkdrive.P2FileTitleUpdatedV1) error {
			fmt.Printf("[ OnP2FileTitleUpdatedV1 access ], data: %s\n", larkcore.Prettify(event))
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

def do_p2_drive_file_title_updated_v1(data: lark.drive.v1.P2DriveFileTitleUpdatedV1) -> None:
    print(f'[ do_p2_drive_file_title_updated_v1 access ], data: {lark.JSON.marshal(data, indent=4)}')

# 注册事件 Register event
event_handler = lark.EventDispatcherHandler.builder("", "") \
    .register_p2_drive_file_title_updated_v1(do_p2_drive_file_title_updated_v1) \
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
import com.lark.oapi.service.drive.v1.model.P2FileTitleUpdatedV1;
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
            .onP2FileTitleUpdatedV1(new DriveService.P2FileTitleUpdatedV1Handler() {
                @Override
                public void handle(P2FileTitleUpdatedV1 event) throws Exception {
                    System.out.printf("[ onP2FileTitleUpdatedV1 access ], data: %s\n", Jsons.DEFAULT.toJson(event.getEvent()));
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
    'drive.file.title_updated_v1': async (data) => {
        console.log(data);
        return 'success';
    },
});

const server = http.createServer();
// 创建路由处理器 Create route handler
server.on('request', lark.adaptDefault('/webhook/event', eventDispatcher));
server.listen(3000);
