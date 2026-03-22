---
title: "Associate external options"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uADM4QjLwADO04CMwgDN"
section: "Approval"
updated: "1675167395000"
---

# Associate external options

## Instructions

### When can I use this API?

You can use this function when your company uses multiple systems (Lark Approval, HR systems, and sales management systems) and you must sync data from other systems to use as options in approval forms. You can configure external data sources as single option/multiple options widget options, so you don't have to manually ensure the consistency across multiple systems.

For example: If you configure a sales-related approval process in Lark Approval and sales staff must enter external customer names when submitting approvals but the customer name list is maintained in the sales management system, you can configure this information as external data for single option/multiple options. Then, when sales employees submit approvals, they can select their own customers from a names list taken from the sales management system. The data in the approval system can be automatically synced and doesn't have to be maintained.

### What problems can this API solve?
- With this API, you don't have to repeatedly update and modify data to maintain data consistency across different systems.
- No matter how many total options there are, employees can only select relevant options after configuration.

### Complete the configuration in three steps
- Company developers develop relevant data APIs based on the development requirements in "Developer documentation".
- Approval administrators specify the developed APIs as data source APIs. After they are verified, approval initiators can directly select the options synced through the APIs when initiating approvals.
- To ensure data security, developers can set API tokens and keys to be input by administrators during configuration. This can protect against data leakage (key is optional).

## Developer documentation
Programs an HTTP or HTTPS API and sets the token and key (in any format). The API implementation has no programming language limitations. The token is used to verify the request source, and the key is used for encryption and decryption (the encryption key is optional; if none is provided, encryption will not be performed.) The call method, returned parameter formats, and encryption and decryption methods are described below.

### 1. Call method

If the form is in editing status and the data source is an external system widget, the approval system will initiate an **HTTP or HTTPS** request to the external data source API address when a user clicks to validate data or initiates a request. 

**Request address**: The request address configured by the user
**Request mode**: POST
**Request timeout duration**: 3 sec
**Request header**:
key|value
--|--
Content-Type|application/json
Currently, you can request different data to be input to single-option and multi-option widgets through the **user ID and employee ID fields or associated extra fields in a form**. If user_id and employee_id are both empty, all options will be returned. All widget HTTP requests use the following input parameter format:

```js 
{
        "user_id": "123",
        "employee_id": "abc",
        "token":"1e8e999f580e7a202dbe1e5103c5e4c58ecc757e",
        "linkage_params":{
          "key1":"value1", // key1 is the field code of the linked field, and value1 is the linked widget value.
          "key2":"value2", // key2 is the field code of the linked field, and value2 is the linked widget value.
        }
} 
```
**Request parameter description**:  
|Parameter|Type|Required|Description|
|-|-|-|-|
|user_id|String|Yes|The user's user_id|
|employee_id|String|Yes|The employee's employee_id|
|token|String|Yes|The token used to verify that the request is from a valid source|
|linkage_params|Map|No|Linked parameters(please return all options when linkage_params is empty)| Note: Currently, the Open Platform considers a user ID to be the same as an employee ID.

### 2. Returned parameter format

Format of output parameters before encryption:

```js 
{
    "code":0,
    "msg":"success!",
    "data":{
        "result":{
            "options":[
                {
                    "id":"id1",
                    "value":"name1",
                    "isDefault":true
                },
                {
                    "id":"id2",
                    "value":"name2"
                },
                {
                    "id":"id3",
                    "value":"name3"
                }
            ],
            "i18nResources":[
                {
                    "locale":"zh_cn",
                    "isDefault":true,
                    "texts":{
                        "name1":"Value 1",
                        "name2":"Value 2",
                        "name3":"Value 3"
                    }
                },
                {
                    "locale":"en_us",
                    "isDefault":false,
                    "texts":{
                        "name1":"value1",
                        "name2":"value2",
                        "name3":"value3"
                    }
                }
            ]
        }
    }
} 
```

**Return parameter description**: (i18nResources Required to pass value)
|Parameter|Type|Description|
|-|-|-|
|code|int|Error code. A value other than 0 indicates failure.|
|msg|string|Return code description|
|data|object|Returned business information|
|  ∟result|object|Request result content|
|    ∟options|list|Option list|
|    ∟i18nResources|list|Internationalized text| **externalData structure**: 
|Parameter|Type|Description|
|-|-|-|
|id|string|Unique ID of an option|
|value|string|Option(It is required to ensure the uniqueness)| |isDefault|bool|Default Option or Not| **i18nResource structure**: 
|Parameter|Type|Description|
|-|-|-|
|locale|string|Language (zh_cn: Chinese, en_us: English, ja_jp: Japanese)
|isDefault|bool|Whether it is a default option
|texts|map[string]string|Internationalized text map. key is the unique value of an internationalized option. This value is the same in different language environments. value is the text for a specific language environment.

  
Format of output parameters after encryption (Encrypt the result content and convert it to base64 output, and return it in plaintext without configuring the key):

```js 
{
    "code":0,
    "msg":"success!",
    "data":{
        "result":"tKqgkBNFEzakJAeS/ySKS7j7YoX2rKVuzLJbG44xHsz0eHaqLx6ZLsAQ/ljfK9mDi0F/32UVXM3gUQaczHbR2upD/EStb+O26FApdvNKm0yvKG0WrhFIe7UCMkrxPnegBqqgqcMHLCZQZ2uh/2k5dDlhReT6fxm/bAR4ZwgyvvshqudakKigshSK0Aq25IQ0H65PS/5iRHgk2b06sahZuvH6b9yrfBXJqHdhztvPkPW2FkipbvLMrzQdXz+deBm2DTJ5W53f2QKOxk7szaXKOr1+u1MyCIkjldPcAHqPYRiOzx6iXQPJ6hMj7MHex08amm44d5T3Z2jzCoinkGSrhpusTcmhHmQnjDjl51a2LqBlty1L9yHuMaED+al2lTUhlzGHqhITCQBJLZraOkXYcR6oOXAV3gP4towZw5G/zeeEtXYZvWUvTZ9F3UAXM4jP"
    }
} 
```

**Return parameter description**: 
|Parameter|Type|Description|
|-|-|-|
|code|int|Error code. A value other than 0 indicates failure.|
|msg|string|Return code description|
|data|string|Returned business data|
|  ∟result|string|Request result content encrypted in base64| ### 3. Encryption/Decryption method

Golang encryption code is shown below:

```js 
//AES CBC encryption
func CBCEncrypter(buf []byte, keyStr string) ([]byte, error) {
	key := sha256.Sum256([]byte(keyStr))
	plaintext := standardizeDataEn(buf)

	if len(plaintext)%aes.BlockSize != 0 {
		return nil, errors.New("plaintext is not a multiple of the block size")
	}

	block, err := aes.NewCipher(key[:sha256.Size])
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext, nil
}

func standardizeDataEn(data []byte) []byte {
	appendingLen := aes.BlockSize - (len(data) % aes.BlockSize)
	sd := make([]byte, len(data)+appendingLen)
	copy(sd, data)
	for i := 0; i 
 Java encryption code example:
  
 
```js 
 public String CBCEncrypter(String key, String source){
        try {
            MessageDigest messageDigest = MessageDigest.getInstance("SHA-256");
            messageDigest.reset();
            messageDigest.update(key.getBytes());

            SecretKeySpec skeySpec = new SecretKeySpec(messageDigest.digest(), "AES");
            Cipher cipher = Cipher.getInstance("AES/CBC/PKCS5Padding");//"Algorithm/Mode/Complement method"
            byte[] sSrcBytes = source.getBytes();
            byte[] newSrc =  new byte[sSrcBytes.length + 16];
            byte[] cSrc = new byte[16];
            System.arraycopy(cSrc, 0, newSrc, 0, cSrc.length);
            System.arraycopy(sSrcBytes, 0, newSrc, 16, sSrcBytes.length);
            IvParameterSpec iv = new IvParameterSpec(cSrc);//Using CBC mode, a vector iv is required, which can increase the strength of the encryption algorithm
            cipher.init(Cipher.ENCRYPT_MODE, skeySpec, iv);
            byte[] encrypted = cipher.doFinal(newSrc);
            return Base64.getEncoder().encodeToString(encrypted);//Here, Base64 is used for transcoding and can be used for secondary encryption
        } catch (Exception e) {
            //handle Exception
        }
        return null;
    } 
```

The decryption code is shown below:

```js 
//AES CBC decryption
func CBCDecrypter(buf []byte, keyStr string) ([]byte, error) {
	key := sha256.Sum256([]byte(keyStr))
	if len(buf)%aes.BlockSize != 0 {
		return nil, errors.New("plaintext is not a multiple of the block size")
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

func RandKey256() (string, error) {
	key := make([]byte, 32)

	if _, err := rand.Read(key); err != nil {
		return "", err
	} else {
		return  string(key), nil
	}
} 
```
