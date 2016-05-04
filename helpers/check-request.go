package helpers

import "errors"

// CheckRequest to make sure request is actually a valid request
func CheckRequest(req Request) error {
	if len(req.CallbackAddress) < 1 || len(req.File) < 1 || len(req.Path) < 1 || len(req.IPAddressHostname) < 1 {
		return errors.New("Invalid payload")
	}

	return nil
}
