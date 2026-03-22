---
title: "Float image user guide"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-float_image/float-image-user-guide"
service: "sheets-v3"
resource: "spreadsheet-sheet-float_image"
section: "Docs"
updated: "1647001421000"
---

# Floating image user guide

## Scenarios

These APIs are used to operate on floating images in sheets.

## Image description

Spreadsheets only store image tokens. Before creating a floating image, you must upload the image to the spreadsheet (see Upload a material and Upload a material in blocks). When you query or obtain images, only the token is returned. To download images, see Download a material and Obtain a temporary material download URL.

## Supported APIs

The token of a single image can be placed in different locations in a spreadsheet. The number of floating images with different tokens cannot exceed 4,000 in a spreadsheet. You can use the following APIs to manage floating images:

1. Obtain a floating image: Obtains floating image information of the specified ID.
2. Create a floating image: Creates a floating image based on the parameters passed in.
3. Update a floating image: Updates the position and size of a floating image.
4. Delete a floating image: Deletes a floating image.
5. Query floating images: Queries the information on all floating images in a sheet.

## Floating image parameters
### Example
![image.png](//sf16-sg.larksuitecdn.com/obj/open-platform-opendoc-sg/5ba581c9134323943e6d5de6f45bc58e_tdh4ZQGMsG.png?lazyload=true&width=634&height=438)

### **float_image_id**

The unique ID of a floating image in a sub-sheet. You can specify a custom ID when creating an image, which will be verified for validity. Otherwise, a default ID is automatically generated. Valid IDs must be 10 characters long and composed of numbers (0-9) and upper and lowercase English letters.

### **float_image_token**

The token of a floating image, which is obtained through the Upload imagesAPI. This field is required for creating a floating image.

### **range**

The position of the cell where the top-left corner of a floating image is located. You can only specify a single cell, such as "ahgsch!A1:A1". This field is required for creating a floating image.

### **width**

The display width of a floating image, which must be greater than or equal to 20 pixels. If this field is not specified when an image is created, the actual width of the image is used.

### **height**

The display height of a floating image, which must be greater than or equal to 20 pixels. If this field is not specified when an image is created, the actual height of the image is used.

### **offset_x**

The horizontal offset of the top-left corner of a floating image from the top-left corner of the cell where it is located, which must be greater than or equal to 0 and less than the width of the cell where the top-left corner of the floating image is located. It defaults to 0.

### **offset_y**

The vertical offset of the top-left corner of a floating image from the top-left corner of the cell where it is located, which must be greater than or equal to 0 and less than the height of the cell in the top-left corner of the floating image. It defaults to 0.
