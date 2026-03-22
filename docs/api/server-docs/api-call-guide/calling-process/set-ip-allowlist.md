---
title: "Set IP allowlist"
url: "https://open.larksuite.com/document/ukTMukTMukTM/ucTMxYjL3ETM24yNxEjN"
section: "API Call Guide"
updated: "1699602836000"
---

# Set IP allowlist

## Introduction

In order to improve the security of app access, developers can set the IP allowlist that are allowed to access the Lark server APIs. Requests from IP addresses that are not in the allowlist will be rejected and error messages will be returned. The error code is  `99991401`. For details about this error code, see Common error codes.

> **Note:** - IP allowlist does not take effect on access_token related interfaces.
> - It takes effect immediately after the IP allowlist is configured.
> - Multiple IP addresses can be configured.
> - The IP allowlist function is disabled by default. It needs to be manually configured to be enabled.

## Restrictions

- IP addresses can only be IPv4 addresses, not IPv6 addresses, and must be public IP addresses.
   
   To learn how to obtain a public IP address, you can refer to the troubleshooting suggestions for error code `99991401` in the Common error codes.
   
- Only a single IP address can be configured, not an IP segment. For example, you cannot configure `172.170.0.0`, which means setting the `172.170.x.x` range as an IP whitelist.

## Procedure

1. Log in to the [Developer Console](https://open.larksuite.com/app).
2. In the app list, click the specified app to enter the app details page.
3. Click **Security Settings** and set it to **IP allowlist** section.

     ![](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/d784660e694a6aebd54b6b865f4fe614_BzIrJAK10G.png?height=408&lazyload=true&maxWidth=650&width=1280)
     
4. Fill in a valid IPv4 address and click **Add**.
