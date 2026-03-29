---
title: "Send callbacks to developer's server"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/event-subscription-guide/callback-subscription/step-1-choose-a-subscription-mode/send-callbacks-to-developers-server"
section: "Events and callbacks"
updated: "1773297872000"
---

# Send callbacks to developer's server

You can set up your own server to receive callback subscription notifications from the Lark Open Platform via the webhook method. This requires you to provide a public access address for your server. When the callback is triggered, the Lark Open Platform will send an HTTP POST request to this public address, containing the callback data.

## Step 1 (Optional): Configure encryption strategy

You can configure an encryption strategy according to your enterprise's data security requirements. This strategy encrypts the transmission of callback data and verifies the request source to assess risks. The encryption strategy includes two parameters: **Encrypt Key** and **Verification Token**. **Encrypt Key** is used for encrypting the callback request transmission, while **Verification Token** verifies if the request is from the Lark Open Platform.

### Encrypt Key

This parameter is used to encrypt the transmission of callback subscription requests.

- If you configure the Encrypt Key, the Lark Open Platform will encrypt the callback data before pushing it to your request address. This encryption enhances data security, so it is recommended to configure this parameter.
- If you do not configure the Encrypt Key, the Lark Open Platform will push the callback data directly in plaintext.

The following is an example of an encrypted callback. When your local server receives the encrypted callback body, it needs to decrypt it. For decryption steps, refer to Receive callbacks.

```
{
    "encrypt": "FIAfJPGRmFZWkaxPQ1XrJZVbv2JwdjfLk4jx0k/U1deAqYK3AXOZ5zcHt/cC4ZNTqYwWUW/EoL+b2hW/C4zoAQQ5CeMtbxX2zHjm+E4nX/Aww+FHUL6iuIMaeL2KLxqdtbHRC50vgC2YI7xohnb3KuCNBMUzLiPeNIpVdnYaeteCmSaESb+AZpJB9PExzTpRDzCRv+T6o5vlzaE8UgIneC1sYu85BnPBEMTSuj1ZZzfdQi7ZW992Z4dmJxn9e8FL2VArNm99f5Io3c2O4AcNsQENNKtfAAxVjCqc3mg5jF0123123flX1UOF5fzJ0sApG2OEn9wfyPDRBsApn9o+fceF9hNrYBGsdtZrZYyGG387CGOtKsuj8e2E8SNp+Pn4E9oYejOTR+ZNLNi+twxaXVlJhr6l+RXYwEiMGQE9zGFBD6h2dOhKh3W84p1GEYnSRIz1+9/Hp66arjC7RCrhuW5OjCj4QFEQJiwgL123123VsL03CxxFZarzxzffryrWUG3VkRdHRHbTsC34+ScoL5MTDU1QAWdqUC1T7xT0lCvQELaIhBTXAYrznJl6PlA83oqlMxpHh0gZBB1jFbfoUr7OQbBs1xqzpYK6Yjux6diwpQB1zlZErYJUfCqK7G/zI9yK/60b4HW0123+123123=" 
}
```

#### Applicable scenarios

- The application verifies that the received callback push is from the Lark Open Platform and not forged.
- The application prevents replay attacks, where the callback pushed by the Lark Open Platform to the application is intercepted by a third party and then sent multiple times to the application. This can pose data security risks and affect the performance of the application's server.

After configuring the Encrypt Key, your local server can perform signature verification based on the Encrypt Key to ensure that the received callbacks are legitimate and sent by the Lark Open Platform. For detailed information on signature verification, refer to Receive callbacks.

#### Encryption principle

The content of the callback request is encrypted using AES-256-CBC. The encryption principle is as follows:
1. Use SHA256 to hash the `Encrypt Key` to get the key `key`.
1. Use PKCS7Padding to pad the callback content.
1. Generate a 16-byte random number as the initial vector `iv`.
1. Use `iv` and `key` to encrypt the callback content to get `encrypted_event`.
1. The ciphertext received by the application is `base64(iv+encrypted_event)`.

### Verification Token

The Verification Token is the application's verification identifier. The developer backend will automatically generate a Verification Token for the application. When the Lark Open Platform pushes callback data, it will include the Verification Token value, which you can use to verify if the pushed callback belongs to the current application.

### Steps to configure

1. Log in to the [Developer Console](https://open.larksuite.com/app).
2. Click on the specific application in the application list to enter the application management details page.
3. In the left navigation bar, select **Development Configuration** > **Events & Callbacks**.
4. In the **Encryption Strategy** tab, configure the Encrypt Key or Verification Token.
    

	![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/ead1a4405163fdf4b1c2fd8d42c0a628_0g0u3AMDPM.png?height=678&lazyload=true&maxWidth=600&width=1684)
    
    - Click the **Reset** icon to reset the parameter value.
        
		![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/8ab5f99e9e57edeef43d9f40b2b572a4_saxyEsC49N.png?height=338&lazyload=true&maxWidth=400&width=832)

    
    - Click the **Edit** icon to edit the parameter value.
        
		![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/b47621e75faa436ff83f499cb0b97db8_yBDdUWN0YY.png?height=344&lazyload=true&maxWidth=400&width=840)

        
## Step 2: Configure the callback subscription method

1. Log in to the [Developer Console](https://open.larksuite.com/app).
2. Enter the specified application and select **Development Configuration** > **Events & Callbacks** in the left navigation bar on the application details page.
3. In the **Events & Callbacks** page, click the **Callback Configuration** tab.

4. Click the edit button on the right side of **Subscription mode**.

    ![](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/4ea3abe6eb6e39ed5a02016e6390eba8_qSLpe1F7H1.png?height=590&lazyload=true&maxWidth=600&width=1734)

5. Select **Send callbacks to developer's server**, configure the **Request URL**, and click **Save**.

    :::note
    Each application can only configure one callback request address, which must be an IPv4 public address. All subsequent callbacks subscribed by the application will send callback data to this request address when triggered.
    :::

	![](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/d879e4149154a750a4a232799778298a_BlQQpHw143.png?height=542&lazyload=true&maxWidth=600&width=1626)

	When you fill in the address and save it, the Lark server will send an HTTP POST request to the request address to verify its validity. The request format is JSON, containing the `challenge` parameter. The developer's server needs to receive this HTTP POST request and return a response containing the `challenge` parameter within 1 second to successfully complete the verification. For details on how to handle and respond to the HTTP POST request, refer to the **Respond to POST validation request** section below.
    
    
### Responding to POST verification requests

Depending on whether the Lark application has configured an Encrypt Key, there are two ways to respond to POST verification requests.

#### Method 1: Application without Encrypt Key

The verification request body sent by the Lark server is as follows:

```json
{ 
    "challenge": "1b6aef1a-401f-406a-be41-f48911eabcef",    // The value needs to be returned as it is in the response.
    "token": "xxxxxx",    // The application's Verification Token, which you can use to determine if the request is from the specified application of Lark's open platform.
    "type": "url_verification"    // Fixed value, indicating this is a verification request.
}
```

The parameter descriptions are as follows:

| Parameter | Type | Description |
| --- | --- | --- |
| challenge | String | The field for verification, which needs to be returned as it is in the response body. Example value: 1b6aef1a-401f-406a-be41-f48911eabcef |
| token | String | Application verification identifier **Verification Token**. You can use this token to verify whether the request belongs to the current application. In the **Developer Console** > **Application Details** > **Development Configuration** > **Events & Callbacks** > **Encryption Strategy** module, you can view the application's **Verification Token**. |
| type | String | Callback type. The fixed value for verification callback type is `url_verification`, indicating this request is for verifying URL legitimacy. | When the request address receives this POST verification request, you need to extract the `challenge` value and return the response data containing the `challenge` value within 1 second. The response body example is as follows:

```json
{ 
    "challenge": "1b6aef1a-401f-406a-be41-f48911eabcef"
}
```

#### Method 2: Application with Encrypt Key

If the application has configured an Encrypt Key, encrypted requests will be pushed. The request body is as follows:

```json
{
    "encrypt": "ds3da3sj32421lkkld4xxxx" // Encrypted string.
}
```

You need to decrypt the `encrypt` string on the server that receives the request and return the extracted `challenge` value within 1 second. The decryption method can be found in Receiving callbacks.

- The decrypted data structure is as follows:

    ```json
    { 
        "challenge": "1b6aef1a-401f-406a-be41-f48911eabcef",    // The value that needs to be returned as it is in the response.
        "token": "xxxxxx",    // The application's Verification Token, which you can use to determine if the request is from the specified application of Lark's open platform.
        "type": "url_verification"    // Fixed value, indicating this is a verification request.
    }
    ```

- The response body example is as follows:

    ```json
    { 
        "challenge": "1b6aef1a-401f-406a-be41-f48911eabcef"
    }
    ```
