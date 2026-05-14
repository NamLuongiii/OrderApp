package mail

import (
	"bytes"
	"html/template"

	"gopkg.in/gomail.v2"
)

type Service interface {
	SendNewOrderPlayed(to string, data SendNewOrderPlayedCommand)
}

type ServiceImpl struct {
	templates map[string]*template.Template
	d         *gomail.Dialer
}

func NewMailService() (Service, error) {
	t, e := loadTemplates()
	if e != nil {
		return nil, e
	}
	d := gomail.NewDialer(
		"smtp.gmail.com",
		587,
		"luongkhacnam222@gmail.com",
		"qoto iesp pzte lixx")

	return &ServiceImpl{
		templates: t,
		d:         d,
	}, e

}

func (s *ServiceImpl) sendGMail(to string, subject string, template string, data interface{}) {
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

func loadTemplates() (map[string]*template.Template, error) {
	templateNames := []string{
		"template/send_new_order_played.html",
	}

	templates := make(map[string]*template.Template, len(templateNames))
	for _, name := range templateNames {
		templates[name] = template.Must(template.ParseFiles(name))
	}
	return templates, nil
}
