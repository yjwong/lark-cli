---
title: "Attendance Statistics Development Guide"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_stats_data/attendance-statistic-reference"
service: "attendance-v1"
resource: "user_stats_data"
section: "Attendance"
updated: "1648474780000"
---

# Developer's guide on attendance statistics

Developers can customize API response data via the attendance statistics API to obtain data content. As in the daily statistics and monthly summary page of [Attendance Admin](https://oa.larksuite.com/attendance/manage/statistics/day), you can select fields in the "Report Settings", and the saved page will display the statistics data of the selected fields.
## Call APIs
1. Call the "Query statistics header" API to query all statistics header fields available in current attendance statistics. Developers can select fields as needed.
2. Call the "Query statistics settings" API to query the statistics header field saved in the current developer's account. When the developer calls the "Query statistics data" API, the statistics data returned are those contained in these header fields.
3. Call the "Update statistics settings" API to save or update the statistics header field to query. Developers can use this API to update header fields if they have new statistical fields to focus on or want to remove statistical fields that are not of interest. After update, call the "Query statistics data" API again to get statistics containing new fields or excluding unnecessary fields.
4. Call the "Query statistics data" API to query statistics data.
