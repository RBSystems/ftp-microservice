package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// CallCallback notifies an external service that the FTP file transfer is complete
func CallCallback(request Request, err string) {
	request.CompletionTime = time.Now()

	if err != "" {
		request.Status = "error"
		request.Error = err
	} else {
		request.Status = "success"
	}

	bits, _ := json.Marshal(request)

	http.Post(request.CallbackAddress, "application/json", bytes.NewBuffer(bits))
	fmt.Printf("Response sent: %+v", request)
}
