package dao

import (
	"database/sql"
	"fmt"

	"github.com/manicar2093/guianetThea/app/entities"
)

const (
	insertDetailsHosting = `INSERT INTO manager."THEA_DETAILS_HOSTING"(id_details_hosting, id_user, host, session_start, session_closure, type_log_out, creation_date, edit_date, status, uuid) VALUES (nextval('manager."THEA_DETAILS_HOSTING_id_details_hosting_seq"'::regclass), $1, $2, $3, $4, 'AUTOMATIC', now(), null, true, $5) RETURNING id_details_hosting`

	updateDetailsHosting = `UPDATE manager."THEA_DETAILS_HOSTING" SET session_closure=$1, type_log_out=$2, edit_date=now() WHERE id_details_hosting = $3`

	findDetailsHostingByUUID = `SELECT id_details_hosting, id_user, host, session_start, session_closure, type_log_out, creation_date, edit_date, status, uuid FROM manager."THEA_DETAILS_HOSTING" WHERE uuid = $1`
)

type DetailsHostingDao interface {
	// Save realiza el guardado de un DetailsHostingDaoImpl. Si el DetailsHostingDaoImpl no contiene un ID se guardará un nuevo registro. Si ID va lleno realizará el update del registro.
	//
	// Se debe considerar que el salvado de información solo contempla los campos id_user, host, session_start, session_closure y uuid. El update solo modifica los campos session_closure y  type_log_out
	Save(details *entities.DetailsHosting) error
	FindDetailsHostingByUUID(uuid string) (entities.DetailsHosting, error)
}

type DetailsHostingDaoImpl struct {
	db *sql.DB
}

func NewDetailsHostingDao(db *sql.DB) DetailsHostingDao {
	return &DetailsHostingDaoImpl{db}
}

// Save realiza el guardado de un DetailsHostingDaoImpl. Si el DetailsHostingDaoImpl no contiene un ID se guardará un nuevo registro. Si ID va lleno realizará el update del registro.
//
// Se debe considerar que el salvado de información solo contempla los campos id_user, host, session_start, session_closure y uuid. El update solo modifica los campos session_closure y  type_log_out
func (d *DetailsHostingDaoImpl) Save(details *entities.DetailsHosting) error {
	switch {
	case details.ID <= 0:
		return d.insert(details)
	default:
		return d.update(details)
	}
}

func (d *DetailsHostingDaoImpl) FindDetailsHostingByUUID(uuid string) (entities.DetailsHosting, error) {
	r := d.db.QueryRow(findDetailsHostingByUUID, &uuid)
	if r.Err() != nil {
		return entities.DetailsHosting{}, r.Err()
	}
	var details entities.DetailsHosting
	e := r.Scan(&details.ID, &details.UserID, &details.Host, &details.SessionStart, &details.SessionClosure, &details.TypeLogOut, &details.CreationDate, &details.EditDate, &details.Status, &details.UUID)
	if e != nil {
		return entities.DetailsHosting{}, e
	}

	return details, nil
}

func (d *DetailsHostingDaoImpl) update(details *entities.DetailsHosting) error {
	r, e := d.db.Exec(updateDetailsHosting, details.SessionClosure, details.TypeLogOut, details.ID)
	if e != nil {
		return e
	}
	affected, e := r.RowsAffected()
	if e != nil {
		return e
	}
	if affected <= 0 {
		return fmt.Errorf("No se actualizó el registro con id %d", details.ID)
	}
	return nil
}

func (d *DetailsHostingDaoImpl) insert(details *entities.DetailsHosting) error {
	stmt, e := d.db.Prepare(insertDetailsHosting)
	if e != nil {
		return e
	}
	r := stmt.QueryRow(&details.UserID, &details.Host, &details.SessionStart, &details.SessionClosure.Time, &details.UUID)
	if r.Err() != nil {
		return r.Err()
	}

	r.Scan(&details.ID)

	return nil
}
