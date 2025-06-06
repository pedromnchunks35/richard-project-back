package routes

import "richard-project-back/controllers"

type ProductRoute struct{}

func (h ProductRoute) RegisterProductRoute(controller controllers.ProductController) {
	routeObject.Group("/product")
	{
		routeObject.GET("/get", controller.GetProduct)
		routeObject.GET("/insert", controller.InsertProduct)
		routeObject.GET("/delete", controller.DeleteProduct)
		routeObject.GET("/update", controller.UpdateProduct)
		routeObject.GET("/test", controller.Teste)
	}
}
