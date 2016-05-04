package main

import (
	"fmt"

	"github.com/byuoitav/av-api/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
	"github.com/labstack/echo/middleware"
)

func main() {
	port := ":8002"
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	e.Get("/health", controllers.Health)
	e.Post("/send", controllers.SendFile)

	fmt.Printf("The FTP microservice is listening on %s\n", port)
	e.Run(fasthttp.New(port))
}
