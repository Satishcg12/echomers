package services

import (
	"gopkg.in/gomail.v2"
)

type EmailService struct {
	SMTPHost string
	SMTPPort int
	Username string
	Password string
}

func NewEmailService(smtpHost string, smtpPort int, username, password string) *EmailService {
	return &EmailService{SMTPHost: smtpHost, SMTPPort: smtpPort, Username: username, Password: password}
}

func (s *EmailService) SendEmail(to, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", s.Username)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(s.SMTPHost, s.SMTPPort, s.Username, s.Password)
	return d.DialAndSend(m)
}
