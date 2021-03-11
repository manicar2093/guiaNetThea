package web

import (
	"net/http"

	"github.com/gorilla/mux"
)

var userDao UserDao
var detailsHostingDao DetailsHostingDao
var passwordUtils PasswordUtils

var loginService LoginService

var pageController *PageController
var loginController *LoginController
var middlewareProvider *MiddlewareProvider

func RegistryHandlers(r *mux.Router) {
	r.HandleFunc("/", pageController.GetLoginPage).Methods(http.MethodGet)
	r.HandleFunc("/{page}", middlewareProvider.NeedsLoggedIn(pageController.GetRequestedPage)).Methods(http.MethodGet)

	r.HandleFunc("/login", loginController.Login).Methods(http.MethodPost)
}

func init() {
	userDao = NewUserDao(DB)
	detailsHostingDao = NewDetailsHostingDao(DB)

	passwordUtils = NewPasswordUtils()

	loginService = NewLoginService(userDao, passwordUtils, Session, detailsHostingDao)

	pageController = NewPageController(Session)
	loginController = NewLoginController(loginService)
	middlewareProvider = NewMiddlewareProvider(Session)
}
