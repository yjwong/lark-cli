---
title: "Conditional Formatting User Guide"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/conditionformat/condition-format-guide"
section: "Docs"
updated: "1647001341000"
---

# Conditional Formatting User Guide

## Scenarios

Conditional formatting is used to change the appearance of a cell based on specified conditions.

## Supported APIs

Conditional formatting currently uses 4 APIs. You can set up to 20 conditional formats for a single sheet.

1. Get conditional formatting: Obtains detailed conditional formatting information for a sheet. One operation can obtain information for up to 10 sheets.
2. Create conditional formatting: Batch sets conditional formatting. One operation can set up to 10 conditional formats.
3. Update conditional formatting: Batch updates conditional formatting. One operation can update up to 10 conditional formats.
4. Delete conditional formatting: Deletes conditional formatting. One operation can delete up to 10 conditional formats.

## Conditional formatting restrictions

### **ranges**

The following five types of application scopes are supported:

1. sheetId: Indicates a whole sheet
2. sheetId!1:2: Indicates a whole row
3. sheetId!A:B: Indicates a whole column
4. sheetId!A1:B2: Indicates a normal range
5. sheetId!A1:C: Omits the end row to extend the range to the last row of the sheet

### **style**

The following style parameters are supported and all are optional and are not set if no value is provided. However, a style must have at least one set parameter.

1. font: bold, type: bool; italic, type: bool
2. text_decoration: type: int, values: 0: default, 1: underline, 2: strikethrough, 3: underline and strikethrough
3. fore_color: The text color, type: string
4. back_color: The background color, type: string

### **rule_type & attrs**

- rule_type, seven types in total. Four types do not require attrs parameters: ***containsBlanks (contains blanks), notContainsBlanks (does not contain blanks), duplicateValues (duplicate values), and uniqueValues (unique values)***. Three types are restricted by attrs parameters as described below: ***cellIs (limited value range), containsText (contains text content), and timePeriod (dates)***.
  - formula: string array, only required when rule_type is ***cellIs***. Two elements must be specified when operator is between or notBetween. Otherwise, one element is required. Values are specified by the user and can be numeric (**"1"**) or text (**"\"a\""**).
  - text: string, only required when rule_type is ***containsText***. Values are specified by the user.
  - timeperiod: string, only required when rule_type is ***timePeriod***. The operator must be "is", and the time_period field can be yesterday, today, tomorrow, or last7Days.
- The user can enter content of up to 1,000 characters.
- attrs parameter array, only one parameter is required in the array.

***Three rule_types are restricted by attrs parameters as follows:***

***cellIs (limited value range)***

In attrs parameter, formula is an array. Here, formula1 and formula2 indicate elements in the formula array.

| operator           | formula1 | formula2 | Remarks          |
| ------------------ | -------- | -------- | ----------- |
| equal              | Required       |          | Limited value range: is equal to    |
| notEqual              | Required       |          | Limited value range: not equal to    |
| greaterThan              | Required       |          | Limited value range: greater than    |
| greaterThanOrEqual              | Required       |          | Limited value range: greater than or equal to    |
| lessThan              | Required       |          | Limited value range: less than    |
| lessThanOrEqual              | Required       |          | Limited value range: less than or equal to    |
| between              | Required       |          | Limited value range: is between    |
| notBetween              | Required       |          | Limited value range: is not between    |
```json
{
  "condition_format": {
    "rule_type": "cellIs",
    "attrs": [
      {
        "operator": "equal",
        "formula": [
              "\"aaaaa\""  //Text must be enclosed in quotation marks ("")
        ]
      }
    ]
  },
  "condition_format": {
    "rule_type": "cellIs",
    "attrs": [
      {
        "operator": "between",
        "formula": [
              "1",
              "10"
        ]
      }
    ]
  }
}
```
***containsText (contains text content)***

| operator     | text | Remarks           |
| ------------ | ---- | ------------ |
| containsText | Required   | Contains the following content: text contains  |
| notContains | Required   | Contains the following content: text does not contain  |
| is | Required   | Contains the following content: text is  |
| beginsWith | Required   | Contains the following content: begins with |
| endsWith | Required   | Contains the following content: ends with  |
```json
{
  "condition_format": {
    "rule_type": "containsText",
    "attrs": [
      {
        "operator": "is",
        "text": "******"
      }
    ]
  }
}
```
***timePeriod (dates)***

| operator | timePeriod | Remarks       |
| -------- | ---------- | -------- |
| is       | yesterday  | Date is: yesterday   |
| is       | today  | Date is: today   |
| is       | tomorrow   | Date is: tomorrow   |
| Is       | last7Days  | Date is: within the last 7 days |
```json
{
  "condition_format": {
    "rule_type": "timePeriod",
    "attrs": [
      {
        "operator": "is",
        "time_period": "today"
      }
    ]
  }
}
```
