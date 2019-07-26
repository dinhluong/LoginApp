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
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
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
	//Delete from dynamodb
	// Create DynamoDB client
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB clientd
	svc := dynamodb.New(sess)
	tableName := os.Getenv("USER_TABLE")
	adminGroup := os.Getenv("ADMIN_GROUP")
	userPoolID := os.Getenv("USERPOOLID")
	user := User{}
	err := json.Unmarshal([]byte(request.Body), &user)
	if err != nil {
		fmt.Println(err)
	}

	userJSON, _ := json.Marshal(&user)
	user.Role = "admin"
	fmt.Println("Got user: " + string(userJSON))

	if err != nil {
		fmt.Println("Got error marshalling user:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":r": {
				S: aws.String(user.Role),
			},
			":u": {
				S: aws.String(user.UserID),
			},
		},
		ExpressionAttributeNames: map[string]*string{
			"#d": aws.String("Role"),
		},
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"UserID": {
				S: aws.String(user.UserID),
			},
		},
		ConditionExpression: aws.String("UserID = :u"),
		UpdateExpression:    aws.String("set #d = :r"),
		ReturnValues:        aws.String("UPDATED_NEW"),
	}

	_, err = svc.UpdateItem(input)

	if err != nil {
		fmt.Println("Got error calling Update item")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	params := &cognitoidentityprovider.AdminAddUserToGroupInput{
		GroupName:  aws.String(adminGroup),
		UserPoolId: aws.String(userPoolID),
		Username:   aws.String(user.Email),
	}
	svc1 := cognitoidentityprovider.New(sess)
	_, err = svc1.AdminAddUserToGroup(params)
	if err != nil {
		fmt.Println("Got error calling coginito add user " + user.UserID + " To group " + adminGroup)
		fmt.Println(err.Error())
		os.Exit(1)
	}
	headers := map[string]string{"Access-Control-Allow-Origin": "*", "Access-Control-Allow-Headers": "Origin, X-Requested-With, Content-Type, Accept, Access-Control-Allow-Methods", "Access-Control-Allow-Methods": "*"}
	return &events.APIGatewayProxyResponse{
		Headers:    headers,
		StatusCode: 204}, nil

}
func main() {
	lambda.Start(ContextHandler)
}
