package controllers

import (
	"fmt"
	"net/http"

	"github.com/byuoitav/ftp-microservice/structs"
	"github.com/byuoitav/hateoas"
	"github.com/labstack/echo"
)

// Root offers HATEOAS links from the root endpoint of the API
func Root(c echo.Context) error {
	hateoasObject := hateoas.GetInfo()

	links, err := hateoas.AddLinks(c, []string{})
	if err != nil {
		response := &structs.Response{
			Message: "An error occurred: " + err.Error(),
		}

		return c.JSON(http.StatusBadRequest, *response)
	}

	fmt.Println("ONE")

	hateoasObject.Links = links

	return c.JSON(http.StatusOK, hateoasObject)
}
