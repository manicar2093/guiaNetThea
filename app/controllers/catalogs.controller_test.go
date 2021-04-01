package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/manicar2093/guianetThea/app/models"
	"github.com/manicar2093/guianetThea/app/services"
	"github.com/stretchr/testify/assert"
)

// TestGetCatalog valida el happy path al obtener el catalogo
func TestGetCatalog(t *testing.T) {
	setUp()

	catalog := "rol"
	rolMock := []models.CatalogModel{
		{ID: 1, Description: "rol1"},
		{ID: 2, Description: "rol2"},
		{ID: 3, Description: "rol3"},
		{ID: 4, Description: "rol4"},
	}

	w, r := httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/admin/catalogs/rol", nil)

	catalogsService.On("CreateCatalog", catalog).Return(rolMock, nil)

	controller := NewCatalogController(catalogsService)

	router.HandleFunc("/admin/catalogs/{catalog}", controller.GetCatalog).Methods(http.MethodGet)
	router.ServeHTTP(w, r)

	catalogsService.AssertExpectations(t)

	var res map[string]interface{}

	json.NewDecoder(w.Body).Decode(&res)

	assert.Len(t, res["data"].([]interface{}), 4, "No se tienen los registros necesarios")

	assert.Equal(t, http.StatusOK, w.Code, "No se obtubo el HTTP status correcto")

}

// TestGetCatalogNoCatalog valida la respuesta cuando no existe el catalogo
func TestGetCatalogNoCatalog(t *testing.T) {
	setUp()

	catalog := "rol"
	rolMock := []models.CatalogModel{}

	w, r := httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/admin/catalogs/rol", nil)

	catalogsService.On("CreateCatalog", catalog).Return(rolMock, services.ErrNoCatalog)

	controller := NewCatalogController(catalogsService)

	router.HandleFunc("/admin/catalogs/{catalog}", controller.GetCatalog).Methods(http.MethodGet)
	router.ServeHTTP(w, r)

	catalogsService.AssertExpectations(t)

	var res map[string]interface{}

	json.NewDecoder(w.Body).Decode(&res)

	assert.Equal(t, res["message"].(string), "There is no catalog 'rol'", "No se tienen los registros necesarios")

	assert.Equal(t, http.StatusNotFound, w.Code, "No se obtubo el HTTP status correcto")

}

// TestGetCatalogCatalogServiceError valida el error que se regresa cuando el servicio tiene un error inesperado
func TestGetCatalogCatalogServiceError(t *testing.T) {
	setUp()

	catalog := "rol"
	rolMock := []models.CatalogModel{}

	w, r := httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/admin/catalogs/rol", nil)

	catalogsService.On("CreateCatalog", catalog).Return(rolMock, fmt.Errorf("Un error inesperado"))

	controller := NewCatalogController(catalogsService)

	router.HandleFunc("/admin/catalogs/{catalog}", controller.GetCatalog).Methods(http.MethodGet)
	router.ServeHTTP(w, r)

	catalogsService.AssertExpectations(t)

	var res map[string]interface{}

	json.NewDecoder(w.Body).Decode(&res)

	assert.Equal(t, res["message"].(string), "An internal error has occured. Please contact support", "No se tienen los registros necesarios")

	assert.Equal(t, http.StatusInternalServerError, w.Code, "No se obtubo el HTTP status correcto")

}
