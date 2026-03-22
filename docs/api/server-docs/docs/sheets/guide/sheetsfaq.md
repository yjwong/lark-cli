---
title: "sheets-faq"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/guide/sheets-faq"
section: "Docs"
updated: "1647001326000"
---

#  Sheets FAQ
##  Explanation of nouns
####  DateTimeRenderOption
Determines how dates should be rendered in the output.
By default, instructs date, time, datetime, and duration fields to be output as doubles in "serial number" format, as popularized by Lotus 1-2-3. The whole number portion of the value (left of the decimal) counts the days since December 30th 1899. The fractional portion (right of the decimal) counts the time as a fraction of the day. For example, January 1st 1900 at noon would be 2.5, 2 because it's 2 days after December 30st 1899, and .5 because noon is half a day. February 1st 1900 at 3pm would be 33.625. This correctly treats the year 1900 as not a leap year.

| Enums         |                |
| --------- | --------------- | -------   | ----------- | --------- | 
|FormattedString | Instructs date, time, datetime, and duration fields to be output as strings in their given number format (which is dependent on the spreadsheet locale). | ## Interface current limiting
| Interface                                                         | Current Limit                                                       |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| Create spreadsheet | 20 times/minute per application per tenant                   |
| Obtain spreadsheet meta data | 100 times/sec per application per tenant                     |
| Update spreadsheet properties | 100 times/sec per application per tenant；single document can only be called serially |
| Operate sheets | 100 times/sec per application per tenant；single document can only be called serially |
| Update sheet properties | 100 times/sec per application per tenant；single document can only be called serially |
| Move dimension | 100 times/minute per application per tenant                  |
| Prepend data | 100 times/sec per application per tenant；single document can only be called serially |
| Append data  | 100 times/sec per application per tenant；single document can only be called serially |
| Insert rows/columns | 100 times/sec per application per tenant；single document can only be called serially |
| Add rows/columns | 100 times/sec per application per tenant；single document can only be called serially |
| Update rows/columns | 100 times/sec per application per tenant；single document can only be called serially |
| Delete rows/columns | 100 times/sec per application per tenant；single document can only be called serially |
| Reading a single range | 100 times/sec per application per tenant                     |
| Reading multiple ranges | 100 times/sec per application per tenant                     |
| Write data to a single range | 100 times/sec per application per tenant；single document can only be called serially |
| Write data to multiple ranges | 100 times/sec per application per tenant；single document can only be called serially |
| Set cell style | 100 times/sec per application per tenant；single document can only be called serially |
| Batch set cell style | 100 times/sec per application per tenant；single document can only be called serially |
| Merge cells  | 100 times/sec per application per tenant；single document can only be called serially |
| Split cells  | 100 times/sec per application per tenant；single document can only be called serially |
| Write images | 100 times/sec per application per tenant；single document can only be called serially |
| Find cells | 100 times/minute per application per tenant                  |
| Replace cells | 20 times/minute per application per tenant                   |
| Create conditional formatting rule | 100 times/sec per application per tenant；single document can only be called serially |
| Obtain conditional formatting rules | 100 times/sec per application per tenant；single document can only be called serially |
| Update conditional formatting rules | 100 times/sec per application per tenant；single document can only be called serially |
| Remove conditional formatting rules | 100 times/sec per application per tenant；single document can only be called serially |
| Add locked cells | 100 times/sec per application per tenant；single document can only be called serially |
| Retrieve protection scopes | 100 times/sec per application per tenant；single document can only be called serially |
| Modify protection scopes | 100 times/sec per application per tenant；single document can only be called serially |
| Delete protection scopes | 100 times/sec per application per tenant；single document can only be called serially |
| Set drop-down list | 100 times/sec per application per tenant；single document can only be called serially |
| Remove the drop-down list in the specific range | 100 times/sec per application per tenant；single document can only be called serially |
| Update drop-down list settings | 100 times/sec per application per tenant；single document can only be called serially |
| Query drop-down list in the specific range | 100 times/sec per application per tenant；single document can only be called serially |
| Get Filter | 100 times/minute per application per tenant                  |
| Create Filter | 20 times/minute per application per tenant                   |
| Update filter | 20 times/minute per application per tenant                   |
| Delete filter | 100 times/minute per application per tenant                  |
| Creating a filter view | 100 times/minute per application per tenant                  |
| Obtain filtere view | 100 times/minute per application per tenant                  |
| Query filter view | 100 times/minute per application per tenant                  |
| Update filter view | 100 times/minute per application per tenant                  |
| Delete filter view | 100 times/minute per application per tenant                  |
| Creating filter condition | 100 times/minute per application per tenant                  |
| Obtain filter condition | 100 times/minute per application per tenant                  |
| Query filter condition | 100 times/minute per application per tenant                  |
| Update filter criteria | 100 times/minute per application per tenant                  |
| Delete filter condition | 100 times/minute per application per tenant                  |
| Create floating image | 100 times/minute per application per tenant                  |
| Obtain floating image | 100 times/minute per application per tenant                  |
| Query floating image | 100 times/minute per application per tenant                  |
| Update floating image | 100 times/minute per application per tenant                  |
| Delete floating image | 100 times/minute per application per tenant                  |
## Spreadsheet Limitation
| Restricted items                                | Restricted values                                            |
| ----------------------------------------------- | ------------------------------------------------------------ |
| The total number of cells in a single worksheet | Less than or equal to 2,000,000 (including empty rows and empty columns) |
| Number of sheets in the spreadsheet | Less than or equal to 300                                             |
| Number of spreadsheets in the document | Less than or equal to 1500 |
| Character limit for a single cell | Less than or equal to 45000 characters |
| Number of columns in a single worksheet | Less than or equal to 13000 columns |
