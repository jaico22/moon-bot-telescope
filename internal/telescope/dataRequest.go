package telescope

import (
	"time"
)

// DataRequest is a structure that contains query parameters for
// data requests from telescope
type DataRequest struct {
	StartDate time.Time `json:"StartDate"`
	EndDate   time.Time `json:"EndDate"`
	Version   int       `json:"Version"`
}
