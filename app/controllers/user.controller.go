package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/manicar2093/guianetThea/app/dao"
	"github.com/manicar2093/guianetThea/app/models"
	"github.com/manicar2093/guianetThea/app/services"
	"github.com/manicar2093/guianetThea/app/utils"
)

// TODO: Realizar refactorización de código
type UserController interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	//GetUserByID(w http.ResponseWriter, r *http.Request)
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
		utils.Error.Printf("Error al decodificar los datos del body. Detalles: \n\t%v", e)
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

	_, e := u.userDao.SaveFromModel(data)
	if e != nil {
		utils.Error.Printf("Error al registrar nuevo usuario. Detalles: \n\t%v", e)
		utils.JSON(map[string]interface{}{
			"message": "Hubo un error al guardar el usuario. Favor de contactar a soporte",
		}, http.StatusInternalServerError, w)
		return
	}

	utils.JSON(nil, http.StatusCreated, w)
}

/*func (u UserControllerImpl) GetUserByID(w http.ResponseWriter, r *http.Request) {
	panic("not implemented") // TODO: Implement
}*/

func (u UserControllerImpl) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var data models.UpdateUserData
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

	stored, e := u.userDao.FindUserByID(int32(data.ID))
	if e != nil {
		switch {
		case e == sql.ErrNoRows:
			utils.JSON(map[string]interface{}{
				"message": "No se encontro el usuario que se quiere editar",
			}, http.StatusNotFound, w)
			return
		default:
			utils.Error.Printf("Error inesperado al buscar el usuario. Detalles: \n\t%v", e)
			utils.JSON(map[string]interface{}{
				"message": "Hubo un error inesperado al buscar al usuario",
			}, http.StatusInternalServerError, w)
			return
		}
	}

	stored.Name = data.Name
	stored.PaternalSureName = data.PaternalSureName
	stored.MaternalSureName.String = data.MaternalSureName
	stored.Email = data.Email
	stored.RolID.Int64 = int64(data.RolID)

	e = u.userDao.Save(&stored)
	if e != nil {
		utils.Error.Printf("Error inesperado al gurdar el usuario. Detalles: \n\t%v", e)
		utils.JSON(map[string]interface{}{
			"message": "Hubo un error inesperado al editar el usuario",
		}, http.StatusInternalServerError, w)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func (u UserControllerImpl) RestorePassword(w http.ResponseWriter, r *http.Request) {
	var data models.RestorePasswordData

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

	stored, e := u.userDao.FindUserByID(int32(data.ID))
	if e != nil {
		switch {
		case e == sql.ErrNoRows:
			utils.JSON(map[string]interface{}{
				"message": "No se encontro el usuario que se quiere editar",
			}, http.StatusNotFound, w)
			return
		default:
			utils.Error.Printf("Error inesperado al buscar el usuario. Detalles: \n\t%v", e)
			utils.JSON(map[string]interface{}{
				"message": "Hubo un error inesperado al buscar al usuario",
			}, http.StatusInternalServerError, w)
			return
		}
	}

	stored.Password = data.Password

	e = u.userDao.Save(&stored)
	if e != nil {
		utils.Error.Printf("Error inesperado al gurdar el usuario. Detalles: \n\t%v", e)
		utils.JSON(map[string]interface{}{
			"message": "Hubo un error inesperado al editar el usuario",
		}, http.StatusInternalServerError, w)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func (u UserControllerImpl) DeleteUser(w http.ResponseWriter, r *http.Request) {
	idUser, e := strconv.Atoi(mux.Vars(r)["idUser"])

	if e != nil {
		utils.Error.Printf("Hubo un error al obtener el idUser. Detalles: \n\t%v", e)
		utils.JSON(map[string]interface{}{
			"message": "Quizá lo que enviaste en el path no es un numero. Favor de revisar la documentación",
		}, http.StatusBadRequest, w)
		return
	}

	e = u.userDao.Delete(int32(idUser))
	if e != nil {
		utils.Error.Printf("Hubo un error al eliminar el usuario. Detalles: \n\t%v", e)
		utils.JSON(map[string]interface{}{
			"message": "Hubo un problema inesperado al borrar al usuario",
		}, http.StatusInternalServerError, w)
		return
	}
}
