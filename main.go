package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"echo-sandbox/di"
)

func main() {

	userController := di.InjectUserController()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})
	e.GET("/user", userController.GetUser())

	e.Logger.Fatal(e.Start(":8080"))
}
