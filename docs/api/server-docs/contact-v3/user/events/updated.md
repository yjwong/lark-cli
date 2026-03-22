---
title: "Employee change"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/events/updated"
service: "contact-v3"
resource: "user"
section: "Contacts"
scopes:
  - "contact:contact.base:readonly"
  - "contact:contact:readonly_as_app"
  - "contact:contact:access_as_app"
field_scopes:
  - "contact:contact:readonly_as_app"
  - "contact:user.base:readonly"
  - "contact:user.department:readonly"
  - "contact:user.dotted_line_leader_info.read"
  - "contact:user.employee:readonly"
  - "contact:user.employee_number:read"
  - "contact:user.gender:readonly"
  - "contact:user.job_level:readonly"
  - "contact:user.employee_id:readonly"
  - "contact:user.phone:readonly"
  - "contact:user.email:readonly"
  - "contact:user.job_family:readonly"
  - "contact:contact:access_as_app"
  - "contact:contact:readonly"
updated: "1692861219000"
---

# Employee change

This event is triggered in case of employee changes. The old_object only displays the original values of updated fields.{Usage Examples}(url=/api/tools/api_explore/api_explore_config?project=contact&version=v3&resource=user&event=updated)

> An app will receive the event only when it has the data scope for the fields being changed. For the relationship between data scopes and fields, refer to App scopes, or view the field descriptions in the parameter table for the event body.

## Event
| Facts |  |
| --- | --- |
| Event type | contact.user.updated_v3 |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `contact:contact.base:readonly` `contact:contact:readonly_as_app` `contact:contact:access_as_app` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:contact:readonly_as_app` `contact:user.base:readonly` `contact:user.department:readonly` `contact:user.dotted_line_leader_info.read` `contact:user.employee:readonly` `contact:user.employee_number:read` `contact:user.gender:readonly` `contact:user.job_level:readonly` `contact:user.employee_id:readonly` `contact:user.phone:readonly` `contact:user.email:readonly` `contact:user.job_family:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
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
|   `object` | `user_event` | New information |
|     `open_id` | `string` | User's open_id |
|     `union_id` | `string` | the union id of user User-related ID concepts |
|     `user_id` | `string` | Unique identifier of the user in a tenant **Required field scopes**: `contact:user.employee_id:readonly` |
|     `name` | `string` | User name **Data validation rules**: - Minimum length: `1` characters **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.base:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|     `en_name` | `string` | English name **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.base:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|     `nickname` | `string` | Alias **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.base:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|     `email` | `string` | Email **Required field scopes**: `contact:user.email:readonly` |
|     `enterprise_email` | `string` | enterprise email **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.base:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|     `job_title` | `string` | Job Title |
|     `mobile` | `string` | Mobile number **Required field scopes**: `contact:user.phone:readonly` |
|     `gender` | `int` | Gender **Optional values are**:  - `0`: Unknown - `1`: Male - `2`: Female  **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.gender:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|     `avatar` | `avatar_info` | User's profile photo **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.base:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|       `avatar_72` | `string` | Link of profile photo in 72*72 px |
|       `avatar_240` | `string` | Link of profile photo in 240*240 px |
|       `avatar_640` | `string` | Link of profile photo in 640*640 px |
|       `avatar_origin` | `string` | Link of the original profile photo |
|     `status` | `user_status` | User status **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|       `is_frozen` | `boolean` | Whether the user is frozen. |
|       `is_resigned` | `boolean` | Whether the user is terminated. |
|       `is_activated` | `boolean` | Whether the user is activated. |
|       `is_exited` | `boolean` | Whether the user is voluntary. The user's status will be automatically changed to terminated after the user has left for a specific time of period. |
|       `is_unjoin` | `boolean` | Whether the user has not joined. The user must confirm in person to join the team. |
|     `department_ids` | `string[]` | List of IDs of the user's departments **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.department:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|     `leader_user_id` | `string` | User ID of the user's direct manager **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|     `city` | `string` | City **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|     `country` | `string` | Country **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|     `work_station` | `string` | Seat ID **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|     `join_time` | `int` | Date of employment **Data validation rules**: - Value range: `1` ～ `2147483647` **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|     `employee_no` | `string` | Employee ID **Required field scopes (Satisfy any)**: `contact:user.employee:readonly` `contact:user.employee_number:read` `contact:contact:readonly_as_app` `contact:contact:access_as_app` `contact:contact:readonly` |
|     `employee_type` | `int` | Employee type **Optional values are**:  - `1`: Regular - `2`: Intern - `3`: Outsourcing - `4`: Contractor - `5`: Consultant  **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|     `orders` | `user_order[]` | User sorting information **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.department:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|       `department_id` | `string` | Department ID for the sorting information. The ID value matches the department_id_type in the query parameter. For information about different IDs, see Department IDs. |
|       `user_order` | `int` | The user's ranking in their direct department. A larger value indicates a higher ranking. |
|       `department_order` | `int` | Ranking of the user's departments. A larger value indicates a higher ranking. |
|       `is_primary_dept` | `boolean` | Identifies the user's unique main department. The main department is the department that ranks first among the user's departments (department_order is the max) |
|     `custom_attrs` | `user_custom_attr[]` | Custom properties **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|       `type` | `string` | Custom field type - `TEXT`: Text - `HREF`: Webpage - `ENUMERATION`: Enumeration - `PICTURE_ENUM`: Image - `GENERIC_USER`: User FAQs about custom fields |
|       `id` | `string` | Custom field ID |
|       `value` | `user_custom_attr_value` | Custom field value |
|         `text` | `string` | When the field type is `TEXT`, this parameter defines the field value and is required; when the field type is `HREF`, this parameter defines the webpage title and is required. |
|         `url` | `string` | When the field type is HREF, this parameter defines the default URL. , such as mobile phone jump to applet, PC jump to webpage . |
|         `pc_url` | `string` | When the field type is HREF, this parameter defines URL on the PC. |
|         `option_id` | `string` | When the field type is ENUMERATION or PICTURE_ENUM, this parameter defines the option value. |
|         `option_value` | `string` | Option value |
|         `name` | `string` | Name |
|         `picture_url` | `string` | Image link |
|         `generic_user` | `custom_attr_generic_user` | When the field type is GENERIC_USER, this parameter defines the user referenced. |
|           `id` | `string` | User's user_id |
|           `type` | `int` | User type    1: User |
|     `job_level_id` | `string` | Job level ID **Required field scopes**: `contact:user.job_level:readonly` |
|     `job_family_id` | `string` | Job family ID **Required field scopes**: `contact:user.job_family:readonly` |
|     `dotted_line_leader_user_ids` | `string[]` | dotted_line_leader_info_ids **Required field scopes**: `contact:user.dotted_line_leader_info.read` |
|   `old_object` | `user_event` | Original information |
|     `open_id` | `string` | User's open_id |
|     `union_id` | `string` | the union id of user User-related ID concepts |
|     `user_id` | `string` | Unique identifier of the user in a tenant **Required field scopes**: `contact:user.employee_id:readonly` |
|     `name` | `string` | User name **Data validation rules**: - Minimum length: `1` characters **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.base:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|     `en_name` | `string` | English name **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.base:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|     `nickname` | `string` | Alias **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.base:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|     `email` | `string` | Email **Required field scopes**: `contact:user.email:readonly` |
|     `job_title` | `string` | Job Title |
|     `mobile` | `string` | Mobile number **Required field scopes**: `contact:user.phone:readonly` |
|     `gender` | `int` | Gender **Optional values are**:  - `0`: Unknown - `1`: Male - `2`: Female  **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.gender:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|     `avatar` | `avatar_info` | User's profile photo **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.base:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|       `avatar_72` | `string` | Link of profile photo in 72*72 px |
|       `avatar_240` | `string` | Link of profile photo in 240*240 px |
|       `avatar_640` | `string` | Link of profile photo in 640*640 px |
|       `avatar_origin` | `string` | Link of the original profile photo |
|     `status` | `user_status` | User status **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|       `is_frozen` | `boolean` | Whether the user is frozen. |
|       `is_resigned` | `boolean` | Whether the user is terminated. |
|       `is_activated` | `boolean` | Whether the user is activated. |
|       `is_exited` | `boolean` | Whether the user is voluntary. The user's status will be automatically changed to terminated after the user has left for a specific time of period. |
|       `is_unjoin` | `boolean` | Whether the user has not joined. The user must confirm in person to join the team. |
|     `department_ids` | `string[]` | List of IDs of the user's departments **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.department:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|     `leader_user_id` | `string` | User ID of the user's direct manager **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|     `city` | `string` | City **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|     `country` | `string` | Country **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|     `work_station` | `string` | Seat ID **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|     `join_time` | `int` | Date of employment **Data validation rules**: - Value range: `1` ～ `2147483647` **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|     `employee_no` | `string` | Employee ID **Required field scopes (Satisfy any)**: `contact:user.employee:readonly` `contact:user.employee_number:read` `contact:contact:readonly_as_app` `contact:contact:access_as_app` `contact:contact:readonly` |
|     `employee_type` | `int` | Employee type **Optional values are**:  - `1`: Regular - `2`: Intern - `3`: Outsourcing - `4`: Contractor - `5`: Consultant  **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|     `orders` | `user_order[]` | User sorting information **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.department:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|       `department_id` | `string` | Department ID for the sorting information. The ID value matches the department_id_type in the query parameter. For information about different IDs, see Department IDs. |
|       `user_order` | `int` | The user's ranking in their direct department. A larger value indicates a higher ranking. |
|       `department_order` | `int` | Ranking of the user's departments. A larger value indicates a higher ranking. |
|       `is_primary_dept` | `boolean` | Identifies the user's unique main department. The main department is the department that ranks first among the user's departments (department_order is the max) |
|     `custom_attrs` | `user_custom_attr[]` | Custom properties **Required field scopes (Satisfy any)**: `contact:contact:readonly_as_app` `contact:user.employee:readonly` `contact:contact:access_as_app` `contact:contact:readonly` |
|       `type` | `string` | Custom field type - `TEXT`: Text - `HREF`: Webpage - `ENUMERATION`: Enumeration - `PICTURE_ENUM`: Image - `GENERIC_USER`: User FAQs about custom fields |
|       `id` | `string` | Custom field ID |
|       `value` | `user_custom_attr_value` | Custom field value |
|         `text` | `string` | When the field type is `TEXT`, this parameter defines the field value and is required; when the field type is `HREF`, this parameter defines the webpage title and is required. |
|         `url` | `string` | When the field type is HREF, this parameter defines the default URL. , such as mobile phone jump to applet, PC jump to webpage . |
|         `pc_url` | `string` | When the field type is HREF, this parameter defines URL on the PC. |
|         `option_id` | `string` | When the field type is ENUMERATION or PICTURE_ENUM, this parameter defines the option value. |
|         `option_value` | `string` | Option value |
|         `name` | `string` | Name |
|         `picture_url` | `string` | Image link |
|         `generic_user` | `custom_attr_generic_user` | When the field type is GENERIC_USER, this parameter defines the user referenced. |
|           `id` | `string` | User's user_id |
|           `type` | `int` | User type    1: User |
|     `job_level_id` | `string` | Job level ID **Required field scopes**: `contact:user.job_level:readonly` |
|     `job_family_id` | `string` | Job family ID **Required field scopes**: `contact:user.job_family:readonly` |
|     `dotted_line_leader_user_ids` | `string[]` | dotted_line_leader_info_ids **Required field scopes**: `contact:user.dotted_line_leader_info.read` | ### Event body example

{
    "schema": "2.0",
    "header": {
        "event_id": "5e3702a84e847582be8db7fb73283c02",
        "event_type": "contact.user.updated_v3",
        "create_time": "1608725989000",
        "token": "rvaYgkND1GOiu5MM0E1rncYC6PLtF7JV",
        "app_id": "cli_9f5343c580712544",
        "tenant_key": "2ca1d211f64f6438"
    },
    "event": {
        "object": {
            "open_id": "ou_7dab8a3d3cdcc9da365777c7ad535d62",
            "union_id": "on_2b832a8c219b087c57208039282c0aa3",
            "user_id": "e33ggbyz",
            "name": "John",
            "en_name": "San Zhang",
            "nickname": "Sunny Zhang",
            "email": "zhangsan@gmail.com",
            "enterprise_email": "demo@mail.com",
            "job_title": "software engineer",
            "mobile": "12345678910",
            "gender": 1,
            "avatar": {
                "avatar_72": "https://foo.icon.com/xxxx",
                "avatar_240": "https://foo.icon.com/xxxx",
                "avatar_640": "https://foo.icon.com/xxxx",
                "avatar_origin": "https://foo.icon.com/xxxx"
            },
            "status": {
                "is_frozen": false,
                "is_resigned": false,
                "is_activated": true,
                "is_exited": false,
                "is_unjoin": false
            },
            "department_ids": [
                "od-4e6ac4d14bcd5071a37a39de902c7141"
            ],
            "leader_user_id": "ou_3ghm8a2u0eftg0ff377125s5dd275z09",
            "city": "Hangzhou",
            "country": "China",
            "work_station": "Hangzhou",
            "join_time": 1615381702,
            "employee_no": "e33ggbyz",
            "employee_type": 1,
            "orders": [
                {
                    "department_id": "od-4e6ac4d14bcd5071a37a39de902c7141",
                    "user_order": 100,
                    "department_order": 100,
                    "is_primary_dept": true
                }
            ],
            "custom_attrs": [
                {
                    "type": "TEXT",
                    "id": "DemoId",
                    "value": {
                        "text": "DemoText",
                        "url": "http://www.fs.cn",
                        "pc_url": "http://www.fs.cn",
                        "option_id": "edcvfrtg",
                        "option_value": "option",
                        "name": "name",
                        "picture_url": "https://xxxxxxxxxxxxxxxxxx",
                        "generic_user": {
                            "id": "9b2fabg5",
                            "type": 1
                        }
                    }
                }
            ],
            "job_level_id": "mga5oa8ayjlp9rb",
            "job_family_id": "mga5oa8ayjlp9rb",
            "dotted_line_leader_user_ids": [
                "od-4e6ac4d14bcd5071a37a39de902c7141"
            ]
        },
        "old_object": {
            "open_id": "ou_7dab8a3d3cdcc9da365777c7ad535d62",
            "union_id": "on_2b832a8c219b087c57208039282c0aa3",
            "user_id": "e33ggbyz",
            "name": "John",
            "en_name": "San Zhang",
            "nickname": "Sunny Zhang",
            "email": "zhangsan@gmail.com",
            "job_title": "software engineer",
            "mobile": "12345678910",
            "gender": 1,
            "avatar": {
                "avatar_72": "https://foo.icon.com/xxxx",
                "avatar_240": "https://foo.icon.com/xxxx",
                "avatar_640": "https://foo.icon.com/xxxx",
                "avatar_origin": "https://foo.icon.com/xxxx"
            },
            "status": {
                "is_frozen": false,
                "is_resigned": false,
                "is_activated": true,
                "is_exited": false,
                "is_unjoin": false
            },
            "department_ids": [
                "od-4e6ac4d14bcd5071a37a39de902c7141"
            ],
            "leader_user_id": "ou_3ghm8a2u0eftg0ff377125s5dd275z09",
            "city": "Hangzhou",
            "country": "China",
            "work_station": "Hangzhou",
            "join_time": 1615381702,
            "employee_no": "e33ggbyz",
            "employee_type": 1,
            "orders": [
                {
                    "department_id": "od-4e6ac4d14bcd5071a37a39de902c7141",
                    "user_order": 100,
                    "department_order": 100,
                    "is_primary_dept": true
                }
            ],
            "custom_attrs": [
                {
                    "type": "TEXT",
                    "id": "DemoId",
                    "value": {
                        "text": "DemoText",
                        "url": "http://www.fs.cn",
                        "pc_url": "http://www.fs.cn",
                        "option_id": "edcvfrtg",
                        "option_value": "option",
                        "name": "name",
                        "picture_url": "https://xxxxxxxxxxxxxxxxxx",
                        "generic_user": {
                            "id": "9b2fabg5",
                            "type": 1
                        }
                    }
                }
            ],
            "job_level_id": "mga5oa8ayjlp9rb",
            "job_family_id": "mga5oa8ayjlp9rb",
            "dotted_line_leader_user_ids": [
                "od-4e6ac4d14bcd5071a37a39de902c7141"
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
		OnP2UserUpdatedV3(func(ctx context.Context, event *larkcontact.P2UserUpdatedV3) error {
			fmt.Printf("[ OnP2UserUpdatedV3 access ], data: %s\n", larkcore.Prettify(event))
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

def do_p2_contact_user_updated_v3(data: lark.contact.v3.P2ContactUserUpdatedV3) -> None:
    print(f'[ do_p2_contact_user_updated_v3 access ], data: {lark.JSON.marshal(data, indent=4)}')

# 注册事件 Register event
event_handler = lark.EventDispatcherHandler.builder("", "") \
    .register_p2_contact_user_updated_v3(do_p2_contact_user_updated_v3) \
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
import com.lark.oapi.service.contact.v3.model.P2UserUpdatedV3;
import com.lark.oapi.event.EventDispatcher;
import com.lark.oapi.ws.Client;

// SDK 使用说明 SDK user guide：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/java-sdk-guide/preparations
public class Sample {
    // 注册事件 Register event
    private static final EventDispatcher EVENT_HANDLER = EventDispatcher.newBuilder("", "")
            .onP2UserUpdatedV3(new ContactService.P2UserUpdatedV3Handler() {
                @Override
                public void handle(P2UserUpdatedV3 event) throws Exception {
                    System.out.printf("[ onP2UserUpdatedV3 access ], data: %s\n", Jsons.DEFAULT.toJson(event.getEvent()));
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
        'contact.user.updated_v3': async (data) => {
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
		OnP2UserUpdatedV3(func(ctx context.Context, event *larkcontact.P2UserUpdatedV3) error {
			fmt.Printf("[ OnP2UserUpdatedV3 access ], data: %s\n", larkcore.Prettify(event))
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

def do_p2_contact_user_updated_v3(data: lark.contact.v3.P2ContactUserUpdatedV3) -> None:
    print(f'[ do_p2_contact_user_updated_v3 access ], data: {lark.JSON.marshal(data, indent=4)}')

# 注册事件 Register event
event_handler = lark.EventDispatcherHandler.builder("", "") \
    .register_p2_contact_user_updated_v3(do_p2_contact_user_updated_v3) \
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
import com.lark.oapi.service.contact.v3.model.P2UserUpdatedV3;
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
            .onP2UserUpdatedV3(new ContactService.P2UserUpdatedV3Handler() {
                @Override
                public void handle(P2UserUpdatedV3 event) throws Exception {
                    System.out.printf("[ onP2UserUpdatedV3 access ], data: %s\n", Jsons.DEFAULT.toJson(event.getEvent()));
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
    'contact.user.updated_v3': async (data) => {
        console.log(data);
        return 'success';
    },
});

const server = http.createServer();
// 创建路由处理器 Create route handler
server.on('request', lark.adaptDefault('/webhook/event', eventDispatcher));
server.listen(3000);
