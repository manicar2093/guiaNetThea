package controllers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/manicar2093/guianetThea/app/entities"
	"github.com/manicar2093/guianetThea/app/mocks"
	"github.com/manicar2093/guianetThea/app/models"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v4/zero"
)

// Mocks
var (
	router           *mux.Router
	validatorService mocks.ValidatorServiceMock
	userDao          mocks.UserDaoMock
)

// Data for testing
var (
	creationData = map[string]interface{}{
		"name":              "Test",
		"paternal_surename": "TestPS",
		"maternal_surename": "TestMS",
		"email":             "TestEmail@mail.com",
		"rol_id":            12,
		"password":          "pass1",
		"password_confirm":  "pass1",
	}

	userCreationModel = models.CreateUserData{
		RolID:            12,
		Name:             "Test",
		PaternalSureName: "TestPS",
		MaternalSureName: "TestMS",
		Email:            "TestEmail@mail.com",
		Password:         "pass1",
		PasswordConfirm:  "pass1",
	}

	userUpdateEntityMock = entities.User{
		UserID:           userUpdateModel.ID,
		RolID:            zero.IntFrom(int64(userUpdateModel.RolID)),
		Name:             userUpdateModel.Name,
		PaternalSureName: userUpdateModel.PaternalSureName,
		MaternalSureName: zero.NewString(userUpdateModel.MaternalSureName, true),
		Email:            userUpdateModel.Email,
		CreationDate:     time.Now(),
		EditDate:         zero.NewTime(time.Now(), true),
	}

	updateData = map[string]interface{}{
		"id":                1,
		"name":              "Test",
		"paternal_surename": "TestPS",
		"maternal_surename": "TestMS",
		"email":             "TestEmail@mail.com",
		"rol_id":            12,
	}

	userUpdateModel = models.UpdateUserData{
		ID:               1,
		RolID:            12,
		Name:             "Test",
		PaternalSureName: "TestPS",
		MaternalSureName: "TestMS",
		Email:            "TestEmail@mail.com",
	}

	userRestorePassEntityMock = entities.User{
		UserID:           1,
		RolID:            zero.IntFrom(int64(1)),
		Name:             "TestName",
		PaternalSureName: "PaternalSureName",
		MaternalSureName: zero.NewString("MaternalSureName", true),
		Email:            "Email@mail.com",
		CreationDate:     time.Now(),
		EditDate:         zero.NewTime(time.Now(), true),
		Password:         restorePasswordModel.Password,
	}

	restorePasswordData = map[string]interface{}{
		"id":               1,
		"password":         "TestPass",
		"password_confirm": "TestPass",
	}

	restorePasswordModel = models.RestorePasswordData{
		ID:              1,
		Password:        "TestPass",
		PasswordConfirm: "TestPass",
	}
)

func setUp() {
	router = mux.NewRouter()
	validatorService = mocks.ValidatorServiceMock{}
	userDao = mocks.UserDaoMock{}
}

/*
	Create User
*/

// TestCreateUser valida el happy path del proceso de guardado de un usuario
func TestCreateUser(t *testing.T) {
	setUp()
	target := "/admin/user/registry"

	w, r := httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, target, serialize(t, &creationData))

	validatorService.On("Validate", userCreationModel).Return([]models.ErrorValidationDetail{}, true)
	userDao.On("SaveFromModel", userCreationModel).Return(1, nil)

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

	w, r := httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, target, serialize(t, &creationData))

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

	w, r := httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, target, serialize(t, &creationData))

	validatorService.On("Validate", userCreationModel).Return([]models.ErrorValidationDetail{}, true)
	userDao.On("SaveFromModel", userCreationModel).Return(0, fmt.Errorf("Un error en el UserDao"))

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

/*
	Upsdate User
*/

// TestUpdateUser valida el happy path del proceso
func TestUpdateUser(t *testing.T) {

	setUp()
	target := "/admin/user/update"

	w, r := httptest.NewRecorder(), httptest.NewRequest(http.MethodPut, target, serialize(t, &updateData))

	validatorService.On("Validate", userUpdateModel).Return([]models.ErrorValidationDetail{}, true)
	userDao.On("FindUserByID", userUpdateModel.ID).Return(userUpdateEntityMock, nil)
	userDao.On("Save", &userUpdateEntityMock).Return(nil)

	controller := NewUserController(userDao, validatorService)

	router.HandleFunc(target, controller.UpdateUser)

	router.ServeHTTP(w, r)

	validatorService.AssertExpectations(t)
	userDao.AssertExpectations(t)

	assert.Equal(t, http.StatusOK, w.Code, "Debió regresar un status 200")

}

// TestUpdateUserValidationError valida que se regresen los errores cuando la data no cumple lo requerido
func TestUpdateUserValidationError(t *testing.T) {

	setUp()
	target := "/admin/user/update"

	w, r := httptest.NewRecorder(), httptest.NewRequest(http.MethodPut, target, serialize(t, &updateData))

	validatorService.On("Validate", userUpdateModel).Return([]models.ErrorValidationDetail{
		{Field: "Name", Constraints: []string{
			"required",
		}},
		{Field: "PaternalSureName", Constraints: []string{
			"required",
		}},
	}, false)
	// userDao.On("FindUserByID", userUpdateModel.ID).Return(userEntityMock, nil)
	// userDao.On("Save", &userEntityMock).Return(nil)

	controller := NewUserController(userDao, validatorService)

	router.HandleFunc(target, controller.UpdateUser)

	router.ServeHTTP(w, r)

	validatorService.AssertExpectations(t)
	userDao.AssertExpectations(t)

	assert.Equal(t, http.StatusBadRequest, w.Code, "Debió regresar un status 400")

}

// TestUpdateUserUserNotFound valida que se mande error cuando el usuario no existe
func TestUpdateUserUserNotFound(t *testing.T) {

	setUp()
	target := "/admin/user/update"

	w, r := httptest.NewRecorder(), httptest.NewRequest(http.MethodPut, target, serialize(t, &updateData))

	validatorService.On("Validate", userUpdateModel).Return([]models.ErrorValidationDetail{}, true)
	userDao.On("FindUserByID", userUpdateModel.ID).Return(entities.User{}, sql.ErrNoRows)
	// userDao.On("Save", &userEntityMock).Return(nil)

	controller := NewUserController(userDao, validatorService)

	router.HandleFunc(target, controller.UpdateUser)

	router.ServeHTTP(w, r)

	var res map[string]interface{}

	json.NewDecoder(w.Body).Decode(&res)

	validatorService.AssertExpectations(t)
	userDao.AssertExpectations(t)

	assert.Equal(t, "No se encontro el usuario que se quiere editar", res["message"].(string), "El mensaje no es el correcto")
	assert.Equal(t, http.StatusNotFound, w.Code, "Debió regresar un status 404")

}

// TestUpdateUserUserDaoError valida la respuesta cuando un error inesperado sucede al buscar un usuario
func TestUpdateUserUserDaoError(t *testing.T) {

	setUp()
	target := "/admin/user/update"

	w, r := httptest.NewRecorder(), httptest.NewRequest(http.MethodPut, target, serialize(t, &updateData))

	validatorService.On("Validate", userUpdateModel).Return([]models.ErrorValidationDetail{}, true)
	userDao.On("FindUserByID", userUpdateModel.ID).Return(entities.User{}, fmt.Errorf("Un error al buscar el usuario"))
	// userDao.On("Save", &userEntityMock).Return(nil)

	controller := NewUserController(userDao, validatorService)

	router.HandleFunc(target, controller.UpdateUser)

	router.ServeHTTP(w, r)

	var res map[string]interface{}

	json.NewDecoder(w.Body).Decode(&res)

	validatorService.AssertExpectations(t)
	userDao.AssertExpectations(t)

	assert.Equal(t, "Hubo un error inesperado al buscar al usuario", res["message"].(string), "El mensaje no es el correcto")
	assert.Equal(t, http.StatusInternalServerError, w.Code, "Debió regresar un status 500")

}

// TestUpdateUserUserDaoSaveError valida se mande el mensaje correcto cuando hay un error al salvar el usuario
func TestUpdateUserUserDaoSaveError(t *testing.T) {

	setUp()
	target := "/admin/user/update"

	w, r := httptest.NewRecorder(), httptest.NewRequest(http.MethodPut, target, serialize(t, &updateData))

	validatorService.On("Validate", userUpdateModel).Return([]models.ErrorValidationDetail{}, true)
	userDao.On("FindUserByID", userUpdateModel.ID).Return(userUpdateEntityMock, nil)
	userDao.On("Save", &userUpdateEntityMock).Return(fmt.Errorf("Un error en el metodo Save"))

	controller := NewUserController(userDao, validatorService)

	router.HandleFunc(target, controller.UpdateUser)

	router.ServeHTTP(w, r)

	var res map[string]interface{}

	json.NewDecoder(w.Body).Decode(&res)

	validatorService.AssertExpectations(t)
	userDao.AssertExpectations(t)

	assert.Equal(t, "Hubo un error inesperado al editar el usuario", res["message"].(string), "El mensaje no es el correcto")
	assert.Equal(t, http.StatusInternalServerError, w.Code, "Debió regresar un status 500")

}

/*
	Restore password
*/

func TestRestorePassword(t *testing.T) {
	setUp()
	target := "/admin/user/restore_password"

	w, r := httptest.NewRecorder(), httptest.NewRequest(http.MethodPut, target, serialize(t, &restorePasswordData))

	validatorService.On("Validate", restorePasswordModel).Return([]models.ErrorValidationDetail{}, true)
	userDao.On("FindUserByID", restorePasswordModel.ID).Return(userRestorePassEntityMock, nil)
	userDao.On("Save", &userRestorePassEntityMock).Return(nil)

	controller := NewUserController(userDao, validatorService)

	router.HandleFunc(target, controller.RestorePassword)

	router.ServeHTTP(w, r)

	validatorService.AssertExpectations(t)
	userDao.AssertExpectations(t)

	assert.Equal(t, http.StatusOK, w.Code, "Código http incorrecto")

}

func TestDeleteUser(t *testing.T) {
	setUp()
	endpoint := "/admin/user/delete/{idUser}"
	target := "/admin/user/delete/1"
	idUser := int32(1)

	w, r := httptest.NewRecorder(), httptest.NewRequest(http.MethodDelete, target, serialize(t, &restorePasswordData))

	userDao.On("Delete", idUser).Return(nil)

	controller := NewUserController(userDao, validatorService)

	router.HandleFunc(endpoint, controller.DeleteUser)

	router.ServeHTTP(w, r)

	validatorService.AssertExpectations(t)
	userDao.AssertExpectations(t)

	assert.Equal(t, http.StatusOK, w.Code, "Código http incorrecto")
}

func serialize(t *testing.T, i interface{}) *bytes.Buffer {

	jsonB, e := json.Marshal(i)
	if e != nil {
		t.Fatal(e)
	}
	return bytes.NewBuffer(jsonB)

}
