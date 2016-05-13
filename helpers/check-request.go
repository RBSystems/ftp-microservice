package helpers

import (
	"errors"
)

// CheckRequest to make sure request is actually a valid request
func CheckRequest(request Request) error {
	if len(request.CallbackAddress) < 1 || len(request.FileLocation) < 1 || len(request.DestinationDirectory) < 1 || len(request.DestinationAddress) < 1 {
		return errors.New("Invalid payload")
	}

	return nil
}
