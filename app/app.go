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
	loginRegistryDao             dao.LoginRegistryDao
	uuidGeneratorUtils           utils.UUIDGeneratorUtils
	passwordUtils                utils.PasswordUtils
	templateUtils                utils.TemplateUtils
	loginService                 services.LoginService
	recordService                services.RecordService
	validatorService             services.ValidatorService
	catalogsService              services.CatalogsService
	loginRegistryService         services.LoginRegistryService
	pageController               *controllers.PageController
	loginController              *controllers.LoginController
	adminController              controllers.AdminController
	userController               controllers.UserController
	catalogsController           controllers.CatalogsController
	loginRegistryContoller       controllers.LoginRegistryController
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
	r.HandleFunc("/logout/", loginController.Logout).Methods(http.MethodGet)
}

func adminHandlers(r *mux.Router) {
	adminRouter := r.PathPrefix("/admin").Subrouter()
	adminRouter.Use(csrfMiddleware)

	adminRouter.HandleFunc("/", middleware.MultipleMiddle(
		adminController.GetAdminIndex,
		middlewareProvider.NeedsLoggedIn)).Methods(http.MethodGet)

	adminRouter.HandleFunc("/user/all", middleware.MultipleMiddle(
		adminController.GetGeneralUsersView,
		middlewareProvider.NeedsLoggedIn)).Methods(http.MethodGet)

	adminRouter.HandleFunc("/user/registry", middleware.MultipleMiddle(
		adminController.GetUserRegistry,
		middlewareProvider.NeedsLoggedIn)).Methods(http.MethodGet)

	adminRouter.HandleFunc("/logginRegistry", middleware.MultipleMiddle(
		adminController.GetLogRegistyView,
		middlewareProvider.NeedsLoggedIn)).Methods(http.MethodGet)

	adminRouter.HandleFunc("/user/{idUser}", middleware.MultipleMiddle(
		adminController.GetUpdateUserForm,
		middlewareProvider.NeedsLoggedIn)).Methods(http.MethodGet)
	// User endpoints
	adminRouter.HandleFunc("/user/registry", middleware.MultipleMiddle(
		userController.CreateUser,
		middlewareProvider.NeedsLoggedIn)).Methods(http.MethodPost)

	adminRouter.HandleFunc("/user/delete/{idUser}", middleware.MultipleMiddle(
		userController.DeleteUser,
		middlewareProvider.NeedsLoggedIn)).Methods(http.MethodDelete)

	adminRouter.HandleFunc("/user/restore_password", middleware.MultipleMiddle(
		userController.RestorePassword,
		middlewareProvider.NeedsLoggedIn)).Methods(http.MethodPut)

	adminRouter.HandleFunc("/user/update", middleware.MultipleMiddle(
		userController.UpdateUser,
		middlewareProvider.NeedsLoggedIn)).Methods(http.MethodPut)

	adminRouter.HandleFunc("/catalogs/{catalog}", middleware.MultipleMiddle(
		catalogsController.GetCatalog,
		middlewareProvider.NeedsLoggedIn)).Methods(http.MethodGet)

	adminRouter.HandleFunc("/login_registry/create", middleware.MultipleMiddle(
		loginRegistryContoller.LoginRegistryInform,
		middlewareProvider.NeedsLoggedIn)).Methods(http.MethodPost)

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
	loginRegistryDao = dao.NewLogRegistryDao(connections.DB)

	uuidGeneratorUtils = utils.NewUUIDGeneratorUtils()
	passwordUtils = utils.NewPasswordUtils()
	templateUtils = utils.NewTemplateUtils()

	loginService = services.NewLoginService(userDao, passwordUtils, sessions.Session, detailsHostingDao, uuidGeneratorUtils)
	recordService = services.NewRecordService(detailsEndpointAndHostingDao, detailsHostingDao, endpointDao, sessions.Session)
	validatorService = services.NewValidatorService()
	catalogsService = services.NewCatalogService(roleDao)
	loginRegistryService = services.NewLoginRegistryService(loginRegistryDao)

	pageController = controllers.NewPageController(sessions.Session, recordService, templateUtils)
	loginController = controllers.NewLoginController(loginService, sessions.Session, recordService)
	adminController = controllers.NewAdminController(templateUtils, userDao, catalogsService)
	userController = controllers.NewUserController(userDao, validatorService, passwordUtils)
	catalogsController = controllers.NewCatalogController(catalogsService)
	loginRegistryContoller = controllers.NewLoginRegistryController(loginRegistryService, validatorService)

	middlewareProvider = middleware.NewMiddlewareProvider(sessions.Session)
	csrfMiddleware = csrf.Protect([]byte("a-key-word"))
}
