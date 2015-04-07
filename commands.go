package main

import (
	"fmt"

	"strings"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/aws/awsutil"
	"github.com/awslabs/aws-sdk-go/service/sqs"
	"github.com/codegangsta/cli"
)

func getCommands() []cli.Command {
	return []cli.Command{
		{
			Name:    "poll",
			Aliases: []string{"p"},
			Usage:   "--poll",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "region, r",
					Value: "us-east-1",
					Usage: "region for aws",
				},
				cli.StringFlag{
					Name:  "url, u",
					Value: "https://sqs.us-east-1.amazonaws.com/123/queue1",
					Usage: "queue url for SQS",
				},
			},
			Action: func(c *cli.Context) {
				reg := c.String("region")
				queueURL := c.String("url")
				svc := sqs.New(&aws.Config{Region: reg})
				params := &sqs.ReceiveMessageInput{
					QueueURL:            aws.String(queueURL),
					MaxNumberOfMessages: aws.Long(1),
					VisibilityTimeout:   aws.Long(1),
					WaitTimeSeconds:     aws.Long(1),
				}
				resp, err := svc.ReceiveMessage(params)

				if awserr := aws.Error(err); awserr != nil {
					fmt.Println("Error:", awserr.Code, awserr.Message)
				} else if err != nil {
					panic(err)
				}
				handle := awsutil.StringValue(resp.Messages[0].ReceiptHandle)
				handle = strings.Replace(handle, "\"", "", -1)

				params2 := &sqs.DeleteMessageInput{
					QueueURL:      aws.String(queueURL), // Required
					ReceiptHandle: aws.String(handle),   // Required
				}

				_, e := svc.DeleteMessage(params2)

				if awserr := aws.Error(e); awserr != nil {
					fmt.Println("Error:", awserr.Code, awserr.Message)
				} else if err != nil {
					panic(err)
				}

			},
		},
	}
}
