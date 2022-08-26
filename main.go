package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"echo-sandbox/di"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// チュートリアルサンプル
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	// コントローラーをDIして、HandlerFuncを呼び出してみた。
	userController := di.InjectUserController()
	e.GET("/user", userController.GetUser())

	e.Logger.Fatal(e.Start(":8080"))
}
