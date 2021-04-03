package models

type CreateUserData struct {
	RolID            int32  `json:"rol_id" validate:"required"`
	Name             string `json:"name" validate:"required"`
	PaternalSureName string `json:"paternal_surename" validate:"required"`
	MaternalSureName string `json:"maternal_surename"`
	Email            string `json:"email" validate:"required,email"`
	Password         string `json:"password" validate:"required"`
	PasswordConfirm  string `json:"password_confirm" validate:"required,eqfield=Password"`
}

func (c CreateUserData) GetValidableData() interface{} {
	return &c
}

type UpdateUserData struct {
	ID               int32  `json:"id" validate:"required"`
	RolID            int32  `json:"rol_id" validate:"required"`
	Name             string `json:"name" validate:"required"`
	PaternalSureName string `json:"paternal_surename" validate:"required"`
	MaternalSureName string `json:"maternal_surename"`
	Email            string `json:"email" validate:"required,email"`
}

func (u UpdateUserData) GetValidableData() interface{} {
	return &u
}

type RestorePasswordData struct {
	ID              int32  `json:"id" validate:"required"`
	Password        string `json:"password" validate:"required"`
	PasswordConfirm string `json:"password_confirm" validate:"required,eqfield=Password"`
}

func (r RestorePasswordData) GetValidableData() interface{} {
	return &r
}
