package controllers

import (
	"github.com/gin-gonic/gin"
	"richard-project-back/dtos"
	"richard-project-back/services"
	"richard-project-back/utils/apiResponses"
)

type HelloWorldControllerImpl struct {
	HelloWorldService services.HelloWorldService
}

func (h HelloWorldControllerImpl) GetHelloWorld(context *gin.Context) {
	helloWorld := &dtos.HelloWorld{}
	err := context.BindJSON(helloWorld)
	if err != nil {
		apiResponses.BadArgumentsResponse(context, err.Error())
		return
	}

	result := h.HelloWorldService.HelloWorld(helloWorld.Msg)
	apiResponses.SuccessResponse(context, result, "SUCCESS")
}

func RegisterHelloWorldControllerImpl(service services.HelloWorldService) *HelloWorldControllerImpl {
	return &HelloWorldControllerImpl{
		HelloWorldService: service,
	}
}
