package email

import (
	"api-project/configs"
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

type Email struct {
	config *configs.EmailConfig
}

func New(config *configs.EmailConfig) *Email {
	return &Email{
		config: config,
	}
}
func (_email *Email) Send(to string, subject string, text string) error {
	senderEmail := _email.config.Email
	senderPassword := _email.config.Password
	smtpHost := "smtp.yandex.ru"
	smtpPort := "587"

	recipientEmail := to

	e := email.NewEmail()
	e.From = fmt.Sprintf("%v <%s>", _email.config.Address, senderEmail)
	e.To = []string{recipientEmail}
	e.Subject = subject
	e.HTML = []byte(text)

	err := e.Send(fmt.Sprintf("%s:%s", smtpHost, smtpPort), smtp.PlainAuth("", senderEmail, senderPassword, smtpHost))
	if err != nil {
		fmt.Printf("Error sending email: %v\n", err)
	} else {
		fmt.Println("Email sent successfully!")
	}
	return err
}
