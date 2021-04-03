package middleware

import (
	"net/http"

	"github.com/manicar2093/guianetThea/app/dao"
	"github.com/manicar2093/guianetThea/app/sessions"
	"github.com/manicar2093/guianetThea/app/utils"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

type MiddlewareProvider interface {
	NeedsLoggedIn(h http.HandlerFunc) http.HandlerFunc
	IsAdmin(h http.HandlerFunc) http.HandlerFunc
}

// MiddlewareProviderImpl contiene los middlewares del sistema
type MiddlewareProviderImpl struct {
	session           sessions.SessionHandler
	rolDao            dao.RolDao
	detailsHostingDao dao.DetailsHostingDao
}

func NewMiddlewareProvider(session sessions.SessionHandler, rolDao dao.RolDao, detailsHostingDao dao.DetailsHostingDao) MiddlewareProvider {
	return &MiddlewareProviderImpl{session: session, rolDao: rolDao, detailsHostingDao: detailsHostingDao}
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

func (m MiddlewareProviderImpl) IsAdmin(h http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		uuid, e := m.session.GetSessionUUID(w, r)
		if e != nil {
			utils.Error.Printf("Error al obtener el userID de la sesión. Detalles: \n\t%v", e)
			http.Error(w, "Internal Server error", http.StatusInternalServerError)
			return
		}

		details, e := m.detailsHostingDao.FindDetailsHostingByUUID(uuid)
		if e != nil {
			utils.Error.Printf("Error al obtener el detalle de la sesión. Detalles: \n\t%v", e)
			http.Error(w, "Internal Server error", http.StatusInternalServerError)
			return
		}

		hasRol, e := m.rolDao.UserHasRol(int(details.UserID), "ADMIN")
		if e != nil {
			utils.Error.Printf("Error al convertir el userID a int. Detalles: \n\t%v", e)
			http.Error(w, "Internal Server error", http.StatusInternalServerError)
			return
		}
		if !hasRol {
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
