---
title: "Get docs content"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/docs-v1/content/get"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/docs/v1/content"
service: "docs-v1"
resource: "content"
section: "Docs"
rate_limit: "5 per second"
scopes:
  - "docs:document.content:read"
updated: "1755227000000"
---

# Get docs content

You can obtain the docs content. Currently, only upgraded document content in markdown format is supported.

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/docs/v1/content |
| HTTP Method | GET |
| Rate Limit | 5 per second |
| Supported app types | custom,isv |
| Required scopes | `docs:document.content:read` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Query parameters
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `doc_token` | `string` | Yes | The unique identification of the docs. Click here to learn how to get `doc_token` **Example value**: B4EPdAYx8oi8HRxgPQQbM15UcBf **Data validation rules**: - Length range: `22` ～ `27` characters |
| `doc_type` | `string` | Yes | Docs type **Example value**: docx **Optional values are**:  - `docx`: Upgraded Document  |
| `content_type` | `string` | Yes | Content type **Example value**: markdown **Optional values are**:  - `markdown`: Markdown format  |
| `lang` | `string` | No | Specifies the language of the user name when the @user element exists in the docs. Default `zh` **Example value**: zh **Optional values are**:  - `zh`: Chinese - `en`: English - `ja`: Japanese  | ### Request Example
```bash 
Curl --location --request GET https://open.larksuite.com/open-apis/docs/v1/content?doc_token=E5UEdFiudogel8x0l8WbR0abcef&doc_type=docx&content_type=markdown&lang=en '\
--Header 'Authorization: Bearer u-xxx'
#Please replace'u-xxx 'in'Authorization: Bearer u-xxx' with the real access token before calling
```

## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `content` | `string` | Content | ### Response body example

{
    "code": 0,
    "msg": "success",
    "data": {
        "content": "# Markdown Export Example Document\n\n# Text\n\n**Bold**\n\n*Italic*\n\n~~Strikethrough~~\n\nHyperlink: [Lark Open Platform](https%3A%2F%2Fopen.larksuite.com)\n\nInline code: `inline code`\n\nUnderline\n\nBackground color + font color\n\nPlain text with markdown syntax needs to be escaped, such as \\*\\*bold syntax\\*\\*, \\*italic syntax\\*, ~~strikethrough syntax~~, \\[hyperlink syntax\\]\\(https://open\\.larksuite\\.com\\), \\`inline code syntax\\`, etc.\n\nMentionDoc: [Markdown Export Example Document](https://example.larksuite.com/docx/FlYadoUfloTbYcxoJcccEoabcef)\n\nMentionUser: @Zhang San\n\nDate reminder: ⏰2024-04-27 15:30\n\nButton: Follow document updates open hyperlink - google\n\nInline file: \\[Markdown Export Example Document.md\\]\n\n# Bullet\n\n- **Bold**\n\n    - *Italic*\n\n    - ~~Strikethrough~~\n\n- Hyperlink: [Lark Open Platform](https%3A%2F%2Fopen.larksuite.com)\n\n    - Inline code: `inline code`\n\n- Underline\n\n    - Background color + font color\n\n# Ordered\n\n1. **Bold**\n\n    1. *Italic*\n\n    2. ~~Strikethrough~~\n\n2. Hyperlink: [Lark Open Platform](https%3A%2F%2Fopen.larksuite.com)\n\n    1. Inline code: `inline code`\n\n        1. Underline\n\n        2. Background color + font color\n\n# Code\n\n```JavaScript\nfunction greeting() {\n    console.log(\"Hello, World!\");\n}\n```\n\n# Quote\n\n> **Bold**\n> \n> *Italic*\n> \n> ~~Strikethrough~~\n> \n> Hyperlink: [Lark Open Platform](https%3A%2F%2Fopen.larksuite.com)\n> \n> Inline code: `inline code`\n> \n> Underline\n> \n> Background color + font color\n> \n\n# Task\n\n* [ ] ~~This is an incomplete task list [Markdown Export Example Document](https://example.larksuite.com/docx/FlYadoUfloTbYcxoJcccEoabcef)⏰2024-04-27 15:30~~\n\n* [x] ~~This is a completed task list~~\n\n# Divider (Horizontal Rule)\n\n---\n\n# Table\n\n\n\n\n\n**Location**\n\n\n\n\n**Features**\n\n\n\n\n**Cuisine**\n\n\n\n\n**Price**\n\n\n\n\n\n\nShenzhen Old Street Snacks\n\n\n\n\nTraditional Cuisine\n\n\n\n\n- Roast Meat\n\n    - Roast Duck\n\n    - Roast Goose\n\n    - Roast Pork\n\n- Dim Sum\n\n    - Rice Noodle Roll\n\n    - Glutinous Rice Chicken\n\n- Seafood\n\n    - Shark Fin\n\n    - Seafood Congee\n\n\n\n\nMedium\n\n\n\n\n\n\nDameisha Seafood Street\n\n\n\n\nSeafood Market\n\n\n\n\n- Fresh Seafood\n\n    - Lobster\n\n    - Crab\n\n    - Scallops\n\n\n\n\nMedium-High\n\n\n\n\n\n\nNanshan District Seafood Street\n\n\n\n\nSeafood Market\n\n\n\n\n- Crab Roe Bun\n\n    - Crab Roe Soup Dumpling\n\n    - Crab Roe Steamed Bun\n\n- Stir-Fried Snail Rice Noodle\n\n    - Stir-Fried Snail Rice Noodle\n\n    - Stir-Fried Rice Noodle\n\n\n\n\nMedium\n\n\n\n\n\n\nLianhua Mountain Restaurant\n\n\n\n\nMountain Cuisine\n\n\n\n\n- Wild Game\n\n    - Wild Boar Meat\n\n    - Goat Meat\n\n    - Venison\n\n- Farmer's Dish\n\n    - Farmer's Stir-Fry\n\n    - Farmer's Claypot Rice\n\n\n\n\nMedium-High\n\n\n\n\n\n\nShenzhen Huaqiangbei Food Street\n\n\n\n\nFood Street\n\n\n\n\n- Hot Pot\n\n    - Spicy Hot Pot\n\n    - Clear Soup Hot Pot\n\n- Skewered Snacks\n\n    - Spicy Skewered Snacks\n\n    - Sour and Spicy Skewered Snacks\n\n- Barbecue\n\n    - Grilled Fish\n\n    - Grilled Meat\n\n\n\n\nLow to Medium\n\n\n\n\n\n# Grid\n\n\n\n\n**Image One**\n\n\n\n\n**Image Two**\n\n\n\n\n**Image Three**\n\n\n\n\n# Callout\n\n\n\nPlain text with markdown syntax needs to be escaped, such as \\*\\*bold syntax\\*\\*, \\*italic syntax\\*, ~~strikethrough syntax~~, \\[hyperlink syntax\\]\\(https://open\\.larksuite\\.com\\), \\`inline code syntax\\`, etc.\n\n\n\n# SyncedSource\n\n**Bold**\n\n*Italic*\n\n~~Strikethrough~~\n\nHyperlink: [Lark Open Platform](https%3A%2F%2Fopen.larksuite.com)\n\nInline code: `inline code`\n\nUnderline\n\n# SyncedReference\n\n# File\n\n\\[Markdown Export Example Document.md\\]\n\n# Bookmark\n\n[https://open.larksuite.com/]()\n\n# Poll\n\n# Agenda\n\nThis is the content of the agenda\n\n# Sheet\n\n# Bitable\n\n# Chart\n\n# Group Name Card\n\n# Whiteboard\n\n# Widget"
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 500 | 2889901 | hybrid resource expired | Check whether the hybrid resources have expired |
| 403 | 2889902 | no permission | - Confirm whether the current access identity has permission to read or edit documents. Please refer to the following methods to resolve this: - If you are using a `tenant_access_token`, it means the current application does not have permission to read or edit documents. You need to add document permissions for the application through the cloud document webpage by navigating to the top right corner **"..."** -> **"... More"** -> **"Add applications"**. **Note**: Before adding a document application, you need to ensure that the target application has at least one cloud document API permission enabled. Otherwise, you will not be able to search for the target application in the document application window. ![](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/bb60f97ebb402475f2af1d3131d4914f_sLOzoqYRXX.png?height=1992&maxWidth=550&width=3278) - If you are using a `user_access_token`, it means the current user does not have permission to read or edit documents. You need to add document permissions for the current user through the **Share** entry in the top right corner of the cloud document webpage. ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/caceea2ac91c170555194d7a8dc2a317_GfTRc9xLAt.png?height=1992&maxWidth=550&width=3278) For more details on the specific steps or other methods to add permissions, refer to Cloud Document FAQ 3. - Confirm whether the current access identity has permission to read MentionDoc, i.e., @document. - Confirm whether the user in MentionUser, i.e., @user, is currently employed and is a contact of the current access identity. - Confirm whether the current access identity has permission to view and share group cards. - Confirm whether the current access identity has permission to access specific Wiki subdirectories. - Confirm whether the current access identity has permission to view document blocks such as OKR, ISV, Add-Ons, etc. |
| 400 | 2889904 | invalid param | Verify that the parameters passed in are valid |
| 500 | 2889905 | internal error | Service internal error, [Contact customer service](https://applink.larksuite.com/client/helpdesk/open?id=6626260912531570952) |
| 403 | 2889906 | docs deleted | Verify that the document has been deleted |
| 400 | 2889914 | doc token not found | Verify that the document Token is correct |
| 500 | 2889925 | content size exceed limit | Document content over limit (10MB) |
| 500 | 2889980 | operation denied on copying page | Verify that the document is copying page |
