package utils

import (
	"bytes"
	"crypto/tls"
	"html/template"
	"log"

	"github.com/k3a/html2text"
	"github.com/mas-wig/post-api-1/config"
	"github.com/mas-wig/post-api-1/types"
	"gopkg.in/gomail.v2"
)

type EmailData struct {
	URL       string
	FirstName string
	Subject   string
}

func SendEmail(user *types.DBResponse, data *EmailData, tmpl *template.Template, tmplName string) error {
	config, _ := config.LoadConfig("..")

	var body bytes.Buffer
	if err := tmpl.ExecuteTemplate(&body, tmplName, &data); err != nil {
		log.Fatal("could not find any template : %w", err)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", config.EmailFrom)
	m.SetHeader("To", user.Email)
	m.SetHeader("Subject", data.Subject)
	m.SetBody("text/html", body.String())
	m.AddAlternative("text/plain", html2text.HTML2Text(body.String()))

	d := gomail.NewDialer(config.SMTPHost, config.SMTPPort, config.SMTPUser, config.SMTPPass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
