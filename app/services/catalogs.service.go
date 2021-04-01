package services

import (
	"errors"

	"github.com/manicar2093/guianetThea/app/dao"
	"github.com/manicar2093/guianetThea/app/models"
)

var (
	ErrNoCatalog = errors.New("no catalog registred for this request")
)

type CatalogsService interface {
	// CreateCatalog crea el slice necesario con los datos del catalogo solicitado.
	CreateCatalog(catalog string) ([]models.CatalogModel, error)
}

type CatalogsServiceImpl struct {
	rolDao dao.RolDao
}

func NewCatalogService(rolDao dao.RolDao) CatalogsService {
	return &CatalogsServiceImpl{rolDao}
}

// FIXME: Este puede ser creado en conjunto con un ORM. Se puede recibir un struct que implemente Catalogable.
// CreateCatalog crea el slice necesario con los datos del catalogo solicitado.
func (c CatalogsServiceImpl) CreateCatalog(catalog string) ([]models.CatalogModel, error) {
	var res []models.CatalogModel
	switch catalog {
	case "rol":
		roles, e := c.rolDao.FindAllByStatus(true)
		if e != nil {
			return res, e
		}
		for _, v := range roles {
			res = append(res, v.GetCatalogInstance())
		}
	default:
		return res, ErrNoCatalog
	}
	return res, nil
}
