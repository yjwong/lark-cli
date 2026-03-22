---
title: "Base overview"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/bitable-overview"
section: "Docs"
updated: "1753668066000"
---

# Base overview

Lark Base is a powerful data management platform, which can help you build applications and re-organize the online data collaboration. With Lark Base, you can transform complex data into actionable insights, build up customized applications and keep everything on tracked. Lark Base APIs enable you to easily integrate Lark Base with other systems and platforms, ensuring real-time updates for all your data.
## Typical cases

The open platform provides customer practice cases that integrate Base capabilities:

- [Sesame Open, building an access control management system in Lark in five days](https://open.larksuite.com/solutions/detail/wiseasy)
- [Lark Integration Platform, upgrading project management and sales experience in the enterprise service industry](https://open.larksuite.com/solutions/detail/cloudwise)
## Development tutorials

Explore the following Base-related tutorials to learn how to use the Base API to facilitate efficient enterprise collaboration.

- Quick access to Base
- Agile project management with Base
## Authentication instructions
Before using the `tenant_access_token` to access Base resources, ensure your application is either an owner or collaborator of the Base; otherwise, the call will fail. You can add the app as a collaborator by adding it to the document application, or create a Base with the app's identity and use the `tenant_access_token` to call the API.

## Usage limitations

The Base API has the following restrictions:
- For batch operations, a maximum of 1,000 records can be processed at a time, with the response status being entirely successful or failed.
- To ensure stability, it is recommended to perform only **one** API write operation at a time on a single Base.
Resource quantity limits within a single Base are listed below:

| **Resource**             | **Maximum Limit**             |
|--------------------------|-------------------------------|
| Records                  | Varies by tenant; no extra restrictions from the platform. Refer to the Base UI for details. ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/c05d45fab4f5c8862624821d9370d2dc_1iQ0TinUhV.png?height=582&lazyload=true&width=1128) |
| Fields                   | 300, with up to 100 formula fields |
| Views                    | 200                           |
| Tables                   | 100                           |
| Custom Roles             | 30                            |
| Advanced Permission Collaborators | 200                   | ## Resource introduction

Currently, Base provides interfaces for eight resource types: Base App, Tables, Views, Records, Fields, Dashboards, Custom Roles, and Advanced Permission Collaborators. This section introduces the meaning of these resources. 

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/80c83de730f0d48a6e2347009c6b2633_N8DO0TPC1z.png?height=719&lazyload=true&maxWidth=700&width=1280)

### Base App

A Base can be considered an application (not to be confused with developer-created apps), identified by a unique `app_token`. A Base can exist as a standalone app or be integrated as a block within documents and spreadsheets.

#### Forms of Base

| **Base form** | **Resource definition** | **Meaning** |
| ---------- | -------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| Base in folder | Base app | Base stored in a folder in the cloud space (cloud drive). The URL starts with **larksuite.com/base**. |
| Base under wiki | Base app and wiki node | Base placed in a wiki. The URL starts with **larksuite.com/wiki**. |
| Base embedded in doc | Base docx block | Base embedded in a "Doc". The URL starts with **larksuite.com/docx**. |
| Base embedded in sheet | Base sheet block | Base embedded in a sheet. The URL starts with **larksuite.com/sheets**. 
#### Obtaining Base `app_token`

The method for obtaining the `app_token` varies depending on the form of the Base.

##### **Base in a Folder**

The `app_token` for this form of Base is highlighted in the URL.

![Base Folder URL](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/6916f8cfac4045ba6585b90e3afdfb0a_sTn7sVvhOB.png?height=766&lazyload=true&maxWidth=700&width=3004)

##### **Base in Knowledge Base**

Call the Get Node Information API to obtain the `app_token`. For example, when the `obj_type` is `bitable`, the `obj_token` value is the `app_token`.
```json
{
    "node":{
        "space_id":"6946843325487912356",
        "node_token":"wikcnKQ1k3p******8Vabcef",
        "obj_token":"AW3Qbtr2cakCnesXzXVbbsrIcVT",
        "obj_type":"bitable"
    }
}
```

##### **Base Embedded in Document**
To obtain the app_token for a Base embedded in a document, call the Get Document Blocks API.

```json
{
  "bitable": {
    "token": "AW3Qbtr2cakCnesXzXVbbsrIcVT_tblkIYhz52o6G5nx"
  }
}
```
##### **Base Embedded in Spreadsheet**
For Bases embedded in spreadsheets, call the Get Spreadsheet Metadata API.

```json
{
  "blockInfo": {
    "blockToken": "AW3Qbtr2cakCnesXzXVbbsrIcVT_tblkIYhz52o6G5nx",
    "blockType": "BITABLE_BLOCK"
  }
}
```

### Table

The data container in a Base. Each Base has at least one data table, but it can have multiple data tables. Each data table has a unique identifier called `table_id`. The `table_id` is unique within a single Base app but not necessarily unique globally.
You can obtain the `table_id` from the Base URL; the highlighted part in the image below shows the unique identifier for the current data table. You can also use the List Tables API to get the `table_id`.

![](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/18741fe2a0d3cafafaf9949b263bb57d_yD1wkOrSju.png?height=746&lazyload=true&maxWidth=700&width=2976)

### View

Refers to the summary and presentation format of the data. There are various types of views, including Table View, Kanban View, Gallery View, Gantt View, and Form View. Each data table has at least one view and may have multiple views. Each view has a unique identifier called `view_id`, which is unique within a single Base but not necessarily unique globally.
You can obtain the `view_id` from the Base URL; the highlighted part in the image below shows the unique identifier for the current view. You can also use the List Views API to get the `view_id`. It is currently not possible to retrieve the `view_id` for embedded Bases.

![](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/140668632c97e0095832219001d17c54_c76RMwZAgW.png?height=748&lazyload=true&maxWidth=700&width=2998)

#### **Form View**

Form View is a type of view in a Base, similar to a questionnaire, which can be used to collect information and data. Each form has a unique identifier called `form_id`, which corresponds to the current view's `view_id`. The method for obtaining `form_id` is the same as for `view_id`.

![](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/b8d9bd2d7e1ca7d0cc955ef7f1f4fe16_zDd5TqOh3Q.png?height=942&lazyload=true&maxWidth=600&width=1327)

### Record

Each row of data in a data table is a record. Every record has a unique identifier called `record_id`, which is unique within a single Base but not necessarily unique globally. The `record_id` can be obtained using the Search Records API.

![](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/abc84b39be159ccdcafa707ee141144d_3fcTz7mMP5.png?height=503&lazyload=true&maxWidth=700&width=1536)

### Field

Fields in a Base are equivalent to columns. Bases offer a variety of field types. Each field has a unique identifier called `field_id`, which is unique within a single Base but not necessarily unique globally. The `field_id` can be obtained using the List Fields API. For more details on fields, refer to the Field Editing Guide.

![](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/6fb2359552ed15433289fbd0d9fc53c1_mGnTo91OJa.png?height=656&lazyload=true&maxWidth=700&width=1170)

### Dashboard

Dashboards are similar to data boards and can be used to analyze data in a Base from different dimensions.

![](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/fae3c8a886a0075fdeeefe5b74f274e5_WepjCfS7Hx.png?height=1490&lazyload=true&maxWidth=600&width=2794)

The unique identifier for a dashboard is called `block_id`, which starts with `blk`. You can get the `block_id` from the Base URL; the highlighted part in the image below shows the unique identifier for the current dashboard. You can also use the List Dashboards API to get it.

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/a966d15323ee73c66b1e9a31d34ae6c7_x3ctncH2nO.png?height=575&lazyload=true&maxWidth=600&width=1397)

### Advanced Permissions

Advanced permissions allow users to set which users can view or edit specific rows in a data table or to specify which columns can be edited by a user. The advanced permission interfaces are divided into **Custom Roles** and **Collaborators**. The **owner** of the Base or users with **management permissions** can set advanced permissions and manage collaborators via the interface. For more details, see the Advanced Permission Overview.

#### **Custom Role**

A role added in advanced permissions with set permissions is a custom role. Each custom role has a unique identifier called `role_id`. The `role_id` can be obtained using the List Custom Roles API.

#### **member**

A member of a custom role in the advanced permissions settings is a collaborator. Each collaborator has a unique identifier called `member_id`, which can be obtained using the List Collaborators API.
### Workflows

Automated workflows are rules set by the user for automatic operation in the Base. After setting "trigger conditions" and "execute actions," the Base will automatically execute the next step based on data changes. You can obtain the ID of automated workflows, i.e., `workflow_id`, through the List automated workflows interface.

### Method list

#### Base
> "Store" represents Store apps; "Custom" represents Custom apps

| `Get App Info `GET` /open-apis/bitable/v1/apps/:app_token ` | `bitable:app:readonly` `bitable:app` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| --- | --- | --- | --- | --- |
| `Update App Info `GET` /open-apis/bitable/v1/apps/:app_token ` | `bitable:app` | `tenant_access_token` `user_access_token` | **✓** | **✓** | #### Table

| `List Tables `GET` /open-apis/bitable/v1/apps/:app_token/tables ` | `bitable:app:readonly` `bitable:app` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| --- | --- | --- | --- | --- |
| `Create Table` `POST` /open-apis/bitable/v1/apps/:app_token/tables | `bitable:app` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Create Tables` `POST` /open-apis/bitable/v1/apps/:app_token/tables/batch_create | `bitable:app` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Delete Table` `DELETE` /open-apis/bitable/v1/apps/:app_token/tables/:table_id | `bitable:app` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Delete Tables` `POST` /open-apis/bitable/v1/apps/:app_token/tables/batch_delete | `bitable:app` | `tenant_access_token` `user_access_token` | **✓** | **✓** | #### View

| `List Views `GET` /open-apis/bitable/v1/apps/:app_token/tables/:table_id/views ` | `bitable:app:readonly` `bitable:app` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| --- | --- | --- | --- | --- |
| `Create View` `POST` /open-apis/bitable/v1/apps/:app_token/tables/:table_id/views | `bitable:app` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Delete View` `DELETE` /open-apis/bitable/v1/apps/:app_token/tables/:table_id/views/:view_id | `bitable:app` | `tenant_access_token` `user_access_token` | **✓** | **✓** | #### Record

| `List records `GET` /open-apis/bitable/v1/apps/:app_token/tables/:table_id/records ` | `bitable:app:readonly` `bitable:app` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| --- | --- | --- | --- | --- |
| `Get Record` `GET` /open-apis/bitable/v1/apps/:app_token/tables/:table_id/records/:record_id | `bitable:app:readonly` `bitable:app` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Create record` `POST` /open-apis/bitable/v1/apps/:app_token/tables/:table_id/records | `bitable:app` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Create Records` `POST` /open-apis/bitable/v1/apps/:app_token/tables/:table_id/records/batch_create | `bitable:app` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Update Record` `PUT` /open-apis/bitable/v1/apps/:app_token/tables/:table_id/records/:record_id | `bitable:app` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Update Records` `POST` /open-apis/bitable/v1/apps/:app_token/tables/:table_id/records/batch_update | `bitable:app` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Delete Record` `DELETE` /open-apis/bitable/v1/apps/:app_token/tables/:table_id/records/:record_id | `bitable:app` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Delete Records` `POST` /open-apis/bitable/v1/apps/:app_token/tables/:table_id/records/batch_delete | `bitable:app` | `tenant_access_token` `user_access_token` | **✓** | **✓** | #### Field

| `List Fields `GET` /open-apis/bitable/v1/apps/:app_token/tables/:table_id/fields ` | `bitable:app:readonly` `bitable:app` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| --- | --- | --- | --- | --- |
| `Create Field` `POST` /open-apis/bitable/v1/apps/:app_token/tables/:table_id/fields | `bitable:app` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Update Field` `PUT` /open-apis/bitable/v1/apps/:app_token/tables/:table_id/fields/:field_id | `bitable:app` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Delete Field` `DELETE` /open-apis/bitable/v1/apps/:app_token/tables/:table_id/fields/:field_id | `bitable:app` | `tenant_access_token` `user_access_token` | **✓** | **✓** | #### Role

| `List Roles `GET` /open-apis/bitable/v1/apps/:app_token/roles ` | `bitable:app` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| --- | --- | --- | --- | --- |
| `Add Role` `POST` /open-apis/bitable/v1/apps/:app_token/roles/:role_id | `bitable:app` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Update Role` `PUT` /open-apis/bitable/v1/apps/:app_token/roles/:role_id | `bitable:app` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Delete Role` `DELETE` /open-apis/bitable/v1/apps/:app_token/roles/:role_id | `bitable:app` | `tenant_access_token` `user_access_token` | **✓** | **✓** | #### Member

| `List Members `GET` /open-apis/bitable/v1/apps/:app_token/roles/:role_id/members ` | `bitable:app` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| --- | --- | --- | --- | --- |
| `Add Member` `POST` /open-apis/bitable/v1/apps/:app_token/roles/:role_id/members | `bitable:app` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Delete Member` `DELETE` /open-apis/bitable/v1/apps/:app_token/roles/:role_id/members/:member_id | `bitable:app` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
