---
title: "Record filter guide"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-record/record-filter-guide"
service: "bitable-v1"
resource: "app-table-record"
section: "Docs"
updated: "1752650893000"
---

# Record filter development guide

In some Base APIs, you can set filtering conditions using request parameters such as `filter` to retrieve the records you need. This document explains how to fill in the filtering parameters.
## Filter description

The description and structure of the `filter` parameter are as follows. For more information, refer to search records.
| Parameter Name | Data Type | Description |
| --- | --- | --- |
| filter | filter_info | An object containing filter information. |
| └ conjunction | string | Represents the logical conjunction between conditions, which can be "and" or "or". |
| └ conditions | condition[] | A collection of filtering conditions. |
| └ └ field_name | string | The name of the condition field. |
| └ └ operator | string | The condition operator. Its optional values are: - `is`: is - `isNot`: is not - `contains`: contains - `doesNotContain`: does not contain - `isEmpty`: is empty - `isNotEmpty`: is not empty - `isGreater`: greater than - `isGreaterEqual`: greater than or equal to - `isLess`: less than - `isLessEqual`: less than or equal to - `like`: LIKE operator. Not supported yet - `in`: IN operator. Not supported yet |
| └ └ value | string[] | The value of the condition, which can be a single value or an array of multiple values. Different field types and different operators can have different fillable values. For details, refer to the instructions below. | The structure of filter is as follows:

```json
{
  "filter": {
    "conjunction": "and",
    "conditions": [
      {
        "field_name": "字段1",
        "operator": "is",
        "value": [
          "文本内容"
        ]
      }
    ]
  }
}
```

## filter usage example

Below is a table of employee sales figures. This section provides examples of using the `filter` parameter based on this table.
| Employee Name   | Position           | Sales        |
|-----------------|--------------------|--------------|
| John Smith      | Junior Salesperson | 10,000.0    |
| Emily Johnson   | Junior Salesperson | 15,000.0    |
| Michael Brown   | Junior Salesperson | 20,000.0    |
| Linda Davis     | Senior Salesperson | 30,000.0    |
| James Wilson    | Senior Salesperson | 50,000.0    |
| Jennifer Miller | Sales Manager      | 100,000.0   | - To filter records where the position is "Junior Salesperson" **and** sales are greater than 10000, the example of the filter parameter is shown below:

  ```JSON
  {
    "filter": {
      "conjunction": "and",
      "conditions": [
        {
          "field_name": "Position",
          "operator": "is",
          "value": [
            "Junior Salesperson"
          ]
        },
        {
          "field_name": "Sales",
          "operator": "isGreater",
          "value": [
            "10000.0"
          ]
        }
      ]
    }
  }

- To filter records where the position is "Senior Salesperson" or sales are greater than 20000, the example of the filter parameter is shown below:
  ```json
  {
    "filter": {
      "conjunction": "or",
      "conditions": [
        {
          "field_name": "Position",
          "operator": "is",
          "value": [
            "Senior Salesperson"
          ]
        },
        {
          "field_name": "Sales",
          "operator": "isGreater",
          "value": [
            "20000.0"
          ]
        }
      ]
    }
  }
  ```
- To filter records where the position is "Junior Salesperson" **or** "Senior Salesperson" **and** the sales amount is 10000 **or** 20000, the filter parameter example is as follows:

> **Note:** As shown in the example, currently only one layer of children parameters is supported, and nested use is not supported.
  ```json
  {
    "filter": {
      "conjunction": "and",
      "children": [
        {
          "conjunction": "or",
          "conditions": [
            {
              "field_name": "职位",
              "operator": "is",
              "value": [
                "Senior Salesperson"
              ]
            },
            {
              "field_name": "职位",
              "operator": "is",
              "value": [
                "Junior Salesperson"
              ]
            }
          ]
        },
        {
          "conjunction": "or",
          "conditions": [
            {
              "field_name": "销售额",
              "operator": "is",
              "value": [
                "10000.0"
              ]
            },
            {
              "field_name": "销售额",
              "operator": "is",
              "value": [
                "20000.0"
              ]
            }
          ]
        }
      ]
    }
  }
  ```
## Field value filling instructions
Base supports the following types of fields as filter conditions. Currently, formula or lookup reference field types are not supported as filter conditions.
> When the value is empty [], please ensure to pass the value in the format `"value":[]`, otherwise a missing value error will be reported.

| Field type | Value example | Description | Restrictions |
| --- | --- | --- | --- |
| Multi-line text | ["text content"] | Fill in the corresponding text content | - The list can only have one element or no elements - When the operator is `isEmpty` or `isNotEmpty`, fill in the empty value `[]` |
| Barcode | ["barcode content"] | Fill in the corresponding barcode content | - The list can only have one element or no elements - When the operator is `isEmpty` or `isNotEmpty`, fill in the empty value `[]` |
| Number | ["23.4"] | Fill in the corresponding number in string form | - The list can only have one element or no elements - When the operator is `isEmpty` or `isNotEmpty`, fill in the empty value `[]` |
| Currency | ["23.4"] | Fill in the corresponding number in string form | - The list can only have one element or no elements - When the operator is `isEmpty` or `isNotEmpty`, fill in the empty value `[]` |
| Progress | ["0.34"] | Fill in the corresponding number in string form | - The list can only have one element or no elements - When the operator is `isEmpty` or `isNotEmpty`, fill in the empty value `[]` |
| Rating | ["1"] | Fill in the corresponding number in string form | - The list can only have one element or no elements - When the operator is `isEmpty` or `isNotEmpty`, fill in the empty value `[]` |
| Single choice | ["a","b"] | Fill in the options in the list | The list may contain multiple elements: - When the operator is `is` or `isNot`, fill in one element - When the operator is `isEmpty` or `isNotEmpty`, fill in the empty value `[]` - Other operators can fill in multiple elements |
| Multiple choice | ["a","b"] | An array containing multiple option name strings | The list may contain multiple elements: - When the operator is `is` or `isNot`, fill in one element - When the operator is `isEmpty` or `isNotEmpty`, fill in the empty value `[]` - Other operators can fill in multiple elements |
| Date | ["ExactDate","1702449755000"] | Unix timestamp in milliseconds | The list may contain multiple elements, refer to the detailed instructions below for filling in date fields |
| Checkbox | ["true"] or ["false"] | Fill in the corresponding boolean content | The list can only have one element, the operator only supports `is` |
| Person | ["ou_9a971ded01b4ca66f4798549878abcef"] | Fill in the corresponding user ID. The user ID type must match the type specified by the `user_id_type` parameter in the query records, the default type is open_id | The list may contain multiple elements: - When the operator is `is` or `isNot`, fill in one element - When the operator is `isEmpty` or `isNotEmpty`, fill in the empty value `[]` - Other operators can fill in multiple elements |
| Phone number | ["131xxxx6666"] | Fill in the corresponding phone number | - The list can only have one element or no elements - When the operator is `isEmpty` or `isNotEmpty`, fill in the empty value `[]` |
| Hyperlink | ["link display name"] | Fill in the corresponding hyperlink name | - The list can only have one element or no elements - When the operator is `isEmpty` or `isNotEmpty`, fill in the empty value `[]` |
| Attachment | [] | Only supports `isEmpty` or `isNotEmpty` | Fill in the empty value `[]` |
| Single-link relation | ["recnVYsuqV"] | Fill in the corresponding record ID | The list may contain multiple elements: - When the operator is `is` or `isNot`, fill in one element - When the operator is `isEmpty` or `isNotEmpty`, fill in the empty value `[]` - Other operators can fill in multiple elements |
| Double-link relation | ["recnVYsuqV"] | Fill in the corresponding record ID | The list may contain multiple elements: - When the operator is `is` or `isNot`, fill in one element - When the operator is `isEmpty` or `isNotEmpty`, fill in the empty value `[]` - Other operators can fill in multiple elements |
| Location | ["Tiananmen Square, Dongcheng District, Beijing"] | Fill in the corresponding address | The list can only have one element or no elements, fill in the empty value `[]` when the operator is `isEmpty`or `isNotEmpty` |
| Group | ["oc_cd07f55f14d6f4a4f1b51504e7e97f48"] | Fill in the corresponding group ID | The list may contain multiple elements: - When the operator is `is` or `isNot`, fill in one element - When the operator is `isEmpty` or `isNotEmpty`, fill in the empty value `[]` - Other operators can fill in multiple elements |
| Creation time | ["ExactDate","1702449755000"] | Unix timestamp in milliseconds | The list may contain multiple elements, refer to the detailed instructions below for filling in date fields |
| Last update time | ["ExactDate","1702449755000"] | Unix timestamp in milliseconds | The list may contain multiple elements, refer to the detailed instructions below for filling in date fields |
| Creator | ["ou_9a971ded01b4ca66f4798549878abcef"] | Fill in the corresponding user ID. The user ID type must match the type specified by the `user_id_type` parameter in the query records, the default type is open_id | The list may contain multiple elements: - When the operator is `is` or `isNot`, fill in one element - When the operator is `isEmpty` or `isNotEmpty`, fill in the empty value `[]` - Other operators can fill in multiple elements |
| Modifier | ["ou_9a971ded01b4ca66f4798549878abcef"] | Fill in the corresponding user ID. The user ID type must match the type specified by the `user_id_type` parameter in the query records, the default type is open_id | The list may contain multiple elements: - When the operator is `is` or `isNot`, fill in one element - When the operator is `isEmpty` or `isNotEmpty`, fill in the empty value `[]` - Other operators can fill in multiple elements |
| Auto number | ["1"] | Fill in the corresponding auto number value | - The list can only have one element or no elements - When the operator is `isEmpty` or `isNotEmpty`, fill in the empty value `[]` | ## Date field filling instructions
When filtering dates, the operator only supports five values: `is`, `isEmpty`, `isNotEmpty`, `isGreater`, and `isLess`.

When the operator is `isEmpty` or `isNotEmpty`, the value should be empty: `"value":[]`.

When the operator is `is`, `isGreater`, or `isLess`, refer to the table below to fill in the date field.

  | Optional value elements | Description | Example target value | Notes |
| --- | --- | --- | --- |
| ExactDate | Specific date | ["ExactDate","1702449755000"] | - Requires 2 elements. The second element needs to be the timestamp of the specific date. - Although the second element is a timestamp, it will be converted to the zero point of the day in the document time zone during actual filtering. - For formula date fields, the second element needs to be filled with the date text in yyyy/MM/dd format, such as 2025/01/07. |
| Today | Today | ["Today"] | Requires 1 element |
| Tomorrow | Tomorrow | ["Tomorrow"] | Requires 1 element |
| Yesterday | Yesterday | ["Yesterday"] | Requires 1 element |
| CurrentWeek | This week | ["CurrentWeek"] | - Requires 1 element - The operator only supports `is` |
| LastWeek | Last week | ["LastWeek"] | - Requires 1 element - The operator only supports `is` |
| CurrentMonth | This month | ["CurrentMonth"] | - Requires 1 element - The operator only supports `is` |
| LastMonth | Last month | ["LastMonth"] | - Requires 1 element - The operator only supports `is` |
| TheLastWeek | In the last seven days | ["TheLastWeek"] | - Requires 1 element - The operator only supports is |
| TheNextWeek | In the next seven days | ["TheNextWeek"] | - Requires 1 element - The operator only supports is |
| TheLastMonth | In the last thirty days | ["TheLastMonth"] | - Requires 1 element - The operator only supports is |
| TheNextMonth | In the next thirty days | ["TheNextMonth"] | - Requires 1 element - The operator only supports is | :::
