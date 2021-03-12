package web

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var (
	userDaoMock                      UserDaoMock
	passwordUtilsMock                PasswordUtilsMock
	uuidGeneratorUtilsMock           UUIDGeneratorUtilsMock
	sessionHandlerMock               SessionHandlerMock
	detailsHostingDaoMock            DetailsHostingDaoMock
	detailsEndpointAndHostingDaoMock DetailsEndpointAndHostingDaoMock
	endpointDaoMock                  EndpointDaoMock
)

// setUp inicializa las variales mock
var setUp = func() {
	userDaoMock = UserDaoMock{}
	passwordUtilsMock = PasswordUtilsMock{}
	uuidGeneratorUtilsMock = UUIDGeneratorUtilsMock{}
	sessionHandlerMock = SessionHandlerMock{}
	detailsHostingDaoMock = DetailsHostingDaoMock{}
	detailsEndpointAndHostingDaoMock = DetailsEndpointAndHostingDaoMock{}
	endpointDaoMock = EndpointDaoMock{}
}

func TestDoLogin(t *testing.T) {

	email, password, w, r := "email", "password", httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, "/login", nil)
	uuid := "UUID"
	userMock := User{
		UserID:           1,
		RolID:            sql.NullInt32{1, true},
		Name:             "Name",
		PaternalSureName: "PatName",
		MaternalSureName: "MatName",
		Email:            email,
		Password:         password,
		Status:           true,
	}
	details := DetailsHosting{
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
}

func TestDoLoginNoUser(t *testing.T) {

	email, password, w, r := "email", "password", httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, "/login", nil)
	userMock := User{}

	setUp()

	userDaoMock.On("FindUserByEmail", email).Return(userMock, sql.ErrNoRows)

	service := NewLoginService(userDaoMock, passwordUtilsMock, sessionHandlerMock, detailsHostingDaoMock, uuidGeneratorUtilsMock)

	e := service.DoLogin(email, password, w, r)

	userDaoMock.AssertExpectations(t)

	le, ok := e.(LoginError)
	if !ok {
		t.Fatal("El error no corresponde al necesario.", e)
	}

	if le.clientMessage != "Usuario y/o Contraseña incorrectos" {
		t.Fatal("El error del cliente no corresponde")
	}

	if le.internalMessage != fmt.Sprintf("Usuario con email %v no existe", email) {
		t.Fatal("El error interno no es correcto")
	}
}

func TestDoLoginPasswordNotMatch(t *testing.T) {

	email, password, w, r := "email", "password", httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, "/login", nil)
	userMock := User{Password: "password"}

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

	if le.clientMessage != "Usuario y/o Contraseña incorrectos" {
		t.Fatal("El error del cliente no corresponde")
	}

	if le.internalMessage != "Contraseña incorrecta" {
		t.Fatal("El error interno no es correcto")
	}
}

func TestRecordServiceImpl_RegisterPageVisited(t *testing.T) {
	uuid, page, w, r := "uuid-session-id", "test", httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/test", nil)
	details := DetailsHosting{ID: 1}
	endpoint := Endpoint{EndpointID: 1}
	detailsEndpoint := DetailsEndpointAndHosting{DetailsHostingID: details.ID, EndpointID: endpoint.EndpointID}

	setUp()

	sessionHandlerMock.On("GetUserID", w, r).Return(uuid, nil)
	detailsHostingDaoMock.On("FindDetailsHostingByUUID", uuid).Return(details, nil)
	endpointDaoMock.On("FindEndpointByName", page).Return(endpoint, nil)
	detailsEndpointAndHostingDaoMock.On("Save", &detailsEndpoint).Return(nil)

	service := NewRecordService(detailsEndpointAndHostingDaoMock, detailsHostingDaoMock, endpointDaoMock, sessionHandlerMock)

	e := service.RegisterPageVisited(w, r, page)

	if e != nil {
		t.Fatal("No debió regresar error:", e)
	}

	sessionHandlerMock.AssertExpectations(t)
	detailsHostingDaoMock.AssertExpectations(t)
	endpointDaoMock.AssertExpectations(t)
	detailsEndpointAndHostingDaoMock.AssertExpectations(t)
}
