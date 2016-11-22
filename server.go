package main

import (
	"log"
	"net/http"

	"github.com/byuoitav/ftp-microservice/handlers"
	"github.com/byuoitav/hateoas"
	"github.com/byuoitav/wso2jwt"
	"github.com/jessemillar/health"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	err := hateoas.Load("https://raw.githubusercontent.com/byuoitav/ftp-microservice/master/swagger.json")
	if err != nil {
		log.Fatalln("Could not load Swagger file. Error: " + err.Error())
	}

	port := ":8002"
	router := echo.New()
	router.Pre(middleware.RemoveTrailingSlash())
	router.Use(echo.WrapMiddleware(wso2jwt.ValidateJWT))

	router.GET("/", echo.WrapHandler(http.HandlerFunc(hateoas.RootResponse)))
	router.GET("/health", echo.WrapHandler(http.HandlerFunc(health.Check)))

	router.GET("/send", handlers.SendInfo)
	router.POST("/send", handlers.Send)

	server := http.Server{
		Addr:           port,
		MaxHeaderBytes: 1024 * 10,
	}

	router.StartServer(&server)
}
