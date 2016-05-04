package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/byuoitav/ftp-microservice/helpers"
)

// SendResponse does something cryptic
func SendResponse(req helpers.Request, err error, errorString string) {
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
	fmt.Println("Response sent")
}
