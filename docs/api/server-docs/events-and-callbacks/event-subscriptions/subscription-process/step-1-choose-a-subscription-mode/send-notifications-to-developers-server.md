---
title: "Send notifications to developer's server"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uYDNxYjL2QTM24iN0EjN/event-subscription-configure-/choose-a-subscription-mode/send-notifications-to-developers-server"
section: "Events and callbacks"
updated: "1773297865000"
---

# Send notifications to developer's server

You can build your own server to receive event subscription notifications from the Lark Open Platform via the Webhook mode. This mode requires you to provide a public access address for your server. When an event is triggered, the Lark Open Platform will send an HTTP POST request to the public address containing the event data.

## Step 1 (Optional): Configure encryption strategy

According to your enterprise data security requirements, you can configure an encryption strategy to encrypt transmitted event data and verify whether the request source is at risk. The encryption strategy includes two parameters: **Encrypt Key** and **Verification Token**. The **Encrypt Key** is used to encrypt the transmission of event requests, and the **Verification Token** is used to verify whether the request originates from the Lark Open Platform.

### Encrypt Key

This parameter is used to encrypt the transmission of event subscription requests.

- If you configure the Encrypt Key, the Lark Open Platform will encrypt the event data when pushing it to the request address. Encrypted push can enhance data security, so it's recommended to configure this parameter.
- If you do not configure the Encrypt Key, the Lark Open Platform will push the event data in plain text without encryption.

An example of an encrypted push event is as follows. After your local server receives the encrypted event body, it needs to perform decryption. Refer to Receiving Events for the decryption steps.

```json
{
    "encrypt": "FIAfJPGRmFZWkaxPQ1XrJZVbv2JwdjfLk4jx0k/U1deAqYK3AXOZ5zcHt/cC4ZNTqYwWUW/EoL+b2hW/C4zoAQQ5CeMtbxX2zHjm+E4nX/Aww+FHUL6iuIMaeL2KLxqdtbHRC50vgC2YI7xohnb3KuCNBMUzLiPeNIpVdnYaeteCmSaESb+AZpJB9PExzTpRDzCRv+T6o5vlzaE8UgIneC1sYu85BnPBEMTSuj1ZZzfdQi7ZW992Z4dmJxn9e8FL2VArNm99f5Io3c2O4AcNsQENNKtfAAxVjCqc3mg5jF0123123flX1UOF5fzJ0sApG2OEn9wfyPDRBsApn9o+fceF9hNrYBGsdtZrZYyGG387CGOtKsuj8e2E8SNp+Pn4E9oYejOTR+ZNLNi+twxaXVlJhr6l+RXYwEiMGQE9zGFBD6h2dOhKh3W84p1GEYnSRIz1+9/Hp66arjC7RCrhuW5OjCj4QFEQJiwgL123123VsL03CxxFZarzxzffryrWUG3VkRdHRHbTsC34+ScoL5MTDU1QAWdqUC1T7xT0lCvQELaIhBTXAYrznJl6PlA83oqlMxpHh0gZBB1jFbfoUr7OQbBs1xqzpYK6Yjux6diwpQB1zlZErYJUfCqK7G/zI9yK/60b4HW0123+123123=" 
}
```

#### Applicable scenarios

- The application verifies that the received event push is from the Lark Open Platform rather than forged.
- The application prevents replay attacks. Replay attacks refer to a third party intercepting event data pushed by the Lark Open Platform to the application, then resending the event multiple times in its original form to the application. This may cause data security concerns and impact the performance of the application server.

After configuring the Encrypt Key, your local server can perform signature verification based on the Encrypt Key to ensure that the application receives only legitimate events pushed by the Lark Open Platform. For detailed information on signature verification, refer to Receiving Events.

#### Encryption principle

The event content is encrypted using AES-256-CBC. The encryption principle is as follows:

1. Hash the `Encrypt Key` with SHA256 to get the key `key`.
2. Pad the event content using the PKCS7Padding method.
3. Generate a 16-byte random number as the initial vector `iv`.
4. Encrypt the event content using `iv` and `key` to get the `encrypted_event`.
5. The ciphertext received by the application is `base64(iv+encrypted_event)`.

### Verification Token

The Verification Token is an authentication identifier for the application. The developer backend will automatically generate a Verification Token for the application. When the Lark Open Platform pushes event data, it will carry the Verification Token, allowing you to verify whether the pushed event belongs to the current application.

### Operation steps

1. Log in to the [Developer Console](https://open.larksuite.com/app).
2. In the application list, click on a specific application to enter the application management details page.
3. In the left navigation bar, select **Development Configuration** > **Event & Callbacks**.
4. Enter the **Encryption Strategy** tab to configure the Encrypt Key or Verification Token.

	![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/a48b2fe3bab8e2d2188b7a54a92af016_WXmDxUhWob.png?height=426&lazyload=true&maxWidth=500&width=1190)

    
    - Click the **Reset** icon to reset the parameter values.
        

		![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/df1fbfb68b3768c3c30f6844d0bb1492_3qjhKxB7Pi.png?height=318&lazyload=true&maxWidth=400&width=798)
    
    - Click the **Edit** icon to edit the parameter values.
    
		![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/1e580a2daa8535a5b3e97bc1a67cbebb_2eNFGFUTWF.png?height=312&lazyload=true&maxWidth=400&width=794)    

        
## Step 2: Set up subscription method

1. Log in to the [Developer Console](https://open.larksuite.com/app).
2. Click on a specific application in the application list to enter the application management details page.
3. In the left navigation bar, select **Development Configuration** > **Events & Callbacks**.
4. Go to the **Event Configuration** tab, and click the edit icon in the **Subscription mode** area.

    ![](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/9f08f8a5dd68573cc8c8d1cc3bb675fc_SYWyqSSsX4.png?height=522&lazyload=true&maxWidth=600&width=1674)

5. Select **Send notifications to developer's server**, enter the request URL, and click **Save**.
    
    :::warning
    - Each application can only configure one request URL, and all subscribed events of this application will be sent to this URL.
    - The request URL must be a public IPv4 address.
    :::

    ![](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/a721337266e3edc32584873d3c40344e_2CQD53uvRk.png?height=888&lazyload=true&maxWidth=600&width=1732)

    After clicking save, the Lark server will send a POST request in `application/json;charset=utf-8` format to the request URL to verify the validity of the configured URL. This POST request will contain a `challenge` field, which the application needs to return as is to the Lark open platform within 1 second. Otherwise, the platform will report a "Challenge code not returned" error. For detailed information, please refer to the **Response handling example** below.

    :::note
    When configuring the server to receive event subscriptions locally, it is recommended to directly print out the HTTP POST requests from the open platform to confirm whether the server has received the requests and what specific content is included in the requests.
    :::

### Response handling example

Based on the configuration of the **Encrypt Key** encryption strategy parameter, the processed POST requests will differ.

- **No Encrypt Key Configured**

    If no Encrypt Key is set, the Lark open platform will send a plain text POST request. The request body example is as follows:

    ```json
    { 
        "challenge": "ajls384kdjxxxx",  // The value that the application needs to return as is in the response 
        "token": "xxxxxx",              // Verification Token
        "type": "url_verification"      // Indicates that this is a verification request 
    }
    ```

    When the request URL receives the POST verification request from the open platform, it needs to parse out the `challenge` value and return it as is within 1 second as a response.

    ```json
    { 
        "challenge": "ajls384kdjxxxx"
    }
    ```

- **Encrypt Key Configured**
  
    If an Encrypt Key is configured, the Lark open platform will send an encrypted POST request.

    ```json
    {
        "encrypt": "ds3da3sj32421lkkld4xxxx" // Encrypted string
    }
    ```

    The application needs to decrypt it first, then extract the `challenge` value from the decrypted content, and return it as is within 1 second as a response. For detailed decryption methods, please refer to [Receiving Events](https://ssl:ttdoc/ukTMukTMukTM/uYDNxYjL2QTM24iN0EjN/event-subscription-configure-/encrypt-key-encryption-configuration-case). The decrypted POST request example is as follows:

    ```json
    { 
        "challenge": "ajls384kdjxxxx",  // The value that the application needs to return as is in the response 
        "token": "xxxxxx",              // Verification Token
        "type": "url_verification"      // Indicates that this is a verification request 
    }
    ```

    Response example:

    ```json
    { 
        "challenge": "ajls384kdjxxxx"
    }
    ```
