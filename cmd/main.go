package main

import (
	"fmt"
	"log"
	"os"
	sql "richard-project-back/configs/connector/postgresql"
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

	fmt.Println("----------------STARTING MAIN--------------")
	err := godotenv.Load(".env")
	fmt.Println(".ENV:", err)
	if err != nil {
		log.Fatalf("error loading the environment variables %s\n", err.Error())
	}
	dir, err := os.Getwd()
	fmt.Println("Current working directory:", dir)

	vault := &sql.PostgresDbVault{}
	connStr := os.Getenv("DATABASE_URL")
	vault.InitDbConnection(connStr)

	helloWorldService = services.HelloWorldServiceImpl{}
	helloWorldController = controllers.RegisterHelloWorldControllerImpl(helloWorldService)
	helloWorldRoute = routes.HelloWorldRouteImpl{}

	productService = services.NewProductService(vault)
	productController = *controllers.RegisterProductController(productService)
	productRoute = routes.ProductRoute{}

	formService = services.NewFormService(vault)
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
