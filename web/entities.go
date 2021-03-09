package web

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
	MaternalSureName string
	Email            string
	Password         string
	CreationDate     time.Time
	EditDate         sql.NullTime
	Status           bool
}

// DetailsHosting es la representación de un registro en la base de datos
type DetailsHosting struct {
	DetailsHostingID int32
	UserID           int32
	Host             string
	SessionStart     time.Time
	SessionClosure   sql.NullTime
	CreationDate     time.Time
	EditDate         sql.NullTime
	Status           bool
}
