package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// User type is a profile user.
type User struct {
	UserID    string
	Email     string
	Role      string
	AvatarUrl string
}

// ContextHandler proccess get User Info
func ContextHandler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	// Create DynamoDB client
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)
	tableName := os.Getenv("USER_TABLE")

	// if request.HTTPMethod == "GET" {
	UserID := request.QueryStringParameters["UserID"]

	if UserID == "" {
		email := request.QueryStringParameters["Email"]
		var queryInput = &dynamodb.QueryInput{
			Limit:     aws.Int64(1),
			TableName: aws.String(tableName),
			IndexName: aws.String("Email-index"),
			KeyConditions: map[string]*dynamodb.Condition{
				"Email": {
					ComparisonOperator: aws.String("EQ"),
					AttributeValueList: []*dynamodb.AttributeValue{
						{
							S: aws.String(email),
						},
					},
				},
			},
		}
		result, err := svc.Query(queryInput)

		if err != nil {
			fmt.Println(err.Error())
		}
		item := User{}
		err = dynamodbattribute.UnmarshalMap(result.Items[0], &item)
		if err != nil {
			mess := fmt.Sprintf("Failed to unmarshal Record, %v", err)
			return &events.APIGatewayProxyResponse{Body: mess, StatusCode: 500}, nil
		}

		if item.Email == "" {
			fmt.Println("Could not find '" + email + "'")
		}
		itemData, err := json.Marshal(item)
		if err != nil {
			mess := fmt.Sprintf("Failed to marshal Record, %v", err)
			return &events.APIGatewayProxyResponse{Body: mess, StatusCode: 500}, nil
		}
		headers := map[string]string{"Access-Control-Allow-Origin": "*", "Access-Control-Allow-Headers": "Origin, X-Requested-With, Content-Type, Accept"}
		return &events.APIGatewayProxyResponse{
			Body:       string(itemData),
			Headers:    headers,
			StatusCode: 200}, nil
	}
	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"UserID": {
				S: aws.String(UserID),
			},
		},
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	item := User{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		mess := fmt.Sprintf("Failed to unmarshal Record, %v", err)
		return &events.APIGatewayProxyResponse{Body: mess, StatusCode: 500}, nil
	}

	if item.UserID == "" {
		fmt.Println("Could not find '" + UserID + "'")
	}
	itemData, _ := json.Marshal(item)
	headers := map[string]string{"Access-Control-Allow-Origin": "*", "Access-Control-Allow-Headers": "Origin, X-Requested-With, Content-Type, Accept"}
	return &events.APIGatewayProxyResponse{
		Body:       string(itemData),
		Headers:    headers,
		StatusCode: 200}, nil
}
func main() {
	lambda.Start(ContextHandler)
}
