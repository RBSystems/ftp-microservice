package controllers

import (
	"net/http"

	"github.com/byuoitav/ftp-microservice/helpers"
	"github.com/byuoitav/ftp-microservice/structs"
	"github.com/labstack/echo"
)

// SendFile initiates an FTP file transfer
func SendFile(c echo.Context) error {
	request := &structs.Request{}

	err := c.Bind(request)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Could not read request body: "+err.Error())
	}

	err = helpers.CheckRequest(*request)
	if err != nil {
		return c.String(http.StatusBadRequest, `Request must be in the form of:
	  {
	  	"IPAddressHostname": "string",
	  	"CallbackAddress":"",
	  	"Path": "string",
	  	"File": "./test.txt"
	  }`)
	}

	go helpers.SendFile(*request) // Start sending the file asynchronously

	return c.String(http.StatusOK, "File transfer started")
}
