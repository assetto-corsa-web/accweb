package pages

import (
	"github.com/assetto-corsa-web/accweb/auth"
	"net/http"
)

func NewUser(w http.ResponseWriter, r *http.Request, claims *auth.TokenClaims) {
	if err := executeTemplate(w, "new_user_page.html", nil); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
