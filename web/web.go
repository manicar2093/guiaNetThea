package web

import (
	"net/http"

	"github.com/gorilla/mux"
)

var pageController *PageController
var loginController *LoginController

func RegistryHandlers(r *mux.Router) {
	r.HandleFunc("/login", pageController.GetLoginPage).Methods(http.MethodGet)
	r.HandleFunc("/{page}", pageController.GetRequestedPage).Methods(http.MethodGet)

	r.HandleFunc("/login", loginController.Login).Methods(http.MethodPost)
}

func init() {
	pageController = NewPageController()
	loginController = NewLoginController()
}
