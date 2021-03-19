package web

import "net/http"

type MiddlewareProvider interface {
	NeedsLoggedIn(h http.HandlerFunc) http.HandlerFunc
}

// MiddlewareProviderImpl contiene los middlewares del sistema
type MiddlewareProviderImpl struct {
	session SessionHandler
}

func NewMiddlewareProvider(session SessionHandler) MiddlewareProvider {
	return &MiddlewareProviderImpl{session: session}
}

// NeedsLoggedIn valida que exista alguna sesión activa. Si está no existe se redirige al login.
func (m MiddlewareProviderImpl) NeedsLoggedIn(h http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		isLoggin := m.session.IsLoggedIn(w, r)
		Info.Println("Usuario esta logeado?", isLoggin)
		if !isLoggin {
			m.session.AddFlashMessage(FlashMessage{Type: "info", Value: "Favor de iniciar sesión."}, w, r)
			http.Redirect(w, r, "/index", http.StatusSeeOther)
			return
		}

		h.ServeHTTP(w, r)

	})

}
