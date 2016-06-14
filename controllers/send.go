package controllers

import (
	"net/http"

	"github.com/byuoitav/ftp-microservice/helpers"
	"github.com/jessemillar/jsonresp"
	"github.com/labstack/echo"
)

// SendInfo returns information about the /send endpoint
func SendInfo(context echo.Context) error {
	return jsonresp.New(context, http.StatusBadRequest, "Send a POST request to the /send endpoint with a body including at least DestinationAddress, DestinationDirectory, and CallbackAddress tokens")
}

// Send initiates an FTP file transfer
func Send(context echo.Context) error {
	request := helpers.Request{}

	err := context.Bind(&request)
	if err != nil {
		return jsonresp.New(context, http.StatusBadRequest, "Could not read request body: "+err.Error())
	}

	if len(request.CallbackAddress) < 1 || len(request.DestinationDirectory) < 1 || len(request.DestinationAddress) < 1 {
		return jsonresp.New(context, http.StatusBadRequest, "Requests must include at least DestinationAddress, DestinationDirectory, and CallbackAddress tokens")
	}

	go helpers.DownloadFile(request) // Download and send the file

	return jsonresp.New(context, http.StatusBadRequest, "File transfer started")
}
