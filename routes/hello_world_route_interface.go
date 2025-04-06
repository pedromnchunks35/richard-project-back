package routes

import "richard-project-back/controllers"

type HelloWorldRoute interface {
	RegisterHelloWorldRoutes(controller controllers.HelloWorldController)
}
