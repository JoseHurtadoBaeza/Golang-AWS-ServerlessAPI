package main

import (
	"os"

	"github.com/JoseHurtadoBaeza/Golang-AWS-ServerlessAPI-freeCodeCamp/pkg/handlers"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

// Declare a global variable to hold the DynamoDB client
var (
	dynaClient dynamodbiface.DynamoDBAPI
)

func main() {

	// Retrieve the AWS region from environment variables
	region := os.Getenv("AWS_REGION")

	// Create a new AWS session with the specified region
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})

	// If there is an error creating the session, exit the program
	if err != nil {
		return
	}

	// Initialize the DynamoDB client using the created session
	dynaClient = dynamodb.New(awsSession)

	// Start the Lambda function and associate it with the handler function
	lambda.Start(handler)

}

// Define the table name used in DynamoDB
const tableName = "go-serverless-api"

// The handler function processes API Gateway requests
func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	// Switch on the HTTP method of the request to handle different types of operations
	switch req.HTTPMethod {

	case "GET":
		// Handle GET request to retrieve user data
		return handlers.GetUser(req, tableName, dynaClient)

	case "POST":
		// Handle POST request to create a new user
		return handlers.CreateUser(req, tableName, dynaClient)

	case "PUT":
		// Handle PUT request to update existing user data
		return handlers.UpdateUser(req, tableName, dynaClient)

	case "DELETE":
		// Handle DELETE request to remove a user
		return handlers.DeleteUser(req, tableName, dynaClient)

	default:
		// Handle unsupported HTTP methods
		return handlers.UnhandledMethod()

	}

}
