package main

import (
	"fmt"

	"github.com/byuoitav/ftp-microservice/controllers"
	"github.com/byuoitav/hateoas"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
	"github.com/labstack/echo/middleware"
)

func main() {
	err := hateoas.Load("https://raw.githubusercontent.com/byuoitav/ftp-microservice/master/swagger.yml")
	if err != nil {
		fmt.Printf("Could not load swagger.yaml file. Error: %s", err.Error())
		panic(err)
	}

	port := ":8002"
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	e.Get("/", controllers.Root)
	e.Get("/health", controllers.Health)
	e.Get("/send", controllers.SendInfo)
	e.Post("/send", controllers.Send)

	fmt.Printf("The FTP microservice is listening on %s\n", port)
	e.Run(fasthttp.New(port))
}
