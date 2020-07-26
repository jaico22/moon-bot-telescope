package telescope

import (
	"log"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/jaico22/moonbot-telescope/internal/database"
	"github.com/jaico22/moonbot-telescope/pkg/kraken"
)

var svc *dynamodb.DynamoDB

// Setup initializes all dependencies
func Setup() {
	svc = database.Initialize()
}

// Trigger gets the current asking price and records into the database
func Trigger() {
	currentAskingPrice := kraken.GetDogePrice()
	log.Printf("Current Asking Price: %1.8f\n", currentAskingPrice)
	recordCurrentAskingPrice(currentAskingPrice)
}

func recordCurrentAskingPrice(askingPrice float32) {
	priceRecord := database.PriceRecord{
		DateTime:    time.Now(),
		AskingPrice: askingPrice,
	}
	database.RecordPrice(svc, priceRecord)
}
