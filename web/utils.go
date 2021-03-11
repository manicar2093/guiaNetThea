package web

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"golang.org/x/crypto/bcrypt"
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

type PasswordUtils interface {
	HashPassword(password []byte) (string, error)
	ValidatePassword(hashed, password string) error
}

type PasswordUtilsImpl struct{}

func NewPasswordUtils() *PasswordUtilsImpl {
	return &PasswordUtilsImpl{}
}

// HashPassword crea un hash con el algoritmo seleccionado para el sistema
func (p PasswordUtilsImpl) HashPassword(password []byte) (string, error) {
	b, e := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if e != nil {
		return "", fmt.Errorf("An error occurred creating user's password")
	}
	return string(b), nil
}

// ValidatePassword valida que la contraseña recibida coincida con el hash proporcionado
func (p PasswordUtilsImpl) ValidatePassword(hashed, password string) error {
	e := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	if e != nil {
		return fmt.Errorf("Invalid password")
	}
	return nil
}
