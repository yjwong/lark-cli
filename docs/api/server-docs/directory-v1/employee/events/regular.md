---
title: "To be resigned to regular"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/directory-v1/employee/events/regular"
service: "directory-v1"
resource: "employee"
section: "Organization"
scopes:
  - "directory:employee:read"
field_scopes:
  - "directory:department.base:read"
  - "directory:department.external_id:read"
  - "directory:employee.base.active_status:read"
  - "directory:employee.base.avatar:read"
  - "directory:employee.base.background_image:read"
  - "directory:employee.base.base:read"
  - "directory:employee.base.custom_field:read"
  - "directory:employee.base.department:read"
  - "directory:employee.base.dept_order:read"
  - "directory:employee.base.description:read"
  - "directory:employee.base.dotted_line_leaders:read"
  - "directory:employee.base.email:read"
  - "directory:employee.base.gender:read"
  - "directory:employee.base.is_resigned:read"
  - "directory:employee.base.leader:read"
  - "directory:employee.base.leader_id:read"
  - "directory:employee.base.mobile:read"
  - "directory:employee.base.name.another_name:read"
  - "directory:employee.base.name.name:read"
  - "directory:employee.base.resign_time:read"
  - "directory:employee.base.status:read"
  - "directory:employee.work.base_work:read"
  - "directory:employee.work.work_station:read"
  - "directory:employee.work.employment_type:read"
  - "directory:employee.work.extension_number:read"
  - "directory:employee.work.job_number:read"
  - "directory:employee.work.job_title:read"
  - "directory:employee.work.join_date:read"
  - "directory:employee.work.resign_date:read"
  - "directory:employee.work.resign_reason:read"
  - "directory:employee.work.resign_remark:read"
  - "directory:employee.work.resign_type:read"
  - "directory:employee.work.staff_status:read"
  - "directory:employee.work.work_country_or_region:read"
  - "directory:employee.work.work_place:read"
  - "directory:employee.work.employment:read"
updated: "1756439278000"
---

# To be resigned to regular

To receive the event notification when an employee cancelled resign request. 

## Prerequisites
You need to configure event subscriptions in the application so that you can receive event data when the event is triggered. For more information about event subscriptions, see Event Overview. 

## Precautions
Some fields of this event require permissions. You can refer to the corresponding parameter descriptions to obtain the permissions required for the parameters. Only when the application has the corresponding field permissions enabled can it successfully receive the complete event body data. For the specific operations to apply for permission, see Apply for API Permissions. 

{Usage Examples}(url=/api/tools/api_explore/api_explore_config?project=directory&version=v1&resource=employee&event=regular)

## Event
| Facts |  |
| --- | --- |
| Event type | directory.employee.regular_v1 |
| Supported app types | custom,isv |
| Required scopes | `directory:employee:read` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `directory:department.base:read` `directory:department.external_id:read` `directory:employee.base.active_status:read` `directory:employee.base.avatar:read` `directory:employee.base.background_image:read` `directory:employee.base.base:read` `directory:employee.base.custom_field:read` `directory:employee.base.department:read` `directory:employee.base.dept_order:read` `directory:employee.base.description:read` `directory:employee.base.dotted_line_leaders:read` `directory:employee.base.email:read` `directory:employee.base.gender:read` `directory:employee.base.is_resigned:read` `directory:employee.base.leader:read` `directory:employee.base.leader_id:read` `directory:employee.base.mobile:read` `directory:employee.base.name.another_name:read` `directory:employee.base.name.name:read` `directory:employee.base.resign_time:read` `directory:employee.base.status:read` `directory:employee.work.base_work:read` `directory:employee.work.work_station:read` `directory:employee.work.employment_type:read` `directory:employee.work.extension_number:read` `directory:employee.work.job_number:read` `directory:employee.work.job_title:read` `directory:employee.work.join_date:read` `directory:employee.work.resign_date:read` `directory:employee.work.resign_reason:read` `directory:employee.work.resign_remark:read` `directory:employee.work.resign_type:read` `directory:employee.work.staff_status:read` `directory:employee.work.work_country_or_region:read` `directory:employee.work.work_place:read` `directory:employee.work.employment:read` |
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
|   `employee` | `employee_entity` | Employee entity |
|     `base_info` | `employee_base_entity` | Basic employee information |
|       `employee_id` | `string` | The user's open_id. For the ID type, see User Identity Overview |
|       `name` | `name` | Name |
|         `name` | `i18n_text` | I18n text **Required field scopes (Satisfy any)**: `directory:employee.base.base:read` `directory:employee.base.name.name:read` |
|           `default_value` | `string` | Default value |
|           `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
|             `key` | `string` | International locale |
|             `value` | `string` | value |
|         `another_name` | `string` | Alias **Required field scopes (Satisfy any)**: `directory:employee.base.base:read` `directory:employee.base.name.another_name:read` |
|       `mobile` | `string` | Mobile phone number **Data validation rules**: - Maximum length: `255` characters **Required field scopes**: `directory:employee.base.mobile:read` |
|       `email` | `string` | Contact email **Data validation rules**: - Maximum length: `255` characters **Required field scopes**: `directory:employee.base.email:read` |
|       `gender` | `int` | Gender **Optional values are**:  - `0`: Unknown - `1`: male - `2`: female - `3`: other  **Required field scopes**: `directory:employee.base.gender:read` |
|       `departments` | `department[]` | Department information **Data validation rules**: - Length range: `0` ～ `20` **Required field scopes**: `directory:employee.base.department:read` |
|         `department_id` | `string` | Department ID. The department ID type is open_department_id. For details about the department ID, see Department Resource Introduction. **Data validation rules**: - Maximum length: `64` characters **Required field scopes (Satisfy any)**: `directory:department.base:read` `directory:department.external_id:read` |
|       `employee_order_in_departments` | `user_department_sort_info[]` | The ranking information of users within the department, the first department is the main department **Data validation rules**: - Length range: `0` ～ `20` **Required field scopes (Satisfy any)**: `directory:employee.base.department:read` `directory:employee.base.dept_order:read` |
|         `department_id` | `string` | Department ID. The department ID type is open_department_id. For details about the department ID, see Department Resource Introduction. |
|         `order_weight_in_deparment` | `string` | The ranking weight of users within the department |
|         `order_weight_among_deparments` | `string` | User ranking weights across multiple departments |
|       `description` | `string` | Personal signature **Required field scopes (Satisfy any)**: `directory:employee.base.base:read` `directory:employee.base.description:read` |
|       `active_status` | `int` | User active status **Optional values are**:  - `1`: not activated - `2`: activate - `3`: freeze - `4`: active exit - `5`: not joined  **Required field scopes (Satisfy any)**: `directory:employee.base.active_status:read` `directory:employee.base.status:read` |
|       `is_resigned` | `boolean` | Whether to quit **Required field scopes (Satisfy any)**: `directory:employee.base.is_resigned:read` `directory:employee.base.status:read` |
|       `leader_id` | `string` | The open_id of the user's line manager. For more information on user IDs, see Overview of User Identity. **Required field scopes (Satisfy any)**: `directory:employee.base.leader:read` `directory:employee.base.leader_id:read` |
|       `dotted_line_leader_ids` | `string[]` | The open_id of the user's dotted-line manager. For more information about user IDs, see User Identity Overview. **Data validation rules**: - Length range: `0` ～ `10` **Required field scopes (Satisfy any)**: `directory:employee.base.dotted_line_leaders:read` `directory:employee.base.leader:read` |
|       `custom_field_values` | `custom_field_value[]` | Custom field value **Data validation rules**: - Length range: `0` ～ `100` **Required field scopes**: `directory:employee.base.custom_field:read` |
|         `field_key` | `string` | Custom field key |
|         `field_type` | `string` | Custom field type **Optional values are**:  - `1`: multiline text - `2`: web link - `3`: Enumeration options - `4`: personnel - `10`: Multiple-choice enumeration types (currently only text types are supported) - `11`: personnel list  |
|         `text_value` | `i18n_text` | I18n text |
|           `default_value` | `string` | Default value |
|           `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
|             `key` | `string` | International locale |
|             `value` | `string` | value |
|         `url_value` | `url_value` | Web link field value |
|           `link_text` | `i18n_text` | I18n text |
|             `default_value` | `string` | Default value |
|             `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
|               `key` | `string` | International locale |
|               `value` | `string` | value |
|           `url` | `string` | Mobile end web link |
|           `pcurl` | `string` | Desktop web link |
|         `enum_value` | `enum_value` | enumeration |
|           `enum_ids` | `string[]` | Option result ID **Data validation rules**: - Length range: `0` ～ `100` |
|           `enum_type` | `string` | option type **Optional values are**:  - `1`: Text - `2`: picture  |
|         `user_values` | `user_value[]` | Person field value **Data validation rules**: - Length range: `0` ～ `100` |
|           `ids` | `string[]` | List of user IDs. For more information about user IDs, see User Identity Overview. **Data validation rules**: - Length range: `0` ～ `100` |
|           `user_type` | `string` | Personnel type **Optional values are**:  - `1`: Employee  |
|       `resign_time` | `string` | Turnover time **Required field scopes**: `directory:employee.base.resign_time:read` |
|       `avatar` | `image_link` | Avatar url **Required field scopes (Satisfy any)**: `directory:employee.base.avatar:read` `directory:employee.base.base:read` |
|         `avatar_72` | `string` | 72 * 72 pixel avatar link |
|         `avatar_240` | `string` | 240 * 240 pixel avatar link |
|         `avatar_640` | `string` | 640 * 640 pixel avatar link |
|         `avatar_origin` | `string` | Original avatar link |
|       `background_image` | `string` | Custom background cover url **Required field scopes (Satisfy any)**: `directory:employee.base.background_image:read` `directory:employee.base.base:read` |
|     `work_info` | `employee_work_entity` | Employee job information |
|       `work_country_or_region` | `string` | Work country/region **Required field scopes (Satisfy any)**: `directory:employee.work.base_work:read` `directory:employee.work.work_country_or_region:read` |
|       `work_place` | `place` | Work location, which is the City property in user's Job details **Required field scopes (Satisfy any)**: `directory:employee.work.base_work:read` `directory:employee.work.work_place:read` |
|         `place_id` | `string` | ID |
|       `work_station` | `i18n_text` | Physical seat ID **Required field scopes (Satisfy any)**: `directory:employee.work.base_work:read` `directory:employee.work.work_station:read` |
|         `default_value` | `string` | Default value |
|         `i18n_value` | `map<string, string>` | Internationalized value, key is zh_cn, ja_jp, en_us, value is the corresponding value |
|           `key` | `string` | International locale |
|           `value` | `string` | value |
|       `job_number` | `string` | job number **Data validation rules**: - Maximum length: `255` characters **Required field scopes (Satisfy any)**: `directory:employee.work.base_work:read` `directory:employee.work.job_number:read` |
|       `extension_number` | `string` | Extension number **Data validation rules**: - Maximum length: `99` characters **Required field scopes (Satisfy any)**: `directory:employee.work.base_work:read` `directory:employee.work.extension_number:read` |
|       `join_date` | `string` | Join date, e.g., 2007-03-20 **Required field scopes (Satisfy any)**: `directory:employee.work.join_date:read` `directory:employee.work.employment:read` |
|       `employment_type` | `int` | Employee type **Required field scopes (Satisfy any)**: `directory:employee.work.employment_type:read` `directory:employee.work.employment:read` |
|       `staff_status` | `int` | Employee personnel status **Optional values are**:  - `1`: on the job - `2`: leave - `3`: to be hired - `4`: Cancel onboarding - `5`: to leave  **Required field scopes (Satisfy any)**: `directory:employee.work.staff_status:read` `directory:employee.work.employment:read` |
|       `job_title` | `job_title` | Position **Required field scopes**: `directory:employee.work.job_title:read` |
|         `job_title_id` | `string` | ID **Data validation rules**: - Maximum length: `64` characters |
|       `resign_date` | `string` | Date of resignation, e.g., 2007-03-20 **Required field scopes (Satisfy any)**: `directory:employee.work.resign_date:read` `directory:employee.work.employment:read` |
|       `resign_reason` | `string` | Reason for leaving **Required field scopes (Satisfy any)**: `directory:employee.work.resign_reason:read` `directory:employee.work.employment:read` |
|       `resign_remark` | `string` | Resignation remarks **Data validation rules**: - Maximum length: `5000` characters **Required field scopes (Satisfy any)**: `directory:employee.work.resign_remark:read` `directory:employee.work.employment:read` |
|       `resign_type` | `string` | Type of turnover **Required field scopes (Satisfy any)**: `directory:employee.work.resign_type:read` `directory:employee.work.employment:read` |
|   `abnormal` | `abnormal_record` | Field exception information |
|     `id` | `string` | abnormal entity id |
|     `row_error` | `int` | row-level exception **Optional values are**:  - `0`: success - `1000`: No permission  **Data validation rules**: - Value range: `0` ～ `100` |
|     `field_errors` | `map<string, abnormal_code>` | Column-level exception, key is the field name, value is the following enumeration |
|       `key` | `string` | field name |
|       `value` | `int` | error code **Optional values are**:  - `0`: success - `1000`: No permission - `2000`: field query exception - `2003`: Field does not exist  **Data validation rules**: - Value range: `0` ～ `3000` | ### Event body example

{
    "schema": "2.0",
    "header": {
        "event_id": "cb1315992a8545190aece434f3c5a082",
        "token": "",
        "create_time": "1726146320000",
        "event_type": "directory.employee.regular_v1",
        "tenant_key": "133c1eae3c0f1748",
        "app_id": "cli_a23f3400fe78901b"
    },
    "event": {
        "abnormal": {
            "row_error": 0
        },
        "employee": {
            "base_info": {
                "active_status": 5,
                "avatar": {
                    "avatar_origin": "v3_00el_a1a8606b-fe92-4426-b1d3-7c999eb578dj"
                },
                "custom_field_values": [
                    {
                        "field_key": "C-xxxx",
                        "field_type": 1,
                        "text_value": {
                            "default_value": "213123",
                            "i18n_value": {
                            }
                        }
                    },
                    {
                        "field_key": "C-xxxx",
                        "field_type": 1,
                        "text_value": {
                            "default_value": "123123",
                            "i18n_value": {
                            }
                        }
                    },
                    {
                        "field_key": "C-xxxx",
                        "field_type": 2,
                        "url_value": {
                            "link_text": {
                                "default_value": "Lark",
                                "i18n_value": {
                                }
                            },
                            "url": "https://www.larksuite.com/"
                        }
                    },
                    {
                        "field_key": "C-7090440828405268500",
                        "field_type": 1,
                        "text_value": {
                            "default_value": "21312",
                            "i18n_value": {
                            }
                        }
                    },
                    {
                        "enum_value": {
                            "enum_ids": [
                                "E-xxxx"
                            ],
                            "enum_type": 1
                        },
                        "field_key": "C-xxxx",
                        "field_type": 3
                    },
                    {
                        "field_key": "C-xxxx",
                        "field_type": 1,
                        "text_value": {
                            "default_locale": "ms_my",
                            "default_value": "21321312",
                            "i18n_value": {
                                "ms_my": "21321312"
                            }
                        }
                    },
                    {
                        "enum_value": {
                            "enum_ids": [
                                "E-7128392834016657427"
                            ],
                            "enum_type": 1
                        },
                        "field_key": "C-xxxx",
                        "field_type": 3
                    },
                    {
                        "enum_value": {
                            "enum_ids": [
                                "E-xxxx"
                            ],
                            "enum_type": 2
                        },
                        "field_key": "C-xxxx",
                        "field_type": 3
                    },
                    {
                        "field_key": "C-xxxx",
                        "field_type": 1,
                        "text_value": {
                            "default_locale": "ms_my",
                            "default_value": "23131",
                            "i18n_value": {
                                "ms_my": "23131"
                            }
                        }
                    },
                    {
                        "field_key": "C-xxxx",
                        "field_type": 1,
                        "text_value": {
                            "default_locale": "ms_my",
                            "default_value": "231231",
                            "i18n_value": {
                                "ms_my": "231231"
                            }
                        }
                    },
                    {
                        "field_key": "C-xxxx",
                        "field_type": 1,
                        "text_value": {
                            "default_locale": "ms_my",
                            "default_value": "213213",
                            "i18n_value": {
                                "ms_my": "213213"
                            }
                        }
                    },
                    {
                        "enum_value": {
                            "enum_ids": [
                                "E-xxxx"
                            ],
                            "enum_type": 1
                        },
                        "field_key": "C-xxxx",
                        "field_type": 3
                    },
                    {
                        "enum_value": {
                            "enum_ids": [
                                "E-xxxx"
                            ],
                            "enum_type": 1
                        },
                        "field_key": "C-xxxx",
                        "field_type": 3
                    }
                ],
                "description": "",
                "dotted_line_leader_ids": [
                    "ou_xxxx"
                ],
                "email": "test1@qq.com",
                "employee_id": "ou_xxxx",
                "employee_order_in_departments": [
                    {
                        "department_id": "od-xxxx",
                        "order_weight_among_deparments": 1,
                        "order_weight_in_deparment": 0
                    }
                ],
                "departments": [
                    {
                        "department_id": "od-xxxxx"
                    }
                ],
                "is_resigned": false,
                "leader_id": "ou_xxxx",
                "mobile": "+86130xxxxxxxx",
                "name": {
                    "another_name": "xxxx",
                    "name": {
                        "default_value": "xxx",
                        "i18n_value": {
                            "en_us": "",
                            "ja_jp": "",
                            "zh_cn": "xxxx"
                        }
                    }
                }
            },
            "work_info": {
                "employment_type": 1,
                "job_number": "1111",
                "job_title": {
                    "job_title_id": "8rrl2m8mn7jb157"
                },
                "join_date": "2024-09-12",
                "resign_date": "2024-12-22",
                "resign_reason": 1,
                "resign_remark": "3123213",
                "resign_type": 1,
                "staff_status": 1,
                "work_country_or_region": "MDCT00000040",
                "work_place": {
                    "place_id": "z1rxe3peznm3nk0"
                },
                "work_station": {
                    "default_value": "11111",
                    "i18n_value": {
                        "en_us": "",
                        "ja_jp": "",
                        "zh_cn": "11111"
                    }
                }
            }
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
		OnP2EmployeeRegularV1(func(ctx context.Context, event *larkdirectory.P2EmployeeRegularV1) error {
			fmt.Printf("[ OnP2EmployeeRegularV1 access ], data: %s\n", larkcore.Prettify(event))
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

def do_p2_directory_employee_regular_v1(data: lark.directory.v1.P2DirectoryEmployeeRegularV1) -> None:
    print(f'[ do_p2_directory_employee_regular_v1 access ], data: {lark.JSON.marshal(data, indent=4)}')

# 注册事件 Register event
event_handler = lark.EventDispatcherHandler.builder("", "") \
    .register_p2_directory_employee_regular_v1(do_p2_directory_employee_regular_v1) \
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
import com.lark.oapi.service.directory.v1.model.P2EmployeeRegularV1;
import com.lark.oapi.event.EventDispatcher;
import com.lark.oapi.ws.Client;

// SDK 使用说明 SDK user guide：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/java-sdk-guide/preparations
public class Sample {
    // 注册事件 Register event
    private static final EventDispatcher EVENT_HANDLER = EventDispatcher.newBuilder("", "")
            .onP2EmployeeRegularV1(new DirectoryService.P2EmployeeRegularV1Handler() {
                @Override
                public void handle(P2EmployeeRegularV1 event) throws Exception {
                    System.out.printf("[ onP2EmployeeRegularV1 access ], data: %s\n", Jsons.DEFAULT.toJson(event.getEvent()));
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
        'directory.employee.regular_v1': async (data) => {
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
		OnP2EmployeeRegularV1(func(ctx context.Context, event *larkdirectory.P2EmployeeRegularV1) error {
			fmt.Printf("[ OnP2EmployeeRegularV1 access ], data: %s\n", larkcore.Prettify(event))
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

def do_p2_directory_employee_regular_v1(data: lark.directory.v1.P2DirectoryEmployeeRegularV1) -> None:
    print(f'[ do_p2_directory_employee_regular_v1 access ], data: {lark.JSON.marshal(data, indent=4)}')

# 注册事件 Register event
event_handler = lark.EventDispatcherHandler.builder("", "") \
    .register_p2_directory_employee_regular_v1(do_p2_directory_employee_regular_v1) \
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
import com.lark.oapi.service.directory.v1.model.P2EmployeeRegularV1;
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
            .onP2EmployeeRegularV1(new DirectoryService.P2EmployeeRegularV1Handler() {
                @Override
                public void handle(P2EmployeeRegularV1 event) throws Exception {
                    System.out.printf("[ onP2EmployeeRegularV1 access ], data: %s\n", Jsons.DEFAULT.toJson(event.getEvent()));
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
    'directory.employee.regular_v1': async (data) => {
        console.log(data);
        return 'success';
    },
});

const server = http.createServer();
// 创建路由处理器 Create route handler
server.on('request', lark.adaptDefault('/webhook/event', eventDispatcher));
server.listen(3000);
