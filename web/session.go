package web

import (
	"errors"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
)

const (
	dayInSeconds  = 86400
	hourInSeconds = 3600
)

// Session es el objeto con el que se puede realizar el manejo de las sesiones
var Session SessionHandler
var defaultSessionTime int

// SessionDuration indica el time.Duration de la sesión
var SessionDuration time.Duration

// SessionError es una estructura con la cual se declaran errores generales en el servicio de sesiones
var (
	ErrGetSession  = errors.New("Error al obtener la sessión")
	ErrSaveSession = errors.New("Error al guardar la sessión")
	sessionName    = "guianetthea"
)

const sessionValue = "session_id"

type SessionHandler interface {
	IsLoggedIn(w http.ResponseWriter, r *http.Request) bool
	GetUserID(w http.ResponseWriter, r *http.Request) (string, error)
	GetCurrentSession(w http.ResponseWriter, r *http.Request) (*sessions.Session, error)
	CreateNewSession(w http.ResponseWriter, r *http.Request, uuid string) error
}

// SessionHandlerImpl servicio para el manejo de sesiones
type SessionHandlerImpl struct {
	session *sessions.CookieStore
}

// IsLoggedIn indica si hay una sesión activa. Aun cuando haya un error al obtener la sesión redirigira al login
func (s *SessionHandlerImpl) IsLoggedIn(w http.ResponseWriter, r *http.Request) bool {
	current, e := s.GetUserID(w, r)
	if e != nil {
		return false
	}
	if current == "" {
		return false
	}
	return true
}

// GetUserID valida si hay una sesión activa. Si es así, regresa el ID del usuario guardado
func (s *SessionHandlerImpl) GetUserID(w http.ResponseWriter, r *http.Request) (string, error) {
	current, e := s.GetCurrentSession(w, r)
	if e != nil {
		return "", e
	}
	d, ok := current.Values[sessionValue]
	if !ok {
		return "", nil
	}
	sessionUUID := d.(string)
	return sessionUUID, nil
}

// GetCurrentSession obtiene la sesión actual.
// Este metodo respeta los detalles del metodo Get de sessions.Get() que son los siguientes:
//
// Get returns a session for the given name after adding it to the registry.
//
// It returns a new session if the sessions doesn't exist. Access IsNew on the session to check if it is an existing session or a new one.
//
// It returns a new session and an error if the session exists but could not be decoded.
func (s *SessionHandlerImpl) GetCurrentSession(w http.ResponseWriter, r *http.Request) (*sessions.Session, error) {
	current, e := s.session.Get(r, sessionName)
	if e != nil {
		return current, ErrGetSession
	}
	return current, nil
}

// CreateNewSession crea una nueva sesión con el userID que se recibe
func (s *SessionHandlerImpl) CreateNewSession(w http.ResponseWriter, r *http.Request, uuid string) error {
	session, e := s.GetCurrentSession(w, r)
	if e != nil {
		return e
	}
	session.Values[sessionValue] = uuid
	e = session.Save(r, w)
	if e != nil {
		return ErrSaveSession
	}

	return nil
}

func init() {
	SessionDuration = 8 * time.Hour
	defaultSessionTime = int(SessionDuration.Seconds())
	instance := &SessionHandlerImpl{session: sessions.NewCookieStore([]byte(GetEnvVar("SECRET-KEY", "a-session-key")))}
	instance.session.MaxAge(defaultSessionTime)
	Session = instance
}
