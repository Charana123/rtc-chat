package main

import (
	"context"

	utils "github.com/Charana123/rtc-chat/backend/utils"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	ddbClient *utils.DDBClient
)

func init() {
	ddbClient = utils.NewDDBClient()
}

func Handler(ctx context.Context, req events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {

	if req.RequestContext.EventType == "CONNECT" {
	} else if req.RequestContext.EventType == "DISCONNECT" {
	}
	resp := events.APIGatewayProxyResponse{StatusCode: 200}
	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
