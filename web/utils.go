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
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

const (
	logFileTrace = "trace.log"
	logFileInfo  = "info.log"
	logFileWarn  = "warn.log"
	logFileError = "error.log"
	logFileGral  = "logs.log"
)

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

	logTrace, err := os.OpenFile(logFileTrace, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("Failed to open log file %s: Detalles: %v", logFileTrace, err))
	}
	logInfo, err := os.OpenFile(logFileInfo, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("Failed to open log file %s: Detalles: %v", logFileInfo, err))
	}
	logWarn, err := os.OpenFile(logFileWarn, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("Failed to open log file %s: Detalles: %v", logFileWarn, err))
	}
	logError, err := os.OpenFile(logFileError, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("Failed to open log file %s: Detalles: %v", logFileError, err))
	}
	logGral, err := os.OpenFile(logFileGral, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("Failed to open log file %s: Detalles: %v", logFileGral, err))
	}

	Trace = log.New(io.MultiWriter(logTrace, logGral, os.Stdout),
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(io.MultiWriter(logInfo, logGral, os.Stdout),
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(io.MultiWriter(logWarn, logGral, os.Stdout),
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(io.MultiWriter(logError, logGral, os.Stderr),
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}
