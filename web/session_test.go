package web

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSessionHandler_GetUserId(t *testing.T) {
	w, r := httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/page", nil)

	userID, e := Session.GetUserID(w, r)
	if e != nil {
		t.Error("No debió haber error: ", e)
	}

	if userID > 0 {
		t.Error("No debe haber userID ya que no hay sesión")
	}

}

func TestSessionHandler_CreateNewSession(t *testing.T) {
	w, r, userID := httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/page", nil), 1

	e := Session.CreateNewSession(w, r, userID)
	if e != nil {
		t.Error("No debió haber error: ", e)
	}

	s, e := Session.session.Get(r, sessionName)
	if e != nil {
		t.Error("No debió haber error al obtener la sesión: ", e)
	}

	d, ok := s.Values["userId"]
	if !ok {
		t.Fatal("Debe encontrarse el userID en la sesión")
	}

	userIDData := d.(int)
	if userIDData != userID {
		t.Fatal("El userID no corresponde")
	}
}

func TestSessionHandler_GetCurrentSession(t *testing.T) {
	w, r, userID := httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/page", nil), 1

	e := Session.CreateNewSession(w, r, userID)
	if e != nil {
		t.Error("No debió haber error: ", e)
	}

	s, e := Session.GetCurrentSession(w, r)
	if e != nil {
		t.Error("No debió haber error al obtener la sesión: ", e)
	}

	d, ok := s.Values["userId"]
	if !ok {
		t.Fatal("Debe encontrarse el userID en la sesión")
	}

	userIDData := d.(int)
	if userIDData != userID {
		t.Fatal("El userID no corresponde")
	}
}
func TestSessionHandler_GetCurrentSessionFail(t *testing.T) {
	w, r := httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/page", nil)

	s, e := Session.GetCurrentSession(w, r)
	if e != nil {
		t.Error("No debió haber error al obtener la sesión: ", e)
	}

	d, ok := s.Values["userId"]
	if ok {
		t.Fatal("No debe encontrarse el userID en la sesión")
	}

	if d != nil {
		t.Fatal("El userID debe ser nil porque no hay sesión activa")
	}
}

func TestSessionHandler_IsLoggedIn(t *testing.T) {

}
