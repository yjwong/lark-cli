---
title: "Filter User Guide"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter/filter-user-guide"
service: "sheets-v3"
resource: "spreadsheet-sheet-filter"
section: "Docs"
updated: "1647001356000"
---

# Data filter user guide

## Scenarios

This is used to add filter conditions to a sheet.

## Supported APIs

A single sheet can only have one range. You can set a filter condition for each column in the range. Note: The first row in the range is not filtered. Use the following APIs to manage data filter conditions:

1. Obtain a filter: Obtains the filter details for a sheet.
2. Create a filter: Creates a filter.
3. Update a filter: Updates an existing filter.
4. Delete filter: Deletes a filter.

## Filter parameter
### **filtered_out_rows**
Rows filtered out will be hidden (returned when the filter is obtained)

### **range**

The following five types of application scopes are supported (see Overview for details):

1. sheetId: Indicates a whole spreadsheet
2. sheetId!1:2: Indicates a whole row
3. sheetId!A:B: Indicates a whole column
4. sheetId!A1:B2: Indicates a normal range
5. sheetId!A1:C: Omits the end row to extend the range to the last row of the spreadsheet

### **col**

Column letters in the range.

### **filter_type**

Five filter_types are supported. Each filter_type must have a comparison type (compare_type) and value (expected).

1. multiValue: Multivalue filter
2. number: Numerical filter
3. text: Text filter
4. color: Color filter
5. clear: Clear filter conditions of a column

***multiValue***

Multivalue filter

compare_type: Left blank

Expected: The value to be displayed, an array with a length greater than 0 and less than the number of rows in the range. A single value can't exceed 50,000 characters.
```json
{
  "condition": {
    "filter_type": "multiValue",
    "expected": ["100", "200", "300"]
  }
}
```
***number***

Numerical filter

| compare_type   | expected length | Notes        |
| -------------- | ----------- | --------- |
| equal          | 1           | Numerical filter: is equal to   |
| notEqual       | 1           | Numerical filter: is not equal to  |
| greater        | 1           | Numerical filter: greater than   |
| greaterOrEqual | 1           | Numerical filter: greater than or equal to |
| less           | 1           | Numerical filter: less than   |
| lessOrEqual    | 1           | Numerical filter: less than or equal to |
| between        | 2           | Numerical filter: is between   |
| notBetween     | 2           | Numerical filter: is not between  |
```json
{
  "condition": {
    "filter_type": "number",
    "compare_type": "equal",
    "expected": ["100"]
  },
  "condition": {
    "filter_type": "number",
    "compare_type": "between",
    "expected": ["100", "200"]
  }
}
```
***text***

Text filter
Text length can't exceed 1,000 characters

| compare_type     | expected length | Notes        |
| ---------------- | ----------- | --------- |
| beginsWith       | 1           | Text filter: text starts with  |
| notBeginsWith | 1           | Text filter: text starts without |
| endsWith         | 1           | Text filter: text ends with  |
| notEndsWith   | 1           | Text filter: text ends without |
| contains         | 1           | Text filter: text contains   |
| notContains   | 1           | Text filter: text does not contain  |
```json
{
  "condition": {
    "filter_type": "text",
    "compare_type": "contains",
    "expected": ["abc"]
  }
}
```
***color***

Color filter

| compare_type     | expected length | Notes        |
| ------------ | ----------- | ---------- |
| backColor    | 1           | Color filter: cell color |
| foreColor    | 1           | Color filter: text color  |
```json
{
  "condition": {
    "filter_type": "color",
    "compare_type": "backColor",
    "expected": ["#ffffff"]
  }
}
```
***clear***

Clear filter. Clears the filter conditions of a column. No values are specified for compare_type and expected.
```json
{
  "condition": {
    "filter_type": "clear",
  }
}
```
