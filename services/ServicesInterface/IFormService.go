package iService

import (
	"richard-project-back/dtos"
)

type IFormService interface {
	TesteForm() string
	UpdateForm(id int64, form *dtos.Form) bool
	DeleteForm(id int64) bool
	CreateForm(form *dtos.Form) bool
	GetForm(id int64) string
}
