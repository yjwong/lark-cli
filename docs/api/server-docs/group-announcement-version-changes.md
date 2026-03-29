---
title: "Group announcement version changes"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/group/upgraded-group-announcement/group-announcement-version-changes"
service: "group-announcement-version-changes"
section: "Group Chat"
updated: "1739935006000"
---

# Group announcement version changes

In December 2024, the Lark Open Platform introduced an upgraded version of the group announcement APIs. This article outlines how to distinguish between the old and upgraded versions of group announcement, as well as the differences in the APIs between the old and upgraded versions.

## Distinguish between upgraded and old versions

You can distinguish between the upgraded and old group announcements by calling the Obtain the basic information of a group announcement API and checking the `announcement_type` field in the response.

| **Categories** | **Announcement type related fields** | **Overview** |
| -------- | ------------ | ----------|
| Upgraded group announcement | docx         | Upgraded group announcement overview|
| Old group announcement | doc          |Old group announcement overview| ## API changes
For the following capabilities, the APIs of the upgraded version and the old version are different. You cannot mix and use them.

| **Capability** | **Upgraded version APIs** | **Old version APIs** |
|---|---|---|
| Obtain group announcement metadata|Obtain the basic information of a group announcement| Obtain group announcement information  |
| edit group announcement | Create blocks in group announcement  Batch update blocks in group announcement  Delete blocks in group announcement| Update group announcement info |
| Get group announcement rich-text content | Obtain all blocks of a group announcementObtain the block content in group announcementObtain all the child blocks | Obtain group announcement information | ## FAQs

### 1. Why not make the upgraded group announcement compatible on the basis of the old version API?

The formats and protocols of the upgraded group announcement are completely different from the old version. The original data structure and API protocols cannot be made compatible. Thus, when operating the upgraded group announcement, the new version API needs to be used.

### 2. Does the upgraded group announcement use Location to mark the position as the old group announcement?

Only the old version use Location to mark the position. For details, see Document Structure Introduction.

There is no Location concept in the upgraded version of the group announcement APIs, which is also one of the main differences between the data structure of the upgraded version and the old version of the group announcement. The upgraded version of the group announcement is based on Block ID and Parent Block ID for positioning, and it can be viewed as a tree structure of blocks.
