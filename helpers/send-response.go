package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/byuoitav/ftp-microservice/structs"
)

// SendResponse does something cryptic
func SendResponse(request structs.Request, err error, errorString string) {
	request.CompletionTime = time.Now()

	if err != nil {
		request.Status = "error"
		errStr := errorString + ": " + err.Error()
		request.Error = errStr
	} else {
		request.Status = "success"
	}

	bits, _ := json.Marshal(request)

	http.Post(request.CallbackAddress, "application/json", bytes.NewBuffer(bits))
	fmt.Println("Response sent")
}
