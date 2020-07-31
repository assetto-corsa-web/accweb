package pages

import (
	"github.com/assetto-corsa-web/accweb/auth"
	"github.com/emvi/logbuch"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		renderLogin(w, "")
	} else {
		if err := r.ParseForm(); err != nil {
			logbuch.Error("Error parsing login form", logbuch.Fields{"err": err})
			w.WriteHeader(http.StatusBadRequest)
			renderLogin(w, "Bad request.")
			return
		}

		username := r.FormValue("username")
		password := r.FormValue("password")
		user := auth.GetUser(username, password)

		if user == nil {
			w.WriteHeader(http.StatusBadRequest)
			renderLogin(w, "Password incorrect.")
			return
		}

		isAdmin := user.Role == auth.RoleAdmin
		isMod := user.Role == auth.RoleMod || isAdmin
		isRO := user.Role == auth.RoleReadOnly || isMod
		token, expires, err := auth.NewToken(isAdmin, isMod, isRO)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			renderLogin(w, "Error generating token.")
			return
		}

		cookie := &http.Cookie{
			Name:     auth.TokenCookie,
			Value:    token,
			Expires:  expires,
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
		}
		http.SetCookie(w, cookie)
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func renderLogin(w http.ResponseWriter, message string) {
	data := struct {
		Message string
	}{
		message,
	}

	if err := executeTemplate(w, "login_page.html", data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
