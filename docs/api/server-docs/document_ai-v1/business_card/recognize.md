---
title: "Recognize business cards in pictures"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/document_ai-v1/business_card/recognize"
method: "POST"
api_path: "https://open.larksuite.com/open-apis/document_ai/v1/business_card/recognize"
service: "document_ai-v1"
resource: "business_card"
section: "AI"
rate_limit: "10 per second"
scopes:
  - "document_ai:business_card:recognize"
updated: "1698807495000"
---

# Identify business cards in files

Business card recognition interface, one-time business card recognition by uploading file types such as JPG/PNG/PDF. The interface is suitable for files under 20MB, and is suitable for business cards in English and Japanese.

> The limited viewership of single tenant: 10QPS, there is no limited viewership of applications under the same tenant, and the 10QPS limited viewership of this tenant is shared

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/document_ai/v1/business_card/recognize |
| HTTP Method | POST |
| Rate Limit | 10 per second |
| Supported app types | custom,isv |
| Required scopes | `document_ai:business_card:recognize` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer t-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token |
| Content-Type | string | Yes | **Example value**: "multipart/form-data; boundary=---7MA4YWxkTrZu0gW" | ### Request body

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `file` | `file` | Yes | Identify the source file of the business card (supports JPG/PNG/PDF) **Example value**: File binary | ### Request body example

```HTTP
---7MA4YWxkTrZu0gW
Content-Disposition: form-data; name="file";
Content-Type: application/octet-stream

File binary
---7MA4YWxkTrZu0gW
```

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `business_cards` | `recognized_entities[]` | Business card information |
|     `entities` | `recognized_entity[]` | List of recognized entities |
|       `type` | `string` | Identified field types **Optional values are**:  - `contact_names`: Contact name - `company_names`: Company name - `departments`: department - `job_titles`: position - `emails`: Email - `websites`: website - `addresses`: address - `mobile_phones`: Mobile phone - `work_phones`: Work phone - `other_phones`: Other phone calls - `faxes`: fax  |
|       `value` | `string` | Recognize the text information of the field | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "business_cards": [
            {
                "entities": [
                    {
                        "type": "contact_names",
                        "value": "Tom"
                    }
                ]
            }
        ]
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 2110001 | invalid file | Input file error, refer to the documentation to check the input parameters |
| 400 | 2110002 | no business card detected | The business card information is not detected, refer to the document to check whether the input file is valid. |
| 500 | 2110010 | network anomaly, please try again | Backend service exception or network exception, can be rerequested |
| 400 | 2110003 | You have reached the Intelligent document parsing limit. To continue using this function, please contact sales to purchase more. | You have reached the Intelligent document parsing limit. To continue using this function, please contact sales to purchase more. |
