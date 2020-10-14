package pages

import (
	"github.com/assetto-corsa-web/accweb/auth"
	"net/http"
)

func DeleteUser(w http.ResponseWriter, r *http.Request, claims *auth.TokenClaims) {
	if !claims.IsAdmin {
		http.Redirect(w, r, "/not-found", http.StatusNotFound)
		return
	}

	data := UserPageData{}
	username := r.URL.Query().Get("username")
	data.User = auth.GetUserList().Get(username)

	if data.User != nil && r.Method == http.MethodPost {
		auth.GetUserList().Remove(username)
		http.Redirect(w, r, "/user", http.StatusFound)
		return
	}

	if err := executeTemplate(w, "delete_user_page.html", data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
