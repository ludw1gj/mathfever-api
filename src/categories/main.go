package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/ludw1gj/mathfever-api/src/api"
	"github.com/ludw1gj/mathfever-api/src/categories/config"
)

func Handler(request events.APIGatewayProxyRequest) (api.Response, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-southeast-2"),
		Credentials: credentials.NewStaticCredentials(config.AWSAccessKeyID, config.AWSSecretAccessKey, ""),
	})
	if err != nil {
		return api.HandleErrorResponse(err, "New Session error", 500)
	}

	svc := dynamodb.New(sess)
	params := &dynamodb.ScanInput{
		TableName: aws.String("MathFeverCategoriesTable"),
	}
	result, err := svc.Scan(params)
	if err != nil {
		return api.HandleErrorResponse(err, "failed to make Query API call", 500)
	}

	var categories []api.Category
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &categories)
	if err != nil {
		return api.HandleErrorResponse(err, "failed to unmarshal list", 500)
	}

	payload, err := json.Marshal(categories)
	if err != nil {
		return api.HandleErrorResponse(err, "failed to marshal payload", 500)
	}
	return api.HandleResponse(string(payload))
}

func main() {
	lambda.Start(Handler)
}
