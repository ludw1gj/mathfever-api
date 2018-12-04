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
	"github.com/ludw1gj/mathfever/src/categories/config"
	"github.com/ludw1gj/mathfever/src/model"
	"github.com/ludw1gj/mathfever/src/util"
)

func Handler(request events.APIGatewayProxyRequest) (util.Response, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-southeast-2"),
		Credentials: credentials.NewStaticCredentials(config.AWSAccessKeyID, config.AWSSecretAccessKey, ""),
	})
	if err != nil {
		return util.HandleErrorResponse(err, "New Session error", 500)
	}

	svc := dynamodb.New(sess)

	params := &dynamodb.ScanInput{
		TableName: aws.String("MathFeverCategoriesTable"),
	}
	result, err := svc.Scan(params)
	if err != nil {
		return util.HandleErrorResponse(err, "failed to make Query API call", 500)
	}

	var categories []model.Category
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &categories)
	if err != nil {
		return util.HandleErrorResponse(err, "failed to unmarshal list", 500)
	}

	payload, _ := json.Marshal(categories)
	resp := util.Response{
		IsBase64Encoded: true,
		StatusCode:      200,
		Body:            string(payload),
		Headers: map[string]string{
			"Content-Type":                     "application/json",
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Credentials": "true",
		},
	}
	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
