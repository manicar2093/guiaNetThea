package services

import (
	"net/http"

	"github.com/manicar2093/guianetThea/app/dao"
	"github.com/manicar2093/guianetThea/app/entities"
	"github.com/manicar2093/guianetThea/app/sessions"
)

// RecordService se encarga de realizar el guardado de las paginas que se visitan con una sesión
type RecordService interface {
	RegisterPageVisited(w http.ResponseWriter, req *http.Request, page string) error
}

type RecordServiceImpl struct {
	detailsEndpointAndHostingDao dao.DetailsEndpointAndHostingDao
	detailsHostingDao            dao.DetailsHostingDao
	endpointDao                  dao.EndpointDao
	sessionHandler               sessions.SessionHandler
}

func NewRecordService(detailsEndpointAndHostingDao dao.DetailsEndpointAndHostingDao, detailsHostingDao dao.DetailsHostingDao, endpointDao dao.EndpointDao, sessionHandler sessions.SessionHandler) RecordService {
	return &RecordServiceImpl{detailsEndpointAndHostingDao, detailsHostingDao, endpointDao, sessionHandler}
}

func (r RecordServiceImpl) RegisterPageVisited(w http.ResponseWriter, req *http.Request, page string) error {
	uuid, e := r.sessionHandler.GetUserID(w, req)
	if e != nil {
		return e
	}

	details, e := r.detailsHostingDao.FindDetailsHostingByUUID(uuid)
	if e != nil {
		return e
	}

	endpoint, e := r.endpointDao.FindEndpointByName(page)
	if e != nil {
		return e
	}

	e = r.detailsEndpointAndHostingDao.Save(&entities.DetailsEndpointAndHosting{DetailsHostingID: details.ID, EndpointID: endpoint.EndpointID})
	if e != nil {
		return e
	}

	return nil
}
