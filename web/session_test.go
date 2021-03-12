package web

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSessionHandler_GetUserId(t *testing.T) {
	w, r := httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/page", nil)

	sessionUUID, e := Session.GetUserID(w, r)
	if e != nil {
		t.Error("No debió haber error: ", e)
	}

	if sessionUUID > "" {
		t.Error("No debe haber UUID ya que no hay sesión")
	}

}

func TestSessionHandler_CreateNewSession(t *testing.T) {
	w, r, uuid := httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/page", nil), "UUID"

	e := Session.CreateNewSession(w, r, uuid)
	if e != nil {
		t.Error("No debió haber error: ", e)
	}

	store, _ := Session.(*SessionHandlerImpl)

	s, e := store.session.Get(r, sessionName)
	if e != nil {
		t.Error("No debió haber error al obtener la sesión: ", e)
	}

	d, ok := s.Values[sessionValue]
	if !ok {
		t.Fatal("Debe encontrarse el UUID en la sesión")
	}

	storedUuid := d.(string)
	if storedUuid != uuid {
		t.Fatal("El userID no corresponde")
	}
}

func TestSessionHandler_GetCurrentSession(t *testing.T) {
	w, r, uuid := httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/page", nil), "UUID"

	e := Session.CreateNewSession(w, r, uuid)
	if e != nil {
		t.Error("No debió haber error: ", e)
	}

	s, e := Session.GetCurrentSession(w, r)
	if e != nil {
		t.Error("No debió haber error al obtener la sesión: ", e)
	}

	d, ok := s.Values[sessionValue]
	if !ok {
		t.Fatal("Debe encontrarse el UUID en la sesión")
	}

	uuidStored := d.(string)
	if uuidStored != uuid {
		t.Fatal("El userID no corresponde")
	}
}
func TestSessionHandler_GetCurrentSessionFail(t *testing.T) {
	w, r := httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/page", nil)

	s, e := Session.GetCurrentSession(w, r)
	if e != nil {
		t.Error("No debió haber error al obtener la sesión: ", e)
	}

	d, ok := s.Values[sessionValue]
	if ok {
		t.Fatal("No debe encontrarse el UUID en la sesión")
	}

	if d != nil {
		t.Fatal("El UUID debe ser nil porque no hay sesión activa")
	}
}

func TestSessionHandler_IsLoggedIn(t *testing.T) {

}
