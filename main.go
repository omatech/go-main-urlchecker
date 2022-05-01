package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/omatech/urlchecker"
)

func main() {
	lambda.Start(Handler)
}

func Handler(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	ApiResponse := events.APIGatewayProxyResponse{}

	path := request.RawPath
	tokenToCheck := request.QueryStringParameters["token"]
	timestamp := urlchecker.GetTimestamp()
	tokenFromUrl := urlchecker.GenerateToken(path, timestamp)
	generatedToken := urlchecker.GenerateToken(path, timestamp)
	err := urlchecker.Check(path, timestamp, tokenToCheck)
	if err == nil {
		ApiResponse = events.APIGatewayProxyResponse{Body: urlchecker.GetRemoteFileInBase64(path), IsBase64Encoded: true, Headers: map[string]string{"Content-Type": "application/pdf"}, StatusCode: 200}

	} else {
		message := urlchecker.Debug(path, timestamp, tokenToCheck)
		message += fmt.Sprintf("TokenToCheck=%s\n", tokenToCheck)
		message += fmt.Sprintf("TokenFromUrl=%s\n", tokenFromUrl)
		message += fmt.Sprintf("GeneratedToken=%s\n", generatedToken)
		message += fmt.Sprintf("Path=%s\n", path)
		message += fmt.Sprintf("ERROR:%v\n", err)
		ApiResponse = events.APIGatewayProxyResponse{Body: message, Headers: map[string]string{"Content-Type": "application/json"}, StatusCode: 200}
	}

	return ApiResponse, nil
}
