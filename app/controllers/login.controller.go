package controllers

import (
	"net/http"

	"github.com/manicar2093/guianetThea/app/services"
	"github.com/manicar2093/guianetThea/app/sessions"
	"github.com/manicar2093/guianetThea/app/utils"
)

type LoginController struct {
	loginService   services.LoginService
	sessionHandler sessions.SessionHandler
	recordService  services.RecordService
}

func NewLoginController(loginService services.LoginService, sessionHandler sessions.SessionHandler, recordService services.RecordService) *LoginController {
	return &LoginController{loginService, sessionHandler, recordService}
}

func (l *LoginController) Login(w http.ResponseWriter, r *http.Request) {
	username, password := r.FormValue("username"), r.FormValue("password")

	e := l.loginService.DoLogin(username, password, w, r)

	if e != nil {
		if el, ok := e.(services.LoginError); ok {
			utils.Error.Println("Usuario:", username, "Detalles:", el.InternalMessage)
			l.sessionHandler.AddFlashMessage(sessions.FlashMessage{Type: "danger", Value: el.ClientMessage}, w, r)
			http.Redirect(w, r, "/index", http.StatusSeeOther)
			return
		}
		// FIXME arreglar el manejo de este error
		utils.Error.Println("Hubo un error al realizar el login:", e)
	}

	utils.Info.Println("Usuario", username, "logueado exitosamente")

	http.Redirect(w, r, "/inicio", http.StatusSeeOther)

}

func (l *LoginController) Logout(w http.ResponseWriter, r *http.Request) {
	e := l.recordService.RegisterManualLogout(w, r)
	if e != nil {
		utils.Error.Printf("Error al registrar el cierre manual de la sesión. Detalles: \n\t%v", e)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	e = l.sessionHandler.DeleteSession(w, r)
	if e != nil {
		utils.Error.Printf("Error al eliminar la sesión. Detalles: \n\t%v", e)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	l.sessionHandler.AddFlashMessage(sessions.FlashMessage{Type: "info", Value: "Sesión terminada"}, w, r)
	http.Redirect(w, r, "/index", http.StatusSeeOther)
	return
}
