// controllers_test almacena todos los mock que se generan para el testing de los controlles.
package controllers

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/gorilla/mux"
	"github.com/manicar2093/guianetThea/app/mocks"
)

// Mocks
var (
	router               *mux.Router
	validatorService     mocks.ValidatorServiceMock
	userDao              mocks.UserDaoMock
	passwordUtils        mocks.PasswordUtilsMock
	catalogsService      mocks.CatalogsServiceMock
	loginRegistryService mocks.LoginRegistryServiceMock
)

func setUp() {
	router = mux.NewRouter()
	validatorService = mocks.ValidatorServiceMock{}
	userDao = mocks.UserDaoMock{}
	passwordUtils = mocks.PasswordUtilsMock{}
	catalogsService = mocks.CatalogsServiceMock{}
	loginRegistryService = mocks.LoginRegistryServiceMock{}
}

func serialize(t *testing.T, i interface{}) *bytes.Buffer {

	jsonB, e := json.Marshal(i)
	if e != nil {
		t.Fatal(e)
	}
	return bytes.NewBuffer(jsonB)

}
