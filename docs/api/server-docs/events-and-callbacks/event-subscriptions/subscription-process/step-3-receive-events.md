---
title: "Step 3: Receive events"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uYDNxYjL2QTM24iN0EjN/event-subscription-configure-/encrypt-key-encryption-configuration-case"
section: "Events and callbacks"
updated: "1750063787000"
---

# Step 3: Receive events

This article describes how developers can use their own servers to receive events.

## Procedure

If the application does not receive the subscribed events in time, you can log in to the [Developer Console](https://open.larksuite.com/app) and go to the application details page. On the **Log Search > Event Log Search** page, check the log information to confirm whether the Lark open platform has pushed the subscribed events.

![](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/6c64dd780235500dbf9194f01cf662f9_CAgVPct6nM.png?height=1046&lazyload=true&maxWidth=600&width=2514)

## Send notifications to developer's server

You can perform security verification after receiving the events pushed by the Lark open platform. If the event is encrypted, you need to decrypt the event before parsing the event details.

### Security verification

When your local server receives events pushed by the Open Platform (excluding request URL verification), if you need to ensure that the request originates from the Lark Open Platform rather than a forged source, there are two ways to perform security checks: signature verification and Verification Token verification.

| Verification Method | Instructions |
| --- | --- |
| **Signature Verification** | If you have configured the Encrypt Key encryption strategy in your Lark application, signature verification must be used. This method is relatively complex but highly secure, and it completes the security check without needing to decrypt and parse the event. The verification process is as follows: 1. Obtain `encrypt_key`. You can view `encrypt_key` on the **Events & Callbacks > Encryption Strategy** page in the application management platform. ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/27ed366b84787bd940fef14eafc730cd_jeXQnHnNXl.png?height=500&lazyload=true&maxWidth=400&width=1618) 2. Verify the request source, refer to the **Signature verification example code** below for sample code. - Concatenate the request headers `X-Lark-Request-Timestamp`, `X-Lark-Request-Nonce` with `encrypt_key`, encode the concatenated string using `encode('utf-8')` to get `byte[] b1`, and then concatenate with the original request body (the original body defined by the event structure, which varies for different events), resulting in `byte[] b`. - Use the sha256 algorithm to hash `b` and obtain the string `s`. Verify whether `s` matches the request header `X-Lark-Signature`. > This method completes the security check without decrypting the event; however, decryption is still required to obtain the event content. For event decryption operations, refer to **Event decryption** below. |
| **Verification Token Verification** | Lark applications are configured with a default Verification Token. You can receive event requests on your business server, retrieve the Verification Token value from the request body, and compare it with the Verification Token value within your Lark application. Matching values indicate that the request originates from the specified application on the Lark Open Platform. - This verification method is simple but has lower security. Without the Encrypt Key encryption strategy, the Verification Token will be transmitted in plain text, posing a data leakage risk. - The Verification Token can be obtained from the **Event & Callback > Encryption Strategy** page on the application management platform and compared with the Verification Token extracted from the event. ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/69ac49bb94a83ee15150b4cfcf88ee95_6nsboEctql.png?height=472&lazyload=true&maxWidth=400&width=1506) > If the application has configured the Encrypt Key encryption strategy, it is recommended to use signature verification. If you still need to retrieve the Verification Token, you must decrypt the event first to obtain it. For event decryption operations, refer to **Event decryption** below. | #### Signature verification example code

Below are example codes in Python 3, Java, Golang, Node.js, C#, and PHP, which include the following parameters:
- `timestamp`: Corresponding to `X-Lark-Request-Timestamp` in the request header.
- `nonce`: Corresponding to `X-Lark-Request-Nonce` in the request header.
- `encrypt_key`: The Encrypt Key obtained from the encryption strategy in the developer console application.

**Python 3**

```python
import hashlib
bytes_b1 = (timestamp + nonce + encrypt_key).encode('utf-8')
bytes_b = bytes_b1 + body
h = hashlib.sha256(bytes_b)
signature = h.hexdigest()

# check if request headers['X-Lark-Signature'] equals to signature
```

**Java**

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

**Golang**

```go
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

**Node.js**

```js
var crypto = require('crypto');
function calculateSignature(timestamp, nonce, encryptKey, body) {
        const content = timestamp + nonce + encryptKey + body
        const sign = crypto.createHash('sha256').update(content).digest('hex');
        return sign
}
```

**C#**

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

**PHP**

```php
 **Note:** Please execute `pip install pycryptodome` first to support the introduction of the AES method.

```python
import hashlib
import base64
from Crypto.Cipher import AES
class  AESCipher(object):
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
        return  self._unpad(cipher.decrypt(enc[AES.block_size:]))
    def decrypt_string(self, enc):
        enc = base64.b64decode(enc)
        return  self.decrypt(enc).decode('utf8')
encrypt = "P37w+VZImNgPEO1RBhJ6RtKl7n6zymIbEG1pReEzghk="
cipher = AESCipher("test key")
print("明文:\n{}".format(cipher.decrypt_string(encrypt)))
```

**Java**

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
        if (r.length > 0) {
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

**Golang**

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

**Node.js**

```js
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

**C#**

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

```
<?php
$encrypt_data = ""; // The information to be decrypted
$encrypt_key = ""; // Encrypt Key from the developer console
$base64_decode_message = base64_decode($encrypt_data);
$iv = substr($base64_decode_message, 0, 16);
$encrypted_event = substr($base64_decode_message, 16);
$decrypt = openssl_decrypt($encrypted_event, 'AES-256-CBC', hash('sha256', $encrypt_key, true), OPENSSL_RAW_DATA, $iv);
print($decrypt); 
// get the real event
```
