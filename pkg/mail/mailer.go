package mail

import (
	"fmt"
	"log"

	gomail "gopkg.in/gomail.v2"
)

// Config mewakili konfigurasi SMTP dasar.
type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

// Mailer bertanggung jawab mengirim email via SMTP.
type Mailer struct {
	dialer  *gomail.Dialer
	from    string
	enabled bool
}

// New membuat instance Mailer baru.
func New(cfg Config) *Mailer {
	enabled := cfg.Host != "" && cfg.Port > 0 && cfg.Username != ""
	var dialer *gomail.Dialer
	if enabled {
		dialer = gomail.NewDialer(cfg.Host, cfg.Port, cfg.Username, cfg.Password)
	}

	return &Mailer{
		dialer:  dialer,
		from:    cfg.From,
		enabled: enabled,
	}
}

// Send mengirim email sederhana.
func (m *Mailer) Send(to, subject, body string) error {
	if !m.enabled {
		log.Printf("mailer disabled, skip sending email to %s with subject %s", to, subject)
		return nil
	}

	msg := gomail.NewMessage()
	msg.SetHeader("From", m.from)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/plain", body)

	if err := m.dialer.DialAndSend(msg); err != nil {
		return fmt.Errorf("send mail: %w", err)
	}
	return nil
}
