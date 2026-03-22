---
title: "Overview"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file-version/overview"
service: "drive-v1"
resource: "file-version"
section: "Docs"
updated: "1711433292000"
---

# Overview
## Function introduction

Support users to manually generate document versions, and efficiently manage these versions, including:

**Create Version:** Add a new document version.

**Manage Versions:** Can delete document versions.

**Version reading:** Get version document content.

## Resource definition

| field name         | type           | describe        | Remark         |
| --------- | --------------- | -------   | ----------- |
|`name` | `string` | Version document title | -- |
|`version` | `string` | Version document version number | -- |
|`parent_token` | `string` | source document token | -- |
|`owner_id` | `string` | Version document owner id | -- |
|`creator_id` | `string` | Version document creator id | -- |
|`create_time` | `int` | Version document creation time | -- |
|`update_time` | `int` | Version document update time | -- |
|`status` | `string` | Version document status |Enumeration value:- `StatusExist`: normal status;- `StatusDeleted`: delete status;- `StatusTrash`: recycle bin status |
|`obj_type` | `string` | Version document type |Enumeration value:- Currently supported: docx, sheet|
|`parent_type` | `string` | Version document parent_type |Enumeration value:- Currently supported: docx, sheet |
