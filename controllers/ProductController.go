package controllers

import (
	"richard-project-back/convert"
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

	result := h.ProductService.GetProduct(1)
	apiResponses.SuccessResponse(context, result, "SUCCESS")
}
func (h ProductControllerImpl) InsertProduct(context *gin.Context) {
	product := &dtos.Product{}
	err := context.BindJSON(product)
	if err != nil {
		apiResponses.BadArgumentsResponse(context, err.Error())
		return
	}

	result := h.ProductService.InsertProduct(convert.ConvertProductDtoToRepository(product))
	apiResponses.SuccessResponse(context, result, "SUCCESS")
}
func (h ProductControllerImpl) RemoveProduct(context *gin.Context) {
	product := &dtos.Product{}
	err := context.BindJSON(product)
	if err != nil {
		apiResponses.BadArgumentsResponse(context, err.Error())
		return
	}

	result := h.ProductService.RemoveProduct(1)
	apiResponses.SuccessResponse(context, result, "SUCCESS")
}
func (h ProductControllerImpl) UpdateProduct(context *gin.Context) {
	product := &dtos.Product{}
	err := context.BindJSON(product)
	if err != nil {
		apiResponses.BadArgumentsResponse(context, err.Error())
		return
	}

	result := h.ProductService.UpdateProduct(1, convert.ConvertProductDtoToRepository(product))
	apiResponses.SuccessResponse(context, result, "SUCCESS")
}
func (h ProductControllerImpl) Teste(context *gin.Context) {
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
