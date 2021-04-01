// controllers_test almacena todos los mock que se generan para el testing de los controlles.
package controllers

import (
	"github.com/gorilla/mux"
	"github.com/manicar2093/guianetThea/app/mocks"
)

// Mocks
var (
	router           *mux.Router
	validatorService mocks.ValidatorServiceMock
	userDao          mocks.UserDaoMock
	passwordUtils    mocks.PasswordUtilsMock
	catalogsService  mocks.CatalogsServiceMock
)

func setUp() {
	router = mux.NewRouter()
	validatorService = mocks.ValidatorServiceMock{}
	userDao = mocks.UserDaoMock{}
	passwordUtils = mocks.PasswordUtilsMock{}
	catalogsService = mocks.CatalogsServiceMock{}
}
