## Go and AWS - Code and Deploy a Serverless API

### Introduction

This project demonstrates how to create a serverless API using Go and AWS. It includes a comprehensive guide on setting up the environment, deploying the application, and testing the API.

### Architecture

![System Architecture](images/architecture.png)

### Technologies Used

- Go
- AWS
- AWS Lambda
- API Gateway
- DynamoDB
- GORM
- JWT

### Project Structure

The project is structured as follows:

```
.
├── handlers
│   ├── api_response.go
│   ├── create_user.go
│   ├── delete_user.go
│   ├── get_user.go
│   ├── put_user.go
│   └── unhandled_method.go
├── main.go
├── models
│   └── user.go
├── user.go
├── validators
│   └── is_email_valid.go
└── go.mod
```

### Getting Started

1. Install Go and set up your development environment.
2. Install the required dependencies:
   - AWS SDK for Go: `go get github.com/aws/aws-sdk-go/aws`
   - AWS Lambda Go SDK: `go get github.com/aws/aws-lambda-go/lambda`
   - GORM: `go get gorm.io/gorm` and `go get gorm.io/driver/mysql`
   - JWT: `go get github.com/golang-jwt/jwt/v5`
   - Bcrypt: `go get golang.org/x/crypto/bcrypt`
3. Create a DynamoDB table named `go-serverless-api`.
4. Update the DynamoDB client in `main.go` to match your DynamoDB configuration.
5. Run the application:
   ```
   go run main.go
   ```
6. The application will start running on `http://localhost:8000`.

### API Endpoints

- `POST /api/register`: Register a new user.
- `POST /api/login`: Log in a user and receive a JWT cookie.
- `GET /api/user`: Retrieve the authenticated user's information.
- `POST /api/logout`: Log out the user by removing the JWT cookie.

To deploy the API in AWS and everything works fine you need to build the executable using this command: 

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main main.go

And later compress in a zip file for the AWS upload option in the Lambda function.
