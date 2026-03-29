---
title: "Overview"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/overview"
service: "attendance-v1"
resource: "overview"
section: "Attendance"
updated: "1737860692000"
---

# Overview
## Attendance group owner
The attendance group owner can modify the schedule of the attendance group and view the attendance data of the attendance group.

If the attendance group owner is assigned the role of attendance administrator by the company administrator, the owner also has the permission scope of the attendance administrator to edit and delete attendance rules.

## Attendance group member
You can set members who need to attend attendance or who don't need to attend attendance by two dimensions: department and employee.

- If a member of the attendance group is added according to the department dimension, when a new employee joins the department, they will automatically join the attendance group.
- If a member of the attendance group is added according to the employee dimension, the employee will not change the attendance group when their superior department is added to another attendance group.

## Attendance group type
There are 3 types of attendance: fixed shift, scheduled shift, and flexible shift.

- Fixed shift: The working time for each member in the attendance group is same. It is suitable for attendance groups with fixed working time or without multiple shifts.
- Scheduled shift: The working time for members in the attendance group are not exactly the same, and can be customized. It is suitable for attendance groups with multiple shifts, such as morning shift and evening shifts.
- Flexible shift: There is no specific shift, and the attendance group members can clock in/out freely during the attendance period, and the working hours are counted according to the attendance period.

## Attendance shift
- Under the fixed shift, you need to arrange shifts from Monday to Sunday and for special days.
- Under the scheduled shift, you need to specify the shift for each member in the attendance group every day.
- Under the flexible shift, you need to set the earliest and the latest time of the day for attendance, and the days of the week when attendance is required.

## Attendance method
There are 2 available attendance methods: GPS attendance and Wi-Fi attendance.

- GPS attendance: It requires latitude and longitude information and attendance location name.
- Wi-Fi attendance: It requires Wi-Fi name and Wi-Fi MAC address.

## Other attendance settings
- Rule settings: Set whether to allow offsite attendance, whether to allow correction and the correction count allowed in one month, and whether to allow attendance on PC.
- Security settings: Set whether to enable face recognition attendance, and under what circumstances to enable face recognition.
- Statistic settings: Set whether the attendance group member can view the statistical data of certain dimensions.
- Overtime settings: Configure rules for calculating overtime hours.
