package routes

import "richard-project-back/controllers"

type FormRoute struct{}

func (h FormRoute) RegisterFormRoute(controller controllers.FormController) {
	routeObject.Group("/form")
	{
		routeObject.GET("/get", controller.GetForm)
		routeObject.GET("/insert", controller.InsertForm)
		routeObject.GET("/delete", controller.DeleteForm)
		routeObject.GET("/update", controller.UpdateForm)
		routeObject.GET("/test", controller.Teste)
	}
}
