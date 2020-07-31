package pages

import (
	"github.com/assetto-corsa-web/accweb/auth"
	"net/http"
	"time"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:     auth.TokenCookie,
		Expires:  time.Now().Add(-time.Minute),
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/login", http.StatusFound)
}
