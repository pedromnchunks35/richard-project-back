package services

import (
	"richard-project-back/dtos"
)

type FormService interface {
	Teste() string
	UpdateForm() bool
	DeleteForm() bool
	CreateForm() bool
	GetForm(id int64) *dtos.Form
}
