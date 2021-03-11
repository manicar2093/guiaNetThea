package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type PageController struct {
	session SessionHandler
}

func NewPageController(session SessionHandler) *PageController {
	return &PageController{session: session}
}

// GetLoginPage valida que si hay una sesión activa manda a /inicio. De lo contrario renderiza el template de login
func (p *PageController) GetLoginPage(w http.ResponseWriter, r *http.Request) {
	if !p.session.IsLoggedIn(w, r) {
		RenderTemplateToWriter("templates/login.html", w, nil)
		return
	}

	http.Redirect(w, r, "/inicio", http.StatusSeeOther)
}

func (p *PageController) GetRequestedPage(w http.ResponseWriter, r *http.Request) {
	page := mux.Vars(r)["page"]
	if page == "favicon.ico" {
		log.Println("Recurso omitido", page)
		return
	}
	pagePath := fmt.Sprintf("templates/%s.html", page)
	RenderTemplateToWriter(pagePath, w, nil)
}

type LoginController struct {
	loginService LoginService
}

func NewLoginController(loginService LoginService) *LoginController {
	return &LoginController{loginService}
}

func (l *LoginController) Login(w http.ResponseWriter, r *http.Request) {
	username, password := r.FormValue("username"), r.FormValue("password")

	e := l.loginService.DoLogin(username, password, w, r)

	if e != nil {
		if el, ok := e.(LoginError); ok {
			fmt.Println(el) // TODO crear los mensajes flash
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}

	http.Redirect(w, r, "/inicio", http.StatusSeeOther)
	return

}

func (l *LoginController) Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Login out")
}
