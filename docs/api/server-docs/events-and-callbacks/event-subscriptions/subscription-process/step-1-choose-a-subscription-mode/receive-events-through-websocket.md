---
title: "Receive events through websocket"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uYDNxYjL2QTM24iN0EjN/event-subscription-configure-/use-websocket"
section: "Events and callbacks"
updated: "1773324708000"
---

# Receive events through websocket

Websockets are capabilities provided within the Lark SDK. You can integrate the Lark SDK into your local server to establish a full-duplex WebSocket channel with the open platform (your server needs to be able to access the public internet). Subsequently, when subscribed events occur, the open platform will send event messages to your server through this channel.

## Advantages

Compared to the traditional Webhook mode, the persistent connection mode significantly reduces integration costs, shortening the development cycle from around one week to just five minutes. Specific advantages include:

- **Convenient and Quick Development**

    It only requires that the runtime environment can access the public internet. There's no need to provide a public IP or domain name, nor use an internal network penetration tool. Events can be received in the local development environment through the persistent connection mode, and once the local service is deployed online, it can work immediately.
- **Encrypted Transmission, No Need for Extra Encryption/Decryption Logic**

    The persistent connection method includes built-in communication encryption and authentication logic. Event data is encrypted during transmission over the network, requiring no additional decryption or signature verification logic from the developers, nor the deployment of firewalls and configuration of whitelists.

## Precautions

Before starting the configuration, you need to ensure you understand the following points:

- The persistent connection mode only supports enterprise self-built applications. For store application event subscription methods, refer to Send notifications to developer's server.
- After receiving a message in the persistent connection mode, it needs to be processed within 3 seconds without throwing an exception, otherwise it will trigger the timeout retry mechanism.
- Each application can establish up to 50 connections (each initialization of a client in persistent connection configuration counts as one connection).
- Message push in the persistent connection mode is **cluster mode** and does not support broadcasting. That is, if the same application deploys multiple clients, only **one random** client among them will receive the message.

## Step 1: Integrate SDK

Below are example codes for the Lark SDK in Golang, Python, Java, and Node.js. The configuration within the code uses the Receive Messages event as an example. In addition to event subscription, Lark SDK can also achieve API calls and subscription callbacks. For a detailed introduction to the SDK, refer to Server-side SDK.

### Golang

#### Install SDK

```
go get -u github.com/larksuite/oapi-sdk-go/v3
```

#### Example Code

```go
package main
import (
   "context"
   "fmt"
   lark "github.com/larksuite/oapi-sdk-go/v3"
   larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
   larkevent "github.com/larksuite/oapi-sdk-go/v3/event"
   "github.com/larksuite/oapi-sdk-go/v3/event/dispatcher"
   larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
   larkws "github.com/larksuite/oapi-sdk-go/v3/ws"
)
func main() {
   // Register event callback, OnP2MessageReceiveV1 is for receiving messages v2.0; message within OnCustomizedEvent is for receiving messages v1.0.
   eventHandler := dispatcher.NewEventDispatcher("", "").
      OnP2MessageReceiveV1(func(ctx context.Context, event *larkim.P2MessageReceiveV1) error {
         fmt.Printf("[ OnP2MessageReceiveV1 access ], data: %s\n", larkcore.Prettify(event))
         return nil
      }).
      OnCustomizedEvent("insert your customized subscribed event key here, such as out_approval", func(ctx context.Context, event *larkevent.EventReq) error {
         fmt.Printf("[ OnCustomizedEvent access ], type: message, data: %s\n", string(event.Body))
         return nil
      })
   // Create Client
   cli := larkws.NewClient("YOUR_APP_ID", "YOUR_APP_SECRET",
      larkws.WithDomain(lark.LarkBaseUrl),
      larkws.WithEventHandler(eventHandler),
      larkws.WithLogLevel(larkcore.LogLevelDebug),
   )
   // Start client
   err := cli.Start(context.Background())
   if err != nil {
      panic(err)
   }
}
```

Explanation of code implementation:
1. Initialize the event handler (eventHandler) through dispatcher.NewEventDispatcher(), note that **both parameters must be empty strings**.
1. Use the OnXXXX() methods of eventHandler to listen for different types of events; the example above listens for "receive message v1.0" and "receive message v2.0" events respectively.
   
   **Note**: There are two types of event structures, v1.0 version event structure and v2.0 version event structure. When handling events, pay attention to which event structure is required. When adding events for an application, you can confirm the event version through the version identifier (as shown in the figure below, the event version is v2.0).
   
   ![](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/8bf0504742fd12133246221d42432191_sjzTH7TScQ.png?height=368&lazyload=true&maxWidth=750&width=1630)
   
1. Initialize the long connection client through larkws.NewClient(). The required parameters are the application's APP_ID and APP_SECRET, which can be obtained by logging in to the [Developer Console](https://open.larksuite.com/app), in the **Credentials & Basic Info** > **Credentials** section of the application details page.
    
    ![](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/369e277da843a70b55db209353b98c97_inP55mum9X.png?height=358&lazyload=true&maxWidth=750&width=2500)
1. Pass the optional parameters eventHandler and log level settings.
1. Start the client using cli.Start(). If the connection is successful, the console will print "connected to wss://xxxxx", and the main thread will block until the process ends.
   
   ![](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/b1b3d6b4d3a724fa7408f5c71a11d19e_ysuov6r3Bn.png?height=338&lazyload=true&maxWidth=750&width=2286)
   
### Python

#### Install SDK

```
pip install lark-oapi -U
```

#### Example Code

```python
import lark_oapi as lark

## P2ImMessageReceiveV1 is for receiving messages v2.0; message within CustomizedEvent is for receiving messages v1.0.
def do_p2_im_message_receive_v1(data: lark.im.v1.P2ImMessageReceiveV1) -> None:
    print(f'[ do_p2_im_message_receive_v1 access ], data: {lark.JSON.marshal(data, indent=4)}')
    
def do_message_event(data: lark.CustomizedEvent) -> None:
    print(f'[ do_customized_event access ], type: message, data: {lark.JSON.marshal(data, indent=4)}')
event_handler = lark.EventDispatcherHandler.builder("", "") \
    .register_p2_im_message_receive_v1(do_p2_im_message_receive_v1) \
    .register_p1_customized_event("Here fill in the key of the event you want to customize for subscription, such as out_approval", do_message_event) \
    .build()

def main():
    cli = lark.ws.Client("YOUR_APP_ID", "YOUR_APP_SECRET",
                         event_handler=event_handler,
                         log_level=lark.LogLevel.DEBUG,
                         domain=lark.LARK_DOMAIN)
    cli.start()
    
if __name__ == "__main__":
    main()
```

Code Implementation Explanation:

1. Initialize the event handler (event_handler) using `lark.EventDispatcherHandler.builder()`, and both parameters of this method must be empty strings.

2. Handle different events through the `register_{event_version}_{event_type or customized_event}` method of the event_handler.

    When adding events on the application details page in the developer backend (specific operations can be seen in Adding Events), you can obtain the type and version information of the events.

    ![](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/a881d5c42f70c633891106b846734e8e_Q3fdvEc9KA.png?height=410&lazyload=true&maxWidth=600&width=1660)

    **Events of version v2.0**
    
    - Use `data: lark.im.v1.P2ImMessageReceiveV1` to specify the selected event, where:
    
        - `lark.im.v1` represents the corresponding business domain and version, and you can get the business domain code prefix from the event type, such as `im.message.receive_v1` corresponds to the business domain code `im`.
    
        - In `P2ImMessageReceiveV1`, **P2** indicates the event version as v2.0, and **ImMessageReceiveV1** corresponds to the event type `im.message.receive_v1`.

    - Use `.register_p2_im_message_receive_v1()` to handle the event, where:
    
        - **p2** indicates the event version as v2.0.
    
        - **im_message_receive_v1** corresponds to the event type `im.message.receive_v1`.

    :::note
    **Note**: The method format in the code differs from the event type, be aware of inconsistencies in format to avoid program failure due to format issues.
    :::
    
    **Events of version v1.0**

    - Always use the `data: lark.CustomizedEvent` method to handle customized events.
    - Use `.register_p1_customized_event("insert your event type here, such as out_approval", do_message_event)` to handle the event:
        - **p1** indicates the event version as v1.0.
        - **customized_event** is the fixed method name.
        - Inside the customized_event method, the event Key (of type String) needs to be passed in, for example, the event type for Out of Office Approval is out_approval.
    
    You can also refer to the comments in the code snippet to understand how to configure it:
    
    ```python
    ## P2ImMessageReceiveV1:
    ## P2 corresponds to version 2.0,
    ## ImMessageReceiveV1 corresponds to the receive message event type im.message.receive_v1.
    def do_p2_im_message_receive_v1(data: lark.im.v1.P2ImMessageReceiveV1) -> None:
        print(f'[ do_p2_im_message_receive_v1 access ], data: {lark.JSON.marshal(data, indent=4)}')

    ## For v1.0 events, use lark.CustomizedEvent
    def do_message_event(data: lark.CustomizedEvent) -> None:
        print(f'[ do_customized_event access ], type: message, data: {lark.JSON.marshal(data, indent=4)}')

    event_handler = lark.EventDispatcherHandler.builder("", "") \
        ## register_p2 corresponds to version 2.0,
        ## im_message_receive_v1 corresponds to the receive message event type im.message.receive_v1.
        .register_p2_im_message_receive_v1(do_p2_im_message_receive_v1) \
        ## register_p1 corresponds to version 1.0,
        ## the parameter value needs to be the event type, such as out_approval.
        .register_p1_customized_event("Here fill in the key of the event you want to customize for subscription, such as out_approval", do_message_event) \
        .build()
    ```

3. Initialize the long connection client using `lark.ws.Client()`, with the mandatory parameters being the application's APP_ID and APP_SECRET, which can be obtained from the **Credentials & Basic Info** area of the application details page on the [Developer Console](https://open.larksuite.com/app).
    
    ![](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/eb3e4629fcf7165b605b27745c48766b_qtlk84KIfC.png?height=382&lazyload=true&maxWidth=600&width=2602)

4. Pass the event_handler as an optional parameter, and you can also set the log level.
5. Start the client with `cli.start()`. If the connection is successful, the console will print "connected to wss://xxxxx," and the main thread will block until the process ends.
    
    ![](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/2085839331117b6f9817482a5493eba6_mDSqdoypa5.png?height=264&lazyload=true&maxWidth=600&width=2488
    
### Java

#### Install SDK

The latest version of the SDK can be found [here](https://mvnrepository.com/artifact/com.larksuite.oapi/oapi-sdk).

```xml

  ...
  
    com.larksuite.oapi
    oapi-sdk
    {latest version}
  
  ...

```

#### Example Code

```java
package com.lark.oapi.sample.ws;
import com.lark.oapi.core.enums.BaseUrlEnum;
import com.lark.oapi.core.request.EventReq;
import com.lark.oapi.core.utils.Jsons;
import com.lark.oapi.event.CustomEventHandler;
import com.lark.oapi.event.EventDispatcher;
import com.lark.oapi.service.im.ImService;
import com.lark.oapi.service.im.v1.model.P2MessageReceiveV1;
import com.lark.oapi.ws.Client;
import java.nio.charset.StandardCharsets;

public class Sample {
    // onP2MessageReceiveV1 is used to receive messages v2.0; onCustomizedEvent's message receives messages v1.0
    private static final EventDispatcher EVENT_HANDLER = EventDispatcher.newBuilder("", "")
            .onP2MessageReceiveV1(new ImService.P2MessageReceiveV1Handler() {
                @Override
                public void handle(P2MessageReceiveV1 event) throws Exception {
                    System.out.printf("[ onP2MessageReceiveV1 access ], data: %s\n", Jsons.DEFAULT.toJson(event.getEvent()));
                }
            })
            .onCustomizedEvent("your_event_key_here", new CustomEventHandler() {
                @Override
                public void handle(EventReq event) throws Exception {
                    System.out.printf("[ onCustomizedEvent access ], type: message, data: %s\n", new String(event.getBody(), StandardCharsets.UTF_8));
                }
            })
            .build();

    public static void main(String[] args) {
        Client cli = new Client.Builder("YOUR_APP_ID", "YOUR_APP_SECRET")
                .domain(BaseUrlEnum.LarkSuite.getUrl())
                .eventHandler(EVENT_HANDLER)
                .build();
        cli.start();
    }
}
```

Explanation of the code:
1. Initialize the event handler (`EVENT_HANDLER`) using `EventDispatcher.newBuilder()` (note: the two parameters should be empty strings).
2. Use `onXXXX()` methods of the `EVENT_HANDLER` to handle different events. The example above listens to both "receive message v1.0" and "receive message v2.0" events.

   **Note**: Events come in two structures: v1.0 and v2.0. You need to pay attention to which event structure is required when handling the event. When adding an event to the application, you can identify the event version via version tags in the Event Subscription Configuration.

3. Initialize the long connection client using `new Client.Builder()` with the required parameters `APP_ID` and `APP_SECRET`, which can be found in the **Credentials & Basic Info** > **App Credentials** section of the developer console [here](https://open.larksuite.com/app).

4. Optionally, pass the `EVENT_HANDLER` and set the log level, which can be configured in the project's logging framework (e.g., log4j2, logback).
5. Use `cli.start()` to start the client. If the connection is successful, the console will display "connected to wss://xxxxx", and the main thread will block until the process ends.

   ![](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/90d8aab594e3fa8ca22bf4b291723efd_8BitTMiPGo.png?height=176&lazyload=true&maxWidth=692&width=2992)
   
### Node.js

#### Install SDK

The version must be greater than or equal to `1.24.0`.

```
// Use the package manager of your project, such as yarn
yarn add @larksuiteoapi/node-sdk
```

#### Example Code

```javascript
import * as Lark from '@larksuiteoapi/node-sdk';

const baseConfig = {
  appId: 'xxx',
  appSecret: 'xxx'
}

const client = new Lark.Client(baseConfig);
const wsClient = new Lark.WSClient({...baseConfig, loggerLevel: Lark.LoggerLevel.debug, domain: Lark.Domain.Lark});

wsClient.start({
  // Handle "receive message" event, event type is im.message.receive_v1
  eventDispatcher: new Lark.EventDispatcher({}).register({
    'im.message.receive_v1': async (data) => {
      const {
        message: { chat_id, content}
      } = data;
      // Example operation: After receiving the message, call the "send message" API to reply to the message.
      await client.im.v1.message.create({
        params: {
          receive_id_type: "chat_id"
        },
        data: {
          receive_id: chat_id,
          content: Lark.messageCard.defaultCard({
            title: `Reply: ${JSON.parse(content).text}`,
            content: 'Happy New Year'
          }),
          msg_type: 'interactive'
        }
      });
    }
  })
});
```

Description of the code implementation:

1. Initialize WSClient.
    
    1. In baseConfig, pass in the app's appId and appSecret. To obtain: Log in to the [Developer Console](https://open.larksuite.com/app), and on the **Credentials and Basic Info** page in the app details, retrieve the app credentials (including App ID and App Secret).
        
        ![](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/3c9c300e0a50875fa74070945bc38bdf_CuNGaPfWjp.png?height=364&lazyload=true&maxWidth=600&width=2612)
    
    2. Optionally add other parameters such as domain, with the default domain being `open.feishu.cn`.
        

2. When the initialization is completed based on the Client, the start method on the instance is used to start the long connection client.
    
    In the start function, the eventDispatcher parameter accepts the EventDispatcher instance, and once the long connection client is successfully started, the relevant event handlers registered by eventDispatcher.register will be executed.
    
    In the following example code, replace `'im.message.receive_v1'` with the **event type** of the event you need to handle. The **event type** can be found in the documentation for the specific event, for example, check the Receive message event documentation to know that the event type is **im.message.receive_v1**.
    
    ```javascript
    wsClient.start({
      eventDispatcher: new Lark.EventDispatcher({}).register({
        'im.message.receive_v1': async (data) => {
          const {
            message: { chat_id, content}
          } = data;
    ```
    
    If the long connection is successfully started, the following information will be printed.
    
    ![](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/a5139264b8e6fcbbeafce287aa86efec_YGMzsCtK7W.png?height=694&lazyload=true&maxWidth=600&width=3124)
    
## Step 2: Set up subscription method

1. Start the local long connection client.
   
   Refer to the specific operations in the above **Step 1: Integrate SDK**.

2. Log in to the [Developer Console](https://open.larksuite.com/app), and select an enterprise self-built application.
   
   Currently, long connection mode does not support store applications.

3. Go to the **Events & Callbacks > Event Configuration** page.

4. Edit the subscription method, select **Receive events through persistent connection**, and save.

    :::warning
    Make sure the local client is running normally and the long connection is online, in order to save successfully.
    :::

    ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/b0485c4604662c05eb0a6f5c9e53dcfe_mA2ROSe8XR.png?height=886&lazyload=true&width=1702)

## FAQ

If you encounter problems during configuration, you can refer to the FAQ for troubleshooting solutions.
