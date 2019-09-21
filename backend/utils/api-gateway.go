package utils

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

func ConstructResponse(statusCode int, bodyMap map[string]interface{}) (events.APIGatewayProxyResponse, error) {
	body, err := json.Marshal(bodyMap)
	fmt.Println("here1")
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, err
	}
	fmt.Println("here2")
	buf := &bytes.Buffer{}
	buf.Write(body)
	// json.HTMLEscape(buf, body)

	return events.APIGatewayProxyResponse{
		StatusCode:      statusCode,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}
