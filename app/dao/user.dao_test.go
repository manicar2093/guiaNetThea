package dao

import (
	"database/sql"
	"testing"

	"github.com/manicar2093/guianetThea/app/connections"
	"github.com/manicar2093/guianetThea/app/entities"
	"github.com/manicar2093/guianetThea/app/models"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v4/zero"
)

func TestSaveUser(t *testing.T) {
	dao := NewUserDao(connections.DB)
	user := entities.User{
		RolID:            zero.NewInt(1, true),
		Name:             "Test",
		PaternalSureName: "Test",
		MaternalSureName: zero.NewString("Test", true),
		Email:            "test@test.com",
		Password:         "password",
		Status:           true,
	}

	e := dao.Save(&user)
	if e != nil {
		t.Fatal("No debió regresar error: ", e)
	}

	if user.UserID <= 0 {
		t.Fatal("El id ingresado debe ser mayor a 0")
	}
}

func TestDeleteUser(t *testing.T) {
	dao := NewUserDao(connections.DB)
	user := entities.User{
		RolID:            zero.NewInt(1, true),
		Name:             "Test DEL",
		PaternalSureName: "Test DEL",
		MaternalSureName: zero.NewString("Test DEL", true),
		Email:            "test@test.com",
		Password:         "password",
		Status:           true,
	}
	dao.Save(&user)

	e := dao.Delete(user.UserID)

	if e != nil {
		t.Fatal("No debió regresar error:", e)
	}
}

func TestUpdateUser(t *testing.T) {
	dao := NewUserDao(connections.DB)
	newEmail, newName := "Changed", "Changed"
	user := entities.User{
		RolID:            zero.NewInt(1, true),
		Name:             "Test",
		PaternalSureName: "Test",
		MaternalSureName: zero.NewString("Test", true),
		Email:            "test@test.com",
		Password:         "password",
		Status:           true,
	}
	dao.Save(&user)

	user.RolID = zero.IntFrom(2)
	user.Email = newEmail
	user.Name = newName

	e := dao.Save(&user)
	if e != nil {
		t.Fatal("No debió regresar error: ", e)
	}

	if user.Name != newName || user.Email != newEmail {
		t.Fatal("Los datos no fueron cambiados")
	}
}

func TestFindUserByID(t *testing.T) {
	dao := NewUserDao(connections.DB)
	id := int32(2)
	user, e := dao.FindUserByID(id)
	if e != nil {
		t.Fatal("No debió haber error. El registro existe: ", e)
	}
	if user.UserID != id {
		t.Fatal("El id no corresponde. El registro no es correcto")
	}
}

func TestFindUserByIDNotExists(t *testing.T) {
	dao := NewUserDao(connections.DB)
	id := int32(99999)
	_, e := dao.FindUserByID(id)
	if e != sql.ErrNoRows {
		t.Fatal("Error inesperado. Debió ser sql.ErrNoRows", e)
	}

}

func TestFindUserByEmail(t *testing.T) {
	dao := NewUserDao(connections.DB)
	email := "email1@email.com"
	user, e := dao.FindUserByEmail(email)
	if e != nil {
		t.Fatal("No debió haber error. El registro existe: ", e)
	}
	if user.Email != email {
		t.Fatal("El email no corresponde. El registro no es correcto")
	}
}

func TestSaveFromModel(t *testing.T) {
	dao := NewUserDao(connections.DB)
	model := models.CreateUserData{
		RolID:            1,
		Name:             "ModelTest",
		PaternalSureName: "ModelTest",
		MaternalSureName: "ModelTest",
		Email:            "email@test.com",
		Password:         "pass",
		PasswordConfirm:  "pass",
	}
	id, err := dao.SaveFromModel(model)

	assert.Nil(t, err, "No debió regresar error")
	assert.Greater(t, id, 0, "No se recibio el id")
}

func TestFindAll(t *testing.T) {
	dao := NewUserDao(connections.DB)
	users, e := dao.FindAll()

	assert.Nil(t, e, "No debió regresar error")
	assert.Greater(t, len(users), 0, "No hay usuarios en el slice")
}
