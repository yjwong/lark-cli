---
title: "Server SDK"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uETO1YjLxkTN24SM5UjN"
section: "API Call Guide"
updated: "1699608965000"
---

# Server SDK

To help developers conveniently develop apps using Lark's open capabilities and simplify the operational flow, the open platform provides a unified server SDK. Developers can use the SDK to develop features quickly and improve development efficiency.

The main capabilities provided by the SDK include:

- The SDK provides structured input parameters for API requests. For example, for the send message API, the SDK provides structured encapsulation for various types of messages.

	![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/2970b1abc05de926538f3f58465ef603_mkrkodCXSt.png?height=426&lazyload=true&maxWidth=700&width=1246)

- The SDK provides complete access token lifecycle management capabilities, eliminating the need for developers to obtain and refresh tokens themselves.
- The SDK provides concise, clear, and easy-to-understand documentation.
    - [Go SDK Documentation](https://github.com/larksuite/oapi-sdk-go/blob/v3_main/README.md)
    - [Python SDK Documentation](https://github.com/larksuite/oapi-sdk-python/blob/v2_main/README.md)
    - [Java SDK Documentation](https://github.com/larksuite/oapi-sdk-java/blob/v2_main/README.md)
    - [NodeJS SDK Documentation](https://github.com/larksuite/node-sdk/blob/main/README.zh.md)

- The SDK includes text annotations for APIs and events, as well as links to usage demos and official documentation.

Currently, we provide SDKs in four languages: Go, Python, Java, and NodeJS. You can check the details on the GitHub project page. If you encounter any issues while using the SDK, you can submit an issue to us or join the discussion group.

**GitHub Projects**                                          | **Issues**                                               | **Demo** | **Language**    |
| -------------------------------------------------------- | -------------------------------------------------------- | -------- | ------------ |
| [oapi-sdk-go](https://github.com/larksuite/oapi-sdk-go)   | [Issues](https://github.com/larksuite/oapi-sdk-go/issues)   | [oapi-sdk-go-demo](https://github.com/larksuite/oapi-sdk-go-demo)      | Golang >= 1.5 |
| [oapi-sdk-python](https://github.com/larksuite/oapi-sdk-python) | [Issues](https://github.com/larksuite/oapi-sdk-python/issues) | [oapi-sdk-python-demo](https://github.com/larksuite/oapi-sdk-python-demo)     | Python >= 3.8    |
| [oapi-sdk-java](https://github.com/larksuite/oapi-sdk-java) | [Issues](https://github.com/larksuite/oapi-sdk-java/issues) | [oapi-sdk-java-demo](https://github.com/larksuite/oapi-sdk-java-demo)     | Java >= 1.8    |
| [oapi-sdk-nodejs](https://github.com/larksuite/node-sdk)  | [Issues](https://github.com/larksuite/node-sdk/issues)      | -   | NodeJS

## Server SDK Downloads

You can refer to the content of this section to download the SDK for various platforms.

### Go

   Execute the following command to install the latest version of the Go SDK

  ```shell
  go get -u github.com/larksuite/oapi-sdk-go/v3
  ``` 
  
### Python

  Use pip to install the latest version of Python SDK

  ```shell
  pip install lark-oapi -U
  ``` 

### Java

   Add the following dependency to the pom.xml file:

  ```xml
  
      com.larksuite.oapi
      oapi-sdk
      {latest version}
  
  ```

> **Note:** You can check the latest version [here](https://mvnrepository.com/artifact/com.larksuite.oapi/oapi-sdk).

  ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/2748a74e144085f38b3b185186c63abb_DpreW7xUML.png?height=1664&lazyload=true&maxWidth=750&width=3168)

### Node.js
- Installation using npm

  ```shell
  npm install @larksuiteoapi/node-sdk
  ```

- Installation using yarn

  ```shell
  yarn add @larksuiteoapi/node-sdk
  ```

## Related Links

- Server API List
