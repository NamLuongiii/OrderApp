package gmail

import (
	"html/template"

	"gopkg.in/gomail.v2"
)

func Config() (*gomail.Dialer, map[string]*template.Template, error) {
	t, e := loadTemplates()
	if e != nil {
		return nil, nil, e
	}
	return gomail.NewDialer(
		"smtp.gmail.com",
		587,
		"luongkhacnam222@gmail.com",
		"qoto iesp pzte lixx"), t, nil
}

func loadTemplates() (map[string]*template.Template, error) {
	path := "service/notification/gmail/template/"
	templateNames := []string{
		"successfully_order_played",
	}

	templates := make(map[string]*template.Template, len(templateNames))
	for _, name := range templateNames {
		templates[name] = template.Must(template.ParseFiles(path + name + ".html"))
	}
	return templates, nil
}
