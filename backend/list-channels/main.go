package main

import (
	"context"
	"fmt"

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

	items, err := ddbClient.ListChannels()
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, err
	}
	fmt.Println(items)
	return utils.ConstructResponse(200, map[string]interface{}{
		"channels": items,
	})
}

func main() {
	lambda.Start(Handler)
}
