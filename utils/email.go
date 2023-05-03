package utils

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"gin-gorm-postgres/initializers"
	"gin-gorm-postgres/models"
	"github.com/k3a/html2text"
	"gopkg.in/gomail.v2"
	"html/template"
	"io/fs"
	"log"
	"path/filepath"
)

type EmailData struct {
	URL       string
	FirstName string
	Subject   string
}

func ParseTemplateDir(dir string) (*template.Template, error) {
	var paths []string
	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
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

func SendEmail(user *models.User, data *EmailData) {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("could not load config", err)
	}

	from := config.EmailFrom
	smtpHost := config.SMTPHost
	smtpUser := config.SMTPUser
	smtpPass := config.SMTPPass
	SmtpPort := config.SMTPPort
	to := user.Email

	var body bytes.Buffer
	template, err := ParseTemplateDir("templates")
	if err != nil {
		log.Fatal("could not load template", err)
	}

	template.ExecuteTemplate(&body, "verificationCode.html", &data)

	m := gomail.NewMessage()

	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", data.Subject)
	m.SetBody("text/html", body.String())
	m.AddAlternative("text/plain", html2text.HTML2Text(body.String()))

	fmt.Println(m)
	fmt.Println(smtpHost, "host", smtpPass, "pass", smtpUser, "user")
	d := gomail.NewDialer(smtpHost, SmtpPort, smtpUser, smtpPass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		log.Fatal("Could not send email :", err)
	}
}
