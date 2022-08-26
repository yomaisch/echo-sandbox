package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/yomaisch/echo-sandbox/internal/di"
)

func main() {

	userController := di.InjectUserController()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})
	e.GET("/user", userController.getUser())

	e.Logger.Fatal(e.Start(":8080"))
}
