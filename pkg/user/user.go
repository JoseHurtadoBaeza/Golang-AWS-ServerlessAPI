package user

import (
	"encoding/json" // Importing JSON encoding/decoding package
	"errors"        // Importing errors package for error handling

	"github.com/JoseHurtadoBaeza/Golang-AWS-ServerlessAPI-freeCodeCamp/pkg/validators" // Importing custom validators
	"github.com/aws/aws-lambda-go/events"                                              // Importing AWS Lambda event definitions
	"github.com/aws/aws-sdk-go/aws"                                                    // Importing AWS SDK base functionality
	"github.com/aws/aws-sdk-go/service/dynamodb"                                       // Importing DynamoDB service client
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"                     // Importing DynamoDB attribute marshaling/unmarshaling
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"                         // Importing DynamoDB API interface
)

// Define error messages for various scenarios
var (
	ErrorFailedToFetchRecord     = "failed to fetch record"
	ErrorFailedToUnmarshalRecord = "failed to unmarshal record"
	ErrorInvalidUserData         = "invalid user data"
	ErrorInvalidEmail            = "invalid email"
	ErrorCouldNotMarshalItem     = "could not marshal item"
	ErrorCouldNotDeleteItem      = "could not delete item"
	ErrorCouldNotDynamoPutItem   = "could not dynamo put item"
	ErrorUserAlreadyExists       = "user.User already exists"
	ErrorUserDoesNotExist        = "user.User does not exist"
)

// User struct represents a user with email, first name, and last name fields
type User struct {
	Email     string `json:"email"`     // Email of the user
	FirstName string `json:"firstName"` // First name of the user
	LastName  string `json:"lastName"`  // Last name of the user
}

// FetchUser retrieves a user by email from DynamoDB
// email: Email of the user to be fetched
// tableName: Name of the DynamoDB table
// dynaClient: DynamoDB client interface
func FetchUser(email, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*User, error) {

	// Create a DynamoDB GetItemInput with the specified email as the key
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"email": {
				S: aws.String(email),
			},
		},
		TableName: aws.String(tableName),
	}

	// Fetch the item from DynamoDB
	result, err := dynaClient.GetItem(input)

	if err != nil {
		// Return error if the fetch fails
		return nil, errors.New(ErrorFailedToFetchRecord)
	}

	// Unmarshal the DynamoDB item to the User struct
	item := new(User)
	err = dynamodbattribute.UnmarshalMap(result.Item, item)
	if err != nil {
		return nil, errors.New(ErrorFailedToUnmarshalRecord)
	}

	return item, nil

}

// FetchUsers retrieves all users from DynamoDB
// tableName: Name of the DynamoDB table
// dynaClient: DynamoDB client interface
func FetchUsers(tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*[]User, error) {

	// Create a DynamoDB ScanInput to retrieve all items
	input := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}

	// Scan the table to retrieve all items
	result, err := dynaClient.Scan(input)
	if err != nil {
		// Return error if the scan fails
		return nil, errors.New(ErrorFailedToFetchRecord)
	}

	// Unmarshal the list of items to a slice of User structs
	item := new([]User)
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, item)

	return item, nil
}

// CreateUser creates a new user in DynamoDB
// req: API Gateway proxy request containing the user data
// tableName: Name of the DynamoDB table
// dynaClient: DynamoDB client interface
func CreateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*User, error) {

	var u User

	// Unmarshal the request body to the User struct
	if err := json.Unmarshal([]byte(req.Body), &u); err != nil {
		return nil, errors.New(ErrorInvalidUserData)
	}

	// Validate the user's email
	if !validators.IsEmailValid(u.Email) {
		return nil, errors.New(ErrorInvalidEmail)
	}

	// Check if the user already exists
	currentUser, _ := FetchUser(u.Email, tableName, dynaClient)
	if currentUser != nil && len(currentUser.Email) != 0 {
		return nil, errors.New(ErrorUserAlreadyExists)
	}

	// Marshal the User struct to a DynamoDB map
	av, err := dynamodbattribute.MarshalMap(u)
	if err != nil {
		return nil, errors.New(ErrorCouldNotMarshalItem)
	}

	// Create a DynamoDB PutItemInput with the marshaled user item
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	// Put the item into the DynamoDB table
	_, err = dynaClient.PutItem(input)
	if err != nil {
		return nil, errors.New(ErrorCouldNotDynamoPutItem)
	}

	return &u, nil
}

// UpdateUser updates an existing user in DynamoDB
// req: API Gateway proxy request containing the updated user data
// tableName: Name of the DynamoDB table
// dynaClient: DynamoDB client interface
func UpdateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*User, error) {

	var u User

	// Unmarshal the request body to the User struct
	if err := json.Unmarshal([]byte(req.Body), &u); err != nil {
		return nil, errors.New(ErrorInvalidEmail)
	}

	// Check if the user exists
	currentUser, _ := FetchUser(u.Email, tableName, dynaClient)
	if currentUser != nil && len(currentUser.Email) == 0 {
		return nil, errors.New(ErrorUserDoesNotExist)
	}

	// Marshal the User struct to a DynamoDB map
	av, err := dynamodbattribute.MarshalMap(u)
	if err != nil {
		return nil, errors.New(ErrorCouldNotMarshalItem)
	}

	// Create a DynamoDB PutItemInput with the marshaled user item
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	// Put the item into the DynamoDB table
	_, err = dynaClient.PutItem(input)
	if err != nil {
		return nil, errors.New(ErrorCouldNotDynamoPutItem)
	}

	return &u, nil
}

// DeleteUser deletes a user from DynamoDB by email
// req: API Gateway proxy request containing the email as a query parameter
// tableName: Name of the DynamoDB table
// dynaClient: DynamoDB client interface
func DeleteUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) error {

	// Extract the email from the query parameters
	email := req.QueryStringParameters["email"]

	// Create a DynamoDB DeleteItemInput with the specified email as the key
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"email": {
				S: aws.String(email),
			},
		},
		TableName: aws.String(tableName),
	}

	// Delete the item from DynamoDB
	_, err := dynaClient.DeleteItem(input)
	if err != nil {
		// Return error if the delete fails
		return errors.New(ErrorCouldNotDeleteItem)
	}

	return nil
}
