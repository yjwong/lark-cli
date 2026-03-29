---
title: "Upload Image"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uEDO04SM4QjLxgDN"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/image/v4/put/"
section: "Deprecated Version (Not Recommended)"
updated: "1626158294000"
---

# Upload image
Uploads an image, and obtains the image_key.

> Bot capability needs to be enabled. Image size cannot exceed 10MB

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/image/v4/put/ |
| HTTP Method | POST | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. |
| Content-Type | string | Yes | **Example value**: "multipart/form-data; boundary=---7MA4YWxkTrZu0gW" | ### Request body
|Parameter|Type|Mandatory|Description|Default|Demo|
-- | -- | -- | -- | -- | --
image | binary | Yes |Image file||image binary data|
image_type | string | Yes |Image type, could be "message", "avatar".||message| ### Request body example

```HTTP
---7MA4YWxkTrZu0gW
Content-Disposition: form-data; name="image";
Content-Type: application/octet-stream

image binary data
---7MA4YWxkTrZu0gW
Content-Disposition: form-data; name="image_type";

message
---7MA4YWxkTrZu0gW
```

## Response
### Response boby
|Parameters|Type|Description|
|-|-|-|
code |int| Error code, anything but 0 indicates failure
msg |string| Error code description.
data | - | -
  ∟image_key |string| Image key

### Response body example
```json
{
	"code": 0,
	"data": {
    	"image_key": "fd3a475a-0c27-4071-a9a6-8712b84b0cb6"
	},
	"msg": "ok"
}
```

### Demo - Java

```java
package com.byted.ee.oapi.bot;

import org.apache.http.Header;
import org.apache.http.HttpEntity;
import org.apache.http.client.methods.CloseableHttpResponse;
import org.apache.http.client.methods.HttpPost;
import org.apache.http.config.Registry;
import org.apache.http.config.RegistryBuilder;
import org.apache.http.conn.socket.ConnectionSocketFactory;
import org.apache.http.conn.socket.PlainConnectionSocketFactory;
import org.apache.http.conn.ssl.NoopHostnameVerifier;
import org.apache.http.conn.ssl.SSLConnectionSocketFactory;
import org.apache.http.conn.ssl.TrustSelfSignedStrategy;
import org.apache.http.entity.mime.MultipartEntityBuilder;
import org.apache.http.entity.mime.content.FileBody;
import org.apache.http.impl.client.BasicCookieStore;
import org.apache.http.impl.client.CloseableHttpClient;
import org.apache.http.impl.client.HttpClients;
import org.apache.http.impl.conn.PoolingHttpClientConnectionManager;
import org.apache.http.ssl.SSLContextBuilder;
import org.apache.http.util.EntityUtils;

import java.io.File;
import java.io.IOException;
import java.security.KeyManagementException;
import java.security.KeyStoreException;
import java.security.NoSuchAlgorithmException;

public class SendImageExample {

    public static CloseableHttpClient getHttpClient() {
        try {
            SSLContextBuilder builder = new SSLContextBuilder();
            builder.loadTrustMaterial(null, new TrustSelfSignedStrategy());
            SSLConnectionSocketFactory sslConnectionSocketFactory = new SSLConnectionSocketFactory(builder.build(),
                    NoopHostnameVerifier.INSTANCE);
            Registry registry = RegistryBuilder.create()
                    .register("http", new PlainConnectionSocketFactory())
                    .register("https", sslConnectionSocketFactory)
                    .build();

            PoolingHttpClientConnectionManager cm = new PoolingHttpClientConnectionManager(registry);
            cm.setMaxTotal(100);
            CloseableHttpClient httpclient = HttpClients.custom()
                    .setSSLSocketFactory(sslConnectionSocketFactory)
                    .setDefaultCookieStore(new BasicCookieStore())
                    .setConnectionManager(cm).build();
            return httpclient;
        } catch (KeyManagementException e) {
            e.printStackTrace();
        } catch (NoSuchAlgorithmException e) {
            e.printStackTrace();
        } catch (KeyStoreException e) {
            e.printStackTrace();
        }
        return HttpClients.createDefault();
    }

    public static String SendImageByApacheHttpClient(File file) throws IOException {

        CloseableHttpClient client = getHttpClient();

        HttpPost post = new HttpPost("https://open.larksuite.com/open-apis/image/v4/put/");

        final MultipartEntityBuilder builder = MultipartEntityBuilder.create();
        FileBody bin = new FileBody(file);
        builder.addPart("image", bin);
        builder.addTextBody("image_type", "message");
        HttpEntity multiPartEntity = builder.build();

        post.setEntity(multiPartEntity);
        post.setHeader("Authorization", "Bearer t-84d31e38a0415bd2db0ee0e8f1dbdb011b151d0f");

        CloseableHttpResponse response = client.execute(post);

        System.out.println("http response code:" + response.getStatusLine().getStatusCode());
        for (Header header: response.getAllHeaders()) {
            System.out.println(header.toString());
        }
        HttpEntity resEntity = response.getEntity();

        if (resEntity == null) {
            System.out.println("never here?");
            return "";
        }
        System.out.println("Response content length: " + resEntity.getContentLength());
        return EntityUtils.toString(resEntity);
    }

    public static void main(String[] args) throws IOException {
        File file = new File("/Users/chengzhiwang/Downloads/IMG_0008.jpg");
        String result = SendImageByApacheHttpClient(file);
        System.out.println(result);

    }

}

```

### Demo - Golang

```go
package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"unsafe"
)

func main()  {
	data := make(map[string]string)
	newFileUploadRequest("https://open.larksuite.com/open-apis/image/v4/put/",
		data,"image","/Users/fanlv/Desktop/IMG_2871.JPG")

}

func newFileUploadRequest(uri string, params map[string]string, paramName, path string)  error {
	file, err := os.Open(path)
	if err != nil {
		return  err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, path)
	if err != nil {
		return  err
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
    writer.WriterField("image_type", "message")
	err = writer.Close()
	if err != nil {
		return err
	}
	request, err := http.NewRequest("POST", uri, body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	request.Header.Set("Authorization", "Bearer t-6fd30ee3784982904166bab48b55e416c68a3d78")

	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	str := (*string)(unsafe.Pointer(&respBytes))
	fmt.Println(*str)

	return  err
}

```
### Error code

For details, please refer to: Service-side error codes
