package controllers

import (
	"fmt"
	"html/template"
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
	switch page {
	case "favicon.ico":
		return
	case "admin":
		http.Redirect(w, r, "/admin/", http.StatusSeeOther)
		return
	}

	pagePath := fmt.Sprintf("templates/%s.html", page)

	logout := template.HTML(`<a href="/logout/" style="background-color: red; padding: .5rem; width: auto; position: fixed; z-index: 90;">LOGOUT</a>`)

	e := p.templateUtils.RenderTemplateToResponseWriter(pagePath, w, map[string]interface{}{
		"logout": logout,
	})
	if e != nil {
		utils.Error.Printf("Error al ingresar a la pagina '%s'. Detalles: \n\t%v", page, e)
	}
	e = p.recordService.RegisterPageVisited(w, r, page)
	if e != nil {
		http.Error(w, "Error Interno del Servidor", http.StatusInternalServerError)
		return
	}
}
