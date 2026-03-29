---
title: "Subscribe Docs events"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file/subscribe"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/drive/v1/files/:file_token/subscribe"
service: "drive-v1"
resource: "file"
section: "Docs"
rate_limit: "1000 per minute、50 per second"
scopes:
  - "docs:doc"
  - "docs:event:subscribe"
  - "drive:drive"
  - "sheets:spreadsheet"
updated: "1768985675000"
---

# Subscribe Docs events

This interface is used to subscribe to various notification events of cloud documents. To understand the configuration process and usage scenarios for event subscriptions, refer to Event Overview. To learn about the types of events supported by cloud documents, refer to Event List.

## Notes
- Document managers can only receiveFile edited, Base app field changed, and Base app record changed events.
- If the application subscribes to events using the application identity (`tenant_access_token`), it is necessary to request permissions for both the application and user identities to receive events. For example, to subscribe to the file created in folder event, you need to enable the `space:document.event:read` permission for both identities in the Developer console.

## Prerequisites
- Notification events of documents can only be subscribed to by document owners and document managers. Before calling the interface, ensure that the application or user has the relevant permissions.
- Before calling this interface, ensure that the subscription method is correctly configured and events have been added. For details, refer to Configure Subscription Method and Add Events.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/drive/v1/files/:file_token/subscribe |
| HTTP Method | POST |
| Rate Limit | 1000 per minute、50 per second |
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
| `file_type` | `string` | Yes | Document type **Example value**: docx **Optional values are**:  - `doc`: Depracated version of document. - `docx`: Upgraded Docs - `sheet`: Spreadsheet - `bitable`: Bitable - `file`: File - `folder`: Folder - `slides`: Slides  |
| `event_type` | `string` | No | Event type. - If `file_type` is `folder`, this field must be filled in and set to `file.created_in_folder_v1`, indicating the subscription to the file created in folder event. - If `file_type` is not `folder`, do not fill in this field. For cloud document types such as Docs, Sheets, and Base, only subscribing to all related cloud document events is currently supported. Subscribing to specific events under a particular cloud document type is not yet supported. **Example value**: file.created_in_folder_v1 | ## Response

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
| 500 | 1069601 | fail | Please try again. If the problem persists, contact [technical support](https://applink.larksuite.com/client/helpdesk) |
| 400 | 1069602 | param error | Check whether the parameter is valid. |
| 403 | 1069603 | forbidden | Unauthorized operation may be due to the following reasons: * The calling identity has no manage permission * The tenant has not turned on the document access recording function Refer to the following methods to authorize the calling identity manage permission: - If you are using a `tenant_access_token`, it means the current application does not have the manage permission. You need to add permissions for the application through the cloud document webpage by navigating to the top right corner **"..."** -> **"... More"** -> **"Add Application"**. ![](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/bb60f97ebb402475f2af1d3131d4914f_sLOzoqYRXX.png?height=278&maxWidth=550&width=1383) **Note**: Before adding a document application, you need to ensure that the target application has at least one cloud document API permission enabled. Otherwise, you will not be able to search for the target application in the document application window. ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/95bc56b2990541ad83cc65ddaa6d1e3a_V24oiRrz7V.png?maxWidth=550) - If you are using a `user_access_token`, it means the current user does not have the manage permission. You need to add permissions for the current user through the **Share** entry in the top right corner of the cloud document webpage. ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/caceea2ac91c170555194d7a8dc2a317_GfTRc9xLAt.png?height=1992&maxWidth=550&width=3278) For more details on the specific steps or other methods to add permissions, refer to Cloud Document FAQ 3. |
| 400 | 1069604 | document not found | Check whether the file exists. |
