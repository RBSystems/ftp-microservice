package controllers

import (
	"net/http"

	"github.com/byuoitav/ftp-microservice/helpers"
	"github.com/byuoitav/ftp-microservice/structs"
	"github.com/labstack/echo"
)

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
			Message: `Request must be in the form of:
  		{
  			"DestinationAddress": "string",
  			"CallbackAddress":"",
  			"Path": "string",
  			"File": "./test.txt"
  		}`,
		}

		return c.JSON(http.StatusOK, *response)
	}

	go helpers.SendFile(*request) // Start sending the file asynchronously

	response := &structs.Response{
		Message: "File transfer started",
	}

	return c.JSON(http.StatusOK, *response)
}
