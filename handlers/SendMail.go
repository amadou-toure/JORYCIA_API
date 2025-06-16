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

	d := gomail.NewDialer("smtp.jorycia.ca", 587, "support@jorycia.ca", "TON_MOT_DE_PASSE")

	return d.DialAndSend(m)
}