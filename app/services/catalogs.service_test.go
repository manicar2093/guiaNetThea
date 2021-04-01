package services

import (
	"testing"

	"github.com/manicar2093/guianetThea/app/entities"
	"github.com/stretchr/testify/assert"
)

// TestCreateCatalog valida el happy path para obtener el catalogo
func TestCreateCatalog(t *testing.T) {
	setUp()
	catalog := "rol"

	rols := []entities.Rol{
		{RolID: 1, Name: "rol1", Code: "RL1", Status: true},
		{RolID: 2, Name: "rol2", Code: "RL2", Status: true},
		{RolID: 3, Name: "rol3", Code: "RL3", Status: true},
	}

	rolDao.On("FindAllRolByStatus", true).Return(rols, nil)

	service := NewCatalogService(rolDao)

	c, e := service.CreateCatalog(catalog)

	rolDao.AssertExpectations(t)

	assert.Nil(t, e, "No debió regresar error")
	assert.Len(t, c, 3, "No se regresaron la cantidad de elementos necesarios")

}

func TestCreateCatalogNotCatalog(t *testing.T) {
	setUp()
	catalog := "dont exists"

	service := NewCatalogService(rolDao)

	c, e := service.CreateCatalog(catalog)

	rolDao.AssertExpectations(t)

	assert.NotNil(t, e, "Debió regresar error")
	assert.Len(t, c, 0, "No se regresaron la cantidad de elementos necesarios")

}
