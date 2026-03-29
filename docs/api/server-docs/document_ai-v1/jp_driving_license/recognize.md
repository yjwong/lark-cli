---
title: "Identify Japanese driving license in documents"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/document_ai-v1/jp_driving_license/recognize"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/document_ai/v1/jp_driving_license/recognize"
service: "document_ai-v1"
resource: "jp_driving_license"
section: "AI"
rate_limit: "10 per second"
scopes:
  - "document_ai:jp_driving_license"
updated: "1710921009000"
---

# Identify the ticket in the file

Japan driving license identification interface, supporting JPG/JPEG/PNG/BMP four file types of one-time recognition.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/document_ai/v1/jp_driving_license/recognize |
| HTTP Method | POST |
| Rate Limit | 10 per second |
| Supported app types | custom,isv |
| Required scopes | `document_ai:jp_driving_license` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Example value**: "multipart/form-data; boundary=---7MA4YWxkTrZu0gW" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `file` | `file` | Yes | Identification of Japanese driver's license source documents **Example value**: file binary | ### Request body example

```HTTP
---7MA4YWxkTrZu0gW
Content-Disposition: form-data; name="file";
Content-Type: application/octet-stream

---7MA4YWxkTrZu0gW
```

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `jp_driving_license` | `jp_driving_license` | Japanese driving license information |
|     `entities` | `jp_driving_license_entity[]` | The type of entity identified |
|       `type` | `string` | The type of field identified **Optional values are**:  - `card_number`: Driving license number - `full_name`: Name - `birthday`: Birthday - `birthday_2nd`: Birthday (not a common expression in the country) - `address`: Address - `effective_date`: Period of validity - `issue_date`: Date of issue  |
|       `value` | `string` | Identify the text information of the field | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "jp_driving_license": {
            "entities": [
                {
                    "type": "card_number",
                    "value": "12**56"
                }
            ]
        }
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 2110001 | Param is invalid | Input file error, refer to the documentation to check the input parameters |
| 400 | 2110002 | No valid entity | Invoice information is not detected, refer to the document to check whether the input file is valid |
| 500 | 2110010 | Internal error, please try later. | Backend service exception or network exception, can be rerequested |
| 400 | 2110003 | You have reached the Intelligent document parsing limit. To continue using this function, please contact sales to purchase more. | You have reached the Intelligent document parsing limit. To continue using this function, please contact sales to purchase more. |
