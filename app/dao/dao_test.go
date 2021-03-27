package dao

import (
	"database/sql"
	"testing"
	"time"

	"github.com/manicar2093/guianetThea/app/connections"
	"github.com/manicar2093/guianetThea/app/entities"
	uuid "github.com/satori/go.uuid"
)

func TestSaveUser(t *testing.T) {
	dao := NewUserDao(connections.DB)
	user := entities.User{
		RolID:            sql.NullInt32{1, true},
		Name:             "Test",
		PaternalSureName: "Test",
		MaternalSureName: sql.NullString{Valid: true, String: "Test"},
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
		RolID:            sql.NullInt32{1, true},
		Name:             "Test DEL",
		PaternalSureName: "Test DEL",
		MaternalSureName: sql.NullString{Valid: true, String: "Test DEL"},
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
		RolID:            sql.NullInt32{1, true},
		Name:             "Test",
		PaternalSureName: "Test",
		MaternalSureName: sql.NullString{Valid: true, String: "Test"},
		Email:            "test@test.com",
		Password:         "password",
		Status:           true,
	}
	dao.Save(&user)

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

func TestFindEndpointByName(t *testing.T) {
	dao := NewEndpointDao(connections.DB)
	name := "inicio"
	endpoint, e := dao.FindEndpointByName(name)

	if e != nil {
		t.Fatal("No debió haber error: ", e)
	}

	if endpoint.EndpointID <= 0 {
		t.Fatal("No hay ID en la estructura")
	}
}

func TestFindEndpointByID(t *testing.T) {
	dao := NewEndpointDao(connections.DB)
	id := int32(1)
	endpoint, e := dao.FindEndpointByID(id)

	if e != nil {
		t.Fatal("No debió haber error: ", e)
	}

	if endpoint.Name == "" {
		t.Fatal("No hay Name en la estructura")
	}
}

func TestDetailsHostingDaoImpl_Save(t *testing.T) {
	dao := NewDetailsHostingDao(connections.DB)
	closure := sql.NullTime{Time: time.Now().Add(1 * time.Hour), Valid: true}
	u1 := uuid.NewV4()
	details := entities.DetailsHosting{UserID: 1, Host: "HOST", SessionStart: time.Now(), SessionClosure: closure, UUID: u1.String()}

	e := dao.Save(&details)
	if e != nil {
		t.Fatal("No debió regresar error:", e)
	}

	if details.ID == 0 {
		t.Fatal("No se tiene el ID del registro guardado")
	}
}

func TestDetailsHostingDaoImpl_Update(t *testing.T) {
	dao := NewDetailsHostingDao(connections.DB)
	closure := sql.NullTime{Time: time.Now().Add(1 * time.Hour), Valid: true}
	u1 := uuid.NewV4()
	details := entities.DetailsHosting{UserID: 1, Host: "HOST", SessionStart: time.Now(), SessionClosure: closure, UUID: u1.String()}

	dao.Save(&details)

	details.TypeLogOut = "MANUAL"
	details.SessionClosure.Time = details.SessionStart.Add(5 * time.Minute)

	e := dao.Save(&details)
	if e != nil {
		t.Fatal("No debió regresar error:", e)
	}

}

func TestDetailsHostingDaoImpl_FindDetailsHostingByUUID(t *testing.T) {
	dao := NewDetailsHostingDao(connections.DB)
	closure := sql.NullTime{Time: time.Now().Add(1 * time.Hour), Valid: true}
	u1 := uuid.NewV4()
	details := entities.DetailsHosting{UserID: 1, Host: "HOST", SessionStart: time.Now(), SessionClosure: closure, UUID: u1.String()}

	dao.Save(&details)

	saved, e := dao.FindDetailsHostingByUUID(details.UUID)
	if e != nil {
		t.Fatal("No debió regresar error:", e)
	}

	if saved.ID != details.ID {
		t.Fatal("No se recupero el registro de la base de datos")
	}
}

func TestDetailsEndpointAndHostingDaoImpl_Save(t *testing.T) {
	detailsHostingDao := NewDetailsHostingDao(connections.DB)
	detailsEndpointAndHostingDao := NewDetailsEndpointAndHostingDao(connections.DB)

	closure := sql.NullTime{Time: time.Now().Add(1 * time.Hour), Valid: true}
	u1 := uuid.NewV4()
	details := entities.DetailsHosting{UserID: 1, Host: "HOST", SessionStart: time.Now(), SessionClosure: closure, UUID: u1.String()}

	detailsHostingDao.Save(&details)

	detailsAndEnpoint := entities.DetailsEndpointAndHosting{EndpointID: 1, DetailsHostingID: details.ID}

	e := detailsEndpointAndHostingDao.Save(&detailsAndEnpoint)

	if e != nil {
		t.Fatal("No debió regresar error: ", e)
	}

}
