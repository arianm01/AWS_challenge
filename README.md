# AWS_challenge
This project is made with golang and aws lambda
## Preview
URL of deployed api : [https://69mband2xc.execute-api.us-east-1.amazonaws.com/dev](https://69mband2xc.execute-api.us-east-1.amazonaws.com/dev/)
## Description
### main.go
this file is where we set lambda handler and then connect our handler with gorilla mux to it
### Controllers
here there is two controller function :
GetDevice that find device by id and CreateDevice that make the device and save it to dynamoDB.
### DynamoDB
we have an interface of dynamo db here for dependency injection purposes
there is there functions in this folder :
GetDynamoDB returns an instance of DynamoDB,
PutItemInDB is a function which put the desired device in dynamoDB and GetDevice is a function which returns the desired device with requested Id
### mocks
There are auto-Generated mock interfaces of dependency for unit tests, that created by mockery
### models
there is the device struct that shows how a device object should look like
### Services
there is two function here for sending two kinds of responses
### Serverless.yml
this file contains of configuration serverless framework that is used for connecting to aws platform and it have some enviroment variables like:

-AWS_TABLE_NAME : name of dynamoDB table

-ACCESS_KEY_ID : AWS_ACCESS_KEY_ID for aws service

-SECRET_ACCESS_KEY : AWS_SECRET_ACCESS_KEY for aws service

-Region : AWS_REGION to manage on which region it should be deployed
## Test
there are two test files in controller and dynamoDB folders:
deviceController_test have tests for GetDevice and CreateDevice,
dynamoDB_test have tests for PutItemInDB and GetDevice functions
