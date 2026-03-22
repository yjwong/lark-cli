---
title: "Board overview"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/board-v1/overview"
service: "board-v1"
resource: "overview"
section: "Docs"
updated: "1744960855000"
---

#  Board overview
Board is a brand-new graphic creation tool that has the characteristics of low cost, efficiency, and convenient collaboration. Using the board, you can easily draw perfect flow charts, planning diagrams and solution diagrams, and you can work with your team on the board for real-time graphic collaboration. Through the board API, the board can be connected to the internal business system, making the board a part of your business process.

##  Basic concepts
There are two basic concepts in the board APIs: board and node.

### Board
A board is an online board created by the user in the document. Each board is identified by a unique whiteboard_id.
> For whiteboard_id, you can call the Obtain all the child blocks API to get. In this context, the block.token which block type is 43 serves as the whiteboard_id for that specific board.

The basic metadata structure of a board is as follows:

```json 
 "whiteboard": {
   "id": string, // The unique identifier of the board.
   "title": string, // The title of the board.
   "theme": int, // The theme of the board.
 }
``` 

###  Node
Within a board, there are multiple types of graphics, each defined as a node. Every board consists of a tree of nodes. The node points to its parent node through parent_id. When the parent_id is empty, it indicates that the parent node does not exist and the node is directly mounted under the board. The node points to its child node through children_id. Node has many properties associated with them. Some of these are base properties and exist on every node, while some node properties will be specific to the node type. For specific node properties, see data structure。

## Method list
> There is currently no board level API interface, only node level API interface.

| Method (API) | Permission requirements (meet either) | Access Credentials | Store | Custom |
| --- | --- | --- | --- | --- |
| GET**Obtain all nodes of a board** /open-apis/board/v1/whiteboards/:whiteboard_id/nodes | :::html `board:whiteboard:node:read` ::: | `tenant_access_token` `user_access_token` | **✓** | **✓** |
