---
title: "Overview"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/getting-start"
service: "wiki-v2"
resource: "getting-start"
section: "Docs"
updated: "1755153169000"
---

# Overview

You can automatically manage your workspace through the Wiki API.
> **Note:** Before calling Wiki API, please make sure your app has required permission.
> - `wiki:wiki`: Can view or edit Wiki.
> - `wiki:wiki.readonly`: Can read Wiki content only.
> 
> See more details: App scopes.

# Introduction

## Workspace

Workspace is the basic unit of workspace. It is a different type of knowledge system built by enterprises according to their needs. It is composed of multiple document pages with hierarchies and affiliations. Each workspace has a unique space_id.

**The space_id of a workspace can be obtained by any of the following methods:**

-   Call [get workspace list](https://open.larksuite.com/ssl:ttdoc/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space/list), get from the return value.

-   If you are a workspace administrator, you can go to the workspace settings page and copy the number part of the address bar.

### Methods

-   [Create workspace](https://open.larksuite.com/ssl:ttdoc/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space/create)

-   [Get a list of workspace](https://open.larksuite.com/ssl:ttdoc/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space/list)

-   [Get workspace information](https://open.larksuite.com/ssl:ttdoc/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space/get)

-   [Update workspace settings](https://open.larksuite.com/ssl:ttdoc/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space-setting/update)

## Node

Each knowledge space is composed of a series of document pages. The affiliation of these document pages will be displayed in a tree shape, which is called a page tree. Each node has a unique wikiToken as its identifier.

### Methods

-   [Create node](https://open.larksuite.com/ssl:ttdoc/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space-node/create)

-   [Get node information](https://open.larksuite.com/ssl:ttdoc/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space/get_node)

-   [Get a list of child nodes](https://open.larksuite.com/ssl:ttdoc/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space-node/list)

-   [Add cloud documents to the workspace](https://open.larksuite.com/ssl:ttdoc/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space-node/move_docs_to_wiki)

## Members

There are two types of [members of workspace](https://open.larksuite.com/ssl:ttdoc/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space-member/create): administrator and ordinary member. It is supported to set an application as a knowledge base member by adding knowledge space members.
