---
title: "Batch remove user group members"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/group-member/batch_remove"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/contact/v3/group/:group_id/member/batch_remove"
service: "contact-v3"
resource: "group-member"
section: "Contacts"
rate_limit: "100 per minute"
scopes:
  - "contact:group"
updated: "1721372748000"
---

# Batch remove user group members

Call this interface to remove one or more members from the specified common user group.

## Precautions

- Currently, only user-type members can be removed, and department-type members cannot be removed.

- If the address book permission scope of the application is **All Employees**, any user can be removed from any user group. If the address book permission scope of the application is not **All Employees**, the user to be removed and the corresponding user group must be within the address book permission scope of the application. For information about the address book permission scope, please refer to Introduction to permission scope resources.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/contact/v3/group/:group_id/member/batch_remove |
| HTTP Method | POST |
| Rate Limit | 100 per minute |
| Supported app types | custom |
| Required scopes | `contact:group` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `group_id` | `string` | User group ID. The user group ID can be obtained from the return value when creating the user group. You can also call the Query User Group List interface to obtain the user group ID. **Example value**: "test_group" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `members` | `memberlist[]` | Yes | Member information to be removed. **Data validation rules**: - Length range: `1` ～ `100` |
|   `member_id` | `string` | Yes | The ID of the user to be removed, the ID type is consistent with the value of member_id_type. You can call the Query User Group Member List interface to obtain the member IDs in the user group, and pass the member ID to be removed into the current parameter. Note that only members of the user type can be removed, and the same user ID type must be used, otherwise an error will be reported. **Example value**: "u287xj12" |
|   `member_type` | `string` | Yes | The type of user group member. Currently, only user is supported, indicating the user type. **Example value**: "user" |
|   `member_id_type` | `string` | No | When `member_type` is set to `user`, this parameter is required and the user ID type must be set through this parameter, including: - open_id: Identifies the identity of a user in a certain application. The Open ID of the same user in different applications is different. - union_id: Identifies the identity of a user under a certain application developer. The Union ID of the same user in applications under the same developer is the same, and the Union ID of applications under different developers is different. Through the Union ID, application developers can associate the identities of the same user in multiple applications. - user_id: Identifies the identity of a user in a certain tenant. The User ID of the same user in tenant A and tenant B is different. In the same tenant, the User ID of a user is consistent in all applications. User ID is mainly used to connect user data between different applications. **Example value**: "user_id" | ### Request body example

{
    "members": [
        {
            "member_id": "u287xj12",
            "member_type": "user",
            "member_id_type": "user_id"
        }
    ]
}

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {}
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 40001 | param error | Parameter error. Please check whether the parameters passed in during the request are correct. |
| 500 | 40003 | internal error | Internal error. Please obtain the X-Request-Id of the request and provide feedback to [Technical Support](https://applink.larksuite.com/TLJpeNdW). |
| 400 | 40022 | invalid member_id_type | Invalid member ID type. You need to set the correct value according to the member_id_type parameter description in the interface documentation. |
| 400 | 41073 | invalid member_id | The member ID is invalid. You need to refer to the member_id parameter description in the interface documentation and enter the correct member ID. |
| 400 | 41074 | invalid member_type | Invalid member type. You need to set the correct value according to the member_type parameter description in the interface documentation. |
| 400 | 42002 | invalid group_id | The user group ID is invalid. You can call the query user group list interface to obtain the correct user group ID. |
| 403 | 42009 | no group authority | The user group permission is missing. The address book permission scope of the application must include the user group of the current operation. To learn about the address book permission scope of the application, please refer to Introduction to permission scope resources. |
| 403 | 41050 | no user authority error | Lack of user permissions. The app's address book permission scope needs to include the current user. To learn about the app's address book permission scope, see Introduction to permission scope resources. | For more error code information, see General error codes.
