---
title: "Upload files"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUDOyUjL1gjM14SN4ITN"
method: "POST"
api_path: "https://www.larksuite.com/approval/openapi/v2/file/upload"
section: "Approval"
scopes:
  - "approval:approval:readonly"
updated: "1675167411000"
---

# Upload files

When images or attachment widgets are used in an approval form, developers must use this API to upload necessary files to the approval system before an approval instance is created.Attachments must be 50 MB or less. Images must be 10 MB or less.

##  Request
| HTTP URL | https://www.larksuite.com/approval/openapi/v2/file/upload |
| --- | --- |
| HTTP Method | POST |
| Supported app types | custom,isv |
| Required scopes | `approval:approval:readonly` | ###  Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes  `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about how to obtain and use access_token |  |
| Content-Type | string | No | user-specified content-type not supported,filling in this parameter may result in an error | ### Request body

| Parameter         | Type           | Required        | Description        |
| --------- | --------------- | -------   | --------- |
|name | string | Yes |  File name (must contain the file extension, such as "document.doc") |
|type |string | Yes | File type (image or attachment) |
|content | file | Yes |File | ::: note
**Note**: Only 1 file can be uploaded each time. To upload multiple files, perform multiple uploads. The file type (image or attachment) is determined by the widget type in the approval definition form.

### Request body example

```json
{
	"name":"123.doc",
	"type":"attachment",
	"content":123.doc
}
````

##  Response

###  Response body

| Parameter         |Type         |Required  | Description        |
| --------- | ----------|----- | --------- |
|code |int |Yes |Error code, a value other than 0 indicates failure. |
|msg | string |Yes| Return code description|
|data | map |Yes| Returned business information |
|  ∟code|string|Yes| File ID code (used to create approval instances)|
|  ∟url|string|Yes| File URL| ::: note
**Note**: The returned URL is valid for 12 hours. After an approval is initiated, a new URL is returned each time details are obtained.

###  Response body example

```json
{
    "code":0,
    "msg":"success",
    "data": {
        "code": "D93653C3-2609-4EE0-8041-61DC1D84F0B5",
        "url": "https://p3-approval-sign.byteimg.com/lark-approval-attachment/image/20210819/a8c1a1f1-47ae-4147-9deb-a8bf2cd833b1.jpg~tplv-ottatrvjsm-image.image?x-expires=1634941752&x-signature=oaZ6Tfv50ryUesNwKTUTnBlJivY%3D#.jpg"
    }
}
```
