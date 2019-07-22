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
	user := User{}
	user.UserID = request.QueryStringParameters["UserID"]
	user.Email = request.QueryStringParameters["Email"]
	user.AvatarUrl = request.QueryStringParameters["AvatarUrl"]

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
	// }
	// if request.HTTPMethod == "PUT" {
	// 	user := new(User)
	// 	user.UserID = request.PathParameters["UserID"]
	// 	user.UserName = request.QueryStringParameters["UserName"]
	// 	user.AvatarUrl = request.QueryStringParameters["AvatarUrl"]
	// 	input := &dynamodb.UpdateItemInput{
	// 		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
	// 			":r":{
	// 				N: aws.String(user.UserName)
	// 			}
	// 		}
	// 	}
	// }
	// return nil, nil
}
func main() {
	lambda.Start(ContextHandler)
}
