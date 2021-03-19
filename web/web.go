package web

import (
	"net/http"

	"github.com/gorilla/mux"
)

var (
	userDao                      UserDao
	detailsHostingDao            DetailsHostingDao
	detailsEndpointAndHostingDao DetailsEndpointAndHostingDao
	endpointDao                  EndpointDao
	uuidGeneratorUtils           UUIDGeneratorUtils
	passwordUtils                PasswordUtils
	loginService                 LoginService
	recordService                RecordService
	pageController               *PageController
	loginController              *LoginController
	middlewareProvider           MiddlewareProvider
)

func RegistryHandlers(r *mux.Router) {
	r.HandleFunc("/", pageController.GetOnDevTemplate).Methods(http.MethodGet)
	r.HandleFunc("/index", pageController.GetLoginPage).Methods(http.MethodGet)
	r.HandleFunc("/{page}", middlewareProvider.NeedsLoggedIn(pageController.GetRequestedPage)).Methods(http.MethodGet)

	r.HandleFunc("/login", loginController.Login).Methods(http.MethodPost)
}

func init() {
	userDao = NewUserDao(DB)
	detailsHostingDao = NewDetailsHostingDao(DB)
	detailsEndpointAndHostingDao = NewDetailsEndpointAndHostingDao(DB)
	endpointDao = NewEndpointDao(DB)

	uuidGeneratorUtils = NewUUIDGeneratorUtils()
	passwordUtils = NewPasswordUtils()

	loginService = NewLoginService(userDao, passwordUtils, Session, detailsHostingDao, uuidGeneratorUtils)
	recordService = NewRecordService(detailsEndpointAndHostingDao, detailsHostingDao, endpointDao, Session)

	pageController = NewPageController(Session, recordService)
	loginController = NewLoginController(loginService, Session)
	middlewareProvider = NewMiddlewareProvider(Session)
}
