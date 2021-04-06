package dao

import (
	"testing"
	"time"

	"github.com/manicar2093/guianetThea/app/connections"
	"github.com/manicar2093/guianetThea/app/entities"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/guregu/null.v4/zero"
)

func TestDetailsEndpointAndHostingDaoImpl_Save(t *testing.T) {
	detailsHostingDao := NewDetailsHostingDao(connections.DB)
	detailsEndpointAndHostingDao := NewDetailsEndpointAndHostingDao(connections.DB)

	closure := zero.NewTime(time.Now().Add(1*time.Hour), true)
	u1 := uuid.NewV4()
	details := entities.DetailsHosting{UserID: 1, Host: "HOST", SessionStart: time.Now(), SessionClosure: closure, UUID: u1.String()}

	detailsHostingDao.Save(&details)

	detailsAndEnpoint := entities.DetailsEndpointAndHosting{EndpointID: 1, DetailsHostingID: details.ID}

	e := detailsEndpointAndHostingDao.Save(&detailsAndEnpoint)

	if e != nil {
		t.Fatal("No debió regresar error: ", e)
	}

}
