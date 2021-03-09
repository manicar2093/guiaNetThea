package web

import (
	"fmt"
	"log"
	"net/http"
	"os"
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

// GetEnvVar obtiene una variable de entorno o regresa el string especificado
func GetEnvVar(env string, possible string) string {
	d, ok := os.LookupEnv(env)
	if !ok {
		return possible
	}
	return d
}

// GetPortFromEnvVar valida la existencia del puerto en el enviroment y lo regresa con el formato necesario
func GetPortFromEnvVar(envVar, possible string) string {
	p := GetEnvVar(envVar, possible)
	if p != possible {

		return fmt.Sprintf(":%s", p)
	}
	return fmt.Sprintf(":%s", possible)
}
