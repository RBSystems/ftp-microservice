package helpers

import (
	"io"
	"net/http"
	"os"
	"strings"
)

// DownloadFile downloads the file reference in the POST body
func DownloadFile(request Request) (Request, error) {
	urlTokens := strings.Split(request.FileLocation, "/")
	request.Filename = urlTokens[len(urlTokens)-1]

	// Download the referenced file
	output, err := os.Create("downloads/" + request.Filename)
	if err != nil {
		return Request{}, err
	}

	defer output.Close()

	response, err := http.Get(request.FileLocation)
	if err != nil {
		return Request{}, err
	}

	defer response.Body.Close()

	_, err = io.Copy(output, response.Body)
	if err != nil {
		return Request{}, err
	}

	return request, nil
}
