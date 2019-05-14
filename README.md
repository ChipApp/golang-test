# Chip Golang test assignment
Create a RESTful API server which implements 1 endpoint:

Checks whether a date is a working day in UK or not.
## URL
/working_day/:date
## Method:
GET

## URL Params

Required

date=[string:YYYY-m-d]
## Data Params
None

## Success Response:
Code: 200 

Content: ```{ working_day : true }```

## Error Response:
Code: 400 BAD REQUEST 

Content: ```{ error : { code: 400, message: "wrong date format" }}```

## Sample Call:
curl http://127.0.0.1:8181/working_day/2019-01-01

