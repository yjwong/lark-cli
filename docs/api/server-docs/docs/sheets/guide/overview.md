---
title: "Overview"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/overview"
section: "Docs"
updated: "1648474091000"
---

# Spreadsheet overview
Supports operations such as automatically creating tables, reading or editing worksheets by calling the table API.
# Glossary
## Spreadsheet
A table is a container that carries data and also provides functions for data processing and presentation. A table may contain one or more worksheets, and when we process data, we all operate on a worksheet.
> **Note:** Each document has a spreadsheetToken as a unique identifier, and each worksheet has a sheetId as a unique identifier.
## spreadsheetToken and sheetId
spreadsheetToken is the unique identifier of a table, you can get the spreadsheetToken of a table in any of the following ways:
- Obtained through the URL of the sheet: https://sample.larksuite.com/sheets/shtcnmBA\*****yGehy8
- Obtain the spreadsheetToken of the table through the return value of Get the document list under the folder API

sheetId is the unique identifier of a worksheet, you can get the sheetId of a worksheet in any of the following ways:
- Obtain through the URL of the sheet: https://sample.larksuite.com/sheets/shtcnmBA\*****yGehy8?sheet=0b\**12
- Obtain the sheetId of the worksheet through the return value of Get Sheet Metadata API
> **Note:** Almost all table operation methods need to pass in spreadsheetToken to specify the table to be operated.
## Range
Range describes a range of the worksheet. In data reading and writing, it can help users filter the operation range of data.
The description method of range is \!:, there are 4 description methods, they are:
- \!:
For example: 0b\**12!A1:B5 means the area of A1:B5 in the worksheet of 0b\**12, as shown in the following figure:
![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/76498891d78bff326b0bcffe43427fa9_fUFkCG77Vw.png?lazyload=true&width=722&height=484)
- \!:, such as: 0b\**12!A:B
- \!:, eg: 0b\**12!A1:B
- \, the area is left blank, such as: 0b\**12, which represents the data within the maximum row and column range that is not empty in this sheet
When using the interface for reading table data, such as "reading a single range" or "reading multiple ranges", the above description methods represent the area of A1:B5 in this worksheet, as shown in the following figure:
![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/dca12da354cb99d01c2893283893ae2a_Y9COFAZsea.png?lazyload=true&width=856&height=576)
# Method collection
## Sheet
Contains interfaces related to table creation, acquisition and update.
### Method list
> "Store" stands for Store; "Custom" stands for Custom apps
| `create sheet `POST` /open-apis/sheets/v3/spreadsheets ` | `drive:drive` View, comment, edit and manage spreadsheets | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| --- | --- | --- | --- | --- |
| `Get table metadata `GET` /open-apis/sheets/v2/spreadsheets/:spreadsheetToken/metainfo ` | `drive:drive` View, comment and download all files in cloud space View, comment, edit and manage spreadsheets `sheets:spreadsheet:readonly` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Update table properties `PUT` /open-apis/sheets/v2/spreadsheets/:spreadsheetToken/properties ` | `drive:drive` View, comment, edit and manage spreadsheets | `tenant_access_token` `user_access_token` | **✓** | **✓** | ## Worksheet
Contains interfaces related to worksheet operations and updating properties.
### Method list
> "Store" stands for Store apps; "Custom" stands for Custom apps
| `Update sheet properties `POST` /open-apis/sheets/v2/spreadsheets/:spreadsheetToken/sheets_batch_update ` | `drive:drive` View, comment, edit and manage spreadsheets | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| --- | --- | --- | --- | --- |
| `operation sheet `POST` /open-apis/sheets/v2/spreadsheets/:spreadsheetToken/sheets_batch_update ` | `drive:drive` View, comment, edit and manage spreadsheets | `tenant_access_token` `user_access_token` | **✓** | **✓** | ## Worksheet - Conditional Formatting
Contains interfaces for creating, deleting, getting, and updating conditional formatting in the worksheet.
### Method list
> "Store" stands for Store apps; "Custom" stands for Custom apps
| `Create Conditional Format `POST` /open-apis/sheets/v2/spreadsheets/:spreadsheetToken/condition_formats/batch_create ` | `drive:drive` View, comment, edit and manage spreadsheets | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| --- | --- | --- | --- | --- |
| `Get conditional format `GET` /open-apis/sheets/v2/spreadsheets/:spreadsheetToken/condition_formats ` | `drive:drive` View, comment and download all files in cloud space View, comment, edit and manage spreadsheets `sheets:spreadsheet:readonly` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `update conditional format `POST` /open-apis/sheets/v2/spreadsheets/:spreadsheetToken/condition_formats/batch_update ` | `drive:drive` View, comment, edit and manage spreadsheets | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `delete conditional format `DELETE` /open-apis/sheets/v2/spreadsheets/:spreadsheetToken/condition_formats/batch_delete ` | `drive:drive` View, comment, edit and manage spreadsheets | `tenant_access_token` `user_access_token` | **✓** | **✓** | ## Worksheet - Filter
Contains worksheet filter create, delete, get and update related interfaces.
### Method list
> "Store" stands for Store apps; "Custom" stands for Custom apps
| `get filter `GET` /open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter ` | `sheets:spreadsheet` `sheets:spreadsheet:readonly` `drive:drive` `drive:drive:readonly` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| --- | --- | --- | --- | --- |
| `create filter `POST` /open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter ` | `sheets:spreadsheet` `drive:drive` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `update filter `PUT` /open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter ` | `drive:drive` View, comment, edit and manage spreadsheets | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `delete filter `DELETE` /open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter ` | `sheets:spreadsheet` `drive:drive` | `tenant_access_token` `user_access_token` | **✓** | **✓** | ## Worksheet - Filter View
Contains worksheet filter view create, delete, get and update related interfaces.
### Method list
> "Store" stands for Store apps; "Custom" stands for Custom apps
| `delete filter view `DELETE` /open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/:filter_view_id ` | `drive:drive` `sheets:spreadsheet` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| --- | --- | --- | --- | --- |
| `Update Filter View `PATCH` /open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/:filter_view_id ` | `sheets:spreadsheet` `drive:drive` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `query filter view `GET` /open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/query ` | `sheets:spreadsheet:readonly` `drive:drive` `drive:drive:readonly` `sheets:spreadsheet` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Get Filter View `GET` /open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/:filter_view_id ` | `sheets:spreadsheet` `sheets:spreadsheet:readonly` `drive:drive` `drive:drive:readonly` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `create filter view `POST` /open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views ` | `sheets:spreadsheet` `drive:drive` | `tenant_access_token` `user_access_token` | **✓** | **✓** | ## Filter View - Filter Criteria
Contains the filter criteria create, delete, get, update and query related interfaces of the filter view.
### Method list
> "Store" stands for Store apps; "Custom" stands for Custom apps
| `delete filter condition `DELETE` /open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/:filter_view_id/conditions/:condition_id ` | `sheets:spreadsheet` `drive:drive` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| --- | --- | --- | --- | --- |
| `Update filter condition `PUT` /open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/:filter_view_id/conditions/:condition_id ` | `sheets:spreadsheet` `drive:drive` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `query filter condition `GET` /open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/:filter_view_id/conditions/query ` | `sheets:spreadsheet` `sheets:spreadsheet:readonly` `drive:drive` `drive:drive:readonly` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Get filter condition `GET` /open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/:filter_view_id/conditions/:condition_id ` | `drive:drive:readonly` `sheets:spreadsheet` `sheets:spreadsheet:readonly` `drive:drive` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `create filter condition `POST` /open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/:filter_view_id/conditions ` | `sheets:spreadsheet` `drive:drive` | `tenant_access_token` `user_access_token` | **✓** | **✓** | ## Worksheet - Row and Column
Contains interfaces related to adding, inserting, deleting, moving and updating worksheet rows and columns.
### Method list
> "Store" stands for Store apps; "Custom" stands for Custom apps
| `delete row and column `DELETE` /open-apis/sheets/v2/spreadsheets/:spreadsheetToken/dimension_range ` | `drive:drive` View, comment, edit and manage spreadsheets | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| --- | --- | --- | --- | --- |
| `update row `PUT` /open-apis/sheets/v2/spreadsheets/:spreadsheetToken/dimension_range ` | `drive:drive` View, comment, edit and manage spreadsheets | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Add row and column `POST` /open-apis/sheets/v2/spreadsheets/:spreadsheetToken/dimension_range ` | `drive:drive` View, comment, edit and manage spreadsheets | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `insert row and column `POST` /open-apis/sheets/v2/spreadsheets/:spreadsheetToken/insert_dimension_range ` | `drive:drive` View, comment, edit and manage spreadsheets | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Move row and column `POST` /open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/move_dimension ` | `drive:drive` `sheets:spreadsheet` | `tenant_access_token` `user_access_token` | **✓** | **✓** | ## Row and Column - Range of Protection
The protection scope including the row and column adds, deletes, obtains and modifies related interfaces.
### Method list
> "Store" stands for Store apps; "Custom" stands for Custom apps
| `Add protection scope `POST` /open-apis/sheets/v2/spreadsheets/:spreadsheetToken/protected_dimension ` | `drive:drive` View, comment, edit and manage spreadsheets | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| --- | --- | --- | --- | --- |
| `Get protection scope `GET` /open-apis/sheets/v2/spreadsheets/:spreadsheetToken/protected_range_batch_get ` | `drive:drive` View, comment, edit and manage spreadsheets | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Modify protection scope `POST` /open-apis/sheets/v2/spreadsheets/:spreadsheetToken/protected_range_batch_update ` | `drive:drive` View, comment, edit and manage spreadsheets | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `delete protection scope `DELETE` /open-apis/sheets/v2/spreadsheets/:spreadsheetToken/protected_range_batch_del ` | `drive:drive` View, comment, edit and manage spreadsheets | `tenant_access_token` `user_access_token` | **✓** | **✓** | ## Worksheet - Data
Contains interfaces related to writing and reading cell data in worksheets, inserting pictures, setting styles, merging and splitting cells, and finding and replacing content.
### Method list
> "Store" stands for Store apps; "Custom" stands for Custom apps
| `insert data `POST` /open-apis/sheets/v2/spreadsheets/:spreadsheetToken/values_prepend ` | `drive:drive` View, comment, edit and manage spreadsheets | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| --- | --- | --- | --- | --- |
| `append data `POST` /open-apis/sheets/v2/spreadsheets/:spreadsheetToken/values_append ` | `drive:drive` View, comment, edit and manage spreadsheets | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `read single range `GET` /open-apis/sheets/v2/spreadsheets/:spreadsheetToken/values/:range ` | `drive:drive` View, comment and download all files in cloud space View, comment, edit and manage spreadsheets `sheets:spreadsheet:readonly` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `read multiple ranges `GET` /open-apis/sheets/v2/spreadsheets/:spreadsheetToken/values_batch_get ` | `drive:drive` View, comment and download all files in cloud space View, comment, edit and manage spreadsheets `sheets:spreadsheet:readonly` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Write data to a single range `PUT` /open-apis/sheets/v2/spreadsheets/:spreadsheetToken/values ` | `drive:drive` View, comment, edit and manage spreadsheets | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Write data to multiple ranges `POST` /open-apis/sheets/v2/spreadsheets/:spreadsheetToken/values_batch_update ` | `drive:drive` View, comment, edit and manage spreadsheets | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `set cell style `PUT` /open-apis/sheets/v2/spreadsheets/:spreadsheetToken/style ` | `drive:drive` View, comment, edit and manage spreadsheets | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Batch cell styles `PUT` /open-apis/sheets/v2/spreadsheets/:spreadsheetToken/styles_batch_update ` | `drive:drive` View, comment, edit and manage spreadsheets | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `merge cells `POST` /open-apis/sheets/v2/spreadsheets/:spreadsheetToken/merge_cells ` | `drive:drive` View, comment, edit and manage spreadsheets | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `split cell `POST` /open-apis/sheets/v2/spreadsheets/:spreadsheetToken/unmerge_cells ` | `drive:drive` View, comment, edit and manage spreadsheets | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `write image `POST` /open-apis/sheets/v2/spreadsheets/:spreadsheetToken/values_image ` | `drive:drive` View, comment, edit and manage spreadsheets | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Find Cell `POST` /open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/find ` | `drive:drive` View, comment, edit and manage spreadsheets View, comment and download all files in cloud space `sheets:spreadsheet:readonly` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `replace cell `POST` /open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/replace ` | `drive:drive` View, comment, edit and manage spreadsheets | `tenant_access_token` `user_access_token` | **✓** | **✓** | ## Worksheet - Floating Image
Contains interfaces related to creating, deleting, querying, getting and updating worksheet floating pictures.
### Method list
> "Store" stands for Store apps; "Custom" stands for Custom apps
| `create floating image `POST` /open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/float_images ` | `drive:drive` View, comment, edit and manage spreadsheets | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| --- | --- | --- | --- | --- |
| `Get floating image `GET` /open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/float_images/:float_image_id ` | `drive:drive` View, comment, edit and manage spreadsheets View, comment and download all files in cloud space `sheets:spreadsheet:readonly` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `query floating image `GET` /open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/float_images/query ` | `drive:drive` View, comment, edit and manage spreadsheets View, comment and download all files in cloud space `sheets:spreadsheet:readonly` | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Update floating image `PATCH` /open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/float_images/:float_image_id ` | `drive:drive` View, comment, edit and manage spreadsheets | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `delete floating image `DELETE` /open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/float_images/:float_image_id ` | `drive:drive` View, comment, edit and manage spreadsheets | `tenant_access_token` `user_access_token` | **✓** | **✓** | ## Worksheet - Data Check
Contains worksheet data validation settings, delete, query and update related interfaces.
### Method list
> "Store" stands for Store apps; "Custom" stands for Custom apps
| `set dropdown `POST` /open-apis/sheets/v2/spreadsheets/:spreadsheetToken/dataValidation ` | `drive:drive` View, comment, edit and manage spreadsheets | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| --- | --- | --- | --- | --- |
| `delete drop-down list setting `DELETE` /open-apis/sheets/v2/spreadsheets/:spreadsheetToken/dataValidation ` | `drive:drive` View, comment, edit and manage spreadsheets | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Update dropdown list settings `PUT` /open-apis/sheets/v2/spreadsheets/:spreadsheetToken/dataValidation/:sheetId/:dataValidationId ` | `drive:drive` View, comment, edit and manage spreadsheets | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Query drop-down list settings `GET` /open-apis/sheets/v2/spreadsheets/:spreadsheetToken/dataValidation ` | `drive:drive` View, comment, edit and manage spreadsheets | `tenant_access_token` `user_access_token` | **✓** | **✓** |
