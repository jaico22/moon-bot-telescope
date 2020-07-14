package telescope

import (
	"log"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/jaico22/moonbot-telescope/internal/database"
	"github.com/jaico22/moonbot-telescope/pkg/kraken"
)

// Config contains all of the necesary configuration information
// to configure Telescope
type Config struct {
	SampleTime time.Duration
}

var configuration Config
var svc *dynamodb.DynamoDB

// Setup initializes all dependencies
func Setup(config Config) {
	configuration = config
	svc = database.Initialize()
}

// Run starts the sniffing processes
func Run() {
	log.Println("Telescope Is Running")
	for {
		currentAskingPrice := kraken.GetDogePrice()
		log.Printf("Current Asking Price: %1.8f\n", currentAskingPrice)
		recordCurrentAskingPrice(currentAskingPrice)
		time.Sleep(configuration.SampleTime)
	}
}

func recordCurrentAskingPrice(askingPrice float32) {
	priceRecord := database.PriceRecord{
		DateTime:    time.Now(),
		AskingPrice: askingPrice,
	}
	database.RecordPrice(svc, priceRecord)
}
