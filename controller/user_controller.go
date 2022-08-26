package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

// echo.Get()メソッドの第2引数にはecho.HandlerFuncでないとコンパイルエラーになるため、
// 本来のコントローラメソッドをラップしている。
func (uc *UserController) GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := new(User)
		// if err := c.Bind(user); err != nil {
		// 	return err
		// }
		return c.JSON(http.StatusOK, user)
	}
}

type User struct {
	Name   string `json:"name"`
	Adress string `json:"adress"`
}
