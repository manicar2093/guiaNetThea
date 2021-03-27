package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/manicar2093/guianetThea/app/sessions"
)

func TestIsLoggedIn(t *testing.T) {

	w, r := httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/resource", nil)
	middl := NewMiddlewareProvider(sessions.Session)

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

	w, r := httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/resource", nil)
	middl := NewMiddlewareProvider(sessions.Session)

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
