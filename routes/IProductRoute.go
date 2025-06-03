package routes

import "richard-project-back/controllers"

type IProductRoute interface {
	RegisterProductRoute(controller controllers.ProductController)
}
