package main

import (
	"encoding/json"
	"fmt"
	"log"

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
		log.Println("New Session error: ", err.Error)
	}

	svc := dynamodb.New(sess)

	params := &dynamodb.ScanInput{
		TableName: aws.String("MathFeverCategoriesTable"),
	}
	result, err := svc.Scan(params)
	if err != nil {
		fmt.Errorf("failed to make Query API call, %v", err.Error)
	}

	var categories []model.Category
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &categories)
	if err != nil {
		errResp, _ := json.Marshal(util.ErrorJSONResponse{Error: err.Error()})
		return util.Response{StatusCode: 400, Body: string(errResp)}, err
	}

	payload, _ := json.Marshal(categories)
	resp := util.Response{
		StatusCode: 200,
		Body:       string(payload),
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
