package dao

import (
	"database/sql"

	"github.com/manicar2093/guianetThea/app/entities"
)

const (
	findAllByStatus = `SELECT id_role, name, code, creation_date, edit_date, status FROM manager."CTHEA_ROLE" WHERE status=$1`
)

type RolDao interface {
	FindAllByStatus(status bool) ([]entities.Rol, error)
}

type RolDaoImpl struct {
	db *sql.DB
}

func NewRolDao(db *sql.DB) RolDao {
	return &RolDaoImpl{db}
}

func (r RolDaoImpl) FindAllByStatus(status bool) (found []entities.Rol, e error) {

	stmt, e := r.db.Prepare(findAllByStatus)
	if e != nil {
		return
	}

	res, e := stmt.Query(&status)
	if e != nil {
		return
	}

	for res.Next() {
		var temp entities.Rol
		if e = res.Scan(&temp.RolID, &temp.Name, &temp.Code, &temp.CreationDate, &temp.EditDate, &temp.Status); e != nil {
			return
		}
		found = append(found, temp)
	}

	return

}
