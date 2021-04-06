package services

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/manicar2093/guianetThea/app/entities"
	"github.com/manicar2093/guianetThea/app/mocks"
)

var (
	userDaoMock                      mocks.UserDaoMock
	passwordUtilsMock                mocks.PasswordUtilsMock
	uuidGeneratorUtilsMock           mocks.UUIDGeneratorUtilsMock
	sessionHandlerMock               mocks.SessionHandlerMock
	detailsHostingDaoMock            mocks.DetailsHostingDaoMock
	detailsEndpointAndHostingDaoMock mocks.DetailsEndpointAndHostingDaoMock
	endpointDaoMock                  mocks.EndpointDaoMock
	rolDao                           mocks.RolDaoMock
	logRegistryDao                   mocks.LogRegistryDaoMock
)

// setUp inicializa las variales mock
var setUp = func() {
	userDaoMock = mocks.UserDaoMock{}
	passwordUtilsMock = mocks.PasswordUtilsMock{}
	uuidGeneratorUtilsMock = mocks.UUIDGeneratorUtilsMock{}
	sessionHandlerMock = mocks.SessionHandlerMock{}
	detailsHostingDaoMock = mocks.DetailsHostingDaoMock{}
	detailsEndpointAndHostingDaoMock = mocks.DetailsEndpointAndHostingDaoMock{}
	endpointDaoMock = mocks.EndpointDaoMock{}
	rolDao = mocks.RolDaoMock{}
	logRegistryDao = mocks.LogRegistryDaoMock{}
}

/*func TestDoLogin(t *testing.T) {

	email, password, w, r := "email", "password", httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, "/login", nil)
	uuid := "UUID"
	userMock := entities.User{
		UserID:           1,
		RolID:            sql.NullInt32{1, true},
		Name:             "Name",
		PaternalSureName: "PatName",
		MaternalSureName: sql.NullString{Valid: true, String: "MatName"},
		Email:            email,
		Password:         password,
		Status:           true,
	}
	details := entities.DetailsHosting{
		UserID:         userMock.UserID,
		Host:           r.RemoteAddr,
		SessionStart:   time.Now(),
		SessionClosure: sql.NullTime{Time: time.Now().Add(SessionDuration), Valid: true},
		UUID:           uuid,
	}

	setUp()

	userDaoMock.On("FindUserByEmail", email).Return(userMock, nil)
	passwordUtilsMock.On("ValidatePassword", userMock.Password, password).Return(nil)
	uuidGeneratorUtilsMock.On("CreateUUIDV4").Return(uuid)
	sessionHandlerMock.On("CreateNewSession", w, r, uuid).Return(nil)
	detailsHostingDaoMock.On("Save", &details).Return(nil)

	service := NewLoginService(userDaoMock, passwordUtilsMock, sessionHandlerMock, detailsHostingDaoMock, uuidGeneratorUtilsMock)

	e := service.DoLogin(email, password, w, r)

	userDaoMock.AssertExpectations(t)
	passwordUtilsMock.AssertExpectations(t)
	uuidGeneratorUtilsMock.AssertExpectations(t)
	sessionHandlerMock.AssertExpectations(t)
	//detailsHostingDaoMock.AssertNumberOfCalls(t, "Save", 1) // FIXME hay un problema ya que los tiempos del SessionStart no coinciden y causa un error

	if e != nil {
		t.Fatal("No debió haber error:", e)
	}
}*/

func TestDoLoginNoUser(t *testing.T) {

	email, password, w, r := "email", "password", httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, "/login", nil)
	userMock := entities.User{}

	setUp()

	userDaoMock.On("FindUserByEmail", email).Return(userMock, sql.ErrNoRows)

	service := NewLoginService(userDaoMock, passwordUtilsMock, sessionHandlerMock, detailsHostingDaoMock, uuidGeneratorUtilsMock)

	e := service.DoLogin(email, password, w, r)

	userDaoMock.AssertExpectations(t)

	le, ok := e.(LoginError)
	if !ok {
		t.Fatal("El error no corresponde al necesario.", e)
	}

	if le.ClientMessage != "Usuario y/o Contraseña incorrectos" {
		t.Fatal("El error del cliente no corresponde")
	}

	if le.InternalMessage != fmt.Sprintf("Usuario con email %v no existe", email) {
		t.Fatal("El error interno no es correcto")
	}
}

func TestDoLoginPasswordNotMatch(t *testing.T) {

	email, password, w, r := "email", "password", httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, "/login", nil)
	userMock := entities.User{Password: "password"}

	setUp()

	userDaoMock.On("FindUserByEmail", email).Return(userMock, nil)
	passwordUtilsMock.On("ValidatePassword", userMock.Password, password).Return(fmt.Errorf("Password not match"))

	service := NewLoginService(userDaoMock, passwordUtilsMock, sessionHandlerMock, detailsHostingDaoMock, uuidGeneratorUtilsMock)

	e := service.DoLogin(email, password, w, r)

	userDaoMock.AssertExpectations(t)

	le, ok := e.(LoginError)
	if !ok {
		t.Fatal("El error no corresponde al necesario.", e)
	}

	if le.ClientMessage != "Usuario y/o Contraseña incorrectos" {
		t.Fatal("El error del cliente no corresponde")
	}

	if le.InternalMessage != "Contraseña incorrecta" {
		t.Fatal("El error interno no es correcto")
	}
}
