---
title: "Obtain user ID via email or mobile number"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/batch_get_id"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/contact/v3/users/batch_get_id"
service: "contact-v3"
resource: "user"
section: "Contacts"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "contact:user.id:readonly"
field_scopes:
  - "contact:user.employee:readonly"
  - "contact:user.employee_id:readonly"
  - "contact:contact:access_as_app"
  - "contact:contact:readonly"
  - "contact:contact:readonly_as_app"
updated: "1710473207000"
---

# Obtain user ID via email or mobile number

This API is used to obtain user ID information via mobile number or email. The supported ID types are open_id, user_id, and union_id, which can be specified by using the query parameters.

> Cases where the user ID is not returned:
> - The queried phone number or email does not exist.
> - The **Obtain user ID** permission is not enabled, resulting in the inability to return the user_id.
> - There is no permission to view user information. You can specify user permission scope in **Permissions & Scopes > Data Permissions > Range of contacts data to access** on the developer backstage application details page.
> - Querying with a corporate email will not return the user ID; the user's email address should be used instead.
> - The Token provided in the Authorization header is incorrect. For example, the Token corresponds to a different application than the one needed.
> - The queried user has left the company.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/contact/v3/users/batch_get_id |
| HTTP Method | POST |
| Rate Limit | 1000 per minute、50 per second |
| Supported app types | custom,isv |
| Required scopes | `contact:user.id:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `contact:user.employee:readonly` `contact:user.employee_id:readonly` `contact:contact:access_as_app` `contact:contact:readonly` `contact:contact:readonly_as_app` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `user_id_type` | `string` | No | User ID categories **Example value**: user_id **Optional values are**:  - `open_id`: Identifies a user to an app. The same user has different Open IDs in different apps. How to get Open ID - `union_id`: Identifies a user to a tenant that acts as a developer. A user has the same Union ID in apps developed by the same developer, and has different Union IDs in apps developed by different developers. A developer can use Union ID to link the same user's identities in multiple apps.How to get Union ID - `user_id`: Identifies a user to a tenant. The same user has different User IDs in different tenants. In one single tenant, a user has the same User ID in all apps （including store apps）. User ID is usually used to communicate user data between different apps. How to get User ID  **Default value**: `open_id` **When the value is `user_id`, the following field scopes are required**: `contact:user.employee_id:readonly` | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `emails` | `string[]` | No | Emails of the users to be queried( Enterprise email is not supported ). A maximum of 50 emails are supported. Note that emails and mobiles are independent of each other, and each user mailbox returns the corresponding user ID. The number of user IDs returned by this interface is the sum of the number of emails and the number of mobiles. **Example value**: ["abc@z.com"] **Data validation rules**: - Maximum length: `50` |
| `mobiles` | `string[]` | No | The mobile phone number of the user to be queried, up to 50 entries. Notice 1. Emails and mobiles are independent of each other, and each user's mobile phone number returns the corresponding user ID. 2. Mobile phone numbers outside of mainland China need to add a country/region code beginning with "+". **Example value**: ["13812345678"] **Data validation rules**: - Maximum length: `50` |
| `include_resigned` | `boolean` | No | Check whether the query results include resigned employees. If the value is true, the ID of the resigned user can be queried. **Example value**: true **Default value**: `false` | ### Request body example

{
    "emails": [
"zhangsan@z.com","lisi@a.com"
    ],
    "mobiles": [
"13812345678","13812345679"
    ],
"include_resigned":true
}

**Go Request Example**
```go
import (
	"context"

	"github.com/larksuite/oapi-sdk-go/v3"
	"github.com/larksuite/oapi-sdk-go/v3/service/contact/v3"
)

func main() {
	client := lark.NewClient("appID", "appSecret")

	req := larkcontact.NewBatchGetIdUserReqBuilder().
		UserIdType(`open_id`).
		Body(larkcontact.NewBatchGetIdUserReqBodyBuilder().
			Emails([]string{`zhangsan@z.com`, `lisi@a.com`}).
			Mobiles([]string{`13812345678`, `13812345679`}).
			Build()).
		Build()

	resp, err := client.Contact.User.BatchGetId(context.Background(), req)
}
```

**Java Request Example**
```java
import com.lark.oapi.Client;
import com.lark.oapi.core.request.RequestOptions;
import com.lark.oapi.service.contact.v3.model.BatchGetIdUserReq;
import com.lark.oapi.service.contact.v3.model.BatchGetIdUserReqBody;
import com.lark.oapi.service.contact.v3.model.BatchGetIdUserResp;

public class Main {

    public static void main(String arg[]) throws Exception {
        Client client = Client.newBuilder("appId", "appSecret").build();

        BatchGetIdUserReq req = BatchGetIdUserReq.newBuilder()
                .userIdType("open_id")
                .batchGetIdUserReqBody(BatchGetIdUserReqBody.newBuilder()
                        .emails(new String[]{"zhangsan@z.com", "lisi@a.com"})
                        .mobiles(new String[]{"13812345678", "13812345679"})
                        .build())
                .build();

        BatchGetIdUserResp resp = client.contact().user().batchGetId(req, RequestOptions.newBuilder().build());
    }
}
```

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `user_list` | `user_contact_info[]` | User ID for the mobile number or email. |
|     `user_id` | `string` | User ID, whose value is the type specified by user_id_type. If the queried mobile number or email does not exist, or if you have no scope to view the user, this field is empty. |
|     `mobile` | `string` | Mobile number, which will be returned when the query is performed by mobile number. |
|     `email` | `string` | Email, which will be returned when the query is performed by email. |
|     `status` | `user_status` | user status **Required field scopes (Satisfy any)**: `contact:user.employee:readonly` `contact:contact:access_as_app` `contact:contact:readonly` `contact:contact:readonly_as_app` |
|       `is_frozen` | `boolean` | Is it frozen? |
|       `is_resigned` | `boolean` | whether to quit |
|       `is_activated` | `boolean` | Whether to activate |
|       `is_exited` | `boolean` | Whether to take the initiative to quit, the user will automatically turn to have left after taking the initiative to quit for a period of time |
|       `is_unjoin` | `boolean` | Whether you have not joined, you need to confirm by the user before joining the team. | ### Response body example

{
	"code": 0,
	"msg": "success",
	"data": {
		"user_list": [{
				"user_id": "ou_979112345678741d29069abcdef089d4",
				"email": "zhanxxxxx@a.com",
				"status": {
					"is_frozen": false,
					"is_resigned": false,
					"is_activated": true,
					"is_exited": false,
					"is_unjoin": false
				}
			}, {
				"user_id": "ou_919112245678741d29069abcdef096af",
				"email": "lisixxxx@a.com",
				"status": {
					"is_frozen": false,
					"is_resigned": false,
					"is_activated": true,
					"is_exited": false,
					"is_unjoin": false
				}
			},
			{
				"user_id": "ou_46a087654321a1dc920ffab8fedc823f",
				"mobile": "13xxxxxx678",
				"status": {
					"is_frozen": false,
					"is_resigned": false,
					"is_activated": true,
					"is_exited": false,
					"is_unjoin": false
				}
			}, {
				"user_id": "ou_01b081675121a1dc920ffab97cdc017a",
				"mobile": "138xxxxxx79",
				"status": {
					"is_frozen": false,
					"is_resigned": false,
					"is_activated": true,
					"is_exited": false,
					"is_unjoin": false
				}
			}
		]
	}
}
