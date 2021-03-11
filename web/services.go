package web

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

type RecordService interface{}

type LoginError struct {
	clientMessage, internalMessage string
}

func (l LoginError) Error() string {
	return l.internalMessage
}

type LoginService interface {
	// DoLogin realiza el login del usuario. Regresa error tipo LoginError si hubo un problema con los datos de login. Regresa un error si hubo un error general
	DoLogin(email, password string, w http.ResponseWriter, r *http.Request) error
}

type LoginServiceImpl struct {
	userDao           UserDao
	passwordUtils     PasswordUtils
	session           *SessionHandler
	detailsHostingDao DetailsHostingDao
}

func NewLoginService(userDao UserDao, passwordUtils PasswordUtils, session *SessionHandler, detailsHostingDao DetailsHostingDao) LoginService {
	return &LoginServiceImpl{userDao, passwordUtils, session, detailsHostingDao}
}

func (l LoginServiceImpl) DoLogin(email, password string, w http.ResponseWriter, r *http.Request) error {
	// Validamos si usuario existe
	saved, e := l.userDao.FindUserByEmail(email)
	if e != nil {
		switch {
		case e == sql.ErrNoRows:
			return LoginError{clientMessage: "Usuario y/o Contraseña incorrectos", internalMessage: fmt.Sprintf("Usuario con email %v no existe", email)}
		default:
			return e
		}
	}

	// Validamos la contraseña correcta
	e = l.passwordUtils.ValidatePassword(saved.Password, password)
	if e != nil {
		return LoginError{clientMessage: "Usuario y/o Contraseña incorrectos", internalMessage: "Contraseña incorrecta"}
	}

	// Creamos el UUID para identificar la sesión
	u := uuid.NewV4().String()

	// Creamos la sesión
	e = l.session.CreateNewSession(w, r, u)
	if e != nil {
		return e
	}

	// Registramos la nueva sesión
	e = l.detailsHostingDao.Save(&DetailsHosting{UserID: saved.UserID, Host: r.RemoteAddr, SessionStart: time.Now(), SessionClosure: sql.NullTime{Time: time.Now().Add(SessionDuration), Valid: true}, UUID: u})
	if e != nil {
		return e
	}
	return nil

}
