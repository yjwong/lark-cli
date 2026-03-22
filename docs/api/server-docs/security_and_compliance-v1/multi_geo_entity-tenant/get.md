---
title: "Get a list of geographic locations"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/security_and_compliance-v1/multi_geo_entity-tenant/get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/security_and_compliance/v1/multi_geo_entity/tenant"
service: "security_and_compliance-v1"
resource: "multi_geo_entity-tenant"
section: "security_and_compliance"
scopes:
  - "security_and_compliance:multi_geo_entity.tenant:readonly"
  - "security_and_compliance:user_migration:multi-geo"
updated: "1752734935000"
---

# Get a list of geographic locations

Get a list of geographic locations available to your organization.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/security_and_compliance/v1/multi_geo_entity/tenant |
| HTTP Method | GET |
| Supported app types | custom |
| Required scopes Enable any scope from the list | `security_and_compliance:multi_geo_entity.tenant:readonly` `security_and_compliance:user_migration:multi-geo` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `tenant` | `tenant` | Tenant |
|     `available_geo_locations` | `string[]` | A list of available geographic locations | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "tenant": {
            "available_geo_locations": [
                "us"
            ]
        }
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 1781001 | Your request contains an invalid request parameter. | Parameter error. Check the input parameters according to the appropriate document. |
| 403 | 1781002 | Lack of necessary permissions. | Apply for Multi-Geo permission. |
| 400 | 1781003 | Multi-Geo is not enabled for this tenant. | Activate Multi-Geo service for your organization. |
| 500 | 1782001 | Internal service error. | Contact customer service for assistance. |
