package app

import (
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"github.com/manicar2093/guianetThea/app/connections"
	"github.com/manicar2093/guianetThea/app/controllers"
	"github.com/manicar2093/guianetThea/app/dao"
	"github.com/manicar2093/guianetThea/app/middleware"
	"github.com/manicar2093/guianetThea/app/services"
	"github.com/manicar2093/guianetThea/app/sessions"
	"github.com/manicar2093/guianetThea/app/utils"
)

var (
	userDao                      dao.UserDao
	detailsHostingDao            dao.DetailsHostingDao
	detailsEndpointAndHostingDao dao.DetailsEndpointAndHostingDao
	endpointDao                  dao.EndpointDao
	roleDao                      dao.RolDao
	uuidGeneratorUtils           utils.UUIDGeneratorUtils
	passwordUtils                utils.PasswordUtils
	templateUtils                utils.TemplateUtils
	loginService                 services.LoginService
	recordService                services.RecordService
	validatorService             services.ValidatorService
	catalogsService              services.CatalogsService
	pageController               *controllers.PageController
	loginController              *controllers.LoginController
	adminController              controllers.AdminController
	userController               controllers.UserController
	catalogsController           controllers.CatalogsController
	middlewareProvider           middleware.MiddlewareProvider
	csrfMiddleware               func(http.Handler) http.Handler
)

func RegistryHandlers(r *mux.Router) {
	webHandlers(r)
	adminHandlers(r)
	registryStaticHandlers(r)
}

func webHandlers(r *mux.Router) {
	r.HandleFunc("/", pageController.GetOnDevTemplate).Methods(http.MethodGet)
	r.HandleFunc("/index", pageController.GetLoginPage).Methods(http.MethodGet)
	r.HandleFunc("/{page}", middlewareProvider.NeedsLoggedIn(pageController.GetRequestedPage)).Methods(http.MethodGet)

	r.HandleFunc("/login", loginController.Login).Methods(http.MethodPost)
}

func adminHandlers(r *mux.Router) {
	adminRouter := r.PathPrefix("/admin").Subrouter()
	adminRouter.Use(csrfMiddleware)

	adminRouter.HandleFunc("/", middlewareProvider.NeedsLoggedIn(adminController.GetAdminIndex)).Methods(http.MethodGet)
	adminRouter.HandleFunc("/user/all", middlewareProvider.NeedsLoggedIn(adminController.GetGeneralUsersView)).Methods(http.MethodGet)
	adminRouter.HandleFunc("/user/registry", middlewareProvider.NeedsLoggedIn(adminController.GetUserRegistry)).Methods(http.MethodGet)
	adminRouter.HandleFunc("/logginRegistry", middlewareProvider.NeedsLoggedIn(adminController.GetLogRegistyView)).Methods(http.MethodGet)
	adminRouter.HandleFunc("/user/{idUser}", middlewareProvider.NeedsLoggedIn(adminController.GetUpdateUserForm)).Methods(http.MethodGet)

	adminRouter.HandleFunc("/user/registry", userController.CreateUser).Methods(http.MethodPost)
	adminRouter.HandleFunc("/user/delete/{idUser}", userController.DeleteUser).Methods(http.MethodDelete)
	adminRouter.HandleFunc("/user/restore_password", userController.RestorePassword).Methods(http.MethodPut)
	adminRouter.HandleFunc("/user/update", userController.UpdateUser).Methods(http.MethodPut)

	adminRouter.HandleFunc("/catalogs/{catalog}", catalogsController.GetCatalog).Methods(http.MethodGet)

}

func registryStaticHandlers(r *mux.Router) {

	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("./static/css"))))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./static/assets"))))
	r.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("./static/images"))))
	r.PathPrefix("/scripts/").Handler(http.StripPrefix("/scripts/", http.FileServer(http.Dir("./static/scripts"))))

}

func init() {
	userDao = dao.NewUserDao(connections.DB)
	detailsHostingDao = dao.NewDetailsHostingDao(connections.DB)
	detailsEndpointAndHostingDao = dao.NewDetailsEndpointAndHostingDao(connections.DB)
	endpointDao = dao.NewEndpointDao(connections.DB)
	roleDao = dao.NewRolDao(connections.DB)

	uuidGeneratorUtils = utils.NewUUIDGeneratorUtils()
	passwordUtils = utils.NewPasswordUtils()
	templateUtils = utils.NewTemplateUtils()

	loginService = services.NewLoginService(userDao, passwordUtils, sessions.Session, detailsHostingDao, uuidGeneratorUtils)
	recordService = services.NewRecordService(detailsEndpointAndHostingDao, detailsHostingDao, endpointDao, sessions.Session)
	validatorService = services.NewValidatorService()
	catalogsService = services.NewCatalogService(roleDao)

	pageController = controllers.NewPageController(sessions.Session, recordService, templateUtils)
	loginController = controllers.NewLoginController(loginService, sessions.Session)
	adminController = controllers.NewAdminController(templateUtils, userDao, catalogsService)
	userController = controllers.NewUserController(userDao, validatorService, passwordUtils)
	catalogsController = controllers.NewCatalogController(catalogsService)

	middlewareProvider = middleware.NewMiddlewareProvider(sessions.Session)
	csrfMiddleware = csrf.Protect([]byte("a-key-word"))
}
