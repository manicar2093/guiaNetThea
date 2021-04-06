package dao

import (
	"database/sql"

	"github.com/manicar2093/guianetThea/app/entities"
)

const (
	insertDetailsEndpointAndHosting = `INSERT INTO manager."THEA_DETAILS_ENDPOINT_AND_HOSTIN" (id_details_endpoint_and_hostin, id_details_hosting, id_endpoint, creation_date, edit_date, status)
		VALUES (nextval('manager."THEA_DETAILS_id_details_endpoint_and_hostin_seq"'::regclass), $1, $2, now(), null, true) RETURNING id_details_endpoint_and_hostin`
)

type DetailsEndpointAndHostingDao interface {
	// Save guarda la instancia y coloca el ID. Solo se requiere el DetailsHostingID y EndpointID
	Save(details *entities.DetailsEndpointAndHosting) error
}

type DetailsEndpointAndHostingDaoImpl struct {
	db *sql.DB
}

func NewDetailsEndpointAndHostingDao(db *sql.DB) DetailsEndpointAndHostingDao {
	return &DetailsEndpointAndHostingDaoImpl{db}
}

func (d *DetailsEndpointAndHostingDaoImpl) Save(details *entities.DetailsEndpointAndHosting) error {
	stmt, e := d.db.Prepare(insertDetailsEndpointAndHosting)
	if e != nil {
		return e
	}

	e = stmt.QueryRow(&details.DetailsHostingID, &details.EndpointID).Scan(&details.ID)
	if e != nil {
		return e
	}
	return nil
}
