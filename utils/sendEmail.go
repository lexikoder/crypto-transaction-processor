package utils

import (
	// "log"
	"os"
	"github.com/go-mail/mail"
)

func SendOtp(toemailaddress string, otp string, expires int) error {
	email := os.Getenv("EMAIL")
	email_password := os.Getenv("EMAIL_APP_PASSWORD")
	htmlBody := GenerateOtpEmail("", otp, expires)
	// Create a new message
	m := mail.NewMessage()
	m.SetHeader("From", email)
	m.SetHeader("To", toemailaddress)
	m.SetHeader("Subject", "Hello from Go-Mail")
	m.SetBody("text/html", htmlBody)

	// Dialer configuration (SMTP)
	d := mail.NewDialer("smtp.gmail.com", 587, email, email_password)

	// Send email
	 err := d.DialAndSend(m)
	 if err != nil {
	  return err
	}
    
	return nil


	
}