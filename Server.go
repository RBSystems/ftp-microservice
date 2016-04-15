package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/jlaffaye/ftp"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func sendResponse(req request, err error, errorString string) {
	req.CompletionTime = time.Now()

	if err != nil {
		req.Status = "error"
		errStr := errorString + ": " + err.Error()
		req.Error = errStr
	} else {
		req.Status = "success"
	}

	bits, _ := json.Marshal(req)

	http.Post(req.CallbackAddress, "application/json", bytes.NewBuffer(bits))
	fmt.Printf("Response Sent. \n")
}

func sendFileHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	bits, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Could not read request body: %s\n", err.Error())
		return
	}

	var req request

	err = json.Unmarshal(bits, &req)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error with the request body: %s", err.Error())
		return
	}

	err = checkReq(req)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `Invalid request. Request must be in form of:
    {
    	"IPAddressHostname": "string",
    	"CallbackAddress":"",
    	"Path": "string",
    	"File": "./test.txt"
    }`)
		return
	}

	go sendFile(req) //start sending the fiel asyncronously.

	fmt.Fprintf(w, "File transfer started.")
}

func checkReq(req request) error {
	if len(req.CallbackAddress) < 1 || len(req.File) < 1 || len(req.Path) < 1 || len(req.IPAddressHostname) < 1 {
		return errors.New("Invalid Payload.")
	}
	return nil
}

func sendFile(req request) {
	fmt.Printf("Sending file %s \n", req.File)

	req.SubmissionTime = time.Now()

	if req.Timeout == 0 {
		req.Timeout = 30
	}

	if len(req.Username) < 1 {
		req.Username = "anonymous"
		req.Password = "anonymous"
	}

	timeout := time.Duration(req.Timeout) * time.Second

	conn, err := ftp.DialTimeout(req.IPAddressHostname+":21", timeout)

	if err != nil {
		sendResponse(req, err, "Error connecting to the client device")
		return
	}
	fmt.Printf("Connection opened.\n")

	err = conn.Login(req.Username, req.Password)
	if err != nil {
		sendResponse(req, err, "There was an error connecting to the device")
		return
	}
	fmt.Printf("Authenticated Succesfully.\n")

	file, err := os.Open(req.File)

	if err != nil {
		sendResponse(req, err, "There was an error opening the file")
		return
	}
	fmt.Printf("File opened.\n Starting transfer \n")

	pathToStore := req.Path + "/" + filepath.Base(req.File) //since the FTP package doesn't do this for us, we add the filename to the dest directory.

	err = conn.Stor(pathToStore, file)

	if err != nil {
		sendResponse(req, err, "There was an error storing the file.")
		return
	}

	fmt.Printf("Transfer Completed. \n")

	sendResponse(req, nil, "")

	return
}

func main() {

	goji.Post("/send", sendFileHandler)

	goji.Serve()
}
