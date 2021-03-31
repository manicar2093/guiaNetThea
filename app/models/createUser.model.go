package models

type CreateUserData struct {
	RolID            int    `json:"rol_id" validate:"required"`
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
