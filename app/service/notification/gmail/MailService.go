package gmail

import (
	"OrderApp/service/notification/gmail/model"
	"OrderApp/service/notification/port"
	"bytes"
	"html/template"

	"gopkg.in/gomail.v2"
)

type MailServicePort struct {
	d         *gomail.Dialer
	templates map[string]*template.Template
}

func NewMailService() (port.MailServicePort, error) {
	d, t, e := Config()
	if e != nil {
		return nil, e
	}
	return &MailServicePort{
		d:         d,
		templates: t,
	}, nil
}

func (s *MailServicePort) SendSuccessfullyOrderPlayed(to string, data model.SuccessfullyOrderPlayedData) {
	s.sendGMail(
		to,
		"Order Updated",
		"successfully_order_played",
		data)
}

func (s *MailServicePort) sendGMail(to string, subject string, template string, data interface{}) {
	tmpl := s.templates[template]
	if tmpl == nil {
		panic("template not found: " + template)
	}
	var body bytes.Buffer
	if e := tmpl.Execute(&body, data); e != nil {
		panic(e)
	}
	m := gomail.NewMessage()
	m.SetHeader("From", "luongkhacnam222@gmail.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body.String())
	e := s.d.DialAndSend(m)
	if e != nil {
		panic(e)
	}
}
