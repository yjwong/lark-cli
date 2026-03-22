---
title: "Post and comment content format conversion"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/moments-v1/moments-explanation"
service: "moments-v1"
resource: "moments-explanation"
section: "Moments"
updated: "1760520773000"
---

# Post and comment content format conversion
In order to support AT and links in posts or comments, the moment uses a set of custom rich text schemes. After external simplification, users only need to understand the content as an array of different nodes.

1.  In terms of the overall design idea, the content is a two-dimensional array, the array of inner elements represents a paragraph, and the array of paragraphs is combined to form the content

1.  The paragraph array can contain three types of elements, which are represented by objects, and the types are distinguished by the tag attribute

    1.  Text

          ` { "tag": "text", "text": "Some awesome text content." } `

    1.  Href

          `{ "tag": "a", "text": "a link", "href": "https://..." }`

    1.  At

          `{ "tag": "at", "user_id": "ou_xxxxxx" , "user_name": "Abracadabra Cheng" }`

    1.  Hashtag

          `{ "tag": "hashtag", "text": "#ExampleHashtag" }`
        1.  Hashtag is only supported in post, comment does not support Hashtag
        1.  Hashtag needs to meet the rules: start with #, followed by 1-50 characters, no special characters such as punctuation marks and spaces

**For example:**

> Some awesome text content. [a nice link](https://www.larksuite.com/) @Abracadabra Cheng

> Another line of text. #ExampleHashtag

**Convert to:**

[
    [
        {
            "tag":"text",
            "text":"Some awesome text content."
        },
        {
            "tag":"a",
            "text":"a nice link",
            "href":"https://www.larksuite.com/"
        },
        {
            "tag":"at",
            "user_id":"ou_xxxxx",
            "user_name":"Abracadabra Cheng"
        }
    ],
    [
        {
            "tag":"text",
            "text":"Another line of text."
        },
        {
            "tag":"hashtag",
            "text":"#ExampleHashtag"
        }
    ]
]

**For example:**

> Hello

> [A nice link](https://www.larksuite.com/) @Abracadabra Cheng

**Convert to:**

[
    [
        {
            "tag":"text",
            "text":"Hello"
        }
    ],
    [
        {
            "tag":"a",
            "text":"A nick link",
            "href":"https://www.larksuite.com/"
        },
        {
            "tag":"at",
            "user_id":"ou_xxxxx",
            "user_name":"Abracadabra Cheng"
        }
    ]
]
