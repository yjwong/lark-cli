---
title: "Step 3: Receive callbacks"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/event-subscription-guide/callback-subscription/receive-and-handle-callbacks"
section: "Events and callbacks"
updated: "1744337108000"
---

# Step 3: Receive callbacks

After configuring the callback subscription request URL, upon receiving the callback pushed by the Lark open platform, you can perform a security verification. If the callback is encrypted, you will need to decrypt the callback first and then parse the callback details.

## Process overview

When the Lark server pushes the callback to the request address, the corresponding business server needs to receive the callback request, complete the business processing and return the response result within 3 seconds. During this process, you need to complete the following operations based on the actual configuration.

> **Note:** Callback is a synchronous operation and does not provide a pushback mechanism. If your business server times out and does not respond, the system will determine that the callback failed and display an error message in the Lark client.

| Operation | Required | Description |
| --- | --- | --- |
| Security check | No | The security check is used to confirm that the request received by the business server comes from the Lark Open Platform and is not a forged risky request. |
| Callback decryption | No | It is recommended to configure Encrypt Key for the application. After configuration, the callback request pushed is encrypted data, which can ensure the security of the requested data. Accordingly, after receiving the request, the business server needs to decrypt it before it can obtain the real callback data. |
| Response callback request | Yes | After receiving the callback request, the business server must return the response result, otherwise the system will judge that the callback failed. | ## Security check

When your business server receives a callback pushed from Lark Open Platform, you can ensure that the request comes from Lark Open Platform in the following ways.

> **Warning:** The security checks provided in this article do not apply to the **Message card callback communication (legacy)** (card.action.trigger_v1) callback. If you need to configure security verification for **Message card callback communication (legacy)**, you need to refer to Configuring callback request address.

### Method 1: Verification Token verification

The Lark application is configured with Verification Token by default. You can receive a callback request in the business server and obtain the Verification Token value in the request body. Compare this value with the Verification Token value in the Lark application. If the values are the same, it means The request comes from a designated application on the Lark open platform.

> **Note:** - This method is relatively simple to operate, but less safe. For example, transmitting the Verification Token value in clear text risks data leakage.
> - If the application is configured with an Encrypt Key encryption policy, you need to decrypt it in the business server before you can obtain the Verification Token value. For the decryption operation, see **Perform callback decryption** below.

### Method 2: Signature verification

If you configure the Encrypt Key encryption policy in the Lark application, after the business server receives the callback request, it does not need to perform a decryption operation and can directly complete the signature verification.

> **Note:** This method is relatively complex to operate, but it is highly safe.

**Verification process**

1. Enter the designated application details page of [Developer Console](https://open.larksuite.com/app).
2. Select **Development Configuration** > **Events and Callbacks** in the left navigation bar, and select **Encryption Strategy**.
3. On the **Encryption Strategy** tab, obtain the configured **Encrypt Key**.
    
     ![](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/1ef66d46517e46554aed864a49233a57_qBkR4em3PK.png?height=696&lazyload=true&maxWidth=600&width=1742)

4. Verify the source of the request in the local business server.
    
     1. After splicing the request headers `X-Lark-Request-Timestamp`, `X-Lark-Request-Nonce` and `encrypt_key`, encode it according to `encode('utf-8')` to get `byte[] b1` , and then splice the original body of the request to get a `byte[] b`.
     2. Use sha256 algorithm on `b` to get the string `s`, and verify whether `s` is consistent with the request header `X-Lark-Signature`.

**Sample code**

The following provides sample code for signature verification in each development language. The parameter descriptions included are shown in the table below.
| **Parameters** | **Description** |
| ------------- | ---------------------------------- |
| `timestamp` | Corresponds to `X-Lark-Request-Timestamp` in the request header |
| `nonce` | Corresponds to `X-Lark-Request-Nonce` in the request header |
| `encrypt_key` | The application Encrypt Key obtained from the developer backend | - **Python 3**
    
```python
import hashlib
bytes_b1 = (timestamp + nonce + encrypt_key).encode('utf-8')
bytes_b = bytes_b1 + body
h = hashlib.sha256(bytes_b)
signature = h.hexdigest()

# check if request headers['X-Lark-Signature'] equals to signature
```

- **Java**

```java
import org.apache.commons.codec.binary.Hex;
import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;
class Main {
   public String calculateSignature(String timestamp, String nonce, String encryptKey, String bodyString) throws NoSuchAlgorithmException {
StringBuilder content = new StringBuilder();
content.append(timestamp).append(nonce).append(encryptKey).append(bodyString);
MessageDigest alg = MessageDigest.getInstance("SHA-256");
String sign = Hex.encodeHexString(alg.digest(content.toString().getBytes()));
return sign;
   }
}
```

- **Golang**

```
import (
    "crypto/sha256"
    "fmt"
)
func calculateSignature(timestamp, nonce, encryptKey, bodystring string) string {
    var b strings.Builder
    b.WriteString(timestamp)
    b.WriteString(nonce)
    b.WriteString(encryptKey)
    b.WriteString(bodystring) //bodystring refers to the entire request body, do not calculate it after deserialization
    bs := []byte(b.String())
    h := sha256.New()
    h.Write(bs)
    bs = h.Sum(nil)
    sig := fmt.Sprintf("%x", bs)
    return sig
}
```

- **node.js**

```javascript
var crypto = require('crypto');
function calculateSignature(timestamp, nonce, encryptKey, body) {
const content = timestamp + nonce + encryptKey + body
const sign = crypto.createHash('sha256').update(content).digest('hex');
return sign
}
```

- **C#**

```c#
using System.Security.Cryptography;
public static string calculateSignature(string timestamp, string nonce, string encryptKey, string body) {
StringBuilder content = new StringBuilder();
content.Append(timestamp);
content.Append(nonce);
content.Append(encryptKey);
content.Append(body);
SHA256 sha256 = new SHA256CryptoServiceProvider();
byte[] bytes_out = sha256.ComputeHash(Encoding.UTF8.GetBytes(content.ToString()));
string result = BitConverter.ToString(bytes_out);
result = result.Replace("-", "");
return result;
}
```

- **PHP**

```
 **Warning:** The callback decryption operation does not apply to the **Message card callback communication (legacy)** (card.action.trigger_v1) callback, so you can ignore the operations in this chapter when using the **Message card callback communication (legacy)** callback.

### Decrypt sample code

You can refer to the following sample codes in each development language to decrypt the callback data.

- **Python 3**

> **Note:** Please execute `pip install pycryptodome` first to support the introduction of AES method.

```python
import hashlib
import base64
from Crypto.Cipher import AES
class AESCipher(object):
     def __init__(self, key):
         self.bs = AES.block_size
         self.key=hashlib.sha256(AESCipher.str_to_bytes(key)).digest()
     @staticmethod
     def str_to_bytes(data):
         u_type = type(b"".decode('utf8'))
         if isinstance(data, u_type):
             return data.encode('utf8')
         return data
     @staticmethod
     def _unpad(s):
         return s[:-ord(s[len(s) - 1:])]
     def decrypt(self, enc):
         iv = enc[:AES.block_size]
         cipher = AES.new(self.key, AES.MODE_CBC, iv)
         return self._unpad(cipher.decrypt(enc[AES.block_size:]))
     def decrypt_string(self, enc):
         enc = base64.b64decode(enc)
         return self.decrypt(enc).decode('utf8')
encrypt = "P37w+VZImNgPEO1RBhJ6RtKl7n6zymIbEG1pReEzghk="
cipher = AESCipher("test key")
print("plaintext:\n{}".format(cipher.decrypt_string(encrypt)))
```

- **Java**

```java
package com.larksuite.oapi.sample;
import javax.crypto.Cipher;
import javax.crypto.spec.IvParameterSpec;
import javax.crypto.spec.SecretKeySpec;
import java.nio.charset.StandardCharsets;
import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;
import java.util.Base64;
public class Decrypt {
     public static void main(String[] args) throws Exception {
         Decrypt d = new Decrypt("test key");
         System.out.println(d.decrypt("P37w+VZImNgPEO1RBhJ6RtKl7n6zymIbEG1pReEzghk=")); //hello world
     }
     private byte[] keyBs;
     public Decrypt(String key) {
         MessageDigest digest = null;
         try {
             digest = MessageDigest.getInstance("SHA-256");
         } catch (NoSuchAlgorithmException e) {
             // won't happen
         }
         keyBs = digest.digest(key.getBytes(StandardCharsets.UTF_8));
     }
     public String decrypt(String base64) throws Exception {
         byte[] decode = Base64.getDecoder().decode(base64);
         Cipher cipher = Cipher.getInstance("AES/CBC/NOPADDING");
         byte[] iv = new byte[16];
         System.arraycopy(decode, 0, iv, 0, 16);
         byte[] data = new byte[decode.length - 16];
         System.arraycopy(decode, 16, data, 0, data.length);
         cipher.init(Cipher.DECRYPT_MODE, new SecretKeySpec(keyBs, "AES"), new IvParameterSpec(iv));
         byte[] r = cipher.doFinal(data);
         if (r. length > 0) {
             int p = r.length - 1;
             for (; p >= 0 && r[p] <= 16; p--) {
             }
             if (p != r.length - 1) {
                 byte[] rr = new byte[p + 1];
                 System.arraycopy(r, 0, rr, 0, p + 1);
                 r = rr;
             }
         }
         return new String(r, StandardCharsets.UTF_8);
     }
}
```

- **Golang**

```go
package main
import (
        "crypto/aes"
        "crypto/cipher"
        "crypto/sha256"
        "encoding/base64"
        "errors"
        "fmt"
        "strings"
)
func main() {
        s, err := Decrypt("P37w+VZImNgPEO1RBhJ6RtKl7n6zymIbEG1pReEzghk=", "test key")
        if err != nil {
                panic(err)
        }
        fmt.Println(s) //hello world
}
func Decrypt(encrypt string, key string) (string, error) {
        buf, err := base64.StdEncoding.DecodeString(encrypt)
        if err != nil {
                return "", fmt.Errorf("base64StdEncode Error[%v]", err)
        }
        if len(buf) < aes.BlockSize {
                return "", errors.New("cipher  too short")
        }
        keyBs := sha256.Sum256([]byte(key))
        block, err := aes.NewCipher(keyBs[:sha256.Size])
        if err != nil {
                return "", fmt.Errorf("AESNewCipher Error[%v]", err)
        }
        iv := buf[:aes.BlockSize]
        buf = buf[aes.BlockSize:]
        // CBC mode always works in whole blocks.
        if len(buf)%aes.BlockSize != 0 {
                return "", errors.New("ciphertext is not a multiple of the block size")
        }
        mode := cipher.NewCBCDecrypter(block, iv)
        mode.CryptBlocks(buf, buf)
        n := strings.Index(string(buf), "{")
        if n == -1 {
                n = 0
        }
        m := strings.LastIndex(string(buf), "}")
        if m == -1 {
                m = len(buf) - 1
        }
        return string(buf[n : m+1]), nil
}
```

- **Node.js**

```javascript
const crypto = require("crypto");
class AESCipher {
    constructor(key) {
        const hash = crypto.createHash('sha256');
        hash.update(key);
        this.key = hash.digest();
    }
    decrypt(encrypt) {
        const encryptBuffer = Buffer.from(encrypt, 'base64');
        const decipher = crypto.createDecipheriv('aes-256-cbc', this.key, encryptBuffer.slice(0, 16));
        let decrypted = decipher.update(encryptBuffer.slice(16).toString('hex'), 'hex', 'utf8');
        decrypted += decipher.final('utf8');
        return decrypted;
    }
}
encrypt = "P37w+VZImNgPEO1RBhJ6RtKl7n6zymIbEG1pReEzghk="
cipher = new AESCipher("test key")
console.log(cipher.decrypt(encrypt))
// hello world
```

- **C#**

```c#
using System;
using System.Linq;
using System.Security.Cryptography;
using System.Text;
namespace decrypt
{
        class AESCipher
        {
                const int BlockSize = 16;
                private byte[] key;
                public AESCipher(string key)
                {
                        this.key = SHA256Hash(key);
                }
                public string DecryptString(string enc)
                {
                        byte[] encBytes = Convert.FromBase64String(enc);
                        RijndaelManaged rijndaelManaged = new RijndaelManaged();
                        rijndaelManaged.Key = this.key;
                        rijndaelManaged.Mode = CipherMode.CBC;
                        rijndaelManaged.IV = encBytes.Take(BlockSize).ToArray();
                        ICryptoTransform transform = rijndaelManaged.CreateDecryptor();
                        byte[] blockBytes = transform.TransformFinalBlock(encBytes, BlockSize, encBytes.Length - BlockSize);
                        return System.Text.Encoding.UTF8.GetString(blockBytes);
                }
                public static byte[] SHA256Hash(string str)
                {
                        byte[] bytes = Encoding.UTF8.GetBytes(str);
                        SHA256 shaManaged = new SHA256Managed();
                        return shaManaged.ComputeHash(bytes);
                }
                public static void Main(string[] args)
                {
                        string encrypt = "P37w+VZImNgPEO1RBhJ6RtKl7n6zymIbEG1pReEzghk=";
                        AESCipher cipher = new AESCipher("test key");
                        Console.WriteLine(cipher.DecryptString(encrypt));
                }
        }
}
```

**PHP**

```php
<?php
$encrypt_data = ""; // Information to be decrypted
$encrypt_key = ""; // Get the Encrypt Key from the developer backend
$base64_decode_message = base64_decode($encrypt_data);
$iv = substr($base64_decode_message, 0, 16);
$encrypted_event = substr($base64_decode_message, 16);
$decrypt = openssl_decrypt($encrypted_event, 'AES-256-CBC', hash('sha256', $encrypt_key, true), OPENSSL_RAW_DATA, $iv);
print($decrypt);
// get the real event
```

## Respond to callback requests

After your business server receives the callback request, it needs to respond to the callback within 3 seconds to complete the interactive behavior of the Lark client (front-end). Currently, please refer to the table below for the functions that need to be subscribed to callbacks and the corresponding callback structure and usage.

| Function | Callback structure | Related documentation |
| --- | --- | --- |
| Link preview | To implement the link preview function, you must subscribe to the **Pull link preview data** callback. For the callback parameters and response parameter descriptions corresponding to this callback, please refer to Pull link preview data. | To understand the link preview feature and how to configure link preview, refer to the Link preview development guide. |
