package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/manicar2093/guianetThea/app/services"
	"github.com/manicar2093/guianetThea/app/sessions"
	"github.com/manicar2093/guianetThea/app/utils"
)

type PageController struct {
	session       sessions.SessionHandler
	recordService services.RecordService
	templateUtils utils.TemplateUtils
}

func NewPageController(session sessions.SessionHandler, recordService services.RecordService, templateUtils utils.TemplateUtils) *PageController {
	return &PageController{session, recordService, templateUtils}
}

// GetLoginPage valida que si hay una sesión activa manda a /inicio. De lo contrario renderiza el template de login
func (p *PageController) GetLoginPage(w http.ResponseWriter, r *http.Request) {
	if !p.session.IsLoggedIn(w, r) {
		flash := p.session.GetFlashMessages(w, r)
		p.templateUtils.RenderTemplateToResponseWriter("templates/login.html", w, flash)
		return
	}

	http.Redirect(w, r, "/inicio", http.StatusSeeOther)
}

// GetOnDevTemplate regresa el template on_dev para motivos de despliegue
func (p *PageController) GetOnDevTemplate(w http.ResponseWriter, r *http.Request) {
	// FIXME: Puede que esto solo sea temporal
	p.templateUtils.RenderTemplateToResponseWriter("templates/on_dev.html", w, nil)

}

func (p *PageController) GetRequestedPage(w http.ResponseWriter, r *http.Request) {
	page := mux.Vars(r)["page"]
	if page == "favicon.ico" {
		utils.Info.Println("Recurso omitido", page)
		return
	}
	pagePath := fmt.Sprintf("templates/%s.html", page)
	e := p.templateUtils.RenderTemplateToResponseWriter(pagePath, w, nil)
	if e != nil {
		utils.Error.Printf("Error al ingresar a la pagina '%s'. Detalles: \n\t%v", page, e)
	}
	e = p.recordService.RegisterPageVisited(w, r, page)
	if e != nil {
		http.Error(w, "Error Interno del Servidor", http.StatusInternalServerError)
		return
	}
}

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
