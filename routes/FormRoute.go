package routes

import "richard-project-back/controllers"

type FormRoute struct{}

func (h FormRoute) RegisterFormRoute(controller controllers.FormController) {
	routeObject.Group("/")
	{
		routeObject.GET("/", controller.GetForm)
		routeObject.GET("/", controller.InsertForm)
		routeObject.GET("/", controller.DeleteForm)
		routeObject.GET("/", controller.UpdateForm)
		routeObject.GET("/", controller.Teste)
	}
}
