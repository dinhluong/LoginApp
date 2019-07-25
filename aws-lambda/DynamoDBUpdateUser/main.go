package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
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

	user := User{}
	err := json.Unmarshal([]byte(request.Body), &user)
	if err != nil {
		fmt.Println(err)
	}

	userJSON, _ := json.Marshal(&user)
	user.Role = "user"
	fmt.Println("Got user: " + string(userJSON))

	if err != nil {
		fmt.Println("Got error marshalling user:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"UserID": {
				N: aws.String(user.UserID),
			},
		},
		TableName: aws.String(tableName),
	}

	_, err = svc.DeleteItem(input)

	if err != nil {
		fmt.Println("Got error calling DeleteItem")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	message := "Deleted item from dynamodb"
	headers := map[string]string{"Access-Control-Allow-Origin": "*", "Access-Control-Allow-Headers": "Origin, X-Requested-With, Content-Type, Accept"}
	return &events.APIGatewayProxyResponse{
		Body:       message,
		Headers:    headers,
		StatusCode: 204}, nil

}
func main() {
	lambda.Start(ContextHandler)
}
