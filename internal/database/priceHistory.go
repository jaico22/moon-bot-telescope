package database

import (
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// PriceRecord is a container containg data relating to a price recording
type PriceRecord struct {
	DateTime    time.Time `json:"DateTime"`
	AskingPrice float32   `json:"AskingPrice"`
}

const tableName = "PriceHistory"

// CreatePricesTable initialized the users table
func CreatePricesTable(svc *dynamodb.DynamoDB) {
	log.Println("Creating prices table")
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("DateTime"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("DateTime"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String(tableName),
	}
	createTable(tableName, svc, input)
}

// RecordPrice takes a price and records it into dynamo
func RecordPrice(svc *dynamodb.DynamoDB, priceRecord PriceRecord) {
	askingPrice := fmt.Sprintf("%1.8f", priceRecord.AskingPrice)
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":p": {
				N: aws.String(askingPrice),
			},
		},
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"DateTime": {
				S: aws.String(priceRecord.DateTime.String()),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set AskingPrice = :p"),
	}
	svc.UpdateItem(input)
	_, err := svc.UpdateItem(input)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("Price recorded")
}
