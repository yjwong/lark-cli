---
title: "Overview"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/import_task/import-user-guide"
service: "drive-v1"
resource: "import_task"
section: "Docs"
updated: "1684896004000"
---

# Import Usage Guide
This guide describes how to import files in specified formats as Upgraded Docs, Sheets, and Bitable, etc., and mount them in specific folder in Space.

## Step 1: Upload Local Files
> **Note:** If the files to be imported have already been uploaded to the space, you can skip this step.
The following upload methods are currently supported:
**Way 1:** Uploading material through Upload material
. The main parameters involved in the request are:
- `parent_type`: a fixed value `ccm_import_open needs` to be passed in.
- `parent_node`: no need to fill in.
- `extra`: refer to Materials Overview - Types and the Extra parameter of upload points to fill in.

As an example of importing a local file named "demo.txt" as an Upgraded Docs:
```JavaScript 
curl --location --request POST 'https://{domain}/open-apis/drive/v1/medias/upload_all' \ // domain needs to be replaced with the real domain name
--header 'Content-Type: multipart/form-data' \
--header 'Authorization: Bearer u-{xxxxx}' \ // needs to be replaced with real Authorization
--form 'file_name="demo.txt"' \ // local file name
--form 'parent_type="ccm_import_open"' \ // use fixed value: ccm_import_open
--form 'size="5"' \ // local file size (in bytes)
--form 'extra={"obj_type":"docx","file_extension":"txt"}' \ // import local txt format files as Upgraded Docs
--form 'file=@"demo.txt"' // local file path
```
**Way 2:** Upload files through Upload file.
> **Note:** For the above two ways, developers can choose on demand.
> - Way 1: the uploaded material is a temporary file, which will be deleted by the system actively after the import is completed.
> - Way 2: The uploaded files are uploaded to the space, the files belong to the uploader and occupy the uploader's space usage, the import system will not do active deletion.
## Step 2: Create import task
Call the Create an import task, the successful request will return the `ticket` (import task ID), the main parameters involved in the request are:
- `file_token`: `Token` of the material or file to be imported.
- `point`: space mount point, where `mount_key` is the target mount directory. If it is empty, it will be mounted to the root directory of the importer's space.
- `type`: the type of cloud document generated after importing. Currently, the following types are supported.
  + `docx`: Upgraded Docs
  + `sheet`: Sheets
  + `bitable`: Sheets
  + `doc`: Docs, it is no longer recommended, please use Upgraded Docs
- `file_extension`: The type of file extension that needs to be imported, i.e., the type of file corresponding to the `file_token`.
> **Note:** `type` and `file_extension` need to be seen together:
> - When `type` takes the value of `docx`, `file_extension` supports the following values:
>   + `docx` (Microsoft Word Document)
>   + `doc` (Microsoft Word 97-2004 Document)
>   + `txt` (text file)
>   + `md` (Markdown)
>   + `mark` (Markdown)
>   + `markdown` (Markdown)
>   + `html` (HTML)
> - When `type` takes the value of `sheet`, `file_extension` supports the following values:
>   + `xlsx` (Microsoft Excel Workbook)
>   + `xls` (Microsoft Excel 97- Excel 2003 Workbook)
>   + `csv` (Comma Separated Values)
> - When `type` takes the value of `bitable`, `file_extension` supports the following values:
>   + `xlsx` (Microsoft Excel Workbook)
>   + `csv` (Comma Separated Values)
> - When `type` takes the value `doc`, `file_extension` supports the following values:
>   + `docx` (Microsoft Word Document)
>   + `doc` (Microsoft Word 97-2004 Document)
>   + `txt` (text file)
>   + `md` (Markdown)
>   + `mark` (Markdown)
>   + `markdown` (Markdown)
Take a plain text file with a `file_extension` of ".txt" and a `Token` of "Nlg1b8TZQoVDy8xcSzJcPPjinzf" and import it as Upgraded Docs as an example.
```javascript 
curl -i -X POST 'https://{domain}/open-apis/drive/v1/import_tasks' \ // domain needs to be replaced with the real domain name
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer u-{xxxxx}' \ // needs to be replaced with real Authorization
-d '{
	"file_extension": "txt", // source file is in "txt" format
	"file_name": "demo",
	"file_token": "Nlg1b8TZQoVDy8xcSzJcPPjinzf",
	"point": {
		"mount_key": "VsGvfNmCAl3t62dIppwcGohNnzc",
		"mount_type": 1 // currently only mounts to space are supported
	},
	"type": "docx" // import as Upgraded Docs
}'
``` 
## Step 3: Query the import task results
Call Query the import task result and poll the import result on demand according to the `ticket` (import job ID) returned in step 2.

- `job_status`: the status of the import job. For example, `0` means the import job was executed successfully.
- `extra`: extra information. Sometimes the import is successful but the content of the source file exceeds the system limit and is truncated by the system, so this parameter is returned for additional prompting:
  + **2000**: the number of columns in the Sheet exceeds the maximum limit, the excess will be truncated and discarded.
  + **2001**: the number of Sheet cells exceeds the maximum limit, the excess will be truncated and discarded.
  + **2002**: the number of Sheet pictures exceeds 4000, and the exceeded pictures are discarded.
  + **2003**: Insufficient storage space in space, please contact the enterprise administrator.
  + **2004**: some pictures of the Sheet failed to be uploaded.
  + **2005**: the Sheet cell character length exceeds the maximum limit, the exceeded part will be truncated and discarded.
  + **3000**: Bitable pictures exceeding the data row area will be discarded.
  + **3001**: Bitable pictures exceeding the data column area will be discarded.
  + **3003**: Bitable columns exceed the maximum limit, the excess will be truncated and discarded.
  + **3004**: the number of cells in a Bitable exceeds the maximum limit, and the excess will be truncated and discarded.
  + **3005**: the number of pictures of the Bitable exceeds 4000, the exceeding pictures are discarded.
  + **3006**: Some pictures of the Bitable failed to be uploaded.
