---
title: "Error Codes"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uMDOyUjLzgjM14yM4ITN"
section: "Deprecated Version (Not Recommended)"
updated: "1646916351000"
---

# Error codes

If a request fails, the API will return an error code (anything but 0). The meanings are as follows:

| Error code | Description                                                  | Error information                        |
| ---------- | ------------------------------------------------------------ | ---------------------------------------- |
| 100001     | Invalid page token format                                    | page token invalid                       |
| 100002     | Invalid field names exist under fields                       | Invalid field selection [illegal field]  |
| 100003     | Time is not in [RFC3339](https://tools.ietf.org/html/rfc3339) format | time format must follow RFC3339 standard |
| 100004     | Invalid building ID                                          | building id invalid                      |
| 100005     | Invalid room ID                                              | room id invalid                          |
| 105001     | Internal error                                               | internal error                           |
