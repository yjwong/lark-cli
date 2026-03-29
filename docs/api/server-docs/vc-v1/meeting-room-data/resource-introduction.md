---
title: "Resource introduction"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting-room-data/resource-introduction"
service: "vc-v1"
resource: "meeting-room-data"
section: "Video Conferencing"
updated: "1689326174000"
---

# Resource introduction

## Resource

It is used to query the conference data of the tenant within a period of time by page, including: get meeting list details、get participant  list details、get participant meeting quality data、get meeting room reservation data.

## Permission
> **「get meeting list details、get participant  list details、get participant meeting quality data、get meeting room reservation data」** these four Get Apis at least need users who create self-built applications have the Lark backstage-video conferencing-meeting management module permission！

## Resource：Meeting List Detail

### Field

| name | description |
| --- | --- |
| `meeting id` | 9-digit conference number **example value**："123456789" |
| `meeting topic` | meeting topic **example value**："xx's meeting" |
| `meeting organizer` | - Schedule meeting：meeting organizer - Instant meeting（1v1 call included）：meeting/call sponser/meeting room； - Interview meeting：interviewer - Open platform for scheduled meetings：meeting booker **example value**："xx" |
| `department` | department name **example value**："Personnel Department" |
| `user id` | user employ_id **example value**："1a2b3c4d" |
| `job number` | job number **example value**："12345670000" |
| `email` | email **example value**："xx@email.com" |
| `phone number` | phone number **example value**："+86123****8910" |
| `meeting start time` | meeting start time **example value**："2022.07.21 11:01:40 (GMT+08:00)" |
| `meeting end time` | meeting end time **example value**："2022.07.21 12:01:40 (GMT+08:00)" |
| `meeting duration` | meeting duration **example value**："01:00:00" |
| `Number of participants` | Number of participants **example value**："2" |
| `audio` | Whether a microphone/speaker was used in the meeting **example value**："yes" |
| `video` | Whether a camera was used in the meeting **example value**："yes" |
| `share` | Whether a share-screen/magic-share was used in the meeting **example value**："no" |
| `recording` | Whether a recording was used in the meeting **example value**："yes" |
| `phone` | Whether phone was used in the meeting **example value**："no" | ## Resource：participant list detail

### Field

| name | description |
| --- | --- |
| `participant` | participant name **example value**："xx" |
| `department` | departmentname **example value**："Personnel Department" |
| `user id` | user employ_id **example value**："1a2b3c4d" |
| `job number` | job number **example value**："12345670000" |
| `phone number` | phone number **example value**："+86123****8910" |
| `email` | email **example value**："xx@email.com" |
| `device` | device_type **example value**："WINDOWS" |
| `app version` | app version **example value**："5.18.0" |
| `public network ip` | public network ip **example value**："170.40.80.80" |
| `internal network ip` | internal network ip **example value**："192.168.1.1" |
| `proxy` | Whether the participant has enabled proxy service settings **example value**："no" |
| `location` | location **example value**："china mainland" |
| `network type` | LAN、wifi **example value**："wifi" |
| `microphone` | microphone **example value**："consistent with the system" |
| `speaker` | speaker **example value**："consistent with the system" |
| `camera` | camera **example value**："Integrated Camera" |
| `audio` | Whether a microphone/speaker was used in the meeting **example value**："yes" |
| `video` | Whether a camera was used in the meeting **example value**："yes" |
| `share` | Whether a share-screen/magic-share was used in the meeting **example value**："no" |
| `join time` | first join time **example value**："2022.07.21 11:04:52 (GMT+08:00)" |
| `leave time` | last leave time **example value**："2022.07.21 12:04:52 (GMT+08:00)" |
| `duration` | duration **example value**："00:01:58" |
| `leave reason` | leave reason **example value**："The host ends the meeting" | ## Resource：Participant meeting quality data

### Field

| name | description |
| --- | --- |
| `time` | time **example value**："2022.07.21 11:30:00 (GMT+08:00)" |
| `audio-birate(receiver)` | audio birate of receiver **example value**："50kbps" |
| `audio-lag(receiver)` | audio lag of receiver **example value**："5ms" |
| `audio-jitter(receiver)` | audio jitter of receiver **example value**："1ms" |
| `audio-loss avg(receiver)` | audio loss avg of receiver **example value**："2%" |
| `audio-loss max(receiver)` | audio loss max of receiver **example value**："5%" |
| `audio-birate(sender)` | audio birate of sender **example value**："50kbps" |
| `audio-lag(sender)` | audio lag of sender **example value**："5ms" |
| `audio-jitter(sender)` | audio jitter of sender **example value**："1ms" |
| `audio-loss avg(sender)` | audio loss avg of sender **example value**："2%" |
| `audio-loss max(sender)` | audio loss max of sender **example value**："5%" |
| `video-birate(receiver)` | video birate of receiver **example value**："50kbps" |
| `video-lag(receiver)` | video lag of receiver **example value**："5ms" |
| `video-jitter(receiver)` | video jitter of receiver **example value**："1ms" |
| `video-loss avg(receiver)` | video loss avg of receiver **example value**："2%" |
| `video-loss max(receiver)` | videoloss max of receiver **example value**："5%" |
| `video-birate(sender)` | video birate of sender **example value**："50kbps" |
| `video-lag(sender)` | video lag of sender **example value**："5ms" |
| `video-jitter(sender)` | video jitter of sender **example value**："1ms" |
| `video-loss avg(sender)` | video loss avg of sender **example value**："2%" |
| `video-loss max(sender)` | video loss max of sender **example value**："5%" |
| `share-screen-birate(receiver)` | share-screen birate of receiver **example value**："50kbps" |
| `share-screen-lag(receiver)` | share-screen lag of receiver **example value**："5ms" |
| `share-screen-jitter(receiver)` | share-screen jitter of receiver **example value**："1ms" |
| `share-screen-loss avg(receiver)` | share-screenloss avg of receiver **example value**："2%" |
| `share-screen-loss max(receiver)` | share-screenloss max of receiver **example value**："5%" |
| `share-screen-birate(sender)` | share-screenbirate of sender **example value**："50kbps" |
| `share-screen-lag(sender)` | share-screen lag of sender **example value**："5ms" |
| `share-screen-jitter(sender)` | share-screen jitter of sender **example value**："1ms" |
| `share-screen-loss avg(sender)` | share-screen loss avg of sender **example value**："2%" |
| `share-screen-loss max(sender)` | share-screen loss max of sender **example value**："5%" |
| `cpu min usage of client` | cpu min usage of client **example value**："2%" |
| `cpu avg usage of client` | cpu avg usage of client **example value**："5%" |
| `cpu max usage of client` | cpu max usage of client **example value**："10%" |
| `cpu max usage of system` | cpu max usage of system **example value**："20%" | ## Resource：meeting room reservation data

### Field

| name | description |
| --- | --- |
| `meeting room name` | meeting room name **example value**："test rooms" |
| `meeting topic` | meeting topic **example value**："xx's meeting" |
| `booker` | meeting roombooker **example value**："xx" |
| `department of booker` | departmentname **example value**：" Personnel Department" |
| `Number of invitees` | Number of invitees **example value**："10" |
| `Number of accept` | Number of accept **example value**："8" |
| `meeting start time` | meeting start time **example value**："2022.07.21 11:01:40 (GMT+08:00)" |
| `meeting end time` | meeting end time **example value**："2022.07.21 12:01:40 (GMT+08:00)" |
| `meeting duration` | meeting duration **example value**："01:00:00" |
| `status of meeting room reservation ` | status of meeting room reservation optional value： - under approval - success - Released (manual release) - Released (Automatic release after the meeting ends) - Released (Released without sign in) **example value**："success" |
| `check-in device` | check-in device optional value： - checkboard - check QR-code - controller - touch screen - When the device is offline, check in automatically **example value**："checkboard" |
| `check-in status of meeting room ` | check-in status of meeting room Optional value： - No sign-in required (instant booking) - No sign-in required (Sign-in settings are not turned on) - No sign-in required - Timeout not signed in - Sign-in time is not over - No sign-in required (no check-in device installed) - signed in **example value**："No sign-in required" |
| `check-in time of meeting room` | check-in time of meeting room **example value**："2022.07.04 21:00:00 (GMT+08:00)" |
| `releaser` | releaser **example value**："xx" |
| `release time` | release time **example value**："2022.07.04 21:00:00 (GMT+08:00)" |
