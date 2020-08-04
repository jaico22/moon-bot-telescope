package telescope

import (
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
	currentDogePriceData := kraken.GetDogePrice()
	recordPriceData(currentDogePriceData)
}

// GetData parses the request agianst the database and returns matching records
func GetData(request DataRequest) []database.PriceRecord {
	return database.GetPricesByDateRange(svc, request.StartDate, request.EndDate, request.Version)
}

func recordPriceData(currentDogePriceData kraken.DogePriceData) {
	priceRecord := database.PriceRecord{
		DateTime:     time.Now(),
		AskingPrice:  currentDogePriceData.AskingPrice,
		BiddingPrice: currentDogePriceData.BiddingPrice,
	}
	database.RecordPrice(svc, priceRecord)
}
