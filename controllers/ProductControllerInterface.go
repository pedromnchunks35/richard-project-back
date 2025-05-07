package controllers

import "github.com/gin-gonic/gin"

type ProductController interface {
	Teste(context *gin.Context)
	GetProduct(context *gin.Context)
	InsertProduct(context *gin.Context)
	RemoveProduct(context *gin.Context)
	UpdateProduct(context *gin.Context)
}
