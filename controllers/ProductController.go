package controllers

import (
	"richard-project-back/dtos"
	"richard-project-back/services"
	"richard-project-back/utils/apiResponses"

	"github.com/gin-gonic/gin"
)

type ProductControllerImpl struct {
	ProductService services.ProductService
}

func (h ProductControllerImpl) GetProduct(context *gin.Context) {
	product := &dtos.Product{}
	err := context.BindJSON(product)
	if err != nil {
		apiResponses.BadArgumentsResponse(context, err.Error())
		return
	}

	result := h.ProductService.Teste()
	apiResponses.SuccessResponse(context, result, "SUCCESS")
}

func RegisterProductControllerImpl(service services.ProductService) *ProductControllerImpl {
	return &ProductControllerImpl{
		ProductService: service,
	}
}
