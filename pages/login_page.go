package pages

import (
	"net/http"
)

func RenderLoginPage(w http.ResponseWriter, r *http.Request) {
	if err := executeTemplate(w, "login_page.html", nil); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
