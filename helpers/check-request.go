package helpers

import (
	"errors"

	"github.com/byuoitav/ftp-microservice/structs"
)

// CheckRequest to make sure request is actually a valid request
func CheckRequest(request structs.Request) error {
	if len(request.CallbackAddress) < 1 || len(request.File) < 1 || len(request.Path) < 1 || len(request.IPAddressHostname) < 1 {
		return errors.New("Invalid payload")
	}

	return nil
}
