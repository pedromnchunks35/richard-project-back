package routes

import "richard-project-back/controllers"

type ProductRouteImpl struct{}

func (h ProductRouteImpl) RegisterProductRoute(controller controllers.ProductController) {
	routeObject.Group("/")
	{
		routeObject.GET("/", controller.GetProduct)
	}
}
