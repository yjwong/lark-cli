---
title: "Getting Started with Docs API"
url: "https://open.larksuite.com/document/ukTMukTMukTM/ugzNzUjL4czM14CO3MTN/guide/getting-start"
section: "Deprecated Version (Not Recommended)"
updated: "1695369893000"
---

# Prepare to integrate the Docs API
Before using the Docs API, you should learn about the relevant rules and restrictions so you can more efficiently integrate the Docs API.

## Make sure your app has Docs scopes
To ensure Docs data is not leaked and protect information security, we use the API Scope to restrict app scopes:
- To read Docs, the app must have the scope to view Docs.
- To edit Docs, the app must have the scopes to view and edit Docs.

If an app doesn't have the necessary scopes, an error will be returned if it calls the API. For details, see App scopes.

## API rules
### QPS limits
Each API has a call frequency limit.
- A single app/user can't call the create API more than 60 times per minute.
- A single app/user can't call the batch_update API more than 60 times per minute.
- A single app/user can't call the content API more than 300 times per minute.  
- A single app/user can't call other APIs more than 60 times per minute.  

Different APIs calculate call frequencies separately. For example, your calls to the create API won't affect how many times you can call the batch_update and content APIs.

### Error codes
An error will be reported when requests don't meet specific conditions. If an error is reported because some parameters don't meet the requirements, it means that all the contents of the request fail to be submitted, but the online data won't be affected.

For details, refer to Server-side error codes.
