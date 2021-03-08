package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/manicar2093/guianetThea/web"
)

func main() {
	r := mux.NewRouter()
	web.RegistryHandlers(r)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	fmt.Println("Servidor iniciado")
	http.ListenAndServe(":8000", r)
}
