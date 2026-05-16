package utils

import (
	"net/smtp"
)

// SendEmail sends a plain text email via SMTP.
func SendEmail(cfg *SMTPConfig, to, subject, body string) error {
	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Content-Type: text/plain; charset=\"UTF-8\"\r\n" +
		"\r\n" + body)
	auth := smtp.PlainAuth("", cfg.Username, cfg.Password, cfg.Host)
	return smtp.SendMail(cfg.Host+":"+cfg.Port, auth, cfg.From, []string{to}, msg)
}

// SendHTMLEmail sends an HTML email via SMTP.
func SendHTMLEmail(cfg *SMTPConfig, to, subject, htmlBody string) error {
	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
		"\r\n" + htmlBody)
	auth := smtp.PlainAuth("", cfg.Username, cfg.Password, cfg.Host)
	return smtp.SendMail(cfg.Host+":"+cfg.Port, auth, cfg.From, []string{to}, msg)
}

// IsValidEmail performs a basic email format check.
func IsValidEmail(email string) bool {
	if len(email) < 5 || len(email) > 254 {
		return false
	}
	at := -1
	for i, c := range email {
		if c == '@' {
			if at >= 0 {
				return false
			}
			at = i
		}
	}
	return at > 0 && at < len(email)-1
}
