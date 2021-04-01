package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/manicar2093/guianetThea/app/services"
	"github.com/manicar2093/guianetThea/app/utils"
)

type CatalogsController interface {
	// GetCatalog es el controller que proveera de los catalogos que contega la bbase de datos.
	GetCatalog(w http.ResponseWriter, r *http.Request)
}

type CatalogsControllerImpl struct {
	catalogsService services.CatalogsService
}

func NewCatalogController(catalogsService services.CatalogsService) CatalogsController {
	return &CatalogsControllerImpl{catalogsService}
}

func (c CatalogsControllerImpl) GetCatalog(w http.ResponseWriter, r *http.Request) {
	catalog := mux.Vars(r)["catalog"]

	res, e := c.catalogsService.CreateCatalog(catalog)
	if e != nil {
		if e == services.ErrNoCatalog {
			utils.JSON(map[string]interface{}{
				"message": fmt.Sprintf("There is no catalog '%s'", catalog),
			}, http.StatusNotFound, w)
			return
		}
		utils.Error.Printf("Error al buscar catalogo '%s'. Detalles: \n\t%v", catalog, e)
		utils.JSON(map[string]interface{}{
			"message": "An internal error has occured. Please contact support",
		}, http.StatusInternalServerError, w)
		return
	}

	utils.JSON(map[string]interface{}{
		"message": "ok",
		"data":    res,
	}, http.StatusOK, w)

}
