package controllers

import "github.com/gin-gonic/gin"

type IFormController interface {
	GetForm(context *gin.Context)
	InsertForm(context *gin.Context)
	UpdateForm(context *gin.Context)
	DeleteForm(context *gin.Context)
	Teste(context *gin.Context)
}
