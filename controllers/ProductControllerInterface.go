package controllers

import "github.com/gin-gonic/gin"

type ProductController interface {
	GetProduct(context *gin.Context)
}
