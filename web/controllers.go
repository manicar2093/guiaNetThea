package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type PageController struct {
	session       SessionHandler
	recordService RecordService
}

func NewPageController(session SessionHandler, recordService RecordService) *PageController {
	return &PageController{session, recordService}
}

// GetLoginPage valida que si hay una sesión activa manda a /inicio. De lo contrario renderiza el template de login
func (p *PageController) GetLoginPage(w http.ResponseWriter, r *http.Request) {
	if !p.session.IsLoggedIn(w, r) {
		flash := p.session.GetFlashMessages(w, r)
		RenderTemplateToWriter("templates/login.html", w, flash)
		return
	}

	http.Redirect(w, r, "/inicio", http.StatusSeeOther)
}

// GetOnDevTemplate regresa el template on_dev para motivos de despliegue
func (p *PageController) GetOnDevTemplate(w http.ResponseWriter, r *http.Request) {

	RenderTemplateToWriter("templates/on_dev.html", w, nil)

}

func (p *PageController) GetRequestedPage(w http.ResponseWriter, r *http.Request) {
	page := mux.Vars(r)["page"]
	if page == "favicon.ico" {
		log.Println("Recurso omitido", page)
		return
	}
	pagePath := fmt.Sprintf("templates/%s.html", page)
	e := RenderTemplateToWriter(pagePath, w, nil)
	if e != nil {
		panic(e)
	}
	e = p.recordService.RegisterPageVisited(w, r, page)
	if e != nil {
		panic(e)
	}
}

type LoginController struct {
	loginService   LoginService
	sessionHandler SessionHandler
}

func NewLoginController(loginService LoginService, sessionHandler SessionHandler) *LoginController {
	return &LoginController{loginService, sessionHandler}
}

func (l *LoginController) Login(w http.ResponseWriter, r *http.Request) {
	username, password := r.FormValue("username"), r.FormValue("password")

	e := l.loginService.DoLogin(username, password, w, r)

	if e != nil {
		if el, ok := e.(LoginError); ok {
			Error.Println(el.internalMessage)
			l.sessionHandler.AddFlashMessage(FlashMessage{Type: "danger", Value: el.clientMessage}, w, r)
			http.Redirect(w, r, "/index", http.StatusSeeOther)
			return
		}
		Error.Println("Hubo un error al realizar el login:", e)
	}

	Info.Println("Usuario logueado?", l.sessionHandler.IsLoggedIn(w, r))

	http.Redirect(w, r, "/inicio", http.StatusSeeOther)

}

func (l *LoginController) Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "LOGINOUT")
}
