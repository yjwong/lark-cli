---
title: "Delete User Group"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/group/delete"
method: "DELETE"
api_path: "https://open.larksuite.com/open-apis/contact/v3/group/:group_id"
service: "contact-v3"
resource: "group"
section: "Contacts"
scopes:
  - "contact:group"
updated: "1647437749000"
---

# Delete a user group

This API is used to delete a user group from a company. Note that the app's contact scope must be "All employees" so as to delete a user group. Click to learn about the range of contacts data that an app can access.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/contact/v3/group/:group_id |
| HTTP Method | DELETE |
| Supported app types | custom |
| Required scopes | `contact:group` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `group_id` | `string` | ID of the user group to be deleted **Example value**: "g1837191" | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions | ### Response body example

```json
{
    "code": 0,
    "msg": "success",
    "data": {}
}
```

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 500 | 40003 | internal error | Internal error. Please provide the X-Request-Id to our agent. [Contact support](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 400 | 42002 | invalid group_id | Invalid user group ID. |
| 403 | 42009 | no user group authority error | No user group scope. The app's contact scope must include this user group or be "All employees". Click to learn more |
| 400 | 42017 | group has member not allow delete | group has member not allow delete |
| 400 | 42015 | user group disable | The user group function is not enabled. Please contact the Lark agent. [Contact support](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
