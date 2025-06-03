package services

import (
	"richard-project-back/dtos"
	queryList "richard-project-back/repositories/Queries"
	iService "richard-project-back/services/ServicesInterface"
)

type FormService struct{}

func NewFormService() iService.IFormService {
	return &FormService{}
}

func (s *FormService) TesteForm() string {
	return "GOOD"
}

func (s *FormService) CreateForm(form *dtos.Form) bool {
	if form == nil {
		return false
	}
	query := queryList.QueryInsertForm
	result := ExecuteQuery(query)
	return result == ""
}

func (s *FormService) UpdateForm(id int64, form *dtos.Form) bool {
	if id == 0 || form == nil {
		return false
	}
	query := queryList.QueryUpdateForm
	result := ExecuteQuery(query)
	return result == ""
}

func (s *FormService) DeleteForm(id int64) bool {
	if id == 0 {
		return false
	}
	query := queryList.QueryDeleteForm
	result := ExecuteQuery(query)
	return result == ""
}

func (s *FormService) GetForm(id int64) string {
	if id == 0 {
		return ""
	}
	query := queryList.QueryGetFormWithProducts
	result := ExecuteQuery(query)
	return result
}
