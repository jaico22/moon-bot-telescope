package kraken

import (
	"github.com/jaico22/moonbot-telescope/pkg/kraken/ticker"
)

// DogePriceData contains relevant data parsed from kraken ticker api
type DogePriceData struct {
	AskingPrice  float32
	BiddingPrice float32
}

// GetDogePrice returns the latest asking price for Doge
func GetDogePrice() DogePriceData {
	tickerData := ticker.GetTicker("xdgusd")
	return DogePriceData{
		AskingPrice:  tickerData.Ask.Price,
		BiddingPrice: tickerData.Bid.Price,
	}
}
