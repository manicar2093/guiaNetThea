package services

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/manicar2093/guianetThea/app/entities"
	"github.com/stretchr/testify/mock"
)

func TestRecordServiceImpl_RegisterPageVisited(t *testing.T) {
	uuid, page, w, r := "uuid-session-id", "test", httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/test", nil)
	details := entities.DetailsHosting{ID: 1}
	endpoint := entities.Endpoint{EndpointID: 1}
	detailsEndpoint := entities.DetailsEndpointAndHosting{DetailsHostingID: details.ID, EndpointID: endpoint.EndpointID}

	setUp()

	sessionHandlerMock.On("GetSessionUUID", w, r).Return(uuid, nil)
	detailsHostingDaoMock.On("FindDetailsHostingByUUID", uuid).Return(details, nil)
	endpointDaoMock.On("FindEndpointByName", page).Return(endpoint, nil)
	detailsEndpointAndHostingDaoMock.On("Save", &detailsEndpoint).Return(nil)

	service := NewRecordService(detailsEndpointAndHostingDaoMock, detailsHostingDaoMock, endpointDaoMock, sessionHandlerMock)

	e := service.RegisterPageVisited(w, r, page)

	if e != nil {
		t.Fatal("No debió regresar error:", e)
	}

	sessionHandlerMock.AssertExpectations(t)
	detailsHostingDaoMock.AssertExpectations(t)
	endpointDaoMock.AssertExpectations(t)
	detailsEndpointAndHostingDaoMock.AssertExpectations(t)
}

func TestRecordServiceImpl_RegisterManualLogout(t *testing.T) {
	uuid, w, r := "uuid-session-id", httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/test", nil)
	details := entities.DetailsHosting{ID: 1}

	setUp()

	sessionHandlerMock.On("GetSessionUUID", w, r).Return(uuid, nil)
	detailsHostingDaoMock.On("FindDetailsHostingByUUID", uuid).Return(details, nil)
	detailsHostingDaoMock.On("Save", mock.AnythingOfType("*entities.DetailsHosting")).Return(nil)

	service := NewRecordService(detailsEndpointAndHostingDaoMock, detailsHostingDaoMock, endpointDaoMock, sessionHandlerMock)

	e := service.RegisterManualLogout(w, r)

	if e != nil {
		t.Fatal("No debió regresar error:", e)
	}

	sessionHandlerMock.AssertExpectations(t)
	detailsHostingDaoMock.AssertExpectations(t)
	endpointDaoMock.AssertExpectations(t)
}
