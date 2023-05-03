package main

import (
	"gopkg.in/gomail.v2"
)

func main() {
	// Set up the email message
	m := gomail.NewMessage()
	m.SetHeader("From", "your_email@example.com")
	m.SetHeader("To", "recipient@example.com")
	m.SetHeader("Subject", "Test email from Go")
	m.SetBody("text/plain", "This is the message body.")

	// Set up the SMTP dialer
	d := gomail.NewDialer("smtp.mailtrap.io", 2525, "6ce172b050b8d6", "f43d75802e2f0f")

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
