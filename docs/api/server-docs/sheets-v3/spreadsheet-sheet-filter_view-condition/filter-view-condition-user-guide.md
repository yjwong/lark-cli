---
title: "Filter View Condition User Guide"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view-condition/filter-view-condition-user-guide"
service: "sheets-v3"
resource: "spreadsheet-sheet-filter_view-condition"
section: "Docs"
updated: "1647001360000"
---

# User guide for using filter conditions in the filter view

## Scenarios

Applies filter conditions to a column in the filter area to display specific data.

## Supported APIs

Five APIs are provided for filter conditions in filter views. A single filter view has only one range. The first row in a range is not filtered. You can set filter conditions for each column in the range.

1. Obtain filter conditions: Obtains the filter conditions of a specified column in the filter view.
2. Create filter conditions: Creates filter conditions for a specified column in the filter view.
3. Update filter conditions: Updates filter conditions of a specified column in the filter view.
4. Delete filter conditions: Deletes filter conditions of a specified column in the filter view.
5. Query filter conditions: Queries all filter conditions in the filter view.

## Filter parameter
### **filter_view_id**
What is a filter_view_id?
A filter_view_id is the unique identifier of a filter view. It is obtained as follows:
1. Use the Obtain filter view API.
2. Open the filter view and find the ID in the link:
![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/d70575214389a7ebec394427787771aa_3D0OWqknnU.png?lazyload=true&width=1148&height=74)

### **range**

The following five types of application scopes are supported (see Overview for details):

1. sheetId: Indicates a whole spreadsheet
2. sheetId!1:2: Indicates a whole row
3. sheetId!A:B: Indicates a whole column
4. sheetId!A1:B2: Indicates a normal range
5. sheetId!A1:C: Omits the end row to extend the range to the last row of the spreadsheet

### **condition_id**

Column letters in the range.

### **filter_type**

Four filter_types are supported. Each filter_type must have a comparison type (compare_type) and value (expected).

1. hiddenValue: Hidden value filter
2. number: Numerical filter
3. text: Text filter
4. color: Color filter

***hiddenValue***

Value to hide

compare_type: Left blank

Expected: The value to be hidden, an array with a length greater than 0 and less than the number of rows in the range. A single value can't exceed 50,000 characters.
```json
{
    "filter_type": "hiddenValue",
    "expected": ["100", "200", "300"]
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
    "filter_type": "number",
    "compare_type": "between",
    "expected": ["100", "200"]
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
    "filter_type": "text",
    "compare_type": "contains",
    "expected": ["abc"]
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
    "filter_type": "color",
    "compare_type": "backColor",
    "expected": ["#ffffff"]
}
```
