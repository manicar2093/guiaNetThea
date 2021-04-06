package dao

import (
	"testing"

	"github.com/manicar2093/guianetThea/app/connections"
	"github.com/stretchr/testify/assert"
)

// TestFindAllByStatus valida el happy path
func TestFindAllByStatus(t *testing.T) {
	dao := NewRolDao(connections.DB)

	roles, e := dao.FindAllByStatus(true)

	assert.Nil(t, e, "No debió regresar error.")
	assert.Len(t, roles, 2, "No se regresó la cantidad de roles registrados")

}

func TestUserHasRol(t *testing.T) {
	dao := NewRolDao(connections.DB)

	hasRole, e := dao.UserHasRol(1, "GENERAL")
	assert.Nil(t, e, "No debió regresar error")

	assert.True(t, hasRole, "El usuario si cuenta con el rol")
}

func TestUserHasRol_NoRol(t *testing.T) {
	dao := NewRolDao(connections.DB)

	hasRole, e := dao.UserHasRol(1, "ADMIN")
	assert.Nil(t, e, "No debió regresar error")

	assert.False(t, hasRole, "El usuario si cuenta con el rol")
}
