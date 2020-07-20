package pages

import "net/http"

func Status(w http.ResponseWriter, r *http.Request) {
	if err := executeTemplate(w, "status_page.html", nil); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
