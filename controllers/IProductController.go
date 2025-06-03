package controllers

import "github.com/gin-gonic/gin"

type IProductController interface {
	Teste(context *gin.Context)
	GetProduct(context *gin.Context)
	GetProductDetail(context *gin.Context)
	InsertProduct(context *gin.Context)
	DeleteProduct(context *gin.Context)
	UpdateProduct(context *gin.Context)
}
