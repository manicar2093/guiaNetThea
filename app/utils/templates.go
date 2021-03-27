package utils

import (
	"net/http"
	"text/template"
)

type TemplateUtils interface {
}

// RenderTemplateToWriter realiza el render del template que se encuentre en el path especificado
func RenderTemplateToWriter(templatePath string, w http.ResponseWriter, data interface{}) error {
	t := template.Must(template.ParseFiles(templatePath))
	e := t.Execute(w, data)
	if e != nil {
		return e
	}
	return nil
}
