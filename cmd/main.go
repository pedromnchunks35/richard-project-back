package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"richard-project-back/controllers"
	"richard-project-back/routes"
	"richard-project-back/services"
)

var helloWorldService services.HelloWorldService
var helloWorldController controllers.HelloWorldController
var helloWorldRoute routes.HelloWorldRoute

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error loading the environment variables %s\n", err.Error())
	}

	helloWorldService = services.HelloWorldServiceImpl{}
	helloWorldController = controllers.RegisterHelloWorldControllerImpl(helloWorldService)
	helloWorldRoute = routes.HelloWorldRouteImpl{}

	helloWorldRoute.RegisterHelloWorldRoutes(helloWorldController)

	port := os.Getenv("SERVER_PORT")
	routeObject := routes.GetRouteObject()
	fmt.Printf("Starting the server in the port :%s \n", port)

	err = routeObject.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("some error occured trying to start the server at port :%s.error:%s\n", port, err.Error())
	}
}
