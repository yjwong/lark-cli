---
title: "Function introduction"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/personal_settings-v1/system_status/overview"
service: "personal_settings-v1"
resource: "system_status"
section: "Personal Settings"
updated: "1672726533000"
---

# Function introduction
## 1. Business introduction
**Who should use these APIs?**

For the tenant-dimensional data of API operations in the system state, the user should be **the tenant system state manager**, not the individual developer in the tenant.

**System status field display**

| 名称         | 描述        |
| --------- | --------------- | -------   | ----------- | --------- |
|`title` | See the serial number below ① |
|`i18n_title` | See the serial number below ①，when both i18n_title and title exist, i18n_title has a higher priority |
|`icon_key` | See the serial number below ② |
|`color` | See the serial number below ③ |
|`priority` | The value is not displayed on the client. The smaller the value, the higher the priority of displaying the system status on the client. |
|`sync_setting->is_open_by_default` | See the serial number below ④ |
|`sync_setting->title` | See the serial number below ⑤ |
|`sync_setting->i18n_title` | See the serial number below ⑤，when both i18n_title and title exist, i18n_title has a higher priority |
|`sync_setting->explain` | See the serial number below ⑥ |
|`sync_setting->i18n_explain` | See the serial number below ⑥，when both i18n_explain and explain exist, i18n_explain has a higher priority | ![20221102-154822.jpeg](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/db953a79b29dc2e0ad6bcf27f003afbc_58VW85WAyC.jpeg?lazyload=true&width=696&height=439)

![20221102-154826.jpeg](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/d713c64473854894527df8524a2039b2_jv1jTmKeMO.jpeg?lazyload=true&width=352&height=592)

## 2. Field enumeration
### icon_key
| emoji | emoji_type | emoji | emoji_type | emoji | emoji_type |
| --- | --- | --- | --- | --- | --- |
| ![GeneralDoNotDisturb](https://sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/ark/16794aef3919f3f20d46d0bdf5c89945.png?lazyload=true&width=96&height=96) | GeneralDoNotDisturb | ![GeneralInMeetingBusy](https://sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/ark/f6cd7c731ec50668b488c9abe1608efc.png?lazyload=true&width=96&height=96) | GeneralInMeetingBusy | ![Coffee](https://sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/ark/f6c935685cf07b57fa39473ca6dd0bc7.png?lazyload=true&width=104&height=96) | Coffee |
| ![GeneralBusinessTrip](https://sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/ark/e8e1285751c447e158c1db03645a5392.png?lazyload=true&width=96&height=96) | GeneralBusinessTrip | ![GeneralWorkFromHome](https://sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/ark/619c9d5d48dd945dc71922a3e4d46db7.png?lazyload=true&width=96&height=96) | GeneralWorkFromHome | ![StatusEnjoyLife](https://sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/ark/9ca77c1047e0a768ddd00941eb64f97b.png?lazyload=true&width=96&height=96) | StatusEnjoyLife |
| ![GeneralTravellingCar](https://sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/ark/26d4dbd2ce407bc241955fc443a3a1fe.png?lazyload=true&width=96&height=96) | GeneralTravellingCar | ![StatusBus](https://sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/ark/290f72459efa88a464220d7c2ebda57d.png?lazyload=true&width=96&height=96) | StatusBus | ![StatusInFlight](https://sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/ark/6ce4bfb3a212e3ec82cd64c208dbb0d3.png?lazyload=true&width=96&height=96) | StatusInFlight |
| ![Typing](https://sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/ark/54509a85bf422471d1f6426ea72445de.png?lazyload=true&width=96&height=96) | Typing | ![EatingFood](https://sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/ark/87b0cc5df92bbb51a471294703e65fc9.png?lazyload=true&width=96&height=96) | EatingFood | ![SICK](https://sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/ark/bff2f1050fa6d35b39c64a550b0f02df.png?lazyload=true&width=96&height=96) | SICK |
| ![GeneralSun](https://sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/ark/8d4d5d2e59bbdb8869413b0c67e38cd2.png?lazyload=true&width=96&height=96) | GeneralSun | ![GeneralMoonRest](https://sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/ark/d4cf6c0f8875dce0f42853ee624eae32.png?lazyload=true&width=96&height=96) | GeneralMoonRest | ![StatusReading](https://sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/ark/fc49412c59b174f47fa417cb4a70b739.png?lazyload=true&width=96&height=96) | StatusReading |
| ![Status_PrivateMessage](https://sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/ark/048f8c167cbc3d276442b05f0e694cfb.png?lazyload=true&width=96&height=96) | Status_PrivateMessage | ![StatusFlashOfInspiration](https://sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/ark/69ac9c6e44ffbbaf3324404611510ccb.png?lazyload=true&width=96&height=96) | StatusFlashOfInspiration | ![StatusFlashOfInspiration](https://sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/ark/6b1ac06bc77bb5c6a57c0e133c30f50c.png?lazyload=true&width=96&height=96) | GeneralVacation |
