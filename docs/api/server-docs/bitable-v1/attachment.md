---
title: "Attachment Field"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/attachment"
service: "bitable-v1"
resource: "attachment"
section: "Docs"
updated: "1753668098000"
---

# Attachment field

> Before using this API, carefully read Materials Overview.

## Data structure

| Field name | Field description |
|:---|:---| |file_token | file token| | name | file name | |size | file size | |tmp_url|A temporary download link| |type|file type| |url| URL for download the file| **Response body**

```json 
{
  "Attachment": [
    {
      "file_token": "boxcnzm3dPEcutYDPplx5iDak4b",
      "name": "Hawaii_1_15Retina_R.jpg",
      "size": 5069121,
      "tmp_url": "https://open.larksuite.com/open-apis/drive/v1/medias/batch_get_tmp_download_url?file_tokens=boxcnzm3dPEcutYDPplx5iDak4b",
      "type": "image/jpeg",
      "url": "https://open.larksuite.com/open-apis/drive/v1/medias/boxcnzm3dPEcutYDPplx5iDak4b/download"
    }
]
}
``` 

## Upload attachments

Upload attachments, upload attachments in a multi-dimensional table in 2 steps

1.CallUpload media or Upload media in blocks interface upload file, upload success after obtaining the file fille_token

2.Call CreateRecord or UpdateRecord to update files to records; 

**Request body**

```json 
 {
    "records": [
        {
            "fields": {
            "Attachment": [
                {"file_token": "boxbcCFb2dBwMK9S8kDILk1tayh"},
                {"file_token": "boxbcCFb2dBwMK9S8kDILk1tayh"}
                ]
             }
        },
        {
            "fields": {
            "Attachment": [
                {"file_token": "boxbcCFb2dBwMK9S8kDILk1tayh"},
                {"file_token": "boxbcCFb2dBwMK9S8kDILk1tayh"}
                ]
             }
        }

     ]

}
``` 
**Response body：**
```json 
{
    "code": 0,
    "data": {
        "records": [
            {
                "fields": {
                    "Attachment": [
                        {
                            "file_token": "boxbcCFb2dBwMK9S8kDILk1tayh"
                        },
                        {
                            "file_token": "boxbcCFb2dBwMK9S8kDILk1tayh"
                        }
                    ]
                },
                "id": "recxgOKlB0"
            },
            
            {
                "fields": {
                    "Attachment": [
                        {
                            "file_token": "boxbcCFb2dBwMK9S8kDILk1tayh"
                        },
                        {
                            "file_token": "boxbcCFb2dBwMK9S8kDILk1tayh"
                        }
                    ]
                },
                "id": "reciGVHpI8"
            }
        ]
    },
    "msg": "Success"
}

``` 

## Download Attachment

1.Call ListRecord  to query the attachments file token

2.CallDownload media . or Get temporary download URL of media These interface already supports Base.
