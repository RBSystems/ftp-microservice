package main

import (
	"fmt"

	"github.com/byuoitav/ftp-microservice/handlers"
	"github.com/byuoitav/hateoas"
	"github.com/byuoitav/wso2jwt"
	"github.com/jessemillar/health"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
	"github.com/labstack/echo/middleware"
)

func main() {
	err := hateoas.Load("https://raw.githubusercontent.com/byuoitav/ftp-microservice/master/swagger.yml")
	if err != nil {
		fmt.Println("Could not load swagger.yaml file. Error: " + err.Error())
		panic(err)
	}

	port := ":8002"
	router := echo.New()
	router.Pre(middleware.RemoveTrailingSlash())

	router.Get("/", hateoas.RootResponse)
	router.Get("/health", health.Check)
	router.Get("/send", handlers.SendInfo, wso2jwt.ValidateJWT())
	router.Post("/send", handlers.Send, wso2jwt.ValidateJWT())

	fmt.Println("The FTP microservice is listening on " + port)
	server := fasthttp.New(port)
	server.ReadBufferSize = 1024 * 10 // Needed to interface properly with WSO2
	router.Run(server)
}
