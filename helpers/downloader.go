package helpers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// DownloadFile downloads the file reference in the POST body
func DownloadFile(request Request) (Request, error) {
	urlTokens := strings.Split(request.FileLocation, "/")
	request.FileLocation = urlTokens[len(urlTokens)-1]

	fmt.Println(request.FileLocation)

	// Download the referenced file
	output, err := os.Create(request.FileLocation)
	if err != nil {
		return Request{}, err
	}

	defer output.Close()

	response, err := http.Get(request.FileLocation)
	if err != nil {
		return Request{}, err
	}

	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)
	if err != nil {
		return Request{}, err
	}

	fmt.Println(n, "bytes downloaded")

	return request, nil
}
