---
title: "Function introduction"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uQjM5YjL0ITO24CNykjN/audit_log/"
section: "security_and_compliance"
updated: "1756885046000"
---

# Function introduction
> **Warning:** It's a beta feature and currently available only to Enterprise Premium plan users. If interested, please contact your Customer Success Manager for application details.

Internal audits of organizations should be conducted in accordance with national laws and regulations. Lark provides a notification feature that allows organizations to send notifications to the audited units or auditees according to clear audit policies and procedures. Organizations should strictly control audit permissions, and only grant the audit administrator permissions to appropriate audit personnel ([Admin | Add administrator roles and assign permissions](https://www.larksuite.com/hc/en-US/articles/360043595213 )) and the API audit permissions to appropriate applications (Apply for API permission).

## Overview
Behavioral audit API is an API opened by the open platform based on the behavior events of Lark users. Users can use this API to connect with custom apps or SIEM tools, and conduct audit and statistical analysis on sensitive usage behaviors of members, so as to ensure the privacy compliance and trace security events.

## What can the behavioral audit log API do?
- This API only provides the behavioral logs of users **(excluding those of administrators) and the audit capability to messages and Docs**
- This API can query the audit logs of members in Lark. Therefore, it is read-only and can't write behavioral audit events
- This API supports all subsets of behavior events in Lark. We will support more behavior events and event information in the future to meet broader audit needs

## Terms
### Behavioral audit log
A behavioral audit log records who did what on an object at what time under what context.

## Restriction strategies
- Retention period: Behavioral audit logs of the last 180 days can be queried
- Query condition restrictions: Query requests with a time span of up to 30 days are supported
