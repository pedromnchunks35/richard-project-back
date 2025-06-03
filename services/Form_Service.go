package services

import (
	"richard-project-back/dtos"
)

type FormServiceImpl struct{}

func NewFormServiceImpl() *FormServiceImpl {
	return &FormServiceImpl{}
}

func (h FormServiceImpl) Teste() string {
	return "GOOD"
}
func (h FormServiceImpl) UpdateForm() bool {
	return false
}

func (h FormServiceImpl) DeleteForm() bool {
	return false
}

func (h FormServiceImpl) CreateForm() bool {
	return false
}

func (h FormServiceImpl) GetForm(id int64) *dtos.Form {
	return &dtos.Form{}
}
