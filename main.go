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

	fmt.Println("Servidor iniciado")
	http.ListenAndServe(utils.GetPortFromEnvVar("PORT", "8000"), r)
}
