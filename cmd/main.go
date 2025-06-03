package main

import (
	"fmt"
	"log"
	"os"
	"richard-project-back/controllers"
	"richard-project-back/routes"
	"richard-project-back/services"
	iService "richard-project-back/services/ServicesInterface"

	"github.com/joho/godotenv"
)

var helloWorldService iService.IHelloWorldService
var helloWorldController controllers.HelloWorldController
var helloWorldRoute routes.HelloWorldRoute

var productService iService.IProductService
var productController controllers.ProductController
var productRoute routes.ProductRoute

var formService iService.IFormService
var formController controllers.FormController
var formRoute routes.FormRoute

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error loading the environment variables %s\n", err.Error())
	}

	helloWorldService = services.HelloWorldServiceImpl{}
	helloWorldController = controllers.RegisterHelloWorldControllerImpl(helloWorldService)
	helloWorldRoute = routes.HelloWorldRouteImpl{}

	productService = services.NewProductService()
	productController = *controllers.RegisterProductController(productService)
	productRoute = routes.ProductRoute{}

	formService = services.NewFormService()
	formController = *controllers.RegisterFormController(formService)
	formRoute = routes.FormRoute{}

	productRoute.RegisterProductRoute(productController)
	helloWorldRoute.RegisterHelloWorldRoutes(helloWorldController)

	port := os.Getenv("SERVER_PORT")
	routeObject := routes.GetRouteObject()
	fmt.Printf("Starting the server in the port :%s \n", port)

	err = routeObject.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("some error occured trying to start the server at port :%s.error:%s\n", port, err.Error())
	}
}
