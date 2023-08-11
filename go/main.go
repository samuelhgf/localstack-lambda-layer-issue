package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"log"
)

type event struct {
	Url    string `json:"url"`
	Number string `json:"number"`
}

func main() {
	cfg, err := config.LoadDefaultConfig(context.Background(), func(options *config.LoadOptions) error {
		options.EndpointResolverWithOptions = aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
			return aws.Endpoint{
				PartitionID:       "aws",
				URL:               "http://localstack:4566",
				SigningRegion:     "us-east-1",
				HostnameImmutable: true,
			}, nil
		})

		return nil
	})
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	lambdaClient := lambda.NewFromConfig(cfg)

	functionName := "url-to-pdf-s3"

	logType := types.LogTypeNone
	if true {
		logType = types.LogTypeTail
	}
	payload, err := json.Marshal(event{
		Url:    "https://exmaple.com",
		Number: "12345",
	})
	if err != nil {
		log.Panicf("Couldn't marshal parameters to JSON. Here's why %v\n", err)
	}
	invokeOutput, err := lambdaClient.Invoke(context.TODO(), &lambda.InvokeInput{
		FunctionName: aws.String(functionName),
		LogType:      logType,
		Payload:      payload,
	})
	if err != nil {
		log.Panicf("Couldn't invoke function %v. Here's why: %v\n", functionName, err)
	}

	fmt.Println(invokeOutput)
}
