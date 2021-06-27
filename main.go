package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var appEnv = os.Getenv("APP_ENV")

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (string, error) {
	var param MyEvent
	err := json.Unmarshal([]byte(request.Body), &param)
	if err != nil {
		return err.Error(), err
	}
	return fmt.Sprintf("Hello %s! You are in %s", param.Name, appEnv), nil
}

func main() {
	lambda.Start(HandleRequest)
}
