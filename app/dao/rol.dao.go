package dao

import (
	"database/sql"

	"github.com/manicar2093/guianetThea/app/entities"
)

const (
	findAllByStatus = `SELECT id_role, name, code, creation_date, edit_date, status FROM manager."CTHEA_ROLE" WHERE status=$1`
	findUserHasRol  = `SELECT
	tuser.id_user,
	trole.name
	FROM manager."THEA_USER" tuser
	INNER JOIN manager."CTHEA_ROLE" trole ON tuser.id_role = trole.id_role
	WHERE tuser.id_user=$1 AND trole.name=$2`
)

type RolDao interface {
	FindAllByStatus(status bool) ([]entities.Rol, error)
	UserHasRol(userID int, rolName string) (bool, error)
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

func (r RolDaoImpl) UserHasRol(userID int, rolName string) (bool, error) {
	stmt, e := r.db.Prepare(findUserHasRol)
	if e != nil {
		return false, e
	}

	var res struct {
		ID  int
		Rol string
	}

	e = stmt.QueryRow(&userID, &rolName).Scan(&res.ID, &res.Rol)
	if e != nil {
		if e == sql.ErrNoRows {
			return false, nil
		}
		return false, e
	}
	return true, nil
}
