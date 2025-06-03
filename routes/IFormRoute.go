package routes

import "richard-project-back/controllers"

type IFormRoute interface {
	RegisterFormRoutes(controller controllers.FormController)
}
