---
title: "Update User Group"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/group/patch"
method: "PATCH"
api_path: "https://open.larksuite.com/open-apis/contact/v3/group/:group_id"
service: "contact-v3"
resource: "group"
section: "Contacts"
scopes:
  - "contact:group"
updated: "1647437749000"
---

# Update a user group

This API is used to update user group information. Note that the app's contact scope must be "All employees" so as to update a user group. Click to learn about the range of contacts data that an app can access.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/contact/v3/group/:group_id |
| HTTP Method | PATCH |
| Supported app types | custom |
| Required scopes | `contact:group` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `group_id` | `string` | User group ID **Example value**: "g187131" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `name` | `string` | No | User group name, which is unique in a company. The maximum length is 100 characters. **Example value**: "IT outsourcing user group" |
| `description` | `string` | No | User group description Maximum length: 500 characters **Example value**: "IT outsourcing user group, for which fine-grained permission control is required" | ### Request body example

```json
{
    "name": "IT outsourcing user group",
    "description": "IT outsourcing user group, for which fine-grained permission control is required"
}
```

## Response

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
| 400 | 42013 | group name exceed limit | User group name exceeds the maximum length of 100 characters. |
| 400 | 42014 | group description exceed limit | User group description exceeds the maximum length of 500 characters. |
| 403 | 42009 | no user group authority error | No user group scope. The app's contact scope must include this user group or be "All employees". Click to learn more |
| 400 | 42015 | user group disable | The user group function is not enabled. Please contact the Lark agent. [Contact support](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952&extra=%7B%22channel%22%3A14%2C%22created_at%22%3A1614493146%2C%22scenario_id%22%3A6885151765134622721%2C%22signature%22%3A%22ca94c408b966dc1de2083e5bbcd418294c146e98%22%7D) |
| 400 | 47009 | duplicated name error | User group name must not be duplicated in a company. |
