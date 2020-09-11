package pages

import (
	"github.com/assetto-corsa-web/accweb/auth"
	"net/http"
)

func User(w http.ResponseWriter, r *http.Request, claims *auth.TokenClaims) {
	data := struct {
		IsAdmin bool
		User    []auth.User
	}{
		claims.IsAdmin,
		auth.GetUserList().GetAll(),
	}

	if err := executeTemplate(w, "user_page.html", data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
