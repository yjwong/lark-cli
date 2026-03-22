---
title: "Feishu Card resource overview"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/cardkit-v1/feishu-card-resource-overview"
service: "cardkit-v1"
resource: "feishu-card-resource-overview"
section: "Lark Card"
updated: "1751026382000"
---

# Lark card resources overview
Lark cards are a capability of the application, including the components needed to build card content and the capabilities needed to send cards, and provide a [visual building tool](https://open.larksuite.com/cardkit?from=open_docs_overview). The Lark Open Platform provides a series of OpenAPIs for Lark cards, using these OpenAPIs you can partially or streamingly update cards at the card and component levels.

## Typical cases

The Open Platform provides cases containing Lark cards, for details, please refer to:
- [Smart approval management, helping to improve corporate efficiency](https://open.larksuite.com/solutions/detail/automation)
- [Smart dispatch of operation and maintenance work orders, smooth and accurate information flow](https://open.larksuite.com/solutions/detail/ticke)

## Access process

The basic access process of the card API is shown in the figure below. For detailed API access process, please refer to process overview.

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/17bab04d388fb1ebec5e96ff95833e91_xYFPzWS4ZV.png?height=214&lazyload=true&width=2920)

## Development guide

- Refer to the Streaming updates OpenAPI calling guide to learn how to call the card interface.

- Visit the Introduction of Message cards to learn how to build, send, and update cards.

## Resource introduction

In Lark card OpenAPI, interfaces are opened centered on card and component resources.
| Resource | Description |
| --- | --- |
| Card | Lark cards can embed structured content in the form of cards into Lark collaboration scenarios such as chat messages, group pinned messages, and link previews, improving information transmission efficiency. For more information, see Introduction of Message cards. Through the card entity API, you can create and update cards from the card dimension. |
| Component | Components in Lark cards can be divided into container, display, and interactive components. For more information, see component overview. Through the component API, you can add and modify components in a card. | ## Method list

The following provides the API and event list for card business. The following interfaces only support card JSON 2.0 structure.
> **Note:** - The message business also provides a series of APIs for sending cards, etc., for card-type messages. For details, refer to message card resources overview.
> - In the table below, **store** refers to store applications, and **custom** refers to enterprise self-built applications. For a description of application types, see application types introduction.

### Card level

    

    Method (API)
    Permission requirements (meet any one)
    Access token
    Store
    Self-built

    
    

    
``POST`Create card entity open-apis/cardkit/v1/cards`
    
    
 `cardkit:card:write`
    
    
`tenant_access_token`

    
    **✓**
    **✓**

    
``PUT` Update card entity in full /open-apis/cardkit/v1/cards/:card_id`
    
    
 `cardkit:card:write`
    
    
`tenant_access_token`

    
    **✓**
    **✓**

    
``PATCH` Update card settings /open-apis/cardkit/v1/cards/:card_id/settings`
    
    
 `cardkit:card:write`
      
    
    
`tenant_access_token`

    
    **✓**
    **✓**

    
``POST` Batch update card entities /open-apis/cardkit/v1/cards/:card_id/batch_update`
    
    
 `cardkit:card:write`
    
    
`tenant_access_token`

    
    **✓**
    **✓**

### Component level

    

    Method (API)
    Permission requirements (meet any one)
    Access token
    Store
    Self-built

    
    

    
``POST`Add component /open-apis/cardkit/v1/cards/:card_id/elements`
    
    
 `cardkit:card:write`
    
    
`tenant_access_token`

    
    **✓**
    **✓**

    
``PUT` Update component /open-apis/cardkit/v1/cards/:card_id/elements/:element_id`
    
    
 `cardkit:card:write`
    
    
`tenant_access_token`

    
    **✓**
    **✓**

    
``PATCH` Update component properties /open-apis/cardkit/v1/cards/:card_id/elements/:element_id`
    
    
 `cardkit:card:write`
      
    
    
`tenant_access_token`

    
    **✓**
    **✓**

    
``PUT` Stream update text /open-apis/cardkit/v1/cards/:card_id/elements/:element_id/content`
    
    
 `cardkit:card:write`
    
    
`tenant_access_token`

    
    **✓**
    **✓**

    
``DELETE` Delete component /open-apis/cardkit/v1/cards/:card_id/elements/:element_id`
    
    
           `cardkit:card:write`

     
    
`tenant_access_token`

    
    **✓**
    **✓**
