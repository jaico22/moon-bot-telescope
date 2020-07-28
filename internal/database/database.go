package database

import (
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Initialize initializes the dynamo client
func Initialize() *dynamodb.DynamoDB {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	log.Println("Creating Session")
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	log.Println("Create DynamoDB Client...")
	svc := dynamodb.New(sess)
	CreatePricesTable(svc)
	return svc
}
