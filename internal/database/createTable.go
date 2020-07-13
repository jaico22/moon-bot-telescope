package database

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// createTable creates a table and handles errors
func createTable(tableName string, svc *dynamodb.DynamoDB, tableInput *dynamodb.CreateTableInput) {
	_, err := svc.CreateTable(tableInput)
	if err != nil {
		if err.Error() != resourceInUseException {
			log.Println("An unexpected error occured creating the table:")
			log.Println(err.Error())
			os.Exit(1)
		}
		log.Println("Table already created")
	} else {
		log.Println("Created the table", tableName)
	}
}
