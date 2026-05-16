package worker

import (
	"bytes"
	"html/template"
	"net/smtp"

	"/pkg/utils"
)

// EmailJob represents an email to be sent asynchronously.
type EmailJob struct {
	To       string
	Subject  string
	Body     string
	HTML     bool
	Template string
	Data     interface{}
}

// Worker processes email jobs asynchronously via a buffered channel.
type Worker struct {
	queue  chan EmailJob
	utils  *utils.Utils
}

// NewWorker creates a new Worker with the given buffer size.
func NewWorker(bufferSize int) *Worker {
	return &Worker{
		queue: make(chan EmailJob, bufferSize),
		utils: utils.GetUtils(),
	}
}

// Start begins processing the email queue. Run in a goroutine.
func (w *Worker) Start() {
	for job := range w.queue {
		w.processEmailJob(job)
	}
}

// SendEmailAsync enqueues a plain text email.
func (w *Worker) SendEmailAsync(to, subject, body string) {
	w.queue <- EmailJob{To: to, Subject: subject, Body: body, HTML: false}
}

// SendEmailHTMLAsync enqueues an HTML email rendered from a template.
func (w *Worker) SendEmailHTMLAsync(to, subject, tmplName string, data interface{}) {
	w.queue <- EmailJob{To: to, Subject: subject, Template: tmplName, Data: data, HTML: true}
}

func (w *Worker) processEmailJob(job EmailJob) {
	u := w.utils
	smtpCfg := u.Config.SMTPs.Default()
	if smtpCfg == nil {
		u.Logger.Errorf("No default SMTP config found")
		return
	}

	body := job.Body
	if job.HTML && job.Template != "" {
		tmpl, err := template.New(job.Template).Parse(job.Template) // simplified
		if err != nil {
			u.Logger.Errorf("Parse email template: %s", err.Error())
			return
		}
		var buf bytes.Buffer
		if err := tmpl.Execute(&buf, job.Data); err != nil {
			u.Logger.Errorf("Execute email template: %s", err.Error())
			return
		}
		body = buf.String()
	}

	msg := []byte("To: " + job.To + "\r\n" +
		"Subject: " + job.Subject + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
		"\r\n" + body)

	addr := smtpCfg.Host + ":" + smtpCfg.Port
	auth := smtp.PlainAuth("", smtpCfg.Username, smtpCfg.Password, smtpCfg.Host)

	if err := smtp.SendMail(addr, auth, smtpCfg.From, []string{job.To}, msg); err != nil {
		u.Logger.Errorf("Send email to %s: %s", job.To, err.Error())
	} else {
		u.Logger.Infof("Email sent to %s: %s", job.To, job.Subject)
	}
}
