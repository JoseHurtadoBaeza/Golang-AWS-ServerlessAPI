# Golang-AWS-ServerlessAPI
Repository to store all the code from the freecodecamp tutorial named "Go and AWS - Code and Deploy a Serverless API"

To deploy the API in AWS and everything works fine you need to build the executable using this command: 

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main main.go

And later compress in a zip file for the AWS upload option in the Lambda function.