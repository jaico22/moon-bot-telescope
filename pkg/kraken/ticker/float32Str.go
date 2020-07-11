package ticker

import (
	"encoding/json"
	"strconv"
)

// float32Str is a wrapper around float32
type float32Str float32

// UnmarshalJSON extends float32Str to allow for JSON umarshalling
func (f *float32Str) UnmarshalJSON(b []byte) error {
	// Try string first
	var s string
	if err := json.Unmarshal(b, &s); err == nil {
		value, err := strconv.ParseFloat(s, 10)
		if err != nil {
			return err
		}
		*f = float32Str(value)
		return nil
	}

	// Fallback to number
	return json.Unmarshal(b, (*float32)(f))
}
