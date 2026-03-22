---
title: "Department Deleted"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/directory-v1/department/events/deleted"
service: "directory-v1"
resource: "department"
section: "Organization"
scopes:
  - "directory:department:read"
field_scopes:
  - "directory:department.base:read"
  - "directory:department.custom_field:read"
  - "directory:department.external_id:read"
  - "directory:department.leader:read"
  - "directory:department.name:read"
  - "directory:department.order_weight:read"
  - "directory:department.organization:read"
  - "directory:department.parent_id:read"
  - "directory:department.status:read"
updated: "1756439308000"
---

# Department Deleted

To receive the event notification when a department is deleted. 

## Prerequisites
You need to configure event subscriptions in the application so that you can receive event data when the event is triggered. For more information about event subscriptions, see Event Overview. 

## Precautions
Some fields of this event require permissions. You can refer to the corresponding parameter descriptions to obtain the permissions required for the parameters. Only when the application has the corresponding field permissions enabled can it successfully receive the complete event body data. For the specific operations to apply for permission, see Apply for API Permissions. 

{Usage Examples}(url=/api/tools/api_explore/api_explore_config?project=directory&version=v1&resource=department&event=deleted)

## Event
| Facts |  |
| --- | --- |
| Event type | directory.department.deleted_v1 |
| Supported app types | custom,isv |
| Required scopes | `directory:department:read` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `directory:department.base:read` `directory:department.custom_field:read` `directory:department.external_id:read` `directory:department.leader:read` `directory:department.name:read` `directory:department.order_weight:read` `directory:department.organization:read` `directory:department.parent_id:read` `directory:department.status:read` |
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
|   `department` | `department` | Department information |
|     `department_id` | `string` | Department ID. The department ID type is open_department_id. For details about the department ID, see Department Resource Introduction. **Data validation rules**: - Maximum length: `64` characters **Required field scopes (Satisfy any)**: `directory:department.base:read` `directory:department.external_id:read` |
|     `leaders` | `department_leader[]` | department head **Data validation rules**: - Length range: `0` ～ `100` **Required field scopes**: `directory:department.leader:read` |
|       `leader_type` | `int` | Type of department head **Optional values are**:  - `1`: main - `2`: vice  |
|       `leader_id` | `string` | Department head ID |
|     `parent_department_id` | `string` | Parent Department ID **Data validation rules**: - Maximum length: `64` characters **Required field scopes (Satisfy any)**: `directory:department.organization:read` `directory:department.parent_id:read` |
|     `name` | `i18n_text` | I18n text **Required field scopes (Satisfy any)**: `directory:department.base:read` `directory:department.name:read` |
|       `default_value` | `string` | Default value |
|       `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
|         `key` | `string` | International locale |
|         `value` | `string` | value |
|     `enabled_status` | `boolean` | Whether to enable **Required field scopes**: `directory:department.status:read` |
|     `order_weight` | `string` | Department ranking weight **Required field scopes (Satisfy any)**: `directory:department.order_weight:read` `directory:department.organization:read` |
|     `custom_field_values` | `custom_field_value[]` | Custom field **Data validation rules**: - Length range: `0` ～ `100` **Required field scopes**: `directory:department.custom_field:read` |
|       `field_key` | `string` | Custom field key |
|       `field_type` | `string` | Custom field type **Optional values are**:  - `1`: multiline text - `2`: web link - `3`: Enumeration options - `4`: personnel - `10`: Multiple-choice enumeration types (currently only text types are supported) - `11`: personnel list  |
|       `text_value` | `i18n_text` | I18n text |
|         `default_value` | `string` | Default value |
|         `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
|           `key` | `string` | International locale |
|           `value` | `string` | value |
|       `url_value` | `url_value` | Web link field value |
|         `link_text` | `i18n_text` | I18n text |
|           `default_value` | `string` | Default value |
|           `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
|             `key` | `string` | International locale |
|             `value` | `string` | value |
|         `url` | `string` | Mobile end web link |
|         `pcurl` | `string` | Desktop web link |
|       `enum_value` | `enum_value` | enumeration |
|         `enum_ids` | `string[]` | Option result ID **Data validation rules**: - Length range: `0` ～ `100` |
|         `enum_type` | `string` | option type **Optional values are**:  - `1`: Text - `2`: picture  |
|       `user_values` | `user_value[]` | Person field value **Data validation rules**: - Length range: `0` ～ `100` |
|         `ids` | `string[]` | List of user IDs. For more information about user IDs, see User Identity Overview. **Data validation rules**: - Length range: `0` ～ `100` |
|         `user_type` | `string` | Personnel type **Optional values are**:  - `1`: Employee  |
|   `abnormal` | `abnormal_record` | Field exception information |
|     `id` | `string` | abnormal entity id |
|     `row_error` | `int` | row-level exception **Optional values are**:  - `0`: success - `1000`: No permission  **Data validation rules**: - Value range: `0` ～ `100` |
|     `field_errors` | `map<string, abnormal_code>` | Column-level exception, key is the field name, value is the following enumeration |
|       `key` | `string` | field name |
|       `value` | `int` | error code **Optional values are**:  - `0`: success - `1000`: No permission - `2000`: field query exception - `2003`: Field does not exist  **Data validation rules**: - Value range: `0` ～ `3000` | ### Event body example

{
    "schema": "2.0",
    "header": {
        "event_id": "b47727586e16bc609ca6d3b4c356a5d9",
        "token": "",
        "create_time": "1726211309000",
        "event_type": "directory.department.deleted_v1",
        "tenant_key": "133c1eae3c0f1748",
        "app_id": "cli_a23f3400fe78901b"
    },
    "event": {
        "abnormal": {
            "row_error": 0
        },
        "department": {
            "enabled_status": true,
            "leaders": [
                {
                    "leader_id": "ou_xxxx",
                    "leader_type": 1
                },
                {
                    "leader_id": "ou_xxxx",
                    "leader_type": 2
                }
            ],
            "name": {
                "default_value": "11111223231",
                "i18n_value": {
                    "en_us": "",
                    "ja_jp": "",
                    "zh_cn": "11111223231"
                }
            },
            "order_weight": 4000,
            "parent_department_id": "od-xxxx"
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
	"github.com/larksuite/oapi-sdk-go/v3/service/directory/v1"
	larkws "github.com/larksuite/oapi-sdk-go/v3/ws"
)

// SDK 使用说明 SDK user guide：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/golang-sdk-guide/preparations
func main() {
	// 注册事件 Register event
	eventHandler := dispatcher.NewEventDispatcher("", "").
		OnP2DepartmentDeletedV1(func(ctx context.Context, event *larkdirectory.P2DepartmentDeletedV1) error {
			fmt.Printf("[ OnP2DepartmentDeletedV1 access ], data: %s\n", larkcore.Prettify(event))
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

def do_p2_directory_department_deleted_v1(data: lark.directory.v1.P2DirectoryDepartmentDeletedV1) -> None:
    print(f'[ do_p2_directory_department_deleted_v1 access ], data: {lark.JSON.marshal(data, indent=4)}')

# 注册事件 Register event
event_handler = lark.EventDispatcherHandler.builder("", "") \
    .register_p2_directory_department_deleted_v1(do_p2_directory_department_deleted_v1) \
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
import com.lark.oapi.service.directory.DirectoryService;
import com.lark.oapi.service.directory.v1.model.P2DepartmentDeletedV1;
import com.lark.oapi.event.EventDispatcher;
import com.lark.oapi.ws.Client;

// SDK 使用说明 SDK user guide：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/java-sdk-guide/preparations
public class Sample {
    // 注册事件 Register event
    private static final EventDispatcher EVENT_HANDLER = EventDispatcher.newBuilder("", "")
            .onP2DepartmentDeletedV1(new DirectoryService.P2DepartmentDeletedV1Handler() {
                @Override
                public void handle(P2DepartmentDeletedV1 event) throws Exception {
                    System.out.printf("[ onP2DepartmentDeletedV1 access ], data: %s\n", Jsons.DEFAULT.toJson(event.getEvent()));
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
        'directory.department.deleted_v1': async (data) => {
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
	"github.com/larksuite/oapi-sdk-go/v3/service/directory/v1"
)

// SDK 使用说明 SDK user guide：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/golang-sdk-guide/preparations
func main() {
	// 注册事件 Register event
	eventHandler := dispatcher.NewEventDispatcher("", "").
		OnP2DepartmentDeletedV1(func(ctx context.Context, event *larkdirectory.P2DepartmentDeletedV1) error {
			fmt.Printf("[ OnP2DepartmentDeletedV1 access ], data: %s\n", larkcore.Prettify(event))
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

def do_p2_directory_department_deleted_v1(data: lark.directory.v1.P2DirectoryDepartmentDeletedV1) -> None:
    print(f'[ do_p2_directory_department_deleted_v1 access ], data: {lark.JSON.marshal(data, indent=4)}')

# 注册事件 Register event
event_handler = lark.EventDispatcherHandler.builder("", "") \
    .register_p2_directory_department_deleted_v1(do_p2_directory_department_deleted_v1) \
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
import com.lark.oapi.service.directory.DirectoryService;
import com.lark.oapi.service.directory.v1.model.P2DepartmentDeletedV1;
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
            .onP2DepartmentDeletedV1(new DirectoryService.P2DepartmentDeletedV1Handler() {
                @Override
                public void handle(P2DepartmentDeletedV1 event) throws Exception {
                    System.out.printf("[ onP2DepartmentDeletedV1 access ], data: %s\n", Jsons.DEFAULT.toJson(event.getEvent()));
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
    'directory.department.deleted_v1': async (data) => {
        console.log(data);
        return 'success';
    },
});

const server = http.createServer();
// 创建路由处理器 Create route handler
server.on('request', lark.adaptDefault('/webhook/event', eventDispatcher));
server.listen(3000);
