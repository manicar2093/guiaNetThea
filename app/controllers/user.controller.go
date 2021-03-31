package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/manicar2093/guianetThea/app/dao"
	"github.com/manicar2093/guianetThea/app/models"
	"github.com/manicar2093/guianetThea/app/services"
	"github.com/manicar2093/guianetThea/app/utils"
)

type UserController interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	GetUserByID(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	RestorePassword(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

type UserControllerImpl struct {
	userDao          dao.UserDao
	validatorService services.ValidatorService
}

func NewUserController(userDao dao.UserDao, validatorService services.ValidatorService) UserController {
	return &UserControllerImpl{userDao, validatorService}
}

func (u UserControllerImpl) CreateUser(w http.ResponseWriter, r *http.Request) {
	var data models.CreateUserData
	if e := json.NewDecoder(r.Body).Decode(&data); e != nil {
		utils.JSON(map[string]interface{}{"message": "No se logró obtener la información necesaria. Valida la documentación"}, http.StatusInternalServerError, w)
		return
	}

	err, valid := u.validatorService.Validate(data)
	if !valid {
		utils.JSON(map[string]interface{}{
			"message": "La información es incorrecta. Favor de revisar la documentación",
			"errores": err,
		}, http.StatusBadRequest, w)
		return
	}

	e := u.userDao.SaveFromModel(data)
	if e != nil {
		utils.Error.Printf("Error al registrar nuevo usuario. Detalles: \n\t%v", e)
		utils.JSON(map[string]interface{}{
			"message": "Hubo un error al guardar el usuario. Favor de contactar a soporte",
		}, http.StatusInternalServerError, w)
		return
	}

	utils.JSON(nil, http.StatusCreated, w)
}

func (u UserControllerImpl) GetUserByID(w http.ResponseWriter, r *http.Request) {
	panic("not implemented") // TODO: Implement
}

func (u UserControllerImpl) UpdateUser(w http.ResponseWriter, r *http.Request) {
	panic("not implemented") // TODO: Implement
}

func (u UserControllerImpl) RestorePassword(w http.ResponseWriter, r *http.Request) {
	panic("not implemented") // TODO: Implement
}

func (u UserControllerImpl) DeleteUser(w http.ResponseWriter, r *http.Request) {
	panic("not implemented") // TODO: Implement
}
