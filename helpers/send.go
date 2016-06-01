package helpers

import (
	"os"
	"time"

	"github.com/jlaffaye/ftp"
)

// SendFile actually sends a file via FTP
func SendFile(request Request) {
	request.SubmissionTime = time.Now()

	if request.Timeout == 0 {
		request.Timeout = 30
	}

	if len(request.UsernameFTP) < 1 {
		request.UsernameFTP = "anonymous"
		request.PasswordFTP = "anonymous"
	}

	timeout := time.Duration(request.Timeout) * time.Second

	connection, err := ftp.DialTimeout(request.DestinationAddress+":21", timeout)
	if err != nil {
		CallCallback(request, "Error connectionecting to the client device: "+err.Error())
		return
	}

	err = connection.Login(request.UsernameFTP, request.PasswordFTP)
	if err != nil {
		CallCallback(request, "There was an error connectionecting to the device: "+err.Error())
		return
	}

	file, err := os.Open("downloads/" + request.Filename)
	if err != nil {
		CallCallback(request, "There was an error opening the file: "+err.Error())
		return
	}

	defer file.Close()

	pathToStore := request.DestinationDirectory + "/" + request.Filename // Since the FTP package doesn't do this for us, we add the filename to the destination directory

	err = connection.Stor(pathToStore, file)
	if err != nil {
		CallCallback(request, "There was an error storing the file: "+err.Error())
		return
	}

	CallCallback(request, "")
}
