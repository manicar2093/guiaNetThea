package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"github.com/manicar2093/guianetThea/app/dao"
	"github.com/manicar2093/guianetThea/app/services"
	"github.com/manicar2093/guianetThea/app/utils"
)

const (
	index         = "templates/admin/index.html"
	editUser      = "templates/admin/edit_user.html"
	gralUsersView = "templates/admin/users.html"
	logRegistry   = "templates/admin/bitacora.html"
	registryUser  = "templates/admin/registry_user.html"
)

type AdminController interface {
	GetAdminIndex(w http.ResponseWriter, r *http.Request)
	GetUpdateUserForm(w http.ResponseWriter, r *http.Request)
	GetGeneralUsersView(w http.ResponseWriter, r *http.Request)
	GetLogRegistyView(w http.ResponseWriter, r *http.Request)
	GetUserRegistry(w http.ResponseWriter, r *http.Request)
}

type AdminControllerImpl struct {
	templateUtils   utils.TemplateUtils
	userDao         dao.UserDao
	catalogsService services.CatalogsService
}

func NewAdminController(templateUtils utils.TemplateUtils, userDao dao.UserDao, catalogsService services.CatalogsService) AdminController {
	return &AdminControllerImpl{templateUtils, userDao, catalogsService}
}

func (a AdminControllerImpl) GetAdminIndex(w http.ResponseWriter, r *http.Request) {

	a.renderTemplate(index, w, r, false, map[string]interface{}{})

}

func (a AdminControllerImpl) GetUpdateUserForm(w http.ResponseWriter, r *http.Request) {

	id, e := strconv.Atoi(mux.Vars(r)["idUser"])
	if e != nil {
		utils.Error.Printf("No se encontrĂ³ el idUser en el path. Detalles: \n\t%v", e)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	user, e := a.userDao.FindUserByID(int32(id))
	if e != nil {
		utils.Error.Printf("Error al buscar el usuario con id '%v'. Detalles: \n\t%v", id, e)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	roles, e := a.catalogsService.CreateCatalog("rol")
	if e != nil {
		utils.Error.Printf("Error al buscar todos los roles. Detalles: \n\t%v", e)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	a.renderTemplate(editUser, w, r, true, map[string]interface{}{
		"user":  user,
		"roles": roles,
	})

}

func (a AdminControllerImpl) GetGeneralUsersView(w http.ResponseWriter, r *http.Request) {

	users, e := a.userDao.FindAll()

	if e != nil {
		utils.Error.Printf("Error al buscar todos los usuarios. Detalles: \n\t%v", e)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	roles, e := a.catalogsService.CreateCatalog("rol")
	if e != nil {
		utils.Error.Printf("Error al buscar todos los roles. Detalles: \n\t%v", e)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	a.renderTemplate(gralUsersView, w, r, true, map[string]interface{}{
		"users": users,
		"roles": roles,
	})

}

func (a AdminControllerImpl) GetLogRegistyView(w http.ResponseWriter, r *http.Request) {

	a.renderTemplate(logRegistry, w, r, true, map[string]interface{}{})

}

func (a AdminControllerImpl) GetUserRegistry(w http.ResponseWriter, r *http.Request) {

	roles, e := a.catalogsService.CreateCatalog("rol")
	if e != nil {
		utils.Error.Printf("Error al buscar todos los roles. Detalles: \n\t%v", e)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	a.renderTemplate(registryUser, w, r, true, map[string]interface{}{
		"roles": roles,
	})

}

// renderTemplate realiza el render del template. Se puede indicar si lleva CSRF o no
func (a AdminControllerImpl) renderTemplate(pagePath string, w http.ResponseWriter, r *http.Request, hasCsrf bool, data map[string]interface{}) {
	if !hasCsrf {
		e := a.templateUtils.RenderTemplateToResponseWriter(pagePath, w, data)
		if e != nil {
			utils.Error.Printf("Error al ingresar a la pagina '%s'. Detalles: \n\t%v", pagePath, e)
		}
		return
	}

	data[csrf.TemplateTag] = csrf.TemplateField(r)
	e := a.templateUtils.RenderTemplateToResponseWriter(pagePath, w, data)
	if e != nil {
		utils.Error.Printf("Error al ingresar a la pagina '%s'. Detalles: \n\t%v", pagePath, e)
	}
}
