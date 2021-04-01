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
