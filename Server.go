package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/jlaffaye/ftp"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

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

	go sendFile(req) //start sending the fiel asyncronously.

	fmt.Fprintf(w, "File transfer started.")
}

func sendFile(req request) {
	fmt.Printf("Sending file %s \n", req.File)

	req.SubmissionTime = time.Now()

	if req.Timeout == 0 {
		req.Timeout = 30
	}

	if strings.EqualFold(req.Username, "") {
		req.Username = "anonymous"
		req.Password = "anonymous"
	}

	timeout := time.Duration(req.Timeout) * time.Second

	conn, err := ftp.DialTimeout(req.IPAddressHostname+":21", timeout)

	if err != nil {
		http.Post(req.CallbackAddress, "text/plain", bytes.NewBufferString("There was an error connecting to the device. \n"+err.Error()))
		return
	}
	fmt.Printf("Connection opened.\n")

	err = conn.Login(req.Username, req.Password)
	if err != nil {
		http.Post(req.CallbackAddress, "text/plain", bytes.NewBufferString("There was an error connecting to the device. \n"+err.Error()))
		return
	}
	fmt.Printf("Authenticated Succesfully.\n")

	file, err := os.Open(req.File)

	conn.ChangeDir("/FIRMWARE")

	if err != nil {
		http.Post(req.CallbackAddress, "text/plain", bytes.NewBufferString("There was an error opening the file. \n "+err.Error()))
		return
	}
	fmt.Printf("File opened.\n Starting transfer \n")
	err = conn.Stor(req.Path, file)

	if err != nil {
		http.Post(req.CallbackAddress, "text/plain", bytes.NewBufferString("There was an error storing the file. \n "+err.Error()))
		return
	}
	fmt.Printf("Transfer Completed. \n")
	req.CompletionTime = time.Now()
	req.Status = "Success"
	bits, _ := json.Marshal(req)

	http.Post(req.CallbackAddress, "application/json", bytes.NewBuffer(bits))
	fmt.Printf("Response Sent. \n")
	return
}

func main() {

	goji.Post("/send", sendFileHandler)

	goji.Serve()
}
