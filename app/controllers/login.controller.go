package controllers

import (
	"fmt"
	"net/http"

	"github.com/manicar2093/guianetThea/app/services"
	"github.com/manicar2093/guianetThea/app/sessions"
	"github.com/manicar2093/guianetThea/app/utils"
)

type LoginController struct {
	loginService   services.LoginService
	sessionHandler sessions.SessionHandler
}

func NewLoginController(loginService services.LoginService, sessionHandler sessions.SessionHandler) *LoginController {
	return &LoginController{loginService, sessionHandler}
}

func (l *LoginController) Login(w http.ResponseWriter, r *http.Request) {
	username, password := r.FormValue("username"), r.FormValue("password")

	e := l.loginService.DoLogin(username, password, w, r)

	if e != nil {
		if el, ok := e.(services.LoginError); ok {
			utils.Error.Println(el.InternalMessage)
			l.sessionHandler.AddFlashMessage(sessions.FlashMessage{Type: "danger", Value: el.ClientMessage}, w, r)
			http.Redirect(w, r, "/index", http.StatusSeeOther)
			return
		}
		// FIXME arreglar el manejo de este error
		utils.Error.Println("Hubo un error al realizar el login:", e)
	}

	utils.Info.Println("Usuario logueado?", l.sessionHandler.IsLoggedIn(w, r))

	http.Redirect(w, r, "/inicio", http.StatusSeeOther)

}

func (l *LoginController) Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "LOGINOUT")
}
