package handlers

import (
	"encoding/json" // Importing JSON encoding/decoding package

	"github.com/aws/aws-lambda-go/events" // Importing AWS Lambda event definitions
)

// apiResponse constructs a standardized HTTP response for API Gateway
// status: HTTP status code to be returned
// body: Response body, which can be any data type
func apiResponse(status int, body interface{}) (*events.APIGatewayProxyResponse, error) {

	// Create an API Gateway proxy response with a JSON content type header
	resp := events.APIGatewayProxyResponse{Headers: map[string]string{"Content-Type": "application/json"}}

	// Set the HTTP status code for the response
	resp.StatusCode = status

	// Marshal the response body into a JSON string
	stringBody, _ := json.Marshal(body) // Ignoring error handling as this is assumed to always succeed
	resp.Body = string(stringBody)      // Assign the JSON string to the response body

	// Return the constructed response
	return &resp, nil

}
