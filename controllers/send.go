package controllers

import (
	"net/http"

	"github.com/byuoitav/ftp-microservice/helpers"
	"github.com/labstack/echo"
)

// SendFile initiates an FTP file transfer
func SendFile(c echo.Context) error {
	req := &helpers.Request{}

	err := c.Bind(req)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Could not read request body: "+err.Error())
	}

	err = helpers.CheckRequest(*req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, `Request must be in form of:
	  {
	  	"IPAddressHostname": "string",
	  	"CallbackAddress":"",
	  	"Path": "string",
	  	"File": "./test.txt"
	  }`)
	}

	go helpers.SendFile(*req) // Start sending the file asynchronously

	return c.String(http.StatusOK, "File transfer started")
}
