---
title: "Record filter development guide"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/filter"
service: "bitable-v1"
resource: "filter"
section: "Deprecated Version (Not Recommended)"
updated: "1752650888000"
---

# Record filter development guide
> This document is deprecated. For the latest version of this document，see Record filter guide.

In the ListRecord, you can filter out the records you need through the query parameter filter.
The filtered expression supports the formula grammar writing of multi-dimensional tables, which can be combined by AND or OR. 

The grammar of the expression is currentValue.[field name]. Due to the limitation of the URL support character set, the request parameter needs to be encoded using URL before transmission. For example, the "+" number needs to be encoded as "% 2B".

## Common Filter formula

Logical Name | formula
| --------- | --------------- |
| is empty | `CurrentValue.[order No.] =""` |
| is not empty | `NOT(CurrentValue.[order No.] ="")` |
| equal | `CurrentValue.[order No.] = "003"` |
| Not equal to | `CurrentValue.[order No.] != "003"` |
| Greater than or equal to| `CurrentValue.[price] >= 10.5` |
| Less than or equal to| `CurrentValue.[order No.]  AND (CurrentValue. [Order Number] .contains ("004"), CurrentValue. [Order Date] = TODAY ())

Or
> OR (CurrentValue. [Order Number] .contains ("004"), CurrentValue. [Order Number] .contains ("009")

## Common date formula
Today
> CurrentValue. [Order Date] = TODAY ()

Yesterday
> CurrentValue. [Order Date] = TODAY () -1

Tomorrow
> CurrentValue. [Order Date] = TODAY () % 2B1

This week
> AND (TODAY () - (WEEKDAY (TODAY (), 2) -1)  AND(DATE(YEAR(TODAY()),MONTH(TODAY()),1)<=CurrentValue.[Order Date], CurrentValue.[Order Date]<=DATE(YEAR(TODAY()),MONTH(TODAY())%2B1,0))
