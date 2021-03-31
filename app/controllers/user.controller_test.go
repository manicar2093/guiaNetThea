package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/manicar2093/guianetThea/app/mocks"
	"github.com/manicar2093/guianetThea/app/models"
	"github.com/stretchr/testify/assert"
)

var (
	router           *mux.Router
	validatorService mocks.ValidatorServiceMock
	userDao          mocks.UserDaoMock
)

func setUp() {
	router = mux.NewRouter()
	validatorService = mocks.ValidatorServiceMock{}
	userDao = mocks.UserDaoMock{}
}

// TestCreateUser valida el happy path del proceso de guardado de un usuario
func TestCreateUser(t *testing.T) {
	setUp()
	target := "/admin/user/registry"
	data := map[string]interface{}{
		"name":              "Test",
		"paternal_surename": "TestPS",
		"maternal_surename": "TestMS",
		"email":             "TestEmail@mail.com",
		"rol_id":            12,
		"password":          "pass1",
		"password_confirm":  "pass1",
	}

	userCreationModel := models.CreateUserData{
		RolID:            12,
		Name:             "Test",
		PaternalSureName: "TestPS",
		MaternalSureName: "TestMS",
		Email:            "TestEmail@mail.com",
		Password:         "pass1",
		PasswordConfirm:  "pass1",
	}

	w, r := httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, target, serialize(t, &data))

	validatorService.On("Validate", userCreationModel).Return([]models.ErrorValidationDetail{}, true)
	userDao.On("SaveFromModel", userCreationModel).Return(nil)

	controller := NewUserController(userDao, validatorService)

	router.HandleFunc(target, controller.CreateUser)

	router.ServeHTTP(w, r)

	validatorService.AssertExpectations(t)
	userDao.AssertExpectations(t)

	assert.Equal(t, http.StatusCreated, w.Code, "Debió regresar un status 201")
}

// TestCreateUserValidatorError valida que se envien correctamente los mensajes de error cuando el body no cumple con lo requerido
func TestCreateUserValidatorError(t *testing.T) {
	setUp()
	target := "/admin/user/registry"
	data := map[string]interface{}{
		"name":              "",
		"paternal_surename": "",
		"maternal_surename": "TestMS",
		"email":             "TestEmail@mail.com",
		"rol_id":            12,
		"password":          "pass1",
		"password_confirm":  "pass1",
	}

	userCreationModel := models.CreateUserData{
		RolID:            12,
		Name:             "",
		PaternalSureName: "",
		MaternalSureName: "TestMS",
		Email:            "TestEmail@mail.com",
		Password:         "pass1",
		PasswordConfirm:  "pass1",
	}

	w, r := httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, target, serialize(t, &data))

	validatorService.On("Validate", userCreationModel).Return([]models.ErrorValidationDetail{
		{Field: "Name", Constraints: []string{
			"required",
		}},
		{Field: "PaternalSureName", Constraints: []string{
			"required",
		}},
	}, false)

	controller := NewUserController(userDao, validatorService)

	router.HandleFunc(target, controller.CreateUser)

	router.ServeHTTP(w, r)

	var res map[string]interface{}

	json.NewDecoder(w.Body).Decode(&res)

	assert.Equal(t, 2, len(res["errores"].([]interface{})), "La cantidad de errores no corresponde")
	assert.Equal(t, "La información es incorrecta. Favor de revisar la documentación", res["message"].(string), "La cantidad de errores no corresponde")

	validatorService.AssertExpectations(t)
	userDao.AssertExpectations(t)

	assert.Equal(t, http.StatusBadRequest, w.Code, "Debió regresar un status 400")
}

// TestCreateUserUserDaoError valida que se mande el error correcto cuando falla el UserDao
func TestCreateUserUserDaoError(t *testing.T) {
	setUp()
	target := "/admin/user/registry"
	data := map[string]interface{}{
		"name":              "Test",
		"paternal_surename": "TestPS",
		"maternal_surename": "TestMS",
		"email":             "TestEmail@mail.com",
		"rol_id":            12,
		"password":          "pass1",
		"password_confirm":  "pass1",
	}

	userCreationModel := models.CreateUserData{
		RolID:            12,
		Name:             "Test",
		PaternalSureName: "TestPS",
		MaternalSureName: "TestMS",
		Email:            "TestEmail@mail.com",
		Password:         "pass1",
		PasswordConfirm:  "pass1",
	}

	w, r := httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, target, serialize(t, &data))

	validatorService.On("Validate", userCreationModel).Return([]models.ErrorValidationDetail{}, true)
	userDao.On("SaveFromModel", userCreationModel).Return(fmt.Errorf("Un error en el UserDao"))

	controller := NewUserController(userDao, validatorService)

	router.HandleFunc(target, controller.CreateUser)

	router.ServeHTTP(w, r)

	validatorService.AssertExpectations(t)
	userDao.AssertExpectations(t)

	var res map[string]interface{}

	json.NewDecoder(w.Body).Decode(&res)

	assert.Equal(t, "Hubo un error al guardar el usuario. Favor de contactar a soporte", res["message"].(string), "El mensaje no es el correcto")

	assert.Equal(t, http.StatusInternalServerError, w.Code, "Debió regresar un status 500")
}

func serialize(t *testing.T, i interface{}) *bytes.Buffer {

	jsonB, e := json.Marshal(i)
	if e != nil {
		t.Fatal(e)
	}
	return bytes.NewBuffer(jsonB)

}
