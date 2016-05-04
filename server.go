package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/byuoitav/av-api/controllers"
	"github.com/byuoitav/ftp-microservice/helpers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
	"github.com/labstack/echo/middleware"
)

func sendResponse(req helpers.Request, err error, errorString string) {
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

func main() {
	port := ":8002"
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	e.Get("/health", controllers.Health)
	e.Post("/send", controllers.SendFile)

	fmt.Printf("The FTP microservice is listening on %s\n", port)
	e.Run(fasthttp.New(port))
}
