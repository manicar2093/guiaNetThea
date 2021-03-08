package web

import (
	"log"
	"net/http"
	"text/template"
)

// RenderTemplateToWriter realiza el render del template que se encuentre en el path especificado
func RenderTemplateToWriter(templatePath string, w http.ResponseWriter, data interface{}) {
	t := template.Must(template.ParseFiles(templatePath))
	e := t.Execute(w, nil)
	if e != nil {
		log.Println(e)
	}
}
