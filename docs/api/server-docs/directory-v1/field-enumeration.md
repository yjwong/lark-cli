---
title: "Field enumeration"
url: "https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/directory-v1/field-enumeration"
service: "directory-v1"
resource: "field-enumeration"
section: "Organization"
updated: "1756439236000"
---

# Query interface field enumeration

## Employee field enum

| enumeration | description |
| --------- | --------------- | 
|`base_info.*` | All basic information of employees**Note**1. Wildcards will obtain all fields, which will have a certain impact on interface performance;2. Only the dependent object `base_info.departments` can be obtained ID, detailed information needs to specify specific fields |
|`base_info.employee_id` | Unique identification of employees within the enterprise | 
|`base_info.name` | Employee's name | 
|`base_info.avatar` | Employee avatar | 
|`base_info.background_image` | Employee custom background image | 
|`base_info.description` | Employee's personalized signature | 
|`base_info.mobile` | Employee’s mobile phone number | 
|`base_info.email` | Employee contact email at work | 
|`base_info.enterprise_email` | Employee’s business email | 
|`base_info.enterprise_email_alias` | Employee’s business email alias | 
|`base_info.gender` | Employee's gender | 
|`base_info.employee_order_in_departments` | Sorting information of employees within their departments | 
|`base_info.department_path_infos` | Department path information of the employee’s department | 
|`base_info.leader_id` | Employee’s immediate supervisor ID | 
|`base_info.dotted_line_leader_ids` | Employee’s dotted superior ID | 
|`base_info.active_status` | Employee account active status | 
|`base_info.is_resigned` | Whether to resign | 
|`base_info.is_primary_admin` | Is it an enterprise super administrator? | 
|`base_info.is_admin` | Is it an ordinary administrator? | 
|`base_info.custom_field_values` | Custom fields | 
|`base_info.resign_time` | Resignation time | 
|`base_info.data_source` | Data Sources | 
|`base_info.geo_name` | Where employee data resides | 
|`base_info.subscription_ids` | List of seat IDs assigned to employees | 
|`base_info.departments.*` | All information about the department to which the employee belongs | 
|`base_info.departments.department_id` | Department ID of the department to which the employee belongs | 
|`base_info.departments.name` | The department name of the department to which the employee belongs | 
|`base_info.departments.parent_department_id` | The parent department ID of the department to which the employee belongs | 
|`base_info.departments.leaders` | Department head information of the department to which the employee belongs | 
|`base_info.departments.has_child` | Whether the department to which the employee belongs has sub-departments |
|`base_info.departments.department_count` | Count of department members and sub-departments of the department to which the employee belongs |
|`base_info.departments.order_weight` | The sorting weight of the department to which the employee belongs under the parent department |
|`base_info.departments.department_path_infos` | Department path information of the department to which the employee belongs |
|`base_info.departments.custom_field_values` | Custom fields for the department to which the employee belongs |
|`base_info.departments.data_source` | The data source of the department to which the employee belongs |
|`work_info.*` | All work information of employees,**Note**1. Wildcards will obtain all fields, which will have a certain impact on interface performance;2. For the dependent object `work_info .job_title`, `work_info.work_place` can only obtain the ID, and detailed information needs to specify specific fields |
|`work_info.work_country_or_region` | Work country/region code |
|`work_info.work_station` | work station |
|`work_info.job_number` | Job number |
|`work_info.extension_number` | extension number |
|`work_info.join_date` | Join date |
|`work_info.employment_type` | Employee type |
|`work_info.staff_status` | Personnel status |
|`work_info.positions` | Position information |
|`work_info.work_place.*` | Work place information |
|`work_info.work_place.place_id` | Work place ID |
|`work_info.work_place.place_name` | Work place name |
|`work_info.work_place.is_enable` | Whether the work place is enabled |
|`work_info.work_place.description` | Work place description |
|`work_info.job_title.*` | Job information |
|`work_info.job_title.job_title_id` | Job ID |
|`work_info.job_title.job_title_name` | Job title |
|`work_info.job_title.is_enable` | Whether the job is enabled |
|`work_info.job_title.description` | Job description |
|`work_info.job_family.*` |  Job family information | 
|`work_info.job_family.job_family_id` | Job family ID | 
|`work_info.job_family.job_family_name` | Name of the job family | 
|`work_info.job_family.is_enabled` | Enablement of the job family | 
|`work_info.job_family.parent_job_family_id` | ID of the parent job family  | 
|`work_info.job_family.description` | Job family description| ## Department field enumeration

| enumeration | description |
| ---------- | --------------- |
|`*` | All information of the department**Note** Wildcards will obtain all fields, which will have a certain impact on interface performance |
|`department_id` | Department ID |
|`name` | Department name |
|`parent_department_id` | Parent department ID |
|`leaders` | Department head information |
|`has_child` | Whether there is a sub-department |
|`department_count` | Count of department members and sub-departments |
|`order_weight` | Sorting weight under the parent department |
|`department_path_infos` | Department path information |
|`custom_field_values` | Custom fields |
|`data_source` | data source |
