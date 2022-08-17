# AWS Lambda Function in Go

Small and simple aws functions in go

## Depedencies
```bash
go get github.com/aws/aws-lambda-go/lambda
```

## Lambda CLI commands
```bash
#Create execution role
aws configure

aws iam create-role --role-name lambda-ex --assume-role-policy-document file://trust-policy.json

aws iam attach-role-policy --role-name lambda-ex --policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole
```

## Push to Lambda

1. Build
2. Compress Executable to main.zip
3. Run `aws lambda create-function --function-name go-aws-lambda --zip-file fileb://main.zip --handler main --runtime go1.x --role arn:aws:iam::<account-id>:role/lambda-ex` to create the function in AWS Lambda

    3.1. Run `aws sts get-caller-identity` to obtain account-id
4. Run `aws lambda invoke --function-name go-aws-lambda --cli-binary-format raw-in-base64-out --payload '{"What is your name?":"Jhon", "How old are you?":348}' output.txt` to invoke the function 

