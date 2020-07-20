package pages

import (
	"github.com/assetto-corsa-web/accweb/auth"
	"net/http"
)

func Server(w http.ResponseWriter, r *http.Request, claims *auth.TokenClaims) {
	data := struct {
		IsAdmin bool
	}{
		true, // TODO
	}

	if err := executeTemplate(w, "server_page.html", data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
