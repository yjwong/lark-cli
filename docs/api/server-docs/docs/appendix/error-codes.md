---
title: "Error Codes"
url: "https://open.larksuite.com/document/ukTMukTMukTM/uIzM0UjLyMDN14iMzQTN"
section: "Docs"
updated: "1622208198000"
---

# Error Codes

## Sheets Open API Error Codes 
|name|	 err code         | msg           | 
|  --------- | --------- | --------------- | 
WrongRequestJson|	90201|	The request body is not a JSON
WrongRange|	90202|	Wrong range format in request
Fail|	90203|	Not expected fail
WrongRequestBody|	90204|	Wrong request body
InvalidUsers|	90205|	Invalid users
EmptySheetId|	90206|	SheetId is empty
EmptySheetTitle|	90207|	Sheet title is empty
SameSheetIdOrTitle|	90208|	Same sheetId or title in request
ExistSheetId|	90209|	SheetId already exists
ExistSheetTitle|	90210|	Sheet title already exists
WrongSheetId|	90211|	Wrong sheetId
WrongRowOrCol|	90212|	Wrong row or col
PermissionFail|	90213|	No permission for file
SpreadSheetNotFound|	90214|	SpreadSheet not found
SheetIdNotFound|	90215|	Sheetid not found
EmptyValue|	90216|	Null value in request
TooManyRequest|	90217|	Too many requests
LockedCell|90218|Cells are protected
CellExcess|90219|The number of cells exceeds the 1 million limit
TooLargeResponse|90221|The data to be returned is too large
TooLargeCell|90222|Cell content exceeds 50,000 characters
ColIdNotFound|90223|ColId not set
RowIdNotFound|90224|RowId not set
NotLinkSpreadSheet|90225|ISV not associated
ExcessLimit|90226|Excess limit
TooLargeRequest|90227|Request too large
ImportProcessing|90228|Importing
WrongURLParam|90229|URL parameter error
|InternalError|	95201—95209|	Internal error, please consult customer service for details	
|DefaultCode|95299|Other errors, please contact customer service for details

## explorer Open API Error Codes 
|name|	 err code         | msg           | 
|  --------- | --------- | --------------- | 
FAILED|	91201|	Not expected failed
PARAMERR|91202|Parameters error
NOTEXIST|91203|Not found 
FORBIDDEN|91204|No permission 
DELETED|91205|Data has been deleted 
OUT_OF_LIMIT|91206|Files number out of limit
DUPLICATE|91207|Same data has already existed
REVIEW|91208|Content reviewed failed
LOCK|96201|Internal error, please consult customer service for details
RECOVER|96202|Internal error, please consult customer service for details

## Permission Open API Error Codes 
|name|	 err code         | msg           | 
|  --------- | --------- | --------------- | 
OPEN_CODE_PARAM_ERROR|	91001|	Parameters error
OPEN_CODE_FORBIDDEN|91002|Not permitted
OPEN_CODE_INVALID_OPERATION|91003|Operate Error 
OPEN_CODE_USER_NO_SHARE_PERM|91004|User does not have permission to share
OPEN_CODE_INTERNAL_ERROR|96001|Internal error, please consult customer service for details

## Comment Open API Error Codes
|name|	 err code         | msg           | 
|  --------- | --------- | --------------- | 
FAILED|	1069301|fail
PARAM_ERROR|1069302|param error
FORBIDDEN|1069303|forbidden
DOCS_HAD_BEEN_DELETED|1069304|docs had been deleted
DOCS_NOT_EXIST|1069305|docs not exist
CONTENT_REVIEW_NOT_PASS|1069306|content review not pass
NOT_EXIST|1069307|not exist
EXCEEDED_LIMIT|1069308|exceeded limit
INTERNAL_ERROR|1069399|internal error

## Comment Open API Error Codes （Post Revisions）
|name|	 err code         | msg           | 
|  --------- | --------- | --------------- | 
FAILED|	90301|fail
PARAM_ERROR|90302|param error
FORBIDDEN|90303|forbidden
META_DELETED|90304|meta had been deleted
META_NOT_EXIST|90305|meta not exist
REVIEW_NOT_PASS|90306|review not pass
INTERNAL_ERROR|90399|internal error
