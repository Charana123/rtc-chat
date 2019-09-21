package main

import (
	"context"

	utils "github.com/Charana123/rtc-chat/backend/utils"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var (
	sess      *session.Session
	ddbClient *utils.DDBClient
)

func init() {
	sess = session.Must(session.NewSessionWithOptions(
		session.Options{
			SharedConfigState: session.SharedConfigEnable,
		},
	))
	ddbClient = utils.NewDDBClient(dynamodb.New(sess))
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
