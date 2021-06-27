package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

var appEnv = os.Getenv("appEnv")

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	return fmt.Sprintf("Hello %s! You are in %s", name.Name, appEnv), nil
}

func main() {
	lambda.Start(HandleRequest)
}
