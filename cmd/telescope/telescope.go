package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jaico22/moonbot-telescope/internal/telescope"
)

func main() {
	lambda.Start(HandleRequest)
}

// RecordEvent is a placeholder for the lambda request
type RecordEvent struct {
}

// HandleRequest is the default lambda endpoint handling
// TODO: Make a more meaningful responce
func HandleRequest(ctx context.Context, recordEvent RecordEvent) (string, error) {
	telescope.Setup()
	telescope.Trigger()
	return fmt.Sprintf("Recorded", nil), nil
}
