package entities

import (
	"time"

	"github.com/manicar2093/guianetThea/app/models"
	"gopkg.in/guregu/null.v4/zero"
)

// User es la representación de un registro en la base de datos
type User struct {
	UserID           int32       `json:"id"`
	RolID            zero.Int    `json:"rol_id"`
	Name             string      `json:"name"`
	PaternalSureName string      `json:"paternal_surename"`
	MaternalSureName zero.String `json:"maternal_surename"`
	Email            string      `json:"email"`
	Password         string      `json:"-"`
	CreationDate     time.Time   `json:"creation_date"`
	EditDate         zero.Time   `json:"update_date"`
	Status           bool        `json:"status"`
}

// DetailsHosting es la representación de un registro en la base de datos
type DetailsHosting struct {
	ID             int32
	UserID         int32
	Host           string
	SessionStart   time.Time
	SessionClosure zero.Time
	TypeLogOut     string
	CreationDate   time.Time
	EditDate       zero.Time
	Status         bool
	UUID           string
}

// Endpoint es la representación del catalogo
type Endpoint struct {
	EndpointID int32
	Name       string
	//Description  string
	//CreationDate time.Time
	//EditDate     sql.NullTime
	//Status       bool
}

// Rol es la representación de la tabla
type Rol struct {
	RolID        int32
	Name         string
	Code         string
	CreationDate time.Time
	EditDate     zero.Time
	Status       bool
}

func (r Rol) GetCatalogInstance() models.CatalogModel {
	return models.CatalogModel{ID: int(r.RolID), Description: r.Name}
}

// DetailsEndpointAndHosting es la representación de la relación entre Endpoint y DetailsHosting
type DetailsEndpointAndHosting struct {
	ID               int32
	DetailsHostingID int32
	EndpointID       int32
	CreationDate     time.Time
	EditDate         zero.Time
	Status           bool
}
