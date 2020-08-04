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
	DateTime     time.Time `json:"DateTime"`
	AskingPrice  float32   `json:"AskingPrice"`
	BiddingPrice float32   `json:"BiddingPrice"`
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
	log.Println("Adding record to database...")
	askingPrice := fmt.Sprintf("%1.8f", priceRecord.AskingPrice)
	biddingPrice := fmt.Sprintf("%1.8f", priceRecord.BiddingPrice)
	log.Printf("BiddingPrice: %s AskingPrice: %s\n", biddingPrice, askingPrice)
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":ap": {
				N: aws.String(askingPrice),
			},
			":bp": {
				N: aws.String((biddingPrice)),
			},
		},
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"DateTime": {
				S: aws.String(priceRecord.DateTime.String()),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("SET AskingPrice = :ap, BiddingPrice = :bp"),
	}
	_, err := svc.UpdateItem(input)
	if err != nil {
		log.Println(err.Error())
		return
	} else {
		log.Println("Price Data Recorded Successfully")
	}
}
