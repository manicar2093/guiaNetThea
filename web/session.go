package web

import (
	"net/http"

	"github.com/gorilla/sessions"
)

// Session es el objeto con el que se puede realizar el manejo de las sesiones
var Session *SessionHandler
var sessionName = "guianetthea-session"

// SessionError es una estructura con la cual se declaran errores generales en el servicio de sesiones
type SessionError struct {
	message string
	e       error
}

func (se SessionError) Error() string {
	return se.message
}

// SessionHandler servicio para el manejo de sesiones
type SessionHandler struct {
	session *sessions.CookieStore
}

// IsLoggedIn valida si hay una sesión activa. Si es así, regresa el ID del usuario guardado
func (s *SessionHandler) IsLoggedIn(w http.ResponseWriter, r *http.Request) (int, error) {
	current, e := s.GetCurrentSession(w, r)
	if e != nil {
		return 0, e
	}
	d, ok := current.Values["userId"]
	if !ok {
		return 0, nil
	}
	userID := d.(int)
	return userID, nil
}

// GetCurrentSession obtiene la sesión actual.
// Este metodo respeta los detalles del metodo Get de sessions.Get() que son los siguientes:
//
// Get returns a session for the given name after adding it to the registry.
//
// It returns a new session if the sessions doesn't exist. Access IsNew on the session to check if it is an existing session or a new one.
//
// It returns a new session and an error if the session exists but could not be decoded.
func (s *SessionHandler) GetCurrentSession(w http.ResponseWriter, r *http.Request) (*sessions.Session, error) {
	current, e := s.session.Get(r, sessionName)
	if e != nil {
		return current, SessionError{"Error al obtener la sessión", e}
	}
	return current, nil
}

// CreateNewSession crea una nueva sesión con el userID que se recibe
func (s *SessionHandler) CreateNewSession(w http.ResponseWriter, r *http.Request, userId int) error {
	session, e := s.GetCurrentSession(w, r)
	if e != nil {
		return e
	}
	session.Values["userId"] = userId
	e = session.Save(r, w)
	if e != nil {
		return SessionError{"Error al guardar la sessión", e}
	}

	return nil
}

func init() {
	Session = &SessionHandler{session: sessions.NewCookieStore([]byte("a-session-key"))}
}
