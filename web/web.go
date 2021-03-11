package web

import (
	"net/http"

	"github.com/gorilla/mux"
)

var pageController *PageController
var loginController *LoginController
var middlewareProvider *MiddlewareProvider

func RegistryHandlers(r *mux.Router) {
	r.HandleFunc("/", pageController.GetLoginPage).Methods(http.MethodGet)
	r.HandleFunc("/{page}", middlewareProvider.NeedsLoggedIn(pageController.GetRequestedPage)).Methods(http.MethodGet)

	r.HandleFunc("/login", loginController.Login).Methods(http.MethodPost)
}

func init() {
	pageController = NewPageController(Session)
	loginController = NewLoginController()
	middlewareProvider = NewMiddlewareProvider(Session)
}
