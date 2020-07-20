package pages

import (
	"github.com/assetto-corsa-web/accweb/auth"
	"net/http"
)

func Logs(w http.ResponseWriter, r *http.Request, claims *auth.TokenClaims) {
	data := struct {
		Servername string
	}{
		"TODO servername", // TODO
	}

	if err := executeTemplate(w, "logs_page.html", data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
