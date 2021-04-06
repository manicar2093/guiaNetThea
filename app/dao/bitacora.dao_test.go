package dao

import (
	"testing"
	"time"

	"github.com/manicar2093/guianetThea/app/connections"
	"github.com/stretchr/testify/assert"
)

// TestLogRegistrySearch valida el happy path de la toma de datos
func TestLogRegistrySearch(t *testing.T) {
	dao := NewLogRegistryDao(connections.DB)

	initDate := time.Date(2021, time.March, 10, 00, 00, 00, 00, time.UTC)
	finalDate := time.Date(2021, time.March, 19, 00, 00, 00, 00, time.UTC)

	data, e := dao.LogRegistrySearch(initDate, finalDate)

	assert.Nil(t, e, "No debió regresar error")
	assert.Greater(t, len(data), 0, "No se tienen registros")
	t.Log(data[0])
}
