package utils

// SMTPConfigProvider provides SMTP config access methods.
type SMTPConfigProvider interface {
	Default() *SMTPConfig
	Config(name string) *SMTPConfig
}
