package utils

import (
	"errors"
	"net/http"
	"text/template"
)

var (
	ErrExecution        = errors.New("error al ejecutar el template")
	ErrTemplateNotFound = errors.New("template was not found")
)

type TemplateUtils interface {
	// RenderTemplateToWriter renderiza el template en el path especificado. Si este no se encuentra el ResponseWriter se le colocará el status 404.
	RenderTemplateToResponseWriter(templatePath string, w http.ResponseWriter, data interface{}) error
}

type TemplateUtilsImpl struct{}

func NewTemplateUtils() TemplateUtils {
	return &TemplateUtilsImpl{}
}

func (t TemplateUtilsImpl) RenderTemplateToResponseWriter(templatePath string, w http.ResponseWriter, data interface{}) error {
	tpl, e := template.ParseFiles(templatePath)
	if e != nil {
		Error.Printf("Error al parsear el template '%s': Detalles: \n\t%v", templatePath, e)
		return ErrTemplateNotFound
	}

	if e = tpl.Execute(w, data); e != nil {
		Error.Printf("Error al ejecutar el template '%s': Detalles: \n\t%v", templatePath, e)
		return ErrExecution
	}
	return nil
}
