---
title: "Get Tenant Information"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/tenant-v2/tenant/query"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/tenant/v2/tenant/query"
service: "tenant-v2"
resource: "tenant"
section: "Company Information"
scopes:
  - "tenant:tenant:readonly"
field_scopes:
  - "tenant:tenant.domain:read"
updated: "1701656970000"
---

# Obtain company information

Obtains company information such as the company name and the company ID.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/tenant/v2/tenant/query |
| HTTP Method | GET |
| Supported app types | custom,isv |
| Required scopes | `tenant:tenant:readonly` |
| Required field scopes | > The response body of the API contains the following sensitive fields, and they will be returned only after corresponding scopes are added. If you do not need the fields, it is not recommended that you request the scopes. `tenant:tenant.domain:read` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `tenant` | `tenant` | Company information |
|     `name` | `string` | Name of the company |
|     `display_id` | `string` | Company ID. It uniquely identifies the company on the platform. |
|     `tenant_tag` | `int` | A tag that identifies Personal Edition or Team Edition **Optional values are**:  - `0`: Team Edition - `2`: Personal Edition  |
|     `tenant_key` | `string` | Company identifier |
|     `avatar` | `avatar` | Profile photo of the company |
|       `avatar_origin` | `string` | Profile photo of the company |
|       `avatar_72` | `string` | Profile photo of the company in 72x72 pixels |
|       `avatar_240` | `string` | Profile photo of the company in 240x240 pixels |
|       `avatar_640` | `string` | Profile photo of the company in 640x640 pixels |
|     `domain` | `string` | The complete domain name of the enterprise. Enterprise domain names can be used by enterprise members to access web pages containing URL addresses such as the management backend and cloud documents. **Required field scopes**: `tenant:tenant.domain:read` | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "tenant": {
            "name": "Name of the company",
            "display_id": "F123456789",
            "tenant_tag": 0,
            "tenant_key": "abcdefghi",
            "avatar": {
                "avatar_origin": "https://foo.icon.com/xxxx",
                "avatar_72": "https://foo.icon.com/xxxx",
                "avatar_240": "https://foo.icon.com/xxxx",
                "avatar_640": "https://foo.icon.com/xxxx"
            },
            "domain": "example.larksuite.com"
        }
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 404 | 1184000 | The enterprise does not exist | Confirm whether tenant_access_key is correct |
| 403 | 1184001 | No permission to obtain corporate information | If the ISV application is pre-installed and the enterprise has not used this application within 180 days, the enterprise information cannot be obtained through this interface. |
