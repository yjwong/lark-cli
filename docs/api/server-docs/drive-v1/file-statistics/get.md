---
title: "Get File Statistics"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file-statistics/get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/drive/v1/files/:file_token/statistics"
service: "drive-v1"
resource: "file-statistics"
section: "Docs"
scopes:
  - "drive:drive.metadata:readonly"
  - "drive:drive:readonly"
  - "drive:drive"
updated: "1647178821000"
---

# Obtain statistics of a file

This API is used to obtain statistics of a file, including the number of unique visitors (UVs), the number of page views (PVs), and the number of likes.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/drive/v1/files/:file_token/statistics |
| HTTP Method | GET |
| Supported app types | custom,isv |
| Required scopes Enable any scope from the list | `drive:drive.metadata:readonly` `drive:drive:readonly` `drive:drive` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" Learn more about obtaining and using access_token. | ::: note
For more information about how to call the AccessToken in the Docs API, see Get started with the Docs API.

### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `file_token` | `string` | Token of the file **Example value**: "doccnRs*******" | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `file_type` | `string` | Yes | Type of the document **Example value**: "doc" **Optional values are**: - `doc`: Doc - `sheet`: Sheet - `mindnote`: MindNotes - `bitable`: Bitable - `wiki`: Wiki - `file`: File | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
| ∟ `file_token` | `string` | Token of the file |
| ∟ `file_type` | `string` | File type |
| ∟ `statistics` | `file_statistics` | Statistics of the file |
| ∟ `uv` | `int` | Number of historical visitors of the file. Multiple visits by the same user with the same user_id are counted as one UV. |
| ∟ `pv` | `int` | Number of historical visits of the file. Multiple visits by the same user with the same user_id are counted as multiple PVs. (Note: Two consecutive visits by the same user within half an hour is regarded as one PV.) |
| ∟ `like_count` | `int` | Total number of historical likes of the file. If the specified document type doesn't support liking, the value –1 is returned. |
| ∟ `timestamp` | `int` | Timestamp in seconds | ### Response body example

```json
{
    "code": 0,
    "msg": "success",
    "data": {
        "file_token": "doccnRs*******",
        "file_type": "doc",
        "statistics": {
            "uv": 10,
            "pv": 15,
            "like_count": 2,
            "timestamp": 1627367349
        }
    }
}
```

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 500 | 1069601 | fail | Please try again. If the problem persists, contact the oncall personnel of the related service party. |
| 400 | 1069602 | param error | Check whether the parameter is valid. |
| 403 | 1069603 | forbidden | Check whether the user has the permission to read the file. |
| 400 | 1069604 | document not found | Check whether the file exists. |
