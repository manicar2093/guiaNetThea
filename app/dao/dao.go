package dao

import (
	"database/sql"
	"fmt"

	"github.com/manicar2093/guianetThea/app/entities"
	"github.com/manicar2093/guianetThea/app/models"
)

// User queries
const (
	insertUser = `INSERT INTO manager."THEA_USER" (id_user, id_role, name, paternal_surname, maternal_surname, email, pasword, creation_date, edit_date, status) VALUES (nextval('manager."THEA_USER_id_user_seq"'::regclass) ,1 ,'%v' ,'%v' ,'%v' ,'%v' ,'%v' ,now(), null, true) RETURNING id_user`

	updateUser = `UPDATE manager."THEA_USER" SET name = $1, paternal_surname = $2, maternal_surname = $3, email = $4, edit_date = now() WHERE id_user = $5`

	deleteUser = `UPDATE manager."THEA_USER" SET status = false WHERE id_user = $1`

	findUserByID = `SELECT id_user, id_role, name, paternal_surname, maternal_surname, email, pasword, status FROM manager."THEA_USER" WHERE id_user= $1 group by id_user, id_role, name, paternal_surname, maternal_surname, email, pasword, status`

	findUserByEmail = `SELECT id_user, id_role, name, paternal_surname, maternal_surname, email, pasword, status FROM manager."THEA_USER" WHERE email= $1 group by id_user, id_role, name, paternal_surname, maternal_surname, email, pasword, status`
)

// Endpoint queries
const (
	findEndpointByName = `SELECT id_endpoint,name FROM manager."CTHEA_ENDPOINT" WHERE name=$1 ORDER BY 1;`

	findEndpointByID = `SELECT id_endpoint,name FROM manager."CTHEA_ENDPOINT" WHERE id_endpoint=$1 ORDER BY 1;`
)

// DetailsHosting queries
const (
	insertDetailsHosting = `INSERT INTO manager."THEA_DETAILS_HOSTING"(id_details_hosting, id_user, host, session_start, session_closure, type_log_out, creation_date, edit_date, status, uuid) VALUES (nextval('manager."THEA_DETAILS_HOSTING_id_details_hosting_seq"'::regclass), $1, $2, $3, $4, 'AUTOMATIC', now(), null, true, $5) RETURNING id_details_hosting`

	updateDetailsHosting = `UPDATE manager."THEA_DETAILS_HOSTING" SET session_closure=$1, type_log_out=$2, edit_date=now() WHERE id_details_hosting = $3`

	findDetailsHostingByUUID = `SELECT id_details_hosting, id_user, host, session_start, session_closure, type_log_out, creation_date, edit_date, status, uuid FROM manager."THEA_DETAILS_HOSTING" WHERE uuid = $1`
)

// DetailsEnpointAndHosting queries
const (
	insertDetailsEndpointAndHosting = `INSERT INTO manager."THEA_DETAILS_ENDPOINT_AND_HOSTIN" (id_details_endpoint_and_hostin, id_details_hosting, id_endpoint, creation_date, edit_date, status)
		VALUES (nextval('manager."THEA_DETAILS_id_details_endpoint_and_hostin_seq"'::regclass), $1, $2, now(), null, true) RETURNING id_details_endpoint_and_hostin`
)

type UserDao interface {
	Save(user *entities.User) error
	Delete(userID int32) error
	FindUserByID(id int32) (entities.User, error)
	FindUserByEmail(email string) (entities.User, error)
	SaveFromModel(user models.CreateUserData) error
}

type UserDaoImpl struct {
	db *sql.DB
}

func NewUserDao(db *sql.DB) *UserDaoImpl {
	return &UserDaoImpl{db}
}

// Save realiza el guardado de un usuario. Si el usuario no contiene un UserID se guardará un nuevo registro. Si UserID va lleno realizará el update del registro
func (u *UserDaoImpl) Save(user *entities.User) error {

	switch {
	case user.UserID <= 0:
		return u.insert(user)
	default:
		return u.update(user)
	}
}

// Delete realiza un update del registro con el idUser a false
func (u *UserDaoImpl) Delete(userID int32) error {
	r, e := u.db.Exec(deleteUser, userID)
	if e != nil {
		return e
	}
	affected, e := r.RowsAffected()
	if e != nil {
		return e
	}
	if affected <= 0 {
		return fmt.Errorf("No hubo ningun registro afectado")
	}

	return nil
}

// FindUserByID busca el usuario con el id proporcionado
func (u *UserDaoImpl) FindUserByID(id int32) (entities.User, error) {
	r := u.db.QueryRow(findUserByID, id)
	if r.Err() != nil {
		return entities.User{}, r.Err()
	}
	var user entities.User
	e := r.Scan(&user.UserID, &user.RolID, &user.Name, &user.PaternalSureName, &user.MaternalSureName, &user.Email, &user.Password, &user.Status)
	if e != nil {
		return entities.User{}, e
	}

	return user, nil
}

// FindUserByEmail busca un usuario por email
func (u *UserDaoImpl) FindUserByEmail(email string) (entities.User, error) {
	r := u.db.QueryRow(findUserByEmail, email)
	if r.Err() != nil {
		return entities.User{}, r.Err()
	}
	var user entities.User
	e := r.Scan(&user.UserID, &user.RolID, &user.Name, &user.PaternalSureName, &user.MaternalSureName, &user.Email, &user.Password, &user.Status)
	if e != nil {
		return entities.User{}, e
	}

	return user, nil
}

func (u *UserDaoImpl) SaveFromModel(user models.CreateUserData) error {
	panic("Not implemented yet")
}

func (u *UserDaoImpl) update(user *entities.User) error {
	r, e := u.db.Exec(updateUser, user.Name, user.PaternalSureName, user.MaternalSureName, user.Email, user.UserID)
	if e != nil {
		return e
	}
	affected, e := r.RowsAffected()
	if e != nil {
		return e
	}
	if affected <= 0 {
		return fmt.Errorf("No se actualizó el registro con id %d", user.UserID)
	}
	return nil
}

func (u *UserDaoImpl) insert(user *entities.User) error {
	query := fmt.Sprintf(insertUser, user.Name, user.PaternalSureName, user.MaternalSureName, user.Email, user.Password) // FIXME cambiar a stmt
	r := u.db.QueryRow(query)

	if r.Err() != nil {
		return r.Err()
	}

	r.Scan(&user.UserID)

	return nil
}

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
