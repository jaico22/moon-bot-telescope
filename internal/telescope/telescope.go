package telescope

import (
	"log"
	"time"

	"github.com/jaico22/moonbot-telescope/internal/database"
	"github.com/jaico22/moonbot-telescope/pkg/kraken"
)

// Config contains all of the necesary configuration information
// to configure Telescope
type Config struct {
	SampleTime time.Duration
}

var configuration Config

// Setup initializes all dependencies
func Setup(config Config) {
	configuration = config
	database.Initialize()
}

// Run starts the sniffing processes
func Run() {
	log.Println("Telescope Is Running")
	for {
		currentAskingPrice := kraken.GetDogePrice()
		log.Printf("Current Asking Price: %1.8f\n", currentAskingPrice)
		time.Sleep(configuration.SampleTime)
	}
}
