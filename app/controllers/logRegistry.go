package controllers

import "net/http"

type LogRegistryController interface {
	CreateLogginRegistry(w http.ResponseWriter, r *http.Request)
}
