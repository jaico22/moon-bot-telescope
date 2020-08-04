package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/kjk/betterguid"
)

// PriceRecord is a container containg data relating to a price recording
type PriceRecord struct {
	DateTime     time.Time `json:"DateTime"`
	AskingPrice  float32   `json:"AskingPrice"`
	BiddingPrice float32   `json:"BiddingPrice"`
}

const tableName = "MoonBot-PriceData"

// CreatePricesTable initialized the users table
func CreatePricesTable(svc *dynamodb.DynamoDB) {
	log.Println("Creating prices table")
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("ItemId"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("DateTime"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("ItemId"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("DateTime"),
				KeyType:       aws.String("RANGE"),
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
				N: aws.String(biddingPrice),
			},
			":v": {
				N: aws.String("2"),
			},
		},
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"ItemId": {
				S: aws.String(betterguid.New()),
			},
			"DateTime": {
				S: aws.String(priceRecord.DateTime.Format("2006-01-02T15:04:05Z07:00")),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("SET AskingPrice = :ap, BiddingPrice = :bp, Version = :v"),
	}
	_, err := svc.UpdateItem(input)
	if err != nil {
		log.Println(err.Error())
		return
	} else {
		log.Println("Price Data Recorded Successfully")
	}
}

// GetPricesByDateRange returns PriceRecords between startDate and endDate
func GetPricesByDateRange(svc *dynamodb.DynamoDB, startDate time.Time, endDate time.Time, version int) []PriceRecord {
	// Build expression and scan object
	log.Println("Version: ", version)
	versionFilter := expression.Name("Version").AttributeExists().And(expression.Name("Version").Equal(expression.Value(version)))
	dateTimeFilter := expression.Name("DateTime").Between(expression.Value(startDate), expression.Value(endDate))
	filter := versionFilter.And(dateTimeFilter)
	projection := expression.NamesList(
		expression.Name("DateTime"),
		expression.Name("AskingPrice"),
		expression.Name("BiddingPrice"))
	expr, err := expression.NewBuilder().WithFilter(filter).WithProjection(projection).Build()
	if err != nil {
		log.Println("An error occured building the expression: ", err.Error())
		os.Exit(1)
	}
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(tableName),
	}
	result, err := svc.Scan(params)
	if err != nil {
		log.Println("An error occured executing the query: ", err.Error())
		os.Exit(1)
	}
	// Unmarshal and return
	var priceRecords []PriceRecord
	for _, item := range result.Items {
		priceRecord := PriceRecord{}
		err = dynamodbattribute.UnmarshalMap(item, &priceRecord)
		if err != nil {
			log.Println("An error occured during unmarshalling:", err.Error())
			os.Exit(1)
		}
		priceRecords = append(priceRecords, priceRecord)
	}
	return priceRecords
}
