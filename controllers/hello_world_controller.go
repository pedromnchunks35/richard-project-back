package controllers

import (
	"richard-project-back/dtos"
	iService "richard-project-back/services/ServicesInterface"
	"richard-project-back/utils/apiResponses"

	"github.com/gin-gonic/gin"
)

type HelloWorldControllerImpl struct {
	HelloWorldService iService.IHelloWorldService
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

func RegisterHelloWorldControllerImpl(service iService.IHelloWorldService) *HelloWorldControllerImpl {
	return &HelloWorldControllerImpl{
		HelloWorldService: service,
	}
}
