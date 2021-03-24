package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/manicar2093/guianetThea/app/utils"
	"github.com/manicar2093/guianetThea/app/web"
)

var (
	userDao                      web.UserDao
	detailsHostingDao            web.DetailsHostingDao
	detailsEndpointAndHostingDao web.DetailsEndpointAndHostingDao
	endpointDao                  web.EndpointDao
	uuidGeneratorUtils           utils.UUIDGeneratorUtils
	passwordUtils                utils.PasswordUtils
	loginService                 web.LoginService
	recordService                web.RecordService
	pageController               *web.PageController
	loginController              *web.LoginController
	middlewareProvider           web.MiddlewareProvider
)

func RegistryHandlers(r *mux.Router) {
	r.HandleFunc("/", pageController.GetOnDevTemplate).Methods(http.MethodGet)
	r.HandleFunc("/index", pageController.GetLoginPage).Methods(http.MethodGet)
	r.HandleFunc("/{page}", middlewareProvider.NeedsLoggedIn(pageController.GetRequestedPage)).Methods(http.MethodGet)

	r.HandleFunc("/login", loginController.Login).Methods(http.MethodPost)
}

func init() {
	userDao = web.NewUserDao(web.DB)
	detailsHostingDao = web.NewDetailsHostingDao(web.DB)
	detailsEndpointAndHostingDao = web.NewDetailsEndpointAndHostingDao(web.DB)
	endpointDao = web.NewEndpointDao(web.DB)

	uuidGeneratorUtils = utils.NewUUIDGeneratorUtils()
	passwordUtils = utils.NewPasswordUtils()

	loginService = web.NewLoginService(userDao, passwordUtils, web.Session, detailsHostingDao, uuidGeneratorUtils)
	recordService = web.NewRecordService(detailsEndpointAndHostingDao, detailsHostingDao, endpointDao, web.Session)

	pageController = web.NewPageController(web.Session, recordService)
	loginController = web.NewLoginController(loginService, web.Session)
	middlewareProvider = web.NewMiddlewareProvider(web.Session)
}
