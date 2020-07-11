package kraken

import (
	"github.com/jaico22/moonbot-telescope/pkg/kraken/ticker"
)

// GetDogePrice returns the latest asking price for Doge
func GetDogePrice() float32 {
	tickerData := ticker.GetTicker("xdgusd")
	return tickerData.Ask.Price
}
