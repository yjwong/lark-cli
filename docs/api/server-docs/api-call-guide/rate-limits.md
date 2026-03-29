---
title: "Rate limits"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUzN04SN3QjL1cDN"
section: "API Call Guide"
updated: "1699602849000"
---

# Rate limits

Lark Open Platform has set different levels of frequency control strategies for different OpenAPIs to ensure system stability and provide developers with optimal performance and a premium development experience. The specific limitations for different frequency control levels are listed in this article, along with instructions on how to handle frequency control limitations when calling OpenAPI.

## Response for limitations

When the limit threshold is exceeded, Lark will respond with an HTTP status code 429, which restricts further operations from the client for a specific duration. The suggested waiting time will be provided in the response header of the failed request.

In case of limitations, the HTTP response will resemble the following examples.

```http
HTTP/1.1 429 Too Many Requests
Content-Type: application/json
x-ogw-ratelimit-limit: 100 //Window period limit, unit: second
x-ogw-ratelimit-reset: 52 //Recovery period for limit, unit: second

{
    "code": 99991400,
    "msg": "request trigger frequency limit"
}
```

For some older versions of OpenAPIs, you may receive a slightly different HTTP response (HTTP status code 400). An example is shown below:

```
HTTP/1.1 400 Bad Request
Content-Type: application/json
x-ogw-ratelimit-limit: 100 //Window period limit, unit: second
x-ogw-ratelimit-reset: 52 //Recovery period for limit, unit: second

{
    "code": 99991400,
    "msg": "request trigger frequency limit"
}
```

## Handling limitation scenarios

When handling errors, you can use the HTTP 429 error code to identify the occurrence of rate limiting. The response header of the failed request includes `x-ogw-ratelimit-reset`,  which can be used to delay the subsequent request and effectively lift the rate limit.

Follow these steps to handle rate limiting:

1. Wait for the specified number of seconds indicated in `x-ogw-ratelimit-reset`.
2. Retry the request.
3. If the request fails again with an HTTP 429 error code, the rate limit is still in effect. Continue to delay and retry the request using the recommended x-ogw-ratelimit-reset value until the request succeeds.

## Frequency control strategy levels

In general, frequency control strategy levels are applied on a per-API, per-application, and per-tenant basis. Different operations on the same resource may have varying frequency control levels depending on the scenario. For instance, write interfaces may have a slightly lower frequency control level compared to read interfaces, ensuring the system's overall reliability.

> **Note:** - Custom apps have different frequency control levels depending on the package provided by the associated organization. Store apps, being independent of the hosting organization, follow their own frequency control policies.
> - Custom bots have different frequency control limits: 100 times per minute and 5 times per second.
> - Most OpenAPIs are categorized under standard frequency control strategy levels. However, certain OpenAPIs are more complex and fall outside the scope of standard frequency control strategies. In such cases, please contact technical support to obtain specific frequency control strategies.
> - Exceeding the maximum request rate (QPS or QPM) of different time periods (e.g., 1000 times/minute & 50 times/second) will trigger rate limiting.
> - The specific limitation values for each frequency control strategy level are subject to change. Any updates will be communicated through the Lark Open Platform changelog.

  
    
      Level
      Description
      Custom app limit（Basic Edition）
      Custom app limit (Business Edition)
      Store app limit
    
  
  
    
      1
      The current API has a limit of 10 times per application and tenant per minute
      10 times/min
      10 times/min
      10 times/min
    
    
      2
      The current API has a limit of 20 times per application and tenant per minute
      20 times/min
      20 times/min
      20 times/min
    
    
      3
      The current API has a limit of 100 times per application and tenant per minute
      100 times/min
      100 times/min
      100 times/min
    
    
      4
      The current API has a limit of 1000 times per minute and 50 times per second per application and tenant
      1000 times/min & 50 times/sec
      1000 times/min & 50 times/sec
      1000 times/min & 50 times/sec
    
    
      5
      The current API has a limit of 1 time per second per application and tenant
      1 time/sec
      1 time/sec
      1 time/sec
    
    
      6
      The current API has a limit of 5 times per second per application and tenant
      5 times/sec
      5 times/sec
      5 times/sec
    
    
      7
      The current API has a limit of 10 times per second per application and tenant
      10 times/sec
      10 times/sec
      10 times/sec
    
    
      8
      The current API has a limit of 20 times per second per application and tenant
      20 times/sec
      20 times/sec
      20 times/sec
    
    
      9
      The current API has a limit of 50 times per second per application and tenant
      50 times/sec
      50 times/sec
      50 times/sec
    
    
      10
      The current API has a limit of 50 times per second per application and tenant
      50 times/sec
      100 times/sec
      50 times/sec
    
    
      11
      The current API has a limit of 100 times per second per application and tenant
      100 times/sec
      100 times/sec
      100 times/sec
    
    
      Special Frequency Control
      Non-standard frequency control. Please click on "Technical Support" in the bottom right corner to learn about specific configurations.
      
      
      
    
  

## How to improve API frequency control for apps and organizations

At present, Lark Open Platform does not provide a self-adjustment feature for API interface frequency control. However, if you require a temporary increase in frequency control for scenarios such as data migration or large-scale activities, you can contact your designated CSM (customer success manager) to submit a request.

> **Note:** The frequency control feature is currently not available for
> Message, Group, Base businesses.
