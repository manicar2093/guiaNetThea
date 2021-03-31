package dao

import (
	"database/sql"
	"fmt"

	"github.com/manicar2093/guianetThea/app/entities"
	"github.com/manicar2093/guianetThea/app/models"
)

const (
	insertUser = `INSERT INTO manager."THEA_USER" (id_user, id_role, name, paternal_surname, maternal_surname, email, pasword, creation_date, edit_date, status) VALUES (nextval('manager."THEA_USER_id_user_seq"'::regclass) ,1 ,'%v' ,'%v' ,'%v' ,'%v' ,'%v' ,now(), null, true) RETURNING id_user`

	updateUser = `UPDATE manager."THEA_USER" SET name = $1, paternal_surname = $2, maternal_surname = $3, email = $4, edit_date = now(), pasword = $5 WHERE id_user = $6`

	deleteUser = `UPDATE manager."THEA_USER" SET status = false WHERE id_user = $1`

	findUserByID = `SELECT id_user, id_role, name, paternal_surname, maternal_surname, email, pasword, status FROM manager."THEA_USER" WHERE id_user= $1 group by id_user, id_role, name, paternal_surname, maternal_surname, email, pasword, status`

	findUserByEmail = `SELECT id_user, id_role, name, paternal_surname, maternal_surname, email, pasword, status FROM manager."THEA_USER" WHERE email= $1 group by id_user, id_role, name, paternal_surname, maternal_surname, email, pasword, status`
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
	r, e := u.db.Exec(updateUser, user.Name, user.PaternalSureName, user.MaternalSureName, user.Email, user.Password, user.UserID)
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
