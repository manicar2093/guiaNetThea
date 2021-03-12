package web

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"
)

type LoginError struct {
	clientMessage, internalMessage string
}

func (l LoginError) Error() string {
	return l.internalMessage
}

type LoginService interface {
	// DoLogin realiza el login del usuario. Regresa error tipo LoginError si hubo un problema con los datos de login. Regresa un error si hubo un error general
	DoLogin(email, password string, w http.ResponseWriter, r *http.Request) error
	DoLogout(w http.ResponseWriter, r *http.Request) error
}

type LoginServiceImpl struct {
	userDao            UserDao
	passwordUtils      PasswordUtils
	session            SessionHandler
	detailsHostingDao  DetailsHostingDao
	uuidGeneratorUtils UUIDGeneratorUtils
}

func NewLoginService(userDao UserDao, passwordUtils PasswordUtils, session SessionHandler, detailsHostingDao DetailsHostingDao, uuidGeneratorUtils UUIDGeneratorUtils) LoginService {
	return &LoginServiceImpl{userDao, passwordUtils, session, detailsHostingDao, uuidGeneratorUtils}
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
	u := l.uuidGeneratorUtils.CreateUUIDV4()

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

func (l LoginServiceImpl) DoLogout(w http.ResponseWriter, r *http.Request) error {
	panic("Not implemented")
}

// RecordService se encarga de realizar el guardado de las paginas que se visitan con una sesión
type RecordService interface {
	RegisterPageVisited(w http.ResponseWriter, req *http.Request, page string) error
}

type RecordServiceImpl struct {
	detailsEndpointAndHostingDao DetailsEndpointAndHostingDao
	detailsHostingDao            DetailsHostingDao
	endpointDao                  EndpointDao
	sessionHandler               SessionHandler
}

func NewRecordService(detailsEndpointAndHostingDao DetailsEndpointAndHostingDao, detailsHostingDao DetailsHostingDao, endpointDao EndpointDao, sessionHandler SessionHandler) RecordService {
	return &RecordServiceImpl{detailsEndpointAndHostingDao, detailsHostingDao, endpointDao, sessionHandler}
}

func (r RecordServiceImpl) RegisterPageVisited(w http.ResponseWriter, req *http.Request, page string) error {
	uuid, e := r.sessionHandler.GetUserID(w, req)
	if e != nil {
		return e
	}

	details, e := r.detailsHostingDao.FindDetailsHostingByUUID(uuid)
	if e != nil {
		return e
	}

	endpoint, e := r.endpointDao.FindEndpointByName(page)
	if e != nil {
		return e
	}

	e = r.detailsEndpointAndHostingDao.Save(&DetailsEndpointAndHosting{DetailsHostingID: details.ID, EndpointID: endpoint.EndpointID})
	if e != nil {
		return e
	}

	return nil
}
