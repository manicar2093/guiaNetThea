package dao

import (
	"database/sql"
	"time"

	"github.com/manicar2093/guianetThea/app/entities"
)

const loginRegistryByDate = `SELECT CONCAT (tuser.name,' ',tuser.paternal_surname,' ',tuser.maternal_surname) AS fullname,
	tuser.email AS email,
	trole.name AS rol,
	(session_closure-session_start) As time,
	ths.host AS host,
	ths.type_log_out AS type_log_out,
	ctep.name page 
FROM manager."THEA_DETAILS_HOSTING" ths
	INNER JOIN manager."THEA_USER" tuser ON tuser.id_user = ths.id_user
	INNER JOIN manager."CTHEA_ROLE" trole ON tuser.id_role = trole.id_role
	INNER JOIN manager."THEA_DETAILS_ENDPOINT_AND_HOSTIN" thss ON thss.id_details_hosting = ths.id_details_hosting
	INNER JOIN manager."CTHEA_ENDPOINT" ctep ON ctep.id_endpoint = thss.id_endpoint
	WHERE ths.creation_date BETWEEN $1 AND $2
order by 1`

type LoginRegistryDao interface {
	LogRegistrySearch(init, final time.Time) ([]entities.LoginRegistry, error)
}

type LoginRegistryDaoImpl struct {
	db *sql.DB
}

func NewLogRegistryDao(db *sql.DB) LoginRegistryDao {
	return &LoginRegistryDaoImpl{db}
}

func (l LoginRegistryDaoImpl) LogRegistrySearch(init, final time.Time) (res []entities.LoginRegistry, e error) {
	stmt, e := l.db.Prepare(loginRegistryByDate)
	if e != nil {
		return
	}
	r, e := stmt.Query(&init, &final)
	if e != nil {
		return
	}

	for r.Next() {
		var temp entities.LoginRegistry
		if e = r.Scan(&temp.Name, &temp.Email, &temp.Rol, &temp.Time, &temp.Host, &temp.TypeLogOut, &temp.Page); e != nil {
			return
		}
		res = append(res, temp)
	}
	return
}
