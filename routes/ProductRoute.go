package routes

import "richard-project-back/controllers"

type ProductRoute struct{}

func (h ProductRoute) RegisterProductRoute(controller controllers.ProductController) {
	routeObject.Group("/")
	{
		routeObject.GET("/", controller.GetProduct)
		routeObject.GET("/", controller.InsertProduct)
		routeObject.GET("/", controller.DeleteProduct)
		routeObject.GET("/", controller.UpdateProduct)
		routeObject.GET("/", controller.Teste)
	}
}
