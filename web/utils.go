package web

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"text/template"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	LogTrace   *log.Logger
	LogInfo    *log.Logger
	LogWarning *log.Logger
	LogError   *log.Logger
)

const logFileName = "logs.log"

// RenderTemplateToWriter realiza el render del template que se encuentre en el path especificado
func RenderTemplateToWriter(templatePath string, w http.ResponseWriter, data interface{}) error {
	t := template.Must(template.ParseFiles(templatePath))
	e := t.Execute(w, data)
	if e != nil {
		return e
	}
	return nil
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

type UUIDGeneratorUtils interface {
	// CreateUUIDV4 crea un UUID V4 con el package uuid
	CreateUUIDV4() string
}

type UUIDGeneratorUtilsImpl struct{}

func NewUUIDGeneratorUtils() UUIDGeneratorUtils {
	return UUIDGeneratorUtilsImpl{}
}
func (u UUIDGeneratorUtilsImpl) CreateUUIDV4() string {
	return uuid.NewV4().String()
}

func init() {

	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("Failed to open log file %s: Detalles: %v", logFileName, err))
	}

	LogTrace = log.New(os.Stdout,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	LogInfo = log.New(os.Stdout,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	LogWarning = log.New(io.MultiWriter(logFile, os.Stdout),
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	LogError = log.New(io.MultiWriter(logFile, os.Stderr),
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}
