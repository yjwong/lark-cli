---
title: "Resource introduction"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/passport-v1/session/usum"
service: "passport-v1"
resource: "session"
section: "Authenticate and Authorize"
updated: "1744168056000"
---

#  Resource introduction
##  Resource definition
It is used to identify the identity, device and login status of the current Lark user.​

##  Field description

  
      
      Name
      Type
      Description
      
  
  

        
        `Authorization`
        
        
        `string`
        

        
         `tenant_access_token`

**Value format**："Bearer `access_token`"

**Example value**："Bearer t-7f1bcd13fc57d46bac21793a18e560"

How to choose and get access token
          
        

        
        `Content-Type`
        
        
        `string`
        

        
        

**Fixed value**："application/json; charset=utf-8"

        
  

        
        `user_id_type`
        
        
        `string`
        

        
               User ID categories

**Example value**："open_id"

**Optional values are**：
- `open_id`：user's open id
- `union_id`：user's union id
- `user_id`：user's user id

**Default value**：`open_id`

**When the value is `user_id`，the following field scopes are required**：
`contact:user.employee_id:readonly`
        
  
    

        
        `user_ids`
        
        
        `string`
        

        
        用户ID​
          
​**Example value**：["47f621ff"]
          
          
**Data validation rules**：
          
Maximum length：`100`

        
  

        
        `code`
        
        
        `int`
        

        
       Error codes, fail if not zero​
        
  
    

        
        `msg`
        
        
        `string`
        

        
        Error descriptions​
        
  
    

        
        `data`
        
        
        `\-`
        

        
         \-`
        
  

	
	  ∟ `mask_sessions`
	
	
	`mask_session[]`
	
	
	User Login Information​
	
    
    

	
	    ∟ `create_time`
	
	
	`string`
	
	
	 Creation time​
	
    
    

	
	    ∟ `terminal_type`
	
	
	`int`
	
	
	Client side type

**Optional values are**：
- `0`：Unknown
- `1`：Personal computer
- `2`：Browser
- `3`：Android phone
- `4`：Apple phone
- `5`：Server level
	
    

	
	    ∟ `user_id`
	
	
	`string`
	
	
	User ID​
	
    

##  Data example​
```json
{​
    "code": 0,​
    "data": {​
        "mask_sessions": [​
            {​
                "create_time": "1644980493",​
                "terminal_type": 2,​
                "user_id": "47f183f1f1"​
            },​
            {​
                "create_time": "1644983127",​
                "terminal_type": 2,​
                "user_id": "47f183ff1"​
            },​
            {​
                "create_time": "1644983127",​
                "terminal_type": 2,​
                "user_id": "47f183ff2"​
            }​
        ]​
    },​
    "msg": ""​
}
```
