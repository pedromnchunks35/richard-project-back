package services

import (
	"log"
	"net/smtp"
	model "richard-project-back/repositories/Models"
	iService "richard-project-back/services/ServicesInterface"
)

type EmailService struct{}

func NewEmailService() iService.IEmailService {
	return &EmailService{}
}

func (s *EmailService) SendEmail(mail *model.Email) bool {
	if mail == nil {
		return false
	}
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	password := "kzqh ljye asqw tdlj"

	auth := smtp.PlainAuth("", mail.From, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, "richardprojectback@gmail.com", mail.To, []byte(mail.Message))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Email sent successfully!")
	return true
}
