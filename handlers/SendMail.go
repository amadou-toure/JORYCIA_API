package handlers

import (
	"gopkg.in/gomail.v2"
)

func SendMail(From,to, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", From)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer("mail.jorycia.ca", 465, "support@jorycia.ca", "i_eat_ass")

	return d.DialAndSend(m)
}