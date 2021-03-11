package web

import "net/http"

// MiddlewareProvider contiene los middlewares del sistema
type MiddlewareProvider struct {
	session SessionHandler
}

func NewMiddlewareProvider(session SessionHandler) *MiddlewareProvider {
	return &MiddlewareProvider{session: session}
}

// NeedsLoggedIn valida que exista alguna sesión activa. Si está no existe se redirige al login.
func (m MiddlewareProvider) NeedsLoggedIn(h http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		isLoggin := m.session.IsLoggedIn(w, r)
		if !isLoggin {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		h.ServeHTTP(w, r)

	})

}
