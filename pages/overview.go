package pages

import (
	"github.com/assetto-corsa-web/accweb/auth"
	"net/http"
)

func Overview(w http.ResponseWriter, r *http.Request, claims *auth.TokenClaims) {
	data := struct {
		IsAdmin bool
	}{
		claims.IsAdmin,
	}

	if err := executeTemplate(w, "overview.html", data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
