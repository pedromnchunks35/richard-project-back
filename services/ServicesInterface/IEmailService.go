package iService

import (
	model "richard-project-back/repositories/Models"
)

type IEmailService interface {
	SendEmail(Mail *model.Email) bool
}
