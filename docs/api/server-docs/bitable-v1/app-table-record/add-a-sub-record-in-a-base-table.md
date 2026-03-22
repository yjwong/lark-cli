---
title: "Add a sub-record in a Base table"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-record/add-a-sub-record-in-a-base-table"
service: "bitable-v1"
resource: "app-table-record"
section: "Docs"
updated: "1753668090000"
---

# Add a sub-record in a Base table

Adding sub-records in Base tables essentially involves setting up one-way or two-way link fields between sub-records and parent records. The sub-records are mapped to the parent records through association fields, thereby establishing a connection. This document explains how to add a sub-record to a record in a Base table using OpenAPI.

## Prerequisite

You have already created a Base table, and there is an existing record in the table serving as the parent record.

## Process overview

The overall process of adding sub-records is as follows:
1. Call the Create Field API to create a one-way or two-way link field in the Base table to establish the relationship between records.
1. Call the Search Records API to get the ID of the existing parent record.
1. Call the Create Record API to add a new record as a sub-record and ensure that the record is associated with the existing parent record.
1. Call the List Fields API to get the field ID of the "one-way link field."
1. Call the Update View API and set the `field_id` parameter in the `hierarchy_config` to the field ID of the "one-way link field" to update the hierarchical structure style of the table view.

## Steps

This section takes an existing record as an example to illustrate how to add a sub-record to it.

1. Call the Create Field API to create a one-way link field in the Base table to establish the relationship between records. The request body is as follows:
    
    
    ```json
    {
      "field_name": "one-way link Field",
      "property": {
        "multiple": true,
        "table_id": "tblY2ha8xGSabcef"
      },
      "type": 18
    }
    ```
     
2. Call the Search Records API to get the ID of the existing parent record. The request body can be empty:

    ```json
    {}
    ```
   
   If the call is successful, the API will return data in the following structure, where `rec9k8PAbR` is the ID of the parent record.
   
    ```json
    {
      "code": 0,
      "data": {
        "has_more": false,
        "items": [
          {
            "fields": {
              "one-way link Field": {},
              "Text": [
                {
                  "text": "Parent Record",
                  "type": "text"
                }
              ]
            },
            "record_id": "rec9k8PAbR"   // Parent record ID
          },
          {
            "fields": {
              "one-way link Field": {}
            },
            "record_id": "recb9nHBYR"
          },
          {
            "fields": {
              "one-way link Field": {}
            },
            "record_id": "recwG1hh0g"
          },
          {
            "fields": {
              "one-way link Field": {}
            },
            "record_id": "recBlfgGRO"
          },
          {
            "fields": {
              "one-way link Field": {}
            },
            "record_id": "recKZHTepH"
          },
          {
            "fields": {
              "one-way link Field": {}
            },
            "record_id": "recDZXc9fs"
          },
          {
            "fields": {
              "one-way link Field": {}
            },
            "record_id": "recX9dPV90"
          },
          {
            "fields": {
              "one-way link Field": {}
            },
            "record_id": "rec6cq2RIk"
          },
          {
            "fields": {
              "one-way link Field": {}
            },
            "record_id": "recuK6kUA1"
          },
          {
            "fields": {
              "one-way link Field": {}
            },
            "record_id": "recLjQD5Eo"
          }
        ],
        "total": 10
      },
      "msg": "success"
    }
    ```
   
1. Call the Create Record API to add a new record as a sub-record and ensure that the record is associated with the existing parent record. The request body is as follows:

    ```json
    {
      "fields": {
        "one-way link Field": [
          "rec9k8PAbR"
        ],
        "Text": "Sub-record"
      }
    }
    ```
        

2. Call the List Fields API to get the field ID of the "one-way link field." If the call is successful, the API will return data in the following structure, where `fldfASqam8` is the field ID of the "one-way link field."
    
    
    ```json
    {
      "code": 0,
      "data": {
        "has_more": false,
        "items": [
          {
            "field_id": "fldplsqa97",
            "field_name": "Text",
            "is_hidden": false,
            "is_primary": true,
            "property": null,
            "type": 1,
            "ui_type": "Text"
          },
          {
            "field_id": "fldfASqam8",   // Field ID of the one-way link field
            "field_name": "one-way link Field",
            "is_hidden": false,
            "is_primary": false,
            "property": {
              "multiple": true,
              "table_id": "tblmyHKpQG3k1kSD",
              "table_name": "Data Table"
            },
            "type": 18,
            "ui_type": "SingleLink"
          }
        ],
        "page_token": "fldfASqam
        "total": 2
      },
      "msg": "success"
    }
    ```
    
1. Call the Update View API, and refer to the following request example. Set the `field_id` parameter in the `hierarchy_config` to the field ID of the "one-way link field" to update the hierarchical structure style of the table view.

    ```json
    {
      "property": {
        "hierarchy_config": {
          "field_id": "fldfASqam8"
        }
      },
      "view_name": "table"
    }
    ```
