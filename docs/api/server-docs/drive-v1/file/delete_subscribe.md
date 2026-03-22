---
title: "Unsubscribe Docs events"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file/delete_subscribe"
method: "DELETE"
api_path: "https://open.larksuite.com/open-apis/drive/v1/files/:file_token/delete_subscribe"
service: "drive-v1"
resource: "file"
section: "Docs"
rate_limit: "10 per minute"
scopes:
  - "docs:doc"
  - "docs:event:subscribe"
  - "drive:drive"
  - "sheets:spreadsheet"
updated: "1768984539000"
---

# Unsubscribe Document events

This interface is used to unsubscribe from notification events of cloud documents. To understand the configuration process and usage scenarios for event subscriptions, refer to Event Overview. To learn about the types of events supported by cloud documents, refer to Event List.

## Precautions
Currently, it only supports unsubscribing from all document events in the event list; it does not support specifying events to unsubscribe from.

## Prerequisites
- Before calling the interface, ensure that the application or user is the document owner or document manager.
- Before calling this interface, ensure that the subscription method is correctly configured and events have been added. For details, refer to Configure Subscription Method and Add Events.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/drive/v1/files/:file_token/delete_subscribe |
| HTTP Method | DELETE |
| Rate Limit | 10 per minute |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `docs:doc` `docs:event:subscribe` `drive:drive` `sheets:spreadsheet` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `file_token` | `string` | Document token. For details, see FAQs. **Example value**: "doccnfYZzTlvXqZIGTdAHKabcef" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `file_type` | `string` | Yes | Document type **Example value**: docx **Optional values are**:  - `doc`: Depracated version of document - `docx`: Upgraded Docs - `sheet`: Spreadsheet - `bitable`: Bitable - `file`: File - `folder`: Folder - `slides`: Slides  |
| `event_type` | `string` | No | Event type, required when subscribing to the folder type and required to be `file.created_in_folder_v1` **Example value**: file.created_in_folder_v1 | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {

    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 500 | 1069601 | fail | Please try again. If the problem persists, contact [technical support](https://applink.larksuite.com/client/helpdesk) |
| 400 | 1069602 | param error | Check whether the parameter is valid. |
| 403 | 1069603 | forbidden | Unauthorized operation may be due to the following reasons: * The calling identity has no manage permission * The tenant has not turned on the document access recording function Refer to the following methods to authorize the calling identity manage permission: - If you are using a `tenant_access_token`, it means the current application does not have the manage permission. You need to add permissions for the application through the cloud document webpage by navigating to the top right corner **"..."** -> **"... More"** -> **"Add Application"**. ![](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/bb60f97ebb402475f2af1d3131d4914f_sLOzoqYRXX.png?height=278&maxWidth=550&width=1383) **Note**: Before adding a document application, you need to ensure that the target application has at least one cloud document API permission enabled. Otherwise, you will not be able to search for the target application in the document application window. ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/95bc56b2990541ad83cc65ddaa6d1e3a_V24oiRrz7V.png?maxWidth=550) - If you are using a `user_access_token`, it means the current user does not have the manage permission. You need to add permissions for the current user through the **Share** entry in the top right corner of the cloud document webpage. ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/caceea2ac91c170555194d7a8dc2a317_GfTRc9xLAt.png?height=1992&maxWidth=550&width=3278) For more details on the specific steps or other methods to add permissions, refer to Cloud Document FAQ 3. |
| 400 | 1069604 | document not found | Check whether the file exists. |
| 400 | 1069605 | Unsubscribed | No subscription document |
| 400 | 1064230 | locked_for_data_migration | Data migrating, temporarily unable to upload. |
