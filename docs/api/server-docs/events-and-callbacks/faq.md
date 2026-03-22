---
title: "FAQ"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/event-subscription-guide/event-subscriptions/faq"
section: "Events and callbacks"
updated: "1718768781000"
---

# FAQ

This article summarizes the frequently asked questions related to event subscription.

## Q: What should I do if an error occurs when configuring the event subscription request address and the request cannot be saved?

You can follow the steps below to check if there is a problem with the configuration.

1. The request address must be a request address that is accessible to the public network using the IPv4 protocol (must start with `http://` or `https://`).

	For example, if the request address is generated using an intranet penetration tool, the development platform will not be able to connect to the request address you configured because the tool is unstable or the local enterprise has set up a firewall.

2. The server where the request address is located can successfully receive the HTTP POST request sent by the open platform to verify the legitimacy of the request address, and must return the received challenge value in JSON format to the open platform within 1 second.

	If the application is configured with an Encrypt Key, it needs to be decrypted first, and then the challenge value must be extracted from the decrypted content, and the value must be returned as a response within 1 second. You can refer to the steps in the **Send Events to Developer Server** section of this article to check if there is a problem with the configuration.

You can use the following curl command line tool to test whether the configured request address can successfully return the challenge value.

```bash
# Note that when using the following command to test, please make sure that the application is not configured with an encryption key. If the application is configured with an encryption key, the HTTP request packet will be encrypted.

curl -v --location '{YOUR_CALLBACK_URL}' \
--header 'Content-Type: application/json' \
--data '{
"challenge": "",
"type": "url_verification",
"token": ""
}'
```

Where:

- `{YOUR_CALLBACK_URL}`: Replace with the request address you configured, starting with `http://` or `https://`.

- ``: Replace with the challenge code used for testing, such as `ajls384kdj1234`.
- ``: Replace with the **Verification Token** value in the [Developer Console](https://open.larksuite.com/app) > **Application Details Page** > **Events & Callbacks** > **Encryption Strategy** page.

The successful return result is shown in the figure below.

![](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/b2ba36fd3588c069867a8ac5b62144f9_jaLkd92ft2.png?height=732&lazyload=true&maxWidth=600&width=2780)

## Q: When configuring the event subscription request address, the prompt "Please enter valid HTTP/HTTPS URL" is displayed. How to deal with it?

**Problem phenomenon**

As shown in the figure below, after configuring the event subscription request address, the prompt "Please enter valid HTTP/HTTPS URL" is displayed when saving.

![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/a5200fc63b3370cde1b1951e48a87e26_XHHornu7Tf.png?height=1018&lazyload=true&maxWidth=600&width=1754)

**Troubleshooting suggestions**

> **Note:** When you configure the server for receiving event subscriptions locally, it is recommended to print out the HTTP POST request from the open platform directly to confirm whether the server has received the request and what the specific content of the request contains.

1. First use the following curl command line tool to test whether the configured request address can successfully return the challenge value.

      ```bash
      # Note that when using the following command to test, please confirm that the application is not configured with an Encrypt Key. If the application is configured with an Encrypt Key, the HTTP request packet will be encrypted.

      curl -v --location '{YOUR_CALLBACK_URL}' \
      --header 'Content-Type: application/json' \
      --data '{
      "challenge": "",
      "type": "url_verification",
      "token": ""
      }'
      ```
      Where:

      - `{YOUR_CALLBACK_URL}`: Replace with the request address you configured, starting with `http://` or `https://`.

      - ``: Replace with the challenge code used for testing, such as `ajls384kdj1234`.

      - ``: Replace with the **Verification Token** value in the [Developer Console](https://open.larksuite.com/app) > **Application Details Page** > **Events & Callbacks** > **Encryption Strategy** page.

     
      The successful return result is shown in the figure below.

  	![](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/b2ba36fd3588c069867a8ac5b62144f9_jaLkd92ft2.png?height=732&lazyload=true&maxWidth=600&width=2780)

2. If the test command can return the same challenge code as in the request body, then:

	- Check whether the request from the open platform has accessed the callback server corresponding to the request address. If there is no record of accessing the callback server, please check whether the callback server is inaccessible due to the configuration of the firewall policy.
	- Check whether the configured request address is accessible from the public network. For example, you can use another external network device to re-execute the curl command in step 1 to test whether the challenge value can be returned normally.

3. If the test command cannot return the same challenge code as in the request body, please refer to the Configure Subscription Method document first, write the correct code, and then try again.

4. If the application is configured with Encrypt Key, the open platform will push the encrypted POST request. In the callback server corresponding to the request address, you need to follow the Event Decryption document to decrypt and then return the challenge code in the message body.

      ```json
      # When the application is configured with Encrypt Key, the challenge code initiated by the open platform is encrypted.
      # Developers need to decrypt according to the event decryption document.

      {
          "encrypt": "ds3da3sj32421lkkld4xxxx" // Encrypted string
      }
      ```
