package routes

import "richard-project-back/controllers"

type ProductRoute interface {
	RegisterProductRoute(controller controllers.ProductController)
}
