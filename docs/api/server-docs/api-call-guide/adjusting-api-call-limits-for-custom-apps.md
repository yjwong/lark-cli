---
title: "Adjusting API call limits for custom apps"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/api-call-guide/api-billing"
section: "API Call Guide"
updated: "1741064486000"
---

# Adjusting API call limits for custom apps

The open platform is making adjustments to improve service stability and user experience. Starting from January 8, 2025, the server-side API will be categorized into **Unlimited API**, **Basic API**, and **Advanced API** categories, with different billing rules based on these categories.

## Adjustment details

The server-side API will be divided into **Unlimited API**, **Basic API**, and **Advanced API** categories, with different billing rules for each type.

### Unlimited API

API calls for the following services by enterprise custom apps will not be metered or charged and can be used for free.
If historical versions of free APIs exist, calls to those historical versions will also not be metered or charged.
- Authentication and Authorization
- Event Subscription
- Contacts
- AI

### Basic API

Basic APIs refer to APIs other than unlimited APIs and advanced APIs. Calls to these APIs by enterprise custom apps will be metered. The total number of API calls allowed varies by [Lark version](https://www.larksuite.com/zh_cn/plans?from=footer), detailed as follows.

| Lark version | API call quota description |
| --- | --- |
| Starter | The total number of basic API calls for all enterprise self-built applications under a single tenant is limited to **10,000 times per month**, which resets on the 1st of each calendar month. > The total limit of basic API calls for self-built applications for both uncertified and certified enterprises in the Standard Edition is 10,000 times per month. |
| - Basic - PRO - Enterprise | Unlimited API call quota. | For example, if an enterprise uses the Starter edition (basic API call limit is 10,000 times per month), and custom app A makes 300 basic API calls and custom app B makes 1,000 basic API calls in the current month, the total number of basic API calls for that enterprise's custom appss is 1,300 times. The remaining basic API call quota for the month would be 8,700 times.

### Advanced API

Advanced APIs refer to APIs with relatively complex functionalities and service costs. If an enterprise custom app needs to call a advanced API, it must first [contact us](https://www.larksuite.com/global/salessupport?tracking_code=701TL00000HcrRPYAZ&lang=en-US) to purchase a advanced API quota package, which provides quota. Calling a advanced API will consume quota. Different APIs consume different amounts of quota per price unit as detailed in the table below. If you have any questions, you can [contact support](https://applink.larksuite.com/client/web_app/open?appId=cli_a4517c8461f8100a&mode=sidebar&channel=Boss+Openapi+Call+Limit+Paid).

> **Warning:** - The quota package can only be purchased by Lark Basic, PRO, or Enterprise. Please upgrade to the appropriate version before purchasing the package. For version introduction, refer to [Lark Versions](https://www.larksuite.com/zh_cn/plans?from=footer).
> - Quota packages can be stacked for purchase, and each package is valid for 1 year from the time of purchase.
> - As business upgrades iterate, advanced APIs will continue to increase.

| API Name | Price unit | Quota consumption |
| --- | --- | --- |
| Recognize resume information in documents | page | 6 |
| Streaming speech recognition | s | 1 |
| Basic image recognition | 1,000 calls | 1,000 | For example, calling the Basic image recognition interface 1,000 times will deduct 1,000 benefit points.

## Effective scope

API metering and billing only apply to enterprise custom apps. There are currently no restrictions on store apps calling APIs.

## Effective dates

- **From January 8, 2025, API usage can be viewed, but usage control is not enforced**

	- In the standard version of Lark, the total limit of **Basic API** calls for all custom apps under a single tenant will be adjusted to **10,000 calls/month**, but exceeding this limit is allowed.
	- All versions of Lark will allow viewing **API calls** on the [Admin](https://www.larksuite.com/admin) > **Billing** > **My Quota** page (only basic API call counts will be recorded).

- **From March 3, 2025, the basic API call limit will take effect**

	For custom apps on the standard version of Lark, once the monthly **Basic API** call limit is exceeded, further **Basic API** calls will fail and return error code `99991403`.
    
- **From April 1, 2025, advanced APIs will begin to consume quota points**
    
	**Advanced API** will start to consume points from the measurement package. Once the points are exhausted, developers of custom apps will not be able to make further **Advanced API** calls, and calls will fail, returning error code `99991406`.

  :::warning
  **Advanced API** control will be gradually enforced based on actual business needs.
  :::
  
## Viewing API call quota

### Viewing API call quota via the Admin

Enterprise administrators can log into the [Admin](https://www.larksuite.com/admin) and view the **API calls** on the **Billing** > **My Quota** page.

> **Note:** **API calls** are used to track **Basic API** usage, not other API types. All versions of Lark can view the total API call count used. In the standard version of Lark, the API call count limit is displayed as **10,000**; for other versions, it is displayed as **Unlimited**.

![](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/a6d74d5d4432bd5792bfb0dc9e90fd11_J45Wlg4wXV.png?height=1394&lazyload=true&maxWidth=600&width=2882)

### Receiving over-usage alerts

- **Basic API**: When the total **Basic API** call count of all custom apps reaches 90% or 100% of the limit, the platform will send alert notifications to the enterprise administrator and the developers of custom apps.

| Notification method | Description |
| --- | --- |
| Admin console | Enterprise administrators can view API call usage information in the **Quota alert** section on the homepage when logged into the [Admin](https://www.larksuite.com/admin). |
| Lark client Bot notification | The platform will send card message notifications via the Lark client: - When the **Basic API** call usage reaches 90% of the limit, the card message notification example is as follows. ![](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/53b6aa01d474743ad14b81167048c17b_wWoTyFpF7u.png?height=1166&lazyload=true&maxWidth=350&width=2122) - When the **Basic API** call usage reaches 100% of the limit, the card message notification example is as follows. ![](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/4d15a61061063d69a6443b4b912cd5e0_sZWAShxm4d.png?height=1118&lazyload=true&maxWidth=350&width=2342) | - **Advanced API**: When the advanced API measurement package usage reaches 90% or 100%, the platform will send alert notifications to the developers of custom apps, reminding them to promptly purchase more advanced API measurement packages. Additionally, on the 1st of each calendar month, the developer assistant will send a consumption report of the advanced API measurement package to the developers of custom apps.

  :::warning
  Currently, it is not supported to view the detailed consumption records of the advanced API measurement package on the platform. Please pay attention to the alert notifications sent by the Lark client bot.
  :::

## Increasing API call limits

The total limit of basic API calls can be lifted by upgrading the version of Lark; advanced API rights points can be replenished by purchasing measurement packages.

### Upgrading Basic API call limits

When an enterprise's **Basic API** call usage reaches the limit, custom apps will not be able to make further **Basic API** calls, and attempts will fail with the following error information:

- HTTP Status Code: `429`
- Error Code: `99991403`
- Error Message: `This month's API call quota has been exceeded`

You can contact the enterprise administrator to upgrade the version of Lark, which will remove the limit on **Basic API** calls. For more information about Lark versions and upgrade entry, refer to [Lark versions](https://www.larksuite.com/en_us/plans?from=footer). If you have any questions, you can [contact support](https://applink.larksuite.com/client/web_app/open?appId=cli_a4517c8461f8100a&mode=sidebar&channel=Boss+Openapi+Call+Limit+Paid).

### Purchasing advanced API measurement packs

When the advanced API measurement package is exhausted, custom apps will not be able to make further **Advanced API** calls, and attempts will fail with the following error information:

- HTTP Status Code: `429`
- Error Code: `99991406`
- Error Message: `Advanced API package has been depleted.`

Additionally, advanced API measurement packages are valid for 1 year from the date of purchase. After this period, they will no longer be usable. If you need to purchase advanced API measurement packages, please [contact us](https://www.larksuite.com/global/salessupport?tracking_code=701TL00000HcrRPYAZ&lang=en-US). If you have any questions, you can [contact support](https://applink.larksuite.com/client/web_app/open?appId=cli_a4517c8461f8100a&mode=sidebar&channel=Boss+Openapi+Call+Limit+Paid).

## FAQ

### Do 4xx/5xx errors returned from API calls count towards the enterprise's API measurement?

No, they do not.

### Do changes in OpenAPI metering and charging rules take effect immediately?

No, changes to API metering and charging rules take effect the following month.

### Does posting messages in a group with a custom bot via webhook consume API metering credits?

No. Posting messages in a group using custom bots via webhook does not consume API metering credits.
