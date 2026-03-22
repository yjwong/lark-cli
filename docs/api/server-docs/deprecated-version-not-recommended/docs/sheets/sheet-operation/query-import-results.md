---
title: "Query Import Results"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uETO2YjLxkjN24SM5YjN"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/sheets/v2/import/result"
section: "Deprecated Version (Not Recommended)"
scopes:
  - "drive:drive"
  - "sheets:spreadsheet"
updated: "1647001467000"
---

# Query import result
Use this API to query the file import result. If no result is shown within 30 minutes, the import failed.

## Request
| HTTP URL | https://open.larksuite.com/open-apis/sheets/v2/import/result |
| --- | --- |
| HTTP Method | GET |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `drive:drive` `sheets:spreadsheet` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `user_access_token`  or  `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about how to obtain and use access_token |
| Content-Type | string | Yes | **Fixed value**: "application/json; charset=utf-8" | ###  Query parameters    

| Parameter          | Type   | Required | Description                                                         | 
| ------------- | ------ | ---- | ------------------------------------------------------------ | 
| ticket        | string | Yes   | Credential obtained during import                                             | 
### cURL Request example
``` 
curl --location --request GET 'https://open.larksuite.com/open-apis/sheets/v2/import/result?ticket=6948310613016144800' \
--header 'Authorization: Bearer t-e346617a4acfc3a11d4ed24dca0d0c0fc8e0067e' \
```

## Response  

### Response body
| Parameter        |Type| Description               |
| ----------- |-----| ------------------ |
| ticket      |string| Credential obtained during import   |
| url         |string| Link of the imported document     |
| warningCode |int| Status code of the imported document | warningCode information is stored in binary digits. The meaning of each bit is as follows:

| 5            | 4            | 3              | 2            | 1          | 0      |
| ------------ | ------------ | -------------- | ------------ | ---------- | ------ |
| Folder does not exist | Failed to import image | Exceeds My Space image limit | Image exceeds limit | Cell truncated | Column truncated | For warningCode=18, 18 (decimal)=10010 (binary) . This corresponds to "Cell truncated" and "Failed to import image".

| 0            | 1            | 0              | 0            | 1          | 0      |
| ------------ | ------------ | -------------- | ------------ | ---------- | ------ |
| Folder does not exist | Failed to import image | Exceeds My Space image limit | Image exceeds limit | Cell truncated | Column truncated | ### Response body example    

```json
{
    "code": 0,
    "msg": "Success",
    "data": {
        "ticket": "6891610404246520328",
        "url": "https://bytedance.feishu.cn/sheets/shtcnaryaxxxx",
        "warningCode": 32
    }
}
```
### Error code

For details, refer to Server-side error codes.
