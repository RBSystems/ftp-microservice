package controllers

import (
	"net/http"

	"github.com/byuoitav/ftp-microservice/helpers"
	"github.com/byuoitav/ftp-microservice/structs"
	"github.com/labstack/echo"
)

// SendInfo returns information about the /send endpoint
func SendInfo(c echo.Context) error {
	response := &structs.Response{
		Message: "Send a POST request to the /send endpoint with a body including at least FileLocation, DestinationAddress, DestinationDirectory, and CallbackAddress tokens",
	}

	return c.JSON(http.StatusOK, *response)
}

// Send initiates an FTP file transfer
func Send(c echo.Context) error {
	request := &structs.Request{}

	err := c.Bind(request)
	if err != nil {
		response := &structs.Response{
			Message: "Could not read request body: " + err.Error(),
		}

		return c.JSON(http.StatusOK, *response)
	}

	err = helpers.CheckRequest(*request)
	if err != nil {
		response := &structs.Response{
			Message: "Requests must include at least FileLocation, DestinationAddress, DestinationDirectory, and CallbackAddress tokens",
		}

		return c.JSON(http.StatusOK, *response)
	}

	go helpers.SendFile(*request) // Start sending the file asynchronously

	response := &structs.Response{
		Message: "File transfer started",
	}

	return c.JSON(http.StatusOK, *response)
}
