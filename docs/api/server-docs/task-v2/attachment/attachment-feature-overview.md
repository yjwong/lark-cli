---
title: "Attachment Feature Overview"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/attachment/attachment-feature-overview"
service: "task-v2"
resource: "attachment"
section: "Tasks v2"
updated: "1699521697000"
---

# Attachment Feature Overview

Some resources, such as tasks, can have attachments. An attachment can be any type of file, such as an image, a PDF document, a zip file, etc. An attachment cannot exist independently and must be associated with a resource. Currently, the resource type that supports attachment association is "task" only. Since an attachment cannot exist independently, to add an attachment to a new task, you need first invoke the Create Task API to create the task, and then invoke the Upload Attachment API to upload files and associate them with the created task.

* Using Upload Attachment can upload one or more attachments and associate them with a certain resource at once.
* Using List Attachments can list all attachments of a resource (supporting pagination).
* Using Get Attachment can get a specific attachment information.
* Using Delete Attachment can completely delete an attachment. Deleted data cannot be restored.

For attachments belonging to the task, the authorization depends on how the calling identity is authorized by the task:
- If the calling identity has read permission on the task, it can get or download the attachment.
- If the calling identity has edit permission on the task, it can upload attachments and associate them with the task.
- If the calling identity has edit permission on the task, or the calling identity is exatly the uploader of the attachment, it can delete the attachment.

For information on task authorization, please refer to the "How are Tasks Authorized?" section in Task Feature Overview.
