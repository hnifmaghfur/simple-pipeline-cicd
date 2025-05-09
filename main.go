package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// new route instance
	route := echo.New()

	// route
	route.GET("/hi", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Hello Hanif !",
		})
	})

	// start server
	route.Start(":3000")
}
