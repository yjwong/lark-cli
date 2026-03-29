---
title: "data structure"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/board-v1/data-structure"
service: "board-v1"
resource: "data-structure"
section: "Docs"
updated: "1744960855000"
---

#  Data structure
This document explains the data structure of Board node.

## Node type
Board node types can be divided into container types and non-container types. The container type can be used as a parent node, and the node can be mounted under the container type when creating a node. The node types currently supported by the board are:
| Type | Container type |
| --- | --- |
| composite_shape | no |
| text_shape | no |
| connector | no |
| section | yes |
| table | yes |
| group | no |
| table_uml | no |
| table_er | no |
| sticky_note | no |
| mind_map | no |
| paint | no |
| image | no |
| svg | no |
| life_line | no |
| activation | no | ## Subtype Of composite_shape
The composite_shape include the following graphics:
| type | shape |
| --- | --- |
| round_rect | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/e89fe4d56c4622797fb87e6e0bc967e4_6QB3LXBKA3.png?height=56&lazyload=true&width=56) |
| round_rect2 | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/9b2380734995c7864ece5bc38dfc504f_hpWOfrLP9n.png?height=56&lazyload=true&width=56) |
| diamond | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/4abbe6d46931ec03a7451e1120fcc748_9iBv0pZ32z.png?height=56&lazyload=true&width=56) |
| rect | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/d01f30fce05b8a3d47e46598e8339d9b_99qYNc3pX6.png?height=56&lazyload=true&width=56) |
| ellipse | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/6f4b604fe9cdf05c993db9434c33158d_QOBgE2WUM1.png?height=72&lazyload=true&width=72) |
| cylinder | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/f2882c6a8750c3b41f5e4223b9914da2_hGdaXz6vvR.png?height=56&lazyload=true&width=56) |
| step | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/7b2bd10dabdd7b6865bbabc9934ec919_FWKOBSCnT7.png?height=56&lazyload=true&width=56) |
| step2 | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/e83264fb4600d32b7b4c89b03b439696_KmLENhH3hK.png?height=72&lazyload=true&width=72) |
| parallelogram | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/5228bd4d11e4bbd9c820347234812511_3pPUVi164k.png?height=56&lazyload=true&width=56) |
| trapezoid | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/b5f80cf9f9edaa077e83ce83d0119f06_888lvZL50o.png?height=72&lazyload=true&width=72) |
| bubble | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/133bb9e448221a42bcd94a8d6109b88f_PwuVDwiOoL.png?height=56&lazyload=true&width=56) |
| rect_bubble | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/d8533da9a4a65e5cfbbeb3258da74e8c_JdRBRTBDfd.png?height=72&lazyload=true&width=72) |
| right_triangle | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/df87c55997c3fdcfb6982bf09d6fd6a7_oYUulBYYy7.png?height=56&lazyload=true&width=56) |
| triangle | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/ae8fbe8c507e1ce8b2cc7ec1f6ebf3f6_VkoQ7Kj4f6.png?height=56&lazyload=true&width=56) |
| star | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/d3e6816752b7eaec30e4510b564469b6_XTstffzFmP.png?height=56&lazyload=true&width=56) |
| hexagon | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/1be5b5e8d40fcafb61ab473d671a6f7f_zcPpy4uC4x.png?height=44&lazyload=true&width=51) |
| pentagon | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/b9064c3bb4bfad130ddaf5aa2a60b4c0_5ZQABnougy.png?height=56&lazyload=true&width=56) |
| octagon | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/1a494a503d7ecef27198472f7fdc716d_WDuCTeMiDK.png?height=56&lazyload=true&width=54) |
| backward_arrow | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/80039dbb3c64e6bdd68aa9df87d86c54_MAEfXZKNMs.png?height=56&lazyload=true&width=56) |
| forward_arrow | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/701cc15422164a77d87b83f624d4c71c_C5P2Bqbp9I.png?height=72&lazyload=true&width=72) |
| double_arrow | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/ddde4f2fede9a2848074a455cd53378f_KkPINJ8dH8.png?height=56&lazyload=true&width=56) |
| cloud | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/8af52c54054f64557634547cdeca39f1_c9YwQbm6pR.png?height=56&lazyload=true&width=56) |
| brace_reverse | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/331086cec5cff6f1afa9dcac37b4f0db_hu2LqhlbIo.png?height=58&lazyload=true&width=58) |
| brace | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/9fd68bf5ff611f5ad3b73d122a709522_JsbhRuZ2jI.png?height=56&lazyload=true&width=56) |
| cross | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/59aab2414b0b7e7dd158074a82593871_bdV4EuoSN7.png?height=56&lazyload=true&width=56) |
| circular_ring | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/fef93dc4c240b9a8d90af1935f16497a_ab7koV7cKP.png?height=56&lazyload=true&width=56) |
| pie | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/0cdc9f720ce99bf257a836356b5cdd1c_qZiHC15QDj.png?height=62&lazyload=true&width=64) |
| cube | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/8bd434d3a0de73b50a0fc03e65b93bbb_yT5rboM17C.png?height=58&lazyload=true&width=72) |
| flow_chart_round_rect | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/020e7158ddb9cbc7c158a5be7d240d38_CoFpKVPIWE.png?height=56&lazyload=true&width=56) |
| flow_chart_round_rect2 | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/e30440ce5b94627e4da714fad8524262_kMNvuYTV6z.png?height=38&lazyload=true&width=47) |
| flow_chart_diamond | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/a918c52f477cf4d603c52c0d78d2c97e_a9bOcxmYWn.png?height=56&lazyload=true&width=56) |
| document_shape | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/3f5bf7fcec25bd85c438d79dbebddbb8_u92Qz6x2QY.png?height=56&lazyload=true&width=56) |
| flow_chart_parallelogram | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/05fd068cbd572a5ed0c139274bb695e5_V3Njqte4q2.png?height=38&lazyload=true&width=50) |
| flow_chart_cylinder | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/576a1779a27254f3c3c8cf5e2ff64940_38zCSCpIVM.png?height=72&lazyload=true&width=72) |
| horiz_cylinder | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/276f1a03e661995b6d8ec97f69668986_g4KXt50PpF.png?height=56&lazyload=true&width=56) |
| predefined_process | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/3753b198d89ddd7d3cf52a73c01c5899_X8ux8jkuLr.png?height=56&lazyload=true&width=56) |
| manual_input | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/f707ede301dccc0c21a50e611193e3d8_81xZqCBcs8.png?height=72&lazyload=true&width=72) |
| flow_chart_trapezoid | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/92ffb9cbe307947ab9985550917ffb59_qzGu1yOVJ3.png?height=56&lazyload=true&width=56) |
| delay_shape | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/d3ef866544c0027e965a86d53da6c0fe_PgeYyktvJI.png?height=56&lazyload=true&width=56) |
| flow_chart_hexagon | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/942ae87876d045820c749577fb25c9ec_62mUwW13hi.png?height=72&lazyload=true&width=72) |
| off_page_connector | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/01f6092f7aa535c4bcef242b29c61775_VFLfc7TCLL.png?height=56&lazyload=true&width=56) |
| flow_chart_mq | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/57773fa9f00fca3d84f67823793fca5e_rakxFvTTsb.png?height=56&lazyload=true&width=56) |
| class_interface | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/4ae742676cb550f0bc4cfd4708f9b200_GhSjzsx0wx.png?height=72&lazyload=true&width=72) |
| classifier | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/3753b198d89ddd7d3cf52a73c01c5899_nIv4bKpXw1.png?height=56&lazyload=true&width=56) |
| note_shape | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/77f97c7e425787f27761acc02ee65b13_hErOoYWGwm.png?height=56&lazyload=true&width=56) |
| actor | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/bfb91f6fadb1232959fd97cab43c706b_Iw8Ozq0VQ7.png?height=72&lazyload=true&width=72) |
| condition_shape | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/f8283a34b913a60f469f393b209668ff_z7KoFzOHCQ.png?height=72&lazyload=true&width=72) |
| condition_shape2 | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/0d99c670660443f29bf3b9ae7bb7959a_Bm7G7LJ9Xs.png?height=72&lazyload=true&width=72) |
| data_flow_round_rect | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/4ae742676cb550f0bc4cfd4708f9b200_k5qzdJo82g.png?height=72&lazyload=true&width=72) |
| data_process | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/a2eb70a65bd30bc533bb94cdf7041389_zmWR5QERIp.png?height=56&lazyload=true&width=56) |
| data_flow_ellipse | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/c7d761ef50af557afa74e0fec8a498d0_mYuT79HkTc.png?height=56&lazyload=true&width=56) |
| data_flow_round_rect3 | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/b8bf530cae5494962cf49820d352b682_3EFJdV3r1X.png?height=56&lazyload=true&width=56) |
| data_store | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/f7a25142c74d6bff23b0c1d3861ab93e_UYVPz16JUy.png?height=72&lazyload=true&width=72) |
| data_store2 | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/cd6b2f5daeb81f1cc0f1fb8823896445_VqD9BKjE92.png?height=72&lazyload=true&width=72) |
| data_store3 | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/87dfcbb078c973109ad73f8ca92876bd_4bCLzGa0X8.png?height=56&lazyload=true&width=56) |
| component_shape | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/c261db8d68d236af4e13a67fb95c9825_Eetr6f2vIv.png?height=48&lazyload=true&width=52) |
| component_shape2 | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/daf528c8a8ead5ecb118cf46a3bad2d3_EohdvHJN59.png?height=46&lazyload=true&width=56) |
| component_interface | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/4f4683b703f54b62773d5e62bcfca402_WYXXntXLYz.png?height=52&lazyload=true&width=54) |
| component_required_interface | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/d1dcda843cea906f4d59e45d6a73dfa8_50RuHxwemr.png?height=50&lazyload=true&width=44) |
| component_assembly | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/374e8f0dabcee0a0d98f1fdc0f430947_TrBkUPr3DH.png?height=50&lazyload=true&width=58) |
| state_start | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/8957de2d692f5501e9bde341fb36b248_KeQexO93vP.png?height=52&lazyload=true&width=52) |
| state_end | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/d7df40d7501c751acf7004b74dfd751b_5qf2Pd6bAa.png?height=54&lazyload=true&width=56) |
| state_concurrence | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/5c37c12898159dd9e2ba4c82500fe167_36LDgy2s51.png?height=56&lazyload=true&width=58) |
| star3 | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/2aaf9e4ba53bc1b34c0412e8cb2913fb_by2c250juN.png?height=56&lazyload=true&width=56) |
| star4 | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/943a7fa46ef913c1e4cd7cee9faa286d_g8gsfRs0Gl.png?height=56&lazyload=true&width=56) |
| star2 | ![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/569baa2d341cc9c274bd06a55f9da713_1MaMNjRjFi.png?height=72&lazyload=true&width=72) | ## Node properties
Node properties can be divided into basic properties and unique properties to specific node.

### Basic properties
#### Text
Text properties of node.
| Property name | Data type | Description | Attribute |
| --- | --- | --- | --- |
| text | string | Text content | required |
| font_weight | string | Font weight  - `regular`: regular - `bold`: bold  | optional, default regular |
| font_size | int | Font size | optional, default 14 |
| horizontal_align | string | Horizontal alignment  - `left`: left typesetting - `center`: center typesetting - `right`: right typesetting  | optional, default center |
| vertical_align | string | Vertical alignment  - `top`: top typesetting - `mid`: middle typesetting - `bottom`: bottom typesetting  | optional, default mid | #### Style
Style properties of node.
| Property name | Data type | Description | Attribute |
| --- | --- | --- | --- |
| fill_opacity | float | Fill opacity | optional, default 100range 0～100 |
| border_style | string | Border style  - `solid`: solid border - `none`: none border - `dash`: dash border - `dot`: dot border  | optional, default solid |
| border_width | string | Border width  - `extra_narrow`: extra narrow - `narrow`: narrow - `medium`: medium - `wide`: wide  | optional, default narrow |
| border_opacity | float | Border opacity | optional, default 100range 0～100 |
| h_flip | bool | Horizontal flip | optional, deault false |
| v_flip | bool | Vertical flip | optional, default false | ### Properties of spcific node
#### Image
Image properties of image shape.
| Property name | Data type | Description | Attribute |
| --- | --- | --- | --- |
| token | string | Image token | required | #### CompositeShape
Properties of composite shape.
| Property name | Data type | Description | Attribute |
| --- | --- | --- | --- |
| type | string | Subtype of composite_shape | required | #### Connector
Properties of connector.
| Property name | Data type | Description | Attribute |
| --- | --- | --- | --- |
| `start_object` | `connector.attached_object` | `Start object of connector` | `optional` |
|   `id` | `string` | `Node id` | `optional` |
| `end_object` | `connector.attached_object` | `End object of connector` | `optional` |
|   `id` | `string` | `Node id` | `optional` |
| `captions` | `connector.caption` | `Text of connector` | `optional` |
|   `data` | `object(Text)[]` | `Text` | `optional` | #### Section
Properties of section.
| Property name | Data type | Description | Attribute |
| --- | --- | --- | --- |
| title | string | Section title | required | #### Table
Properties of table.
| Property name | Data type | Description | Attribute |
| --- | --- | --- | --- |
| `title` | `string` | `Table tile` | `optional` |
| `meta` | `table.meta` | `Table meta` | `required` |
|   `row_num` | `int` | `Table row number` | `required` |
|   `col_num` | `int` | `Table column number` | `required` |
| `cells` | `table.cell[]` | `Table cells` | `optional` |
|   `row_index` | `int` | `Row index, start from 1` | `required` |
|   `col_index` | `int` | `Column index，start from 1` | `required` |
|   `merge_info` | `table.cell.merge_info` | `Meger info of cell` | `optional` |
|     `row_span` | `int` | `The number of consecutive rows to merge starting from the current row index` | `required` |
|     `col_span` | `int` | `The number of consecutive columns to merge starting from the current column index` | `required` |
|   `children` | `string[]` | `Nodes that in the cell` | `optional` |
|   `text` | `object(Text)` | `Text of the cell` | `optional` | #### MindMap
Properties of mind map.
| Property name | Data type | Description | Attribute |
| --- | --- | --- | --- |
| parent_id | string | Parent of the mind map, empty represents root of the mind map | optional | ## Node data structure
| Property name | Data type | Description | Attribute |
| --- | --- | --- | --- |
| id | string | Node id | read only |
| type | string | Node type | required |
| parent_id | string | Parent id | optional |
| children | []string | Children ids | optional |
| x | float | The x-axis position of the node relative to the board (coordinates relative to the parent container if there is a parent container), measured in pixels | required |
| y | float | The y-axis position of the node relative to the board (coordinates relative to the parent container if there is a parent container), measured in pixels | required |
| angle | float | Angle of node | optional, default 0 |
| width | float | Width of node, measured in pixels | required |
| height | float | Height of node, measured in pixels | required |
| text | object(Text) | Text of node | optional |
| style | object(Style) | Style of node | optional |
| image | object(Image) | Image properties of image shape, for image type | optional |
| composite_shape | object(CompositeShape) | Composite shape properties, for composite_shape type | optional |
| connector | object(Connector) | Connector properties, for connector type | optional |
| section | object(Section) | Section properties, for section type | optional |
| table | object(Table) | Table properties, for table, table_er and table_uml type | optional |
| mind_map | object(MindMap) | Mind map properties, for mind_map type | optional |
