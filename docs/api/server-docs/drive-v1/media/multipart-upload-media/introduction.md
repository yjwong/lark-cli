---
title: "Introduction"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/media/multipart-upload-media/introduction"
service: "drive-v1"
resource: "media"
section: "Docs"
updated: "1665221865000"
---

# Overview

This document provides an overview of the API used to upload a material in blocks.

This upload method is recommended for files larger than 20 MB and scenarios with high network interruption probabilities. You can upload files in fixed-length blocks to improve the success rate and reduce bandwidth usage, while displaying the upload progress of blocks. With this method, you can resume uploading based on an upload transaction ID and an upload block number within the day when you apply for the upload transaction ID.

## Methods

### Upload a material in blocks (Pre­uploading)

Sends an initialization request to obtain an upload transaction ID and a split policy. The current policy is to split a file into 4 MB fixed-length blocks.

### Upload a material in blocks (or Upload a block)

Uploads a block of a specified file.

### Complete uploading a material in blocks

Completes an upload task.
