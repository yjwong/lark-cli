---
title: "Quick approval callback"
url: "https://open.larksuite.com/document/ukTMukTMukTM/ukjNyYjL5YjM24SO2IjN/quick-approval-callback"
section: "Approval"
updated: "1675167418000"
---

# Quick approval callback

When the approver performs [Approve] and [Reject] actions on the approval tasks in the [Pending] list, the Approval Center will call the callback URL configured in the approval definition to notify the third-party system.

After receiving the callback, the third-party system transfers the process and synchronizes the latest task information back to the Approval Center via the approval instance sync API.

The callback timeout is 5 sec, and if it isn't returned after 5 sec, the card will report an error and won't be updated.

### API definition

```json
# Unencrypted   This code is just for example. Enterprises must enter their corresponding enterprise information.
curl --location --request POST 'https://www.larksuite.com/approval/openapi/v2/external/instanceOperate' \
--header 'Content-Type: application/json' \
--data-raw '{
  "action_type": "APPROVE",
  "action_context": "12345",
  "user_id": "b85s39b",
  "approval_code": "A3895051-9B16-4ABA-B785-AD2986177BB1",
  "instance_id": "people_1234",
  "task_id": "step1",
  "reason": "ok",
  "attachments": [
      {
          "file_type": "IMAGE",
          "file_name": "1.png",
          "file_size": 12345,
          "url": "https://sf3-cn.larkcdn.com/obj/lark-approval-attachment/image/20200512/413342ae-957f-4c6f-8d06-7dea05875d8b"
      }
  ],
  "token": "aaaaa"
}'

# Encrypted
curl --location --request POST 'https://www.larksuite.com/approval/openapi/v2/external/instanceOperate' \
--header 'Content-Type: application/json' \
--data-raw '{
    "encrypt": "=sfasdfasfsdafasfaf="
}'
```
```js
# Decryption  Sample code

package decrypt

import (
   "crypto/aes"
   "crypto/cipher"
   "crypto/rand"
   "crypto/sha256"
   "encoding/base64"
   "fmt"
   "io"
)

func CBCDecrypter(decryptContent string, keyStr string) ([]byte, error) {
   buf, err := base64.StdEncoding.DecodeString(decryptContent)
   if err != nil {
      return nil, err
   }
   key := sha256.Sum256([]byte(keyStr))
   if len(buf)%aes.BlockSize != 0 {
      return nil, fmt.Errorf("plaintext is not a multiple of the block size")
   }
   block, err := aes.NewCipher(key[:sha256.Size])
   if err != nil {
      return nil, err
   }
   ciphertext := make([]byte, aes.BlockSize+len(buf))
   iv := ciphertext[:aes.BlockSize]
   if _, err := io.ReadFull(rand.Reader, iv); err != nil {
      return nil, err
   }
   mode := cipher.NewCBCDecrypter(block, iv)
   mode.CryptBlocks(ciphertext[aes.BlockSize:], buf)
   ciphertext = ciphertext[32:]

   plain := standardizeDataDe(ciphertext)
   return plain, nil
}

func standardizeDataDe(origData []byte) []byte {
   length := len(origData)
   unpadding := int(origData[length-1])
   if unpadding > length {
      return nil
   }
   return origData[:(length - unpadding)]
}
```

### API description
|Parameter|Type|Required|Description|
|-|-|-|-|
|action_type|string|Yes|Action typeAPPROVE - Approve REJECT - Reject|
|action_context|string|No|Action context|
|user_id|string|Yes|Operator|
|approval_code|string|Yes|Approval definition code|
|instance_id|string|No|Instance ID (required for list actions)|
|task_id|string|No|Task ID (required for list actions)|
|message_id|int64|No|message_id returned when sending message card (required for card actions)|
|id|string|No|Task ID returned after search (required for actions after search)|
|reason|string|No|Reason|
|attachments|list|No|Attachment|
|  ∟file_type|string|Yes|Attachment type:IMAGE ATTACHMENTS|
|  ∟file_size|int|Yes|File size|
|  ∟file_name|string|Yes|File name|
|  ∟url|string|Yes|Attachment CDN URL |
|token|string|Yes|action_callback_token configured in definition|
|encrypt|string|No|If the definition is configured with action_callback_key, the above parameters will be encrypted and placed in this field, and the business party needs to use key for decryption.|
