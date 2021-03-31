package services

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/manicar2093/guianetThea/app/dao"
	"github.com/manicar2093/guianetThea/app/entities"
	"github.com/manicar2093/guianetThea/app/sessions"
	"github.com/manicar2093/guianetThea/app/utils"
	"gopkg.in/guregu/null.v4/zero"
)

type LoginError struct {
	ClientMessage, InternalMessage string
}

func (l LoginError) Error() string {
	return l.InternalMessage
}

type LoginService interface {
	// DoLogin realiza el login del usuario. Regresa error tipo LoginError si hubo un problema con los datos de login. Regresa un error si hubo un error general
	DoLogin(email, password string, w http.ResponseWriter, r *http.Request) error
	DoLogout(w http.ResponseWriter, r *http.Request) error
}

type LoginServiceImpl struct {
	userDao            dao.UserDao
	passwordUtils      utils.PasswordUtils
	session            sessions.SessionHandler
	detailsHostingDao  dao.DetailsHostingDao
	uuidGeneratorUtils utils.UUIDGeneratorUtils
}

func NewLoginService(userDao dao.UserDao, passwordUtils utils.PasswordUtils, session sessions.SessionHandler, detailsHostingDao dao.DetailsHostingDao, uuidGeneratorUtils utils.UUIDGeneratorUtils) LoginService {
	return &LoginServiceImpl{userDao, passwordUtils, session, detailsHostingDao, uuidGeneratorUtils}
}

func (l LoginServiceImpl) DoLogin(email, password string, w http.ResponseWriter, r *http.Request) error {
	// Validamos si usuario existe
	saved, e := l.userDao.FindUserByEmail(email)
	if e != nil {
		switch {
		case e == sql.ErrNoRows:
			return LoginError{ClientMessage: "Usuario y/o Contraseña incorrectos", InternalMessage: fmt.Sprintf("Usuario con email %v no existe", email)}
		default:
			return e
		}
	}

	// Validamos la contraseña correcta
	e = l.passwordUtils.ValidatePassword(saved.Password, password)
	if e != nil {
		return LoginError{ClientMessage: "Usuario y/o Contraseña incorrectos", InternalMessage: "Contraseña incorrecta"}
	}

	// Creamos el UUID para identificar la sesión
	u := l.uuidGeneratorUtils.CreateUUIDV4()

	// Creamos la sesión
	e = l.session.CreateNewSession(w, r, u)
	if e != nil {
		return e
	}

	// Registramos la nueva sesión
	e = l.detailsHostingDao.Save(&entities.DetailsHosting{UserID: saved.UserID, Host: r.RemoteAddr, SessionStart: time.Now(), SessionClosure: zero.NewTime(time.Now().Add(sessions.SessionDuration), true), UUID: u})
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
	detailsEndpointAndHostingDao dao.DetailsEndpointAndHostingDao
	detailsHostingDao            dao.DetailsHostingDao
	endpointDao                  dao.EndpointDao
	sessionHandler               sessions.SessionHandler
}

func NewRecordService(detailsEndpointAndHostingDao dao.DetailsEndpointAndHostingDao, detailsHostingDao dao.DetailsHostingDao, endpointDao dao.EndpointDao, sessionHandler sessions.SessionHandler) RecordService {
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

	e = r.detailsEndpointAndHostingDao.Save(&entities.DetailsEndpointAndHosting{DetailsHostingID: details.ID, EndpointID: endpoint.EndpointID})
	if e != nil {
		return e
	}

	return nil
}
