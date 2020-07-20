package pages

import (
	"github.com/assetto-corsa-web/accweb/auth"
	"net/http"
)

func Overview(w http.ResponseWriter, r *http.Request, claims *auth.TokenClaims) {
	data := struct {
		IsAdmin bool
	}{
		true, // TODO
	}

	if err := executeTemplate(w, "overview_page.html", data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
