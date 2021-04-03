package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/manicar2093/guianetThea/app/entities"
	"github.com/manicar2093/guianetThea/app/mocks"
	"github.com/manicar2093/guianetThea/app/sessions"
)

var (
	rolDao            mocks.RolDaoMock
	detailsHostingDao mocks.DetailsHostingDaoMock
)

func setUp() {
	rolDao = mocks.RolDaoMock{}
	detailsHostingDao = mocks.DetailsHostingDaoMock{}
}

func TestIsLoggedIn(t *testing.T) {
	setUp()
	w, r := httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/resource", nil)
	middl := NewMiddlewareProvider(sessions.Session, rolDao, detailsHostingDao)

	sessions.Session.CreateNewSession(w, r, "UUID")

	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusAlreadyReported)
		return
	}

	server := mux.NewRouter()
	server.HandleFunc("/resource", middl.NeedsLoggedIn(f))
	server.ServeHTTP(w, r)

	if w.Code != http.StatusAlreadyReported {
		t.Fatal("No se reconocio la sesión")
	}

}

func TestIsLoggedInNoSession(t *testing.T) {
	setUp()
	w, r := httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/resource", nil)
	middl := NewMiddlewareProvider(sessions.Session, rolDao, detailsHostingDao)

	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusAlreadyReported)
		return
	}

	server := mux.NewRouter()
	server.HandleFunc("/resource", middl.NeedsLoggedIn(f))
	server.ServeHTTP(w, r)

	if w.Code != http.StatusSeeOther {
		t.Fatal("No se reconocio la sesión")
	}

	location := w.Header().Get("Location")

	if location != "/index" {
		t.Fatal("El URL destino no es el correcto:", location)
	}

}

func TestIsAdmin(t *testing.T) {
	setUp()
	w, r := httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/resource", nil)
	uuid := "UUID"
	sessions.Session.CreateNewSession(w, r, uuid)
	detailsHostingDao.On("FindDetailsHostingByUUID", uuid).Return(entities.DetailsHosting{UserID: 1}, nil)
	rolDao.On("UserHasRol", 1, "ADMIN").Return(true, nil)
	middl := NewMiddlewareProvider(sessions.Session, rolDao, detailsHostingDao)

	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusAlreadyReported)
		return
	}

	server := mux.NewRouter()
	server.HandleFunc("/resource", middl.IsAdmin(f))
	server.ServeHTTP(w, r)

	detailsHostingDao.AssertExpectations(t)
	rolDao.AssertExpectations(t)

	if w.Code != http.StatusAlreadyReported {
		t.Fatal("No se reconocio la sesión")
	}
}

func TestIsAdmin_NoAdmin(t *testing.T) {
	setUp()
	w, r := httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/resource", nil)
	uuid := "UUID"
	sessions.Session.CreateNewSession(w, r, uuid)
	detailsHostingDao.On("FindDetailsHostingByUUID", uuid).Return(entities.DetailsHosting{UserID: 1}, nil)
	rolDao.On("UserHasRol", 1, "ADMIN").Return(false, nil)
	middl := NewMiddlewareProvider(sessions.Session, rolDao, detailsHostingDao)

	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusAlreadyReported)
		return
	}

	server := mux.NewRouter()
	server.HandleFunc("/resource", middl.IsAdmin(f))
	server.ServeHTTP(w, r)

	detailsHostingDao.AssertExpectations(t)
	rolDao.AssertExpectations(t)

	if w.Code != http.StatusSeeOther {
		t.Fatal("No se reconocio la sesión")
	}

	location := w.Header().Get("Location")

	if location != "/index" {
		t.Fatal("El URL destino no es el correcto:", location)
	}
}
