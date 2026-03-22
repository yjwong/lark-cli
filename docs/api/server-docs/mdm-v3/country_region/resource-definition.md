---
title: "Resource Definition"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/mdm-v3/country_region/resource-definition"
service: "mdm-v3"
resource: "country_region"
section: "Feishu Master Data Management"
updated: "1756439308000"
---

#  Resource Definition
Country/region provides general geographical location country/region level data, supports multiple languages ​​and is updated regularly, mainly including administrative codes and names. For example: Mainland China, Taiwan, the United States, etc.
## 字段说明
| Field Code       | Field Name          | Module        | Description       | 
| --------- | --------------- | -------   | ----------- | --------- |
|mdm_code|	MDM Code		|Basic Information	|Unique identifier of professional data, the format is "MDCT+8 digits"|
|name	|Name|	Basic Information		|Name of Country Region|
|alpha_3_code	|	Alpha3 Code|	Basic Information		|	three letter code|
|alpha_2_code	|Alpha2 Code|	Basic Information		|two letter code|
|numeric_code	|Numeric Code|	Basic Information		|numeric code|
|full_name	|Full Name|	Basic Information		|Country Region Full Name|
|global_code	|Global Code|	Basic Information		|International calling code|
|continents|	Continent|Basic Information	|Continent: 1-Asia, 2-Europe, 3-Africa, 4-North America, 5-South America, 6-Oceania, 7-Antarctica|
|status	|Status|	Basic Information		|Status, enumeration value: 1_valid, 0_invalid| ### 数据示例

```json
{
  "alpha_3_code": "AND",
  "alpha_2_code": "AD",
  "numeric_code": "20",
  "name": {
    "value": "Andorra",
    "multilingual_value": {
      "zh-CN": "Andorra"
    },
    "return_language": "en-US"
  },
  "mdm_code": "MDCT00000001",
  "full_name": {
    "value": "Principality of Andorra",
    "multilingual_value": {
      "zh-CN": "Principality of Andorra"
    },
    "return_language": "en-US"
  },
  "global_code": "+376",
  "status": "1",
  "continents": {
    "value": "2",
    "multilingual_name": {
      "zh-CN": "Europe"
    }
  }
}
```
