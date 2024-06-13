package handlers

import (
	"net/http"

	"github.com/JoseHurtadoBaeza/Golang-AWS-ServerlessAPI-freeCodeCamp/pkg/user"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

// Error message for unsupported HTTP methods
var ErrorMethodNotAllowed = "method not allowed"

// ErrorBody is used to structure error responses
type ErrorBody struct {
	ErrorMsg *string `json:"error,omitempty"`
}

// GetUser handles the GET request to retrieve a user or users
// It supports fetching a single user by email or fetching all users
func GetUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {

	// Extract the 'email' parameter from the query string
	email := req.QueryStringParameters["email"]

	// If an email is provided, fetch the specific user
	if len(email) > 0 {

		// Fetch the user by email
		result, err := user.FetchUser(email, tableName, dynaClient)

		// Handle errors in fetching user
		if err != nil {
			return apiResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
		}

		// Return the fetched user data
		return apiResponse(http.StatusOK, result)
	}

	// If no email is provided, fetch all users
	result, err := user.FetchUsers(tableName, dynaClient)

	// Handle errors in fetching users
	if err != nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}

	// Return the fetched users data
	return apiResponse(http.StatusOK, result)

}

// CreateUser handles the POST request to create a new user
func CreateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {

	// Create a new user with the provided request data
	result, err := user.CreateUser(req, tableName, dynaClient)

	// Handle errors in user creation
	if err != nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}

	// Return the result of user creation
	return apiResponse(http.StatusCreated, result)
}

// UpdateUser handles the PUT request to update an existing user
func UpdateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {

	// Update the user with the provided request data
	result, err := user.UpdateUser(req, tableName, dynaClient)

	// Handle errors in updating user
	if err != nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}

	// Return the result of user update
	return apiResponse(http.StatusOK, result)

}

// DeleteUser handles the DELETE request to remove a user
func DeleteUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {

	// Delete the user with the provided request data
	err := user.DeleteUser(req, tableName, dynaClient)

	// Handle errors in deleting user
	if err != nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}

	// Return a successful deletion response
	return apiResponse(http.StatusOK, nil)
}

// UnhandledMethod handles any unsupported HTTP methods
func UnhandledMethod() (*events.APIGatewayProxyResponse, error) {

	// Return an error response for unsupported methods

	return apiResponse(http.StatusMethodNotAllowed, ErrorMethodNotAllowed)
}
