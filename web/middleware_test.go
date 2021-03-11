package web

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestIsLoggedIn(t *testing.T) {

	w, r := httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/resource", nil)
	middl := NewMiddlewareProvider(Session)

	Session.CreateNewSession(w, r, 2)

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

	w, r := httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/resource", nil)
	middl := NewMiddlewareProvider(Session)

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

	if location != "/" {
		t.Fatal("El URL destino no es el correcto:", location)
	}

}
