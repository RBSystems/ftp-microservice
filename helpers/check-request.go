package helpers

import (
	"errors"

	"github.com/byuoitav/ftp-microservice/structs"
)

// CheckRequest to make sure request is actually a valid request
func CheckRequest(request structs.Request) error {
	if len(request.CallbackAddress) < 1 || len(request.FileLocation) < 1 || len(request.DestinationDirectory) < 1 || len(request.DestinationAddress) < 1 {
		return errors.New("Invalid payload")
	}

	return nil
}
