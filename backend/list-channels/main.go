package main

import (
	"context"
	"encoding/json"
	"fmt"

	utils "github.com/Charana123/rtc-chat/backend/utils"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	ddbClient        *utils.DDBClient
	apiGatewayClient *utils.APIGateway
)

func init() {
	ddbClient = utils.NewDDBClient()
	apiGatewayClient = utils.NewAPIGateway()
}

type ListChannels struct {
	FilterPrefix string `json:"FilterPrefix"`
}

func Handler(ctx context.Context, req events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {

	data := &ListChannels{}
	json.Unmarshal([]byte(req.Body), data)

	fmt.Println("data.filterPrefix", data.FilterPrefix)
	items, err := ddbClient.ListChannels(data.FilterPrefix)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, err
	}
	fmt.Println("items", items)
	err = apiGatewayClient.SendMessageToClient(req.RequestContext.ConnectionID, items)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, nil
	}
	return events.APIGatewayProxyResponse{StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
