package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/manicar2093/guianetThea/app"
	"github.com/manicar2093/guianetThea/app/utils"
)

func main() {
	r := mux.NewRouter()
	app.RegistryHandlers(r)

	registryStaticHandlers(r)

	fmt.Println("Servidor iniciado")
	http.ListenAndServe(utils.GetPortFromEnvVar("PORT", "8000"), r)
}

func registryStaticHandlers(r *mux.Router) {

	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("./static/css"))))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./static/assets"))))
	r.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("./static/images"))))
	r.PathPrefix("/scripts/").Handler(http.StripPrefix("/scripts/", http.FileServer(http.Dir("./static/scripts"))))

}
