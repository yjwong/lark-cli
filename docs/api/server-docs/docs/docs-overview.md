---
title: "Docs Overview"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/docs-overview"
section: "Docs"
updated: "1647001231000"
---

# Docs Overview

## Services

Docs is the general name for the system of online Lark docs, sheets, Bitable, Wiki, and Drive. You can call the relevant Docs APIs to perform the following operations:
- Upload and download files

- Create and edit different types of online documents

- Manage all your files, folders, and Wiki

### How to integrate

|  | Step | Description |
| --- | --- | --- |
| 1 | Create an app | - For creating a custom app, see Custom app development process. -   For creating a store app, see Develop and publish store apps. |
| 2 | Call APIs to operate Contacts. | Before calling APIs, you need to obtain the access token and enable corresponding permission scopes. For details, refer to How to call server-side APIs. You can also quickly debug these APIs using Postman: [![Run in Postman](https://run.pstmn.io/button.svg?lazyload=true&width=123&height=30)](https://god.gw.postman.com/run-collection/17195890-6fd42609-12b3-409b-8cbb-73656cf9d805?action=collection%2Ffork&collection-url=entityId%3D17195890-6fd42609-12b3-409b-8cbb-73656cf9d805%26entityType%3Dcollection%26workspaceId%3D55edcb9c-bbfe-45f5-a74a-efb882fe5384#?env%5Bfeishu-demo%5D=W3sia2V5IjoiYmFzZVVybCIsInZhbHVlIjoib3Blbi5mZWlzaHUuY24iLCJlbmFibGVkIjp0cnVlfSx7ImtleSI6InN0b3JlX2FwcF9pZCIsInZhbHVlIjoiY2xpX2EwYTUwODBjODRiODUwMTMiLCJlbmFibGVkIjp0cnVlfSx7ImtleSI6InN0b3JlX2FwcF9zZWNyZXQiLCJ2YWx1ZSI6InpFb0xLNHBjZE45VXBsNUpDOGNjcGZORlI3Q1FpbmNhIiwiZW5hYmxlZCI6dHJ1ZX0seyJrZXkiOiJhcHBfdGlja2V0IiwidmFsdWUiOiI3YTU1NjFlODdiMjkzYjFjOTEyZWM1NTQ2MDVjNDFlOWZhMjZkYzJmIiwiZW5hYmxlZCI6dHJ1ZX0seyJrZXkiOiJhcHBfYWNjZXNzX3Rva2VuIiwidmFsdWUiOiJhLWNlOTJjZTNhMmRjNmM2ZjQzYTVjNzM2YmRlMzAxM2FkYzdlZGM2MzQiLCJlbmFibGVkIjp0cnVlfSx7ImtleSI6InRlbmFudF9rZXkiLCJ2YWx1ZSI6IjczNjU4OGM5MjYwZjE3NWQiLCJlbmFibGVkIjp0cnVlfSx7ImtleSI6InRlbmFudF9hY2Nlc3NfdG9rZW4iLCJ2YWx1ZSI6InQtMmQ0OWY4ZjMyOTY2YTEzYmMzN2ZiMWJkZWFmZTBkNDdhMjAwZDZkZiIsImVuYWJsZWQiOnRydWV9LHsia2V5IjoiU1RBVEUiLCJ2YWx1ZSI6IjExIiwiZW5hYmxlZCI6dHJ1ZX0seyJrZXkiOiJSRURJUkVDVF9VUkkiLCJ2YWx1ZSI6Imh0dHBzJTNBJTJGJTJGd3d3LmJhaWR1LmNvbSUyRiIsImVuYWJsZWQiOnRydWV9LHsia2V5IjoidXNlcl9hY2Nlc3NfdG9rZW4iLCJ2YWx1ZSI6InUtMDJvYmhid01DSEl5c2ZhNzFWVGNUZCIsImVuYWJsZWQiOnRydWV9LHsia2V5IjoiYXBwX2lkIiwidmFsdWUiOiJjbGlfYTA3ZmM0ZDVhMmY5NTAwYyIsImVuYWJsZWQiOnRydWV9LHsia2V5IjoiYXBwX3NlY3JldCIsInZhbHVlIjoiVlpBcFd0ZXc2UUdHQm1SbmxJNTF2aEZtbUU0bkJScmwiLCJlbmFibGVkIjp0cnVlfV0=). For usage methods, see Postman template manual. |
| 3 | Listen events and obtain changes in Contacts. | You need to apply for relevant permission scope before listening on events. For details, see Event subscription overview. | ## Resources

Docs business domains are centered on resources. The following figure illustrates the relationship among resources:

![Frame 12.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/a8e93536264d0ec8d2c2aca57c3328bd_03F8pHDhZu.png?lazyload=true&width=1640&height=757)

The following table describes different resources.
| Folders | A container used to manage files and other folders. |
| --- | --- |
| Files | A general term for various types of files, referring to all files in My Space. |
| Documents | Online Lark docs. |
| Sheets | Lark sheets. |
| Comments | Comments on online Lark docs. |
