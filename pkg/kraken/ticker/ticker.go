package ticker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// Ticker contains data from Kraken's public Ticker endpoint
type Ticker struct {
	Ask AskBid
	Bid AskBid
}

// AskBid contains data related to Ask and Bid Arrays as per Kraken documentation
type AskBid struct {
	Price          float32
	WholeLotVolume uint32
	LotVolume      uint32
}

// GetTicker sends a request to Kraken's public Ticker endpoint for the given AssetPair
// the format that Kraken expects is "{Symbol}{Currency}" (e.g. "XDGUSD" for Doge Coin and US Dollars)
func GetTicker(AssetPair string) Ticker {
	tickerURL := fmt.Sprintf("https://api.kraken.com/0/public/Ticker?pair=%s", AssetPair)
	resp, err := http.Get(tickerURL)
	if err != nil {
		log.Panicf("Error occured sending request to get ticker data; Err=%s\n", err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicf("Error occured while parsing body: Err=%s\n", err.Error())
	}
	var response map[string]map[string]tickerResponseResult
	json.Unmarshal(body, &response)
	assetPairResult := response["result"][strings.ToUpper(AssetPair)]
	return dataFromAssetPairResult(assetPairResult)
}

func dataFromAssetPairResult(assetPairResult tickerResponseResult) Ticker {
	return Ticker{
		Ask: AskBid{
			Price:          float32(assetPairResult.AskArray[0]),
			WholeLotVolume: uint32(assetPairResult.AskArray[1]),
			LotVolume:      uint32(assetPairResult.AskArray[2]),
		},
		Bid: AskBid{
			Price:          float32(assetPairResult.BidArray[0]),
			WholeLotVolume: uint32(assetPairResult.BidArray[1]),
			LotVolume:      uint32(assetPairResult.BidArray[2]),
		},
	}
}
