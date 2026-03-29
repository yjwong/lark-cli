---
title: "Video Conferencing overview"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/video-conferencing-overview"
service: "vc-v1"
resource: "video-conferencing-overview"
section: "Video Conferencing"
updated: "1649822467000"
---

# Video Conference overview

## Business introduction

Video conference (VC) refers to Lark's audiovisual conferencing business which provide users with comprehensive, convenient, and high-quality audiovisual interactive experience that can satisfy the online, real-time communication and collaboration needs of users across different industries. With the video conferencing APIs, you can access various features, including:

- Meeting reservation, such as scheduled conferences, updated conferences, etc.
- Meeting operation, such as inviting participants, setting the host, ending the conference, etc.
- Meeting recording, such as start/stop recording, obtain recording files, etc.
- Meeting report, such as obtaining meeting data report, Top user list, etc.
- Lark Rooms configuration, such as setting background image, setting digital signage, etc.

### Access process

|  | Steps | Introduction |
| --- | --- | --- |
| 1 | Create an app | - If you would like to create a custom app, visit Custom app development processto learn more - If you would like to create a store app, visit Develop and launch third-party appsto learn more |
| 2 | Call the API and perform actions to the video conference | Before calling an API, you first need to obtain the access credentials and enable the corresponding permissions. To learn more, visit Call server-side APIs |
| 3 | Monitor events and be notified of video conference changes | Before you can start monitoring events, you first need to request the corresponding permissions. To learn more, visit Event subscription overview. Please notice only events from scheduled meetings on open platform could be monitored. To learn more, visit Schedule a meeting | ## Resource introduction

The video conference domain is opened with "resources" as the center. The relationship diagram of resources is as follows:

 

![20220411-183322.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/34ba5cc879684f81df1d65b07e867bd5_78St1FOe1C.png?lazyload=true&width=2458&height=2114)

The definition of resoures is as follows:
| Meeting reservation | Users can schedule a meeting, set participants and meeting permissions ahead of time, and access meeting information |
| --- | --- |
| Meeting operation | Users can invite and remove participants and set meeting hosts during a meeting |
| Meeting recording | Users can record a meeting and get the link to the recording after the meeting is over |
| Meeting reports | The meeting report records a tenant's meeting usage in a period of time. The report includes information such as the number of meetings, the length of the meeting, and the number of participants | The following will introduce the fields, methods, and events of each resource in detail.

### Resource: Meeting reservation
View Resource fields and examples

#### Method list
>  Store refers to app store apps, and custom refers to custom apps.
| `Schedule a meeting` `POST` /open-apis/vc/v1/reserves/apply >Schedule a video meeting  | Update meeting scheduling information | `user_access_token` | **✓** | **✓** |
| --- | --- | --- | --- | --- |
| `Update meeting scheduling` `PUT` /open-apis/vc/v1/reserves/:reserve_id >Update the settings of a scheduled video meeting, including its theme and permission settings | Update meeting scheduling information | `user_access_token` | **✓** | **✓** |
| `Delete a scheduled meeting` `DELETE`/open-apis/vc/v1/reserves/:reserve_id >Delete a scheduled video meeting | Update meeting scheduling information | `user_access_token` | **✓** | **✓** |
| `Access scheduling information` `GET`/open-apis/vc/v1/reserves/:reserve_id >Access the details of a scheduled video meeting, including the meeting ID, meeting link, and meeting permissions | Access meeting scheduling information | `user_access_token` | **✓** | **✓** |
| `Access an active meeting` `GET`/open-apis/vc/v1/reserves/:reserve_id/get_active_meeting >Access the details of an ongoing scheduled meeting, including the number of participants and details of each participant | Access meeting scheduling information | `user_access_token` | **✓** | **✓** | ### Resource: Meeting operation

View Resource fields and examples

#### Method list

| `Access meeting details` `GET` /open-apis/vc/v1/meetings/:meeting_id >Access the detailed data of a meeting, including the meeting theme, meeting ID, meeting link, start time, meeting status, and list of participants | Access meeting information | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| --- | --- | --- | --- | --- |
| `Access the list of meetings associated with the meeting number` `GET` /open-apis/vc/v1/meetings/list_by_no >Access the brief information list of the meetings associated with the meeting numbers in a specified time period | Access meeting information | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `Invite participants` `PATCH` /open-apis/vc/v1/meetings/:meeting_id/invite >Invite participants to a meeting | Update meeting information | `user_access_token` | **✓** | **✓** |
| `Remove participants` `POST` /open-apis/vc/v1/meetings/:meeting_id/kickout >Remove participants from a meeting | Update meeting information | `tenant_access_token` | **✓** | **✓** |
| `Set a host` `PATCH` /open-apis/vc/v1/meetings/:meeting_id/set_host >Designate a participant as the host | Update meeting information | `tenant_access_token` `user_access_token` | **✓** | **✓** |
| `End meetings` `PATCH` /open-apis/vc/v1/meetings/:meeting_id/end >End an ongoing meeting | Update meeting information | `user_access_token` | **✓** | **✓** | #### Event list

| `Meeting starts` | When the meeting is started | Access meeting information | vc.meeting.meeting_started_v1 | **✓** | **✓** |
| --- | --- | --- | --- | --- | --- |
| `Meeting ends` | When the meeting is ended | Access meeting information | vc.meeting.meeting_ended_v1 | **✓** | **✓** |
| `Join a meeting` | When a participant joins the meeting | Access meeting information | vc.meeting.join_meeting_v1 | **✓** | **✓** |
| `Leave a meeting` | When a participant leaves the meeting | Access meeting information | vc.meeting.leave_meeting_v1 | **✓** | **✓** |
| `Recording starts` | When the recording starts | Access meeting information | vc.meeting.recording_started_v1 | **✓** | **✓** |
| `Recording stops` | When the recording ends | Access meeting information | vc.meeting.recording_ended_v1 | **✓** | **✓** |
| `Recording completed` | When the recording file is uploaded | Access meeting information | vc.meeting.recording_ready_v1 | **✓** | **✓** |
| `Screen sharing starts` | When screen sharing starts | Access meeting information | vc.meeting.share_started_v1 | **✓** | **✓** |
| `Screen sharing ends` | When screen sharing ends | Access meeting information | vc.meeting.share_ended_v1 | **✓** | **✓** | ### Resource: Meeting recording

View Resource fields and examples

#### Method list

| `Start recording` `PATCH` /open-apis/vc/v1/meetings/:meeting_id/recording/start >Start recording during the meeting | Update meeting recording information | `user_access_token` | **✓** | **✓** |
| --- | --- | --- | --- | --- |
| `Stop recording` `PATCH` /open-apis/vc/v1/meetings/:meeting_id/recording/stop >Stop recording during the meeting | Update meeting recording information | `user_access_token` | **✓** | **✓** |
| `Access recording file` `GET` /open-apis/vc/v1/meetings/:meeting_id/recording >Access the recording file of a meeting | Access meeting recording information | `user_access_token` | **✓** | **✓** |
| `Authorize access to the recording file` `PATCH` /open-apis/vc/v1/meetings/:meeting_id/recording/set_permission >Authorize the access of the meeting recording file to an organization, user, or make it public | Update meeting recording information | `user_access_token` | **✓** | **✓** | ### Resource: Meeting report
View Resource fields and examples

#### Method list

| `Access the meeting report` `GET` /open-apis/vc/v1/reports/get_daily >Access the daily meeting usage report of an organization during a specified time period, including the total number of meetings, total meeting duration, and total number of participants | Access the meeting report | `tenant_access_token` | **✓** | **✓** |
| --- | --- | --- | --- | --- |
| `Access the list of top users` `GET` /open-apis/vc/v1/reports/get_top_user >Access the list of top meeting users of an organization during a specified time period | Access the meeting report | `tenant_access_token` | **✓** | **✓** |
