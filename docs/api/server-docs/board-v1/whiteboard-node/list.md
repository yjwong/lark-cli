---
title: "list nodes"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/board-v1/whiteboard-node/list"
method: "GET"
api_path: "https://open.larksuite.com/open-apis/board/v1/whiteboards/:whiteboard_id/nodes"
service: "board-v1"
resource: "whiteboard-node"
section: "Docs"
rate_limit: "10 per second"
scopes:
  - "board:whiteboard:node:read"
updated: "1744960859000"
---

# Obtain all nodes of a board

Obtain all nodes of a board

## Request
| Facts |  |
| --- | --- |
| HTTP URL | https://open.larksuite.com/open-apis/board/v1/whiteboards/:whiteboard_id/nodes |
| HTTP Method | GET |
| Rate Limit | 10 per second |
| Supported app types | custom,isv |
| Required scopes | `board:whiteboard:node:read` | ### Request header
| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| Authorization | string | Yes | `tenant_access_token` or `user_access_token` **Value format**: "Bearer `access_token`" **Example value**: "Bearer u-7f1bcd13fc57d46bac21793a18e560" How to choose and get access token | ### Path parameters
| Parameter | Type | Description |
| --- | --- | --- |
| `whiteboard_id` | `string` | The unique identification of the board **Example value**: "Ru8nwrWFOhEmaFbEU2VbPRsHcxb" **Data validation rules**: - Length range: `22` ～ `27` characters | ## Response

### Response body
| Parameter | Type | Description |
| --- | --- | --- |
| `code` | `int` | Error codes, fail if not zero |
| `msg` | `string` | Error descriptions |
| `data` | `\-` | \- |
|   `nodes` | `whiteboard.node[]` | Node information of the board |
|     `id` | `string` | Node ID |
|     `type` | `string` | Node type **Optional values are**:  - `image`: image - `text_shape`: text - `group`: group - `composite_shape`: composite_shape - `svg`: svg - `connector`: connector - `table`: table - `life_line`: life_line - `activation`: activation - `section`: section - `table_uml`: table_uml - `table_er`: table_er - `sticky_note`: sticky_note - `mind_map`: mind_map - `paint`: paint  |
|     `parent_id` | `string` | Parent id |
|     `children` | `string[]` | Children ids |
|     `x` | `number(float)` | The x-axis position of the node relative to the board (coordinates relative to the parent container if there is a parent container), measured in pixels |
|     `y` | `number(float)` | The y-axis position of the node relative to the board (coordinates relative to the parent container if there is a parent container), measured in pixels |
|     `angle` | `number(float)` | Angle of node |
|     `width` | `number(float)` | Width of node, measured in pixels |
|     `height` | `number(float)` | Height of node, measured in pixels |
|     `text` | `text` | Text of node |
|       `text` | `string` | Text content |
|       `font_weight` | `string` | Font weight **Optional values are**:  - `regular`: regular - `bold`: bold  |
|       `font_size` | `int` | Font size |
|       `horizontal_align` | `string` | Horizontal alignment **Optional values are**:  - `left`: left typesetting - `center`: center typesetting - `right`: right typesetting  |
|       `vertical_align` | `string` | Vertical alignment **Optional values are**:  - `top`: top typesetting - `mid`: middle typesetting - `bottom`: bottom typesetting  |
|     `style` | `style` | Style of node |
|       `fill_opacity` | `number(float)` | Fill opacity |
|       `border_style` | `string` | Border style **Optional values are**:  - `solid`: solid border - `none`: none border - `dash`: dash border - `dot`: dot border  |
|       `border_width` | `string` | Border width **Optional values are**:  - `extra_narrow`: extra narrow - `narrow`: narrow - `medium`: medium - `wide`: wide  |
|       `border_opacity` | `number(float)` | Border opacity |
|       `h_flip` | `boolean` | Horizontal flip |
|       `v_flip` | `boolean` | Vertical flip |
|     `image` | `image` | Image properties |
|       `token` | `string` | Image token |
|     `composite_shape` | `composite_shape` | Composite shape properties |
|       `type` | `string` | Subtype of composite_shape **Optional values are**:  - `round_rect2`: round_rect2 - `ellipse`: ellipse - `hexagon`: hexagon - `cylinder`: cylinder - `parallelogram`: parallelogram - `trapezoid`: trapezoidal - `triangle`: triangle - `round_rect`: round_rect - `step`: step - `diamond`: diamond - `rect`: rect - `star`: star - `bubble`: bubble - `pentagon`: pentagon - `forward_arrow`: forward_arrow - `document_shape`: document_shape - `condition_shape`: condition_shape - `cloud`: cloud - `cross`: cross - `step2`: step2 - `predefined_process`: predefined_process - `delay_shape`: delay_shape - `off_page_connector`: off_page_connector - `note_shape`: note_shape - `data_process`: data_process - `data_store`: data_store - `data_store2`: data_store2 - `data_store3`: data_store3 - `star2`: star2 - `star3`: star3 - `star4`: star4 - `actor`: actor - `brace`: brace - `condition_shape2`: condition_shape2 - `double_arrow`: double_arrow - `data_flow_round_rect3`: data_flow_round_rect3 - `rect_bubble`: rect_bubble - `manual_input`: manual_input - `flow_chart_round_rect`: flow_chart_round_rect - `flow_chart_round_rect2`: flow_chart_round_rect2 - `flow_chart_diamond`: flow_chart_diamond - `flow_chart_parallelogram`: flow_chart_parallelogram - `flow_chart_cylinder`: flow_chart_cylinder - `flow_chart_trapezoid`: flow_chart_trapezoid - `flow_chart_hexagon`: flow_chart_hexagon - `data_flow_round_rect`: data_flow_round_rect - `data_flow_ellipse`: data_flow_ellipse - `backward_arrow`: backward_arrow - `brace_reverse`: brace_reverse - `flow_chart_mq`: flow_chart_mq - `horiz_cylinder`: horiz_cylinder - `class_interface`: class_interface - `classifier`: classifier - `circular_ring`: circular_ring - `pie`: pie - `right_triangle`: right_triangle - `octagon`: octagon - `state_start`: state_start - `state_end`: state_end - `state_concurrence`: state_concurrence - `component_shape`: component_shape - `component_shape2`: component_shape2 - `component_interface`: component_interface - `component_required_interface`: component_required_interface - `component_assembly`: component_assembly - `cube`: cube  |
|     `connector` | `connector` | Connector properties |
|       `start_object` | `connector.attached_object` | Start object of connector |
|         `id` | `string` | Node id |
|       `end_object` | `connector.attached_object` | End object of connector |
|         `id` | `string` | Node id |
|       `captions` | `connector.caption` | Text of connector |
|         `data` | `text[]` | Text |
|           `text` | `string` | Text list |
|           `font_weight` | `string` | Font weight **Optional values are**:  - `regular`: regular - `bold`: bold  |
|           `font_size` | `int` | Font size |
|           `horizontal_align` | `string` | Horizontal alignment **Optional values are**:  - `left`: left typesetting - `center`: center typesetting - `right`: right typesetting  |
|           `vertical_align` | `string` | Vertical alignment **Optional values are**:  - `top`: top typesetting - `mid`: middle typesetting - `bottom`: bottom typesetting  |
|     `section` | `section` | Section properties |
|       `title` | `string` | title |
|     `table` | `table` | Table properties |
|       `meta` | `table.meta` | Table meta |
|         `row_num` | `int` | Table row number |
|         `col_num` | `int` | Table column number |
|       `title` | `string` | Table tile |
|       `cells` | `table.cell[]` | cell list |
|         `row_index` | `int` | Row index, start from 1 |
|         `col_index` | `int` | Column index，start from 1 |
|         `merge_info` | `table.cell.merge_info` | Meger info of cell |
|           `row_span` | `int` | the number of consecutive rows to be merged starting from the current row index |
|           `col_span` | `int` | The number of consecutive columns to merge starting from the current column index |
|         `children` | `string[]` | Nodes that in the cell |
|         `text` | `text` | Text of the cell |
|           `text` | `string` | Text content |
|           `font_weight` | `string` | Font weight **Optional values are**:  - `regular`: regular - `bold`: bold  |
|           `font_size` | `int` | Font size |
|           `horizontal_align` | `string` | Horizontal alignment **Optional values are**:  - `left`: left typesetting - `center`: center typesetting - `right`: right typesetting  |
|           `vertical_align` | `string` | Vertical alignment **Optional values are**:  - `top`: top typesetting - `mid`: middle typesetting - `bottom`: bottom typesetting  |
|     `mind_map` | `mind_map` | center |
|       `parent_id` | `string` | mid | ### Response body example

{
    "code": 0,
    "msg": "Success",
    "data": {
        "nodes": [
            {
                "id": "o1:1",
                "type": "composite_shape",
                "parent_id": "o1:1",
                "children": [
                    "o1:1"
                ],
                "x": 100,
                "y": 100,
                "angle": 100,
                "width": 100,
                "height": 100,
                "text": {
                    "text": "Text content",
                    "font_weight": "regular",
                    "font_size": 14,
                    "horizontal_align": "center",
                    "vertical_align": "mid"
                },
                "style": {
                    "fill_opacity": 50,
                    "border_style": "solid",
                    "border_width": "narrow",
                    "border_opacity": 50,
                    "h_flip": false,
                    "v_flip": false
                },
                "image": {
                    "token": "EeSHb3qs9oSBXoxvw33bqtOsczb"
                },
                "composite_shape": {
                    "type": "ellipse"
                },
                "connector": {
                    "start_object": {
                        "id": "o1:1"
                    },
                    "end_object": {
                        "id": "o1:1"
                    },
                    "captions": {
                        "data": [
                            {
                                "text": "Text",
                                "font_weight": "regular",
                                "font_size": 14,
                                "horizontal_align": "center",
                                "vertical_align": "mid"
                            }
                        ]
                    }
                },
                "section": {
                    "title": "Section title"
                },
                "table": {
                    "meta": {
                        "row_num": 3,
                        "col_num": 3
                    },
                    "title": "Table cells",
                    "cells": [
                        {
                            "row_index": 1,
                            "col_index": 1,
                            "merge_info": {
                                "row_span": 2,
                                "col_span": 2
                            },
                            "children": [
                                "o1:1"
                            ],
                            "text": {
                                "text": "Text content",
                                "font_weight": "regular",
                                "font_size": 14,
                                "horizontal_align": "center",
                                "vertical_align": "mid"
                            }
                        }
                    ]
                },
                "mind_map": {
                    "parent_id": "mid"
                }
            }
        ]
    }
}

### Error code
| HTTP status code | Error code | Description | Troubleshooting suggestions |
| --- | --- | --- | --- |
| 400 | 2890001 | invalid format | Confirm whether the incoming parameter format is correct |
| 400 | 2890002 | invalid arg | Confirm whether the parameters passed in are valid |
| 500 | 2891001 | server internal error | server internal error |
| 400 | 2890006 | record missing | record missing |
| 401 | 2890007 | auth failed | auth failed |
| 403 | 2890008 | forbidden | Confirm whether the operator has permissions for the board |
| 429 | 2890009 | too many request | too many request |
