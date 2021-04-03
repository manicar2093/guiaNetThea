package dao

import (
	"testing"
	"time"

	"github.com/manicar2093/guianetThea/app/connections"
	"github.com/manicar2093/guianetThea/app/entities"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/guregu/null.v4/zero"
)

func TestDetailsHostingDaoImpl_Save(t *testing.T) {
	dao := NewDetailsHostingDao(connections.DB)
	closure := zero.NewTime(time.Now().Add(1*time.Hour), true)
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
	closure := zero.NewTime(time.Now().Add(1*time.Hour), true)
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
	closure := zero.NewTime(time.Now().Add(1*time.Hour), true)
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
