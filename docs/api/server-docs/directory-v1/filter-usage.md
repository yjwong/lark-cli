---
title: "Filter usage"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/directory-v1/filter-usage"
service: "directory-v1"
resource: "filter-usage"
section: "Organization"
updated: "1756439236000"
---

# Query condition usage

## Request parameters
| Name | Type | Required | Description |
| --- | --- | --- | --- |
| `filter` | `multi_filter_condition` | no | Query conditions |
|   `conditions` | `filter_condition[]` | yes | comparison expression list **Data verification rules**: - Length range: `0` ~ `10` |
|     `field` | `string` | yes | The left value of the filter condition, optional filtering conditions can be found in the  **Supported Query Conditions**  section below. **Example value**: "base_info.mobile" |
|     `operator` | `string` | yes | comparison operator **Example value**: "eq" |
|     `value` | `string` | yes | rvalue of filter condition | ## Request example
```json
  {
     "filter": {
         "conditions": [
            {
                "field": "work_info.mobile",
                "operator": "eq",
                "value": "\"+8613000000001\""
            },
          	{
                "field": "work_info.staff_status",
                "operator": "eq",
                "value": "1"
            },
            {
                "field": "base_info.departments.department_id",
                "operator": "in",
                "value": "[\"77a83513ge4c9f91\"]"
            }
         ]
     }
}
```

## Support query conditions
| Object | The lvalue of the filter condition (field) | lvalue type | Supported operators (operators) | The rvalue of the filter condition (value) | Description |
| --------- | --------------- | ----   | ----   |----------------- | --------- |
|Employees | `base_info.mobile`| string | `eq` | `\"+8613000000001\"` | Filter by mobile phone number |
|Employees | `base_info.mobile`| string | `in` | `[\"+8613000000000\", \"+8613000000002\"]` | Filter by mobile phone number |
|Employee | `base_info.email`| string | `eq` | `\"111@163.com\"` | Filter by email |
|Employee | `base_info.email`| string | `in` | `[\"111@163.com\", \"222@163.com\"]` | Filter by email |
|Employees | `work_info.staff_status`| int | `eq` | `1` | Filter according to the employee staff status filtering (1-5), **need to cooperate with the department ID combination query**, Employment status (1), does not support the status of leaving the company (2), to be employed (3), cancel the employment status (4), to be left (5)  filtering |
|Employees | `base_info.departments.department_id`| string | `eq` | `\"33wqwdsa\"` | Filter by department ID, need to cooperate with personnel status combination query |
|Department | `parent_department_id`| string | `eq` | `\"33wqwdsa\"` | Filter by parent department ID |
|Department | `parent_department_id`| string | `eq` | `[\"33wqwdsa\",\"33wqwdsb\"]` | Filter by parent department ID | ### Rvalue description
- When the lvalue is of type int and the operator is eq, the value is an int string, that is, `"1"`
- When the lvalue is an int type and the operator is in, the value is the serialized value of the int array, that is, `"[1,2]"`
- When the lvalue is of string type and the operator is eq, the value is a string string, that is, `"\"1\""`
- When the lvalue is a string type and the operator is in, the value is the serialized value of the string array, that is, `"[\"1\",\"2\"]"`
