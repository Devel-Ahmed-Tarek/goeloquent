package goeloquent

import (
	"fmt"
	"log"

	"gopkg.in/gomail.v2"
)

// EmailConfig يحتوي على إعدادات البريد الإلكتروني
type EmailConfig struct {
	SMTPHost    string
	SMTPPort    int
	Username    string
	Password    string
	FromAddress string
	FromName    string
}

// EmailService يمثل خدمة إرسال البريد الإلكتروني
type EmailService struct {
	Config EmailConfig
}

// NewEmailService ينشئ كائن EmailService جديد
func NewEmailService(config EmailConfig) *EmailService {
	return &EmailService{Config: config}
}

// SendEmail يرسل رسالة بريد إلكتروني إلى المستلم المحدد
func (es *EmailService) SendEmail(to, subject, body string) error {
	m := gomail.NewMessage()
	if es.Config.FromName != "" {
		m.SetHeader("From", m.FormatAddress(es.Config.FromAddress, es.Config.FromName))
	} else {
		m.SetHeader("From", es.Config.FromAddress)
	}
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(es.Config.SMTPHost, es.Config.SMTPPort, es.Config.Username, es.Config.Password)
	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("❌ Failed to send email: %v", err)
	}
	log.Println("✅ Email sent successfully!")
	return nil
}
