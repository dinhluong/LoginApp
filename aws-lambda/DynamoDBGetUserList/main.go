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
	params := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}
	result, err := svc.Scan(params)
	if err != nil {
		fmt.Println(err.Error())
	}
	items := []User{}
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &items)
	if err != nil {
		mess := fmt.Sprintf("Failed to unmarshal Record, %v", err)
		return &events.APIGatewayProxyResponse{Body: mess, StatusCode: 500}, nil
	}

	itemData, _ := json.Marshal(items)
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
