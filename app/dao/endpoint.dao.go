package dao

import (
	"database/sql"

	"github.com/manicar2093/guianetThea/app/entities"
)

const (
	findEndpointByName = `SELECT id_endpoint,name FROM manager."CTHEA_ENDPOINT" WHERE name=$1 ORDER BY 1;`

	findEndpointByID = `SELECT id_endpoint,name FROM manager."CTHEA_ENDPOINT" WHERE id_endpoint=$1 ORDER BY 1;`
)

type EndpointDao interface {
	FindEndpointByName(name string) (entities.Endpoint, error)
	FindEndpointByID(id int32) (entities.Endpoint, error)
}

type EndpointDaoImpl struct {
	db *sql.DB
}

func NewEndpointDao(db *sql.DB) EndpointDao {
	return &EndpointDaoImpl{db}
}

func (e EndpointDaoImpl) FindEndpointByName(name string) (entities.Endpoint, error) {
	r := e.db.QueryRow(findEndpointByName, &name)
	if r.Err() != nil {
		return entities.Endpoint{}, r.Err()
	}
	var endpoint entities.Endpoint
	err := r.Scan(&endpoint.EndpointID, &endpoint.Name)
	if err != nil {
		return entities.Endpoint{}, err
	}
	return endpoint, nil
}

func (e EndpointDaoImpl) FindEndpointByID(id int32) (entities.Endpoint, error) {
	r := e.db.QueryRow(findEndpointByID, &id)
	if r.Err() != nil {
		return entities.Endpoint{}, r.Err()
	}
	var endpoint entities.Endpoint
	err := r.Scan(&endpoint.EndpointID, &endpoint.Name)
	if err != nil {
		return entities.Endpoint{}, err
	}
	return endpoint, nil
}
