package controllers

import (
	"richard-project-back/dtos"
	"richard-project-back/helper"
	iService "richard-project-back/services/ServicesInterface"
	"richard-project-back/utils/apiResponses"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductService iService.IProductService
}

func (h ProductController) GetProduct(context *gin.Context) {
	product := &dtos.Product{}
	err := context.BindJSON(product)
	if err != nil {
		apiResponses.BadArgumentsResponse(context, err.Error())
		return
	}

	result := h.ProductService.GetProduct(1)
	apiResponses.SuccessResponse(context, result, "SUCCESS")
}

func (h ProductController) GetProductDetail(context *gin.Context) {
	product := &dtos.Product{}
	err := context.BindJSON(product)
	if err != nil {
		apiResponses.BadArgumentsResponse(context, err.Error())
		return
	}

	result := h.ProductService.GetProductDetail(1)
	apiResponses.SuccessResponse(context, result, "SUCCESS")
}

func (h ProductController) InsertProduct(context *gin.Context) {
	product := &dtos.Product{}
	err := context.BindJSON(product)
	if err != nil {
		apiResponses.BadArgumentsResponse(context, err.Error())
		return
	}

	result := h.ProductService.InsertProduct(helper.ConvertProductDtoToRepository(product))
	apiResponses.SuccessResponse(context, result, "SUCCESS")
}
func (h ProductController) DeleteProduct(context *gin.Context) {
	product := &dtos.Product{}
	err := context.BindJSON(product)
	if err != nil {
		apiResponses.BadArgumentsResponse(context, err.Error())
		return
	}

	result := h.ProductService.DeleteProduct(1)
	apiResponses.SuccessResponse(context, result, "SUCCESS")
}
func (h ProductController) UpdateProduct(context *gin.Context) {
	product := &dtos.Product{}
	err := context.BindJSON(product)
	if err != nil {
		apiResponses.BadArgumentsResponse(context, err.Error())
		return
	}

	result := h.ProductService.UpdateProduct(1, helper.ConvertProductDtoToRepository(product))
	apiResponses.SuccessResponse(context, result, "SUCCESS")
}
func (h ProductController) Teste(context *gin.Context) {
	product := &dtos.Product{}
	err := context.BindJSON(product)
	if err != nil {
		apiResponses.BadArgumentsResponse(context, err.Error())
		return
	}

	result := h.ProductService.TesteProduct()
	apiResponses.SuccessResponse(context, result, "SUCCESS")
}

func RegisterProductController(service iService.IProductService) *ProductController {
	return &ProductController{
		ProductService: service,
	}
}
