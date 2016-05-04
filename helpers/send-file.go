package helpers

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/byuoitav/ftp-microservice/helpers"
	"github.com/jlaffaye/ftp"
)

// SendFile actually sends a file via FTP
func SendFile(request helpers.Request) {
	fmt.Printf("Sending file %s\n", request.File)

	request.SubmissionTime = time.Now()

	if request.Timeout == 0 {
		request.Timeout = 30
	}

	if len(request.Username) < 1 {
		request.Username = "anonymous"
		request.Password = "anonymous"
	}

	timeout := time.Duration(request.Timeout) * time.Second

	conn, err := ftp.DialTimeout(request.IPAddressHostname+":21", timeout)
	if err != nil {
		helpers.SendResponse(request, err, "Error connecting to the client device")
		return
	}

	fmt.Println("Connection opened")

	err = conn.Login(request.Username, request.Password)
	if err != nil {
		helpers.SendResponse(request, err, "There was an error connecting to the device")
		return
	}

	fmt.Println("Authenticated succesfully")

	file, err := os.Open(request.File)
	if err != nil {
		helpers.SendResponse(request, err, "There was an error opening the file")
		return
	}

	defer file.Close()

	fmt.Println("File opened; starting transfer")

	pathToStore := request.Path + "/" + filepath.Base(request.File) // Since the FTP package doesn't do this for us, we add the filename to the dest directory.

	err = conn.Stor(pathToStore, file)
	if err != nil {
		helpers.SendResponse(request, err, "There was an error storing the file")
		return
	}

	fmt.Println("Transfer completed")

	helpers.SendResponse(request, nil, "")

	return
}
