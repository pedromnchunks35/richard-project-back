package routes

import "richard-project-back/controllers"

type HelloWorldRouteImpl struct{}

func (h HelloWorldRouteImpl) RegisterHelloWorldRoutes(controller controllers.HelloWorldController) {
	routeObject.Group("/")
	{
		routeObject.GET("/", controller.GetHelloWorld)
	}
}
