package di

import "echo-sandbox/controller"

func InjectUserController() *controller.UserController {
	return controller.NewUserController()
}
