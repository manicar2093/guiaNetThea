package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type PageController struct {
}

func NewPageController() *PageController {
	return &PageController{}
}

func (p *PageController) GetLoginPage(w http.ResponseWriter, r *http.Request) {
	RenderTemplateToWriter("templates/login.html", w, nil)
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
}

func NewLoginController() *LoginController {
	return &LoginController{}
}

func (l *LoginController) Login(w http.ResponseWriter, r *http.Request) {
	username, password := r.FormValue("username"), r.FormValue("password")
	if username == "manicar2093" && password == "12345678" {
		Session.CreateNewSession(w, r, 4) // TODO recuerda quitar esto.
		http.Redirect(w, r, "/inicio", http.StatusSeeOther)
		return
	}
	fmt.Println(username, password)
	fmt.Fprintln(w, username, password)
}

func (l *LoginController) Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Login out")
}
