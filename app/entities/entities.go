package entities

import (
	"database/sql"
	"time"
)

// User es la representación de un registro en la base de datos
type User struct {
	UserID           int32
	RolID            sql.NullInt32
	Name             string
	PaternalSureName string
	MaternalSureName sql.NullString
	Email            string
	Password         string
	CreationDate     time.Time
	EditDate         sql.NullTime
	Status           bool
}

// DetailsHosting es la representación de un registro en la base de datos
type DetailsHosting struct {
	ID             int32
	UserID         int32
	Host           string
	SessionStart   time.Time
	SessionClosure sql.NullTime
	TypeLogOut     string
	CreationDate   time.Time
	EditDate       sql.NullTime
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
	EditDate     sql.NullTime
	Status       bool
}

// DetailsEndpointAndHosting es la representación de la relación entre Endpoint y DetailsHosting
type DetailsEndpointAndHosting struct {
	ID               int32
	DetailsHostingID int32
	EndpointID       int32
	CreationDate     time.Time
	EditDate         sql.NullTime
	Status           bool
}