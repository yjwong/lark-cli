---
title: "FAQs"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/approval-common-problem/approval-events-faq"
service: "approval-v4"
resource: "approval-common-problem"
section: "Approval"
updated: "1676430227000"
---

# FAQs

Frequently Asked Questions When Using Approval Related APIs

## Approval event not received
**Q：What are the possible reasons for subscribing to the approval event, but not receiving it?？**
- If all events are not found, a step may be missing in the subscription step, please follow the steps again.
- If a certain event is not received, you can go to the [developer background] - [log retrieval] to check the event log retrieval. It may be that the callback address cannot be accessed.
- Because the same approval event is ordered in an approval instance, for example, if the approval event is monitored, if the "PENDING" message of "Approval Instance Status Change" of Approval Instance A is received and not returned in time (please check the push cycle and frequency of the event), then the open API will not continue to send other "Approval Instance Status Change" messages of Approval Instance A.
- If a developer is using multiple apps at the same time (such as a test app and an online app), it may be that he is only subscribing to the test app and not subscribing to the online app.
