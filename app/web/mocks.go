package web

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/mock"
)

// En este modulo se almacenan  todos los mock para el testing del sistema

type UserDaoMock struct {
	mock.Mock
}

func (u UserDaoMock) Save(user *User) error {
	args := u.Called(user)
	return args.Error(0)
}

func (u UserDaoMock) Delete(userID int32) error {
	args := u.Called(userID)
	return args.Error(0)
}

func (u UserDaoMock) FindUserByID(id int32) (User, error) {
	args := u.Called(id)
	return args.Get(0).(User), args.Error(1)
}

func (u UserDaoMock) FindUserByEmail(email string) (User, error) {
	args := u.Called(email)
	return args.Get(0).(User), args.Error(1)
}

type MiddlewareProviderMock struct {
	mock.Mock
}

func (m MiddlewareProviderMock) NeedsLoggedIn(h http.HandlerFunc) http.HandlerFunc {
	args := m.Called(h)
	return args.Get(0).(http.HandlerFunc)
}

type SessionHandlerMock struct {
	mock.Mock
}

func (s SessionHandlerMock) IsLoggedIn(w http.ResponseWriter, r *http.Request) bool {
	args := s.Called(w, r)
	return args.Bool(0)
}

func (s SessionHandlerMock) GetUserID(w http.ResponseWriter, r *http.Request) (string, error) {
	args := s.Called(w, r)
	return args.String(0), args.Error(1)
}

func (s SessionHandlerMock) GetCurrentSession(w http.ResponseWriter, r *http.Request) (*sessions.Session, error) {
	args := s.Called(w, r)
	return args.Get(0).(*sessions.Session), args.Error(1)
}

func (s SessionHandlerMock) CreateNewSession(w http.ResponseWriter, r *http.Request, uuid string) error {
	args := s.Called(w, r, uuid)
	return args.Error(0)
}

func (s SessionHandlerMock) AddFlashMessage(message FlashMessage, w http.ResponseWriter, r *http.Request) {
	s.Called(message, w, r)
}

func (s SessionHandlerMock) GetFlashMessages(w http.ResponseWriter, r *http.Request) []interface{} {
	args := s.Called(w, r)
	return args.Get(0).([]interface{})
}

type EndpointDaoMock struct {
	mock.Mock
}

func (e EndpointDaoMock) FindEndpointByName(name string) (Endpoint, error) {
	args := e.Called(name)
	return args.Get(0).(Endpoint), args.Error(1)
}

func (e EndpointDaoMock) FindEndpointByID(id int32) (Endpoint, error) {
	args := e.Called(id)
	return args.Get(0).(Endpoint), args.Error(1)
}

type DetailsHostingDaoMock struct {
	mock.Mock
}

// Save realiza el guardado de un DetailsHostingDaoImpl. Si el DetailsHostingDaoImpl no contiene un ID se guardará un nuevo registro. Si ID va lleno realizará el update del registro.
//
// Se debe considerar que el salvado de información solo contempla los campos id_user, host, session_start, session_closure y uuid. El update solo modifica los campos session_closure y  type_log_out
func (d DetailsHostingDaoMock) Save(details *DetailsHosting) error {
	args := d.Called(details)
	return args.Error(0)
}

func (d DetailsHostingDaoMock) FindDetailsHostingByUUID(uuid string) (DetailsHosting, error) {
	args := d.Called(uuid)
	return args.Get(0).(DetailsHosting), args.Error(1)
}

type DetailsEndpointAndHostingDaoMock struct {
	mock.Mock
}

func (d DetailsEndpointAndHostingDaoMock) Save(details *DetailsEndpointAndHosting) error {
	args := d.Called(details)
	return args.Error(0)
}

type UUIDGeneratorUtilsMock struct {
	mock.Mock
}

// CreateUUIDV4 crea un UUID V4 con el package uuid
func (u UUIDGeneratorUtilsMock) CreateUUIDV4() string {
	args := u.Called()
	return args.String(0)
}

type PasswordUtilsMock struct {
	mock.Mock
}

func (p PasswordUtilsMock) HashPassword(password []byte) (string, error) {
	args := p.Called(password)
	return args.String(0), args.Error(1)
}

func (p PasswordUtilsMock) ValidatePassword(hashed, password string) error {
	args := p.Called(hashed, password)
	return args.Error(0)
}

type LoginServiceMock struct {
	mock.Mock
}
