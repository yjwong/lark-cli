---
title: "Overview"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/export_task/export-user-guide"
service: "drive-v1"
resource: "export_task"
section: "Docs"
updated: "1684896008000"
---

# Exporting User Guide

This guide describes how to export Upgraded Docs, Sheets, Bitable and Docs to local files, as well as the considerations related to exporting.

## Export Steps

### Step 1: Create an export task

The Create an export task is called exporting Docs to local files in the specified format, which currently supports exporting Upgraded Docs, Sheets, Bitable and Docs. The API is asynchronous and returns immediately when the task is created and does not block until the task is successfully executed.
- `type`: the type of Docs to be exported. Currently, the following types of Docs are supported:
    - `docx`: Upgraded Docs
    - `sheet`: Sheets
    - `bitable`: Bitable
    - `doc`: Docs
- `file_extension`: the file extension of the exported product, such as `pdf`.
> **Note:** `type` and `file_extension` need to be seen together, i.e. the file extension of the exported product must match the Docs type:
> - When `type` takes the value of `docx`, `file_extension` supports the following values:
>     - `docx` (Microsoft Word Document)
>     - `pdf` (Portable Document Format)
> - When `type` takes the value of `sheet`, `file_extension` supports the following values:
>     - `xlsx` (Microsoft Excel Workbook)
>     - `csv` (Comma Separated Values)
> - When `type` takes the value of `bitable`, `file_extension` supports the following values:
>     - `xlsx` (Microsoft Excel Workbook)
>     - `csv` (Comma Separated Values)
> - When `type` takes the value of `doc`, `file_extension` supports the following values:
>     - `docx` (Microsoft Word Document)
>     - `pdf` (Portable Document Format)
> **Warning:** If you want to export `wiki` (Wiki), you need to get the corresponding `obj_token` (Docs Token) and `obj_type` (Docs Type) of the node by Get Wiki node information, and then create the specific export task accordingly.
### Step 2: Query the export task results

Call the Query export task results to get the export result by the `token` and `ticket`(export task ID) of the Docs to be exported.

### Step 3: Download the export file

Call the Download export file and download the corresponding export file according to the export product `token` obtained from the export task result.

## Caution

- The exported file will be deleted 10 minutes after the task ends, and it will not be downloaded after that.
- When exporting as `docx`, it will be limited by the size of the Docs resources (such as image). If the total resources in the Docs exceeds 1GB, the export will fail.
- When exporting as `pdf`, it will be limited by the size of the Docs resources (such as image). If the total resources in the Docs exceeds 128MB, the export will fail.
- When watermarking is enabled within the enterprise.
    - Exporting as an `tenant_access_token` will use the Application ID as the watermark, such as `cli_a2c2xxxxxxxxd01b`.
    - Exporting as a `user_access_token` will use the User Identity as the watermark.
