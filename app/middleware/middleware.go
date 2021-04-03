package middleware

import (
	"net/http"

	"github.com/manicar2093/guianetThea/app/sessions"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

type MiddlewareProvider interface {
	NeedsLoggedIn(h http.HandlerFunc) http.HandlerFunc
}

// MiddlewareProviderImpl contiene los middlewares del sistema
type MiddlewareProviderImpl struct {
	session sessions.SessionHandler
}

func NewMiddlewareProvider(session sessions.SessionHandler) MiddlewareProvider {
	return &MiddlewareProviderImpl{session: session}
}

// NeedsLoggedIn valida que exista alguna sesión activa. Si está no existe se redirige al login.
func (m MiddlewareProviderImpl) NeedsLoggedIn(h http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		isLoggin := m.session.IsLoggedIn(w, r)

		if !isLoggin {
			m.session.AddFlashMessage(sessions.FlashMessage{Type: "info", Value: "Favor de iniciar sesión."}, w, r)
			http.Redirect(w, r, "/index", http.StatusSeeOther)
			return
		}

		h.ServeHTTP(w, r)

	})

}

func MultipleMiddle(h http.HandlerFunc, mid ...Middleware) http.HandlerFunc {
	if len(mid) < 1 {
		return h
	}

	wrapped := h

	for i := len(mid) - 1; i >= 0; i-- {
		wrapped = mid[i](wrapped)
	}

	return wrapped
}
