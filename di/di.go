package di

import "github.com/yomaisch/echo-sandbox/controller"

func InjectUserController() *controller.UserController {
	return controller.NewUserController()
}
