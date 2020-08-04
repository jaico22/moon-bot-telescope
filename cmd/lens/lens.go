package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jaico22/moonbot-telescope/internal/database"
	"github.com/jaico22/moonbot-telescope/internal/telescope"
)

func main() {
	lambda.Start(HandleRequest)
}

// LensResponse encapsolates data from request processing
type LensResponse struct {
	RequestCopy     telescope.DataRequest
	Records         []database.PriceRecord
	NumberOfRecords int
}

// HandleRequest is the default lambda endpoint handling
// TODO: Make a more meaningful responce
func HandleRequest(ctx context.Context, request telescope.DataRequest) (LensResponse, error) {
	log.Println("Handling ViewData request...")
	telescope.Setup()
	log.Printf("Fetching data from dates %s to %s\n", request.StartDate.String(), request.EndDate.String())
	records := telescope.GetData(request)
	response := LensResponse{
		Records:         records,
		RequestCopy:     request,
		NumberOfRecords: len(records),
	}
	return response, nil
}
