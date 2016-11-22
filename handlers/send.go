package handlers

import (
	"net/http"

	"github.com/byuoitav/ftp-microservice/helpers"
	"github.com/jessemillar/jsonresp"
	"github.com/labstack/echo"
)

// SendInfo returns information about the /send endpoint
func SendInfo(context echo.Context) error {
	jsonresp.New(context.Response(), http.StatusBadRequest, "Send a POST request to the /send endpoint with a body including at least DestinationAddress, DestinationDirectory, and CallbackAddress tokens")
	return nil
}

// Send initiates an FTP file transfer
func Send(context echo.Context) error {
	request := helpers.Request{}

	err := context.Bind(&request)
	if err != nil {
		jsonresp.New(context.Response(), http.StatusBadRequest, "Could not read request body: "+err.Error())
		return nil
	}

	if len(request.CallbackAddress) < 1 || len(request.DestinationDirectory) < 1 || len(request.DestinationAddress) < 1 {
		jsonresp.New(context.Response(), http.StatusBadRequest, "Requests must include at least DestinationAddress, DestinationDirectory, and CallbackAddress tokens")
		return nil
	}

	go helpers.DownloadFile(request) // Download and send the file

	jsonresp.New(context.Response(), http.StatusBadRequest, "File transfer started")
	return nil
}
