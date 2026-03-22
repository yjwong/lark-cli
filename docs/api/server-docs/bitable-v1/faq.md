---
title: "Base FAQs"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/faq"
service: "bitable-v1"
resource: "faq"
section: "Docs"
updated: "1753668105000"
---

# Base FAQs

## 1. How to obtain the `file_token` for an attachment when adding a new record in a Base table?

When adding a new record, the `file_token` for attachment-type fields requires the developer to first upload the attachment to the corresponding Base table.

Therefore, you need to call the Upload Media or Upload a file in blocks-Pre­uploading API to upload the attachment to the Base table to obtain the file_token, and then write the record into the Base table.

Note: The file_token is only valid within the current Base table. If you need to upload the attachment to another Base table, you will need to re-upload the media to obtain a new file_token.

## 2. How to download files from records in a Base table?

You can download attachments from the Base table using the Download Media API.

For Base tables with advanced permissions, when calling the Download Media or Get Temporary Media Download Link API, you need to add the `extra` parameter as a URL query parameter for authentication. You can obtain the `extra` parameter as follows:

1. Call the Search Records API. The response example will return the download link for the attachment, as shown below:
```json
{
  "code": 0,
  "data": {
    "has_more": false,
    "items": [
      {
        "fields": {
          "Attachment": [
            {
              "file_token": "RSkabsaphoy6yVxK0mGcggabcef",
              "name": "74d2c703524489dbabcef.png",
              "size": 87123,
              "tmp_url": "https://open.larksuite.com/open-apis/drive/v1/medias/batch_get_tmp_download_url?file_tokens=RSkabsaphoy6yVxK0mGcggabcef&extra={\"bitablePerm\":{\"tableId\":\"tblz8ExGaVhuiSD1\",\"rev\":32}}",  // Temporary download link for the media, the URL value must be URL-encoded
              "type": "image/png",
              "url": "https://open.larksuite.com/open-apis/drive/v1/medias/RSkabsaphoy6yVxK0mGcggabcef/download?extra={\"bitablePerm\":{\"tableId\":\"tblz8ExGaVhuiSD1\",\"rev\":32}}"   // Download URL for the media, the URL value must be URL-encoded
            }
          ]
        },
        "record_id": "recbMzD0zT"
      }
    ],
    "total": 1
  },
  "msg": "success"
}
```

Example of constructing the URL:

```JSON
{"bitablePerm":{"tableId":"tblO6OeNZxfabcef","attachments":{"fld32zZi5I":{"rec0BuOHq":["boxbcsQNT0JsmrztOnX530abcef"]}}}}
// After escaping
https://open.larksuite.com/open-apis/drive/v1/medias/boxbcsQNT0JsmrztOnX530abcef/download?extra=%7B%22bitablePerm%22%3A%7B%22tableId%22%3A%22tblO6OeNZxfabcef%22%2C%22attachments%22%3A%7B%22fld32zZi5I%22%3A%7B%22rec0BuOHq%22%3A%5B%22boxbcsQNT0JsmrztOnX530abcef%22%5D%7D%7D%7D%7D
```

## 3. What to do if you encounter error codes like 1254040, 1254041, 1254042, 1254043, or 1254044 (NotFound) when calling Base table APIs?

These error codes indicate that the resource corresponding to the provided ID could not be found. You can try the following methods to resolve the issue:

1. Verify that the ID is correct. You can refer to the Overview of Base tables to find the correct IDs for different resources.
2. If this error occurs shortly after creating a resource, it may be caused by server replication delays. Try waiting for a few seconds and then retrying the request.

## 4. How to resolve the issue when the query record interface is called successfully but the data returned is empty, even though there is data in the multi-dimensional table?

This is usually due to the advanced permissions enabled on the multi-dimensional table, resulting in insufficient calling identity permissions. You need to grant the calling identity **Manage** permissions for the data table or **Manage** permissions for the multi-dimensional table, and then call again. The specific steps are as follows:

- To grant the user manage permissions, you can add the user in the **Multi-dimensional Table Advanced Permissions Settings** and grant sufficient permissions to the user. 

- To grant the application manage permissions, you need to add manage permissions for the application through the **"..."** -> **"...More"** -> **"Add Document Application"** entry in the upper right corner of the multi-dimensional table page.

    ![](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/22c027f63c540592d3ca8f41d48bb107_CSas7OYJBR.png?height=1994&lazyload=true&maxWidth=550&width=3278)
   
    ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/9f3353931fafeea16a39f0eb887db175_0tjzC9P3zU.png?height=728&lazyload=true&maxWidth=550&width=890)
    
    **Note**:
    Before adding a document application, you need to ensure that the target application has at least one multi-dimensional table API permission enabled. Otherwise, you will not be able to search for the target application in the document application window.

- You can also add a user or a group containing the application in the **Multi-dimensional Table Advanced Permissions Settings**, and grant custom read-write permissions to this group.

## 5. How to resolve the issue when the query record interface returns empty data for a checkbox?

In a multi-dimensional table, if a field is empty, the query record interface does not return data. Accordingly, if the checkbox field is empty (i.e., the user has not selected or deselected the checkbox), the query record interface also does not return data. In this scenario, since the effect of an empty value is the same as the effect of the checkbox being `false`, developers need to handle this empty value scenario on their own.

## 6. How to filter custom number type auto-number fields?

To filter auto-number fields containing fixed characters, you need to remove the custom fixed characters before filtering; otherwise, empty data will be returned.

## 7. How to filter personnel fields when calling the query record interface?

For personnel fields, you need to pass the user's ID in the value. The type of user ID is consistent with the type specified by the `user_id_type` query parameter in the query record interface, which defaults to the Open ID type. The following is a request example for filtering personnel IDs `ou_00d9ea2d7bcd6b6aa7d71dd88deabcef` and `ou_5fdfc3d312b24d28224769baf52abcef`:

```json
{
  "view_id": "vewfrk8iX4",
  "field_names": [
    "creator"
  ],
  "filter": {
    "conjunction": "and",
    "conditions": [
      {
        "field_name": "creator",
        "operator": "contains",
        "value": [
          "ou_00d9ea2d7bcd6b6aa7d71dd88deabcef",
          "ou_5fdfc3d312b24d28224769baf52abcef"
        ]
      }
    ]
  }
}
```

## 8. Does the query record interface support querying specific rows, such as rows 10 to 20 in the data table?

1. Add an auto-numbering field in the multidimensional table, and select auto-increment numbers as the numbering type.

2. Filter the auto-numbering. For example, to query rows 10 to 20, filter the numbers less than 21 and greater than 9, and then call the Query Records interface. The request body is as follows:

    ```json
    {
      "view_id": "vewfrk8iX4",
      "field_names": [],
      "filter": {
        "conjunction": "and",
        "conditions": [
          {
            "field_name": "number",
            "operator": "isGreater",
            "value": [
              "9"
            ]
          },
          {
            "field_name": "number",
            "operator": "isLess",
            "value": [
              "21"
            ]
          }
        ]
      }
    }
    ```

You can also call the Batch Get Records interface to query using the Record ID of the records to get multiple records' data.

## 9. How to get the total number of records (or total rows) of a specified data table in the multidimensional table?

You can call the Query Records interface and set the `conditions` field to empty or not set it in the request body. Here is a request body example:

```json
{
  "view_id": "vewfrk8iX4",  // Please replace with the actual view_id
  "filter": {
    "conjunction": "and",
    "conditions": []
  }
}
```
If the request is successful, the response body will return the `total` field, which is the total number of records.
