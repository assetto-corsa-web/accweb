package pages

import "net/http"

func NotFound(w http.ResponseWriter, r *http.Request) {
	if err := executeTemplate(w, "404_page.html", nil); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
