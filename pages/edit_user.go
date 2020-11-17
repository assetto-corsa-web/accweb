package pages

import (
	"github.com/assetto-corsa-web/accweb/auth"
	"net/http"
)

func EditUser(w http.ResponseWriter, r *http.Request, claims *auth.TokenClaims) {
	if !claims.IsAdmin {
		http.Redirect(w, r, "/not-found", http.StatusNotFound)
		return
	}

	data := UserPageData{}
	username := r.URL.Query().Get("username")
	data.User = auth.GetUserList().Get(username)

	if data.User != nil && r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		password := r.Form.Get("password")
		repeatPassword := r.Form.Get("repeat_password")
		data.User.Role = r.Form.Get("role")

		if password != "" && repeatPassword != "" {
			data.User.Password = password
			data.PasswordError = isValidPassword(password, repeatPassword)
		}

		data.RoleError = isValidRole(data.User.Role)

		if data.PasswordError == "" && data.RoleError == "" {
			auth.GetUserList().Set(data.User.Username, data.User.Password, data.User.Role)
			http.Redirect(w, r, "/user", http.StatusFound)
			return
		}
	}

	if err := executeTemplate(w, "edit_user.html", data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
