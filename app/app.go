package app

import (
	"net/http"

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
	uuidGeneratorUtils           utils.UUIDGeneratorUtils
	passwordUtils                utils.PasswordUtils
	loginService                 services.LoginService
	recordService                services.RecordService
	pageController               *controllers.PageController
	loginController              *controllers.LoginController
	middlewareProvider           middleware.MiddlewareProvider
)

func RegistryHandlers(r *mux.Router) {
	webHandlers(r)
}

func webHandlers(r *mux.Router) {
	r.HandleFunc("/", pageController.GetOnDevTemplate).Methods(http.MethodGet)
	r.HandleFunc("/index", pageController.GetLoginPage).Methods(http.MethodGet)
	r.HandleFunc("/{page}", middlewareProvider.NeedsLoggedIn(pageController.GetRequestedPage)).Methods(http.MethodGet)

	r.HandleFunc("/login", loginController.Login).Methods(http.MethodPost)
}

func init() {
	userDao = dao.NewUserDao(connections.DB)
	detailsHostingDao = dao.NewDetailsHostingDao(connections.DB)
	detailsEndpointAndHostingDao = dao.NewDetailsEndpointAndHostingDao(connections.DB)
	endpointDao = dao.NewEndpointDao(connections.DB)

	uuidGeneratorUtils = utils.NewUUIDGeneratorUtils()
	passwordUtils = utils.NewPasswordUtils()

	loginService = services.NewLoginService(userDao, passwordUtils, sessions.Session, detailsHostingDao, uuidGeneratorUtils)
	recordService = services.NewRecordService(detailsEndpointAndHostingDao, detailsHostingDao, endpointDao, sessions.Session)

	pageController = controllers.NewPageController(sessions.Session, recordService)
	loginController = controllers.NewLoginController(loginService, sessions.Session)
	middlewareProvider = middleware.NewMiddlewareProvider(sessions.Session)
}
