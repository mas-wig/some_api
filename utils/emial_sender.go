package utils

import (
	"bytes"
	"crypto/tls"
	"html/template"
	"log"
	"os"
	"path/filepath"

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

func ParseTemplateDir(dir string) (*template.Template, error) {
	var paths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return template.ParseFiles(paths...)
}

func SendEmail(user *types.DBResponse, data *EmailData, tmplName string) error {
	var body bytes.Buffer
	config, _ := config.LoadConfig("..")
	template, err := ParseTemplateDir("public/templates")
	if err != nil {
		log.Fatal("Could not parse template", err)
	}

	template = template.Lookup(tmplName)
	template.Execute(&body, &data)

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
