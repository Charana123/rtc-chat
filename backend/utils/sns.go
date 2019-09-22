package utils

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

type SNSClient struct {
	svc      *sns.SNS
	TopicArn string
}

func NewSNSClient() *SNSClient {
	sess := session.Must(session.NewSessionWithOptions(
		session.Options{
			SharedConfigState: session.SharedConfigEnable,
		},
	))
	return &SNSClient{
		svc:      sns.New(sess),
		TopicArn: "arn:aws:sns:us-east-1:650894155690:MessageTopic",
	}
}

func (snsc *SNSClient) SendMessage(message string) error {
	input := &sns.PublishInput{
		Message:  aws.String(message),
		TopicArn: aws.String(snsc.TopicArn),
	}
	_, err := snsc.svc.Publish(input)
	return err
}
