package ticker

type tickerResponseResult struct {
	AskArray                        []float32Str `json:"a,[]string"`
	BidArray                        []float32Str `json:"b"`
	CloseArray                      []float32Str `json:"c"`
	VolumeArray                     []float32Str `json:"v"`
	VolumeWeightedAveragePriceArray []float32Str `json:"p"`
	TotalNumberOfTrades             []int32      `json:"t"`
	LowArray                        []float32Str `json:"l"`
	HighArray                       []float32Str `json:"h"`
	OpeningPrice                    float32Str   `json:"o"`
}
