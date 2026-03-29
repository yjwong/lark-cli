---
title: "FAQ"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/faq"
service: "drive-v1"
resource: "faq"
section: "Docs"
updated: "1665221847000"
---

# Space FAQs

## 1. Cloud space folder single-layer file limit.
| The upper limit of a single-layer node in a folder in cloud space is 1500. If the limit is exceeded, interfaces such asCreate online document、Copy File、Move File、Create FolderandUpload Filewill return a failure, and the error code is 1062507, if there is such a requirement, you can consider creating a new file in a different folder. |
| --- | ## 2. Cloud space file interface concurrency limit.
| The cloud space folder does not support concurrent calls to interfaces such asCreate online document、Copy File、Move File、Create Folder、Delete File以及Upload File.It will return a failure with an error code of 1061045. Retry can be successful. Users should try to avoid concurrent calls. |
| --- | ## 3. How to allow an app (tenant_access_token) to access a folder in a personal cloud space?
| Requires the app to enable robotics capabilities. Open Lark software, create a new group, and add the app as a group robot. Find the corresponding folder in Lark Cloud Documents, My Space, and share the folder with the newly created group. |
| --- | ## 4. How to get folder, file and various types of online document token information?
| Reference DocumentationDocs FAQs。 |
| --- |
