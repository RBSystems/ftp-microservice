package helpers

import (
	"net/http"

	"github.com/byuoitav/ftp-microservice/structs"
	"github.com/labstack/echo"
)

// SendResponse returns a message in proper JSON format
func SendResponse(c echo.Context, message string) error {
	response := &structs.Response{
		Message: message,
	}

	return c.JSON(http.StatusOK, *response)
}
