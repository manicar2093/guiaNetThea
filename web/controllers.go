package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type PageController struct {
	session *SessionHandler
}

func NewPageController(session *SessionHandler) *PageController {
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
	userDao *UserDao
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
