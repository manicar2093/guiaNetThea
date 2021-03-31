package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	ErrExecution        = errors.New("error al ejecutar el template")
	ErrTemplateNotFound = errors.New("template was not found")
)

type TemplateUtils interface {
	// RenderTemplateToWriter renderiza el template en el path especificado. Si no se encuentra la pagina o hay un error al realizar el render se usa el metodo http.Error para especificar en el ResponseWriter el error con su
	// respectivo StatusCode y, además, regresa el error que se presentó
	RenderTemplateToResponseWriter(templatePath string, w http.ResponseWriter, data interface{}) error
}

type TemplateUtilsImpl struct{}

func NewTemplateUtils() TemplateUtils {
	return &TemplateUtilsImpl{}
}

func (t TemplateUtilsImpl) RenderTemplateToResponseWriter(templatePath string, w http.ResponseWriter, data interface{}) error {
	name := filepath.Base(templatePath)

	tpl, e := t.createTemplate(name, templatePath)

	if e != nil {
		Error.Printf("Error al parsear el template '%s': Detalles: \n\t%v", templatePath, e)
		http.Error(w, "No se encontró la ruta especificada", http.StatusNotFound)
		return ErrTemplateNotFound
	}

	if e = tpl.Execute(w, data); e != nil {
		Error.Printf("Error al ejecutar el template '%s': Detalles: \n\t%v", templatePath, e)
		http.Error(w, "Error Interno del Servidor", http.StatusInternalServerError)
		return ErrExecution
	}
	return nil
}

func (t TemplateUtilsImpl) createTemplate(name, templatePath string) (*template.Template, error) {
	return template.New(name).Funcs(
		template.FuncMap{
			"ToJSON": func(i interface{}) string {
				b, e := json.Marshal(i)
				if e != nil {
					panic(fmt.Sprintf("Error al realizar el Marshal a JSON del Usuario: Detalles: \n\t%v", e))
				}
				return string(b)
			},
		},
	).ParseFiles(templatePath)
}
