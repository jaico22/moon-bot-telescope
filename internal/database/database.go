package database

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Initialize initializes the dynamo client
func Initialize() *dynamodb.DynamoDB {
	// Create Session
	log.Println("Creating Session")
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	log.Println("Create DynamoDB Client...")
	var svc *dynamodb.DynamoDB
	if os.Getenv("AWS_SAM_LOCAL") == "true" {
		// Use local endpoint if invoked from SAM
		log.Println("Local invoke detected; Using local resources")
		localCfg := aws.Config{
			Endpoint: aws.String("http://172.16.123.1:8000"),
		}
		svc = dynamodb.New(sess, &localCfg)
	} else {
		svc = dynamodb.New(sess)
	}

	// Create Prices Table
	CreatePricesTable(svc)

	return svc
}
