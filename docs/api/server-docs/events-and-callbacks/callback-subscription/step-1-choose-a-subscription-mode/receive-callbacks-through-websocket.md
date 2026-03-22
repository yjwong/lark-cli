---
title: "Receive callbacks through websocket"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/event-subscription-guide/callback-subscription/configure-callback-request-address"
section: "Events and callbacks"
updated: "1773324712000"
---

# Receive callbacks through websocket

A websocket is a capability provided by the Lark SDK. You can integrate the Lark SDK on your local server and establish a WebSocket full-duplex channel with the open platform (your server needs to be able to access the public internet). Subsequently, when the subscribed callbacks of the application occur, the open platform will send callback messages to your server through this channel.

## Advantages

Compared to the traditional Webhook model, the persistent connection mode greatly reduces the access cost, decreasing the original development cycle of approximately one week to five minutes. The specific advantages are as follows:

- **Convenient and efficient development calls**
  
  You only need to ensure that the operating environment has the capability to access the public internet; there is no need to provide a public IP or domain name, and no need to use intranet penetration tools. You can receive callback messages in the local development environment via persistent connection mode, and it will directly take effect even after deploying the local service online later.

- **Encrypted transmission, no need for additional encryption and decryption logic**
  
  The persistent connection method has built-in communication encryption and authentication logic. The callback data is transmitted encrypted over the network, so developers don't need additional decryption or signature verification logic. There's also no need to deploy firewalls and configure whitelists.

Configure the callback subscription method to receive callback messages pushed by the open platform to the application. When the subscribed callbacks of the application (such as Lark card interaction, link preview) occur, the open platform will perform callbacks in the specified manner.

### Usage restrictions

- Through persistent connection mode only supports enterprise self-built applications.

- Message card return interaction (old) callback does not support the **Receive callbacks through persistent connection** subscription method. You can only choose the **Send callbacks to developer's server** subscription method.

- Each application can establish up to 50 connections (when configuring a long connection, each client initialized is a connection).

### Precautions

- The same requirements as **Send callbacks to developer's server**, after receiving a message in the long connection mode, it also needs to be processed within 3 seconds.

- The message push in the long connection mode is **cluster mode**, and broadcasting is not supported. That is, if multiple clients are deployed for the same application, only one random client will receive the message.

## Step 1: Integrate SDK

The following sample code takes Card callback interaction and Pull link preview data callbacks as examples.

### Go SDK

#### Installation

```
go get -u github.com/larksuite/oapi-sdk-go/v3@latest
```

#### Sample code

```Go
package main

import (
    "context"
    "fmt"
    lark "github.com/larksuite/oapi-sdk-go/v3"
    larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
    "github.com/larksuite/oapi-sdk-go/v3/event/dispatcher"
    "github.com/larksuite/oapi-sdk-go/v3/event/dispatcher/callback"
    larkws "github.com/larksuite/oapi-sdk-go/v3/ws"
)
func main() {
    // Register callback
    eventHandler := dispatcher.NewEventDispatcher("", "").
       // Listen for the 「card return interaction card.action.trigger」 event
       OnP2CardActionTrigger(func(ctx context.Context, event *callback.CardActionTriggerEvent) (*callback.CardActionTriggerResponse, error) {
          fmt.Printf("[ OnP2CardActionTrigger access ], data: %s\n", larkcore.Prettify(event))
          return nil, nil
       }).
       // Listen for the 「pull link preview data url.preview.get」 event
       OnP2CardURLPreviewGet(func(ctx context.Context, event *callback.URLPreviewGetEvent) (*callback.URLPreviewGetResponse, error) {
          fmt.Printf("[ OnP2URLPreviewAction access ], data: %s\n", larkcore.Prettify(event))
          return nil, nil
       })
    // Create client
    cli := larkws.NewClient("YOUR_APP_ID", "YOUR_APP_SECRET",
       larkws.WithDomain(lark.LarkBaseUrl),
       larkws.WithEventHandler(eventHandler),
       larkws.WithLogLevel(larkcore.LogLevelDebug),
    )
    // Establishing a persistent connection
    err := cli.Start(context.Background())
    if err != nil {
       panic(err)
    }
}
```

Steps:
1. Initialize the event handler (eventHandler) through dispatcher.NewEventDispatcher(). **Note that the two parameters must be filled with empty strings**.
2. Listen to different callback types through the OnXXXX() method of eventHandler. In the above example, **card return interaction card.action.trigger** and **pull link preview data url.preview.get** callbacks are listened to respectively.
3. Initialize the long connection client through larkws.NewClient(). The required parameters are the APP_ID and APP_SECRET of the application. You can go to **Basic Info > Credentials & Basic Info** in the application details page of [Developer Console](https://open.larksuite.com/app) to obtain the APP_ID and APP_SECRET of the application.
4. Optional parameters are passed to eventHandler, and the log level can be set at the same time.
5. Start the client through cli.Start(). If the connection is successful, the console will print `connected to wss://xxxxx`, and the main thread will be blocked until the process ends.

	![](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/ee46554b73c969db98c96c31a4d85d13_ez0VWNNi3V.png?height=188&lazyload=true&maxWidth=600&width=2796)

### Python SDK

#### Installation

```
pip install lark-oapi==1.4.0
```

#### Sample code

```python
import lark_oapi as lark
from lark_oapi.event.callback.model.p2_card_action_trigger import P2CardActionTrigger, P2CardActionTriggerResponse
from lark_oapi.event.callback.model.p2_url_preview_get import P2URLPreviewGet, P2URLPreviewGetResponse

# Listen for the 「card return interaction card.action.trigger」 event
def do_card_action_trigger(data: P2CardActionTrigger) -> P2CardActionTriggerResponse:
    print(lark.JSON.marshal(data))
    resp = {
        "toast": {
            "type": "info",
            "content": "Card sent back successfully from python sdk"
        }
    }
    return P2CardActionTriggerResponse(resp)

# Listen for the 「pull link preview data url.preview.get」 event
def do_url_preview_get(data: P2URLPreviewGet) -> P2URLPreviewGetResponse:
    print(lark.JSON.marshal(data))
    resp = {
        "inline": {
            "title": "Link Preview Test",
        }
    }
    return P2URLPreviewGetResponse(resp)
event_handler = lark.EventDispatcherHandler.builder("", "") \
    .register_p2_card_action_trigger(do_card_action_trigger) \
    .register_p2_url_preview_get(do_url_preview_get) \
    .build()
def main():
    cli = lark.ws.Client(lark.APP_ID, lark.APP_SECRET,
                         event_handler=event_handler, log_level=lark.LogLevel.DEBUG, domain=lark.LARK_DOMAIN)
    cli.start()
if __name__ == "__main__":
    main()
```

Steps:
1. Initialize the event handler (event_handler) through lark.EventDispatcherHandler.builder(). **Two parameters must be filled with empty strings**.
2. Listen to different event types through the register_xxxx() method of event_handler. In the above example, **card return interaction card.action.trigger** and **pull link preview data url.preview.get** callbacks are listened.
3. Register the card callback processing function through register_p2_card_action_trigger.
4. Register the link preview callback processing function through register_p2_url_preview_get.
5. Initialize the long connection client through lark.ws.Client(). The required parameters are the application's APP_ID and APP_SECRET. The application's APP_ID and APP_SECRET can be obtained in the [Developer Backend](https://open.larksuite.com/app).
6. Optional parameters are passed to event_handler, and the log level can be set at the same time.
7. Start the client through cli.start(). If the connection is successful, the console will print `connected to wss://xxxxx`, and the main thread will be blocked until the process ends.

	![](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/ac590caeb5cc3fe42ca3855696f42c6e_4ZrUub9NAL.png?height=264&lazyload=true&maxWidth=600&width=2488)

### Java SDK

#### Installation

```xml

  ...
  
    com.larksuite.oapi
    oapi-sdk
    2.4.0
  
  ...

```

#### Sample code

```java
package com.lark.oapi.sample.ws;
import com.lark.oapi.core.enums.BaseUrlEnum;
import com.lark.oapi.core.request.EventReq;
import com.lark.oapi.core.utils.Jsons;
import com.lark.oapi.event.CustomEventHandler;
import com.lark.oapi.event.EventDispatcher;
import com.lark.oapi.event.cardcallback.P2CardActionTriggerHandler;
import com.lark.oapi.event.cardcallback.P2URLPreviewGetHandler;
import com.lark.oapi.event.cardcallback.model.*;
import com.lark.oapi.service.im.ImService;
import com.lark.oapi.service.im.v1.model.P2MessageReceiveV1;
import com.lark.oapi.ws.Client;
import java.nio.charset.StandardCharsets;

public class Sample {
    private static final EventDispatcher EVENT_HANDLER = EventDispatcher.newBuilder("", "")  // Long connection does not require these two parameters, please keep them empty strings
            // Listen for the 「card return interaction card.action.trigger」 event
            .onP2CardActionTrigger(new P2CardActionTriggerHandler() {
                @Override
                public P2CardActionTriggerResponse handle(P2CardActionTrigger event) throws Exception {
                    System.out.printf("[ P2CardActionTrigger access ], data: %s\n", Jsons.DEFAULT.toJson(event.getEvent()));
                    P2CardActionTriggerResponse resp = new P2CardActionTriggerResponse();
                    CallBackToast toast = new CallBackToast();
                    toast.setType("info");
                    toast.setContent("Card interaction successful from Java SDK");
                    resp.setToast(toast);
                    return resp;
                }
            })
            // Listen for the 「pull link preview data url.preview.get」 event
            .onP2URLPreviewGet(new P2URLPreviewGetHandler() {
                @Override
                public P2URLPreviewGetResponse handle(P2URLPreviewGet event) throws Exception {
                    System.out.printf("[ P2URLPreviewGet access ], data: %s\n", Jsons.DEFAULT.toJson(event.getEvent()));
                    P2URLPreviewGetResponse resp = new P2URLPreviewGetResponse();
                    URLPreviewGetInline inline = new URLPreviewGetInline();
                    inline.setTitle("Link preview test from Java SDK");
                    resp.setInline(inline);
                    return resp;
                }
            })
            .build();
    public static void main(String[] args) {
        Client client = new Client.Builder("", "")
                .domain(BaseUrlEnum.LarkSuite.getUrl())
                .eventHandler(EVENT_HANDLER)
                .build();
        client.start();
    }
}
```

### Node SDK

#### Installation

```
npm install @larksuiteoapi/node-sdk@1.36.0
```

#### Sample code

```Javascript
import * as Lark from "@larksuiteoapi/node-sdk";
const wsClient = new Lark.WSClient({
  appId: "YOUR_APP_ID",
  appSecret: "YOUR_APP_SECRET",
  domain: Lark.Domain.Lark
});
const eventDispatcher = new Lark.EventDispatcher({}).register({
  "card.action.trigger": async (data) => {
    console.log(data);
    return {
      toast: {
        type: "success",
        content: "Card interaction successful",
        i18n: {
          zh_cn: "Card interaction successful",
          en_us: "card action success",
        },
      },
    };
  },
});
wsClient.start({ eventDispatcher });
```

Steps:
1. Initialize the persistent client through new Lark.WSClient(). The required parameters are the APP_ID and APP_SECRET of the application. You can go to **Basic Info > Credentials & Basic Info** in the application details page of [Developer Backstage](https://open.larksuite.com/app) to obtain the APP_ID and APP_SECRET of the application.
2. Initialize the event handler (eventDispatcher) through new Lark.EventDispatcher.
3. Use the register method of eventDispatcher to listen to different callback types. In the above example, **Card return interaction card.action.trigger** callback is listened.
4. Start the client through wsClient.start. If the connection is successful, the console will print `[info]: [ "[ws]", "ws client ready" ]`, and the main thread will be blocked until the process ends.
	![](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/a94ad9fd204d20ebadd0df1091cc47cf_3ZL3KlUTlh.png?height=228&lazyload=true&maxWidth=600&width=4130)
    
    
## Step 2: Switch subscription method

1. Log in to [Developer Console](https://open.larksuite.com/app) and enter the specified enterprise self-built application details page.

	Currently, the long connection mode does not support store applications.

2. Enter the **Events & Callbacks > Callback Configuration** page.

3. Edit the subscription method, select **Receive callbacks through persistent connection**, and click **Save**.

	:::warning
	At this time, the local long connection program built based on the SDK must be connected and running to be successfully saved.
    :::

	![](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/6c768016bc6700d6e1acb6b270b5e84c_JNM0kGF2Wr.png?height=508&lazyload=true&maxWidth=600&width=1616)
