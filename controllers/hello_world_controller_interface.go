package controllers

import "github.com/gin-gonic/gin"

type HelloWorldController interface {
	GetHelloWorld(context *gin.Context)
}
