package controllers

import (
	"richard-project-back/dtos"
	iService "richard-project-back/services/ServicesInterface"
	"richard-project-back/utils/apiResponses"

	"github.com/gin-gonic/gin"
)

type FormController struct {
	FormService iService.IFormService
}

func (h FormController) GetForm(context *gin.Context) {
	form := &dtos.Form{}
	err := context.BindJSON(form)
	if err != nil {
		apiResponses.BadArgumentsResponse(context, err.Error())
		return
	}

	result := h.FormService.GetForm(form.Id)
	apiResponses.SuccessResponse(context, result, "SUCCESS")
}

func (h FormController) InsertForm(context *gin.Context) {
	form := &dtos.Form{}
	err := context.BindJSON(form)
	if err != nil {
		apiResponses.BadArgumentsResponse(context, err.Error())
		return
	}

	result := h.FormService.CreateForm(form)
	apiResponses.SuccessResponse(context, result, "SUCCESS")
}

func (h FormController) UpdateForm(context *gin.Context) {
	form := &dtos.Form{}
	err := context.BindJSON(form)
	if err != nil {
		apiResponses.BadArgumentsResponse(context, err.Error())
		return
	}

	result := h.FormService.UpdateForm(form.Id, form)
	apiResponses.SuccessResponse(context, result, "SUCCESS")
}

func (h FormController) DeleteForm(context *gin.Context) {
	form := &dtos.Form{}
	err := context.BindJSON(form)
	if err != nil {
		apiResponses.BadArgumentsResponse(context, err.Error())
		return
	}

	result := h.FormService.DeleteForm(form.Id)
	apiResponses.SuccessResponse(context, result, "SUCCESS")
}

func (h FormController) Teste(context *gin.Context) {
	result := h.FormService.TesteForm()
	apiResponses.SuccessResponse(context, result, "SUCCESS")
}

func RegisterFormController(service iService.IFormService) *FormController {
	return &FormController{
		FormService: service,
	}
}
